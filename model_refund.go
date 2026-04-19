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

// checks if the Refund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Refund{}

// Refund A refund represents money returned to a customer after a completed payment. Refunds create a child transaction linked to the original payment for audit trail purposes.
type Refund struct {
	Id string `json:"id"`
	PublicId string `json:"public_id"`
	// The original payment ID this refund is for (alias for parent_transaction_id)
	PaymentId string `json:"payment_id"`
	// Parent transaction ID in unified transaction model. Same as payment_id for refunds.
	ParentTransactionId string `json:"parent_transaction_id"`
	OrderId NullableString `json:"order_id"`
	// Name of the payment processor that handled this refund
	Gateway string `json:"gateway"`
	GatewayRefundId NullableString `json:"gateway_refund_id"`
	// Refund amount in cents
	AmountMinor float32 `json:"amount_minor"`
	Currency string `json:"currency"`
	Reason NullableString `json:"reason"`
	ReasonDetails NullableString `json:"reason_details"`
	Status string `json:"status"`
	FailureReason NullableString `json:"failure_reason"`
	FailureCode NullableString `json:"failure_code"`
	ProcessedAt NullableString `json:"processed_at"`
	CreatedAt string `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _Refund Refund

// NewRefund instantiates a new Refund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefund(id string, publicId string, paymentId string, parentTransactionId string, orderId NullableString, gateway string, gatewayRefundId NullableString, amountMinor float32, currency string, reason NullableString, reasonDetails NullableString, status string, failureReason NullableString, failureCode NullableString, processedAt NullableString, createdAt string) *Refund {
	this := Refund{}
	this.Id = id
	this.PublicId = publicId
	this.PaymentId = paymentId
	this.ParentTransactionId = parentTransactionId
	this.OrderId = orderId
	this.Gateway = gateway
	this.GatewayRefundId = gatewayRefundId
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.Reason = reason
	this.ReasonDetails = reasonDetails
	this.Status = status
	this.FailureReason = failureReason
	this.FailureCode = failureCode
	this.ProcessedAt = processedAt
	this.CreatedAt = createdAt
	return &this
}

// NewRefundWithDefaults instantiates a new Refund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefundWithDefaults() *Refund {
	this := Refund{}
	return &this
}

// GetId returns the Id field value
func (o *Refund) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Refund) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Refund) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
func (o *Refund) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *Refund) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *Refund) SetPublicId(v string) {
	o.PublicId = v
}

// GetPaymentId returns the PaymentId field value
func (o *Refund) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *Refund) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *Refund) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetParentTransactionId returns the ParentTransactionId field value
func (o *Refund) GetParentTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentTransactionId
}

// GetParentTransactionIdOk returns a tuple with the ParentTransactionId field value
// and a boolean to check if the value has been set.
func (o *Refund) GetParentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTransactionId, true
}

// SetParentTransactionId sets field value
func (o *Refund) SetParentTransactionId(v string) {
	o.ParentTransactionId = v
}

// GetOrderId returns the OrderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetOrderId() string {
	if o == nil || o.OrderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OrderId.Get()
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderId.Get(), o.OrderId.IsSet()
}

// SetOrderId sets field value
func (o *Refund) SetOrderId(v string) {
	o.OrderId.Set(&v)
}

// GetGateway returns the Gateway field value
func (o *Refund) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *Refund) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *Refund) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayRefundId returns the GatewayRefundId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetGatewayRefundId() string {
	if o == nil || o.GatewayRefundId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayRefundId.Get()
}

// GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetGatewayRefundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayRefundId.Get(), o.GatewayRefundId.IsSet()
}

// SetGatewayRefundId sets field value
func (o *Refund) SetGatewayRefundId(v string) {
	o.GatewayRefundId.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *Refund) GetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *Refund) GetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *Refund) SetAmountMinor(v float32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *Refund) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Refund) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Refund) SetCurrency(v string) {
	o.Currency = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *Refund) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetReasonDetails returns the ReasonDetails field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetReasonDetails() string {
	if o == nil || o.ReasonDetails.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReasonDetails.Get()
}

// GetReasonDetailsOk returns a tuple with the ReasonDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetReasonDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReasonDetails.Get(), o.ReasonDetails.IsSet()
}

// SetReasonDetails sets field value
func (o *Refund) SetReasonDetails(v string) {
	o.ReasonDetails.Set(&v)
}

// GetStatus returns the Status field value
func (o *Refund) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Refund) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Refund) SetStatus(v string) {
	o.Status = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *Refund) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}

// GetFailureCode returns the FailureCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetFailureCode() string {
	if o == nil || o.FailureCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// SetFailureCode sets field value
func (o *Refund) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}

// GetProcessedAt returns the ProcessedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Refund) GetProcessedAt() string {
	if o == nil || o.ProcessedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProcessedAt.Get()
}

// GetProcessedAtOk returns a tuple with the ProcessedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Refund) GetProcessedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessedAt.Get(), o.ProcessedAt.IsSet()
}

// SetProcessedAt sets field value
func (o *Refund) SetProcessedAt(v string) {
	o.ProcessedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Refund) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Refund) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Refund) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Refund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["parent_transaction_id"] = o.ParentTransactionId
	toSerialize["order_id"] = o.OrderId.Get()
	toSerialize["gateway"] = o.Gateway
	toSerialize["gateway_refund_id"] = o.GatewayRefundId.Get()
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["reason_details"] = o.ReasonDetails.Get()
	toSerialize["status"] = o.Status
	toSerialize["failure_reason"] = o.FailureReason.Get()
	toSerialize["failure_code"] = o.FailureCode.Get()
	toSerialize["processed_at"] = o.ProcessedAt.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Refund) UnmarshalJSON(data []byte) (err error) {
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
		"gateway_refund_id",
		"amount_minor",
		"currency",
		"reason",
		"reason_details",
		"status",
		"failure_reason",
		"failure_code",
		"processed_at",
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

	varRefund := _Refund{}

	err = json.Unmarshal(data, &varRefund)

	if err != nil {
		return err
	}

	*o = Refund(varRefund)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "parent_transaction_id")
		delete(additionalProperties, "order_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_refund_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "reason_details")
		delete(additionalProperties, "status")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "failure_code")
		delete(additionalProperties, "processed_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefund struct {
	value *Refund
	isSet bool
}

func (v NullableRefund) Get() *Refund {
	return v.value
}

func (v *NullableRefund) Set(val *Refund) {
	v.value = val
	v.isSet = true
}

func (v NullableRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefund(val *Refund) *NullableRefund {
	return &NullableRefund{value: val, isSet: true}
}

func (v NullableRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


