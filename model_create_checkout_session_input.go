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

// checks if the CreateCheckoutSessionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCheckoutSessionInput{}

// CreateCheckoutSessionInput Parameters for creating a hosted checkout session. Customers are redirected to complete payment on a RevKeen-hosted page.
type CreateCheckoutSessionInput struct {
	InvoiceId *string `json:"invoiceId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	AmountMinor *int32 `json:"amountMinor,omitempty"`
	Currency *string `json:"currency,omitempty"`
	SuccessUrl *string `json:"successUrl,omitempty"`
	CancelUrl *string `json:"cancelUrl,omitempty"`
	// Payment methods to offer. Intersected with merchant capabilities. Defaults to merchant config.
	AllowedMethods []string `json:"allowedMethods,omitempty"`
	// Target a registered companion device. Session is pushed via SSE to the device.
	CompanionDeviceId *string `json:"companionDeviceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateCheckoutSessionInput CreateCheckoutSessionInput

// NewCreateCheckoutSessionInput instantiates a new CreateCheckoutSessionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCheckoutSessionInput() *CreateCheckoutSessionInput {
	this := CreateCheckoutSessionInput{}
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// NewCreateCheckoutSessionInputWithDefaults instantiates a new CreateCheckoutSessionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCheckoutSessionInputWithDefaults() *CreateCheckoutSessionInput {
	this := CreateCheckoutSessionInput{}
	var currency string = "USD"
	this.Currency = &currency
	return &this
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *CreateCheckoutSessionInput) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *CreateCheckoutSessionInput) SetProductId(v string) {
	o.ProductId = &v
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetAmountMinor() int32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret int32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given int32 and assigns it to the AmountMinor field.
func (o *CreateCheckoutSessionInput) SetAmountMinor(v int32) {
	o.AmountMinor = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateCheckoutSessionInput) SetCurrency(v string) {
	o.Currency = &v
}

// GetSuccessUrl returns the SuccessUrl field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetSuccessUrl() string {
	if o == nil || IsNil(o.SuccessUrl) {
		var ret string
		return ret
	}
	return *o.SuccessUrl
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetSuccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SuccessUrl) {
		return nil, false
	}
	return o.SuccessUrl, true
}

// HasSuccessUrl returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasSuccessUrl() bool {
	if o != nil && !IsNil(o.SuccessUrl) {
		return true
	}

	return false
}

// SetSuccessUrl gets a reference to the given string and assigns it to the SuccessUrl field.
func (o *CreateCheckoutSessionInput) SetSuccessUrl(v string) {
	o.SuccessUrl = &v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *CreateCheckoutSessionInput) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetAllowedMethods() []string {
	if o == nil || IsNil(o.AllowedMethods) {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMethods) {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasAllowedMethods() bool {
	if o != nil && !IsNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *CreateCheckoutSessionInput) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetCompanionDeviceId returns the CompanionDeviceId field value if set, zero value otherwise.
func (o *CreateCheckoutSessionInput) GetCompanionDeviceId() string {
	if o == nil || IsNil(o.CompanionDeviceId) {
		var ret string
		return ret
	}
	return *o.CompanionDeviceId
}

// GetCompanionDeviceIdOk returns a tuple with the CompanionDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCheckoutSessionInput) GetCompanionDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanionDeviceId) {
		return nil, false
	}
	return o.CompanionDeviceId, true
}

// HasCompanionDeviceId returns a boolean if a field has been set.
func (o *CreateCheckoutSessionInput) HasCompanionDeviceId() bool {
	if o != nil && !IsNil(o.CompanionDeviceId) {
		return true
	}

	return false
}

// SetCompanionDeviceId gets a reference to the given string and assigns it to the CompanionDeviceId field.
func (o *CreateCheckoutSessionInput) SetCompanionDeviceId(v string) {
	o.CompanionDeviceId = &v
}

func (o CreateCheckoutSessionInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCheckoutSessionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.AmountMinor) {
		toSerialize["amountMinor"] = o.AmountMinor
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.SuccessUrl) {
		toSerialize["successUrl"] = o.SuccessUrl
	}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if !IsNil(o.AllowedMethods) {
		toSerialize["allowedMethods"] = o.AllowedMethods
	}
	if !IsNil(o.CompanionDeviceId) {
		toSerialize["companionDeviceId"] = o.CompanionDeviceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCheckoutSessionInput) UnmarshalJSON(data []byte) (err error) {
	varCreateCheckoutSessionInput := _CreateCheckoutSessionInput{}

	err = json.Unmarshal(data, &varCreateCheckoutSessionInput)

	if err != nil {
		return err
	}

	*o = CreateCheckoutSessionInput(varCreateCheckoutSessionInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invoiceId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "successUrl")
		delete(additionalProperties, "cancelUrl")
		delete(additionalProperties, "allowedMethods")
		delete(additionalProperties, "companionDeviceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCheckoutSessionInput struct {
	value *CreateCheckoutSessionInput
	isSet bool
}

func (v NullableCreateCheckoutSessionInput) Get() *CreateCheckoutSessionInput {
	return v.value
}

func (v *NullableCreateCheckoutSessionInput) Set(val *CreateCheckoutSessionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCheckoutSessionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCheckoutSessionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCheckoutSessionInput(val *CreateCheckoutSessionInput) *NullableCreateCheckoutSessionInput {
	return &NullableCreateCheckoutSessionInput{value: val, isSet: true}
}

func (v NullableCreateCheckoutSessionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCheckoutSessionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


