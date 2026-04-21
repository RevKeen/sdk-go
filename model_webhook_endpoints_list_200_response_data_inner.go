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

// checks if the WebhookEndpointsList200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEndpointsList200ResponseDataInner{}

// WebhookEndpointsList200ResponseDataInner struct for WebhookEndpointsList200ResponseDataInner
type WebhookEndpointsList200ResponseDataInner struct {
	// Unique identifier for the webhook endpoint
	Id string `json:"id"`
	// The URL that receives webhook POST requests
	Url string `json:"url"`
	// Optional human-readable description
	Description NullableString `json:"description"`
	// Event types this endpoint subscribes to
	EnabledEvents []string `json:"enabled_events"`
	// Whether the endpoint is currently receiving events
	Status string `json:"status"`
	// Circuit breaker status. Opens after repeated delivery failures.
	CircuitBreakerState string `json:"circuit_breaker_state"`
	// Total number of delivery attempts
	TotalDeliveries float32 `json:"total_deliveries"`
	// Number of successful deliveries
	SuccessfulDeliveries float32 `json:"successful_deliveries"`
	// Number of failed deliveries
	FailedDeliveries float32 `json:"failed_deliveries"`
	// Timestamp of the most recent delivery attempt
	LastDeliveryAt NullableString `json:"last_delivery_at"`
	// When the endpoint was created
	CreatedAt string `json:"created_at"`
	// When the endpoint was last modified
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _WebhookEndpointsList200ResponseDataInner WebhookEndpointsList200ResponseDataInner

// NewWebhookEndpointsList200ResponseDataInner instantiates a new WebhookEndpointsList200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEndpointsList200ResponseDataInner(id string, url string, description NullableString, enabledEvents []string, status string, circuitBreakerState string, totalDeliveries float32, successfulDeliveries float32, failedDeliveries float32, lastDeliveryAt NullableString, createdAt string, updatedAt string) *WebhookEndpointsList200ResponseDataInner {
	this := WebhookEndpointsList200ResponseDataInner{}
	this.Id = id
	this.Url = url
	this.Description = description
	this.EnabledEvents = enabledEvents
	this.Status = status
	this.CircuitBreakerState = circuitBreakerState
	this.TotalDeliveries = totalDeliveries
	this.SuccessfulDeliveries = successfulDeliveries
	this.FailedDeliveries = failedDeliveries
	this.LastDeliveryAt = lastDeliveryAt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewWebhookEndpointsList200ResponseDataInnerWithDefaults instantiates a new WebhookEndpointsList200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEndpointsList200ResponseDataInnerWithDefaults() *WebhookEndpointsList200ResponseDataInner {
	this := WebhookEndpointsList200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookEndpointsList200ResponseDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WebhookEndpointsList200ResponseDataInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetUrl(v string) {
	o.Url = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookEndpointsList200ResponseDataInner) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEndpointsList200ResponseDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetEnabledEvents returns the EnabledEvents field value
func (o *WebhookEndpointsList200ResponseDataInner) GetEnabledEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnabledEvents
}

// GetEnabledEventsOk returns a tuple with the EnabledEvents field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetEnabledEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnabledEvents, true
}

// SetEnabledEvents sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetEnabledEvents(v []string) {
	o.EnabledEvents = v
}

// GetStatus returns the Status field value
func (o *WebhookEndpointsList200ResponseDataInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetStatus(v string) {
	o.Status = v
}

// GetCircuitBreakerState returns the CircuitBreakerState field value
func (o *WebhookEndpointsList200ResponseDataInner) GetCircuitBreakerState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CircuitBreakerState
}

// GetCircuitBreakerStateOk returns a tuple with the CircuitBreakerState field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetCircuitBreakerStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CircuitBreakerState, true
}

// SetCircuitBreakerState sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetCircuitBreakerState(v string) {
	o.CircuitBreakerState = v
}

// GetTotalDeliveries returns the TotalDeliveries field value
func (o *WebhookEndpointsList200ResponseDataInner) GetTotalDeliveries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalDeliveries
}

// GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetTotalDeliveriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDeliveries, true
}

// SetTotalDeliveries sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetTotalDeliveries(v float32) {
	o.TotalDeliveries = v
}

// GetSuccessfulDeliveries returns the SuccessfulDeliveries field value
func (o *WebhookEndpointsList200ResponseDataInner) GetSuccessfulDeliveries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SuccessfulDeliveries
}

// GetSuccessfulDeliveriesOk returns a tuple with the SuccessfulDeliveries field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetSuccessfulDeliveriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessfulDeliveries, true
}

// SetSuccessfulDeliveries sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetSuccessfulDeliveries(v float32) {
	o.SuccessfulDeliveries = v
}

// GetFailedDeliveries returns the FailedDeliveries field value
func (o *WebhookEndpointsList200ResponseDataInner) GetFailedDeliveries() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FailedDeliveries
}

// GetFailedDeliveriesOk returns a tuple with the FailedDeliveries field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetFailedDeliveriesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedDeliveries, true
}

// SetFailedDeliveries sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetFailedDeliveries(v float32) {
	o.FailedDeliveries = v
}

// GetLastDeliveryAt returns the LastDeliveryAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebhookEndpointsList200ResponseDataInner) GetLastDeliveryAt() string {
	if o == nil || o.LastDeliveryAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastDeliveryAt.Get()
}

// GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookEndpointsList200ResponseDataInner) GetLastDeliveryAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastDeliveryAt.Get(), o.LastDeliveryAt.IsSet()
}

// SetLastDeliveryAt sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetLastDeliveryAt(v string) {
	o.LastDeliveryAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookEndpointsList200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WebhookEndpointsList200ResponseDataInner) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookEndpointsList200ResponseDataInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WebhookEndpointsList200ResponseDataInner) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o WebhookEndpointsList200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEndpointsList200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["description"] = o.Description.Get()
	toSerialize["enabled_events"] = o.EnabledEvents
	toSerialize["status"] = o.Status
	toSerialize["circuit_breaker_state"] = o.CircuitBreakerState
	toSerialize["total_deliveries"] = o.TotalDeliveries
	toSerialize["successful_deliveries"] = o.SuccessfulDeliveries
	toSerialize["failed_deliveries"] = o.FailedDeliveries
	toSerialize["last_delivery_at"] = o.LastDeliveryAt.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookEndpointsList200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"description",
		"enabled_events",
		"status",
		"circuit_breaker_state",
		"total_deliveries",
		"successful_deliveries",
		"failed_deliveries",
		"last_delivery_at",
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

	varWebhookEndpointsList200ResponseDataInner := _WebhookEndpointsList200ResponseDataInner{}

	err = json.Unmarshal(data, &varWebhookEndpointsList200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = WebhookEndpointsList200ResponseDataInner(varWebhookEndpointsList200ResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled_events")
		delete(additionalProperties, "status")
		delete(additionalProperties, "circuit_breaker_state")
		delete(additionalProperties, "total_deliveries")
		delete(additionalProperties, "successful_deliveries")
		delete(additionalProperties, "failed_deliveries")
		delete(additionalProperties, "last_delivery_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookEndpointsList200ResponseDataInner struct {
	value *WebhookEndpointsList200ResponseDataInner
	isSet bool
}

func (v NullableWebhookEndpointsList200ResponseDataInner) Get() *WebhookEndpointsList200ResponseDataInner {
	return v.value
}

func (v *NullableWebhookEndpointsList200ResponseDataInner) Set(val *WebhookEndpointsList200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEndpointsList200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEndpointsList200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEndpointsList200ResponseDataInner(val *WebhookEndpointsList200ResponseDataInner) *NullableWebhookEndpointsList200ResponseDataInner {
	return &NullableWebhookEndpointsList200ResponseDataInner{value: val, isSet: true}
}

func (v NullableWebhookEndpointsList200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEndpointsList200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


