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

// checks if the SubscriptionsCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsCreateRequest{}

// SubscriptionsCreateRequest struct for SubscriptionsCreateRequest
type SubscriptionsCreateRequest struct {
	CustomerId string `json:"customerId"`
	BusinessId *string `json:"businessId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	PlanId *string `json:"planId,omitempty"`
	PriceId *string `json:"priceId,omitempty"`
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	TrialEnd *time.Time `json:"trialEnd,omitempty"`
	BillingAnchorDay *int32 `json:"billingAnchorDay,omitempty"`
	PauseAtPeriodEnd *bool `json:"pauseAtPeriodEnd,omitempty"`
	CancelAtPeriodEnd *bool `json:"cancelAtPeriodEnd,omitempty"`
	BillingStartsOn *time.Time `json:"billingStartsOn,omitempty"`
	BillingEndStrategy *string `json:"billingEndStrategy,omitempty"`
	BillingEndsOn *time.Time `json:"billingEndsOn,omitempty"`
	BillingMaxCycles *int32 `json:"billingMaxCycles,omitempty"`
	ProrationMode *string `json:"prorationMode,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionsCreateRequest SubscriptionsCreateRequest

// NewSubscriptionsCreateRequest instantiates a new SubscriptionsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsCreateRequest(customerId string) *SubscriptionsCreateRequest {
	this := SubscriptionsCreateRequest{}
	this.CustomerId = customerId
	var currency string = "USD"
	this.Currency = &currency
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// NewSubscriptionsCreateRequestWithDefaults instantiates a new SubscriptionsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsCreateRequestWithDefaults() *SubscriptionsCreateRequest {
	this := SubscriptionsCreateRequest{}
	var currency string = "USD"
	this.Currency = &currency
	var quantity int32 = 1
	this.Quantity = &quantity
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *SubscriptionsCreateRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *SubscriptionsCreateRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *SubscriptionsCreateRequest) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *SubscriptionsCreateRequest) SetProductId(v string) {
	o.ProductId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionsCreateRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetPriceId() string {
	if o == nil || IsNil(o.PriceId) {
		var ret string
		return ret
	}
	return *o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetPriceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PriceId) {
		return nil, false
	}
	return o.PriceId, true
}

// HasPriceId returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasPriceId() bool {
	if o != nil && !IsNil(o.PriceId) {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given string and assigns it to the PriceId field.
func (o *SubscriptionsCreateRequest) SetPriceId(v string) {
	o.PriceId = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *SubscriptionsCreateRequest) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SubscriptionsCreateRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SubscriptionsCreateRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SubscriptionsCreateRequest) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetTrialEnd() time.Time {
	if o == nil || IsNil(o.TrialEnd) {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetTrialEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given time.Time and assigns it to the TrialEnd field.
func (o *SubscriptionsCreateRequest) SetTrialEnd(v time.Time) {
	o.TrialEnd = &v
}

// GetBillingAnchorDay returns the BillingAnchorDay field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBillingAnchorDay() int32 {
	if o == nil || IsNil(o.BillingAnchorDay) {
		var ret int32
		return ret
	}
	return *o.BillingAnchorDay
}

// GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBillingAnchorDayOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingAnchorDay) {
		return nil, false
	}
	return o.BillingAnchorDay, true
}

// HasBillingAnchorDay returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBillingAnchorDay() bool {
	if o != nil && !IsNil(o.BillingAnchorDay) {
		return true
	}

	return false
}

// SetBillingAnchorDay gets a reference to the given int32 and assigns it to the BillingAnchorDay field.
func (o *SubscriptionsCreateRequest) SetBillingAnchorDay(v int32) {
	o.BillingAnchorDay = &v
}

// GetPauseAtPeriodEnd returns the PauseAtPeriodEnd field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetPauseAtPeriodEnd() bool {
	if o == nil || IsNil(o.PauseAtPeriodEnd) {
		var ret bool
		return ret
	}
	return *o.PauseAtPeriodEnd
}

// GetPauseAtPeriodEndOk returns a tuple with the PauseAtPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetPauseAtPeriodEndOk() (*bool, bool) {
	if o == nil || IsNil(o.PauseAtPeriodEnd) {
		return nil, false
	}
	return o.PauseAtPeriodEnd, true
}

// HasPauseAtPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasPauseAtPeriodEnd() bool {
	if o != nil && !IsNil(o.PauseAtPeriodEnd) {
		return true
	}

	return false
}

// SetPauseAtPeriodEnd gets a reference to the given bool and assigns it to the PauseAtPeriodEnd field.
func (o *SubscriptionsCreateRequest) SetPauseAtPeriodEnd(v bool) {
	o.PauseAtPeriodEnd = &v
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetCancelAtPeriodEnd() bool {
	if o == nil || IsNil(o.CancelAtPeriodEnd) {
		var ret bool
		return ret
	}
	return *o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetCancelAtPeriodEndOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelAtPeriodEnd) {
		return nil, false
	}
	return o.CancelAtPeriodEnd, true
}

// HasCancelAtPeriodEnd returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasCancelAtPeriodEnd() bool {
	if o != nil && !IsNil(o.CancelAtPeriodEnd) {
		return true
	}

	return false
}

// SetCancelAtPeriodEnd gets a reference to the given bool and assigns it to the CancelAtPeriodEnd field.
func (o *SubscriptionsCreateRequest) SetCancelAtPeriodEnd(v bool) {
	o.CancelAtPeriodEnd = &v
}

// GetBillingStartsOn returns the BillingStartsOn field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBillingStartsOn() time.Time {
	if o == nil || IsNil(o.BillingStartsOn) {
		var ret time.Time
		return ret
	}
	return *o.BillingStartsOn
}

// GetBillingStartsOnOk returns a tuple with the BillingStartsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBillingStartsOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BillingStartsOn) {
		return nil, false
	}
	return o.BillingStartsOn, true
}

// HasBillingStartsOn returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBillingStartsOn() bool {
	if o != nil && !IsNil(o.BillingStartsOn) {
		return true
	}

	return false
}

// SetBillingStartsOn gets a reference to the given time.Time and assigns it to the BillingStartsOn field.
func (o *SubscriptionsCreateRequest) SetBillingStartsOn(v time.Time) {
	o.BillingStartsOn = &v
}

// GetBillingEndStrategy returns the BillingEndStrategy field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBillingEndStrategy() string {
	if o == nil || IsNil(o.BillingEndStrategy) {
		var ret string
		return ret
	}
	return *o.BillingEndStrategy
}

// GetBillingEndStrategyOk returns a tuple with the BillingEndStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBillingEndStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.BillingEndStrategy) {
		return nil, false
	}
	return o.BillingEndStrategy, true
}

// HasBillingEndStrategy returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBillingEndStrategy() bool {
	if o != nil && !IsNil(o.BillingEndStrategy) {
		return true
	}

	return false
}

// SetBillingEndStrategy gets a reference to the given string and assigns it to the BillingEndStrategy field.
func (o *SubscriptionsCreateRequest) SetBillingEndStrategy(v string) {
	o.BillingEndStrategy = &v
}

// GetBillingEndsOn returns the BillingEndsOn field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBillingEndsOn() time.Time {
	if o == nil || IsNil(o.BillingEndsOn) {
		var ret time.Time
		return ret
	}
	return *o.BillingEndsOn
}

// GetBillingEndsOnOk returns a tuple with the BillingEndsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBillingEndsOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.BillingEndsOn) {
		return nil, false
	}
	return o.BillingEndsOn, true
}

// HasBillingEndsOn returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBillingEndsOn() bool {
	if o != nil && !IsNil(o.BillingEndsOn) {
		return true
	}

	return false
}

// SetBillingEndsOn gets a reference to the given time.Time and assigns it to the BillingEndsOn field.
func (o *SubscriptionsCreateRequest) SetBillingEndsOn(v time.Time) {
	o.BillingEndsOn = &v
}

// GetBillingMaxCycles returns the BillingMaxCycles field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetBillingMaxCycles() int32 {
	if o == nil || IsNil(o.BillingMaxCycles) {
		var ret int32
		return ret
	}
	return *o.BillingMaxCycles
}

// GetBillingMaxCyclesOk returns a tuple with the BillingMaxCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetBillingMaxCyclesOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingMaxCycles) {
		return nil, false
	}
	return o.BillingMaxCycles, true
}

// HasBillingMaxCycles returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasBillingMaxCycles() bool {
	if o != nil && !IsNil(o.BillingMaxCycles) {
		return true
	}

	return false
}

// SetBillingMaxCycles gets a reference to the given int32 and assigns it to the BillingMaxCycles field.
func (o *SubscriptionsCreateRequest) SetBillingMaxCycles(v int32) {
	o.BillingMaxCycles = &v
}

// GetProrationMode returns the ProrationMode field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetProrationMode() string {
	if o == nil || IsNil(o.ProrationMode) {
		var ret string
		return ret
	}
	return *o.ProrationMode
}

// GetProrationModeOk returns a tuple with the ProrationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetProrationModeOk() (*string, bool) {
	if o == nil || IsNil(o.ProrationMode) {
		return nil, false
	}
	return o.ProrationMode, true
}

// HasProrationMode returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasProrationMode() bool {
	if o != nil && !IsNil(o.ProrationMode) {
		return true
	}

	return false
}

// SetProrationMode gets a reference to the given string and assigns it to the ProrationMode field.
func (o *SubscriptionsCreateRequest) SetProrationMode(v string) {
	o.ProrationMode = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubscriptionsCreateRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsCreateRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubscriptionsCreateRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SubscriptionsCreateRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SubscriptionsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerId"] = o.CustomerId
	if !IsNil(o.BusinessId) {
		toSerialize["businessId"] = o.BusinessId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.PriceId) {
		toSerialize["priceId"] = o.PriceId
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TrialEnd) {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	if !IsNil(o.BillingAnchorDay) {
		toSerialize["billingAnchorDay"] = o.BillingAnchorDay
	}
	if !IsNil(o.PauseAtPeriodEnd) {
		toSerialize["pauseAtPeriodEnd"] = o.PauseAtPeriodEnd
	}
	if !IsNil(o.CancelAtPeriodEnd) {
		toSerialize["cancelAtPeriodEnd"] = o.CancelAtPeriodEnd
	}
	if !IsNil(o.BillingStartsOn) {
		toSerialize["billingStartsOn"] = o.BillingStartsOn
	}
	if !IsNil(o.BillingEndStrategy) {
		toSerialize["billingEndStrategy"] = o.BillingEndStrategy
	}
	if !IsNil(o.BillingEndsOn) {
		toSerialize["billingEndsOn"] = o.BillingEndsOn
	}
	if !IsNil(o.BillingMaxCycles) {
		toSerialize["billingMaxCycles"] = o.BillingMaxCycles
	}
	if !IsNil(o.ProrationMode) {
		toSerialize["prorationMode"] = o.ProrationMode
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionsCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerId",
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

	varSubscriptionsCreateRequest := _SubscriptionsCreateRequest{}

	err = json.Unmarshal(data, &varSubscriptionsCreateRequest)

	if err != nil {
		return err
	}

	*o = SubscriptionsCreateRequest(varSubscriptionsCreateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "businessId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "planId")
		delete(additionalProperties, "priceId")
		delete(additionalProperties, "paymentMethodId")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "trialEnd")
		delete(additionalProperties, "billingAnchorDay")
		delete(additionalProperties, "pauseAtPeriodEnd")
		delete(additionalProperties, "cancelAtPeriodEnd")
		delete(additionalProperties, "billingStartsOn")
		delete(additionalProperties, "billingEndStrategy")
		delete(additionalProperties, "billingEndsOn")
		delete(additionalProperties, "billingMaxCycles")
		delete(additionalProperties, "prorationMode")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionsCreateRequest struct {
	value *SubscriptionsCreateRequest
	isSet bool
}

func (v NullableSubscriptionsCreateRequest) Get() *SubscriptionsCreateRequest {
	return v.value
}

func (v *NullableSubscriptionsCreateRequest) Set(val *SubscriptionsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsCreateRequest(val *SubscriptionsCreateRequest) *NullableSubscriptionsCreateRequest {
	return &NullableSubscriptionsCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


