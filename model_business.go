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
	"time"
	"fmt"
)

// checks if the Business type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Business{}

// Business A business entity (company) associated with a customer for B2B billing.
type Business struct {
	// Unique identifier (UUID)
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Associated customer ID
	CustomerId NullableString `json:"customer_id,omitempty"`
	// Business name
	Name string `json:"name"`
	// Company registration number
	CompanyNumber NullableString `json:"company_number,omitempty"`
	// Tax ID (VAT, EIN, etc.)
	TaxIdentifier NullableString `json:"tax_identifier,omitempty"`
	// Billing email address
	BillingEmail NullableString `json:"billing_email,omitempty"`
	// Billing phone number
	BillingPhone NullableString `json:"billing_phone,omitempty"`
	// Billing address line 1
	BillingAddressLine1 NullableString `json:"billing_address_line1,omitempty"`
	// Billing address line 2
	BillingAddressLine2 NullableString `json:"billing_address_line2,omitempty"`
	// Billing city
	BillingCity NullableString `json:"billing_city,omitempty"`
	// Billing postcode/ZIP
	BillingPostcode NullableString `json:"billing_postcode,omitempty"`
	// Billing country (ISO 3166-1 alpha-2)
	BillingCountry NullableString `json:"billing_country,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Creation timestamp (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp (ISO 8601)
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _Business Business

// NewBusiness instantiates a new Business object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusiness(id string, object string, name string, createdAt time.Time, updatedAt time.Time) *Business {
	this := Business{}
	this.Id = id
	this.Object = object
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewBusinessWithDefaults instantiates a new Business object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessWithDefaults() *Business {
	this := Business{}
	return &this
}

// GetId returns the Id field value
func (o *Business) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Business) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Business) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Business) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Business) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Business) SetObject(v string) {
	o.Object = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Business) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Business) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Business) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Business) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetName returns the Name field value
func (o *Business) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Business) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Business) SetName(v string) {
	o.Name = v
}

// GetCompanyNumber returns the CompanyNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetCompanyNumber() string {
	if o == nil || IsNil(o.CompanyNumber.Get()) {
		var ret string
		return ret
	}
	return *o.CompanyNumber.Get()
}

// GetCompanyNumberOk returns a tuple with the CompanyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetCompanyNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyNumber.Get(), o.CompanyNumber.IsSet()
}

// HasCompanyNumber returns a boolean if a field has been set.
func (o *Business) HasCompanyNumber() bool {
	if o != nil && o.CompanyNumber.IsSet() {
		return true
	}

	return false
}

// SetCompanyNumber gets a reference to the given NullableString and assigns it to the CompanyNumber field.
func (o *Business) SetCompanyNumber(v string) {
	o.CompanyNumber.Set(&v)
}
// SetCompanyNumberNil sets the value for CompanyNumber to be an explicit nil
func (o *Business) SetCompanyNumberNil() {
	o.CompanyNumber.Set(nil)
}

// UnsetCompanyNumber ensures that no value is present for CompanyNumber, not even an explicit nil
func (o *Business) UnsetCompanyNumber() {
	o.CompanyNumber.Unset()
}

// GetTaxIdentifier returns the TaxIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetTaxIdentifier() string {
	if o == nil || IsNil(o.TaxIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.TaxIdentifier.Get()
}

// GetTaxIdentifierOk returns a tuple with the TaxIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetTaxIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxIdentifier.Get(), o.TaxIdentifier.IsSet()
}

// HasTaxIdentifier returns a boolean if a field has been set.
func (o *Business) HasTaxIdentifier() bool {
	if o != nil && o.TaxIdentifier.IsSet() {
		return true
	}

	return false
}

// SetTaxIdentifier gets a reference to the given NullableString and assigns it to the TaxIdentifier field.
func (o *Business) SetTaxIdentifier(v string) {
	o.TaxIdentifier.Set(&v)
}
// SetTaxIdentifierNil sets the value for TaxIdentifier to be an explicit nil
func (o *Business) SetTaxIdentifierNil() {
	o.TaxIdentifier.Set(nil)
}

// UnsetTaxIdentifier ensures that no value is present for TaxIdentifier, not even an explicit nil
func (o *Business) UnsetTaxIdentifier() {
	o.TaxIdentifier.Unset()
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *Business) HasBillingEmail() bool {
	if o != nil && o.BillingEmail.IsSet() {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given NullableString and assigns it to the BillingEmail field.
func (o *Business) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}
// SetBillingEmailNil sets the value for BillingEmail to be an explicit nil
func (o *Business) SetBillingEmailNil() {
	o.BillingEmail.Set(nil)
}

// UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
func (o *Business) UnsetBillingEmail() {
	o.BillingEmail.Unset()
}

// GetBillingPhone returns the BillingPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingPhone() string {
	if o == nil || IsNil(o.BillingPhone.Get()) {
		var ret string
		return ret
	}
	return *o.BillingPhone.Get()
}

// GetBillingPhoneOk returns a tuple with the BillingPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingPhone.Get(), o.BillingPhone.IsSet()
}

// HasBillingPhone returns a boolean if a field has been set.
func (o *Business) HasBillingPhone() bool {
	if o != nil && o.BillingPhone.IsSet() {
		return true
	}

	return false
}

// SetBillingPhone gets a reference to the given NullableString and assigns it to the BillingPhone field.
func (o *Business) SetBillingPhone(v string) {
	o.BillingPhone.Set(&v)
}
// SetBillingPhoneNil sets the value for BillingPhone to be an explicit nil
func (o *Business) SetBillingPhoneNil() {
	o.BillingPhone.Set(nil)
}

// UnsetBillingPhone ensures that no value is present for BillingPhone, not even an explicit nil
func (o *Business) UnsetBillingPhone() {
	o.BillingPhone.Unset()
}

// GetBillingAddressLine1 returns the BillingAddressLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingAddressLine1() string {
	if o == nil || IsNil(o.BillingAddressLine1.Get()) {
		var ret string
		return ret
	}
	return *o.BillingAddressLine1.Get()
}

// GetBillingAddressLine1Ok returns a tuple with the BillingAddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAddressLine1.Get(), o.BillingAddressLine1.IsSet()
}

// HasBillingAddressLine1 returns a boolean if a field has been set.
func (o *Business) HasBillingAddressLine1() bool {
	if o != nil && o.BillingAddressLine1.IsSet() {
		return true
	}

	return false
}

// SetBillingAddressLine1 gets a reference to the given NullableString and assigns it to the BillingAddressLine1 field.
func (o *Business) SetBillingAddressLine1(v string) {
	o.BillingAddressLine1.Set(&v)
}
// SetBillingAddressLine1Nil sets the value for BillingAddressLine1 to be an explicit nil
func (o *Business) SetBillingAddressLine1Nil() {
	o.BillingAddressLine1.Set(nil)
}

// UnsetBillingAddressLine1 ensures that no value is present for BillingAddressLine1, not even an explicit nil
func (o *Business) UnsetBillingAddressLine1() {
	o.BillingAddressLine1.Unset()
}

// GetBillingAddressLine2 returns the BillingAddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingAddressLine2() string {
	if o == nil || IsNil(o.BillingAddressLine2.Get()) {
		var ret string
		return ret
	}
	return *o.BillingAddressLine2.Get()
}

// GetBillingAddressLine2Ok returns a tuple with the BillingAddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingAddressLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingAddressLine2.Get(), o.BillingAddressLine2.IsSet()
}

// HasBillingAddressLine2 returns a boolean if a field has been set.
func (o *Business) HasBillingAddressLine2() bool {
	if o != nil && o.BillingAddressLine2.IsSet() {
		return true
	}

	return false
}

// SetBillingAddressLine2 gets a reference to the given NullableString and assigns it to the BillingAddressLine2 field.
func (o *Business) SetBillingAddressLine2(v string) {
	o.BillingAddressLine2.Set(&v)
}
// SetBillingAddressLine2Nil sets the value for BillingAddressLine2 to be an explicit nil
func (o *Business) SetBillingAddressLine2Nil() {
	o.BillingAddressLine2.Set(nil)
}

// UnsetBillingAddressLine2 ensures that no value is present for BillingAddressLine2, not even an explicit nil
func (o *Business) UnsetBillingAddressLine2() {
	o.BillingAddressLine2.Unset()
}

// GetBillingCity returns the BillingCity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingCity() string {
	if o == nil || IsNil(o.BillingCity.Get()) {
		var ret string
		return ret
	}
	return *o.BillingCity.Get()
}

// GetBillingCityOk returns a tuple with the BillingCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingCity.Get(), o.BillingCity.IsSet()
}

// HasBillingCity returns a boolean if a field has been set.
func (o *Business) HasBillingCity() bool {
	if o != nil && o.BillingCity.IsSet() {
		return true
	}

	return false
}

// SetBillingCity gets a reference to the given NullableString and assigns it to the BillingCity field.
func (o *Business) SetBillingCity(v string) {
	o.BillingCity.Set(&v)
}
// SetBillingCityNil sets the value for BillingCity to be an explicit nil
func (o *Business) SetBillingCityNil() {
	o.BillingCity.Set(nil)
}

// UnsetBillingCity ensures that no value is present for BillingCity, not even an explicit nil
func (o *Business) UnsetBillingCity() {
	o.BillingCity.Unset()
}

// GetBillingPostcode returns the BillingPostcode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingPostcode() string {
	if o == nil || IsNil(o.BillingPostcode.Get()) {
		var ret string
		return ret
	}
	return *o.BillingPostcode.Get()
}

// GetBillingPostcodeOk returns a tuple with the BillingPostcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingPostcode.Get(), o.BillingPostcode.IsSet()
}

// HasBillingPostcode returns a boolean if a field has been set.
func (o *Business) HasBillingPostcode() bool {
	if o != nil && o.BillingPostcode.IsSet() {
		return true
	}

	return false
}

// SetBillingPostcode gets a reference to the given NullableString and assigns it to the BillingPostcode field.
func (o *Business) SetBillingPostcode(v string) {
	o.BillingPostcode.Set(&v)
}
// SetBillingPostcodeNil sets the value for BillingPostcode to be an explicit nil
func (o *Business) SetBillingPostcodeNil() {
	o.BillingPostcode.Set(nil)
}

// UnsetBillingPostcode ensures that no value is present for BillingPostcode, not even an explicit nil
func (o *Business) UnsetBillingPostcode() {
	o.BillingPostcode.Unset()
}

// GetBillingCountry returns the BillingCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Business) GetBillingCountry() string {
	if o == nil || IsNil(o.BillingCountry.Get()) {
		var ret string
		return ret
	}
	return *o.BillingCountry.Get()
}

// GetBillingCountryOk returns a tuple with the BillingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Business) GetBillingCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingCountry.Get(), o.BillingCountry.IsSet()
}

// HasBillingCountry returns a boolean if a field has been set.
func (o *Business) HasBillingCountry() bool {
	if o != nil && o.BillingCountry.IsSet() {
		return true
	}

	return false
}

// SetBillingCountry gets a reference to the given NullableString and assigns it to the BillingCountry field.
func (o *Business) SetBillingCountry(v string) {
	o.BillingCountry.Set(&v)
}
// SetBillingCountryNil sets the value for BillingCountry to be an explicit nil
func (o *Business) SetBillingCountryNil() {
	o.BillingCountry.Set(nil)
}

// UnsetBillingCountry ensures that no value is present for BillingCountry, not even an explicit nil
func (o *Business) UnsetBillingCountry() {
	o.BillingCountry.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Business) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Business) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Business) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Business) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Business) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Business) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Business) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Business) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Business) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Business) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Business) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Business) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	if o.CustomerId.IsSet() {
		toSerialize["customer_id"] = o.CustomerId.Get()
	}
	toSerialize["name"] = o.Name
	if o.CompanyNumber.IsSet() {
		toSerialize["company_number"] = o.CompanyNumber.Get()
	}
	if o.TaxIdentifier.IsSet() {
		toSerialize["tax_identifier"] = o.TaxIdentifier.Get()
	}
	if o.BillingEmail.IsSet() {
		toSerialize["billing_email"] = o.BillingEmail.Get()
	}
	if o.BillingPhone.IsSet() {
		toSerialize["billing_phone"] = o.BillingPhone.Get()
	}
	if o.BillingAddressLine1.IsSet() {
		toSerialize["billing_address_line1"] = o.BillingAddressLine1.Get()
	}
	if o.BillingAddressLine2.IsSet() {
		toSerialize["billing_address_line2"] = o.BillingAddressLine2.Get()
	}
	if o.BillingCity.IsSet() {
		toSerialize["billing_city"] = o.BillingCity.Get()
	}
	if o.BillingPostcode.IsSet() {
		toSerialize["billing_postcode"] = o.BillingPostcode.Get()
	}
	if o.BillingCountry.IsSet() {
		toSerialize["billing_country"] = o.BillingCountry.Get()
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Business) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"name",
		"created_at",
		"updated_at",
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

	varBusiness := _Business{}

	err = json.Unmarshal(data, &varBusiness)

	if err != nil {
		return err
	}

	*o = Business(varBusiness)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "company_number")
		delete(additionalProperties, "tax_identifier")
		delete(additionalProperties, "billing_email")
		delete(additionalProperties, "billing_phone")
		delete(additionalProperties, "billing_address_line1")
		delete(additionalProperties, "billing_address_line2")
		delete(additionalProperties, "billing_city")
		delete(additionalProperties, "billing_postcode")
		delete(additionalProperties, "billing_country")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBusiness struct {
	value *Business
	isSet bool
}

func (v NullableBusiness) Get() *Business {
	return v.value
}

func (v *NullableBusiness) Set(val *Business) {
	v.value = val
	v.isSet = true
}

func (v NullableBusiness) IsSet() bool {
	return v.isSet
}

func (v *NullableBusiness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusiness(val *Business) *NullableBusiness {
	return &NullableBusiness{value: val, isSet: true}
}

func (v NullableBusiness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusiness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


