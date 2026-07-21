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
	"fmt"
)

// checks if the CollectionSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionSchedule{}

// CollectionSchedule A scheduled one-off Direct Debit collection.
type CollectionSchedule struct {
	// Collection schedule ID
	Id NullableString `json:"id"`
	// Mandate the collection is scheduled against
	MandateId string `json:"mandate_id"`
	// Bureau schedule reference
	LzScheduleId string `json:"lz_schedule_id"`
	// Collection amount in minor units (pence)
	AmountMinor int32 `json:"amount_minor"`
	// ISO 4217 currency
	Currency string `json:"currency"`
	// invoice | payment_link
	SourceType string `json:"source_type"`
	// Invoice or payment-link ID the collection settles
	SourceId string `json:"source_id"`
	// Schedule status
	Status string `json:"status"`
	// Collection date (YYYY-MM-DD)
	CollectionDate NullableString `json:"collection_date"`
	// Advance-notice date (YYYY-MM-DD)
	NoticeDate NullableString `json:"notice_date"`
	// Bacs submission date (YYYY-MM-DD)
	SubmissionDate NullableString `json:"submission_date"`
	// True when an existing schedule for this source was returned instead of creating a duplicate
	AlreadyScheduled bool `json:"already_scheduled"`
	AdditionalProperties map[string]interface{}
}

type _CollectionSchedule CollectionSchedule

// NewCollectionSchedule instantiates a new CollectionSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionSchedule(id NullableString, mandateId string, lzScheduleId string, amountMinor int32, currency string, sourceType string, sourceId string, status string, collectionDate NullableString, noticeDate NullableString, submissionDate NullableString, alreadyScheduled bool) *CollectionSchedule {
	this := CollectionSchedule{}
	this.Id = id
	this.MandateId = mandateId
	this.LzScheduleId = lzScheduleId
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.SourceType = sourceType
	this.SourceId = sourceId
	this.Status = status
	this.CollectionDate = collectionDate
	this.NoticeDate = noticeDate
	this.SubmissionDate = submissionDate
	this.AlreadyScheduled = alreadyScheduled
	return &this
}

// NewCollectionScheduleWithDefaults instantiates a new CollectionSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionScheduleWithDefaults() *CollectionSchedule {
	this := CollectionSchedule{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CollectionSchedule) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionSchedule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CollectionSchedule) SetId(v string) {
	o.Id.Set(&v)
}

// GetMandateId returns the MandateId field value
func (o *CollectionSchedule) GetMandateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateId, true
}

// SetMandateId sets field value
func (o *CollectionSchedule) SetMandateId(v string) {
	o.MandateId = v
}

// GetLzScheduleId returns the LzScheduleId field value
func (o *CollectionSchedule) GetLzScheduleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LzScheduleId
}

// GetLzScheduleIdOk returns a tuple with the LzScheduleId field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetLzScheduleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LzScheduleId, true
}

// SetLzScheduleId sets field value
func (o *CollectionSchedule) SetLzScheduleId(v string) {
	o.LzScheduleId = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *CollectionSchedule) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *CollectionSchedule) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *CollectionSchedule) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CollectionSchedule) SetCurrency(v string) {
	o.Currency = v
}

// GetSourceType returns the SourceType field value
func (o *CollectionSchedule) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *CollectionSchedule) SetSourceType(v string) {
	o.SourceType = v
}

// GetSourceId returns the SourceId field value
func (o *CollectionSchedule) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CollectionSchedule) SetSourceId(v string) {
	o.SourceId = v
}

// GetStatus returns the Status field value
func (o *CollectionSchedule) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CollectionSchedule) SetStatus(v string) {
	o.Status = v
}

// GetCollectionDate returns the CollectionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CollectionSchedule) GetCollectionDate() string {
	if o == nil || o.CollectionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.CollectionDate.Get()
}

// GetCollectionDateOk returns a tuple with the CollectionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionSchedule) GetCollectionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionDate.Get(), o.CollectionDate.IsSet()
}

// SetCollectionDate sets field value
func (o *CollectionSchedule) SetCollectionDate(v string) {
	o.CollectionDate.Set(&v)
}

// GetNoticeDate returns the NoticeDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CollectionSchedule) GetNoticeDate() string {
	if o == nil || o.NoticeDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NoticeDate.Get()
}

// GetNoticeDateOk returns a tuple with the NoticeDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionSchedule) GetNoticeDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoticeDate.Get(), o.NoticeDate.IsSet()
}

// SetNoticeDate sets field value
func (o *CollectionSchedule) SetNoticeDate(v string) {
	o.NoticeDate.Set(&v)
}

// GetSubmissionDate returns the SubmissionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CollectionSchedule) GetSubmissionDate() string {
	if o == nil || o.SubmissionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubmissionDate.Get()
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectionSchedule) GetSubmissionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionDate.Get(), o.SubmissionDate.IsSet()
}

// SetSubmissionDate sets field value
func (o *CollectionSchedule) SetSubmissionDate(v string) {
	o.SubmissionDate.Set(&v)
}

// GetAlreadyScheduled returns the AlreadyScheduled field value
func (o *CollectionSchedule) GetAlreadyScheduled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlreadyScheduled
}

// GetAlreadyScheduledOk returns a tuple with the AlreadyScheduled field value
// and a boolean to check if the value has been set.
func (o *CollectionSchedule) GetAlreadyScheduledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlreadyScheduled, true
}

// SetAlreadyScheduled sets field value
func (o *CollectionSchedule) SetAlreadyScheduled(v bool) {
	o.AlreadyScheduled = v
}

func (o CollectionSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["mandate_id"] = o.MandateId
	toSerialize["lz_schedule_id"] = o.LzScheduleId
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["source_type"] = o.SourceType
	toSerialize["source_id"] = o.SourceId
	toSerialize["status"] = o.Status
	toSerialize["collection_date"] = o.CollectionDate.Get()
	toSerialize["notice_date"] = o.NoticeDate.Get()
	toSerialize["submission_date"] = o.SubmissionDate.Get()
	toSerialize["already_scheduled"] = o.AlreadyScheduled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"mandate_id",
		"lz_schedule_id",
		"amount_minor",
		"currency",
		"source_type",
		"source_id",
		"status",
		"collection_date",
		"notice_date",
		"submission_date",
		"already_scheduled",
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

	varCollectionSchedule := _CollectionSchedule{}

	err = json.Unmarshal(data, &varCollectionSchedule)

	if err != nil {
		return err
	}

	*o = CollectionSchedule(varCollectionSchedule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "lz_schedule_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "source_type")
		delete(additionalProperties, "source_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "collection_date")
		delete(additionalProperties, "notice_date")
		delete(additionalProperties, "submission_date")
		delete(additionalProperties, "already_scheduled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionSchedule struct {
	value *CollectionSchedule
	isSet bool
}

func (v NullableCollectionSchedule) Get() *CollectionSchedule {
	return v.value
}

func (v *NullableCollectionSchedule) Set(val *CollectionSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionSchedule(val *CollectionSchedule) *NullableCollectionSchedule {
	return &NullableCollectionSchedule{value: val, isSet: true}
}

func (v NullableCollectionSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


