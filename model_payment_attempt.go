/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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

// checks if the PaymentAttempt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentAttempt{}

// PaymentAttempt A record of an individual payment processing attempt, including gateway response codes, AVS/CVV results, and retry metadata.
type PaymentAttempt struct {
	// Unique identifier for the payment attempt
	Id string `json:"id"`
	// External reference ID (payatt_xxx format). Used as order_id in gateway requests.
	PublicId NullableString `json:"public_id"`
	// Order ID sent to the processor. Format: payatt_<public_id>. Primary correlation key for webhook mapping.
	GatewayOrderId NullableString `json:"gateway_order_id"`
	// Processor-assigned transaction reference.
	GatewayTransactionId NullableString `json:"gateway_transaction_id"`
	// Logical payment (PaymentIntent) this attempt belongs to
	PaymentId NullableString `json:"payment_id"`
	// Associated invoice ID
	InvoiceId NullableString `json:"invoice_id"`
	// Associated subscription ID
	SubscriptionId NullableString `json:"subscription_id"`
	// Associated checkout session ID
	CheckoutSessionId NullableString `json:"checkout_session_id"`
	// Billing run ID for batch reconciliation
	BillingRunId NullableString `json:"billing_run_id"`
	// Name of the payment processor that handled this attempt
	Gateway string `json:"gateway"`
	// Processor-returned response code
	GatewayCode string `json:"gateway_code"`
	// Processor-returned response message
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
	// Payment amount in cents
	AmountCents NullableInt32 `json:"amount_cents"`
	// Three-letter ISO currency code
	Currency NullableString `json:"currency"`
	// Last 4 digits of card
	PaymentMethodLast4 NullableString `json:"payment_method_last4"`
	// Card brand (visa, mastercard, etc.)
	PaymentMethodBrand NullableString `json:"payment_method_brand"`
	// Attempt sequence number (1 = first attempt)
	AttemptNumber int32 `json:"attempt_number"`
	// Legacy retry attempt number
	RetryAttemptNumber int32 `json:"retry_attempt_number"`
	// Maximum retry attempts allowed
	MaxRetryAttempts NullableInt32 `json:"max_retry_attempts"`
	// Next scheduled retry timestamp
	NextRetryAt NullableTime `json:"next_retry_at"`
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

type _PaymentAttempt PaymentAttempt

// NewPaymentAttempt instantiates a new PaymentAttempt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentAttempt(id string, publicId NullableString, gatewayOrderId NullableString, gatewayTransactionId NullableString, paymentId NullableString, invoiceId NullableString, subscriptionId NullableString, checkoutSessionId NullableString, billingRunId NullableString, gateway string, gatewayCode string, gatewayMessage string, category string, severity string, retryability string, customerAction string, merchantAction string, subscriptionDirective string, safeCustomerMessageKey string, declineType NullableString, attemptStatus string, amountCents NullableInt32, currency NullableString, paymentMethodLast4 NullableString, paymentMethodBrand NullableString, attemptNumber int32, retryAttemptNumber int32, maxRetryAttempts NullableInt32, nextRetryAt NullableTime, isScheduledRetry bool, isManualRetry bool, createdAt time.Time, updatedAt time.Time) *PaymentAttempt {
	this := PaymentAttempt{}
	this.Id = id
	this.PublicId = publicId
	this.GatewayOrderId = gatewayOrderId
	this.GatewayTransactionId = gatewayTransactionId
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

// NewPaymentAttemptWithDefaults instantiates a new PaymentAttempt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentAttemptWithDefaults() *PaymentAttempt {
	this := PaymentAttempt{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentAttempt) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentAttempt) SetId(v string) {
	o.Id = v
}

// GetPublicId returns the PublicId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetPublicId() string {
	if o == nil || o.PublicId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PublicId.Get()
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicId.Get(), o.PublicId.IsSet()
}

// SetPublicId sets field value
func (o *PaymentAttempt) SetPublicId(v string) {
	o.PublicId.Set(&v)
}

// GetGatewayOrderId returns the GatewayOrderId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetGatewayOrderId() string {
	if o == nil || o.GatewayOrderId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayOrderId.Get()
}

// GetGatewayOrderIdOk returns a tuple with the GatewayOrderId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetGatewayOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayOrderId.Get(), o.GatewayOrderId.IsSet()
}

// SetGatewayOrderId sets field value
func (o *PaymentAttempt) SetGatewayOrderId(v string) {
	o.GatewayOrderId.Set(&v)
}

// GetGatewayTransactionId returns the GatewayTransactionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetGatewayTransactionId() string {
	if o == nil || o.GatewayTransactionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GatewayTransactionId.Get()
}

// GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetGatewayTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GatewayTransactionId.Get(), o.GatewayTransactionId.IsSet()
}

// SetGatewayTransactionId sets field value
func (o *PaymentAttempt) SetGatewayTransactionId(v string) {
	o.GatewayTransactionId.Set(&v)
}

// GetPaymentId returns the PaymentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetPaymentId() string {
	if o == nil || o.PaymentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentId.Get()
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetPaymentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentId.Get(), o.PaymentId.IsSet()
}

// SetPaymentId sets field value
func (o *PaymentAttempt) SetPaymentId(v string) {
	o.PaymentId.Set(&v)
}

// GetInvoiceId returns the InvoiceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetInvoiceId() string {
	if o == nil || o.InvoiceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InvoiceId.Get()
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvoiceId.Get(), o.InvoiceId.IsSet()
}

// SetInvoiceId sets field value
func (o *PaymentAttempt) SetInvoiceId(v string) {
	o.InvoiceId.Set(&v)
}

// GetSubscriptionId returns the SubscriptionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.SubscriptionId.Get()
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionId.Get(), o.SubscriptionId.IsSet()
}

// SetSubscriptionId sets field value
func (o *PaymentAttempt) SetSubscriptionId(v string) {
	o.SubscriptionId.Set(&v)
}

// GetCheckoutSessionId returns the CheckoutSessionId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetCheckoutSessionId() string {
	if o == nil || o.CheckoutSessionId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CheckoutSessionId.Get()
}

// GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetCheckoutSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckoutSessionId.Get(), o.CheckoutSessionId.IsSet()
}

// SetCheckoutSessionId sets field value
func (o *PaymentAttempt) SetCheckoutSessionId(v string) {
	o.CheckoutSessionId.Set(&v)
}

// GetBillingRunId returns the BillingRunId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetBillingRunId() string {
	if o == nil || o.BillingRunId.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingRunId.Get()
}

// GetBillingRunIdOk returns a tuple with the BillingRunId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetBillingRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingRunId.Get(), o.BillingRunId.IsSet()
}

// SetBillingRunId sets field value
func (o *PaymentAttempt) SetBillingRunId(v string) {
	o.BillingRunId.Set(&v)
}

// GetGateway returns the Gateway field value
func (o *PaymentAttempt) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *PaymentAttempt) SetGateway(v string) {
	o.Gateway = v
}

// GetGatewayCode returns the GatewayCode field value
func (o *PaymentAttempt) GetGatewayCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayCode
}

// GetGatewayCodeOk returns a tuple with the GatewayCode field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetGatewayCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayCode, true
}

// SetGatewayCode sets field value
func (o *PaymentAttempt) SetGatewayCode(v string) {
	o.GatewayCode = v
}

// GetGatewayMessage returns the GatewayMessage field value
func (o *PaymentAttempt) GetGatewayMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayMessage
}

// GetGatewayMessageOk returns a tuple with the GatewayMessage field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetGatewayMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayMessage, true
}

// SetGatewayMessage sets field value
func (o *PaymentAttempt) SetGatewayMessage(v string) {
	o.GatewayMessage = v
}

// GetCategory returns the Category field value
func (o *PaymentAttempt) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PaymentAttempt) SetCategory(v string) {
	o.Category = v
}

// GetSeverity returns the Severity field value
func (o *PaymentAttempt) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *PaymentAttempt) SetSeverity(v string) {
	o.Severity = v
}

// GetRetryability returns the Retryability field value
func (o *PaymentAttempt) GetRetryability() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Retryability
}

// GetRetryabilityOk returns a tuple with the Retryability field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetRetryabilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryability, true
}

// SetRetryability sets field value
func (o *PaymentAttempt) SetRetryability(v string) {
	o.Retryability = v
}

// GetCustomerAction returns the CustomerAction field value
func (o *PaymentAttempt) GetCustomerAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerAction
}

// GetCustomerActionOk returns a tuple with the CustomerAction field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetCustomerActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAction, true
}

// SetCustomerAction sets field value
func (o *PaymentAttempt) SetCustomerAction(v string) {
	o.CustomerAction = v
}

// GetMerchantAction returns the MerchantAction field value
func (o *PaymentAttempt) GetMerchantAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantAction
}

// GetMerchantActionOk returns a tuple with the MerchantAction field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetMerchantActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantAction, true
}

// SetMerchantAction sets field value
func (o *PaymentAttempt) SetMerchantAction(v string) {
	o.MerchantAction = v
}

// GetSubscriptionDirective returns the SubscriptionDirective field value
func (o *PaymentAttempt) GetSubscriptionDirective() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionDirective
}

// GetSubscriptionDirectiveOk returns a tuple with the SubscriptionDirective field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetSubscriptionDirectiveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionDirective, true
}

// SetSubscriptionDirective sets field value
func (o *PaymentAttempt) SetSubscriptionDirective(v string) {
	o.SubscriptionDirective = v
}

// GetSafeCustomerMessageKey returns the SafeCustomerMessageKey field value
func (o *PaymentAttempt) GetSafeCustomerMessageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SafeCustomerMessageKey
}

// GetSafeCustomerMessageKeyOk returns a tuple with the SafeCustomerMessageKey field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetSafeCustomerMessageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SafeCustomerMessageKey, true
}

// SetSafeCustomerMessageKey sets field value
func (o *PaymentAttempt) SetSafeCustomerMessageKey(v string) {
	o.SafeCustomerMessageKey = v
}

// GetDeclineType returns the DeclineType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetDeclineType() string {
	if o == nil || o.DeclineType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeclineType.Get()
}

// GetDeclineTypeOk returns a tuple with the DeclineType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetDeclineTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeclineType.Get(), o.DeclineType.IsSet()
}

// SetDeclineType sets field value
func (o *PaymentAttempt) SetDeclineType(v string) {
	o.DeclineType.Set(&v)
}

// GetAttemptStatus returns the AttemptStatus field value
func (o *PaymentAttempt) GetAttemptStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttemptStatus
}

// GetAttemptStatusOk returns a tuple with the AttemptStatus field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetAttemptStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptStatus, true
}

// SetAttemptStatus sets field value
func (o *PaymentAttempt) SetAttemptStatus(v string) {
	o.AttemptStatus = v
}

// GetAmountCents returns the AmountCents field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaymentAttempt) GetAmountCents() int32 {
	if o == nil || o.AmountCents.Get() == nil {
		var ret int32
		return ret
	}

	return *o.AmountCents.Get()
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountCents.Get(), o.AmountCents.IsSet()
}

// SetAmountCents sets field value
func (o *PaymentAttempt) SetAmountCents(v int32) {
	o.AmountCents.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *PaymentAttempt) SetCurrency(v string) {
	o.Currency.Set(&v)
}

// GetPaymentMethodLast4 returns the PaymentMethodLast4 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetPaymentMethodLast4() string {
	if o == nil || o.PaymentMethodLast4.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethodLast4.Get()
}

// GetPaymentMethodLast4Ok returns a tuple with the PaymentMethodLast4 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetPaymentMethodLast4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodLast4.Get(), o.PaymentMethodLast4.IsSet()
}

// SetPaymentMethodLast4 sets field value
func (o *PaymentAttempt) SetPaymentMethodLast4(v string) {
	o.PaymentMethodLast4.Set(&v)
}

// GetPaymentMethodBrand returns the PaymentMethodBrand field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentAttempt) GetPaymentMethodBrand() string {
	if o == nil || o.PaymentMethodBrand.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethodBrand.Get()
}

// GetPaymentMethodBrandOk returns a tuple with the PaymentMethodBrand field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetPaymentMethodBrandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodBrand.Get(), o.PaymentMethodBrand.IsSet()
}

// SetPaymentMethodBrand sets field value
func (o *PaymentAttempt) SetPaymentMethodBrand(v string) {
	o.PaymentMethodBrand.Set(&v)
}

// GetAttemptNumber returns the AttemptNumber field value
func (o *PaymentAttempt) GetAttemptNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AttemptNumber
}

// GetAttemptNumberOk returns a tuple with the AttemptNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetAttemptNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttemptNumber, true
}

// SetAttemptNumber sets field value
func (o *PaymentAttempt) SetAttemptNumber(v int32) {
	o.AttemptNumber = v
}

// GetRetryAttemptNumber returns the RetryAttemptNumber field value
func (o *PaymentAttempt) GetRetryAttemptNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RetryAttemptNumber
}

// GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetRetryAttemptNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetryAttemptNumber, true
}

// SetRetryAttemptNumber sets field value
func (o *PaymentAttempt) SetRetryAttemptNumber(v int32) {
	o.RetryAttemptNumber = v
}

// GetMaxRetryAttempts returns the MaxRetryAttempts field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaymentAttempt) GetMaxRetryAttempts() int32 {
	if o == nil || o.MaxRetryAttempts.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxRetryAttempts.Get()
}

// GetMaxRetryAttemptsOk returns a tuple with the MaxRetryAttempts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetMaxRetryAttemptsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetryAttempts.Get(), o.MaxRetryAttempts.IsSet()
}

// SetMaxRetryAttempts sets field value
func (o *PaymentAttempt) SetMaxRetryAttempts(v int32) {
	o.MaxRetryAttempts.Set(&v)
}

// GetNextRetryAt returns the NextRetryAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PaymentAttempt) GetNextRetryAt() time.Time {
	if o == nil || o.NextRetryAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.NextRetryAt.Get()
}

// GetNextRetryAtOk returns a tuple with the NextRetryAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentAttempt) GetNextRetryAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRetryAt.Get(), o.NextRetryAt.IsSet()
}

// SetNextRetryAt sets field value
func (o *PaymentAttempt) SetNextRetryAt(v time.Time) {
	o.NextRetryAt.Set(&v)
}

// GetIsScheduledRetry returns the IsScheduledRetry field value
func (o *PaymentAttempt) GetIsScheduledRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScheduledRetry
}

// GetIsScheduledRetryOk returns a tuple with the IsScheduledRetry field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetIsScheduledRetryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScheduledRetry, true
}

// SetIsScheduledRetry sets field value
func (o *PaymentAttempt) SetIsScheduledRetry(v bool) {
	o.IsScheduledRetry = v
}

// GetIsManualRetry returns the IsManualRetry field value
func (o *PaymentAttempt) GetIsManualRetry() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsManualRetry
}

// GetIsManualRetryOk returns a tuple with the IsManualRetry field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetIsManualRetryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsManualRetry, true
}

// SetIsManualRetry sets field value
func (o *PaymentAttempt) SetIsManualRetry(v bool) {
	o.IsManualRetry = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PaymentAttempt) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PaymentAttempt) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PaymentAttempt) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PaymentAttempt) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PaymentAttempt) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PaymentAttempt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentAttempt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["public_id"] = o.PublicId.Get()
	toSerialize["gateway_order_id"] = o.GatewayOrderId.Get()
	toSerialize["gateway_transaction_id"] = o.GatewayTransactionId.Get()
	toSerialize["payment_id"] = o.PaymentId.Get()
	toSerialize["invoice_id"] = o.InvoiceId.Get()
	toSerialize["subscription_id"] = o.SubscriptionId.Get()
	toSerialize["checkout_session_id"] = o.CheckoutSessionId.Get()
	toSerialize["billing_run_id"] = o.BillingRunId.Get()
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
	toSerialize["amount_cents"] = o.AmountCents.Get()
	toSerialize["currency"] = o.Currency.Get()
	toSerialize["payment_method_last4"] = o.PaymentMethodLast4.Get()
	toSerialize["payment_method_brand"] = o.PaymentMethodBrand.Get()
	toSerialize["attempt_number"] = o.AttemptNumber
	toSerialize["retry_attempt_number"] = o.RetryAttemptNumber
	toSerialize["max_retry_attempts"] = o.MaxRetryAttempts.Get()
	toSerialize["next_retry_at"] = o.NextRetryAt.Get()
	toSerialize["is_scheduled_retry"] = o.IsScheduledRetry
	toSerialize["is_manual_retry"] = o.IsManualRetry
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentAttempt) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"public_id",
		"gateway_order_id",
		"gateway_transaction_id",
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

	varPaymentAttempt := _PaymentAttempt{}

	err = json.Unmarshal(data, &varPaymentAttempt)

	if err != nil {
		return err
	}

	*o = PaymentAttempt(varPaymentAttempt)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "gateway_order_id")
		delete(additionalProperties, "gateway_transaction_id")
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

type NullablePaymentAttempt struct {
	value *PaymentAttempt
	isSet bool
}

func (v NullablePaymentAttempt) Get() *PaymentAttempt {
	return v.value
}

func (v *NullablePaymentAttempt) Set(val *PaymentAttempt) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAttempt) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAttempt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAttempt(val *PaymentAttempt) *NullablePaymentAttempt {
	return &NullablePaymentAttempt{value: val, isSet: true}
}

func (v NullablePaymentAttempt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAttempt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


