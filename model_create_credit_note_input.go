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
	"fmt"
)

// checks if the CreateCreditNoteInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCreditNoteInput{}

// CreateCreditNoteInput Parameters for creating a credit note against an invoice, specifying the amount, reason, and whether to issue a refund.
type CreateCreditNoteInput struct {
	// The ID of the invoice to issue a credit note for
	InvoiceId string `json:"invoice_id"`
	// Amount to credit in minor units (cents). Capped at 2,000,000,000 to stay below the 32-bit integer ceiling of the underlying DB column.
	AmountMinor int32 `json:"amount_minor"`
	// Tax portion of the credit in minor units (cents). For UK/EU VAT compliance. Capped at 2,000,000,000.
	TaxAmountMinor *int32 `json:"tax_amount_minor,omitempty"`
	// How the credit should be applied
	CreditMethod *string `json:"credit_method,omitempty"`
	// Reason for the credit note
	Reason *string `json:"reason,omitempty"`
	// Standardized reason code
	ReasonCode *string `json:"reason_code,omitempty"`
	// Whether to cancel the associated subscription after issuing
	CancelSubscription *bool `json:"cancel_subscription,omitempty"`
	// Whether this credit note represents a prorated amount
	IsProrated *bool `json:"is_prorated,omitempty"`
	// Total days in the billing period (for prorated credits). Constrained to 1–366.
	ProrationDaysTotal *int32 `json:"proration_days_total,omitempty"`
	// Unused days in the billing period (for prorated credits). Must be <= proration_days_total. Constrained to 0–366.
	ProrationDaysUnused *int32 `json:"proration_days_unused,omitempty"`
	// Idempotency key to prevent duplicate credit notes
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// Arbitrary key-value metadata
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// When true, automatically trigger the process-credit-note task for multi-gateway reversal routing. Response will include gateway_operations array. Always combine with `idempotency_key` to avoid duplicate gateway operations.
	AutoRoute *bool `json:"auto_route,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateCreditNoteInput CreateCreditNoteInput

// NewCreateCreditNoteInput instantiates a new CreateCreditNoteInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCreditNoteInput(invoiceId string, amountMinor int32) *CreateCreditNoteInput {
	this := CreateCreditNoteInput{}
	this.InvoiceId = invoiceId
	this.AmountMinor = amountMinor
	var creditMethod string = "refund_to_payment_method"
	this.CreditMethod = &creditMethod
	var cancelSubscription bool = false
	this.CancelSubscription = &cancelSubscription
	var autoRoute bool = false
	this.AutoRoute = &autoRoute
	return &this
}

// NewCreateCreditNoteInputWithDefaults instantiates a new CreateCreditNoteInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCreditNoteInputWithDefaults() *CreateCreditNoteInput {
	this := CreateCreditNoteInput{}
	var creditMethod string = "refund_to_payment_method"
	this.CreditMethod = &creditMethod
	var cancelSubscription bool = false
	this.CancelSubscription = &cancelSubscription
	var autoRoute bool = false
	this.AutoRoute = &autoRoute
	return &this
}

// GetInvoiceId returns the InvoiceId field value
func (o *CreateCreditNoteInput) GetInvoiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetInvoiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceId, true
}

// SetInvoiceId sets field value
func (o *CreateCreditNoteInput) SetInvoiceId(v string) {
	o.InvoiceId = v
}

// GetAmountMinor returns the AmountMinor field value
func (o *CreateCreditNoteInput) GetAmountMinor() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountMinor
}

// GetAmountMinorOk returns a tuple with the AmountMinor field value
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetAmountMinorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountMinor, true
}

// SetAmountMinor sets field value
func (o *CreateCreditNoteInput) SetAmountMinor(v int32) {
	o.AmountMinor = v
}

// GetTaxAmountMinor returns the TaxAmountMinor field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetTaxAmountMinor() int32 {
	if o == nil || IsNil(o.TaxAmountMinor) {
		var ret int32
		return ret
	}
	return *o.TaxAmountMinor
}

// GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetTaxAmountMinorOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxAmountMinor) {
		return nil, false
	}
	return o.TaxAmountMinor, true
}

// HasTaxAmountMinor returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasTaxAmountMinor() bool {
	if o != nil && !IsNil(o.TaxAmountMinor) {
		return true
	}

	return false
}

// SetTaxAmountMinor gets a reference to the given int32 and assigns it to the TaxAmountMinor field.
func (o *CreateCreditNoteInput) SetTaxAmountMinor(v int32) {
	o.TaxAmountMinor = &v
}

// GetCreditMethod returns the CreditMethod field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetCreditMethod() string {
	if o == nil || IsNil(o.CreditMethod) {
		var ret string
		return ret
	}
	return *o.CreditMethod
}

// GetCreditMethodOk returns a tuple with the CreditMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetCreditMethodOk() (*string, bool) {
	if o == nil || IsNil(o.CreditMethod) {
		return nil, false
	}
	return o.CreditMethod, true
}

// HasCreditMethod returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasCreditMethod() bool {
	if o != nil && !IsNil(o.CreditMethod) {
		return true
	}

	return false
}

// SetCreditMethod gets a reference to the given string and assigns it to the CreditMethod field.
func (o *CreateCreditNoteInput) SetCreditMethod(v string) {
	o.CreditMethod = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CreateCreditNoteInput) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *CreateCreditNoteInput) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetCancelSubscription returns the CancelSubscription field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetCancelSubscription() bool {
	if o == nil || IsNil(o.CancelSubscription) {
		var ret bool
		return ret
	}
	return *o.CancelSubscription
}

// GetCancelSubscriptionOk returns a tuple with the CancelSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetCancelSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.CancelSubscription) {
		return nil, false
	}
	return o.CancelSubscription, true
}

// HasCancelSubscription returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasCancelSubscription() bool {
	if o != nil && !IsNil(o.CancelSubscription) {
		return true
	}

	return false
}

// SetCancelSubscription gets a reference to the given bool and assigns it to the CancelSubscription field.
func (o *CreateCreditNoteInput) SetCancelSubscription(v bool) {
	o.CancelSubscription = &v
}

// GetIsProrated returns the IsProrated field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetIsProrated() bool {
	if o == nil || IsNil(o.IsProrated) {
		var ret bool
		return ret
	}
	return *o.IsProrated
}

// GetIsProratedOk returns a tuple with the IsProrated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetIsProratedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProrated) {
		return nil, false
	}
	return o.IsProrated, true
}

// HasIsProrated returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasIsProrated() bool {
	if o != nil && !IsNil(o.IsProrated) {
		return true
	}

	return false
}

// SetIsProrated gets a reference to the given bool and assigns it to the IsProrated field.
func (o *CreateCreditNoteInput) SetIsProrated(v bool) {
	o.IsProrated = &v
}

// GetProrationDaysTotal returns the ProrationDaysTotal field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetProrationDaysTotal() int32 {
	if o == nil || IsNil(o.ProrationDaysTotal) {
		var ret int32
		return ret
	}
	return *o.ProrationDaysTotal
}

// GetProrationDaysTotalOk returns a tuple with the ProrationDaysTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetProrationDaysTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.ProrationDaysTotal) {
		return nil, false
	}
	return o.ProrationDaysTotal, true
}

// HasProrationDaysTotal returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasProrationDaysTotal() bool {
	if o != nil && !IsNil(o.ProrationDaysTotal) {
		return true
	}

	return false
}

// SetProrationDaysTotal gets a reference to the given int32 and assigns it to the ProrationDaysTotal field.
func (o *CreateCreditNoteInput) SetProrationDaysTotal(v int32) {
	o.ProrationDaysTotal = &v
}

// GetProrationDaysUnused returns the ProrationDaysUnused field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetProrationDaysUnused() int32 {
	if o == nil || IsNil(o.ProrationDaysUnused) {
		var ret int32
		return ret
	}
	return *o.ProrationDaysUnused
}

// GetProrationDaysUnusedOk returns a tuple with the ProrationDaysUnused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetProrationDaysUnusedOk() (*int32, bool) {
	if o == nil || IsNil(o.ProrationDaysUnused) {
		return nil, false
	}
	return o.ProrationDaysUnused, true
}

// HasProrationDaysUnused returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasProrationDaysUnused() bool {
	if o != nil && !IsNil(o.ProrationDaysUnused) {
		return true
	}

	return false
}

// SetProrationDaysUnused gets a reference to the given int32 and assigns it to the ProrationDaysUnused field.
func (o *CreateCreditNoteInput) SetProrationDaysUnused(v int32) {
	o.ProrationDaysUnused = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *CreateCreditNoteInput) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *CreateCreditNoteInput) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetAutoRoute returns the AutoRoute field value if set, zero value otherwise.
func (o *CreateCreditNoteInput) GetAutoRoute() bool {
	if o == nil || IsNil(o.AutoRoute) {
		var ret bool
		return ret
	}
	return *o.AutoRoute
}

// GetAutoRouteOk returns a tuple with the AutoRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCreditNoteInput) GetAutoRouteOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRoute) {
		return nil, false
	}
	return o.AutoRoute, true
}

// HasAutoRoute returns a boolean if a field has been set.
func (o *CreateCreditNoteInput) HasAutoRoute() bool {
	if o != nil && !IsNil(o.AutoRoute) {
		return true
	}

	return false
}

// SetAutoRoute gets a reference to the given bool and assigns it to the AutoRoute field.
func (o *CreateCreditNoteInput) SetAutoRoute(v bool) {
	o.AutoRoute = &v
}

func (o CreateCreditNoteInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCreditNoteInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice_id"] = o.InvoiceId
	toSerialize["amount_minor"] = o.AmountMinor
	if !IsNil(o.TaxAmountMinor) {
		toSerialize["tax_amount_minor"] = o.TaxAmountMinor
	}
	if !IsNil(o.CreditMethod) {
		toSerialize["credit_method"] = o.CreditMethod
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	if !IsNil(o.CancelSubscription) {
		toSerialize["cancel_subscription"] = o.CancelSubscription
	}
	if !IsNil(o.IsProrated) {
		toSerialize["is_prorated"] = o.IsProrated
	}
	if !IsNil(o.ProrationDaysTotal) {
		toSerialize["proration_days_total"] = o.ProrationDaysTotal
	}
	if !IsNil(o.ProrationDaysUnused) {
		toSerialize["proration_days_unused"] = o.ProrationDaysUnused
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.AutoRoute) {
		toSerialize["auto_route"] = o.AutoRoute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCreditNoteInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_id",
		"amount_minor",
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

	varCreateCreditNoteInput := _CreateCreditNoteInput{}

	err = json.Unmarshal(data, &varCreateCreditNoteInput)

	if err != nil {
		return err
	}

	*o = CreateCreditNoteInput(varCreateCreditNoteInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invoice_id")
		delete(additionalProperties, "amount_minor")
		delete(additionalProperties, "tax_amount_minor")
		delete(additionalProperties, "credit_method")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "reason_code")
		delete(additionalProperties, "cancel_subscription")
		delete(additionalProperties, "is_prorated")
		delete(additionalProperties, "proration_days_total")
		delete(additionalProperties, "proration_days_unused")
		delete(additionalProperties, "idempotency_key")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "auto_route")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCreditNoteInput struct {
	value *CreateCreditNoteInput
	isSet bool
}

func (v NullableCreateCreditNoteInput) Get() *CreateCreditNoteInput {
	return v.value
}

func (v *NullableCreateCreditNoteInput) Set(val *CreateCreditNoteInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCreditNoteInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCreditNoteInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCreditNoteInput(val *CreateCreditNoteInput) *NullableCreateCreditNoteInput {
	return &NullableCreateCreditNoteInput{value: val, isSet: true}
}

func (v NullableCreateCreditNoteInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCreditNoteInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


