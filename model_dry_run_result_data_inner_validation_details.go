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

// checks if the DryRunResultDataInnerValidationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DryRunResultDataInnerValidationDetails{}

// DryRunResultDataInnerValidationDetails struct for DryRunResultDataInnerValidationDetails
type DryRunResultDataInnerValidationDetails struct {
	MeterMatched bool `json:"meterMatched"`
	CustomerMatched bool `json:"customerMatched"`
	BillableQuantity float32 `json:"billableQuantity"`
	FiltersPassed bool `json:"filtersPassed"`
	AdditionalProperties map[string]interface{}
}

type _DryRunResultDataInnerValidationDetails DryRunResultDataInnerValidationDetails

// NewDryRunResultDataInnerValidationDetails instantiates a new DryRunResultDataInnerValidationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDryRunResultDataInnerValidationDetails(meterMatched bool, customerMatched bool, billableQuantity float32, filtersPassed bool) *DryRunResultDataInnerValidationDetails {
	this := DryRunResultDataInnerValidationDetails{}
	this.MeterMatched = meterMatched
	this.CustomerMatched = customerMatched
	this.BillableQuantity = billableQuantity
	this.FiltersPassed = filtersPassed
	return &this
}

// NewDryRunResultDataInnerValidationDetailsWithDefaults instantiates a new DryRunResultDataInnerValidationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDryRunResultDataInnerValidationDetailsWithDefaults() *DryRunResultDataInnerValidationDetails {
	this := DryRunResultDataInnerValidationDetails{}
	return &this
}

// GetMeterMatched returns the MeterMatched field value
func (o *DryRunResultDataInnerValidationDetails) GetMeterMatched() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MeterMatched
}

// GetMeterMatchedOk returns a tuple with the MeterMatched field value
// and a boolean to check if the value has been set.
func (o *DryRunResultDataInnerValidationDetails) GetMeterMatchedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterMatched, true
}

// SetMeterMatched sets field value
func (o *DryRunResultDataInnerValidationDetails) SetMeterMatched(v bool) {
	o.MeterMatched = v
}

// GetCustomerMatched returns the CustomerMatched field value
func (o *DryRunResultDataInnerValidationDetails) GetCustomerMatched() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CustomerMatched
}

// GetCustomerMatchedOk returns a tuple with the CustomerMatched field value
// and a boolean to check if the value has been set.
func (o *DryRunResultDataInnerValidationDetails) GetCustomerMatchedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerMatched, true
}

// SetCustomerMatched sets field value
func (o *DryRunResultDataInnerValidationDetails) SetCustomerMatched(v bool) {
	o.CustomerMatched = v
}

// GetBillableQuantity returns the BillableQuantity field value
func (o *DryRunResultDataInnerValidationDetails) GetBillableQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BillableQuantity
}

// GetBillableQuantityOk returns a tuple with the BillableQuantity field value
// and a boolean to check if the value has been set.
func (o *DryRunResultDataInnerValidationDetails) GetBillableQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableQuantity, true
}

// SetBillableQuantity sets field value
func (o *DryRunResultDataInnerValidationDetails) SetBillableQuantity(v float32) {
	o.BillableQuantity = v
}

// GetFiltersPassed returns the FiltersPassed field value
func (o *DryRunResultDataInnerValidationDetails) GetFiltersPassed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FiltersPassed
}

// GetFiltersPassedOk returns a tuple with the FiltersPassed field value
// and a boolean to check if the value has been set.
func (o *DryRunResultDataInnerValidationDetails) GetFiltersPassedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiltersPassed, true
}

// SetFiltersPassed sets field value
func (o *DryRunResultDataInnerValidationDetails) SetFiltersPassed(v bool) {
	o.FiltersPassed = v
}

func (o DryRunResultDataInnerValidationDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DryRunResultDataInnerValidationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meterMatched"] = o.MeterMatched
	toSerialize["customerMatched"] = o.CustomerMatched
	toSerialize["billableQuantity"] = o.BillableQuantity
	toSerialize["filtersPassed"] = o.FiltersPassed

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DryRunResultDataInnerValidationDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meterMatched",
		"customerMatched",
		"billableQuantity",
		"filtersPassed",
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

	varDryRunResultDataInnerValidationDetails := _DryRunResultDataInnerValidationDetails{}

	err = json.Unmarshal(data, &varDryRunResultDataInnerValidationDetails)

	if err != nil {
		return err
	}

	*o = DryRunResultDataInnerValidationDetails(varDryRunResultDataInnerValidationDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "meterMatched")
		delete(additionalProperties, "customerMatched")
		delete(additionalProperties, "billableQuantity")
		delete(additionalProperties, "filtersPassed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDryRunResultDataInnerValidationDetails struct {
	value *DryRunResultDataInnerValidationDetails
	isSet bool
}

func (v NullableDryRunResultDataInnerValidationDetails) Get() *DryRunResultDataInnerValidationDetails {
	return v.value
}

func (v *NullableDryRunResultDataInnerValidationDetails) Set(val *DryRunResultDataInnerValidationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDryRunResultDataInnerValidationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDryRunResultDataInnerValidationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDryRunResultDataInnerValidationDetails(val *DryRunResultDataInnerValidationDetails) *NullableDryRunResultDataInnerValidationDetails {
	return &NullableDryRunResultDataInnerValidationDetails{value: val, isSet: true}
}

func (v NullableDryRunResultDataInnerValidationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDryRunResultDataInnerValidationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


