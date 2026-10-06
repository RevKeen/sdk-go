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

// checks if the ApiErrorError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorError{}

// ApiErrorError struct for ApiErrorError
type ApiErrorError struct {
	// The type of error returned
	Type string `json:"type"`
	// Machine-readable error code for programmatic handling
	Code string `json:"code"`
	// Human-readable error message
	Message string `json:"message"`
	// The parameter that caused the error
	Param *string `json:"param,omitempty"`
	// HTTP status code
	StatusCode *int32 `json:"status_code,omitempty"`
	// Unique request identifier for debugging
	RequestId *string `json:"request_id,omitempty"`
	// Additional context about the error
	Details map[string]*interface{} `json:"details,omitempty"`
	// Card decline code from the payment processor
	DeclineCode *string `json:"decline_code,omitempty"`
	// URL to documentation about this error
	DocUrl               *string `json:"doc_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiErrorError ApiErrorError

// NewApiErrorError instantiates a new ApiErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorError(type_ string, code string, message string) *ApiErrorError {
	this := ApiErrorError{}
	this.Type = type_
	this.Code = code
	this.Message = message
	return &this
}

// NewApiErrorErrorWithDefaults instantiates a new ApiErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorErrorWithDefaults() *ApiErrorError {
	this := ApiErrorError{}
	return &this
}

// GetType returns the Type field value
func (o *ApiErrorError) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiErrorError) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *ApiErrorError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ApiErrorError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ApiErrorError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiErrorError) SetMessage(v string) {
	o.Message = v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *ApiErrorError) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *ApiErrorError) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *ApiErrorError) SetParam(v string) {
	o.Param = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ApiErrorError) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ApiErrorError) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ApiErrorError) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ApiErrorError) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ApiErrorError) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ApiErrorError) SetRequestId(v string) {
	o.RequestId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ApiErrorError) GetDetails() map[string]*interface{} {
	if o == nil || IsNil(o.Details) {
		var ret map[string]*interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetDetailsOk() (map[string]*interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]*interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ApiErrorError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]*interface{} and assigns it to the Details field.
func (o *ApiErrorError) SetDetails(v map[string]*interface{}) {
	o.Details = v
}

// GetDeclineCode returns the DeclineCode field value if set, zero value otherwise.
func (o *ApiErrorError) GetDeclineCode() string {
	if o == nil || IsNil(o.DeclineCode) {
		var ret string
		return ret
	}
	return *o.DeclineCode
}

// GetDeclineCodeOk returns a tuple with the DeclineCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetDeclineCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DeclineCode) {
		return nil, false
	}
	return o.DeclineCode, true
}

// HasDeclineCode returns a boolean if a field has been set.
func (o *ApiErrorError) HasDeclineCode() bool {
	if o != nil && !IsNil(o.DeclineCode) {
		return true
	}

	return false
}

// SetDeclineCode gets a reference to the given string and assigns it to the DeclineCode field.
func (o *ApiErrorError) SetDeclineCode(v string) {
	o.DeclineCode = &v
}

// GetDocUrl returns the DocUrl field value if set, zero value otherwise.
func (o *ApiErrorError) GetDocUrl() string {
	if o == nil || IsNil(o.DocUrl) {
		var ret string
		return ret
	}
	return *o.DocUrl
}

// GetDocUrlOk returns a tuple with the DocUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorError) GetDocUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocUrl) {
		return nil, false
	}
	return o.DocUrl, true
}

// HasDocUrl returns a boolean if a field has been set.
func (o *ApiErrorError) HasDocUrl() bool {
	if o != nil && !IsNil(o.DocUrl) {
		return true
	}

	return false
}

// SetDocUrl gets a reference to the given string and assigns it to the DocUrl field.
func (o *ApiErrorError) SetDocUrl(v string) {
	o.DocUrl = &v
}

func (o ApiErrorError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.DeclineCode) {
		toSerialize["decline_code"] = o.DeclineCode
	}
	if !IsNil(o.DocUrl) {
		toSerialize["doc_url"] = o.DocUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiErrorError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"code",
		"message",
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

	varApiErrorError := _ApiErrorError{}

	err = json.Unmarshal(data, &varApiErrorError)

	if err != nil {
		return err
	}

	*o = ApiErrorError(varApiErrorError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "param")
		delete(additionalProperties, "status_code")
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "details")
		delete(additionalProperties, "decline_code")
		delete(additionalProperties, "doc_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiErrorError struct {
	value *ApiErrorError
	isSet bool
}

func (v NullableApiErrorError) Get() *ApiErrorError {
	return v.value
}

func (v *NullableApiErrorError) Set(val *ApiErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorError(val *ApiErrorError) *NullableApiErrorError {
	return &NullableApiErrorError{value: val, isSet: true}
}

func (v NullableApiErrorError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
