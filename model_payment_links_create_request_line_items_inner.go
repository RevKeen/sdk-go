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

// checks if the PaymentLinksCreateRequestLineItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinksCreateRequestLineItemsInner{}

// PaymentLinksCreateRequestLineItemsInner struct for PaymentLinksCreateRequestLineItemsInner
type PaymentLinksCreateRequestLineItemsInner struct {
	Price *string `json:"price,omitempty"`
	PriceData *PaymentLinksCreateRequestLineItemsInnerPriceData `json:"price_data,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	MerchantReference NullableString `json:"merchant_reference,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinksCreateRequestLineItemsInner PaymentLinksCreateRequestLineItemsInner

// NewPaymentLinksCreateRequestLineItemsInner instantiates a new PaymentLinksCreateRequestLineItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinksCreateRequestLineItemsInner() *PaymentLinksCreateRequestLineItemsInner {
	this := PaymentLinksCreateRequestLineItemsInner{}
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// NewPaymentLinksCreateRequestLineItemsInnerWithDefaults instantiates a new PaymentLinksCreateRequestLineItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksCreateRequestLineItemsInnerWithDefaults() *PaymentLinksCreateRequestLineItemsInner {
	this := PaymentLinksCreateRequestLineItemsInner{}
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInner) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PaymentLinksCreateRequestLineItemsInner) SetPrice(v string) {
	o.Price = &v
}

// GetPriceData returns the PriceData field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceData() PaymentLinksCreateRequestLineItemsInnerPriceData {
	if o == nil || IsNil(o.PriceData) {
		var ret PaymentLinksCreateRequestLineItemsInnerPriceData
		return ret
	}
	return *o.PriceData
}

// GetPriceDataOk returns a tuple with the PriceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceDataOk() (*PaymentLinksCreateRequestLineItemsInnerPriceData, bool) {
	if o == nil || IsNil(o.PriceData) {
		return nil, false
	}
	return o.PriceData, true
}

// HasPriceData returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) HasPriceData() bool {
	if o != nil && !IsNil(o.PriceData) {
		return true
	}

	return false
}

// SetPriceData gets a reference to the given PaymentLinksCreateRequestLineItemsInnerPriceData and assigns it to the PriceData field.
func (o *PaymentLinksCreateRequestLineItemsInner) SetPriceData(v PaymentLinksCreateRequestLineItemsInnerPriceData) {
	o.PriceData = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PaymentLinksCreateRequestLineItemsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksCreateRequestLineItemsInner) GetMerchantReference() string {
	if o == nil || IsNil(o.MerchantReference.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantReference.Get()
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksCreateRequestLineItemsInner) GetMerchantReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantReference.Get(), o.MerchantReference.IsSet()
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInner) HasMerchantReference() bool {
	if o != nil && o.MerchantReference.IsSet() {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given NullableString and assigns it to the MerchantReference field.
func (o *PaymentLinksCreateRequestLineItemsInner) SetMerchantReference(v string) {
	o.MerchantReference.Set(&v)
}
// SetMerchantReferenceNil sets the value for MerchantReference to be an explicit nil
func (o *PaymentLinksCreateRequestLineItemsInner) SetMerchantReferenceNil() {
	o.MerchantReference.Set(nil)
}

// UnsetMerchantReference ensures that no value is present for MerchantReference, not even an explicit nil
func (o *PaymentLinksCreateRequestLineItemsInner) UnsetMerchantReference() {
	o.MerchantReference.Unset()
}

func (o PaymentLinksCreateRequestLineItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinksCreateRequestLineItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceData) {
		toSerialize["price_data"] = o.PriceData
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if o.MerchantReference.IsSet() {
		toSerialize["merchant_reference"] = o.MerchantReference.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinksCreateRequestLineItemsInner) UnmarshalJSON(data []byte) (err error) {
	varPaymentLinksCreateRequestLineItemsInner := _PaymentLinksCreateRequestLineItemsInner{}

	err = json.Unmarshal(data, &varPaymentLinksCreateRequestLineItemsInner)

	if err != nil {
		return err
	}

	*o = PaymentLinksCreateRequestLineItemsInner(varPaymentLinksCreateRequestLineItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "price_data")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "merchant_reference")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentLinksCreateRequestLineItemsInner struct {
	value *PaymentLinksCreateRequestLineItemsInner
	isSet bool
}

func (v NullablePaymentLinksCreateRequestLineItemsInner) Get() *PaymentLinksCreateRequestLineItemsInner {
	return v.value
}

func (v *NullablePaymentLinksCreateRequestLineItemsInner) Set(val *PaymentLinksCreateRequestLineItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinksCreateRequestLineItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinksCreateRequestLineItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinksCreateRequestLineItemsInner(val *PaymentLinksCreateRequestLineItemsInner) *NullablePaymentLinksCreateRequestLineItemsInner {
	return &NullablePaymentLinksCreateRequestLineItemsInner{value: val, isSet: true}
}

func (v NullablePaymentLinksCreateRequestLineItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinksCreateRequestLineItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


