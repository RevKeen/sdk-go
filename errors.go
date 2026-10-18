package revkeen

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ErrorKind is a stable, transport-independent classification for API errors.
// Callers should use Kind for control flow and retain Code for finer-grained
// diagnostics supplied by the API.
type ErrorKind string

const (
	ErrorKindAuthentication      ErrorKind = "authentication"
	ErrorKindAuthorization       ErrorKind = "authorization"
	ErrorKindValidation          ErrorKind = "validation"
	ErrorKindNotFound            ErrorKind = "not_found"
	ErrorKindRateLimit           ErrorKind = "rate_limit"
	ErrorKindIdempotencyConflict ErrorKind = "idempotency_conflict"
	ErrorKindConflict            ErrorKind = "conflict"
	ErrorKindServer              ErrorKind = "server"
	ErrorKindUnknown             ErrorKind = "unknown"
)

// RateLimit exposes the rate-limit headers RevKeen publishes alongside a
// response. Every field is nil when its header was absent, so a published zero
// (for example "0 remaining") is never confused with an unreported limit.
//
// Units are fixed by the public contract: ResetAt is parsed from
// X-RateLimit-Reset, which is a UNIX timestamp in seconds, and RetryAfter is
// the server-stated Retry-After delay. RetryAfter is never clamped to the SDK
// retry ceiling; that bound belongs to the retry policy alone.
type RateLimit struct {
	// Limit is X-RateLimit-Limit: requests permitted in the current window.
	Limit *int64
	// Remaining is X-RateLimit-Remaining: requests left in the current window.
	Remaining *int64
	// ResetAt is X-RateLimit-Reset resolved from UNIX seconds.
	ResetAt *time.Time
	// RetryAfter is Retry-After: whole seconds, or an HTTP-date resolved
	// against now.
	RetryAfter *time.Duration
}

func parseRateLimit(header http.Header, now time.Time) RateLimit {
	if header == nil {
		return RateLimit{}
	}
	limit := RateLimit{
		Limit:     parseNonNegativeInt64(header.Get("X-RateLimit-Limit")),
		Remaining: parseNonNegativeInt64(header.Get("X-RateLimit-Remaining")),
	}
	if seconds := parseNonNegativeInt64(header.Get("X-RateLimit-Reset")); seconds != nil {
		resetAt := time.Unix(*seconds, 0).UTC()
		limit.ResetAt = &resetAt
	}
	if delay, ok := unboundedRetryAfter(header.Get("Retry-After"), now); ok {
		limit.RetryAfter = &delay
	}
	return limit
}

func parseNonNegativeInt64(value string) *int64 {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil || parsed < 0 {
		return nil
	}
	return &parsed
}

// unboundedRetryAfter reports the delay exactly as the server stated it. The
// retry policy applies its own ceiling; error metadata must not.
func unboundedRetryAfter(value string, now time.Time) (time.Duration, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	if seconds := parseNonNegativeInt64(value); seconds != nil {
		return time.Duration(*seconds) * time.Second, true
	}
	when, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}
	delay := when.Sub(now)
	if delay < 0 {
		delay = 0
	}
	return delay, true
}

// RequestError preserves the generated error while surfacing the identifiers
// needed to correlate every failed SDK call with RevKeen logs.
type RequestError struct {
	Err           error
	Kind          ErrorKind
	RequestID     string
	TraceID       string
	CorrelationID string
	Attempts      int
	StatusCode    int
	RawBody       []byte
	Code          string
	Message       string
	Details       any
	// RateLimit carries the rate-limit headers observed on the response, if
	// any. Its fields stay nil when the corresponding header was absent.
	RateLimit RateLimit
}

func (e *RequestError) Error() string {
	if e == nil || e.Err == nil {
		return "revkeen request failed"
	}
	return fmt.Sprintf("revkeen request failed (request_id=%s trace_id=%s attempts=%d): %v", e.RequestID, e.TraceID, e.Attempts, e.Err)
}

func (e *RequestError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// Body returns a defensive copy of the raw API error body.
func (e *RequestError) Body() []byte {
	if e == nil {
		return nil
	}
	return append([]byte(nil), e.RawBody...)
}

func wrapRequestError(ctx context.Context, response *http.Response, err error) error {
	state := runtimeStateFromContext(ctx)
	if err == nil {
		return nil
	}
	var existing *RequestError
	if errors.As(err, &existing) {
		return err
	}

	requestID, traceID, attempts := "", "", 1
	if state != nil {
		requestID, traceID, attempts = state.requestID, state.traceID, state.attempts
	}
	statusCode := 0
	rateLimit := RateLimit{}
	if response != nil {
		statusCode = response.StatusCode
		rateLimit = parseRateLimit(response.Header, time.Now())
		if value := response.Header.Get("x-request-id"); value != "" {
			requestID = value
		} else if value := response.Header.Get("x-correlation-id"); value != "" {
			// Legacy response-only fallback. The SDK never sends this header.
			requestID = value
		}
		if value := response.Header.Get("x-trace-id"); value != "" {
			traceID = value
		}
	}
	rawBody := errorBody(err)
	payload := parseErrorPayload(rawBody)
	if payload.code == "" {
		payload.code = "unknown_error"
	}
	return &RequestError{
		Err:           err,
		Kind:          classifyError(statusCode, payload.code),
		RequestID:     requestID,
		TraceID:       traceID,
		CorrelationID: requestID,
		Attempts:      attempts,
		StatusCode:    statusCode,
		RawBody:       rawBody,
		Code:          payload.code,
		Message:       payload.message,
		Details:       payload.details,
		RateLimit:     rateLimit,
	}
}

type rawBodyError interface {
	Body() []byte
}

func errorBody(err error) []byte {
	var bodyError rawBodyError
	if !errors.As(err, &bodyError) {
		return nil
	}
	return append([]byte(nil), bodyError.Body()...)
}

type errorPayload struct {
	code    string
	message string
	details any
}

func parseErrorPayload(body []byte) errorPayload {
	if len(bytes.TrimSpace(body)) == 0 {
		return errorPayload{}
	}
	decoder := json.NewDecoder(bytes.NewReader(body))
	decoder.UseNumber()
	var root map[string]any
	if err := decoder.Decode(&root); err != nil {
		return errorPayload{}
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return errorPayload{}
	}

	candidate := root
	nestedEnvelope := false
	if nested, ok := root["error"].(map[string]any); ok {
		candidate = nested
		nestedEnvelope = true
	}
	payload := errorPayload{
		code:    firstString(candidate, "code", "error_code", "errorCode", "machine_code", "machineCode", "type"),
		message: firstString(candidate, "message", "detail", "title"),
		details: firstValue(candidate, "details", "errors"),
	}
	if payload.code == "" && nestedEnvelope {
		payload.code = firstString(root, "code", "error_code", "errorCode", "machine_code", "machineCode", "type")
	}
	if payload.message == "" {
		if message, ok := root["error"].(string); ok {
			payload.message = message
		} else if nestedEnvelope {
			payload.message = firstString(root, "message", "detail", "title")
		}
	}
	if payload.details == nil && nestedEnvelope {
		payload.details = firstValue(root, "details", "errors")
	}
	return payload
}

func firstString(values map[string]any, keys ...string) string {
	for _, key := range keys {
		if value, ok := values[key].(string); ok && value != "" {
			return value
		}
	}
	return ""
}

func firstValue(values map[string]any, keys ...string) any {
	for _, key := range keys {
		if value, ok := values[key]; ok {
			return value
		}
	}
	return nil
}

func classifyError(statusCode int, code string) ErrorKind {
	switch statusCode {
	case http.StatusUnauthorized:
		return ErrorKindAuthentication
	case http.StatusForbidden:
		return ErrorKindAuthorization
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		return ErrorKindValidation
	case http.StatusNotFound:
		return ErrorKindNotFound
	case http.StatusTooManyRequests:
		return ErrorKindRateLimit
	case http.StatusConflict:
		switch strings.ToLower(code) {
		case "idempotency_conflict", "idempotency_key_conflict", "idempotency_key_reused", "idempotency_key_mismatch", "idempotency_mismatch":
			return ErrorKindIdempotencyConflict
		}
		return ErrorKindConflict
	default:
		if statusCode >= http.StatusInternalServerError {
			return ErrorKindServer
		}
		return ErrorKindUnknown
	}
}
