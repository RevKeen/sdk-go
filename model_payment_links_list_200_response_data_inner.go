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

// checks if the PaymentLinksList200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinksList200ResponseDataInner{}

// PaymentLinksList200ResponseDataInner struct for PaymentLinksList200ResponseDataInner
type PaymentLinksList200ResponseDataInner struct {
	Id string `json:"id"`
	MerchantId string `json:"merchant_id"`
	PublicId string `json:"public_id"`
	Name *string `json:"name,omitempty"`
	Slug NullableString `json:"slug,omitempty"`
	Mode string `json:"mode"`
	Status string `json:"status"`
	AmountMinor NullableInt32 `json:"amount_minor,omitempty"`
	Currency *string `json:"currency,omitempty"`
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	SubscriptionId NullableString `json:"subscription_id,omitempty"`
	ProductIds []string `json:"product_ids,omitempty"`
	LineItems []PaymentLinksList200ResponseDataInnerLineItemsInner `json:"line_items,omitempty"`
	// Custom URL to redirect after successful payment. Supports placeholders: {CHECKOUT_ID}, {TRANSACTION_ID}, {PAYMENT_ID}, {CUSTOMER_ID}, {AMOUNT}, {CURRENCY}, {PORTAL_URL}. If not set, redirects to RevKeen's default success page.
	SuccessUrl NullableString `json:"success_url,omitempty"`
	// Custom URL to redirect if customer cancels the checkout. If not set, no cancel redirect is configured.
	CancelUrl NullableString `json:"cancel_url,omitempty"`
	RedirectUrl NullableString `json:"redirect_url,omitempty"`
	CollectShipping *bool `json:"collect_shipping,omitempty"`
	AllowPromotionCodes *bool `json:"allow_promotion_codes,omitempty"`
	TaxBehavior *string `json:"tax_behavior,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// Customer-facing checkout shape. bundle=all products together; single_select=radio picker; multi_select=cart with quantity per line.
	SelectionMode *string `json:"selection_mode,omitempty"`
	// For single_select only — option pre-selected when checkout loads. Must reference a product attached to this payment link.
	DefaultProductId NullableString `json:"default_product_id,omitempty"`
	// For single_select only — option shown with a 'Popular' badge.
	HighlightedProductId NullableString `json:"highlighted_product_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinksList200ResponseDataInner PaymentLinksList200ResponseDataInner

// NewPaymentLinksList200ResponseDataInner instantiates a new PaymentLinksList200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinksList200ResponseDataInner(id string, merchantId string, publicId string, mode string, status string, createdAt time.Time, updatedAt time.Time) *PaymentLinksList200ResponseDataInner {
	this := PaymentLinksList200ResponseDataInner{}
	this.Id = id
	this.MerchantId = merchantId
	this.PublicId = publicId
	this.Mode = mode
	this.Status = status
	var currency string = "USD"
	this.Currency = &currency
	var collectShipping bool = false
	this.CollectShipping = &collectShipping
	var allowPromotionCodes bool = true
	this.AllowPromotionCodes = &allowPromotionCodes
	var taxBehavior string = "unspecified"
	this.TaxBehavior = &taxBehavior
	var selectionMode string = "bundle"
	this.SelectionMode = &selectionMode
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPaymentLinksList200ResponseDataInnerWithDefaults instantiates a new PaymentLinksList200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksList200ResponseDataInnerWithDefaults() *PaymentLinksList200ResponseDataInner {
	this := PaymentLinksList200ResponseDataInner{}
	var currency string = "USD"
	this.Currency = &currency
	var collectShipping bool = false
	this.CollectShipping = &collectShipping
	var allowPromotionCodes bool = true
	this.AllowPromotionCodes = &allowPromotionCodes
	var taxBehavior string = "unspecified"
	this.TaxBehavior = &taxBehavior
	var selectionMode string = "bundle"
	this.SelectionMode = &selectionMode
	return &this
}

// GetId returns the Id field value
func (o *PaymentLinksList200ResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentLinksList200ResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *PaymentLinksList200ResponseDataInner) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *PaymentLinksList200ResponseDataInner) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetPublicId returns the PublicId field value
func (o *PaymentLinksList200ResponseDataInner) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *PaymentLinksList200ResponseDataInner) SetPublicId(v string) {
	o.PublicId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentLinksList200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *PaymentLinksList200ResponseDataInner) SetSlug(v string) {
	o.Slug.Set(&v)
}
// SetSlugNil sets the value for Slug to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetSlug() {
	o.Slug.Unset()
}

// GetMode returns the Mode field value
func (o *PaymentLinksList200ResponseDataInner) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *PaymentLinksList200ResponseDataInner) SetMode(v string) {
	o.Mode = v
}

// GetStatus returns the Status field value
func (o *PaymentLinksList200ResponseDataInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentLinksList200ResponseDataInner) SetStatus(v string) {
	o.Status = v
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountMinor.Get()
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountMinor.Get(), o.AmountMinor.IsSet()
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasAmountMinor() bool {
	if o != nil && o.AmountMinor.IsSet() {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given NullableInt32 and assigns it to the AmountMinor field.
func (o *PaymentLinksList200ResponseDataInner) SetAmountMinor(v int32) {
	o.AmountMinor.Set(&v)
}
// SetAmountMinorNil sets the value for AmountMinor to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetAmountMinorNil() {
	o.AmountMinor.Set(nil)
}

// UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetAmountMinor() {
	o.AmountMinor.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentLinksList200ResponseDataInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *PaymentLinksList200ResponseDataInner) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId.Get()) {
		var ret string
		return ret
	}
	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId.IsSet() {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given NullableString and assigns it to the SubscriptionId field.
func (o *PaymentLinksList200ResponseDataInner) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}
// SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetSubscriptionIdNil() {
	o.SubscriptionId.Set(nil)
}

// UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetSubscriptionId() {
	o.SubscriptionId.Unset()
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *PaymentLinksList200ResponseDataInner) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetLineItems() []PaymentLinksList200ResponseDataInnerLineItemsInner {
	if o == nil || IsNil(o.LineItems) {
		var ret []PaymentLinksList200ResponseDataInnerLineItemsInner
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetLineItemsOk() ([]PaymentLinksList200ResponseDataInnerLineItemsInner, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []PaymentLinksList200ResponseDataInnerLineItemsInner and assigns it to the LineItems field.
func (o *PaymentLinksList200ResponseDataInner) SetLineItems(v []PaymentLinksList200ResponseDataInnerLineItemsInner) {
	o.LineItems = v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SuccessUrl.Get()
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetSuccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessUrl.Get(), o.SuccessUrl.IsSet()
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasSuccessUrl() bool {
	if o != nil && o.SuccessUrl.IsSet() {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given NullableString and assigns it to the SuccessUrl field.
func (o *PaymentLinksList200ResponseDataInner) SetSuccessUrl(v string) {
	o.SuccessUrl.Set(&v)
}
// SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetSuccessUrlNil() {
	o.SuccessUrl.Set(nil)
}

// UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetSuccessUrl() {
	o.SuccessUrl.Unset()
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CancelUrl.Get()
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetCancelUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelUrl.Get(), o.CancelUrl.IsSet()
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasCancelUrl() bool {
	if o != nil && o.CancelUrl.IsSet() {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given NullableString and assigns it to the CancelUrl field.
func (o *PaymentLinksList200ResponseDataInner) SetCancelUrl(v string) {
	o.CancelUrl.Set(&v)
}
// SetCancelUrlNil sets the value for CancelUrl to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetCancelUrlNil() {
	o.CancelUrl.Set(nil)
}

// UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetCancelUrl() {
	o.CancelUrl.Unset()
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RedirectUrl.Get()
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUrl.Get(), o.RedirectUrl.IsSet()
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl.IsSet() {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given NullableString and assigns it to the RedirectUrl field.
func (o *PaymentLinksList200ResponseDataInner) SetRedirectUrl(v string) {
	o.RedirectUrl.Set(&v)
}
// SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetRedirectUrlNil() {
	o.RedirectUrl.Set(nil)
}

// UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetRedirectUrl() {
	o.RedirectUrl.Unset()
}

// GetCollectShipping returns the CollectShipping field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetCollectShipping() bool {
	if o == nil || IsNil(o.CollectShipping) {
		var ret bool
		return ret
	}
	return *o.CollectShipping
}

// GetCollectShippingOk returns a tuple with the CollectShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetCollectShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectShipping) {
		return nil, false
	}
	return o.CollectShipping, true
}

// HasCollectShipping returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasCollectShipping() bool {
	if o != nil && !IsNil(o.CollectShipping) {
		return true
	}

	return false
}

// SetCollectShipping gets a reference to the given bool and assigns it to the CollectShipping field.
func (o *PaymentLinksList200ResponseDataInner) SetCollectShipping(v bool) {
	o.CollectShipping = &v
}

// GetAllowPromotionCodes returns the AllowPromotionCodes field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetAllowPromotionCodes() bool {
	if o == nil || IsNil(o.AllowPromotionCodes) {
		var ret bool
		return ret
	}
	return *o.AllowPromotionCodes
}

// GetAllowPromotionCodesOk returns a tuple with the AllowPromotionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetAllowPromotionCodesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPromotionCodes) {
		return nil, false
	}
	return o.AllowPromotionCodes, true
}

// HasAllowPromotionCodes returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasAllowPromotionCodes() bool {
	if o != nil && !IsNil(o.AllowPromotionCodes) {
		return true
	}

	return false
}

// SetAllowPromotionCodes gets a reference to the given bool and assigns it to the AllowPromotionCodes field.
func (o *PaymentLinksList200ResponseDataInner) SetAllowPromotionCodes(v bool) {
	o.AllowPromotionCodes = &v
}

// GetTaxBehavior returns the TaxBehavior field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetTaxBehavior() string {
	if o == nil || IsNil(o.TaxBehavior) {
		var ret string
		return ret
	}
	return *o.TaxBehavior
}

// GetTaxBehaviorOk returns a tuple with the TaxBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetTaxBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.TaxBehavior) {
		return nil, false
	}
	return o.TaxBehavior, true
}

// HasTaxBehavior returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasTaxBehavior() bool {
	if o != nil && !IsNil(o.TaxBehavior) {
		return true
	}

	return false
}

// SetTaxBehavior gets a reference to the given string and assigns it to the TaxBehavior field.
func (o *PaymentLinksList200ResponseDataInner) SetTaxBehavior(v string) {
	o.TaxBehavior = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentLinksList200ResponseDataInner) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *PaymentLinksList200ResponseDataInner) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetSelectionMode returns the SelectionMode field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInner) GetSelectionMode() string {
	if o == nil || IsNil(o.SelectionMode) {
		var ret string
		return ret
	}
	return *o.SelectionMode
}

// GetSelectionModeOk returns a tuple with the SelectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetSelectionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelectionMode) {
		return nil, false
	}
	return o.SelectionMode, true
}

// HasSelectionMode returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasSelectionMode() bool {
	if o != nil && !IsNil(o.SelectionMode) {
		return true
	}

	return false
}

// SetSelectionMode gets a reference to the given string and assigns it to the SelectionMode field.
func (o *PaymentLinksList200ResponseDataInner) SetSelectionMode(v string) {
	o.SelectionMode = &v
}

// GetDefaultProductId returns the DefaultProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetDefaultProductId() string {
	if o == nil || IsNil(o.DefaultProductId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultProductId.Get()
}

// GetDefaultProductIdOk returns a tuple with the DefaultProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetDefaultProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultProductId.Get(), o.DefaultProductId.IsSet()
}

// HasDefaultProductId returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasDefaultProductId() bool {
	if o != nil && o.DefaultProductId.IsSet() {
		return true
	}

	return false
}

// SetDefaultProductId gets a reference to the given NullableString and assigns it to the DefaultProductId field.
func (o *PaymentLinksList200ResponseDataInner) SetDefaultProductId(v string) {
	o.DefaultProductId.Set(&v)
}
// SetDefaultProductIdNil sets the value for DefaultProductId to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetDefaultProductIdNil() {
	o.DefaultProductId.Set(nil)
}

// UnsetDefaultProductId ensures that no value is present for DefaultProductId, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetDefaultProductId() {
	o.DefaultProductId.Unset()
}

// GetHighlightedProductId returns the HighlightedProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInner) GetHighlightedProductId() string {
	if o == nil || IsNil(o.HighlightedProductId.Get()) {
		var ret string
		return ret
	}
	return *o.HighlightedProductId.Get()
}

// GetHighlightedProductIdOk returns a tuple with the HighlightedProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInner) GetHighlightedProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighlightedProductId.Get(), o.HighlightedProductId.IsSet()
}

// HasHighlightedProductId returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInner) HasHighlightedProductId() bool {
	if o != nil && o.HighlightedProductId.IsSet() {
		return true
	}

	return false
}

// SetHighlightedProductId gets a reference to the given NullableString and assigns it to the HighlightedProductId field.
func (o *PaymentLinksList200ResponseDataInner) SetHighlightedProductId(v string) {
	o.HighlightedProductId.Set(&v)
}
// SetHighlightedProductIdNil sets the value for HighlightedProductId to be an explicit nil
func (o *PaymentLinksList200ResponseDataInner) SetHighlightedProductIdNil() {
	o.HighlightedProductId.Set(nil)
}

// UnsetHighlightedProductId ensures that no value is present for HighlightedProductId, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInner) UnsetHighlightedProductId() {
	o.HighlightedProductId.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentLinksList200ResponseDataInner) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentLinksList200ResponseDataInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaymentLinksList200ResponseDataInner) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaymentLinksList200ResponseDataInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PaymentLinksList200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinksList200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["public_id"] = o.PublicId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	toSerialize["mode"] = o.Mode
	toSerialize["status"] = o.Status
	if o.AmountMinor.IsSet() {
		toSerialize["amount_minor"] = o.AmountMinor.Get()
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
	}
	if o.SubscriptionId.IsSet() {
		toSerialize["subscription_id"] = o.SubscriptionId.Get()
	}
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if o.SuccessUrl.IsSet() {
		toSerialize["success_url"] = o.SuccessUrl.Get()
	}
	if o.CancelUrl.IsSet() {
		toSerialize["cancel_url"] = o.CancelUrl.Get()
	}
	if o.RedirectUrl.IsSet() {
		toSerialize["redirect_url"] = o.RedirectUrl.Get()
	}
	if !IsNil(o.CollectShipping) {
		toSerialize["collect_shipping"] = o.CollectShipping
	}
	if !IsNil(o.AllowPromotionCodes) {
		toSerialize["allow_promotion_codes"] = o.AllowPromotionCodes
	}
	if !IsNil(o.TaxBehavior) {
		toSerialize["tax_behavior"] = o.TaxBehavior
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if !IsNil(o.SelectionMode) {
		toSerialize["selection_mode"] = o.SelectionMode
	}
	if o.DefaultProductId.IsSet() {
		toSerialize["default_product_id"] = o.DefaultProductId.Get()
	}
	if o.HighlightedProductId.IsSet() {
		toSerialize["highlighted_product_id"] = o.HighlightedProductId.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinksList200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"public_id",
		"mode",
		"status",
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

	varPaymentLinksList200ResponseDataInner := _PaymentLinksList200ResponseDataInner{}

	err = json.Unmarshal(data, &varPaymentLinksList200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = PaymentLinksList200ResponseDataInner(varPaymentLinksList200ResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "line_items")
		delete(additionalProperties, "success_url")
		delete(additionalProperties, "cancel_url")
		delete(additionalProperties, "redirect_url")
		delete(additionalProperties, "collect_shipping")
		delete(additionalProperties, "allow_promotion_codes")
		delete(additionalProperties, "tax_behavior")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "selection_mode")
		delete(additionalProperties, "default_product_id")
		delete(additionalProperties, "highlighted_product_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentLinksList200ResponseDataInner struct {
	value *PaymentLinksList200ResponseDataInner
	isSet bool
}

func (v NullablePaymentLinksList200ResponseDataInner) Get() *PaymentLinksList200ResponseDataInner {
	return v.value
}

func (v *NullablePaymentLinksList200ResponseDataInner) Set(val *PaymentLinksList200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinksList200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinksList200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinksList200ResponseDataInner(val *PaymentLinksList200ResponseDataInner) *NullablePaymentLinksList200ResponseDataInner {
	return &NullablePaymentLinksList200ResponseDataInner{value: val, isSet: true}
}

func (v NullablePaymentLinksList200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinksList200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


