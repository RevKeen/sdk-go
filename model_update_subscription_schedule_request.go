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

// checks if the UpdateSubscriptionScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSubscriptionScheduleRequest{}

// UpdateSubscriptionScheduleRequest Parameters for updating a subscription schedule's phases before they take effect.
type UpdateSubscriptionScheduleRequest struct {
	// Replace all phases
	Phases []SchedulePhase `json:"phases,omitempty"`
	// What happens when all phases complete
	EndBehavior *string `json:"end_behavior,omitempty"`
	// Proration behavior when updating phases
	ProrationBehavior *string `json:"proration_behavior,omitempty"`
	// Custom metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateSubscriptionScheduleRequest UpdateSubscriptionScheduleRequest

// NewUpdateSubscriptionScheduleRequest instantiates a new UpdateSubscriptionScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionScheduleRequest() *UpdateSubscriptionScheduleRequest {
	this := UpdateSubscriptionScheduleRequest{}
	return &this
}

// NewUpdateSubscriptionScheduleRequestWithDefaults instantiates a new UpdateSubscriptionScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionScheduleRequestWithDefaults() *UpdateSubscriptionScheduleRequest {
	this := UpdateSubscriptionScheduleRequest{}
	return &this
}

// GetPhases returns the Phases field value if set, zero value otherwise.
func (o *UpdateSubscriptionScheduleRequest) GetPhases() []SchedulePhase {
	if o == nil || IsNil(o.Phases) {
		var ret []SchedulePhase
		return ret
	}
	return o.Phases
}

// GetPhasesOk returns a tuple with the Phases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionScheduleRequest) GetPhasesOk() ([]SchedulePhase, bool) {
	if o == nil || IsNil(o.Phases) {
		return nil, false
	}
	return o.Phases, true
}

// HasPhases returns a boolean if a field has been set.
func (o *UpdateSubscriptionScheduleRequest) HasPhases() bool {
	if o != nil && !IsNil(o.Phases) {
		return true
	}

	return false
}

// SetPhases gets a reference to the given []SchedulePhase and assigns it to the Phases field.
func (o *UpdateSubscriptionScheduleRequest) SetPhases(v []SchedulePhase) {
	o.Phases = v
}

// GetEndBehavior returns the EndBehavior field value if set, zero value otherwise.
func (o *UpdateSubscriptionScheduleRequest) GetEndBehavior() string {
	if o == nil || IsNil(o.EndBehavior) {
		var ret string
		return ret
	}
	return *o.EndBehavior
}

// GetEndBehaviorOk returns a tuple with the EndBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionScheduleRequest) GetEndBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.EndBehavior) {
		return nil, false
	}
	return o.EndBehavior, true
}

// HasEndBehavior returns a boolean if a field has been set.
func (o *UpdateSubscriptionScheduleRequest) HasEndBehavior() bool {
	if o != nil && !IsNil(o.EndBehavior) {
		return true
	}

	return false
}

// SetEndBehavior gets a reference to the given string and assigns it to the EndBehavior field.
func (o *UpdateSubscriptionScheduleRequest) SetEndBehavior(v string) {
	o.EndBehavior = &v
}

// GetProrationBehavior returns the ProrationBehavior field value if set, zero value otherwise.
func (o *UpdateSubscriptionScheduleRequest) GetProrationBehavior() string {
	if o == nil || IsNil(o.ProrationBehavior) {
		var ret string
		return ret
	}
	return *o.ProrationBehavior
}

// GetProrationBehaviorOk returns a tuple with the ProrationBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionScheduleRequest) GetProrationBehaviorOk() (*string, bool) {
	if o == nil || IsNil(o.ProrationBehavior) {
		return nil, false
	}
	return o.ProrationBehavior, true
}

// HasProrationBehavior returns a boolean if a field has been set.
func (o *UpdateSubscriptionScheduleRequest) HasProrationBehavior() bool {
	if o != nil && !IsNil(o.ProrationBehavior) {
		return true
	}

	return false
}

// SetProrationBehavior gets a reference to the given string and assigns it to the ProrationBehavior field.
func (o *UpdateSubscriptionScheduleRequest) SetProrationBehavior(v string) {
	o.ProrationBehavior = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateSubscriptionScheduleRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionScheduleRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateSubscriptionScheduleRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdateSubscriptionScheduleRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UpdateSubscriptionScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSubscriptionScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phases) {
		toSerialize["phases"] = o.Phases
	}
	if !IsNil(o.EndBehavior) {
		toSerialize["end_behavior"] = o.EndBehavior
	}
	if !IsNil(o.ProrationBehavior) {
		toSerialize["proration_behavior"] = o.ProrationBehavior
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateSubscriptionScheduleRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateSubscriptionScheduleRequest := _UpdateSubscriptionScheduleRequest{}

	err = json.Unmarshal(data, &varUpdateSubscriptionScheduleRequest)

	if err != nil {
		return err
	}

	*o = UpdateSubscriptionScheduleRequest(varUpdateSubscriptionScheduleRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "phases")
		delete(additionalProperties, "end_behavior")
		delete(additionalProperties, "proration_behavior")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateSubscriptionScheduleRequest struct {
	value *UpdateSubscriptionScheduleRequest
	isSet bool
}

func (v NullableUpdateSubscriptionScheduleRequest) Get() *UpdateSubscriptionScheduleRequest {
	return v.value
}

func (v *NullableUpdateSubscriptionScheduleRequest) Set(val *UpdateSubscriptionScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionScheduleRequest(val *UpdateSubscriptionScheduleRequest) *NullableUpdateSubscriptionScheduleRequest {
	return &NullableUpdateSubscriptionScheduleRequest{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


