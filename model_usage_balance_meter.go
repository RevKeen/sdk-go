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

// checks if the UsageBalanceMeter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageBalanceMeter{}

// UsageBalanceMeter Current-period usage for a single meter, including consumed units, included allowance, overage, and estimated cost.
type UsageBalanceMeter struct {
	// Meter ID
	MeterId string `json:"meter_id"`
	// Meter display name
	MeterName string `json:"meter_name"`
	// Unit display name
	UnitName NullableString `json:"unit_name"`
	// Aggregated usage in current period
	CurrentValue float32 `json:"current_value"`
	// Units included in plan (0 if pure metered)
	IncludedQuantity float32 `json:"included_quantity"`
	// Usage consumed (same as current_value)
	UsedQuantity float32 `json:"used_quantity"`
	// Remaining included units (max 0)
	RemainingIncluded float32 `json:"remaining_included"`
	// Units over allowance (max 0)
	OverageQuantity float32 `json:"overage_quantity"`
	// Estimated charge for overage in minor units
	EstimatedAmountMinor float32 `json:"estimated_amount_minor"`
	// Sum of event cost annotations (null if no cost data)
	TotalCostMinor NullableFloat32 `json:"total_cost_minor"`
	// Estimated margin: revenue - cost (null if no cost data)
	MarginMinor NullableFloat32 `json:"margin_minor"`
	// Currency code
	Currency string `json:"currency"`
	// Current billing period start (ISO 8601)
	PeriodStart string `json:"period_start"`
	// Current billing period end (ISO 8601)
	PeriodEnd string `json:"period_end"`
	AdditionalProperties map[string]interface{}
}

type _UsageBalanceMeter UsageBalanceMeter

// NewUsageBalanceMeter instantiates a new UsageBalanceMeter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBalanceMeter(meterId string, meterName string, unitName NullableString, currentValue float32, includedQuantity float32, usedQuantity float32, remainingIncluded float32, overageQuantity float32, estimatedAmountMinor float32, totalCostMinor NullableFloat32, marginMinor NullableFloat32, currency string, periodStart string, periodEnd string) *UsageBalanceMeter {
	this := UsageBalanceMeter{}
	this.MeterId = meterId
	this.MeterName = meterName
	this.UnitName = unitName
	this.CurrentValue = currentValue
	this.IncludedQuantity = includedQuantity
	this.UsedQuantity = usedQuantity
	this.RemainingIncluded = remainingIncluded
	this.OverageQuantity = overageQuantity
	this.EstimatedAmountMinor = estimatedAmountMinor
	this.TotalCostMinor = totalCostMinor
	this.MarginMinor = marginMinor
	this.Currency = currency
	this.PeriodStart = periodStart
	this.PeriodEnd = periodEnd
	return &this
}

// NewUsageBalanceMeterWithDefaults instantiates a new UsageBalanceMeter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBalanceMeterWithDefaults() *UsageBalanceMeter {
	this := UsageBalanceMeter{}
	return &this
}

// GetMeterId returns the MeterId field value
func (o *UsageBalanceMeter) GetMeterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetMeterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterId, true
}

// SetMeterId sets field value
func (o *UsageBalanceMeter) SetMeterId(v string) {
	o.MeterId = v
}

// GetMeterName returns the MeterName field value
func (o *UsageBalanceMeter) GetMeterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterName
}

// GetMeterNameOk returns a tuple with the MeterName field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetMeterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterName, true
}

// SetMeterName sets field value
func (o *UsageBalanceMeter) SetMeterName(v string) {
	o.MeterName = v
}

// GetUnitName returns the UnitName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageBalanceMeter) GetUnitName() string {
	if o == nil || o.UnitName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnitName.Get()
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBalanceMeter) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitName.Get(), o.UnitName.IsSet()
}

// SetUnitName sets field value
func (o *UsageBalanceMeter) SetUnitName(v string) {
	o.UnitName.Set(&v)
}

// GetCurrentValue returns the CurrentValue field value
func (o *UsageBalanceMeter) GetCurrentValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentValue
}

// GetCurrentValueOk returns a tuple with the CurrentValue field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetCurrentValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentValue, true
}

// SetCurrentValue sets field value
func (o *UsageBalanceMeter) SetCurrentValue(v float32) {
	o.CurrentValue = v
}

// GetIncludedQuantity returns the IncludedQuantity field value
func (o *UsageBalanceMeter) GetIncludedQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IncludedQuantity
}

// GetIncludedQuantityOk returns a tuple with the IncludedQuantity field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetIncludedQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludedQuantity, true
}

// SetIncludedQuantity sets field value
func (o *UsageBalanceMeter) SetIncludedQuantity(v float32) {
	o.IncludedQuantity = v
}

// GetUsedQuantity returns the UsedQuantity field value
func (o *UsageBalanceMeter) GetUsedQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UsedQuantity
}

// GetUsedQuantityOk returns a tuple with the UsedQuantity field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetUsedQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedQuantity, true
}

// SetUsedQuantity sets field value
func (o *UsageBalanceMeter) SetUsedQuantity(v float32) {
	o.UsedQuantity = v
}

// GetRemainingIncluded returns the RemainingIncluded field value
func (o *UsageBalanceMeter) GetRemainingIncluded() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemainingIncluded
}

// GetRemainingIncludedOk returns a tuple with the RemainingIncluded field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetRemainingIncludedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingIncluded, true
}

// SetRemainingIncluded sets field value
func (o *UsageBalanceMeter) SetRemainingIncluded(v float32) {
	o.RemainingIncluded = v
}

// GetOverageQuantity returns the OverageQuantity field value
func (o *UsageBalanceMeter) GetOverageQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OverageQuantity
}

// GetOverageQuantityOk returns a tuple with the OverageQuantity field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetOverageQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverageQuantity, true
}

// SetOverageQuantity sets field value
func (o *UsageBalanceMeter) SetOverageQuantity(v float32) {
	o.OverageQuantity = v
}

// GetEstimatedAmountMinor returns the EstimatedAmountMinor field value
func (o *UsageBalanceMeter) GetEstimatedAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.EstimatedAmountMinor
}

// GetEstimatedAmountMinorOk returns a tuple with the EstimatedAmountMinor field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetEstimatedAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedAmountMinor, true
}

// SetEstimatedAmountMinor sets field value
func (o *UsageBalanceMeter) SetEstimatedAmountMinor(v float32) {
	o.EstimatedAmountMinor = v
}

// GetTotalCostMinor returns the TotalCostMinor field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UsageBalanceMeter) GetTotalCostMinor() float32 {
	if o == nil || o.TotalCostMinor.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TotalCostMinor.Get()
}

// GetTotalCostMinorOk returns a tuple with the TotalCostMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBalanceMeter) GetTotalCostMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCostMinor.Get(), o.TotalCostMinor.IsSet()
}

// SetTotalCostMinor sets field value
func (o *UsageBalanceMeter) SetTotalCostMinor(v float32) {
	o.TotalCostMinor.Set(&v)
}

// GetMarginMinor returns the MarginMinor field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UsageBalanceMeter) GetMarginMinor() float32 {
	if o == nil || o.MarginMinor.Get() == nil {
		var ret float32
		return ret
	}

	return *o.MarginMinor.Get()
}

// GetMarginMinorOk returns a tuple with the MarginMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBalanceMeter) GetMarginMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarginMinor.Get(), o.MarginMinor.IsSet()
}

// SetMarginMinor sets field value
func (o *UsageBalanceMeter) SetMarginMinor(v float32) {
	o.MarginMinor.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *UsageBalanceMeter) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UsageBalanceMeter) SetCurrency(v string) {
	o.Currency = v
}

// GetPeriodStart returns the PeriodStart field value
func (o *UsageBalanceMeter) GetPeriodStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetPeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodStart, true
}

// SetPeriodStart sets field value
func (o *UsageBalanceMeter) SetPeriodStart(v string) {
	o.PeriodStart = v
}

// GetPeriodEnd returns the PeriodEnd field value
func (o *UsageBalanceMeter) GetPeriodEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceMeter) GetPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodEnd, true
}

// SetPeriodEnd sets field value
func (o *UsageBalanceMeter) SetPeriodEnd(v string) {
	o.PeriodEnd = v
}

func (o UsageBalanceMeter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageBalanceMeter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meter_id"] = o.MeterId
	toSerialize["meter_name"] = o.MeterName
	toSerialize["unit_name"] = o.UnitName.Get()
	toSerialize["current_value"] = o.CurrentValue
	toSerialize["included_quantity"] = o.IncludedQuantity
	toSerialize["used_quantity"] = o.UsedQuantity
	toSerialize["remaining_included"] = o.RemainingIncluded
	toSerialize["overage_quantity"] = o.OverageQuantity
	toSerialize["estimated_amount_minor"] = o.EstimatedAmountMinor
	toSerialize["total_cost_minor"] = o.TotalCostMinor.Get()
	toSerialize["margin_minor"] = o.MarginMinor.Get()
	toSerialize["currency"] = o.Currency
	toSerialize["period_start"] = o.PeriodStart
	toSerialize["period_end"] = o.PeriodEnd

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageBalanceMeter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meter_id",
		"meter_name",
		"unit_name",
		"current_value",
		"included_quantity",
		"used_quantity",
		"remaining_included",
		"overage_quantity",
		"estimated_amount_minor",
		"total_cost_minor",
		"margin_minor",
		"currency",
		"period_start",
		"period_end",
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

	varUsageBalanceMeter := _UsageBalanceMeter{}

	err = json.Unmarshal(data, &varUsageBalanceMeter)

	if err != nil {
		return err
	}

	*o = UsageBalanceMeter(varUsageBalanceMeter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "meter_id")
		delete(additionalProperties, "meter_name")
		delete(additionalProperties, "unit_name")
		delete(additionalProperties, "current_value")
		delete(additionalProperties, "included_quantity")
		delete(additionalProperties, "used_quantity")
		delete(additionalProperties, "remaining_included")
		delete(additionalProperties, "overage_quantity")
		delete(additionalProperties, "estimated_amount_minor")
		delete(additionalProperties, "total_cost_minor")
		delete(additionalProperties, "margin_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "period_start")
		delete(additionalProperties, "period_end")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageBalanceMeter struct {
	value *UsageBalanceMeter
	isSet bool
}

func (v NullableUsageBalanceMeter) Get() *UsageBalanceMeter {
	return v.value
}

func (v *NullableUsageBalanceMeter) Set(val *UsageBalanceMeter) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageBalanceMeter) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageBalanceMeter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageBalanceMeter(val *UsageBalanceMeter) *NullableUsageBalanceMeter {
	return &NullableUsageBalanceMeter{value: val, isSet: true}
}

func (v NullableUsageBalanceMeter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageBalanceMeter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


