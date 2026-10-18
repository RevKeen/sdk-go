package revkeen

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func testRuntimeTransport(base http.RoundTripper) *runtimeRoundTripper {
	return &runtimeRoundTripper{
		apiKey:         "rk_sandbox_test",
		base:           base,
		connectTimeout: defaultConnectTimeout,
		readTimeout:    defaultReadTimeout,
		maxAttempts:    defaultMaxAttempts,
		jitter:         func(time.Duration) time.Duration { return 0 },
		sleep:          func(context.Context, time.Duration) error { return nil },
		now:            time.Now,
	}
}

func runtimeRequest(t *testing.T, method string, retrySafePost bool, idempotencyKey string) *http.Request {
	t.Helper()
	ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{
		WithRequestID("request-fixed"),
		WithTraceID("trace-fixed"),
	}, retrySafePost)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequestWithContext(ctx, method, "https://api.revkeen.test/payment-intents", strings.NewReader(`{"amount":100}`))
	if err != nil {
		t.Fatal(err)
	}
	if idempotencyKey != "" {
		request.Header.Set("Idempotency-Key", idempotencyKey)
	}
	return request
}

func TestUnkeyedPaymentCreationPostTimeoutIsAttemptedOnce(t *testing.T) {
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.DeadlineExceeded
	}))

	_, err := transport.RoundTrip(runtimeRequest(t, http.MethodPost, true, ""))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline error, got %v", err)
	}
	if attempts != 1 {
		t.Fatalf("unkeyed payment-creation POST made %d attempts; want exactly 1", attempts)
	}
}

func TestPaymentCreationTimeoutIsObservedOnceWithoutKeyAndRetriedWithStableKeyAndBody(t *testing.T) {
	for _, testCase := range []struct {
		name           string
		idempotencyKey string
		wantAttempts   int
	}{
		{name: "unkeyed", wantAttempts: 1},
		{name: "keyed", idempotencyKey: "payment_create_123", wantAttempts: defaultMaxAttempts},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			type observation struct {
				body           string
				idempotencyKey string
				requestID      string
				traceID        string
			}
			var mu sync.Mutex
			var observations []observation
			release := make(chan struct{})
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				body, _ := io.ReadAll(request.Body)
				mu.Lock()
				observations = append(observations, observation{
					body:           string(body),
					idempotencyKey: request.Header.Get("Idempotency-Key"),
					requestID:      request.Header.Get("x-request-id"),
					traceID:        request.Header.Get("x-trace-id"),
				})
				mu.Unlock()
				// The API has received the payment creation, but deliberately never
				// sends headers or a body before the SDK's response timeout.
				<-release
			}))
			defer server.Close()
			defer close(release)

			ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{
				WithRequestID("request-fixed"),
				WithTraceID("trace-fixed"),
				WithCallReadTimeout(20 * time.Millisecond),
			}, true)
			if err != nil {
				t.Fatal(err)
			}
			request, err := http.NewRequestWithContext(ctx, http.MethodPost, server.URL+"/payment-intents", strings.NewReader(`{"amount":100}`))
			if err != nil {
				t.Fatal(err)
			}
			if testCase.idempotencyKey != "" {
				request.Header.Set("Idempotency-Key", testCase.idempotencyKey)
			}
			transport := testRuntimeTransport(http.DefaultTransport)
			_, err = transport.RoundTrip(request)
			if err == nil {
				t.Fatal("expected response-header timeout")
			}

			mu.Lock()
			got := append([]observation(nil), observations...)
			mu.Unlock()
			if len(got) != testCase.wantAttempts {
				t.Fatalf("server observed %d payment creations; want %d", len(got), testCase.wantAttempts)
			}
			for _, observed := range got {
				if observed.body != `{"amount":100}` {
					t.Fatalf("retried body changed: %q", observed.body)
				}
				if observed.idempotencyKey != testCase.idempotencyKey {
					t.Fatalf("idempotency key changed: %q", observed.idempotencyKey)
				}
				if observed.requestID != "request-fixed" || observed.traceID != "trace-fixed" {
					t.Fatalf("correlation IDs changed: request=%q trace=%q", observed.requestID, observed.traceID)
				}
			}
		})
	}
}

func TestKeyedOperationProvenPostTimeoutUsesThreeTotalAttempts(t *testing.T) {
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.DeadlineExceeded
	}))

	_, err := transport.RoundTrip(runtimeRequest(t, http.MethodPost, true, "payment_create_123"))
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline error, got %v", err)
	}
	if attempts != defaultMaxAttempts {
		t.Fatalf("keyed, operation-proven POST made %d attempts; want %d", attempts, defaultMaxAttempts)
	}
}

func TestKeyedButUnprovenPostTimeoutIsAttemptedOnce(t *testing.T) {
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.DeadlineExceeded
	}))

	_, _ = transport.RoundTrip(runtimeRequest(t, http.MethodPost, false, "payment_create_123"))
	if attempts != 1 {
		t.Fatalf("unproven POST made %d attempts; want exactly 1", attempts)
	}
}

func TestGetTimeoutUsesThreeTotalAttempts(t *testing.T) {
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.DeadlineExceeded
	}))

	_, _ = transport.RoundTrip(runtimeRequest(t, http.MethodGet, false, ""))
	if attempts != defaultMaxAttempts {
		t.Fatalf("GET made %d attempts; want %d", attempts, defaultMaxAttempts)
	}
}

func TestRetryStatusTable(t *testing.T) {
	retryStatuses := []int{408, 425, 429, 500, 502, 503, 504}
	for _, status := range retryStatuses {
		t.Run(http.StatusText(status), func(t *testing.T) {
			attempts := 0
			transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
				attempts++
				return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader("retry")), Header: make(http.Header)}, nil
			}))
			response, err := transport.RoundTrip(runtimeRequest(t, http.MethodGet, false, ""))
			if err != nil {
				t.Fatal(err)
			}
			_ = response.Body.Close()
			if attempts != defaultMaxAttempts {
				t.Fatalf("status %d made %d attempts; want %d", status, attempts, defaultMaxAttempts)
			}
		})
	}
}

func TestRetryAfterIsHonoredAndBounded(t *testing.T) {
	var sleeps []time.Duration
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Body:       io.NopCloser(strings.NewReader("retry")),
			Header:     http.Header{"Retry-After": []string{"1"}},
		}, nil
	}))
	transport.maxAttempts = 2
	transport.sleep = func(_ context.Context, delay time.Duration) error {
		sleeps = append(sleeps, delay)
		return nil
	}
	response, err := transport.RoundTrip(runtimeRequest(t, http.MethodGet, false, ""))
	if err != nil {
		t.Fatal(err)
	}
	_ = response.Body.Close()
	if len(sleeps) != 1 || sleeps[0] != time.Second {
		t.Fatalf("Retry-After sleeps = %v; want [1s]", sleeps)
	}

	now := time.Date(2026, time.August, 21, 12, 0, 0, 0, time.UTC)
	if delay, ok := parseRetryAfter(now.Add(10*time.Second).Format(http.TimeFormat), now); !ok || delay != retryMaxDelay {
		t.Fatalf("bounded HTTP-date Retry-After = %s, %v", delay, ok)
	}
	if delay, ok := parseRetryAfter("10", now); !ok || delay != retryMaxDelay {
		t.Fatalf("bounded numeric Retry-After = %s, %v", delay, ok)
	}
}

func TestConflictAndMutationMethodsAreNeverRetried(t *testing.T) {
	for _, method := range []string{http.MethodPut, http.MethodPatch, http.MethodDelete} {
		t.Run(method, func(t *testing.T) {
			attempts := 0
			transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
				attempts++
				return nil, context.DeadlineExceeded
			}))
			_, _ = transport.RoundTrip(runtimeRequest(t, method, true, "payment_create_123"))
			if attempts != 1 {
				t.Fatalf("%s made %d attempts; want exactly 1", method, attempts)
			}
		})
	}

	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return &http.Response{StatusCode: http.StatusConflict, Body: io.NopCloser(strings.NewReader("conflict")), Header: make(http.Header)}, nil
	}))
	response, err := transport.RoundTrip(runtimeRequest(t, http.MethodGet, false, ""))
	if err != nil {
		t.Fatal(err)
	}
	_ = response.Body.Close()
	if attempts != 1 {
		t.Fatalf("409 made %d attempts; want exactly 1", attempts)
	}
}

func TestCallerCancellationIsNeverRetried(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	ctx, err := prepareRuntimeRequest(ctx, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.revkeen.test/payment-intents", nil)
	if err != nil {
		t.Fatal(err)
	}
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.Canceled
	}))
	_, err = transport.RoundTrip(request)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
	if attempts != 0 {
		t.Fatalf("cancelled request reached transport %d times; want 0", attempts)
	}
}

func TestFullBodyReadBudgetStartsAfterResponseHeaders(t *testing.T) {
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		writer.WriteHeader(http.StatusOK)
		if flusher, ok := writer.(http.Flusher); ok {
			flusher.Flush()
		}
		<-release
	}))
	defer server.Close()
	defer close(release)

	ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{WithCallReadTimeout(20 * time.Millisecond)}, false)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := testRuntimeTransport(http.DefaultTransport).RoundTrip(request)
	if err != nil {
		t.Fatalf("headers should arrive before timeout: %v", err)
	}
	started := time.Now()
	_, err = io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err == nil {
		t.Fatal("expected full-body read timeout")
	}
	if elapsed := time.Since(started); elapsed > 250*time.Millisecond {
		t.Fatalf("body timeout was not enforced promptly: %s", elapsed)
	}
}

func TestAuthAndCorrelationHeadersAreStableAcrossRetries(t *testing.T) {
	var mu sync.Mutex
	var headers []http.Header
	transport := testRuntimeTransport(roundTripFunc(func(request *http.Request) (*http.Response, error) {
		mu.Lock()
		headers = append(headers, request.Header.Clone())
		mu.Unlock()
		return nil, context.DeadlineExceeded
	}))
	request := runtimeRequest(t, http.MethodGet, false, "")
	request.Header.Set("x-correlation-id", "must-not-leak")
	_, _ = transport.RoundTrip(request)

	if len(headers) != defaultMaxAttempts {
		t.Fatalf("captured %d attempts; want %d", len(headers), defaultMaxAttempts)
	}
	for _, header := range headers {
		if got := header.Get("x-api-key"); got != "rk_sandbox_test" {
			t.Fatalf("x-api-key = %q", got)
		}
		if got := header.Get("x-request-id"); got != "request-fixed" {
			t.Fatalf("x-request-id = %q", got)
		}
		if got := header.Get("x-trace-id"); got != "trace-fixed" {
			t.Fatalf("x-trace-id = %q", got)
		}
		if got := header.Get("x-correlation-id"); got != "" {
			t.Fatalf("SDK sent forbidden x-correlation-id %q", got)
		}
		if got := header.Get("Idempotency-Key"); got != "" {
			t.Fatalf("SDK invented idempotency key %q", got)
		}
	}
}

func TestTraceIDDefaultsToRequestID(t *testing.T) {
	ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{
		WithRequestID("outbound-request"),
	}, false)
	if err != nil {
		t.Fatal(err)
	}
	state := runtimeStateFromContext(ctx)
	if state.traceID != state.requestID {
		t.Fatalf("trace ID %q should default to request ID %q", state.traceID, state.requestID)
	}
}

func TestRequestErrorSurfacesResponseIdentifiers(t *testing.T) {
	ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{
		WithRequestID("outbound-request"), WithTraceID("outbound-trace"),
	}, false)
	if err != nil {
		t.Fatal(err)
	}
	state := runtimeStateFromContext(ctx)
	state.attempts = 2
	response := &http.Response{StatusCode: 503, Header: http.Header{
		"X-Request-Id": []string{"server-request"},
		"X-Trace-Id":   []string{"server-trace"},
	}}
	want := errors.New("upstream unavailable")
	got := wrapRequestError(ctx, response, want)
	var requestError *RequestError
	if !errors.As(got, &requestError) {
		t.Fatalf("expected RequestError, got %T", got)
	}
	if requestError.RequestID != "server-request" || requestError.CorrelationID != "server-request" || requestError.TraceID != "server-trace" || requestError.Attempts != 2 || requestError.StatusCode != 503 {
		t.Fatalf("unexpected request metadata: %+v", requestError)
	}
	if !errors.Is(got, want) {
		t.Fatal("RequestError did not preserve cause")
	}
}

func TestRequestErrorAcceptsLegacyCorrelationIDOnlyAsResponseFallback(t *testing.T) {
	ctx, err := prepareRuntimeRequest(context.Background(), nil, false)
	if err != nil {
		t.Fatal(err)
	}
	response := &http.Response{StatusCode: 500, Header: http.Header{"X-Correlation-Id": []string{"legacy-id"}}}
	var requestError *RequestError
	if !errors.As(wrapRequestError(ctx, response, errors.New("failed")), &requestError) {
		t.Fatal("expected RequestError")
	}
	if requestError.RequestID != "legacy-id" {
		t.Fatalf("legacy fallback request ID = %q", requestError.RequestID)
	}
}

func TestCredentialPrefixValidation(t *testing.T) {
	for _, value := range []string{"rk_live_value", "rk_sandbox_value"} {
		if !hasAnyPrefix(value, "rk_live_", "rk_sandbox_") {
			t.Fatalf("valid secret prefix rejected: %s", value)
		}
	}
	for _, value := range []string{"rk_pk_live_value", "rk_pk_sandbox_value"} {
		if !hasAnyPrefix(value, "rk_pk_live_", "rk_pk_sandbox_") {
			t.Fatalf("valid publishable prefix rejected: %s", value)
		}
	}
	if hasAnyPrefix("rk_live_", "rk_live_") {
		t.Fatal("empty secret suffix accepted")
	}
	if _, err := NewSecretClient("rk_pk_live_value", EnvironmentProduction); err == nil {
		t.Fatal("publishable key was accepted by secret constructor")
	}
	if _, err := NewPublishableClient("rk_live_value", EnvironmentProduction); err == nil {
		t.Fatal("secret key was accepted by publishable constructor")
	}
}

func TestPublishableClientExposesOnlyBrowserSafeServices(t *testing.T) {
	clientType := reflect.TypeOf(PublishableClient{})
	if clientType.NumField() != 2 {
		t.Fatalf("publishable client exposes %d services; want 2", clientType.NumField())
	}
	want := map[string]bool{"CartSessionsAPI": true, "StorefrontAPI": true}
	for index := 0; index < clientType.NumField(); index++ {
		delete(want, clientType.Field(index).Name)
	}
	if len(want) != 0 {
		t.Fatalf("publishable client is missing browser-safe services: %v", want)
	}
}

func TestOptionsRejectUnsafeValuesAndIncompatibleTransport(t *testing.T) {
	for _, opt := range []RequestOption{
		WithRequestID(""),
		WithRequestID("bad\rvalue"),
		WithTraceID("bad\nvalue"),
		WithCallConnectTimeout(0),
		WithCallReadTimeout(-time.Second),
		WithCallMaxAttempts(0),
		WithCallMaxAttempts(4),
	} {
		if _, err := prepareRuntimeRequest(context.Background(), []RequestOption{opt}, false); err == nil {
			t.Fatal("unsafe request option was accepted")
		}
	}

	settings := defaultClientSettings()
	if err := WithMaxAttempts(0)(&settings); err == nil {
		t.Fatal("client max attempts of zero was accepted")
	}
	if err := WithMaxAttempts(3)(&settings); err != nil || settings.maxAttempts != 3 {
		t.Fatalf("valid client max attempts rejected: %v", err)
	}
	_, err := NewSecretClient("rk_sandbox_test", EnvironmentSandbox, WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("unused")
	})}))
	if err == nil {
		t.Fatal("custom non-*http.Transport was accepted without enforceable connect timeout")
	}
}

func TestPerCallMaxAttemptsCanDisableRetries(t *testing.T) {
	attempts := 0
	transport := testRuntimeTransport(roundTripFunc(func(*http.Request) (*http.Response, error) {
		attempts++
		return nil, context.DeadlineExceeded
	}))
	ctx, err := prepareRuntimeRequest(context.Background(), []RequestOption{WithCallMaxAttempts(1)}, false)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.revkeen.test/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = transport.RoundTrip(request)
	if attempts != 1 {
		t.Fatalf("per-call max attempts made %d attempts; want 1", attempts)
	}
}

func TestIdempotencyKeyPattern(t *testing.T) {
	valid := []string{"payment_create_123", "123e4567-e89b-12d3-a456-426614174000", "ABCDEFGH"}
	invalid := []string{"", "short", "contains spaces", strings.Repeat("a", 65)}
	for _, value := range valid {
		if !idempotencyKeyPattern.MatchString(value) {
			t.Fatalf("valid idempotency key rejected: %q", value)
		}
	}
	for _, value := range invalid {
		if idempotencyKeyPattern.MatchString(value) {
			t.Fatalf("invalid idempotency key accepted: %q", value)
		}
	}
}

func TestRetryContractConstants(t *testing.T) {
	if defaultMaxAttempts != 3 || retryDelay(1) != 250*time.Millisecond || retryDelay(2) != 500*time.Millisecond || retryDelay(5) != retryMaxDelay {
		t.Fatalf("retry contract drifted: attempts=%d first=%s second=%s cap=%s", defaultMaxAttempts, retryDelay(1), retryDelay(2), retryDelay(5))
	}
}
