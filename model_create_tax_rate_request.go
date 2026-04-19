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

// checks if the CreateTaxRateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTaxRateRequest{}

// CreateTaxRateRequest Parameters for creating a new tax rate with a percentage and jurisdiction details.
type CreateTaxRateRequest struct {
	// Display name for the tax rate
	DisplayName string `json:"display_name"`
	// Description of the tax rate
	Description *string `json:"description,omitempty"`
	// Tax jurisdiction (e.g., 'US-CA', 'DE', 'GB')
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	// Tax percentage (e.g., 8.25 for 8.25%)
	Percentage float32 `json:"percentage"`
	// Whether tax is included in displayed price
	Inclusive *bool `json:"inclusive,omitempty"`
	// Type of tax
	TaxType *string `json:"tax_type,omitempty"`
	// ISO 2-letter country code
	Country *string `json:"country,omitempty"`
	// State/province code
	State *string `json:"state,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateTaxRateRequest CreateTaxRateRequest

// NewCreateTaxRateRequest instantiates a new CreateTaxRateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTaxRateRequest(displayName string, percentage float32) *CreateTaxRateRequest {
	this := CreateTaxRateRequest{}
	this.DisplayName = displayName
	this.Percentage = percentage
	var inclusive bool = false
	this.Inclusive = &inclusive
	var taxType string = "vat"
	this.TaxType = &taxType
	return &this
}

// NewCreateTaxRateRequestWithDefaults instantiates a new CreateTaxRateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTaxRateRequestWithDefaults() *CreateTaxRateRequest {
	this := CreateTaxRateRequest{}
	var inclusive bool = false
	this.Inclusive = &inclusive
	var taxType string = "vat"
	this.TaxType = &taxType
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CreateTaxRateRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateTaxRateRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTaxRateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetJurisdiction returns the Jurisdiction field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetJurisdiction() string {
	if o == nil || IsNil(o.Jurisdiction) {
		var ret string
		return ret
	}
	return *o.Jurisdiction
}

// GetJurisdictionOk returns a tuple with the Jurisdiction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetJurisdictionOk() (*string, bool) {
	if o == nil || IsNil(o.Jurisdiction) {
		return nil, false
	}
	return o.Jurisdiction, true
}

// HasJurisdiction returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasJurisdiction() bool {
	if o != nil && !IsNil(o.Jurisdiction) {
		return true
	}

	return false
}

// SetJurisdiction gets a reference to the given string and assigns it to the Jurisdiction field.
func (o *CreateTaxRateRequest) SetJurisdiction(v string) {
	o.Jurisdiction = &v
}

// GetPercentage returns the Percentage field value
func (o *CreateTaxRateRequest) GetPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *CreateTaxRateRequest) SetPercentage(v float32) {
	o.Percentage = v
}

// GetInclusive returns the Inclusive field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetInclusive() bool {
	if o == nil || IsNil(o.Inclusive) {
		var ret bool
		return ret
	}
	return *o.Inclusive
}

// GetInclusiveOk returns a tuple with the Inclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetInclusiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inclusive) {
		return nil, false
	}
	return o.Inclusive, true
}

// HasInclusive returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasInclusive() bool {
	if o != nil && !IsNil(o.Inclusive) {
		return true
	}

	return false
}

// SetInclusive gets a reference to the given bool and assigns it to the Inclusive field.
func (o *CreateTaxRateRequest) SetInclusive(v bool) {
	o.Inclusive = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *CreateTaxRateRequest) SetTaxType(v string) {
	o.TaxType = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CreateTaxRateRequest) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CreateTaxRateRequest) SetState(v string) {
	o.State = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateTaxRateRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTaxRateRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateTaxRateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateTaxRateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o CreateTaxRateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTaxRateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Jurisdiction) {
		toSerialize["jurisdiction"] = o.Jurisdiction
	}
	toSerialize["percentage"] = o.Percentage
	if !IsNil(o.Inclusive) {
		toSerialize["inclusive"] = o.Inclusive
	}
	if !IsNil(o.TaxType) {
		toSerialize["tax_type"] = o.TaxType
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateTaxRateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display_name",
		"percentage",
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

	varCreateTaxRateRequest := _CreateTaxRateRequest{}

	err = json.Unmarshal(data, &varCreateTaxRateRequest)

	if err != nil {
		return err
	}

	*o = CreateTaxRateRequest(varCreateTaxRateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "display_name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "jurisdiction")
		delete(additionalProperties, "percentage")
		delete(additionalProperties, "inclusive")
		delete(additionalProperties, "tax_type")
		delete(additionalProperties, "country")
		delete(additionalProperties, "state")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateTaxRateRequest struct {
	value *CreateTaxRateRequest
	isSet bool
}

func (v NullableCreateTaxRateRequest) Get() *CreateTaxRateRequest {
	return v.value
}

func (v *NullableCreateTaxRateRequest) Set(val *CreateTaxRateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTaxRateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTaxRateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTaxRateRequest(val *CreateTaxRateRequest) *NullableCreateTaxRateRequest {
	return &NullableCreateTaxRateRequest{value: val, isSet: true}
}

func (v NullableCreateTaxRateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTaxRateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


