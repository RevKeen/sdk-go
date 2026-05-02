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
	"fmt"
)

// checks if the BillingDetailsAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingDetailsAddress{}

// BillingDetailsAddress Billing address
type BillingDetailsAddress struct {
	Line1 NullableString `json:"line1"`
	Line2 NullableString `json:"line2"`
	City NullableString `json:"city"`
	State NullableString `json:"state"`
	PostalCode NullableString `json:"postal_code"`
	Country NullableString `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _BillingDetailsAddress BillingDetailsAddress

// NewBillingDetailsAddress instantiates a new BillingDetailsAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingDetailsAddress(line1 NullableString, line2 NullableString, city NullableString, state NullableString, postalCode NullableString, country NullableString) *BillingDetailsAddress {
	this := BillingDetailsAddress{}
	this.Line1 = line1
	this.Line2 = line2
	this.City = city
	this.State = state
	this.PostalCode = postalCode
	this.Country = country
	return &this
}

// NewBillingDetailsAddressWithDefaults instantiates a new BillingDetailsAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingDetailsAddressWithDefaults() *BillingDetailsAddress {
	this := BillingDetailsAddress{}
	return &this
}

// GetLine1 returns the Line1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetLine1() string {
	if o == nil || o.Line1.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// SetLine1 sets field value
func (o *BillingDetailsAddress) SetLine1(v string) {
	o.Line1.Set(&v)
}

// GetLine2 returns the Line2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetLine2() string {
	if o == nil || o.Line2.Get() == nil {
		var ret string
		return ret
	}

	return *o.Line2.Get()
}

// GetLine2Ok returns a tuple with the Line2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line2.Get(), o.Line2.IsSet()
}

// SetLine2 sets field value
func (o *BillingDetailsAddress) SetLine2(v string) {
	o.Line2.Set(&v)
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}

	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// SetCity sets field value
func (o *BillingDetailsAddress) SetCity(v string) {
	o.City.Set(&v)
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}

	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// SetState sets field value
func (o *BillingDetailsAddress) SetState(v string) {
	o.State.Set(&v)
}

// GetPostalCode returns the PostalCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// SetPostalCode sets field value
func (o *BillingDetailsAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BillingDetailsAddress) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingDetailsAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *BillingDetailsAddress) SetCountry(v string) {
	o.Country.Set(&v)
}

func (o BillingDetailsAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingDetailsAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["line1"] = o.Line1.Get()
	toSerialize["line2"] = o.Line2.Get()
	toSerialize["city"] = o.City.Get()
	toSerialize["state"] = o.State.Get()
	toSerialize["postal_code"] = o.PostalCode.Get()
	toSerialize["country"] = o.Country.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingDetailsAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"line1",
		"line2",
		"city",
		"state",
		"postal_code",
		"country",
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

	varBillingDetailsAddress := _BillingDetailsAddress{}

	err = json.Unmarshal(data, &varBillingDetailsAddress)

	if err != nil {
		return err
	}

	*o = BillingDetailsAddress(varBillingDetailsAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "line1")
		delete(additionalProperties, "line2")
		delete(additionalProperties, "city")
		delete(additionalProperties, "state")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingDetailsAddress struct {
	value *BillingDetailsAddress
	isSet bool
}

func (v NullableBillingDetailsAddress) Get() *BillingDetailsAddress {
	return v.value
}

func (v *NullableBillingDetailsAddress) Set(val *BillingDetailsAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingDetailsAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingDetailsAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingDetailsAddress(val *BillingDetailsAddress) *NullableBillingDetailsAddress {
	return &NullableBillingDetailsAddress{value: val, isSet: true}
}

func (v NullableBillingDetailsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingDetailsAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


