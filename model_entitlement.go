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
	"time"
	"fmt"
)

// checks if the Entitlement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entitlement{}

// Entitlement struct for Entitlement
type Entitlement struct {
	Id string `json:"id"`
	CustomerId string `json:"customerId"`
	BenefitId string `json:"benefitId"`
	Benefit Benefit `json:"benefit"`
	GrantedAt time.Time `json:"grantedAt"`
	ExpiresAt NullableTime `json:"expiresAt"`
	Metadata map[string]interface{} `json:"metadata"`
	Status string `json:"status"`
	HasAccess bool `json:"hasAccess"`
	AccessLevel string `json:"accessLevel"`
	SubscriptionId NullableString `json:"subscriptionId"`
	SubscriptionStatus NullableString `json:"subscriptionStatus"`
	AdditionalProperties map[string]interface{}
}

type _Entitlement Entitlement

// NewEntitlement instantiates a new Entitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlement(id string, customerId string, benefitId string, benefit Benefit, grantedAt time.Time, expiresAt NullableTime, metadata map[string]interface{}, status string, hasAccess bool, accessLevel string, subscriptionId NullableString, subscriptionStatus NullableString) *Entitlement {
	this := Entitlement{}
	this.Id = id
	this.CustomerId = customerId
	this.BenefitId = benefitId
	this.Benefit = benefit
	this.GrantedAt = grantedAt
	this.ExpiresAt = expiresAt
	this.Metadata = metadata
	this.Status = status
	this.HasAccess = hasAccess
	this.AccessLevel = accessLevel
	this.SubscriptionId = subscriptionId
	this.SubscriptionStatus = subscriptionStatus
	return &this
}

// NewEntitlementWithDefaults instantiates a new Entitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementWithDefaults() *Entitlement {
	this := Entitlement{}
	return &this
}

// GetId returns the Id field value
func (o *Entitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Entitlement) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *Entitlement) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Entitlement) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetBenefitId returns the BenefitId field value
func (o *Entitlement) GetBenefitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BenefitId
}

// GetBenefitIdOk returns a tuple with the BenefitId field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetBenefitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitId, true
}

// SetBenefitId sets field value
func (o *Entitlement) SetBenefitId(v string) {
	o.BenefitId = v
}

// GetBenefit returns the Benefit field value
func (o *Entitlement) GetBenefit() Benefit {
	if o == nil {
		var ret Benefit
		return ret
	}

	return o.Benefit
}

// GetBenefitOk returns a tuple with the Benefit field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetBenefitOk() (*Benefit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Benefit, true
}

// SetBenefit sets field value
func (o *Entitlement) SetBenefit(v Benefit) {
	o.Benefit = v
}

// GetGrantedAt returns the GrantedAt field value
func (o *Entitlement) GetGrantedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.GrantedAt
}

// GetGrantedAtOk returns a tuple with the GrantedAt field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetGrantedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantedAt, true
}

// SetGrantedAt sets field value
func (o *Entitlement) SetGrantedAt(v time.Time) {
	o.GrantedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Entitlement) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlement) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *Entitlement) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *Entitlement) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Entitlement) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetStatus returns the Status field value
func (o *Entitlement) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Entitlement) SetStatus(v string) {
	o.Status = v
}

// GetHasAccess returns the HasAccess field value
func (o *Entitlement) GetHasAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetHasAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAccess, true
}

// SetHasAccess sets field value
func (o *Entitlement) SetHasAccess(v bool) {
	o.HasAccess = v
}

// GetAccessLevel returns the AccessLevel field value
func (o *Entitlement) GetAccessLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value
// and a boolean to check if the value has been set.
func (o *Entitlement) GetAccessLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessLevel, true
}

// SetAccessLevel sets field value
func (o *Entitlement) SetAccessLevel(v string) {
	o.AccessLevel = v
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Entitlement) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlement) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *Entitlement) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetSubscriptionStatus returns the SubscriptionStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Entitlement) GetSubscriptionStatus() string {
	if o == nil || o.SubscriptionStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionStatus.Get()
}

// GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlement) GetSubscriptionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionStatus.Get(), o.SubscriptionStatus.IsSet()
}

// SetSubscriptionStatus sets field value
func (o *Entitlement) SetSubscriptionStatus(v string) {
	o.SubscriptionStatus.Set(&v)
}

func (o Entitlement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entitlement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customerId"] = o.CustomerId
	toSerialize["benefitId"] = o.BenefitId
	toSerialize["benefit"] = o.Benefit
	toSerialize["grantedAt"] = o.GrantedAt
	toSerialize["expiresAt"] = o.ExpiresAt.Get()
	toSerialize["metadata"] = o.Metadata
	toSerialize["status"] = o.Status
	toSerialize["hasAccess"] = o.HasAccess
	toSerialize["accessLevel"] = o.AccessLevel
	toSerialize["subscriptionId"] = o.SubscriptionId.Get()
	toSerialize["subscriptionStatus"] = o.SubscriptionStatus.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Entitlement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customerId",
		"benefitId",
		"benefit",
		"grantedAt",
		"expiresAt",
		"metadata",
		"status",
		"hasAccess",
		"accessLevel",
		"subscriptionId",
		"subscriptionStatus",
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

	varEntitlement := _Entitlement{}

	err = json.Unmarshal(data, &varEntitlement)

	if err != nil {
		return err
	}

	*o = Entitlement(varEntitlement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "benefitId")
		delete(additionalProperties, "benefit")
		delete(additionalProperties, "grantedAt")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "status")
		delete(additionalProperties, "hasAccess")
		delete(additionalProperties, "accessLevel")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "subscriptionStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlement struct {
	value *Entitlement
	isSet bool
}

func (v NullableEntitlement) Get() *Entitlement {
	return v.value
}

func (v *NullableEntitlement) Set(val *Entitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlement(val *Entitlement) *NullableEntitlement {
	return &NullableEntitlement{value: val, isSet: true}
}

func (v NullableEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


