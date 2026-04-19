/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the CardDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardDetails{}

// CardDetails Card details (if type is 'card')
type CardDetails struct {
	// The card brand
	Brand NullableString `json:"brand"`
	// The last 4 digits of the card number
	Last4 NullableString `json:"last4"`
	// Expiration month (1-12)
	ExpMonth NullableInt32 `json:"exp_month"`
	// Expiration year
	ExpYear NullableInt32 `json:"exp_year"`
	// The card funding type
	Funding NullableString `json:"funding"`
	AdditionalProperties map[string]interface{}
}

type _CardDetails CardDetails

// NewCardDetails instantiates a new CardDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDetails(brand NullableString, last4 NullableString, expMonth NullableInt32, expYear NullableInt32, funding NullableString) *CardDetails {
	this := CardDetails{}
	this.Brand = brand
	this.Last4 = last4
	this.ExpMonth = expMonth
	this.ExpYear = expYear
	this.Funding = funding
	return &this
}

// NewCardDetailsWithDefaults instantiates a new CardDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDetailsWithDefaults() *CardDetails {
	this := CardDetails{}
	return &this
}

// GetBrand returns the Brand field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CardDetails) GetBrand() string {
	if o == nil || o.Brand.Get() == nil {
		var ret string
		return ret
	}

	return *o.Brand.Get()
}

// GetBrandOk returns a tuple with the Brand field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardDetails) GetBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brand.Get(), o.Brand.IsSet()
}

// SetBrand sets field value
func (o *CardDetails) SetBrand(v string) {
	o.Brand.Set(&v)
}

// GetLast4 returns the Last4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CardDetails) GetLast4() string {
	if o == nil || o.Last4.Get() == nil {
		var ret string
		return ret
	}

	return *o.Last4.Get()
}

// GetLast4Ok returns a tuple with the Last4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardDetails) GetLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last4.Get(), o.Last4.IsSet()
}

// SetLast4 sets field value
func (o *CardDetails) SetLast4(v string) {
	o.Last4.Set(&v)
}

// GetExpMonth returns the ExpMonth field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CardDetails) GetExpMonth() int32 {
	if o == nil || o.ExpMonth.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpMonth.Get()
}

// GetExpMonthOk returns a tuple with the ExpMonth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardDetails) GetExpMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpMonth.Get(), o.ExpMonth.IsSet()
}

// SetExpMonth sets field value
func (o *CardDetails) SetExpMonth(v int32) {
	o.ExpMonth.Set(&v)
}

// GetExpYear returns the ExpYear field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *CardDetails) GetExpYear() int32 {
	if o == nil || o.ExpYear.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpYear.Get()
}

// GetExpYearOk returns a tuple with the ExpYear field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardDetails) GetExpYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpYear.Get(), o.ExpYear.IsSet()
}

// SetExpYear sets field value
func (o *CardDetails) SetExpYear(v int32) {
	o.ExpYear.Set(&v)
}

// GetFunding returns the Funding field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CardDetails) GetFunding() string {
	if o == nil || o.Funding.Get() == nil {
		var ret string
		return ret
	}

	return *o.Funding.Get()
}

// GetFundingOk returns a tuple with the Funding field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardDetails) GetFundingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Funding.Get(), o.Funding.IsSet()
}

// SetFunding sets field value
func (o *CardDetails) SetFunding(v string) {
	o.Funding.Set(&v)
}

func (o CardDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["brand"] = o.Brand.Get()
	toSerialize["last4"] = o.Last4.Get()
	toSerialize["exp_month"] = o.ExpMonth.Get()
	toSerialize["exp_year"] = o.ExpYear.Get()
	toSerialize["funding"] = o.Funding.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CardDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"brand",
		"last4",
		"exp_month",
		"exp_year",
		"funding",
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

	varCardDetails := _CardDetails{}

	err = json.Unmarshal(data, &varCardDetails)

	if err != nil {
		return err
	}

	*o = CardDetails(varCardDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "brand")
		delete(additionalProperties, "last4")
		delete(additionalProperties, "exp_month")
		delete(additionalProperties, "exp_year")
		delete(additionalProperties, "funding")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCardDetails struct {
	value *CardDetails
	isSet bool
}

func (v NullableCardDetails) Get() *CardDetails {
	return v.value
}

func (v *NullableCardDetails) Set(val *CardDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDetails(val *CardDetails) *NullableCardDetails {
	return &NullableCardDetails{value: val, isSet: true}
}

func (v NullableCardDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


