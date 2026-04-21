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

// checks if the Benefit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Benefit{}

// Benefit struct for Benefit
type Benefit struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	BenefitType string `json:"benefitType"`
	BenefitKey string `json:"benefitKey"`
	Category NullableString `json:"category"`
	IconUrl NullableString `json:"iconUrl"`
	DisplayOrder NullableString `json:"displayOrder"`
	IsActive NullableBool `json:"isActive"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	Config interface{} `json:"config,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Benefit Benefit

// NewBenefit instantiates a new Benefit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBenefit(id string, name string, description NullableString, benefitType string, benefitKey string, category NullableString, iconUrl NullableString, displayOrder NullableString, isActive NullableBool) *Benefit {
	this := Benefit{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.BenefitType = benefitType
	this.BenefitKey = benefitKey
	this.Category = category
	this.IconUrl = iconUrl
	this.DisplayOrder = displayOrder
	this.IsActive = isActive
	return &this
}

// NewBenefitWithDefaults instantiates a new Benefit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBenefitWithDefaults() *Benefit {
	this := Benefit{}
	return &this
}

// GetId returns the Id field value
func (o *Benefit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Benefit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Benefit) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Benefit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Benefit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Benefit) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Benefit) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Benefit) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetBenefitType returns the BenefitType field value
func (o *Benefit) GetBenefitType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BenefitType
}

// GetBenefitTypeOk returns a tuple with the BenefitType field value
// and a boolean to check if the value has been set.
func (o *Benefit) GetBenefitTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitType, true
}

// SetBenefitType sets field value
func (o *Benefit) SetBenefitType(v string) {
	o.BenefitType = v
}

// GetBenefitKey returns the BenefitKey field value
func (o *Benefit) GetBenefitKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BenefitKey
}

// GetBenefitKeyOk returns a tuple with the BenefitKey field value
// and a boolean to check if the value has been set.
func (o *Benefit) GetBenefitKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitKey, true
}

// SetBenefitKey sets field value
func (o *Benefit) SetBenefitKey(v string) {
	o.BenefitKey = v
}

// GetCategory returns the Category field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Benefit) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}

	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// SetCategory sets field value
func (o *Benefit) SetCategory(v string) {
	o.Category.Set(&v)
}

// GetIconUrl returns the IconUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Benefit) GetIconUrl() string {
	if o == nil || o.IconUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// SetIconUrl sets field value
func (o *Benefit) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// GetDisplayOrder returns the DisplayOrder field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Benefit) GetDisplayOrder() string {
	if o == nil || o.DisplayOrder.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayOrder.Get()
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetDisplayOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayOrder.Get(), o.DisplayOrder.IsSet()
}

// SetDisplayOrder sets field value
func (o *Benefit) SetDisplayOrder(v string) {
	o.DisplayOrder.Set(&v)
}

// GetIsActive returns the IsActive field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Benefit) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// SetIsActive sets field value
func (o *Benefit) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *Benefit) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *Benefit) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Benefit) GetConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Benefit) GetConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Benefit) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given interface{} and assigns it to the Config field.
func (o *Benefit) SetConfig(v interface{}) {
	o.Config = v
}

func (o Benefit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Benefit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["benefitType"] = o.BenefitType
	toSerialize["benefitKey"] = o.BenefitKey
	toSerialize["category"] = o.Category.Get()
	toSerialize["iconUrl"] = o.IconUrl.Get()
	toSerialize["displayOrder"] = o.DisplayOrder.Get()
	toSerialize["isActive"] = o.IsActive.Get()
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Benefit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"benefitType",
		"benefitKey",
		"category",
		"iconUrl",
		"displayOrder",
		"isActive",
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

	varBenefit := _Benefit{}

	err = json.Unmarshal(data, &varBenefit)

	if err != nil {
		return err
	}

	*o = Benefit(varBenefit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "benefitType")
		delete(additionalProperties, "benefitKey")
		delete(additionalProperties, "category")
		delete(additionalProperties, "iconUrl")
		delete(additionalProperties, "displayOrder")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "defaultValue")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBenefit struct {
	value *Benefit
	isSet bool
}

func (v NullableBenefit) Get() *Benefit {
	return v.value
}

func (v *NullableBenefit) Set(val *Benefit) {
	v.value = val
	v.isSet = true
}

func (v NullableBenefit) IsSet() bool {
	return v.isSet
}

func (v *NullableBenefit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBenefit(val *Benefit) *NullableBenefit {
	return &NullableBenefit{value: val, isSet: true}
}

func (v NullableBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBenefit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


