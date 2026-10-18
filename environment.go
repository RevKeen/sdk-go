package revkeen

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	sandboxBaseURL    = "https://staging-api.revkeen.com/v2"
	productionBaseURL = "https://api.revkeen.com/v2"
)

// Environment identifies one of RevKeen's named API environments.
type Environment string

const (
	EnvironmentSandbox    Environment = "sandbox"
	EnvironmentProduction Environment = "production"
)

func (environment Environment) baseURL() (string, error) {
	switch environment {
	case EnvironmentSandbox:
		return sandboxBaseURL, nil
	case EnvironmentProduction:
		return productionBaseURL, nil
	default:
		return "", fmt.Errorf("revkeen: unknown environment %q", environment)
	}
}

func validateCustomBaseURL(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("revkeen: base URL must be an absolute URL")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("revkeen: base URL must use http or https")
	}
	if parsed.User != nil {
		return "", fmt.Errorf("revkeen: base URL must not contain credentials")
	}
	if parsed.RawQuery != "" || parsed.Fragment != "" {
		return "", fmt.Errorf("revkeen: base URL must not contain a query or fragment")
	}
	return strings.TrimRight(rawURL, "/"), nil
}
