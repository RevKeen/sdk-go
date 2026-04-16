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

// checks if the ProductsExternalUpsertRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductsExternalUpsertRequest{}

// ProductsExternalUpsertRequest struct for ProductsExternalUpsertRequest
type ProductsExternalUpsertRequest struct {
	// Deprecated: merchantId is derived from API key authentication. This field is ignored.
	// Deprecated
	MerchantId *string `json:"merchantId,omitempty"`
	// Product name
	Name string `json:"name"`
	// Product description
	Description *string `json:"description,omitempty"`
	// Product kind - subscription or one-time
	Kind *string `json:"kind,omitempty"`
	// Price in minor units (cents)
	AmountCents int32 `json:"amountCents"`
	// ISO 4217 currency code
	Currency *string `json:"currency,omitempty"`
	// Whether the product is active
	IsActive *bool `json:"isActive,omitempty"`
	// Billing interval for subscriptions
	Interval *string `json:"interval,omitempty"`
	// Number of intervals between billings
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// External system's last update timestamp for stale update protection
	ExternalUpdatedAt *time.Time `json:"externalUpdatedAt,omitempty"`
	// Additional external reference (e.g., membership ID)
	ExternalRef *string `json:"externalRef,omitempty"`
	// Additional metadata from external system
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductsExternalUpsertRequest ProductsExternalUpsertRequest

// NewProductsExternalUpsertRequest instantiates a new ProductsExternalUpsertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductsExternalUpsertRequest(name string, amountCents int32) *ProductsExternalUpsertRequest {
	this := ProductsExternalUpsertRequest{}
	this.Name = name
	this.AmountCents = amountCents
	var currency string = "USD"
	this.Currency = &currency
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// NewProductsExternalUpsertRequestWithDefaults instantiates a new ProductsExternalUpsertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductsExternalUpsertRequestWithDefaults() *ProductsExternalUpsertRequest {
	this := ProductsExternalUpsertRequest{}
	var currency string = "USD"
	this.Currency = &currency
	var isActive bool = true
	this.IsActive = &isActive
	return &this
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
// Deprecated
func (o *ProductsExternalUpsertRequest) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ProductsExternalUpsertRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
// Deprecated
func (o *ProductsExternalUpsertRequest) SetMerchantId(v string) {
	o.MerchantId = &v
}

// GetName returns the Name field value
func (o *ProductsExternalUpsertRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductsExternalUpsertRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductsExternalUpsertRequest) SetDescription(v string) {
	o.Description = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ProductsExternalUpsertRequest) SetKind(v string) {
	o.Kind = &v
}

// GetAmountCents returns the AmountCents field value
func (o *ProductsExternalUpsertRequest) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *ProductsExternalUpsertRequest) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ProductsExternalUpsertRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ProductsExternalUpsertRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *ProductsExternalUpsertRequest) SetInterval(v string) {
	o.Interval = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *ProductsExternalUpsertRequest) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetExternalUpdatedAt returns the ExternalUpdatedAt field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetExternalUpdatedAt() time.Time {
	if o == nil || IsNil(o.ExternalUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ExternalUpdatedAt
}

// GetExternalUpdatedAtOk returns a tuple with the ExternalUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetExternalUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExternalUpdatedAt) {
		return nil, false
	}
	return o.ExternalUpdatedAt, true
}

// HasExternalUpdatedAt returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasExternalUpdatedAt() bool {
	if o != nil && !IsNil(o.ExternalUpdatedAt) {
		return true
	}

	return false
}

// SetExternalUpdatedAt gets a reference to the given time.Time and assigns it to the ExternalUpdatedAt field.
func (o *ProductsExternalUpsertRequest) SetExternalUpdatedAt(v time.Time) {
	o.ExternalUpdatedAt = &v
}

// GetExternalRef returns the ExternalRef field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetExternalRef() string {
	if o == nil || IsNil(o.ExternalRef) {
		var ret string
		return ret
	}
	return *o.ExternalRef
}

// GetExternalRefOk returns a tuple with the ExternalRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetExternalRefOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalRef) {
		return nil, false
	}
	return o.ExternalRef, true
}

// HasExternalRef returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasExternalRef() bool {
	if o != nil && !IsNil(o.ExternalRef) {
		return true
	}

	return false
}

// SetExternalRef gets a reference to the given string and assigns it to the ExternalRef field.
func (o *ProductsExternalUpsertRequest) SetExternalRef(v string) {
	o.ExternalRef = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProductsExternalUpsertRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductsExternalUpsertRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProductsExternalUpsertRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ProductsExternalUpsertRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ProductsExternalUpsertRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductsExternalUpsertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["amountCents"] = o.AmountCents
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.IntervalCount) {
		toSerialize["intervalCount"] = o.IntervalCount
	}
	if !IsNil(o.ExternalUpdatedAt) {
		toSerialize["externalUpdatedAt"] = o.ExternalUpdatedAt
	}
	if !IsNil(o.ExternalRef) {
		toSerialize["externalRef"] = o.ExternalRef
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductsExternalUpsertRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"amountCents",
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

	varProductsExternalUpsertRequest := _ProductsExternalUpsertRequest{}

	err = json.Unmarshal(data, &varProductsExternalUpsertRequest)

	if err != nil {
		return err
	}

	*o = ProductsExternalUpsertRequest(varProductsExternalUpsertRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "merchantId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "amountCents")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "intervalCount")
		delete(additionalProperties, "externalUpdatedAt")
		delete(additionalProperties, "externalRef")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductsExternalUpsertRequest struct {
	value *ProductsExternalUpsertRequest
	isSet bool
}

func (v NullableProductsExternalUpsertRequest) Get() *ProductsExternalUpsertRequest {
	return v.value
}

func (v *NullableProductsExternalUpsertRequest) Set(val *ProductsExternalUpsertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductsExternalUpsertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductsExternalUpsertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductsExternalUpsertRequest(val *ProductsExternalUpsertRequest) *NullableProductsExternalUpsertRequest {
	return &NullableProductsExternalUpsertRequest{value: val, isSet: true}
}

func (v NullableProductsExternalUpsertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductsExternalUpsertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


