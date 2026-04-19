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

// checks if the OpenDispute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenDispute{}

// OpenDispute An open dispute requiring action, with urgency indicators for evidence submission deadlines.
type OpenDispute struct {
	// Unique identifier for the dispute
	Id string `json:"id"`
	// Public-facing dispute identifier
	PublicId string `json:"public_id"`
	// The original payment ID this dispute is for
	PaymentId NullableString `json:"payment_id"`
	// Parent transaction ID in unified transaction model
	ParentTransactionId NullableString `json:"parent_transaction_id"`
	// Payment gateway that processed the original transaction
	Gateway string `json:"gateway"`
	// Disputed amount in cents
	AmountMinor float32 `json:"amount_minor"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Human-readable reason for the dispute
	Reason NullableString `json:"reason"`
	// Name of the customer who filed the dispute
	CustomerName NullableString `json:"customer_name"`
	// Card brand used for the original payment
	CardBrand NullableString `json:"card_brand"`
	// Last four digits of the card used for the original payment
	CardLast4 NullableString `json:"card_last4"`
	// Deadline for submitting evidence to contest the dispute
	EvidenceDueBy NullableString `json:"evidence_due_by"`
	// Number of days remaining until the evidence deadline
	DaysUntilDue NullableFloat32 `json:"days_until_due"`
	// Whether the evidence submission deadline has passed
	IsOverdue bool `json:"is_overdue"`
	// Timestamp when the dispute was filed
	DisputedAt string `json:"disputed_at"`
	AdditionalProperties map[string]interface{}
}

type _OpenDispute OpenDispute

// NewOpenDispute instantiates a new OpenDispute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenDispute(id string, publicId string, paymentId NullableString, parentTransactionId NullableString, gateway string, amountMinor float32, currency string, reason NullableString, customerName NullableString, cardBrand NullableString, cardLast4 NullableString, evidenceDueBy NullableString, daysUntilDue NullableFloat32, isOverdue bool, disputedAt string) *OpenDispute {
	this := OpenDispute{}
	this.Id = id
	this.PublicId = publicId
	this.PaymentId = paymentId
	this.ParentTransactionId = parentTransactionId
	this.Gateway = gateway
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.Reason = reason
	this.CustomerName = customerName
	this.CardBrand = cardBrand
	this.CardLast4 = cardLast4
	this.EvidenceDueBy = evidenceDueBy
	this.DaysUntilDue = daysUntilDue
	this.IsOverdue = isOverdue
	this.DisputedAt = disputedAt
	return &this
}

// NewOpenDisputeWithDefaults instantiates a new OpenDispute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenDisputeWithDefaults() *OpenDispute {
	this := OpenDispute{}
	return &this
}

// GetId returns the Id field value
func (o *OpenDispute) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenDispute) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
func (o *OpenDispute) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *OpenDispute) SetPublicId(v string) {
	o.PublicId = v
}

// GetPaymentId returns the PaymentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// SetPaymentId sets field value
func (o *OpenDispute) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}

// GetParentTransactionId returns the ParentTransactionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetParentTransactionId() string {
	if o == nil || o.ParentTransactionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentTransactionId.Get()
}

// GetParentTransactionIdOk returns a tuple with the ParentTransactionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetParentTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentTransactionId.Get(), o.ParentTransactionId.IsSet()
}

// SetParentTransactionId sets field value
func (o *OpenDispute) SetParentTransactionId(v string) {
	o.ParentTransactionId.Set(&v)
}

// GetGateway returns the Gateway field value
func (o *OpenDispute) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *OpenDispute) SetGateway(v string) {
	o.Gateway = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *OpenDispute) GetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *OpenDispute) SetAmountMinor(v float32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *OpenDispute) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *OpenDispute) SetCurrency(v string) {
	o.Currency = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *OpenDispute) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetCustomerName returns the CustomerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetCustomerName() string {
	if o == nil || o.CustomerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerName.Get()
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerName.Get(), o.CustomerName.IsSet()
}

// SetCustomerName sets field value
func (o *OpenDispute) SetCustomerName(v string) {
	o.CustomerName.Set(&v)
}

// GetCardBrand returns the CardBrand field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetCardBrand() string {
	if o == nil || o.CardBrand.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardBrand.Get()
}

// GetCardBrandOk returns a tuple with the CardBrand field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetCardBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardBrand.Get(), o.CardBrand.IsSet()
}

// SetCardBrand sets field value
func (o *OpenDispute) SetCardBrand(v string) {
	o.CardBrand.Set(&v)
}

// GetCardLast4 returns the CardLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetCardLast4() string {
	if o == nil || o.CardLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardLast4.Get()
}

// GetCardLast4Ok returns a tuple with the CardLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetCardLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardLast4.Get(), o.CardLast4.IsSet()
}

// SetCardLast4 sets field value
func (o *OpenDispute) SetCardLast4(v string) {
	o.CardLast4.Set(&v)
}

// GetEvidenceDueBy returns the EvidenceDueBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpenDispute) GetEvidenceDueBy() string {
	if o == nil || o.EvidenceDueBy.Get() == nil {
		var ret string
		return ret
	}

	return *o.EvidenceDueBy.Get()
}

// GetEvidenceDueByOk returns a tuple with the EvidenceDueBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetEvidenceDueByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvidenceDueBy.Get(), o.EvidenceDueBy.IsSet()
}

// SetEvidenceDueBy sets field value
func (o *OpenDispute) SetEvidenceDueBy(v string) {
	o.EvidenceDueBy.Set(&v)
}

// GetDaysUntilDue returns the DaysUntilDue field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *OpenDispute) GetDaysUntilDue() float32 {
	if o == nil || o.DaysUntilDue.Get() == nil {
		var ret float32
		return ret
	}

	return *o.DaysUntilDue.Get()
}

// GetDaysUntilDueOk returns a tuple with the DaysUntilDue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenDispute) GetDaysUntilDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DaysUntilDue.Get(), o.DaysUntilDue.IsSet()
}

// SetDaysUntilDue sets field value
func (o *OpenDispute) SetDaysUntilDue(v float32) {
	o.DaysUntilDue.Set(&v)
}

// GetIsOverdue returns the IsOverdue field value
func (o *OpenDispute) GetIsOverdue() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOverdue
}

// GetIsOverdueOk returns a tuple with the IsOverdue field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetIsOverdueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOverdue, true
}

// SetIsOverdue sets field value
func (o *OpenDispute) SetIsOverdue(v bool) {
	o.IsOverdue = v
}

// GetDisputedAt returns the DisputedAt field value
func (o *OpenDispute) GetDisputedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisputedAt
}

// GetDisputedAtOk returns a tuple with the DisputedAt field value
// and a boolean to check if the value has been set.
func (o *OpenDispute) GetDisputedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisputedAt, true
}

// SetDisputedAt sets field value
func (o *OpenDispute) SetDisputedAt(v string) {
	o.DisputedAt = v
}

func (o OpenDispute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenDispute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId
	toSerialize["payment_id"] = o.PaymentId.Get()
	toSerialize["parent_transaction_id"] = o.ParentTransactionId.Get()
	toSerialize["gateway"] = o.Gateway
	toSerialize["amount_minor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["customer_name"] = o.CustomerName.Get()
	toSerialize["card_brand"] = o.CardBrand.Get()
	toSerialize["card_last4"] = o.CardLast4.Get()
	toSerialize["evidence_due_by"] = o.EvidenceDueBy.Get()
	toSerialize["days_until_due"] = o.DaysUntilDue.Get()
	toSerialize["is_overdue"] = o.IsOverdue
	toSerialize["disputed_at"] = o.DisputedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenDispute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"public_id",
		"payment_id",
		"parent_transaction_id",
		"gateway",
		"amount_minor",
		"currency",
		"reason",
		"customer_name",
		"card_brand",
		"card_last4",
		"evidence_due_by",
		"days_until_due",
		"is_overdue",
		"disputed_at",
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

	varOpenDispute := _OpenDispute{}

	err = json.Unmarshal(data, &varOpenDispute)

	if err != nil {
		return err
	}

	*o = OpenDispute(varOpenDispute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "parent_transaction_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "customer_name")
		delete(additionalProperties, "card_brand")
		delete(additionalProperties, "card_last4")
		delete(additionalProperties, "evidence_due_by")
		delete(additionalProperties, "days_until_due")
		delete(additionalProperties, "is_overdue")
		delete(additionalProperties, "disputed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenDispute struct {
	value *OpenDispute
	isSet bool
}

func (v NullableOpenDispute) Get() *OpenDispute {
	return v.value
}

func (v *NullableOpenDispute) Set(val *OpenDispute) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenDispute) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenDispute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenDispute(val *OpenDispute) *NullableOpenDispute {
	return &NullableOpenDispute{value: val, isSet: true}
}

func (v NullableOpenDispute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenDispute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


