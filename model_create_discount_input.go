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

// checks if the CreateDiscountInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDiscountInput{}

// CreateDiscountInput Parameters for creating a new discount code with amount, type, and usage limits.
type CreateDiscountInput struct {
	// Unique discount code (will be uppercased)
	Code string `json:"code"`
	// Human-readable name
	Name string `json:"name"`
	// Optional description
	Description *string `json:"description,omitempty"`
	// Type of discount
	DiscountType string `json:"discount_type"`
	// Discount value - percentage (1-100) or amount in cents
	DiscountValue float32 `json:"discount_value"`
	// Which products the discount applies to
	AppliesTo *string `json:"applies_to,omitempty"`
	// Scope of where discount can be applied
	Scope *string `json:"scope,omitempty"`
	// Product IDs if applies_to is 'specific_products'
	ProductIds []string `json:"product_ids,omitempty"`
	// Maximum total redemptions allowed
	MaxRedemptions *int32 `json:"max_redemptions,omitempty"`
	// Maximum redemptions per customer (0 = unlimited)
	MaxRedemptionsPerUser *int32 `json:"max_redemptions_per_user,omitempty"`
	// Start date (ISO 8601)
	ValidFrom *string `json:"valid_from,omitempty"`
	// End date (ISO 8601)
	ValidUntil *string `json:"valid_until,omitempty"`
	// How the discount applies to recurring payments
	RecurringType *string `json:"recurring_type,omitempty"`
	// Number of billing cycles (only if recurring_type is 'repeating')
	RecurringCycles *int32 `json:"recurring_cycles,omitempty"`
	// Only available to first-time customers
	FirstTimeCustomer *bool `json:"first_time_customer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDiscountInput CreateDiscountInput

// NewCreateDiscountInput instantiates a new CreateDiscountInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDiscountInput(code string, name string, discountType string, discountValue float32) *CreateDiscountInput {
	this := CreateDiscountInput{}
	this.Code = code
	this.Name = name
	this.DiscountType = discountType
	this.DiscountValue = discountValue
	var appliesTo string = "all"
	this.AppliesTo = &appliesTo
	var maxRedemptionsPerUser int32 = 0
	this.MaxRedemptionsPerUser = &maxRedemptionsPerUser
	var firstTimeCustomer bool = false
	this.FirstTimeCustomer = &firstTimeCustomer
	return &this
}

// NewCreateDiscountInputWithDefaults instantiates a new CreateDiscountInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDiscountInputWithDefaults() *CreateDiscountInput {
	this := CreateDiscountInput{}
	var appliesTo string = "all"
	this.AppliesTo = &appliesTo
	var maxRedemptionsPerUser int32 = 0
	this.MaxRedemptionsPerUser = &maxRedemptionsPerUser
	var firstTimeCustomer bool = false
	this.FirstTimeCustomer = &firstTimeCustomer
	return &this
}

// GetCode returns the Code field value
func (o *CreateDiscountInput) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CreateDiscountInput) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value
func (o *CreateDiscountInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDiscountInput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDiscountInput) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountType returns the DiscountType field value
func (o *CreateDiscountInput) GetDiscountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetDiscountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountType, true
}

// SetDiscountType sets field value
func (o *CreateDiscountInput) SetDiscountType(v string) {
	o.DiscountType = v
}

// GetDiscountValue returns the DiscountValue field value
func (o *CreateDiscountInput) GetDiscountValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DiscountValue
}

// GetDiscountValueOk returns a tuple with the DiscountValue field value
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetDiscountValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountValue, true
}

// SetDiscountValue sets field value
func (o *CreateDiscountInput) SetDiscountValue(v float32) {
	o.DiscountValue = v
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetAppliesTo() string {
	if o == nil || IsNil(o.AppliesTo) {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetAppliesToOk() (*string, bool) {
	if o == nil || IsNil(o.AppliesTo) {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasAppliesTo() bool {
	if o != nil && !IsNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *CreateDiscountInput) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *CreateDiscountInput) SetScope(v string) {
	o.Scope = &v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetProductIds() []string {
	if o == nil || IsNil(o.ProductIds) {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *CreateDiscountInput) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetMaxRedemptions returns the MaxRedemptions field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetMaxRedemptions() int32 {
	if o == nil || IsNil(o.MaxRedemptions) {
		var ret int32
		return ret
	}
	return *o.MaxRedemptions
}

// GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetMaxRedemptionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRedemptions) {
		return nil, false
	}
	return o.MaxRedemptions, true
}

// HasMaxRedemptions returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasMaxRedemptions() bool {
	if o != nil && !IsNil(o.MaxRedemptions) {
		return true
	}

	return false
}

// SetMaxRedemptions gets a reference to the given int32 and assigns it to the MaxRedemptions field.
func (o *CreateDiscountInput) SetMaxRedemptions(v int32) {
	o.MaxRedemptions = &v
}

// GetMaxRedemptionsPerUser returns the MaxRedemptionsPerUser field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetMaxRedemptionsPerUser() int32 {
	if o == nil || IsNil(o.MaxRedemptionsPerUser) {
		var ret int32
		return ret
	}
	return *o.MaxRedemptionsPerUser
}

// GetMaxRedemptionsPerUserOk returns a tuple with the MaxRedemptionsPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetMaxRedemptionsPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRedemptionsPerUser) {
		return nil, false
	}
	return o.MaxRedemptionsPerUser, true
}

// HasMaxRedemptionsPerUser returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasMaxRedemptionsPerUser() bool {
	if o != nil && !IsNil(o.MaxRedemptionsPerUser) {
		return true
	}

	return false
}

// SetMaxRedemptionsPerUser gets a reference to the given int32 and assigns it to the MaxRedemptionsPerUser field.
func (o *CreateDiscountInput) SetMaxRedemptionsPerUser(v int32) {
	o.MaxRedemptionsPerUser = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *CreateDiscountInput) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetValidUntil() string {
	if o == nil || IsNil(o.ValidUntil) {
		var ret string
		return ret
	}
	return *o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetValidUntilOk() (*string, bool) {
	if o == nil || IsNil(o.ValidUntil) {
		return nil, false
	}
	return o.ValidUntil, true
}

// HasValidUntil returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasValidUntil() bool {
	if o != nil && !IsNil(o.ValidUntil) {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given string and assigns it to the ValidUntil field.
func (o *CreateDiscountInput) SetValidUntil(v string) {
	o.ValidUntil = &v
}

// GetRecurringType returns the RecurringType field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetRecurringType() string {
	if o == nil || IsNil(o.RecurringType) {
		var ret string
		return ret
	}
	return *o.RecurringType
}

// GetRecurringTypeOk returns a tuple with the RecurringType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetRecurringTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringType) {
		return nil, false
	}
	return o.RecurringType, true
}

// HasRecurringType returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasRecurringType() bool {
	if o != nil && !IsNil(o.RecurringType) {
		return true
	}

	return false
}

// SetRecurringType gets a reference to the given string and assigns it to the RecurringType field.
func (o *CreateDiscountInput) SetRecurringType(v string) {
	o.RecurringType = &v
}

// GetRecurringCycles returns the RecurringCycles field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetRecurringCycles() int32 {
	if o == nil || IsNil(o.RecurringCycles) {
		var ret int32
		return ret
	}
	return *o.RecurringCycles
}

// GetRecurringCyclesOk returns a tuple with the RecurringCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetRecurringCyclesOk() (*int32, bool) {
	if o == nil || IsNil(o.RecurringCycles) {
		return nil, false
	}
	return o.RecurringCycles, true
}

// HasRecurringCycles returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasRecurringCycles() bool {
	if o != nil && !IsNil(o.RecurringCycles) {
		return true
	}

	return false
}

// SetRecurringCycles gets a reference to the given int32 and assigns it to the RecurringCycles field.
func (o *CreateDiscountInput) SetRecurringCycles(v int32) {
	o.RecurringCycles = &v
}

// GetFirstTimeCustomer returns the FirstTimeCustomer field value if set, zero value otherwise.
func (o *CreateDiscountInput) GetFirstTimeCustomer() bool {
	if o == nil || IsNil(o.FirstTimeCustomer) {
		var ret bool
		return ret
	}
	return *o.FirstTimeCustomer
}

// GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountInput) GetFirstTimeCustomerOk() (*bool, bool) {
	if o == nil || IsNil(o.FirstTimeCustomer) {
		return nil, false
	}
	return o.FirstTimeCustomer, true
}

// HasFirstTimeCustomer returns a boolean if a field has been set.
func (o *CreateDiscountInput) HasFirstTimeCustomer() bool {
	if o != nil && !IsNil(o.FirstTimeCustomer) {
		return true
	}

	return false
}

// SetFirstTimeCustomer gets a reference to the given bool and assigns it to the FirstTimeCustomer field.
func (o *CreateDiscountInput) SetFirstTimeCustomer(v bool) {
	o.FirstTimeCustomer = &v
}

func (o CreateDiscountInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDiscountInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["discount_type"] = o.DiscountType
	toSerialize["discount_value"] = o.DiscountValue
	if !IsNil(o.AppliesTo) {
		toSerialize["applies_to"] = o.AppliesTo
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ProductIds) {
		toSerialize["product_ids"] = o.ProductIds
	}
	if !IsNil(o.MaxRedemptions) {
		toSerialize["max_redemptions"] = o.MaxRedemptions
	}
	if !IsNil(o.MaxRedemptionsPerUser) {
		toSerialize["max_redemptions_per_user"] = o.MaxRedemptionsPerUser
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if !IsNil(o.ValidUntil) {
		toSerialize["valid_until"] = o.ValidUntil
	}
	if !IsNil(o.RecurringType) {
		toSerialize["recurring_type"] = o.RecurringType
	}
	if !IsNil(o.RecurringCycles) {
		toSerialize["recurring_cycles"] = o.RecurringCycles
	}
	if !IsNil(o.FirstTimeCustomer) {
		toSerialize["first_time_customer"] = o.FirstTimeCustomer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDiscountInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"name",
		"discount_type",
		"discount_value",
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

	varCreateDiscountInput := _CreateDiscountInput{}

	err = json.Unmarshal(data, &varCreateDiscountInput)

	if err != nil {
		return err
	}

	*o = CreateDiscountInput(varCreateDiscountInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "discount_type")
		delete(additionalProperties, "discount_value")
		delete(additionalProperties, "applies_to")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "max_redemptions")
		delete(additionalProperties, "max_redemptions_per_user")
		delete(additionalProperties, "valid_from")
		delete(additionalProperties, "valid_until")
		delete(additionalProperties, "recurring_type")
		delete(additionalProperties, "recurring_cycles")
		delete(additionalProperties, "first_time_customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDiscountInput struct {
	value *CreateDiscountInput
	isSet bool
}

func (v NullableCreateDiscountInput) Get() *CreateDiscountInput {
	return v.value
}

func (v *NullableCreateDiscountInput) Set(val *CreateDiscountInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDiscountInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDiscountInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDiscountInput(val *CreateDiscountInput) *NullableCreateDiscountInput {
	return &NullableCreateDiscountInput{value: val, isSet: true}
}

func (v NullableCreateDiscountInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDiscountInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


