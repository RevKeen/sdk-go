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
	"fmt"
)

// checks if the SubscriptionSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionSchedule{}

// SubscriptionSchedule A subscription schedule defines future changes to a subscription. Each phase specifies different pricing, quantities, or plans that take effect at specified dates.
type SubscriptionSchedule struct {
	// Public schedule ID (sub_sched_xxx)
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Schedule status. not_started: Scheduled for future. active: Currently executing phases. completed: All phases finished. canceled: Manually canceled. released: Detached from subscription.
	Status string `json:"status"`
	// Customer ID
	Customer NullableString `json:"customer"`
	// Linked subscription ID
	Subscription NullableString `json:"subscription"`
	CurrentPhase SubscriptionScheduleCurrentPhase `json:"current_phase"`
	// All phases in the schedule
	Phases []SchedulePhase `json:"phases"`
	// What happens when the schedule completes. cancel: Cancel the subscription. release: Make subscription standalone.
	EndBehavior string `json:"end_behavior"`
	// When released (Unix timestamp)
	ReleasedAt NullableInt32 `json:"released_at,omitempty"`
	// Subscription ID at release time
	ReleasedSubscription NullableString `json:"released_subscription,omitempty"`
	// Custom metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Unix timestamp of creation
	Created int32 `json:"created"`
	// Whether in live mode
	Livemode bool `json:"livemode"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionSchedule SubscriptionSchedule

// NewSubscriptionSchedule instantiates a new SubscriptionSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionSchedule(id string, object string, status string, customer NullableString, subscription NullableString, currentPhase SubscriptionScheduleCurrentPhase, phases []SchedulePhase, endBehavior string, created int32, livemode bool) *SubscriptionSchedule {
	this := SubscriptionSchedule{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.Customer = customer
	this.Subscription = subscription
	this.CurrentPhase = currentPhase
	this.Phases = phases
	this.EndBehavior = endBehavior
	this.Created = created
	this.Livemode = livemode
	return &this
}

// NewSubscriptionScheduleWithDefaults instantiates a new SubscriptionSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionScheduleWithDefaults() *SubscriptionSchedule {
	this := SubscriptionSchedule{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionSchedule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionSchedule) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *SubscriptionSchedule) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SubscriptionSchedule) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *SubscriptionSchedule) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionSchedule) SetStatus(v string) {
	o.Status = v
}

// GetCustomer returns the Customer field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionSchedule) GetCustomer() string {
	if o == nil || o.Customer.Get() == nil {
		var ret string
		return ret
	}

	return *o.Customer.Get()
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionSchedule) GetCustomerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customer.Get(), o.Customer.IsSet()
}

// SetCustomer sets field value
func (o *SubscriptionSchedule) SetCustomer(v string) {
	o.Customer.Set(&v)
}

// GetSubscription returns the Subscription field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionSchedule) GetSubscription() string {
	if o == nil || o.Subscription.Get() == nil {
		var ret string
		return ret
	}

	return *o.Subscription.Get()
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionSchedule) GetSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscription.Get(), o.Subscription.IsSet()
}

// SetSubscription sets field value
func (o *SubscriptionSchedule) SetSubscription(v string) {
	o.Subscription.Set(&v)
}

// GetCurrentPhase returns the CurrentPhase field value
func (o *SubscriptionSchedule) GetCurrentPhase() SubscriptionScheduleCurrentPhase {
	if o == nil {
		var ret SubscriptionScheduleCurrentPhase
		return ret
	}

	return o.CurrentPhase
}

// GetCurrentPhaseOk returns a tuple with the CurrentPhase field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetCurrentPhaseOk() (*SubscriptionScheduleCurrentPhase, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPhase, true
}

// SetCurrentPhase sets field value
func (o *SubscriptionSchedule) SetCurrentPhase(v SubscriptionScheduleCurrentPhase) {
	o.CurrentPhase = v
}

// GetPhases returns the Phases field value
func (o *SubscriptionSchedule) GetPhases() []SchedulePhase {
	if o == nil {
		var ret []SchedulePhase
		return ret
	}

	return o.Phases
}

// GetPhasesOk returns a tuple with the Phases field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetPhasesOk() ([]SchedulePhase, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phases, true
}

// SetPhases sets field value
func (o *SubscriptionSchedule) SetPhases(v []SchedulePhase) {
	o.Phases = v
}

// GetEndBehavior returns the EndBehavior field value
func (o *SubscriptionSchedule) GetEndBehavior() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndBehavior
}

// GetEndBehaviorOk returns a tuple with the EndBehavior field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetEndBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndBehavior, true
}

// SetEndBehavior sets field value
func (o *SubscriptionSchedule) SetEndBehavior(v string) {
	o.EndBehavior = v
}

// GetReleasedAt returns the ReleasedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionSchedule) GetReleasedAt() int32 {
	if o == nil || IsNil(o.ReleasedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.ReleasedAt.Get()
}

// GetReleasedAtOk returns a tuple with the ReleasedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionSchedule) GetReleasedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleasedAt.Get(), o.ReleasedAt.IsSet()
}

// HasReleasedAt returns a boolean if a field has been set.
func (o *SubscriptionSchedule) HasReleasedAt() bool {
	if o != nil && o.ReleasedAt.IsSet() {
		return true
	}

	return false
}

// SetReleasedAt gets a reference to the given NullableInt32 and assigns it to the ReleasedAt field.
func (o *SubscriptionSchedule) SetReleasedAt(v int32) {
	o.ReleasedAt.Set(&v)
}
// SetReleasedAtNil sets the value for ReleasedAt to be an explicit nil
func (o *SubscriptionSchedule) SetReleasedAtNil() {
	o.ReleasedAt.Set(nil)
}

// UnsetReleasedAt ensures that no value is present for ReleasedAt, not even an explicit nil
func (o *SubscriptionSchedule) UnsetReleasedAt() {
	o.ReleasedAt.Unset()
}

// GetReleasedSubscription returns the ReleasedSubscription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionSchedule) GetReleasedSubscription() string {
	if o == nil || IsNil(o.ReleasedSubscription.Get()) {
		var ret string
		return ret
	}
	return *o.ReleasedSubscription.Get()
}

// GetReleasedSubscriptionOk returns a tuple with the ReleasedSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionSchedule) GetReleasedSubscriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleasedSubscription.Get(), o.ReleasedSubscription.IsSet()
}

// HasReleasedSubscription returns a boolean if a field has been set.
func (o *SubscriptionSchedule) HasReleasedSubscription() bool {
	if o != nil && o.ReleasedSubscription.IsSet() {
		return true
	}

	return false
}

// SetReleasedSubscription gets a reference to the given NullableString and assigns it to the ReleasedSubscription field.
func (o *SubscriptionSchedule) SetReleasedSubscription(v string) {
	o.ReleasedSubscription.Set(&v)
}
// SetReleasedSubscriptionNil sets the value for ReleasedSubscription to be an explicit nil
func (o *SubscriptionSchedule) SetReleasedSubscriptionNil() {
	o.ReleasedSubscription.Set(nil)
}

// UnsetReleasedSubscription ensures that no value is present for ReleasedSubscription, not even an explicit nil
func (o *SubscriptionSchedule) UnsetReleasedSubscription() {
	o.ReleasedSubscription.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubscriptionSchedule) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubscriptionSchedule) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SubscriptionSchedule) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreated returns the Created field value
func (o *SubscriptionSchedule) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *SubscriptionSchedule) SetCreated(v int32) {
	o.Created = v
}

// GetLivemode returns the Livemode field value
func (o *SubscriptionSchedule) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSchedule) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *SubscriptionSchedule) SetLivemode(v bool) {
	o.Livemode = v
}

func (o SubscriptionSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["customer"] = o.Customer.Get()
	toSerialize["subscription"] = o.Subscription.Get()
	toSerialize["current_phase"] = o.CurrentPhase
	toSerialize["phases"] = o.Phases
	toSerialize["end_behavior"] = o.EndBehavior
	if o.ReleasedAt.IsSet() {
		toSerialize["released_at"] = o.ReleasedAt.Get()
	}
	if o.ReleasedSubscription.IsSet() {
		toSerialize["released_subscription"] = o.ReleasedSubscription.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created"] = o.Created
	toSerialize["livemode"] = o.Livemode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"customer",
		"subscription",
		"current_phase",
		"phases",
		"end_behavior",
		"created",
		"livemode",
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

	varSubscriptionSchedule := _SubscriptionSchedule{}

	err = json.Unmarshal(data, &varSubscriptionSchedule)

	if err != nil {
		return err
	}

	*o = SubscriptionSchedule(varSubscriptionSchedule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "status")
		delete(additionalProperties, "customer")
		delete(additionalProperties, "subscription")
		delete(additionalProperties, "current_phase")
		delete(additionalProperties, "phases")
		delete(additionalProperties, "end_behavior")
		delete(additionalProperties, "released_at")
		delete(additionalProperties, "released_subscription")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created")
		delete(additionalProperties, "livemode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionSchedule struct {
	value *SubscriptionSchedule
	isSet bool
}

func (v NullableSubscriptionSchedule) Get() *SubscriptionSchedule {
	return v.value
}

func (v *NullableSubscriptionSchedule) Set(val *SubscriptionSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSchedule(val *SubscriptionSchedule) *NullableSubscriptionSchedule {
	return &NullableSubscriptionSchedule{value: val, isSet: true}
}

func (v NullableSubscriptionSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


