/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-01-15` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-01-15
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Payment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payment{}

// Payment A completed or in-progress money movement. Payments represent funds transferred between a customer and your account across all channels — online, terminal, and recurring.
type Payment struct {
	// Unique identifier for the payment
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Associated invoice ID
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	// Customer ID
	CustomerId NullableString `json:"customer_id,omitempty"`
	// Payment status. pending: Processing. authorized: Funds reserved. captured/succeeded: Funds collected. failed: Payment failed. voided: Canceled before capture. refunded/partially_refunded: Returned to customer.
	Status string `json:"status"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Payment amount in cents
	Amount int32 `json:"amount"`
	// Amount captured in cents (for auth-capture flow)
	AmountCaptured NullableInt32 `json:"amount_captured,omitempty"`
	// Amount refunded in cents
	AmountRefunded *int32 `json:"amount_refunded,omitempty"`
	// Payment gateway (e.g., 'nmi', 'stripe')
	Gateway NullableString `json:"gateway,omitempty"`
	// Gateway's transaction reference
	GatewayTransactionId NullableString `json:"gateway_transaction_id,omitempty"`
	// Gateway response code
	GatewayResponseCode NullableString `json:"gateway_response_code,omitempty"`
	// Gateway response message
	GatewayResponseText NullableString `json:"gateway_response_text,omitempty"`
	// Payment method type (card, ach, wallet)
	PaymentMethodType NullableString `json:"payment_method_type,omitempty"`
	// Card brand (visa, mastercard, etc.)
	CardBrand NullableString `json:"card_brand,omitempty"`
	// Last 4 digits of card
	CardLastFour NullableString `json:"card_last_four,omitempty"`
	// Payment channel. card_present: terminal/POS payment. card_not_present: online or recurring payment. bank_transfer: ACH/direct debit. manual: manually recorded.
	PaymentChannel NullableString `json:"payment_channel,omitempty"`
	// Card entry mode for card-present (terminal) payments. null for online payments.
	EntryMode NullableString `json:"entry_mode,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorizedAt NullableTime `json:"authorized_at,omitempty"`
	CapturedAt NullableTime `json:"captured_at,omitempty"`
	VoidedAt NullableTime `json:"voided_at,omitempty"`
	RefundedAt NullableTime `json:"refunded_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Payment Payment

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment(id string, object string, status string, currency string, amount int32, createdAt time.Time, updatedAt time.Time) *Payment {
	this := Payment{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.Currency = currency
	this.Amount = amount
	var amountRefunded int32 = 0
	this.AmountRefunded = &amountRefunded
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	var amountRefunded int32 = 0
	this.AmountRefunded = &amountRefunded
	return &this
}

// GetId returns the Id field value
func (o *Payment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payment) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Payment) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Payment) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Payment) SetObject(v string) {
	o.Object = v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Payment) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *Payment) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *Payment) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *Payment) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Payment) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Payment) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Payment) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Payment) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetStatus returns the Status field value
func (o *Payment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payment) SetStatus(v string) {
	o.Status = v
}

// GetCurrency returns the Currency field value
func (o *Payment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Payment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Payment) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *Payment) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Payment) SetAmount(v int32) {
	o.Amount = v
}

// GetAmountCaptured returns the AmountCaptured field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetAmountCaptured() int32 {
	if o == nil || IsNil(o.AmountCaptured.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountCaptured.Get()
}

// GetAmountCapturedOk returns a tuple with the AmountCaptured field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetAmountCapturedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCaptured.Get(), o.AmountCaptured.IsSet()
}

// HasAmountCaptured returns a boolean if a field has been set.
func (o *Payment) HasAmountCaptured() bool {
	if o != nil && o.AmountCaptured.IsSet() {
		return true
	}

	return false
}

// SetAmountCaptured gets a reference to the given NullableInt32 and assigns it to the AmountCaptured field.
func (o *Payment) SetAmountCaptured(v int32) {
	o.AmountCaptured.Set(&v)
}
// SetAmountCapturedNil sets the value for AmountCaptured to be an explicit nil
func (o *Payment) SetAmountCapturedNil() {
	o.AmountCaptured.Set(nil)
}

// UnsetAmountCaptured ensures that no value is present for AmountCaptured, not even an explicit nil
func (o *Payment) UnsetAmountCaptured() {
	o.AmountCaptured.Unset()
}

// GetAmountRefunded returns the AmountRefunded field value if set, zero value otherwise.
func (o *Payment) GetAmountRefunded() int32 {
	if o == nil || IsNil(o.AmountRefunded) {
		var ret int32
		return ret
	}
	return *o.AmountRefunded
}

// GetAmountRefundedOk returns a tuple with the AmountRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetAmountRefundedOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountRefunded) {
		return nil, false
	}
	return o.AmountRefunded, true
}

// HasAmountRefunded returns a boolean if a field has been set.
func (o *Payment) HasAmountRefunded() bool {
	if o != nil && !IsNil(o.AmountRefunded) {
		return true
	}

	return false
}

// SetAmountRefunded gets a reference to the given int32 and assigns it to the AmountRefunded field.
func (o *Payment) SetAmountRefunded(v int32) {
	o.AmountRefunded = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *Payment) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *Payment) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *Payment) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *Payment) UnsetGateway() {
	o.Gateway.Unset()
}

// GetGatewayTransactionId returns the GatewayTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetGatewayTransactionId() string {
	if o == nil || IsNil(o.GatewayTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayTransactionId.Get()
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTransactionId.Get(), o.GatewayTransactionId.IsSet()
}

// HasGatewayTransactionId returns a boolean if a field has been set.
func (o *Payment) HasGatewayTransactionId() bool {
	if o != nil && o.GatewayTransactionId.IsSet() {
		return true
	}

	return false
}

// SetGatewayTransactionId gets a reference to the given NullableString and assigns it to the GatewayTransactionId field.
func (o *Payment) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId.Set(&v)
}
// SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil
func (o *Payment) SetGatewayTransactionIdNil() {
	o.GatewayTransactionId.Set(nil)
}

// UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
func (o *Payment) UnsetGatewayTransactionId() {
	o.GatewayTransactionId.Unset()
}

// GetGatewayResponseCode returns the GatewayResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetGatewayResponseCode() string {
	if o == nil || IsNil(o.GatewayResponseCode.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayResponseCode.Get()
}

// GetGatewayResponseCodeOk returns a tuple with the GatewayResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetGatewayResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayResponseCode.Get(), o.GatewayResponseCode.IsSet()
}

// HasGatewayResponseCode returns a boolean if a field has been set.
func (o *Payment) HasGatewayResponseCode() bool {
	if o != nil && o.GatewayResponseCode.IsSet() {
		return true
	}

	return false
}

// SetGatewayResponseCode gets a reference to the given NullableString and assigns it to the GatewayResponseCode field.
func (o *Payment) SetGatewayResponseCode(v string) {
	o.GatewayResponseCode.Set(&v)
}
// SetGatewayResponseCodeNil sets the value for GatewayResponseCode to be an explicit nil
func (o *Payment) SetGatewayResponseCodeNil() {
	o.GatewayResponseCode.Set(nil)
}

// UnsetGatewayResponseCode ensures that no value is present for GatewayResponseCode, not even an explicit nil
func (o *Payment) UnsetGatewayResponseCode() {
	o.GatewayResponseCode.Unset()
}

// GetGatewayResponseText returns the GatewayResponseText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetGatewayResponseText() string {
	if o == nil || IsNil(o.GatewayResponseText.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayResponseText.Get()
}

// GetGatewayResponseTextOk returns a tuple with the GatewayResponseText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetGatewayResponseTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayResponseText.Get(), o.GatewayResponseText.IsSet()
}

// HasGatewayResponseText returns a boolean if a field has been set.
func (o *Payment) HasGatewayResponseText() bool {
	if o != nil && o.GatewayResponseText.IsSet() {
		return true
	}

	return false
}

// SetGatewayResponseText gets a reference to the given NullableString and assigns it to the GatewayResponseText field.
func (o *Payment) SetGatewayResponseText(v string) {
	o.GatewayResponseText.Set(&v)
}
// SetGatewayResponseTextNil sets the value for GatewayResponseText to be an explicit nil
func (o *Payment) SetGatewayResponseTextNil() {
	o.GatewayResponseText.Set(nil)
}

// UnsetGatewayResponseText ensures that no value is present for GatewayResponseText, not even an explicit nil
func (o *Payment) UnsetGatewayResponseText() {
	o.GatewayResponseText.Unset()
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetPaymentMethodType() string {
	if o == nil || IsNil(o.PaymentMethodType.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType.Get()
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodType.Get(), o.PaymentMethodType.IsSet()
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *Payment) HasPaymentMethodType() bool {
	if o != nil && o.PaymentMethodType.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given NullableString and assigns it to the PaymentMethodType field.
func (o *Payment) SetPaymentMethodType(v string) {
	o.PaymentMethodType.Set(&v)
}
// SetPaymentMethodTypeNil sets the value for PaymentMethodType to be an explicit nil
func (o *Payment) SetPaymentMethodTypeNil() {
	o.PaymentMethodType.Set(nil)
}

// UnsetPaymentMethodType ensures that no value is present for PaymentMethodType, not even an explicit nil
func (o *Payment) UnsetPaymentMethodType() {
	o.PaymentMethodType.Unset()
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetCardBrand() string {
	if o == nil || IsNil(o.CardBrand.Get()) {
		var ret string
		return ret
	}
	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// HasCardBrand returns a boolean if a field has been set.
func (o *Payment) HasCardBrand() bool {
	if o != nil && o.CardBrand.IsSet() {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given NullableString and assigns it to the CardBrand field.
func (o *Payment) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}
// SetCardBrandNil sets the value for CardBrand to be an explicit nil
func (o *Payment) SetCardBrandNil() {
	o.CardBrand.Set(nil)
}

// UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
func (o *Payment) UnsetCardBrand() {
	o.CardBrand.Unset()
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetCardLastFour() string {
	if o == nil || IsNil(o.CardLastFour.Get()) {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetCardLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *Payment) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *Payment) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *Payment) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *Payment) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetPaymentChannel returns the PaymentChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetPaymentChannel() string {
	if o == nil || IsNil(o.PaymentChannel.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentChannel.Get()
}

// GetPaymentChannelOk returns a tuple with the PaymentChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetPaymentChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentChannel.Get(), o.PaymentChannel.IsSet()
}

// HasPaymentChannel returns a boolean if a field has been set.
func (o *Payment) HasPaymentChannel() bool {
	if o != nil && o.PaymentChannel.IsSet() {
		return true
	}

	return false
}

// SetPaymentChannel gets a reference to the given NullableString and assigns it to the PaymentChannel field.
func (o *Payment) SetPaymentChannel(v string) {
	o.PaymentChannel.Set(&v)
}
// SetPaymentChannelNil sets the value for PaymentChannel to be an explicit nil
func (o *Payment) SetPaymentChannelNil() {
	o.PaymentChannel.Set(nil)
}

// UnsetPaymentChannel ensures that no value is present for PaymentChannel, not even an explicit nil
func (o *Payment) UnsetPaymentChannel() {
	o.PaymentChannel.Unset()
}

// GetEntryMode returns the EntryMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetEntryMode() string {
	if o == nil || IsNil(o.EntryMode.Get()) {
		var ret string
		return ret
	}
	return *o.EntryMode.Get()
}

// GetEntryModeOk returns a tuple with the EntryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetEntryModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EntryMode.Get(), o.EntryMode.IsSet()
}

// HasEntryMode returns a boolean if a field has been set.
func (o *Payment) HasEntryMode() bool {
	if o != nil && o.EntryMode.IsSet() {
		return true
	}

	return false
}

// SetEntryMode gets a reference to the given NullableString and assigns it to the EntryMode field.
func (o *Payment) SetEntryMode(v string) {
	o.EntryMode.Set(&v)
}
// SetEntryModeNil sets the value for EntryMode to be an explicit nil
func (o *Payment) SetEntryModeNil() {
	o.EntryMode.Set(nil)
}

// UnsetEntryMode ensures that no value is present for EntryMode, not even an explicit nil
func (o *Payment) UnsetEntryMode() {
	o.EntryMode.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Payment) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Payment) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Payment) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Payment) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Payment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Payment) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Payment) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Payment) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Payment) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAuthorizedAt returns the AuthorizedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetAuthorizedAt() time.Time {
	if o == nil || IsNil(o.AuthorizedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AuthorizedAt.Get()
}

// GetAuthorizedAtOk returns a tuple with the AuthorizedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetAuthorizedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedAt.Get(), o.AuthorizedAt.IsSet()
}

// HasAuthorizedAt returns a boolean if a field has been set.
func (o *Payment) HasAuthorizedAt() bool {
	if o != nil && o.AuthorizedAt.IsSet() {
		return true
	}

	return false
}

// SetAuthorizedAt gets a reference to the given NullableTime and assigns it to the AuthorizedAt field.
func (o *Payment) SetAuthorizedAt(v time.Time) {
	o.AuthorizedAt.Set(&v)
}
// SetAuthorizedAtNil sets the value for AuthorizedAt to be an explicit nil
func (o *Payment) SetAuthorizedAtNil() {
	o.AuthorizedAt.Set(nil)
}

// UnsetAuthorizedAt ensures that no value is present for AuthorizedAt, not even an explicit nil
func (o *Payment) UnsetAuthorizedAt() {
	o.AuthorizedAt.Unset()
}

// GetCapturedAt returns the CapturedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetCapturedAt() time.Time {
	if o == nil || IsNil(o.CapturedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CapturedAt.Get()
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetCapturedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CapturedAt.Get(), o.CapturedAt.IsSet()
}

// HasCapturedAt returns a boolean if a field has been set.
func (o *Payment) HasCapturedAt() bool {
	if o != nil && o.CapturedAt.IsSet() {
		return true
	}

	return false
}

// SetCapturedAt gets a reference to the given NullableTime and assigns it to the CapturedAt field.
func (o *Payment) SetCapturedAt(v time.Time) {
	o.CapturedAt.Set(&v)
}
// SetCapturedAtNil sets the value for CapturedAt to be an explicit nil
func (o *Payment) SetCapturedAtNil() {
	o.CapturedAt.Set(nil)
}

// UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
func (o *Payment) UnsetCapturedAt() {
	o.CapturedAt.Unset()
}

// GetVoidedAt returns the VoidedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetVoidedAt() time.Time {
	if o == nil || IsNil(o.VoidedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.VoidedAt.Get()
}

// GetVoidedAtOk returns a tuple with the VoidedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetVoidedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoidedAt.Get(), o.VoidedAt.IsSet()
}

// HasVoidedAt returns a boolean if a field has been set.
func (o *Payment) HasVoidedAt() bool {
	if o != nil && o.VoidedAt.IsSet() {
		return true
	}

	return false
}

// SetVoidedAt gets a reference to the given NullableTime and assigns it to the VoidedAt field.
func (o *Payment) SetVoidedAt(v time.Time) {
	o.VoidedAt.Set(&v)
}
// SetVoidedAtNil sets the value for VoidedAt to be an explicit nil
func (o *Payment) SetVoidedAtNil() {
	o.VoidedAt.Set(nil)
}

// UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil
func (o *Payment) UnsetVoidedAt() {
	o.VoidedAt.Unset()
}

// GetRefundedAt returns the RefundedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Payment) GetRefundedAt() time.Time {
	if o == nil || IsNil(o.RefundedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RefundedAt.Get()
}

// GetRefundedAtOk returns a tuple with the RefundedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payment) GetRefundedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundedAt.Get(), o.RefundedAt.IsSet()
}

// HasRefundedAt returns a boolean if a field has been set.
func (o *Payment) HasRefundedAt() bool {
	if o != nil && o.RefundedAt.IsSet() {
		return true
	}

	return false
}

// SetRefundedAt gets a reference to the given NullableTime and assigns it to the RefundedAt field.
func (o *Payment) SetRefundedAt(v time.Time) {
	o.RefundedAt.Set(&v)
}
// SetRefundedAtNil sets the value for RefundedAt to be an explicit nil
func (o *Payment) SetRefundedAtNil() {
	o.RefundedAt.Set(nil)
}

// UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
func (o *Payment) UnsetRefundedAt() {
	o.RefundedAt.Unset()
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
	if o.AmountCaptured.IsSet() {
		toSerialize["amount_captured"] = o.AmountCaptured.Get()
	}
	if !IsNil(o.AmountRefunded) {
		toSerialize["amount_refunded"] = o.AmountRefunded
	}
	if o.Gateway.IsSet() {
		toSerialize["gateway"] = o.Gateway.Get()
	}
	if o.GatewayTransactionId.IsSet() {
		toSerialize["gateway_transaction_id"] = o.GatewayTransactionId.Get()
	}
	if o.GatewayResponseCode.IsSet() {
		toSerialize["gateway_response_code"] = o.GatewayResponseCode.Get()
	}
	if o.GatewayResponseText.IsSet() {
		toSerialize["gateway_response_text"] = o.GatewayResponseText.Get()
	}
	if o.PaymentMethodType.IsSet() {
		toSerialize["payment_method_type"] = o.PaymentMethodType.Get()
	}
	if o.CardBrand.IsSet() {
		toSerialize["card_brand"] = o.CardBrand.Get()
	}
	if o.CardLastFour.IsSet() {
		toSerialize["card_last_four"] = o.CardLastFour.Get()
	}
	if o.PaymentChannel.IsSet() {
		toSerialize["payment_channel"] = o.PaymentChannel.Get()
	}
	if o.EntryMode.IsSet() {
		toSerialize["entry_mode"] = o.EntryMode.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.AuthorizedAt.IsSet() {
		toSerialize["authorized_at"] = o.AuthorizedAt.Get()
	}
	if o.CapturedAt.IsSet() {
		toSerialize["captured_at"] = o.CapturedAt.Get()
	}
	if o.VoidedAt.IsSet() {
		toSerialize["voided_at"] = o.VoidedAt.Get()
	}
	if o.RefundedAt.IsSet() {
		toSerialize["refunded_at"] = o.RefundedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Payment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"currency",
		"amount",
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

	varPayment := _Payment{}

	err = json.Unmarshal(data, &varPayment)

	if err != nil {
		return err
	}

	*o = Payment(varPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "amount_captured")
		delete(additionalProperties, "amount_refunded")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_transaction_id")
		delete(additionalProperties, "gateway_response_code")
		delete(additionalProperties, "gateway_response_text")
		delete(additionalProperties, "payment_method_type")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "card_last_four")
		delete(additionalProperties, "payment_channel")
		delete(additionalProperties, "entry_mode")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "authorized_at")
		delete(additionalProperties, "captured_at")
		delete(additionalProperties, "voided_at")
		delete(additionalProperties, "refunded_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


