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
)

// checks if the UpdateDiscountInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDiscountInput{}

// UpdateDiscountInput Parameters for updating an existing discount. Active discounts can only have metadata and usage limits modified.
type UpdateDiscountInput struct {
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DiscountValue *float32 `json:"discount_value,omitempty"`
	AppliesTo *string `json:"applies_to,omitempty"`
	Scope *string `json:"scope,omitempty"`
	ProductIds []string `json:"product_ids,omitempty"`
	MaxRedemptions NullableInt32 `json:"max_redemptions,omitempty"`
	MaxRedemptionsPerUser *int32 `json:"max_redemptions_per_user,omitempty"`
	ValidFrom NullableString `json:"valid_from,omitempty"`
	ValidUntil NullableString `json:"valid_until,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	RecurringType NullableString `json:"recurring_type,omitempty"`
	RecurringCycles NullableInt32 `json:"recurring_cycles,omitempty"`
	FirstTimeCustomer *bool `json:"first_time_customer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateDiscountInput UpdateDiscountInput

// NewUpdateDiscountInput instantiates a new UpdateDiscountInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDiscountInput() *UpdateDiscountInput {
	this := UpdateDiscountInput{}
	return &this
}

// NewUpdateDiscountInputWithDefaults instantiates a new UpdateDiscountInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDiscountInputWithDefaults() *UpdateDiscountInput {
	this := UpdateDiscountInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDiscountInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateDiscountInput) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateDiscountInput) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateDiscountInput) UnsetDescription() {
	o.Description.Unset()
}

// GetDiscountValue returns the DiscountValue field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetDiscountValue() float32 {
	if o == nil || IsNil(o.DiscountValue) {
		var ret float32
		return ret
	}
	return *o.DiscountValue
}

// GetDiscountValueOk returns a tuple with the DiscountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetDiscountValueOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountValue) {
		return nil, false
	}
	return o.DiscountValue, true
}

// HasDiscountValue returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasDiscountValue() bool {
	if o != nil && !IsNil(o.DiscountValue) {
		return true
	}

	return false
}

// SetDiscountValue gets a reference to the given float32 and assigns it to the DiscountValue field.
func (o *UpdateDiscountInput) SetDiscountValue(v float32) {
	o.DiscountValue = &v
}

// GetAppliesTo returns the AppliesTo field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetAppliesTo() string {
	if o == nil || IsNil(o.AppliesTo) {
		var ret string
		return ret
	}
	return *o.AppliesTo
}

// GetAppliesToOk returns a tuple with the AppliesTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetAppliesToOk() (*string, bool) {
	if o == nil || IsNil(o.AppliesTo) {
		return nil, false
	}
	return o.AppliesTo, true
}

// HasAppliesTo returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasAppliesTo() bool {
	if o != nil && !IsNil(o.AppliesTo) {
		return true
	}

	return false
}

// SetAppliesTo gets a reference to the given string and assigns it to the AppliesTo field.
func (o *UpdateDiscountInput) SetAppliesTo(v string) {
	o.AppliesTo = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UpdateDiscountInput) SetScope(v string) {
	o.Scope = &v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetProductIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetProductIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []string and assigns it to the ProductIds field.
func (o *UpdateDiscountInput) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetMaxRedemptions returns the MaxRedemptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetMaxRedemptions() int32 {
	if o == nil || IsNil(o.MaxRedemptions.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRedemptions.Get()
}

// GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetMaxRedemptionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRedemptions.Get(), o.MaxRedemptions.IsSet()
}

// HasMaxRedemptions returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasMaxRedemptions() bool {
	if o != nil && o.MaxRedemptions.IsSet() {
		return true
	}

	return false
}

// SetMaxRedemptions gets a reference to the given NullableInt32 and assigns it to the MaxRedemptions field.
func (o *UpdateDiscountInput) SetMaxRedemptions(v int32) {
	o.MaxRedemptions.Set(&v)
}
// SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil
func (o *UpdateDiscountInput) SetMaxRedemptionsNil() {
	o.MaxRedemptions.Set(nil)
}

// UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
func (o *UpdateDiscountInput) UnsetMaxRedemptions() {
	o.MaxRedemptions.Unset()
}

// GetMaxRedemptionsPerUser returns the MaxRedemptionsPerUser field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetMaxRedemptionsPerUser() int32 {
	if o == nil || IsNil(o.MaxRedemptionsPerUser) {
		var ret int32
		return ret
	}
	return *o.MaxRedemptionsPerUser
}

// GetMaxRedemptionsPerUserOk returns a tuple with the MaxRedemptionsPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetMaxRedemptionsPerUserOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRedemptionsPerUser) {
		return nil, false
	}
	return o.MaxRedemptionsPerUser, true
}

// HasMaxRedemptionsPerUser returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasMaxRedemptionsPerUser() bool {
	if o != nil && !IsNil(o.MaxRedemptionsPerUser) {
		return true
	}

	return false
}

// SetMaxRedemptionsPerUser gets a reference to the given int32 and assigns it to the MaxRedemptionsPerUser field.
func (o *UpdateDiscountInput) SetMaxRedemptionsPerUser(v int32) {
	o.MaxRedemptionsPerUser = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom.Get()) {
		var ret string
		return ret
	}
	return *o.ValidFrom.Get()
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetValidFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidFrom.Get(), o.ValidFrom.IsSet()
}

// HasValidFrom returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasValidFrom() bool {
	if o != nil && o.ValidFrom.IsSet() {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given NullableString and assigns it to the ValidFrom field.
func (o *UpdateDiscountInput) SetValidFrom(v string) {
	o.ValidFrom.Set(&v)
}
// SetValidFromNil sets the value for ValidFrom to be an explicit nil
func (o *UpdateDiscountInput) SetValidFromNil() {
	o.ValidFrom.Set(nil)
}

// UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
func (o *UpdateDiscountInput) UnsetValidFrom() {
	o.ValidFrom.Unset()
}

// GetValidUntil returns the ValidUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetValidUntil() string {
	if o == nil || IsNil(o.ValidUntil.Get()) {
		var ret string
		return ret
	}
	return *o.ValidUntil.Get()
}

// GetValidUntilOk returns a tuple with the ValidUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetValidUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidUntil.Get(), o.ValidUntil.IsSet()
}

// HasValidUntil returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasValidUntil() bool {
	if o != nil && o.ValidUntil.IsSet() {
		return true
	}

	return false
}

// SetValidUntil gets a reference to the given NullableString and assigns it to the ValidUntil field.
func (o *UpdateDiscountInput) SetValidUntil(v string) {
	o.ValidUntil.Set(&v)
}
// SetValidUntilNil sets the value for ValidUntil to be an explicit nil
func (o *UpdateDiscountInput) SetValidUntilNil() {
	o.ValidUntil.Set(nil)
}

// UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
func (o *UpdateDiscountInput) UnsetValidUntil() {
	o.ValidUntil.Unset()
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UpdateDiscountInput) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetIsArchived() bool {
	if o == nil || IsNil(o.IsArchived) {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetIsArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *UpdateDiscountInput) SetIsArchived(v bool) {
	o.IsArchived = &v
}

// GetRecurringType returns the RecurringType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetRecurringType() string {
	if o == nil || IsNil(o.RecurringType.Get()) {
		var ret string
		return ret
	}
	return *o.RecurringType.Get()
}

// GetRecurringTypeOk returns a tuple with the RecurringType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetRecurringTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringType.Get(), o.RecurringType.IsSet()
}

// HasRecurringType returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasRecurringType() bool {
	if o != nil && o.RecurringType.IsSet() {
		return true
	}

	return false
}

// SetRecurringType gets a reference to the given NullableString and assigns it to the RecurringType field.
func (o *UpdateDiscountInput) SetRecurringType(v string) {
	o.RecurringType.Set(&v)
}
// SetRecurringTypeNil sets the value for RecurringType to be an explicit nil
func (o *UpdateDiscountInput) SetRecurringTypeNil() {
	o.RecurringType.Set(nil)
}

// UnsetRecurringType ensures that no value is present for RecurringType, not even an explicit nil
func (o *UpdateDiscountInput) UnsetRecurringType() {
	o.RecurringType.Unset()
}

// GetRecurringCycles returns the RecurringCycles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDiscountInput) GetRecurringCycles() int32 {
	if o == nil || IsNil(o.RecurringCycles.Get()) {
		var ret int32
		return ret
	}
	return *o.RecurringCycles.Get()
}

// GetRecurringCyclesOk returns a tuple with the RecurringCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDiscountInput) GetRecurringCyclesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecurringCycles.Get(), o.RecurringCycles.IsSet()
}

// HasRecurringCycles returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasRecurringCycles() bool {
	if o != nil && o.RecurringCycles.IsSet() {
		return true
	}

	return false
}

// SetRecurringCycles gets a reference to the given NullableInt32 and assigns it to the RecurringCycles field.
func (o *UpdateDiscountInput) SetRecurringCycles(v int32) {
	o.RecurringCycles.Set(&v)
}
// SetRecurringCyclesNil sets the value for RecurringCycles to be an explicit nil
func (o *UpdateDiscountInput) SetRecurringCyclesNil() {
	o.RecurringCycles.Set(nil)
}

// UnsetRecurringCycles ensures that no value is present for RecurringCycles, not even an explicit nil
func (o *UpdateDiscountInput) UnsetRecurringCycles() {
	o.RecurringCycles.Unset()
}

// GetFirstTimeCustomer returns the FirstTimeCustomer field value if set, zero value otherwise.
func (o *UpdateDiscountInput) GetFirstTimeCustomer() bool {
	if o == nil || IsNil(o.FirstTimeCustomer) {
		var ret bool
		return ret
	}
	return *o.FirstTimeCustomer
}

// GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDiscountInput) GetFirstTimeCustomerOk() (*bool, bool) {
	if o == nil || IsNil(o.FirstTimeCustomer) {
		return nil, false
	}
	return o.FirstTimeCustomer, true
}

// HasFirstTimeCustomer returns a boolean if a field has been set.
func (o *UpdateDiscountInput) HasFirstTimeCustomer() bool {
	if o != nil && !IsNil(o.FirstTimeCustomer) {
		return true
	}

	return false
}

// SetFirstTimeCustomer gets a reference to the given bool and assigns it to the FirstTimeCustomer field.
func (o *UpdateDiscountInput) SetFirstTimeCustomer(v bool) {
	o.FirstTimeCustomer = &v
}

func (o UpdateDiscountInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDiscountInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.DiscountValue) {
		toSerialize["discount_value"] = o.DiscountValue
	}
	if !IsNil(o.AppliesTo) {
		toSerialize["applies_to"] = o.AppliesTo
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if o.ProductIds != nil {
		toSerialize["product_ids"] = o.ProductIds
	}
	if o.MaxRedemptions.IsSet() {
		toSerialize["max_redemptions"] = o.MaxRedemptions.Get()
	}
	if !IsNil(o.MaxRedemptionsPerUser) {
		toSerialize["max_redemptions_per_user"] = o.MaxRedemptionsPerUser
	}
	if o.ValidFrom.IsSet() {
		toSerialize["valid_from"] = o.ValidFrom.Get()
	}
	if o.ValidUntil.IsSet() {
		toSerialize["valid_until"] = o.ValidUntil.Get()
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.IsArchived) {
		toSerialize["is_archived"] = o.IsArchived
	}
	if o.RecurringType.IsSet() {
		toSerialize["recurring_type"] = o.RecurringType.Get()
	}
	if o.RecurringCycles.IsSet() {
		toSerialize["recurring_cycles"] = o.RecurringCycles.Get()
	}
	if !IsNil(o.FirstTimeCustomer) {
		toSerialize["first_time_customer"] = o.FirstTimeCustomer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateDiscountInput) UnmarshalJSON(data []byte) (err error) {
	varUpdateDiscountInput := _UpdateDiscountInput{}

	err = json.Unmarshal(data, &varUpdateDiscountInput)

	if err != nil {
		return err
	}

	*o = UpdateDiscountInput(varUpdateDiscountInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "discount_value")
		delete(additionalProperties, "applies_to")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "product_ids")
		delete(additionalProperties, "max_redemptions")
		delete(additionalProperties, "max_redemptions_per_user")
		delete(additionalProperties, "valid_from")
		delete(additionalProperties, "valid_until")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "is_archived")
		delete(additionalProperties, "recurring_type")
		delete(additionalProperties, "recurring_cycles")
		delete(additionalProperties, "first_time_customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateDiscountInput struct {
	value *UpdateDiscountInput
	isSet bool
}

func (v NullableUpdateDiscountInput) Get() *UpdateDiscountInput {
	return v.value
}

func (v *NullableUpdateDiscountInput) Set(val *UpdateDiscountInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiscountInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiscountInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiscountInput(val *UpdateDiscountInput) *NullableUpdateDiscountInput {
	return &NullableUpdateDiscountInput{value: val, isSet: true}
}

func (v NullableUpdateDiscountInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiscountInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


