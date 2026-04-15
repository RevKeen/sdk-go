/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-01-15` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-01-15
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the RetryEligibilityResponseDataLastAttempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryEligibilityResponseDataLastAttempt{}

// RetryEligibilityResponseDataLastAttempt struct for RetryEligibilityResponseDataLastAttempt
type RetryEligibilityResponseDataLastAttempt struct {
	// Unique identifier for the payment attempt
	Id string `json:"id"`
	// External reference ID (payatt_xxx format). Used as order_id in gateway requests.
	PublicId string `json:"public_id"`
	// The exact order_id sent to the gateway (NMI). Format: payatt_<public_id>. Primary correlation key for webhook mapping.
	GatewayOrderId string `json:"gateway_order_id"`
	// Transaction ID returned by the gateway (NMI's transaction reference).
	GatewayTransactionId string `json:"gateway_transaction_id"`
	// Legacy transaction reference (deprecated - use gateway_transaction_id)
	TransactionId string `json:"transaction_id"`
	// Logical payment (PaymentIntent) this attempt belongs to
	PaymentId string `json:"payment_id"`
	// Associated invoice ID
	InvoiceId string `json:"invoice_id"`
	// Associated subscription ID
	SubscriptionId string `json:"subscription_id"`
	// Associated checkout session ID
	CheckoutSessionId string `json:"checkout_session_id"`
	// Billing run ID for batch reconciliation
	BillingRunId string `json:"billing_run_id"`
	// Payment gateway (e.g., 'nmi', 'stripe')
	Gateway string `json:"gateway"`
	// Raw gateway response code
	GatewayCode string `json:"gateway_code"`
	// Raw gateway response message
	GatewayMessage string `json:"gateway_message"`
	// Normalized error category for the payment attempt
	Category string `json:"category"`
	// Severity level for logging/monitoring
	Severity string `json:"severity"`
	// Retry behavior recommendation. retry_now: Safe to retry immediately. retry_later: Retry after delay. do_not_retry: Hard decline, don't retry.
	Retryability string `json:"retryability"`
	// Recommended action for the customer
	CustomerAction string `json:"customer_action"`
	// Recommended action for the merchant
	MerchantAction string `json:"merchant_action"`
	// Directive for subscription handling after decline
	SubscriptionDirective string `json:"subscription_directive"`
	// Template key for customer-facing message
	SafeCustomerMessageKey string `json:"safe_customer_message_key"`
	// Decline classification. soft: May succeed on retry. hard: Will not succeed on retry.
	DeclineType NullableString `json:"decline_type"`
	// Final outcome status of the attempt
	AttemptStatus string `json:"attempt_status"`
	// Address Verification System response code
	AvsCode string `json:"avs_code"`
	// CVV verification response code
	CvvCode string `json:"cvv_code"`
	// Issuer-specific response code
	IssuerCode string `json:"issuer_code"`
	// Payment amount in cents
	AmountCents int32 `json:"amount_cents"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Last 4 digits of card
	PaymentMethodLast4 string `json:"payment_method_last4"`
	// Card brand (visa, mastercard, etc.)
	PaymentMethodBrand string `json:"payment_method_brand"`
	// Attempt sequence number (1 = first attempt)
	AttemptNumber int32 `json:"attempt_number"`
	// Legacy retry attempt number
	RetryAttemptNumber int32 `json:"retry_attempt_number"`
	// Maximum retry attempts allowed
	MaxRetryAttempts int32 `json:"max_retry_attempts"`
	// Next scheduled retry timestamp
	NextRetryAt time.Time `json:"next_retry_at"`
	// Whether this was a scheduled retry
	IsScheduledRetry bool `json:"is_scheduled_retry"`
	// Whether this was a manual retry
	IsManualRetry bool `json:"is_manual_retry"`
	// When the attempt was created
	CreatedAt time.Time `json:"created_at"`
	// When the attempt was last updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _RetryEligibilityResponseDataLastAttempt RetryEligibilityResponseDataLastAttempt

// NewRetryEligibilityResponseDataLastAttempt instantiates a new RetryEligibilityResponseDataLastAttempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryEligibilityResponseDataLastAttempt(id string, publicId string, gatewayOrderId string, gatewayTransactionId string, transactionId string, paymentId string, invoiceId string, subscriptionId string, checkoutSessionId string, billingRunId string, gateway string, gatewayCode string, gatewayMessage string, category string, severity string, retryability string, customerAction string, merchantAction string, subscriptionDirective string, safeCustomerMessageKey string, declineType NullableString, attemptStatus string, avsCode string, cvvCode string, issuerCode string, amountCents int32, currency string, paymentMethodLast4 string, paymentMethodBrand string, attemptNumber int32, retryAttemptNumber int32, maxRetryAttempts int32, nextRetryAt time.Time, isScheduledRetry bool, isManualRetry bool, createdAt time.Time, updatedAt time.Time) *RetryEligibilityResponseDataLastAttempt {
	this := RetryEligibilityResponseDataLastAttempt{}
	this.Id = id
	this.PublicId = publicId
	this.GatewayOrderId = gatewayOrderId
	this.GatewayTransactionId = gatewayTransactionId
	this.TransactionId = transactionId
	this.PaymentId = paymentId
	this.InvoiceId = invoiceId
	this.SubscriptionId = subscriptionId
	this.CheckoutSessionId = checkoutSessionId
	this.BillingRunId = billingRunId
	this.Gateway = gateway
	this.GatewayCode = gatewayCode
	this.GatewayMessage = gatewayMessage
	this.Category = category
	this.Severity = severity
	this.Retryability = retryability
	this.CustomerAction = customerAction
	this.MerchantAction = merchantAction
	this.SubscriptionDirective = subscriptionDirective
	this.SafeCustomerMessageKey = safeCustomerMessageKey
	this.DeclineType = declineType
	this.AttemptStatus = attemptStatus
	this.AvsCode = avsCode
	this.CvvCode = cvvCode
	this.IssuerCode = issuerCode
	this.AmountCents = amountCents
	this.Currency = currency
	this.PaymentMethodLast4 = paymentMethodLast4
	this.PaymentMethodBrand = paymentMethodBrand
	this.AttemptNumber = attemptNumber
	this.RetryAttemptNumber = retryAttemptNumber
	this.MaxRetryAttempts = maxRetryAttempts
	this.NextRetryAt = nextRetryAt
	this.IsScheduledRetry = isScheduledRetry
	this.IsManualRetry = isManualRetry
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewRetryEligibilityResponseDataLastAttemptWithDefaults instantiates a new RetryEligibilityResponseDataLastAttempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryEligibilityResponseDataLastAttemptWithDefaults() *RetryEligibilityResponseDataLastAttempt {
	this := RetryEligibilityResponseDataLastAttempt{}
	return &this
}

// GetId returns the Id field value
func (o *RetryEligibilityResponseDataLastAttempt) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetPublicId(v string) {
	o.PublicId = v
}

// GetGatewayOrderId returns the GatewayOrderId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayOrderId
}

// GetGatewayOrderIdOk returns a tuple with the GatewayOrderId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayOrderId, true
}

// SetGatewayOrderId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayOrderId(v string) {
	o.GatewayOrderId = v
}

// GetGatewayTransactionId returns the GatewayTransactionId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayTransactionId
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayTransactionId, true
}

// SetGatewayTransactionId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId = v
}

// GetTransactionId returns the TransactionId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetPaymentId returns the PaymentId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentId(v string) {
	o.PaymentId = v
}

// GetInvoiceId returns the InvoiceId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetCheckoutSessionId returns the CheckoutSessionId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCheckoutSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckoutSessionId
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckoutSessionId, true
}

// SetCheckoutSessionId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId = v
}

// GetBillingRunId returns the BillingRunId field value
func (o *RetryEligibilityResponseDataLastAttempt) GetBillingRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BillingRunId
}

// GetBillingRunIdOk returns a tuple with the BillingRunId field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetBillingRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingRunId, true
}

// SetBillingRunId sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetBillingRunId(v string) {
	o.BillingRunId = v
}

// GetGateway returns the Gateway field value
func (o *RetryEligibilityResponseDataLastAttempt) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayCode returns the GatewayCode field value
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayCode
}

// GetGatewayCodeOk returns a tuple with the GatewayCode field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayCode, true
}

// SetGatewayCode sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayCode(v string) {
	o.GatewayCode = v
}

// GetGatewayMessage returns the GatewayMessage field value
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayMessage
}

// GetGatewayMessageOk returns a tuple with the GatewayMessage field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayMessage, true
}

// SetGatewayMessage sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayMessage(v string) {
	o.GatewayMessage = v
}

// GetCategory returns the Category field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCategory(v string) {
	o.Category = v
}

// GetSeverity returns the Severity field value
func (o *RetryEligibilityResponseDataLastAttempt) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetSeverity(v string) {
	o.Severity = v
}

// GetRetryability returns the Retryability field value
func (o *RetryEligibilityResponseDataLastAttempt) GetRetryability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Retryability
}

// GetRetryabilityOk returns a tuple with the Retryability field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetRetryabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryability, true
}

// SetRetryability sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetRetryability(v string) {
	o.Retryability = v
}

// GetCustomerAction returns the CustomerAction field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCustomerAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerAction
}

// GetCustomerActionOk returns a tuple with the CustomerAction field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCustomerActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAction, true
}

// SetCustomerAction sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCustomerAction(v string) {
	o.CustomerAction = v
}

// GetMerchantAction returns the MerchantAction field value
func (o *RetryEligibilityResponseDataLastAttempt) GetMerchantAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAction
}

// GetMerchantActionOk returns a tuple with the MerchantAction field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetMerchantActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAction, true
}

// SetMerchantAction sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetMerchantAction(v string) {
	o.MerchantAction = v
}

// GetSubscriptionDirective returns the SubscriptionDirective field value
func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionDirective() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionDirective
}

// GetSubscriptionDirectiveOk returns a tuple with the SubscriptionDirective field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionDirectiveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionDirective, true
}

// SetSubscriptionDirective sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetSubscriptionDirective(v string) {
	o.SubscriptionDirective = v
}

// GetSafeCustomerMessageKey returns the SafeCustomerMessageKey field value
func (o *RetryEligibilityResponseDataLastAttempt) GetSafeCustomerMessageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SafeCustomerMessageKey
}

// GetSafeCustomerMessageKeyOk returns a tuple with the SafeCustomerMessageKey field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetSafeCustomerMessageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SafeCustomerMessageKey, true
}

// SetSafeCustomerMessageKey sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetSafeCustomerMessageKey(v string) {
	o.SafeCustomerMessageKey = v
}

// GetDeclineType returns the DeclineType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RetryEligibilityResponseDataLastAttempt) GetDeclineType() string {
	if o == nil || o.DeclineType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeclineType.Get()
}

// GetDeclineTypeOk returns a tuple with the DeclineType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RetryEligibilityResponseDataLastAttempt) GetDeclineTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclineType.Get(), o.DeclineType.IsSet()
}

// SetDeclineType sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetDeclineType(v string) {
	o.DeclineType.Set(&v)
}

// GetAttemptStatus returns the AttemptStatus field value
func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttemptStatus
}

// GetAttemptStatusOk returns a tuple with the AttemptStatus field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptStatus, true
}

// SetAttemptStatus sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetAttemptStatus(v string) {
	o.AttemptStatus = v
}

// GetAvsCode returns the AvsCode field value
func (o *RetryEligibilityResponseDataLastAttempt) GetAvsCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvsCode
}

// GetAvsCodeOk returns a tuple with the AvsCode field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetAvsCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvsCode, true
}

// SetAvsCode sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetAvsCode(v string) {
	o.AvsCode = v
}

// GetCvvCode returns the CvvCode field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCvvCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CvvCode
}

// GetCvvCodeOk returns a tuple with the CvvCode field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCvvCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CvvCode, true
}

// SetCvvCode sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCvvCode(v string) {
	o.CvvCode = v
}

// GetIssuerCode returns the IssuerCode field value
func (o *RetryEligibilityResponseDataLastAttempt) GetIssuerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerCode
}

// GetIssuerCodeOk returns a tuple with the IssuerCode field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetIssuerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerCode, true
}

// SetIssuerCode sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetIssuerCode(v string) {
	o.IssuerCode = v
}

// GetAmountCents returns the AmountCents field value
func (o *RetryEligibilityResponseDataLastAttempt) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetCurrency returns the Currency field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCurrency(v string) {
	o.Currency = v
}

// GetPaymentMethodLast4 returns the PaymentMethodLast4 field value
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodLast4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodLast4
}

// GetPaymentMethodLast4Ok returns a tuple with the PaymentMethodLast4 field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodLast4, true
}

// SetPaymentMethodLast4 sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentMethodLast4(v string) {
	o.PaymentMethodLast4 = v
}

// GetPaymentMethodBrand returns the PaymentMethodBrand field value
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodBrand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMethodBrand
}

// GetPaymentMethodBrandOk returns a tuple with the PaymentMethodBrand field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMethodBrand, true
}

// SetPaymentMethodBrand sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentMethodBrand(v string) {
	o.PaymentMethodBrand = v
}

// GetAttemptNumber returns the AttemptNumber field value
func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttemptNumber
}

// GetAttemptNumberOk returns a tuple with the AttemptNumber field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptNumber, true
}

// SetAttemptNumber sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetAttemptNumber(v int32) {
	o.AttemptNumber = v
}

// GetRetryAttemptNumber returns the RetryAttemptNumber field value
func (o *RetryEligibilityResponseDataLastAttempt) GetRetryAttemptNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetryAttemptNumber
}

// GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetRetryAttemptNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryAttemptNumber, true
}

// SetRetryAttemptNumber sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetRetryAttemptNumber(v int32) {
	o.RetryAttemptNumber = v
}

// GetMaxRetryAttempts returns the MaxRetryAttempts field value
func (o *RetryEligibilityResponseDataLastAttempt) GetMaxRetryAttempts() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxRetryAttempts
}

// GetMaxRetryAttemptsOk returns a tuple with the MaxRetryAttempts field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetMaxRetryAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxRetryAttempts, true
}

// SetMaxRetryAttempts sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetMaxRetryAttempts(v int32) {
	o.MaxRetryAttempts = v
}

// GetNextRetryAt returns the NextRetryAt field value
func (o *RetryEligibilityResponseDataLastAttempt) GetNextRetryAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.NextRetryAt
}

// GetNextRetryAtOk returns a tuple with the NextRetryAt field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetNextRetryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextRetryAt, true
}

// SetNextRetryAt sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetNextRetryAt(v time.Time) {
	o.NextRetryAt = v
}

// GetIsScheduledRetry returns the IsScheduledRetry field value
func (o *RetryEligibilityResponseDataLastAttempt) GetIsScheduledRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScheduledRetry
}

// GetIsScheduledRetryOk returns a tuple with the IsScheduledRetry field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetIsScheduledRetryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScheduledRetry, true
}

// SetIsScheduledRetry sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetIsScheduledRetry(v bool) {
	o.IsScheduledRetry = v
}

// GetIsManualRetry returns the IsManualRetry field value
func (o *RetryEligibilityResponseDataLastAttempt) GetIsManualRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsManualRetry
}

// GetIsManualRetryOk returns a tuple with the IsManualRetry field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetIsManualRetryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsManualRetry, true
}

// SetIsManualRetry sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetIsManualRetry(v bool) {
	o.IsManualRetry = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RetryEligibilityResponseDataLastAttempt) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RetryEligibilityResponseDataLastAttempt) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RetryEligibilityResponseDataLastAttempt) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RetryEligibilityResponseDataLastAttempt) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o RetryEligibilityResponseDataLastAttempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryEligibilityResponseDataLastAttempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId
	toSerialize["gateway_order_id"] = o.GatewayOrderId
	toSerialize["gateway_transaction_id"] = o.GatewayTransactionId
	toSerialize["transaction_id"] = o.TransactionId
	toSerialize["payment_id"] = o.PaymentId
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["subscription_id"] = o.SubscriptionId
	toSerialize["checkout_session_id"] = o.CheckoutSessionId
	toSerialize["billing_run_id"] = o.BillingRunId
	toSerialize["gateway"] = o.Gateway
	toSerialize["gateway_code"] = o.GatewayCode
	toSerialize["gateway_message"] = o.GatewayMessage
	toSerialize["category"] = o.Category
	toSerialize["severity"] = o.Severity
	toSerialize["retryability"] = o.Retryability
	toSerialize["customer_action"] = o.CustomerAction
	toSerialize["merchant_action"] = o.MerchantAction
	toSerialize["subscription_directive"] = o.SubscriptionDirective
	toSerialize["safe_customer_message_key"] = o.SafeCustomerMessageKey
	toSerialize["decline_type"] = o.DeclineType.Get()
	toSerialize["attempt_status"] = o.AttemptStatus
	toSerialize["avs_code"] = o.AvsCode
	toSerialize["cvv_code"] = o.CvvCode
	toSerialize["issuer_code"] = o.IssuerCode
	toSerialize["amount_cents"] = o.AmountCents
	toSerialize["currency"] = o.Currency
	toSerialize["payment_method_last4"] = o.PaymentMethodLast4
	toSerialize["payment_method_brand"] = o.PaymentMethodBrand
	toSerialize["attempt_number"] = o.AttemptNumber
	toSerialize["retry_attempt_number"] = o.RetryAttemptNumber
	toSerialize["max_retry_attempts"] = o.MaxRetryAttempts
	toSerialize["next_retry_at"] = o.NextRetryAt
	toSerialize["is_scheduled_retry"] = o.IsScheduledRetry
	toSerialize["is_manual_retry"] = o.IsManualRetry
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RetryEligibilityResponseDataLastAttempt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"public_id",
		"gateway_order_id",
		"gateway_transaction_id",
		"transaction_id",
		"payment_id",
		"invoice_id",
		"subscription_id",
		"checkout_session_id",
		"billing_run_id",
		"gateway",
		"gateway_code",
		"gateway_message",
		"category",
		"severity",
		"retryability",
		"customer_action",
		"merchant_action",
		"subscription_directive",
		"safe_customer_message_key",
		"decline_type",
		"attempt_status",
		"avs_code",
		"cvv_code",
		"issuer_code",
		"amount_cents",
		"currency",
		"payment_method_last4",
		"payment_method_brand",
		"attempt_number",
		"retry_attempt_number",
		"max_retry_attempts",
		"next_retry_at",
		"is_scheduled_retry",
		"is_manual_retry",
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

	varRetryEligibilityResponseDataLastAttempt := _RetryEligibilityResponseDataLastAttempt{}

	err = json.Unmarshal(data, &varRetryEligibilityResponseDataLastAttempt)

	if err != nil {
		return err
	}

	*o = RetryEligibilityResponseDataLastAttempt(varRetryEligibilityResponseDataLastAttempt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "gateway_order_id")
		delete(additionalProperties, "gateway_transaction_id")
		delete(additionalProperties, "transaction_id")
		delete(additionalProperties, "payment_id")
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "subscription_id")
		delete(additionalProperties, "checkout_session_id")
		delete(additionalProperties, "billing_run_id")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway_code")
		delete(additionalProperties, "gateway_message")
		delete(additionalProperties, "category")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "retryability")
		delete(additionalProperties, "customer_action")
		delete(additionalProperties, "merchant_action")
		delete(additionalProperties, "subscription_directive")
		delete(additionalProperties, "safe_customer_message_key")
		delete(additionalProperties, "decline_type")
		delete(additionalProperties, "attempt_status")
		delete(additionalProperties, "avs_code")
		delete(additionalProperties, "cvv_code")
		delete(additionalProperties, "issuer_code")
		delete(additionalProperties, "amount_cents")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "payment_method_last4")
		delete(additionalProperties, "payment_method_brand")
		delete(additionalProperties, "attempt_number")
		delete(additionalProperties, "retry_attempt_number")
		delete(additionalProperties, "max_retry_attempts")
		delete(additionalProperties, "next_retry_at")
		delete(additionalProperties, "is_scheduled_retry")
		delete(additionalProperties, "is_manual_retry")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRetryEligibilityResponseDataLastAttempt struct {
	value *RetryEligibilityResponseDataLastAttempt
	isSet bool
}

func (v NullableRetryEligibilityResponseDataLastAttempt) Get() *RetryEligibilityResponseDataLastAttempt {
	return v.value
}

func (v *NullableRetryEligibilityResponseDataLastAttempt) Set(val *RetryEligibilityResponseDataLastAttempt) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryEligibilityResponseDataLastAttempt) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryEligibilityResponseDataLastAttempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryEligibilityResponseDataLastAttempt(val *RetryEligibilityResponseDataLastAttempt) *NullableRetryEligibilityResponseDataLastAttempt {
	return &NullableRetryEligibilityResponseDataLastAttempt{value: val, isSet: true}
}

func (v NullableRetryEligibilityResponseDataLastAttempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryEligibilityResponseDataLastAttempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


