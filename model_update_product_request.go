/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-01-15` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-01-15
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
)

// checks if the UpdateProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductRequest{}

// UpdateProductRequest Parameters for updating an existing product. Only the provided fields are changed.
type UpdateProductRequest struct {
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	AmountMinor *int32 `json:"amount_minor,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Interval NullableString `json:"interval,omitempty"`
	IntervalCount NullableInt32 `json:"interval_count,omitempty"`
	TrialDays *int32 `json:"trial_days,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	// Fulfillment type for the product.
	FulfillmentType *string `json:"fulfillment_type,omitempty"`
	// How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month.
	BillingAnchorRule *string `json:"billing_anchor_rule,omitempty"`
	BillingAnchorDay NullableInt32 `json:"billing_anchor_day,omitempty"`
	// When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date.
	FirstChargeBehavior *string `json:"first_charge_behavior,omitempty"`
	// How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles.
	EndBehavior *string `json:"end_behavior,omitempty"`
	MaxPayments NullableInt32 `json:"max_payments,omitempty"`
	// Arbitrary key-value metadata for the product
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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
func (o *UpdateProductRequest) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetAmountMinorOk() (*int32, bool) {
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

// SetAmountMinor gets a reference to the given int32 and assigns it to the AmountMinor field.
func (o *UpdateProductRequest) SetAmountMinor(v int32) {
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
func (o *UpdateProductRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
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

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateProductRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UpdateProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
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
	if !IsNil(o.FulfillmentType) {
		toSerialize["fulfillment_type"] = o.FulfillmentType
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
		delete(additionalProperties, "fulfillment_type")
		delete(additionalProperties, "billing_anchor_rule")
		delete(additionalProperties, "billing_anchor_day")
		delete(additionalProperties, "first_charge_behavior")
		delete(additionalProperties, "end_behavior")
		delete(additionalProperties, "max_payments")
		delete(additionalProperties, "metadata")
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


