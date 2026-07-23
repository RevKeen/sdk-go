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
)

// checks if the CreateAccountingInvoicePaymentRequestInputExternalCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountingInvoicePaymentRequestInputExternalCustomer{}

// CreateAccountingInvoicePaymentRequestInputExternalCustomer struct for CreateAccountingInvoicePaymentRequestInputExternalCustomer
type CreateAccountingInvoicePaymentRequestInputExternalCustomer struct {
	Id NullableString `json:"id,omitempty"`
	Reference NullableString `json:"reference,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Email NullableString `json:"email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccountingInvoicePaymentRequestInputExternalCustomer CreateAccountingInvoicePaymentRequestInputExternalCustomer

// NewCreateAccountingInvoicePaymentRequestInputExternalCustomer instantiates a new CreateAccountingInvoicePaymentRequestInputExternalCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountingInvoicePaymentRequestInputExternalCustomer() *CreateAccountingInvoicePaymentRequestInputExternalCustomer {
	this := CreateAccountingInvoicePaymentRequestInputExternalCustomer{}
	return &this
}

// NewCreateAccountingInvoicePaymentRequestInputExternalCustomerWithDefaults instantiates a new CreateAccountingInvoicePaymentRequestInputExternalCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountingInvoicePaymentRequestInputExternalCustomerWithDefaults() *CreateAccountingInvoicePaymentRequestInputExternalCustomer {
	this := CreateAccountingInvoicePaymentRequestInputExternalCustomer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) UnsetId() {
	o.Id.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetReference() string {
	if o == nil || IsNil(o.Reference.Get()) {
		var ret string
		return ret
	}
	return *o.Reference.Get()
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reference.Get(), o.Reference.IsSet()
}

// HasReference returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) HasReference() bool {
	if o != nil && o.Reference.IsSet() {
		return true
	}

	return false
}

// SetReference gets a reference to the given NullableString and assigns it to the Reference field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetReference(v string) {
	o.Reference.Set(&v)
}
// SetReferenceNil sets the value for Reference to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetReferenceNil() {
	o.Reference.Set(nil)
}

// UnsetReference ensures that no value is present for Reference, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) UnsetReference() {
	o.Reference.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) UnsetEmail() {
	o.Email.Unset()
}

func (o CreateAccountingInvoicePaymentRequestInputExternalCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountingInvoicePaymentRequestInputExternalCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Reference.IsSet() {
		toSerialize["reference"] = o.Reference.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAccountingInvoicePaymentRequestInputExternalCustomer) UnmarshalJSON(data []byte) (err error) {
	varCreateAccountingInvoicePaymentRequestInputExternalCustomer := _CreateAccountingInvoicePaymentRequestInputExternalCustomer{}

	err = json.Unmarshal(data, &varCreateAccountingInvoicePaymentRequestInputExternalCustomer)

	if err != nil {
		return err
	}

	*o = CreateAccountingInvoicePaymentRequestInputExternalCustomer(varCreateAccountingInvoicePaymentRequestInputExternalCustomer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "reference")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer struct {
	value *CreateAccountingInvoicePaymentRequestInputExternalCustomer
	isSet bool
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) Get() *CreateAccountingInvoicePaymentRequestInputExternalCustomer {
	return v.value
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) Set(val *CreateAccountingInvoicePaymentRequestInputExternalCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountingInvoicePaymentRequestInputExternalCustomer(val *CreateAccountingInvoicePaymentRequestInputExternalCustomer) *NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer {
	return &NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer{value: val, isSet: true}
}

func (v NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountingInvoicePaymentRequestInputExternalCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


