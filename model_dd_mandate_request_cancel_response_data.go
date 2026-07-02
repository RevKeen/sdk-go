/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the DdMandateRequestCancelResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdMandateRequestCancelResponseData{}

// DdMandateRequestCancelResponseData struct for DdMandateRequestCancelResponseData
type DdMandateRequestCancelResponseData struct {
	Id string `json:"id"`
	CustomerId string `json:"customer_id"`
	InvoiceId NullableString `json:"invoice_id"`
	Status string `json:"status"`
	CancelledAt NullableString `json:"cancelled_at,omitempty"`
	CancelledBy NullableString `json:"cancelled_by,omitempty"`
	ExpiredAt NullableString `json:"expired_at,omitempty"`
	Reason NullableString `json:"reason,omitempty"`
	Idempotent *bool `json:"idempotent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DdMandateRequestCancelResponseData DdMandateRequestCancelResponseData

// NewDdMandateRequestCancelResponseData instantiates a new DdMandateRequestCancelResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdMandateRequestCancelResponseData(id string, customerId string, invoiceId NullableString, status string) *DdMandateRequestCancelResponseData {
	this := DdMandateRequestCancelResponseData{}
	this.Id = id
	this.CustomerId = customerId
	this.InvoiceId = invoiceId
	this.Status = status
	return &this
}

// NewDdMandateRequestCancelResponseDataWithDefaults instantiates a new DdMandateRequestCancelResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdMandateRequestCancelResponseDataWithDefaults() *DdMandateRequestCancelResponseData {
	this := DdMandateRequestCancelResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *DdMandateRequestCancelResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestCancelResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DdMandateRequestCancelResponseData) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *DdMandateRequestCancelResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestCancelResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DdMandateRequestCancelResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestCancelResponseData) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestCancelResponseData) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *DdMandateRequestCancelResponseData) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetStatus returns the Status field value
func (o *DdMandateRequestCancelResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestCancelResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DdMandateRequestCancelResponseData) SetStatus(v string) {
	o.Status = v
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdMandateRequestCancelResponseData) GetCancelledAt() string {
	if o == nil || IsNil(o.CancelledAt.Get()) {
		var ret string
		return ret
	}
	return *o.CancelledAt.Get()
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestCancelResponseData) GetCancelledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledAt.Get(), o.CancelledAt.IsSet()
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *DdMandateRequestCancelResponseData) HasCancelledAt() bool {
	if o != nil && o.CancelledAt.IsSet() {
		return true
	}

	return false
}

// SetCancelledAt gets a reference to the given NullableString and assigns it to the CancelledAt field.
func (o *DdMandateRequestCancelResponseData) SetCancelledAt(v string) {
	o.CancelledAt.Set(&v)
}
// SetCancelledAtNil sets the value for CancelledAt to be an explicit nil
func (o *DdMandateRequestCancelResponseData) SetCancelledAtNil() {
	o.CancelledAt.Set(nil)
}

// UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
func (o *DdMandateRequestCancelResponseData) UnsetCancelledAt() {
	o.CancelledAt.Unset()
}

// GetCancelledBy returns the CancelledBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdMandateRequestCancelResponseData) GetCancelledBy() string {
	if o == nil || IsNil(o.CancelledBy.Get()) {
		var ret string
		return ret
	}
	return *o.CancelledBy.Get()
}

// GetCancelledByOk returns a tuple with the CancelledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestCancelResponseData) GetCancelledByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledBy.Get(), o.CancelledBy.IsSet()
}

// HasCancelledBy returns a boolean if a field has been set.
func (o *DdMandateRequestCancelResponseData) HasCancelledBy() bool {
	if o != nil && o.CancelledBy.IsSet() {
		return true
	}

	return false
}

// SetCancelledBy gets a reference to the given NullableString and assigns it to the CancelledBy field.
func (o *DdMandateRequestCancelResponseData) SetCancelledBy(v string) {
	o.CancelledBy.Set(&v)
}
// SetCancelledByNil sets the value for CancelledBy to be an explicit nil
func (o *DdMandateRequestCancelResponseData) SetCancelledByNil() {
	o.CancelledBy.Set(nil)
}

// UnsetCancelledBy ensures that no value is present for CancelledBy, not even an explicit nil
func (o *DdMandateRequestCancelResponseData) UnsetCancelledBy() {
	o.CancelledBy.Unset()
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdMandateRequestCancelResponseData) GetExpiredAt() string {
	if o == nil || IsNil(o.ExpiredAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestCancelResponseData) GetExpiredAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *DdMandateRequestCancelResponseData) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableString and assigns it to the ExpiredAt field.
func (o *DdMandateRequestCancelResponseData) SetExpiredAt(v string) {
	o.ExpiredAt.Set(&v)
}
// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *DdMandateRequestCancelResponseData) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *DdMandateRequestCancelResponseData) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DdMandateRequestCancelResponseData) GetReason() string {
	if o == nil || IsNil(o.Reason.Get()) {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestCancelResponseData) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *DdMandateRequestCancelResponseData) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *DdMandateRequestCancelResponseData) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *DdMandateRequestCancelResponseData) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *DdMandateRequestCancelResponseData) UnsetReason() {
	o.Reason.Unset()
}

// GetIdempotent returns the Idempotent field value if set, zero value otherwise.
func (o *DdMandateRequestCancelResponseData) GetIdempotent() bool {
	if o == nil || IsNil(o.Idempotent) {
		var ret bool
		return ret
	}
	return *o.Idempotent
}

// GetIdempotentOk returns a tuple with the Idempotent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdMandateRequestCancelResponseData) GetIdempotentOk() (*bool, bool) {
	if o == nil || IsNil(o.Idempotent) {
		return nil, false
	}
	return o.Idempotent, true
}

// HasIdempotent returns a boolean if a field has been set.
func (o *DdMandateRequestCancelResponseData) HasIdempotent() bool {
	if o != nil && !IsNil(o.Idempotent) {
		return true
	}

	return false
}

// SetIdempotent gets a reference to the given bool and assigns it to the Idempotent field.
func (o *DdMandateRequestCancelResponseData) SetIdempotent(v bool) {
	o.Idempotent = &v
}

func (o DdMandateRequestCancelResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdMandateRequestCancelResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["status"] = o.Status
	if o.CancelledAt.IsSet() {
		toSerialize["cancelled_at"] = o.CancelledAt.Get()
	}
	if o.CancelledBy.IsSet() {
		toSerialize["cancelled_by"] = o.CancelledBy.Get()
	}
	if o.ExpiredAt.IsSet() {
		toSerialize["expired_at"] = o.ExpiredAt.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if !IsNil(o.Idempotent) {
		toSerialize["idempotent"] = o.Idempotent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdMandateRequestCancelResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"customer_id",
		"invoice_id",
		"status",
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

	varDdMandateRequestCancelResponseData := _DdMandateRequestCancelResponseData{}

	err = json.Unmarshal(data, &varDdMandateRequestCancelResponseData)

	if err != nil {
		return err
	}

	*o = DdMandateRequestCancelResponseData(varDdMandateRequestCancelResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "cancelled_at")
		delete(additionalProperties, "cancelled_by")
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "idempotent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdMandateRequestCancelResponseData struct {
	value *DdMandateRequestCancelResponseData
	isSet bool
}

func (v NullableDdMandateRequestCancelResponseData) Get() *DdMandateRequestCancelResponseData {
	return v.value
}

func (v *NullableDdMandateRequestCancelResponseData) Set(val *DdMandateRequestCancelResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDdMandateRequestCancelResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDdMandateRequestCancelResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdMandateRequestCancelResponseData(val *DdMandateRequestCancelResponseData) *NullableDdMandateRequestCancelResponseData {
	return &NullableDdMandateRequestCancelResponseData{value: val, isSet: true}
}

func (v NullableDdMandateRequestCancelResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdMandateRequestCancelResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


