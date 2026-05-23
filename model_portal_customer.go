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

// checks if the PortalCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortalCustomer{}

// PortalCustomer The authenticated customer, as visible through the customer-portal surface. This is a deliberately thinner projection than the merchant-facing `Customer` — it omits tax IDs, internal metadata, and anything that could leak merchant-side context to the customer.
type PortalCustomer struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Email NullableString `json:"email"`
	Name NullableString `json:"name"`
	Phone NullableString `json:"phone"`
	CompanyName NullableString `json:"company_name"`
	DefaultCurrency NullableString `json:"default_currency"`
	Locale NullableString `json:"locale"`
	CreatedAt time.Time `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _PortalCustomer PortalCustomer

// NewPortalCustomer instantiates a new PortalCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalCustomer(id string, object string, email NullableString, name NullableString, phone NullableString, companyName NullableString, defaultCurrency NullableString, locale NullableString, createdAt time.Time) *PortalCustomer {
	this := PortalCustomer{}
	this.Id = id
	this.Object = object
	this.Email = email
	this.Name = name
	this.Phone = phone
	this.CompanyName = companyName
	this.DefaultCurrency = defaultCurrency
	this.Locale = locale
	this.CreatedAt = createdAt
	return &this
}

// NewPortalCustomerWithDefaults instantiates a new PortalCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalCustomerWithDefaults() *PortalCustomer {
	this := PortalCustomer{}
	return &this
}

// GetId returns the Id field value
func (o *PortalCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PortalCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PortalCustomer) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PortalCustomer) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PortalCustomer) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PortalCustomer) SetObject(v string) {
	o.Object = v
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *PortalCustomer) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PortalCustomer) SetName(v string) {
	o.Name.Set(&v)
}

// GetPhone returns the Phone field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}

	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// SetPhone sets field value
func (o *PortalCustomer) SetPhone(v string) {
	o.Phone.Set(&v)
}

// GetCompanyName returns the CompanyName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetCompanyName() string {
	if o == nil || o.CompanyName.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompanyName.Get()
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyName.Get(), o.CompanyName.IsSet()
}

// SetCompanyName sets field value
func (o *PortalCustomer) SetCompanyName(v string) {
	o.CompanyName.Set(&v)
}

// GetDefaultCurrency returns the DefaultCurrency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetDefaultCurrency() string {
	if o == nil || o.DefaultCurrency.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultCurrency.Get()
}

// GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetDefaultCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultCurrency.Get(), o.DefaultCurrency.IsSet()
}

// SetDefaultCurrency sets field value
func (o *PortalCustomer) SetDefaultCurrency(v string) {
	o.DefaultCurrency.Set(&v)
}

// GetLocale returns the Locale field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PortalCustomer) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}

	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PortalCustomer) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// SetLocale sets field value
func (o *PortalCustomer) SetLocale(v string) {
	o.Locale.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *PortalCustomer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PortalCustomer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PortalCustomer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PortalCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortalCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["email"] = o.Email.Get()
	toSerialize["name"] = o.Name.Get()
	toSerialize["phone"] = o.Phone.Get()
	toSerialize["company_name"] = o.CompanyName.Get()
	toSerialize["default_currency"] = o.DefaultCurrency.Get()
	toSerialize["locale"] = o.Locale.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortalCustomer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"email",
		"name",
		"phone",
		"company_name",
		"default_currency",
		"locale",
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

	varPortalCustomer := _PortalCustomer{}

	err = json.Unmarshal(data, &varPortalCustomer)

	if err != nil {
		return err
	}

	*o = PortalCustomer(varPortalCustomer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "company_name")
		delete(additionalProperties, "default_currency")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortalCustomer struct {
	value *PortalCustomer
	isSet bool
}

func (v NullablePortalCustomer) Get() *PortalCustomer {
	return v.value
}

func (v *NullablePortalCustomer) Set(val *PortalCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalCustomer(val *PortalCustomer) *NullablePortalCustomer {
	return &NullablePortalCustomer{value: val, isSet: true}
}

func (v NullablePortalCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


