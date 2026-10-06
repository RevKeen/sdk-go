/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your merchant API key in the `x-api-key` header, or as `Authorization: Bearer <key>` — both are accepted. Create and manage keys in the [Dashboard](https://app.revkeen.com/dashboard/settings/general?tab=developer). Use `rk_sandbox_*` for staging/test and `rk_live_*` for production. Keys are scoped, rollable, IP-lockable and revocable; the raw secret is shown once, at creation.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Access tokens are opaque, prefixed `rkoa_`, and sent via `Authorization: Bearer rkoa_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Errors  Every non-2xx response uses one envelope: `{ \"error\": { \"type\", \"code\", \"message\", ... } }`. Branch on `error.type` and `error.code`; treat `error.message` as display text only. Error bodies never echo the credential you presented.  | Status | `error.type` | Notes | |--------|--------------|-------| | `400` | `invalid_request_error` | Schema validation failed. `details.fields` maps each rejected parameter to its reasons. | | `401` | `authentication_error` | Missing, invalid, expired or revoked credential. Always carries `WWW-Authenticate: Bearer realm=\"revkeen-api\"`. | | `402` | `card_error` | Declined by the issuer or processor; see `decline_code`. | | `403` | `authorization_error` | The caller is identified but the action is denied — scope, IP allowlist, origin, or merchant mismatch. | | `409` | `invalid_request_error` or `idempotency_error` | Resource-state conflict, or an `Idempotency-Key` replayed with a different body. | | `429` | `rate_limit_error` | See Rate Limits below. | | `500` | `api_error` | Retry with exponential backoff, quoting `error.request_id`. |  `401` codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`.  `403` codes: `insufficient_permissions`, `forbidden`, `ip_not_allowed`, `tenant_scope_mismatch`, `no_organization`, `no_merchant`, `merchant_mismatch`, `cart_disabled`, `origin_not_allowed`, `verification_required`, `read_only_impersonation_blocked`, `publishable_key_restricted`, `internal_only`.  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  Each API key (`rk_live_*`, `rk_sandbox_*`, `rk_pk_*`) may make **600 requests per minute**, in production and staging alike. The limit is per key, counted over a 60-second window, with no separate burst allowance. Some endpoints apply their own, lower limits.  An over-limit key is still valid: it receives `429`, never `401`. Wait for `Retry-After` and retry; do not re-authenticate or rotate the key.  A rate-limited request returns `429` with `type: rate_limit_error`, `code: rate_limit_exceeded` and `details.retry_after` in whole seconds.  | Header | Meaning | |--------|---------| | `Retry-After` | Whole seconds to wait before retrying. Present on every `429`. | | `X-RateLimit-Limit` | Requests permitted in the current window, when the limiter publishes one. | | `X-RateLimit-Remaining` | Requests left in the current window. | | `X-RateLimit-Reset` | UNIX timestamp in **seconds** at which the window resets. |

API version: 2026-05-01
Contact: support@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the DdMandateRequestActiveItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdMandateRequestActiveItem{}

// DdMandateRequestActiveItem struct for DdMandateRequestActiveItem
type DdMandateRequestActiveItem struct {
	Id             string `json:"id"`
	MerchantId     string `json:"merchant_id"`
	CustomerId     string `json:"customer_id"`
	InvoiceId      string `json:"invoice_id"`
	SubscriptionId string `json:"subscription_id"`
	MandateId      string `json:"mandate_id"`
	// Internal raw request token. Do not expose directly to customers.
	Token string `json:"token"`
	// Customer-facing tracked URL: pay.revkeen.com/dd_<token>
	SignableUrl                   string                          `json:"signable_url"`
	CreatedAt                     time.Time                       `json:"created_at"`
	ExpiresAt                     time.Time                       `json:"expires_at"`
	ConsumedAt                    time.Time                       `json:"consumed_at"`
	CancelledAt                   time.Time                       `json:"cancelled_at"`
	ExpiredAt                     time.Time                       `json:"expired_at"`
	Status                        DdMandateRequestLifecycleStatus `json:"status"`
	NotificationDeliveryEventType string                          `json:"notification_delivery_event_type"`
	NotificationDeliveryStatus    string                          `json:"notification_delivery_status"`
	NotificationDeliveryError     string                          `json:"notification_delivery_error"`
	NotificationDeliveryUpdatedAt time.Time                       `json:"notification_delivery_updated_at"`
	PdfDocuments                  []DdMandatePdfReference         `json:"pdf_documents"`
	// When the customer first opened the request link (its first capture session). Null until opened.
	OpenedAt             NullableTime `json:"opened_at"`
	AdditionalProperties map[string]interface{}
}

type _DdMandateRequestActiveItem DdMandateRequestActiveItem

// NewDdMandateRequestActiveItem instantiates a new DdMandateRequestActiveItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdMandateRequestActiveItem(id string, merchantId string, customerId string, invoiceId string, subscriptionId string, mandateId string, token string, signableUrl string, createdAt time.Time, expiresAt time.Time, consumedAt time.Time, cancelledAt time.Time, expiredAt time.Time, status DdMandateRequestLifecycleStatus, notificationDeliveryEventType string, notificationDeliveryStatus string, notificationDeliveryError string, notificationDeliveryUpdatedAt time.Time, pdfDocuments []DdMandatePdfReference, openedAt NullableTime) *DdMandateRequestActiveItem {
	this := DdMandateRequestActiveItem{}
	this.Id = id
	this.MerchantId = merchantId
	this.CustomerId = customerId
	this.InvoiceId = invoiceId
	this.SubscriptionId = subscriptionId
	this.MandateId = mandateId
	this.Token = token
	this.SignableUrl = signableUrl
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.ConsumedAt = consumedAt
	this.CancelledAt = cancelledAt
	this.ExpiredAt = expiredAt
	this.Status = status
	this.NotificationDeliveryEventType = notificationDeliveryEventType
	this.NotificationDeliveryStatus = notificationDeliveryStatus
	this.NotificationDeliveryError = notificationDeliveryError
	this.NotificationDeliveryUpdatedAt = notificationDeliveryUpdatedAt
	this.PdfDocuments = pdfDocuments
	this.OpenedAt = openedAt
	return &this
}

// NewDdMandateRequestActiveItemWithDefaults instantiates a new DdMandateRequestActiveItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdMandateRequestActiveItemWithDefaults() *DdMandateRequestActiveItem {
	this := DdMandateRequestActiveItem{}
	return &this
}

// GetId returns the Id field value
func (o *DdMandateRequestActiveItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DdMandateRequestActiveItem) SetId(v string) {
	o.Id = v
}

// GetMerchantId returns the MerchantId field value
func (o *DdMandateRequestActiveItem) GetMerchantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetMerchantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantId, true
}

// SetMerchantId sets field value
func (o *DdMandateRequestActiveItem) SetMerchantId(v string) {
	o.MerchantId = v
}

// GetCustomerId returns the CustomerId field value
func (o *DdMandateRequestActiveItem) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DdMandateRequestActiveItem) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *DdMandateRequestActiveItem) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *DdMandateRequestActiveItem) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *DdMandateRequestActiveItem) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *DdMandateRequestActiveItem) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetMandateId returns the MandateId field value
func (o *DdMandateRequestActiveItem) GetMandateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MandateId
}

// GetMandateIdOk returns a tuple with the MandateId field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetMandateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MandateId, true
}

// SetMandateId sets field value
func (o *DdMandateRequestActiveItem) SetMandateId(v string) {
	o.MandateId = v
}

// GetToken returns the Token field value
func (o *DdMandateRequestActiveItem) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DdMandateRequestActiveItem) SetToken(v string) {
	o.Token = v
}

// GetSignableUrl returns the SignableUrl field value
func (o *DdMandateRequestActiveItem) GetSignableUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableUrl
}

// GetSignableUrlOk returns a tuple with the SignableUrl field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetSignableUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignableUrl, true
}

// SetSignableUrl sets field value
func (o *DdMandateRequestActiveItem) SetSignableUrl(v string) {
	o.SignableUrl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DdMandateRequestActiveItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DdMandateRequestActiveItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *DdMandateRequestActiveItem) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *DdMandateRequestActiveItem) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetConsumedAt returns the ConsumedAt field value
func (o *DdMandateRequestActiveItem) GetConsumedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ConsumedAt
}

// GetConsumedAtOk returns a tuple with the ConsumedAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetConsumedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumedAt, true
}

// SetConsumedAt sets field value
func (o *DdMandateRequestActiveItem) SetConsumedAt(v time.Time) {
	o.ConsumedAt = v
}

// GetCancelledAt returns the CancelledAt field value
func (o *DdMandateRequestActiveItem) GetCancelledAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetCancelledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelledAt, true
}

// SetCancelledAt sets field value
func (o *DdMandateRequestActiveItem) SetCancelledAt(v time.Time) {
	o.CancelledAt = v
}

// GetExpiredAt returns the ExpiredAt field value
func (o *DdMandateRequestActiveItem) GetExpiredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiredAt, true
}

// SetExpiredAt sets field value
func (o *DdMandateRequestActiveItem) SetExpiredAt(v time.Time) {
	o.ExpiredAt = v
}

// GetStatus returns the Status field value
func (o *DdMandateRequestActiveItem) GetStatus() DdMandateRequestLifecycleStatus {
	if o == nil {
		var ret DdMandateRequestLifecycleStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetStatusOk() (*DdMandateRequestLifecycleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DdMandateRequestActiveItem) SetStatus(v DdMandateRequestLifecycleStatus) {
	o.Status = v
}

// GetNotificationDeliveryEventType returns the NotificationDeliveryEventType field value
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationDeliveryEventType
}

// GetNotificationDeliveryEventTypeOk returns a tuple with the NotificationDeliveryEventType field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDeliveryEventType, true
}

// SetNotificationDeliveryEventType sets field value
func (o *DdMandateRequestActiveItem) SetNotificationDeliveryEventType(v string) {
	o.NotificationDeliveryEventType = v
}

// GetNotificationDeliveryStatus returns the NotificationDeliveryStatus field value
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationDeliveryStatus
}

// GetNotificationDeliveryStatusOk returns a tuple with the NotificationDeliveryStatus field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDeliveryStatus, true
}

// SetNotificationDeliveryStatus sets field value
func (o *DdMandateRequestActiveItem) SetNotificationDeliveryStatus(v string) {
	o.NotificationDeliveryStatus = v
}

// GetNotificationDeliveryError returns the NotificationDeliveryError field value
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationDeliveryError
}

// GetNotificationDeliveryErrorOk returns a tuple with the NotificationDeliveryError field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDeliveryError, true
}

// SetNotificationDeliveryError sets field value
func (o *DdMandateRequestActiveItem) SetNotificationDeliveryError(v string) {
	o.NotificationDeliveryError = v
}

// GetNotificationDeliveryUpdatedAt returns the NotificationDeliveryUpdatedAt field value
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NotificationDeliveryUpdatedAt
}

// GetNotificationDeliveryUpdatedAtOk returns a tuple with the NotificationDeliveryUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetNotificationDeliveryUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDeliveryUpdatedAt, true
}

// SetNotificationDeliveryUpdatedAt sets field value
func (o *DdMandateRequestActiveItem) SetNotificationDeliveryUpdatedAt(v time.Time) {
	o.NotificationDeliveryUpdatedAt = v
}

// GetPdfDocuments returns the PdfDocuments field value
func (o *DdMandateRequestActiveItem) GetPdfDocuments() []DdMandatePdfReference {
	if o == nil {
		var ret []DdMandatePdfReference
		return ret
	}

	return o.PdfDocuments
}

// GetPdfDocumentsOk returns a tuple with the PdfDocuments field value
// and a boolean to check if the value has been set.
func (o *DdMandateRequestActiveItem) GetPdfDocumentsOk() ([]DdMandatePdfReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.PdfDocuments, true
}

// SetPdfDocuments sets field value
func (o *DdMandateRequestActiveItem) SetPdfDocuments(v []DdMandatePdfReference) {
	o.PdfDocuments = v
}

// GetOpenedAt returns the OpenedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *DdMandateRequestActiveItem) GetOpenedAt() time.Time {
	if o == nil || o.OpenedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.OpenedAt.Get()
}

// GetOpenedAtOk returns a tuple with the OpenedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DdMandateRequestActiveItem) GetOpenedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenedAt.Get(), o.OpenedAt.IsSet()
}

// SetOpenedAt sets field value
func (o *DdMandateRequestActiveItem) SetOpenedAt(v time.Time) {
	o.OpenedAt.Set(&v)
}

func (o DdMandateRequestActiveItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdMandateRequestActiveItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["merchant_id"] = o.MerchantId
	toSerialize["customer_id"] = o.CustomerId
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["mandate_id"] = o.MandateId
	toSerialize["token"] = o.Token
	toSerialize["signable_url"] = o.SignableUrl
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["consumed_at"] = o.ConsumedAt
	toSerialize["cancelled_at"] = o.CancelledAt
	toSerialize["expired_at"] = o.ExpiredAt
	toSerialize["status"] = o.Status
	toSerialize["notification_delivery_event_type"] = o.NotificationDeliveryEventType
	toSerialize["notification_delivery_status"] = o.NotificationDeliveryStatus
	toSerialize["notification_delivery_error"] = o.NotificationDeliveryError
	toSerialize["notification_delivery_updated_at"] = o.NotificationDeliveryUpdatedAt
	toSerialize["pdf_documents"] = o.PdfDocuments
	toSerialize["opened_at"] = o.OpenedAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DdMandateRequestActiveItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"merchant_id",
		"customer_id",
		"invoice_id",
		"subscription_id",
		"mandate_id",
		"token",
		"signable_url",
		"created_at",
		"expires_at",
		"consumed_at",
		"cancelled_at",
		"expired_at",
		"status",
		"notification_delivery_event_type",
		"notification_delivery_status",
		"notification_delivery_error",
		"notification_delivery_updated_at",
		"pdf_documents",
		"opened_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDdMandateRequestActiveItem := _DdMandateRequestActiveItem{}

	err = json.Unmarshal(data, &varDdMandateRequestActiveItem)

	if err != nil {
		return err
	}

	*o = DdMandateRequestActiveItem(varDdMandateRequestActiveItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "merchant_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "mandate_id")
		delete(additionalProperties, "token")
		delete(additionalProperties, "signable_url")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "consumed_at")
		delete(additionalProperties, "cancelled_at")
		delete(additionalProperties, "expired_at")
		delete(additionalProperties, "status")
		delete(additionalProperties, "notification_delivery_event_type")
		delete(additionalProperties, "notification_delivery_status")
		delete(additionalProperties, "notification_delivery_error")
		delete(additionalProperties, "notification_delivery_updated_at")
		delete(additionalProperties, "pdf_documents")
		delete(additionalProperties, "opened_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDdMandateRequestActiveItem struct {
	value *DdMandateRequestActiveItem
	isSet bool
}

func (v NullableDdMandateRequestActiveItem) Get() *DdMandateRequestActiveItem {
	return v.value
}

func (v *NullableDdMandateRequestActiveItem) Set(val *DdMandateRequestActiveItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDdMandateRequestActiveItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDdMandateRequestActiveItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdMandateRequestActiveItem(val *DdMandateRequestActiveItem) *NullableDdMandateRequestActiveItem {
	return &NullableDdMandateRequestActiveItem{value: val, isSet: true}
}

func (v NullableDdMandateRequestActiveItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdMandateRequestActiveItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
