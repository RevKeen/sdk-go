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

// checks if the UsageBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageBalanceResponse{}

// UsageBalanceResponse Current-period usage balance across all meters for a customer, with per-meter breakdowns and estimated totals.
type UsageBalanceResponse struct {
	Object string `json:"object"`
	Meters []UsageBalanceMeter `json:"meters"`
	// Total estimated charge across all meters
	TotalEstimatedAmountMinor float32 `json:"total_estimated_amount_minor"`
	// Total cost across all meters (null if no cost data)
	TotalCostMinor NullableFloat32 `json:"total_cost_minor"`
	// Currency code
	Currency string `json:"currency"`
	AdditionalProperties map[string]interface{}
}

type _UsageBalanceResponse UsageBalanceResponse

// NewUsageBalanceResponse instantiates a new UsageBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageBalanceResponse(object string, meters []UsageBalanceMeter, totalEstimatedAmountMinor float32, totalCostMinor NullableFloat32, currency string) *UsageBalanceResponse {
	this := UsageBalanceResponse{}
	this.Object = object
	this.Meters = meters
	this.TotalEstimatedAmountMinor = totalEstimatedAmountMinor
	this.TotalCostMinor = totalCostMinor
	this.Currency = currency
	return &this
}

// NewUsageBalanceResponseWithDefaults instantiates a new UsageBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageBalanceResponseWithDefaults() *UsageBalanceResponse {
	this := UsageBalanceResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *UsageBalanceResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *UsageBalanceResponse) SetObject(v string) {
	o.Object = v
}

// GetMeters returns the Meters field value
func (o *UsageBalanceResponse) GetMeters() []UsageBalanceMeter {
	if o == nil {
		var ret []UsageBalanceMeter
		return ret
	}

	return o.Meters
}

// GetMetersOk returns a tuple with the Meters field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceResponse) GetMetersOk() ([]UsageBalanceMeter, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meters, true
}

// SetMeters sets field value
func (o *UsageBalanceResponse) SetMeters(v []UsageBalanceMeter) {
	o.Meters = v
}

// GetTotalEstimatedAmountMinor returns the TotalEstimatedAmountMinor field value
func (o *UsageBalanceResponse) GetTotalEstimatedAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalEstimatedAmountMinor
}

// GetTotalEstimatedAmountMinorOk returns a tuple with the TotalEstimatedAmountMinor field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceResponse) GetTotalEstimatedAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalEstimatedAmountMinor, true
}

// SetTotalEstimatedAmountMinor sets field value
func (o *UsageBalanceResponse) SetTotalEstimatedAmountMinor(v float32) {
	o.TotalEstimatedAmountMinor = v
}

// GetTotalCostMinor returns the TotalCostMinor field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *UsageBalanceResponse) GetTotalCostMinor() float32 {
	if o == nil || o.TotalCostMinor.Get() == nil {
		var ret float32
		return ret
	}

	return *o.TotalCostMinor.Get()
}

// GetTotalCostMinorOk returns a tuple with the TotalCostMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageBalanceResponse) GetTotalCostMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCostMinor.Get(), o.TotalCostMinor.IsSet()
}

// SetTotalCostMinor sets field value
func (o *UsageBalanceResponse) SetTotalCostMinor(v float32) {
	o.TotalCostMinor.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *UsageBalanceResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UsageBalanceResponse) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UsageBalanceResponse) SetCurrency(v string) {
	o.Currency = v
}

func (o UsageBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["meters"] = o.Meters
	toSerialize["total_estimated_amount_minor"] = o.TotalEstimatedAmountMinor
	toSerialize["total_cost_minor"] = o.TotalCostMinor.Get()
	toSerialize["currency"] = o.Currency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"meters",
		"total_estimated_amount_minor",
		"total_cost_minor",
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

	varUsageBalanceResponse := _UsageBalanceResponse{}

	err = json.Unmarshal(data, &varUsageBalanceResponse)

	if err != nil {
		return err
	}

	*o = UsageBalanceResponse(varUsageBalanceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "meters")
		delete(additionalProperties, "total_estimated_amount_minor")
		delete(additionalProperties, "total_cost_minor")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageBalanceResponse struct {
	value *UsageBalanceResponse
	isSet bool
}

func (v NullableUsageBalanceResponse) Get() *UsageBalanceResponse {
	return v.value
}

func (v *NullableUsageBalanceResponse) Set(val *UsageBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageBalanceResponse(val *UsageBalanceResponse) *NullableUsageBalanceResponse {
	return &NullableUsageBalanceResponse{value: val, isSet: true}
}

func (v NullableUsageBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


