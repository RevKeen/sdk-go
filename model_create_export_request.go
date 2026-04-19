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
	"fmt"
)

// checks if the CreateExportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExportRequest{}

// CreateExportRequest Parameters for creating an async data export. The export is processed in the background — poll GET /v2/exports/:id for status.
type CreateExportRequest struct {
	// The type of resource to export.
	ResourceType string `json:"resource_type"`
	// Output file format. csv: Comma-separated values. xlsx: Excel workbook.
	Format *string `json:"format,omitempty"`
	// Filters to apply (e.g., { status: 'paid', created_gte: '2024-01-01' })
	Filters map[string]interface{} `json:"filters,omitempty"`
	// Column set to include. standard: common fields. full: all fields. minimal: IDs and key fields only.
	ColumnSet *string `json:"column_set,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateExportRequest CreateExportRequest

// NewCreateExportRequest instantiates a new CreateExportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExportRequest(resourceType string) *CreateExportRequest {
	this := CreateExportRequest{}
	this.ResourceType = resourceType
	var format string = "csv"
	this.Format = &format
	var columnSet string = "standard"
	this.ColumnSet = &columnSet
	return &this
}

// NewCreateExportRequestWithDefaults instantiates a new CreateExportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExportRequestWithDefaults() *CreateExportRequest {
	this := CreateExportRequest{}
	var format string = "csv"
	this.Format = &format
	var columnSet string = "standard"
	this.ColumnSet = &columnSet
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *CreateExportRequest) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CreateExportRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CreateExportRequest) SetResourceType(v string) {
	o.ResourceType = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateExportRequest) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExportRequest) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateExportRequest) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateExportRequest) SetFormat(v string) {
	o.Format = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CreateExportRequest) GetFilters() map[string]interface{} {
	if o == nil || IsNil(o.Filters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExportRequest) GetFiltersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filters) {
		return map[string]interface{}{}, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CreateExportRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string]interface{} and assigns it to the Filters field.
func (o *CreateExportRequest) SetFilters(v map[string]interface{}) {
	o.Filters = v
}

// GetColumnSet returns the ColumnSet field value if set, zero value otherwise.
func (o *CreateExportRequest) GetColumnSet() string {
	if o == nil || IsNil(o.ColumnSet) {
		var ret string
		return ret
	}
	return *o.ColumnSet
}

// GetColumnSetOk returns a tuple with the ColumnSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExportRequest) GetColumnSetOk() (*string, bool) {
	if o == nil || IsNil(o.ColumnSet) {
		return nil, false
	}
	return o.ColumnSet, true
}

// HasColumnSet returns a boolean if a field has been set.
func (o *CreateExportRequest) HasColumnSet() bool {
	if o != nil && !IsNil(o.ColumnSet) {
		return true
	}

	return false
}

// SetColumnSet gets a reference to the given string and assigns it to the ColumnSet field.
func (o *CreateExportRequest) SetColumnSet(v string) {
	o.ColumnSet = &v
}

func (o CreateExportRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateExportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resource_type"] = o.ResourceType
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.ColumnSet) {
		toSerialize["column_set"] = o.ColumnSet
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateExportRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resource_type",
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

	varCreateExportRequest := _CreateExportRequest{}

	err = json.Unmarshal(data, &varCreateExportRequest)

	if err != nil {
		return err
	}

	*o = CreateExportRequest(varCreateExportRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resource_type")
		delete(additionalProperties, "format")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "column_set")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateExportRequest struct {
	value *CreateExportRequest
	isSet bool
}

func (v NullableCreateExportRequest) Get() *CreateExportRequest {
	return v.value
}

func (v *NullableCreateExportRequest) Set(val *CreateExportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExportRequest(val *CreateExportRequest) *NullableCreateExportRequest {
	return &NullableCreateExportRequest{value: val, isSet: true}
}

func (v NullableCreateExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


