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

// checks if the AnalyticsInvoicesGetArAging200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyticsInvoicesGetArAging200Response{}

// AnalyticsInvoicesGetArAging200Response struct for AnalyticsInvoicesGetArAging200Response
type AnalyticsInvoicesGetArAging200Response struct {
	AsOfDate string `json:"asOfDate"`
	Buckets []AnalyticsInvoicesGetArAging200ResponseBucketsInner `json:"buckets"`
	TotalOutstanding float32 `json:"totalOutstanding"`
	TotalInvoices float32 `json:"totalInvoices"`
	AdditionalProperties map[string]interface{}
}

type _AnalyticsInvoicesGetArAging200Response AnalyticsInvoicesGetArAging200Response

// NewAnalyticsInvoicesGetArAging200Response instantiates a new AnalyticsInvoicesGetArAging200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyticsInvoicesGetArAging200Response(asOfDate string, buckets []AnalyticsInvoicesGetArAging200ResponseBucketsInner, totalOutstanding float32, totalInvoices float32) *AnalyticsInvoicesGetArAging200Response {
	this := AnalyticsInvoicesGetArAging200Response{}
	this.AsOfDate = asOfDate
	this.Buckets = buckets
	this.TotalOutstanding = totalOutstanding
	this.TotalInvoices = totalInvoices
	return &this
}

// NewAnalyticsInvoicesGetArAging200ResponseWithDefaults instantiates a new AnalyticsInvoicesGetArAging200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyticsInvoicesGetArAging200ResponseWithDefaults() *AnalyticsInvoicesGetArAging200Response {
	this := AnalyticsInvoicesGetArAging200Response{}
	return &this
}

// GetAsOfDate returns the AsOfDate field value
func (o *AnalyticsInvoicesGetArAging200Response) GetAsOfDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AsOfDate
}

// GetAsOfDateOk returns a tuple with the AsOfDate field value
// and a boolean to check if the value has been set.
func (o *AnalyticsInvoicesGetArAging200Response) GetAsOfDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsOfDate, true
}

// SetAsOfDate sets field value
func (o *AnalyticsInvoicesGetArAging200Response) SetAsOfDate(v string) {
	o.AsOfDate = v
}

// GetBuckets returns the Buckets field value
func (o *AnalyticsInvoicesGetArAging200Response) GetBuckets() []AnalyticsInvoicesGetArAging200ResponseBucketsInner {
	if o == nil {
		var ret []AnalyticsInvoicesGetArAging200ResponseBucketsInner
		return ret
	}

	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *AnalyticsInvoicesGetArAging200Response) GetBucketsOk() ([]AnalyticsInvoicesGetArAging200ResponseBucketsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Buckets, true
}

// SetBuckets sets field value
func (o *AnalyticsInvoicesGetArAging200Response) SetBuckets(v []AnalyticsInvoicesGetArAging200ResponseBucketsInner) {
	o.Buckets = v
}

// GetTotalOutstanding returns the TotalOutstanding field value
func (o *AnalyticsInvoicesGetArAging200Response) GetTotalOutstanding() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalOutstanding
}

// GetTotalOutstandingOk returns a tuple with the TotalOutstanding field value
// and a boolean to check if the value has been set.
func (o *AnalyticsInvoicesGetArAging200Response) GetTotalOutstandingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalOutstanding, true
}

// SetTotalOutstanding sets field value
func (o *AnalyticsInvoicesGetArAging200Response) SetTotalOutstanding(v float32) {
	o.TotalOutstanding = v
}

// GetTotalInvoices returns the TotalInvoices field value
func (o *AnalyticsInvoicesGetArAging200Response) GetTotalInvoices() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalInvoices
}

// GetTotalInvoicesOk returns a tuple with the TotalInvoices field value
// and a boolean to check if the value has been set.
func (o *AnalyticsInvoicesGetArAging200Response) GetTotalInvoicesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalInvoices, true
}

// SetTotalInvoices sets field value
func (o *AnalyticsInvoicesGetArAging200Response) SetTotalInvoices(v float32) {
	o.TotalInvoices = v
}

func (o AnalyticsInvoicesGetArAging200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyticsInvoicesGetArAging200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asOfDate"] = o.AsOfDate
	toSerialize["buckets"] = o.Buckets
	toSerialize["totalOutstanding"] = o.TotalOutstanding
	toSerialize["totalInvoices"] = o.TotalInvoices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnalyticsInvoicesGetArAging200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asOfDate",
		"buckets",
		"totalOutstanding",
		"totalInvoices",
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

	varAnalyticsInvoicesGetArAging200Response := _AnalyticsInvoicesGetArAging200Response{}

	err = json.Unmarshal(data, &varAnalyticsInvoicesGetArAging200Response)

	if err != nil {
		return err
	}

	*o = AnalyticsInvoicesGetArAging200Response(varAnalyticsInvoicesGetArAging200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asOfDate")
		delete(additionalProperties, "buckets")
		delete(additionalProperties, "totalOutstanding")
		delete(additionalProperties, "totalInvoices")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnalyticsInvoicesGetArAging200Response struct {
	value *AnalyticsInvoicesGetArAging200Response
	isSet bool
}

func (v NullableAnalyticsInvoicesGetArAging200Response) Get() *AnalyticsInvoicesGetArAging200Response {
	return v.value
}

func (v *NullableAnalyticsInvoicesGetArAging200Response) Set(val *AnalyticsInvoicesGetArAging200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsInvoicesGetArAging200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsInvoicesGetArAging200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsInvoicesGetArAging200Response(val *AnalyticsInvoicesGetArAging200Response) *NullableAnalyticsInvoicesGetArAging200Response {
	return &NullableAnalyticsInvoicesGetArAging200Response{value: val, isSet: true}
}

func (v NullableAnalyticsInvoicesGetArAging200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsInvoicesGetArAging200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


