/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-01-15` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-01-15
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the Void type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Void{}

// Void A void cancels an unsettled payment before gateway settlement. Voids are always for the full amount — use a refund for settled payments.
type Void struct {
	Id string `json:"id"`
	PublicId string `json:"public_id"`
	// The original payment ID this void is for (alias for parent_transaction_id)
	PaymentId string `json:"payment_id"`
	// Parent transaction ID in unified transaction model.
	ParentTransactionId string `json:"parent_transaction_id"`
	OrderId NullableString `json:"order_id"`
	Gateway string `json:"gateway"`
	GatewayVoidId NullableString `json:"gateway_void_id"`
	// Void amount in cents (always full amount of original transaction)
	AmountMinor float32 `json:"amount_minor"`
	Currency string `json:"currency"`
	Reason NullableString `json:"reason"`
	ReasonCode NullableString `json:"reason_code"`
	Status string `json:"status"`
	FailureReason NullableString `json:"failure_reason"`
	FailureCode NullableString `json:"failure_code"`
	VoidedAt NullableString `json:"voided_at"`
	CreatedAt string `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _Void Void

// NewVoid instantiates a new Void object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoid(id string, publicId string, paymentId string, parentTransactionId string, orderId NullableString, gateway string, gatewayVoidId NullableString, amountMinor float32, currency string, reason NullableString, reasonCode NullableString, status string, failureReason NullableString, failureCode NullableString, voidedAt NullableString, createdAt string) *Void {
	this := Void{}
	this.Id = id
	this.PublicId = publicId
	this.PaymentId = paymentId
	this.ParentTransactionId = parentTransactionId
	this.OrderId = orderId
	this.Gateway = gateway
	this.GatewayVoidId = gatewayVoidId
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.Reason = reason
	this.ReasonCode = reasonCode
	this.Status = status
	this.FailureReason = failureReason
	this.FailureCode = failureCode
	this.VoidedAt = voidedAt
	this.CreatedAt = createdAt
	return &this
}

// NewVoidWithDefaults instantiates a new Void object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidWithDefaults() *Void {
	this := Void{}
	return &this
}

// GetId returns the Id field value
func (o *Void) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Void) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Void) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
func (o *Void) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *Void) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *Void) SetPublicId(v string) {
	o.PublicId = v
}

// GetPaymentId returns the PaymentId field value
func (o *Void) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *Void) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *Void) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetParentTransactionId returns the ParentTransactionId field value
func (o *Void) GetParentTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentTransactionId
}

// GetParentTransactionIdOk returns a tuple with the ParentTransactionId field value
// and a boolean to check if the value has been set.
func (o *Void) GetParentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTransactionId, true
}

// SetParentTransactionId sets field value
func (o *Void) SetParentTransactionId(v string) {
	o.ParentTransactionId = v
}

// GetOrderId returns the OrderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetOrderId() string {
	if o == nil || o.OrderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrderId.Get()
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderId.Get(), o.OrderId.IsSet()
}

// SetOrderId sets field value
func (o *Void) SetOrderId(v string) {
	o.OrderId.Set(&v)
}

// GetGateway returns the Gateway field value
func (o *Void) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *Void) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *Void) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayVoidId returns the GatewayVoidId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetGatewayVoidId() string {
	if o == nil || o.GatewayVoidId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayVoidId.Get()
}

// GetGatewayVoidIdOk returns a tuple with the GatewayVoidId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetGatewayVoidIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayVoidId.Get(), o.GatewayVoidId.IsSet()
}

// SetGatewayVoidId sets field value
func (o *Void) SetGatewayVoidId(v string) {
	o.GatewayVoidId.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *Void) GetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *Void) GetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *Void) SetAmountMinor(v float32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *Void) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Void) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Void) SetCurrency(v string) {
	o.Currency = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *Void) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetReasonCode returns the ReasonCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetReasonCode() string {
	if o == nil || o.ReasonCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonCode.Get()
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonCode.Get(), o.ReasonCode.IsSet()
}

// SetReasonCode sets field value
func (o *Void) SetReasonCode(v string) {
	o.ReasonCode.Set(&v)
}

// GetStatus returns the Status field value
func (o *Void) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Void) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Void) SetStatus(v string) {
	o.Status = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *Void) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}

// GetFailureCode returns the FailureCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetFailureCode() string {
	if o == nil || o.FailureCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// SetFailureCode sets field value
func (o *Void) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}

// GetVoidedAt returns the VoidedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Void) GetVoidedAt() string {
	if o == nil || o.VoidedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.VoidedAt.Get()
}

// GetVoidedAtOk returns a tuple with the VoidedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Void) GetVoidedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VoidedAt.Get(), o.VoidedAt.IsSet()
}

// SetVoidedAt sets field value
func (o *Void) SetVoidedAt(v string) {
	o.VoidedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Void) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Void) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Void) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o Void) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Void) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["parent_transaction_id"] = o.ParentTransactionId
	toSerialize["order_id"] = o.OrderId.Get()
	toSerialize["gateway"] = o.Gateway
	toSerialize["gateway_void_id"] = o.GatewayVoidId.Get()
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["reason_code"] = o.ReasonCode.Get()
	toSerialize["status"] = o.Status
	toSerialize["failure_reason"] = o.FailureReason.Get()
	toSerialize["failure_code"] = o.FailureCode.Get()
	toSerialize["voided_at"] = o.VoidedAt.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Void) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"public_id",
		"payment_id",
		"parent_transaction_id",
		"order_id",
		"gateway",
		"gateway_void_id",
		"amount_minor",
		"currency",
		"reason",
		"reason_code",
		"status",
		"failure_reason",
		"failure_code",
		"voided_at",
		"created_at",
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

	varVoid := _Void{}

	err = json.Unmarshal(data, &varVoid)

	if err != nil {
		return err
	}

	*o = Void(varVoid)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "parent_transaction_id")
		delete(additionalProperties, "order_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_void_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "reason_code")
		delete(additionalProperties, "status")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "failure_code")
		delete(additionalProperties, "voided_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVoid struct {
	value *Void
	isSet bool
}

func (v NullableVoid) Get() *Void {
	return v.value
}

func (v *NullableVoid) Set(val *Void) {
	v.value = val
	v.isSet = true
}

func (v NullableVoid) IsSet() bool {
	return v.isSet
}

func (v *NullableVoid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoid(val *Void) *NullableVoid {
	return &NullableVoid{value: val, isSet: true}
}

func (v NullableVoid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


