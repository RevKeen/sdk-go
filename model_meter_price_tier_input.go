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
)

// checks if the MeterPriceTierInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeterPriceTierInput{}

// MeterPriceTierInput struct for MeterPriceTierInput
type MeterPriceTierInput struct {
	// Upper bound of this tier (1-based inclusive). null = open-ended final tier. New clients should use this canonical shape; legacy first_unit/last_unit is accepted during the deprecation window.
	UpTo NullableInt32 `json:"up_to,omitempty"`
	// Per-unit price in this tier (minor units). For a sub-penny per-tier rate use `unit_amount_decimal`. Optional when `unit_amount_decimal` or `flat_amount_minor` is supplied.
	UnitAmountMinor *int64 `json:"unit_amount_minor,omitempty"`
	// Exact per-unit rate for this tier in minor units as a decimal string (sub-penny safe, e.g. \"0.0000024\" = 0.0000024 pence per token). Takes precedence over `unit_amount_minor`. Must be > 0 with at most 12 decimal places. Amounts are rounded per tier with half-up (away from zero) before the tier subtotals are summed.
	UnitAmountDecimal *string `json:"unit_amount_decimal,omitempty"`
	// Flat fee for entering this tier (minor units). Always a whole number of minor units — only the per-unit RATE supports sub-penny precision.
	FlatAmountMinor      *int64 `json:"flat_amount_minor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MeterPriceTierInput MeterPriceTierInput

// NewMeterPriceTierInput instantiates a new MeterPriceTierInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeterPriceTierInput() *MeterPriceTierInput {
	this := MeterPriceTierInput{}
	return &this
}

// NewMeterPriceTierInputWithDefaults instantiates a new MeterPriceTierInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeterPriceTierInputWithDefaults() *MeterPriceTierInput {
	this := MeterPriceTierInput{}
	return &this
}

// GetUpTo returns the UpTo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeterPriceTierInput) GetUpTo() int32 {
	if o == nil || IsNil(o.UpTo.Get()) {
		var ret int32
		return ret
	}
	return *o.UpTo.Get()
}

// GetUpToOk returns a tuple with the UpTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MeterPriceTierInput) GetUpToOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpTo.Get(), o.UpTo.IsSet()
}

// HasUpTo returns a boolean if a field has been set.
func (o *MeterPriceTierInput) HasUpTo() bool {
	if o != nil && o.UpTo.IsSet() {
		return true
	}

	return false
}

// SetUpTo gets a reference to the given NullableInt32 and assigns it to the UpTo field.
func (o *MeterPriceTierInput) SetUpTo(v int32) {
	o.UpTo.Set(&v)
}

// SetUpToNil sets the value for UpTo to be an explicit nil
func (o *MeterPriceTierInput) SetUpToNil() {
	o.UpTo.Set(nil)
}

// UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
func (o *MeterPriceTierInput) UnsetUpTo() {
	o.UpTo.Unset()
}

// GetUnitAmountMinor returns the UnitAmountMinor field value if set, zero value otherwise.
func (o *MeterPriceTierInput) GetUnitAmountMinor() int64 {
	if o == nil || IsNil(o.UnitAmountMinor) {
		var ret int64
		return ret
	}
	return *o.UnitAmountMinor
}

// GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterPriceTierInput) GetUnitAmountMinorOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitAmountMinor) {
		return nil, false
	}
	return o.UnitAmountMinor, true
}

// HasUnitAmountMinor returns a boolean if a field has been set.
func (o *MeterPriceTierInput) HasUnitAmountMinor() bool {
	if o != nil && !IsNil(o.UnitAmountMinor) {
		return true
	}

	return false
}

// SetUnitAmountMinor gets a reference to the given int64 and assigns it to the UnitAmountMinor field.
func (o *MeterPriceTierInput) SetUnitAmountMinor(v int64) {
	o.UnitAmountMinor = &v
}

// GetUnitAmountDecimal returns the UnitAmountDecimal field value if set, zero value otherwise.
func (o *MeterPriceTierInput) GetUnitAmountDecimal() string {
	if o == nil || IsNil(o.UnitAmountDecimal) {
		var ret string
		return ret
	}
	return *o.UnitAmountDecimal
}

// GetUnitAmountDecimalOk returns a tuple with the UnitAmountDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterPriceTierInput) GetUnitAmountDecimalOk() (*string, bool) {
	if o == nil || IsNil(o.UnitAmountDecimal) {
		return nil, false
	}
	return o.UnitAmountDecimal, true
}

// HasUnitAmountDecimal returns a boolean if a field has been set.
func (o *MeterPriceTierInput) HasUnitAmountDecimal() bool {
	if o != nil && !IsNil(o.UnitAmountDecimal) {
		return true
	}

	return false
}

// SetUnitAmountDecimal gets a reference to the given string and assigns it to the UnitAmountDecimal field.
func (o *MeterPriceTierInput) SetUnitAmountDecimal(v string) {
	o.UnitAmountDecimal = &v
}

// GetFlatAmountMinor returns the FlatAmountMinor field value if set, zero value otherwise.
func (o *MeterPriceTierInput) GetFlatAmountMinor() int64 {
	if o == nil || IsNil(o.FlatAmountMinor) {
		var ret int64
		return ret
	}
	return *o.FlatAmountMinor
}

// GetFlatAmountMinorOk returns a tuple with the FlatAmountMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterPriceTierInput) GetFlatAmountMinorOk() (*int64, bool) {
	if o == nil || IsNil(o.FlatAmountMinor) {
		return nil, false
	}
	return o.FlatAmountMinor, true
}

// HasFlatAmountMinor returns a boolean if a field has been set.
func (o *MeterPriceTierInput) HasFlatAmountMinor() bool {
	if o != nil && !IsNil(o.FlatAmountMinor) {
		return true
	}

	return false
}

// SetFlatAmountMinor gets a reference to the given int64 and assigns it to the FlatAmountMinor field.
func (o *MeterPriceTierInput) SetFlatAmountMinor(v int64) {
	o.FlatAmountMinor = &v
}

func (o MeterPriceTierInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeterPriceTierInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UpTo.IsSet() {
		toSerialize["up_to"] = o.UpTo.Get()
	}
	if !IsNil(o.UnitAmountMinor) {
		toSerialize["unit_amount_minor"] = o.UnitAmountMinor
	}
	if !IsNil(o.UnitAmountDecimal) {
		toSerialize["unit_amount_decimal"] = o.UnitAmountDecimal
	}
	if !IsNil(o.FlatAmountMinor) {
		toSerialize["flat_amount_minor"] = o.FlatAmountMinor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MeterPriceTierInput) UnmarshalJSON(data []byte) (err error) {
	varMeterPriceTierInput := _MeterPriceTierInput{}

	err = json.Unmarshal(data, &varMeterPriceTierInput)

	if err != nil {
		return err
	}

	*o = MeterPriceTierInput(varMeterPriceTierInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "up_to")
		delete(additionalProperties, "unit_amount_minor")
		delete(additionalProperties, "unit_amount_decimal")
		delete(additionalProperties, "flat_amount_minor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeterPriceTierInput struct {
	value *MeterPriceTierInput
	isSet bool
}

func (v NullableMeterPriceTierInput) Get() *MeterPriceTierInput {
	return v.value
}

func (v *NullableMeterPriceTierInput) Set(val *MeterPriceTierInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMeterPriceTierInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMeterPriceTierInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeterPriceTierInput(val *MeterPriceTierInput) *NullableMeterPriceTierInput {
	return &NullableMeterPriceTierInput{value: val, isSet: true}
}

func (v NullableMeterPriceTierInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeterPriceTierInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
