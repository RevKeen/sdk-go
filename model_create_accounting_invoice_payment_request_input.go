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

// checks if the CreateAccountingInvoicePaymentRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountingInvoicePaymentRequestInput{}

// CreateAccountingInvoicePaymentRequestInput Create an accounting-led invoice payment request and mint/reuse a checkout session.
type CreateAccountingInvoicePaymentRequestInput struct {
	Provider string `json:"provider"`
	ConnectionId string `json:"connection_id"`
	ProviderAccountId string `json:"provider_account_id"`
	ExternalInvoice CreateAccountingInvoicePaymentRequestInputExternalInvoice `json:"external_invoice"`
	ExternalCustomer *CreateAccountingInvoicePaymentRequestInputExternalCustomer `json:"external_customer,omitempty"`
	AmountDueMinor int32 `json:"amount_due_minor"`
	TotalAmountMinor NullableInt32 `json:"total_amount_minor,omitempty"`
	AmountPaidMinor NullableInt32 `json:"amount_paid_minor,omitempty"`
	Currency string `json:"currency"`
	Checkout *CreateAccountingInvoicePaymentRequestInputCheckout `json:"checkout,omitempty"`
	PayloadFingerprint NullableString `json:"payload_fingerprint,omitempty"`
	ProviderMetadata map[string]interface{} `json:"provider_metadata,omitempty"`
	SafeProviderInvoiceSnapshot map[string]interface{} `json:"safe_provider_invoice_snapshot,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAccountingInvoicePaymentRequestInput CreateAccountingInvoicePaymentRequestInput

// NewCreateAccountingInvoicePaymentRequestInput instantiates a new CreateAccountingInvoicePaymentRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountingInvoicePaymentRequestInput(provider string, connectionId string, providerAccountId string, externalInvoice CreateAccountingInvoicePaymentRequestInputExternalInvoice, amountDueMinor int32, currency string) *CreateAccountingInvoicePaymentRequestInput {
	this := CreateAccountingInvoicePaymentRequestInput{}
	this.Provider = provider
	this.ConnectionId = connectionId
	this.ProviderAccountId = providerAccountId
	this.ExternalInvoice = externalInvoice
	this.AmountDueMinor = amountDueMinor
	this.Currency = currency
	return &this
}

// NewCreateAccountingInvoicePaymentRequestInputWithDefaults instantiates a new CreateAccountingInvoicePaymentRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountingInvoicePaymentRequestInputWithDefaults() *CreateAccountingInvoicePaymentRequestInput {
	this := CreateAccountingInvoicePaymentRequestInput{}
	return &this
}

// GetProvider returns the Provider field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetProvider(v string) {
	o.Provider = v
}

// GetConnectionId returns the ConnectionId field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetProviderAccountId returns the ProviderAccountId field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderAccountId
}

// GetProviderAccountIdOk returns a tuple with the ProviderAccountId field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderAccountId, true
}

// SetProviderAccountId sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetProviderAccountId(v string) {
	o.ProviderAccountId = v
}

// GetExternalInvoice returns the ExternalInvoice field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalInvoice() CreateAccountingInvoicePaymentRequestInputExternalInvoice {
	if o == nil {
		var ret CreateAccountingInvoicePaymentRequestInputExternalInvoice
		return ret
	}

	return o.ExternalInvoice
}

// GetExternalInvoiceOk returns a tuple with the ExternalInvoice field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalInvoiceOk() (*CreateAccountingInvoicePaymentRequestInputExternalInvoice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalInvoice, true
}

// SetExternalInvoice sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetExternalInvoice(v CreateAccountingInvoicePaymentRequestInputExternalInvoice) {
	o.ExternalInvoice = v
}

// GetExternalCustomer returns the ExternalCustomer field value if set, zero value otherwise.
func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalCustomer() CreateAccountingInvoicePaymentRequestInputExternalCustomer {
	if o == nil || IsNil(o.ExternalCustomer) {
		var ret CreateAccountingInvoicePaymentRequestInputExternalCustomer
		return ret
	}
	return *o.ExternalCustomer
}

// GetExternalCustomerOk returns a tuple with the ExternalCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalCustomerOk() (*CreateAccountingInvoicePaymentRequestInputExternalCustomer, bool) {
	if o == nil || IsNil(o.ExternalCustomer) {
		return nil, false
	}
	return o.ExternalCustomer, true
}

// HasExternalCustomer returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasExternalCustomer() bool {
	if o != nil && !IsNil(o.ExternalCustomer) {
		return true
	}

	return false
}

// SetExternalCustomer gets a reference to the given CreateAccountingInvoicePaymentRequestInputExternalCustomer and assigns it to the ExternalCustomer field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetExternalCustomer(v CreateAccountingInvoicePaymentRequestInputExternalCustomer) {
	o.ExternalCustomer = &v
}

// GetAmountDueMinor returns the AmountDueMinor field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountDueMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDueMinor
}

// GetAmountDueMinorOk returns a tuple with the AmountDueMinor field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountDueMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDueMinor, true
}

// SetAmountDueMinor sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountDueMinor(v int32) {
	o.AmountDueMinor = v
}

// GetTotalAmountMinor returns the TotalAmountMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInput) GetTotalAmountMinor() int32 {
	if o == nil || IsNil(o.TotalAmountMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalAmountMinor.Get()
}

// GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInput) GetTotalAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalAmountMinor.Get(), o.TotalAmountMinor.IsSet()
}

// HasTotalAmountMinor returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasTotalAmountMinor() bool {
	if o != nil && o.TotalAmountMinor.IsSet() {
		return true
	}

	return false
}

// SetTotalAmountMinor gets a reference to the given NullableInt32 and assigns it to the TotalAmountMinor field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetTotalAmountMinor(v int32) {
	o.TotalAmountMinor.Set(&v)
}
// SetTotalAmountMinorNil sets the value for TotalAmountMinor to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) SetTotalAmountMinorNil() {
	o.TotalAmountMinor.Set(nil)
}

// UnsetTotalAmountMinor ensures that no value is present for TotalAmountMinor, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) UnsetTotalAmountMinor() {
	o.TotalAmountMinor.Unset()
}

// GetAmountPaidMinor returns the AmountPaidMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountPaidMinor() int32 {
	if o == nil || IsNil(o.AmountPaidMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountPaidMinor.Get()
}

// GetAmountPaidMinorOk returns a tuple with the AmountPaidMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountPaidMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountPaidMinor.Get(), o.AmountPaidMinor.IsSet()
}

// HasAmountPaidMinor returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasAmountPaidMinor() bool {
	if o != nil && o.AmountPaidMinor.IsSet() {
		return true
	}

	return false
}

// SetAmountPaidMinor gets a reference to the given NullableInt32 and assigns it to the AmountPaidMinor field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountPaidMinor(v int32) {
	o.AmountPaidMinor.Set(&v)
}
// SetAmountPaidMinorNil sets the value for AmountPaidMinor to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountPaidMinorNil() {
	o.AmountPaidMinor.Set(nil)
}

// UnsetAmountPaidMinor ensures that no value is present for AmountPaidMinor, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) UnsetAmountPaidMinor() {
	o.AmountPaidMinor.Unset()
}

// GetCurrency returns the Currency field value
func (o *CreateAccountingInvoicePaymentRequestInput) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateAccountingInvoicePaymentRequestInput) SetCurrency(v string) {
	o.Currency = v
}

// GetCheckout returns the Checkout field value if set, zero value otherwise.
func (o *CreateAccountingInvoicePaymentRequestInput) GetCheckout() CreateAccountingInvoicePaymentRequestInputCheckout {
	if o == nil || IsNil(o.Checkout) {
		var ret CreateAccountingInvoicePaymentRequestInputCheckout
		return ret
	}
	return *o.Checkout
}

// GetCheckoutOk returns a tuple with the Checkout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetCheckoutOk() (*CreateAccountingInvoicePaymentRequestInputCheckout, bool) {
	if o == nil || IsNil(o.Checkout) {
		return nil, false
	}
	return o.Checkout, true
}

// HasCheckout returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasCheckout() bool {
	if o != nil && !IsNil(o.Checkout) {
		return true
	}

	return false
}

// SetCheckout gets a reference to the given CreateAccountingInvoicePaymentRequestInputCheckout and assigns it to the Checkout field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetCheckout(v CreateAccountingInvoicePaymentRequestInputCheckout) {
	o.Checkout = &v
}

// GetPayloadFingerprint returns the PayloadFingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAccountingInvoicePaymentRequestInput) GetPayloadFingerprint() string {
	if o == nil || IsNil(o.PayloadFingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.PayloadFingerprint.Get()
}

// GetPayloadFingerprintOk returns a tuple with the PayloadFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAccountingInvoicePaymentRequestInput) GetPayloadFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadFingerprint.Get(), o.PayloadFingerprint.IsSet()
}

// HasPayloadFingerprint returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasPayloadFingerprint() bool {
	if o != nil && o.PayloadFingerprint.IsSet() {
		return true
	}

	return false
}

// SetPayloadFingerprint gets a reference to the given NullableString and assigns it to the PayloadFingerprint field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetPayloadFingerprint(v string) {
	o.PayloadFingerprint.Set(&v)
}
// SetPayloadFingerprintNil sets the value for PayloadFingerprint to be an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) SetPayloadFingerprintNil() {
	o.PayloadFingerprint.Set(nil)
}

// UnsetPayloadFingerprint ensures that no value is present for PayloadFingerprint, not even an explicit nil
func (o *CreateAccountingInvoicePaymentRequestInput) UnsetPayloadFingerprint() {
	o.PayloadFingerprint.Unset()
}

// GetProviderMetadata returns the ProviderMetadata field value if set, zero value otherwise.
func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderMetadata() map[string]interface{} {
	if o == nil || IsNil(o.ProviderMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.ProviderMetadata
}

// GetProviderMetadataOk returns a tuple with the ProviderMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ProviderMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ProviderMetadata, true
}

// HasProviderMetadata returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasProviderMetadata() bool {
	if o != nil && !IsNil(o.ProviderMetadata) {
		return true
	}

	return false
}

// SetProviderMetadata gets a reference to the given map[string]interface{} and assigns it to the ProviderMetadata field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetProviderMetadata(v map[string]interface{}) {
	o.ProviderMetadata = v
}

// GetSafeProviderInvoiceSnapshot returns the SafeProviderInvoiceSnapshot field value if set, zero value otherwise.
func (o *CreateAccountingInvoicePaymentRequestInput) GetSafeProviderInvoiceSnapshot() map[string]interface{} {
	if o == nil || IsNil(o.SafeProviderInvoiceSnapshot) {
		var ret map[string]interface{}
		return ret
	}
	return o.SafeProviderInvoiceSnapshot
}

// GetSafeProviderInvoiceSnapshotOk returns a tuple with the SafeProviderInvoiceSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) GetSafeProviderInvoiceSnapshotOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.SafeProviderInvoiceSnapshot) {
		return map[string]interface{}{}, false
	}
	return o.SafeProviderInvoiceSnapshot, true
}

// HasSafeProviderInvoiceSnapshot returns a boolean if a field has been set.
func (o *CreateAccountingInvoicePaymentRequestInput) HasSafeProviderInvoiceSnapshot() bool {
	if o != nil && !IsNil(o.SafeProviderInvoiceSnapshot) {
		return true
	}

	return false
}

// SetSafeProviderInvoiceSnapshot gets a reference to the given map[string]interface{} and assigns it to the SafeProviderInvoiceSnapshot field.
func (o *CreateAccountingInvoicePaymentRequestInput) SetSafeProviderInvoiceSnapshot(v map[string]interface{}) {
	o.SafeProviderInvoiceSnapshot = v
}

func (o CreateAccountingInvoicePaymentRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountingInvoicePaymentRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["provider_account_id"] = o.ProviderAccountId
	toSerialize["external_invoice"] = o.ExternalInvoice
	if !IsNil(o.ExternalCustomer) {
		toSerialize["external_customer"] = o.ExternalCustomer
	}
	toSerialize["amount_due_minor"] = o.AmountDueMinor
	if o.TotalAmountMinor.IsSet() {
		toSerialize["total_amount_minor"] = o.TotalAmountMinor.Get()
	}
	if o.AmountPaidMinor.IsSet() {
		toSerialize["amount_paid_minor"] = o.AmountPaidMinor.Get()
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Checkout) {
		toSerialize["checkout"] = o.Checkout
	}
	if o.PayloadFingerprint.IsSet() {
		toSerialize["payload_fingerprint"] = o.PayloadFingerprint.Get()
	}
	if !IsNil(o.ProviderMetadata) {
		toSerialize["provider_metadata"] = o.ProviderMetadata
	}
	if !IsNil(o.SafeProviderInvoiceSnapshot) {
		toSerialize["safe_provider_invoice_snapshot"] = o.SafeProviderInvoiceSnapshot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAccountingInvoicePaymentRequestInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"connection_id",
		"provider_account_id",
		"external_invoice",
		"amount_due_minor",
		"currency",
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

	varCreateAccountingInvoicePaymentRequestInput := _CreateAccountingInvoicePaymentRequestInput{}

	err = json.Unmarshal(data, &varCreateAccountingInvoicePaymentRequestInput)

	if err != nil {
		return err
	}

	*o = CreateAccountingInvoicePaymentRequestInput(varCreateAccountingInvoicePaymentRequestInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provider")
		delete(additionalProperties, "connection_id")
		delete(additionalProperties, "provider_account_id")
		delete(additionalProperties, "external_invoice")
		delete(additionalProperties, "external_customer")
		delete(additionalProperties, "amount_due_minor")
		delete(additionalProperties, "total_amount_minor")
		delete(additionalProperties, "amount_paid_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "checkout")
		delete(additionalProperties, "payload_fingerprint")
		delete(additionalProperties, "provider_metadata")
		delete(additionalProperties, "safe_provider_invoice_snapshot")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAccountingInvoicePaymentRequestInput struct {
	value *CreateAccountingInvoicePaymentRequestInput
	isSet bool
}

func (v NullableCreateAccountingInvoicePaymentRequestInput) Get() *CreateAccountingInvoicePaymentRequestInput {
	return v.value
}

func (v *NullableCreateAccountingInvoicePaymentRequestInput) Set(val *CreateAccountingInvoicePaymentRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountingInvoicePaymentRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountingInvoicePaymentRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountingInvoicePaymentRequestInput(val *CreateAccountingInvoicePaymentRequestInput) *NullableCreateAccountingInvoicePaymentRequestInput {
	return &NullableCreateAccountingInvoicePaymentRequestInput{value: val, isSet: true}
}

func (v NullableCreateAccountingInvoicePaymentRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountingInvoicePaymentRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


