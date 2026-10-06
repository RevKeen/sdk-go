/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your merchant API key in the `x-api-key` header, or as `Authorization: Bearer <key>` — both are accepted. Create and manage keys in the [Dashboard](https://app.revkeen.com/dashboard/settings/general?tab=developer). Use `rk_sandbox_*` for staging/test and `rk_live_*` for production. Keys are scoped, rollable, IP-lockable and revocable; the raw secret is shown once, at creation.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Access tokens are opaque, prefixed `rkoa_`, and sent via `Authorization: Bearer rkoa_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Errors  Every non-2xx response uses one envelope: `{ \"error\": { \"type\", \"code\", \"message\", ... } }`. Branch on `error.type` and `error.code`; treat `error.message` as display text only. Error bodies never echo the credential you presented.  | Status | `error.type` | Notes | |--------|--------------|-------| | `400` | `invalid_request_error` | Schema validation failed. `details.fields` maps each rejected parameter to its reasons. | | `401` | `authentication_error` | Missing, invalid, expired or revoked credential. Always carries `WWW-Authenticate: Bearer realm=\"revkeen-api\"`. | | `402` | `card_error` | Declined by the issuer or processor; see `decline_code`. | | `403` | `authorization_error` | The caller is identified but the action is denied — scope, IP allowlist, origin, or merchant mismatch. | | `409` | `invalid_request_error` or `idempotency_error` | Resource-state conflict, or an `Idempotency-Key` replayed with a different body. | | `429` | `rate_limit_error` | See Rate Limits below. | | `500` | `api_error` | Retry with exponential backoff, quoting `error.request_id`. |  `401` codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`.  `403` codes: `insufficient_permissions`, `forbidden`, `ip_not_allowed`, `tenant_scope_mismatch`, `no_organization`, `no_merchant`, `merchant_mismatch`, `cart_disabled`, `origin_not_allowed`, `verification_required`, `read_only_impersonation_blocked`, `publishable_key_restricted`, `internal_only`.  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  Each API key (`rk_live_*`, `rk_sandbox_*`, `rk_pk_*`) may make **600 requests per minute**, in production and staging alike. The limit is per key, counted over a 60-second window, with no separate burst allowance. Some endpoints apply their own, lower limits.  An over-limit key is still valid: it receives `429`, never `401`. Wait for `Retry-After` and retry; do not re-authenticate or rotate the key.  A rate-limited request returns `429` with `type: rate_limit_error`, `code: rate_limit_exceeded` and `details.retry_after` in whole seconds.  | Header | Meaning | |--------|---------| | `Retry-After` | Whole seconds to wait before retrying. Present on every `429`. | | `X-RateLimit-Limit` | Requests permitted in the current window, when the limiter publishes one. | | `X-RateLimit-Remaining` | Requests left in the current window. | | `X-RateLimit-Reset` | UNIX timestamp in **seconds** at which the window resets. |

API version: 2026-05-01
Contact: support@revkeen.com
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
	ProductIds []string                                  `json:"product_ids,omitempty"`
	LineItems  []PaymentLinksCreateRequestLineItemsInner `json:"line_items,omitempty"`
	Name       *string                                   `json:"name,omitempty"`
	Slug       *string                                   `json:"slug,omitempty"`
	Mode       *string                                   `json:"mode,omitempty"`
	// Deprecated and ignored. A payment link charges the prices of its products or line_items; an amount sent here is never charged.
	// Deprecated
	AmountMinor *int64 `json:"amount_minor,omitempty"`
	// Not used: a payment link takes its currency from its products. For an ad-hoc line item, set line_items[].price_data.currency.
	Currency *string `json:"currency,omitempty"`
	// Deprecated and ignored. This endpoint does not create a payment link for an invoice.
	// Deprecated
	InvoiceId *string `json:"invoice_id,omitempty"`
	// Deprecated and ignored. This endpoint does not attach a payment link to an existing subscription.
	// Deprecated
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// Custom URL to redirect after successful payment. Supports placeholders: {CHECKOUT_ID}, {TRANSACTION_ID}, {PAYMENT_ID}, {CUSTOMER_ID}, {AMOUNT}, {CURRENCY}, {PORTAL_URL}, {CHECKOUT_STATUS_TOKEN}. If not set, redirects to RevKeen's default success page.
	SuccessUrl *string `json:"success_url,omitempty"`
	// Custom URL to redirect if customer cancels the checkout. If not set, no cancel redirect is configured.
	CancelUrl *string `json:"cancel_url,omitempty"`
	// Deprecated and ignored. Use success_url.
	// Deprecated
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// When true, checkout asks the customer for a delivery address before payment.
	CollectShipping *bool `json:"collect_shipping,omitempty"`
	// Not applied yet: links created through this endpoint do not accept discount codes, whatever this is set to.
	AllowPromotionCodes *bool `json:"allow_promotion_codes,omitempty"`
	// Whether the link's prices include tax. inclusive or exclusive is saved on the link; unspecified keeps the default (exclusive). Tax is only calculated when the link's tax mode is manual or auto, which this endpoint does not set.
	TaxBehavior *string                 `json:"tax_behavior,omitempty"`
	Metadata    map[string]*interface{} `json:"metadata,omitempty"`
	ExpiresAt   *time.Time              `json:"expires_at,omitempty"`
	// Customer-facing checkout shape. bundle=all products together (default); single_select=radio picker; multi_select=cart with quantity per line. Companion FKs (default_product_id, highlighted_product_id) are only valid when selection_mode='single_select'.
	SelectionMode *string `json:"selection_mode,omitempty"`
	// For single_select only — option pre-selected when checkout loads. Must reference a product attached to this payment link.
	DefaultProductId NullableString `json:"default_product_id,omitempty"`
	// For single_select only — option shown with a 'Popular' badge in the picker.
	HighlightedProductId NullableString `json:"highlighted_product_id,omitempty"`
	IdempotencyKey       *string        `json:"idempotency_key,omitempty"`
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

// NewPaymentLinksCreateRequestWithDefaults instantiates a new PaymentLinksCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksCreateRequestWithDefaults() *PaymentLinksCreateRequest {
	this := PaymentLinksCreateRequest{}
	var mode string = "custom"
	this.Mode = &mode
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
// Deprecated
func (o *PaymentLinksCreateRequest) GetAmountMinor() int64 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int64
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PaymentLinksCreateRequest) GetAmountMinorOk() (*int64, bool) {
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

// SetAmountMinor gets a reference to the given int64 and assigns it to the AmountMinor field.
// Deprecated
func (o *PaymentLinksCreateRequest) SetAmountMinor(v int64) {
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
// Deprecated
func (o *PaymentLinksCreateRequest) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
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
// Deprecated
func (o *PaymentLinksCreateRequest) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
// Deprecated
func (o *PaymentLinksCreateRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
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
// Deprecated
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
// Deprecated
func (o *PaymentLinksCreateRequest) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
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
// Deprecated
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
func (o *PaymentLinksCreateRequest) GetMetadata() map[string]*interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]*interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetMetadataOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]*interface{}{}, false
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

// SetMetadata gets a reference to the given map[string]*interface{} and assigns it to the Metadata field.
func (o *PaymentLinksCreateRequest) SetMetadata(v map[string]*interface{}) {
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

// GetSelectionMode returns the SelectionMode field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequest) GetSelectionMode() string {
	if o == nil || IsNil(o.SelectionMode) {
		var ret string
		return ret
	}
	return *o.SelectionMode
}

// GetSelectionModeOk returns a tuple with the SelectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequest) GetSelectionModeOk() (*string, bool) {
	if o == nil || IsNil(o.SelectionMode) {
		return nil, false
	}
	return o.SelectionMode, true
}

// HasSelectionMode returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasSelectionMode() bool {
	if o != nil && !IsNil(o.SelectionMode) {
		return true
	}

	return false
}

// SetSelectionMode gets a reference to the given string and assigns it to the SelectionMode field.
func (o *PaymentLinksCreateRequest) SetSelectionMode(v string) {
	o.SelectionMode = &v
}

// GetDefaultProductId returns the DefaultProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksCreateRequest) GetDefaultProductId() string {
	if o == nil || IsNil(o.DefaultProductId.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultProductId.Get()
}

// GetDefaultProductIdOk returns a tuple with the DefaultProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksCreateRequest) GetDefaultProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultProductId.Get(), o.DefaultProductId.IsSet()
}

// HasDefaultProductId returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasDefaultProductId() bool {
	if o != nil && o.DefaultProductId.IsSet() {
		return true
	}

	return false
}

// SetDefaultProductId gets a reference to the given NullableString and assigns it to the DefaultProductId field.
func (o *PaymentLinksCreateRequest) SetDefaultProductId(v string) {
	o.DefaultProductId.Set(&v)
}

// SetDefaultProductIdNil sets the value for DefaultProductId to be an explicit nil
func (o *PaymentLinksCreateRequest) SetDefaultProductIdNil() {
	o.DefaultProductId.Set(nil)
}

// UnsetDefaultProductId ensures that no value is present for DefaultProductId, not even an explicit nil
func (o *PaymentLinksCreateRequest) UnsetDefaultProductId() {
	o.DefaultProductId.Unset()
}

// GetHighlightedProductId returns the HighlightedProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksCreateRequest) GetHighlightedProductId() string {
	if o == nil || IsNil(o.HighlightedProductId.Get()) {
		var ret string
		return ret
	}
	return *o.HighlightedProductId.Get()
}

// GetHighlightedProductIdOk returns a tuple with the HighlightedProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksCreateRequest) GetHighlightedProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighlightedProductId.Get(), o.HighlightedProductId.IsSet()
}

// HasHighlightedProductId returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequest) HasHighlightedProductId() bool {
	if o != nil && o.HighlightedProductId.IsSet() {
		return true
	}

	return false
}

// SetHighlightedProductId gets a reference to the given NullableString and assigns it to the HighlightedProductId field.
func (o *PaymentLinksCreateRequest) SetHighlightedProductId(v string) {
	o.HighlightedProductId.Set(&v)
}

// SetHighlightedProductIdNil sets the value for HighlightedProductId to be an explicit nil
func (o *PaymentLinksCreateRequest) SetHighlightedProductIdNil() {
	o.HighlightedProductId.Set(nil)
}

// UnsetHighlightedProductId ensures that no value is present for HighlightedProductId, not even an explicit nil
func (o *PaymentLinksCreateRequest) UnsetHighlightedProductId() {
	o.HighlightedProductId.Unset()
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
	toSerialize, err := o.ToMap()
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
	if !IsNil(o.SelectionMode) {
		toSerialize["selection_mode"] = o.SelectionMode
	}
	if o.DefaultProductId.IsSet() {
		toSerialize["default_product_id"] = o.DefaultProductId.Get()
	}
	if o.HighlightedProductId.IsSet() {
		toSerialize["highlighted_product_id"] = o.HighlightedProductId.Get()
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
		delete(additionalProperties, "selection_mode")
		delete(additionalProperties, "default_product_id")
		delete(additionalProperties, "highlighted_product_id")
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
