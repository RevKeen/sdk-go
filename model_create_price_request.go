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
	"fmt"
)

// checks if the CreatePriceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePriceRequest{}

// CreatePriceRequest Parameters for creating a new price. Specify the product, amount, currency, and billing interval.
type CreatePriceRequest struct {
	// ID of the product this price belongs to
	ProductId string `json:"product_id"`
	// Three-letter ISO currency code
	Currency *string `json:"currency,omitempty"`
	// Price in minor units (cents). Required for fixed pricing.
	UnitAmount *int32 `json:"unit_amount,omitempty"`
	// Price type
	Type *string `json:"type,omitempty"`
	// Pricing model
	PricingModel *string `json:"pricing_model,omitempty"`
	// Billing interval (required for recurring)
	Interval *string `json:"interval,omitempty"`
	// Number of intervals between billings (1-12)
	IntervalCount *int32 `json:"interval_count,omitempty"`
	// Trial period in days (0-730)
	TrialPeriodDays *int32 `json:"trial_period_days,omitempty"`
	// PWYW: minimum amount in cents
	MinimumAmount *int32 `json:"minimum_amount,omitempty"`
	// PWYW: maximum amount in cents
	MaximumAmount *int32 `json:"maximum_amount,omitempty"`
	// PWYW: suggested amount in cents
	SuggestedAmount *int32 `json:"suggested_amount,omitempty"`
	// PWYW: quick-select amounts (max 10)
	PresetAmounts []int32 `json:"preset_amounts,omitempty"`
	// Display name (e.g., 'Monthly', 'Annual')
	Nickname *string `json:"nickname,omitempty"`
	// Stable key for API lookups
	LookupKey *string `json:"lookup_key,omitempty"`
	// Billing scheme. Defaults to `per_unit`. Set to `tiered` with `tiers_mode` + `tiers` for graduated or volume pricing.
	BillingScheme *string `json:"billing_scheme,omitempty"`
	// Tiered pricing mode. Required when `billing_scheme` is `tiered`, must be omitted otherwise.
	TiersMode *string `json:"tiers_mode,omitempty"`
	// Price tiers (ordered by `up_to`, min 2, max 50). Required when `billing_scheme` is `tiered`. Final tier MUST have `up_to: null`.
	Tiers []PriceTier `json:"tiers,omitempty"`
	TransformQuantity *TransformQuantity `json:"transform_quantity,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePriceRequest CreatePriceRequest

// NewCreatePriceRequest instantiates a new CreatePriceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePriceRequest(productId string) *CreatePriceRequest {
	this := CreatePriceRequest{}
	this.ProductId = productId
	var currency string = "usd"
	this.Currency = &currency
	var type_ string = "one_time"
	this.Type = &type_
	var pricingModel string = "fixed"
	this.PricingModel = &pricingModel
	return &this
}

// NewCreatePriceRequestWithDefaults instantiates a new CreatePriceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePriceRequestWithDefaults() *CreatePriceRequest {
	this := CreatePriceRequest{}
	var currency string = "usd"
	this.Currency = &currency
	var type_ string = "one_time"
	this.Type = &type_
	var pricingModel string = "fixed"
	this.PricingModel = &pricingModel
	return &this
}

// GetProductId returns the ProductId field value
func (o *CreatePriceRequest) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *CreatePriceRequest) SetProductId(v string) {
	o.ProductId = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreatePriceRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetUnitAmount returns the UnitAmount field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetUnitAmount() int32 {
	if o == nil || IsNil(o.UnitAmount) {
		var ret int32
		return ret
	}
	return *o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetUnitAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitAmount) {
		return nil, false
	}
	return o.UnitAmount, true
}

// HasUnitAmount returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasUnitAmount() bool {
	if o != nil && !IsNil(o.UnitAmount) {
		return true
	}

	return false
}

// SetUnitAmount gets a reference to the given int32 and assigns it to the UnitAmount field.
func (o *CreatePriceRequest) SetUnitAmount(v int32) {
	o.UnitAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreatePriceRequest) SetType(v string) {
	o.Type = &v
}

// GetPricingModel returns the PricingModel field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetPricingModel() string {
	if o == nil || IsNil(o.PricingModel) {
		var ret string
		return ret
	}
	return *o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetPricingModelOk() (*string, bool) {
	if o == nil || IsNil(o.PricingModel) {
		return nil, false
	}
	return o.PricingModel, true
}

// HasPricingModel returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasPricingModel() bool {
	if o != nil && !IsNil(o.PricingModel) {
		return true
	}

	return false
}

// SetPricingModel gets a reference to the given string and assigns it to the PricingModel field.
func (o *CreatePriceRequest) SetPricingModel(v string) {
	o.PricingModel = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *CreatePriceRequest) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *CreatePriceRequest) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriodDays) {
		return nil, false
	}
	return o.TrialPeriodDays, true
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasTrialPeriodDays() bool {
	if o != nil && !IsNil(o.TrialPeriodDays) {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given int32 and assigns it to the TrialPeriodDays field.
func (o *CreatePriceRequest) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetMinimumAmount() int32 {
	if o == nil || IsNil(o.MinimumAmount) {
		var ret int32
		return ret
	}
	return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetMinimumAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumAmount) {
		return nil, false
	}
	return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasMinimumAmount() bool {
	if o != nil && !IsNil(o.MinimumAmount) {
		return true
	}

	return false
}

// SetMinimumAmount gets a reference to the given int32 and assigns it to the MinimumAmount field.
func (o *CreatePriceRequest) SetMinimumAmount(v int32) {
	o.MinimumAmount = &v
}

// GetMaximumAmount returns the MaximumAmount field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetMaximumAmount() int32 {
	if o == nil || IsNil(o.MaximumAmount) {
		var ret int32
		return ret
	}
	return *o.MaximumAmount
}

// GetMaximumAmountOk returns a tuple with the MaximumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetMaximumAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumAmount) {
		return nil, false
	}
	return o.MaximumAmount, true
}

// HasMaximumAmount returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasMaximumAmount() bool {
	if o != nil && !IsNil(o.MaximumAmount) {
		return true
	}

	return false
}

// SetMaximumAmount gets a reference to the given int32 and assigns it to the MaximumAmount field.
func (o *CreatePriceRequest) SetMaximumAmount(v int32) {
	o.MaximumAmount = &v
}

// GetSuggestedAmount returns the SuggestedAmount field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetSuggestedAmount() int32 {
	if o == nil || IsNil(o.SuggestedAmount) {
		var ret int32
		return ret
	}
	return *o.SuggestedAmount
}

// GetSuggestedAmountOk returns a tuple with the SuggestedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetSuggestedAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.SuggestedAmount) {
		return nil, false
	}
	return o.SuggestedAmount, true
}

// HasSuggestedAmount returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasSuggestedAmount() bool {
	if o != nil && !IsNil(o.SuggestedAmount) {
		return true
	}

	return false
}

// SetSuggestedAmount gets a reference to the given int32 and assigns it to the SuggestedAmount field.
func (o *CreatePriceRequest) SetSuggestedAmount(v int32) {
	o.SuggestedAmount = &v
}

// GetPresetAmounts returns the PresetAmounts field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetPresetAmounts() []int32 {
	if o == nil || IsNil(o.PresetAmounts) {
		var ret []int32
		return ret
	}
	return o.PresetAmounts
}

// GetPresetAmountsOk returns a tuple with the PresetAmounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetPresetAmountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.PresetAmounts) {
		return nil, false
	}
	return o.PresetAmounts, true
}

// HasPresetAmounts returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasPresetAmounts() bool {
	if o != nil && !IsNil(o.PresetAmounts) {
		return true
	}

	return false
}

// SetPresetAmounts gets a reference to the given []int32 and assigns it to the PresetAmounts field.
func (o *CreatePriceRequest) SetPresetAmounts(v []int32) {
	o.PresetAmounts = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *CreatePriceRequest) SetNickname(v string) {
	o.Nickname = &v
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey) {
		var ret string
		return ret
	}
	return *o.LookupKey
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetLookupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LookupKey) {
		return nil, false
	}
	return o.LookupKey, true
}

// HasLookupKey returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasLookupKey() bool {
	if o != nil && !IsNil(o.LookupKey) {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given string and assigns it to the LookupKey field.
func (o *CreatePriceRequest) SetLookupKey(v string) {
	o.LookupKey = &v
}

// GetBillingScheme returns the BillingScheme field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetBillingScheme() string {
	if o == nil || IsNil(o.BillingScheme) {
		var ret string
		return ret
	}
	return *o.BillingScheme
}

// GetBillingSchemeOk returns a tuple with the BillingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetBillingSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.BillingScheme) {
		return nil, false
	}
	return o.BillingScheme, true
}

// HasBillingScheme returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasBillingScheme() bool {
	if o != nil && !IsNil(o.BillingScheme) {
		return true
	}

	return false
}

// SetBillingScheme gets a reference to the given string and assigns it to the BillingScheme field.
func (o *CreatePriceRequest) SetBillingScheme(v string) {
	o.BillingScheme = &v
}

// GetTiersMode returns the TiersMode field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetTiersMode() string {
	if o == nil || IsNil(o.TiersMode) {
		var ret string
		return ret
	}
	return *o.TiersMode
}

// GetTiersModeOk returns a tuple with the TiersMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetTiersModeOk() (*string, bool) {
	if o == nil || IsNil(o.TiersMode) {
		return nil, false
	}
	return o.TiersMode, true
}

// HasTiersMode returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasTiersMode() bool {
	if o != nil && !IsNil(o.TiersMode) {
		return true
	}

	return false
}

// SetTiersMode gets a reference to the given string and assigns it to the TiersMode field.
func (o *CreatePriceRequest) SetTiersMode(v string) {
	o.TiersMode = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetTiers() []PriceTier {
	if o == nil || IsNil(o.Tiers) {
		var ret []PriceTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetTiersOk() ([]PriceTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []PriceTier and assigns it to the Tiers field.
func (o *CreatePriceRequest) SetTiers(v []PriceTier) {
	o.Tiers = v
}

// GetTransformQuantity returns the TransformQuantity field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetTransformQuantity() TransformQuantity {
	if o == nil || IsNil(o.TransformQuantity) {
		var ret TransformQuantity
		return ret
	}
	return *o.TransformQuantity
}

// GetTransformQuantityOk returns a tuple with the TransformQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetTransformQuantityOk() (*TransformQuantity, bool) {
	if o == nil || IsNil(o.TransformQuantity) {
		return nil, false
	}
	return o.TransformQuantity, true
}

// HasTransformQuantity returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasTransformQuantity() bool {
	if o != nil && !IsNil(o.TransformQuantity) {
		return true
	}

	return false
}

// SetTransformQuantity gets a reference to the given TransformQuantity and assigns it to the TransformQuantity field.
func (o *CreatePriceRequest) SetTransformQuantity(v TransformQuantity) {
	o.TransformQuantity = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreatePriceRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePriceRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreatePriceRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreatePriceRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CreatePriceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePriceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_id"] = o.ProductId
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.UnitAmount) {
		toSerialize["unit_amount"] = o.UnitAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PricingModel) {
		toSerialize["pricing_model"] = o.PricingModel
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.IntervalCount) {
		toSerialize["interval_count"] = o.IntervalCount
	}
	if !IsNil(o.TrialPeriodDays) {
		toSerialize["trial_period_days"] = o.TrialPeriodDays
	}
	if !IsNil(o.MinimumAmount) {
		toSerialize["minimum_amount"] = o.MinimumAmount
	}
	if !IsNil(o.MaximumAmount) {
		toSerialize["maximum_amount"] = o.MaximumAmount
	}
	if !IsNil(o.SuggestedAmount) {
		toSerialize["suggested_amount"] = o.SuggestedAmount
	}
	if !IsNil(o.PresetAmounts) {
		toSerialize["preset_amounts"] = o.PresetAmounts
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.LookupKey) {
		toSerialize["lookup_key"] = o.LookupKey
	}
	if !IsNil(o.BillingScheme) {
		toSerialize["billing_scheme"] = o.BillingScheme
	}
	if !IsNil(o.TiersMode) {
		toSerialize["tiers_mode"] = o.TiersMode
	}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	if !IsNil(o.TransformQuantity) {
		toSerialize["transform_quantity"] = o.TransformQuantity
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePriceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_id",
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

	varCreatePriceRequest := _CreatePriceRequest{}

	err = json.Unmarshal(data, &varCreatePriceRequest)

	if err != nil {
		return err
	}

	*o = CreatePriceRequest(varCreatePriceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product_id")
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePriceRequest struct {
	value *CreatePriceRequest
	isSet bool
}

func (v NullableCreatePriceRequest) Get() *CreatePriceRequest {
	return v.value
}

func (v *NullableCreatePriceRequest) Set(val *CreatePriceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePriceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePriceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePriceRequest(val *CreatePriceRequest) *NullableCreatePriceRequest {
	return &NullableCreatePriceRequest{value: val, isSet: true}
}

func (v NullableCreatePriceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePriceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


