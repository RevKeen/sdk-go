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
)

// checks if the InvoiceLineUsageEventProvenanceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceLineUsageEventProvenanceItem{}

// InvoiceLineUsageEventProvenanceItem struct for InvoiceLineUsageEventProvenanceItem
type InvoiceLineUsageEventProvenanceItem struct {
	Outcome              string                                           `json:"outcome"`
	ExclusionReason      NullableString                                   `json:"exclusion_reason,omitempty"`
	Entry                NullableInvoiceLineUsageEventProvenanceItemEntry `json:"entry,omitempty"`
	Match                NullableInvoiceLineUsageEventProvenanceItemMatch `json:"match,omitempty"`
	Event                InvoiceLineUsageEventProvenanceItemEvent         `json:"event"`
	AdditionalProperties map[string]interface{}
}

type _InvoiceLineUsageEventProvenanceItem InvoiceLineUsageEventProvenanceItem

// NewInvoiceLineUsageEventProvenanceItem instantiates a new InvoiceLineUsageEventProvenanceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceLineUsageEventProvenanceItem(outcome string, event InvoiceLineUsageEventProvenanceItemEvent) *InvoiceLineUsageEventProvenanceItem {
	this := InvoiceLineUsageEventProvenanceItem{}
	this.Outcome = outcome
	this.Event = event
	return &this
}

// NewInvoiceLineUsageEventProvenanceItemWithDefaults instantiates a new InvoiceLineUsageEventProvenanceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceLineUsageEventProvenanceItemWithDefaults() *InvoiceLineUsageEventProvenanceItem {
	this := InvoiceLineUsageEventProvenanceItem{}
	return &this
}

// GetOutcome returns the Outcome field value
func (o *InvoiceLineUsageEventProvenanceItem) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *InvoiceLineUsageEventProvenanceItem) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *InvoiceLineUsageEventProvenanceItem) SetOutcome(v string) {
	o.Outcome = v
}

// GetExclusionReason returns the ExclusionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUsageEventProvenanceItem) GetExclusionReason() string {
	if o == nil || IsNil(o.ExclusionReason.Get()) {
		var ret string
		return ret
	}
	return *o.ExclusionReason.Get()
}

// GetExclusionReasonOk returns a tuple with the ExclusionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUsageEventProvenanceItem) GetExclusionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExclusionReason.Get(), o.ExclusionReason.IsSet()
}

// HasExclusionReason returns a boolean if a field has been set.
func (o *InvoiceLineUsageEventProvenanceItem) HasExclusionReason() bool {
	if o != nil && o.ExclusionReason.IsSet() {
		return true
	}

	return false
}

// SetExclusionReason gets a reference to the given NullableString and assigns it to the ExclusionReason field.
func (o *InvoiceLineUsageEventProvenanceItem) SetExclusionReason(v string) {
	o.ExclusionReason.Set(&v)
}

// SetExclusionReasonNil sets the value for ExclusionReason to be an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) SetExclusionReasonNil() {
	o.ExclusionReason.Set(nil)
}

// UnsetExclusionReason ensures that no value is present for ExclusionReason, not even an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) UnsetExclusionReason() {
	o.ExclusionReason.Unset()
}

// GetEntry returns the Entry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUsageEventProvenanceItem) GetEntry() InvoiceLineUsageEventProvenanceItemEntry {
	if o == nil || IsNil(o.Entry.Get()) {
		var ret InvoiceLineUsageEventProvenanceItemEntry
		return ret
	}
	return *o.Entry.Get()
}

// GetEntryOk returns a tuple with the Entry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUsageEventProvenanceItem) GetEntryOk() (*InvoiceLineUsageEventProvenanceItemEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entry.Get(), o.Entry.IsSet()
}

// HasEntry returns a boolean if a field has been set.
func (o *InvoiceLineUsageEventProvenanceItem) HasEntry() bool {
	if o != nil && o.Entry.IsSet() {
		return true
	}

	return false
}

// SetEntry gets a reference to the given NullableInvoiceLineUsageEventProvenanceItemEntry and assigns it to the Entry field.
func (o *InvoiceLineUsageEventProvenanceItem) SetEntry(v InvoiceLineUsageEventProvenanceItemEntry) {
	o.Entry.Set(&v)
}

// SetEntryNil sets the value for Entry to be an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) SetEntryNil() {
	o.Entry.Set(nil)
}

// UnsetEntry ensures that no value is present for Entry, not even an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) UnsetEntry() {
	o.Entry.Unset()
}

// GetMatch returns the Match field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineUsageEventProvenanceItem) GetMatch() InvoiceLineUsageEventProvenanceItemMatch {
	if o == nil || IsNil(o.Match.Get()) {
		var ret InvoiceLineUsageEventProvenanceItemMatch
		return ret
	}
	return *o.Match.Get()
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineUsageEventProvenanceItem) GetMatchOk() (*InvoiceLineUsageEventProvenanceItemMatch, bool) {
	if o == nil {
		return nil, false
	}
	return o.Match.Get(), o.Match.IsSet()
}

// HasMatch returns a boolean if a field has been set.
func (o *InvoiceLineUsageEventProvenanceItem) HasMatch() bool {
	if o != nil && o.Match.IsSet() {
		return true
	}

	return false
}

// SetMatch gets a reference to the given NullableInvoiceLineUsageEventProvenanceItemMatch and assigns it to the Match field.
func (o *InvoiceLineUsageEventProvenanceItem) SetMatch(v InvoiceLineUsageEventProvenanceItemMatch) {
	o.Match.Set(&v)
}

// SetMatchNil sets the value for Match to be an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) SetMatchNil() {
	o.Match.Set(nil)
}

// UnsetMatch ensures that no value is present for Match, not even an explicit nil
func (o *InvoiceLineUsageEventProvenanceItem) UnsetMatch() {
	o.Match.Unset()
}

// GetEvent returns the Event field value
func (o *InvoiceLineUsageEventProvenanceItem) GetEvent() InvoiceLineUsageEventProvenanceItemEvent {
	if o == nil {
		var ret InvoiceLineUsageEventProvenanceItemEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *InvoiceLineUsageEventProvenanceItem) GetEventOk() (*InvoiceLineUsageEventProvenanceItemEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *InvoiceLineUsageEventProvenanceItem) SetEvent(v InvoiceLineUsageEventProvenanceItemEvent) {
	o.Event = v
}

func (o InvoiceLineUsageEventProvenanceItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceLineUsageEventProvenanceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["outcome"] = o.Outcome
	if o.ExclusionReason.IsSet() {
		toSerialize["exclusion_reason"] = o.ExclusionReason.Get()
	}
	if o.Entry.IsSet() {
		toSerialize["entry"] = o.Entry.Get()
	}
	if o.Match.IsSet() {
		toSerialize["match"] = o.Match.Get()
	}
	toSerialize["event"] = o.Event

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvoiceLineUsageEventProvenanceItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"outcome",
		"event",
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

	varInvoiceLineUsageEventProvenanceItem := _InvoiceLineUsageEventProvenanceItem{}

	err = json.Unmarshal(data, &varInvoiceLineUsageEventProvenanceItem)

	if err != nil {
		return err
	}

	*o = InvoiceLineUsageEventProvenanceItem(varInvoiceLineUsageEventProvenanceItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "exclusion_reason")
		delete(additionalProperties, "entry")
		delete(additionalProperties, "match")
		delete(additionalProperties, "event")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvoiceLineUsageEventProvenanceItem struct {
	value *InvoiceLineUsageEventProvenanceItem
	isSet bool
}

func (v NullableInvoiceLineUsageEventProvenanceItem) Get() *InvoiceLineUsageEventProvenanceItem {
	return v.value
}

func (v *NullableInvoiceLineUsageEventProvenanceItem) Set(val *InvoiceLineUsageEventProvenanceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceLineUsageEventProvenanceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceLineUsageEventProvenanceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceLineUsageEventProvenanceItem(val *InvoiceLineUsageEventProvenanceItem) *NullableInvoiceLineUsageEventProvenanceItem {
	return &NullableInvoiceLineUsageEventProvenanceItem{value: val, isSet: true}
}

func (v NullableInvoiceLineUsageEventProvenanceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceLineUsageEventProvenanceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
