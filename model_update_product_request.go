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

// checks if the UpdateProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductRequest{}

// UpdateProductRequest Parameters for updating an existing product. Only the provided fields are changed.
type UpdateProductRequest struct {
	Name          *string        `json:"name,omitempty"`
	Description   NullableString `json:"description,omitempty"`
	AmountMinor   *int64         `json:"amount_minor,omitempty"`
	Currency      *string        `json:"currency,omitempty"`
	Interval      NullableString `json:"interval,omitempty"`
	IntervalCount NullableInt32  `json:"interval_count,omitempty"`
	TrialDays     *int32         `json:"trial_days,omitempty"`
	IsActive      *bool          `json:"is_active,omitempty"`
	IsArchived    *bool          `json:"is_archived,omitempty"`
	// URL handle, unique among this merchant's products. Pass null to clear it.
	Slug NullableString `json:"slug,omitempty"`
	// Ordered extra images (at most 20). Replaces the list. The main image stays in image_url.
	Images []ProductImage `json:"images,omitempty"`
	// Search and social title (at most 70 characters). Defaults to the name.
	SeoTitle NullableString `json:"seo_title,omitempty"`
	// Meta description (at most 320 characters).
	SeoDescription NullableString `json:"seo_description,omitempty"`
	// Whether the product is listed on the Storefront. hidden: only payment links and invoices. scheduled: listed from store_visible_from.
	StoreVisibility *string `json:"store_visibility,omitempty"`
	// When a scheduled product appears on the Storefront (ISO 8601). Required when store_visibility is scheduled.
	StoreVisibleFrom NullableTime `json:"store_visible_from,omitempty"`
	// Fulfillment type for the product.
	FulfillmentType *string `json:"fulfillment_type,omitempty"`
	// Turn stock tracking on (requires stock_quantity) or off. One-time products only. Stock is stored in the product's availability record.
	TrackInventory *bool `json:"track_inventory,omitempty"`
	// Deprecated. Sets the units available to sell in the product's availability record. Required (non-negative) when track_inventory is true.
	// Deprecated
	StockQuantity NullableInt32 `json:"stock_quantity,omitempty"`
	// Tags for Storefront filters and search. Replaces the product's tags. At most 50, each up to 50 characters.
	Tags []string `json:"tags,omitempty"`
	// How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month.
	BillingAnchorRule *string       `json:"billing_anchor_rule,omitempty"`
	BillingAnchorDay  NullableInt32 `json:"billing_anchor_day,omitempty"`
	// When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date.
	FirstChargeBehavior *string `json:"first_charge_behavior,omitempty"`
	// How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles.
	EndBehavior *string       `json:"end_behavior,omitempty"`
	MaxPayments NullableInt32 `json:"max_payments,omitempty"`
	// Arbitrary key-value metadata for the product
	Metadata map[string]*interface{} `json:"metadata,omitempty"`
	// Product image S3 key or public URL. Pass null to remove.
	ImageUrl             NullableString `json:"image_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateProductRequest UpdateProductRequest

// NewUpdateProductRequest instantiates a new UpdateProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductRequest() *UpdateProductRequest {
	this := UpdateProductRequest{}
	return &this
}

// NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductRequestWithDefaults() *UpdateProductRequest {
	this := UpdateProductRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateProductRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateProductRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateProductRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateProductRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetAmountMinor() int64 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int64
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetAmountMinorOk() (*int64, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given int64 and assigns it to the AmountMinor field.
func (o *UpdateProductRequest) SetAmountMinor(v int64) {
	o.AmountMinor = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UpdateProductRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetInterval() string {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret string
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableString and assigns it to the Interval field.
func (o *UpdateProductRequest) SetInterval(v string) {
	o.Interval.Set(&v)
}

// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *UpdateProductRequest) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *UpdateProductRequest) UnsetInterval() {
	o.Interval.Unset()
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasIntervalCount() bool {
	if o != nil && o.IntervalCount.IsSet() {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given NullableInt32 and assigns it to the IntervalCount field.
func (o *UpdateProductRequest) SetIntervalCount(v int32) {
	o.IntervalCount.Set(&v)
}

// SetIntervalCountNil sets the value for IntervalCount to be an explicit nil
func (o *UpdateProductRequest) SetIntervalCountNil() {
	o.IntervalCount.Set(nil)
}

// UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
func (o *UpdateProductRequest) UnsetIntervalCount() {
	o.IntervalCount.Unset()
}

// GetTrialDays returns the TrialDays field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetTrialDays() int32 {
	if o == nil || IsNil(o.TrialDays) {
		var ret int32
		return ret
	}
	return *o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetTrialDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialDays) {
		return nil, false
	}
	return o.TrialDays, true
}

// HasTrialDays returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasTrialDays() bool {
	if o != nil && !IsNil(o.TrialDays) {
		return true
	}

	return false
}

// SetTrialDays gets a reference to the given int32 and assigns it to the TrialDays field.
func (o *UpdateProductRequest) SetTrialDays(v int32) {
	o.TrialDays = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateProductRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetIsArchived() bool {
	if o == nil || IsNil(o.IsArchived) {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetIsArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *UpdateProductRequest) SetIsArchived(v bool) {
	o.IsArchived = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *UpdateProductRequest) SetSlug(v string) {
	o.Slug.Set(&v)
}

// SetSlugNil sets the value for Slug to be an explicit nil
func (o *UpdateProductRequest) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *UpdateProductRequest) UnsetSlug() {
	o.Slug.Unset()
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetImages() []ProductImage {
	if o == nil || IsNil(o.Images) {
		var ret []ProductImage
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetImagesOk() ([]ProductImage, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ProductImage and assigns it to the Images field.
func (o *UpdateProductRequest) SetImages(v []ProductImage) {
	o.Images = v
}

// GetSeoTitle returns the SeoTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetSeoTitle() string {
	if o == nil || IsNil(o.SeoTitle.Get()) {
		var ret string
		return ret
	}
	return *o.SeoTitle.Get()
}

// GetSeoTitleOk returns a tuple with the SeoTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetSeoTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeoTitle.Get(), o.SeoTitle.IsSet()
}

// HasSeoTitle returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasSeoTitle() bool {
	if o != nil && o.SeoTitle.IsSet() {
		return true
	}

	return false
}

// SetSeoTitle gets a reference to the given NullableString and assigns it to the SeoTitle field.
func (o *UpdateProductRequest) SetSeoTitle(v string) {
	o.SeoTitle.Set(&v)
}

// SetSeoTitleNil sets the value for SeoTitle to be an explicit nil
func (o *UpdateProductRequest) SetSeoTitleNil() {
	o.SeoTitle.Set(nil)
}

// UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
func (o *UpdateProductRequest) UnsetSeoTitle() {
	o.SeoTitle.Unset()
}

// GetSeoDescription returns the SeoDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetSeoDescription() string {
	if o == nil || IsNil(o.SeoDescription.Get()) {
		var ret string
		return ret
	}
	return *o.SeoDescription.Get()
}

// GetSeoDescriptionOk returns a tuple with the SeoDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetSeoDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeoDescription.Get(), o.SeoDescription.IsSet()
}

// HasSeoDescription returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasSeoDescription() bool {
	if o != nil && o.SeoDescription.IsSet() {
		return true
	}

	return false
}

// SetSeoDescription gets a reference to the given NullableString and assigns it to the SeoDescription field.
func (o *UpdateProductRequest) SetSeoDescription(v string) {
	o.SeoDescription.Set(&v)
}

// SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil
func (o *UpdateProductRequest) SetSeoDescriptionNil() {
	o.SeoDescription.Set(nil)
}

// UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
func (o *UpdateProductRequest) UnsetSeoDescription() {
	o.SeoDescription.Unset()
}

// GetStoreVisibility returns the StoreVisibility field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetStoreVisibility() string {
	if o == nil || IsNil(o.StoreVisibility) {
		var ret string
		return ret
	}
	return *o.StoreVisibility
}

// GetStoreVisibilityOk returns a tuple with the StoreVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetStoreVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.StoreVisibility) {
		return nil, false
	}
	return o.StoreVisibility, true
}

// HasStoreVisibility returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasStoreVisibility() bool {
	if o != nil && !IsNil(o.StoreVisibility) {
		return true
	}

	return false
}

// SetStoreVisibility gets a reference to the given string and assigns it to the StoreVisibility field.
func (o *UpdateProductRequest) SetStoreVisibility(v string) {
	o.StoreVisibility = &v
}

// GetStoreVisibleFrom returns the StoreVisibleFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetStoreVisibleFrom() time.Time {
	if o == nil || IsNil(o.StoreVisibleFrom.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StoreVisibleFrom.Get()
}

// GetStoreVisibleFromOk returns a tuple with the StoreVisibleFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetStoreVisibleFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreVisibleFrom.Get(), o.StoreVisibleFrom.IsSet()
}

// HasStoreVisibleFrom returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasStoreVisibleFrom() bool {
	if o != nil && o.StoreVisibleFrom.IsSet() {
		return true
	}

	return false
}

// SetStoreVisibleFrom gets a reference to the given NullableTime and assigns it to the StoreVisibleFrom field.
func (o *UpdateProductRequest) SetStoreVisibleFrom(v time.Time) {
	o.StoreVisibleFrom.Set(&v)
}

// SetStoreVisibleFromNil sets the value for StoreVisibleFrom to be an explicit nil
func (o *UpdateProductRequest) SetStoreVisibleFromNil() {
	o.StoreVisibleFrom.Set(nil)
}

// UnsetStoreVisibleFrom ensures that no value is present for StoreVisibleFrom, not even an explicit nil
func (o *UpdateProductRequest) UnsetStoreVisibleFrom() {
	o.StoreVisibleFrom.Unset()
}

// GetFulfillmentType returns the FulfillmentType field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetFulfillmentType() string {
	if o == nil || IsNil(o.FulfillmentType) {
		var ret string
		return ret
	}
	return *o.FulfillmentType
}

// GetFulfillmentTypeOk returns a tuple with the FulfillmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetFulfillmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentType) {
		return nil, false
	}
	return o.FulfillmentType, true
}

// HasFulfillmentType returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasFulfillmentType() bool {
	if o != nil && !IsNil(o.FulfillmentType) {
		return true
	}

	return false
}

// SetFulfillmentType gets a reference to the given string and assigns it to the FulfillmentType field.
func (o *UpdateProductRequest) SetFulfillmentType(v string) {
	o.FulfillmentType = &v
}

// GetTrackInventory returns the TrackInventory field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetTrackInventory() bool {
	if o == nil || IsNil(o.TrackInventory) {
		var ret bool
		return ret
	}
	return *o.TrackInventory
}

// GetTrackInventoryOk returns a tuple with the TrackInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetTrackInventoryOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackInventory) {
		return nil, false
	}
	return o.TrackInventory, true
}

// HasTrackInventory returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasTrackInventory() bool {
	if o != nil && !IsNil(o.TrackInventory) {
		return true
	}

	return false
}

// SetTrackInventory gets a reference to the given bool and assigns it to the TrackInventory field.
func (o *UpdateProductRequest) SetTrackInventory(v bool) {
	o.TrackInventory = &v
}

// GetStockQuantity returns the StockQuantity field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *UpdateProductRequest) GetStockQuantity() int32 {
	if o == nil || IsNil(o.StockQuantity.Get()) {
		var ret int32
		return ret
	}
	return *o.StockQuantity.Get()
}

// GetStockQuantityOk returns a tuple with the StockQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *UpdateProductRequest) GetStockQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StockQuantity.Get(), o.StockQuantity.IsSet()
}

// HasStockQuantity returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasStockQuantity() bool {
	if o != nil && o.StockQuantity.IsSet() {
		return true
	}

	return false
}

// SetStockQuantity gets a reference to the given NullableInt32 and assigns it to the StockQuantity field.
// Deprecated
func (o *UpdateProductRequest) SetStockQuantity(v int32) {
	o.StockQuantity.Set(&v)
}

// SetStockQuantityNil sets the value for StockQuantity to be an explicit nil
func (o *UpdateProductRequest) SetStockQuantityNil() {
	o.StockQuantity.Set(nil)
}

// UnsetStockQuantity ensures that no value is present for StockQuantity, not even an explicit nil
func (o *UpdateProductRequest) UnsetStockQuantity() {
	o.StockQuantity.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateProductRequest) SetTags(v []string) {
	o.Tags = v
}

// GetBillingAnchorRule returns the BillingAnchorRule field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetBillingAnchorRule() string {
	if o == nil || IsNil(o.BillingAnchorRule) {
		var ret string
		return ret
	}
	return *o.BillingAnchorRule
}

// GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetBillingAnchorRuleOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAnchorRule) {
		return nil, false
	}
	return o.BillingAnchorRule, true
}

// HasBillingAnchorRule returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasBillingAnchorRule() bool {
	if o != nil && !IsNil(o.BillingAnchorRule) {
		return true
	}

	return false
}

// SetBillingAnchorRule gets a reference to the given string and assigns it to the BillingAnchorRule field.
func (o *UpdateProductRequest) SetBillingAnchorRule(v string) {
	o.BillingAnchorRule = &v
}

// GetBillingAnchorDay returns the BillingAnchorDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetBillingAnchorDay() int32 {
	if o == nil || IsNil(o.BillingAnchorDay.Get()) {
		var ret int32
		return ret
	}
	return *o.BillingAnchorDay.Get()
}

// GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetBillingAnchorDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAnchorDay.Get(), o.BillingAnchorDay.IsSet()
}

// HasBillingAnchorDay returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasBillingAnchorDay() bool {
	if o != nil && o.BillingAnchorDay.IsSet() {
		return true
	}

	return false
}

// SetBillingAnchorDay gets a reference to the given NullableInt32 and assigns it to the BillingAnchorDay field.
func (o *UpdateProductRequest) SetBillingAnchorDay(v int32) {
	o.BillingAnchorDay.Set(&v)
}

// SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil
func (o *UpdateProductRequest) SetBillingAnchorDayNil() {
	o.BillingAnchorDay.Set(nil)
}

// UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
func (o *UpdateProductRequest) UnsetBillingAnchorDay() {
	o.BillingAnchorDay.Unset()
}

// GetFirstChargeBehavior returns the FirstChargeBehavior field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetFirstChargeBehavior() string {
	if o == nil || IsNil(o.FirstChargeBehavior) {
		var ret string
		return ret
	}
	return *o.FirstChargeBehavior
}

// GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetFirstChargeBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.FirstChargeBehavior) {
		return nil, false
	}
	return o.FirstChargeBehavior, true
}

// HasFirstChargeBehavior returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasFirstChargeBehavior() bool {
	if o != nil && !IsNil(o.FirstChargeBehavior) {
		return true
	}

	return false
}

// SetFirstChargeBehavior gets a reference to the given string and assigns it to the FirstChargeBehavior field.
func (o *UpdateProductRequest) SetFirstChargeBehavior(v string) {
	o.FirstChargeBehavior = &v
}

// GetEndBehavior returns the EndBehavior field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetEndBehavior() string {
	if o == nil || IsNil(o.EndBehavior) {
		var ret string
		return ret
	}
	return *o.EndBehavior
}

// GetEndBehaviorOk returns a tuple with the EndBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetEndBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.EndBehavior) {
		return nil, false
	}
	return o.EndBehavior, true
}

// HasEndBehavior returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasEndBehavior() bool {
	if o != nil && !IsNil(o.EndBehavior) {
		return true
	}

	return false
}

// SetEndBehavior gets a reference to the given string and assigns it to the EndBehavior field.
func (o *UpdateProductRequest) SetEndBehavior(v string) {
	o.EndBehavior = &v
}

// GetMaxPayments returns the MaxPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetMaxPayments() int32 {
	if o == nil || IsNil(o.MaxPayments.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPayments.Get()
}

// GetMaxPaymentsOk returns a tuple with the MaxPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetMaxPaymentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPayments.Get(), o.MaxPayments.IsSet()
}

// HasMaxPayments returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasMaxPayments() bool {
	if o != nil && o.MaxPayments.IsSet() {
		return true
	}

	return false
}

// SetMaxPayments gets a reference to the given NullableInt32 and assigns it to the MaxPayments field.
func (o *UpdateProductRequest) SetMaxPayments(v int32) {
	o.MaxPayments.Set(&v)
}

// SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil
func (o *UpdateProductRequest) SetMaxPaymentsNil() {
	o.MaxPayments.Set(nil)
}

// UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
func (o *UpdateProductRequest) UnsetMaxPayments() {
	o.MaxPayments.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateProductRequest) GetMetadata() map[string]*interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]*interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetMetadataOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]*interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]*interface{} and assigns it to the Metadata field.
func (o *UpdateProductRequest) SetMetadata(v map[string]*interface{}) {
	o.Metadata = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProductRequest) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProductRequest) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UpdateProductRequest) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *UpdateProductRequest) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *UpdateProductRequest) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *UpdateProductRequest) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

func (o UpdateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.AmountMinor) {
		toSerialize["amount_minor"] = o.AmountMinor
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.IntervalCount.IsSet() {
		toSerialize["interval_count"] = o.IntervalCount.Get()
	}
	if !IsNil(o.TrialDays) {
		toSerialize["trial_days"] = o.TrialDays
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsArchived) {
		toSerialize["is_archived"] = o.IsArchived
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if o.SeoTitle.IsSet() {
		toSerialize["seo_title"] = o.SeoTitle.Get()
	}
	if o.SeoDescription.IsSet() {
		toSerialize["seo_description"] = o.SeoDescription.Get()
	}
	if !IsNil(o.StoreVisibility) {
		toSerialize["store_visibility"] = o.StoreVisibility
	}
	if o.StoreVisibleFrom.IsSet() {
		toSerialize["store_visible_from"] = o.StoreVisibleFrom.Get()
	}
	if !IsNil(o.FulfillmentType) {
		toSerialize["fulfillment_type"] = o.FulfillmentType
	}
	if !IsNil(o.TrackInventory) {
		toSerialize["track_inventory"] = o.TrackInventory
	}
	if o.StockQuantity.IsSet() {
		toSerialize["stock_quantity"] = o.StockQuantity.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.BillingAnchorRule) {
		toSerialize["billing_anchor_rule"] = o.BillingAnchorRule
	}
	if o.BillingAnchorDay.IsSet() {
		toSerialize["billing_anchor_day"] = o.BillingAnchorDay.Get()
	}
	if !IsNil(o.FirstChargeBehavior) {
		toSerialize["first_charge_behavior"] = o.FirstChargeBehavior
	}
	if !IsNil(o.EndBehavior) {
		toSerialize["end_behavior"] = o.EndBehavior
	}
	if o.MaxPayments.IsSet() {
		toSerialize["max_payments"] = o.MaxPayments.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateProductRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateProductRequest := _UpdateProductRequest{}

	err = json.Unmarshal(data, &varUpdateProductRequest)

	if err != nil {
		return err
	}

	*o = UpdateProductRequest(varUpdateProductRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_count")
		delete(additionalProperties, "trial_days")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_archived")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "images")
		delete(additionalProperties, "seo_title")
		delete(additionalProperties, "seo_description")
		delete(additionalProperties, "store_visibility")
		delete(additionalProperties, "store_visible_from")
		delete(additionalProperties, "fulfillment_type")
		delete(additionalProperties, "track_inventory")
		delete(additionalProperties, "stock_quantity")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "billing_anchor_rule")
		delete(additionalProperties, "billing_anchor_day")
		delete(additionalProperties, "first_charge_behavior")
		delete(additionalProperties, "end_behavior")
		delete(additionalProperties, "max_payments")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "image_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateProductRequest struct {
	value *UpdateProductRequest
	isSet bool
}

func (v NullableUpdateProductRequest) Get() *UpdateProductRequest {
	return v.value
}

func (v *NullableUpdateProductRequest) Set(val *UpdateProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductRequest(val *UpdateProductRequest) *NullableUpdateProductRequest {
	return &NullableUpdateProductRequest{value: val, isSet: true}
}

func (v NullableUpdateProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
