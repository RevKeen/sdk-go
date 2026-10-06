package revkeen

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"
)

type generatedBodyError struct{ body []byte }

func (err *generatedBodyError) Error() string { return "generated API error" }
func (err *generatedBodyError) Body() []byte  { return err.body }

func TestRequestErrorKindTaxonomy(t *testing.T) {
	tests := []struct {
		name   string
		status int
		body   string
		kind   ErrorKind
	}{
		{name: "authentication", status: 401, kind: ErrorKindAuthentication},
		{name: "authorization", status: 403, kind: ErrorKindAuthorization},
		{name: "bad request", status: 400, kind: ErrorKindValidation},
		{name: "unprocessable", status: 422, kind: ErrorKindValidation},
		{name: "not found", status: 404, kind: ErrorKindNotFound},
		{name: "rate limit", status: 429, kind: ErrorKindRateLimit},
		{name: "idempotency mismatch", status: 409, body: `{"code":"IDEMPOTENCY_KEY_MISMATCH"}`, kind: ErrorKindIdempotencyConflict},
		{name: "idempotency in progress", status: 409, body: `{"error":{"code":"idempotency_conflict"}}`, kind: ErrorKindIdempotencyConflict},
		{name: "generic conflict", status: 409, body: `{"code":"STATE_CONFLICT"}`, kind: ErrorKindConflict},
		{name: "server", status: 503, kind: ErrorKindServer},
		{name: "fallback", status: 418, kind: ErrorKindUnknown},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cause := &generatedBodyError{body: []byte(test.body)}
			response := &http.Response{StatusCode: test.status, Header: make(http.Header)}
			var requestError *RequestError
			if !errors.As(wrapRequestError(context.Background(), response, cause), &requestError) {
				t.Fatal("expected RequestError")
			}
			if requestError.Kind != test.kind {
				t.Fatalf("kind = %q; want %q", requestError.Kind, test.kind)
			}
		})
	}
}

func TestRequestErrorParsesNestedAndFlatBodiesWithoutLosingCause(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		wantCode    string
		wantMessage string
		wantDetails any
	}{
		{name: "nested", body: `{"error":{"code":"INVALID_INPUT","message":"amount is invalid","details":{"field":"amount"}}}`, wantCode: "INVALID_INPUT", wantMessage: "amount is invalid", wantDetails: map[string]any{"field": "amount"}},
		{name: "flat aliases", body: `{"error_code":"RATE_LIMITED","detail":"slow down","errors":[{"field":"request"}]}`, wantCode: "RATE_LIMITED", wantMessage: "slow down", wantDetails: []any{map[string]any{"field": "request"}}},
		{name: "string error", body: `{"code":"FAILED","error":"plain message"}`, wantCode: "FAILED", wantMessage: "plain message"},
		{name: "malformed", body: `{not-json`, wantCode: "unknown_error"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cause := &generatedBodyError{body: []byte(test.body)}
			response := &http.Response{StatusCode: 400, Header: http.Header{
				"X-Request-Id": []string{"server-request"},
				"X-Trace-Id":   []string{"server-trace"},
			}}
			got := wrapRequestError(context.Background(), response, cause)
			var requestError *RequestError
			if !errors.As(got, &requestError) {
				t.Fatal("expected RequestError")
			}
			if !errors.Is(got, cause) {
				t.Fatal("generated error compatibility was lost")
			}
			if requestError.Code != test.wantCode || requestError.Message != test.wantMessage {
				t.Fatalf("parsed code/message = %q/%q; want %q/%q", requestError.Code, requestError.Message, test.wantCode, test.wantMessage)
			}
			if !reflect.DeepEqual(requestError.Details, test.wantDetails) {
				t.Fatalf("details = %#v; want %#v", requestError.Details, test.wantDetails)
			}
			if string(requestError.RawBody) != test.body || string(requestError.Body()) != test.body {
				t.Fatal("raw error body was not preserved")
			}
			if requestError.RequestID != "server-request" || requestError.CorrelationID != "server-request" || requestError.TraceID != "server-trace" {
				t.Fatalf("response identifiers not preserved: %+v", requestError)
			}
		})
	}
}

func TestRequestErrorBodyReturnsDefensiveCopy(t *testing.T) {
	requestError := &RequestError{RawBody: []byte("body")}
	copyOfBody := requestError.Body()
	copyOfBody[0] = 'B'
	if string(requestError.RawBody) != "body" {
		t.Fatal("Body exposed mutable RequestError storage")
	}
}

func TestParseRateLimitUsesContractedUnits(t *testing.T) {
	now := time.Date(2026, time.August, 28, 12, 0, 0, 0, time.UTC)
	header := http.Header{}
	header.Set("X-RateLimit-Limit", "1000")
	header.Set("X-RateLimit-Remaining", "0")
	// X-RateLimit-Reset is a UNIX timestamp in seconds, never milliseconds.
	header.Set("X-RateLimit-Reset", "1787000000")
	header.Set("Retry-After", "30")

	limit := parseRateLimit(header, now)

	if limit.Limit == nil || *limit.Limit != 1000 {
		t.Fatalf("Limit = %v; want 1000", limit.Limit)
	}
	if limit.Remaining == nil || *limit.Remaining != 0 {
		t.Fatalf("Remaining = %v; want a published 0", limit.Remaining)
	}
	wantReset := time.Unix(1787000000, 0).UTC()
	if limit.ResetAt == nil || !limit.ResetAt.Equal(wantReset) {
		t.Fatalf("ResetAt = %v; want %v", limit.ResetAt, wantReset)
	}
	if limit.RetryAfter == nil || *limit.RetryAfter != 30*time.Second {
		t.Fatalf("RetryAfter = %v; want 30s", limit.RetryAfter)
	}
}

func TestParseRateLimitResolvesHttpDateAndNeverClampsRetryAfter(t *testing.T) {
	now := time.Date(2026, time.August, 28, 12, 0, 0, 0, time.UTC)

	header := http.Header{}
	header.Set("Retry-After", now.Add(45*time.Second).Format(http.TimeFormat))
	if limit := parseRateLimit(header, now); limit.RetryAfter == nil || *limit.RetryAfter != 45*time.Second {
		t.Fatalf("HTTP-date RetryAfter = %v; want 45s", limit.RetryAfter)
	}

	// The retry policy bounds its own sleep at retryMaxDelay. Error metadata
	// must report what the server actually said.
	header.Set("Retry-After", "600")
	if limit := parseRateLimit(header, now); limit.RetryAfter == nil || *limit.RetryAfter != 600*time.Second {
		t.Fatalf("RetryAfter = %v; want an unclamped 600s", limit.RetryAfter)
	}
}

func TestParseRateLimitLeavesAbsentAndUnparseableHeadersNil(t *testing.T) {
	now := time.Date(2026, time.August, 28, 12, 0, 0, 0, time.UTC)

	if limit := parseRateLimit(http.Header{}, now); limit.Limit != nil || limit.Remaining != nil || limit.ResetAt != nil || limit.RetryAfter != nil {
		t.Fatalf("absent headers produced %+v; want an empty RateLimit", limit)
	}

	header := http.Header{}
	header.Set("X-RateLimit-Limit", "  ")
	header.Set("X-RateLimit-Remaining", "many")
	header.Set("X-RateLimit-Reset", "-1")
	header.Set("Retry-After", "soon")
	if limit := parseRateLimit(header, now); limit.Limit != nil || limit.Remaining != nil || limit.ResetAt != nil || limit.RetryAfter != nil {
		t.Fatalf("unparseable headers produced %+v; want an empty RateLimit", limit)
	}
}

func TestRequestErrorSurfacesRateLimitHeaders(t *testing.T) {
	header := http.Header{}
	header.Set("X-RateLimit-Limit", "1000")
	header.Set("X-RateLimit-Remaining", "0")
	header.Set("Retry-After", "30")
	response := &http.Response{StatusCode: http.StatusTooManyRequests, Header: header}

	var requestError *RequestError
	cause := &generatedBodyError{body: []byte(`{"error":{"code":"rate_limit_exceeded"}}`)}
	if !errors.As(wrapRequestError(context.Background(), response, cause), &requestError) {
		t.Fatal("expected RequestError")
	}

	if requestError.Kind != ErrorKindRateLimit {
		t.Fatalf("kind = %q; want %q", requestError.Kind, ErrorKindRateLimit)
	}
	if requestError.RateLimit.RetryAfter == nil || *requestError.RateLimit.RetryAfter != 30*time.Second {
		t.Fatalf("RetryAfter = %v; want 30s", requestError.RateLimit.RetryAfter)
	}
	if requestError.RateLimit.Remaining == nil || *requestError.RateLimit.Remaining != 0 {
		t.Fatalf("Remaining = %v; want a published 0", requestError.RateLimit.Remaining)
	}

	plain := &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{}}
	var forbidden *RequestError
	if !errors.As(wrapRequestError(context.Background(), plain, &generatedBodyError{}), &forbidden) {
		t.Fatal("expected RequestError")
	}
	if forbidden.Kind != ErrorKindAuthorization {
		t.Fatalf("kind = %q; want %q", forbidden.Kind, ErrorKindAuthorization)
	}
	if forbidden.RateLimit.RetryAfter != nil || forbidden.RateLimit.Limit != nil {
		t.Fatalf("unrated response produced %+v; want an empty RateLimit", forbidden.RateLimit)
	}
}

func TestErrorKindIgnoresProviderNeutralErrorType(t *testing.T) {
	// REV-6760 renamed the 403 envelope type to authorization_error. Kind is
	// derived from the HTTP status, so the rename cannot reclassify it.
	cases := []struct {
		name   string
		status int
		body   string
		kind   ErrorKind
	}{
		{name: "401 authentication_error", status: 401, body: `{"error":{"type":"authentication_error","code":"invalid_api_key"}}`, kind: ErrorKindAuthentication},
		{name: "403 authorization_error", status: 403, body: `{"error":{"type":"authorization_error","code":"insufficient_permissions"}}`, kind: ErrorKindAuthorization},
		{name: "403 stale authentication_error", status: 403, body: `{"error":{"type":"authentication_error","code":"forbidden"}}`, kind: ErrorKindAuthorization},
		{name: "429 rate_limit_error", status: 429, body: `{"error":{"type":"rate_limit_error","code":"rate_limit_exceeded"}}`, kind: ErrorKindRateLimit},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			response := &http.Response{StatusCode: test.status, Header: http.Header{}}
			var requestError *RequestError
			if !errors.As(wrapRequestError(context.Background(), response, &generatedBodyError{body: []byte(test.body)}), &requestError) {
				t.Fatal("expected RequestError")
			}
			if requestError.Kind != test.kind {
				t.Fatalf("kind = %q; want %q", requestError.Kind, test.kind)
			}
		})
	}
}
