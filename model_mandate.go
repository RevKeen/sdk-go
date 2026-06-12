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

// checks if the Mandate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Mandate{}

// Mandate A Direct Debit mandate (customer authorisation to collect via Bacs).
type Mandate struct {
	// Mandate ID
	Id string `json:"id"`
	// RevKeen mandate reference (Bacs DDI reference)
	MandateRef string `json:"mandate_ref"`
	// Human-facing mandate reference
	MandateReference NullableString `json:"mandate_reference"`
	// Customer the mandate belongs to
	CustomerId string `json:"customer_id"`
	// Mandate status
	Status string `json:"status"`
	// Name on the bank account
	AccountHolderName NullableString `json:"account_holder_name"`
	// Masked sort code — never the full value
	SortCode string `json:"sort_code"`
	// Last 4 digits of the account number
	AccountNumberLast4 NullableString `json:"account_number_last4"`
	// Resolved bank name
	BankName NullableString `json:"bank_name"`
	// Bacs statement descriptor (max 18 chars)
	StatementName NullableString `json:"statement_name,omitempty"`
	// Advance-notice days applied
	NoticeDays NullableInt32 `json:"notice_days"`
	// Earliest collection date (YYYY-MM-DD)
	FirstCollectionDate NullableString `json:"first_collection_date,omitempty"`
	// Next scheduled collection date (YYYY-MM-DD)
	NextCollectionDate NullableString `json:"next_collection_date,omitempty"`
	// Recovery fallback card
	BackupPaymentMethodId NullableString `json:"backup_payment_method_id,omitempty"`
	// When the mandate became active (ISO 8601)
	ActivatedAt NullableString `json:"activated_at,omitempty"`
	// When the mandate was suspended (ISO 8601)
	SuspendedAt NullableString `json:"suspended_at,omitempty"`
	// When the mandate was cancelled (ISO 8601)
	CancelledAt NullableString `json:"cancelled_at,omitempty"`
	// Most recent failure reason
	FailureReason NullableString `json:"failure_reason,omitempty"`
	// Mandate-request consumed on creation, if any
	MandateRequestId NullableString `json:"mandate_request_id,omitempty"`
	// Invoice linked via the mandate-request, if any
	InvoiceId NullableString `json:"invoice_id,omitempty"`
	// Creation timestamp (ISO 8601)
	CreatedAt NullableString `json:"created_at,omitempty"`
	// Last-updated timestamp (ISO 8601)
	UpdatedAt NullableString `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Mandate Mandate

// NewMandate instantiates a new Mandate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMandate(id string, mandateRef string, mandateReference NullableString, customerId string, status string, accountHolderName NullableString, sortCode string, accountNumberLast4 NullableString, bankName NullableString, noticeDays NullableInt32) *Mandate {
	this := Mandate{}
	this.Id = id
	this.MandateRef = mandateRef
	this.MandateReference = mandateReference
	this.CustomerId = customerId
	this.Status = status
	this.AccountHolderName = accountHolderName
	this.SortCode = sortCode
	this.AccountNumberLast4 = accountNumberLast4
	this.BankName = bankName
	this.NoticeDays = noticeDays
	return &this
}

// NewMandateWithDefaults instantiates a new Mandate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMandateWithDefaults() *Mandate {
	this := Mandate{}
	return &this
}

// GetId returns the Id field value
func (o *Mandate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Mandate) SetId(v string) {
	o.Id = v
}

// GetMandateRef returns the MandateRef field value
func (o *Mandate) GetMandateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateRef
}

// GetMandateRefOk returns a tuple with the MandateRef field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetMandateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateRef, true
}

// SetMandateRef sets field value
func (o *Mandate) SetMandateRef(v string) {
	o.MandateRef = v
}

// GetMandateReference returns the MandateReference field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Mandate) GetMandateReference() string {
	if o == nil || o.MandateReference.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateReference.Get()
}

// GetMandateReferenceOk returns a tuple with the MandateReference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetMandateReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateReference.Get(), o.MandateReference.IsSet()
}

// SetMandateReference sets field value
func (o *Mandate) SetMandateReference(v string) {
	o.MandateReference.Set(&v)
}

// GetCustomerId returns the CustomerId field value
func (o *Mandate) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *Mandate) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetStatus returns the Status field value
func (o *Mandate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Mandate) SetStatus(v string) {
	o.Status = v
}

// GetAccountHolderName returns the AccountHolderName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Mandate) GetAccountHolderName() string {
	if o == nil || o.AccountHolderName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// SetAccountHolderName sets field value
func (o *Mandate) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}

// GetSortCode returns the SortCode field value
func (o *Mandate) GetSortCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *Mandate) GetSortCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *Mandate) SetSortCode(v string) {
	o.SortCode = v
}

// GetAccountNumberLast4 returns the AccountNumberLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Mandate) GetAccountNumberLast4() string {
	if o == nil || o.AccountNumberLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountNumberLast4.Get()
}

// GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetAccountNumberLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumberLast4.Get(), o.AccountNumberLast4.IsSet()
}

// SetAccountNumberLast4 sets field value
func (o *Mandate) SetAccountNumberLast4(v string) {
	o.AccountNumberLast4.Set(&v)
}

// GetBankName returns the BankName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Mandate) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// SetBankName sets field value
func (o *Mandate) SetBankName(v string) {
	o.BankName.Set(&v)
}

// GetStatementName returns the StatementName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetStatementName() string {
	if o == nil || IsNil(o.StatementName.Get()) {
		var ret string
		return ret
	}
	return *o.StatementName.Get()
}

// GetStatementNameOk returns a tuple with the StatementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetStatementNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatementName.Get(), o.StatementName.IsSet()
}

// HasStatementName returns a boolean if a field has been set.
func (o *Mandate) HasStatementName() bool {
	if o != nil && o.StatementName.IsSet() {
		return true
	}

	return false
}

// SetStatementName gets a reference to the given NullableString and assigns it to the StatementName field.
func (o *Mandate) SetStatementName(v string) {
	o.StatementName.Set(&v)
}
// SetStatementNameNil sets the value for StatementName to be an explicit nil
func (o *Mandate) SetStatementNameNil() {
	o.StatementName.Set(nil)
}

// UnsetStatementName ensures that no value is present for StatementName, not even an explicit nil
func (o *Mandate) UnsetStatementName() {
	o.StatementName.Unset()
}

// GetNoticeDays returns the NoticeDays field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Mandate) GetNoticeDays() int32 {
	if o == nil || o.NoticeDays.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NoticeDays.Get()
}

// GetNoticeDaysOk returns a tuple with the NoticeDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetNoticeDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoticeDays.Get(), o.NoticeDays.IsSet()
}

// SetNoticeDays sets field value
func (o *Mandate) SetNoticeDays(v int32) {
	o.NoticeDays.Set(&v)
}

// GetFirstCollectionDate returns the FirstCollectionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetFirstCollectionDate() string {
	if o == nil || IsNil(o.FirstCollectionDate.Get()) {
		var ret string
		return ret
	}
	return *o.FirstCollectionDate.Get()
}

// GetFirstCollectionDateOk returns a tuple with the FirstCollectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetFirstCollectionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstCollectionDate.Get(), o.FirstCollectionDate.IsSet()
}

// HasFirstCollectionDate returns a boolean if a field has been set.
func (o *Mandate) HasFirstCollectionDate() bool {
	if o != nil && o.FirstCollectionDate.IsSet() {
		return true
	}

	return false
}

// SetFirstCollectionDate gets a reference to the given NullableString and assigns it to the FirstCollectionDate field.
func (o *Mandate) SetFirstCollectionDate(v string) {
	o.FirstCollectionDate.Set(&v)
}
// SetFirstCollectionDateNil sets the value for FirstCollectionDate to be an explicit nil
func (o *Mandate) SetFirstCollectionDateNil() {
	o.FirstCollectionDate.Set(nil)
}

// UnsetFirstCollectionDate ensures that no value is present for FirstCollectionDate, not even an explicit nil
func (o *Mandate) UnsetFirstCollectionDate() {
	o.FirstCollectionDate.Unset()
}

// GetNextCollectionDate returns the NextCollectionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetNextCollectionDate() string {
	if o == nil || IsNil(o.NextCollectionDate.Get()) {
		var ret string
		return ret
	}
	return *o.NextCollectionDate.Get()
}

// GetNextCollectionDateOk returns a tuple with the NextCollectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetNextCollectionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCollectionDate.Get(), o.NextCollectionDate.IsSet()
}

// HasNextCollectionDate returns a boolean if a field has been set.
func (o *Mandate) HasNextCollectionDate() bool {
	if o != nil && o.NextCollectionDate.IsSet() {
		return true
	}

	return false
}

// SetNextCollectionDate gets a reference to the given NullableString and assigns it to the NextCollectionDate field.
func (o *Mandate) SetNextCollectionDate(v string) {
	o.NextCollectionDate.Set(&v)
}
// SetNextCollectionDateNil sets the value for NextCollectionDate to be an explicit nil
func (o *Mandate) SetNextCollectionDateNil() {
	o.NextCollectionDate.Set(nil)
}

// UnsetNextCollectionDate ensures that no value is present for NextCollectionDate, not even an explicit nil
func (o *Mandate) UnsetNextCollectionDate() {
	o.NextCollectionDate.Unset()
}

// GetBackupPaymentMethodId returns the BackupPaymentMethodId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetBackupPaymentMethodId() string {
	if o == nil || IsNil(o.BackupPaymentMethodId.Get()) {
		var ret string
		return ret
	}
	return *o.BackupPaymentMethodId.Get()
}

// GetBackupPaymentMethodIdOk returns a tuple with the BackupPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetBackupPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupPaymentMethodId.Get(), o.BackupPaymentMethodId.IsSet()
}

// HasBackupPaymentMethodId returns a boolean if a field has been set.
func (o *Mandate) HasBackupPaymentMethodId() bool {
	if o != nil && o.BackupPaymentMethodId.IsSet() {
		return true
	}

	return false
}

// SetBackupPaymentMethodId gets a reference to the given NullableString and assigns it to the BackupPaymentMethodId field.
func (o *Mandate) SetBackupPaymentMethodId(v string) {
	o.BackupPaymentMethodId.Set(&v)
}
// SetBackupPaymentMethodIdNil sets the value for BackupPaymentMethodId to be an explicit nil
func (o *Mandate) SetBackupPaymentMethodIdNil() {
	o.BackupPaymentMethodId.Set(nil)
}

// UnsetBackupPaymentMethodId ensures that no value is present for BackupPaymentMethodId, not even an explicit nil
func (o *Mandate) UnsetBackupPaymentMethodId() {
	o.BackupPaymentMethodId.Unset()
}

// GetActivatedAt returns the ActivatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetActivatedAt() string {
	if o == nil || IsNil(o.ActivatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.ActivatedAt.Get()
}

// GetActivatedAtOk returns a tuple with the ActivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetActivatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActivatedAt.Get(), o.ActivatedAt.IsSet()
}

// HasActivatedAt returns a boolean if a field has been set.
func (o *Mandate) HasActivatedAt() bool {
	if o != nil && o.ActivatedAt.IsSet() {
		return true
	}

	return false
}

// SetActivatedAt gets a reference to the given NullableString and assigns it to the ActivatedAt field.
func (o *Mandate) SetActivatedAt(v string) {
	o.ActivatedAt.Set(&v)
}
// SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil
func (o *Mandate) SetActivatedAtNil() {
	o.ActivatedAt.Set(nil)
}

// UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
func (o *Mandate) UnsetActivatedAt() {
	o.ActivatedAt.Unset()
}

// GetSuspendedAt returns the SuspendedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetSuspendedAt() string {
	if o == nil || IsNil(o.SuspendedAt.Get()) {
		var ret string
		return ret
	}
	return *o.SuspendedAt.Get()
}

// GetSuspendedAtOk returns a tuple with the SuspendedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetSuspendedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuspendedAt.Get(), o.SuspendedAt.IsSet()
}

// HasSuspendedAt returns a boolean if a field has been set.
func (o *Mandate) HasSuspendedAt() bool {
	if o != nil && o.SuspendedAt.IsSet() {
		return true
	}

	return false
}

// SetSuspendedAt gets a reference to the given NullableString and assigns it to the SuspendedAt field.
func (o *Mandate) SetSuspendedAt(v string) {
	o.SuspendedAt.Set(&v)
}
// SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil
func (o *Mandate) SetSuspendedAtNil() {
	o.SuspendedAt.Set(nil)
}

// UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
func (o *Mandate) UnsetSuspendedAt() {
	o.SuspendedAt.Unset()
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetCancelledAt() string {
	if o == nil || IsNil(o.CancelledAt.Get()) {
		var ret string
		return ret
	}
	return *o.CancelledAt.Get()
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetCancelledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelledAt.Get(), o.CancelledAt.IsSet()
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *Mandate) HasCancelledAt() bool {
	if o != nil && o.CancelledAt.IsSet() {
		return true
	}

	return false
}

// SetCancelledAt gets a reference to the given NullableString and assigns it to the CancelledAt field.
func (o *Mandate) SetCancelledAt(v string) {
	o.CancelledAt.Set(&v)
}
// SetCancelledAtNil sets the value for CancelledAt to be an explicit nil
func (o *Mandate) SetCancelledAtNil() {
	o.CancelledAt.Set(nil)
}

// UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
func (o *Mandate) UnsetCancelledAt() {
	o.CancelledAt.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason.Get()) {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *Mandate) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *Mandate) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *Mandate) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *Mandate) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetMandateRequestId returns the MandateRequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetMandateRequestId() string {
	if o == nil || IsNil(o.MandateRequestId.Get()) {
		var ret string
		return ret
	}
	return *o.MandateRequestId.Get()
}

// GetMandateRequestIdOk returns a tuple with the MandateRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetMandateRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateRequestId.Get(), o.MandateRequestId.IsSet()
}

// HasMandateRequestId returns a boolean if a field has been set.
func (o *Mandate) HasMandateRequestId() bool {
	if o != nil && o.MandateRequestId.IsSet() {
		return true
	}

	return false
}

// SetMandateRequestId gets a reference to the given NullableString and assigns it to the MandateRequestId field.
func (o *Mandate) SetMandateRequestId(v string) {
	o.MandateRequestId.Set(&v)
}
// SetMandateRequestIdNil sets the value for MandateRequestId to be an explicit nil
func (o *Mandate) SetMandateRequestIdNil() {
	o.MandateRequestId.Set(nil)
}

// UnsetMandateRequestId ensures that no value is present for MandateRequestId, not even an explicit nil
func (o *Mandate) UnsetMandateRequestId() {
	o.MandateRequestId.Unset()
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId.Get()) {
		var ret string
		return ret
	}
	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Mandate) HasInvoiceId() bool {
	if o != nil && o.InvoiceId.IsSet() {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given NullableString and assigns it to the InvoiceId field.
func (o *Mandate) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}
// SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil
func (o *Mandate) SetInvoiceIdNil() {
	o.InvoiceId.Set(nil)
}

// UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
func (o *Mandate) UnsetInvoiceId() {
	o.InvoiceId.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Mandate) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Mandate) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Mandate) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Mandate) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Mandate) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Mandate) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Mandate) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableString and assigns it to the UpdatedAt field.
func (o *Mandate) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Mandate) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Mandate) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Mandate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Mandate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["mandate_ref"] = o.MandateRef
	toSerialize["mandate_reference"] = o.MandateReference.Get()
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["status"] = o.Status
	toSerialize["account_holder_name"] = o.AccountHolderName.Get()
	toSerialize["sort_code"] = o.SortCode
	toSerialize["account_number_last4"] = o.AccountNumberLast4.Get()
	toSerialize["bank_name"] = o.BankName.Get()
	if o.StatementName.IsSet() {
		toSerialize["statement_name"] = o.StatementName.Get()
	}
	toSerialize["notice_days"] = o.NoticeDays.Get()
	if o.FirstCollectionDate.IsSet() {
		toSerialize["first_collection_date"] = o.FirstCollectionDate.Get()
	}
	if o.NextCollectionDate.IsSet() {
		toSerialize["next_collection_date"] = o.NextCollectionDate.Get()
	}
	if o.BackupPaymentMethodId.IsSet() {
		toSerialize["backup_payment_method_id"] = o.BackupPaymentMethodId.Get()
	}
	if o.ActivatedAt.IsSet() {
		toSerialize["activated_at"] = o.ActivatedAt.Get()
	}
	if o.SuspendedAt.IsSet() {
		toSerialize["suspended_at"] = o.SuspendedAt.Get()
	}
	if o.CancelledAt.IsSet() {
		toSerialize["cancelled_at"] = o.CancelledAt.Get()
	}
	if o.FailureReason.IsSet() {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	if o.MandateRequestId.IsSet() {
		toSerialize["mandate_request_id"] = o.MandateRequestId.Get()
	}
	if o.InvoiceId.IsSet() {
		toSerialize["invoice_id"] = o.InvoiceId.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Mandate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"mandate_ref",
		"mandate_reference",
		"customer_id",
		"status",
		"account_holder_name",
		"sort_code",
		"account_number_last4",
		"bank_name",
		"notice_days",
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

	varMandate := _Mandate{}

	err = json.Unmarshal(data, &varMandate)

	if err != nil {
		return err
	}

	*o = Mandate(varMandate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "mandate_ref")
		delete(additionalProperties, "mandate_reference")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "account_holder_name")
		delete(additionalProperties, "sort_code")
		delete(additionalProperties, "account_number_last4")
		delete(additionalProperties, "bank_name")
		delete(additionalProperties, "statement_name")
		delete(additionalProperties, "notice_days")
		delete(additionalProperties, "first_collection_date")
		delete(additionalProperties, "next_collection_date")
		delete(additionalProperties, "backup_payment_method_id")
		delete(additionalProperties, "activated_at")
		delete(additionalProperties, "suspended_at")
		delete(additionalProperties, "cancelled_at")
		delete(additionalProperties, "failure_reason")
		delete(additionalProperties, "mandate_request_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMandate struct {
	value *Mandate
	isSet bool
}

func (v NullableMandate) Get() *Mandate {
	return v.value
}

func (v *NullableMandate) Set(val *Mandate) {
	v.value = val
	v.isSet = true
}

func (v NullableMandate) IsSet() bool {
	return v.isSet
}

func (v *NullableMandate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMandate(val *Mandate) *NullableMandate {
	return &NullableMandate{value: val, isSet: true}
}

func (v NullableMandate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMandate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


