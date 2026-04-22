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

// checks if the Product type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Product{}

// Product A product defines what you sell — the name, billing type, price, and fulfillment method. Products can be one-time, recurring (subscription), or usage-based.
type Product struct {
	// Unique identifier (UUID)
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// User-facing product identifier
	ProductId string `json:"product_id"`
	// Product display name
	Name string `json:"name"`
	// Product description
	Description NullableString `json:"description,omitempty"`
	// Product kind
	Kind string `json:"kind"`
	// How the product is priced
	PricingModel string `json:"pricing_model"`
	// Price in minor units (cents/pence)
	AmountMinor NullableInt32 `json:"amount_minor"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Billing interval (day, week, month, year)
	Interval NullableString `json:"interval,omitempty"`
	// Number of intervals between billings
	IntervalCount NullableInt32 `json:"interval_count,omitempty"`
	// Free trial period in days
	TrialDays int32 `json:"trial_days"`
	// Fulfillment type
	FulfillmentType string `json:"fulfillment_type"`
	// Billing date calculation rule
	BillingAnchorRule NullableString `json:"billing_anchor_rule,omitempty"`
	// Day of month for billing (1-31)
	BillingAnchorDay NullableInt32 `json:"billing_anchor_day,omitempty"`
	// First payment timing
	FirstChargeBehavior NullableString `json:"first_charge_behavior,omitempty"`
	// Subscription end behavior
	EndBehavior NullableString `json:"end_behavior,omitempty"`
	// Max billing cycles for fixed-payment subscriptions
	MaxPayments NullableInt32 `json:"max_payments,omitempty"`
	// Associated usage meter ID
	UsageMeterId NullableString `json:"usage_meter_id,omitempty"`
	// URL-friendly slug
	Slug NullableString `json:"slug,omitempty"`
	// Whether the product is active
	IsActive bool `json:"is_active"`
	// Whether the product is archived
	IsArchived bool `json:"is_archived"`
	// Product image URL
	ImageUrl NullableString `json:"image_url,omitempty"`
	// Tax behavior (exclusive, inclusive, location)
	TaxBehavior NullableString `json:"tax_behavior,omitempty"`
	// Tax code for tax calculation
	TaxCode NullableString `json:"tax_code,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Creation timestamp (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp (ISO 8601)
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Product Product

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(id string, object string, productId string, name string, kind string, pricingModel string, amountMinor NullableInt32, currency string, trialDays int32, fulfillmentType string, isActive bool, isArchived bool, createdAt time.Time, updatedAt time.Time) *Product {
	this := Product{}
	this.Id = id
	this.Object = object
	this.ProductId = productId
	this.Name = name
	this.Kind = kind
	this.PricingModel = pricingModel
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.TrialDays = trialDays
	this.FulfillmentType = fulfillmentType
	this.IsActive = isActive
	this.IsArchived = isArchived
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value
func (o *Product) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Product) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Product) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Product) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Product) SetObject(v string) {
	o.Object = v
}

// GetProductId returns the ProductId field value
func (o *Product) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *Product) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *Product) SetProductId(v string) {
	o.ProductId = v
}

// GetName returns the Name field value
func (o *Product) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Product) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Product) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Product) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Product) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Product) UnsetDescription() {
	o.Description.Unset()
}

// GetKind returns the Kind field value
func (o *Product) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Product) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Product) SetKind(v string) {
	o.Kind = v
}

// GetPricingModel returns the PricingModel field value
func (o *Product) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *Product) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *Product) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetAmountMinor returns the AmountMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Product) GetAmountMinor() int32 {
	if o == nil || o.AmountMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AmountMinor.Get()
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountMinor.Get(), o.AmountMinor.IsSet()
}

// SetAmountMinor sets field value
func (o *Product) SetAmountMinor(v int32) {
	o.AmountMinor.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *Product) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Product) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Product) SetCurrency(v string) {
	o.Currency = v
}

// GetInterval returns the Interval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetInterval() string {
	if o == nil || IsNil(o.Interval.Get()) {
		var ret string
		return ret
	}
	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// HasInterval returns a boolean if a field has been set.
func (o *Product) HasInterval() bool {
	if o != nil && o.Interval.IsSet() {
		return true
	}

	return false
}

// SetInterval gets a reference to the given NullableString and assigns it to the Interval field.
func (o *Product) SetInterval(v string) {
	o.Interval.Set(&v)
}
// SetIntervalNil sets the value for Interval to be an explicit nil
func (o *Product) SetIntervalNil() {
	o.Interval.Set(nil)
}

// UnsetInterval ensures that no value is present for Interval, not even an explicit nil
func (o *Product) UnsetInterval() {
	o.Interval.Unset()
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *Product) HasIntervalCount() bool {
	if o != nil && o.IntervalCount.IsSet() {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given NullableInt32 and assigns it to the IntervalCount field.
func (o *Product) SetIntervalCount(v int32) {
	o.IntervalCount.Set(&v)
}
// SetIntervalCountNil sets the value for IntervalCount to be an explicit nil
func (o *Product) SetIntervalCountNil() {
	o.IntervalCount.Set(nil)
}

// UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
func (o *Product) UnsetIntervalCount() {
	o.IntervalCount.Unset()
}

// GetTrialDays returns the TrialDays field value
func (o *Product) GetTrialDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value
// and a boolean to check if the value has been set.
func (o *Product) GetTrialDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialDays, true
}

// SetTrialDays sets field value
func (o *Product) SetTrialDays(v int32) {
	o.TrialDays = v
}

// GetFulfillmentType returns the FulfillmentType field value
func (o *Product) GetFulfillmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentType
}

// GetFulfillmentTypeOk returns a tuple with the FulfillmentType field value
// and a boolean to check if the value has been set.
func (o *Product) GetFulfillmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentType, true
}

// SetFulfillmentType sets field value
func (o *Product) SetFulfillmentType(v string) {
	o.FulfillmentType = v
}

// GetBillingAnchorRule returns the BillingAnchorRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetBillingAnchorRule() string {
	if o == nil || IsNil(o.BillingAnchorRule.Get()) {
		var ret string
		return ret
	}
	return *o.BillingAnchorRule.Get()
}

// GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetBillingAnchorRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAnchorRule.Get(), o.BillingAnchorRule.IsSet()
}

// HasBillingAnchorRule returns a boolean if a field has been set.
func (o *Product) HasBillingAnchorRule() bool {
	if o != nil && o.BillingAnchorRule.IsSet() {
		return true
	}

	return false
}

// SetBillingAnchorRule gets a reference to the given NullableString and assigns it to the BillingAnchorRule field.
func (o *Product) SetBillingAnchorRule(v string) {
	o.BillingAnchorRule.Set(&v)
}
// SetBillingAnchorRuleNil sets the value for BillingAnchorRule to be an explicit nil
func (o *Product) SetBillingAnchorRuleNil() {
	o.BillingAnchorRule.Set(nil)
}

// UnsetBillingAnchorRule ensures that no value is present for BillingAnchorRule, not even an explicit nil
func (o *Product) UnsetBillingAnchorRule() {
	o.BillingAnchorRule.Unset()
}

// GetBillingAnchorDay returns the BillingAnchorDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetBillingAnchorDay() int32 {
	if o == nil || IsNil(o.BillingAnchorDay.Get()) {
		var ret int32
		return ret
	}
	return *o.BillingAnchorDay.Get()
}

// GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetBillingAnchorDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAnchorDay.Get(), o.BillingAnchorDay.IsSet()
}

// HasBillingAnchorDay returns a boolean if a field has been set.
func (o *Product) HasBillingAnchorDay() bool {
	if o != nil && o.BillingAnchorDay.IsSet() {
		return true
	}

	return false
}

// SetBillingAnchorDay gets a reference to the given NullableInt32 and assigns it to the BillingAnchorDay field.
func (o *Product) SetBillingAnchorDay(v int32) {
	o.BillingAnchorDay.Set(&v)
}
// SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil
func (o *Product) SetBillingAnchorDayNil() {
	o.BillingAnchorDay.Set(nil)
}

// UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
func (o *Product) UnsetBillingAnchorDay() {
	o.BillingAnchorDay.Unset()
}

// GetFirstChargeBehavior returns the FirstChargeBehavior field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetFirstChargeBehavior() string {
	if o == nil || IsNil(o.FirstChargeBehavior.Get()) {
		var ret string
		return ret
	}
	return *o.FirstChargeBehavior.Get()
}

// GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetFirstChargeBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstChargeBehavior.Get(), o.FirstChargeBehavior.IsSet()
}

// HasFirstChargeBehavior returns a boolean if a field has been set.
func (o *Product) HasFirstChargeBehavior() bool {
	if o != nil && o.FirstChargeBehavior.IsSet() {
		return true
	}

	return false
}

// SetFirstChargeBehavior gets a reference to the given NullableString and assigns it to the FirstChargeBehavior field.
func (o *Product) SetFirstChargeBehavior(v string) {
	o.FirstChargeBehavior.Set(&v)
}
// SetFirstChargeBehaviorNil sets the value for FirstChargeBehavior to be an explicit nil
func (o *Product) SetFirstChargeBehaviorNil() {
	o.FirstChargeBehavior.Set(nil)
}

// UnsetFirstChargeBehavior ensures that no value is present for FirstChargeBehavior, not even an explicit nil
func (o *Product) UnsetFirstChargeBehavior() {
	o.FirstChargeBehavior.Unset()
}

// GetEndBehavior returns the EndBehavior field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetEndBehavior() string {
	if o == nil || IsNil(o.EndBehavior.Get()) {
		var ret string
		return ret
	}
	return *o.EndBehavior.Get()
}

// GetEndBehaviorOk returns a tuple with the EndBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetEndBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndBehavior.Get(), o.EndBehavior.IsSet()
}

// HasEndBehavior returns a boolean if a field has been set.
func (o *Product) HasEndBehavior() bool {
	if o != nil && o.EndBehavior.IsSet() {
		return true
	}

	return false
}

// SetEndBehavior gets a reference to the given NullableString and assigns it to the EndBehavior field.
func (o *Product) SetEndBehavior(v string) {
	o.EndBehavior.Set(&v)
}
// SetEndBehaviorNil sets the value for EndBehavior to be an explicit nil
func (o *Product) SetEndBehaviorNil() {
	o.EndBehavior.Set(nil)
}

// UnsetEndBehavior ensures that no value is present for EndBehavior, not even an explicit nil
func (o *Product) UnsetEndBehavior() {
	o.EndBehavior.Unset()
}

// GetMaxPayments returns the MaxPayments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetMaxPayments() int32 {
	if o == nil || IsNil(o.MaxPayments.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPayments.Get()
}

// GetMaxPaymentsOk returns a tuple with the MaxPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetMaxPaymentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPayments.Get(), o.MaxPayments.IsSet()
}

// HasMaxPayments returns a boolean if a field has been set.
func (o *Product) HasMaxPayments() bool {
	if o != nil && o.MaxPayments.IsSet() {
		return true
	}

	return false
}

// SetMaxPayments gets a reference to the given NullableInt32 and assigns it to the MaxPayments field.
func (o *Product) SetMaxPayments(v int32) {
	o.MaxPayments.Set(&v)
}
// SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil
func (o *Product) SetMaxPaymentsNil() {
	o.MaxPayments.Set(nil)
}

// UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
func (o *Product) UnsetMaxPayments() {
	o.MaxPayments.Unset()
}

// GetUsageMeterId returns the UsageMeterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetUsageMeterId() string {
	if o == nil || IsNil(o.UsageMeterId.Get()) {
		var ret string
		return ret
	}
	return *o.UsageMeterId.Get()
}

// GetUsageMeterIdOk returns a tuple with the UsageMeterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetUsageMeterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageMeterId.Get(), o.UsageMeterId.IsSet()
}

// HasUsageMeterId returns a boolean if a field has been set.
func (o *Product) HasUsageMeterId() bool {
	if o != nil && o.UsageMeterId.IsSet() {
		return true
	}

	return false
}

// SetUsageMeterId gets a reference to the given NullableString and assigns it to the UsageMeterId field.
func (o *Product) SetUsageMeterId(v string) {
	o.UsageMeterId.Set(&v)
}
// SetUsageMeterIdNil sets the value for UsageMeterId to be an explicit nil
func (o *Product) SetUsageMeterIdNil() {
	o.UsageMeterId.Set(nil)
}

// UnsetUsageMeterId ensures that no value is present for UsageMeterId, not even an explicit nil
func (o *Product) UnsetUsageMeterId() {
	o.UsageMeterId.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *Product) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *Product) SetSlug(v string) {
	o.Slug.Set(&v)
}
// SetSlugNil sets the value for Slug to be an explicit nil
func (o *Product) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *Product) UnsetSlug() {
	o.Slug.Unset()
}

// GetIsActive returns the IsActive field value
func (o *Product) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *Product) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsArchived returns the IsArchived field value
func (o *Product) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *Product) GetIsArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *Product) SetIsArchived(v bool) {
	o.IsArchived = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Product) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Product) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Product) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Product) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetTaxBehavior returns the TaxBehavior field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetTaxBehavior() string {
	if o == nil || IsNil(o.TaxBehavior.Get()) {
		var ret string
		return ret
	}
	return *o.TaxBehavior.Get()
}

// GetTaxBehaviorOk returns a tuple with the TaxBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetTaxBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxBehavior.Get(), o.TaxBehavior.IsSet()
}

// HasTaxBehavior returns a boolean if a field has been set.
func (o *Product) HasTaxBehavior() bool {
	if o != nil && o.TaxBehavior.IsSet() {
		return true
	}

	return false
}

// SetTaxBehavior gets a reference to the given NullableString and assigns it to the TaxBehavior field.
func (o *Product) SetTaxBehavior(v string) {
	o.TaxBehavior.Set(&v)
}
// SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil
func (o *Product) SetTaxBehaviorNil() {
	o.TaxBehavior.Set(nil)
}

// UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil
func (o *Product) UnsetTaxBehavior() {
	o.TaxBehavior.Unset()
}

// GetTaxCode returns the TaxCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Product) GetTaxCode() string {
	if o == nil || IsNil(o.TaxCode.Get()) {
		var ret string
		return ret
	}
	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Product) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// HasTaxCode returns a boolean if a field has been set.
func (o *Product) HasTaxCode() bool {
	if o != nil && o.TaxCode.IsSet() {
		return true
	}

	return false
}

// SetTaxCode gets a reference to the given NullableString and assigns it to the TaxCode field.
func (o *Product) SetTaxCode(v string) {
	o.TaxCode.Set(&v)
}
// SetTaxCodeNil sets the value for TaxCode to be an explicit nil
func (o *Product) SetTaxCodeNil() {
	o.TaxCode.Set(nil)
}

// UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
func (o *Product) UnsetTaxCode() {
	o.TaxCode.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Product) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Product) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Product) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Product) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Product) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Product) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Product) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Product) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Product) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Product) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["product_id"] = o.ProductId
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["kind"] = o.Kind
	toSerialize["pricing_model"] = o.PricingModel
	toSerialize["amount_minor"] = o.AmountMinor.Get()
	toSerialize["currency"] = o.Currency
	if o.Interval.IsSet() {
		toSerialize["interval"] = o.Interval.Get()
	}
	if o.IntervalCount.IsSet() {
		toSerialize["interval_count"] = o.IntervalCount.Get()
	}
	toSerialize["trial_days"] = o.TrialDays
	toSerialize["fulfillment_type"] = o.FulfillmentType
	if o.BillingAnchorRule.IsSet() {
		toSerialize["billing_anchor_rule"] = o.BillingAnchorRule.Get()
	}
	if o.BillingAnchorDay.IsSet() {
		toSerialize["billing_anchor_day"] = o.BillingAnchorDay.Get()
	}
	if o.FirstChargeBehavior.IsSet() {
		toSerialize["first_charge_behavior"] = o.FirstChargeBehavior.Get()
	}
	if o.EndBehavior.IsSet() {
		toSerialize["end_behavior"] = o.EndBehavior.Get()
	}
	if o.MaxPayments.IsSet() {
		toSerialize["max_payments"] = o.MaxPayments.Get()
	}
	if o.UsageMeterId.IsSet() {
		toSerialize["usage_meter_id"] = o.UsageMeterId.Get()
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	toSerialize["is_active"] = o.IsActive
	toSerialize["is_archived"] = o.IsArchived
	if o.ImageUrl.IsSet() {
		toSerialize["image_url"] = o.ImageUrl.Get()
	}
	if o.TaxBehavior.IsSet() {
		toSerialize["tax_behavior"] = o.TaxBehavior.Get()
	}
	if o.TaxCode.IsSet() {
		toSerialize["tax_code"] = o.TaxCode.Get()
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

func (o *Product) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"product_id",
		"name",
		"kind",
		"pricing_model",
		"amount_minor",
		"currency",
		"trial_days",
		"fulfillment_type",
		"is_active",
		"is_archived",
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

	varProduct := _Product{}

	err = json.Unmarshal(data, &varProduct)

	if err != nil {
		return err
	}

	*o = Product(varProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "pricing_model")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_count")
		delete(additionalProperties, "trial_days")
		delete(additionalProperties, "fulfillment_type")
		delete(additionalProperties, "billing_anchor_rule")
		delete(additionalProperties, "billing_anchor_day")
		delete(additionalProperties, "first_charge_behavior")
		delete(additionalProperties, "end_behavior")
		delete(additionalProperties, "max_payments")
		delete(additionalProperties, "usage_meter_id")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_archived")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "tax_behavior")
		delete(additionalProperties, "tax_code")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


