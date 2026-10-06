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
	"fmt"
)

// checks if the StorefrontPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorefrontPrice{}

// StorefrontPrice Active price summary for storefront display. Amounts are minor units.
type StorefrontPrice struct {
	Id                string         `json:"id"`
	ProductId         string         `json:"product_id"`
	Currency          string         `json:"currency"`
	UnitAmount        NullableInt32  `json:"unit_amount"`
	UnitAmountDecimal NullableString `json:"unit_amount_decimal"`
	Type              NullableString `json:"type"`
	Interval          NullableString `json:"interval"`
	IntervalCount     NullableInt32  `json:"interval_count"`
	BillingScheme     string         `json:"billing_scheme"`
	UsageType         NullableString `json:"usage_type"`
	PackageSize       NullableInt32  `json:"package_size"`
	TrialPeriodDays   NullableInt32  `json:"trial_period_days"`
	// The variant's options, e.g. {\"Size\":\"M\"}. Null when the price is not a variant.
	OptionValues map[string]string `json:"option_values,omitempty"`
	// The variant's SKU. Null when not set.
	Sku NullableString `json:"sku,omitempty"`
	// The variant's image as a browser URL. Null when not set.
	ImageUrl             NullableString `json:"image_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorefrontPrice StorefrontPrice

// NewStorefrontPrice instantiates a new StorefrontPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorefrontPrice(id string, productId string, currency string, unitAmount NullableInt32, unitAmountDecimal NullableString, type_ NullableString, interval NullableString, intervalCount NullableInt32, billingScheme string, usageType NullableString, packageSize NullableInt32, trialPeriodDays NullableInt32) *StorefrontPrice {
	this := StorefrontPrice{}
	this.Id = id
	this.ProductId = productId
	this.Currency = currency
	this.UnitAmount = unitAmount
	this.UnitAmountDecimal = unitAmountDecimal
	this.Type = type_
	this.Interval = interval
	this.IntervalCount = intervalCount
	this.BillingScheme = billingScheme
	this.UsageType = usageType
	this.PackageSize = packageSize
	this.TrialPeriodDays = trialPeriodDays
	return &this
}

// NewStorefrontPriceWithDefaults instantiates a new StorefrontPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorefrontPriceWithDefaults() *StorefrontPrice {
	this := StorefrontPrice{}
	return &this
}

// GetId returns the Id field value
func (o *StorefrontPrice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorefrontPrice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StorefrontPrice) SetId(v string) {
	o.Id = v
}

// GetProductId returns the ProductId field value
func (o *StorefrontPrice) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *StorefrontPrice) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *StorefrontPrice) SetProductId(v string) {
	o.ProductId = v
}

// GetCurrency returns the Currency field value
func (o *StorefrontPrice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *StorefrontPrice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *StorefrontPrice) SetCurrency(v string) {
	o.Currency = v
}

// GetUnitAmount returns the UnitAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StorefrontPrice) GetUnitAmount() int32 {
	if o == nil || o.UnitAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UnitAmount.Get()
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitAmount.Get(), o.UnitAmount.IsSet()
}

// SetUnitAmount sets field value
func (o *StorefrontPrice) SetUnitAmount(v int32) {
	o.UnitAmount.Set(&v)
}

// GetUnitAmountDecimal returns the UnitAmountDecimal field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontPrice) GetUnitAmountDecimal() string {
	if o == nil || o.UnitAmountDecimal.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnitAmountDecimal.Get()
}

// GetUnitAmountDecimalOk returns a tuple with the UnitAmountDecimal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetUnitAmountDecimalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitAmountDecimal.Get(), o.UnitAmountDecimal.IsSet()
}

// SetUnitAmountDecimal sets field value
func (o *StorefrontPrice) SetUnitAmountDecimal(v string) {
	o.UnitAmountDecimal.Set(&v)
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontPrice) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *StorefrontPrice) SetType(v string) {
	o.Type.Set(&v)
}

// GetInterval returns the Interval field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontPrice) GetInterval() string {
	if o == nil || o.Interval.Get() == nil {
		var ret string
		return ret
	}

	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// SetInterval sets field value
func (o *StorefrontPrice) SetInterval(v string) {
	o.Interval.Set(&v)
}

// GetIntervalCount returns the IntervalCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StorefrontPrice) GetIntervalCount() int32 {
	if o == nil || o.IntervalCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// SetIntervalCount sets field value
func (o *StorefrontPrice) SetIntervalCount(v int32) {
	o.IntervalCount.Set(&v)
}

// GetBillingScheme returns the BillingScheme field value
func (o *StorefrontPrice) GetBillingScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingScheme
}

// GetBillingSchemeOk returns a tuple with the BillingScheme field value
// and a boolean to check if the value has been set.
func (o *StorefrontPrice) GetBillingSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingScheme, true
}

// SetBillingScheme sets field value
func (o *StorefrontPrice) SetBillingScheme(v string) {
	o.BillingScheme = v
}

// GetUsageType returns the UsageType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontPrice) GetUsageType() string {
	if o == nil || o.UsageType.Get() == nil {
		var ret string
		return ret
	}

	return *o.UsageType.Get()
}

// GetUsageTypeOk returns a tuple with the UsageType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetUsageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageType.Get(), o.UsageType.IsSet()
}

// SetUsageType sets field value
func (o *StorefrontPrice) SetUsageType(v string) {
	o.UsageType.Set(&v)
}

// GetPackageSize returns the PackageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StorefrontPrice) GetPackageSize() int32 {
	if o == nil || o.PackageSize.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PackageSize.Get()
}

// GetPackageSizeOk returns a tuple with the PackageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetPackageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageSize.Get(), o.PackageSize.IsSet()
}

// SetPackageSize sets field value
func (o *StorefrontPrice) SetPackageSize(v int32) {
	o.PackageSize.Set(&v)
}

// GetTrialPeriodDays returns the TrialPeriodDays field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *StorefrontPrice) GetTrialPeriodDays() int32 {
	if o == nil || o.TrialPeriodDays.Get() == nil {
		var ret int32
		return ret
	}

	return *o.TrialPeriodDays.Get()
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialPeriodDays.Get(), o.TrialPeriodDays.IsSet()
}

// SetTrialPeriodDays sets field value
func (o *StorefrontPrice) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays.Set(&v)
}

// GetOptionValues returns the OptionValues field value if set, zero value otherwise.
func (o *StorefrontPrice) GetOptionValues() map[string]string {
	if o == nil || IsNil(o.OptionValues) {
		var ret map[string]string
		return ret
	}
	return o.OptionValues
}

// GetOptionValuesOk returns a tuple with the OptionValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorefrontPrice) GetOptionValuesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.OptionValues) {
		return map[string]string{}, false
	}
	return o.OptionValues, true
}

// HasOptionValues returns a boolean if a field has been set.
func (o *StorefrontPrice) HasOptionValues() bool {
	if o != nil && !IsNil(o.OptionValues) {
		return true
	}

	return false
}

// SetOptionValues gets a reference to the given map[string]string and assigns it to the OptionValues field.
func (o *StorefrontPrice) SetOptionValues(v map[string]string) {
	o.OptionValues = v
}

// GetSku returns the Sku field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorefrontPrice) GetSku() string {
	if o == nil || IsNil(o.Sku.Get()) {
		var ret string
		return ret
	}
	return *o.Sku.Get()
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sku.Get(), o.Sku.IsSet()
}

// HasSku returns a boolean if a field has been set.
func (o *StorefrontPrice) HasSku() bool {
	if o != nil && o.Sku.IsSet() {
		return true
	}

	return false
}

// SetSku gets a reference to the given NullableString and assigns it to the Sku field.
func (o *StorefrontPrice) SetSku(v string) {
	o.Sku.Set(&v)
}

// SetSkuNil sets the value for Sku to be an explicit nil
func (o *StorefrontPrice) SetSkuNil() {
	o.Sku.Set(nil)
}

// UnsetSku ensures that no value is present for Sku, not even an explicit nil
func (o *StorefrontPrice) UnsetSku() {
	o.Sku.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorefrontPrice) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontPrice) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *StorefrontPrice) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *StorefrontPrice) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *StorefrontPrice) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *StorefrontPrice) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

func (o StorefrontPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorefrontPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["product_id"] = o.ProductId
	toSerialize["currency"] = o.Currency
	toSerialize["unit_amount"] = o.UnitAmount.Get()
	toSerialize["unit_amount_decimal"] = o.UnitAmountDecimal.Get()
	toSerialize["type"] = o.Type.Get()
	toSerialize["interval"] = o.Interval.Get()
	toSerialize["interval_count"] = o.IntervalCount.Get()
	toSerialize["billing_scheme"] = o.BillingScheme
	toSerialize["usage_type"] = o.UsageType.Get()
	toSerialize["package_size"] = o.PackageSize.Get()
	toSerialize["trial_period_days"] = o.TrialPeriodDays.Get()
	if !IsNil(o.OptionValues) {
		toSerialize["option_values"] = o.OptionValues
	}
	if o.Sku.IsSet() {
		toSerialize["sku"] = o.Sku.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorefrontPrice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"product_id",
		"currency",
		"unit_amount",
		"unit_amount_decimal",
		"type",
		"interval",
		"interval_count",
		"billing_scheme",
		"usage_type",
		"package_size",
		"trial_period_days",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varStorefrontPrice := _StorefrontPrice{}

	err = json.Unmarshal(data, &varStorefrontPrice)

	if err != nil {
		return err
	}

	*o = StorefrontPrice(varStorefrontPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "unit_amount")
		delete(additionalProperties, "unit_amount_decimal")
		delete(additionalProperties, "type")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_count")
		delete(additionalProperties, "billing_scheme")
		delete(additionalProperties, "usage_type")
		delete(additionalProperties, "package_size")
		delete(additionalProperties, "trial_period_days")
		delete(additionalProperties, "option_values")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "image_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorefrontPrice struct {
	value *StorefrontPrice
	isSet bool
}

func (v NullableStorefrontPrice) Get() *StorefrontPrice {
	return v.value
}

func (v *NullableStorefrontPrice) Set(val *StorefrontPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableStorefrontPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableStorefrontPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorefrontPrice(val *StorefrontPrice) *NullableStorefrontPrice {
	return &NullableStorefrontPrice{value: val, isSet: true}
}

func (v NullableStorefrontPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorefrontPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
