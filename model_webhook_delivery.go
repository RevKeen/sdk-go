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

// checks if the WebhookDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookDelivery{}

// WebhookDelivery A record of a single webhook delivery attempt chain to one endpoint for one event. Use these records to debug specific delivery failures and to retry individual deliveries.
type WebhookDelivery struct {
	Id string `json:"id"`
	Object string `json:"object"`
	EndpointId string `json:"endpoint_id"`
	// The event this delivery attempted to send.
	EventId string `json:"event_id"`
	EventType NullableString `json:"event_type"`
	// Delivery state. pending: queued or retrying. succeeded: 2xx response received. failed: last attempt failed, may retry based on attempts/max_attempts. dead_lettered: max attempts exhausted, will not retry automatically.
	Status string `json:"status"`
	Attempts int32 `json:"attempts"`
	MaxAttempts int32 `json:"max_attempts"`
	// HTTP status code returned on the last attempt.
	LastStatusCode NullableInt32 `json:"last_status_code"`
	// Truncated error message from the last failed attempt.
	LastError NullableString `json:"last_error"`
	LastErrorCode NullableString `json:"last_error_code"`
	LastDurationMs NullableInt32 `json:"last_duration_ms"`
	LastAttemptAt NullableTime `json:"last_attempt_at"`
	NextRetryAt NullableTime `json:"next_retry_at"`
	DeliveredAt NullableTime `json:"delivered_at"`
	DeadLetteredAt NullableTime `json:"dead_lettered_at"`
	DeadLetterReason NullableString `json:"dead_letter_reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _WebhookDelivery WebhookDelivery

// NewWebhookDelivery instantiates a new WebhookDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookDelivery(id string, object string, endpointId string, eventId string, eventType NullableString, status string, attempts int32, maxAttempts int32, lastStatusCode NullableInt32, lastError NullableString, lastErrorCode NullableString, lastDurationMs NullableInt32, lastAttemptAt NullableTime, nextRetryAt NullableTime, deliveredAt NullableTime, deadLetteredAt NullableTime, deadLetterReason NullableString, createdAt time.Time, updatedAt time.Time) *WebhookDelivery {
	this := WebhookDelivery{}
	this.Id = id
	this.Object = object
	this.EndpointId = endpointId
	this.EventId = eventId
	this.EventType = eventType
	this.Status = status
	this.Attempts = attempts
	this.MaxAttempts = maxAttempts
	this.LastStatusCode = lastStatusCode
	this.LastError = lastError
	this.LastErrorCode = lastErrorCode
	this.LastDurationMs = lastDurationMs
	this.LastAttemptAt = lastAttemptAt
	this.NextRetryAt = nextRetryAt
	this.DeliveredAt = deliveredAt
	this.DeadLetteredAt = deadLetteredAt
	this.DeadLetterReason = deadLetterReason
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookDeliveryWithDefaults() *WebhookDelivery {
	this := WebhookDelivery{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookDelivery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookDelivery) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *WebhookDelivery) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *WebhookDelivery) SetObject(v string) {
	o.Object = v
}

// GetEndpointId returns the EndpointId field value
func (o *WebhookDelivery) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *WebhookDelivery) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetEventId returns the EventId field value
func (o *WebhookDelivery) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *WebhookDelivery) SetEventId(v string) {
	o.EventId = v
}

// GetEventType returns the EventType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookDelivery) GetEventType() string {
	if o == nil || o.EventType.Get() == nil {
		var ret string
		return ret
	}

	return *o.EventType.Get()
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventType.Get(), o.EventType.IsSet()
}

// SetEventType sets field value
func (o *WebhookDelivery) SetEventType(v string) {
	o.EventType.Set(&v)
}

// GetStatus returns the Status field value
func (o *WebhookDelivery) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebhookDelivery) SetStatus(v string) {
	o.Status = v
}

// GetAttempts returns the Attempts field value
func (o *WebhookDelivery) GetAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempts, true
}

// SetAttempts sets field value
func (o *WebhookDelivery) SetAttempts(v int32) {
	o.Attempts = v
}

// GetMaxAttempts returns the MaxAttempts field value
func (o *WebhookDelivery) GetMaxAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxAttempts
}

// GetMaxAttemptsOk returns a tuple with the MaxAttempts field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetMaxAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAttempts, true
}

// SetMaxAttempts sets field value
func (o *WebhookDelivery) SetMaxAttempts(v int32) {
	o.MaxAttempts = v
}

// GetLastStatusCode returns the LastStatusCode field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WebhookDelivery) GetLastStatusCode() int32 {
	if o == nil || o.LastStatusCode.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastStatusCode.Get()
}

// GetLastStatusCodeOk returns a tuple with the LastStatusCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetLastStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastStatusCode.Get(), o.LastStatusCode.IsSet()
}

// SetLastStatusCode sets field value
func (o *WebhookDelivery) SetLastStatusCode(v int32) {
	o.LastStatusCode.Set(&v)
}

// GetLastError returns the LastError field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookDelivery) GetLastError() string {
	if o == nil || o.LastError.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastError.Get()
}

// GetLastErrorOk returns a tuple with the LastError field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetLastErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastError.Get(), o.LastError.IsSet()
}

// SetLastError sets field value
func (o *WebhookDelivery) SetLastError(v string) {
	o.LastError.Set(&v)
}

// GetLastErrorCode returns the LastErrorCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookDelivery) GetLastErrorCode() string {
	if o == nil || o.LastErrorCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastErrorCode.Get()
}

// GetLastErrorCodeOk returns a tuple with the LastErrorCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetLastErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastErrorCode.Get(), o.LastErrorCode.IsSet()
}

// SetLastErrorCode sets field value
func (o *WebhookDelivery) SetLastErrorCode(v string) {
	o.LastErrorCode.Set(&v)
}

// GetLastDurationMs returns the LastDurationMs field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *WebhookDelivery) GetLastDurationMs() int32 {
	if o == nil || o.LastDurationMs.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastDurationMs.Get()
}

// GetLastDurationMsOk returns a tuple with the LastDurationMs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetLastDurationMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastDurationMs.Get(), o.LastDurationMs.IsSet()
}

// SetLastDurationMs sets field value
func (o *WebhookDelivery) SetLastDurationMs(v int32) {
	o.LastDurationMs.Set(&v)
}

// GetLastAttemptAt returns the LastAttemptAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WebhookDelivery) GetLastAttemptAt() time.Time {
	if o == nil || o.LastAttemptAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastAttemptAt.Get()
}

// GetLastAttemptAtOk returns a tuple with the LastAttemptAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetLastAttemptAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAttemptAt.Get(), o.LastAttemptAt.IsSet()
}

// SetLastAttemptAt sets field value
func (o *WebhookDelivery) SetLastAttemptAt(v time.Time) {
	o.LastAttemptAt.Set(&v)
}

// GetNextRetryAt returns the NextRetryAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WebhookDelivery) GetNextRetryAt() time.Time {
	if o == nil || o.NextRetryAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NextRetryAt.Get()
}

// GetNextRetryAtOk returns a tuple with the NextRetryAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetNextRetryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRetryAt.Get(), o.NextRetryAt.IsSet()
}

// SetNextRetryAt sets field value
func (o *WebhookDelivery) SetNextRetryAt(v time.Time) {
	o.NextRetryAt.Set(&v)
}

// GetDeliveredAt returns the DeliveredAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WebhookDelivery) GetDeliveredAt() time.Time {
	if o == nil || o.DeliveredAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeliveredAt.Get()
}

// GetDeliveredAtOk returns a tuple with the DeliveredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetDeliveredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveredAt.Get(), o.DeliveredAt.IsSet()
}

// SetDeliveredAt sets field value
func (o *WebhookDelivery) SetDeliveredAt(v time.Time) {
	o.DeliveredAt.Set(&v)
}

// GetDeadLetteredAt returns the DeadLetteredAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WebhookDelivery) GetDeadLetteredAt() time.Time {
	if o == nil || o.DeadLetteredAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DeadLetteredAt.Get()
}

// GetDeadLetteredAtOk returns a tuple with the DeadLetteredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetDeadLetteredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadLetteredAt.Get(), o.DeadLetteredAt.IsSet()
}

// SetDeadLetteredAt sets field value
func (o *WebhookDelivery) SetDeadLetteredAt(v time.Time) {
	o.DeadLetteredAt.Set(&v)
}

// GetDeadLetterReason returns the DeadLetterReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookDelivery) GetDeadLetterReason() string {
	if o == nil || o.DeadLetterReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeadLetterReason.Get()
}

// GetDeadLetterReasonOk returns a tuple with the DeadLetterReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookDelivery) GetDeadLetterReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeadLetterReason.Get(), o.DeadLetterReason.IsSet()
}

// SetDeadLetterReason sets field value
func (o *WebhookDelivery) SetDeadLetterReason(v string) {
	o.DeadLetterReason.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookDelivery) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookDelivery) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WebhookDelivery) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookDelivery) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WebhookDelivery) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o WebhookDelivery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["endpoint_id"] = o.EndpointId
	toSerialize["event_id"] = o.EventId
	toSerialize["event_type"] = o.EventType.Get()
	toSerialize["status"] = o.Status
	toSerialize["attempts"] = o.Attempts
	toSerialize["max_attempts"] = o.MaxAttempts
	toSerialize["last_status_code"] = o.LastStatusCode.Get()
	toSerialize["last_error"] = o.LastError.Get()
	toSerialize["last_error_code"] = o.LastErrorCode.Get()
	toSerialize["last_duration_ms"] = o.LastDurationMs.Get()
	toSerialize["last_attempt_at"] = o.LastAttemptAt.Get()
	toSerialize["next_retry_at"] = o.NextRetryAt.Get()
	toSerialize["delivered_at"] = o.DeliveredAt.Get()
	toSerialize["dead_lettered_at"] = o.DeadLetteredAt.Get()
	toSerialize["dead_letter_reason"] = o.DeadLetterReason.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookDelivery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"endpoint_id",
		"event_id",
		"event_type",
		"status",
		"attempts",
		"max_attempts",
		"last_status_code",
		"last_error",
		"last_error_code",
		"last_duration_ms",
		"last_attempt_at",
		"next_retry_at",
		"delivered_at",
		"dead_lettered_at",
		"dead_letter_reason",
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

	varWebhookDelivery := _WebhookDelivery{}

	err = json.Unmarshal(data, &varWebhookDelivery)

	if err != nil {
		return err
	}

	*o = WebhookDelivery(varWebhookDelivery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "endpoint_id")
		delete(additionalProperties, "event_id")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "attempts")
		delete(additionalProperties, "max_attempts")
		delete(additionalProperties, "last_status_code")
		delete(additionalProperties, "last_error")
		delete(additionalProperties, "last_error_code")
		delete(additionalProperties, "last_duration_ms")
		delete(additionalProperties, "last_attempt_at")
		delete(additionalProperties, "next_retry_at")
		delete(additionalProperties, "delivered_at")
		delete(additionalProperties, "dead_lettered_at")
		delete(additionalProperties, "dead_letter_reason")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookDelivery struct {
	value *WebhookDelivery
	isSet bool
}

func (v NullableWebhookDelivery) Get() *WebhookDelivery {
	return v.value
}

func (v *NullableWebhookDelivery) Set(val *WebhookDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookDelivery(val *WebhookDelivery) *NullableWebhookDelivery {
	return &NullableWebhookDelivery{value: val, isSet: true}
}

func (v NullableWebhookDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


