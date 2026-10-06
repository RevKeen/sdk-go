/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your merchant API key in the `x-api-key` header, or as `Authorization: Bearer <key>` — both are accepted. Create and manage keys in the [Dashboard](https://app.revkeen.com/dashboard/settings/general?tab=developer). Use `rk_sandbox_*` for staging/test and `rk_live_*` for production. Keys are scoped, rollable, IP-lockable and revocable; the raw secret is shown once, at creation.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Access tokens are opaque, prefixed `rkoa_`, and sent via `Authorization: Bearer rkoa_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Errors  Every non-2xx response uses one envelope: `{ \"error\": { \"type\", \"code\", \"message\", ... } }`. Branch on `error.type` and `error.code`; treat `error.message` as display text only. Error bodies never echo the credential you presented.  | Status | `error.type` | Notes | |--------|--------------|-------| | `400` | `invalid_request_error` | Schema validation failed. `details.fields` maps each rejected parameter to its reasons. | | `401` | `authentication_error` | Missing, invalid, expired or revoked credential. Always carries `WWW-Authenticate: Bearer realm=\"revkeen-api\"`. | | `402` | `card_error` | Declined by the issuer or processor; see `decline_code`. | | `403` | `authorization_error` | The caller is identified but the action is denied — scope, IP allowlist, origin, or merchant mismatch. | | `409` | `invalid_request_error` or `idempotency_error` | Resource-state conflict, or an `Idempotency-Key` replayed with a different body. | | `429` | `rate_limit_error` | See Rate Limits below. | | `500` | `api_error` | Retry with exponential backoff, quoting `error.request_id`. |  `401` codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`.  `403` codes: `insufficient_permissions`, `forbidden`, `ip_not_allowed`, `tenant_scope_mismatch`, `no_organization`, `no_merchant`, `merchant_mismatch`, `cart_disabled`, `origin_not_allowed`, `verification_required`, `read_only_impersonation_blocked`, `publishable_key_restricted`, `internal_only`.  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  Each API key (`rk_live_*`, `rk_sandbox_*`, `rk_pk_*`) may make **600 requests per minute**, in production and staging alike. The limit is per key, counted over a 60-second window, with no separate burst allowance. Some endpoints apply their own, lower limits.  An over-limit key is still valid: it receives `429`, never `401`. Wait for `Retry-After` and retry; do not re-authenticate or rotate the key.  A rate-limited request returns `429` with `type: rate_limit_error`, `code: rate_limit_exceeded` and `details.retry_after` in whole seconds.  | Header | Meaning | |--------|---------| | `Retry-After` | Whole seconds to wait before retrying. Present on every `429`. | | `X-RateLimit-Limit` | Requests permitted in the current window, when the limiter publishes one. | | `X-RateLimit-Remaining` | Requests left in the current window. | | `X-RateLimit-Reset` | UNIX timestamp in **seconds** at which the window resets. |

API version: 2026-05-01
Contact: support@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MandatesAPI interface {

	/*
			MandatesCancel Cancel a Direct Debit mandate

			Cancels a mandate. If it was lodged with the bureau, the London & Zurich customer is cancelled too.

		---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates` — List Direct Debit mandates
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Mandate ID
			@return ApiMandatesCancelRequest
	*/
	MandatesCancel(ctx context.Context, id string) ApiMandatesCancelRequest

	// MandatesCancelExecute executes the request
	//  @return MandateActionResponse
	MandatesCancelExecute(r ApiMandatesCancelRequest) (*MandateActionResponse, *http.Response, error)

	/*
			MandatesCreate Create a Direct Debit mandate

			Creates a Direct Debit mandate for a customer. Bank details are validated with the bureau and stored KMS-encrypted; only masked values are returned.

		---

		**Related endpoints**

		- `GET /mandates` — List Direct Debit mandates
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
		- `410 gone` — the resource has been archived or permanently deleted.
		- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiMandatesCreateRequest
	*/
	MandatesCreate(ctx context.Context) ApiMandatesCreateRequest

	// MandatesCreateExecute executes the request
	//  @return MandateResponse
	MandatesCreateExecute(r ApiMandatesCreateRequest) (*MandateResponse, *http.Response, error)

	/*
			MandatesGet Retrieve a Direct Debit mandate

			---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates` — List Direct Debit mandates
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Mandate ID
			@return ApiMandatesGetRequest
	*/
	MandatesGet(ctx context.Context, id string) ApiMandatesGetRequest

	// MandatesGetExecute executes the request
	//  @return MandateResponse
	MandatesGetExecute(r ApiMandatesGetRequest) (*MandateResponse, *http.Response, error)

	/*
			MandatesList List Direct Debit mandates

			---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiMandatesListRequest
	*/
	MandatesList(ctx context.Context) ApiMandatesListRequest

	// MandatesListExecute executes the request
	//  @return MandateListResponse
	MandatesListExecute(r ApiMandatesListRequest) (*MandateListResponse, *http.Response, error)

	/*
			MandatesReinstate Reinstate a suspended Direct Debit mandate

			Reinstates a suspended mandate. This is NOT a status flip — it re-lodges the mandate to Bacs (a new AUDDIS instruction) using the securely stored bank details, then moves the mandate to pending_lodgement (REV-3123).

		---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates` — List Direct Debit mandates
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
		- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Mandate ID
			@return ApiMandatesReinstateRequest
	*/
	MandatesReinstate(ctx context.Context, id string) ApiMandatesReinstateRequest

	// MandatesReinstateExecute executes the request
	//  @return MandateActionResponse
	MandatesReinstateExecute(r ApiMandatesReinstateRequest) (*MandateActionResponse, *http.Response, error)

	/*
			MandatesScheduleCollection Schedule a one-off Direct Debit collection

			Schedules a one-off Bacs collection against an active mandate for an invoice or payment link. L&Z confirms the permitted collection date, the payer receives the configured regulatory advance notice first, and nothing is marked paid until the bureau confirms collection. Idempotent per source: retrying with the same sourceId returns or reconciles the same provider schedule instead of collecting twice. An invoice source must be issued and still owe money: a draft or pending_approval invoice is refused, and so are paid, void and other closed invoices (409 invoice_not_collectable at the final check before bureau submission).

		---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates` — List Direct Debit mandates
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
		- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Mandate ID
			@return ApiMandatesScheduleCollectionRequest
	*/
	MandatesScheduleCollection(ctx context.Context, id string) ApiMandatesScheduleCollectionRequest

	// MandatesScheduleCollectionExecute executes the request
	//  @return ScheduleCollectionResponse
	MandatesScheduleCollectionExecute(r ApiMandatesScheduleCollectionRequest) (*ScheduleCollectionResponse, *http.Response, error)

	/*
			MandatesSuspend Suspend a Direct Debit mandate

			Suspends an active (or pending-lodgement) mandate.

		---

		**Related endpoints**

		- `POST /mandates` — Create a Direct Debit mandate
		- `GET /mandates` — List Direct Debit mandates
		- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
		- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
		- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
		- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Mandate ID
			@return ApiMandatesSuspendRequest
	*/
	MandatesSuspend(ctx context.Context, id string) ApiMandatesSuspendRequest

	// MandatesSuspendExecute executes the request
	//  @return MandateActionResponse
	MandatesSuspendExecute(r ApiMandatesSuspendRequest) (*MandateActionResponse, *http.Response, error)
}

// MandatesAPIService MandatesAPI service
type MandatesAPIService service

type ApiMandatesCancelRequest struct {
	ctx                  context.Context
	ApiService           MandatesAPI
	id                   string
	idempotencyKey       *string
	mandateActionRequest *MandateActionRequest
	runtimeOptions       []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiMandatesCancelRequest) IdempotencyKey(idempotencyKey string) ApiMandatesCancelRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMandatesCancelRequest) MandateActionRequest(mandateActionRequest MandateActionRequest) ApiMandatesCancelRequest {
	r.mandateActionRequest = &mandateActionRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesCancelRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesCancelRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesCancelRequest) Execute() (*MandateActionResponse, *http.Response, error) {
	var value *MandateActionResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesCancelExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesCancel Cancel a Direct Debit mandate

Cancels a mandate. If it was lodged with the bureau, the London & Zurich customer is cancelled too.

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates` — List Direct Debit mandates
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mandate ID
	@return ApiMandatesCancelRequest
*/
func (a *MandatesAPIService) MandatesCancel(ctx context.Context, id string) ApiMandatesCancelRequest {
	return ApiMandatesCancelRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return MandateActionResponse
func (a *MandatesAPIService) MandatesCancelExecute(r ApiMandatesCancelRequest) (*MandateActionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateActionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.mandateActionRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesCreateRequest struct {
	ctx                  context.Context
	ApiService           MandatesAPI
	idempotencyKey       *string
	createMandateRequest *CreateMandateRequest
	runtimeOptions       []RequestOption
}

// Required stable key for this logical mandate submission. Reuse the same key only with an identical request.
func (r ApiMandatesCreateRequest) IdempotencyKey(idempotencyKey string) ApiMandatesCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMandatesCreateRequest) CreateMandateRequest(createMandateRequest CreateMandateRequest) ApiMandatesCreateRequest {
	r.createMandateRequest = &createMandateRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesCreateRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesCreateRequest) Execute() (*MandateResponse, *http.Response, error) {
	var value *MandateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesCreate Create a Direct Debit mandate

Creates a Direct Debit mandate for a customer. Bank details are validated with the bureau and stored KMS-encrypted; only masked values are returned.

---

**Related endpoints**

- `GET /mandates` — List Direct Debit mandates
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
- `410 gone` — the resource has been archived or permanently deleted.
- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMandatesCreateRequest
*/
func (a *MandatesAPIService) MandatesCreate(ctx context.Context) ApiMandatesCreateRequest {
	return ApiMandatesCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MandateResponse
func (a *MandatesAPIService) MandatesCreateExecute(r ApiMandatesCreateRequest) (*MandateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.idempotencyKey == nil {
		return localVarReturnValue, nil, reportError("idempotencyKey is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	// body params
	localVarPostBody = r.createMandateRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesGetRequest struct {
	ctx            context.Context
	ApiService     MandatesAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesGetRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesGetRequest) Execute() (*MandateResponse, *http.Response, error) {
	var value *MandateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesGet Retrieve a Direct Debit mandate

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates` — List Direct Debit mandates
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mandate ID
	@return ApiMandatesGetRequest
*/
func (a *MandatesAPIService) MandatesGet(ctx context.Context, id string) ApiMandatesGetRequest {
	return ApiMandatesGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return MandateResponse
func (a *MandatesAPIService) MandatesGetExecute(r ApiMandatesGetRequest) (*MandateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesListRequest struct {
	ctx            context.Context
	ApiService     MandatesAPI
	customerId     *string
	status         *string
	runtimeOptions []RequestOption
}

// Filter mandates by customer
func (r ApiMandatesListRequest) CustomerId(customerId string) ApiMandatesListRequest {
	r.customerId = &customerId
	return r
}

// Filter mandates by status (e.g. active, suspended, cancelled)
func (r ApiMandatesListRequest) Status(status string) ApiMandatesListRequest {
	r.status = &status
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesListRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesListRequest) Execute() (*MandateListResponse, *http.Response, error) {
	var value *MandateListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesList List Direct Debit mandates

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMandatesListRequest
*/
func (a *MandatesAPIService) MandatesList(ctx context.Context) ApiMandatesListRequest {
	return ApiMandatesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MandateListResponse
func (a *MandatesAPIService) MandatesListExecute(r ApiMandatesListRequest) (*MandateListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerId", r.customerId, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesReinstateRequest struct {
	ctx                  context.Context
	ApiService           MandatesAPI
	id                   string
	idempotencyKey       *string
	mandateActionRequest *MandateActionRequest
	runtimeOptions       []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiMandatesReinstateRequest) IdempotencyKey(idempotencyKey string) ApiMandatesReinstateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMandatesReinstateRequest) MandateActionRequest(mandateActionRequest MandateActionRequest) ApiMandatesReinstateRequest {
	r.mandateActionRequest = &mandateActionRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesReinstateRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesReinstateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesReinstateRequest) Execute() (*MandateActionResponse, *http.Response, error) {
	var value *MandateActionResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesReinstateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesReinstate Reinstate a suspended Direct Debit mandate

Reinstates a suspended mandate. This is NOT a status flip — it re-lodges the mandate to Bacs (a new AUDDIS instruction) using the securely stored bank details, then moves the mandate to pending_lodgement (REV-3123).

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates` — List Direct Debit mandates
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mandate ID
	@return ApiMandatesReinstateRequest
*/
func (a *MandatesAPIService) MandatesReinstate(ctx context.Context, id string) ApiMandatesReinstateRequest {
	return ApiMandatesReinstateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return MandateActionResponse
func (a *MandatesAPIService) MandatesReinstateExecute(r ApiMandatesReinstateRequest) (*MandateActionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateActionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesReinstate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates/{id}/reinstate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.mandateActionRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesScheduleCollectionRequest struct {
	ctx                       context.Context
	ApiService                MandatesAPI
	id                        string
	idempotencyKey            *string
	scheduleCollectionRequest *ScheduleCollectionRequest
	runtimeOptions            []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiMandatesScheduleCollectionRequest) IdempotencyKey(idempotencyKey string) ApiMandatesScheduleCollectionRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMandatesScheduleCollectionRequest) ScheduleCollectionRequest(scheduleCollectionRequest ScheduleCollectionRequest) ApiMandatesScheduleCollectionRequest {
	r.scheduleCollectionRequest = &scheduleCollectionRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesScheduleCollectionRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesScheduleCollectionRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesScheduleCollectionRequest) Execute() (*ScheduleCollectionResponse, *http.Response, error) {
	var value *ScheduleCollectionResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesScheduleCollectionExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesScheduleCollection Schedule a one-off Direct Debit collection

Schedules a one-off Bacs collection against an active mandate for an invoice or payment link. L&Z confirms the permitted collection date, the payer receives the configured regulatory advance notice first, and nothing is marked paid until the bureau confirms collection. Idempotent per source: retrying with the same sourceId returns or reconciles the same provider schedule instead of collecting twice. An invoice source must be issued and still owe money: a draft or pending_approval invoice is refused, and so are paid, void and other closed invoices (409 invoice_not_collectable at the final check before bureau submission).

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates` — List Direct Debit mandates
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/suspend` — Suspend a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mandate ID
	@return ApiMandatesScheduleCollectionRequest
*/
func (a *MandatesAPIService) MandatesScheduleCollection(ctx context.Context, id string) ApiMandatesScheduleCollectionRequest {
	return ApiMandatesScheduleCollectionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ScheduleCollectionResponse
func (a *MandatesAPIService) MandatesScheduleCollectionExecute(r ApiMandatesScheduleCollectionRequest) (*ScheduleCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ScheduleCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesScheduleCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates/{id}/collections"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.scheduleCollectionRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ForbiddenError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMandatesSuspendRequest struct {
	ctx                  context.Context
	ApiService           MandatesAPI
	id                   string
	idempotencyKey       *string
	mandateActionRequest *MandateActionRequest
	runtimeOptions       []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiMandatesSuspendRequest) IdempotencyKey(idempotencyKey string) ApiMandatesSuspendRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiMandatesSuspendRequest) MandateActionRequest(mandateActionRequest MandateActionRequest) ApiMandatesSuspendRequest {
	r.mandateActionRequest = &mandateActionRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiMandatesSuspendRequest) WithRequestOptions(opts ...RequestOption) ApiMandatesSuspendRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiMandatesSuspendRequest) Execute() (*MandateActionResponse, *http.Response, error) {
	var value *MandateActionResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.MandatesSuspendExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
MandatesSuspend Suspend a Direct Debit mandate

Suspends an active (or pending-lodgement) mandate.

---

**Related endpoints**

- `POST /mandates` — Create a Direct Debit mandate
- `GET /mandates` — List Direct Debit mandates
- `GET /mandates/{id}` — Retrieve a Direct Debit mandate
- `POST /mandates/{id}/cancel` — Cancel a Direct Debit mandate
- `POST /mandates/{id}/reinstate` — Reinstate a suspended Direct Debit mandate
- `POST /mandates/{id}/collections` — Schedule a one-off Direct Debit collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mandate ID
	@return ApiMandatesSuspendRequest
*/
func (a *MandatesAPIService) MandatesSuspend(ctx context.Context, id string) ApiMandatesSuspendRequest {
	return ApiMandatesSuspendRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return MandateActionResponse
func (a *MandatesAPIService) MandatesSuspendExecute(r ApiMandatesSuspendRequest) (*MandateActionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MandateActionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MandatesAPIService.MandatesSuspend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandates/{id}/suspend"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.mandateActionRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
