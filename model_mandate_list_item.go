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

// checks if the MandateListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MandateListItem{}

// MandateListItem Summary representation of a Direct Debit mandate in a list.
type MandateListItem struct {
	Id string `json:"id"`
	MandateRef string `json:"mandate_ref"`
	MandateReference NullableString `json:"mandate_reference"`
	CustomerId string `json:"customer_id"`
	Status string `json:"status"`
	AccountHolderName NullableString `json:"account_holder_name"`
	SortCode string `json:"sort_code"`
	AccountNumberLast4 NullableString `json:"account_number_last4"`
	BankName NullableString `json:"bank_name"`
	NoticeDays NullableInt32 `json:"notice_days"`
	ActivatedAt NullableString `json:"activated_at"`
	CreatedAt NullableString `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _MandateListItem MandateListItem

// NewMandateListItem instantiates a new MandateListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandateListItem(id string, mandateRef string, mandateReference NullableString, customerId string, status string, accountHolderName NullableString, sortCode string, accountNumberLast4 NullableString, bankName NullableString, noticeDays NullableInt32, activatedAt NullableString, createdAt NullableString) *MandateListItem {
	this := MandateListItem{}
	this.Id = id
	this.MandateRef = mandateRef
	this.MandateReference = mandateReference
	this.CustomerId = customerId
	this.Status = status
	this.AccountHolderName = accountHolderName
	this.SortCode = sortCode
	this.AccountNumberLast4 = accountNumberLast4
	this.BankName = bankName
	this.NoticeDays = noticeDays
	this.ActivatedAt = activatedAt
	this.CreatedAt = createdAt
	return &this
}

// NewMandateListItemWithDefaults instantiates a new MandateListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateListItemWithDefaults() *MandateListItem {
	this := MandateListItem{}
	return &this
}

// GetId returns the Id field value
func (o *MandateListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MandateListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MandateListItem) SetId(v string) {
	o.Id = v
}

// GetMandateRef returns the MandateRef field value
func (o *MandateListItem) GetMandateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateRef
}

// GetMandateRefOk returns a tuple with the MandateRef field value
// and a boolean to check if the value has been set.
func (o *MandateListItem) GetMandateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateRef, true
}

// SetMandateRef sets field value
func (o *MandateListItem) SetMandateRef(v string) {
	o.MandateRef = v
}

// GetMandateReference returns the MandateReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetMandateReference() string {
	if o == nil || o.MandateReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateReference.Get()
}

// GetMandateReferenceOk returns a tuple with the MandateReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetMandateReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateReference.Get(), o.MandateReference.IsSet()
}

// SetMandateReference sets field value
func (o *MandateListItem) SetMandateReference(v string) {
	o.MandateReference.Set(&v)
}

// GetCustomerId returns the CustomerId field value
func (o *MandateListItem) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *MandateListItem) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *MandateListItem) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetStatus returns the Status field value
func (o *MandateListItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MandateListItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MandateListItem) SetStatus(v string) {
	o.Status = v
}

// GetAccountHolderName returns the AccountHolderName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetAccountHolderName() string {
	if o == nil || o.AccountHolderName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// SetAccountHolderName sets field value
func (o *MandateListItem) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}

// GetSortCode returns the SortCode field value
func (o *MandateListItem) GetSortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *MandateListItem) GetSortCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *MandateListItem) SetSortCode(v string) {
	o.SortCode = v
}

// GetAccountNumberLast4 returns the AccountNumberLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetAccountNumberLast4() string {
	if o == nil || o.AccountNumberLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountNumberLast4.Get()
}

// GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetAccountNumberLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumberLast4.Get(), o.AccountNumberLast4.IsSet()
}

// SetAccountNumberLast4 sets field value
func (o *MandateListItem) SetAccountNumberLast4(v string) {
	o.AccountNumberLast4.Set(&v)
}

// GetBankName returns the BankName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// SetBankName sets field value
func (o *MandateListItem) SetBankName(v string) {
	o.BankName.Set(&v)
}

// GetNoticeDays returns the NoticeDays field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *MandateListItem) GetNoticeDays() int32 {
	if o == nil || o.NoticeDays.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NoticeDays.Get()
}

// GetNoticeDaysOk returns a tuple with the NoticeDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetNoticeDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoticeDays.Get(), o.NoticeDays.IsSet()
}

// SetNoticeDays sets field value
func (o *MandateListItem) SetNoticeDays(v int32) {
	o.NoticeDays.Set(&v)
}

// GetActivatedAt returns the ActivatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetActivatedAt() string {
	if o == nil || o.ActivatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ActivatedAt.Get()
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetActivatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivatedAt.Get(), o.ActivatedAt.IsSet()
}

// SetActivatedAt sets field value
func (o *MandateListItem) SetActivatedAt(v string) {
	o.ActivatedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MandateListItem) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MandateListItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *MandateListItem) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

func (o MandateListItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MandateListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["mandate_ref"] = o.MandateRef
	toSerialize["mandate_reference"] = o.MandateReference.Get()
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["status"] = o.Status
	toSerialize["account_holder_name"] = o.AccountHolderName.Get()
	toSerialize["sort_code"] = o.SortCode
	toSerialize["account_number_last4"] = o.AccountNumberLast4.Get()
	toSerialize["bank_name"] = o.BankName.Get()
	toSerialize["notice_days"] = o.NoticeDays.Get()
	toSerialize["activated_at"] = o.ActivatedAt.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MandateListItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"mandate_ref",
		"mandate_reference",
		"customer_id",
		"status",
		"account_holder_name",
		"sort_code",
		"account_number_last4",
		"bank_name",
		"notice_days",
		"activated_at",
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

	varMandateListItem := _MandateListItem{}

	err = json.Unmarshal(data, &varMandateListItem)

	if err != nil {
		return err
	}

	*o = MandateListItem(varMandateListItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mandate_ref")
		delete(additionalProperties, "mandate_reference")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "account_holder_name")
		delete(additionalProperties, "sort_code")
		delete(additionalProperties, "account_number_last4")
		delete(additionalProperties, "bank_name")
		delete(additionalProperties, "notice_days")
		delete(additionalProperties, "activated_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMandateListItem struct {
	value *MandateListItem
	isSet bool
}

func (v NullableMandateListItem) Get() *MandateListItem {
	return v.value
}

func (v *NullableMandateListItem) Set(val *MandateListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMandateListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMandateListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandateListItem(val *MandateListItem) *NullableMandateListItem {
	return &NullableMandateListItem{value: val, isSet: true}
}

func (v NullableMandateListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandateListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


