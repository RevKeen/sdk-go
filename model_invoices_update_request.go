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

// checks if the InvoicesUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicesUpdateRequest{}

// InvoicesUpdateRequest struct for InvoicesUpdateRequest
type InvoicesUpdateRequest struct {
	Status *string `json:"status,omitempty"`
	TotalMinor *int32 `json:"total_minor,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	SubscriptionTerms *InvoicesUpdateRequestSubscriptionTerms `json:"subscription_terms,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvoicesUpdateRequest InvoicesUpdateRequest

// NewInvoicesUpdateRequest instantiates a new InvoicesUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesUpdateRequest() *InvoicesUpdateRequest {
	this := InvoicesUpdateRequest{}
	return &this
}

// NewInvoicesUpdateRequestWithDefaults instantiates a new InvoicesUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesUpdateRequestWithDefaults() *InvoicesUpdateRequest {
	this := InvoicesUpdateRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InvoicesUpdateRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InvoicesUpdateRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InvoicesUpdateRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTotalMinor returns the TotalMinor field value if set, zero value otherwise.
func (o *InvoicesUpdateRequest) GetTotalMinor() int32 {
	if o == nil || IsNil(o.TotalMinor) {
		var ret int32
		return ret
	}
	return *o.TotalMinor
}

// GetTotalMinorOk returns a tuple with the TotalMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesUpdateRequest) GetTotalMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalMinor) {
		return nil, false
	}
	return o.TotalMinor, true
}

// HasTotalMinor returns a boolean if a field has been set.
func (o *InvoicesUpdateRequest) HasTotalMinor() bool {
	if o != nil && !IsNil(o.TotalMinor) {
		return true
	}

	return false
}

// SetTotalMinor gets a reference to the given int32 and assigns it to the TotalMinor field.
func (o *InvoicesUpdateRequest) SetTotalMinor(v int32) {
	o.TotalMinor = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoicesUpdateRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesUpdateRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoicesUpdateRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InvoicesUpdateRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetSubscriptionTerms returns the SubscriptionTerms field value if set, zero value otherwise.
func (o *InvoicesUpdateRequest) GetSubscriptionTerms() InvoicesUpdateRequestSubscriptionTerms {
	if o == nil || IsNil(o.SubscriptionTerms) {
		var ret InvoicesUpdateRequestSubscriptionTerms
		return ret
	}
	return *o.SubscriptionTerms
}

// GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesUpdateRequest) GetSubscriptionTermsOk() (*InvoicesUpdateRequestSubscriptionTerms, bool) {
	if o == nil || IsNil(o.SubscriptionTerms) {
		return nil, false
	}
	return o.SubscriptionTerms, true
}

// HasSubscriptionTerms returns a boolean if a field has been set.
func (o *InvoicesUpdateRequest) HasSubscriptionTerms() bool {
	if o != nil && !IsNil(o.SubscriptionTerms) {
		return true
	}

	return false
}

// SetSubscriptionTerms gets a reference to the given InvoicesUpdateRequestSubscriptionTerms and assigns it to the SubscriptionTerms field.
func (o *InvoicesUpdateRequest) SetSubscriptionTerms(v InvoicesUpdateRequestSubscriptionTerms) {
	o.SubscriptionTerms = &v
}

func (o InvoicesUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicesUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalMinor) {
		toSerialize["total_minor"] = o.TotalMinor
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.SubscriptionTerms) {
		toSerialize["subscription_terms"] = o.SubscriptionTerms
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoicesUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	varInvoicesUpdateRequest := _InvoicesUpdateRequest{}

	err = json.Unmarshal(data, &varInvoicesUpdateRequest)

	if err != nil {
		return err
	}

	*o = InvoicesUpdateRequest(varInvoicesUpdateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "total_minor")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "subscription_terms")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoicesUpdateRequest struct {
	value *InvoicesUpdateRequest
	isSet bool
}

func (v NullableInvoicesUpdateRequest) Get() *InvoicesUpdateRequest {
	return v.value
}

func (v *NullableInvoicesUpdateRequest) Set(val *InvoicesUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesUpdateRequest(val *InvoicesUpdateRequest) *NullableInvoicesUpdateRequest {
	return &NullableInvoicesUpdateRequest{value: val, isSet: true}
}

func (v NullableInvoicesUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


