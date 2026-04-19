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

// checks if the SubscriptionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionItem{}

// SubscriptionItem A line item within a subscription, representing a specific product and price that the customer is billed for each cycle.
type SubscriptionItem struct {
	Id string `json:"id"`
	SubscriptionId string `json:"subscription_id"`
	ProductId NullableString `json:"product_id"`
	PriceId NullableString `json:"price_id"`
	Description NullableString `json:"description"`
	// Item quantity
	Quantity int32 `json:"quantity"`
	Currency string `json:"currency"`
	// Unit price in cents
	UnitAmountMinor int32 `json:"unit_amount_minor"`
	// none: service (no order), physical: creates shipping order, digital: creates download order
	FulfillmentType string `json:"fulfillment_type"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionItem SubscriptionItem

// NewSubscriptionItem instantiates a new SubscriptionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionItem(id string, subscriptionId string, productId NullableString, priceId NullableString, description NullableString, quantity int32, currency string, unitAmountMinor int32, fulfillmentType string, createdAt string, updatedAt string) *SubscriptionItem {
	this := SubscriptionItem{}
	this.Id = id
	this.SubscriptionId = subscriptionId
	this.ProductId = productId
	this.PriceId = priceId
	this.Description = description
	this.Quantity = quantity
	this.Currency = currency
	this.UnitAmountMinor = unitAmountMinor
	this.FulfillmentType = fulfillmentType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionItemWithDefaults() *SubscriptionItem {
	this := SubscriptionItem{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionItem) SetId(v string) {
	o.Id = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *SubscriptionItem) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *SubscriptionItem) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetProductId returns the ProductId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionItem) GetProductId() string {
	if o == nil || o.ProductId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionItem) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// SetProductId sets field value
func (o *SubscriptionItem) SetProductId(v string) {
	o.ProductId.Set(&v)
}

// GetPriceId returns the PriceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionItem) GetPriceId() string {
	if o == nil || o.PriceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionItem) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// SetPriceId sets field value
func (o *SubscriptionItem) SetPriceId(v string) {
	o.PriceId.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *SubscriptionItem) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetQuantity returns the Quantity field value
func (o *SubscriptionItem) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SubscriptionItem) SetQuantity(v int32) {
	o.Quantity = v
}

// GetCurrency returns the Currency field value
func (o *SubscriptionItem) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SubscriptionItem) SetCurrency(v string) {
	o.Currency = v
}

// GetUnitAmountMinor returns the UnitAmountMinor field value
func (o *SubscriptionItem) GetUnitAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnitAmountMinor
}

// GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetUnitAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitAmountMinor, true
}

// SetUnitAmountMinor sets field value
func (o *SubscriptionItem) SetUnitAmountMinor(v int32) {
	o.UnitAmountMinor = v
}

// GetFulfillmentType returns the FulfillmentType field value
func (o *SubscriptionItem) GetFulfillmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentType
}

// GetFulfillmentTypeOk returns a tuple with the FulfillmentType field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetFulfillmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentType, true
}

// SetFulfillmentType sets field value
func (o *SubscriptionItem) SetFulfillmentType(v string) {
	o.FulfillmentType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubscriptionItem) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubscriptionItem) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SubscriptionItem) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SubscriptionItem) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SubscriptionItem) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SubscriptionItem) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SubscriptionItem) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o SubscriptionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["product_id"] = o.ProductId.Get()
	toSerialize["price_id"] = o.PriceId.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["quantity"] = o.Quantity
	toSerialize["currency"] = o.Currency
	toSerialize["unit_amount_minor"] = o.UnitAmountMinor
	toSerialize["fulfillment_type"] = o.FulfillmentType
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"subscription_id",
		"product_id",
		"price_id",
		"description",
		"quantity",
		"currency",
		"unit_amount_minor",
		"fulfillment_type",
		"created_at",
		"updated_at",
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

	varSubscriptionItem := _SubscriptionItem{}

	err = json.Unmarshal(data, &varSubscriptionItem)

	if err != nil {
		return err
	}

	*o = SubscriptionItem(varSubscriptionItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "price_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "unit_amount_minor")
		delete(additionalProperties, "fulfillment_type")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionItem struct {
	value *SubscriptionItem
	isSet bool
}

func (v NullableSubscriptionItem) Get() *SubscriptionItem {
	return v.value
}

func (v *NullableSubscriptionItem) Set(val *SubscriptionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionItem(val *SubscriptionItem) *NullableSubscriptionItem {
	return &NullableSubscriptionItem{value: val, isSet: true}
}

func (v NullableSubscriptionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


