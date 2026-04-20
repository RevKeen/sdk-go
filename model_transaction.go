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
	"time"
	"fmt"
)

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction A financial movement in the unified transaction model. Every money movement — payments, refunds, voids, disputes — is recorded as a transaction linked to an invoice.
type Transaction struct {
	// Unique identifier for the transaction
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Transaction type. sale: original payment. refund: money returned. void: cancel before settlement. capture: capture a previous auth. dispute: chargeback. adjustment: manual correction.
	Type string `json:"type"`
	// Transaction status. pending: processing. succeeded: completed. failed: declined or errored. voided: canceled.
	Status string `json:"status"`
	// Associated invoice ID
	InvoiceId string `json:"invoice_id"`
	// Customer ID
	CustomerId NullableString `json:"customer_id,omitempty"`
	// Parent transaction ID (required for refunds, voids, disputes, adjustments)
	ParentTransactionId NullableString `json:"parent_transaction_id,omitempty"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Transaction amount in minor units (cents)
	Amount int32 `json:"amount"`
	// Amount refunded in minor units
	AmountRefunded *int32 `json:"amount_refunded,omitempty"`
	// Payment gateway (e.g., 'nmi')
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
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Creation timestamp (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp (ISO 8601)
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Transaction Transaction

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction(id string, object string, type_ string, status string, invoiceId string, currency string, amount int32, createdAt time.Time, updatedAt time.Time) *Transaction {
	this := Transaction{}
	this.Id = id
	this.Object = object
	this.Type = type_
	this.Status = status
	this.InvoiceId = invoiceId
	this.Currency = currency
	this.Amount = amount
	var amountRefunded int32 = 0
	this.AmountRefunded = &amountRefunded
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	var amountRefunded int32 = 0
	this.AmountRefunded = &amountRefunded
	return &this
}

// GetId returns the Id field value
func (o *Transaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transaction) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Transaction) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Transaction) SetObject(v string) {
	o.Object = v
}

// GetType returns the Type field value
func (o *Transaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Transaction) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *Transaction) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transaction) SetStatus(v string) {
	o.Status = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *Transaction) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *Transaction) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Transaction) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Transaction) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Transaction) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Transaction) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetParentTransactionId returns the ParentTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetParentTransactionId() string {
	if o == nil || IsNil(o.ParentTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.ParentTransactionId.Get()
}

// GetParentTransactionIdOk returns a tuple with the ParentTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetParentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentTransactionId.Get(), o.ParentTransactionId.IsSet()
}

// HasParentTransactionId returns a boolean if a field has been set.
func (o *Transaction) HasParentTransactionId() bool {
	if o != nil && o.ParentTransactionId.IsSet() {
		return true
	}

	return false
}

// SetParentTransactionId gets a reference to the given NullableString and assigns it to the ParentTransactionId field.
func (o *Transaction) SetParentTransactionId(v string) {
	o.ParentTransactionId.Set(&v)
}
// SetParentTransactionIdNil sets the value for ParentTransactionId to be an explicit nil
func (o *Transaction) SetParentTransactionIdNil() {
	o.ParentTransactionId.Set(nil)
}

// UnsetParentTransactionId ensures that no value is present for ParentTransactionId, not even an explicit nil
func (o *Transaction) UnsetParentTransactionId() {
	o.ParentTransactionId.Unset()
}

// GetCurrency returns the Currency field value
func (o *Transaction) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Transaction) SetCurrency(v string) {
	o.Currency = v
}

// GetAmount returns the Amount field value
func (o *Transaction) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *Transaction) SetAmount(v int32) {
	o.Amount = v
}

// GetAmountRefunded returns the AmountRefunded field value if set, zero value otherwise.
func (o *Transaction) GetAmountRefunded() int32 {
	if o == nil || IsNil(o.AmountRefunded) {
		var ret int32
		return ret
	}
	return *o.AmountRefunded
}

// GetAmountRefundedOk returns a tuple with the AmountRefunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetAmountRefundedOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountRefunded) {
		return nil, false
	}
	return o.AmountRefunded, true
}

// HasAmountRefunded returns a boolean if a field has been set.
func (o *Transaction) HasAmountRefunded() bool {
	if o != nil && !IsNil(o.AmountRefunded) {
		return true
	}

	return false
}

// SetAmountRefunded gets a reference to the given int32 and assigns it to the AmountRefunded field.
func (o *Transaction) SetAmountRefunded(v int32) {
	o.AmountRefunded = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetGateway() string {
	if o == nil || IsNil(o.Gateway.Get()) {
		var ret string
		return ret
	}
	return *o.Gateway.Get()
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gateway.Get(), o.Gateway.IsSet()
}

// HasGateway returns a boolean if a field has been set.
func (o *Transaction) HasGateway() bool {
	if o != nil && o.Gateway.IsSet() {
		return true
	}

	return false
}

// SetGateway gets a reference to the given NullableString and assigns it to the Gateway field.
func (o *Transaction) SetGateway(v string) {
	o.Gateway.Set(&v)
}
// SetGatewayNil sets the value for Gateway to be an explicit nil
func (o *Transaction) SetGatewayNil() {
	o.Gateway.Set(nil)
}

// UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
func (o *Transaction) UnsetGateway() {
	o.Gateway.Unset()
}

// GetGatewayTransactionId returns the GatewayTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetGatewayTransactionId() string {
	if o == nil || IsNil(o.GatewayTransactionId.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayTransactionId.Get()
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTransactionId.Get(), o.GatewayTransactionId.IsSet()
}

// HasGatewayTransactionId returns a boolean if a field has been set.
func (o *Transaction) HasGatewayTransactionId() bool {
	if o != nil && o.GatewayTransactionId.IsSet() {
		return true
	}

	return false
}

// SetGatewayTransactionId gets a reference to the given NullableString and assigns it to the GatewayTransactionId field.
func (o *Transaction) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId.Set(&v)
}
// SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil
func (o *Transaction) SetGatewayTransactionIdNil() {
	o.GatewayTransactionId.Set(nil)
}

// UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
func (o *Transaction) UnsetGatewayTransactionId() {
	o.GatewayTransactionId.Unset()
}

// GetGatewayResponseCode returns the GatewayResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetGatewayResponseCode() string {
	if o == nil || IsNil(o.GatewayResponseCode.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayResponseCode.Get()
}

// GetGatewayResponseCodeOk returns a tuple with the GatewayResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetGatewayResponseCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayResponseCode.Get(), o.GatewayResponseCode.IsSet()
}

// HasGatewayResponseCode returns a boolean if a field has been set.
func (o *Transaction) HasGatewayResponseCode() bool {
	if o != nil && o.GatewayResponseCode.IsSet() {
		return true
	}

	return false
}

// SetGatewayResponseCode gets a reference to the given NullableString and assigns it to the GatewayResponseCode field.
func (o *Transaction) SetGatewayResponseCode(v string) {
	o.GatewayResponseCode.Set(&v)
}
// SetGatewayResponseCodeNil sets the value for GatewayResponseCode to be an explicit nil
func (o *Transaction) SetGatewayResponseCodeNil() {
	o.GatewayResponseCode.Set(nil)
}

// UnsetGatewayResponseCode ensures that no value is present for GatewayResponseCode, not even an explicit nil
func (o *Transaction) UnsetGatewayResponseCode() {
	o.GatewayResponseCode.Unset()
}

// GetGatewayResponseText returns the GatewayResponseText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetGatewayResponseText() string {
	if o == nil || IsNil(o.GatewayResponseText.Get()) {
		var ret string
		return ret
	}
	return *o.GatewayResponseText.Get()
}

// GetGatewayResponseTextOk returns a tuple with the GatewayResponseText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetGatewayResponseTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayResponseText.Get(), o.GatewayResponseText.IsSet()
}

// HasGatewayResponseText returns a boolean if a field has been set.
func (o *Transaction) HasGatewayResponseText() bool {
	if o != nil && o.GatewayResponseText.IsSet() {
		return true
	}

	return false
}

// SetGatewayResponseText gets a reference to the given NullableString and assigns it to the GatewayResponseText field.
func (o *Transaction) SetGatewayResponseText(v string) {
	o.GatewayResponseText.Set(&v)
}
// SetGatewayResponseTextNil sets the value for GatewayResponseText to be an explicit nil
func (o *Transaction) SetGatewayResponseTextNil() {
	o.GatewayResponseText.Set(nil)
}

// UnsetGatewayResponseText ensures that no value is present for GatewayResponseText, not even an explicit nil
func (o *Transaction) UnsetGatewayResponseText() {
	o.GatewayResponseText.Unset()
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetPaymentMethodType() string {
	if o == nil || IsNil(o.PaymentMethodType.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType.Get()
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodType.Get(), o.PaymentMethodType.IsSet()
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *Transaction) HasPaymentMethodType() bool {
	if o != nil && o.PaymentMethodType.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given NullableString and assigns it to the PaymentMethodType field.
func (o *Transaction) SetPaymentMethodType(v string) {
	o.PaymentMethodType.Set(&v)
}
// SetPaymentMethodTypeNil sets the value for PaymentMethodType to be an explicit nil
func (o *Transaction) SetPaymentMethodTypeNil() {
	o.PaymentMethodType.Set(nil)
}

// UnsetPaymentMethodType ensures that no value is present for PaymentMethodType, not even an explicit nil
func (o *Transaction) UnsetPaymentMethodType() {
	o.PaymentMethodType.Unset()
}

// GetCardBrand returns the CardBrand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetCardBrand() string {
	if o == nil || IsNil(o.CardBrand.Get()) {
		var ret string
		return ret
	}
	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// HasCardBrand returns a boolean if a field has been set.
func (o *Transaction) HasCardBrand() bool {
	if o != nil && o.CardBrand.IsSet() {
		return true
	}

	return false
}

// SetCardBrand gets a reference to the given NullableString and assigns it to the CardBrand field.
func (o *Transaction) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}
// SetCardBrandNil sets the value for CardBrand to be an explicit nil
func (o *Transaction) SetCardBrandNil() {
	o.CardBrand.Set(nil)
}

// UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
func (o *Transaction) UnsetCardBrand() {
	o.CardBrand.Unset()
}

// GetCardLastFour returns the CardLastFour field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Transaction) GetCardLastFour() string {
	if o == nil || IsNil(o.CardLastFour.Get()) {
		var ret string
		return ret
	}
	return *o.CardLastFour.Get()
}

// GetCardLastFourOk returns a tuple with the CardLastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transaction) GetCardLastFourOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLastFour.Get(), o.CardLastFour.IsSet()
}

// HasCardLastFour returns a boolean if a field has been set.
func (o *Transaction) HasCardLastFour() bool {
	if o != nil && o.CardLastFour.IsSet() {
		return true
	}

	return false
}

// SetCardLastFour gets a reference to the given NullableString and assigns it to the CardLastFour field.
func (o *Transaction) SetCardLastFour(v string) {
	o.CardLastFour.Set(&v)
}
// SetCardLastFourNil sets the value for CardLastFour to be an explicit nil
func (o *Transaction) SetCardLastFourNil() {
	o.CardLastFour.Set(nil)
}

// UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
func (o *Transaction) UnsetCardLastFour() {
	o.CardLastFour.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Transaction) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Transaction) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Transaction) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Transaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Transaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Transaction) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Transaction) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["invoice_id"] = o.InvoiceId
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	if o.ParentTransactionId.IsSet() {
		toSerialize["parent_transaction_id"] = o.ParentTransactionId.Get()
	}
	toSerialize["currency"] = o.Currency
	toSerialize["amount"] = o.Amount
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
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Transaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"type",
		"status",
		"invoice_id",
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

	varTransaction := _Transaction{}

	err = json.Unmarshal(data, &varTransaction)

	if err != nil {
		return err
	}

	*o = Transaction(varTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "parent_transaction_id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "amount_refunded")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_transaction_id")
		delete(additionalProperties, "gateway_response_code")
		delete(additionalProperties, "gateway_response_text")
		delete(additionalProperties, "payment_method_type")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "card_last_four")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


