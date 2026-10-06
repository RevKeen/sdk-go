package revkeen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

type webhookFixture struct {
	Secret              string                  `json:"secret"`
	RawPayload          string                  `json:"rawPayload"`
	ReserializedPayload string                  `json:"reserializedPayload"`
	TamperedPayload     string                  `json:"tamperedPayload"`
	VerificationTime    int64                   `json:"verificationTime"`
	ToleranceSeconds    int64                   `json:"toleranceSeconds"`
	Valid               webhookFixtureSignature `json:"valid"`
	Expired             webhookFixtureSignature `json:"expired"`
	Future              webhookFixtureSignature `json:"future"`
}

type webhookFixtureSignature struct {
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
	Header    string `json:"header"`
}

func loadWebhookFixture(t *testing.T) webhookFixture {
	t.Helper()
	paths := []string{
		"testdata/webhook-signature-fixtures.json",
		"../contracts/webhook-signature-fixtures.json",
	}
	var body []byte
	var err error
	for _, path := range paths {
		body, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	if err != nil {
		t.Fatalf("read shared webhook fixture: %v", err)
	}
	var fixture webhookFixture
	if err := json.Unmarshal(body, &fixture); err != nil {
		t.Fatalf("decode shared webhook fixture: %v", err)
	}
	return fixture
}

func fixtureVerifyOptions(fixture webhookFixture) []WebhookVerifyOption {
	return []WebhookVerifyOption{
		WithWebhookTolerance(time.Duration(fixture.ToleranceSeconds) * time.Second),
		withWebhookVerificationTime(time.Unix(fixture.VerificationTime, 0)),
	}
}

func requireWebhookErrorCode(t *testing.T, err error, want WebhookSignatureVerificationErrorCode) {
	t.Helper()
	var verificationError *WebhookSignatureVerificationError
	if !errors.As(err, &verificationError) {
		t.Fatalf("error type = %T; want *WebhookSignatureVerificationError (%v)", err, err)
	}
	if verificationError.Code != want {
		t.Fatalf("error code = %q; want %q", verificationError.Code, want)
	}
}

func TestVerifyWebhookSignatureUsesSharedContractFixture(t *testing.T) {
	fixture := loadWebhookFixture(t)
	opts := fixtureVerifyOptions(fixture)
	if err := VerifyWebhookSignature([]byte(fixture.RawPayload), fixture.Valid.Header, fixture.Secret, opts...); err != nil {
		t.Fatalf("valid signature rejected: %v", err)
	}

	multiple := fmt.Sprintf("t=%d,v1=%064d,v1=%s", fixture.Valid.Timestamp, 0, fixture.Valid.Signature)
	if err := VerifyWebhookSignature([]byte(fixture.RawPayload), multiple, fixture.Secret, opts...); err != nil {
		t.Fatalf("matching signature among multiple v1 values rejected: %v", err)
	}

	for _, test := range []struct {
		name    string
		payload string
		header  string
		code    WebhookSignatureVerificationErrorCode
	}{
		{name: "tampered", payload: fixture.TamperedPayload, header: fixture.Valid.Header, code: WebhookErrorSignatureMismatch},
		{name: "reserialized", payload: fixture.ReserializedPayload, header: fixture.Valid.Header, code: WebhookErrorSignatureMismatch},
		{name: "expired", payload: fixture.RawPayload, header: fixture.Expired.Header, code: WebhookErrorTimestampOutsideTolerance},
		{name: "future", payload: fixture.RawPayload, header: fixture.Future.Header, code: WebhookErrorTimestampOutsideTolerance},
	} {
		t.Run(test.name, func(t *testing.T) {
			requireWebhookErrorCode(t, VerifyWebhookSignature([]byte(test.payload), test.header, fixture.Secret, opts...), test.code)
		})
	}
}

func TestVerifyWebhookSignatureRejectsMalformedHeadersWithStableCodes(t *testing.T) {
	now := time.Unix(1_700_000_120, 0)
	hexSignature := make([]byte, 64)
	for index := range hexSignature {
		hexSignature[index] = '0'
	}
	for _, test := range []struct {
		name   string
		header string
		code   WebhookSignatureVerificationErrorCode
	}{
		{name: "missing", header: "", code: WebhookErrorMissingSignature},
		{name: "missing separator", header: "t=1700000120,v1", code: WebhookErrorInvalidSignatureHeader},
		{name: "missing signature", header: "t=1700000120", code: WebhookErrorInvalidSignatureHeader},
		{name: "duplicate timestamp", header: "t=1700000120,t=1700000120,v1=" + string(hexSignature), code: WebhookErrorInvalidSignatureHeader},
		{name: "non decimal timestamp", header: "t=1e9,v1=" + string(hexSignature), code: WebhookErrorInvalidTimestamp},
		{name: "bad hex length", header: "t=1700000120,v1=abcd", code: WebhookErrorInvalidSignatureEncoding},
		{name: "bad hex character", header: "t=1700000120,v1=" + string(hexSignature[:63]) + "z", code: WebhookErrorInvalidSignatureEncoding},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := VerifyWebhookSignature(nil, test.header, "rk_wh_test", withWebhookVerificationTime(now))
			requireWebhookErrorCode(t, err, test.code)
		})
	}
}

type webhookUnmarshalProbe struct {
	called *bool
}

func (probe *webhookUnmarshalProbe) UnmarshalJSON([]byte) error {
	*probe.called = true
	return nil
}

func TestConstructWebhookEventVerifiesRawBytesBeforeJSONUnmarshal(t *testing.T) {
	fixture := loadWebhookFixture(t)
	called := false
	target := webhookUnmarshalProbe{called: &called}
	err := ConstructWebhookEvent(
		[]byte(fixture.TamperedPayload),
		fixture.Valid.Header,
		fixture.Secret,
		&target,
		fixtureVerifyOptions(fixture)...,
	)
	requireWebhookErrorCode(t, err, WebhookErrorSignatureMismatch)
	if called {
		t.Fatal("JSON unmarshal ran before signature verification")
	}
}

func TestConstructWebhookEventReturnsTypedInvalidPayloadAfterVerification(t *testing.T) {
	timestamp := int64(1_700_000_120)
	payload := []byte(`{"id":`)
	secret := "rk_wh_test"
	header := signWebhookPayload(timestamp, payload, secret)
	var target map[string]any
	err := ConstructWebhookEvent(payload, header, secret, &target, withWebhookVerificationTime(time.Unix(timestamp, 0)))
	requireWebhookErrorCode(t, err, WebhookErrorInvalidPayload)
}

func TestWebhookToleranceRejectsNegativeDuration(t *testing.T) {
	err := VerifyWebhookSignature(nil, "unused", "unused", WithWebhookTolerance(-time.Second))
	requireWebhookErrorCode(t, err, WebhookErrorInvalidTimestamp)
}

func signWebhookPayload(timestamp int64, payload []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	_, _ = fmt.Fprintf(mac, "%d.", timestamp)
	_, _ = mac.Write(payload)
	return fmt.Sprintf("t=%d,v1=%s", timestamp, hex.EncodeToString(mac.Sum(nil)))
}
