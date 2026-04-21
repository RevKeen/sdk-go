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
	"time"
	"fmt"
)

// checks if the BillingPreviewResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPreviewResponseData{}

// BillingPreviewResponseData struct for BillingPreviewResponseData
type BillingPreviewResponseData struct {
	Items []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner `json:"items"`
	Summary string `json:"summary"`
	TotalAmountMinor int32 `json:"totalAmountMinor"`
	FirstChargeDate time.Time `json:"firstChargeDate"`
	TrialEndDate NullableTime `json:"trialEndDate"`
	ScheduleEndDate NullableTime `json:"scheduleEndDate"`
	Timezone string `json:"timezone"`
	AdditionalProperties map[string]interface{}
}

type _BillingPreviewResponseData BillingPreviewResponseData

// NewBillingPreviewResponseData instantiates a new BillingPreviewResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPreviewResponseData(items []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, summary string, totalAmountMinor int32, firstChargeDate time.Time, trialEndDate NullableTime, scheduleEndDate NullableTime, timezone string) *BillingPreviewResponseData {
	this := BillingPreviewResponseData{}
	this.Items = items
	this.Summary = summary
	this.TotalAmountMinor = totalAmountMinor
	this.FirstChargeDate = firstChargeDate
	this.TrialEndDate = trialEndDate
	this.ScheduleEndDate = scheduleEndDate
	this.Timezone = timezone
	return &this
}

// NewBillingPreviewResponseDataWithDefaults instantiates a new BillingPreviewResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPreviewResponseDataWithDefaults() *BillingPreviewResponseData {
	this := BillingPreviewResponseData{}
	return &this
}

// GetItems returns the Items field value
func (o *BillingPreviewResponseData) GetItems() []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner {
	if o == nil {
		var ret []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *BillingPreviewResponseData) GetItemsOk() ([]SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *BillingPreviewResponseData) SetItems(v []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) {
	o.Items = v
}

// GetSummary returns the Summary field value
func (o *BillingPreviewResponseData) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *BillingPreviewResponseData) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value
func (o *BillingPreviewResponseData) SetSummary(v string) {
	o.Summary = v
}

// GetTotalAmountMinor returns the TotalAmountMinor field value
func (o *BillingPreviewResponseData) GetTotalAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalAmountMinor
}

// GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field value
// and a boolean to check if the value has been set.
func (o *BillingPreviewResponseData) GetTotalAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountMinor, true
}

// SetTotalAmountMinor sets field value
func (o *BillingPreviewResponseData) SetTotalAmountMinor(v int32) {
	o.TotalAmountMinor = v
}

// GetFirstChargeDate returns the FirstChargeDate field value
func (o *BillingPreviewResponseData) GetFirstChargeDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.FirstChargeDate
}

// GetFirstChargeDateOk returns a tuple with the FirstChargeDate field value
// and a boolean to check if the value has been set.
func (o *BillingPreviewResponseData) GetFirstChargeDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstChargeDate, true
}

// SetFirstChargeDate sets field value
func (o *BillingPreviewResponseData) SetFirstChargeDate(v time.Time) {
	o.FirstChargeDate = v
}

// GetTrialEndDate returns the TrialEndDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BillingPreviewResponseData) GetTrialEndDate() time.Time {
	if o == nil || o.TrialEndDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.TrialEndDate.Get()
}

// GetTrialEndDateOk returns a tuple with the TrialEndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingPreviewResponseData) GetTrialEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrialEndDate.Get(), o.TrialEndDate.IsSet()
}

// SetTrialEndDate sets field value
func (o *BillingPreviewResponseData) SetTrialEndDate(v time.Time) {
	o.TrialEndDate.Set(&v)
}

// GetScheduleEndDate returns the ScheduleEndDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *BillingPreviewResponseData) GetScheduleEndDate() time.Time {
	if o == nil || o.ScheduleEndDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ScheduleEndDate.Get()
}

// GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingPreviewResponseData) GetScheduleEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleEndDate.Get(), o.ScheduleEndDate.IsSet()
}

// SetScheduleEndDate sets field value
func (o *BillingPreviewResponseData) SetScheduleEndDate(v time.Time) {
	o.ScheduleEndDate.Set(&v)
}

// GetTimezone returns the Timezone field value
func (o *BillingPreviewResponseData) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *BillingPreviewResponseData) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *BillingPreviewResponseData) SetTimezone(v string) {
	o.Timezone = v
}

func (o BillingPreviewResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPreviewResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["summary"] = o.Summary
	toSerialize["totalAmountMinor"] = o.TotalAmountMinor
	toSerialize["firstChargeDate"] = o.FirstChargeDate
	toSerialize["trialEndDate"] = o.TrialEndDate.Get()
	toSerialize["scheduleEndDate"] = o.ScheduleEndDate.Get()
	toSerialize["timezone"] = o.Timezone

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BillingPreviewResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"summary",
		"totalAmountMinor",
		"firstChargeDate",
		"trialEndDate",
		"scheduleEndDate",
		"timezone",
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

	varBillingPreviewResponseData := _BillingPreviewResponseData{}

	err = json.Unmarshal(data, &varBillingPreviewResponseData)

	if err != nil {
		return err
	}

	*o = BillingPreviewResponseData(varBillingPreviewResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		delete(additionalProperties, "summary")
		delete(additionalProperties, "totalAmountMinor")
		delete(additionalProperties, "firstChargeDate")
		delete(additionalProperties, "trialEndDate")
		delete(additionalProperties, "scheduleEndDate")
		delete(additionalProperties, "timezone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBillingPreviewResponseData struct {
	value *BillingPreviewResponseData
	isSet bool
}

func (v NullableBillingPreviewResponseData) Get() *BillingPreviewResponseData {
	return v.value
}

func (v *NullableBillingPreviewResponseData) Set(val *BillingPreviewResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPreviewResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPreviewResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPreviewResponseData(val *BillingPreviewResponseData) *NullableBillingPreviewResponseData {
	return &NullableBillingPreviewResponseData{value: val, isSet: true}
}

func (v NullableBillingPreviewResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPreviewResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


