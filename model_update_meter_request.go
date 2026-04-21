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
)

// checks if the UpdateMeterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMeterRequest{}

// UpdateMeterRequest Parameters for updating an existing meter. Only display name and metadata can be changed.
type UpdateMeterRequest struct {
	// Updated meter name
	Name *string `json:"name,omitempty"`
	// Updated slug
	Slug *string `json:"slug,omitempty"`
	// Updated description
	Description *string `json:"description,omitempty"`
	// Meter status
	Status *string `json:"status,omitempty"`
	// Updated value key
	ValueKey *string `json:"value_key,omitempty"`
	// Updated filter conditions
	FilterConditions []map[string]interface{} `json:"filter_conditions,omitempty"`
	// Updated unique count key
	UniqueCountKey *string `json:"unique_count_key,omitempty"`
	// Updated unit name
	UnitName *string `json:"unit_name,omitempty"`
	// Updated carry forward setting
	CarryForward *bool `json:"carry_forward,omitempty"`
	// Updated metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Updated threshold percentages for usage.threshold.reached webhooks.
	AlertThresholds []float32 `json:"alert_thresholds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateMeterRequest UpdateMeterRequest

// NewUpdateMeterRequest instantiates a new UpdateMeterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMeterRequest() *UpdateMeterRequest {
	this := UpdateMeterRequest{}
	return &this
}

// NewUpdateMeterRequestWithDefaults instantiates a new UpdateMeterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMeterRequestWithDefaults() *UpdateMeterRequest {
	this := UpdateMeterRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateMeterRequest) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *UpdateMeterRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateMeterRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateMeterRequest) SetStatus(v string) {
	o.Status = &v
}

// GetValueKey returns the ValueKey field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetValueKey() string {
	if o == nil || IsNil(o.ValueKey) {
		var ret string
		return ret
	}
	return *o.ValueKey
}

// GetValueKeyOk returns a tuple with the ValueKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetValueKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ValueKey) {
		return nil, false
	}
	return o.ValueKey, true
}

// HasValueKey returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasValueKey() bool {
	if o != nil && !IsNil(o.ValueKey) {
		return true
	}

	return false
}

// SetValueKey gets a reference to the given string and assigns it to the ValueKey field.
func (o *UpdateMeterRequest) SetValueKey(v string) {
	o.ValueKey = &v
}

// GetFilterConditions returns the FilterConditions field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetFilterConditions() []map[string]interface{} {
	if o == nil || IsNil(o.FilterConditions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.FilterConditions
}

// GetFilterConditionsOk returns a tuple with the FilterConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetFilterConditionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.FilterConditions) {
		return nil, false
	}
	return o.FilterConditions, true
}

// HasFilterConditions returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasFilterConditions() bool {
	if o != nil && !IsNil(o.FilterConditions) {
		return true
	}

	return false
}

// SetFilterConditions gets a reference to the given []map[string]interface{} and assigns it to the FilterConditions field.
func (o *UpdateMeterRequest) SetFilterConditions(v []map[string]interface{}) {
	o.FilterConditions = v
}

// GetUniqueCountKey returns the UniqueCountKey field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetUniqueCountKey() string {
	if o == nil || IsNil(o.UniqueCountKey) {
		var ret string
		return ret
	}
	return *o.UniqueCountKey
}

// GetUniqueCountKeyOk returns a tuple with the UniqueCountKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetUniqueCountKeyOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueCountKey) {
		return nil, false
	}
	return o.UniqueCountKey, true
}

// HasUniqueCountKey returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasUniqueCountKey() bool {
	if o != nil && !IsNil(o.UniqueCountKey) {
		return true
	}

	return false
}

// SetUniqueCountKey gets a reference to the given string and assigns it to the UniqueCountKey field.
func (o *UpdateMeterRequest) SetUniqueCountKey(v string) {
	o.UniqueCountKey = &v
}

// GetUnitName returns the UnitName field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetUnitName() string {
	if o == nil || IsNil(o.UnitName) {
		var ret string
		return ret
	}
	return *o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetUnitNameOk() (*string, bool) {
	if o == nil || IsNil(o.UnitName) {
		return nil, false
	}
	return o.UnitName, true
}

// HasUnitName returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasUnitName() bool {
	if o != nil && !IsNil(o.UnitName) {
		return true
	}

	return false
}

// SetUnitName gets a reference to the given string and assigns it to the UnitName field.
func (o *UpdateMeterRequest) SetUnitName(v string) {
	o.UnitName = &v
}

// GetCarryForward returns the CarryForward field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetCarryForward() bool {
	if o == nil || IsNil(o.CarryForward) {
		var ret bool
		return ret
	}
	return *o.CarryForward
}

// GetCarryForwardOk returns a tuple with the CarryForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetCarryForwardOk() (*bool, bool) {
	if o == nil || IsNil(o.CarryForward) {
		return nil, false
	}
	return o.CarryForward, true
}

// HasCarryForward returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasCarryForward() bool {
	if o != nil && !IsNil(o.CarryForward) {
		return true
	}

	return false
}

// SetCarryForward gets a reference to the given bool and assigns it to the CarryForward field.
func (o *UpdateMeterRequest) SetCarryForward(v bool) {
	o.CarryForward = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateMeterRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAlertThresholds returns the AlertThresholds field value if set, zero value otherwise.
func (o *UpdateMeterRequest) GetAlertThresholds() []float32 {
	if o == nil || IsNil(o.AlertThresholds) {
		var ret []float32
		return ret
	}
	return o.AlertThresholds
}

// GetAlertThresholdsOk returns a tuple with the AlertThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMeterRequest) GetAlertThresholdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.AlertThresholds) {
		return nil, false
	}
	return o.AlertThresholds, true
}

// HasAlertThresholds returns a boolean if a field has been set.
func (o *UpdateMeterRequest) HasAlertThresholds() bool {
	if o != nil && !IsNil(o.AlertThresholds) {
		return true
	}

	return false
}

// SetAlertThresholds gets a reference to the given []float32 and assigns it to the AlertThresholds field.
func (o *UpdateMeterRequest) SetAlertThresholds(v []float32) {
	o.AlertThresholds = v
}

func (o UpdateMeterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMeterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

func (o *UpdateMeterRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateMeterRequest := _UpdateMeterRequest{}

	err = json.Unmarshal(data, &varUpdateMeterRequest)

	if err != nil {
		return err
	}

	*o = UpdateMeterRequest(varUpdateMeterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
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

type NullableUpdateMeterRequest struct {
	value *UpdateMeterRequest
	isSet bool
}

func (v NullableUpdateMeterRequest) Get() *UpdateMeterRequest {
	return v.value
}

func (v *NullableUpdateMeterRequest) Set(val *UpdateMeterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMeterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMeterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMeterRequest(val *UpdateMeterRequest) *NullableUpdateMeterRequest {
	return &NullableUpdateMeterRequest{value: val, isSet: true}
}

func (v NullableUpdateMeterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMeterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


