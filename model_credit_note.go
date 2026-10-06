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
	"time"
)

// checks if the CreditNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNote{}

// CreditNote A formal accounting document that reduces the amount owed on a paid or partially paid invoice, used for refunds, billing corrections, and prorated cancellations.
type CreditNote struct {
	// Unique credit note identifier
	Id string `json:"id"`
	// Sequential credit note number for accounting reference
	CreditNoteNumber NullableString `json:"credit_note_number"`
	// ID of the invoice this credit note applies to
	InvoiceId string `json:"invoice_id"`
	// ID of the customer who received the credit
	CustomerId NullableString `json:"customer_id"`
	// Credit amount in minor units (cents)
	AmountMinor int64 `json:"amount_minor"`
	// Tax portion of the credit in minor units
	TaxAmountMinor NullableInt64 `json:"tax_amount_minor"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Credit note status: draft, issued, or void
	Status string `json:"status"`
	// Human-readable reason for the credit
	Reason NullableString `json:"reason"`
	// Machine-readable reason code (e.g., billing_error, customer_request)
	ReasonCode NullableString `json:"reason_code"`
	// How the credit is applied: refund_to_payment_method, customer_balance, or external
	CreditMethod NullableString `json:"credit_method"`
	// URL to the credit note PDF document
	PdfUrl NullableString `json:"pdf_url"`
	// When the credit note was issued (ISO 8601)
	IssuedAt NullableTime `json:"issued_at"`
	// When the credit note was created (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// When the credit note was last updated (ISO 8601)
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CreditNote CreditNote

// NewCreditNote instantiates a new CreditNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNote(id string, creditNoteNumber NullableString, invoiceId string, customerId NullableString, amountMinor int64, taxAmountMinor NullableInt64, currency string, status string, reason NullableString, reasonCode NullableString, creditMethod NullableString, pdfUrl NullableString, issuedAt NullableTime, createdAt time.Time, updatedAt time.Time) *CreditNote {
	this := CreditNote{}
	this.Id = id
	this.CreditNoteNumber = creditNoteNumber
	this.InvoiceId = invoiceId
	this.CustomerId = customerId
	this.AmountMinor = amountMinor
	this.TaxAmountMinor = taxAmountMinor
	this.Currency = currency
	this.Status = status
	this.Reason = reason
	this.ReasonCode = reasonCode
	this.CreditMethod = creditMethod
	this.PdfUrl = pdfUrl
	this.IssuedAt = issuedAt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCreditNoteWithDefaults instantiates a new CreditNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNoteWithDefaults() *CreditNote {
	this := CreditNote{}
	return &this
}

// GetId returns the Id field value
func (o *CreditNote) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreditNote) SetId(v string) {
	o.Id = v
}

// GetCreditNoteNumber returns the CreditNoteNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetCreditNoteNumber() string {
	if o == nil || o.CreditNoteNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreditNoteNumber.Get()
}

// GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetCreditNoteNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditNoteNumber.Get(), o.CreditNoteNumber.IsSet()
}

// SetCreditNoteNumber sets field value
func (o *CreditNote) SetCreditNoteNumber(v string) {
	o.CreditNoteNumber.Set(&v)
}

// GetInvoiceId returns the InvoiceId field value
func (o *CreditNote) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *CreditNote) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *CreditNote) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *CreditNote) GetAmountMinor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetAmountMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *CreditNote) SetAmountMinor(v int64) {
	o.AmountMinor = v
}

// GetTaxAmountMinor returns the TaxAmountMinor field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CreditNote) GetTaxAmountMinor() int64 {
	if o == nil || o.TaxAmountMinor.Get() == nil {
		var ret int64
		return ret
	}

	return *o.TaxAmountMinor.Get()
}

// GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetTaxAmountMinorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxAmountMinor.Get(), o.TaxAmountMinor.IsSet()
}

// SetTaxAmountMinor sets field value
func (o *CreditNote) SetTaxAmountMinor(v int64) {
	o.TaxAmountMinor.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *CreditNote) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreditNote) SetCurrency(v string) {
	o.Currency = v
}

// GetStatus returns the Status field value
func (o *CreditNote) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreditNote) SetStatus(v string) {
	o.Status = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *CreditNote) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetReasonCode returns the ReasonCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetReasonCode() string {
	if o == nil || o.ReasonCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonCode.Get()
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonCode.Get(), o.ReasonCode.IsSet()
}

// SetReasonCode sets field value
func (o *CreditNote) SetReasonCode(v string) {
	o.ReasonCode.Set(&v)
}

// GetCreditMethod returns the CreditMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetCreditMethod() string {
	if o == nil || o.CreditMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreditMethod.Get()
}

// GetCreditMethodOk returns a tuple with the CreditMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetCreditMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreditMethod.Get(), o.CreditMethod.IsSet()
}

// SetCreditMethod sets field value
func (o *CreditNote) SetCreditMethod(v string) {
	o.CreditMethod.Set(&v)
}

// GetPdfUrl returns the PdfUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNote) GetPdfUrl() string {
	if o == nil || o.PdfUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PdfUrl.Get()
}

// GetPdfUrlOk returns a tuple with the PdfUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetPdfUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PdfUrl.Get(), o.PdfUrl.IsSet()
}

// SetPdfUrl sets field value
func (o *CreditNote) SetPdfUrl(v string) {
	o.PdfUrl.Set(&v)
}

// GetIssuedAt returns the IssuedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CreditNote) GetIssuedAt() time.Time {
	if o == nil || o.IssuedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.IssuedAt.Get()
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNote) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedAt.Get(), o.IssuedAt.IsSet()
}

// SetIssuedAt sets field value
func (o *CreditNote) SetIssuedAt(v time.Time) {
	o.IssuedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreditNote) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreditNote) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CreditNote) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CreditNote) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CreditNote) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CreditNote) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["credit_note_number"] = o.CreditNoteNumber.Get()
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["tax_amount_minor"] = o.TaxAmountMinor.Get()
	toSerialize["currency"] = o.Currency
	toSerialize["status"] = o.Status
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["reason_code"] = o.ReasonCode.Get()
	toSerialize["credit_method"] = o.CreditMethod.Get()
	toSerialize["pdf_url"] = o.PdfUrl.Get()
	toSerialize["issued_at"] = o.IssuedAt.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNote) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"credit_note_number",
		"invoice_id",
		"customer_id",
		"amount_minor",
		"tax_amount_minor",
		"currency",
		"status",
		"reason",
		"reason_code",
		"credit_method",
		"pdf_url",
		"issued_at",
		"created_at",
		"updated_at",
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

	varCreditNote := _CreditNote{}

	err = json.Unmarshal(data, &varCreditNote)

	if err != nil {
		return err
	}

	*o = CreditNote(varCreditNote)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "credit_note_number")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "tax_amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "status")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "reason_code")
		delete(additionalProperties, "credit_method")
		delete(additionalProperties, "pdf_url")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNote struct {
	value *CreditNote
	isSet bool
}

func (v NullableCreditNote) Get() *CreditNote {
	return v.value
}

func (v *NullableCreditNote) Set(val *CreditNote) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNote) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNote(val *CreditNote) *NullableCreditNote {
	return &NullableCreditNote{value: val, isSet: true}
}

func (v NullableCreditNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
