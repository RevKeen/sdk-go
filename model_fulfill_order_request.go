/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
)

// checks if the FulfillOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillOrderRequest{}

// FulfillOrderRequest Parameters for marking an order as fulfilled, with optional tracking and shipping details.
type FulfillOrderRequest struct {
	// Shipping tracking number
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// Shipping carrier (e.g., 'ups', 'fedex', 'usps')
	Carrier *string `json:"carrier,omitempty"`
	// Specific line items to fulfill (defaults to all)
	LineItemIds []string `json:"line_item_ids,omitempty"`
	// Send fulfillment notification to customer
	NotifyCustomer *bool `json:"notify_customer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FulfillOrderRequest FulfillOrderRequest

// NewFulfillOrderRequest instantiates a new FulfillOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillOrderRequest() *FulfillOrderRequest {
	this := FulfillOrderRequest{}
	var notifyCustomer bool = true
	this.NotifyCustomer = &notifyCustomer
	return &this
}

// NewFulfillOrderRequestWithDefaults instantiates a new FulfillOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillOrderRequestWithDefaults() *FulfillOrderRequest {
	this := FulfillOrderRequest{}
	var notifyCustomer bool = true
	this.NotifyCustomer = &notifyCustomer
	return &this
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *FulfillOrderRequest) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillOrderRequest) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *FulfillOrderRequest) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *FulfillOrderRequest) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *FulfillOrderRequest) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillOrderRequest) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *FulfillOrderRequest) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *FulfillOrderRequest) SetCarrier(v string) {
	o.Carrier = &v
}

// GetLineItemIds returns the LineItemIds field value if set, zero value otherwise.
func (o *FulfillOrderRequest) GetLineItemIds() []string {
	if o == nil || IsNil(o.LineItemIds) {
		var ret []string
		return ret
	}
	return o.LineItemIds
}

// GetLineItemIdsOk returns a tuple with the LineItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillOrderRequest) GetLineItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LineItemIds) {
		return nil, false
	}
	return o.LineItemIds, true
}

// HasLineItemIds returns a boolean if a field has been set.
func (o *FulfillOrderRequest) HasLineItemIds() bool {
	if o != nil && !IsNil(o.LineItemIds) {
		return true
	}

	return false
}

// SetLineItemIds gets a reference to the given []string and assigns it to the LineItemIds field.
func (o *FulfillOrderRequest) SetLineItemIds(v []string) {
	o.LineItemIds = v
}

// GetNotifyCustomer returns the NotifyCustomer field value if set, zero value otherwise.
func (o *FulfillOrderRequest) GetNotifyCustomer() bool {
	if o == nil || IsNil(o.NotifyCustomer) {
		var ret bool
		return ret
	}
	return *o.NotifyCustomer
}

// GetNotifyCustomerOk returns a tuple with the NotifyCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillOrderRequest) GetNotifyCustomerOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyCustomer) {
		return nil, false
	}
	return o.NotifyCustomer, true
}

// HasNotifyCustomer returns a boolean if a field has been set.
func (o *FulfillOrderRequest) HasNotifyCustomer() bool {
	if o != nil && !IsNil(o.NotifyCustomer) {
		return true
	}

	return false
}

// SetNotifyCustomer gets a reference to the given bool and assigns it to the NotifyCustomer field.
func (o *FulfillOrderRequest) SetNotifyCustomer(v bool) {
	o.NotifyCustomer = &v
}

func (o FulfillOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.LineItemIds) {
		toSerialize["line_item_ids"] = o.LineItemIds
	}
	if !IsNil(o.NotifyCustomer) {
		toSerialize["notify_customer"] = o.NotifyCustomer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FulfillOrderRequest) UnmarshalJSON(data []byte) (err error) {
	varFulfillOrderRequest := _FulfillOrderRequest{}

	err = json.Unmarshal(data, &varFulfillOrderRequest)

	if err != nil {
		return err
	}

	*o = FulfillOrderRequest(varFulfillOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tracking_number")
		delete(additionalProperties, "carrier")
		delete(additionalProperties, "line_item_ids")
		delete(additionalProperties, "notify_customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFulfillOrderRequest struct {
	value *FulfillOrderRequest
	isSet bool
}

func (v NullableFulfillOrderRequest) Get() *FulfillOrderRequest {
	return v.value
}

func (v *NullableFulfillOrderRequest) Set(val *FulfillOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillOrderRequest(val *FulfillOrderRequest) *NullableFulfillOrderRequest {
	return &NullableFulfillOrderRequest{value: val, isSet: true}
}

func (v NullableFulfillOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


