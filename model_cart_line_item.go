/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the CartLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CartLineItem{}

// CartLineItem struct for CartLineItem
type CartLineItem struct {
	Id string `json:"id"`
	ProductId string `json:"product_id"`
	Name string `json:"name"`
	Quantity int32 `json:"quantity"`
	UnitPriceMinor int32 `json:"unit_price_minor"`
	Currency string `json:"currency"`
	Recurring CartLineItemRecurring `json:"recurring"`
	BillingMaxCycles NullableInt32 `json:"billing_max_cycles,omitempty"`
	TrialPeriodDays NullableInt32 `json:"trial_period_days,omitempty"`
	StartRule NullableCartStartRule `json:"start_rule,omitempty"`
	BillingAnchorRule NullableCartBillingAnchorRule `json:"billing_anchor_rule,omitempty"`
	BillingAnchorDay NullableInt32 `json:"billing_anchor_day,omitempty"`
	DueTodayMinor NullableInt32 `json:"due_today_minor,omitempty"`
	FirstChargeMinor NullableInt32 `json:"first_charge_minor,omitempty"`
	FirstRenewalAt NullableTime `json:"first_renewal_at,omitempty"`
	EffectiveStartRule NullableCartStartRule `json:"effective_start_rule,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CartLineItem CartLineItem

// NewCartLineItem instantiates a new CartLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartLineItem(id string, productId string, name string, quantity int32, unitPriceMinor int32, currency string, recurring CartLineItemRecurring) *CartLineItem {
	this := CartLineItem{}
	this.Id = id
	this.ProductId = productId
	this.Name = name
	this.Quantity = quantity
	this.UnitPriceMinor = unitPriceMinor
	this.Currency = currency
	this.Recurring = recurring
	return &this
}

// NewCartLineItemWithDefaults instantiates a new CartLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartLineItemWithDefaults() *CartLineItem {
	this := CartLineItem{}
	return &this
}

// GetId returns the Id field value
func (o *CartLineItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CartLineItem) SetId(v string) {
	o.Id = v
}

// GetProductId returns the ProductId field value
func (o *CartLineItem) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *CartLineItem) SetProductId(v string) {
	o.ProductId = v
}

// GetName returns the Name field value
func (o *CartLineItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CartLineItem) SetName(v string) {
	o.Name = v
}

// GetQuantity returns the Quantity field value
func (o *CartLineItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *CartLineItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetUnitPriceMinor returns the UnitPriceMinor field value
func (o *CartLineItem) GetUnitPriceMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitPriceMinor
}

// GetUnitPriceMinorOk returns a tuple with the UnitPriceMinor field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetUnitPriceMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitPriceMinor, true
}

// SetUnitPriceMinor sets field value
func (o *CartLineItem) SetUnitPriceMinor(v int32) {
	o.UnitPriceMinor = v
}

// GetCurrency returns the Currency field value
func (o *CartLineItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CartLineItem) SetCurrency(v string) {
	o.Currency = v
}

// GetRecurring returns the Recurring field value
func (o *CartLineItem) GetRecurring() CartLineItemRecurring {
	if o == nil {
		var ret CartLineItemRecurring
		return ret
	}

	return o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetRecurringOk() (*CartLineItemRecurring, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurring, true
}

// SetRecurring sets field value
func (o *CartLineItem) SetRecurring(v CartLineItemRecurring) {
	o.Recurring = v
}

// GetBillingMaxCycles returns the BillingMaxCycles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetBillingMaxCycles() int32 {
	if o == nil || IsNil(o.BillingMaxCycles.Get()) {
		var ret int32
		return ret
	}
	return *o.BillingMaxCycles.Get()
}

// GetBillingMaxCyclesOk returns a tuple with the BillingMaxCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetBillingMaxCyclesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingMaxCycles.Get(), o.BillingMaxCycles.IsSet()
}

// HasBillingMaxCycles returns a boolean if a field has been set.
func (o *CartLineItem) HasBillingMaxCycles() bool {
	if o != nil && o.BillingMaxCycles.IsSet() {
		return true
	}

	return false
}

// SetBillingMaxCycles gets a reference to the given NullableInt32 and assigns it to the BillingMaxCycles field.
func (o *CartLineItem) SetBillingMaxCycles(v int32) {
	o.BillingMaxCycles.Set(&v)
}
// SetBillingMaxCyclesNil sets the value for BillingMaxCycles to be an explicit nil
func (o *CartLineItem) SetBillingMaxCyclesNil() {
	o.BillingMaxCycles.Set(nil)
}

// UnsetBillingMaxCycles ensures that no value is present for BillingMaxCycles, not even an explicit nil
func (o *CartLineItem) UnsetBillingMaxCycles() {
	o.BillingMaxCycles.Unset()
}

// GetTrialPeriodDays returns the TrialPeriodDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetTrialPeriodDays() int32 {
	if o == nil || IsNil(o.TrialPeriodDays.Get()) {
		var ret int32
		return ret
	}
	return *o.TrialPeriodDays.Get()
}

// GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetTrialPeriodDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialPeriodDays.Get(), o.TrialPeriodDays.IsSet()
}

// HasTrialPeriodDays returns a boolean if a field has been set.
func (o *CartLineItem) HasTrialPeriodDays() bool {
	if o != nil && o.TrialPeriodDays.IsSet() {
		return true
	}

	return false
}

// SetTrialPeriodDays gets a reference to the given NullableInt32 and assigns it to the TrialPeriodDays field.
func (o *CartLineItem) SetTrialPeriodDays(v int32) {
	o.TrialPeriodDays.Set(&v)
}
// SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil
func (o *CartLineItem) SetTrialPeriodDaysNil() {
	o.TrialPeriodDays.Set(nil)
}

// UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
func (o *CartLineItem) UnsetTrialPeriodDays() {
	o.TrialPeriodDays.Unset()
}

// GetStartRule returns the StartRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetStartRule() CartStartRule {
	if o == nil || IsNil(o.StartRule.Get()) {
		var ret CartStartRule
		return ret
	}
	return *o.StartRule.Get()
}

// GetStartRuleOk returns a tuple with the StartRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetStartRuleOk() (*CartStartRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartRule.Get(), o.StartRule.IsSet()
}

// HasStartRule returns a boolean if a field has been set.
func (o *CartLineItem) HasStartRule() bool {
	if o != nil && o.StartRule.IsSet() {
		return true
	}

	return false
}

// SetStartRule gets a reference to the given NullableCartStartRule and assigns it to the StartRule field.
func (o *CartLineItem) SetStartRule(v CartStartRule) {
	o.StartRule.Set(&v)
}
// SetStartRuleNil sets the value for StartRule to be an explicit nil
func (o *CartLineItem) SetStartRuleNil() {
	o.StartRule.Set(nil)
}

// UnsetStartRule ensures that no value is present for StartRule, not even an explicit nil
func (o *CartLineItem) UnsetStartRule() {
	o.StartRule.Unset()
}

// GetBillingAnchorRule returns the BillingAnchorRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetBillingAnchorRule() CartBillingAnchorRule {
	if o == nil || IsNil(o.BillingAnchorRule.Get()) {
		var ret CartBillingAnchorRule
		return ret
	}
	return *o.BillingAnchorRule.Get()
}

// GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetBillingAnchorRuleOk() (*CartBillingAnchorRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAnchorRule.Get(), o.BillingAnchorRule.IsSet()
}

// HasBillingAnchorRule returns a boolean if a field has been set.
func (o *CartLineItem) HasBillingAnchorRule() bool {
	if o != nil && o.BillingAnchorRule.IsSet() {
		return true
	}

	return false
}

// SetBillingAnchorRule gets a reference to the given NullableCartBillingAnchorRule and assigns it to the BillingAnchorRule field.
func (o *CartLineItem) SetBillingAnchorRule(v CartBillingAnchorRule) {
	o.BillingAnchorRule.Set(&v)
}
// SetBillingAnchorRuleNil sets the value for BillingAnchorRule to be an explicit nil
func (o *CartLineItem) SetBillingAnchorRuleNil() {
	o.BillingAnchorRule.Set(nil)
}

// UnsetBillingAnchorRule ensures that no value is present for BillingAnchorRule, not even an explicit nil
func (o *CartLineItem) UnsetBillingAnchorRule() {
	o.BillingAnchorRule.Unset()
}

// GetBillingAnchorDay returns the BillingAnchorDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetBillingAnchorDay() int32 {
	if o == nil || IsNil(o.BillingAnchorDay.Get()) {
		var ret int32
		return ret
	}
	return *o.BillingAnchorDay.Get()
}

// GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetBillingAnchorDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAnchorDay.Get(), o.BillingAnchorDay.IsSet()
}

// HasBillingAnchorDay returns a boolean if a field has been set.
func (o *CartLineItem) HasBillingAnchorDay() bool {
	if o != nil && o.BillingAnchorDay.IsSet() {
		return true
	}

	return false
}

// SetBillingAnchorDay gets a reference to the given NullableInt32 and assigns it to the BillingAnchorDay field.
func (o *CartLineItem) SetBillingAnchorDay(v int32) {
	o.BillingAnchorDay.Set(&v)
}
// SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil
func (o *CartLineItem) SetBillingAnchorDayNil() {
	o.BillingAnchorDay.Set(nil)
}

// UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
func (o *CartLineItem) UnsetBillingAnchorDay() {
	o.BillingAnchorDay.Unset()
}

// GetDueTodayMinor returns the DueTodayMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetDueTodayMinor() int32 {
	if o == nil || IsNil(o.DueTodayMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.DueTodayMinor.Get()
}

// GetDueTodayMinorOk returns a tuple with the DueTodayMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetDueTodayMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueTodayMinor.Get(), o.DueTodayMinor.IsSet()
}

// HasDueTodayMinor returns a boolean if a field has been set.
func (o *CartLineItem) HasDueTodayMinor() bool {
	if o != nil && o.DueTodayMinor.IsSet() {
		return true
	}

	return false
}

// SetDueTodayMinor gets a reference to the given NullableInt32 and assigns it to the DueTodayMinor field.
func (o *CartLineItem) SetDueTodayMinor(v int32) {
	o.DueTodayMinor.Set(&v)
}
// SetDueTodayMinorNil sets the value for DueTodayMinor to be an explicit nil
func (o *CartLineItem) SetDueTodayMinorNil() {
	o.DueTodayMinor.Set(nil)
}

// UnsetDueTodayMinor ensures that no value is present for DueTodayMinor, not even an explicit nil
func (o *CartLineItem) UnsetDueTodayMinor() {
	o.DueTodayMinor.Unset()
}

// GetFirstChargeMinor returns the FirstChargeMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetFirstChargeMinor() int32 {
	if o == nil || IsNil(o.FirstChargeMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.FirstChargeMinor.Get()
}

// GetFirstChargeMinorOk returns a tuple with the FirstChargeMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetFirstChargeMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstChargeMinor.Get(), o.FirstChargeMinor.IsSet()
}

// HasFirstChargeMinor returns a boolean if a field has been set.
func (o *CartLineItem) HasFirstChargeMinor() bool {
	if o != nil && o.FirstChargeMinor.IsSet() {
		return true
	}

	return false
}

// SetFirstChargeMinor gets a reference to the given NullableInt32 and assigns it to the FirstChargeMinor field.
func (o *CartLineItem) SetFirstChargeMinor(v int32) {
	o.FirstChargeMinor.Set(&v)
}
// SetFirstChargeMinorNil sets the value for FirstChargeMinor to be an explicit nil
func (o *CartLineItem) SetFirstChargeMinorNil() {
	o.FirstChargeMinor.Set(nil)
}

// UnsetFirstChargeMinor ensures that no value is present for FirstChargeMinor, not even an explicit nil
func (o *CartLineItem) UnsetFirstChargeMinor() {
	o.FirstChargeMinor.Unset()
}

// GetFirstRenewalAt returns the FirstRenewalAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetFirstRenewalAt() time.Time {
	if o == nil || IsNil(o.FirstRenewalAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.FirstRenewalAt.Get()
}

// GetFirstRenewalAtOk returns a tuple with the FirstRenewalAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetFirstRenewalAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstRenewalAt.Get(), o.FirstRenewalAt.IsSet()
}

// HasFirstRenewalAt returns a boolean if a field has been set.
func (o *CartLineItem) HasFirstRenewalAt() bool {
	if o != nil && o.FirstRenewalAt.IsSet() {
		return true
	}

	return false
}

// SetFirstRenewalAt gets a reference to the given NullableTime and assigns it to the FirstRenewalAt field.
func (o *CartLineItem) SetFirstRenewalAt(v time.Time) {
	o.FirstRenewalAt.Set(&v)
}
// SetFirstRenewalAtNil sets the value for FirstRenewalAt to be an explicit nil
func (o *CartLineItem) SetFirstRenewalAtNil() {
	o.FirstRenewalAt.Set(nil)
}

// UnsetFirstRenewalAt ensures that no value is present for FirstRenewalAt, not even an explicit nil
func (o *CartLineItem) UnsetFirstRenewalAt() {
	o.FirstRenewalAt.Unset()
}

// GetEffectiveStartRule returns the EffectiveStartRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CartLineItem) GetEffectiveStartRule() CartStartRule {
	if o == nil || IsNil(o.EffectiveStartRule.Get()) {
		var ret CartStartRule
		return ret
	}
	return *o.EffectiveStartRule.Get()
}

// GetEffectiveStartRuleOk returns a tuple with the EffectiveStartRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CartLineItem) GetEffectiveStartRuleOk() (*CartStartRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectiveStartRule.Get(), o.EffectiveStartRule.IsSet()
}

// HasEffectiveStartRule returns a boolean if a field has been set.
func (o *CartLineItem) HasEffectiveStartRule() bool {
	if o != nil && o.EffectiveStartRule.IsSet() {
		return true
	}

	return false
}

// SetEffectiveStartRule gets a reference to the given NullableCartStartRule and assigns it to the EffectiveStartRule field.
func (o *CartLineItem) SetEffectiveStartRule(v CartStartRule) {
	o.EffectiveStartRule.Set(&v)
}
// SetEffectiveStartRuleNil sets the value for EffectiveStartRule to be an explicit nil
func (o *CartLineItem) SetEffectiveStartRuleNil() {
	o.EffectiveStartRule.Set(nil)
}

// UnsetEffectiveStartRule ensures that no value is present for EffectiveStartRule, not even an explicit nil
func (o *CartLineItem) UnsetEffectiveStartRule() {
	o.EffectiveStartRule.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CartLineItem) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartLineItem) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CartLineItem) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CartLineItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CartLineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["product_id"] = o.ProductId
	toSerialize["name"] = o.Name
	toSerialize["quantity"] = o.Quantity
	toSerialize["unit_price_minor"] = o.UnitPriceMinor
	toSerialize["currency"] = o.Currency
	toSerialize["recurring"] = o.Recurring
	if o.BillingMaxCycles.IsSet() {
		toSerialize["billing_max_cycles"] = o.BillingMaxCycles.Get()
	}
	if o.TrialPeriodDays.IsSet() {
		toSerialize["trial_period_days"] = o.TrialPeriodDays.Get()
	}
	if o.StartRule.IsSet() {
		toSerialize["start_rule"] = o.StartRule.Get()
	}
	if o.BillingAnchorRule.IsSet() {
		toSerialize["billing_anchor_rule"] = o.BillingAnchorRule.Get()
	}
	if o.BillingAnchorDay.IsSet() {
		toSerialize["billing_anchor_day"] = o.BillingAnchorDay.Get()
	}
	if o.DueTodayMinor.IsSet() {
		toSerialize["due_today_minor"] = o.DueTodayMinor.Get()
	}
	if o.FirstChargeMinor.IsSet() {
		toSerialize["first_charge_minor"] = o.FirstChargeMinor.Get()
	}
	if o.FirstRenewalAt.IsSet() {
		toSerialize["first_renewal_at"] = o.FirstRenewalAt.Get()
	}
	if o.EffectiveStartRule.IsSet() {
		toSerialize["effective_start_rule"] = o.EffectiveStartRule.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CartLineItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"product_id",
		"name",
		"quantity",
		"unit_price_minor",
		"currency",
		"recurring",
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

	varCartLineItem := _CartLineItem{}

	err = json.Unmarshal(data, &varCartLineItem)

	if err != nil {
		return err
	}

	*o = CartLineItem(varCartLineItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unit_price_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "recurring")
		delete(additionalProperties, "billing_max_cycles")
		delete(additionalProperties, "trial_period_days")
		delete(additionalProperties, "start_rule")
		delete(additionalProperties, "billing_anchor_rule")
		delete(additionalProperties, "billing_anchor_day")
		delete(additionalProperties, "due_today_minor")
		delete(additionalProperties, "first_charge_minor")
		delete(additionalProperties, "first_renewal_at")
		delete(additionalProperties, "effective_start_rule")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCartLineItem struct {
	value *CartLineItem
	isSet bool
}

func (v NullableCartLineItem) Get() *CartLineItem {
	return v.value
}

func (v *NullableCartLineItem) Set(val *CartLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCartLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCartLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartLineItem(val *CartLineItem) *NullableCartLineItem {
	return &NullableCartLineItem{value: val, isSet: true}
}

func (v NullableCartLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


