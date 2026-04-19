/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the SubscriptionChangePlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionChangePlanResponse{}

// SubscriptionChangePlanResponse The subscription after a plan change, with proration details and updated billing schedule.
type SubscriptionChangePlanResponse struct {
	Subscription SubscriptionCreateResponseData `json:"subscription"`
	// Proration details
	Proration map[string]interface{} `json:"proration,omitempty"`
	PreviousPlanId *string `json:"previousPlanId,omitempty"`
	InvoiceCreated bool `json:"invoiceCreated"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionChangePlanResponse SubscriptionChangePlanResponse

// NewSubscriptionChangePlanResponse instantiates a new SubscriptionChangePlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionChangePlanResponse(subscription SubscriptionCreateResponseData, invoiceCreated bool) *SubscriptionChangePlanResponse {
	this := SubscriptionChangePlanResponse{}
	this.Subscription = subscription
	this.InvoiceCreated = invoiceCreated
	return &this
}

// NewSubscriptionChangePlanResponseWithDefaults instantiates a new SubscriptionChangePlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionChangePlanResponseWithDefaults() *SubscriptionChangePlanResponse {
	this := SubscriptionChangePlanResponse{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *SubscriptionChangePlanResponse) GetSubscription() SubscriptionCreateResponseData {
	if o == nil {
		var ret SubscriptionCreateResponseData
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChangePlanResponse) GetSubscriptionOk() (*SubscriptionCreateResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SubscriptionChangePlanResponse) SetSubscription(v SubscriptionCreateResponseData) {
	o.Subscription = v
}

// GetProration returns the Proration field value if set, zero value otherwise.
func (o *SubscriptionChangePlanResponse) GetProration() map[string]interface{} {
	if o == nil || IsNil(o.Proration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Proration
}

// GetProrationOk returns a tuple with the Proration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionChangePlanResponse) GetProrationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Proration) {
		return map[string]interface{}{}, false
	}
	return o.Proration, true
}

// HasProration returns a boolean if a field has been set.
func (o *SubscriptionChangePlanResponse) HasProration() bool {
	if o != nil && !IsNil(o.Proration) {
		return true
	}

	return false
}

// SetProration gets a reference to the given map[string]interface{} and assigns it to the Proration field.
func (o *SubscriptionChangePlanResponse) SetProration(v map[string]interface{}) {
	o.Proration = v
}

// GetPreviousPlanId returns the PreviousPlanId field value if set, zero value otherwise.
func (o *SubscriptionChangePlanResponse) GetPreviousPlanId() string {
	if o == nil || IsNil(o.PreviousPlanId) {
		var ret string
		return ret
	}
	return *o.PreviousPlanId
}

// GetPreviousPlanIdOk returns a tuple with the PreviousPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionChangePlanResponse) GetPreviousPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPlanId) {
		return nil, false
	}
	return o.PreviousPlanId, true
}

// HasPreviousPlanId returns a boolean if a field has been set.
func (o *SubscriptionChangePlanResponse) HasPreviousPlanId() bool {
	if o != nil && !IsNil(o.PreviousPlanId) {
		return true
	}

	return false
}

// SetPreviousPlanId gets a reference to the given string and assigns it to the PreviousPlanId field.
func (o *SubscriptionChangePlanResponse) SetPreviousPlanId(v string) {
	o.PreviousPlanId = &v
}

// GetInvoiceCreated returns the InvoiceCreated field value
func (o *SubscriptionChangePlanResponse) GetInvoiceCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvoiceCreated
}

// GetInvoiceCreatedOk returns a tuple with the InvoiceCreated field value
// and a boolean to check if the value has been set.
func (o *SubscriptionChangePlanResponse) GetInvoiceCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceCreated, true
}

// SetInvoiceCreated sets field value
func (o *SubscriptionChangePlanResponse) SetInvoiceCreated(v bool) {
	o.InvoiceCreated = v
}

func (o SubscriptionChangePlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionChangePlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.Proration) {
		toSerialize["proration"] = o.Proration
	}
	if !IsNil(o.PreviousPlanId) {
		toSerialize["previousPlanId"] = o.PreviousPlanId
	}
	toSerialize["invoiceCreated"] = o.InvoiceCreated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionChangePlanResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscription",
		"invoiceCreated",
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

	varSubscriptionChangePlanResponse := _SubscriptionChangePlanResponse{}

	err = json.Unmarshal(data, &varSubscriptionChangePlanResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionChangePlanResponse(varSubscriptionChangePlanResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscription")
		delete(additionalProperties, "proration")
		delete(additionalProperties, "previousPlanId")
		delete(additionalProperties, "invoiceCreated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionChangePlanResponse struct {
	value *SubscriptionChangePlanResponse
	isSet bool
}

func (v NullableSubscriptionChangePlanResponse) Get() *SubscriptionChangePlanResponse {
	return v.value
}

func (v *NullableSubscriptionChangePlanResponse) Set(val *SubscriptionChangePlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionChangePlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionChangePlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionChangePlanResponse(val *SubscriptionChangePlanResponse) *NullableSubscriptionChangePlanResponse {
	return &NullableSubscriptionChangePlanResponse{value: val, isSet: true}
}

func (v NullableSubscriptionChangePlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionChangePlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


