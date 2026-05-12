/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the CustomerMeter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerMeter{}

// CustomerMeter An on-the-fly aggregate of a customer's consumption for a single meter. RevKeen computes this live from `usage_events` — there is no separate `customer_meters` table to keep in sync.
type CustomerMeter struct {
	Object string `json:"object"`
	CustomerId string `json:"customer_id"`
	MeterId string `json:"meter_id"`
	MeterName string `json:"meter_name"`
	// The `event_name` the merchant's code posts to /v2/usage-events for this meter.
	MeterEventName string `json:"meter_event_name"`
	// How quantities are aggregated. Typical values: `sum`, `count`, `max`, `unique`.
	Aggregation string `json:"aggregation"`
	// Human-readable unit for display (e.g. `tokens`, `requests`, `GB`). Merchant-configurable.
	UnitName NullableString `json:"unit_name"`
	// Aggregate across all usage events for this (customer, meter) pair. The aggregation function applied matches the meter's `aggregation` field.
	TotalQuantity float32 `json:"total_quantity"`
	// Number of `usage_events` rows that contributed to the aggregate. Diagnostic — not the aggregate itself.
	EventCount int32 `json:"event_count"`
	// Timestamp of the most recent usage event. Null if no events have been recorded.
	LastEventAt NullableTime `json:"last_event_at"`
	AdditionalProperties map[string]interface{}
}

type _CustomerMeter CustomerMeter

// NewCustomerMeter instantiates a new CustomerMeter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerMeter(object string, customerId string, meterId string, meterName string, meterEventName string, aggregation string, unitName NullableString, totalQuantity float32, eventCount int32, lastEventAt NullableTime) *CustomerMeter {
	this := CustomerMeter{}
	this.Object = object
	this.CustomerId = customerId
	this.MeterId = meterId
	this.MeterName = meterName
	this.MeterEventName = meterEventName
	this.Aggregation = aggregation
	this.UnitName = unitName
	this.TotalQuantity = totalQuantity
	this.EventCount = eventCount
	this.LastEventAt = lastEventAt
	return &this
}

// NewCustomerMeterWithDefaults instantiates a new CustomerMeter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerMeterWithDefaults() *CustomerMeter {
	this := CustomerMeter{}
	return &this
}

// GetObject returns the Object field value
func (o *CustomerMeter) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CustomerMeter) SetObject(v string) {
	o.Object = v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomerMeter) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomerMeter) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetMeterId returns the MeterId field value
func (o *CustomerMeter) GetMeterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetMeterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterId, true
}

// SetMeterId sets field value
func (o *CustomerMeter) SetMeterId(v string) {
	o.MeterId = v
}

// GetMeterName returns the MeterName field value
func (o *CustomerMeter) GetMeterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterName
}

// GetMeterNameOk returns a tuple with the MeterName field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetMeterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterName, true
}

// SetMeterName sets field value
func (o *CustomerMeter) SetMeterName(v string) {
	o.MeterName = v
}

// GetMeterEventName returns the MeterEventName field value
func (o *CustomerMeter) GetMeterEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterEventName
}

// GetMeterEventNameOk returns a tuple with the MeterEventName field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetMeterEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterEventName, true
}

// SetMeterEventName sets field value
func (o *CustomerMeter) SetMeterEventName(v string) {
	o.MeterEventName = v
}

// GetAggregation returns the Aggregation field value
func (o *CustomerMeter) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *CustomerMeter) SetAggregation(v string) {
	o.Aggregation = v
}

// GetUnitName returns the UnitName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerMeter) GetUnitName() string {
	if o == nil || o.UnitName.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnitName.Get()
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerMeter) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitName.Get(), o.UnitName.IsSet()
}

// SetUnitName sets field value
func (o *CustomerMeter) SetUnitName(v string) {
	o.UnitName.Set(&v)
}

// GetTotalQuantity returns the TotalQuantity field value
func (o *CustomerMeter) GetTotalQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetTotalQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalQuantity, true
}

// SetTotalQuantity sets field value
func (o *CustomerMeter) SetTotalQuantity(v float32) {
	o.TotalQuantity = v
}

// GetEventCount returns the EventCount field value
func (o *CustomerMeter) GetEventCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *CustomerMeter) GetEventCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value
func (o *CustomerMeter) SetEventCount(v int32) {
	o.EventCount = v
}

// GetLastEventAt returns the LastEventAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomerMeter) GetLastEventAt() time.Time {
	if o == nil || o.LastEventAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastEventAt.Get()
}

// GetLastEventAtOk returns a tuple with the LastEventAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerMeter) GetLastEventAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastEventAt.Get(), o.LastEventAt.IsSet()
}

// SetLastEventAt sets field value
func (o *CustomerMeter) SetLastEventAt(v time.Time) {
	o.LastEventAt.Set(&v)
}

func (o CustomerMeter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerMeter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["meter_id"] = o.MeterId
	toSerialize["meter_name"] = o.MeterName
	toSerialize["meter_event_name"] = o.MeterEventName
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["unit_name"] = o.UnitName.Get()
	toSerialize["total_quantity"] = o.TotalQuantity
	toSerialize["event_count"] = o.EventCount
	toSerialize["last_event_at"] = o.LastEventAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerMeter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"customer_id",
		"meter_id",
		"meter_name",
		"meter_event_name",
		"aggregation",
		"unit_name",
		"total_quantity",
		"event_count",
		"last_event_at",
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

	varCustomerMeter := _CustomerMeter{}

	err = json.Unmarshal(data, &varCustomerMeter)

	if err != nil {
		return err
	}

	*o = CustomerMeter(varCustomerMeter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "meter_id")
		delete(additionalProperties, "meter_name")
		delete(additionalProperties, "meter_event_name")
		delete(additionalProperties, "aggregation")
		delete(additionalProperties, "unit_name")
		delete(additionalProperties, "total_quantity")
		delete(additionalProperties, "event_count")
		delete(additionalProperties, "last_event_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerMeter struct {
	value *CustomerMeter
	isSet bool
}

func (v NullableCustomerMeter) Get() *CustomerMeter {
	return v.value
}

func (v *NullableCustomerMeter) Set(val *CustomerMeter) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerMeter) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerMeter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerMeter(val *CustomerMeter) *NullableCustomerMeter {
	return &NullableCustomerMeter{value: val, isSet: true}
}

func (v NullableCustomerMeter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerMeter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


