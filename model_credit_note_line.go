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
	"time"
	"fmt"
)

// checks if the CreditNoteLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNoteLine{}

// CreditNoteLine A single line item on a credit note. Each line represents a portion of the invoice that is being credited — either a whole invoice line or a prorated slice of one.
type CreditNoteLine struct {
	Id string `json:"id"`
	Object string `json:"object"`
	CreditNoteId string `json:"credit_note_id"`
	// The invoice line item this credit line applies to, when the credit was scoped to a specific line. Null for invoice-level credits.
	InvoiceLineItemId NullableString `json:"invoice_line_item_id"`
	Description string `json:"description"`
	Quantity int32 `json:"quantity"`
	// Per-unit credit amount in cents
	UnitAmountMinor int32 `json:"unit_amount_minor"`
	// Total credit for this line in cents
	TotalAmountMinor int32 `json:"total_amount_minor"`
	Metadata map[string]interface{} `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _CreditNoteLine CreditNoteLine

// NewCreditNoteLine instantiates a new CreditNoteLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNoteLine(id string, object string, creditNoteId string, invoiceLineItemId NullableString, description string, quantity int32, unitAmountMinor int32, totalAmountMinor int32, metadata map[string]interface{}, createdAt time.Time) *CreditNoteLine {
	this := CreditNoteLine{}
	this.Id = id
	this.Object = object
	this.CreditNoteId = creditNoteId
	this.InvoiceLineItemId = invoiceLineItemId
	this.Description = description
	this.Quantity = quantity
	this.UnitAmountMinor = unitAmountMinor
	this.TotalAmountMinor = totalAmountMinor
	this.Metadata = metadata
	this.CreatedAt = createdAt
	return &this
}

// NewCreditNoteLineWithDefaults instantiates a new CreditNoteLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNoteLineWithDefaults() *CreditNoteLine {
	this := CreditNoteLine{}
	return &this
}

// GetId returns the Id field value
func (o *CreditNoteLine) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreditNoteLine) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CreditNoteLine) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CreditNoteLine) SetObject(v string) {
	o.Object = v
}

// GetCreditNoteId returns the CreditNoteId field value
func (o *CreditNoteLine) GetCreditNoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreditNoteId
}

// GetCreditNoteIdOk returns a tuple with the CreditNoteId field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetCreditNoteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditNoteId, true
}

// SetCreditNoteId sets field value
func (o *CreditNoteLine) SetCreditNoteId(v string) {
	o.CreditNoteId = v
}

// GetInvoiceLineItemId returns the InvoiceLineItemId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNoteLine) GetInvoiceLineItemId() string {
	if o == nil || o.InvoiceLineItemId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceLineItemId.Get()
}

// GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNoteLine) GetInvoiceLineItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceLineItemId.Get(), o.InvoiceLineItemId.IsSet()
}

// SetInvoiceLineItemId sets field value
func (o *CreditNoteLine) SetInvoiceLineItemId(v string) {
	o.InvoiceLineItemId.Set(&v)
}

// GetDescription returns the Description field value
func (o *CreditNoteLine) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreditNoteLine) SetDescription(v string) {
	o.Description = v
}

// GetQuantity returns the Quantity field value
func (o *CreditNoteLine) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CreditNoteLine) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitAmountMinor returns the UnitAmountMinor field value
func (o *CreditNoteLine) GetUnitAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmountMinor
}

// GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetUnitAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmountMinor, true
}

// SetUnitAmountMinor sets field value
func (o *CreditNoteLine) SetUnitAmountMinor(v int32) {
	o.UnitAmountMinor = v
}

// GetTotalAmountMinor returns the TotalAmountMinor field value
func (o *CreditNoteLine) GetTotalAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmountMinor
}

// GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetTotalAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountMinor, true
}

// SetTotalAmountMinor sets field value
func (o *CreditNoteLine) SetTotalAmountMinor(v int32) {
	o.TotalAmountMinor = v
}

// GetMetadata returns the Metadata field value
func (o *CreditNoteLine) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *CreditNoteLine) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreditNoteLine) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreditNoteLine) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreditNoteLine) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o CreditNoteLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNoteLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["credit_note_id"] = o.CreditNoteId
	toSerialize["invoice_line_item_id"] = o.InvoiceLineItemId.Get()
	toSerialize["description"] = o.Description
	toSerialize["quantity"] = o.Quantity
	toSerialize["unit_amount_minor"] = o.UnitAmountMinor
	toSerialize["total_amount_minor"] = o.TotalAmountMinor
	toSerialize["metadata"] = o.Metadata
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNoteLine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"credit_note_id",
		"invoice_line_item_id",
		"description",
		"quantity",
		"unit_amount_minor",
		"total_amount_minor",
		"metadata",
		"created_at",
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

	varCreditNoteLine := _CreditNoteLine{}

	err = json.Unmarshal(data, &varCreditNoteLine)

	if err != nil {
		return err
	}

	*o = CreditNoteLine(varCreditNoteLine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "credit_note_id")
		delete(additionalProperties, "invoice_line_item_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unit_amount_minor")
		delete(additionalProperties, "total_amount_minor")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNoteLine struct {
	value *CreditNoteLine
	isSet bool
}

func (v NullableCreditNoteLine) Get() *CreditNoteLine {
	return v.value
}

func (v *NullableCreditNoteLine) Set(val *CreditNoteLine) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNoteLine) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNoteLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNoteLine(val *CreditNoteLine) *NullableCreditNoteLine {
	return &NullableCreditNoteLine{value: val, isSet: true}
}

func (v NullableCreditNoteLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNoteLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


