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
	"time"
	"fmt"
)

// checks if the PortalSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortalSubscription{}

// PortalSubscription A customer-facing subscription projection. Pricing and cadence fields are exposed; merchant-only bookkeeping fields (plan_id, external IDs) are not.
type PortalSubscription struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Status string `json:"status"`
	Currency NullableString `json:"currency"`
	AmountMinor NullableInt32 `json:"amount_minor"`
	Interval NullableString `json:"interval"`
	IntervalCount NullableInt32 `json:"interval_count"`
	CurrentPeriodStart NullableTime `json:"current_period_start"`
	CurrentPeriodEnd NullableTime `json:"current_period_end"`
	TrialEnd NullableTime `json:"trial_end"`
	CancelAt NullableTime `json:"cancel_at"`
	CanceledAt NullableTime `json:"canceled_at"`
	StartedAt NullableTime `json:"started_at"`
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _PortalSubscription PortalSubscription

// NewPortalSubscription instantiates a new PortalSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalSubscription(id string, object string, status string, currency NullableString, amountMinor NullableInt32, interval NullableString, intervalCount NullableInt32, currentPeriodStart NullableTime, currentPeriodEnd NullableTime, trialEnd NullableTime, cancelAt NullableTime, canceledAt NullableTime, startedAt NullableTime, createdAt time.Time) *PortalSubscription {
	this := PortalSubscription{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.Currency = currency
	this.AmountMinor = amountMinor
	this.Interval = interval
	this.IntervalCount = intervalCount
	this.CurrentPeriodStart = currentPeriodStart
	this.CurrentPeriodEnd = currentPeriodEnd
	this.TrialEnd = trialEnd
	this.CancelAt = cancelAt
	this.CanceledAt = canceledAt
	this.StartedAt = startedAt
	this.CreatedAt = createdAt
	return &this
}

// NewPortalSubscriptionWithDefaults instantiates a new PortalSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalSubscriptionWithDefaults() *PortalSubscription {
	this := PortalSubscription{}
	return &this
}

// GetId returns the Id field value
func (o *PortalSubscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PortalSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PortalSubscription) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PortalSubscription) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PortalSubscription) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PortalSubscription) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *PortalSubscription) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PortalSubscription) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PortalSubscription) SetStatus(v string) {
	o.Status = v
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalSubscription) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *PortalSubscription) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalSubscription) GetAmountMinor() int32 {
	if o == nil || o.AmountMinor.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AmountMinor.Get()
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountMinor.Get(), o.AmountMinor.IsSet()
}

// SetAmountMinor sets field value
func (o *PortalSubscription) SetAmountMinor(v int32) {
	o.AmountMinor.Set(&v)
}

// GetInterval returns the Interval field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalSubscription) GetInterval() string {
	if o == nil || o.Interval.Get() == nil {
		var ret string
		return ret
	}

	return *o.Interval.Get()
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interval.Get(), o.Interval.IsSet()
}

// SetInterval sets field value
func (o *PortalSubscription) SetInterval(v string) {
	o.Interval.Set(&v)
}

// GetIntervalCount returns the IntervalCount field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PortalSubscription) GetIntervalCount() int32 {
	if o == nil || o.IntervalCount.Get() == nil {
		var ret int32
		return ret
	}

	return *o.IntervalCount.Get()
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetIntervalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalCount.Get(), o.IntervalCount.IsSet()
}

// SetIntervalCount sets field value
func (o *PortalSubscription) SetIntervalCount(v int32) {
	o.IntervalCount.Set(&v)
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetCurrentPeriodStart() time.Time {
	if o == nil || o.CurrentPeriodStart.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CurrentPeriodStart.Get()
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetCurrentPeriodStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodStart.Get(), o.CurrentPeriodStart.IsSet()
}

// SetCurrentPeriodStart sets field value
func (o *PortalSubscription) SetCurrentPeriodStart(v time.Time) {
	o.CurrentPeriodStart.Set(&v)
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetCurrentPeriodEnd() time.Time {
	if o == nil || o.CurrentPeriodEnd.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CurrentPeriodEnd.Get()
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetCurrentPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPeriodEnd.Get(), o.CurrentPeriodEnd.IsSet()
}

// SetCurrentPeriodEnd sets field value
func (o *PortalSubscription) SetCurrentPeriodEnd(v time.Time) {
	o.CurrentPeriodEnd.Set(&v)
}

// GetTrialEnd returns the TrialEnd field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetTrialEnd() time.Time {
	if o == nil || o.TrialEnd.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.TrialEnd.Get()
}

// GetTrialEndOk returns a tuple with the TrialEnd field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetTrialEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEnd.Get(), o.TrialEnd.IsSet()
}

// SetTrialEnd sets field value
func (o *PortalSubscription) SetTrialEnd(v time.Time) {
	o.TrialEnd.Set(&v)
}

// GetCancelAt returns the CancelAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetCancelAt() time.Time {
	if o == nil || o.CancelAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CancelAt.Get()
}

// GetCancelAtOk returns a tuple with the CancelAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetCancelAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelAt.Get(), o.CancelAt.IsSet()
}

// SetCancelAt sets field value
func (o *PortalSubscription) SetCancelAt(v time.Time) {
	o.CancelAt.Set(&v)
}

// GetCanceledAt returns the CanceledAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetCanceledAt() time.Time {
	if o == nil || o.CanceledAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// SetCanceledAt sets field value
func (o *PortalSubscription) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}

// GetStartedAt returns the StartedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PortalSubscription) GetStartedAt() time.Time {
	if o == nil || o.StartedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalSubscription) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// SetStartedAt sets field value
func (o *PortalSubscription) SetStartedAt(v time.Time) {
	o.StartedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *PortalSubscription) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PortalSubscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PortalSubscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PortalSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortalSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["amount_minor"] = o.AmountMinor.Get()
	toSerialize["interval"] = o.Interval.Get()
	toSerialize["interval_count"] = o.IntervalCount.Get()
	toSerialize["current_period_start"] = o.CurrentPeriodStart.Get()
	toSerialize["current_period_end"] = o.CurrentPeriodEnd.Get()
	toSerialize["trial_end"] = o.TrialEnd.Get()
	toSerialize["cancel_at"] = o.CancelAt.Get()
	toSerialize["canceled_at"] = o.CanceledAt.Get()
	toSerialize["started_at"] = o.StartedAt.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortalSubscription) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"currency",
		"amount_minor",
		"interval",
		"interval_count",
		"current_period_start",
		"current_period_end",
		"trial_end",
		"cancel_at",
		"canceled_at",
		"started_at",
		"created_at",
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

	varPortalSubscription := _PortalSubscription{}

	err = json.Unmarshal(data, &varPortalSubscription)

	if err != nil {
		return err
	}

	*o = PortalSubscription(varPortalSubscription)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "status")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "interval")
		delete(additionalProperties, "interval_count")
		delete(additionalProperties, "current_period_start")
		delete(additionalProperties, "current_period_end")
		delete(additionalProperties, "trial_end")
		delete(additionalProperties, "cancel_at")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "started_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortalSubscription struct {
	value *PortalSubscription
	isSet bool
}

func (v NullablePortalSubscription) Get() *PortalSubscription {
	return v.value
}

func (v *NullablePortalSubscription) Set(val *PortalSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalSubscription(val *PortalSubscription) *NullablePortalSubscription {
	return &NullablePortalSubscription{value: val, isSet: true}
}

func (v NullablePortalSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


