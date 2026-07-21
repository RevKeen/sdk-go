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

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription A subscription with full plan, billing, and status details.
type Subscription struct {
	// Subscription UUID
	Id string `json:"id"`
	// Owning merchant UUID
	MerchantId string `json:"merchant_id"`
	// Customer UUID
	CustomerId string `json:"customer_id"`
	// Product UUID
	ProductId NullableString `json:"product_id,omitempty"`
	// Price UUID
	PriceId NullableString `json:"price_id,omitempty"`
	// Plan UUID
	PlanId NullableString `json:"plan_id,omitempty"`
	// Subscription status
	Status string `json:"status"`
	// Subscribed quantity
	Quantity int32 `json:"quantity"`
	// Recurring amount in minor units
	AmountMinor int32 `json:"amount_minor"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Billing interval (day, week, month, year)
	BillingInterval string `json:"billing_interval"`
	// Current period start (ISO 8601)
	CurrentPeriodStart time.Time `json:"current_period_start"`
	// Current period end (ISO 8601)
	CurrentPeriodEnd time.Time `json:"current_period_end"`
	// Trial end (ISO 8601)
	TrialEnd NullableTime `json:"trial_end,omitempty"`
	// Cancellation timestamp (ISO 8601)
	CanceledAt NullableTime `json:"canceled_at,omitempty"`
	// Creation timestamp (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp (ISO 8601)
	UpdatedAt time.Time `json:"updated_at"`
	Dunning *SubscriptionDunning `json:"dunning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Subscription Subscription

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(id string, merchantId string, customerId string, status string, quantity int32, amountMinor int32, currency string, billingInterval string, currentPeriodStart time.Time, currentPeriodEnd time.Time, createdAt time.Time, updatedAt time.Time) *Subscription {
	this := Subscription{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.Status = status
	this.Quantity = quantity
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.BillingInterval = billingInterval
	this.CurrentPeriodStart = currentPeriodStart
	this.CurrentPeriodEnd = currentPeriodEnd
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetId returns the Id field value
func (o *Subscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Subscription) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *Subscription) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *Subscription) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *Subscription) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Subscription) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetProductId() string {
	if o == nil || IsNil(o.ProductId.Get()) {
		var ret string
		return ret
	}
	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// HasProductId returns a boolean if a field has been set.
func (o *Subscription) HasProductId() bool {
	if o != nil && o.ProductId.IsSet() {
		return true
	}

	return false
}

// SetProductId gets a reference to the given NullableString and assigns it to the ProductId field.
func (o *Subscription) SetProductId(v string) {
	o.ProductId.Set(&v)
}
// SetProductIdNil sets the value for ProductId to be an explicit nil
func (o *Subscription) SetProductIdNil() {
	o.ProductId.Set(nil)
}

// UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
func (o *Subscription) UnsetProductId() {
	o.ProductId.Unset()
}

// GetPriceId returns the PriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetPriceId() string {
	if o == nil || IsNil(o.PriceId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// HasPriceId returns a boolean if a field has been set.
func (o *Subscription) HasPriceId() bool {
	if o != nil && o.PriceId.IsSet() {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given NullableString and assigns it to the PriceId field.
func (o *Subscription) SetPriceId(v string) {
	o.PriceId.Set(&v)
}
// SetPriceIdNil sets the value for PriceId to be an explicit nil
func (o *Subscription) SetPriceIdNil() {
	o.PriceId.Set(nil)
}

// UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
func (o *Subscription) UnsetPriceId() {
	o.PriceId.Unset()
}

// GetPlanId returns the PlanId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetPlanId() string {
	if o == nil || IsNil(o.PlanId.Get()) {
		var ret string
		return ret
	}
	return *o.PlanId.Get()
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId.Get(), o.PlanId.IsSet()
}

// HasPlanId returns a boolean if a field has been set.
func (o *Subscription) HasPlanId() bool {
	if o != nil && o.PlanId.IsSet() {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given NullableString and assigns it to the PlanId field.
func (o *Subscription) SetPlanId(v string) {
	o.PlanId.Set(&v)
}
// SetPlanIdNil sets the value for PlanId to be an explicit nil
func (o *Subscription) SetPlanIdNil() {
	o.PlanId.Set(nil)
}

// UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil
func (o *Subscription) UnsetPlanId() {
	o.PlanId.Unset()
}

// GetStatus returns the Status field value
func (o *Subscription) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Subscription) SetStatus(v string) {
	o.Status = v
}

// GetQuantity returns the Quantity field value
func (o *Subscription) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Subscription) SetQuantity(v int32) {
	o.Quantity = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *Subscription) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *Subscription) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *Subscription) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Subscription) SetCurrency(v string) {
	o.Currency = v
}

// GetBillingInterval returns the BillingInterval field value
func (o *Subscription) GetBillingInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingInterval
}

// GetBillingIntervalOk returns a tuple with the BillingInterval field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetBillingIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingInterval, true
}

// SetBillingInterval sets field value
func (o *Subscription) SetBillingInterval(v string) {
	o.BillingInterval = v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
func (o *Subscription) GetCurrentPeriodStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodStart
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodStart, true
}

// SetCurrentPeriodStart sets field value
func (o *Subscription) SetCurrentPeriodStart(v time.Time) {
	o.CurrentPeriodStart = v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
func (o *Subscription) GetCurrentPeriodEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCurrentPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodEnd, true
}

// SetCurrentPeriodEnd sets field value
func (o *Subscription) SetCurrentPeriodEnd(v time.Time) {
	o.CurrentPeriodEnd = v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetTrialEnd() time.Time {
	if o == nil || IsNil(o.TrialEnd.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetTrialEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *Subscription) HasTrialEnd() bool {
	if o != nil && o.TrialEnd.IsSet() {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given NullableTime and assigns it to the TrialEnd field.
func (o *Subscription) SetTrialEnd(v time.Time) {
	o.TrialEnd.Set(&v)
}
// SetTrialEndNil sets the value for TrialEnd to be an explicit nil
func (o *Subscription) SetTrialEndNil() {
	o.TrialEnd.Set(nil)
}

// UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
func (o *Subscription) UnsetTrialEnd() {
	o.TrialEnd.Unset()
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Subscription) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Subscription) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *Subscription) HasCanceledAt() bool {
	if o != nil && o.CanceledAt.IsSet() {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given NullableTime and assigns it to the CanceledAt field.
func (o *Subscription) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}
// SetCanceledAtNil sets the value for CanceledAt to be an explicit nil
func (o *Subscription) SetCanceledAtNil() {
	o.CanceledAt.Set(nil)
}

// UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
func (o *Subscription) UnsetCanceledAt() {
	o.CanceledAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *Subscription) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Subscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Subscription) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Subscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDunning returns the Dunning field value if set, zero value otherwise.
func (o *Subscription) GetDunning() SubscriptionDunning {
	if o == nil || IsNil(o.Dunning) {
		var ret SubscriptionDunning
		return ret
	}
	return *o.Dunning
}

// GetDunningOk returns a tuple with the Dunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetDunningOk() (*SubscriptionDunning, bool) {
	if o == nil || IsNil(o.Dunning) {
		return nil, false
	}
	return o.Dunning, true
}

// HasDunning returns a boolean if a field has been set.
func (o *Subscription) HasDunning() bool {
	if o != nil && !IsNil(o.Dunning) {
		return true
	}

	return false
}

// SetDunning gets a reference to the given SubscriptionDunning and assigns it to the Dunning field.
func (o *Subscription) SetDunning(v SubscriptionDunning) {
	o.Dunning = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["customer_id"] = o.CustomerId
	if o.ProductId.IsSet() {
		toSerialize["product_id"] = o.ProductId.Get()
	}
	if o.PriceId.IsSet() {
		toSerialize["price_id"] = o.PriceId.Get()
	}
	if o.PlanId.IsSet() {
		toSerialize["plan_id"] = o.PlanId.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["quantity"] = o.Quantity
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["billing_interval"] = o.BillingInterval
	toSerialize["current_period_start"] = o.CurrentPeriodStart
	toSerialize["current_period_end"] = o.CurrentPeriodEnd
	if o.TrialEnd.IsSet() {
		toSerialize["trial_end"] = o.TrialEnd.Get()
	}
	if o.CanceledAt.IsSet() {
		toSerialize["canceled_at"] = o.CanceledAt.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Dunning) {
		toSerialize["dunning"] = o.Dunning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Subscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"customer_id",
		"status",
		"quantity",
		"amount_minor",
		"currency",
		"billing_interval",
		"current_period_start",
		"current_period_end",
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

	varSubscription := _Subscription{}

	err = json.Unmarshal(data, &varSubscription)

	if err != nil {
		return err
	}

	*o = Subscription(varSubscription)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "price_id")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "billing_interval")
		delete(additionalProperties, "current_period_start")
		delete(additionalProperties, "current_period_end")
		delete(additionalProperties, "trial_end")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "dunning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


