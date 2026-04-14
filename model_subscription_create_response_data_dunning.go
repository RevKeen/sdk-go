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

// checks if the SubscriptionCreateResponseDataDunning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateResponseDataDunning{}

// SubscriptionCreateResponseDataDunning struct for SubscriptionCreateResponseDataDunning
type SubscriptionCreateResponseDataDunning struct {
	IsInDunning bool `json:"isInDunning"`
	Phase NullableFloat32 `json:"phase"`
	PhaseLabel NullableString `json:"phaseLabel"`
	PhaseSeverity NullableString `json:"phaseSeverity"`
	RetryCount float32 `json:"retryCount"`
	TotalPossibleRetries float32 `json:"totalPossibleRetries"`
	NextRetryAt NullableTime `json:"nextRetryAt"`
	DaysInDunning float32 `json:"daysInDunning"`
	AccessRestricted bool `json:"accessRestricted"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionCreateResponseDataDunning SubscriptionCreateResponseDataDunning

// NewSubscriptionCreateResponseDataDunning instantiates a new SubscriptionCreateResponseDataDunning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateResponseDataDunning(isInDunning bool, phase NullableFloat32, phaseLabel NullableString, phaseSeverity NullableString, retryCount float32, totalPossibleRetries float32, nextRetryAt NullableTime, daysInDunning float32, accessRestricted bool) *SubscriptionCreateResponseDataDunning {
	this := SubscriptionCreateResponseDataDunning{}
	this.IsInDunning = isInDunning
	this.Phase = phase
	this.PhaseLabel = phaseLabel
	this.PhaseSeverity = phaseSeverity
	this.RetryCount = retryCount
	this.TotalPossibleRetries = totalPossibleRetries
	this.NextRetryAt = nextRetryAt
	this.DaysInDunning = daysInDunning
	this.AccessRestricted = accessRestricted
	return &this
}

// NewSubscriptionCreateResponseDataDunningWithDefaults instantiates a new SubscriptionCreateResponseDataDunning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateResponseDataDunningWithDefaults() *SubscriptionCreateResponseDataDunning {
	this := SubscriptionCreateResponseDataDunning{}
	return &this
}

// GetIsInDunning returns the IsInDunning field value
func (o *SubscriptionCreateResponseDataDunning) GetIsInDunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInDunning
}

// GetIsInDunningOk returns a tuple with the IsInDunning field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseDataDunning) GetIsInDunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInDunning, true
}

// SetIsInDunning sets field value
func (o *SubscriptionCreateResponseDataDunning) SetIsInDunning(v bool) {
	o.IsInDunning = v
}

// GetPhase returns the Phase field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhase() float32 {
	if o == nil || o.Phase.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Phase.Get()
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhaseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase.Get(), o.Phase.IsSet()
}

// SetPhase sets field value
func (o *SubscriptionCreateResponseDataDunning) SetPhase(v float32) {
	o.Phase.Set(&v)
}

// GetPhaseLabel returns the PhaseLabel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhaseLabel() string {
	if o == nil || o.PhaseLabel.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhaseLabel.Get()
}

// GetPhaseLabelOk returns a tuple with the PhaseLabel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhaseLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhaseLabel.Get(), o.PhaseLabel.IsSet()
}

// SetPhaseLabel sets field value
func (o *SubscriptionCreateResponseDataDunning) SetPhaseLabel(v string) {
	o.PhaseLabel.Set(&v)
}

// GetPhaseSeverity returns the PhaseSeverity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhaseSeverity() string {
	if o == nil || o.PhaseSeverity.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhaseSeverity.Get()
}

// GetPhaseSeverityOk returns a tuple with the PhaseSeverity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionCreateResponseDataDunning) GetPhaseSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhaseSeverity.Get(), o.PhaseSeverity.IsSet()
}

// SetPhaseSeverity sets field value
func (o *SubscriptionCreateResponseDataDunning) SetPhaseSeverity(v string) {
	o.PhaseSeverity.Set(&v)
}

// GetRetryCount returns the RetryCount field value
func (o *SubscriptionCreateResponseDataDunning) GetRetryCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseDataDunning) GetRetryCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryCount, true
}

// SetRetryCount sets field value
func (o *SubscriptionCreateResponseDataDunning) SetRetryCount(v float32) {
	o.RetryCount = v
}

// GetTotalPossibleRetries returns the TotalPossibleRetries field value
func (o *SubscriptionCreateResponseDataDunning) GetTotalPossibleRetries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPossibleRetries
}

// GetTotalPossibleRetriesOk returns a tuple with the TotalPossibleRetries field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseDataDunning) GetTotalPossibleRetriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPossibleRetries, true
}

// SetTotalPossibleRetries sets field value
func (o *SubscriptionCreateResponseDataDunning) SetTotalPossibleRetries(v float32) {
	o.TotalPossibleRetries = v
}

// GetNextRetryAt returns the NextRetryAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SubscriptionCreateResponseDataDunning) GetNextRetryAt() time.Time {
	if o == nil || o.NextRetryAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NextRetryAt.Get()
}

// GetNextRetryAtOk returns a tuple with the NextRetryAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionCreateResponseDataDunning) GetNextRetryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRetryAt.Get(), o.NextRetryAt.IsSet()
}

// SetNextRetryAt sets field value
func (o *SubscriptionCreateResponseDataDunning) SetNextRetryAt(v time.Time) {
	o.NextRetryAt.Set(&v)
}

// GetDaysInDunning returns the DaysInDunning field value
func (o *SubscriptionCreateResponseDataDunning) GetDaysInDunning() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysInDunning
}

// GetDaysInDunningOk returns a tuple with the DaysInDunning field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseDataDunning) GetDaysInDunningOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysInDunning, true
}

// SetDaysInDunning sets field value
func (o *SubscriptionCreateResponseDataDunning) SetDaysInDunning(v float32) {
	o.DaysInDunning = v
}

// GetAccessRestricted returns the AccessRestricted field value
func (o *SubscriptionCreateResponseDataDunning) GetAccessRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessRestricted
}

// GetAccessRestrictedOk returns a tuple with the AccessRestricted field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseDataDunning) GetAccessRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRestricted, true
}

// SetAccessRestricted sets field value
func (o *SubscriptionCreateResponseDataDunning) SetAccessRestricted(v bool) {
	o.AccessRestricted = v
}

func (o SubscriptionCreateResponseDataDunning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateResponseDataDunning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isInDunning"] = o.IsInDunning
	toSerialize["phase"] = o.Phase.Get()
	toSerialize["phaseLabel"] = o.PhaseLabel.Get()
	toSerialize["phaseSeverity"] = o.PhaseSeverity.Get()
	toSerialize["retryCount"] = o.RetryCount
	toSerialize["totalPossibleRetries"] = o.TotalPossibleRetries
	toSerialize["nextRetryAt"] = o.NextRetryAt.Get()
	toSerialize["daysInDunning"] = o.DaysInDunning
	toSerialize["accessRestricted"] = o.AccessRestricted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionCreateResponseDataDunning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isInDunning",
		"phase",
		"phaseLabel",
		"phaseSeverity",
		"retryCount",
		"totalPossibleRetries",
		"nextRetryAt",
		"daysInDunning",
		"accessRestricted",
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

	varSubscriptionCreateResponseDataDunning := _SubscriptionCreateResponseDataDunning{}

	err = json.Unmarshal(data, &varSubscriptionCreateResponseDataDunning)

	if err != nil {
		return err
	}

	*o = SubscriptionCreateResponseDataDunning(varSubscriptionCreateResponseDataDunning)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isInDunning")
		delete(additionalProperties, "phase")
		delete(additionalProperties, "phaseLabel")
		delete(additionalProperties, "phaseSeverity")
		delete(additionalProperties, "retryCount")
		delete(additionalProperties, "totalPossibleRetries")
		delete(additionalProperties, "nextRetryAt")
		delete(additionalProperties, "daysInDunning")
		delete(additionalProperties, "accessRestricted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionCreateResponseDataDunning struct {
	value *SubscriptionCreateResponseDataDunning
	isSet bool
}

func (v NullableSubscriptionCreateResponseDataDunning) Get() *SubscriptionCreateResponseDataDunning {
	return v.value
}

func (v *NullableSubscriptionCreateResponseDataDunning) Set(val *SubscriptionCreateResponseDataDunning) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateResponseDataDunning) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateResponseDataDunning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateResponseDataDunning(val *SubscriptionCreateResponseDataDunning) *NullableSubscriptionCreateResponseDataDunning {
	return &NullableSubscriptionCreateResponseDataDunning{value: val, isSet: true}
}

func (v NullableSubscriptionCreateResponseDataDunning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateResponseDataDunning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


