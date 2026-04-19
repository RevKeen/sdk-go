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
)

// checks if the SubscriptionsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsUpdateRequest{}

// SubscriptionsUpdateRequest struct for SubscriptionsUpdateRequest
type SubscriptionsUpdateRequest struct {
	Status *string `json:"status,omitempty"`
	AmountMinor *int32 `json:"amountMinor,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	NextBillingDate *string `json:"nextBillingDate,omitempty"`
	BillingAnchorDay *int32 `json:"billingAnchorDay,omitempty"`
	PauseAtPeriodEnd *bool `json:"pauseAtPeriodEnd,omitempty"`
	CancelAtPeriodEnd *bool `json:"cancelAtPeriodEnd,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionsUpdateRequest SubscriptionsUpdateRequest

// NewSubscriptionsUpdateRequest instantiates a new SubscriptionsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsUpdateRequest() *SubscriptionsUpdateRequest {
	this := SubscriptionsUpdateRequest{}
	return &this
}

// NewSubscriptionsUpdateRequestWithDefaults instantiates a new SubscriptionsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsUpdateRequestWithDefaults() *SubscriptionsUpdateRequest {
	this := SubscriptionsUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubscriptionsUpdateRequest) SetStatus(v string) {
	o.Status = &v
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given int32 and assigns it to the AmountMinor field.
func (o *SubscriptionsUpdateRequest) SetAmountMinor(v int32) {
	o.AmountMinor = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SubscriptionsUpdateRequest) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetNextBillingDate returns the NextBillingDate field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetNextBillingDate() string {
	if o == nil || IsNil(o.NextBillingDate) {
		var ret string
		return ret
	}
	return *o.NextBillingDate
}

// GetNextBillingDateOk returns a tuple with the NextBillingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetNextBillingDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextBillingDate) {
		return nil, false
	}
	return o.NextBillingDate, true
}

// HasNextBillingDate returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasNextBillingDate() bool {
	if o != nil && !IsNil(o.NextBillingDate) {
		return true
	}

	return false
}

// SetNextBillingDate gets a reference to the given string and assigns it to the NextBillingDate field.
func (o *SubscriptionsUpdateRequest) SetNextBillingDate(v string) {
	o.NextBillingDate = &v
}

// GetBillingAnchorDay returns the BillingAnchorDay field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetBillingAnchorDay() int32 {
	if o == nil || IsNil(o.BillingAnchorDay) {
		var ret int32
		return ret
	}
	return *o.BillingAnchorDay
}

// GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetBillingAnchorDayOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingAnchorDay) {
		return nil, false
	}
	return o.BillingAnchorDay, true
}

// HasBillingAnchorDay returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasBillingAnchorDay() bool {
	if o != nil && !IsNil(o.BillingAnchorDay) {
		return true
	}

	return false
}

// SetBillingAnchorDay gets a reference to the given int32 and assigns it to the BillingAnchorDay field.
func (o *SubscriptionsUpdateRequest) SetBillingAnchorDay(v int32) {
	o.BillingAnchorDay = &v
}

// GetPauseAtPeriodEnd returns the PauseAtPeriodEnd field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetPauseAtPeriodEnd() bool {
	if o == nil || IsNil(o.PauseAtPeriodEnd) {
		var ret bool
		return ret
	}
	return *o.PauseAtPeriodEnd
}

// GetPauseAtPeriodEndOk returns a tuple with the PauseAtPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetPauseAtPeriodEndOk() (*bool, bool) {
	if o == nil || IsNil(o.PauseAtPeriodEnd) {
		return nil, false
	}
	return o.PauseAtPeriodEnd, true
}

// HasPauseAtPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasPauseAtPeriodEnd() bool {
	if o != nil && !IsNil(o.PauseAtPeriodEnd) {
		return true
	}

	return false
}

// SetPauseAtPeriodEnd gets a reference to the given bool and assigns it to the PauseAtPeriodEnd field.
func (o *SubscriptionsUpdateRequest) SetPauseAtPeriodEnd(v bool) {
	o.PauseAtPeriodEnd = &v
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetCancelAtPeriodEnd() bool {
	if o == nil || IsNil(o.CancelAtPeriodEnd) {
		var ret bool
		return ret
	}
	return *o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetCancelAtPeriodEndOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelAtPeriodEnd) {
		return nil, false
	}
	return o.CancelAtPeriodEnd, true
}

// HasCancelAtPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasCancelAtPeriodEnd() bool {
	if o != nil && !IsNil(o.CancelAtPeriodEnd) {
		return true
	}

	return false
}

// SetCancelAtPeriodEnd gets a reference to the given bool and assigns it to the CancelAtPeriodEnd field.
func (o *SubscriptionsUpdateRequest) SetCancelAtPeriodEnd(v bool) {
	o.CancelAtPeriodEnd = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubscriptionsUpdateRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsUpdateRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubscriptionsUpdateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SubscriptionsUpdateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SubscriptionsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AmountMinor) {
		toSerialize["amountMinor"] = o.AmountMinor
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.NextBillingDate) {
		toSerialize["nextBillingDate"] = o.NextBillingDate
	}
	if !IsNil(o.BillingAnchorDay) {
		toSerialize["billingAnchorDay"] = o.BillingAnchorDay
	}
	if !IsNil(o.PauseAtPeriodEnd) {
		toSerialize["pauseAtPeriodEnd"] = o.PauseAtPeriodEnd
	}
	if !IsNil(o.CancelAtPeriodEnd) {
		toSerialize["cancelAtPeriodEnd"] = o.CancelAtPeriodEnd
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionsUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varSubscriptionsUpdateRequest := _SubscriptionsUpdateRequest{}

	err = json.Unmarshal(data, &varSubscriptionsUpdateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionsUpdateRequest(varSubscriptionsUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "nextBillingDate")
		delete(additionalProperties, "billingAnchorDay")
		delete(additionalProperties, "pauseAtPeriodEnd")
		delete(additionalProperties, "cancelAtPeriodEnd")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionsUpdateRequest struct {
	value *SubscriptionsUpdateRequest
	isSet bool
}

func (v NullableSubscriptionsUpdateRequest) Get() *SubscriptionsUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionsUpdateRequest) Set(val *SubscriptionsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsUpdateRequest(val *SubscriptionsUpdateRequest) *NullableSubscriptionsUpdateRequest {
	return &NullableSubscriptionsUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


