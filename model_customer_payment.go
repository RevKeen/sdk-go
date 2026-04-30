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

// checks if the CustomerPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerPayment{}

// CustomerPayment A payment transaction belonging to a specific customer, with amount, status, and gateway details.
type CustomerPayment struct {
	Id string `json:"id"`
	CustomerId NullableString `json:"customer_id"`
	AmountMinor int32 `json:"amount_minor"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	PaymentMethodType NullableString `json:"payment_method_type"`
	GatewayTransactionId NullableString `json:"gateway_transaction_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _CustomerPayment CustomerPayment

// NewCustomerPayment instantiates a new CustomerPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPayment(id string, customerId NullableString, amountMinor int32, currency string, status string, paymentMethodType NullableString, gatewayTransactionId NullableString, createdAt string, updatedAt string) *CustomerPayment {
	this := CustomerPayment{}
	this.Id = id
	this.CustomerId = customerId
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.Status = status
	this.PaymentMethodType = paymentMethodType
	this.GatewayTransactionId = gatewayTransactionId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomerPaymentWithDefaults instantiates a new CustomerPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentWithDefaults() *CustomerPayment {
	this := CustomerPayment{}
	return &this
}

// GetId returns the Id field value
func (o *CustomerPayment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerPayment) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerPayment) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPayment) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *CustomerPayment) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *CustomerPayment) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *CustomerPayment) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *CustomerPayment) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CustomerPayment) SetCurrency(v string) {
	o.Currency = v
}

// GetStatus returns the Status field value
func (o *CustomerPayment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CustomerPayment) SetStatus(v string) {
	o.Status = v
}

// GetPaymentMethodType returns the PaymentMethodType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerPayment) GetPaymentMethodType() string {
	if o == nil || o.PaymentMethodType.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethodType.Get()
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPayment) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodType.Get(), o.PaymentMethodType.IsSet()
}

// SetPaymentMethodType sets field value
func (o *CustomerPayment) SetPaymentMethodType(v string) {
	o.PaymentMethodType.Set(&v)
}

// GetGatewayTransactionId returns the GatewayTransactionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustomerPayment) GetGatewayTransactionId() string {
	if o == nil || o.GatewayTransactionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayTransactionId.Get()
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerPayment) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTransactionId.Get(), o.GatewayTransactionId.IsSet()
}

// SetGatewayTransactionId sets field value
func (o *CustomerPayment) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerPayment) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerPayment) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomerPayment) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerPayment) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomerPayment) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o CustomerPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["status"] = o.Status
	toSerialize["payment_method_type"] = o.PaymentMethodType.Get()
	toSerialize["gateway_transaction_id"] = o.GatewayTransactionId.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customer_id",
		"amount_minor",
		"currency",
		"status",
		"payment_method_type",
		"gateway_transaction_id",
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

	varCustomerPayment := _CustomerPayment{}

	err = json.Unmarshal(data, &varCustomerPayment)

	if err != nil {
		return err
	}

	*o = CustomerPayment(varCustomerPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "status")
		delete(additionalProperties, "payment_method_type")
		delete(additionalProperties, "gateway_transaction_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerPayment struct {
	value *CustomerPayment
	isSet bool
}

func (v NullableCustomerPayment) Get() *CustomerPayment {
	return v.value
}

func (v *NullableCustomerPayment) Set(val *CustomerPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPayment(val *CustomerPayment) *NullableCustomerPayment {
	return &NullableCustomerPayment{value: val, isSet: true}
}

func (v NullableCustomerPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


