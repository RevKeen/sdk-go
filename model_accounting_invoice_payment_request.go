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

// checks if the AccountingInvoicePaymentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountingInvoicePaymentRequest{}

// AccountingInvoicePaymentRequest struct for AccountingInvoicePaymentRequest
type AccountingInvoicePaymentRequest struct {
	Id string `json:"id"`
	Object string `json:"object"`
	MerchantId string `json:"merchant_id"`
	Provider string `json:"provider"`
	ConnectionId string `json:"connection_id"`
	ProviderAccountId string `json:"provider_account_id"`
	ExternalInvoiceId string `json:"external_invoice_id"`
	ExternalInvoiceNumber NullableString `json:"external_invoice_number,omitempty"`
	ExternalInvoiceStatus NullableString `json:"external_invoice_status,omitempty"`
	ExternalInvoiceUrl NullableString `json:"external_invoice_url,omitempty"`
	ExternalInvoiceIssuedAt NullableString `json:"external_invoice_issued_at,omitempty"`
	ExternalInvoiceDueAt NullableString `json:"external_invoice_due_at,omitempty"`
	ExternalInvoiceUpdatedAt NullableString `json:"external_invoice_updated_at,omitempty"`
	ExternalCustomerId NullableString `json:"external_customer_id,omitempty"`
	ExternalCustomerReference NullableString `json:"external_customer_reference,omitempty"`
	ExternalCustomerName NullableString `json:"external_customer_name,omitempty"`
	ExternalCustomerEmail NullableString `json:"external_customer_email,omitempty"`
	AmountDueMinor int32 `json:"amount_due_minor"`
	TotalAmountMinor NullableInt32 `json:"total_amount_minor,omitempty"`
	AmountPaidMinor NullableInt32 `json:"amount_paid_minor,omitempty"`
	Currency string `json:"currency"`
	CheckoutSessionId NullableString `json:"checkout_session_id,omitempty"`
	CheckoutUrl NullableString `json:"checkout_url,omitempty"`
	CheckoutExpiresAt NullableString `json:"checkout_expires_at,omitempty"`
	CheckoutSuccessUrl NullableString `json:"checkout_success_url,omitempty"`
	CheckoutCancelUrl NullableString `json:"checkout_cancel_url,omitempty"`
	CheckoutAllowedMethods []string `json:"checkout_allowed_methods,omitempty"`
	Status string `json:"status"`
	SyncStatus string `json:"sync_status"`
	IdempotencyKey string `json:"idempotency_key"`
	PayloadFingerprint string `json:"payload_fingerprint"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _AccountingInvoicePaymentRequest AccountingInvoicePaymentRequest

// NewAccountingInvoicePaymentRequest instantiates a new AccountingInvoicePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountingInvoicePaymentRequest(id string, object string, merchantId string, provider string, connectionId string, providerAccountId string, externalInvoiceId string, amountDueMinor int32, currency string, status string, syncStatus string, idempotencyKey string, payloadFingerprint string, createdAt string, updatedAt string) *AccountingInvoicePaymentRequest {
	this := AccountingInvoicePaymentRequest{}
	this.Id = id
	this.Object = object
	this.MerchantId = merchantId
	this.Provider = provider
	this.ConnectionId = connectionId
	this.ProviderAccountId = providerAccountId
	this.ExternalInvoiceId = externalInvoiceId
	this.AmountDueMinor = amountDueMinor
	this.Currency = currency
	this.Status = status
	this.SyncStatus = syncStatus
	this.IdempotencyKey = idempotencyKey
	this.PayloadFingerprint = payloadFingerprint
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAccountingInvoicePaymentRequestWithDefaults instantiates a new AccountingInvoicePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountingInvoicePaymentRequestWithDefaults() *AccountingInvoicePaymentRequest {
	this := AccountingInvoicePaymentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *AccountingInvoicePaymentRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountingInvoicePaymentRequest) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *AccountingInvoicePaymentRequest) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AccountingInvoicePaymentRequest) SetObject(v string) {
	o.Object = v
}

// GetMerchantId returns the MerchantId field value
func (o *AccountingInvoicePaymentRequest) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *AccountingInvoicePaymentRequest) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetProvider returns the Provider field value
func (o *AccountingInvoicePaymentRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AccountingInvoicePaymentRequest) SetProvider(v string) {
	o.Provider = v
}

// GetConnectionId returns the ConnectionId field value
func (o *AccountingInvoicePaymentRequest) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *AccountingInvoicePaymentRequest) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetProviderAccountId returns the ProviderAccountId field value
func (o *AccountingInvoicePaymentRequest) GetProviderAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderAccountId
}

// GetProviderAccountIdOk returns a tuple with the ProviderAccountId field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetProviderAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderAccountId, true
}

// SetProviderAccountId sets field value
func (o *AccountingInvoicePaymentRequest) SetProviderAccountId(v string) {
	o.ProviderAccountId = v
}

// GetExternalInvoiceId returns the ExternalInvoiceId field value
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalInvoiceId
}

// GetExternalInvoiceIdOk returns a tuple with the ExternalInvoiceId field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalInvoiceId, true
}

// SetExternalInvoiceId sets field value
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceId(v string) {
	o.ExternalInvoiceId = v
}

// GetExternalInvoiceNumber returns the ExternalInvoiceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceNumber() string {
	if o == nil || IsNil(o.ExternalInvoiceNumber.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceNumber.Get()
}

// GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceNumber.Get(), o.ExternalInvoiceNumber.IsSet()
}

// HasExternalInvoiceNumber returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceNumber() bool {
	if o != nil && o.ExternalInvoiceNumber.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceNumber gets a reference to the given NullableString and assigns it to the ExternalInvoiceNumber field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceNumber(v string) {
	o.ExternalInvoiceNumber.Set(&v)
}
// SetExternalInvoiceNumberNil sets the value for ExternalInvoiceNumber to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceNumberNil() {
	o.ExternalInvoiceNumber.Set(nil)
}

// UnsetExternalInvoiceNumber ensures that no value is present for ExternalInvoiceNumber, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceNumber() {
	o.ExternalInvoiceNumber.Unset()
}

// GetExternalInvoiceStatus returns the ExternalInvoiceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceStatus() string {
	if o == nil || IsNil(o.ExternalInvoiceStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceStatus.Get()
}

// GetExternalInvoiceStatusOk returns a tuple with the ExternalInvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceStatus.Get(), o.ExternalInvoiceStatus.IsSet()
}

// HasExternalInvoiceStatus returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceStatus() bool {
	if o != nil && o.ExternalInvoiceStatus.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceStatus gets a reference to the given NullableString and assigns it to the ExternalInvoiceStatus field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceStatus(v string) {
	o.ExternalInvoiceStatus.Set(&v)
}
// SetExternalInvoiceStatusNil sets the value for ExternalInvoiceStatus to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceStatusNil() {
	o.ExternalInvoiceStatus.Set(nil)
}

// UnsetExternalInvoiceStatus ensures that no value is present for ExternalInvoiceStatus, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceStatus() {
	o.ExternalInvoiceStatus.Unset()
}

// GetExternalInvoiceUrl returns the ExternalInvoiceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUrl() string {
	if o == nil || IsNil(o.ExternalInvoiceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceUrl.Get()
}

// GetExternalInvoiceUrlOk returns a tuple with the ExternalInvoiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceUrl.Get(), o.ExternalInvoiceUrl.IsSet()
}

// HasExternalInvoiceUrl returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceUrl() bool {
	if o != nil && o.ExternalInvoiceUrl.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceUrl gets a reference to the given NullableString and assigns it to the ExternalInvoiceUrl field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUrl(v string) {
	o.ExternalInvoiceUrl.Set(&v)
}
// SetExternalInvoiceUrlNil sets the value for ExternalInvoiceUrl to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUrlNil() {
	o.ExternalInvoiceUrl.Set(nil)
}

// UnsetExternalInvoiceUrl ensures that no value is present for ExternalInvoiceUrl, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceUrl() {
	o.ExternalInvoiceUrl.Unset()
}

// GetExternalInvoiceIssuedAt returns the ExternalInvoiceIssuedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIssuedAt() string {
	if o == nil || IsNil(o.ExternalInvoiceIssuedAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceIssuedAt.Get()
}

// GetExternalInvoiceIssuedAtOk returns a tuple with the ExternalInvoiceIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIssuedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceIssuedAt.Get(), o.ExternalInvoiceIssuedAt.IsSet()
}

// HasExternalInvoiceIssuedAt returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceIssuedAt() bool {
	if o != nil && o.ExternalInvoiceIssuedAt.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceIssuedAt gets a reference to the given NullableString and assigns it to the ExternalInvoiceIssuedAt field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceIssuedAt(v string) {
	o.ExternalInvoiceIssuedAt.Set(&v)
}
// SetExternalInvoiceIssuedAtNil sets the value for ExternalInvoiceIssuedAt to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceIssuedAtNil() {
	o.ExternalInvoiceIssuedAt.Set(nil)
}

// UnsetExternalInvoiceIssuedAt ensures that no value is present for ExternalInvoiceIssuedAt, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceIssuedAt() {
	o.ExternalInvoiceIssuedAt.Unset()
}

// GetExternalInvoiceDueAt returns the ExternalInvoiceDueAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceDueAt() string {
	if o == nil || IsNil(o.ExternalInvoiceDueAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceDueAt.Get()
}

// GetExternalInvoiceDueAtOk returns a tuple with the ExternalInvoiceDueAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceDueAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceDueAt.Get(), o.ExternalInvoiceDueAt.IsSet()
}

// HasExternalInvoiceDueAt returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceDueAt() bool {
	if o != nil && o.ExternalInvoiceDueAt.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceDueAt gets a reference to the given NullableString and assigns it to the ExternalInvoiceDueAt field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceDueAt(v string) {
	o.ExternalInvoiceDueAt.Set(&v)
}
// SetExternalInvoiceDueAtNil sets the value for ExternalInvoiceDueAt to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceDueAtNil() {
	o.ExternalInvoiceDueAt.Set(nil)
}

// UnsetExternalInvoiceDueAt ensures that no value is present for ExternalInvoiceDueAt, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceDueAt() {
	o.ExternalInvoiceDueAt.Unset()
}

// GetExternalInvoiceUpdatedAt returns the ExternalInvoiceUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUpdatedAt() string {
	if o == nil || IsNil(o.ExternalInvoiceUpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceUpdatedAt.Get()
}

// GetExternalInvoiceUpdatedAtOk returns a tuple with the ExternalInvoiceUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalInvoiceUpdatedAt.Get(), o.ExternalInvoiceUpdatedAt.IsSet()
}

// HasExternalInvoiceUpdatedAt returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceUpdatedAt() bool {
	if o != nil && o.ExternalInvoiceUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetExternalInvoiceUpdatedAt gets a reference to the given NullableString and assigns it to the ExternalInvoiceUpdatedAt field.
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUpdatedAt(v string) {
	o.ExternalInvoiceUpdatedAt.Set(&v)
}
// SetExternalInvoiceUpdatedAtNil sets the value for ExternalInvoiceUpdatedAt to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUpdatedAtNil() {
	o.ExternalInvoiceUpdatedAt.Set(nil)
}

// UnsetExternalInvoiceUpdatedAt ensures that no value is present for ExternalInvoiceUpdatedAt, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceUpdatedAt() {
	o.ExternalInvoiceUpdatedAt.Unset()
}

// GetExternalCustomerId returns the ExternalCustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerId() string {
	if o == nil || IsNil(o.ExternalCustomerId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerId.Get()
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalCustomerId.Get(), o.ExternalCustomerId.IsSet()
}

// HasExternalCustomerId returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalCustomerId() bool {
	if o != nil && o.ExternalCustomerId.IsSet() {
		return true
	}

	return false
}

// SetExternalCustomerId gets a reference to the given NullableString and assigns it to the ExternalCustomerId field.
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerId(v string) {
	o.ExternalCustomerId.Set(&v)
}
// SetExternalCustomerIdNil sets the value for ExternalCustomerId to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerIdNil() {
	o.ExternalCustomerId.Set(nil)
}

// UnsetExternalCustomerId ensures that no value is present for ExternalCustomerId, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerId() {
	o.ExternalCustomerId.Unset()
}

// GetExternalCustomerReference returns the ExternalCustomerReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerReference() string {
	if o == nil || IsNil(o.ExternalCustomerReference.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerReference.Get()
}

// GetExternalCustomerReferenceOk returns a tuple with the ExternalCustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalCustomerReference.Get(), o.ExternalCustomerReference.IsSet()
}

// HasExternalCustomerReference returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalCustomerReference() bool {
	if o != nil && o.ExternalCustomerReference.IsSet() {
		return true
	}

	return false
}

// SetExternalCustomerReference gets a reference to the given NullableString and assigns it to the ExternalCustomerReference field.
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerReference(v string) {
	o.ExternalCustomerReference.Set(&v)
}
// SetExternalCustomerReferenceNil sets the value for ExternalCustomerReference to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerReferenceNil() {
	o.ExternalCustomerReference.Set(nil)
}

// UnsetExternalCustomerReference ensures that no value is present for ExternalCustomerReference, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerReference() {
	o.ExternalCustomerReference.Unset()
}

// GetExternalCustomerName returns the ExternalCustomerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerName() string {
	if o == nil || IsNil(o.ExternalCustomerName.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerName.Get()
}

// GetExternalCustomerNameOk returns a tuple with the ExternalCustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalCustomerName.Get(), o.ExternalCustomerName.IsSet()
}

// HasExternalCustomerName returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalCustomerName() bool {
	if o != nil && o.ExternalCustomerName.IsSet() {
		return true
	}

	return false
}

// SetExternalCustomerName gets a reference to the given NullableString and assigns it to the ExternalCustomerName field.
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerName(v string) {
	o.ExternalCustomerName.Set(&v)
}
// SetExternalCustomerNameNil sets the value for ExternalCustomerName to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerNameNil() {
	o.ExternalCustomerName.Set(nil)
}

// UnsetExternalCustomerName ensures that no value is present for ExternalCustomerName, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerName() {
	o.ExternalCustomerName.Unset()
}

// GetExternalCustomerEmail returns the ExternalCustomerEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerEmail() string {
	if o == nil || IsNil(o.ExternalCustomerEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerEmail.Get()
}

// GetExternalCustomerEmailOk returns a tuple with the ExternalCustomerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetExternalCustomerEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalCustomerEmail.Get(), o.ExternalCustomerEmail.IsSet()
}

// HasExternalCustomerEmail returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasExternalCustomerEmail() bool {
	if o != nil && o.ExternalCustomerEmail.IsSet() {
		return true
	}

	return false
}

// SetExternalCustomerEmail gets a reference to the given NullableString and assigns it to the ExternalCustomerEmail field.
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerEmail(v string) {
	o.ExternalCustomerEmail.Set(&v)
}
// SetExternalCustomerEmailNil sets the value for ExternalCustomerEmail to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetExternalCustomerEmailNil() {
	o.ExternalCustomerEmail.Set(nil)
}

// UnsetExternalCustomerEmail ensures that no value is present for ExternalCustomerEmail, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerEmail() {
	o.ExternalCustomerEmail.Unset()
}

// GetAmountDueMinor returns the AmountDueMinor field value
func (o *AccountingInvoicePaymentRequest) GetAmountDueMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountDueMinor
}

// GetAmountDueMinorOk returns a tuple with the AmountDueMinor field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetAmountDueMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountDueMinor, true
}

// SetAmountDueMinor sets field value
func (o *AccountingInvoicePaymentRequest) SetAmountDueMinor(v int32) {
	o.AmountDueMinor = v
}

// GetTotalAmountMinor returns the TotalAmountMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetTotalAmountMinor() int32 {
	if o == nil || IsNil(o.TotalAmountMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalAmountMinor.Get()
}

// GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetTotalAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalAmountMinor.Get(), o.TotalAmountMinor.IsSet()
}

// HasTotalAmountMinor returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasTotalAmountMinor() bool {
	if o != nil && o.TotalAmountMinor.IsSet() {
		return true
	}

	return false
}

// SetTotalAmountMinor gets a reference to the given NullableInt32 and assigns it to the TotalAmountMinor field.
func (o *AccountingInvoicePaymentRequest) SetTotalAmountMinor(v int32) {
	o.TotalAmountMinor.Set(&v)
}
// SetTotalAmountMinorNil sets the value for TotalAmountMinor to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetTotalAmountMinorNil() {
	o.TotalAmountMinor.Set(nil)
}

// UnsetTotalAmountMinor ensures that no value is present for TotalAmountMinor, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetTotalAmountMinor() {
	o.TotalAmountMinor.Unset()
}

// GetAmountPaidMinor returns the AmountPaidMinor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetAmountPaidMinor() int32 {
	if o == nil || IsNil(o.AmountPaidMinor.Get()) {
		var ret int32
		return ret
	}
	return *o.AmountPaidMinor.Get()
}

// GetAmountPaidMinorOk returns a tuple with the AmountPaidMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetAmountPaidMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountPaidMinor.Get(), o.AmountPaidMinor.IsSet()
}

// HasAmountPaidMinor returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasAmountPaidMinor() bool {
	if o != nil && o.AmountPaidMinor.IsSet() {
		return true
	}

	return false
}

// SetAmountPaidMinor gets a reference to the given NullableInt32 and assigns it to the AmountPaidMinor field.
func (o *AccountingInvoicePaymentRequest) SetAmountPaidMinor(v int32) {
	o.AmountPaidMinor.Set(&v)
}
// SetAmountPaidMinorNil sets the value for AmountPaidMinor to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetAmountPaidMinorNil() {
	o.AmountPaidMinor.Set(nil)
}

// UnsetAmountPaidMinor ensures that no value is present for AmountPaidMinor, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetAmountPaidMinor() {
	o.AmountPaidMinor.Unset()
}

// GetCurrency returns the Currency field value
func (o *AccountingInvoicePaymentRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AccountingInvoicePaymentRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCheckoutSessionId returns the CheckoutSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutSessionId() string {
	if o == nil || IsNil(o.CheckoutSessionId.Get()) {
		var ret string
		return ret
	}
	return *o.CheckoutSessionId.Get()
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutSessionId.Get(), o.CheckoutSessionId.IsSet()
}

// HasCheckoutSessionId returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutSessionId() bool {
	if o != nil && o.CheckoutSessionId.IsSet() {
		return true
	}

	return false
}

// SetCheckoutSessionId gets a reference to the given NullableString and assigns it to the CheckoutSessionId field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId.Set(&v)
}
// SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetCheckoutSessionIdNil() {
	o.CheckoutSessionId.Set(nil)
}

// UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetCheckoutSessionId() {
	o.CheckoutSessionId.Unset()
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutUrl() string {
	if o == nil || IsNil(o.CheckoutUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CheckoutUrl.Get()
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutUrl.Get(), o.CheckoutUrl.IsSet()
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutUrl() bool {
	if o != nil && o.CheckoutUrl.IsSet() {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given NullableString and assigns it to the CheckoutUrl field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutUrl(v string) {
	o.CheckoutUrl.Set(&v)
}
// SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetCheckoutUrlNil() {
	o.CheckoutUrl.Set(nil)
}

// UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetCheckoutUrl() {
	o.CheckoutUrl.Unset()
}

// GetCheckoutExpiresAt returns the CheckoutExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutExpiresAt() string {
	if o == nil || IsNil(o.CheckoutExpiresAt.Get()) {
		var ret string
		return ret
	}
	return *o.CheckoutExpiresAt.Get()
}

// GetCheckoutExpiresAtOk returns a tuple with the CheckoutExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutExpiresAt.Get(), o.CheckoutExpiresAt.IsSet()
}

// HasCheckoutExpiresAt returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutExpiresAt() bool {
	if o != nil && o.CheckoutExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetCheckoutExpiresAt gets a reference to the given NullableString and assigns it to the CheckoutExpiresAt field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutExpiresAt(v string) {
	o.CheckoutExpiresAt.Set(&v)
}
// SetCheckoutExpiresAtNil sets the value for CheckoutExpiresAt to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetCheckoutExpiresAtNil() {
	o.CheckoutExpiresAt.Set(nil)
}

// UnsetCheckoutExpiresAt ensures that no value is present for CheckoutExpiresAt, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetCheckoutExpiresAt() {
	o.CheckoutExpiresAt.Unset()
}

// GetCheckoutSuccessUrl returns the CheckoutSuccessUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutSuccessUrl() string {
	if o == nil || IsNil(o.CheckoutSuccessUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CheckoutSuccessUrl.Get()
}

// GetCheckoutSuccessUrlOk returns a tuple with the CheckoutSuccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutSuccessUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutSuccessUrl.Get(), o.CheckoutSuccessUrl.IsSet()
}

// HasCheckoutSuccessUrl returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutSuccessUrl() bool {
	if o != nil && o.CheckoutSuccessUrl.IsSet() {
		return true
	}

	return false
}

// SetCheckoutSuccessUrl gets a reference to the given NullableString and assigns it to the CheckoutSuccessUrl field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutSuccessUrl(v string) {
	o.CheckoutSuccessUrl.Set(&v)
}
// SetCheckoutSuccessUrlNil sets the value for CheckoutSuccessUrl to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetCheckoutSuccessUrlNil() {
	o.CheckoutSuccessUrl.Set(nil)
}

// UnsetCheckoutSuccessUrl ensures that no value is present for CheckoutSuccessUrl, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetCheckoutSuccessUrl() {
	o.CheckoutSuccessUrl.Unset()
}

// GetCheckoutCancelUrl returns the CheckoutCancelUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutCancelUrl() string {
	if o == nil || IsNil(o.CheckoutCancelUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CheckoutCancelUrl.Get()
}

// GetCheckoutCancelUrlOk returns a tuple with the CheckoutCancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutCancelUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutCancelUrl.Get(), o.CheckoutCancelUrl.IsSet()
}

// HasCheckoutCancelUrl returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutCancelUrl() bool {
	if o != nil && o.CheckoutCancelUrl.IsSet() {
		return true
	}

	return false
}

// SetCheckoutCancelUrl gets a reference to the given NullableString and assigns it to the CheckoutCancelUrl field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutCancelUrl(v string) {
	o.CheckoutCancelUrl.Set(&v)
}
// SetCheckoutCancelUrlNil sets the value for CheckoutCancelUrl to be an explicit nil
func (o *AccountingInvoicePaymentRequest) SetCheckoutCancelUrlNil() {
	o.CheckoutCancelUrl.Set(nil)
}

// UnsetCheckoutCancelUrl ensures that no value is present for CheckoutCancelUrl, not even an explicit nil
func (o *AccountingInvoicePaymentRequest) UnsetCheckoutCancelUrl() {
	o.CheckoutCancelUrl.Unset()
}

// GetCheckoutAllowedMethods returns the CheckoutAllowedMethods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountingInvoicePaymentRequest) GetCheckoutAllowedMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CheckoutAllowedMethods
}

// GetCheckoutAllowedMethodsOk returns a tuple with the CheckoutAllowedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountingInvoicePaymentRequest) GetCheckoutAllowedMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.CheckoutAllowedMethods) {
		return nil, false
	}
	return o.CheckoutAllowedMethods, true
}

// HasCheckoutAllowedMethods returns a boolean if a field has been set.
func (o *AccountingInvoicePaymentRequest) HasCheckoutAllowedMethods() bool {
	if o != nil && !IsNil(o.CheckoutAllowedMethods) {
		return true
	}

	return false
}

// SetCheckoutAllowedMethods gets a reference to the given []string and assigns it to the CheckoutAllowedMethods field.
func (o *AccountingInvoicePaymentRequest) SetCheckoutAllowedMethods(v []string) {
	o.CheckoutAllowedMethods = v
}

// GetStatus returns the Status field value
func (o *AccountingInvoicePaymentRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountingInvoicePaymentRequest) SetStatus(v string) {
	o.Status = v
}

// GetSyncStatus returns the SyncStatus field value
func (o *AccountingInvoicePaymentRequest) GetSyncStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetSyncStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncStatus, true
}

// SetSyncStatus sets field value
func (o *AccountingInvoicePaymentRequest) SetSyncStatus(v string) {
	o.SyncStatus = v
}

// GetIdempotencyKey returns the IdempotencyKey field value
func (o *AccountingInvoicePaymentRequest) GetIdempotencyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdempotencyKey, true
}

// SetIdempotencyKey sets field value
func (o *AccountingInvoicePaymentRequest) SetIdempotencyKey(v string) {
	o.IdempotencyKey = v
}

// GetPayloadFingerprint returns the PayloadFingerprint field value
func (o *AccountingInvoicePaymentRequest) GetPayloadFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadFingerprint
}

// GetPayloadFingerprintOk returns a tuple with the PayloadFingerprint field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetPayloadFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadFingerprint, true
}

// SetPayloadFingerprint sets field value
func (o *AccountingInvoicePaymentRequest) SetPayloadFingerprint(v string) {
	o.PayloadFingerprint = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccountingInvoicePaymentRequest) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccountingInvoicePaymentRequest) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AccountingInvoicePaymentRequest) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountingInvoicePaymentRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AccountingInvoicePaymentRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o AccountingInvoicePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountingInvoicePaymentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["provider"] = o.Provider
	toSerialize["connection_id"] = o.ConnectionId
	toSerialize["provider_account_id"] = o.ProviderAccountId
	toSerialize["external_invoice_id"] = o.ExternalInvoiceId
	if o.ExternalInvoiceNumber.IsSet() {
		toSerialize["external_invoice_number"] = o.ExternalInvoiceNumber.Get()
	}
	if o.ExternalInvoiceStatus.IsSet() {
		toSerialize["external_invoice_status"] = o.ExternalInvoiceStatus.Get()
	}
	if o.ExternalInvoiceUrl.IsSet() {
		toSerialize["external_invoice_url"] = o.ExternalInvoiceUrl.Get()
	}
	if o.ExternalInvoiceIssuedAt.IsSet() {
		toSerialize["external_invoice_issued_at"] = o.ExternalInvoiceIssuedAt.Get()
	}
	if o.ExternalInvoiceDueAt.IsSet() {
		toSerialize["external_invoice_due_at"] = o.ExternalInvoiceDueAt.Get()
	}
	if o.ExternalInvoiceUpdatedAt.IsSet() {
		toSerialize["external_invoice_updated_at"] = o.ExternalInvoiceUpdatedAt.Get()
	}
	if o.ExternalCustomerId.IsSet() {
		toSerialize["external_customer_id"] = o.ExternalCustomerId.Get()
	}
	if o.ExternalCustomerReference.IsSet() {
		toSerialize["external_customer_reference"] = o.ExternalCustomerReference.Get()
	}
	if o.ExternalCustomerName.IsSet() {
		toSerialize["external_customer_name"] = o.ExternalCustomerName.Get()
	}
	if o.ExternalCustomerEmail.IsSet() {
		toSerialize["external_customer_email"] = o.ExternalCustomerEmail.Get()
	}
	toSerialize["amount_due_minor"] = o.AmountDueMinor
	if o.TotalAmountMinor.IsSet() {
		toSerialize["total_amount_minor"] = o.TotalAmountMinor.Get()
	}
	if o.AmountPaidMinor.IsSet() {
		toSerialize["amount_paid_minor"] = o.AmountPaidMinor.Get()
	}
	toSerialize["currency"] = o.Currency
	if o.CheckoutSessionId.IsSet() {
		toSerialize["checkout_session_id"] = o.CheckoutSessionId.Get()
	}
	if o.CheckoutUrl.IsSet() {
		toSerialize["checkout_url"] = o.CheckoutUrl.Get()
	}
	if o.CheckoutExpiresAt.IsSet() {
		toSerialize["checkout_expires_at"] = o.CheckoutExpiresAt.Get()
	}
	if o.CheckoutSuccessUrl.IsSet() {
		toSerialize["checkout_success_url"] = o.CheckoutSuccessUrl.Get()
	}
	if o.CheckoutCancelUrl.IsSet() {
		toSerialize["checkout_cancel_url"] = o.CheckoutCancelUrl.Get()
	}
	if o.CheckoutAllowedMethods != nil {
		toSerialize["checkout_allowed_methods"] = o.CheckoutAllowedMethods
	}
	toSerialize["status"] = o.Status
	toSerialize["sync_status"] = o.SyncStatus
	toSerialize["idempotency_key"] = o.IdempotencyKey
	toSerialize["payload_fingerprint"] = o.PayloadFingerprint
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountingInvoicePaymentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"merchant_id",
		"provider",
		"connection_id",
		"provider_account_id",
		"external_invoice_id",
		"amount_due_minor",
		"currency",
		"status",
		"sync_status",
		"idempotency_key",
		"payload_fingerprint",
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

	varAccountingInvoicePaymentRequest := _AccountingInvoicePaymentRequest{}

	err = json.Unmarshal(data, &varAccountingInvoicePaymentRequest)

	if err != nil {
		return err
	}

	*o = AccountingInvoicePaymentRequest(varAccountingInvoicePaymentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "connection_id")
		delete(additionalProperties, "provider_account_id")
		delete(additionalProperties, "external_invoice_id")
		delete(additionalProperties, "external_invoice_number")
		delete(additionalProperties, "external_invoice_status")
		delete(additionalProperties, "external_invoice_url")
		delete(additionalProperties, "external_invoice_issued_at")
		delete(additionalProperties, "external_invoice_due_at")
		delete(additionalProperties, "external_invoice_updated_at")
		delete(additionalProperties, "external_customer_id")
		delete(additionalProperties, "external_customer_reference")
		delete(additionalProperties, "external_customer_name")
		delete(additionalProperties, "external_customer_email")
		delete(additionalProperties, "amount_due_minor")
		delete(additionalProperties, "total_amount_minor")
		delete(additionalProperties, "amount_paid_minor")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "checkout_session_id")
		delete(additionalProperties, "checkout_url")
		delete(additionalProperties, "checkout_expires_at")
		delete(additionalProperties, "checkout_success_url")
		delete(additionalProperties, "checkout_cancel_url")
		delete(additionalProperties, "checkout_allowed_methods")
		delete(additionalProperties, "status")
		delete(additionalProperties, "sync_status")
		delete(additionalProperties, "idempotency_key")
		delete(additionalProperties, "payload_fingerprint")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountingInvoicePaymentRequest struct {
	value *AccountingInvoicePaymentRequest
	isSet bool
}

func (v NullableAccountingInvoicePaymentRequest) Get() *AccountingInvoicePaymentRequest {
	return v.value
}

func (v *NullableAccountingInvoicePaymentRequest) Set(val *AccountingInvoicePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountingInvoicePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountingInvoicePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountingInvoicePaymentRequest(val *AccountingInvoicePaymentRequest) *NullableAccountingInvoicePaymentRequest {
	return &NullableAccountingInvoicePaymentRequest{value: val, isSet: true}
}

func (v NullableAccountingInvoicePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountingInvoicePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


