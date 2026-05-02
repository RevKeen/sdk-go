/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the CreditNotePreviewInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNotePreviewInvoice{}

// CreditNotePreviewInvoice struct for CreditNotePreviewInvoice
type CreditNotePreviewInvoice struct {
	Id string `json:"id"`
	TotalMinor int32 `json:"total_minor"`
	AmountPaidMinor int32 `json:"amount_paid_minor"`
	AmountDueMinor int32 `json:"amount_due_minor"`
	PreviouslyCreditedMinor int32 `json:"previously_credited_minor"`
	MaxCreditableMinor int32 `json:"max_creditable_minor"`
	AdditionalProperties map[string]interface{}
}

type _CreditNotePreviewInvoice CreditNotePreviewInvoice

// NewCreditNotePreviewInvoice instantiates a new CreditNotePreviewInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNotePreviewInvoice(id string, totalMinor int32, amountPaidMinor int32, amountDueMinor int32, previouslyCreditedMinor int32, maxCreditableMinor int32) *CreditNotePreviewInvoice {
	this := CreditNotePreviewInvoice{}
	this.Id = id
	this.TotalMinor = totalMinor
	this.AmountPaidMinor = amountPaidMinor
	this.AmountDueMinor = amountDueMinor
	this.PreviouslyCreditedMinor = previouslyCreditedMinor
	this.MaxCreditableMinor = maxCreditableMinor
	return &this
}

// NewCreditNotePreviewInvoiceWithDefaults instantiates a new CreditNotePreviewInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNotePreviewInvoiceWithDefaults() *CreditNotePreviewInvoice {
	this := CreditNotePreviewInvoice{}
	return &this
}

// GetId returns the Id field value
func (o *CreditNotePreviewInvoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreditNotePreviewInvoice) SetId(v string) {
	o.Id = v
}

// GetTotalMinor returns the TotalMinor field value
func (o *CreditNotePreviewInvoice) GetTotalMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMinor
}

// GetTotalMinorOk returns a tuple with the TotalMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetTotalMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinor, true
}

// SetTotalMinor sets field value
func (o *CreditNotePreviewInvoice) SetTotalMinor(v int32) {
	o.TotalMinor = v
}

// GetAmountPaidMinor returns the AmountPaidMinor field value
func (o *CreditNotePreviewInvoice) GetAmountPaidMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountPaidMinor
}

// GetAmountPaidMinorOk returns a tuple with the AmountPaidMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetAmountPaidMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountPaidMinor, true
}

// SetAmountPaidMinor sets field value
func (o *CreditNotePreviewInvoice) SetAmountPaidMinor(v int32) {
	o.AmountPaidMinor = v
}

// GetAmountDueMinor returns the AmountDueMinor field value
func (o *CreditNotePreviewInvoice) GetAmountDueMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDueMinor
}

// GetAmountDueMinorOk returns a tuple with the AmountDueMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetAmountDueMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDueMinor, true
}

// SetAmountDueMinor sets field value
func (o *CreditNotePreviewInvoice) SetAmountDueMinor(v int32) {
	o.AmountDueMinor = v
}

// GetPreviouslyCreditedMinor returns the PreviouslyCreditedMinor field value
func (o *CreditNotePreviewInvoice) GetPreviouslyCreditedMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PreviouslyCreditedMinor
}

// GetPreviouslyCreditedMinorOk returns a tuple with the PreviouslyCreditedMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetPreviouslyCreditedMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviouslyCreditedMinor, true
}

// SetPreviouslyCreditedMinor sets field value
func (o *CreditNotePreviewInvoice) SetPreviouslyCreditedMinor(v int32) {
	o.PreviouslyCreditedMinor = v
}

// GetMaxCreditableMinor returns the MaxCreditableMinor field value
func (o *CreditNotePreviewInvoice) GetMaxCreditableMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxCreditableMinor
}

// GetMaxCreditableMinorOk returns a tuple with the MaxCreditableMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNotePreviewInvoice) GetMaxCreditableMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCreditableMinor, true
}

// SetMaxCreditableMinor sets field value
func (o *CreditNotePreviewInvoice) SetMaxCreditableMinor(v int32) {
	o.MaxCreditableMinor = v
}

func (o CreditNotePreviewInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNotePreviewInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["total_minor"] = o.TotalMinor
	toSerialize["amount_paid_minor"] = o.AmountPaidMinor
	toSerialize["amount_due_minor"] = o.AmountDueMinor
	toSerialize["previously_credited_minor"] = o.PreviouslyCreditedMinor
	toSerialize["max_creditable_minor"] = o.MaxCreditableMinor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNotePreviewInvoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"total_minor",
		"amount_paid_minor",
		"amount_due_minor",
		"previously_credited_minor",
		"max_creditable_minor",
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

	varCreditNotePreviewInvoice := _CreditNotePreviewInvoice{}

	err = json.Unmarshal(data, &varCreditNotePreviewInvoice)

	if err != nil {
		return err
	}

	*o = CreditNotePreviewInvoice(varCreditNotePreviewInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "total_minor")
		delete(additionalProperties, "amount_paid_minor")
		delete(additionalProperties, "amount_due_minor")
		delete(additionalProperties, "previously_credited_minor")
		delete(additionalProperties, "max_creditable_minor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNotePreviewInvoice struct {
	value *CreditNotePreviewInvoice
	isSet bool
}

func (v NullableCreditNotePreviewInvoice) Get() *CreditNotePreviewInvoice {
	return v.value
}

func (v *NullableCreditNotePreviewInvoice) Set(val *CreditNotePreviewInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNotePreviewInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNotePreviewInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNotePreviewInvoice(val *CreditNotePreviewInvoice) *NullableCreditNotePreviewInvoice {
	return &NullableCreditNotePreviewInvoice{value: val, isSet: true}
}

func (v NullableCreditNotePreviewInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNotePreviewInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


