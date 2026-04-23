/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the TerminalPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalPayment{}

// TerminalPayment A card-present payment processed through a POS terminal, with device details and card entry mode.
type TerminalPayment struct {
	// Unique identifier for this terminal payment attempt
	Id string `json:"id"`
	// Associated invoice ID, or null for walk-in/ad-hoc payments
	InvoiceId NullableString `json:"invoice_id"`
	// The terminal device that processed (or is processing) this payment
	DeviceId NullableString `json:"device_id"`
	// Transaction type. sale: original charge. refund: money returned. void: pre-settlement cancellation.
	Type string `json:"type"`
	// Payment lifecycle status. requested: command sent, awaiting card. in_progress: terminal processing. approved: payment succeeded. declined: issuer declined. cancelled: merchant cancelled. error: terminal error. timed_out: no response within 3 minutes.
	Status string `json:"status"`
	// Amount in minor units (e.g., pence for GBP, cents for USD)
	AmountMinor float32 `json:"amount_minor"`
	// ISO 4217 currency code
	Currency string `json:"currency"`
	// Payment reference (invoice number or custom reference)
	Reference NullableString `json:"reference"`
	// Serial number of the PAX terminal that processed this payment
	TerminalSerial NullableString `json:"terminal_serial"`
	// Unique Transaction Identifier from the terminal
	Uti NullableString `json:"uti"`
	// Authorization code from the payment processor. Present when approved.
	AuthCode NullableString `json:"auth_code"`
	// Terminal response code. '00' indicates approval.
	ResponseCode NullableString `json:"response_code"`
	// Retrieval Reference Number for settlement reconciliation
	Rrn NullableString `json:"rrn"`
	// Card network (e.g., VISA, MASTERCARD, AMEX)
	CardScheme NullableString `json:"card_scheme"`
	// Masked card number — only the last 4 digits are visible
	MaskedPan NullableString `json:"masked_pan"`
	// How the card was read: contactless (NFC tap), emv_chip (chip insert), magnetic_stripe (swipe), manual_entry (keyed), or fallback (chip-to-swipe)
	EntryMode NullableString `json:"entry_mode"`
	// Human-readable error message when status is error or timed_out
	ErrorMessage NullableString `json:"error_message"`
	// ISO 8601 timestamp when the payment was initiated
	CreatedAt string `json:"created_at"`
	// ISO 8601 timestamp when the terminal returned a result, or null if still in progress
	CompletedAt NullableString `json:"completed_at"`
	AdditionalProperties map[string]interface{}
}

type _TerminalPayment TerminalPayment

// NewTerminalPayment instantiates a new TerminalPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalPayment(id string, invoiceId NullableString, deviceId NullableString, type_ string, status string, amountMinor float32, currency string, reference NullableString, terminalSerial NullableString, uti NullableString, authCode NullableString, responseCode NullableString, rrn NullableString, cardScheme NullableString, maskedPan NullableString, entryMode NullableString, errorMessage NullableString, createdAt string, completedAt NullableString) *TerminalPayment {
	this := TerminalPayment{}
	this.Id = id
	this.InvoiceId = invoiceId
	this.DeviceId = deviceId
	this.Type = type_
	this.Status = status
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.Reference = reference
	this.TerminalSerial = terminalSerial
	this.Uti = uti
	this.AuthCode = authCode
	this.ResponseCode = responseCode
	this.Rrn = rrn
	this.CardScheme = cardScheme
	this.MaskedPan = maskedPan
	this.EntryMode = entryMode
	this.ErrorMessage = errorMessage
	this.CreatedAt = createdAt
	this.CompletedAt = completedAt
	return &this
}

// NewTerminalPaymentWithDefaults instantiates a new TerminalPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalPaymentWithDefaults() *TerminalPayment {
	this := TerminalPayment{}
	return &this
}

// GetId returns the Id field value
func (o *TerminalPayment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TerminalPayment) SetId(v string) {
	o.Id = v
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *TerminalPayment) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetDeviceId returns the DeviceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetDeviceId() string {
	if o == nil || o.DeviceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// SetDeviceId sets field value
func (o *TerminalPayment) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}

// GetType returns the Type field value
func (o *TerminalPayment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerminalPayment) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *TerminalPayment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TerminalPayment) SetStatus(v string) {
	o.Status = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *TerminalPayment) GetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *TerminalPayment) SetAmountMinor(v float32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *TerminalPayment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *TerminalPayment) SetCurrency(v string) {
	o.Currency = v
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetReference() string {
	if o == nil || o.Reference.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// SetReference sets field value
func (o *TerminalPayment) SetReference(v string) {
	o.Reference.Set(&v)
}

// GetTerminalSerial returns the TerminalSerial field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetTerminalSerial() string {
	if o == nil || o.TerminalSerial.Get() == nil {
		var ret string
		return ret
	}

	return *o.TerminalSerial.Get()
}

// GetTerminalSerialOk returns a tuple with the TerminalSerial field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetTerminalSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminalSerial.Get(), o.TerminalSerial.IsSet()
}

// SetTerminalSerial sets field value
func (o *TerminalPayment) SetTerminalSerial(v string) {
	o.TerminalSerial.Set(&v)
}

// GetUti returns the Uti field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetUti() string {
	if o == nil || o.Uti.Get() == nil {
		var ret string
		return ret
	}

	return *o.Uti.Get()
}

// GetUtiOk returns a tuple with the Uti field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetUtiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uti.Get(), o.Uti.IsSet()
}

// SetUti sets field value
func (o *TerminalPayment) SetUti(v string) {
	o.Uti.Set(&v)
}

// GetAuthCode returns the AuthCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetAuthCode() string {
	if o == nil || o.AuthCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthCode.Get()
}

// GetAuthCodeOk returns a tuple with the AuthCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetAuthCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCode.Get(), o.AuthCode.IsSet()
}

// SetAuthCode sets field value
func (o *TerminalPayment) SetAuthCode(v string) {
	o.AuthCode.Set(&v)
}

// GetResponseCode returns the ResponseCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetResponseCode() string {
	if o == nil || o.ResponseCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResponseCode.Get()
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCode.Get(), o.ResponseCode.IsSet()
}

// SetResponseCode sets field value
func (o *TerminalPayment) SetResponseCode(v string) {
	o.ResponseCode.Set(&v)
}

// GetRrn returns the Rrn field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetRrn() string {
	if o == nil || o.Rrn.Get() == nil {
		var ret string
		return ret
	}

	return *o.Rrn.Get()
}

// GetRrnOk returns a tuple with the Rrn field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetRrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rrn.Get(), o.Rrn.IsSet()
}

// SetRrn sets field value
func (o *TerminalPayment) SetRrn(v string) {
	o.Rrn.Set(&v)
}

// GetCardScheme returns the CardScheme field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetCardScheme() string {
	if o == nil || o.CardScheme.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardScheme.Get()
}

// GetCardSchemeOk returns a tuple with the CardScheme field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetCardSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardScheme.Get(), o.CardScheme.IsSet()
}

// SetCardScheme sets field value
func (o *TerminalPayment) SetCardScheme(v string) {
	o.CardScheme.Set(&v)
}

// GetMaskedPan returns the MaskedPan field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetMaskedPan() string {
	if o == nil || o.MaskedPan.Get() == nil {
		var ret string
		return ret
	}

	return *o.MaskedPan.Get()
}

// GetMaskedPanOk returns a tuple with the MaskedPan field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetMaskedPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaskedPan.Get(), o.MaskedPan.IsSet()
}

// SetMaskedPan sets field value
func (o *TerminalPayment) SetMaskedPan(v string) {
	o.MaskedPan.Set(&v)
}

// GetEntryMode returns the EntryMode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetEntryMode() string {
	if o == nil || o.EntryMode.Get() == nil {
		var ret string
		return ret
	}

	return *o.EntryMode.Get()
}

// GetEntryModeOk returns a tuple with the EntryMode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetEntryModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryMode.Get(), o.EntryMode.IsSet()
}

// SetEntryMode sets field value
func (o *TerminalPayment) SetEntryMode(v string) {
	o.EntryMode.Set(&v)
}

// GetErrorMessage returns the ErrorMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// SetErrorMessage sets field value
func (o *TerminalPayment) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *TerminalPayment) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TerminalPayment) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TerminalPayment) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalPayment) GetCompletedAt() string {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalPayment) GetCompletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *TerminalPayment) SetCompletedAt(v string) {
	o.CompletedAt.Set(&v)
}

func (o TerminalPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["device_id"] = o.DeviceId.Get()
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["reference"] = o.Reference.Get()
	toSerialize["terminal_serial"] = o.TerminalSerial.Get()
	toSerialize["uti"] = o.Uti.Get()
	toSerialize["auth_code"] = o.AuthCode.Get()
	toSerialize["response_code"] = o.ResponseCode.Get()
	toSerialize["rrn"] = o.Rrn.Get()
	toSerialize["card_scheme"] = o.CardScheme.Get()
	toSerialize["masked_pan"] = o.MaskedPan.Get()
	toSerialize["entry_mode"] = o.EntryMode.Get()
	toSerialize["error_message"] = o.ErrorMessage.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["completed_at"] = o.CompletedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TerminalPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"invoice_id",
		"device_id",
		"type",
		"status",
		"amount_minor",
		"currency",
		"reference",
		"terminal_serial",
		"uti",
		"auth_code",
		"response_code",
		"rrn",
		"card_scheme",
		"masked_pan",
		"entry_mode",
		"error_message",
		"created_at",
		"completed_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTerminalPayment := _TerminalPayment{}

	err = json.Unmarshal(data, &varTerminalPayment)

	if err != nil {
		return err
	}

	*o = TerminalPayment(varTerminalPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "device_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "terminal_serial")
		delete(additionalProperties, "uti")
		delete(additionalProperties, "auth_code")
		delete(additionalProperties, "response_code")
		delete(additionalProperties, "rrn")
		delete(additionalProperties, "card_scheme")
		delete(additionalProperties, "masked_pan")
		delete(additionalProperties, "entry_mode")
		delete(additionalProperties, "error_message")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerminalPayment struct {
	value *TerminalPayment
	isSet bool
}

func (v NullableTerminalPayment) Get() *TerminalPayment {
	return v.value
}

func (v *NullableTerminalPayment) Set(val *TerminalPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalPayment(val *TerminalPayment) *NullableTerminalPayment {
	return &NullableTerminalPayment{value: val, isSet: true}
}

func (v NullableTerminalPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


