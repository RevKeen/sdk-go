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

// checks if the CreditNotePreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNotePreview{}

// CreditNotePreview A non-persistent preview of a credit note. Use this to show a customer (or merchant operator) what the credit would do before committing to creating it.
type CreditNotePreview struct {
	Object string `json:"object"`
	InvoiceId string `json:"invoice_id"`
	ProposedAmountMinor int32 `json:"proposed_amount_minor"`
	ProposedTaxAmountMinor NullableInt32 `json:"proposed_tax_amount_minor"`
	Currency string `json:"currency"`
	Invoice CreditNotePreviewInvoice `json:"invoice"`
	AfterCredit CreditNotePreviewAfterCredit `json:"after_credit"`
	// True if `amount_minor` exceeds the invoice's remaining creditable amount. A real create call would be rejected.
	ExceedsMaxCreditable bool `json:"exceeds_max_creditable"`
	CreditMethod string `json:"credit_method"`
	AdditionalProperties map[string]interface{}
}

type _CreditNotePreview CreditNotePreview

// NewCreditNotePreview instantiates a new CreditNotePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNotePreview(object string, invoiceId string, proposedAmountMinor int32, proposedTaxAmountMinor NullableInt32, currency string, invoice CreditNotePreviewInvoice, afterCredit CreditNotePreviewAfterCredit, exceedsMaxCreditable bool, creditMethod string) *CreditNotePreview {
	this := CreditNotePreview{}
	this.Object = object
	this.InvoiceId = invoiceId
	this.ProposedAmountMinor = proposedAmountMinor
	this.ProposedTaxAmountMinor = proposedTaxAmountMinor
	this.Currency = currency
	this.Invoice = invoice
	this.AfterCredit = afterCredit
	this.ExceedsMaxCreditable = exceedsMaxCreditable
	this.CreditMethod = creditMethod
	return &this
}

// NewCreditNotePreviewWithDefaults instantiates a new CreditNotePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNotePreviewWithDefaults() *CreditNotePreview {
	this := CreditNotePreview{}
	return &this
}

// GetObject returns the Object field value
func (o *CreditNotePreview) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreditNotePreview) SetObject(v string) {
	o.Object = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *CreditNotePreview) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *CreditNotePreview) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetProposedAmountMinor returns the ProposedAmountMinor field value
func (o *CreditNotePreview) GetProposedAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProposedAmountMinor
}

// GetProposedAmountMinorOk returns a tuple with the ProposedAmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetProposedAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProposedAmountMinor, true
}

// SetProposedAmountMinor sets field value
func (o *CreditNotePreview) SetProposedAmountMinor(v int32) {
	o.ProposedAmountMinor = v
}

// GetProposedTaxAmountMinor returns the ProposedTaxAmountMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CreditNotePreview) GetProposedTaxAmountMinor() int32 {
	if o == nil || o.ProposedTaxAmountMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ProposedTaxAmountMinor.Get()
}

// GetProposedTaxAmountMinorOk returns a tuple with the ProposedTaxAmountMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNotePreview) GetProposedTaxAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProposedTaxAmountMinor.Get(), o.ProposedTaxAmountMinor.IsSet()
}

// SetProposedTaxAmountMinor sets field value
func (o *CreditNotePreview) SetProposedTaxAmountMinor(v int32) {
	o.ProposedTaxAmountMinor.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *CreditNotePreview) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreditNotePreview) SetCurrency(v string) {
	o.Currency = v
}

// GetInvoice returns the Invoice field value
func (o *CreditNotePreview) GetInvoice() CreditNotePreviewInvoice {
	if o == nil {
		var ret CreditNotePreviewInvoice
		return ret
	}

	return o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetInvoiceOk() (*CreditNotePreviewInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Invoice, true
}

// SetInvoice sets field value
func (o *CreditNotePreview) SetInvoice(v CreditNotePreviewInvoice) {
	o.Invoice = v
}

// GetAfterCredit returns the AfterCredit field value
func (o *CreditNotePreview) GetAfterCredit() CreditNotePreviewAfterCredit {
	if o == nil {
		var ret CreditNotePreviewAfterCredit
		return ret
	}

	return o.AfterCredit
}

// GetAfterCreditOk returns a tuple with the AfterCredit field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetAfterCreditOk() (*CreditNotePreviewAfterCredit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfterCredit, true
}

// SetAfterCredit sets field value
func (o *CreditNotePreview) SetAfterCredit(v CreditNotePreviewAfterCredit) {
	o.AfterCredit = v
}

// GetExceedsMaxCreditable returns the ExceedsMaxCreditable field value
func (o *CreditNotePreview) GetExceedsMaxCreditable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExceedsMaxCreditable
}

// GetExceedsMaxCreditableOk returns a tuple with the ExceedsMaxCreditable field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetExceedsMaxCreditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceedsMaxCreditable, true
}

// SetExceedsMaxCreditable sets field value
func (o *CreditNotePreview) SetExceedsMaxCreditable(v bool) {
	o.ExceedsMaxCreditable = v
}

// GetCreditMethod returns the CreditMethod field value
func (o *CreditNotePreview) GetCreditMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreditMethod
}

// GetCreditMethodOk returns a tuple with the CreditMethod field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreview) GetCreditMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditMethod, true
}

// SetCreditMethod sets field value
func (o *CreditNotePreview) SetCreditMethod(v string) {
	o.CreditMethod = v
}

func (o CreditNotePreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNotePreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["proposed_amount_minor"] = o.ProposedAmountMinor
	toSerialize["proposed_tax_amount_minor"] = o.ProposedTaxAmountMinor.Get()
	toSerialize["currency"] = o.Currency
	toSerialize["invoice"] = o.Invoice
	toSerialize["after_credit"] = o.AfterCredit
	toSerialize["exceeds_max_creditable"] = o.ExceedsMaxCreditable
	toSerialize["credit_method"] = o.CreditMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNotePreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"invoice_id",
		"proposed_amount_minor",
		"proposed_tax_amount_minor",
		"currency",
		"invoice",
		"after_credit",
		"exceeds_max_creditable",
		"credit_method",
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

	varCreditNotePreview := _CreditNotePreview{}

	err = json.Unmarshal(data, &varCreditNotePreview)

	if err != nil {
		return err
	}

	*o = CreditNotePreview(varCreditNotePreview)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "proposed_amount_minor")
		delete(additionalProperties, "proposed_tax_amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "invoice")
		delete(additionalProperties, "after_credit")
		delete(additionalProperties, "exceeds_max_creditable")
		delete(additionalProperties, "credit_method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNotePreview struct {
	value *CreditNotePreview
	isSet bool
}

func (v NullableCreditNotePreview) Get() *CreditNotePreview {
	return v.value
}

func (v *NullableCreditNotePreview) Set(val *CreditNotePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNotePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNotePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNotePreview(val *CreditNotePreview) *NullableCreditNotePreview {
	return &NullableCreditNotePreview{value: val, isSet: true}
}

func (v NullableCreditNotePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNotePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


