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

// checks if the AnalyticsCustomersGetLtv200ResponseCustomersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsCustomersGetLtv200ResponseCustomersInner{}

// AnalyticsCustomersGetLtv200ResponseCustomersInner struct for AnalyticsCustomersGetLtv200ResponseCustomersInner
type AnalyticsCustomersGetLtv200ResponseCustomersInner struct {
	CustomerId string `json:"customerId"`
	CustomerName string `json:"customerName"`
	Ltv float32 `json:"ltv"`
	TotalRevenue float32 `json:"totalRevenue"`
	AvgOrderValue float32 `json:"avgOrderValue"`
	OrderCount float32 `json:"orderCount"`
	AdditionalProperties map[string]interface{}
}

type _AnalyticsCustomersGetLtv200ResponseCustomersInner AnalyticsCustomersGetLtv200ResponseCustomersInner

// NewAnalyticsCustomersGetLtv200ResponseCustomersInner instantiates a new AnalyticsCustomersGetLtv200ResponseCustomersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsCustomersGetLtv200ResponseCustomersInner(customerId string, customerName string, ltv float32, totalRevenue float32, avgOrderValue float32, orderCount float32) *AnalyticsCustomersGetLtv200ResponseCustomersInner {
	this := AnalyticsCustomersGetLtv200ResponseCustomersInner{}
	this.CustomerId = customerId
	this.CustomerName = customerName
	this.Ltv = ltv
	this.TotalRevenue = totalRevenue
	this.AvgOrderValue = avgOrderValue
	this.OrderCount = orderCount
	return &this
}

// NewAnalyticsCustomersGetLtv200ResponseCustomersInnerWithDefaults instantiates a new AnalyticsCustomersGetLtv200ResponseCustomersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsCustomersGetLtv200ResponseCustomersInnerWithDefaults() *AnalyticsCustomersGetLtv200ResponseCustomersInner {
	this := AnalyticsCustomersGetLtv200ResponseCustomersInner{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetCustomerName returns the CustomerName field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetCustomerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerName, true
}

// SetCustomerName sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetCustomerName(v string) {
	o.CustomerName = v
}

// GetLtv returns the Ltv field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetLtv() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Ltv
}

// GetLtvOk returns a tuple with the Ltv field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetLtvOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ltv, true
}

// SetLtv sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetLtv(v float32) {
	o.Ltv = v
}

// GetTotalRevenue returns the TotalRevenue field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetTotalRevenue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetTotalRevenueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRevenue, true
}

// SetTotalRevenue sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetTotalRevenue(v float32) {
	o.TotalRevenue = v
}

// GetAvgOrderValue returns the AvgOrderValue field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetAvgOrderValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AvgOrderValue
}

// GetAvgOrderValueOk returns a tuple with the AvgOrderValue field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetAvgOrderValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgOrderValue, true
}

// SetAvgOrderValue sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetAvgOrderValue(v float32) {
	o.AvgOrderValue = v
}

// GetOrderCount returns the OrderCount field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetOrderCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrderCount
}

// GetOrderCountOk returns a tuple with the OrderCount field value
// and a boolean to check if the value has been set.
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) GetOrderCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderCount, true
}

// SetOrderCount sets field value
func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) SetOrderCount(v float32) {
	o.OrderCount = v
}

func (o AnalyticsCustomersGetLtv200ResponseCustomersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsCustomersGetLtv200ResponseCustomersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerId"] = o.CustomerId
	toSerialize["customerName"] = o.CustomerName
	toSerialize["ltv"] = o.Ltv
	toSerialize["totalRevenue"] = o.TotalRevenue
	toSerialize["avgOrderValue"] = o.AvgOrderValue
	toSerialize["orderCount"] = o.OrderCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnalyticsCustomersGetLtv200ResponseCustomersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerId",
		"customerName",
		"ltv",
		"totalRevenue",
		"avgOrderValue",
		"orderCount",
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

	varAnalyticsCustomersGetLtv200ResponseCustomersInner := _AnalyticsCustomersGetLtv200ResponseCustomersInner{}

	err = json.Unmarshal(data, &varAnalyticsCustomersGetLtv200ResponseCustomersInner)

	if err != nil {
		return err
	}

	*o = AnalyticsCustomersGetLtv200ResponseCustomersInner(varAnalyticsCustomersGetLtv200ResponseCustomersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "customerName")
		delete(additionalProperties, "ltv")
		delete(additionalProperties, "totalRevenue")
		delete(additionalProperties, "avgOrderValue")
		delete(additionalProperties, "orderCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnalyticsCustomersGetLtv200ResponseCustomersInner struct {
	value *AnalyticsCustomersGetLtv200ResponseCustomersInner
	isSet bool
}

func (v NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) Get() *AnalyticsCustomersGetLtv200ResponseCustomersInner {
	return v.value
}

func (v *NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) Set(val *AnalyticsCustomersGetLtv200ResponseCustomersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsCustomersGetLtv200ResponseCustomersInner(val *AnalyticsCustomersGetLtv200ResponseCustomersInner) *NullableAnalyticsCustomersGetLtv200ResponseCustomersInner {
	return &NullableAnalyticsCustomersGetLtv200ResponseCustomersInner{value: val, isSet: true}
}

func (v NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsCustomersGetLtv200ResponseCustomersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


