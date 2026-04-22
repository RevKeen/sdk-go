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

// checks if the CreateMeterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMeterRequest{}

// CreateMeterRequest Parameters for creating a new usage meter with aggregation rules and event matching.
type CreateMeterRequest struct {
	// Human-readable meter name
	Name string `json:"name"`
	// Event name to match against
	EventName string `json:"event_name"`
	// Aggregation method
	Aggregation *string `json:"aggregation,omitempty"`
	// URL-safe identifier
	Slug *string `json:"slug,omitempty"`
	// Meter description
	Description *string `json:"description,omitempty"`
	// Property key for sum/max/last aggregations
	ValueKey *string `json:"value_key,omitempty"`
	// Filter conditions for matching events
	FilterConditions []map[string]interface{} `json:"filter_conditions,omitempty"`
	// Property key for count_unique aggregation
	UniqueCountKey *string `json:"unique_count_key,omitempty"`
	// Display unit name
	UnitName *string `json:"unit_name,omitempty"`
	// Carry forward last value across periods
	CarryForward *bool `json:"carry_forward,omitempty"`
	// Arbitrary key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Threshold percentages that trigger usage.threshold.reached webhooks.
	AlertThresholds []float32 `json:"alert_thresholds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMeterRequest CreateMeterRequest

// NewCreateMeterRequest instantiates a new CreateMeterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMeterRequest(name string, eventName string) *CreateMeterRequest {
	this := CreateMeterRequest{}
	this.Name = name
	this.EventName = eventName
	var aggregation string = "count"
	this.Aggregation = &aggregation
	return &this
}

// NewCreateMeterRequestWithDefaults instantiates a new CreateMeterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMeterRequestWithDefaults() *CreateMeterRequest {
	this := CreateMeterRequest{}
	var aggregation string = "count"
	this.Aggregation = &aggregation
	return &this
}

// GetName returns the Name field value
func (o *CreateMeterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMeterRequest) SetName(v string) {
	o.Name = v
}

// GetEventName returns the EventName field value
func (o *CreateMeterRequest) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *CreateMeterRequest) SetEventName(v string) {
	o.EventName = v
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetAggregation() string {
	if o == nil || IsNil(o.Aggregation) {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *CreateMeterRequest) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *CreateMeterRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateMeterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetValueKey returns the ValueKey field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetValueKey() string {
	if o == nil || IsNil(o.ValueKey) {
		var ret string
		return ret
	}
	return *o.ValueKey
}

// GetValueKeyOk returns a tuple with the ValueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetValueKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ValueKey) {
		return nil, false
	}
	return o.ValueKey, true
}

// HasValueKey returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasValueKey() bool {
	if o != nil && !IsNil(o.ValueKey) {
		return true
	}

	return false
}

// SetValueKey gets a reference to the given string and assigns it to the ValueKey field.
func (o *CreateMeterRequest) SetValueKey(v string) {
	o.ValueKey = &v
}

// GetFilterConditions returns the FilterConditions field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetFilterConditions() []map[string]interface{} {
	if o == nil || IsNil(o.FilterConditions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.FilterConditions
}

// GetFilterConditionsOk returns a tuple with the FilterConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetFilterConditionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.FilterConditions) {
		return nil, false
	}
	return o.FilterConditions, true
}

// HasFilterConditions returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasFilterConditions() bool {
	if o != nil && !IsNil(o.FilterConditions) {
		return true
	}

	return false
}

// SetFilterConditions gets a reference to the given []map[string]interface{} and assigns it to the FilterConditions field.
func (o *CreateMeterRequest) SetFilterConditions(v []map[string]interface{}) {
	o.FilterConditions = v
}

// GetUniqueCountKey returns the UniqueCountKey field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetUniqueCountKey() string {
	if o == nil || IsNil(o.UniqueCountKey) {
		var ret string
		return ret
	}
	return *o.UniqueCountKey
}

// GetUniqueCountKeyOk returns a tuple with the UniqueCountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetUniqueCountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueCountKey) {
		return nil, false
	}
	return o.UniqueCountKey, true
}

// HasUniqueCountKey returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasUniqueCountKey() bool {
	if o != nil && !IsNil(o.UniqueCountKey) {
		return true
	}

	return false
}

// SetUniqueCountKey gets a reference to the given string and assigns it to the UniqueCountKey field.
func (o *CreateMeterRequest) SetUniqueCountKey(v string) {
	o.UniqueCountKey = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetUnitName() string {
	if o == nil || IsNil(o.UnitName) {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetUnitNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitName) {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasUnitName() bool {
	if o != nil && !IsNil(o.UnitName) {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *CreateMeterRequest) SetUnitName(v string) {
	o.UnitName = &v
}

// GetCarryForward returns the CarryForward field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetCarryForward() bool {
	if o == nil || IsNil(o.CarryForward) {
		var ret bool
		return ret
	}
	return *o.CarryForward
}

// GetCarryForwardOk returns a tuple with the CarryForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetCarryForwardOk() (*bool, bool) {
	if o == nil || IsNil(o.CarryForward) {
		return nil, false
	}
	return o.CarryForward, true
}

// HasCarryForward returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasCarryForward() bool {
	if o != nil && !IsNil(o.CarryForward) {
		return true
	}

	return false
}

// SetCarryForward gets a reference to the given bool and assigns it to the CarryForward field.
func (o *CreateMeterRequest) SetCarryForward(v bool) {
	o.CarryForward = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateMeterRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAlertThresholds returns the AlertThresholds field value if set, zero value otherwise.
func (o *CreateMeterRequest) GetAlertThresholds() []float32 {
	if o == nil || IsNil(o.AlertThresholds) {
		var ret []float32
		return ret
	}
	return o.AlertThresholds
}

// GetAlertThresholdsOk returns a tuple with the AlertThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterRequest) GetAlertThresholdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.AlertThresholds) {
		return nil, false
	}
	return o.AlertThresholds, true
}

// HasAlertThresholds returns a boolean if a field has been set.
func (o *CreateMeterRequest) HasAlertThresholds() bool {
	if o != nil && !IsNil(o.AlertThresholds) {
		return true
	}

	return false
}

// SetAlertThresholds gets a reference to the given []float32 and assigns it to the AlertThresholds field.
func (o *CreateMeterRequest) SetAlertThresholds(v []float32) {
	o.AlertThresholds = v
}

func (o CreateMeterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMeterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["event_name"] = o.EventName
	if !IsNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ValueKey) {
		toSerialize["value_key"] = o.ValueKey
	}
	if !IsNil(o.FilterConditions) {
		toSerialize["filter_conditions"] = o.FilterConditions
	}
	if !IsNil(o.UniqueCountKey) {
		toSerialize["unique_count_key"] = o.UniqueCountKey
	}
	if !IsNil(o.UnitName) {
		toSerialize["unit_name"] = o.UnitName
	}
	if !IsNil(o.CarryForward) {
		toSerialize["carry_forward"] = o.CarryForward
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.AlertThresholds) {
		toSerialize["alert_thresholds"] = o.AlertThresholds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMeterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"event_name",
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

	varCreateMeterRequest := _CreateMeterRequest{}

	err = json.Unmarshal(data, &varCreateMeterRequest)

	if err != nil {
		return err
	}

	*o = CreateMeterRequest(varCreateMeterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "event_name")
		delete(additionalProperties, "aggregation")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "value_key")
		delete(additionalProperties, "filter_conditions")
		delete(additionalProperties, "unique_count_key")
		delete(additionalProperties, "unit_name")
		delete(additionalProperties, "carry_forward")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "alert_thresholds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMeterRequest struct {
	value *CreateMeterRequest
	isSet bool
}

func (v NullableCreateMeterRequest) Get() *CreateMeterRequest {
	return v.value
}

func (v *NullableCreateMeterRequest) Set(val *CreateMeterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMeterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMeterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMeterRequest(val *CreateMeterRequest) *NullableCreateMeterRequest {
	return &NullableCreateMeterRequest{value: val, isSet: true}
}

func (v NullableCreateMeterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMeterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


