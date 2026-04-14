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

// checks if the SubscriptionCreateResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateResponseData{}

// SubscriptionCreateResponseData struct for SubscriptionCreateResponseData
type SubscriptionCreateResponseData struct {
	Id string `json:"id"`
	MerchantId string `json:"merchantId"`
	CustomerId string `json:"customerId"`
	ProductId *string `json:"productId,omitempty"`
	PriceId *string `json:"priceId,omitempty"`
	PlanId *string `json:"planId,omitempty"`
	Status string `json:"status"`
	Quantity int32 `json:"quantity"`
	AmountMinor int32 `json:"amountMinor"`
	Currency string `json:"currency"`
	BillingInterval string `json:"billingInterval"`
	CurrentPeriodStart time.Time `json:"currentPeriodStart"`
	CurrentPeriodEnd time.Time `json:"currentPeriodEnd"`
	TrialEnd *time.Time `json:"trialEnd,omitempty"`
	CanceledAt *time.Time `json:"canceledAt,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Dunning *SubscriptionCreateResponseDataDunning `json:"dunning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionCreateResponseData SubscriptionCreateResponseData

// NewSubscriptionCreateResponseData instantiates a new SubscriptionCreateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateResponseData(id string, merchantId string, customerId string, status string, quantity int32, amountMinor int32, currency string, billingInterval string, currentPeriodStart time.Time, currentPeriodEnd time.Time, createdAt time.Time, updatedAt time.Time) *SubscriptionCreateResponseData {
	this := SubscriptionCreateResponseData{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.Status = status
	this.Quantity = quantity
	this.AmountMinor = amountMinor
	this.Currency = currency
	this.BillingInterval = billingInterval
	this.CurrentPeriodStart = currentPeriodStart
	this.CurrentPeriodEnd = currentPeriodEnd
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSubscriptionCreateResponseDataWithDefaults instantiates a new SubscriptionCreateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateResponseDataWithDefaults() *SubscriptionCreateResponseData {
	this := SubscriptionCreateResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionCreateResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionCreateResponseData) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *SubscriptionCreateResponseData) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *SubscriptionCreateResponseData) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *SubscriptionCreateResponseData) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *SubscriptionCreateResponseData) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *SubscriptionCreateResponseData) SetProductId(v string) {
	o.ProductId = &v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetPriceId() string {
	if o == nil || IsNil(o.PriceId) {
		var ret string
		return ret
	}
	return *o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetPriceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PriceId) {
		return nil, false
	}
	return o.PriceId, true
}

// HasPriceId returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasPriceId() bool {
	if o != nil && !IsNil(o.PriceId) {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given string and assigns it to the PriceId field.
func (o *SubscriptionCreateResponseData) SetPriceId(v string) {
	o.PriceId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionCreateResponseData) SetPlanId(v string) {
	o.PlanId = &v
}

// GetStatus returns the Status field value
func (o *SubscriptionCreateResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionCreateResponseData) SetStatus(v string) {
	o.Status = v
}

// GetQuantity returns the Quantity field value
func (o *SubscriptionCreateResponseData) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SubscriptionCreateResponseData) SetQuantity(v int32) {
	o.Quantity = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *SubscriptionCreateResponseData) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *SubscriptionCreateResponseData) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *SubscriptionCreateResponseData) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *SubscriptionCreateResponseData) SetCurrency(v string) {
	o.Currency = v
}

// GetBillingInterval returns the BillingInterval field value
func (o *SubscriptionCreateResponseData) GetBillingInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingInterval
}

// GetBillingIntervalOk returns a tuple with the BillingInterval field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetBillingIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingInterval, true
}

// SetBillingInterval sets field value
func (o *SubscriptionCreateResponseData) SetBillingInterval(v string) {
	o.BillingInterval = v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
func (o *SubscriptionCreateResponseData) GetCurrentPeriodStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodStart
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCurrentPeriodStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodStart, true
}

// SetCurrentPeriodStart sets field value
func (o *SubscriptionCreateResponseData) SetCurrentPeriodStart(v time.Time) {
	o.CurrentPeriodStart = v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
func (o *SubscriptionCreateResponseData) GetCurrentPeriodEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCurrentPeriodEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodEnd, true
}

// SetCurrentPeriodEnd sets field value
func (o *SubscriptionCreateResponseData) SetCurrentPeriodEnd(v time.Time) {
	o.CurrentPeriodEnd = v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetTrialEnd() time.Time {
	if o == nil || IsNil(o.TrialEnd) {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetTrialEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given time.Time and assigns it to the TrialEnd field.
func (o *SubscriptionCreateResponseData) SetTrialEnd(v time.Time) {
	o.TrialEnd = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CanceledAt) {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasCanceledAt() bool {
	if o != nil && !IsNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *SubscriptionCreateResponseData) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SubscriptionCreateResponseData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SubscriptionCreateResponseData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SubscriptionCreateResponseData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SubscriptionCreateResponseData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDunning returns the Dunning field value if set, zero value otherwise.
func (o *SubscriptionCreateResponseData) GetDunning() SubscriptionCreateResponseDataDunning {
	if o == nil || IsNil(o.Dunning) {
		var ret SubscriptionCreateResponseDataDunning
		return ret
	}
	return *o.Dunning
}

// GetDunningOk returns a tuple with the Dunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateResponseData) GetDunningOk() (*SubscriptionCreateResponseDataDunning, bool) {
	if o == nil || IsNil(o.Dunning) {
		return nil, false
	}
	return o.Dunning, true
}

// HasDunning returns a boolean if a field has been set.
func (o *SubscriptionCreateResponseData) HasDunning() bool {
	if o != nil && !IsNil(o.Dunning) {
		return true
	}

	return false
}

// SetDunning gets a reference to the given SubscriptionCreateResponseDataDunning and assigns it to the Dunning field.
func (o *SubscriptionCreateResponseData) SetDunning(v SubscriptionCreateResponseDataDunning) {
	o.Dunning = &v
}

func (o SubscriptionCreateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchantId"] = o.MerchantId
	toSerialize["customerId"] = o.CustomerId
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.PriceId) {
		toSerialize["priceId"] = o.PriceId
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	toSerialize["status"] = o.Status
	toSerialize["quantity"] = o.Quantity
	toSerialize["amountMinor"] = o.AmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["billingInterval"] = o.BillingInterval
	toSerialize["currentPeriodStart"] = o.CurrentPeriodStart
	toSerialize["currentPeriodEnd"] = o.CurrentPeriodEnd
	if !IsNil(o.TrialEnd) {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	if !IsNil(o.CanceledAt) {
		toSerialize["canceledAt"] = o.CanceledAt
	}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	if !IsNil(o.Dunning) {
		toSerialize["dunning"] = o.Dunning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionCreateResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchantId",
		"customerId",
		"status",
		"quantity",
		"amountMinor",
		"currency",
		"billingInterval",
		"currentPeriodStart",
		"currentPeriodEnd",
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

	varSubscriptionCreateResponseData := _SubscriptionCreateResponseData{}

	err = json.Unmarshal(data, &varSubscriptionCreateResponseData)

	if err != nil {
		return err
	}

	*o = SubscriptionCreateResponseData(varSubscriptionCreateResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchantId")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "priceId")
		delete(additionalProperties, "planId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "amountMinor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "billingInterval")
		delete(additionalProperties, "currentPeriodStart")
		delete(additionalProperties, "currentPeriodEnd")
		delete(additionalProperties, "trialEnd")
		delete(additionalProperties, "canceledAt")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "dunning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionCreateResponseData struct {
	value *SubscriptionCreateResponseData
	isSet bool
}

func (v NullableSubscriptionCreateResponseData) Get() *SubscriptionCreateResponseData {
	return v.value
}

func (v *NullableSubscriptionCreateResponseData) Set(val *SubscriptionCreateResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateResponseData(val *SubscriptionCreateResponseData) *NullableSubscriptionCreateResponseData {
	return &NullableSubscriptionCreateResponseData{value: val, isSet: true}
}

func (v NullableSubscriptionCreateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


