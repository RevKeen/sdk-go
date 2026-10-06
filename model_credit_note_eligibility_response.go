/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your merchant API key in the `x-api-key` header, or as `Authorization: Bearer <key>` — both are accepted. Create and manage keys in the [Dashboard](https://app.revkeen.com/dashboard/settings/general?tab=developer). Use `rk_sandbox_*` for staging/test and `rk_live_*` for production. Keys are scoped, rollable, IP-lockable and revocable; the raw secret is shown once, at creation.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Access tokens are opaque, prefixed `rkoa_`, and sent via `Authorization: Bearer rkoa_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Errors  Every non-2xx response uses one envelope: `{ \"error\": { \"type\", \"code\", \"message\", ... } }`. Branch on `error.type` and `error.code`; treat `error.message` as display text only. Error bodies never echo the credential you presented.  | Status | `error.type` | Notes | |--------|--------------|-------| | `400` | `invalid_request_error` | Schema validation failed. `details.fields` maps each rejected parameter to its reasons. | | `401` | `authentication_error` | Missing, invalid, expired or revoked credential. Always carries `WWW-Authenticate: Bearer realm=\"revkeen-api\"`. | | `402` | `card_error` | Declined by the issuer or processor; see `decline_code`. | | `403` | `authorization_error` | The caller is identified but the action is denied — scope, IP allowlist, origin, or merchant mismatch. | | `409` | `invalid_request_error` or `idempotency_error` | Resource-state conflict, or an `Idempotency-Key` replayed with a different body. | | `429` | `rate_limit_error` | See Rate Limits below. | | `500` | `api_error` | Retry with exponential backoff, quoting `error.request_id`. |  `401` codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`.  `403` codes: `insufficient_permissions`, `forbidden`, `ip_not_allowed`, `tenant_scope_mismatch`, `no_organization`, `no_merchant`, `merchant_mismatch`, `cart_disabled`, `origin_not_allowed`, `verification_required`, `read_only_impersonation_blocked`, `publishable_key_restricted`, `internal_only`.  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  Each API key (`rk_live_*`, `rk_sandbox_*`, `rk_pk_*`) may make **600 requests per minute**, in production and staging alike. The limit is per key, counted over a 60-second window, with no separate burst allowance. Some endpoints apply their own, lower limits.  An over-limit key is still valid: it receives `429`, never `401`. Wait for `Retry-After` and retry; do not re-authenticate or rotate the key.  A rate-limited request returns `429` with `type: rate_limit_error`, `code: rate_limit_exceeded` and `details.retry_after` in whole seconds.  | Header | Meaning | |--------|---------| | `Retry-After` | Whole seconds to wait before retrying. Present on every `429`. | | `X-RateLimit-Limit` | Requests permitted in the current window, when the limiter publishes one. | | `X-RateLimit-Remaining` | Requests left in the current window. | | `X-RateLimit-Reset` | UNIX timestamp in **seconds** at which the window resets. |

API version: 2026-05-01
Contact: support@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the CreditNoteEligibilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNoteEligibilityResponse{}

// CreditNoteEligibilityResponse Eligibility details showing whether a credit note can be issued for the specified invoice and the maximum creditable amount.
type CreditNoteEligibilityResponse struct {
	InvoiceId            string                                       `json:"invoice_id"`
	TotalAmountMinor     int64                                        `json:"total_amount_minor"`
	TotalPaidMinor       int64                                        `json:"total_paid_minor"`
	TotalCreditedMinor   int64                                        `json:"total_credited_minor"`
	MaxCreditableMinor   int64                                        `json:"max_creditable_minor"`
	Eligible             bool                                         `json:"eligible"`
	Payments             []CreditNoteEligibilityResponsePaymentsInner `json:"payments"`
	AdditionalProperties map[string]interface{}
}

type _CreditNoteEligibilityResponse CreditNoteEligibilityResponse

// NewCreditNoteEligibilityResponse instantiates a new CreditNoteEligibilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNoteEligibilityResponse(invoiceId string, totalAmountMinor int64, totalPaidMinor int64, totalCreditedMinor int64, maxCreditableMinor int64, eligible bool, payments []CreditNoteEligibilityResponsePaymentsInner) *CreditNoteEligibilityResponse {
	this := CreditNoteEligibilityResponse{}
	this.InvoiceId = invoiceId
	this.TotalAmountMinor = totalAmountMinor
	this.TotalPaidMinor = totalPaidMinor
	this.TotalCreditedMinor = totalCreditedMinor
	this.MaxCreditableMinor = maxCreditableMinor
	this.Eligible = eligible
	this.Payments = payments
	return &this
}

// NewCreditNoteEligibilityResponseWithDefaults instantiates a new CreditNoteEligibilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNoteEligibilityResponseWithDefaults() *CreditNoteEligibilityResponse {
	this := CreditNoteEligibilityResponse{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *CreditNoteEligibilityResponse) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *CreditNoteEligibilityResponse) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetTotalAmountMinor returns the TotalAmountMinor field value
func (o *CreditNoteEligibilityResponse) GetTotalAmountMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalAmountMinor
}

// GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetTotalAmountMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountMinor, true
}

// SetTotalAmountMinor sets field value
func (o *CreditNoteEligibilityResponse) SetTotalAmountMinor(v int64) {
	o.TotalAmountMinor = v
}

// GetTotalPaidMinor returns the TotalPaidMinor field value
func (o *CreditNoteEligibilityResponse) GetTotalPaidMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalPaidMinor
}

// GetTotalPaidMinorOk returns a tuple with the TotalPaidMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetTotalPaidMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPaidMinor, true
}

// SetTotalPaidMinor sets field value
func (o *CreditNoteEligibilityResponse) SetTotalPaidMinor(v int64) {
	o.TotalPaidMinor = v
}

// GetTotalCreditedMinor returns the TotalCreditedMinor field value
func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCreditedMinor
}

// GetTotalCreditedMinorOk returns a tuple with the TotalCreditedMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCreditedMinor, true
}

// SetTotalCreditedMinor sets field value
func (o *CreditNoteEligibilityResponse) SetTotalCreditedMinor(v int64) {
	o.TotalCreditedMinor = v
}

// GetMaxCreditableMinor returns the MaxCreditableMinor field value
func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxCreditableMinor
}

// GetMaxCreditableMinorOk returns a tuple with the MaxCreditableMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCreditableMinor, true
}

// SetMaxCreditableMinor sets field value
func (o *CreditNoteEligibilityResponse) SetMaxCreditableMinor(v int64) {
	o.MaxCreditableMinor = v
}

// GetEligible returns the Eligible field value
func (o *CreditNoteEligibilityResponse) GetEligible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetEligibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Eligible, true
}

// SetEligible sets field value
func (o *CreditNoteEligibilityResponse) SetEligible(v bool) {
	o.Eligible = v
}

// GetPayments returns the Payments field value
func (o *CreditNoteEligibilityResponse) GetPayments() []CreditNoteEligibilityResponsePaymentsInner {
	if o == nil {
		var ret []CreditNoteEligibilityResponsePaymentsInner
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponse) GetPaymentsOk() ([]CreditNoteEligibilityResponsePaymentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *CreditNoteEligibilityResponse) SetPayments(v []CreditNoteEligibilityResponsePaymentsInner) {
	o.Payments = v
}

func (o CreditNoteEligibilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNoteEligibilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["total_amount_minor"] = o.TotalAmountMinor
	toSerialize["total_paid_minor"] = o.TotalPaidMinor
	toSerialize["total_credited_minor"] = o.TotalCreditedMinor
	toSerialize["max_creditable_minor"] = o.MaxCreditableMinor
	toSerialize["eligible"] = o.Eligible
	toSerialize["payments"] = o.Payments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNoteEligibilityResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_id",
		"total_amount_minor",
		"total_paid_minor",
		"total_credited_minor",
		"max_creditable_minor",
		"eligible",
		"payments",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreditNoteEligibilityResponse := _CreditNoteEligibilityResponse{}

	err = json.Unmarshal(data, &varCreditNoteEligibilityResponse)

	if err != nil {
		return err
	}

	*o = CreditNoteEligibilityResponse(varCreditNoteEligibilityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "total_amount_minor")
		delete(additionalProperties, "total_paid_minor")
		delete(additionalProperties, "total_credited_minor")
		delete(additionalProperties, "max_creditable_minor")
		delete(additionalProperties, "eligible")
		delete(additionalProperties, "payments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNoteEligibilityResponse struct {
	value *CreditNoteEligibilityResponse
	isSet bool
}

func (v NullableCreditNoteEligibilityResponse) Get() *CreditNoteEligibilityResponse {
	return v.value
}

func (v *NullableCreditNoteEligibilityResponse) Set(val *CreditNoteEligibilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNoteEligibilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNoteEligibilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNoteEligibilityResponse(val *CreditNoteEligibilityResponse) *NullableCreditNoteEligibilityResponse {
	return &NullableCreditNoteEligibilityResponse{value: val, isSet: true}
}

func (v NullableCreditNoteEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNoteEligibilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
