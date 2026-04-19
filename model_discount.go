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

// checks if the Discount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Discount{}

// Discount A discount code that applies a percentage or fixed amount reduction to charges, invoices, or subscriptions.
type Discount struct {
	Id string `json:"id"`
	Code string `json:"code"`
	Name NullableString `json:"name"`
	Description NullableString `json:"description"`
	DiscountType string `json:"discount_type"`
	// Amount off (cents) or percentage
	DiscountValue float32 `json:"discount_value"`
	AppliesTo NullableString `json:"applies_to"`
	ProductIds []string `json:"product_ids"`
	MaxRedemptions NullableFloat32 `json:"max_redemptions"`
	CurrentRedemptions NullableFloat32 `json:"current_redemptions"`
	ValidFrom NullableString `json:"valid_from"`
	ValidUntil NullableString `json:"valid_until"`
	RecurringType NullableString `json:"recurring_type"`
	RecurringCycles NullableFloat32 `json:"recurring_cycles"`
	FirstTimeCustomer NullableBool `json:"first_time_customer"`
	IsActive NullableBool `json:"is_active"`
	CreatedAt NullableString `json:"created_at"`
	UpdatedAt NullableString `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Discount Discount

// NewDiscount instantiates a new Discount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscount(id string, code string, name NullableString, description NullableString, discountType string, discountValue float32, appliesTo NullableString, productIds []string, maxRedemptions NullableFloat32, currentRedemptions NullableFloat32, validFrom NullableString, validUntil NullableString, recurringType NullableString, recurringCycles NullableFloat32, firstTimeCustomer NullableBool, isActive NullableBool, createdAt NullableString, updatedAt NullableString) *Discount {
	this := Discount{}
	this.Id = id
	this.Code = code
	this.Name = name
	this.Description = description
	this.DiscountType = discountType
	this.DiscountValue = discountValue
	this.AppliesTo = appliesTo
	this.ProductIds = productIds
	this.MaxRedemptions = maxRedemptions
	this.CurrentRedemptions = currentRedemptions
	this.ValidFrom = validFrom
	this.ValidUntil = validUntil
	this.RecurringType = recurringType
	this.RecurringCycles = recurringCycles
	this.FirstTimeCustomer = firstTimeCustomer
	this.IsActive = isActive
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDiscountWithDefaults instantiates a new Discount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscountWithDefaults() *Discount {
	this := Discount{}
	return &this
}

// GetId returns the Id field value
func (o *Discount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Discount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Discount) SetId(v string) {
	o.Id = v
}

// GetCode returns the Code field value
func (o *Discount) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Discount) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Discount) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *Discount) SetName(v string) {
	o.Name.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Discount) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetDiscountType returns the DiscountType field value
func (o *Discount) GetDiscountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value
// and a boolean to check if the value has been set.
func (o *Discount) GetDiscountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountType, true
}

// SetDiscountType sets field value
func (o *Discount) SetDiscountType(v string) {
	o.DiscountType = v
}

// GetDiscountValue returns the DiscountValue field value
func (o *Discount) GetDiscountValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiscountValue
}

// GetDiscountValueOk returns a tuple with the DiscountValue field value
// and a boolean to check if the value has been set.
func (o *Discount) GetDiscountValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountValue, true
}

// SetDiscountValue sets field value
func (o *Discount) SetDiscountValue(v float32) {
	o.DiscountValue = v
}

// GetAppliesTo returns the AppliesTo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetAppliesTo() string {
	if o == nil || o.AppliesTo.Get() == nil {
		var ret string
		return ret
	}

	return *o.AppliesTo.Get()
}

// GetAppliesToOk returns a tuple with the AppliesTo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetAppliesToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppliesTo.Get(), o.AppliesTo.IsSet()
}

// SetAppliesTo sets field value
func (o *Discount) SetAppliesTo(v string) {
	o.AppliesTo.Set(&v)
}

// GetProductIds returns the ProductIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *Discount) GetProductIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// SetProductIds sets field value
func (o *Discount) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetMaxRedemptions returns the MaxRedemptions field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Discount) GetMaxRedemptions() float32 {
	if o == nil || o.MaxRedemptions.Get() == nil {
		var ret float32
		return ret
	}

	return *o.MaxRedemptions.Get()
}

// GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetMaxRedemptionsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRedemptions.Get(), o.MaxRedemptions.IsSet()
}

// SetMaxRedemptions sets field value
func (o *Discount) SetMaxRedemptions(v float32) {
	o.MaxRedemptions.Set(&v)
}

// GetCurrentRedemptions returns the CurrentRedemptions field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Discount) GetCurrentRedemptions() float32 {
	if o == nil || o.CurrentRedemptions.Get() == nil {
		var ret float32
		return ret
	}

	return *o.CurrentRedemptions.Get()
}

// GetCurrentRedemptionsOk returns a tuple with the CurrentRedemptions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetCurrentRedemptionsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentRedemptions.Get(), o.CurrentRedemptions.IsSet()
}

// SetCurrentRedemptions sets field value
func (o *Discount) SetCurrentRedemptions(v float32) {
	o.CurrentRedemptions.Set(&v)
}

// GetValidFrom returns the ValidFrom field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetValidFrom() string {
	if o == nil || o.ValidFrom.Get() == nil {
		var ret string
		return ret
	}

	return *o.ValidFrom.Get()
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetValidFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidFrom.Get(), o.ValidFrom.IsSet()
}

// SetValidFrom sets field value
func (o *Discount) SetValidFrom(v string) {
	o.ValidFrom.Set(&v)
}

// GetValidUntil returns the ValidUntil field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetValidUntil() string {
	if o == nil || o.ValidUntil.Get() == nil {
		var ret string
		return ret
	}

	return *o.ValidUntil.Get()
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetValidUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidUntil.Get(), o.ValidUntil.IsSet()
}

// SetValidUntil sets field value
func (o *Discount) SetValidUntil(v string) {
	o.ValidUntil.Set(&v)
}

// GetRecurringType returns the RecurringType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetRecurringType() string {
	if o == nil || o.RecurringType.Get() == nil {
		var ret string
		return ret
	}

	return *o.RecurringType.Get()
}

// GetRecurringTypeOk returns a tuple with the RecurringType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetRecurringTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringType.Get(), o.RecurringType.IsSet()
}

// SetRecurringType sets field value
func (o *Discount) SetRecurringType(v string) {
	o.RecurringType.Set(&v)
}

// GetRecurringCycles returns the RecurringCycles field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Discount) GetRecurringCycles() float32 {
	if o == nil || o.RecurringCycles.Get() == nil {
		var ret float32
		return ret
	}

	return *o.RecurringCycles.Get()
}

// GetRecurringCyclesOk returns a tuple with the RecurringCycles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetRecurringCyclesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCycles.Get(), o.RecurringCycles.IsSet()
}

// SetRecurringCycles sets field value
func (o *Discount) SetRecurringCycles(v float32) {
	o.RecurringCycles.Set(&v)
}

// GetFirstTimeCustomer returns the FirstTimeCustomer field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Discount) GetFirstTimeCustomer() bool {
	if o == nil || o.FirstTimeCustomer.Get() == nil {
		var ret bool
		return ret
	}

	return *o.FirstTimeCustomer.Get()
}

// GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetFirstTimeCustomerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstTimeCustomer.Get(), o.FirstTimeCustomer.IsSet()
}

// SetFirstTimeCustomer sets field value
func (o *Discount) SetFirstTimeCustomer(v bool) {
	o.FirstTimeCustomer.Set(&v)
}

// GetIsActive returns the IsActive field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Discount) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// SetIsActive sets field value
func (o *Discount) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Discount) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Discount) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Discount) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *Discount) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

func (o Discount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Discount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["discount_type"] = o.DiscountType
	toSerialize["discount_value"] = o.DiscountValue
	toSerialize["applies_to"] = o.AppliesTo.Get()
	if o.ProductIds != nil {
		toSerialize["product_ids"] = o.ProductIds
	}
	toSerialize["max_redemptions"] = o.MaxRedemptions.Get()
	toSerialize["current_redemptions"] = o.CurrentRedemptions.Get()
	toSerialize["valid_from"] = o.ValidFrom.Get()
	toSerialize["valid_until"] = o.ValidUntil.Get()
	toSerialize["recurring_type"] = o.RecurringType.Get()
	toSerialize["recurring_cycles"] = o.RecurringCycles.Get()
	toSerialize["first_time_customer"] = o.FirstTimeCustomer.Get()
	toSerialize["is_active"] = o.IsActive.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Discount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"code",
		"name",
		"description",
		"discount_type",
		"discount_value",
		"applies_to",
		"product_ids",
		"max_redemptions",
		"current_redemptions",
		"valid_from",
		"valid_until",
		"recurring_type",
		"recurring_cycles",
		"first_time_customer",
		"is_active",
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

	varDiscount := _Discount{}

	err = json.Unmarshal(data, &varDiscount)

	if err != nil {
		return err
	}

	*o = Discount(varDiscount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "discount_type")
		delete(additionalProperties, "discount_value")
		delete(additionalProperties, "applies_to")
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "max_redemptions")
		delete(additionalProperties, "current_redemptions")
		delete(additionalProperties, "valid_from")
		delete(additionalProperties, "valid_until")
		delete(additionalProperties, "recurring_type")
		delete(additionalProperties, "recurring_cycles")
		delete(additionalProperties, "first_time_customer")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscount struct {
	value *Discount
	isSet bool
}

func (v NullableDiscount) Get() *Discount {
	return v.value
}

func (v *NullableDiscount) Set(val *Discount) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscount(val *Discount) *NullableDiscount {
	return &NullableDiscount{value: val, isSet: true}
}

func (v NullableDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


