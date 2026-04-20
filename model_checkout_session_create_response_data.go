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

// checks if the CheckoutSessionCreateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSessionCreateResponseData{}

// CheckoutSessionCreateResponseData struct for CheckoutSessionCreateResponseData
type CheckoutSessionCreateResponseData struct {
	Id string `json:"id"`
	Url string `json:"url"`
	PublicToken string `json:"publicToken"`
	AmountMinor *float32 `json:"amountMinor,omitempty"`
	Currency string `json:"currency"`
	CustomerId *string `json:"customerId,omitempty"`
	ExpiresAt string `json:"expiresAt"`
	AllowedMethods []string `json:"allowedMethods,omitempty"`
	SelectedMethod NullableString `json:"selectedMethod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CheckoutSessionCreateResponseData CheckoutSessionCreateResponseData

// NewCheckoutSessionCreateResponseData instantiates a new CheckoutSessionCreateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionCreateResponseData(id string, url string, publicToken string, currency string, expiresAt string) *CheckoutSessionCreateResponseData {
	this := CheckoutSessionCreateResponseData{}
	this.Id = id
	this.Url = url
	this.PublicToken = publicToken
	this.Currency = currency
	this.ExpiresAt = expiresAt
	return &this
}

// NewCheckoutSessionCreateResponseDataWithDefaults instantiates a new CheckoutSessionCreateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionCreateResponseDataWithDefaults() *CheckoutSessionCreateResponseData {
	this := CheckoutSessionCreateResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *CheckoutSessionCreateResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutSessionCreateResponseData) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CheckoutSessionCreateResponseData) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CheckoutSessionCreateResponseData) SetUrl(v string) {
	o.Url = v
}

// GetPublicToken returns the PublicToken field value
func (o *CheckoutSessionCreateResponseData) GetPublicToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicToken
}

// GetPublicTokenOk returns a tuple with the PublicToken field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetPublicTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicToken, true
}

// SetPublicToken sets field value
func (o *CheckoutSessionCreateResponseData) SetPublicToken(v string) {
	o.PublicToken = v
}

// GetAmountMinor returns the AmountMinor field value if set, zero value otherwise.
func (o *CheckoutSessionCreateResponseData) GetAmountMinor() float32 {
	if o == nil || IsNil(o.AmountMinor) {
		var ret float32
		return ret
	}
	return *o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetAmountMinorOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountMinor) {
		return nil, false
	}
	return o.AmountMinor, true
}

// HasAmountMinor returns a boolean if a field has been set.
func (o *CheckoutSessionCreateResponseData) HasAmountMinor() bool {
	if o != nil && !IsNil(o.AmountMinor) {
		return true
	}

	return false
}

// SetAmountMinor gets a reference to the given float32 and assigns it to the AmountMinor field.
func (o *CheckoutSessionCreateResponseData) SetAmountMinor(v float32) {
	o.AmountMinor = &v
}

// GetCurrency returns the Currency field value
func (o *CheckoutSessionCreateResponseData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CheckoutSessionCreateResponseData) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CheckoutSessionCreateResponseData) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CheckoutSessionCreateResponseData) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CheckoutSessionCreateResponseData) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *CheckoutSessionCreateResponseData) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *CheckoutSessionCreateResponseData) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise.
func (o *CheckoutSessionCreateResponseData) GetAllowedMethods() []string {
	if o == nil || IsNil(o.AllowedMethods) {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionCreateResponseData) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMethods) {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *CheckoutSessionCreateResponseData) HasAllowedMethods() bool {
	if o != nil && !IsNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *CheckoutSessionCreateResponseData) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetSelectedMethod returns the SelectedMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckoutSessionCreateResponseData) GetSelectedMethod() string {
	if o == nil || IsNil(o.SelectedMethod.Get()) {
		var ret string
		return ret
	}
	return *o.SelectedMethod.Get()
}

// GetSelectedMethodOk returns a tuple with the SelectedMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSessionCreateResponseData) GetSelectedMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedMethod.Get(), o.SelectedMethod.IsSet()
}

// HasSelectedMethod returns a boolean if a field has been set.
func (o *CheckoutSessionCreateResponseData) HasSelectedMethod() bool {
	if o != nil && o.SelectedMethod.IsSet() {
		return true
	}

	return false
}

// SetSelectedMethod gets a reference to the given NullableString and assigns it to the SelectedMethod field.
func (o *CheckoutSessionCreateResponseData) SetSelectedMethod(v string) {
	o.SelectedMethod.Set(&v)
}
// SetSelectedMethodNil sets the value for SelectedMethod to be an explicit nil
func (o *CheckoutSessionCreateResponseData) SetSelectedMethodNil() {
	o.SelectedMethod.Set(nil)
}

// UnsetSelectedMethod ensures that no value is present for SelectedMethod, not even an explicit nil
func (o *CheckoutSessionCreateResponseData) UnsetSelectedMethod() {
	o.SelectedMethod.Unset()
}

func (o CheckoutSessionCreateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSessionCreateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["publicToken"] = o.PublicToken
	if !IsNil(o.AmountMinor) {
		toSerialize["amountMinor"] = o.AmountMinor
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	toSerialize["expiresAt"] = o.ExpiresAt
	if !IsNil(o.AllowedMethods) {
		toSerialize["allowedMethods"] = o.AllowedMethods
	}
	if o.SelectedMethod.IsSet() {
		toSerialize["selectedMethod"] = o.SelectedMethod.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckoutSessionCreateResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"publicToken",
		"currency",
		"expiresAt",
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

	varCheckoutSessionCreateResponseData := _CheckoutSessionCreateResponseData{}

	err = json.Unmarshal(data, &varCheckoutSessionCreateResponseData)

	if err != nil {
		return err
	}

	*o = CheckoutSessionCreateResponseData(varCheckoutSessionCreateResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "publicToken")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "allowedMethods")
		delete(additionalProperties, "selectedMethod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckoutSessionCreateResponseData struct {
	value *CheckoutSessionCreateResponseData
	isSet bool
}

func (v NullableCheckoutSessionCreateResponseData) Get() *CheckoutSessionCreateResponseData {
	return v.value
}

func (v *NullableCheckoutSessionCreateResponseData) Set(val *CheckoutSessionCreateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionCreateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionCreateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionCreateResponseData(val *CheckoutSessionCreateResponseData) *NullableCheckoutSessionCreateResponseData {
	return &NullableCheckoutSessionCreateResponseData{value: val, isSet: true}
}

func (v NullableCheckoutSessionCreateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionCreateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


