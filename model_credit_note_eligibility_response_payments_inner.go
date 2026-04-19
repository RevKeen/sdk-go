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

// checks if the CreditNoteEligibilityResponsePaymentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditNoteEligibilityResponsePaymentsInner{}

// CreditNoteEligibilityResponsePaymentsInner struct for CreditNoteEligibilityResponsePaymentsInner
type CreditNoteEligibilityResponsePaymentsInner struct {
	TransactionId string `json:"transaction_id"`
	Gateway string `json:"gateway"`
	PaymentMethod NullableString `json:"payment_method"`
	AmountMinor float32 `json:"amount_minor"`
	AmountRefundedMinor float32 `json:"amount_refunded_minor"`
	AmountAvailableMinor float32 `json:"amount_available_minor"`
	CardBrand NullableString `json:"card_brand"`
	CardLast4 NullableString `json:"card_last4"`
	TerminalSerial NullableString `json:"terminal_serial"`
	AvailableOperations []string `json:"available_operations"`
	Constraints CreditNoteEligibilityResponsePaymentsInnerConstraints `json:"constraints"`
	AdditionalProperties map[string]interface{}
}

type _CreditNoteEligibilityResponsePaymentsInner CreditNoteEligibilityResponsePaymentsInner

// NewCreditNoteEligibilityResponsePaymentsInner instantiates a new CreditNoteEligibilityResponsePaymentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditNoteEligibilityResponsePaymentsInner(transactionId string, gateway string, paymentMethod NullableString, amountMinor float32, amountRefundedMinor float32, amountAvailableMinor float32, cardBrand NullableString, cardLast4 NullableString, terminalSerial NullableString, availableOperations []string, constraints CreditNoteEligibilityResponsePaymentsInnerConstraints) *CreditNoteEligibilityResponsePaymentsInner {
	this := CreditNoteEligibilityResponsePaymentsInner{}
	this.TransactionId = transactionId
	this.Gateway = gateway
	this.PaymentMethod = paymentMethod
	this.AmountMinor = amountMinor
	this.AmountRefundedMinor = amountRefundedMinor
	this.AmountAvailableMinor = amountAvailableMinor
	this.CardBrand = cardBrand
	this.CardLast4 = cardLast4
	this.TerminalSerial = terminalSerial
	this.AvailableOperations = availableOperations
	this.Constraints = constraints
	return &this
}

// NewCreditNoteEligibilityResponsePaymentsInnerWithDefaults instantiates a new CreditNoteEligibilityResponsePaymentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditNoteEligibilityResponsePaymentsInnerWithDefaults() *CreditNoteEligibilityResponsePaymentsInner {
	this := CreditNoteEligibilityResponsePaymentsInner{}
	return &this
}

// GetTransactionId returns the TransactionId field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetGateway returns the Gateway field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetGateway(v string) {
	o.Gateway = v
}

// GetPaymentMethod returns the PaymentMethod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetPaymentMethod() string {
	if o == nil || o.PaymentMethod.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// SetPaymentMethod sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetPaymentMethod(v string) {
	o.PaymentMethod.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountMinor(v float32) {
	o.AmountMinor = v
}

// GetAmountRefundedMinor returns the AmountRefundedMinor field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountRefundedMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountRefundedMinor
}

// GetAmountRefundedMinorOk returns a tuple with the AmountRefundedMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountRefundedMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountRefundedMinor, true
}

// SetAmountRefundedMinor sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountRefundedMinor(v float32) {
	o.AmountRefundedMinor = v
}

// GetAmountAvailableMinor returns the AmountAvailableMinor field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountAvailableMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountAvailableMinor
}

// GetAmountAvailableMinorOk returns a tuple with the AmountAvailableMinor field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountAvailableMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountAvailableMinor, true
}

// SetAmountAvailableMinor sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountAvailableMinor(v float32) {
	o.AmountAvailableMinor = v
}

// GetCardBrand returns the CardBrand field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardBrand() string {
	if o == nil || o.CardBrand.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// SetCardBrand sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}

// GetCardLast4 returns the CardLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardLast4() string {
	if o == nil || o.CardLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardLast4.Get()
}

// GetCardLast4Ok returns a tuple with the CardLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLast4.Get(), o.CardLast4.IsSet()
}

// SetCardLast4 sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardLast4(v string) {
	o.CardLast4.Set(&v)
}

// GetTerminalSerial returns the TerminalSerial field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetTerminalSerial() string {
	if o == nil || o.TerminalSerial.Get() == nil {
		var ret string
		return ret
	}

	return *o.TerminalSerial.Get()
}

// GetTerminalSerialOk returns a tuple with the TerminalSerial field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreditNoteEligibilityResponsePaymentsInner) GetTerminalSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminalSerial.Get(), o.TerminalSerial.IsSet()
}

// SetTerminalSerial sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetTerminalSerial(v string) {
	o.TerminalSerial.Set(&v)
}

// GetAvailableOperations returns the AvailableOperations field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAvailableOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableOperations
}

// GetAvailableOperationsOk returns a tuple with the AvailableOperations field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetAvailableOperationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableOperations, true
}

// SetAvailableOperations sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetAvailableOperations(v []string) {
	o.AvailableOperations = v
}

// GetConstraints returns the Constraints field value
func (o *CreditNoteEligibilityResponsePaymentsInner) GetConstraints() CreditNoteEligibilityResponsePaymentsInnerConstraints {
	if o == nil {
		var ret CreditNoteEligibilityResponsePaymentsInnerConstraints
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *CreditNoteEligibilityResponsePaymentsInner) GetConstraintsOk() (*CreditNoteEligibilityResponsePaymentsInnerConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraints, true
}

// SetConstraints sets field value
func (o *CreditNoteEligibilityResponsePaymentsInner) SetConstraints(v CreditNoteEligibilityResponsePaymentsInnerConstraints) {
	o.Constraints = v
}

func (o CreditNoteEligibilityResponsePaymentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditNoteEligibilityResponsePaymentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["gateway"] = o.Gateway
	toSerialize["payment_method"] = o.PaymentMethod.Get()
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["amount_refunded_minor"] = o.AmountRefundedMinor
	toSerialize["amount_available_minor"] = o.AmountAvailableMinor
	toSerialize["card_brand"] = o.CardBrand.Get()
	toSerialize["card_last4"] = o.CardLast4.Get()
	toSerialize["terminal_serial"] = o.TerminalSerial.Get()
	toSerialize["available_operations"] = o.AvailableOperations
	toSerialize["constraints"] = o.Constraints

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreditNoteEligibilityResponsePaymentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction_id",
		"gateway",
		"payment_method",
		"amount_minor",
		"amount_refunded_minor",
		"amount_available_minor",
		"card_brand",
		"card_last4",
		"terminal_serial",
		"available_operations",
		"constraints",
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

	varCreditNoteEligibilityResponsePaymentsInner := _CreditNoteEligibilityResponsePaymentsInner{}

	err = json.Unmarshal(data, &varCreditNoteEligibilityResponsePaymentsInner)

	if err != nil {
		return err
	}

	*o = CreditNoteEligibilityResponsePaymentsInner(varCreditNoteEligibilityResponsePaymentsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "amount_refunded_minor")
		delete(additionalProperties, "amount_available_minor")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "card_last4")
		delete(additionalProperties, "terminal_serial")
		delete(additionalProperties, "available_operations")
		delete(additionalProperties, "constraints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditNoteEligibilityResponsePaymentsInner struct {
	value *CreditNoteEligibilityResponsePaymentsInner
	isSet bool
}

func (v NullableCreditNoteEligibilityResponsePaymentsInner) Get() *CreditNoteEligibilityResponsePaymentsInner {
	return v.value
}

func (v *NullableCreditNoteEligibilityResponsePaymentsInner) Set(val *CreditNoteEligibilityResponsePaymentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNoteEligibilityResponsePaymentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNoteEligibilityResponsePaymentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNoteEligibilityResponsePaymentsInner(val *CreditNoteEligibilityResponsePaymentsInner) *NullableCreditNoteEligibilityResponsePaymentsInner {
	return &NullableCreditNoteEligibilityResponsePaymentsInner{value: val, isSet: true}
}

func (v NullableCreditNoteEligibilityResponsePaymentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNoteEligibilityResponsePaymentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


