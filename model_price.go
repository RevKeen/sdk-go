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

// checks if the Price type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Price{}

// Price A price defines how much and how often to charge for a product. A single product can have multiple prices for different currencies, intervals, or tiers.
type Price struct {
	// Unique identifier for the price
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// ID of the product this price belongs to
	ProductId string `json:"product_id"`
	// Whether the price is active
	Active bool `json:"active"`
	// Three-letter ISO currency code (lowercase)
	Currency string `json:"currency"`
	// Price in minor units (cents)
	UnitAmount NullableInt32 `json:"unit_amount"`
	// Price type
	Type string `json:"type"`
	// Pricing model
	PricingModel string `json:"pricing_model"`
	// Billing interval (recurring only)
	Interval NullableString `json:"interval,omitempty"`
	// Number of intervals between billings
	IntervalCount NullableInt32 `json:"interval_count,omitempty"`
	// Trial period in days (recurring only)
	TrialPeriodDays NullableInt32 `json:"trial_period_days,omitempty"`
	// PWYW: minimum amount in cents
	MinimumAmount NullableInt32 `json:"minimum_amount,omitempty"`
	// PWYW: maximum amount in cents
	MaximumAmount NullableInt32 `json:"maximum_amount,omitempty"`
	// PWYW: suggested amount in cents
	SuggestedAmount NullableInt32 `json:"suggested_amount,omitempty"`
	// PWYW: quick-select amounts
	PresetAmounts []int32 `json:"preset_amounts,omitempty"`
	// Display name (e.g., 'Monthly', 'Annual - Save 17%')
	Nickname NullableString `json:"nickname,omitempty"`
	// Stable key for API lookups
	LookupKey NullableString `json:"lookup_key,omitempty"`
	// Billing scheme. Defaults to `per_unit`. Set to `tiered` along with `tiers_mode` + `tiers` to use graduated or volume pricing.
	BillingScheme string `json:"billing_scheme"`
	// Tiered pricing mode. Required when `billing_scheme` is `tiered`, must be null otherwise.
	TiersMode NullableString `json:"tiers_mode,omitempty"`
	// Price tiers (ordered by `up_to`). Present only when `billing_scheme` is `tiered`. At least 2 tiers, final tier MUST have `up_to: null`. Immutable after creation.
	Tiers []PriceTier `json:"tiers,omitempty"`
	TransformQuantity *PriceTransformQuantity `json:"transform_quantity,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Price Price

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(id string, object string, productId string, active bool, currency string, unitAmount NullableInt32, type_ string, pricingModel string, billingScheme string, createdAt time.Time, updatedAt time.Time) *Price {
	this := Price{}
	this.Id = id
	this.Object = object
	this.ProductId = productId
	this.Active = active
	this.Currency = currency
	this.UnitAmount = unitAmount
	this.Type = type_
	this.PricingModel = pricingModel
	this.BillingScheme = billingScheme
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetId returns the Id field value
func (o *Price) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Price) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Price) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Price) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Price) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Price) SetObject(v string) {
	o.Object = v
}

// GetProductId returns the ProductId field value
func (o *Price) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *Price) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *Price) SetProductId(v string) {
	o.ProductId = v
}

// GetActive returns the Active field value
func (o *Price) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *Price) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *Price) SetActive(v bool) {
	o.Active = v
}

// GetCurrency returns the Currency field value
func (o *Price) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Price) SetCurrency(v string) {
	o.Currency = v
}

// GetUnitAmount returns the UnitAmount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Price) GetUnitAmount() int32 {
	if o == nil || o.UnitAmount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.UnitAmount.Get()
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitAmount.Get(), o.UnitAmount.IsSet()
}

// SetUnitAmount sets field value
func (o *Price) SetUnitAmount(v int32) {
	o.UnitAmount.Set(&v)
}

// GetType returns the Type field value
func (o *Price) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Price) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Price) SetType(v string) {
	o.Type = v
}

// GetPricingModel returns the PricingModel field value
func (o *Price) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *Price) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *Price) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetInterval() string {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret string
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Price) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableString and assigns it to the Interval field.
func (o *Price) SetInterval(v string) {
	o.Interval.Set(&v)
}
// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *Price) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *Price) UnsetInterval() {
	o.Interval.Unset()
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *Price) HasIntervalCount() bool {
	if o != nil && o.IntervalCount.IsSet() {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given NullableInt32 and assigns it to the IntervalCount field.
func (o *Price) SetIntervalCount(v int32) {
	o.IntervalCount.Set(&v)
}
// SetIntervalCountNil sets the value for IntervalCount to be an explicit nil
func (o *Price) SetIntervalCountNil() {
	o.IntervalCount.Set(nil)
}

// UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
func (o *Price) UnsetIntervalCount() {
	o.IntervalCount.Unset()
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays.Get()) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays.Get()
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialPeriodDays.Get(), o.TrialPeriodDays.IsSet()
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *Price) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays.IsSet() {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given NullableInt32 and assigns it to the TrialPeriodDays field.
func (o *Price) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays.Set(&v)
}
// SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil
func (o *Price) SetTrialPeriodDaysNil() {
	o.TrialPeriodDays.Set(nil)
}

// UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
func (o *Price) UnsetTrialPeriodDays() {
	o.TrialPeriodDays.Unset()
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetMinimumAmount() int32 {
	if o == nil || IsNil(o.MinimumAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.MinimumAmount.Get()
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetMinimumAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumAmount.Get(), o.MinimumAmount.IsSet()
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *Price) HasMinimumAmount() bool {
	if o != nil && o.MinimumAmount.IsSet() {
		return true
	}

	return false
}

// SetMinimumAmount gets a reference to the given NullableInt32 and assigns it to the MinimumAmount field.
func (o *Price) SetMinimumAmount(v int32) {
	o.MinimumAmount.Set(&v)
}
// SetMinimumAmountNil sets the value for MinimumAmount to be an explicit nil
func (o *Price) SetMinimumAmountNil() {
	o.MinimumAmount.Set(nil)
}

// UnsetMinimumAmount ensures that no value is present for MinimumAmount, not even an explicit nil
func (o *Price) UnsetMinimumAmount() {
	o.MinimumAmount.Unset()
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetMaximumAmount() int32 {
	if o == nil || IsNil(o.MaximumAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumAmount.Get()
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetMaximumAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumAmount.Get(), o.MaximumAmount.IsSet()
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *Price) HasMaximumAmount() bool {
	if o != nil && o.MaximumAmount.IsSet() {
		return true
	}

	return false
}

// SetMaximumAmount gets a reference to the given NullableInt32 and assigns it to the MaximumAmount field.
func (o *Price) SetMaximumAmount(v int32) {
	o.MaximumAmount.Set(&v)
}
// SetMaximumAmountNil sets the value for MaximumAmount to be an explicit nil
func (o *Price) SetMaximumAmountNil() {
	o.MaximumAmount.Set(nil)
}

// UnsetMaximumAmount ensures that no value is present for MaximumAmount, not even an explicit nil
func (o *Price) UnsetMaximumAmount() {
	o.MaximumAmount.Unset()
}

// GetSuggestedAmount returns the SuggestedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetSuggestedAmount() int32 {
	if o == nil || IsNil(o.SuggestedAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.SuggestedAmount.Get()
}

// GetSuggestedAmountOk returns a tuple with the SuggestedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetSuggestedAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuggestedAmount.Get(), o.SuggestedAmount.IsSet()
}

// HasSuggestedAmount returns a boolean if a field has been set.
func (o *Price) HasSuggestedAmount() bool {
	if o != nil && o.SuggestedAmount.IsSet() {
		return true
	}

	return false
}

// SetSuggestedAmount gets a reference to the given NullableInt32 and assigns it to the SuggestedAmount field.
func (o *Price) SetSuggestedAmount(v int32) {
	o.SuggestedAmount.Set(&v)
}
// SetSuggestedAmountNil sets the value for SuggestedAmount to be an explicit nil
func (o *Price) SetSuggestedAmountNil() {
	o.SuggestedAmount.Set(nil)
}

// UnsetSuggestedAmount ensures that no value is present for SuggestedAmount, not even an explicit nil
func (o *Price) UnsetSuggestedAmount() {
	o.SuggestedAmount.Unset()
}

// GetPresetAmounts returns the PresetAmounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetPresetAmounts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.PresetAmounts
}

// GetPresetAmountsOk returns a tuple with the PresetAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetPresetAmountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.PresetAmounts) {
		return nil, false
	}
	return o.PresetAmounts, true
}

// HasPresetAmounts returns a boolean if a field has been set.
func (o *Price) HasPresetAmounts() bool {
	if o != nil && !IsNil(o.PresetAmounts) {
		return true
	}

	return false
}

// SetPresetAmounts gets a reference to the given []int32 and assigns it to the PresetAmounts field.
func (o *Price) SetPresetAmounts(v []int32) {
	o.PresetAmounts = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetNickname() string {
	if o == nil || IsNil(o.Nickname.Get()) {
		var ret string
		return ret
	}
	return *o.Nickname.Get()
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetNicknameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nickname.Get(), o.Nickname.IsSet()
}

// HasNickname returns a boolean if a field has been set.
func (o *Price) HasNickname() bool {
	if o != nil && o.Nickname.IsSet() {
		return true
	}

	return false
}

// SetNickname gets a reference to the given NullableString and assigns it to the Nickname field.
func (o *Price) SetNickname(v string) {
	o.Nickname.Set(&v)
}
// SetNicknameNil sets the value for Nickname to be an explicit nil
func (o *Price) SetNicknameNil() {
	o.Nickname.Set(nil)
}

// UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
func (o *Price) UnsetNickname() {
	o.Nickname.Unset()
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey.Get()) {
		var ret string
		return ret
	}
	return *o.LookupKey.Get()
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetLookupKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LookupKey.Get(), o.LookupKey.IsSet()
}

// HasLookupKey returns a boolean if a field has been set.
func (o *Price) HasLookupKey() bool {
	if o != nil && o.LookupKey.IsSet() {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given NullableString and assigns it to the LookupKey field.
func (o *Price) SetLookupKey(v string) {
	o.LookupKey.Set(&v)
}
// SetLookupKeyNil sets the value for LookupKey to be an explicit nil
func (o *Price) SetLookupKeyNil() {
	o.LookupKey.Set(nil)
}

// UnsetLookupKey ensures that no value is present for LookupKey, not even an explicit nil
func (o *Price) UnsetLookupKey() {
	o.LookupKey.Unset()
}

// GetBillingScheme returns the BillingScheme field value
func (o *Price) GetBillingScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingScheme
}

// GetBillingSchemeOk returns a tuple with the BillingScheme field value
// and a boolean to check if the value has been set.
func (o *Price) GetBillingSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingScheme, true
}

// SetBillingScheme sets field value
func (o *Price) SetBillingScheme(v string) {
	o.BillingScheme = v
}

// GetTiersMode returns the TiersMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetTiersMode() string {
	if o == nil || IsNil(o.TiersMode.Get()) {
		var ret string
		return ret
	}
	return *o.TiersMode.Get()
}

// GetTiersModeOk returns a tuple with the TiersMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetTiersModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TiersMode.Get(), o.TiersMode.IsSet()
}

// HasTiersMode returns a boolean if a field has been set.
func (o *Price) HasTiersMode() bool {
	if o != nil && o.TiersMode.IsSet() {
		return true
	}

	return false
}

// SetTiersMode gets a reference to the given NullableString and assigns it to the TiersMode field.
func (o *Price) SetTiersMode(v string) {
	o.TiersMode.Set(&v)
}
// SetTiersModeNil sets the value for TiersMode to be an explicit nil
func (o *Price) SetTiersModeNil() {
	o.TiersMode.Set(nil)
}

// UnsetTiersMode ensures that no value is present for TiersMode, not even an explicit nil
func (o *Price) UnsetTiersMode() {
	o.TiersMode.Unset()
}

// GetTiers returns the Tiers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Price) GetTiers() []PriceTier {
	if o == nil {
		var ret []PriceTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Price) GetTiersOk() ([]PriceTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *Price) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []PriceTier and assigns it to the Tiers field.
func (o *Price) SetTiers(v []PriceTier) {
	o.Tiers = v
}

// GetTransformQuantity returns the TransformQuantity field value if set, zero value otherwise.
func (o *Price) GetTransformQuantity() PriceTransformQuantity {
	if o == nil || IsNil(o.TransformQuantity) {
		var ret PriceTransformQuantity
		return ret
	}
	return *o.TransformQuantity
}

// GetTransformQuantityOk returns a tuple with the TransformQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetTransformQuantityOk() (*PriceTransformQuantity, bool) {
	if o == nil || IsNil(o.TransformQuantity) {
		return nil, false
	}
	return o.TransformQuantity, true
}

// HasTransformQuantity returns a boolean if a field has been set.
func (o *Price) HasTransformQuantity() bool {
	if o != nil && !IsNil(o.TransformQuantity) {
		return true
	}

	return false
}

// SetTransformQuantity gets a reference to the given PriceTransformQuantity and assigns it to the TransformQuantity field.
func (o *Price) SetTransformQuantity(v PriceTransformQuantity) {
	o.TransformQuantity = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Price) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Price) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Price) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Price) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Price) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Price) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Price) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Price) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Price) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Price) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["product_id"] = o.ProductId
	toSerialize["active"] = o.Active
	toSerialize["currency"] = o.Currency
	toSerialize["unit_amount"] = o.UnitAmount.Get()
	toSerialize["type"] = o.Type
	toSerialize["pricing_model"] = o.PricingModel
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.IntervalCount.IsSet() {
		toSerialize["interval_count"] = o.IntervalCount.Get()
	}
	if o.TrialPeriodDays.IsSet() {
		toSerialize["trial_period_days"] = o.TrialPeriodDays.Get()
	}
	if o.MinimumAmount.IsSet() {
		toSerialize["minimum_amount"] = o.MinimumAmount.Get()
	}
	if o.MaximumAmount.IsSet() {
		toSerialize["maximum_amount"] = o.MaximumAmount.Get()
	}
	if o.SuggestedAmount.IsSet() {
		toSerialize["suggested_amount"] = o.SuggestedAmount.Get()
	}
	if o.PresetAmounts != nil {
		toSerialize["preset_amounts"] = o.PresetAmounts
	}
	if o.Nickname.IsSet() {
		toSerialize["nickname"] = o.Nickname.Get()
	}
	if o.LookupKey.IsSet() {
		toSerialize["lookup_key"] = o.LookupKey.Get()
	}
	toSerialize["billing_scheme"] = o.BillingScheme
	if o.TiersMode.IsSet() {
		toSerialize["tiers_mode"] = o.TiersMode.Get()
	}
	if o.Tiers != nil {
		toSerialize["tiers"] = o.Tiers
	}
	if !IsNil(o.TransformQuantity) {
		toSerialize["transform_quantity"] = o.TransformQuantity
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

func (o *Price) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"product_id",
		"active",
		"currency",
		"unit_amount",
		"type",
		"pricing_model",
		"billing_scheme",
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

	varPrice := _Price{}

	err = json.Unmarshal(data, &varPrice)

	if err != nil {
		return err
	}

	*o = Price(varPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "active")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "unit_amount")
		delete(additionalProperties, "type")
		delete(additionalProperties, "pricing_model")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_count")
		delete(additionalProperties, "trial_period_days")
		delete(additionalProperties, "minimum_amount")
		delete(additionalProperties, "maximum_amount")
		delete(additionalProperties, "suggested_amount")
		delete(additionalProperties, "preset_amounts")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "lookup_key")
		delete(additionalProperties, "billing_scheme")
		delete(additionalProperties, "tiers_mode")
		delete(additionalProperties, "tiers")
		delete(additionalProperties, "transform_quantity")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


