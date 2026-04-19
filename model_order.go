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

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order An order represents a one-off purchase with lifecycle operations: create, pay, cancel, and fulfill.
type Order struct {
	// Unique identifier for the order
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Public-facing order ID
	PublicId string `json:"public_id"`
	// ID of the customer
	CustomerId NullableString `json:"customer_id,omitempty"`
	// Order status. draft: Being built. pending: Awaiting payment. paid: Fully paid. partially_paid: Partial payment received. refunded: Refunded. canceled: Canceled. fulfilled: Shipped/delivered.
	Status string `json:"status"`
	// Billing type. one_time: Single purchase. recurring: Subscription-based.
	BillingType string `json:"billing_type"`
	// Fulfillment status for physical goods.
	FulfillmentStatus NullableString `json:"fulfillment_status,omitempty"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Subtotal in cents
	Subtotal int32 `json:"subtotal"`
	// Total discount in cents
	Discount int32 `json:"discount"`
	// Total tax in cents
	Tax int32 `json:"tax"`
	// Total amount in cents
	Total int32 `json:"total"`
	// Amount paid in cents
	AmountPaid int32 `json:"amount_paid"`
	// Amount refunded in cents
	AmountRefunded int32 `json:"amount_refunded"`
	// Amount still due in cents
	AmountDue int32 `json:"amount_due"`
	LineItems []OrderLineItem `json:"line_items,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Notes NullableString `json:"notes,omitempty"`
	// ID of the linked invoice (Commercial Truth)
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	// ID of the linked subscription (for recurring fulfillment)
	SubscriptionId NullableString `json:"subscription_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PaidAt NullableTime `json:"paid_at,omitempty"`
	CanceledAt NullableTime `json:"canceled_at,omitempty"`
	FulfilledAt NullableTime `json:"fulfilled_at,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(id string, object string, publicId string, status string, billingType string, currency string, subtotal int32, discount int32, tax int32, total int32, amountPaid int32, amountRefunded int32, amountDue int32, createdAt time.Time, updatedAt time.Time) *Order {
	this := Order{}
	this.Id = id
	this.Object = object
	this.PublicId = publicId
	this.Status = status
	this.BillingType = billingType
	this.Currency = currency
	this.Subtotal = subtotal
	this.Discount = discount
	this.Tax = tax
	this.Total = total
	this.AmountPaid = amountPaid
	this.AmountRefunded = amountRefunded
	this.AmountDue = amountDue
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetId returns the Id field value
func (o *Order) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Order) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Order) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Order) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Order) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Order) SetObject(v string) {
	o.Object = v
}

// GetPublicId returns the PublicId field value
func (o *Order) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *Order) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *Order) SetPublicId(v string) {
	o.PublicId = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Order) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Order) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Order) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Order) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetStatus returns the Status field value
func (o *Order) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Order) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Order) SetStatus(v string) {
	o.Status = v
}

// GetBillingType returns the BillingType field value
func (o *Order) GetBillingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
func (o *Order) GetBillingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingType, true
}

// SetBillingType sets field value
func (o *Order) SetBillingType(v string) {
	o.BillingType = v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetFulfillmentStatus() string {
	if o == nil || IsNil(o.FulfillmentStatus.Get()) {
		var ret string
		return ret
	}
	return *o.FulfillmentStatus.Get()
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfillmentStatus.Get(), o.FulfillmentStatus.IsSet()
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *Order) HasFulfillmentStatus() bool {
	if o != nil && o.FulfillmentStatus.IsSet() {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given NullableString and assigns it to the FulfillmentStatus field.
func (o *Order) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus.Set(&v)
}
// SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil
func (o *Order) SetFulfillmentStatusNil() {
	o.FulfillmentStatus.Set(nil)
}

// UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
func (o *Order) UnsetFulfillmentStatus() {
	o.FulfillmentStatus.Unset()
}

// GetCurrency returns the Currency field value
func (o *Order) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Order) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Order) SetCurrency(v string) {
	o.Currency = v
}

// GetSubtotal returns the Subtotal field value
func (o *Order) GetSubtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *Order) GetSubtotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *Order) SetSubtotal(v int32) {
	o.Subtotal = v
}

// GetDiscount returns the Discount field value
func (o *Order) GetDiscount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value
// and a boolean to check if the value has been set.
func (o *Order) GetDiscountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discount, true
}

// SetDiscount sets field value
func (o *Order) SetDiscount(v int32) {
	o.Discount = v
}

// GetTax returns the Tax field value
func (o *Order) GetTax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tax
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
func (o *Order) GetTaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tax, true
}

// SetTax sets field value
func (o *Order) SetTax(v int32) {
	o.Tax = v
}

// GetTotal returns the Total field value
func (o *Order) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Order) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Order) SetTotal(v int32) {
	o.Total = v
}

// GetAmountPaid returns the AmountPaid field value
func (o *Order) GetAmountPaid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountPaid
}

// GetAmountPaidOk returns a tuple with the AmountPaid field value
// and a boolean to check if the value has been set.
func (o *Order) GetAmountPaidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountPaid, true
}

// SetAmountPaid sets field value
func (o *Order) SetAmountPaid(v int32) {
	o.AmountPaid = v
}

// GetAmountRefunded returns the AmountRefunded field value
func (o *Order) GetAmountRefunded() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountRefunded
}

// GetAmountRefundedOk returns a tuple with the AmountRefunded field value
// and a boolean to check if the value has been set.
func (o *Order) GetAmountRefundedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountRefunded, true
}

// SetAmountRefunded sets field value
func (o *Order) SetAmountRefunded(v int32) {
	o.AmountRefunded = v
}

// GetAmountDue returns the AmountDue field value
func (o *Order) GetAmountDue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDue
}

// GetAmountDueOk returns a tuple with the AmountDue field value
// and a boolean to check if the value has been set.
func (o *Order) GetAmountDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDue, true
}

// SetAmountDue sets field value
func (o *Order) SetAmountDue(v int32) {
	o.AmountDue = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Order) GetLineItems() []OrderLineItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []OrderLineItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetLineItemsOk() ([]OrderLineItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *Order) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []OrderLineItem and assigns it to the LineItems field.
func (o *Order) SetLineItems(v []OrderLineItem) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Order) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Order) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Order) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *Order) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *Order) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *Order) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *Order) UnsetNotes() {
	o.Notes.Unset()
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Order) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *Order) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *Order) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *Order) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Order) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *Order) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *Order) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *Order) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *Order) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Order) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Order) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Order) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Order) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Order) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetPaidAt() time.Time {
	if o == nil || IsNil(o.PaidAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PaidAt.Get()
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetPaidAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaidAt.Get(), o.PaidAt.IsSet()
}

// HasPaidAt returns a boolean if a field has been set.
func (o *Order) HasPaidAt() bool {
	if o != nil && o.PaidAt.IsSet() {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given NullableTime and assigns it to the PaidAt field.
func (o *Order) SetPaidAt(v time.Time) {
	o.PaidAt.Set(&v)
}
// SetPaidAtNil sets the value for PaidAt to be an explicit nil
func (o *Order) SetPaidAtNil() {
	o.PaidAt.Set(nil)
}

// UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
func (o *Order) UnsetPaidAt() {
	o.PaidAt.Unset()
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *Order) HasCanceledAt() bool {
	if o != nil && o.CanceledAt.IsSet() {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given NullableTime and assigns it to the CanceledAt field.
func (o *Order) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}
// SetCanceledAtNil sets the value for CanceledAt to be an explicit nil
func (o *Order) SetCanceledAtNil() {
	o.CanceledAt.Set(nil)
}

// UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
func (o *Order) UnsetCanceledAt() {
	o.CanceledAt.Unset()
}

// GetFulfilledAt returns the FulfilledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetFulfilledAt() time.Time {
	if o == nil || IsNil(o.FulfilledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FulfilledAt.Get()
}

// GetFulfilledAtOk returns a tuple with the FulfilledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetFulfilledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfilledAt.Get(), o.FulfilledAt.IsSet()
}

// HasFulfilledAt returns a boolean if a field has been set.
func (o *Order) HasFulfilledAt() bool {
	if o != nil && o.FulfilledAt.IsSet() {
		return true
	}

	return false
}

// SetFulfilledAt gets a reference to the given NullableTime and assigns it to the FulfilledAt field.
func (o *Order) SetFulfilledAt(v time.Time) {
	o.FulfilledAt.Set(&v)
}
// SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil
func (o *Order) SetFulfilledAtNil() {
	o.FulfilledAt.Set(nil)
}

// UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil
func (o *Order) UnsetFulfilledAt() {
	o.FulfilledAt.Unset()
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Order) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *Order) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *Order) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *Order) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["public_id"] = o.PublicId
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["billing_type"] = o.BillingType
	if o.FulfillmentStatus.IsSet() {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus.Get()
	}
	toSerialize["currency"] = o.Currency
	toSerialize["subtotal"] = o.Subtotal
	toSerialize["discount"] = o.Discount
	toSerialize["tax"] = o.Tax
	toSerialize["total"] = o.Total
	toSerialize["amount_paid"] = o.AmountPaid
	toSerialize["amount_refunded"] = o.AmountRefunded
	toSerialize["amount_due"] = o.AmountDue
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscription_id"] = o.SubscriptionId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.PaidAt.IsSet() {
		toSerialize["paid_at"] = o.PaidAt.Get()
	}
	if o.CanceledAt.IsSet() {
		toSerialize["canceled_at"] = o.CanceledAt.Get()
	}
	if o.FulfilledAt.IsSet() {
		toSerialize["fulfilled_at"] = o.FulfilledAt.Get()
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Order) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"public_id",
		"status",
		"billing_type",
		"currency",
		"subtotal",
		"discount",
		"tax",
		"total",
		"amount_paid",
		"amount_refunded",
		"amount_due",
		"created_at",
		"updated_at",
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

	varOrder := _Order{}

	err = json.Unmarshal(data, &varOrder)

	if err != nil {
		return err
	}

	*o = Order(varOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "billing_type")
		delete(additionalProperties, "fulfillment_status")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "subtotal")
		delete(additionalProperties, "discount")
		delete(additionalProperties, "tax")
		delete(additionalProperties, "total")
		delete(additionalProperties, "amount_paid")
		delete(additionalProperties, "amount_refunded")
		delete(additionalProperties, "amount_due")
		delete(additionalProperties, "line_items")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "paid_at")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "fulfilled_at")
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


