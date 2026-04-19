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

// checks if the PaymentLinksCreateRequestLineItemsInnerPriceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentLinksCreateRequestLineItemsInnerPriceData{}

// PaymentLinksCreateRequestLineItemsInnerPriceData struct for PaymentLinksCreateRequestLineItemsInnerPriceData
type PaymentLinksCreateRequestLineItemsInnerPriceData struct {
	Currency *string `json:"currency,omitempty"`
	UnitAmount int32 `json:"unit_amount"`
	ProductData *PaymentLinksCreateRequestLineItemsInnerPriceDataProductData `json:"product_data,omitempty"`
	Product *string `json:"product,omitempty"`
	Recurring *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring `json:"recurring,omitempty"`
	TaxBehavior NullableString `json:"tax_behavior,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentLinksCreateRequestLineItemsInnerPriceData PaymentLinksCreateRequestLineItemsInnerPriceData

// NewPaymentLinksCreateRequestLineItemsInnerPriceData instantiates a new PaymentLinksCreateRequestLineItemsInnerPriceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinksCreateRequestLineItemsInnerPriceData(unitAmount int32) *PaymentLinksCreateRequestLineItemsInnerPriceData {
	this := PaymentLinksCreateRequestLineItemsInnerPriceData{}
	var currency string = "USD"
	this.Currency = &currency
	this.UnitAmount = unitAmount
	return &this
}

// NewPaymentLinksCreateRequestLineItemsInnerPriceDataWithDefaults instantiates a new PaymentLinksCreateRequestLineItemsInnerPriceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinksCreateRequestLineItemsInnerPriceDataWithDefaults() *PaymentLinksCreateRequestLineItemsInnerPriceData {
	this := PaymentLinksCreateRequestLineItemsInnerPriceData{}
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetCurrency(v string) {
	o.Currency = &v
}

// GetUnitAmount returns the UnitAmount field value
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetUnitAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetUnitAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmount, true
}

// SetUnitAmount sets field value
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetUnitAmount(v int32) {
	o.UnitAmount = v
}

// GetProductData returns the ProductData field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductData() PaymentLinksCreateRequestLineItemsInnerPriceDataProductData {
	if o == nil || IsNil(o.ProductData) {
		var ret PaymentLinksCreateRequestLineItemsInnerPriceDataProductData
		return ret
	}
	return *o.ProductData
}

// GetProductDataOk returns a tuple with the ProductData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductDataOk() (*PaymentLinksCreateRequestLineItemsInnerPriceDataProductData, bool) {
	if o == nil || IsNil(o.ProductData) {
		return nil, false
	}
	return o.ProductData, true
}

// HasProductData returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasProductData() bool {
	if o != nil && !IsNil(o.ProductData) {
		return true
	}

	return false
}

// SetProductData gets a reference to the given PaymentLinksCreateRequestLineItemsInnerPriceDataProductData and assigns it to the ProductData field.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetProductData(v PaymentLinksCreateRequestLineItemsInnerPriceDataProductData) {
	o.ProductData = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetProduct(v string) {
	o.Product = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetRecurring() PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring {
	if o == nil || IsNil(o.Recurring) {
		var ret PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetRecurringOk() (*PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring, bool) {
	if o == nil || IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasRecurring() bool {
	if o != nil && !IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring and assigns it to the Recurring field.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetRecurring(v PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring) {
	o.Recurring = &v
}

// GetTaxBehavior returns the TaxBehavior field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetTaxBehavior() string {
	if o == nil || IsNil(o.TaxBehavior.Get()) {
		var ret string
		return ret
	}
	return *o.TaxBehavior.Get()
}

// GetTaxBehaviorOk returns a tuple with the TaxBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetTaxBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxBehavior.Get(), o.TaxBehavior.IsSet()
}

// HasTaxBehavior returns a boolean if a field has been set.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasTaxBehavior() bool {
	if o != nil && o.TaxBehavior.IsSet() {
		return true
	}

	return false
}

// SetTaxBehavior gets a reference to the given NullableString and assigns it to the TaxBehavior field.
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetTaxBehavior(v string) {
	o.TaxBehavior.Set(&v)
}
// SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetTaxBehaviorNil() {
	o.TaxBehavior.Set(nil)
}

// UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil
func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) UnsetTaxBehavior() {
	o.TaxBehavior.Unset()
}

func (o PaymentLinksCreateRequestLineItemsInnerPriceData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentLinksCreateRequestLineItemsInnerPriceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	toSerialize["unit_amount"] = o.UnitAmount
	if !IsNil(o.ProductData) {
		toSerialize["product_data"] = o.ProductData
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	if o.TaxBehavior.IsSet() {
		toSerialize["tax_behavior"] = o.TaxBehavior.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unit_amount",
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

	varPaymentLinksCreateRequestLineItemsInnerPriceData := _PaymentLinksCreateRequestLineItemsInnerPriceData{}

	err = json.Unmarshal(data, &varPaymentLinksCreateRequestLineItemsInnerPriceData)

	if err != nil {
		return err
	}

	*o = PaymentLinksCreateRequestLineItemsInnerPriceData(varPaymentLinksCreateRequestLineItemsInnerPriceData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currency")
		delete(additionalProperties, "unit_amount")
		delete(additionalProperties, "product_data")
		delete(additionalProperties, "product")
		delete(additionalProperties, "recurring")
		delete(additionalProperties, "tax_behavior")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentLinksCreateRequestLineItemsInnerPriceData struct {
	value *PaymentLinksCreateRequestLineItemsInnerPriceData
	isSet bool
}

func (v NullablePaymentLinksCreateRequestLineItemsInnerPriceData) Get() *PaymentLinksCreateRequestLineItemsInnerPriceData {
	return v.value
}

func (v *NullablePaymentLinksCreateRequestLineItemsInnerPriceData) Set(val *PaymentLinksCreateRequestLineItemsInnerPriceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinksCreateRequestLineItemsInnerPriceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinksCreateRequestLineItemsInnerPriceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinksCreateRequestLineItemsInnerPriceData(val *PaymentLinksCreateRequestLineItemsInnerPriceData) *NullablePaymentLinksCreateRequestLineItemsInnerPriceData {
	return &NullablePaymentLinksCreateRequestLineItemsInnerPriceData{value: val, isSet: true}
}

func (v NullablePaymentLinksCreateRequestLineItemsInnerPriceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinksCreateRequestLineItemsInnerPriceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


