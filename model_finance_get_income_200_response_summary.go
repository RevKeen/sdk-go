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
	"fmt"
)

// checks if the FinanceGetIncome200ResponseSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinanceGetIncome200ResponseSummary{}

// FinanceGetIncome200ResponseSummary struct for FinanceGetIncome200ResponseSummary
type FinanceGetIncome200ResponseSummary struct {
	TotalGrossCents float32 `json:"total_gross_cents"`
	TotalFeesCents float32 `json:"total_fees_cents"`
	TotalNetCents float32 `json:"total_net_cents"`
	TotalTxCount float32 `json:"total_tx_count"`
	TotalRefundCount float32 `json:"total_refund_count"`
	AdditionalProperties map[string]interface{}
}

type _FinanceGetIncome200ResponseSummary FinanceGetIncome200ResponseSummary

// NewFinanceGetIncome200ResponseSummary instantiates a new FinanceGetIncome200ResponseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinanceGetIncome200ResponseSummary(totalGrossCents float32, totalFeesCents float32, totalNetCents float32, totalTxCount float32, totalRefundCount float32) *FinanceGetIncome200ResponseSummary {
	this := FinanceGetIncome200ResponseSummary{}
	this.TotalGrossCents = totalGrossCents
	this.TotalFeesCents = totalFeesCents
	this.TotalNetCents = totalNetCents
	this.TotalTxCount = totalTxCount
	this.TotalRefundCount = totalRefundCount
	return &this
}

// NewFinanceGetIncome200ResponseSummaryWithDefaults instantiates a new FinanceGetIncome200ResponseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinanceGetIncome200ResponseSummaryWithDefaults() *FinanceGetIncome200ResponseSummary {
	this := FinanceGetIncome200ResponseSummary{}
	return &this
}

// GetTotalGrossCents returns the TotalGrossCents field value
func (o *FinanceGetIncome200ResponseSummary) GetTotalGrossCents() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalGrossCents
}

// GetTotalGrossCentsOk returns a tuple with the TotalGrossCents field value
// and a boolean to check if the value has been set.
func (o *FinanceGetIncome200ResponseSummary) GetTotalGrossCentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalGrossCents, true
}

// SetTotalGrossCents sets field value
func (o *FinanceGetIncome200ResponseSummary) SetTotalGrossCents(v float32) {
	o.TotalGrossCents = v
}

// GetTotalFeesCents returns the TotalFeesCents field value
func (o *FinanceGetIncome200ResponseSummary) GetTotalFeesCents() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalFeesCents
}

// GetTotalFeesCentsOk returns a tuple with the TotalFeesCents field value
// and a boolean to check if the value has been set.
func (o *FinanceGetIncome200ResponseSummary) GetTotalFeesCentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalFeesCents, true
}

// SetTotalFeesCents sets field value
func (o *FinanceGetIncome200ResponseSummary) SetTotalFeesCents(v float32) {
	o.TotalFeesCents = v
}

// GetTotalNetCents returns the TotalNetCents field value
func (o *FinanceGetIncome200ResponseSummary) GetTotalNetCents() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalNetCents
}

// GetTotalNetCentsOk returns a tuple with the TotalNetCents field value
// and a boolean to check if the value has been set.
func (o *FinanceGetIncome200ResponseSummary) GetTotalNetCentsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalNetCents, true
}

// SetTotalNetCents sets field value
func (o *FinanceGetIncome200ResponseSummary) SetTotalNetCents(v float32) {
	o.TotalNetCents = v
}

// GetTotalTxCount returns the TotalTxCount field value
func (o *FinanceGetIncome200ResponseSummary) GetTotalTxCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalTxCount
}

// GetTotalTxCountOk returns a tuple with the TotalTxCount field value
// and a boolean to check if the value has been set.
func (o *FinanceGetIncome200ResponseSummary) GetTotalTxCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTxCount, true
}

// SetTotalTxCount sets field value
func (o *FinanceGetIncome200ResponseSummary) SetTotalTxCount(v float32) {
	o.TotalTxCount = v
}

// GetTotalRefundCount returns the TotalRefundCount field value
func (o *FinanceGetIncome200ResponseSummary) GetTotalRefundCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalRefundCount
}

// GetTotalRefundCountOk returns a tuple with the TotalRefundCount field value
// and a boolean to check if the value has been set.
func (o *FinanceGetIncome200ResponseSummary) GetTotalRefundCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRefundCount, true
}

// SetTotalRefundCount sets field value
func (o *FinanceGetIncome200ResponseSummary) SetTotalRefundCount(v float32) {
	o.TotalRefundCount = v
}

func (o FinanceGetIncome200ResponseSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinanceGetIncome200ResponseSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total_gross_cents"] = o.TotalGrossCents
	toSerialize["total_fees_cents"] = o.TotalFeesCents
	toSerialize["total_net_cents"] = o.TotalNetCents
	toSerialize["total_tx_count"] = o.TotalTxCount
	toSerialize["total_refund_count"] = o.TotalRefundCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FinanceGetIncome200ResponseSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total_gross_cents",
		"total_fees_cents",
		"total_net_cents",
		"total_tx_count",
		"total_refund_count",
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

	varFinanceGetIncome200ResponseSummary := _FinanceGetIncome200ResponseSummary{}

	err = json.Unmarshal(data, &varFinanceGetIncome200ResponseSummary)

	if err != nil {
		return err
	}

	*o = FinanceGetIncome200ResponseSummary(varFinanceGetIncome200ResponseSummary)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total_gross_cents")
		delete(additionalProperties, "total_fees_cents")
		delete(additionalProperties, "total_net_cents")
		delete(additionalProperties, "total_tx_count")
		delete(additionalProperties, "total_refund_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFinanceGetIncome200ResponseSummary struct {
	value *FinanceGetIncome200ResponseSummary
	isSet bool
}

func (v NullableFinanceGetIncome200ResponseSummary) Get() *FinanceGetIncome200ResponseSummary {
	return v.value
}

func (v *NullableFinanceGetIncome200ResponseSummary) Set(val *FinanceGetIncome200ResponseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableFinanceGetIncome200ResponseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableFinanceGetIncome200ResponseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinanceGetIncome200ResponseSummary(val *FinanceGetIncome200ResponseSummary) *NullableFinanceGetIncome200ResponseSummary {
	return &NullableFinanceGetIncome200ResponseSummary{value: val, isSet: true}
}

func (v NullableFinanceGetIncome200ResponseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinanceGetIncome200ResponseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


