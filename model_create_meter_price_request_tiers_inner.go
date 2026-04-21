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

// checks if the CreateMeterPriceRequestTiersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMeterPriceRequestTiersInner{}

// CreateMeterPriceRequestTiersInner struct for CreateMeterPriceRequestTiersInner
type CreateMeterPriceRequestTiersInner struct {
	// Start of tier range
	FirstUnit int32 `json:"first_unit"`
	// End of tier range (null = unbounded)
	LastUnit NullableInt32 `json:"last_unit,omitempty"`
	// Per-unit price in this tier (minor units)
	UnitAmountMinor int32 `json:"unit_amount_minor"`
	// Flat fee for entering this tier (minor units)
	FlatAmountMinor *int32 `json:"flat_amount_minor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMeterPriceRequestTiersInner CreateMeterPriceRequestTiersInner

// NewCreateMeterPriceRequestTiersInner instantiates a new CreateMeterPriceRequestTiersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMeterPriceRequestTiersInner(firstUnit int32, unitAmountMinor int32) *CreateMeterPriceRequestTiersInner {
	this := CreateMeterPriceRequestTiersInner{}
	this.FirstUnit = firstUnit
	this.UnitAmountMinor = unitAmountMinor
	return &this
}

// NewCreateMeterPriceRequestTiersInnerWithDefaults instantiates a new CreateMeterPriceRequestTiersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMeterPriceRequestTiersInnerWithDefaults() *CreateMeterPriceRequestTiersInner {
	this := CreateMeterPriceRequestTiersInner{}
	return &this
}

// GetFirstUnit returns the FirstUnit field value
func (o *CreateMeterPriceRequestTiersInner) GetFirstUnit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstUnit
}

// GetFirstUnitOk returns a tuple with the FirstUnit field value
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequestTiersInner) GetFirstUnitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstUnit, true
}

// SetFirstUnit sets field value
func (o *CreateMeterPriceRequestTiersInner) SetFirstUnit(v int32) {
	o.FirstUnit = v
}

// GetLastUnit returns the LastUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMeterPriceRequestTiersInner) GetLastUnit() int32 {
	if o == nil || IsNil(o.LastUnit.Get()) {
		var ret int32
		return ret
	}
	return *o.LastUnit.Get()
}

// GetLastUnitOk returns a tuple with the LastUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMeterPriceRequestTiersInner) GetLastUnitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUnit.Get(), o.LastUnit.IsSet()
}

// HasLastUnit returns a boolean if a field has been set.
func (o *CreateMeterPriceRequestTiersInner) HasLastUnit() bool {
	if o != nil && o.LastUnit.IsSet() {
		return true
	}

	return false
}

// SetLastUnit gets a reference to the given NullableInt32 and assigns it to the LastUnit field.
func (o *CreateMeterPriceRequestTiersInner) SetLastUnit(v int32) {
	o.LastUnit.Set(&v)
}
// SetLastUnitNil sets the value for LastUnit to be an explicit nil
func (o *CreateMeterPriceRequestTiersInner) SetLastUnitNil() {
	o.LastUnit.Set(nil)
}

// UnsetLastUnit ensures that no value is present for LastUnit, not even an explicit nil
func (o *CreateMeterPriceRequestTiersInner) UnsetLastUnit() {
	o.LastUnit.Unset()
}

// GetUnitAmountMinor returns the UnitAmountMinor field value
func (o *CreateMeterPriceRequestTiersInner) GetUnitAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmountMinor
}

// GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequestTiersInner) GetUnitAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmountMinor, true
}

// SetUnitAmountMinor sets field value
func (o *CreateMeterPriceRequestTiersInner) SetUnitAmountMinor(v int32) {
	o.UnitAmountMinor = v
}

// GetFlatAmountMinor returns the FlatAmountMinor field value if set, zero value otherwise.
func (o *CreateMeterPriceRequestTiersInner) GetFlatAmountMinor() int32 {
	if o == nil || IsNil(o.FlatAmountMinor) {
		var ret int32
		return ret
	}
	return *o.FlatAmountMinor
}

// GetFlatAmountMinorOk returns a tuple with the FlatAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequestTiersInner) GetFlatAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.FlatAmountMinor) {
		return nil, false
	}
	return o.FlatAmountMinor, true
}

// HasFlatAmountMinor returns a boolean if a field has been set.
func (o *CreateMeterPriceRequestTiersInner) HasFlatAmountMinor() bool {
	if o != nil && !IsNil(o.FlatAmountMinor) {
		return true
	}

	return false
}

// SetFlatAmountMinor gets a reference to the given int32 and assigns it to the FlatAmountMinor field.
func (o *CreateMeterPriceRequestTiersInner) SetFlatAmountMinor(v int32) {
	o.FlatAmountMinor = &v
}

func (o CreateMeterPriceRequestTiersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMeterPriceRequestTiersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first_unit"] = o.FirstUnit
	if o.LastUnit.IsSet() {
		toSerialize["last_unit"] = o.LastUnit.Get()
	}
	toSerialize["unit_amount_minor"] = o.UnitAmountMinor
	if !IsNil(o.FlatAmountMinor) {
		toSerialize["flat_amount_minor"] = o.FlatAmountMinor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMeterPriceRequestTiersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_unit",
		"unit_amount_minor",
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

	varCreateMeterPriceRequestTiersInner := _CreateMeterPriceRequestTiersInner{}

	err = json.Unmarshal(data, &varCreateMeterPriceRequestTiersInner)

	if err != nil {
		return err
	}

	*o = CreateMeterPriceRequestTiersInner(varCreateMeterPriceRequestTiersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "first_unit")
		delete(additionalProperties, "last_unit")
		delete(additionalProperties, "unit_amount_minor")
		delete(additionalProperties, "flat_amount_minor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMeterPriceRequestTiersInner struct {
	value *CreateMeterPriceRequestTiersInner
	isSet bool
}

func (v NullableCreateMeterPriceRequestTiersInner) Get() *CreateMeterPriceRequestTiersInner {
	return v.value
}

func (v *NullableCreateMeterPriceRequestTiersInner) Set(val *CreateMeterPriceRequestTiersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMeterPriceRequestTiersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMeterPriceRequestTiersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMeterPriceRequestTiersInner(val *CreateMeterPriceRequestTiersInner) *NullableCreateMeterPriceRequestTiersInner {
	return &NullableCreateMeterPriceRequestTiersInner{value: val, isSet: true}
}

func (v NullableCreateMeterPriceRequestTiersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMeterPriceRequestTiersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


