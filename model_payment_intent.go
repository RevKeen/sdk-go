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

// checks if the PaymentIntent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentIntent{}

// PaymentIntent A PaymentIntent tracks the lifecycle of a payment from creation through authorization, optional 3DS authentication, and final capture.
type PaymentIntent struct {
	// Public payment intent ID (pi_xxx)
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Amount in cents
	Amount int32 `json:"amount"`
	// Amount that can be captured (for manual capture)
	AmountCapturable int32 `json:"amount_capturable"`
	// Amount actually received
	AmountReceived int32 `json:"amount_received"`
	// Three-letter ISO currency code
	Currency string `json:"currency"`
	// Customer ID
	Customer string `json:"customer"`
	// Description for merchant reference
	Description NullableString `json:"description,omitempty"`
	LastPaymentError *PaymentIntentLastPaymentError `json:"last_payment_error,omitempty"`
	// Custom key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	NextAction *PaymentIntentNextAction `json:"next_action,omitempty"`
	// Payment method ID
	PaymentMethod NullableString `json:"payment_method,omitempty"`
	// Email for receipt
	ReceiptEmail NullableString `json:"receipt_email,omitempty"`
	// Statement descriptor
	StatementDescriptor NullableString `json:"statement_descriptor,omitempty"`
	// Statement descriptor suffix
	StatementDescriptorSuffix NullableString `json:"statement_descriptor_suffix,omitempty"`
	// Payment intent status. requires_payment_method: Needs payment method. requires_confirmation: Ready to confirm. requires_action: Requires customer action (3DS). processing: Being processed. succeeded: Payment complete. canceled: Canceled.
	Status string `json:"status"`
	// Capture method. automatic: Capture immediately on confirmation. manual: Authorize then capture separately.
	CaptureMethod string `json:"capture_method"`
	// Client secret for frontend confirmation
	ClientSecret string `json:"client_secret"`
	// When the intent was canceled
	CanceledAt NullableTime `json:"canceled_at,omitempty"`
	// Why the intent was canceled
	CancellationReason NullableString `json:"cancellation_reason,omitempty"`
	// Unix timestamp of creation
	Created int32 `json:"created"`
	// Whether in live mode
	Livemode bool `json:"livemode"`
	AdditionalProperties map[string]interface{}
}

type _PaymentIntent PaymentIntent

// NewPaymentIntent instantiates a new PaymentIntent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentIntent(id string, object string, amount int32, amountCapturable int32, amountReceived int32, currency string, customer string, status string, captureMethod string, clientSecret string, created int32, livemode bool) *PaymentIntent {
	this := PaymentIntent{}
	this.Id = id
	this.Object = object
	this.Amount = amount
	this.AmountCapturable = amountCapturable
	this.AmountReceived = amountReceived
	this.Currency = currency
	this.Customer = customer
	this.Status = status
	this.CaptureMethod = captureMethod
	this.ClientSecret = clientSecret
	this.Created = created
	this.Livemode = livemode
	return &this
}

// NewPaymentIntentWithDefaults instantiates a new PaymentIntent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentIntentWithDefaults() *PaymentIntent {
	this := PaymentIntent{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentIntent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentIntent) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *PaymentIntent) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *PaymentIntent) SetObject(v string) {
	o.Object = v
}

// GetAmount returns the Amount field value
func (o *PaymentIntent) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentIntent) SetAmount(v int32) {
	o.Amount = v
}

// GetAmountCapturable returns the AmountCapturable field value
func (o *PaymentIntent) GetAmountCapturable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCapturable
}

// GetAmountCapturableOk returns a tuple with the AmountCapturable field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetAmountCapturableOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCapturable, true
}

// SetAmountCapturable sets field value
func (o *PaymentIntent) SetAmountCapturable(v int32) {
	o.AmountCapturable = v
}

// GetAmountReceived returns the AmountReceived field value
func (o *PaymentIntent) GetAmountReceived() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountReceived
}

// GetAmountReceivedOk returns a tuple with the AmountReceived field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetAmountReceivedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountReceived, true
}

// SetAmountReceived sets field value
func (o *PaymentIntent) SetAmountReceived(v int32) {
	o.AmountReceived = v
}

// GetCurrency returns the Currency field value
func (o *PaymentIntent) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentIntent) SetCurrency(v string) {
	o.Currency = v
}

// GetCustomer returns the Customer field value
func (o *PaymentIntent) GetCustomer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCustomerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *PaymentIntent) SetCustomer(v string) {
	o.Customer = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentIntent) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PaymentIntent) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PaymentIntent) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PaymentIntent) UnsetDescription() {
	o.Description.Unset()
}

// GetLastPaymentError returns the LastPaymentError field value if set, zero value otherwise.
func (o *PaymentIntent) GetLastPaymentError() PaymentIntentLastPaymentError {
	if o == nil || IsNil(o.LastPaymentError) {
		var ret PaymentIntentLastPaymentError
		return ret
	}
	return *o.LastPaymentError
}

// GetLastPaymentErrorOk returns a tuple with the LastPaymentError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetLastPaymentErrorOk() (*PaymentIntentLastPaymentError, bool) {
	if o == nil || IsNil(o.LastPaymentError) {
		return nil, false
	}
	return o.LastPaymentError, true
}

// HasLastPaymentError returns a boolean if a field has been set.
func (o *PaymentIntent) HasLastPaymentError() bool {
	if o != nil && !IsNil(o.LastPaymentError) {
		return true
	}

	return false
}

// SetLastPaymentError gets a reference to the given PaymentIntentLastPaymentError and assigns it to the LastPaymentError field.
func (o *PaymentIntent) SetLastPaymentError(v PaymentIntentLastPaymentError) {
	o.LastPaymentError = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentIntent) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentIntent) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentIntent) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetNextAction returns the NextAction field value if set, zero value otherwise.
func (o *PaymentIntent) GetNextAction() PaymentIntentNextAction {
	if o == nil || IsNil(o.NextAction) {
		var ret PaymentIntentNextAction
		return ret
	}
	return *o.NextAction
}

// GetNextActionOk returns a tuple with the NextAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetNextActionOk() (*PaymentIntentNextAction, bool) {
	if o == nil || IsNil(o.NextAction) {
		return nil, false
	}
	return o.NextAction, true
}

// HasNextAction returns a boolean if a field has been set.
func (o *PaymentIntent) HasNextAction() bool {
	if o != nil && !IsNil(o.NextAction) {
		return true
	}

	return false
}

// SetNextAction gets a reference to the given PaymentIntentNextAction and assigns it to the NextAction field.
func (o *PaymentIntent) SetNextAction(v PaymentIntentNextAction) {
	o.NextAction = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod.Get()) {
		var ret string
		return ret
	}
	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentIntent) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given NullableString and assigns it to the PaymentMethod field.
func (o *PaymentIntent) SetPaymentMethod(v string) {
	o.PaymentMethod.Set(&v)
}
// SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil
func (o *PaymentIntent) SetPaymentMethodNil() {
	o.PaymentMethod.Set(nil)
}

// UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
func (o *PaymentIntent) UnsetPaymentMethod() {
	o.PaymentMethod.Unset()
}

// GetReceiptEmail returns the ReceiptEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetReceiptEmail() string {
	if o == nil || IsNil(o.ReceiptEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiptEmail.Get()
}

// GetReceiptEmailOk returns a tuple with the ReceiptEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetReceiptEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiptEmail.Get(), o.ReceiptEmail.IsSet()
}

// HasReceiptEmail returns a boolean if a field has been set.
func (o *PaymentIntent) HasReceiptEmail() bool {
	if o != nil && o.ReceiptEmail.IsSet() {
		return true
	}

	return false
}

// SetReceiptEmail gets a reference to the given NullableString and assigns it to the ReceiptEmail field.
func (o *PaymentIntent) SetReceiptEmail(v string) {
	o.ReceiptEmail.Set(&v)
}
// SetReceiptEmailNil sets the value for ReceiptEmail to be an explicit nil
func (o *PaymentIntent) SetReceiptEmailNil() {
	o.ReceiptEmail.Set(nil)
}

// UnsetReceiptEmail ensures that no value is present for ReceiptEmail, not even an explicit nil
func (o *PaymentIntent) UnsetReceiptEmail() {
	o.ReceiptEmail.Unset()
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetStatementDescriptor() string {
	if o == nil || IsNil(o.StatementDescriptor.Get()) {
		var ret string
		return ret
	}
	return *o.StatementDescriptor.Get()
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetStatementDescriptorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatementDescriptor.Get(), o.StatementDescriptor.IsSet()
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *PaymentIntent) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor.IsSet() {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given NullableString and assigns it to the StatementDescriptor field.
func (o *PaymentIntent) SetStatementDescriptor(v string) {
	o.StatementDescriptor.Set(&v)
}
// SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil
func (o *PaymentIntent) SetStatementDescriptorNil() {
	o.StatementDescriptor.Set(nil)
}

// UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
func (o *PaymentIntent) UnsetStatementDescriptor() {
	o.StatementDescriptor.Unset()
}

// GetStatementDescriptorSuffix returns the StatementDescriptorSuffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetStatementDescriptorSuffix() string {
	if o == nil || IsNil(o.StatementDescriptorSuffix.Get()) {
		var ret string
		return ret
	}
	return *o.StatementDescriptorSuffix.Get()
}

// GetStatementDescriptorSuffixOk returns a tuple with the StatementDescriptorSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetStatementDescriptorSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatementDescriptorSuffix.Get(), o.StatementDescriptorSuffix.IsSet()
}

// HasStatementDescriptorSuffix returns a boolean if a field has been set.
func (o *PaymentIntent) HasStatementDescriptorSuffix() bool {
	if o != nil && o.StatementDescriptorSuffix.IsSet() {
		return true
	}

	return false
}

// SetStatementDescriptorSuffix gets a reference to the given NullableString and assigns it to the StatementDescriptorSuffix field.
func (o *PaymentIntent) SetStatementDescriptorSuffix(v string) {
	o.StatementDescriptorSuffix.Set(&v)
}
// SetStatementDescriptorSuffixNil sets the value for StatementDescriptorSuffix to be an explicit nil
func (o *PaymentIntent) SetStatementDescriptorSuffixNil() {
	o.StatementDescriptorSuffix.Set(nil)
}

// UnsetStatementDescriptorSuffix ensures that no value is present for StatementDescriptorSuffix, not even an explicit nil
func (o *PaymentIntent) UnsetStatementDescriptorSuffix() {
	o.StatementDescriptorSuffix.Unset()
}

// GetStatus returns the Status field value
func (o *PaymentIntent) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentIntent) SetStatus(v string) {
	o.Status = v
}

// GetCaptureMethod returns the CaptureMethod field value
func (o *PaymentIntent) GetCaptureMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaptureMethod
}

// GetCaptureMethodOk returns a tuple with the CaptureMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCaptureMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptureMethod, true
}

// SetCaptureMethod sets field value
func (o *PaymentIntent) SetCaptureMethod(v string) {
	o.CaptureMethod = v
}

// GetClientSecret returns the ClientSecret field value
func (o *PaymentIntent) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *PaymentIntent) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *PaymentIntent) HasCanceledAt() bool {
	if o != nil && o.CanceledAt.IsSet() {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given NullableTime and assigns it to the CanceledAt field.
func (o *PaymentIntent) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}
// SetCanceledAtNil sets the value for CanceledAt to be an explicit nil
func (o *PaymentIntent) SetCanceledAtNil() {
	o.CanceledAt.Set(nil)
}

// UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
func (o *PaymentIntent) UnsetCanceledAt() {
	o.CanceledAt.Unset()
}

// GetCancellationReason returns the CancellationReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentIntent) GetCancellationReason() string {
	if o == nil || IsNil(o.CancellationReason.Get()) {
		var ret string
		return ret
	}
	return *o.CancellationReason.Get()
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentIntent) GetCancellationReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancellationReason.Get(), o.CancellationReason.IsSet()
}

// HasCancellationReason returns a boolean if a field has been set.
func (o *PaymentIntent) HasCancellationReason() bool {
	if o != nil && o.CancellationReason.IsSet() {
		return true
	}

	return false
}

// SetCancellationReason gets a reference to the given NullableString and assigns it to the CancellationReason field.
func (o *PaymentIntent) SetCancellationReason(v string) {
	o.CancellationReason.Set(&v)
}
// SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil
func (o *PaymentIntent) SetCancellationReasonNil() {
	o.CancellationReason.Set(nil)
}

// UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
func (o *PaymentIntent) UnsetCancellationReason() {
	o.CancellationReason.Unset()
}

// GetCreated returns the Created field value
func (o *PaymentIntent) GetCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCreatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *PaymentIntent) SetCreated(v int32) {
	o.Created = v
}

// GetLivemode returns the Livemode field value
func (o *PaymentIntent) GetLivemode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Livemode
}

// GetLivemodeOk returns a tuple with the Livemode field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetLivemodeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Livemode, true
}

// SetLivemode sets field value
func (o *PaymentIntent) SetLivemode(v bool) {
	o.Livemode = v
}

func (o PaymentIntent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentIntent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["amount"] = o.Amount
	toSerialize["amount_capturable"] = o.AmountCapturable
	toSerialize["amount_received"] = o.AmountReceived
	toSerialize["currency"] = o.Currency
	toSerialize["customer"] = o.Customer
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.LastPaymentError) {
		toSerialize["last_payment_error"] = o.LastPaymentError
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.NextAction) {
		toSerialize["next_action"] = o.NextAction
	}
	if o.PaymentMethod.IsSet() {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	if o.ReceiptEmail.IsSet() {
		toSerialize["receipt_email"] = o.ReceiptEmail.Get()
	}
	if o.StatementDescriptor.IsSet() {
		toSerialize["statement_descriptor"] = o.StatementDescriptor.Get()
	}
	if o.StatementDescriptorSuffix.IsSet() {
		toSerialize["statement_descriptor_suffix"] = o.StatementDescriptorSuffix.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["capture_method"] = o.CaptureMethod
	toSerialize["client_secret"] = o.ClientSecret
	if o.CanceledAt.IsSet() {
		toSerialize["canceled_at"] = o.CanceledAt.Get()
	}
	if o.CancellationReason.IsSet() {
		toSerialize["cancellation_reason"] = o.CancellationReason.Get()
	}
	toSerialize["created"] = o.Created
	toSerialize["livemode"] = o.Livemode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentIntent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"amount",
		"amount_capturable",
		"amount_received",
		"currency",
		"customer",
		"status",
		"capture_method",
		"client_secret",
		"created",
		"livemode",
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

	varPaymentIntent := _PaymentIntent{}

	err = json.Unmarshal(data, &varPaymentIntent)

	if err != nil {
		return err
	}

	*o = PaymentIntent(varPaymentIntent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "amount_capturable")
		delete(additionalProperties, "amount_received")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "customer")
		delete(additionalProperties, "description")
		delete(additionalProperties, "last_payment_error")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "next_action")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "receipt_email")
		delete(additionalProperties, "statement_descriptor")
		delete(additionalProperties, "statement_descriptor_suffix")
		delete(additionalProperties, "status")
		delete(additionalProperties, "capture_method")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "cancellation_reason")
		delete(additionalProperties, "created")
		delete(additionalProperties, "livemode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentIntent struct {
	value *PaymentIntent
	isSet bool
}

func (v NullablePaymentIntent) Get() *PaymentIntent {
	return v.value
}

func (v *NullablePaymentIntent) Set(val *PaymentIntent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentIntent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentIntent(val *PaymentIntent) *NullablePaymentIntent {
	return &NullablePaymentIntent{value: val, isSet: true}
}

func (v NullablePaymentIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


