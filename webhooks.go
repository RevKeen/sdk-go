package revkeen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DefaultWebhookTolerance = 5 * time.Minute

// WebhookSignatureVerificationErrorCode is a stable reason callers can use to
// distinguish malformed, stale, mismatched, and invalid webhook deliveries.
type WebhookSignatureVerificationErrorCode string

const (
	WebhookErrorMissingSignature          WebhookSignatureVerificationErrorCode = "missing_signature"
	WebhookErrorInvalidSignatureHeader    WebhookSignatureVerificationErrorCode = "invalid_signature_header"
	WebhookErrorInvalidTimestamp          WebhookSignatureVerificationErrorCode = "invalid_timestamp"
	WebhookErrorTimestampOutsideTolerance WebhookSignatureVerificationErrorCode = "timestamp_outside_tolerance"
	WebhookErrorInvalidSignatureEncoding  WebhookSignatureVerificationErrorCode = "invalid_signature_encoding"
	WebhookErrorSignatureMismatch         WebhookSignatureVerificationErrorCode = "signature_mismatch"
	WebhookErrorInvalidPayload            WebhookSignatureVerificationErrorCode = "invalid_payload"
)

// WebhookSignatureVerificationError reports a stable verification reason
// without exposing the webhook secret or expected signature.
type WebhookSignatureVerificationError struct {
	Code WebhookSignatureVerificationErrorCode
	Err  error
}

func (err *WebhookSignatureVerificationError) Error() string {
	if err == nil {
		return "revkeen: webhook signature verification failed"
	}
	if err.Err != nil {
		return fmt.Sprintf("revkeen: webhook verification failed (%s): %v", err.Code, err.Err)
	}
	return fmt.Sprintf("revkeen: webhook verification failed (%s)", err.Code)
}

func (err *WebhookSignatureVerificationError) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.Err
}

type webhookVerifySettings struct {
	tolerance time.Duration
	now       func() time.Time
}

// WebhookVerifyOption configures webhook verification.
type WebhookVerifyOption func(*webhookVerifySettings) error

// WithWebhookTolerance changes the maximum accepted timestamp skew.
func WithWebhookTolerance(tolerance time.Duration) WebhookVerifyOption {
	return func(settings *webhookVerifySettings) error {
		if tolerance < 0 {
			return webhookVerificationError(WebhookErrorInvalidTimestamp, "timestamp tolerance must not be negative")
		}
		settings.tolerance = tolerance
		return nil
	}
}

func withWebhookVerificationTime(now time.Time) WebhookVerifyOption {
	return func(settings *webhookVerifySettings) error {
		settings.now = func() time.Time { return now }
		return nil
	}
}

type parsedWebhookSignature struct {
	timestamp  int64
	signatures [][]byte
}

// VerifyWebhookSignature verifies a signature over timestamp + "." + the
// untouched request body. Callers must pass bytes captured before JSON parsing.
func VerifyWebhookSignature(payload []byte, signature, secret string, opts ...WebhookVerifyOption) error {
	settings := webhookVerifySettings{tolerance: DefaultWebhookTolerance, now: time.Now}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(&settings); err != nil {
			return err
		}
	}

	parsed, err := parseWebhookSignatureHeader(signature)
	if err != nil {
		return err
	}

	delta := settings.now().Sub(time.Unix(parsed.timestamp, 0))
	if delta < -settings.tolerance || delta > settings.tolerance {
		return webhookVerificationError(WebhookErrorTimestampOutsideTolerance, "timestamp is outside the tolerance window")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = mac.Write([]byte(strconv.FormatInt(parsed.timestamp, 10)))
	_, _ = mac.Write([]byte("."))
	_, _ = mac.Write(payload)
	expected := mac.Sum(nil)
	for _, candidate := range parsed.signatures {
		if hmac.Equal(candidate, expected) {
			return nil
		}
	}
	return webhookVerificationError(WebhookErrorSignatureMismatch, "signature did not match")
}

// ConstructWebhookEvent authenticates the raw body before unmarshalling it
// into target. target follows encoding/json's pointer requirements.
func ConstructWebhookEvent(payload []byte, signature, secret string, target any, opts ...WebhookVerifyOption) error {
	if err := VerifyWebhookSignature(payload, signature, secret, opts...); err != nil {
		return err
	}
	if err := json.Unmarshal(payload, target); err != nil {
		return &WebhookSignatureVerificationError{Code: WebhookErrorInvalidPayload, Err: err}
	}
	return nil
}

func parseWebhookSignatureHeader(header string) (parsedWebhookSignature, error) {
	if header == "" {
		return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorMissingSignature, "signature header is missing")
	}

	var timestampText string
	var encodedSignatures []string
	for _, part := range strings.Split(header, ",") {
		separator := strings.IndexByte(part, '=')
		if separator <= 0 || separator == len(part)-1 {
			return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidSignatureHeader, "signature header is malformed")
		}
		key, value := part[:separator], part[separator+1:]
		switch key {
		case "t":
			if timestampText != "" {
				return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidSignatureHeader, "signature header contains multiple timestamps")
			}
			timestampText = value
		case "v1":
			encodedSignatures = append(encodedSignatures, value)
		}
	}
	if timestampText == "" || len(encodedSignatures) == 0 {
		return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidSignatureHeader, "signature header is missing a timestamp or v1 signature")
	}

	for _, character := range timestampText {
		if character < '0' || character > '9' {
			return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidTimestamp, "timestamp must contain decimal digits only")
		}
	}
	timestamp, err := strconv.ParseInt(timestampText, 10, 64)
	if err != nil || timestamp <= 0 {
		return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidTimestamp, "timestamp is invalid")
	}

	signatures := make([][]byte, 0, len(encodedSignatures))
	for _, encoded := range encodedSignatures {
		if len(encoded) != sha256.Size*2 {
			return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidSignatureEncoding, "v1 signature must contain 64 hexadecimal characters")
		}
		decoded, err := hex.DecodeString(encoded)
		if err != nil {
			return parsedWebhookSignature{}, webhookVerificationError(WebhookErrorInvalidSignatureEncoding, "v1 signature must contain 64 hexadecimal characters")
		}
		signatures = append(signatures, decoded)
	}
	return parsedWebhookSignature{timestamp: timestamp, signatures: signatures}, nil
}

func webhookVerificationError(code WebhookSignatureVerificationErrorCode, message string) error {
	return &WebhookSignatureVerificationError{Code: code, Err: fmt.Errorf("%s", message)}
}
