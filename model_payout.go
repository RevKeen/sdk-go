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

// checks if the Payout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Payout{}

// Payout A payout represents funds settled from processed payments to your bank account. Track settlement batches and reconcile amounts.
type Payout struct {
	Id string `json:"id"`
	PublicId string `json:"public_id"`
	// Name of the payment processor that produced this payout batch
	Gateway string `json:"gateway"`
	GatewayPayoutId string `json:"gateway_payout_id"`
	// Gross amount in cents
	GrossAmountMinor float32 `json:"gross_amount_minor"`
	// Fees in cents
	FeesAmountMinor float32 `json:"fees_amount_minor"`
	// Net amount in cents
	NetAmountMinor float32 `json:"net_amount_minor"`
	Currency string `json:"currency"`
	ChargesCount NullableFloat32 `json:"charges_count"`
	RefundsCount NullableFloat32 `json:"refunds_count"`
	ChargebacksCount NullableFloat32 `json:"chargebacks_count"`
	Status string `json:"status"`
	ArrivalDate NullableString `json:"arrival_date"`
	SettledAt NullableString `json:"settled_at"`
	CreatedAt string `json:"created_at"`
	AdditionalProperties map[string]interface{}
}

type _Payout Payout

// NewPayout instantiates a new Payout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayout(id string, publicId string, gateway string, gatewayPayoutId string, grossAmountMinor float32, feesAmountMinor float32, netAmountMinor float32, currency string, chargesCount NullableFloat32, refundsCount NullableFloat32, chargebacksCount NullableFloat32, status string, arrivalDate NullableString, settledAt NullableString, createdAt string) *Payout {
	this := Payout{}
	this.Id = id
	this.PublicId = publicId
	this.Gateway = gateway
	this.GatewayPayoutId = gatewayPayoutId
	this.GrossAmountMinor = grossAmountMinor
	this.FeesAmountMinor = feesAmountMinor
	this.NetAmountMinor = netAmountMinor
	this.Currency = currency
	this.ChargesCount = chargesCount
	this.RefundsCount = refundsCount
	this.ChargebacksCount = chargebacksCount
	this.Status = status
	this.ArrivalDate = arrivalDate
	this.SettledAt = settledAt
	this.CreatedAt = createdAt
	return &this
}

// NewPayoutWithDefaults instantiates a new Payout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutWithDefaults() *Payout {
	this := Payout{}
	return &this
}

// GetId returns the Id field value
func (o *Payout) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payout) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payout) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
func (o *Payout) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *Payout) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *Payout) SetPublicId(v string) {
	o.PublicId = v
}

// GetGateway returns the Gateway field value
func (o *Payout) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *Payout) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *Payout) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayPayoutId returns the GatewayPayoutId field value
func (o *Payout) GetGatewayPayoutId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayPayoutId
}

// GetGatewayPayoutIdOk returns a tuple with the GatewayPayoutId field value
// and a boolean to check if the value has been set.
func (o *Payout) GetGatewayPayoutIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayPayoutId, true
}

// SetGatewayPayoutId sets field value
func (o *Payout) SetGatewayPayoutId(v string) {
	o.GatewayPayoutId = v
}

// GetGrossAmountMinor returns the GrossAmountMinor field value
func (o *Payout) GetGrossAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GrossAmountMinor
}

// GetGrossAmountMinorOk returns a tuple with the GrossAmountMinor field value
// and a boolean to check if the value has been set.
func (o *Payout) GetGrossAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrossAmountMinor, true
}

// SetGrossAmountMinor sets field value
func (o *Payout) SetGrossAmountMinor(v float32) {
	o.GrossAmountMinor = v
}

// GetFeesAmountMinor returns the FeesAmountMinor field value
func (o *Payout) GetFeesAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.FeesAmountMinor
}

// GetFeesAmountMinorOk returns a tuple with the FeesAmountMinor field value
// and a boolean to check if the value has been set.
func (o *Payout) GetFeesAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeesAmountMinor, true
}

// SetFeesAmountMinor sets field value
func (o *Payout) SetFeesAmountMinor(v float32) {
	o.FeesAmountMinor = v
}

// GetNetAmountMinor returns the NetAmountMinor field value
func (o *Payout) GetNetAmountMinor() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetAmountMinor
}

// GetNetAmountMinorOk returns a tuple with the NetAmountMinor field value
// and a boolean to check if the value has been set.
func (o *Payout) GetNetAmountMinorOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetAmountMinor, true
}

// SetNetAmountMinor sets field value
func (o *Payout) SetNetAmountMinor(v float32) {
	o.NetAmountMinor = v
}

// GetCurrency returns the Currency field value
func (o *Payout) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Payout) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Payout) SetCurrency(v string) {
	o.Currency = v
}

// GetChargesCount returns the ChargesCount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Payout) GetChargesCount() float32 {
	if o == nil || o.ChargesCount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ChargesCount.Get()
}

// GetChargesCountOk returns a tuple with the ChargesCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payout) GetChargesCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargesCount.Get(), o.ChargesCount.IsSet()
}

// SetChargesCount sets field value
func (o *Payout) SetChargesCount(v float32) {
	o.ChargesCount.Set(&v)
}

// GetRefundsCount returns the RefundsCount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Payout) GetRefundsCount() float32 {
	if o == nil || o.RefundsCount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.RefundsCount.Get()
}

// GetRefundsCountOk returns a tuple with the RefundsCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payout) GetRefundsCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefundsCount.Get(), o.RefundsCount.IsSet()
}

// SetRefundsCount sets field value
func (o *Payout) SetRefundsCount(v float32) {
	o.RefundsCount.Set(&v)
}

// GetChargebacksCount returns the ChargebacksCount field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Payout) GetChargebacksCount() float32 {
	if o == nil || o.ChargebacksCount.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ChargebacksCount.Get()
}

// GetChargebacksCountOk returns a tuple with the ChargebacksCount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payout) GetChargebacksCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChargebacksCount.Get(), o.ChargebacksCount.IsSet()
}

// SetChargebacksCount sets field value
func (o *Payout) SetChargebacksCount(v float32) {
	o.ChargebacksCount.Set(&v)
}

// GetStatus returns the Status field value
func (o *Payout) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payout) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payout) SetStatus(v string) {
	o.Status = v
}

// GetArrivalDate returns the ArrivalDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payout) GetArrivalDate() string {
	if o == nil || o.ArrivalDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.ArrivalDate.Get()
}

// GetArrivalDateOk returns a tuple with the ArrivalDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payout) GetArrivalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArrivalDate.Get(), o.ArrivalDate.IsSet()
}

// SetArrivalDate sets field value
func (o *Payout) SetArrivalDate(v string) {
	o.ArrivalDate.Set(&v)
}

// GetSettledAt returns the SettledAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Payout) GetSettledAt() string {
	if o == nil || o.SettledAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.SettledAt.Get()
}

// GetSettledAtOk returns a tuple with the SettledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Payout) GetSettledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettledAt.Get(), o.SettledAt.IsSet()
}

// SetSettledAt sets field value
func (o *Payout) SetSettledAt(v string) {
	o.SettledAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Payout) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Payout) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Payout) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o Payout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Payout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId
	toSerialize["gateway"] = o.Gateway
	toSerialize["gateway_payout_id"] = o.GatewayPayoutId
	toSerialize["gross_amount_minor"] = o.GrossAmountMinor
	toSerialize["fees_amount_minor"] = o.FeesAmountMinor
	toSerialize["net_amount_minor"] = o.NetAmountMinor
	toSerialize["currency"] = o.Currency
	toSerialize["charges_count"] = o.ChargesCount.Get()
	toSerialize["refunds_count"] = o.RefundsCount.Get()
	toSerialize["chargebacks_count"] = o.ChargebacksCount.Get()
	toSerialize["status"] = o.Status
	toSerialize["arrival_date"] = o.ArrivalDate.Get()
	toSerialize["settled_at"] = o.SettledAt.Get()
	toSerialize["created_at"] = o.CreatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Payout) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"public_id",
		"gateway",
		"gateway_payout_id",
		"gross_amount_minor",
		"fees_amount_minor",
		"net_amount_minor",
		"currency",
		"charges_count",
		"refunds_count",
		"chargebacks_count",
		"status",
		"arrival_date",
		"settled_at",
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

	varPayout := _Payout{}

	err = json.Unmarshal(data, &varPayout)

	if err != nil {
		return err
	}

	*o = Payout(varPayout)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_payout_id")
		delete(additionalProperties, "gross_amount_minor")
		delete(additionalProperties, "fees_amount_minor")
		delete(additionalProperties, "net_amount_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "charges_count")
		delete(additionalProperties, "refunds_count")
		delete(additionalProperties, "chargebacks_count")
		delete(additionalProperties, "status")
		delete(additionalProperties, "arrival_date")
		delete(additionalProperties, "settled_at")
		delete(additionalProperties, "created_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePayout struct {
	value *Payout
	isSet bool
}

func (v NullablePayout) Get() *Payout {
	return v.value
}

func (v *NullablePayout) Set(val *Payout) {
	v.value = val
	v.isSet = true
}

func (v NullablePayout) IsSet() bool {
	return v.isSet
}

func (v *NullablePayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayout(val *Payout) *NullablePayout {
	return &NullablePayout{value: val, isSet: true}
}

func (v NullablePayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


