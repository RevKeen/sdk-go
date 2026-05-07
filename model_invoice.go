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

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoice{}

// Invoice The canonical record of what a customer owes or has paid. Invoices track line items, amounts, and payment status through their lifecycle: draft → open → finalized → paid.
type Invoice struct {
	// Unique invoice identifier
	Id string `json:"id"`
	// ID of the customer this invoice belongs to
	CustomerUuid string `json:"customer_uuid"`
	// Merchant-assigned invoice number for reference
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	// Total amount in minor units (cents)
	TotalMinor int32 `json:"total_minor"`
	// Three-letter ISO 4217 currency code
	Currency string `json:"currency"`
	// Allowed payment methods for checkout sessions against this invoice
	AllowedMethods []string `json:"allowed_methods,omitempty"`
	// Invoice lifecycle status: draft, open, finalized, paid, void, or uncollectible
	Status string `json:"status"`
	// Payment due date in ISO 8601 format
	DueDate NullableString `json:"due_date,omitempty"`
	// Key-value pairs for custom fields
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	// External integration source (e.g., practicehub, wodify)
	ExternalSource NullableString `json:"external_source,omitempty"`
	// Type within external system (e.g., appointment, membership)
	ExternalType NullableString `json:"external_type,omitempty"`
	// ID from external system
	ExternalId NullableString `json:"external_id,omitempty"`
	SubscriptionTerms *InvoiceSubscriptionTerms `json:"subscription_terms,omitempty"`
	// When the invoice was created (ISO 8601)
	CreatedAt string `json:"created_at"`
	// When the invoice was last updated (ISO 8601)
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Invoice Invoice

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice(id string, customerUuid string, totalMinor int32, currency string, status string, createdAt string, updatedAt string) *Invoice {
	this := Invoice{}
	this.Id = id
	this.CustomerUuid = customerUuid
	this.TotalMinor = totalMinor
	this.Currency = currency
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value
func (o *Invoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Invoice) SetId(v string) {
	o.Id = v
}

// GetCustomerUuid returns the CustomerUuid field value
func (o *Invoice) GetCustomerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerUuid
}

// GetCustomerUuidOk returns a tuple with the CustomerUuid field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerUuid, true
}

// SetCustomerUuid sets field value
func (o *Invoice) SetCustomerUuid(v string) {
	o.CustomerUuid = v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *Invoice) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *Invoice) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *Invoice) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetTotalMinor returns the TotalMinor field value
func (o *Invoice) GetTotalMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalMinor
}

// GetTotalMinorOk returns a tuple with the TotalMinor field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetTotalMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalMinor, true
}

// SetTotalMinor sets field value
func (o *Invoice) SetTotalMinor(v int32) {
	o.TotalMinor = v
}

// GetCurrency returns the Currency field value
func (o *Invoice) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Invoice) SetCurrency(v string) {
	o.Currency = v
}

// GetAllowedMethods returns the AllowedMethods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoice) GetAllowedMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedMethods
}

// GetAllowedMethodsOk returns a tuple with the AllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoice) GetAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMethods) {
		return nil, false
	}
	return o.AllowedMethods, true
}

// HasAllowedMethods returns a boolean if a field has been set.
func (o *Invoice) HasAllowedMethods() bool {
	if o != nil && !IsNil(o.AllowedMethods) {
		return true
	}

	return false
}

// SetAllowedMethods gets a reference to the given []string and assigns it to the AllowedMethods field.
func (o *Invoice) SetAllowedMethods(v []string) {
	o.AllowedMethods = v
}

// GetStatus returns the Status field value
func (o *Invoice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Invoice) SetStatus(v string) {
	o.Status = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoice) GetDueDate() string {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret string
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoice) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *Invoice) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableString and assigns it to the DueDate field.
func (o *Invoice) SetDueDate(v string) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *Invoice) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *Invoice) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Invoice) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Invoice) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Invoice) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoice) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalSource.Get()
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoice) GetExternalSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalSource.Get(), o.ExternalSource.IsSet()
}

// HasExternalSource returns a boolean if a field has been set.
func (o *Invoice) HasExternalSource() bool {
	if o != nil && o.ExternalSource.IsSet() {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given NullableString and assigns it to the ExternalSource field.
func (o *Invoice) SetExternalSource(v string) {
	o.ExternalSource.Set(&v)
}
// SetExternalSourceNil sets the value for ExternalSource to be an explicit nil
func (o *Invoice) SetExternalSourceNil() {
	o.ExternalSource.Set(nil)
}

// UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
func (o *Invoice) UnsetExternalSource() {
	o.ExternalSource.Unset()
}

// GetExternalType returns the ExternalType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoice) GetExternalType() string {
	if o == nil || IsNil(o.ExternalType.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalType.Get()
}

// GetExternalTypeOk returns a tuple with the ExternalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoice) GetExternalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalType.Get(), o.ExternalType.IsSet()
}

// HasExternalType returns a boolean if a field has been set.
func (o *Invoice) HasExternalType() bool {
	if o != nil && o.ExternalType.IsSet() {
		return true
	}

	return false
}

// SetExternalType gets a reference to the given NullableString and assigns it to the ExternalType field.
func (o *Invoice) SetExternalType(v string) {
	o.ExternalType.Set(&v)
}
// SetExternalTypeNil sets the value for ExternalType to be an explicit nil
func (o *Invoice) SetExternalTypeNil() {
	o.ExternalType.Set(nil)
}

// UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
func (o *Invoice) UnsetExternalType() {
	o.ExternalType.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invoice) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invoice) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *Invoice) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *Invoice) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *Invoice) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *Invoice) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetSubscriptionTerms returns the SubscriptionTerms field value if set, zero value otherwise.
func (o *Invoice) GetSubscriptionTerms() InvoiceSubscriptionTerms {
	if o == nil || IsNil(o.SubscriptionTerms) {
		var ret InvoiceSubscriptionTerms
		return ret
	}
	return *o.SubscriptionTerms
}

// GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetSubscriptionTermsOk() (*InvoiceSubscriptionTerms, bool) {
	if o == nil || IsNil(o.SubscriptionTerms) {
		return nil, false
	}
	return o.SubscriptionTerms, true
}

// HasSubscriptionTerms returns a boolean if a field has been set.
func (o *Invoice) HasSubscriptionTerms() bool {
	if o != nil && !IsNil(o.SubscriptionTerms) {
		return true
	}

	return false
}

// SetSubscriptionTerms gets a reference to the given InvoiceSubscriptionTerms and assigns it to the SubscriptionTerms field.
func (o *Invoice) SetSubscriptionTerms(v InvoiceSubscriptionTerms) {
	o.SubscriptionTerms = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Invoice) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Invoice) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Invoice) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Invoice) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Invoice) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customer_uuid"] = o.CustomerUuid
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoice_number"] = o.InvoiceNumber
	}
	toSerialize["total_minor"] = o.TotalMinor
	toSerialize["currency"] = o.Currency
	if o.AllowedMethods != nil {
		toSerialize["allowed_methods"] = o.AllowedMethods
	}
	toSerialize["status"] = o.Status
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.ExternalSource.IsSet() {
		toSerialize["external_source"] = o.ExternalSource.Get()
	}
	if o.ExternalType.IsSet() {
		toSerialize["external_type"] = o.ExternalType.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if !IsNil(o.SubscriptionTerms) {
		toSerialize["subscription_terms"] = o.SubscriptionTerms
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Invoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customer_uuid",
		"total_minor",
		"currency",
		"status",
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

	varInvoice := _Invoice{}

	err = json.Unmarshal(data, &varInvoice)

	if err != nil {
		return err
	}

	*o = Invoice(varInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_uuid")
		delete(additionalProperties, "invoice_number")
		delete(additionalProperties, "total_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "allowed_methods")
		delete(additionalProperties, "status")
		delete(additionalProperties, "due_date")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "external_source")
		delete(additionalProperties, "external_type")
		delete(additionalProperties, "external_id")
		delete(additionalProperties, "subscription_terms")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


