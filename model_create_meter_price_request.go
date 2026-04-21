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

// checks if the CreateMeterPriceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMeterPriceRequest{}

// CreateMeterPriceRequest Parameters for creating a price attached to a usage meter, defining per-unit cost and billing thresholds.
type CreateMeterPriceRequest struct {
	// Pricing model
	PricingModel string `json:"pricing_model"`
	// ISO 4217 currency code
	Currency string `json:"currency"`
	// Price per unit in minor units (for per_unit)
	UnitAmountMinor *int32 `json:"unit_amount_minor,omitempty"`
	// Base charge per period in minor units
	FlatFeeMinor *int32 `json:"flat_fee_minor,omitempty"`
	// Units per package (for package model)
	PackageSize *int32 `json:"package_size,omitempty"`
	// Merchant's cost per unit for margin intelligence
	CostPerUnitMinor *int32 `json:"cost_per_unit_minor,omitempty"`
	// Tier configuration (for graduated and volume)
	Tiers []CreateMeterPriceRequestTiersInner `json:"tiers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMeterPriceRequest CreateMeterPriceRequest

// NewCreateMeterPriceRequest instantiates a new CreateMeterPriceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMeterPriceRequest(pricingModel string, currency string) *CreateMeterPriceRequest {
	this := CreateMeterPriceRequest{}
	this.PricingModel = pricingModel
	this.Currency = currency
	return &this
}

// NewCreateMeterPriceRequestWithDefaults instantiates a new CreateMeterPriceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMeterPriceRequestWithDefaults() *CreateMeterPriceRequest {
	this := CreateMeterPriceRequest{}
	return &this
}

// GetPricingModel returns the PricingModel field value
func (o *CreateMeterPriceRequest) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *CreateMeterPriceRequest) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetCurrency returns the Currency field value
func (o *CreateMeterPriceRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateMeterPriceRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetUnitAmountMinor returns the UnitAmountMinor field value if set, zero value otherwise.
func (o *CreateMeterPriceRequest) GetUnitAmountMinor() int32 {
	if o == nil || IsNil(o.UnitAmountMinor) {
		var ret int32
		return ret
	}
	return *o.UnitAmountMinor
}

// GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetUnitAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitAmountMinor) {
		return nil, false
	}
	return o.UnitAmountMinor, true
}

// HasUnitAmountMinor returns a boolean if a field has been set.
func (o *CreateMeterPriceRequest) HasUnitAmountMinor() bool {
	if o != nil && !IsNil(o.UnitAmountMinor) {
		return true
	}

	return false
}

// SetUnitAmountMinor gets a reference to the given int32 and assigns it to the UnitAmountMinor field.
func (o *CreateMeterPriceRequest) SetUnitAmountMinor(v int32) {
	o.UnitAmountMinor = &v
}

// GetFlatFeeMinor returns the FlatFeeMinor field value if set, zero value otherwise.
func (o *CreateMeterPriceRequest) GetFlatFeeMinor() int32 {
	if o == nil || IsNil(o.FlatFeeMinor) {
		var ret int32
		return ret
	}
	return *o.FlatFeeMinor
}

// GetFlatFeeMinorOk returns a tuple with the FlatFeeMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetFlatFeeMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.FlatFeeMinor) {
		return nil, false
	}
	return o.FlatFeeMinor, true
}

// HasFlatFeeMinor returns a boolean if a field has been set.
func (o *CreateMeterPriceRequest) HasFlatFeeMinor() bool {
	if o != nil && !IsNil(o.FlatFeeMinor) {
		return true
	}

	return false
}

// SetFlatFeeMinor gets a reference to the given int32 and assigns it to the FlatFeeMinor field.
func (o *CreateMeterPriceRequest) SetFlatFeeMinor(v int32) {
	o.FlatFeeMinor = &v
}

// GetPackageSize returns the PackageSize field value if set, zero value otherwise.
func (o *CreateMeterPriceRequest) GetPackageSize() int32 {
	if o == nil || IsNil(o.PackageSize) {
		var ret int32
		return ret
	}
	return *o.PackageSize
}

// GetPackageSizeOk returns a tuple with the PackageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetPackageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageSize) {
		return nil, false
	}
	return o.PackageSize, true
}

// HasPackageSize returns a boolean if a field has been set.
func (o *CreateMeterPriceRequest) HasPackageSize() bool {
	if o != nil && !IsNil(o.PackageSize) {
		return true
	}

	return false
}

// SetPackageSize gets a reference to the given int32 and assigns it to the PackageSize field.
func (o *CreateMeterPriceRequest) SetPackageSize(v int32) {
	o.PackageSize = &v
}

// GetCostPerUnitMinor returns the CostPerUnitMinor field value if set, zero value otherwise.
func (o *CreateMeterPriceRequest) GetCostPerUnitMinor() int32 {
	if o == nil || IsNil(o.CostPerUnitMinor) {
		var ret int32
		return ret
	}
	return *o.CostPerUnitMinor
}

// GetCostPerUnitMinorOk returns a tuple with the CostPerUnitMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetCostPerUnitMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.CostPerUnitMinor) {
		return nil, false
	}
	return o.CostPerUnitMinor, true
}

// HasCostPerUnitMinor returns a boolean if a field has been set.
func (o *CreateMeterPriceRequest) HasCostPerUnitMinor() bool {
	if o != nil && !IsNil(o.CostPerUnitMinor) {
		return true
	}

	return false
}

// SetCostPerUnitMinor gets a reference to the given int32 and assigns it to the CostPerUnitMinor field.
func (o *CreateMeterPriceRequest) SetCostPerUnitMinor(v int32) {
	o.CostPerUnitMinor = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *CreateMeterPriceRequest) GetTiers() []CreateMeterPriceRequestTiersInner {
	if o == nil || IsNil(o.Tiers) {
		var ret []CreateMeterPriceRequestTiersInner
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMeterPriceRequest) GetTiersOk() ([]CreateMeterPriceRequestTiersInner, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *CreateMeterPriceRequest) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []CreateMeterPriceRequestTiersInner and assigns it to the Tiers field.
func (o *CreateMeterPriceRequest) SetTiers(v []CreateMeterPriceRequestTiersInner) {
	o.Tiers = v
}

func (o CreateMeterPriceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMeterPriceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pricing_model"] = o.PricingModel
	toSerialize["currency"] = o.Currency
	if !IsNil(o.UnitAmountMinor) {
		toSerialize["unit_amount_minor"] = o.UnitAmountMinor
	}
	if !IsNil(o.FlatFeeMinor) {
		toSerialize["flat_fee_minor"] = o.FlatFeeMinor
	}
	if !IsNil(o.PackageSize) {
		toSerialize["package_size"] = o.PackageSize
	}
	if !IsNil(o.CostPerUnitMinor) {
		toSerialize["cost_per_unit_minor"] = o.CostPerUnitMinor
	}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMeterPriceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pricing_model",
		"currency",
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

	varCreateMeterPriceRequest := _CreateMeterPriceRequest{}

	err = json.Unmarshal(data, &varCreateMeterPriceRequest)

	if err != nil {
		return err
	}

	*o = CreateMeterPriceRequest(varCreateMeterPriceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pricing_model")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "unit_amount_minor")
		delete(additionalProperties, "flat_fee_minor")
		delete(additionalProperties, "package_size")
		delete(additionalProperties, "cost_per_unit_minor")
		delete(additionalProperties, "tiers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMeterPriceRequest struct {
	value *CreateMeterPriceRequest
	isSet bool
}

func (v NullableCreateMeterPriceRequest) Get() *CreateMeterPriceRequest {
	return v.value
}

func (v *NullableCreateMeterPriceRequest) Set(val *CreateMeterPriceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMeterPriceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMeterPriceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMeterPriceRequest(val *CreateMeterPriceRequest) *NullableCreateMeterPriceRequest {
	return &NullableCreateMeterPriceRequest{value: val, isSet: true}
}

func (v NullableCreateMeterPriceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMeterPriceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


