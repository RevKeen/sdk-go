/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the CheckoutSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckoutSession{}

// CheckoutSession A checkout session represents a customer's intent to pay.
type CheckoutSession struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Status string `json:"status"`
	Mode NullableString `json:"mode"`
	AmountTotal NullableFloat32 `json:"amount_total"`
	Currency NullableString `json:"currency"`
	CustomerId NullableString `json:"customer_id"`
	PaymentIntentId NullableString `json:"payment_intent_id"`
	InvoiceId NullableString `json:"invoice_id"`
	SubscriptionId NullableString `json:"subscription_id"`
	Url NullableString `json:"url"`
	SuccessUrl NullableString `json:"success_url"`
	CancelUrl NullableString `json:"cancel_url"`
	ExpiresAt NullableString `json:"expires_at"`
	CreatedAt NullableString `json:"created_at"`
	AllowedMethods []string `json:"allowed_methods"`
	SelectedMethod NullableString `json:"selected_method"`
	AdditionalProperties map[string]interface{}
}

type _CheckoutSession CheckoutSession

// NewCheckoutSession instantiates a new CheckoutSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSession(id string, object string, status string, mode NullableString, amountTotal NullableFloat32, currency NullableString, customerId NullableString, paymentIntentId NullableString, invoiceId NullableString, subscriptionId NullableString, url NullableString, successUrl NullableString, cancelUrl NullableString, expiresAt NullableString, createdAt NullableString, allowedMethods []string, selectedMethod NullableString) *CheckoutSession {
	this := CheckoutSession{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.Mode = mode
	this.AmountTotal = amountTotal
	this.Currency = currency
	this.CustomerId = customerId
	this.PaymentIntentId = paymentIntentId
	this.InvoiceId = invoiceId
	this.SubscriptionId = subscriptionId
	this.Url = url
	this.SuccessUrl = successUrl
	this.CancelUrl = cancelUrl
	this.ExpiresAt = expiresAt
	this.CreatedAt = createdAt
	this.AllowedMethods = allowedMethods
	this.SelectedMethod = selectedMethod
	return &this
}

// NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionWithDefaults() *CheckoutSession {
	this := CheckoutSession{}
	return &this
}

// GetId returns the Id field value
func (o *CheckoutSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CheckoutSession) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *CheckoutSession) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *CheckoutSession) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *CheckoutSession) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CheckoutSession) SetStatus(v string) {
	o.Status = v
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// SetMode sets field value
func (o *CheckoutSession) SetMode(v string) {
	o.Mode.Set(&v)
}

// GetAmountTotal returns the AmountTotal field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *CheckoutSession) GetAmountTotal() float32 {
	if o == nil || o.AmountTotal.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AmountTotal.Get()
}

// GetAmountTotalOk returns a tuple with the AmountTotal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetAmountTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountTotal.Get(), o.AmountTotal.IsSet()
}

// SetAmountTotal sets field value
func (o *CheckoutSession) SetAmountTotal(v float32) {
	o.AmountTotal.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *CheckoutSession) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *CheckoutSession) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetPaymentIntentId returns the PaymentIntentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetPaymentIntentId() string {
	if o == nil || o.PaymentIntentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentIntentId.Get()
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentIntentId.Get(), o.PaymentIntentId.IsSet()
}

// SetPaymentIntentId sets field value
func (o *CheckoutSession) SetPaymentIntentId(v string) {
	o.PaymentIntentId.Set(&v)
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *CheckoutSession) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *CheckoutSession) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *CheckoutSession) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetSuccessUrl returns the SuccessUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetSuccessUrl() string {
	if o == nil || o.SuccessUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.SuccessUrl.Get()
}

// GetSuccessUrlOk returns a tuple with the SuccessUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetSuccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessUrl.Get(), o.SuccessUrl.IsSet()
}

// SetSuccessUrl sets field value
func (o *CheckoutSession) SetSuccessUrl(v string) {
	o.SuccessUrl.Set(&v)
}

// GetCancelUrl returns the CancelUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetCancelUrl() string {
	if o == nil || o.CancelUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CancelUrl.Get()
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetCancelUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelUrl.Get(), o.CancelUrl.IsSet()
}

// SetCancelUrl sets field value
func (o *CheckoutSession) SetCancelUrl(v string) {
	o.CancelUrl.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *CheckoutSession) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *CheckoutSession) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetAllowedMethods returns the AllowedMethods field value
func (o *CheckoutSession) GetAllowedMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value
// and a boolean to check if the value has been set.
func (o *CheckoutSession) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedMethods, true
}

// SetAllowedMethods sets field value
func (o *CheckoutSession) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetSelectedMethod returns the SelectedMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckoutSession) GetSelectedMethod() string {
	if o == nil || o.SelectedMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.SelectedMethod.Get()
}

// GetSelectedMethodOk returns a tuple with the SelectedMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckoutSession) GetSelectedMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedMethod.Get(), o.SelectedMethod.IsSet()
}

// SetSelectedMethod sets field value
func (o *CheckoutSession) SetSelectedMethod(v string) {
	o.SelectedMethod.Set(&v)
}

func (o CheckoutSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckoutSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["mode"] = o.Mode.Get()
	toSerialize["amount_total"] = o.AmountTotal.Get()
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["payment_intent_id"] = o.PaymentIntentId.Get()
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	toSerialize["url"] = o.Url.Get()
	toSerialize["success_url"] = o.SuccessUrl.Get()
	toSerialize["cancel_url"] = o.CancelUrl.Get()
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["allowed_methods"] = o.AllowedMethods
	toSerialize["selected_method"] = o.SelectedMethod.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CheckoutSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"mode",
		"amount_total",
		"currency",
		"customer_id",
		"payment_intent_id",
		"invoice_id",
		"subscription_id",
		"url",
		"success_url",
		"cancel_url",
		"expires_at",
		"created_at",
		"allowed_methods",
		"selected_method",
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

	varCheckoutSession := _CheckoutSession{}

	err = json.Unmarshal(data, &varCheckoutSession)

	if err != nil {
		return err
	}

	*o = CheckoutSession(varCheckoutSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "status")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "amount_total")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "payment_intent_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "success_url")
		delete(additionalProperties, "cancel_url")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "allowed_methods")
		delete(additionalProperties, "selected_method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCheckoutSession struct {
	value *CheckoutSession
	isSet bool
}

func (v NullableCheckoutSession) Get() *CheckoutSession {
	return v.value
}

func (v *NullableCheckoutSession) Set(val *CheckoutSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSession(val *CheckoutSession) *NullableCheckoutSession {
	return &NullableCheckoutSession{value: val, isSet: true}
}

func (v NullableCheckoutSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


