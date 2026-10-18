package revkeen

import (
	"context"
	cryptorand "crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"math/big"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	retryInitialDelay = 250 * time.Millisecond
	retryMaxDelay     = 2 * time.Second
)

var idempotencyKeyPattern = regexp.MustCompile(`^(?:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}|[A-Za-z0-9_-]{8,64})$`)

type runtimeRoundTripper struct {
	apiKey         string
	base           http.RoundTripper
	connectTimeout time.Duration
	readTimeout    time.Duration
	maxAttempts    int
	jitter         func(time.Duration) time.Duration
	sleep          func(context.Context, time.Duration) error
	now            func() time.Time
}

func (transport *runtimeRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	state := runtimeStateFromContext(request.Context())
	if state == nil {
		preparedContext, _ := prepareRuntimeRequest(request.Context(), nil, false)
		request = request.Clone(preparedContext)
		state = runtimeStateFromContext(request.Context())
	}

	connectTimeout := transport.connectTimeout
	if state.connectTimeout != nil {
		connectTimeout = *state.connectTimeout
	}
	readTimeout := transport.readTimeout
	if state.readTimeout != nil {
		readTimeout = *state.readTimeout
	}
	maxAttempts := transport.maxAttempts
	if maxAttempts == 0 {
		maxAttempts = defaultMaxAttempts
	}
	if state.maxAttempts != nil {
		maxAttempts = *state.maxAttempts
	}
	base := configuredRoundTripper(transport.base, connectTimeout, readTimeout)
	retryableMethod := request.Method == http.MethodGet ||
		(request.Method == http.MethodPost && state.retrySafePost && idempotencyKeyPattern.MatchString(request.Header.Get("Idempotency-Key")))
	replayableBody := request.Body == nil || request.GetBody != nil

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		if err := request.Context().Err(); err != nil {
			return nil, err
		}
		state.attempts = attempt

		attemptRequest, err := cloneRequestForAttempt(request, request.Context(), attempt)
		if err != nil {
			return nil, err
		}
		attemptRequest.Header.Set("x-api-key", transport.apiKey)
		attemptRequest.Header.Set("x-request-id", state.requestID)
		attemptRequest.Header.Set("x-trace-id", state.traceID)
		attemptRequest.Header.Del("x-correlation-id")

		response, err := base.RoundTrip(attemptRequest)
		if err == nil && response == nil {
			err = errors.New("revkeen: transport returned an empty response")
		}
		if response != nil {
			if response.Body == nil {
				response.Body = http.NoBody
			}
			response.Body = responseBodyWithTimeout(response.Body, readTimeout)
		}
		if err == nil && !shouldRetryStatus(response.StatusCode) {
			return response, nil
		}

		callerCancelled := request.Context().Err() != nil
		canRetry := retryableMethod && replayableBody && !callerCancelled && attempt < maxAttempts
		if err == nil && response != nil {
			if !canRetry {
				return response, nil
			}
			drainAndClose(response.Body)
		} else {
			if !canRetry {
				return nil, err
			}
		}

		delay := transport.jitter(retryDelay(attempt))
		if response != nil {
			now := time.Now
			if transport.now != nil {
				now = transport.now
			}
			if retryAfter, ok := parseRetryAfter(response.Header.Get("Retry-After"), now()); ok {
				delay = retryAfter
			}
		}
		if err := transport.sleep(request.Context(), delay); err != nil {
			return nil, err
		}
	}
	return nil, errors.New("revkeen: retry loop exhausted")
}

func configuredRoundTripper(base http.RoundTripper, connectTimeout, readTimeout time.Duration) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	if standard, ok := base.(*http.Transport); ok {
		clone := standard.Clone()
		clone.DialContext = (&net.Dialer{Timeout: connectTimeout, KeepAlive: 30 * time.Second}).DialContext
		clone.ResponseHeaderTimeout = readTimeout
		return clone
	}
	return base
}

func cloneRequestForAttempt(request *http.Request, ctx context.Context, attempt int) (*http.Request, error) {
	clone := request.Clone(ctx)
	if request.Body == nil {
		return clone, nil
	}
	if request.GetBody != nil {
		body, err := request.GetBody()
		if err != nil {
			return nil, err
		}
		clone.Body = body
		return clone, nil
	}
	if attempt == 1 {
		clone.Body = request.Body
		return clone, nil
	}
	return nil, errors.New("revkeen: request body cannot be replayed")
}

func shouldRetryStatus(status int) bool {
	switch status {
	case http.StatusRequestTimeout, http.StatusTooEarly, http.StatusTooManyRequests,
		http.StatusInternalServerError, http.StatusBadGateway,
		http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	default:
		return false
	}
}

func retryDelay(completedAttempt int) time.Duration {
	delay := retryInitialDelay << (completedAttempt - 1)
	if delay > retryMaxDelay {
		return retryMaxDelay
	}
	return delay
}

func fullJitter(limit time.Duration) time.Duration {
	if limit <= 0 {
		return 0
	}
	value, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(limit)+1))
	if err != nil {
		return limit / 2
	}
	return time.Duration(value.Int64())
}

func sleepWithContext(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func drainAndClose(body io.ReadCloser) {
	if body == nil {
		return
	}
	_, _ = io.Copy(io.Discard, io.LimitReader(body, 64<<10))
	_ = body.Close()
}

type cancelOnCloseReadCloser struct {
	io.ReadCloser
	timer *time.Timer
}

func (body *cancelOnCloseReadCloser) Close() error {
	if body.timer != nil {
		body.timer.Stop()
	}
	return body.ReadCloser.Close()
}

func responseBodyWithTimeout(body io.ReadCloser, timeout time.Duration) io.ReadCloser {
	wrapper := &cancelOnCloseReadCloser{ReadCloser: body}
	if timeout > 0 {
		wrapper.timer = time.AfterFunc(timeout, func() {
			_ = body.Close()
		})
	}
	return wrapper
}

func parseRetryAfter(value string, now time.Time) (time.Duration, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.Atoi(value); err == nil && seconds >= 0 {
		return boundedRetryAfter(time.Duration(seconds) * time.Second), true
	}
	when, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}
	return boundedRetryAfter(when.Sub(now)), true
}

func boundedRetryAfter(delay time.Duration) time.Duration {
	if delay < 0 {
		return 0
	}
	if delay > retryMaxDelay {
		return retryMaxDelay
	}
	return delay
}

func newRuntimeID() string {
	var raw [16]byte
	if _, err := io.ReadFull(cryptorand.Reader, raw[:]); err != nil {
		return ""
	}
	raw[6] = (raw[6] & 0x0f) | 0x40
	raw[8] = (raw[8] & 0x3f) | 0x80
	encoded := make([]byte, 36)
	hex.Encode(encoded[0:8], raw[0:4])
	encoded[8] = '-'
	hex.Encode(encoded[9:13], raw[4:6])
	encoded[13] = '-'
	hex.Encode(encoded[14:18], raw[6:8])
	encoded[18] = '-'
	hex.Encode(encoded[19:23], raw[8:10])
	encoded[23] = '-'
	hex.Encode(encoded[24:36], raw[10:16])
	return string(encoded)
}
