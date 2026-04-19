/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the PortalInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortalInvoice{}

// PortalInvoice A customer-facing invoice projection. Lists and retrievals are scoped to invoices owned by the authenticated customer.
type PortalInvoice struct {
	Id string `json:"id"`
	Object string `json:"object"`
	InvoiceNumber NullableString `json:"invoice_number"`
	Status string `json:"status"`
	Currency NullableString `json:"currency"`
	Subtotal NullableInt32 `json:"subtotal"`
	TaxAmount NullableInt32 `json:"tax_amount"`
	DiscountAmount NullableInt32 `json:"discount_amount"`
	Total int32 `json:"total"`
	AmountPaid NullableInt32 `json:"amount_paid"`
	AmountDue NullableInt32 `json:"amount_due"`
	InvoiceDate NullableTime `json:"invoice_date"`
	DueDate NullableTime `json:"due_date"`
	PaidAt NullableTime `json:"paid_at"`
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _PortalInvoice PortalInvoice

// NewPortalInvoice instantiates a new PortalInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalInvoice(id string, object string, invoiceNumber NullableString, status string, currency NullableString, subtotal NullableInt32, taxAmount NullableInt32, discountAmount NullableInt32, total int32, amountPaid NullableInt32, amountDue NullableInt32, invoiceDate NullableTime, dueDate NullableTime, paidAt NullableTime, createdAt time.Time) *PortalInvoice {
	this := PortalInvoice{}
	this.Id = id
	this.Object = object
	this.InvoiceNumber = invoiceNumber
	this.Status = status
	this.Currency = currency
	this.Subtotal = subtotal
	this.TaxAmount = taxAmount
	this.DiscountAmount = discountAmount
	this.Total = total
	this.AmountPaid = amountPaid
	this.AmountDue = amountDue
	this.InvoiceDate = invoiceDate
	this.DueDate = dueDate
	this.PaidAt = paidAt
	this.CreatedAt = createdAt
	return &this
}

// NewPortalInvoiceWithDefaults instantiates a new PortalInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalInvoiceWithDefaults() *PortalInvoice {
	this := PortalInvoice{}
	return &this
}

// GetId returns the Id field value
func (o *PortalInvoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PortalInvoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PortalInvoice) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PortalInvoice) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PortalInvoice) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PortalInvoice) SetObject(v string) {
	o.Object = v
}

// GetInvoiceNumber returns the InvoiceNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalInvoice) GetInvoiceNumber() string {
	if o == nil || o.InvoiceNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceNumber.Get()
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceNumber.Get(), o.InvoiceNumber.IsSet()
}

// SetInvoiceNumber sets field value
func (o *PortalInvoice) SetInvoiceNumber(v string) {
	o.InvoiceNumber.Set(&v)
}

// GetStatus returns the Status field value
func (o *PortalInvoice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PortalInvoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PortalInvoice) SetStatus(v string) {
	o.Status = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalInvoice) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *PortalInvoice) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetSubtotal returns the Subtotal field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalInvoice) GetSubtotal() int32 {
	if o == nil || o.Subtotal.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Subtotal.Get()
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetSubtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subtotal.Get(), o.Subtotal.IsSet()
}

// SetSubtotal sets field value
func (o *PortalInvoice) SetSubtotal(v int32) {
	o.Subtotal.Set(&v)
}

// GetTaxAmount returns the TaxAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalInvoice) GetTaxAmount() int32 {
	if o == nil || o.TaxAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TaxAmount.Get()
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetTaxAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxAmount.Get(), o.TaxAmount.IsSet()
}

// SetTaxAmount sets field value
func (o *PortalInvoice) SetTaxAmount(v int32) {
	o.TaxAmount.Set(&v)
}

// GetDiscountAmount returns the DiscountAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalInvoice) GetDiscountAmount() int32 {
	if o == nil || o.DiscountAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DiscountAmount.Get()
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetDiscountAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountAmount.Get(), o.DiscountAmount.IsSet()
}

// SetDiscountAmount sets field value
func (o *PortalInvoice) SetDiscountAmount(v int32) {
	o.DiscountAmount.Set(&v)
}

// GetTotal returns the Total field value
func (o *PortalInvoice) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PortalInvoice) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PortalInvoice) SetTotal(v int32) {
	o.Total = v
}

// GetAmountPaid returns the AmountPaid field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalInvoice) GetAmountPaid() int32 {
	if o == nil || o.AmountPaid.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AmountPaid.Get()
}

// GetAmountPaidOk returns a tuple with the AmountPaid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetAmountPaidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountPaid.Get(), o.AmountPaid.IsSet()
}

// SetAmountPaid sets field value
func (o *PortalInvoice) SetAmountPaid(v int32) {
	o.AmountPaid.Set(&v)
}

// GetAmountDue returns the AmountDue field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalInvoice) GetAmountDue() int32 {
	if o == nil || o.AmountDue.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AmountDue.Get()
}

// GetAmountDueOk returns a tuple with the AmountDue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetAmountDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountDue.Get(), o.AmountDue.IsSet()
}

// SetAmountDue sets field value
func (o *PortalInvoice) SetAmountDue(v int32) {
	o.AmountDue.Set(&v)
}

// GetInvoiceDate returns the InvoiceDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalInvoice) GetInvoiceDate() time.Time {
	if o == nil || o.InvoiceDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.InvoiceDate.Get()
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetInvoiceDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceDate.Get(), o.InvoiceDate.IsSet()
}

// SetInvoiceDate sets field value
func (o *PortalInvoice) SetInvoiceDate(v time.Time) {
	o.InvoiceDate.Set(&v)
}

// GetDueDate returns the DueDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalInvoice) GetDueDate() time.Time {
	if o == nil || o.DueDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// SetDueDate sets field value
func (o *PortalInvoice) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}

// GetPaidAt returns the PaidAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalInvoice) GetPaidAt() time.Time {
	if o == nil || o.PaidAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalInvoice) GetPaidAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// SetPaidAt sets field value
func (o *PortalInvoice) SetPaidAt(v time.Time) {
	o.PaidAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *PortalInvoice) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PortalInvoice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PortalInvoice) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PortalInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortalInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["invoice_number"] = o.InvoiceNumber.Get()
	toSerialize["status"] = o.Status
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["subtotal"] = o.Subtotal.Get()
	toSerialize["tax_amount"] = o.TaxAmount.Get()
	toSerialize["discount_amount"] = o.DiscountAmount.Get()
	toSerialize["total"] = o.Total
	toSerialize["amount_paid"] = o.AmountPaid.Get()
	toSerialize["amount_due"] = o.AmountDue.Get()
	toSerialize["invoice_date"] = o.InvoiceDate.Get()
	toSerialize["due_date"] = o.DueDate.Get()
	toSerialize["paid_at"] = o.PaidAt.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortalInvoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"invoice_number",
		"status",
		"currency",
		"subtotal",
		"tax_amount",
		"discount_amount",
		"total",
		"amount_paid",
		"amount_due",
		"invoice_date",
		"due_date",
		"paid_at",
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

	varPortalInvoice := _PortalInvoice{}

	err = json.Unmarshal(data, &varPortalInvoice)

	if err != nil {
		return err
	}

	*o = PortalInvoice(varPortalInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "invoice_number")
		delete(additionalProperties, "status")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "subtotal")
		delete(additionalProperties, "tax_amount")
		delete(additionalProperties, "discount_amount")
		delete(additionalProperties, "total")
		delete(additionalProperties, "amount_paid")
		delete(additionalProperties, "amount_due")
		delete(additionalProperties, "invoice_date")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "paid_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortalInvoice struct {
	value *PortalInvoice
	isSet bool
}

func (v NullablePortalInvoice) Get() *PortalInvoice {
	return v.value
}

func (v *NullablePortalInvoice) Set(val *PortalInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalInvoice(val *PortalInvoice) *NullablePortalInvoice {
	return &NullablePortalInvoice{value: val, isSet: true}
}

func (v NullablePortalInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


