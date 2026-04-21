/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the CustomerCreateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCreateResponseData{}

// CustomerCreateResponseData struct for CustomerCreateResponseData
type CustomerCreateResponseData struct {
	// Unique customer identifier
	Id string `json:"id"`
	// ID of the merchant this customer belongs to
	MerchantId string `json:"merchantId"`
	// Merchant-assigned reference ID for external system mapping
	MerchantRefId *string `json:"merchantRefId,omitempty"`
	// Customer's email address
	Email string `json:"email"`
	// Customer's full name
	Name *string `json:"name,omitempty"`
	// Customer's phone number
	Phone *string `json:"phone,omitempty"`
	// Linked Better Auth user ID for portal access
	AuthUserId *string `json:"authUserId,omitempty"`
	// Arbitrary key-value metadata attached to this customer
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Merchant-defined custom fields displayed as columns in the dashboard
	CustomFields map[string]interface{} `json:"customFields,omitempty"`
	// When the customer was created (ISO 8601)
	CreatedAt time.Time `json:"createdAt"`
	// When the customer was last updated (ISO 8601)
	UpdatedAt time.Time `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _CustomerCreateResponseData CustomerCreateResponseData

// NewCustomerCreateResponseData instantiates a new CustomerCreateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateResponseData(id string, merchantId string, email string, createdAt time.Time, updatedAt time.Time) *CustomerCreateResponseData {
	this := CustomerCreateResponseData{}
	this.Id = id
	this.MerchantId = merchantId
	this.Email = email
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCustomerCreateResponseDataWithDefaults instantiates a new CustomerCreateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateResponseDataWithDefaults() *CustomerCreateResponseData {
	this := CustomerCreateResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *CustomerCreateResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerCreateResponseData) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *CustomerCreateResponseData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *CustomerCreateResponseData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetMerchantRefId returns the MerchantRefId field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetMerchantRefId() string {
	if o == nil || IsNil(o.MerchantRefId) {
		var ret string
		return ret
	}
	return *o.MerchantRefId
}

// GetMerchantRefIdOk returns a tuple with the MerchantRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetMerchantRefIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantRefId) {
		return nil, false
	}
	return o.MerchantRefId, true
}

// HasMerchantRefId returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasMerchantRefId() bool {
	if o != nil && !IsNil(o.MerchantRefId) {
		return true
	}

	return false
}

// SetMerchantRefId gets a reference to the given string and assigns it to the MerchantRefId field.
func (o *CustomerCreateResponseData) SetMerchantRefId(v string) {
	o.MerchantRefId = &v
}

// GetEmail returns the Email field value
func (o *CustomerCreateResponseData) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CustomerCreateResponseData) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomerCreateResponseData) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CustomerCreateResponseData) SetPhone(v string) {
	o.Phone = &v
}

// GetAuthUserId returns the AuthUserId field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetAuthUserId() string {
	if o == nil || IsNil(o.AuthUserId) {
		var ret string
		return ret
	}
	return *o.AuthUserId
}

// GetAuthUserIdOk returns a tuple with the AuthUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetAuthUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUserId) {
		return nil, false
	}
	return o.AuthUserId, true
}

// HasAuthUserId returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasAuthUserId() bool {
	if o != nil && !IsNil(o.AuthUserId) {
		return true
	}

	return false
}

// SetAuthUserId gets a reference to the given string and assigns it to the AuthUserId field.
func (o *CustomerCreateResponseData) SetAuthUserId(v string) {
	o.AuthUserId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CustomerCreateResponseData) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *CustomerCreateResponseData) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CustomerCreateResponseData) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CustomerCreateResponseData) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerCreateResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerCreateResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomerCreateResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerCreateResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomerCreateResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CustomerCreateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCreateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchantId"] = o.MerchantId
	if !IsNil(o.MerchantRefId) {
		toSerialize["merchantRefId"] = o.MerchantRefId
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.AuthUserId) {
		toSerialize["authUserId"] = o.AuthUserId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerCreateResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchantId",
		"email",
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

	varCustomerCreateResponseData := _CustomerCreateResponseData{}

	err = json.Unmarshal(data, &varCustomerCreateResponseData)

	if err != nil {
		return err
	}

	*o = CustomerCreateResponseData(varCustomerCreateResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchantId")
		delete(additionalProperties, "merchantRefId")
		delete(additionalProperties, "email")
		delete(additionalProperties, "name")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "authUserId")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "customFields")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerCreateResponseData struct {
	value *CustomerCreateResponseData
	isSet bool
}

func (v NullableCustomerCreateResponseData) Get() *CustomerCreateResponseData {
	return v.value
}

func (v *NullableCustomerCreateResponseData) Set(val *CustomerCreateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateResponseData(val *CustomerCreateResponseData) *NullableCustomerCreateResponseData {
	return &NullableCustomerCreateResponseData{value: val, isSet: true}
}

func (v NullableCustomerCreateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


