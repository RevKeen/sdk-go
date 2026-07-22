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

// checks if the DdMandatePdfReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdMandatePdfReference{}

// DdMandatePdfReference struct for DdMandatePdfReference
type DdMandatePdfReference struct {
	Id string `json:"id"`
	MandateRequestId NullableString `json:"mandate_request_id"`
	MandateId string `json:"mandate_id"`
	CustomerId string `json:"customer_id"`
	MerchantId string `json:"merchant_id"`
	PdfType string `json:"pdf_type"`
	StorageKey string `json:"storage_key"`
	PdfUrl NullableString `json:"pdf_url"`
	GeneratedAt string `json:"generated_at"`
	TemplateVersion string `json:"template_version"`
	ComplianceVersion string `json:"compliance_version"`
	AdditionalProperties map[string]interface{}
}

type _DdMandatePdfReference DdMandatePdfReference

// NewDdMandatePdfReference instantiates a new DdMandatePdfReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdMandatePdfReference(id string, mandateRequestId NullableString, mandateId string, customerId string, merchantId string, pdfType string, storageKey string, pdfUrl NullableString, generatedAt string, templateVersion string, complianceVersion string) *DdMandatePdfReference {
	this := DdMandatePdfReference{}
	this.Id = id
	this.MandateRequestId = mandateRequestId
	this.MandateId = mandateId
	this.CustomerId = customerId
	this.MerchantId = merchantId
	this.PdfType = pdfType
	this.StorageKey = storageKey
	this.PdfUrl = pdfUrl
	this.GeneratedAt = generatedAt
	this.TemplateVersion = templateVersion
	this.ComplianceVersion = complianceVersion
	return &this
}

// NewDdMandatePdfReferenceWithDefaults instantiates a new DdMandatePdfReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdMandatePdfReferenceWithDefaults() *DdMandatePdfReference {
	this := DdMandatePdfReference{}
	return &this
}

// GetId returns the Id field value
func (o *DdMandatePdfReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DdMandatePdfReference) SetId(v string) {
	o.Id = v
}

// GetMandateRequestId returns the MandateRequestId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandatePdfReference) GetMandateRequestId() string {
	if o == nil || o.MandateRequestId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateRequestId.Get()
}

// GetMandateRequestIdOk returns a tuple with the MandateRequestId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandatePdfReference) GetMandateRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateRequestId.Get(), o.MandateRequestId.IsSet()
}

// SetMandateRequestId sets field value
func (o *DdMandatePdfReference) SetMandateRequestId(v string) {
	o.MandateRequestId.Set(&v)
}

// GetMandateId returns the MandateId field value
func (o *DdMandatePdfReference) GetMandateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateId, true
}

// SetMandateId sets field value
func (o *DdMandatePdfReference) SetMandateId(v string) {
	o.MandateId = v
}

// GetCustomerId returns the CustomerId field value
func (o *DdMandatePdfReference) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DdMandatePdfReference) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetMerchantId returns the MerchantId field value
func (o *DdMandatePdfReference) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *DdMandatePdfReference) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetPdfType returns the PdfType field value
func (o *DdMandatePdfReference) GetPdfType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PdfType
}

// GetPdfTypeOk returns a tuple with the PdfType field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetPdfTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PdfType, true
}

// SetPdfType sets field value
func (o *DdMandatePdfReference) SetPdfType(v string) {
	o.PdfType = v
}

// GetStorageKey returns the StorageKey field value
func (o *DdMandatePdfReference) GetStorageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageKey
}

// GetStorageKeyOk returns a tuple with the StorageKey field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetStorageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageKey, true
}

// SetStorageKey sets field value
func (o *DdMandatePdfReference) SetStorageKey(v string) {
	o.StorageKey = v
}

// GetPdfUrl returns the PdfUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdMandatePdfReference) GetPdfUrl() string {
	if o == nil || o.PdfUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PdfUrl.Get()
}

// GetPdfUrlOk returns a tuple with the PdfUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandatePdfReference) GetPdfUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PdfUrl.Get(), o.PdfUrl.IsSet()
}

// SetPdfUrl sets field value
func (o *DdMandatePdfReference) SetPdfUrl(v string) {
	o.PdfUrl.Set(&v)
}

// GetGeneratedAt returns the GeneratedAt field value
func (o *DdMandatePdfReference) GetGeneratedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetGeneratedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratedAt, true
}

// SetGeneratedAt sets field value
func (o *DdMandatePdfReference) SetGeneratedAt(v string) {
	o.GeneratedAt = v
}

// GetTemplateVersion returns the TemplateVersion field value
func (o *DdMandatePdfReference) GetTemplateVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateVersion
}

// GetTemplateVersionOk returns a tuple with the TemplateVersion field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetTemplateVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateVersion, true
}

// SetTemplateVersion sets field value
func (o *DdMandatePdfReference) SetTemplateVersion(v string) {
	o.TemplateVersion = v
}

// GetComplianceVersion returns the ComplianceVersion field value
func (o *DdMandatePdfReference) GetComplianceVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceVersion
}

// GetComplianceVersionOk returns a tuple with the ComplianceVersion field value
// and a boolean to check if the value has been set.
func (o *DdMandatePdfReference) GetComplianceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceVersion, true
}

// SetComplianceVersion sets field value
func (o *DdMandatePdfReference) SetComplianceVersion(v string) {
	o.ComplianceVersion = v
}

func (o DdMandatePdfReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdMandatePdfReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["mandate_request_id"] = o.MandateRequestId.Get()
	toSerialize["mandate_id"] = o.MandateId
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["pdf_type"] = o.PdfType
	toSerialize["storage_key"] = o.StorageKey
	toSerialize["pdf_url"] = o.PdfUrl.Get()
	toSerialize["generated_at"] = o.GeneratedAt
	toSerialize["template_version"] = o.TemplateVersion
	toSerialize["compliance_version"] = o.ComplianceVersion

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdMandatePdfReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"mandate_request_id",
		"mandate_id",
		"customer_id",
		"merchant_id",
		"pdf_type",
		"storage_key",
		"pdf_url",
		"generated_at",
		"template_version",
		"compliance_version",
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

	varDdMandatePdfReference := _DdMandatePdfReference{}

	err = json.Unmarshal(data, &varDdMandatePdfReference)

	if err != nil {
		return err
	}

	*o = DdMandatePdfReference(varDdMandatePdfReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mandate_request_id")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "pdf_type")
		delete(additionalProperties, "storage_key")
		delete(additionalProperties, "pdf_url")
		delete(additionalProperties, "generated_at")
		delete(additionalProperties, "template_version")
		delete(additionalProperties, "compliance_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdMandatePdfReference struct {
	value *DdMandatePdfReference
	isSet bool
}

func (v NullableDdMandatePdfReference) Get() *DdMandatePdfReference {
	return v.value
}

func (v *NullableDdMandatePdfReference) Set(val *DdMandatePdfReference) {
	v.value = val
	v.isSet = true
}

func (v NullableDdMandatePdfReference) IsSet() bool {
	return v.isSet
}

func (v *NullableDdMandatePdfReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdMandatePdfReference(val *DdMandatePdfReference) *NullableDdMandatePdfReference {
	return &NullableDdMandatePdfReference{value: val, isSet: true}
}

func (v NullableDdMandatePdfReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdMandatePdfReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


