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

// checks if the InvoicesMarginEstimate200ResponseDataEstimatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicesMarginEstimate200ResponseDataEstimatesInner{}

// InvoicesMarginEstimate200ResponseDataEstimatesInner struct for InvoicesMarginEstimate200ResponseDataEstimatesInner
type InvoicesMarginEstimate200ResponseDataEstimatesInner struct {
	Rail string `json:"rail"`
	CardScheme NullableString `json:"card_scheme"`
	GrossMinor int32 `json:"gross_minor"`
	EstimatedFeeMinor int32 `json:"estimated_fee_minor"`
	NetAfterFeesMinor int32 `json:"net_after_fees_minor"`
	FeeBreakdown InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown `json:"fee_breakdown"`
	PricingSource string `json:"pricing_source"`
	MatrixVersion string `json:"matrix_version"`
	AdditionalProperties map[string]interface{}
}

type _InvoicesMarginEstimate200ResponseDataEstimatesInner InvoicesMarginEstimate200ResponseDataEstimatesInner

// NewInvoicesMarginEstimate200ResponseDataEstimatesInner instantiates a new InvoicesMarginEstimate200ResponseDataEstimatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicesMarginEstimate200ResponseDataEstimatesInner(rail string, cardScheme NullableString, grossMinor int32, estimatedFeeMinor int32, netAfterFeesMinor int32, feeBreakdown InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown, pricingSource string, matrixVersion string) *InvoicesMarginEstimate200ResponseDataEstimatesInner {
	this := InvoicesMarginEstimate200ResponseDataEstimatesInner{}
	this.Rail = rail
	this.CardScheme = cardScheme
	this.GrossMinor = grossMinor
	this.EstimatedFeeMinor = estimatedFeeMinor
	this.NetAfterFeesMinor = netAfterFeesMinor
	this.FeeBreakdown = feeBreakdown
	this.PricingSource = pricingSource
	this.MatrixVersion = matrixVersion
	return &this
}

// NewInvoicesMarginEstimate200ResponseDataEstimatesInnerWithDefaults instantiates a new InvoicesMarginEstimate200ResponseDataEstimatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesMarginEstimate200ResponseDataEstimatesInnerWithDefaults() *InvoicesMarginEstimate200ResponseDataEstimatesInner {
	this := InvoicesMarginEstimate200ResponseDataEstimatesInner{}
	return &this
}

// GetRail returns the Rail field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetRail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rail
}

// GetRailOk returns a tuple with the Rail field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetRailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rail, true
}

// SetRail sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetRail(v string) {
	o.Rail = v
}

// GetCardScheme returns the CardScheme field value
// If the value is explicit nil, the zero value for string will be returned
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetCardScheme() string {
	if o == nil || o.CardScheme.Get() == nil {
		var ret string
		return ret
	}

	return *o.CardScheme.Get()
}

// GetCardSchemeOk returns a tuple with the CardScheme field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetCardSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CardScheme.Get(), o.CardScheme.IsSet()
}

// SetCardScheme sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetCardScheme(v string) {
	o.CardScheme.Set(&v)
}

// GetGrossMinor returns the GrossMinor field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetGrossMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GrossMinor
}

// GetGrossMinorOk returns a tuple with the GrossMinor field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetGrossMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrossMinor, true
}

// SetGrossMinor sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetGrossMinor(v int32) {
	o.GrossMinor = v
}

// GetEstimatedFeeMinor returns the EstimatedFeeMinor field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetEstimatedFeeMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedFeeMinor
}

// GetEstimatedFeeMinorOk returns a tuple with the EstimatedFeeMinor field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetEstimatedFeeMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedFeeMinor, true
}

// SetEstimatedFeeMinor sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetEstimatedFeeMinor(v int32) {
	o.EstimatedFeeMinor = v
}

// GetNetAfterFeesMinor returns the NetAfterFeesMinor field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetNetAfterFeesMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NetAfterFeesMinor
}

// GetNetAfterFeesMinorOk returns a tuple with the NetAfterFeesMinor field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetNetAfterFeesMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetAfterFeesMinor, true
}

// SetNetAfterFeesMinor sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetNetAfterFeesMinor(v int32) {
	o.NetAfterFeesMinor = v
}

// GetFeeBreakdown returns the FeeBreakdown field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetFeeBreakdown() InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown {
	if o == nil {
		var ret InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown
		return ret
	}

	return o.FeeBreakdown
}

// GetFeeBreakdownOk returns a tuple with the FeeBreakdown field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetFeeBreakdownOk() (*InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeBreakdown, true
}

// SetFeeBreakdown sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetFeeBreakdown(v InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown) {
	o.FeeBreakdown = v
}

// GetPricingSource returns the PricingSource field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetPricingSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingSource
}

// GetPricingSourceOk returns a tuple with the PricingSource field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetPricingSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingSource, true
}

// SetPricingSource sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetPricingSource(v string) {
	o.PricingSource = v
}

// GetMatrixVersion returns the MatrixVersion field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetMatrixVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatrixVersion
}

// GetMatrixVersionOk returns a tuple with the MatrixVersion field value
// and a boolean to check if the value has been set.
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetMatrixVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatrixVersion, true
}

// SetMatrixVersion sets field value
func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetMatrixVersion(v string) {
	o.MatrixVersion = v
}

func (o InvoicesMarginEstimate200ResponseDataEstimatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicesMarginEstimate200ResponseDataEstimatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rail"] = o.Rail
	toSerialize["card_scheme"] = o.CardScheme.Get()
	toSerialize["gross_minor"] = o.GrossMinor
	toSerialize["estimated_fee_minor"] = o.EstimatedFeeMinor
	toSerialize["net_after_fees_minor"] = o.NetAfterFeesMinor
	toSerialize["fee_breakdown"] = o.FeeBreakdown
	toSerialize["pricing_source"] = o.PricingSource
	toSerialize["matrix_version"] = o.MatrixVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rail",
		"card_scheme",
		"gross_minor",
		"estimated_fee_minor",
		"net_after_fees_minor",
		"fee_breakdown",
		"pricing_source",
		"matrix_version",
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

	varInvoicesMarginEstimate200ResponseDataEstimatesInner := _InvoicesMarginEstimate200ResponseDataEstimatesInner{}

	err = json.Unmarshal(data, &varInvoicesMarginEstimate200ResponseDataEstimatesInner)

	if err != nil {
		return err
	}

	*o = InvoicesMarginEstimate200ResponseDataEstimatesInner(varInvoicesMarginEstimate200ResponseDataEstimatesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rail")
		delete(additionalProperties, "card_scheme")
		delete(additionalProperties, "gross_minor")
		delete(additionalProperties, "estimated_fee_minor")
		delete(additionalProperties, "net_after_fees_minor")
		delete(additionalProperties, "fee_breakdown")
		delete(additionalProperties, "pricing_source")
		delete(additionalProperties, "matrix_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoicesMarginEstimate200ResponseDataEstimatesInner struct {
	value *InvoicesMarginEstimate200ResponseDataEstimatesInner
	isSet bool
}

func (v NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) Get() *InvoicesMarginEstimate200ResponseDataEstimatesInner {
	return v.value
}

func (v *NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) Set(val *InvoicesMarginEstimate200ResponseDataEstimatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicesMarginEstimate200ResponseDataEstimatesInner(val *InvoicesMarginEstimate200ResponseDataEstimatesInner) *NullableInvoicesMarginEstimate200ResponseDataEstimatesInner {
	return &NullableInvoicesMarginEstimate200ResponseDataEstimatesInner{value: val, isSet: true}
}

func (v NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicesMarginEstimate200ResponseDataEstimatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


