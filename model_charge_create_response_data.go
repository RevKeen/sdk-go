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
	"time"
	"fmt"
)

// checks if the ChargeCreateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeCreateResponseData{}

// ChargeCreateResponseData struct for ChargeCreateResponseData
type ChargeCreateResponseData struct {
	Id string `json:"id"`
	CustomerId string `json:"customerId"`
	InvoiceId NullableString `json:"invoiceId"`
	PaymentMethodId NullableString `json:"paymentMethodId"`
	AmountMinor int32 `json:"amountMinor"`
	AmountCapturedMinor int32 `json:"amountCapturedMinor"`
	AmountRefundedMinor int32 `json:"amountRefundedMinor"`
	Currency string `json:"currency"`
	Status string `json:"status"`
	Description string `json:"description"`
	StatementDescriptor NullableString `json:"statementDescriptor"`
	FailureCode NullableString `json:"failureCode"`
	FailureMessage NullableString `json:"failureMessage"`
	GatewayTransactionId NullableString `json:"gatewayTransactionId"`
	ReceiptUrl NullableString `json:"receiptUrl"`
	Captured bool `json:"captured"`
	Metadata map[string]interface{} `json:"metadata"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _ChargeCreateResponseData ChargeCreateResponseData

// NewChargeCreateResponseData instantiates a new ChargeCreateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeCreateResponseData(id string, customerId string, invoiceId NullableString, paymentMethodId NullableString, amountMinor int32, amountCapturedMinor int32, amountRefundedMinor int32, currency string, status string, description string, statementDescriptor NullableString, failureCode NullableString, failureMessage NullableString, gatewayTransactionId NullableString, receiptUrl NullableString, captured bool, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time) *ChargeCreateResponseData {
	this := ChargeCreateResponseData{}
	this.Id = id
	this.CustomerId = customerId
	this.InvoiceId = invoiceId
	this.PaymentMethodId = paymentMethodId
	this.AmountMinor = amountMinor
	this.AmountCapturedMinor = amountCapturedMinor
	this.AmountRefundedMinor = amountRefundedMinor
	this.Currency = currency
	this.Status = status
	this.Description = description
	this.StatementDescriptor = statementDescriptor
	this.FailureCode = failureCode
	this.FailureMessage = failureMessage
	this.GatewayTransactionId = gatewayTransactionId
	this.ReceiptUrl = receiptUrl
	this.Captured = captured
	this.Metadata = metadata
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewChargeCreateResponseDataWithDefaults instantiates a new ChargeCreateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeCreateResponseDataWithDefaults() *ChargeCreateResponseData {
	this := ChargeCreateResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *ChargeCreateResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChargeCreateResponseData) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *ChargeCreateResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ChargeCreateResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *ChargeCreateResponseData) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetPaymentMethodId returns the PaymentMethodId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// SetPaymentMethodId sets field value
func (o *ChargeCreateResponseData) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}

// GetAmountMinor returns the AmountMinor field value
func (o *ChargeCreateResponseData) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *ChargeCreateResponseData) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetAmountCapturedMinor returns the AmountCapturedMinor field value
func (o *ChargeCreateResponseData) GetAmountCapturedMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCapturedMinor
}

// GetAmountCapturedMinorOk returns a tuple with the AmountCapturedMinor field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetAmountCapturedMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCapturedMinor, true
}

// SetAmountCapturedMinor sets field value
func (o *ChargeCreateResponseData) SetAmountCapturedMinor(v int32) {
	o.AmountCapturedMinor = v
}

// GetAmountRefundedMinor returns the AmountRefundedMinor field value
func (o *ChargeCreateResponseData) GetAmountRefundedMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountRefundedMinor
}

// GetAmountRefundedMinorOk returns a tuple with the AmountRefundedMinor field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetAmountRefundedMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountRefundedMinor, true
}

// SetAmountRefundedMinor sets field value
func (o *ChargeCreateResponseData) SetAmountRefundedMinor(v int32) {
	o.AmountRefundedMinor = v
}

// GetCurrency returns the Currency field value
func (o *ChargeCreateResponseData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *ChargeCreateResponseData) SetCurrency(v string) {
	o.Currency = v
}

// GetStatus returns the Status field value
func (o *ChargeCreateResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChargeCreateResponseData) SetStatus(v string) {
	o.Status = v
}

// GetDescription returns the Description field value
func (o *ChargeCreateResponseData) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ChargeCreateResponseData) SetDescription(v string) {
	o.Description = v
}

// GetStatementDescriptor returns the StatementDescriptor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetStatementDescriptor() string {
	if o == nil || o.StatementDescriptor.Get() == nil {
		var ret string
		return ret
	}

	return *o.StatementDescriptor.Get()
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetStatementDescriptorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatementDescriptor.Get(), o.StatementDescriptor.IsSet()
}

// SetStatementDescriptor sets field value
func (o *ChargeCreateResponseData) SetStatementDescriptor(v string) {
	o.StatementDescriptor.Set(&v)
}

// GetFailureCode returns the FailureCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetFailureCode() string {
	if o == nil || o.FailureCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureCode.Get()
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetFailureCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCode.Get(), o.FailureCode.IsSet()
}

// SetFailureCode sets field value
func (o *ChargeCreateResponseData) SetFailureCode(v string) {
	o.FailureCode.Set(&v)
}

// GetFailureMessage returns the FailureMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetFailureMessage() string {
	if o == nil || o.FailureMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.FailureMessage.Get()
}

// GetFailureMessageOk returns a tuple with the FailureMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetFailureMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureMessage.Get(), o.FailureMessage.IsSet()
}

// SetFailureMessage sets field value
func (o *ChargeCreateResponseData) SetFailureMessage(v string) {
	o.FailureMessage.Set(&v)
}

// GetGatewayTransactionId returns the GatewayTransactionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetGatewayTransactionId() string {
	if o == nil || o.GatewayTransactionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayTransactionId.Get()
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTransactionId.Get(), o.GatewayTransactionId.IsSet()
}

// SetGatewayTransactionId sets field value
func (o *ChargeCreateResponseData) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId.Set(&v)
}

// GetReceiptUrl returns the ReceiptUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ChargeCreateResponseData) GetReceiptUrl() string {
	if o == nil || o.ReceiptUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ReceiptUrl.Get()
}

// GetReceiptUrlOk returns a tuple with the ReceiptUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChargeCreateResponseData) GetReceiptUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiptUrl.Get(), o.ReceiptUrl.IsSet()
}

// SetReceiptUrl sets field value
func (o *ChargeCreateResponseData) SetReceiptUrl(v string) {
	o.ReceiptUrl.Set(&v)
}

// GetCaptured returns the Captured field value
func (o *ChargeCreateResponseData) GetCaptured() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Captured
}

// GetCapturedOk returns a tuple with the Captured field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetCapturedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Captured, true
}

// SetCaptured sets field value
func (o *ChargeCreateResponseData) SetCaptured(v bool) {
	o.Captured = v
}

// GetMetadata returns the Metadata field value
func (o *ChargeCreateResponseData) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ChargeCreateResponseData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ChargeCreateResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ChargeCreateResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ChargeCreateResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ChargeCreateResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ChargeCreateResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ChargeCreateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeCreateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customerId"] = o.CustomerId
	toSerialize["invoiceId"] = o.InvoiceId.Get()
	toSerialize["paymentMethodId"] = o.PaymentMethodId.Get()
	toSerialize["amountMinor"] = o.AmountMinor
	toSerialize["amountCapturedMinor"] = o.AmountCapturedMinor
	toSerialize["amountRefundedMinor"] = o.AmountRefundedMinor
	toSerialize["currency"] = o.Currency
	toSerialize["status"] = o.Status
	toSerialize["description"] = o.Description
	toSerialize["statementDescriptor"] = o.StatementDescriptor.Get()
	toSerialize["failureCode"] = o.FailureCode.Get()
	toSerialize["failureMessage"] = o.FailureMessage.Get()
	toSerialize["gatewayTransactionId"] = o.GatewayTransactionId.Get()
	toSerialize["receiptUrl"] = o.ReceiptUrl.Get()
	toSerialize["captured"] = o.Captured
	toSerialize["metadata"] = o.Metadata
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChargeCreateResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customerId",
		"invoiceId",
		"paymentMethodId",
		"amountMinor",
		"amountCapturedMinor",
		"amountRefundedMinor",
		"currency",
		"status",
		"description",
		"statementDescriptor",
		"failureCode",
		"failureMessage",
		"gatewayTransactionId",
		"receiptUrl",
		"captured",
		"metadata",
		"createdAt",
		"updatedAt",
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

	varChargeCreateResponseData := _ChargeCreateResponseData{}

	err = json.Unmarshal(data, &varChargeCreateResponseData)

	if err != nil {
		return err
	}

	*o = ChargeCreateResponseData(varChargeCreateResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "invoiceId")
		delete(additionalProperties, "paymentMethodId")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "amountCapturedMinor")
		delete(additionalProperties, "amountRefundedMinor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "status")
		delete(additionalProperties, "description")
		delete(additionalProperties, "statementDescriptor")
		delete(additionalProperties, "failureCode")
		delete(additionalProperties, "failureMessage")
		delete(additionalProperties, "gatewayTransactionId")
		delete(additionalProperties, "receiptUrl")
		delete(additionalProperties, "captured")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChargeCreateResponseData struct {
	value *ChargeCreateResponseData
	isSet bool
}

func (v NullableChargeCreateResponseData) Get() *ChargeCreateResponseData {
	return v.value
}

func (v *NullableChargeCreateResponseData) Set(val *ChargeCreateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeCreateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeCreateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeCreateResponseData(val *ChargeCreateResponseData) *NullableChargeCreateResponseData {
	return &NullableChargeCreateResponseData{value: val, isSet: true}
}

func (v NullableChargeCreateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeCreateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


