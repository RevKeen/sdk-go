/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the SubscriptionDunning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDunning{}

// SubscriptionDunning Dunning summary when past due
type SubscriptionDunning struct {
	IsInDunning bool `json:"is_in_dunning"`
	Phase NullableFloat32 `json:"phase"`
	PhaseLabel NullableString `json:"phase_label"`
	PhaseSeverity NullableString `json:"phase_severity"`
	RetryCount float32 `json:"retry_count"`
	TotalPossibleRetries float32 `json:"total_possible_retries"`
	NextRetryAt NullableTime `json:"next_retry_at"`
	DaysInDunning float32 `json:"days_in_dunning"`
	AccessRestricted bool `json:"access_restricted"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionDunning SubscriptionDunning

// NewSubscriptionDunning instantiates a new SubscriptionDunning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDunning(isInDunning bool, phase NullableFloat32, phaseLabel NullableString, phaseSeverity NullableString, retryCount float32, totalPossibleRetries float32, nextRetryAt NullableTime, daysInDunning float32, accessRestricted bool) *SubscriptionDunning {
	this := SubscriptionDunning{}
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

// NewSubscriptionDunningWithDefaults instantiates a new SubscriptionDunning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDunningWithDefaults() *SubscriptionDunning {
	this := SubscriptionDunning{}
	return &this
}

// GetIsInDunning returns the IsInDunning field value
func (o *SubscriptionDunning) GetIsInDunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInDunning
}

// GetIsInDunningOk returns a tuple with the IsInDunning field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDunning) GetIsInDunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsInDunning, true
}

// SetIsInDunning sets field value
func (o *SubscriptionDunning) SetIsInDunning(v bool) {
	o.IsInDunning = v
}

// GetPhase returns the Phase field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *SubscriptionDunning) GetPhase() float32 {
	if o == nil || o.Phase.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Phase.Get()
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionDunning) GetPhaseOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phase.Get(), o.Phase.IsSet()
}

// SetPhase sets field value
func (o *SubscriptionDunning) SetPhase(v float32) {
	o.Phase.Set(&v)
}

// GetPhaseLabel returns the PhaseLabel field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionDunning) GetPhaseLabel() string {
	if o == nil || o.PhaseLabel.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhaseLabel.Get()
}

// GetPhaseLabelOk returns a tuple with the PhaseLabel field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionDunning) GetPhaseLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhaseLabel.Get(), o.PhaseLabel.IsSet()
}

// SetPhaseLabel sets field value
func (o *SubscriptionDunning) SetPhaseLabel(v string) {
	o.PhaseLabel.Set(&v)
}

// GetPhaseSeverity returns the PhaseSeverity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionDunning) GetPhaseSeverity() string {
	if o == nil || o.PhaseSeverity.Get() == nil {
		var ret string
		return ret
	}

	return *o.PhaseSeverity.Get()
}

// GetPhaseSeverityOk returns a tuple with the PhaseSeverity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionDunning) GetPhaseSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhaseSeverity.Get(), o.PhaseSeverity.IsSet()
}

// SetPhaseSeverity sets field value
func (o *SubscriptionDunning) SetPhaseSeverity(v string) {
	o.PhaseSeverity.Set(&v)
}

// GetRetryCount returns the RetryCount field value
func (o *SubscriptionDunning) GetRetryCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDunning) GetRetryCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryCount, true
}

// SetRetryCount sets field value
func (o *SubscriptionDunning) SetRetryCount(v float32) {
	o.RetryCount = v
}

// GetTotalPossibleRetries returns the TotalPossibleRetries field value
func (o *SubscriptionDunning) GetTotalPossibleRetries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalPossibleRetries
}

// GetTotalPossibleRetriesOk returns a tuple with the TotalPossibleRetries field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDunning) GetTotalPossibleRetriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPossibleRetries, true
}

// SetTotalPossibleRetries sets field value
func (o *SubscriptionDunning) SetTotalPossibleRetries(v float32) {
	o.TotalPossibleRetries = v
}

// GetNextRetryAt returns the NextRetryAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SubscriptionDunning) GetNextRetryAt() time.Time {
	if o == nil || o.NextRetryAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NextRetryAt.Get()
}

// GetNextRetryAtOk returns a tuple with the NextRetryAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionDunning) GetNextRetryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRetryAt.Get(), o.NextRetryAt.IsSet()
}

// SetNextRetryAt sets field value
func (o *SubscriptionDunning) SetNextRetryAt(v time.Time) {
	o.NextRetryAt.Set(&v)
}

// GetDaysInDunning returns the DaysInDunning field value
func (o *SubscriptionDunning) GetDaysInDunning() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysInDunning
}

// GetDaysInDunningOk returns a tuple with the DaysInDunning field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDunning) GetDaysInDunningOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysInDunning, true
}

// SetDaysInDunning sets field value
func (o *SubscriptionDunning) SetDaysInDunning(v float32) {
	o.DaysInDunning = v
}

// GetAccessRestricted returns the AccessRestricted field value
func (o *SubscriptionDunning) GetAccessRestricted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessRestricted
}

// GetAccessRestrictedOk returns a tuple with the AccessRestricted field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDunning) GetAccessRestrictedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessRestricted, true
}

// SetAccessRestricted sets field value
func (o *SubscriptionDunning) SetAccessRestricted(v bool) {
	o.AccessRestricted = v
}

func (o SubscriptionDunning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDunning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_in_dunning"] = o.IsInDunning
	toSerialize["phase"] = o.Phase.Get()
	toSerialize["phase_label"] = o.PhaseLabel.Get()
	toSerialize["phase_severity"] = o.PhaseSeverity.Get()
	toSerialize["retry_count"] = o.RetryCount
	toSerialize["total_possible_retries"] = o.TotalPossibleRetries
	toSerialize["next_retry_at"] = o.NextRetryAt.Get()
	toSerialize["days_in_dunning"] = o.DaysInDunning
	toSerialize["access_restricted"] = o.AccessRestricted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionDunning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"is_in_dunning",
		"phase",
		"phase_label",
		"phase_severity",
		"retry_count",
		"total_possible_retries",
		"next_retry_at",
		"days_in_dunning",
		"access_restricted",
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

	varSubscriptionDunning := _SubscriptionDunning{}

	err = json.Unmarshal(data, &varSubscriptionDunning)

	if err != nil {
		return err
	}

	*o = SubscriptionDunning(varSubscriptionDunning)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_in_dunning")
		delete(additionalProperties, "phase")
		delete(additionalProperties, "phase_label")
		delete(additionalProperties, "phase_severity")
		delete(additionalProperties, "retry_count")
		delete(additionalProperties, "total_possible_retries")
		delete(additionalProperties, "next_retry_at")
		delete(additionalProperties, "days_in_dunning")
		delete(additionalProperties, "access_restricted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionDunning struct {
	value *SubscriptionDunning
	isSet bool
}

func (v NullableSubscriptionDunning) Get() *SubscriptionDunning {
	return v.value
}

func (v *NullableSubscriptionDunning) Set(val *SubscriptionDunning) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDunning) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDunning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDunning(val *SubscriptionDunning) *NullableSubscriptionDunning {
	return &NullableSubscriptionDunning{value: val, isSet: true}
}

func (v NullableSubscriptionDunning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDunning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


