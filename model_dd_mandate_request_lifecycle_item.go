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

// checks if the DdMandateRequestLifecycleItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdMandateRequestLifecycleItem{}

// DdMandateRequestLifecycleItem struct for DdMandateRequestLifecycleItem
type DdMandateRequestLifecycleItem struct {
	Id string `json:"id"`
	MerchantId string `json:"merchant_id"`
	CustomerId string `json:"customer_id"`
	InvoiceId NullableString `json:"invoice_id"`
	SubscriptionId NullableString `json:"subscription_id"`
	MandateId NullableString `json:"mandate_id"`
	// Internal raw request token. Do not expose directly to customers.
	Token string `json:"token"`
	// Customer-facing tracked URL: pay.revkeen.com/dd_<token>
	SignableUrl string `json:"signable_url"`
	CreatedAt NullableString `json:"created_at"`
	ExpiresAt NullableString `json:"expires_at"`
	ConsumedAt NullableString `json:"consumed_at"`
	CancelledAt NullableString `json:"cancelled_at"`
	ExpiredAt NullableString `json:"expired_at"`
	Status DdMandateRequestLifecycleStatus `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DdMandateRequestLifecycleItem DdMandateRequestLifecycleItem

// NewDdMandateRequestLifecycleItem instantiates a new DdMandateRequestLifecycleItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdMandateRequestLifecycleItem(id string, merchantId string, customerId string, invoiceId NullableString, subscriptionId NullableString, mandateId NullableString, token string, signableUrl string, createdAt NullableString, expiresAt NullableString, consumedAt NullableString, cancelledAt NullableString, expiredAt NullableString, status DdMandateRequestLifecycleStatus) *DdMandateRequestLifecycleItem {
	this := DdMandateRequestLifecycleItem{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.InvoiceId = invoiceId
	this.SubscriptionId = subscriptionId
	this.MandateId = mandateId
	this.Token = token
	this.SignableUrl = signableUrl
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.ConsumedAt = consumedAt
	this.CancelledAt = cancelledAt
	this.ExpiredAt = expiredAt
	this.Status = status
	return &this
}

// NewDdMandateRequestLifecycleItemWithDefaults instantiates a new DdMandateRequestLifecycleItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdMandateRequestLifecycleItemWithDefaults() *DdMandateRequestLifecycleItem {
	this := DdMandateRequestLifecycleItem{}
	return &this
}

// GetId returns the Id field value
func (o *DdMandateRequestLifecycleItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DdMandateRequestLifecycleItem) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *DdMandateRequestLifecycleItem) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *DdMandateRequestLifecycleItem) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *DdMandateRequestLifecycleItem) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DdMandateRequestLifecycleItem) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *DdMandateRequestLifecycleItem) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *DdMandateRequestLifecycleItem) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetMandateId returns the MandateId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetMandateId() string {
	if o == nil || o.MandateId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateId.Get()
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateId.Get(), o.MandateId.IsSet()
}

// SetMandateId sets field value
func (o *DdMandateRequestLifecycleItem) SetMandateId(v string) {
	o.MandateId.Set(&v)
}

// GetToken returns the Token field value
func (o *DdMandateRequestLifecycleItem) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DdMandateRequestLifecycleItem) SetToken(v string) {
	o.Token = v
}

// GetSignableUrl returns the SignableUrl field value
func (o *DdMandateRequestLifecycleItem) GetSignableUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableUrl
}

// GetSignableUrlOk returns a tuple with the SignableUrl field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetSignableUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignableUrl, true
}

// SetSignableUrl sets field value
func (o *DdMandateRequestLifecycleItem) SetSignableUrl(v string) {
	o.SignableUrl = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *DdMandateRequestLifecycleItem) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *DdMandateRequestLifecycleItem) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetConsumedAt returns the ConsumedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetConsumedAt() string {
	if o == nil || o.ConsumedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ConsumedAt.Get()
}

// GetConsumedAtOk returns a tuple with the ConsumedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetConsumedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumedAt.Get(), o.ConsumedAt.IsSet()
}

// SetConsumedAt sets field value
func (o *DdMandateRequestLifecycleItem) SetConsumedAt(v string) {
	o.ConsumedAt.Set(&v)
}

// GetCancelledAt returns the CancelledAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetCancelledAt() string {
	if o == nil || o.CancelledAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CancelledAt.Get()
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetCancelledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledAt.Get(), o.CancelledAt.IsSet()
}

// SetCancelledAt sets field value
func (o *DdMandateRequestLifecycleItem) SetCancelledAt(v string) {
	o.CancelledAt.Set(&v)
}

// GetExpiredAt returns the ExpiredAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandateRequestLifecycleItem) GetExpiredAt() string {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestLifecycleItem) GetExpiredAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// SetExpiredAt sets field value
func (o *DdMandateRequestLifecycleItem) SetExpiredAt(v string) {
	o.ExpiredAt.Set(&v)
}

// GetStatus returns the Status field value
func (o *DdMandateRequestLifecycleItem) GetStatus() DdMandateRequestLifecycleStatus {
	if o == nil {
		var ret DdMandateRequestLifecycleStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestLifecycleItem) GetStatusOk() (*DdMandateRequestLifecycleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DdMandateRequestLifecycleItem) SetStatus(v DdMandateRequestLifecycleStatus) {
	o.Status = v
}

func (o DdMandateRequestLifecycleItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdMandateRequestLifecycleItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	toSerialize["mandate_id"] = o.MandateId.Get()
	toSerialize["token"] = o.Token
	toSerialize["signable_url"] = o.SignableUrl
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["consumed_at"] = o.ConsumedAt.Get()
	toSerialize["cancelled_at"] = o.CancelledAt.Get()
	toSerialize["expired_at"] = o.ExpiredAt.Get()
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdMandateRequestLifecycleItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"customer_id",
		"invoice_id",
		"subscription_id",
		"mandate_id",
		"token",
		"signable_url",
		"created_at",
		"expires_at",
		"consumed_at",
		"cancelled_at",
		"expired_at",
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

	varDdMandateRequestLifecycleItem := _DdMandateRequestLifecycleItem{}

	err = json.Unmarshal(data, &varDdMandateRequestLifecycleItem)

	if err != nil {
		return err
	}

	*o = DdMandateRequestLifecycleItem(varDdMandateRequestLifecycleItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "token")
		delete(additionalProperties, "signable_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "consumed_at")
		delete(additionalProperties, "cancelled_at")
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdMandateRequestLifecycleItem struct {
	value *DdMandateRequestLifecycleItem
	isSet bool
}

func (v NullableDdMandateRequestLifecycleItem) Get() *DdMandateRequestLifecycleItem {
	return v.value
}

func (v *NullableDdMandateRequestLifecycleItem) Set(val *DdMandateRequestLifecycleItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDdMandateRequestLifecycleItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDdMandateRequestLifecycleItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdMandateRequestLifecycleItem(val *DdMandateRequestLifecycleItem) *NullableDdMandateRequestLifecycleItem {
	return &NullableDdMandateRequestLifecycleItem{value: val, isSet: true}
}

func (v NullableDdMandateRequestLifecycleItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdMandateRequestLifecycleItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


