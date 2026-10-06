package revkeen

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	defaultConnectTimeout = 10 * time.Second
	defaultReadTimeout    = 30 * time.Second
	defaultMaxAttempts    = 3
	minimumMaxAttempts    = 1
	maximumMaxAttempts    = 3
)

type clientSettings struct {
	httpClient     *http.Client
	connectTimeout time.Duration
	readTimeout    time.Duration
	maxAttempts    int
}

func defaultClientSettings() clientSettings {
	return clientSettings{
		connectTimeout: defaultConnectTimeout,
		readTimeout:    defaultReadTimeout,
		maxAttempts:    defaultMaxAttempts,
	}
}

// ClientOption configures a generated API client and its runtime transport.
type ClientOption func(*clientSettings) error

// WithHTTPClient supplies the underlying HTTP client. It is copied before the
// RevKeen transport is installed, so the caller's client is not mutated.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(settings *clientSettings) error {
		if client == nil {
			return fmt.Errorf("revkeen: HTTP client must not be nil")
		}
		settings.httpClient = client
		return nil
	}
}

// WithConnectTimeout sets the client-level connection timeout.
func WithConnectTimeout(timeout time.Duration) ClientOption {
	return func(settings *clientSettings) error {
		if timeout <= 0 {
			return fmt.Errorf("revkeen: connect timeout must be positive")
		}
		settings.connectTimeout = timeout
		return nil
	}
}

// WithReadTimeout sets the client-level response timeout.
func WithReadTimeout(timeout time.Duration) ClientOption {
	return func(settings *clientSettings) error {
		if timeout <= 0 {
			return fmt.Errorf("revkeen: read timeout must be positive")
		}
		settings.readTimeout = timeout
		return nil
	}
}

// WithMaxAttempts bounds total attempts, including the initial request.
func WithMaxAttempts(attempts int) ClientOption {
	return func(settings *clientSettings) error {
		if attempts < minimumMaxAttempts || attempts > maximumMaxAttempts {
			return fmt.Errorf("revkeen: max attempts must be between %d and %d", minimumMaxAttempts, maximumMaxAttempts)
		}
		settings.maxAttempts = attempts
		return nil
	}
}

type requestSettings struct {
	requestID      string
	traceID        string
	connectTimeout *time.Duration
	readTimeout    *time.Duration
	maxAttempts    *int
}

// RequestOption applies to one generated request via WithRequestOptions.
type RequestOption func(*requestSettings) error

// WithRequestID uses a caller-provided request identifier for one call.
func WithRequestID(requestID string) RequestOption {
	return func(settings *requestSettings) error {
		if err := validateCorrelationID("request ID", requestID); err != nil {
			return err
		}
		settings.requestID = requestID
		return nil
	}
}

// WithTraceID uses a caller-provided trace identifier for one call.
func WithTraceID(traceID string) RequestOption {
	return func(settings *requestSettings) error {
		if err := validateCorrelationID("trace ID", traceID); err != nil {
			return err
		}
		settings.traceID = traceID
		return nil
	}
}

// WithCallConnectTimeout overrides the connection timeout for one call.
func WithCallConnectTimeout(timeout time.Duration) RequestOption {
	return func(settings *requestSettings) error {
		if timeout <= 0 {
			return fmt.Errorf("revkeen: per-call connect timeout must be positive")
		}
		settings.connectTimeout = &timeout
		return nil
	}
}

// WithCallReadTimeout overrides the response timeout for one call.
func WithCallReadTimeout(timeout time.Duration) RequestOption {
	return func(settings *requestSettings) error {
		if timeout <= 0 {
			return fmt.Errorf("revkeen: per-call read timeout must be positive")
		}
		settings.readTimeout = &timeout
		return nil
	}
}

// WithCallMaxAttempts overrides total attempts for one call.
func WithCallMaxAttempts(attempts int) RequestOption {
	return func(settings *requestSettings) error {
		if attempts < minimumMaxAttempts || attempts > maximumMaxAttempts {
			return fmt.Errorf("revkeen: per-call max attempts must be between %d and %d", minimumMaxAttempts, maximumMaxAttempts)
		}
		settings.maxAttempts = &attempts
		return nil
	}
}

type runtimeContextKey struct{}

type runtimeState struct {
	requestID      string
	traceID        string
	retrySafePost  bool
	connectTimeout *time.Duration
	readTimeout    *time.Duration
	maxAttempts    *int
	attempts       int
}

func prepareRuntimeRequest(ctx context.Context, opts []RequestOption, retrySafePost bool) (context.Context, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if existing, ok := ctx.Value(runtimeContextKey{}).(*runtimeState); ok {
		existing.retrySafePost = existing.retrySafePost || retrySafePost
		return ctx, nil
	}

	settings := requestSettings{}
	var optionErr error
	for _, opt := range opts {
		if opt != nil {
			if err := opt(&settings); err != nil {
				optionErr = err
				break
			}
		}
	}
	if settings.requestID == "" {
		settings.requestID = newRuntimeID()
		if settings.requestID == "" {
			return ctx, fmt.Errorf("revkeen: could not generate a cryptographically secure request ID")
		}
	}
	if settings.traceID == "" {
		settings.traceID = settings.requestID
	}
	state := &runtimeState{
		requestID:      settings.requestID,
		traceID:        settings.traceID,
		retrySafePost:  retrySafePost,
		connectTimeout: settings.connectTimeout,
		readTimeout:    settings.readTimeout,
		maxAttempts:    settings.maxAttempts,
	}
	return context.WithValue(ctx, runtimeContextKey{}, state), optionErr
}

func runtimeStateFromContext(ctx context.Context) *runtimeState {
	if ctx == nil {
		return nil
	}
	state, _ := ctx.Value(runtimeContextKey{}).(*runtimeState)
	return state
}

func validateCorrelationID(name, value string) error {
	if value == "" {
		return fmt.Errorf("revkeen: %s must not be empty", name)
	}
	if strings.ContainsAny(value, "\r\n") {
		return fmt.Errorf("revkeen: %s must not contain CR or LF", name)
	}
	return nil
}
