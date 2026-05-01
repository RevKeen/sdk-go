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

// checks if the PaymentLinksList200ResponseDataInnerLineItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinksList200ResponseDataInnerLineItemsInner{}

// PaymentLinksList200ResponseDataInnerLineItemsInner struct for PaymentLinksList200ResponseDataInnerLineItemsInner
type PaymentLinksList200ResponseDataInnerLineItemsInner struct {
	ProductId string `json:"product_id"`
	PriceId NullableString `json:"price_id,omitempty"`
	Quantity int32 `json:"quantity"`
	MerchantReference NullableString `json:"merchant_reference,omitempty"`
	IsAdhoc bool `json:"is_adhoc"`
	ProductDetails *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails `json:"product_details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinksList200ResponseDataInnerLineItemsInner PaymentLinksList200ResponseDataInnerLineItemsInner

// NewPaymentLinksList200ResponseDataInnerLineItemsInner instantiates a new PaymentLinksList200ResponseDataInnerLineItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinksList200ResponseDataInnerLineItemsInner(productId string, quantity int32, isAdhoc bool) *PaymentLinksList200ResponseDataInnerLineItemsInner {
	this := PaymentLinksList200ResponseDataInnerLineItemsInner{}
	this.ProductId = productId
	this.Quantity = quantity
	this.IsAdhoc = isAdhoc
	return &this
}

// NewPaymentLinksList200ResponseDataInnerLineItemsInnerWithDefaults instantiates a new PaymentLinksList200ResponseDataInnerLineItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksList200ResponseDataInnerLineItemsInnerWithDefaults() *PaymentLinksList200ResponseDataInnerLineItemsInner {
	this := PaymentLinksList200ResponseDataInnerLineItemsInner{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetProductId(v string) {
	o.ProductId = v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetPriceId() string {
	if o == nil || IsNil(o.PriceId.Get()) {
		var ret string
		return ret
	}
	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// HasPriceId returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) HasPriceId() bool {
	if o != nil && o.PriceId.IsSet() {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given NullableString and assigns it to the PriceId field.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetPriceId(v string) {
	o.PriceId.Set(&v)
}
// SetPriceIdNil sets the value for PriceId to be an explicit nil
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetPriceIdNil() {
	o.PriceId.Set(nil)
}

// UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) UnsetPriceId() {
	o.PriceId.Unset()
}

// GetQuantity returns the Quantity field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetQuantity(v int32) {
	o.Quantity = v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetMerchantReference() string {
	if o == nil || IsNil(o.MerchantReference.Get()) {
		var ret string
		return ret
	}
	return *o.MerchantReference.Get()
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetMerchantReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MerchantReference.Get(), o.MerchantReference.IsSet()
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) HasMerchantReference() bool {
	if o != nil && o.MerchantReference.IsSet() {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given NullableString and assigns it to the MerchantReference field.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetMerchantReference(v string) {
	o.MerchantReference.Set(&v)
}
// SetMerchantReferenceNil sets the value for MerchantReference to be an explicit nil
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetMerchantReferenceNil() {
	o.MerchantReference.Set(nil)
}

// UnsetMerchantReference ensures that no value is present for MerchantReference, not even an explicit nil
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) UnsetMerchantReference() {
	o.MerchantReference.Unset()
}

// GetIsAdhoc returns the IsAdhoc field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetIsAdhoc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAdhoc
}

// GetIsAdhocOk returns a tuple with the IsAdhoc field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetIsAdhocOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAdhoc, true
}

// SetIsAdhoc sets field value
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetIsAdhoc(v bool) {
	o.IsAdhoc = v
}

// GetProductDetails returns the ProductDetails field value if set, zero value otherwise.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetProductDetails() PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails {
	if o == nil || IsNil(o.ProductDetails) {
		var ret PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails
		return ret
	}
	return *o.ProductDetails
}

// GetProductDetailsOk returns a tuple with the ProductDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) GetProductDetailsOk() (*PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails, bool) {
	if o == nil || IsNil(o.ProductDetails) {
		return nil, false
	}
	return o.ProductDetails, true
}

// HasProductDetails returns a boolean if a field has been set.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) HasProductDetails() bool {
	if o != nil && !IsNil(o.ProductDetails) {
		return true
	}

	return false
}

// SetProductDetails gets a reference to the given PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails and assigns it to the ProductDetails field.
func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) SetProductDetails(v PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) {
	o.ProductDetails = &v
}

func (o PaymentLinksList200ResponseDataInnerLineItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinksList200ResponseDataInnerLineItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_id"] = o.ProductId
	if o.PriceId.IsSet() {
		toSerialize["price_id"] = o.PriceId.Get()
	}
	toSerialize["quantity"] = o.Quantity
	if o.MerchantReference.IsSet() {
		toSerialize["merchant_reference"] = o.MerchantReference.Get()
	}
	toSerialize["is_adhoc"] = o.IsAdhoc
	if !IsNil(o.ProductDetails) {
		toSerialize["product_details"] = o.ProductDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinksList200ResponseDataInnerLineItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_id",
		"quantity",
		"is_adhoc",
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

	varPaymentLinksList200ResponseDataInnerLineItemsInner := _PaymentLinksList200ResponseDataInnerLineItemsInner{}

	err = json.Unmarshal(data, &varPaymentLinksList200ResponseDataInnerLineItemsInner)

	if err != nil {
		return err
	}

	*o = PaymentLinksList200ResponseDataInnerLineItemsInner(varPaymentLinksList200ResponseDataInnerLineItemsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "price_id")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "merchant_reference")
		delete(additionalProperties, "is_adhoc")
		delete(additionalProperties, "product_details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentLinksList200ResponseDataInnerLineItemsInner struct {
	value *PaymentLinksList200ResponseDataInnerLineItemsInner
	isSet bool
}

func (v NullablePaymentLinksList200ResponseDataInnerLineItemsInner) Get() *PaymentLinksList200ResponseDataInnerLineItemsInner {
	return v.value
}

func (v *NullablePaymentLinksList200ResponseDataInnerLineItemsInner) Set(val *PaymentLinksList200ResponseDataInnerLineItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinksList200ResponseDataInnerLineItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinksList200ResponseDataInnerLineItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinksList200ResponseDataInnerLineItemsInner(val *PaymentLinksList200ResponseDataInnerLineItemsInner) *NullablePaymentLinksList200ResponseDataInnerLineItemsInner {
	return &NullablePaymentLinksList200ResponseDataInnerLineItemsInner{value: val, isSet: true}
}

func (v NullablePaymentLinksList200ResponseDataInnerLineItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinksList200ResponseDataInnerLineItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


