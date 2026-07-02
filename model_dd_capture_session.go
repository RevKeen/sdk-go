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

// checks if the DdCaptureSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdCaptureSession{}

// DdCaptureSession struct for DdCaptureSession
type DdCaptureSession struct {
	Id string `json:"id"`
	// Plaintext capture-session token. Only present in the response that issues it (create / token re-issue); never returned on subsequent reads.
	SessionToken *string `json:"session_token,omitempty"`
	MerchantId string `json:"merchant_id"`
	CustomerId NullableString `json:"customer_id"`
	CheckoutSessionId NullableString `json:"checkout_session_id"`
	MandateRequestId NullableString `json:"mandate_request_id"`
	Source string `json:"source"`
	Status string `json:"status"`
	AccountHolderName NullableString `json:"account_holder_name"`
	// Last 2 digits of the sort code — never the full value
	SortCodeLast2 NullableString `json:"sort_code_last2"`
	// Last 4 digits of the account number
	AccountNumberLast4 NullableString `json:"account_number_last4"`
	BankName NullableString `json:"bank_name"`
	// Advance-notice date (YYYY-MM-DD)
	NoticeDate NullableString `json:"notice_date"`
	// Bacs submission date (YYYY-MM-DD)
	SubmissionDate NullableString `json:"submission_date"`
	// Collection date (YYYY-MM-DD)
	CollectionDate NullableString `json:"collection_date"`
	// Settlement date (YYYY-MM-DD)
	SettlementDate NullableString `json:"settlement_date"`
	NoticeDays NullableInt32 `json:"notice_days"`
	// Mandate created from this capture session, once completed
	MandateId NullableString `json:"mandate_id"`
	ExpiresAt NullableString `json:"expires_at"`
	CompletedAt NullableString `json:"completed_at"`
	AbandonedAt NullableString `json:"abandoned_at"`
	CreatedAt NullableString `json:"created_at"`
	UpdatedAt NullableString `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _DdCaptureSession DdCaptureSession

// NewDdCaptureSession instantiates a new DdCaptureSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdCaptureSession(id string, merchantId string, customerId NullableString, checkoutSessionId NullableString, mandateRequestId NullableString, source string, status string, accountHolderName NullableString, sortCodeLast2 NullableString, accountNumberLast4 NullableString, bankName NullableString, noticeDate NullableString, submissionDate NullableString, collectionDate NullableString, settlementDate NullableString, noticeDays NullableInt32, mandateId NullableString, expiresAt NullableString, completedAt NullableString, abandonedAt NullableString, createdAt NullableString, updatedAt NullableString) *DdCaptureSession {
	this := DdCaptureSession{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.CheckoutSessionId = checkoutSessionId
	this.MandateRequestId = mandateRequestId
	this.Source = source
	this.Status = status
	this.AccountHolderName = accountHolderName
	this.SortCodeLast2 = sortCodeLast2
	this.AccountNumberLast4 = accountNumberLast4
	this.BankName = bankName
	this.NoticeDate = noticeDate
	this.SubmissionDate = submissionDate
	this.CollectionDate = collectionDate
	this.SettlementDate = settlementDate
	this.NoticeDays = noticeDays
	this.MandateId = mandateId
	this.ExpiresAt = expiresAt
	this.CompletedAt = completedAt
	this.AbandonedAt = abandonedAt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDdCaptureSessionWithDefaults instantiates a new DdCaptureSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdCaptureSessionWithDefaults() *DdCaptureSession {
	this := DdCaptureSession{}
	return &this
}

// GetId returns the Id field value
func (o *DdCaptureSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DdCaptureSession) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DdCaptureSession) SetId(v string) {
	o.Id = v
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *DdCaptureSession) GetSessionToken() string {
	if o == nil || IsNil(o.SessionToken) {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdCaptureSession) GetSessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.SessionToken) {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *DdCaptureSession) HasSessionToken() bool {
	if o != nil && !IsNil(o.SessionToken) {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *DdCaptureSession) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetMerchantId returns the MerchantId field value
func (o *DdCaptureSession) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *DdCaptureSession) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *DdCaptureSession) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *DdCaptureSession) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetCheckoutSessionId returns the CheckoutSessionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetCheckoutSessionId() string {
	if o == nil || o.CheckoutSessionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CheckoutSessionId.Get()
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutSessionId.Get(), o.CheckoutSessionId.IsSet()
}

// SetCheckoutSessionId sets field value
func (o *DdCaptureSession) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId.Set(&v)
}

// GetMandateRequestId returns the MandateRequestId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetMandateRequestId() string {
	if o == nil || o.MandateRequestId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateRequestId.Get()
}

// GetMandateRequestIdOk returns a tuple with the MandateRequestId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetMandateRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateRequestId.Get(), o.MandateRequestId.IsSet()
}

// SetMandateRequestId sets field value
func (o *DdCaptureSession) SetMandateRequestId(v string) {
	o.MandateRequestId.Set(&v)
}

// GetSource returns the Source field value
func (o *DdCaptureSession) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DdCaptureSession) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DdCaptureSession) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value
func (o *DdCaptureSession) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DdCaptureSession) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DdCaptureSession) SetStatus(v string) {
	o.Status = v
}

// GetAccountHolderName returns the AccountHolderName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetAccountHolderName() string {
	if o == nil || o.AccountHolderName.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountHolderName.Get()
}

// GetAccountHolderNameOk returns a tuple with the AccountHolderName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetAccountHolderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHolderName.Get(), o.AccountHolderName.IsSet()
}

// SetAccountHolderName sets field value
func (o *DdCaptureSession) SetAccountHolderName(v string) {
	o.AccountHolderName.Set(&v)
}

// GetSortCodeLast2 returns the SortCodeLast2 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetSortCodeLast2() string {
	if o == nil || o.SortCodeLast2.Get() == nil {
		var ret string
		return ret
	}

	return *o.SortCodeLast2.Get()
}

// GetSortCodeLast2Ok returns a tuple with the SortCodeLast2 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetSortCodeLast2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortCodeLast2.Get(), o.SortCodeLast2.IsSet()
}

// SetSortCodeLast2 sets field value
func (o *DdCaptureSession) SetSortCodeLast2(v string) {
	o.SortCodeLast2.Set(&v)
}

// GetAccountNumberLast4 returns the AccountNumberLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetAccountNumberLast4() string {
	if o == nil || o.AccountNumberLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.AccountNumberLast4.Get()
}

// GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetAccountNumberLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumberLast4.Get(), o.AccountNumberLast4.IsSet()
}

// SetAccountNumberLast4 sets field value
func (o *DdCaptureSession) SetAccountNumberLast4(v string) {
	o.AccountNumberLast4.Set(&v)
}

// GetBankName returns the BankName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetBankName() string {
	if o == nil || o.BankName.Get() == nil {
		var ret string
		return ret
	}

	return *o.BankName.Get()
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetBankNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BankName.Get(), o.BankName.IsSet()
}

// SetBankName sets field value
func (o *DdCaptureSession) SetBankName(v string) {
	o.BankName.Set(&v)
}

// GetNoticeDate returns the NoticeDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetNoticeDate() string {
	if o == nil || o.NoticeDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.NoticeDate.Get()
}

// GetNoticeDateOk returns a tuple with the NoticeDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetNoticeDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoticeDate.Get(), o.NoticeDate.IsSet()
}

// SetNoticeDate sets field value
func (o *DdCaptureSession) SetNoticeDate(v string) {
	o.NoticeDate.Set(&v)
}

// GetSubmissionDate returns the SubmissionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetSubmissionDate() string {
	if o == nil || o.SubmissionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubmissionDate.Get()
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetSubmissionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionDate.Get(), o.SubmissionDate.IsSet()
}

// SetSubmissionDate sets field value
func (o *DdCaptureSession) SetSubmissionDate(v string) {
	o.SubmissionDate.Set(&v)
}

// GetCollectionDate returns the CollectionDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetCollectionDate() string {
	if o == nil || o.CollectionDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.CollectionDate.Get()
}

// GetCollectionDateOk returns a tuple with the CollectionDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetCollectionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionDate.Get(), o.CollectionDate.IsSet()
}

// SetCollectionDate sets field value
func (o *DdCaptureSession) SetCollectionDate(v string) {
	o.CollectionDate.Set(&v)
}

// GetSettlementDate returns the SettlementDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetSettlementDate() string {
	if o == nil || o.SettlementDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.SettlementDate.Get()
}

// GetSettlementDateOk returns a tuple with the SettlementDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetSettlementDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SettlementDate.Get(), o.SettlementDate.IsSet()
}

// SetSettlementDate sets field value
func (o *DdCaptureSession) SetSettlementDate(v string) {
	o.SettlementDate.Set(&v)
}

// GetNoticeDays returns the NoticeDays field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *DdCaptureSession) GetNoticeDays() int32 {
	if o == nil || o.NoticeDays.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NoticeDays.Get()
}

// GetNoticeDaysOk returns a tuple with the NoticeDays field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetNoticeDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoticeDays.Get(), o.NoticeDays.IsSet()
}

// SetNoticeDays sets field value
func (o *DdCaptureSession) SetNoticeDays(v int32) {
	o.NoticeDays.Set(&v)
}

// GetMandateId returns the MandateId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetMandateId() string {
	if o == nil || o.MandateId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MandateId.Get()
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MandateId.Get(), o.MandateId.IsSet()
}

// SetMandateId sets field value
func (o *DdCaptureSession) SetMandateId(v string) {
	o.MandateId.Set(&v)
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetExpiresAt() string {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *DdCaptureSession) SetExpiresAt(v string) {
	o.ExpiresAt.Set(&v)
}

// GetCompletedAt returns the CompletedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetCompletedAt() string {
	if o == nil || o.CompletedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetCompletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// SetCompletedAt sets field value
func (o *DdCaptureSession) SetCompletedAt(v string) {
	o.CompletedAt.Set(&v)
}

// GetAbandonedAt returns the AbandonedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetAbandonedAt() string {
	if o == nil || o.AbandonedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.AbandonedAt.Get()
}

// GetAbandonedAtOk returns a tuple with the AbandonedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetAbandonedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AbandonedAt.Get(), o.AbandonedAt.IsSet()
}

// SetAbandonedAt sets field value
func (o *DdCaptureSession) SetAbandonedAt(v string) {
	o.AbandonedAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *DdCaptureSession) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DdCaptureSession) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdCaptureSession) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *DdCaptureSession) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

func (o DdCaptureSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdCaptureSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.SessionToken) {
		toSerialize["session_token"] = o.SessionToken
	}
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["checkout_session_id"] = o.CheckoutSessionId.Get()
	toSerialize["mandate_request_id"] = o.MandateRequestId.Get()
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	toSerialize["account_holder_name"] = o.AccountHolderName.Get()
	toSerialize["sort_code_last2"] = o.SortCodeLast2.Get()
	toSerialize["account_number_last4"] = o.AccountNumberLast4.Get()
	toSerialize["bank_name"] = o.BankName.Get()
	toSerialize["notice_date"] = o.NoticeDate.Get()
	toSerialize["submission_date"] = o.SubmissionDate.Get()
	toSerialize["collection_date"] = o.CollectionDate.Get()
	toSerialize["settlement_date"] = o.SettlementDate.Get()
	toSerialize["notice_days"] = o.NoticeDays.Get()
	toSerialize["mandate_id"] = o.MandateId.Get()
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["completed_at"] = o.CompletedAt.Get()
	toSerialize["abandoned_at"] = o.AbandonedAt.Get()
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["updated_at"] = o.UpdatedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdCaptureSession) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"customer_id",
		"checkout_session_id",
		"mandate_request_id",
		"source",
		"status",
		"account_holder_name",
		"sort_code_last2",
		"account_number_last4",
		"bank_name",
		"notice_date",
		"submission_date",
		"collection_date",
		"settlement_date",
		"notice_days",
		"mandate_id",
		"expires_at",
		"completed_at",
		"abandoned_at",
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

	varDdCaptureSession := _DdCaptureSession{}

	err = json.Unmarshal(data, &varDdCaptureSession)

	if err != nil {
		return err
	}

	*o = DdCaptureSession(varDdCaptureSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "session_token")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "checkout_session_id")
		delete(additionalProperties, "mandate_request_id")
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "account_holder_name")
		delete(additionalProperties, "sort_code_last2")
		delete(additionalProperties, "account_number_last4")
		delete(additionalProperties, "bank_name")
		delete(additionalProperties, "notice_date")
		delete(additionalProperties, "submission_date")
		delete(additionalProperties, "collection_date")
		delete(additionalProperties, "settlement_date")
		delete(additionalProperties, "notice_days")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "abandoned_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdCaptureSession struct {
	value *DdCaptureSession
	isSet bool
}

func (v NullableDdCaptureSession) Get() *DdCaptureSession {
	return v.value
}

func (v *NullableDdCaptureSession) Set(val *DdCaptureSession) {
	v.value = val
	v.isSet = true
}

func (v NullableDdCaptureSession) IsSet() bool {
	return v.isSet
}

func (v *NullableDdCaptureSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdCaptureSession(val *DdCaptureSession) *NullableDdCaptureSession {
	return &NullableDdCaptureSession{value: val, isSet: true}
}

func (v NullableDdCaptureSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdCaptureSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


