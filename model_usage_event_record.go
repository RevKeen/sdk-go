/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the UsageEventRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageEventRecord{}

// UsageEventRecord A persisted usage event record.
type UsageEventRecord struct {
	// Usage event ID
	Id string `json:"id"`
	// Owning merchant ID
	MerchantId string `json:"merchant_id"`
	// Customer ID
	CustomerId NullableString `json:"customer_id"`
	// Subscription ID
	SubscriptionId NullableString `json:"subscription_id"`
	// Meter ID
	MeterId NullableString `json:"meter_id"`
	// Event quantity
	Quantity float32 `json:"quantity"`
	// Event timestamp (ISO 8601)
	EventTime time.Time `json:"event_time"`
	// Deduplication key
	IdempotencyKey NullableString `json:"idempotency_key"`
	// External event identifier
	ExternalId NullableString `json:"external_id"`
	// Ingestion source
	Source string `json:"source"`
	// Arbitrary event metadata
	Metadata map[string]interface{} `json:"metadata"`
	// Arbitrary event properties used for filtering/aggregation
	Properties map[string]interface{} `json:"properties"`
	// When the event was ingested (ISO 8601)
	IngestionTimestamp NullableTime `json:"ingestion_timestamp"`
	// Creation timestamp (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _UsageEventRecord UsageEventRecord

// NewUsageEventRecord instantiates a new UsageEventRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageEventRecord(id string, merchantId string, customerId NullableString, subscriptionId NullableString, meterId NullableString, quantity float32, eventTime time.Time, idempotencyKey NullableString, externalId NullableString, source string, metadata map[string]interface{}, properties map[string]interface{}, ingestionTimestamp NullableTime, createdAt time.Time) *UsageEventRecord {
	this := UsageEventRecord{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.SubscriptionId = subscriptionId
	this.MeterId = meterId
	this.Quantity = quantity
	this.EventTime = eventTime
	this.IdempotencyKey = idempotencyKey
	this.ExternalId = externalId
	this.Source = source
	this.Metadata = metadata
	this.Properties = properties
	this.IngestionTimestamp = ingestionTimestamp
	this.CreatedAt = createdAt
	return &this
}

// NewUsageEventRecordWithDefaults instantiates a new UsageEventRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageEventRecordWithDefaults() *UsageEventRecord {
	this := UsageEventRecord{}
	return &this
}

// GetId returns the Id field value
func (o *UsageEventRecord) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UsageEventRecord) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *UsageEventRecord) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *UsageEventRecord) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageEventRecord) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *UsageEventRecord) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageEventRecord) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *UsageEventRecord) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetMeterId returns the MeterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageEventRecord) GetMeterId() string {
	if o == nil || o.MeterId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MeterId.Get()
}

// GetMeterIdOk returns a tuple with the MeterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetMeterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MeterId.Get(), o.MeterId.IsSet()
}

// SetMeterId sets field value
func (o *UsageEventRecord) SetMeterId(v string) {
	o.MeterId.Set(&v)
}

// GetQuantity returns the Quantity field value
func (o *UsageEventRecord) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *UsageEventRecord) SetQuantity(v float32) {
	o.Quantity = v
}

// GetEventTime returns the EventTime field value
func (o *UsageEventRecord) GetEventTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetEventTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTime, true
}

// SetEventTime sets field value
func (o *UsageEventRecord) SetEventTime(v time.Time) {
	o.EventTime = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageEventRecord) GetIdempotencyKey() string {
	if o == nil || o.IdempotencyKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.IdempotencyKey.Get()
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdempotencyKey.Get(), o.IdempotencyKey.IsSet()
}

// SetIdempotencyKey sets field value
func (o *UsageEventRecord) SetIdempotencyKey(v string) {
	o.IdempotencyKey.Set(&v)
}

// GetExternalId returns the ExternalId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UsageEventRecord) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// SetExternalId sets field value
func (o *UsageEventRecord) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}

// GetSource returns the Source field value
func (o *UsageEventRecord) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *UsageEventRecord) SetSource(v string) {
	o.Source = v
}

// GetMetadata returns the Metadata field value
func (o *UsageEventRecord) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *UsageEventRecord) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetProperties returns the Properties field value
func (o *UsageEventRecord) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *UsageEventRecord) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetIngestionTimestamp returns the IngestionTimestamp field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *UsageEventRecord) GetIngestionTimestamp() time.Time {
	if o == nil || o.IngestionTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.IngestionTimestamp.Get()
}

// GetIngestionTimestampOk returns a tuple with the IngestionTimestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UsageEventRecord) GetIngestionTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestionTimestamp.Get(), o.IngestionTimestamp.IsSet()
}

// SetIngestionTimestamp sets field value
func (o *UsageEventRecord) SetIngestionTimestamp(v time.Time) {
	o.IngestionTimestamp.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *UsageEventRecord) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UsageEventRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UsageEventRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o UsageEventRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageEventRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	toSerialize["meter_id"] = o.MeterId.Get()
	toSerialize["quantity"] = o.Quantity
	toSerialize["event_time"] = o.EventTime
	toSerialize["idempotency_key"] = o.IdempotencyKey.Get()
	toSerialize["external_id"] = o.ExternalId.Get()
	toSerialize["source"] = o.Source
	toSerialize["metadata"] = o.Metadata
	toSerialize["properties"] = o.Properties
	toSerialize["ingestion_timestamp"] = o.IngestionTimestamp.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UsageEventRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"customer_id",
		"subscription_id",
		"meter_id",
		"quantity",
		"event_time",
		"idempotency_key",
		"external_id",
		"source",
		"metadata",
		"properties",
		"ingestion_timestamp",
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

	varUsageEventRecord := _UsageEventRecord{}

	err = json.Unmarshal(data, &varUsageEventRecord)

	if err != nil {
		return err
	}

	*o = UsageEventRecord(varUsageEventRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "meter_id")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "event_time")
		delete(additionalProperties, "idempotency_key")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "source")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "ingestion_timestamp")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUsageEventRecord struct {
	value *UsageEventRecord
	isSet bool
}

func (v NullableUsageEventRecord) Get() *UsageEventRecord {
	return v.value
}

func (v *NullableUsageEventRecord) Set(val *UsageEventRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageEventRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageEventRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageEventRecord(val *UsageEventRecord) *NullableUsageEventRecord {
	return &NullableUsageEventRecord{value: val, isSet: true}
}

func (v NullableUsageEventRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageEventRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


