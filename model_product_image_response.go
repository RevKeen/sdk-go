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

// checks if the ProductImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductImageResponse{}

// ProductImageResponse One extra product image, in display order.
type ProductImageResponse struct {
	// Storage key or https URL
	Key string `json:"key"`
	// Media library asset ID
	AssetId NullableString `json:"asset_id"`
	// Alt text
	Alt NullableString `json:"alt"`
	// Where the image came from
	Source string `json:"source"`
	// Photographer credit
	Photographer NullableString `json:"photographer"`
	// Link for the credit
	AttributionUrl       NullableString `json:"attribution_url"`
	AdditionalProperties map[string]interface{}
}

type _ProductImageResponse ProductImageResponse

// NewProductImageResponse instantiates a new ProductImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductImageResponse(key string, assetId NullableString, alt NullableString, source string, photographer NullableString, attributionUrl NullableString) *ProductImageResponse {
	this := ProductImageResponse{}
	this.Key = key
	this.AssetId = assetId
	this.Alt = alt
	this.Source = source
	this.Photographer = photographer
	this.AttributionUrl = attributionUrl
	return &this
}

// NewProductImageResponseWithDefaults instantiates a new ProductImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductImageResponseWithDefaults() *ProductImageResponse {
	this := ProductImageResponse{}
	return &this
}

// GetKey returns the Key field value
func (o *ProductImageResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ProductImageResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ProductImageResponse) SetKey(v string) {
	o.Key = v
}

// GetAssetId returns the AssetId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductImageResponse) GetAssetId() string {
	if o == nil || o.AssetId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AssetId.Get()
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductImageResponse) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetId.Get(), o.AssetId.IsSet()
}

// SetAssetId sets field value
func (o *ProductImageResponse) SetAssetId(v string) {
	o.AssetId.Set(&v)
}

// GetAlt returns the Alt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductImageResponse) GetAlt() string {
	if o == nil || o.Alt.Get() == nil {
		var ret string
		return ret
	}

	return *o.Alt.Get()
}

// GetAltOk returns a tuple with the Alt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductImageResponse) GetAltOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alt.Get(), o.Alt.IsSet()
}

// SetAlt sets field value
func (o *ProductImageResponse) SetAlt(v string) {
	o.Alt.Set(&v)
}

// GetSource returns the Source field value
func (o *ProductImageResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ProductImageResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ProductImageResponse) SetSource(v string) {
	o.Source = v
}

// GetPhotographer returns the Photographer field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductImageResponse) GetPhotographer() string {
	if o == nil || o.Photographer.Get() == nil {
		var ret string
		return ret
	}

	return *o.Photographer.Get()
}

// GetPhotographerOk returns a tuple with the Photographer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductImageResponse) GetPhotographerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Photographer.Get(), o.Photographer.IsSet()
}

// SetPhotographer sets field value
func (o *ProductImageResponse) SetPhotographer(v string) {
	o.Photographer.Set(&v)
}

// GetAttributionUrl returns the AttributionUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProductImageResponse) GetAttributionUrl() string {
	if o == nil || o.AttributionUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.AttributionUrl.Get()
}

// GetAttributionUrlOk returns a tuple with the AttributionUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductImageResponse) GetAttributionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributionUrl.Get(), o.AttributionUrl.IsSet()
}

// SetAttributionUrl sets field value
func (o *ProductImageResponse) SetAttributionUrl(v string) {
	o.AttributionUrl.Set(&v)
}

func (o ProductImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["asset_id"] = o.AssetId.Get()
	toSerialize["alt"] = o.Alt.Get()
	toSerialize["source"] = o.Source
	toSerialize["photographer"] = o.Photographer.Get()
	toSerialize["attribution_url"] = o.AttributionUrl.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductImageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"asset_id",
		"alt",
		"source",
		"photographer",
		"attribution_url",
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

	varProductImageResponse := _ProductImageResponse{}

	err = json.Unmarshal(data, &varProductImageResponse)

	if err != nil {
		return err
	}

	*o = ProductImageResponse(varProductImageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "asset_id")
		delete(additionalProperties, "alt")
		delete(additionalProperties, "source")
		delete(additionalProperties, "photographer")
		delete(additionalProperties, "attribution_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProductImageResponse struct {
	value *ProductImageResponse
	isSet bool
}

func (v NullableProductImageResponse) Get() *ProductImageResponse {
	return v.value
}

func (v *NullableProductImageResponse) Set(val *ProductImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductImageResponse(val *ProductImageResponse) *NullableProductImageResponse {
	return &NullableProductImageResponse{value: val, isSet: true}
}

func (v NullableProductImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
