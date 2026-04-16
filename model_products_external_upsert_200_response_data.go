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
	"time"
	"fmt"
)

// checks if the ProductsExternalUpsert200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductsExternalUpsert200ResponseData{}

// ProductsExternalUpsert200ResponseData struct for ProductsExternalUpsert200ResponseData
type ProductsExternalUpsert200ResponseData struct {
	Id string `json:"id"`
	MerchantId string `json:"merchantId"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Kind NullableString `json:"kind"`
	AmountCents NullableFloat32 `json:"amountCents"`
	Currency NullableString `json:"currency"`
	IsActive NullableBool `json:"isActive"`
	ExternalId NullableString `json:"externalId"`
	ExternalSystem NullableString `json:"externalSystem"`
	ExternalRef NullableString `json:"externalRef"`
	Interval NullableString `json:"interval"`
	IntervalCount NullableFloat32 `json:"intervalCount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _ProductsExternalUpsert200ResponseData ProductsExternalUpsert200ResponseData

// NewProductsExternalUpsert200ResponseData instantiates a new ProductsExternalUpsert200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductsExternalUpsert200ResponseData(id string, merchantId string, name string, description NullableString, kind NullableString, amountCents NullableFloat32, currency NullableString, isActive NullableBool, externalId NullableString, externalSystem NullableString, externalRef NullableString, interval NullableString, intervalCount NullableFloat32, createdAt time.Time, updatedAt time.Time) *ProductsExternalUpsert200ResponseData {
	this := ProductsExternalUpsert200ResponseData{}
	this.Id = id
	this.MerchantId = merchantId
	this.Name = name
	this.Description = description
	this.Kind = kind
	this.AmountCents = amountCents
	this.Currency = currency
	this.IsActive = isActive
	this.ExternalId = externalId
	this.ExternalSystem = externalSystem
	this.ExternalRef = externalRef
	this.Interval = interval
	this.IntervalCount = intervalCount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewProductsExternalUpsert200ResponseDataWithDefaults instantiates a new ProductsExternalUpsert200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductsExternalUpsert200ResponseDataWithDefaults() *ProductsExternalUpsert200ResponseData {
	this := ProductsExternalUpsert200ResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *ProductsExternalUpsert200ResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsert200ResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductsExternalUpsert200ResponseData) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *ProductsExternalUpsert200ResponseData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsert200ResponseData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *ProductsExternalUpsert200ResponseData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetName returns the Name field value
func (o *ProductsExternalUpsert200ResponseData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsert200ResponseData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductsExternalUpsert200ResponseData) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *ProductsExternalUpsert200ResponseData) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetKind returns the Kind field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetKind() string {
	if o == nil || o.Kind.Get() == nil {
		var ret string
		return ret
	}

	return *o.Kind.Get()
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kind.Get(), o.Kind.IsSet()
}

// SetKind sets field value
func (o *ProductsExternalUpsert200ResponseData) SetKind(v string) {
	o.Kind.Set(&v)
}

// GetAmountCents returns the AmountCents field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ProductsExternalUpsert200ResponseData) GetAmountCents() float32 {
	if o == nil || o.AmountCents.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AmountCents.Get()
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetAmountCentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCents.Get(), o.AmountCents.IsSet()
}

// SetAmountCents sets field value
func (o *ProductsExternalUpsert200ResponseData) SetAmountCents(v float32) {
	o.AmountCents.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *ProductsExternalUpsert200ResponseData) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetIsActive returns the IsActive field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ProductsExternalUpsert200ResponseData) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// SetIsActive sets field value
func (o *ProductsExternalUpsert200ResponseData) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *ProductsExternalUpsert200ResponseData) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetExternalSystem returns the ExternalSystem field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalSystem() string {
	if o == nil || o.ExternalSystem.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalSystem.Get()
}

// GetExternalSystemOk returns a tuple with the ExternalSystem field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalSystem.Get(), o.ExternalSystem.IsSet()
}

// SetExternalSystem sets field value
func (o *ProductsExternalUpsert200ResponseData) SetExternalSystem(v string) {
	o.ExternalSystem.Set(&v)
}

// GetExternalRef returns the ExternalRef field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalRef() string {
	if o == nil || o.ExternalRef.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalRef.Get()
}

// GetExternalRefOk returns a tuple with the ExternalRef field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetExternalRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalRef.Get(), o.ExternalRef.IsSet()
}

// SetExternalRef sets field value
func (o *ProductsExternalUpsert200ResponseData) SetExternalRef(v string) {
	o.ExternalRef.Set(&v)
}

// GetInterval returns the Interval field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductsExternalUpsert200ResponseData) GetInterval() string {
	if o == nil || o.Interval.Get() == nil {
		var ret string
		return ret
	}

	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// SetInterval sets field value
func (o *ProductsExternalUpsert200ResponseData) SetInterval(v string) {
	o.Interval.Set(&v)
}

// GetIntervalCount returns the IntervalCount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ProductsExternalUpsert200ResponseData) GetIntervalCount() float32 {
	if o == nil || o.IntervalCount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductsExternalUpsert200ResponseData) GetIntervalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// SetIntervalCount sets field value
func (o *ProductsExternalUpsert200ResponseData) SetIntervalCount(v float32) {
	o.IntervalCount.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProductsExternalUpsert200ResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsert200ResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProductsExternalUpsert200ResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProductsExternalUpsert200ResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsert200ResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProductsExternalUpsert200ResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ProductsExternalUpsert200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductsExternalUpsert200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["kind"] = o.Kind.Get()
	toSerialize["amountCents"] = o.AmountCents.Get()
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["isActive"] = o.IsActive.Get()
	toSerialize["externalId"] = o.ExternalId.Get()
	toSerialize["externalSystem"] = o.ExternalSystem.Get()
	toSerialize["externalRef"] = o.ExternalRef.Get()
	toSerialize["interval"] = o.Interval.Get()
	toSerialize["intervalCount"] = o.IntervalCount.Get()
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductsExternalUpsert200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchantId",
		"name",
		"description",
		"kind",
		"amountCents",
		"currency",
		"isActive",
		"externalId",
		"externalSystem",
		"externalRef",
		"interval",
		"intervalCount",
		"createdAt",
		"updatedAt",
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

	varProductsExternalUpsert200ResponseData := _ProductsExternalUpsert200ResponseData{}

	err = json.Unmarshal(data, &varProductsExternalUpsert200ResponseData)

	if err != nil {
		return err
	}

	*o = ProductsExternalUpsert200ResponseData(varProductsExternalUpsert200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchantId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "amountCents")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "externalSystem")
		delete(additionalProperties, "externalRef")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalCount")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductsExternalUpsert200ResponseData struct {
	value *ProductsExternalUpsert200ResponseData
	isSet bool
}

func (v NullableProductsExternalUpsert200ResponseData) Get() *ProductsExternalUpsert200ResponseData {
	return v.value
}

func (v *NullableProductsExternalUpsert200ResponseData) Set(val *ProductsExternalUpsert200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableProductsExternalUpsert200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableProductsExternalUpsert200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductsExternalUpsert200ResponseData(val *ProductsExternalUpsert200ResponseData) *NullableProductsExternalUpsert200ResponseData {
	return &NullableProductsExternalUpsert200ResponseData{value: val, isSet: true}
}

func (v NullableProductsExternalUpsert200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductsExternalUpsert200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


