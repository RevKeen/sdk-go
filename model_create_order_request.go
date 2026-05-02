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

// checks if the CreateOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrderRequest{}

// CreateOrderRequest Parameters for creating a new order with line items and customer association.
type CreateOrderRequest struct {
	// Customer ID (optional for guest orders)
	CustomerId *string `json:"customer_id,omitempty"`
	// Three-letter ISO currency code
	Currency *string `json:"currency,omitempty"`
	// Billing type. one_time: Single purchase. recurring: Subscription-based.
	BillingType *string `json:"billing_type,omitempty"`
	// Order line items (at least one required)
	LineItems []CreateOrderRequestLineItemsInner `json:"line_items"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Notes *string `json:"notes,omitempty"`
	// Expiration timestamp for the order
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ID of the linked invoice (Commercial Truth)
	InvoiceId *string `json:"invoice_id,omitempty"`
	// ID of the linked subscription (for recurring fulfillment)
	SubscriptionId *string `json:"subscription_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateOrderRequest CreateOrderRequest

// NewCreateOrderRequest instantiates a new CreateOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrderRequest(lineItems []CreateOrderRequestLineItemsInner) *CreateOrderRequest {
	this := CreateOrderRequest{}
	var currency string = "usd"
	this.Currency = &currency
	var billingType string = "one_time"
	this.BillingType = &billingType
	this.LineItems = lineItems
	return &this
}

// NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrderRequestWithDefaults() *CreateOrderRequest {
	this := CreateOrderRequest{}
	var currency string = "usd"
	this.Currency = &currency
	var billingType string = "one_time"
	this.BillingType = &billingType
	return &this
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CreateOrderRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CreateOrderRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetBillingType returns the BillingType field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetBillingType() string {
	if o == nil || IsNil(o.BillingType) {
		var ret string
		return ret
	}
	return *o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetBillingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BillingType) {
		return nil, false
	}
	return o.BillingType, true
}

// HasBillingType returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasBillingType() bool {
	if o != nil && !IsNil(o.BillingType) {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given string and assigns it to the BillingType field.
func (o *CreateOrderRequest) SetBillingType(v string) {
	o.BillingType = &v
}

// GetLineItems returns the LineItems field value
func (o *CreateOrderRequest) GetLineItems() []CreateOrderRequestLineItemsInner {
	if o == nil {
		var ret []CreateOrderRequestLineItemsInner
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetLineItemsOk() ([]CreateOrderRequestLineItemsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *CreateOrderRequest) SetLineItems(v []CreateOrderRequestLineItemsInner) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateOrderRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CreateOrderRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *CreateOrderRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *CreateOrderRequest) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *CreateOrderRequest) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrderRequest) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateOrderRequest) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *CreateOrderRequest) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o CreateOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.BillingType) {
		toSerialize["billing_type"] = o.BillingType
	}
	toSerialize["line_items"] = o.LineItems
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"line_items",
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

	varCreateOrderRequest := _CreateOrderRequest{}

	err = json.Unmarshal(data, &varCreateOrderRequest)

	if err != nil {
		return err
	}

	*o = CreateOrderRequest(varCreateOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "billing_type")
		delete(additionalProperties, "line_items")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateOrderRequest struct {
	value *CreateOrderRequest
	isSet bool
}

func (v NullableCreateOrderRequest) Get() *CreateOrderRequest {
	return v.value
}

func (v *NullableCreateOrderRequest) Set(val *CreateOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrderRequest(val *CreateOrderRequest) *NullableCreateOrderRequest {
	return &NullableCreateOrderRequest{value: val, isSet: true}
}

func (v NullableCreateOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


