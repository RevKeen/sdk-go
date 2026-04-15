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
	"time"
	"fmt"
)

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethod{}

// PaymentMethod A saved payment instrument (card, bank account, etc.) attached to a customer for future charges.
type PaymentMethod struct {
	// Unique identifier for the payment method
	Id string `json:"id"`
	// Object type, always 'payment_method'
	Object string `json:"object"`
	// Public ID visible in API responses (pm_xxx format)
	PublicId NullableString `json:"public_id"`
	// The type of payment method
	Type string `json:"type"`
	// The status of the payment method
	Status string `json:"status"`
	// ID of the customer this payment method belongs to
	CustomerId string `json:"customer_id"`
	// Whether this is the customer's default payment method
	IsDefault bool `json:"is_default"`
	Card CardDetails `json:"card"`
	UsBankAccount BankAccountDetails `json:"us_bank_account"`
	BillingDetails BillingDetails `json:"billing_details"`
	// Custom metadata attached to the payment method
	Metadata map[string]interface{} `json:"metadata"`
	// When the payment method was created
	CreatedAt time.Time `json:"created_at"`
	// When the payment method was last updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethod PaymentMethod

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod(id string, object string, publicId NullableString, type_ string, status string, customerId string, isDefault bool, card CardDetails, usBankAccount BankAccountDetails, billingDetails BillingDetails, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time) *PaymentMethod {
	this := PaymentMethod{}
	this.Id = id
	this.Object = object
	this.PublicId = publicId
	this.Type = type_
	this.Status = status
	this.CustomerId = customerId
	this.IsDefault = isDefault
	this.Card = card
	this.UsBankAccount = usBankAccount
	this.BillingDetails = billingDetails
	this.Metadata = metadata
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentMethod) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentMethod) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentMethod) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentMethod) SetObject(v string) {
	o.Object = v
}

// GetPublicId returns the PublicId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentMethod) GetPublicId() string {
	if o == nil || o.PublicId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicId.Get()
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicId.Get(), o.PublicId.IsSet()
}

// SetPublicId sets field value
func (o *PaymentMethod) SetPublicId(v string) {
	o.PublicId.Set(&v)
}

// GetType returns the Type field value
func (o *PaymentMethod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaymentMethod) SetType(v string) {
	o.Type = v
}

// GetStatus returns the Status field value
func (o *PaymentMethod) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentMethod) SetStatus(v string) {
	o.Status = v
}

// GetCustomerId returns the CustomerId field value
func (o *PaymentMethod) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *PaymentMethod) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetIsDefault returns the IsDefault field value
func (o *PaymentMethod) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *PaymentMethod) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetCard returns the Card field value
func (o *PaymentMethod) GetCard() CardDetails {
	if o == nil {
		var ret CardDetails
		return ret
	}

	return o.Card
}

// GetCardOk returns a tuple with the Card field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCardOk() (*CardDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Card, true
}

// SetCard sets field value
func (o *PaymentMethod) SetCard(v CardDetails) {
	o.Card = v
}

// GetUsBankAccount returns the UsBankAccount field value
func (o *PaymentMethod) GetUsBankAccount() BankAccountDetails {
	if o == nil {
		var ret BankAccountDetails
		return ret
	}

	return o.UsBankAccount
}

// GetUsBankAccountOk returns a tuple with the UsBankAccount field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUsBankAccountOk() (*BankAccountDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsBankAccount, true
}

// SetUsBankAccount sets field value
func (o *PaymentMethod) SetUsBankAccount(v BankAccountDetails) {
	o.UsBankAccount = v
}

// GetBillingDetails returns the BillingDetails field value
func (o *PaymentMethod) GetBillingDetails() BillingDetails {
	if o == nil {
		var ret BillingDetails
		return ret
	}

	return o.BillingDetails
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetBillingDetailsOk() (*BillingDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingDetails, true
}

// SetBillingDetails sets field value
func (o *PaymentMethod) SetBillingDetails(v BillingDetails) {
	o.BillingDetails = v
}

// GetMetadata returns the Metadata field value
func (o *PaymentMethod) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *PaymentMethod) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentMethod) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentMethod) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaymentMethod) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaymentMethod) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["public_id"] = o.PublicId.Get()
	toSerialize["type"] = o.Type
	toSerialize["status"] = o.Status
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["is_default"] = o.IsDefault
	toSerialize["card"] = o.Card
	toSerialize["us_bank_account"] = o.UsBankAccount
	toSerialize["billing_details"] = o.BillingDetails
	toSerialize["metadata"] = o.Metadata
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"public_id",
		"type",
		"status",
		"customer_id",
		"is_default",
		"card",
		"us_bank_account",
		"billing_details",
		"metadata",
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

	varPaymentMethod := _PaymentMethod{}

	err = json.Unmarshal(data, &varPaymentMethod)

	if err != nil {
		return err
	}

	*o = PaymentMethod(varPaymentMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "is_default")
		delete(additionalProperties, "card")
		delete(additionalProperties, "us_bank_account")
		delete(additionalProperties, "billing_details")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


