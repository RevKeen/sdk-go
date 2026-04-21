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

// checks if the AnalyticsRevenueMrrSummary200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsRevenueMrrSummary200Response{}

// AnalyticsRevenueMrrSummary200Response struct for AnalyticsRevenueMrrSummary200Response
type AnalyticsRevenueMrrSummary200Response struct {
	CurrentMrr float32 `json:"currentMrr"`
	PreviousMrr float32 `json:"previousMrr"`
	Growth float32 `json:"growth"`
	NewMrr float32 `json:"newMrr"`
	ChurnedMrr float32 `json:"churnedMrr"`
	ExpansionMrr float32 `json:"expansionMrr"`
	ContractionMrr float32 `json:"contractionMrr"`
	AsOfDate string `json:"asOfDate"`
	AdditionalProperties map[string]interface{}
}

type _AnalyticsRevenueMrrSummary200Response AnalyticsRevenueMrrSummary200Response

// NewAnalyticsRevenueMrrSummary200Response instantiates a new AnalyticsRevenueMrrSummary200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsRevenueMrrSummary200Response(currentMrr float32, previousMrr float32, growth float32, newMrr float32, churnedMrr float32, expansionMrr float32, contractionMrr float32, asOfDate string) *AnalyticsRevenueMrrSummary200Response {
	this := AnalyticsRevenueMrrSummary200Response{}
	this.CurrentMrr = currentMrr
	this.PreviousMrr = previousMrr
	this.Growth = growth
	this.NewMrr = newMrr
	this.ChurnedMrr = churnedMrr
	this.ExpansionMrr = expansionMrr
	this.ContractionMrr = contractionMrr
	this.AsOfDate = asOfDate
	return &this
}

// NewAnalyticsRevenueMrrSummary200ResponseWithDefaults instantiates a new AnalyticsRevenueMrrSummary200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsRevenueMrrSummary200ResponseWithDefaults() *AnalyticsRevenueMrrSummary200Response {
	this := AnalyticsRevenueMrrSummary200Response{}
	return &this
}

// GetCurrentMrr returns the CurrentMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetCurrentMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CurrentMrr
}

// GetCurrentMrrOk returns a tuple with the CurrentMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetCurrentMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentMrr, true
}

// SetCurrentMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetCurrentMrr(v float32) {
	o.CurrentMrr = v
}

// GetPreviousMrr returns the PreviousMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetPreviousMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PreviousMrr
}

// GetPreviousMrrOk returns a tuple with the PreviousMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetPreviousMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousMrr, true
}

// SetPreviousMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetPreviousMrr(v float32) {
	o.PreviousMrr = v
}

// GetGrowth returns the Growth field value
func (o *AnalyticsRevenueMrrSummary200Response) GetGrowth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Growth
}

// GetGrowthOk returns a tuple with the Growth field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetGrowthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Growth, true
}

// SetGrowth sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetGrowth(v float32) {
	o.Growth = v
}

// GetNewMrr returns the NewMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetNewMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewMrr
}

// GetNewMrrOk returns a tuple with the NewMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetNewMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewMrr, true
}

// SetNewMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetNewMrr(v float32) {
	o.NewMrr = v
}

// GetChurnedMrr returns the ChurnedMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetChurnedMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ChurnedMrr
}

// GetChurnedMrrOk returns a tuple with the ChurnedMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetChurnedMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChurnedMrr, true
}

// SetChurnedMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetChurnedMrr(v float32) {
	o.ChurnedMrr = v
}

// GetExpansionMrr returns the ExpansionMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetExpansionMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ExpansionMrr
}

// GetExpansionMrrOk returns a tuple with the ExpansionMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetExpansionMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpansionMrr, true
}

// SetExpansionMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetExpansionMrr(v float32) {
	o.ExpansionMrr = v
}

// GetContractionMrr returns the ContractionMrr field value
func (o *AnalyticsRevenueMrrSummary200Response) GetContractionMrr() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ContractionMrr
}

// GetContractionMrrOk returns a tuple with the ContractionMrr field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetContractionMrrOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractionMrr, true
}

// SetContractionMrr sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetContractionMrr(v float32) {
	o.ContractionMrr = v
}

// GetAsOfDate returns the AsOfDate field value
func (o *AnalyticsRevenueMrrSummary200Response) GetAsOfDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AsOfDate
}

// GetAsOfDateOk returns a tuple with the AsOfDate field value
// and a boolean to check if the value has been set.
func (o *AnalyticsRevenueMrrSummary200Response) GetAsOfDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsOfDate, true
}

// SetAsOfDate sets field value
func (o *AnalyticsRevenueMrrSummary200Response) SetAsOfDate(v string) {
	o.AsOfDate = v
}

func (o AnalyticsRevenueMrrSummary200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsRevenueMrrSummary200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentMrr"] = o.CurrentMrr
	toSerialize["previousMrr"] = o.PreviousMrr
	toSerialize["growth"] = o.Growth
	toSerialize["newMrr"] = o.NewMrr
	toSerialize["churnedMrr"] = o.ChurnedMrr
	toSerialize["expansionMrr"] = o.ExpansionMrr
	toSerialize["contractionMrr"] = o.ContractionMrr
	toSerialize["asOfDate"] = o.AsOfDate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnalyticsRevenueMrrSummary200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentMrr",
		"previousMrr",
		"growth",
		"newMrr",
		"churnedMrr",
		"expansionMrr",
		"contractionMrr",
		"asOfDate",
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

	varAnalyticsRevenueMrrSummary200Response := _AnalyticsRevenueMrrSummary200Response{}

	err = json.Unmarshal(data, &varAnalyticsRevenueMrrSummary200Response)

	if err != nil {
		return err
	}

	*o = AnalyticsRevenueMrrSummary200Response(varAnalyticsRevenueMrrSummary200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currentMrr")
		delete(additionalProperties, "previousMrr")
		delete(additionalProperties, "growth")
		delete(additionalProperties, "newMrr")
		delete(additionalProperties, "churnedMrr")
		delete(additionalProperties, "expansionMrr")
		delete(additionalProperties, "contractionMrr")
		delete(additionalProperties, "asOfDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnalyticsRevenueMrrSummary200Response struct {
	value *AnalyticsRevenueMrrSummary200Response
	isSet bool
}

func (v NullableAnalyticsRevenueMrrSummary200Response) Get() *AnalyticsRevenueMrrSummary200Response {
	return v.value
}

func (v *NullableAnalyticsRevenueMrrSummary200Response) Set(val *AnalyticsRevenueMrrSummary200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsRevenueMrrSummary200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsRevenueMrrSummary200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsRevenueMrrSummary200Response(val *AnalyticsRevenueMrrSummary200Response) *NullableAnalyticsRevenueMrrSummary200Response {
	return &NullableAnalyticsRevenueMrrSummary200Response{value: val, isSet: true}
}

func (v NullableAnalyticsRevenueMrrSummary200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsRevenueMrrSummary200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


