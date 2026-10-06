package revkeen

import (
	"fmt"
	"net/http"
	"strings"
)

// NewClient creates a secret-key client for a named RevKeen environment. The
// key is sent as a raw x-api-key value; it is never converted to a bearer token.
func NewClient(apiKey string, environment Environment, opts ...ClientOption) (*APIClient, error) {
	return NewSecretClient(apiKey, environment, opts...)
}

// NewSecretClient creates a client authenticated with an rk_live_ or
// rk_sandbox_ secret key for a named RevKeen environment.
func NewSecretClient(apiKey string, environment Environment, opts ...ClientOption) (*APIClient, error) {
	baseURL, err := environment.baseURL()
	if err != nil {
		return nil, err
	}
	return newSecretClient(apiKey, baseURL, opts...)
}

// NewClientWithCustomBaseURL creates a secret-key client for an explicitly
// supplied endpoint. Prefer NewClient with a named environment.
func NewClientWithCustomBaseURL(apiKey, rawURL string, opts ...ClientOption) (*APIClient, error) {
	return NewSecretClientWithCustomBaseURL(apiKey, rawURL, opts...)
}

// NewSecretClientWithCustomBaseURL creates a secret-key client for an
// explicitly supplied endpoint. Prefer NewSecretClient with a named environment.
func NewSecretClientWithCustomBaseURL(apiKey, rawURL string, opts ...ClientOption) (*APIClient, error) {
	baseURL, err := validateCustomBaseURL(rawURL)
	if err != nil {
		return nil, err
	}
	return newSecretClient(apiKey, baseURL, opts...)
}

func newSecretClient(apiKey, baseURL string, opts ...ClientOption) (*APIClient, error) {
	if !hasAnyPrefix(apiKey, "rk_live_", "rk_sandbox_") {
		return nil, fmt.Errorf("revkeen: secret API key must start with rk_live_ or rk_sandbox_")
	}
	return newRuntimeClient(apiKey, baseURL, opts...)
}

// PublishableClient exposes only browser-safe generated services.
type PublishableClient struct {
	CartSessionsAPI CartSessionsAPI
	StorefrontAPI   StorefrontAPI
}

// NewPublishableClient creates a restricted browser-safe client for a named
// RevKeen environment, authenticated with a publishable key.
func NewPublishableClient(apiKey string, environment Environment, opts ...ClientOption) (*PublishableClient, error) {
	baseURL, err := environment.baseURL()
	if err != nil {
		return nil, err
	}
	return newPublishableClient(apiKey, baseURL, opts...)
}

// NewPublishableClientWithCustomBaseURL creates a restricted browser-safe
// client for an explicitly supplied endpoint. Prefer NewPublishableClient with
// a named environment.
func NewPublishableClientWithCustomBaseURL(apiKey, rawURL string, opts ...ClientOption) (*PublishableClient, error) {
	baseURL, err := validateCustomBaseURL(rawURL)
	if err != nil {
		return nil, err
	}
	return newPublishableClient(apiKey, baseURL, opts...)
}

func newPublishableClient(apiKey, baseURL string, opts ...ClientOption) (*PublishableClient, error) {
	if !hasAnyPrefix(apiKey, "rk_pk_live_", "rk_pk_sandbox_") {
		return nil, fmt.Errorf("revkeen: publishable API key must start with rk_pk_live_ or rk_pk_sandbox_")
	}
	client, err := newRuntimeClient(apiKey, baseURL, opts...)
	if err != nil {
		return nil, err
	}
	return &PublishableClient{
		CartSessionsAPI: client.CartSessionsAPI,
		StorefrontAPI:   client.StorefrontAPI,
	}, nil
}

func newRuntimeClient(apiKey, baseURL string, opts ...ClientOption) (*APIClient, error) {
	settings := defaultClientSettings()
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(&settings); err != nil {
			return nil, err
		}
	}

	cfg := NewConfiguration()
	if len(cfg.Servers) == 0 {
		return nil, fmt.Errorf("revkeen: generated configuration has no server entry")
	}
	cfg.Servers[0].URL = strings.TrimRight(baseURL, "/")

	httpClient := settings.httpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	clientCopy := *httpClient
	base := clientCopy.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	standardTransport, ok := base.(*http.Transport)
	if !ok {
		return nil, fmt.Errorf("revkeen: custom HTTP transport must be *http.Transport so connect and response timeouts remain enforceable")
	}
	if standardTransport.DialTLS != nil || standardTransport.DialTLSContext != nil {
		return nil, fmt.Errorf("revkeen: custom TLS dialers are incompatible with enforceable connect timeouts")
	}
	clientCopy.Transport = &runtimeRoundTripper{
		apiKey:         apiKey,
		base:           base,
		connectTimeout: settings.connectTimeout,
		readTimeout:    settings.readTimeout,
		maxAttempts:    settings.maxAttempts,
		jitter:         fullJitter,
		sleep:          sleepWithContext,
	}
	cfg.HTTPClient = &clientCopy
	return NewAPIClient(cfg), nil
}

func hasAnyPrefix(value string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(value, prefix) && len(value) > len(prefix) {
			return true
		}
	}
	return false
}
