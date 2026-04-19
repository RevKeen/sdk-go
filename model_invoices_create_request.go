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

// checks if the InvoicesCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicesCreateRequest{}

// InvoicesCreateRequest struct for InvoicesCreateRequest
type InvoicesCreateRequest struct {
	CustomerUuid string `json:"customer_uuid"`
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	TotalMinor int32 `json:"total_minor"`
	Currency *string `json:"currency,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Status *string `json:"status,omitempty"`
	// Key-value pairs for custom fields
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	SubscriptionTerms *SubscriptionTerms `json:"subscription_terms,omitempty"`
	// Restrict checkout payment methods for this invoice. When set, narrows available rails for checkout sessions created against this invoice. Omit to use merchant default.
	AllowedMethods []string `json:"allowed_methods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvoicesCreateRequest InvoicesCreateRequest

// NewInvoicesCreateRequest instantiates a new InvoicesCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesCreateRequest(customerUuid string, totalMinor int32) *InvoicesCreateRequest {
	this := InvoicesCreateRequest{}
	this.CustomerUuid = customerUuid
	this.TotalMinor = totalMinor
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// NewInvoicesCreateRequestWithDefaults instantiates a new InvoicesCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesCreateRequestWithDefaults() *InvoicesCreateRequest {
	this := InvoicesCreateRequest{}
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// GetCustomerUuid returns the CustomerUuid field value
func (o *InvoicesCreateRequest) GetCustomerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerUuid
}

// GetCustomerUuidOk returns a tuple with the CustomerUuid field value
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetCustomerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerUuid, true
}

// SetCustomerUuid sets field value
func (o *InvoicesCreateRequest) SetCustomerUuid(v string) {
	o.CustomerUuid = v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *InvoicesCreateRequest) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetTotalMinor returns the TotalMinor field value
func (o *InvoicesCreateRequest) GetTotalMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMinor
}

// GetTotalMinorOk returns a tuple with the TotalMinor field value
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetTotalMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinor, true
}

// SetTotalMinor sets field value
func (o *InvoicesCreateRequest) SetTotalMinor(v int32) {
	o.TotalMinor = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *InvoicesCreateRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *InvoicesCreateRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InvoicesCreateRequest) SetStatus(v string) {
	o.Status = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *InvoicesCreateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetSubscriptionTerms returns the SubscriptionTerms field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetSubscriptionTerms() SubscriptionTerms {
	if o == nil || IsNil(o.SubscriptionTerms) {
		var ret SubscriptionTerms
		return ret
	}
	return *o.SubscriptionTerms
}

// GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetSubscriptionTermsOk() (*SubscriptionTerms, bool) {
	if o == nil || IsNil(o.SubscriptionTerms) {
		return nil, false
	}
	return o.SubscriptionTerms, true
}

// HasSubscriptionTerms returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasSubscriptionTerms() bool {
	if o != nil && !IsNil(o.SubscriptionTerms) {
		return true
	}

	return false
}

// SetSubscriptionTerms gets a reference to the given SubscriptionTerms and assigns it to the SubscriptionTerms field.
func (o *InvoicesCreateRequest) SetSubscriptionTerms(v SubscriptionTerms) {
	o.SubscriptionTerms = &v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *InvoicesCreateRequest) GetAllowedMethods() []string {
	if o == nil || IsNil(o.AllowedMethods) {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicesCreateRequest) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMethods) {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *InvoicesCreateRequest) HasAllowedMethods() bool {
	if o != nil && !IsNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *InvoicesCreateRequest) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

func (o InvoicesCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicesCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer_uuid"] = o.CustomerUuid
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoice_number"] = o.InvoiceNumber
	}
	toSerialize["total_minor"] = o.TotalMinor
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.SubscriptionTerms) {
		toSerialize["subscription_terms"] = o.SubscriptionTerms
	}
	if !IsNil(o.AllowedMethods) {
		toSerialize["allowed_methods"] = o.AllowedMethods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoicesCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer_uuid",
		"total_minor",
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

	varInvoicesCreateRequest := _InvoicesCreateRequest{}

	err = json.Unmarshal(data, &varInvoicesCreateRequest)

	if err != nil {
		return err
	}

	*o = InvoicesCreateRequest(varInvoicesCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_uuid")
		delete(additionalProperties, "invoice_number")
		delete(additionalProperties, "total_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "status")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "subscription_terms")
		delete(additionalProperties, "allowed_methods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoicesCreateRequest struct {
	value *InvoicesCreateRequest
	isSet bool
}

func (v NullableInvoicesCreateRequest) Get() *InvoicesCreateRequest {
	return v.value
}

func (v *NullableInvoicesCreateRequest) Set(val *InvoicesCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesCreateRequest(val *InvoicesCreateRequest) *NullableInvoicesCreateRequest {
	return &NullableInvoicesCreateRequest{value: val, isSet: true}
}

func (v NullableInvoicesCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


