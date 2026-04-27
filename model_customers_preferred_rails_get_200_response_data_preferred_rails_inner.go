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

// checks if the CustomersPreferredRailsGet200ResponseDataPreferredRailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomersPreferredRailsGet200ResponseDataPreferredRailsInner{}

// CustomersPreferredRailsGet200ResponseDataPreferredRailsInner struct for CustomersPreferredRailsGet200ResponseDataPreferredRailsInner
type CustomersPreferredRailsGet200ResponseDataPreferredRailsInner struct {
	Rail string `json:"rail"`
	Available bool `json:"available"`
	ReasonIfNot NullableString `json:"reason_if_not"`
	LastUsedAt NullableTime `json:"last_used_at"`
	MandateId *string `json:"mandate_id,omitempty"`
	Rank int32 `json:"rank"`
	Score float32 `json:"score"`
	Reason string `json:"reason"`
	Reasons []string `json:"reasons"`
	EstimatedFeeMinor *int32 `json:"estimated_fee_minor,omitempty"`
	NetAfterFeesMinor *int32 `json:"net_after_fees_minor,omitempty"`
	Mandate *CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate `json:"mandate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomersPreferredRailsGet200ResponseDataPreferredRailsInner CustomersPreferredRailsGet200ResponseDataPreferredRailsInner

// NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInner instantiates a new CustomersPreferredRailsGet200ResponseDataPreferredRailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInner(rail string, available bool, reasonIfNot NullableString, lastUsedAt NullableTime, rank int32, score float32, reason string, reasons []string) *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner {
	this := CustomersPreferredRailsGet200ResponseDataPreferredRailsInner{}
	this.Rail = rail
	this.Available = available
	this.ReasonIfNot = reasonIfNot
	this.LastUsedAt = lastUsedAt
	this.Rank = rank
	this.Score = score
	this.Reason = reason
	this.Reasons = reasons
	return &this
}

// NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInnerWithDefaults instantiates a new CustomersPreferredRailsGet200ResponseDataPreferredRailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInnerWithDefaults() *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner {
	this := CustomersPreferredRailsGet200ResponseDataPreferredRailsInner{}
	return &this
}

// GetRail returns the Rail field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rail
}

// GetRailOk returns a tuple with the Rail field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rail, true
}

// SetRail sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetRail(v string) {
	o.Rail = v
}

// GetAvailable returns the Available field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetAvailable(v bool) {
	o.Available = v
}

// GetReasonIfNot returns the ReasonIfNot field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonIfNot() string {
	if o == nil || o.ReasonIfNot.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonIfNot.Get()
}

// GetReasonIfNotOk returns a tuple with the ReasonIfNot field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonIfNotOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonIfNot.Get(), o.ReasonIfNot.IsSet()
}

// SetReasonIfNot sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReasonIfNot(v string) {
	o.ReasonIfNot.Set(&v)
}

// GetLastUsedAt returns the LastUsedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// SetLastUsedAt sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// GetMandateId returns the MandateId field value if set, zero value otherwise.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateId() string {
	if o == nil || IsNil(o.MandateId) {
		var ret string
		return ret
	}
	return *o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateIdOk() (*string, bool) {
	if o == nil || IsNil(o.MandateId) {
		return nil, false
	}
	return o.MandateId, true
}

// HasMandateId returns a boolean if a field has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasMandateId() bool {
	if o != nil && !IsNil(o.MandateId) {
		return true
	}

	return false
}

// SetMandateId gets a reference to the given string and assigns it to the MandateId field.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetMandateId(v string) {
	o.MandateId = &v
}

// GetRank returns the Rank field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetRank(v int32) {
	o.Rank = v
}

// GetScore returns the Score field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetScore(v float32) {
	o.Score = v
}

// GetReason returns the Reason field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReason(v string) {
	o.Reason = v
}

// GetReasons returns the Reasons field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReasons(v []string) {
	o.Reasons = v
}

// GetEstimatedFeeMinor returns the EstimatedFeeMinor field value if set, zero value otherwise.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetEstimatedFeeMinor() int32 {
	if o == nil || IsNil(o.EstimatedFeeMinor) {
		var ret int32
		return ret
	}
	return *o.EstimatedFeeMinor
}

// GetEstimatedFeeMinorOk returns a tuple with the EstimatedFeeMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetEstimatedFeeMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedFeeMinor) {
		return nil, false
	}
	return o.EstimatedFeeMinor, true
}

// HasEstimatedFeeMinor returns a boolean if a field has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasEstimatedFeeMinor() bool {
	if o != nil && !IsNil(o.EstimatedFeeMinor) {
		return true
	}

	return false
}

// SetEstimatedFeeMinor gets a reference to the given int32 and assigns it to the EstimatedFeeMinor field.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetEstimatedFeeMinor(v int32) {
	o.EstimatedFeeMinor = &v
}

// GetNetAfterFeesMinor returns the NetAfterFeesMinor field value if set, zero value otherwise.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetNetAfterFeesMinor() int32 {
	if o == nil || IsNil(o.NetAfterFeesMinor) {
		var ret int32
		return ret
	}
	return *o.NetAfterFeesMinor
}

// GetNetAfterFeesMinorOk returns a tuple with the NetAfterFeesMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetNetAfterFeesMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.NetAfterFeesMinor) {
		return nil, false
	}
	return o.NetAfterFeesMinor, true
}

// HasNetAfterFeesMinor returns a boolean if a field has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasNetAfterFeesMinor() bool {
	if o != nil && !IsNil(o.NetAfterFeesMinor) {
		return true
	}

	return false
}

// SetNetAfterFeesMinor gets a reference to the given int32 and assigns it to the NetAfterFeesMinor field.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetNetAfterFeesMinor(v int32) {
	o.NetAfterFeesMinor = &v
}

// GetMandate returns the Mandate field value if set, zero value otherwise.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandate() CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate {
	if o == nil || IsNil(o.Mandate) {
		var ret CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate
		return ret
	}
	return *o.Mandate
}

// GetMandateOk returns a tuple with the Mandate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateOk() (*CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate, bool) {
	if o == nil || IsNil(o.Mandate) {
		return nil, false
	}
	return o.Mandate, true
}

// HasMandate returns a boolean if a field has been set.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasMandate() bool {
	if o != nil && !IsNil(o.Mandate) {
		return true
	}

	return false
}

// SetMandate gets a reference to the given CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate and assigns it to the Mandate field.
func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetMandate(v CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate) {
	o.Mandate = &v
}

func (o CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rail"] = o.Rail
	toSerialize["available"] = o.Available
	toSerialize["reason_if_not"] = o.ReasonIfNot.Get()
	toSerialize["last_used_at"] = o.LastUsedAt.Get()
	if !IsNil(o.MandateId) {
		toSerialize["mandate_id"] = o.MandateId
	}
	toSerialize["rank"] = o.Rank
	toSerialize["score"] = o.Score
	toSerialize["reason"] = o.Reason
	toSerialize["reasons"] = o.Reasons
	if !IsNil(o.EstimatedFeeMinor) {
		toSerialize["estimated_fee_minor"] = o.EstimatedFeeMinor
	}
	if !IsNil(o.NetAfterFeesMinor) {
		toSerialize["net_after_fees_minor"] = o.NetAfterFeesMinor
	}
	if !IsNil(o.Mandate) {
		toSerialize["mandate"] = o.Mandate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rail",
		"available",
		"reason_if_not",
		"last_used_at",
		"rank",
		"score",
		"reason",
		"reasons",
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

	varCustomersPreferredRailsGet200ResponseDataPreferredRailsInner := _CustomersPreferredRailsGet200ResponseDataPreferredRailsInner{}

	err = json.Unmarshal(data, &varCustomersPreferredRailsGet200ResponseDataPreferredRailsInner)

	if err != nil {
		return err
	}

	*o = CustomersPreferredRailsGet200ResponseDataPreferredRailsInner(varCustomersPreferredRailsGet200ResponseDataPreferredRailsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rail")
		delete(additionalProperties, "available")
		delete(additionalProperties, "reason_if_not")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "score")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "reasons")
		delete(additionalProperties, "estimated_fee_minor")
		delete(additionalProperties, "net_after_fees_minor")
		delete(additionalProperties, "mandate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner struct {
	value *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner
	isSet bool
}

func (v NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) Get() *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner {
	return v.value
}

func (v *NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) Set(val *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner(val *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) *NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner {
	return &NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner{value: val, isSet: true}
}

func (v NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomersPreferredRailsGet200ResponseDataPreferredRailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


