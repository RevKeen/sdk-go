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

// checks if the CustomerSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSubscription{}

// CustomerSubscription A subscription belonging to a specific customer, with plan and status details.
type CustomerSubscription struct {
	Id string `json:"id"`
	CustomerId string `json:"customer_id"`
	PriceId NullableString `json:"price_id"`
	Status string `json:"status"`
	CurrentPeriodStart NullableString `json:"current_period_start"`
	CurrentPeriodEnd NullableString `json:"current_period_end"`
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	CanceledAt NullableString `json:"canceled_at"`
	TrialEnd NullableString `json:"trial_end"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CustomerSubscription CustomerSubscription

// NewCustomerSubscription instantiates a new CustomerSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSubscription(id string, customerId string, priceId NullableString, status string, currentPeriodStart NullableString, currentPeriodEnd NullableString, cancelAtPeriodEnd bool, canceledAt NullableString, trialEnd NullableString, createdAt string, updatedAt string) *CustomerSubscription {
	this := CustomerSubscription{}
	this.Id = id
	this.CustomerId = customerId
	this.PriceId = priceId
	this.Status = status
	this.CurrentPeriodStart = currentPeriodStart
	this.CurrentPeriodEnd = currentPeriodEnd
	this.CancelAtPeriodEnd = cancelAtPeriodEnd
	this.CanceledAt = canceledAt
	this.TrialEnd = trialEnd
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomerSubscriptionWithDefaults instantiates a new CustomerSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSubscriptionWithDefaults() *CustomerSubscription {
	this := CustomerSubscription{}
	return &this
}

// GetId returns the Id field value
func (o *CustomerSubscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerSubscription) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *CustomerSubscription) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomerSubscription) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetPriceId returns the PriceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerSubscription) GetPriceId() string {
	if o == nil || o.PriceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PriceId.Get()
}

// GetPriceIdOk returns a tuple with the PriceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerSubscription) GetPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceId.Get(), o.PriceId.IsSet()
}

// SetPriceId sets field value
func (o *CustomerSubscription) SetPriceId(v string) {
	o.PriceId.Set(&v)
}

// GetStatus returns the Status field value
func (o *CustomerSubscription) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CustomerSubscription) SetStatus(v string) {
	o.Status = v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerSubscription) GetCurrentPeriodStart() string {
	if o == nil || o.CurrentPeriodStart.Get() == nil {
		var ret string
		return ret
	}

	return *o.CurrentPeriodStart.Get()
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerSubscription) GetCurrentPeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodStart.Get(), o.CurrentPeriodStart.IsSet()
}

// SetCurrentPeriodStart sets field value
func (o *CustomerSubscription) SetCurrentPeriodStart(v string) {
	o.CurrentPeriodStart.Set(&v)
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerSubscription) GetCurrentPeriodEnd() string {
	if o == nil || o.CurrentPeriodEnd.Get() == nil {
		var ret string
		return ret
	}

	return *o.CurrentPeriodEnd.Get()
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerSubscription) GetCurrentPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd.Get(), o.CurrentPeriodEnd.IsSet()
}

// SetCurrentPeriodEnd sets field value
func (o *CustomerSubscription) SetCurrentPeriodEnd(v string) {
	o.CurrentPeriodEnd.Set(&v)
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value
func (o *CustomerSubscription) GetCancelAtPeriodEnd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetCancelAtPeriodEndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelAtPeriodEnd, true
}

// SetCancelAtPeriodEnd sets field value
func (o *CustomerSubscription) SetCancelAtPeriodEnd(v bool) {
	o.CancelAtPeriodEnd = v
}

// GetCanceledAt returns the CanceledAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerSubscription) GetCanceledAt() string {
	if o == nil || o.CanceledAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerSubscription) GetCanceledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// SetCanceledAt sets field value
func (o *CustomerSubscription) SetCanceledAt(v string) {
	o.CanceledAt.Set(&v)
}

// GetTrialEnd returns the TrialEnd field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerSubscription) GetTrialEnd() string {
	if o == nil || o.TrialEnd.Get() == nil {
		var ret string
		return ret
	}

	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerSubscription) GetTrialEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// SetTrialEnd sets field value
func (o *CustomerSubscription) SetTrialEnd(v string) {
	o.TrialEnd.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerSubscription) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerSubscription) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomerSubscription) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerSubscription) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomerSubscription) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o CustomerSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["price_id"] = o.PriceId.Get()
	toSerialize["status"] = o.Status
	toSerialize["current_period_start"] = o.CurrentPeriodStart.Get()
	toSerialize["current_period_end"] = o.CurrentPeriodEnd.Get()
	toSerialize["cancel_at_period_end"] = o.CancelAtPeriodEnd
	toSerialize["canceled_at"] = o.CanceledAt.Get()
	toSerialize["trial_end"] = o.TrialEnd.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customer_id",
		"price_id",
		"status",
		"current_period_start",
		"current_period_end",
		"cancel_at_period_end",
		"canceled_at",
		"trial_end",
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

	varCustomerSubscription := _CustomerSubscription{}

	err = json.Unmarshal(data, &varCustomerSubscription)

	if err != nil {
		return err
	}

	*o = CustomerSubscription(varCustomerSubscription)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "price_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "current_period_start")
		delete(additionalProperties, "current_period_end")
		delete(additionalProperties, "cancel_at_period_end")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "trial_end")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerSubscription struct {
	value *CustomerSubscription
	isSet bool
}

func (v NullableCustomerSubscription) Get() *CustomerSubscription {
	return v.value
}

func (v *NullableCustomerSubscription) Set(val *CustomerSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSubscription(val *CustomerSubscription) *NullableCustomerSubscription {
	return &NullableCustomerSubscription{value: val, isSet: true}
}

func (v NullableCustomerSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


