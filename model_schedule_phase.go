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
	"fmt"
)

// checks if the SchedulePhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePhase{}

// SchedulePhase A time-bounded phase within a subscription schedule, with its own pricing, quantities, and billing configuration.
type SchedulePhase struct {
	// Products/prices included in this phase
	Items []PhaseItem `json:"items"`
	// When this phase starts (ISO 8601 date)
	StartDate string `json:"start_date"`
	// When this phase ends (null for indefinite)
	EndDate NullableString `json:"end_date,omitempty"`
	// How to handle prorations when entering this phase
	ProrationBehavior *string `json:"proration_behavior,omitempty"`
	// Override billing cycle anchor for this phase
	BillingCycleAnchor NullableString `json:"billing_cycle_anchor,omitempty"`
	// Payment method to use for this phase
	DefaultPaymentMethod NullableString `json:"default_payment_method,omitempty"`
	// How to collect payment for this phase
	CollectionMethod *string `json:"collection_method,omitempty"`
	// Coupon code to apply during this phase
	Coupon NullableString `json:"coupon,omitempty"`
	// End of trial period (if applicable)
	TrialEnd NullableString `json:"trial_end,omitempty"`
	// Phase-specific metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SchedulePhase SchedulePhase

// NewSchedulePhase instantiates a new SchedulePhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePhase(items []PhaseItem, startDate string) *SchedulePhase {
	this := SchedulePhase{}
	this.Items = items
	this.StartDate = startDate
	return &this
}

// NewSchedulePhaseWithDefaults instantiates a new SchedulePhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePhaseWithDefaults() *SchedulePhase {
	this := SchedulePhase{}
	return &this
}

// GetItems returns the Items field value
func (o *SchedulePhase) GetItems() []PhaseItem {
	if o == nil {
		var ret []PhaseItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SchedulePhase) GetItemsOk() ([]PhaseItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SchedulePhase) SetItems(v []PhaseItem) {
	o.Items = v
}

// GetStartDate returns the StartDate field value
func (o *SchedulePhase) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *SchedulePhase) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *SchedulePhase) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulePhase) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulePhase) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *SchedulePhase) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *SchedulePhase) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *SchedulePhase) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *SchedulePhase) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetProrationBehavior returns the ProrationBehavior field value if set, zero value otherwise.
func (o *SchedulePhase) GetProrationBehavior() string {
	if o == nil || IsNil(o.ProrationBehavior) {
		var ret string
		return ret
	}
	return *o.ProrationBehavior
}

// GetProrationBehaviorOk returns a tuple with the ProrationBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePhase) GetProrationBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.ProrationBehavior) {
		return nil, false
	}
	return o.ProrationBehavior, true
}

// HasProrationBehavior returns a boolean if a field has been set.
func (o *SchedulePhase) HasProrationBehavior() bool {
	if o != nil && !IsNil(o.ProrationBehavior) {
		return true
	}

	return false
}

// SetProrationBehavior gets a reference to the given string and assigns it to the ProrationBehavior field.
func (o *SchedulePhase) SetProrationBehavior(v string) {
	o.ProrationBehavior = &v
}

// GetBillingCycleAnchor returns the BillingCycleAnchor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulePhase) GetBillingCycleAnchor() string {
	if o == nil || IsNil(o.BillingCycleAnchor.Get()) {
		var ret string
		return ret
	}
	return *o.BillingCycleAnchor.Get()
}

// GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulePhase) GetBillingCycleAnchorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingCycleAnchor.Get(), o.BillingCycleAnchor.IsSet()
}

// HasBillingCycleAnchor returns a boolean if a field has been set.
func (o *SchedulePhase) HasBillingCycleAnchor() bool {
	if o != nil && o.BillingCycleAnchor.IsSet() {
		return true
	}

	return false
}

// SetBillingCycleAnchor gets a reference to the given NullableString and assigns it to the BillingCycleAnchor field.
func (o *SchedulePhase) SetBillingCycleAnchor(v string) {
	o.BillingCycleAnchor.Set(&v)
}
// SetBillingCycleAnchorNil sets the value for BillingCycleAnchor to be an explicit nil
func (o *SchedulePhase) SetBillingCycleAnchorNil() {
	o.BillingCycleAnchor.Set(nil)
}

// UnsetBillingCycleAnchor ensures that no value is present for BillingCycleAnchor, not even an explicit nil
func (o *SchedulePhase) UnsetBillingCycleAnchor() {
	o.BillingCycleAnchor.Unset()
}

// GetDefaultPaymentMethod returns the DefaultPaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulePhase) GetDefaultPaymentMethod() string {
	if o == nil || IsNil(o.DefaultPaymentMethod.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultPaymentMethod.Get()
}

// GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulePhase) GetDefaultPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPaymentMethod.Get(), o.DefaultPaymentMethod.IsSet()
}

// HasDefaultPaymentMethod returns a boolean if a field has been set.
func (o *SchedulePhase) HasDefaultPaymentMethod() bool {
	if o != nil && o.DefaultPaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetDefaultPaymentMethod gets a reference to the given NullableString and assigns it to the DefaultPaymentMethod field.
func (o *SchedulePhase) SetDefaultPaymentMethod(v string) {
	o.DefaultPaymentMethod.Set(&v)
}
// SetDefaultPaymentMethodNil sets the value for DefaultPaymentMethod to be an explicit nil
func (o *SchedulePhase) SetDefaultPaymentMethodNil() {
	o.DefaultPaymentMethod.Set(nil)
}

// UnsetDefaultPaymentMethod ensures that no value is present for DefaultPaymentMethod, not even an explicit nil
func (o *SchedulePhase) UnsetDefaultPaymentMethod() {
	o.DefaultPaymentMethod.Unset()
}

// GetCollectionMethod returns the CollectionMethod field value if set, zero value otherwise.
func (o *SchedulePhase) GetCollectionMethod() string {
	if o == nil || IsNil(o.CollectionMethod) {
		var ret string
		return ret
	}
	return *o.CollectionMethod
}

// GetCollectionMethodOk returns a tuple with the CollectionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePhase) GetCollectionMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionMethod) {
		return nil, false
	}
	return o.CollectionMethod, true
}

// HasCollectionMethod returns a boolean if a field has been set.
func (o *SchedulePhase) HasCollectionMethod() bool {
	if o != nil && !IsNil(o.CollectionMethod) {
		return true
	}

	return false
}

// SetCollectionMethod gets a reference to the given string and assigns it to the CollectionMethod field.
func (o *SchedulePhase) SetCollectionMethod(v string) {
	o.CollectionMethod = &v
}

// GetCoupon returns the Coupon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulePhase) GetCoupon() string {
	if o == nil || IsNil(o.Coupon.Get()) {
		var ret string
		return ret
	}
	return *o.Coupon.Get()
}

// GetCouponOk returns a tuple with the Coupon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulePhase) GetCouponOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Coupon.Get(), o.Coupon.IsSet()
}

// HasCoupon returns a boolean if a field has been set.
func (o *SchedulePhase) HasCoupon() bool {
	if o != nil && o.Coupon.IsSet() {
		return true
	}

	return false
}

// SetCoupon gets a reference to the given NullableString and assigns it to the Coupon field.
func (o *SchedulePhase) SetCoupon(v string) {
	o.Coupon.Set(&v)
}
// SetCouponNil sets the value for Coupon to be an explicit nil
func (o *SchedulePhase) SetCouponNil() {
	o.Coupon.Set(nil)
}

// UnsetCoupon ensures that no value is present for Coupon, not even an explicit nil
func (o *SchedulePhase) UnsetCoupon() {
	o.Coupon.Unset()
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchedulePhase) GetTrialEnd() string {
	if o == nil || IsNil(o.TrialEnd.Get()) {
		var ret string
		return ret
	}
	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchedulePhase) GetTrialEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SchedulePhase) HasTrialEnd() bool {
	if o != nil && o.TrialEnd.IsSet() {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given NullableString and assigns it to the TrialEnd field.
func (o *SchedulePhase) SetTrialEnd(v string) {
	o.TrialEnd.Set(&v)
}
// SetTrialEndNil sets the value for TrialEnd to be an explicit nil
func (o *SchedulePhase) SetTrialEndNil() {
	o.TrialEnd.Set(nil)
}

// UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
func (o *SchedulePhase) UnsetTrialEnd() {
	o.TrialEnd.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SchedulePhase) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePhase) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SchedulePhase) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SchedulePhase) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SchedulePhase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["start_date"] = o.StartDate
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if !IsNil(o.ProrationBehavior) {
		toSerialize["proration_behavior"] = o.ProrationBehavior
	}
	if o.BillingCycleAnchor.IsSet() {
		toSerialize["billing_cycle_anchor"] = o.BillingCycleAnchor.Get()
	}
	if o.DefaultPaymentMethod.IsSet() {
		toSerialize["default_payment_method"] = o.DefaultPaymentMethod.Get()
	}
	if !IsNil(o.CollectionMethod) {
		toSerialize["collection_method"] = o.CollectionMethod
	}
	if o.Coupon.IsSet() {
		toSerialize["coupon"] = o.Coupon.Get()
	}
	if o.TrialEnd.IsSet() {
		toSerialize["trial_end"] = o.TrialEnd.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SchedulePhase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"start_date",
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

	varSchedulePhase := _SchedulePhase{}

	err = json.Unmarshal(data, &varSchedulePhase)

	if err != nil {
		return err
	}

	*o = SchedulePhase(varSchedulePhase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "start_date")
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "proration_behavior")
		delete(additionalProperties, "billing_cycle_anchor")
		delete(additionalProperties, "default_payment_method")
		delete(additionalProperties, "collection_method")
		delete(additionalProperties, "coupon")
		delete(additionalProperties, "trial_end")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSchedulePhase struct {
	value *SchedulePhase
	isSet bool
}

func (v NullableSchedulePhase) Get() *SchedulePhase {
	return v.value
}

func (v *NullableSchedulePhase) Set(val *SchedulePhase) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePhase) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePhase(val *SchedulePhase) *NullableSchedulePhase {
	return &NullableSchedulePhase{value: val, isSet: true}
}

func (v NullableSchedulePhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


