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
	"time"
	"fmt"
)

// checks if the CreateAccountingInvoicePaymentRequestInputExternalInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountingInvoicePaymentRequestInputExternalInvoice{}

// CreateAccountingInvoicePaymentRequestInputExternalInvoice struct for CreateAccountingInvoicePaymentRequestInputExternalInvoice
type CreateAccountingInvoicePaymentRequestInputExternalInvoice struct {
	Id string `json:"id"`
	Number NullableString `json:"number,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Url NullableString `json:"url,omitempty"`
	IssuedAt NullableTime `json:"issued_at,omitempty"`
	DueAt NullableTime `json:"due_at,omitempty"`
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccountingInvoicePaymentRequestInputExternalInvoice CreateAccountingInvoicePaymentRequestInputExternalInvoice

// NewCreateAccountingInvoicePaymentRequestInputExternalInvoice instantiates a new CreateAccountingInvoicePaymentRequestInputExternalInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountingInvoicePaymentRequestInputExternalInvoice(id string) *CreateAccountingInvoicePaymentRequestInputExternalInvoice {
	this := CreateAccountingInvoicePaymentRequestInputExternalInvoice{}
	this.Id = id
	return &this
}

// NewCreateAccountingInvoicePaymentRequestInputExternalInvoiceWithDefaults instantiates a new CreateAccountingInvoicePaymentRequestInputExternalInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountingInvoicePaymentRequestInputExternalInvoiceWithDefaults() *CreateAccountingInvoicePaymentRequestInputExternalInvoice {
	this := CreateAccountingInvoicePaymentRequestInputExternalInvoice{}
	return &this
}

// GetId returns the Id field value
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetId(v string) {
	o.Id = v
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetNumber() string {
	if o == nil || IsNil(o.Number.Get()) {
		var ret string
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableString and assigns it to the Number field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetNumber(v string) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetNumber() {
	o.Number.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetStatus() {
	o.Status.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetUrl() {
	o.Url.Unset()
}

// GetIssuedAt returns the IssuedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIssuedAt() time.Time {
	if o == nil || IsNil(o.IssuedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.IssuedAt.Get()
}

// GetIssuedAtOk returns a tuple with the IssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIssuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuedAt.Get(), o.IssuedAt.IsSet()
}

// HasIssuedAt returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasIssuedAt() bool {
	if o != nil && o.IssuedAt.IsSet() {
		return true
	}

	return false
}

// SetIssuedAt gets a reference to the given NullableTime and assigns it to the IssuedAt field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetIssuedAt(v time.Time) {
	o.IssuedAt.Set(&v)
}
// SetIssuedAtNil sets the value for IssuedAt to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetIssuedAtNil() {
	o.IssuedAt.Set(nil)
}

// UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetIssuedAt() {
	o.IssuedAt.Unset()
}

// GetDueAt returns the DueAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetDueAt() time.Time {
	if o == nil || IsNil(o.DueAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueAt.Get()
}

// GetDueAtOk returns a tuple with the DueAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetDueAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueAt.Get(), o.DueAt.IsSet()
}

// HasDueAt returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasDueAt() bool {
	if o != nil && o.DueAt.IsSet() {
		return true
	}

	return false
}

// SetDueAt gets a reference to the given NullableTime and assigns it to the DueAt field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetDueAt(v time.Time) {
	o.DueAt.Set(&v)
}
// SetDueAtNil sets the value for DueAt to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetDueAtNil() {
	o.DueAt.Set(nil)
}

// UnsetDueAt ensures that no value is present for DueAt, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetDueAt() {
	o.DueAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o CreateAccountingInvoicePaymentRequestInputExternalInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountingInvoicePaymentRequestInputExternalInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.IssuedAt.IsSet() {
		toSerialize["issued_at"] = o.IssuedAt.Get()
	}
	if o.DueAt.IsSet() {
		toSerialize["due_at"] = o.DueAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCreateAccountingInvoicePaymentRequestInputExternalInvoice := _CreateAccountingInvoicePaymentRequestInputExternalInvoice{}

	err = json.Unmarshal(data, &varCreateAccountingInvoicePaymentRequestInputExternalInvoice)

	if err != nil {
		return err
	}

	*o = CreateAccountingInvoicePaymentRequestInputExternalInvoice(varCreateAccountingInvoicePaymentRequestInputExternalInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "number")
		delete(additionalProperties, "status")
		delete(additionalProperties, "url")
		delete(additionalProperties, "issued_at")
		delete(additionalProperties, "due_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice struct {
	value *CreateAccountingInvoicePaymentRequestInputExternalInvoice
	isSet bool
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) Get() *CreateAccountingInvoicePaymentRequestInputExternalInvoice {
	return v.value
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) Set(val *CreateAccountingInvoicePaymentRequestInputExternalInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountingInvoicePaymentRequestInputExternalInvoice(val *CreateAccountingInvoicePaymentRequestInputExternalInvoice) *NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice {
	return &NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice{value: val, isSet: true}
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


