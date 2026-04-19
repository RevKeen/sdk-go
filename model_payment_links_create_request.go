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
)

// checks if the PaymentLinksCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinksCreateRequest{}

// PaymentLinksCreateRequest struct for PaymentLinksCreateRequest
type PaymentLinksCreateRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
	LineItems []PaymentLinksCreateRequestLineItemsInner `json:"line_items,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Mode *string `json:"mode,omitempty"`
	AmountMinor *int32 `json:"amount_minor,omitempty"`
	Currency *string `json:"currency,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// Custom URL to redirect after successful payment. Supports placeholders: {CHECKOUT_ID}, {TRANSACTION_ID}, {PAYMENT_ID}, {CUSTOMER_ID}, {AMOUNT}, {CURRENCY}, {PORTAL_URL}. If not set, redirects to RevKeen's default success page.
	SuccessUrl *string `json:"success_url,omitempty"`
	// Custom URL to redirect if customer cancels the checkout. If not set, no cancel redirect is configured.
	CancelUrl *string `json:"cancel_url,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	CollectShipping *bool `json:"collect_shipping,omitempty"`
	AllowPromotionCodes *bool `json:"allow_promotion_codes,omitempty"`
	TaxBehavior *string `json:"tax_behavior,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinksCreateRequest PaymentLinksCreateRequest

// NewPaymentLinksCreateRequest instantiates a new PaymentLinksCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinksCreateRequest() *PaymentLinksCreateRequest {
	this := PaymentLinksCreateRequest{}
	var mode string = "custom"
	this.Mode = &mode
	var currency string = "USD"
	this.Currency = &currency
	var collectShipping bool = false
	this.CollectShipping = &collectShipping
	var allowPromotionCodes bool = true
	this.AllowPromotionCodes = &allowPromotionCodes
	var taxBehavior string = "unspecified"
	this.TaxBehavior = &taxBehavior
	return &this
}

// NewPaymentLinksCreateRequestWithDefaults instantiates a new PaymentLinksCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksCreateRequestWithDefaults() *PaymentLinksCreateRequest {
	this := PaymentLinksCreateRequest{}
	var mode string = "custom"
	this.Mode = &mode
	var currency string = "USD"
	this.Currency = &currency
	var collectShipping bool = false
	this.CollectShipping = &collectShipping
	var allowPromotionCodes bool = true
	this.AllowPromotionCodes = &allowPromotionCodes
	var taxBehavior string = "unspecified"
	this.TaxBehavior = &taxBehavior
	return &this
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *PaymentLinksCreateRequest) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetLineItems() []PaymentLinksCreateRequestLineItemsInner {
	if o == nil || IsNil(o.LineItems) {
		var ret []PaymentLinksCreateRequestLineItemsInner
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetLineItemsOk() ([]PaymentLinksCreateRequestLineItemsInner, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []PaymentLinksCreateRequestLineItemsInner and assigns it to the LineItems field.
func (o *PaymentLinksCreateRequest) SetLineItems(v []PaymentLinksCreateRequestLineItemsInner) {
	o.LineItems = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentLinksCreateRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *PaymentLinksCreateRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PaymentLinksCreateRequest) SetMode(v string) {
	o.Mode = &v
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given int32 and assigns it to the AmountMinor field.
func (o *PaymentLinksCreateRequest) SetAmountMinor(v int32) {
	o.AmountMinor = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentLinksCreateRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *PaymentLinksCreateRequest) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *PaymentLinksCreateRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *PaymentLinksCreateRequest) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *PaymentLinksCreateRequest) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *PaymentLinksCreateRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetCollectShipping returns the CollectShipping field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetCollectShipping() bool {
	if o == nil || IsNil(o.CollectShipping) {
		var ret bool
		return ret
	}
	return *o.CollectShipping
}

// GetCollectShippingOk returns a tuple with the CollectShipping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetCollectShippingOk() (*bool, bool) {
	if o == nil || IsNil(o.CollectShipping) {
		return nil, false
	}
	return o.CollectShipping, true
}

// HasCollectShipping returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasCollectShipping() bool {
	if o != nil && !IsNil(o.CollectShipping) {
		return true
	}

	return false
}

// SetCollectShipping gets a reference to the given bool and assigns it to the CollectShipping field.
func (o *PaymentLinksCreateRequest) SetCollectShipping(v bool) {
	o.CollectShipping = &v
}

// GetAllowPromotionCodes returns the AllowPromotionCodes field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetAllowPromotionCodes() bool {
	if o == nil || IsNil(o.AllowPromotionCodes) {
		var ret bool
		return ret
	}
	return *o.AllowPromotionCodes
}

// GetAllowPromotionCodesOk returns a tuple with the AllowPromotionCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetAllowPromotionCodesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPromotionCodes) {
		return nil, false
	}
	return o.AllowPromotionCodes, true
}

// HasAllowPromotionCodes returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasAllowPromotionCodes() bool {
	if o != nil && !IsNil(o.AllowPromotionCodes) {
		return true
	}

	return false
}

// SetAllowPromotionCodes gets a reference to the given bool and assigns it to the AllowPromotionCodes field.
func (o *PaymentLinksCreateRequest) SetAllowPromotionCodes(v bool) {
	o.AllowPromotionCodes = &v
}

// GetTaxBehavior returns the TaxBehavior field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetTaxBehavior() string {
	if o == nil || IsNil(o.TaxBehavior) {
		var ret string
		return ret
	}
	return *o.TaxBehavior
}

// GetTaxBehaviorOk returns a tuple with the TaxBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetTaxBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.TaxBehavior) {
		return nil, false
	}
	return o.TaxBehavior, true
}

// HasTaxBehavior returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasTaxBehavior() bool {
	if o != nil && !IsNil(o.TaxBehavior) {
		return true
	}

	return false
}

// SetTaxBehavior gets a reference to the given string and assigns it to the TaxBehavior field.
func (o *PaymentLinksCreateRequest) SetTaxBehavior(v string) {
	o.TaxBehavior = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentLinksCreateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *PaymentLinksCreateRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *PaymentLinksCreateRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

func (o PaymentLinksCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinksCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.AmountMinor) {
		toSerialize["amount_minor"] = o.AmountMinor
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["success_url"] = o.SuccessUrl
	}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirect_url"] = o.RedirectUrl
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
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinksCreateRequest) UnmarshalJSON(data []byte) (err error) {
	varPaymentLinksCreateRequest := _PaymentLinksCreateRequest{}

	err = json.Unmarshal(data, &varPaymentLinksCreateRequest)

	if err != nil {
		return err
	}

	*o = PaymentLinksCreateRequest(varPaymentLinksCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "line_items")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "success_url")
		delete(additionalProperties, "cancel_url")
		delete(additionalProperties, "redirect_url")
		delete(additionalProperties, "collect_shipping")
		delete(additionalProperties, "allow_promotion_codes")
		delete(additionalProperties, "tax_behavior")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "idempotency_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentLinksCreateRequest struct {
	value *PaymentLinksCreateRequest
	isSet bool
}

func (v NullablePaymentLinksCreateRequest) Get() *PaymentLinksCreateRequest {
	return v.value
}

func (v *NullablePaymentLinksCreateRequest) Set(val *PaymentLinksCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinksCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinksCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinksCreateRequest(val *PaymentLinksCreateRequest) *NullablePaymentLinksCreateRequest {
	return &NullablePaymentLinksCreateRequest{value: val, isSet: true}
}

func (v NullablePaymentLinksCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinksCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


