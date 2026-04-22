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
)

// checks if the UpdatePaymentIntentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePaymentIntentRequest{}

// UpdatePaymentIntentRequest Parameters for updating a payment intent before confirmation. Only amount, description, payment method, and metadata can be changed.
type UpdatePaymentIntentRequest struct {
	// Amount in cents
	Amount *int32 `json:"amount,omitempty"`
	// Three-letter ISO currency code
	Currency *string `json:"currency,omitempty"`
	// Payment method ID
	PaymentMethod *string `json:"payment_method,omitempty"`
	// Description for merchant reference
	Description *string `json:"description,omitempty"`
	// Statement descriptor
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Statement descriptor suffix
	StatementDescriptorSuffix *string `json:"statement_descriptor_suffix,omitempty"`
	// Email to send receipt to
	ReceiptEmail *string `json:"receipt_email,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdatePaymentIntentRequest UpdatePaymentIntentRequest

// NewUpdatePaymentIntentRequest instantiates a new UpdatePaymentIntentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePaymentIntentRequest() *UpdatePaymentIntentRequest {
	this := UpdatePaymentIntentRequest{}
	return &this
}

// NewUpdatePaymentIntentRequestWithDefaults instantiates a new UpdatePaymentIntentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePaymentIntentRequestWithDefaults() *UpdatePaymentIntentRequest {
	this := UpdatePaymentIntentRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *UpdatePaymentIntentRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UpdatePaymentIntentRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *UpdatePaymentIntentRequest) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePaymentIntentRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetStatementDescriptor() string {
	if o == nil || IsNil(o.StatementDescriptor) {
		var ret string
		return ret
	}
	return *o.StatementDescriptor
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetStatementDescriptorOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDescriptor) {
		return nil, false
	}
	return o.StatementDescriptor, true
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasStatementDescriptor() bool {
	if o != nil && !IsNil(o.StatementDescriptor) {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given string and assigns it to the StatementDescriptor field.
func (o *UpdatePaymentIntentRequest) SetStatementDescriptor(v string) {
	o.StatementDescriptor = &v
}

// GetStatementDescriptorSuffix returns the StatementDescriptorSuffix field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetStatementDescriptorSuffix() string {
	if o == nil || IsNil(o.StatementDescriptorSuffix) {
		var ret string
		return ret
	}
	return *o.StatementDescriptorSuffix
}

// GetStatementDescriptorSuffixOk returns a tuple with the StatementDescriptorSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetStatementDescriptorSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDescriptorSuffix) {
		return nil, false
	}
	return o.StatementDescriptorSuffix, true
}

// HasStatementDescriptorSuffix returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasStatementDescriptorSuffix() bool {
	if o != nil && !IsNil(o.StatementDescriptorSuffix) {
		return true
	}

	return false
}

// SetStatementDescriptorSuffix gets a reference to the given string and assigns it to the StatementDescriptorSuffix field.
func (o *UpdatePaymentIntentRequest) SetStatementDescriptorSuffix(v string) {
	o.StatementDescriptorSuffix = &v
}

// GetReceiptEmail returns the ReceiptEmail field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetReceiptEmail() string {
	if o == nil || IsNil(o.ReceiptEmail) {
		var ret string
		return ret
	}
	return *o.ReceiptEmail
}

// GetReceiptEmailOk returns a tuple with the ReceiptEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetReceiptEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptEmail) {
		return nil, false
	}
	return o.ReceiptEmail, true
}

// HasReceiptEmail returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasReceiptEmail() bool {
	if o != nil && !IsNil(o.ReceiptEmail) {
		return true
	}

	return false
}

// SetReceiptEmail gets a reference to the given string and assigns it to the ReceiptEmail field.
func (o *UpdatePaymentIntentRequest) SetReceiptEmail(v string) {
	o.ReceiptEmail = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdatePaymentIntentRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePaymentIntentRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdatePaymentIntentRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UpdatePaymentIntentRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o UpdatePaymentIntentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePaymentIntentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StatementDescriptor) {
		toSerialize["statement_descriptor"] = o.StatementDescriptor
	}
	if !IsNil(o.StatementDescriptorSuffix) {
		toSerialize["statement_descriptor_suffix"] = o.StatementDescriptorSuffix
	}
	if !IsNil(o.ReceiptEmail) {
		toSerialize["receipt_email"] = o.ReceiptEmail
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdatePaymentIntentRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdatePaymentIntentRequest := _UpdatePaymentIntentRequest{}

	err = json.Unmarshal(data, &varUpdatePaymentIntentRequest)

	if err != nil {
		return err
	}

	*o = UpdatePaymentIntentRequest(varUpdatePaymentIntentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "description")
		delete(additionalProperties, "statement_descriptor")
		delete(additionalProperties, "statement_descriptor_suffix")
		delete(additionalProperties, "receipt_email")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdatePaymentIntentRequest struct {
	value *UpdatePaymentIntentRequest
	isSet bool
}

func (v NullableUpdatePaymentIntentRequest) Get() *UpdatePaymentIntentRequest {
	return v.value
}

func (v *NullableUpdatePaymentIntentRequest) Set(val *UpdatePaymentIntentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePaymentIntentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePaymentIntentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePaymentIntentRequest(val *UpdatePaymentIntentRequest) *NullableUpdatePaymentIntentRequest {
	return &NullableUpdatePaymentIntentRequest{value: val, isSet: true}
}

func (v NullableUpdatePaymentIntentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePaymentIntentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


