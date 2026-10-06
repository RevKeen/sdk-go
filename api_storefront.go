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

type StorefrontAPI interface {

	/*
			StorefrontOriginsCreate Register a storefront origin

			Register an exact web origin (scheme://host[:port]) for publishable-key storefront calls. Wildcards are rejected; http is allowed only for localhost.

		---

		**Related endpoints**

		- `GET /storefront/products` — List storefront products
		- `GET /storefront/products/{productId}` — Get a storefront product
		- `GET /storefront/origins` — List storefront origins
		- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
		- `GET /storefront/status` — Get storefront integration status

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiStorefrontOriginsCreateRequest
	*/
	StorefrontOriginsCreate(ctx context.Context) ApiStorefrontOriginsCreateRequest

	// StorefrontOriginsCreateExecute executes the request
	//  @return StorefrontOriginCreateResponse
	StorefrontOriginsCreateExecute(r ApiStorefrontOriginsCreateRequest) (*StorefrontOriginCreateResponse, *http.Response, error)

	/*
			StorefrontOriginsDelete Remove a storefront origin

			Remove a registered storefront origin. Browser calls from it fail closed afterwards.

		---

		**Related endpoints**

		- `GET /storefront/products` — List storefront products
		- `GET /storefront/products/{productId}` — Get a storefront product
		- `GET /storefront/origins` — List storefront origins
		- `POST /storefront/origins` — Register a storefront origin
		- `GET /storefront/status` — Get storefront integration status

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param originId Storefront origin id.
			@return ApiStorefrontOriginsDeleteRequest
	*/
	StorefrontOriginsDelete(ctx context.Context, originId string) ApiStorefrontOriginsDeleteRequest

	// StorefrontOriginsDeleteExecute executes the request
	//  @return StorefrontOriginDeleteResponse
	StorefrontOriginsDeleteExecute(r ApiStorefrontOriginsDeleteRequest) (*StorefrontOriginDeleteResponse, *http.Response, error)

	/*
			StorefrontOriginsList List storefront origins

			List the browser origins registered for publishable-key storefront calls.

		---

		**Related endpoints**

		- `GET /storefront/products` — List storefront products
		- `GET /storefront/products/{productId}` — Get a storefront product
		- `POST /storefront/origins` — Register a storefront origin
		- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
		- `GET /storefront/status` — Get storefront integration status

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiStorefrontOriginsListRequest
	*/
	StorefrontOriginsList(ctx context.Context) ApiStorefrontOriginsListRequest

	// StorefrontOriginsListExecute executes the request
	//  @return StorefrontOriginListResponse
	StorefrontOriginsListExecute(r ApiStorefrontOriginsListRequest) (*StorefrontOriginListResponse, *http.Response, error)

	/*
			StorefrontProductsGet Get a storefront product

			Fetch one active, cart-eligible product by UUID, URL handle (slug), or product reference. A UUID also finds a product hidden from the Storefront; a handle or reference only finds listed products, and the handle wins over another product's reference. Same browser-safe projection as the list endpoint.

		---

		**Related endpoints**

		- `GET /storefront/products` — List storefront products
		- `GET /storefront/origins` — List storefront origins
		- `POST /storefront/origins` — Register a storefront origin
		- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
		- `GET /storefront/status` — Get storefront integration status

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param productId Product UUID, URL handle (slug), or merchant product reference.
			@return ApiStorefrontProductsGetRequest
	*/
	StorefrontProductsGet(ctx context.Context, productId string) ApiStorefrontProductsGetRequest

	// StorefrontProductsGetExecute executes the request
	//  @return StorefrontProductResponse
	StorefrontProductsGetExecute(r ApiStorefrontProductsGetRequest) (*StorefrontProductResponse, *http.Response, error)

	/*
			StorefrontProductsList List storefront products

			List active, cart-eligible products with browser-safe display data, active prices, and derived availability. Only products listed on the Storefront are returned (store_visibility visible, or scheduled and due). Intended for publishable-key storefront and CMS-picker use; never returns product metadata or internal fields.

		---

		**Related endpoints**

		- `GET /storefront/products/{productId}` — Get a storefront product
		- `GET /storefront/origins` — List storefront origins
		- `POST /storefront/origins` — Register a storefront origin
		- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
		- `GET /storefront/status` — Get storefront integration status

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiStorefrontProductsListRequest
	*/
	StorefrontProductsList(ctx context.Context) ApiStorefrontProductsListRequest

	// StorefrontProductsListExecute executes the request
	//  @return StorefrontProductListResponse
	StorefrontProductsListExecute(r ApiStorefrontProductsListRequest) (*StorefrontProductListResponse, *http.Response, error)

	/*
			StorefrontStatusGet Get storefront integration status

			Readiness report for the Cart/headless-storefront setup: activation, managed keys, registered origins, product-read readiness, webhook health, and availability tracking. CART_DISABLED / KEYS_MISSING / ORIGIN_MISSING / PRODUCT_READ_UNAVAILABLE are fail-level; WEBHOOK_MISSING / WEBHOOK_UNREACHABLE are warn-level. Never returns key material.

		---

		**Related endpoints**

		- `GET /storefront/products` — List storefront products
		- `GET /storefront/products/{productId}` — Get a storefront product
		- `GET /storefront/origins` — List storefront origins
		- `POST /storefront/origins` — Register a storefront origin
		- `DELETE /storefront/origins/{originId}` — Remove a storefront origin

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiStorefrontStatusGetRequest
	*/
	StorefrontStatusGet(ctx context.Context) ApiStorefrontStatusGetRequest

	// StorefrontStatusGetExecute executes the request
	//  @return StorefrontStatusResponse
	StorefrontStatusGetExecute(r ApiStorefrontStatusGetRequest) (*StorefrontStatusResponse, *http.Response, error)
}

// StorefrontAPIService StorefrontAPI service
type StorefrontAPIService service

type ApiStorefrontOriginsCreateRequest struct {
	ctx                           context.Context
	ApiService                    StorefrontAPI
	storefrontOriginCreateRequest *StorefrontOriginCreateRequest
	runtimeOptions                []RequestOption
}

func (r ApiStorefrontOriginsCreateRequest) StorefrontOriginCreateRequest(storefrontOriginCreateRequest StorefrontOriginCreateRequest) ApiStorefrontOriginsCreateRequest {
	r.storefrontOriginCreateRequest = &storefrontOriginCreateRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontOriginsCreateRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontOriginsCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontOriginsCreateRequest) Execute() (*StorefrontOriginCreateResponse, *http.Response, error) {
	var value *StorefrontOriginCreateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontOriginsCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontOriginsCreate Register a storefront origin

Register an exact web origin (scheme://host[:port]) for publishable-key storefront calls. Wildcards are rejected; http is allowed only for localhost.

---

**Related endpoints**

- `GET /storefront/products` — List storefront products
- `GET /storefront/products/{productId}` — Get a storefront product
- `GET /storefront/origins` — List storefront origins
- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
- `GET /storefront/status` — Get storefront integration status

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorefrontOriginsCreateRequest
*/
func (a *StorefrontAPIService) StorefrontOriginsCreate(ctx context.Context) ApiStorefrontOriginsCreateRequest {
	return ApiStorefrontOriginsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StorefrontOriginCreateResponse
func (a *StorefrontAPIService) StorefrontOriginsCreateExecute(r ApiStorefrontOriginsCreateRequest) (*StorefrontOriginCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontOriginCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontOriginsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/origins"

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
	// body params
	localVarPostBody = r.storefrontOriginCreateRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v StorefrontOriginError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v StorefrontOriginError
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
			var v StorefrontOriginError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v StorefrontOriginError
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

type ApiStorefrontOriginsDeleteRequest struct {
	ctx            context.Context
	ApiService     StorefrontAPI
	originId       string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontOriginsDeleteRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontOriginsDeleteRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontOriginsDeleteRequest) Execute() (*StorefrontOriginDeleteResponse, *http.Response, error) {
	var value *StorefrontOriginDeleteResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontOriginsDeleteExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontOriginsDelete Remove a storefront origin

Remove a registered storefront origin. Browser calls from it fail closed afterwards.

---

**Related endpoints**

- `GET /storefront/products` — List storefront products
- `GET /storefront/products/{productId}` — Get a storefront product
- `GET /storefront/origins` — List storefront origins
- `POST /storefront/origins` — Register a storefront origin
- `GET /storefront/status` — Get storefront integration status

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param originId Storefront origin id.
	@return ApiStorefrontOriginsDeleteRequest
*/
func (a *StorefrontAPIService) StorefrontOriginsDelete(ctx context.Context, originId string) ApiStorefrontOriginsDeleteRequest {
	return ApiStorefrontOriginsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		originId:   originId,
	}
}

// Execute executes the request
//
//	@return StorefrontOriginDeleteResponse
func (a *StorefrontAPIService) StorefrontOriginsDeleteExecute(r ApiStorefrontOriginsDeleteRequest) (*StorefrontOriginDeleteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontOriginDeleteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontOriginsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/origins/{originId}"
	localVarPath = strings.Replace(localVarPath, "{"+"originId"+"}", url.PathEscape(parameterValueToString(r.originId, "originId")), -1)

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
			var v StorefrontOriginError
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
			var v StorefrontOriginError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v StorefrontOriginError
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

type ApiStorefrontOriginsListRequest struct {
	ctx            context.Context
	ApiService     StorefrontAPI
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontOriginsListRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontOriginsListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontOriginsListRequest) Execute() (*StorefrontOriginListResponse, *http.Response, error) {
	var value *StorefrontOriginListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontOriginsListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontOriginsList List storefront origins

List the browser origins registered for publishable-key storefront calls.

---

**Related endpoints**

- `GET /storefront/products` — List storefront products
- `GET /storefront/products/{productId}` — Get a storefront product
- `POST /storefront/origins` — Register a storefront origin
- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
- `GET /storefront/status` — Get storefront integration status

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorefrontOriginsListRequest
*/
func (a *StorefrontAPIService) StorefrontOriginsList(ctx context.Context) ApiStorefrontOriginsListRequest {
	return ApiStorefrontOriginsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StorefrontOriginListResponse
func (a *StorefrontAPIService) StorefrontOriginsListExecute(r ApiStorefrontOriginsListRequest) (*StorefrontOriginListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontOriginListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontOriginsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/origins"

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
			var v StorefrontOriginError
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
			var v StorefrontOriginError
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

type ApiStorefrontProductsGetRequest struct {
	ctx            context.Context
	ApiService     StorefrontAPI
	productId      string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontProductsGetRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontProductsGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontProductsGetRequest) Execute() (*StorefrontProductResponse, *http.Response, error) {
	var value *StorefrontProductResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontProductsGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontProductsGet Get a storefront product

Fetch one active, cart-eligible product by UUID, URL handle (slug), or product reference. A UUID also finds a product hidden from the Storefront; a handle or reference only finds listed products, and the handle wins over another product's reference. Same browser-safe projection as the list endpoint.

---

**Related endpoints**

- `GET /storefront/products` — List storefront products
- `GET /storefront/origins` — List storefront origins
- `POST /storefront/origins` — Register a storefront origin
- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
- `GET /storefront/status` — Get storefront integration status

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param productId Product UUID, URL handle (slug), or merchant product reference.
	@return ApiStorefrontProductsGetRequest
*/
func (a *StorefrontAPIService) StorefrontProductsGet(ctx context.Context, productId string) ApiStorefrontProductsGetRequest {
	return ApiStorefrontProductsGetRequest{
		ApiService: a,
		ctx:        ctx,
		productId:  productId,
	}
}

// Execute executes the request
//
//	@return StorefrontProductResponse
func (a *StorefrontAPIService) StorefrontProductsGetExecute(r ApiStorefrontProductsGetRequest) (*StorefrontProductResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontProductResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontProductsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/products/{productId}"
	localVarPath = strings.Replace(localVarPath, "{"+"productId"+"}", url.PathEscape(parameterValueToString(r.productId, "productId")), -1)

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
			var v StorefrontProductError
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
			var v StorefrontProductError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v StorefrontProductError
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

type ApiStorefrontProductsListRequest struct {
	ctx            context.Context
	ApiService     StorefrontAPI
	limit          *int32
	offset         *int32
	collection     *string
	tag            *string
	search         *string
	sort           *string
	runtimeOptions []RequestOption
}

// Maximum products to return (default 50, max 100).
func (r ApiStorefrontProductsListRequest) Limit(limit int32) ApiStorefrontProductsListRequest {
	r.limit = &limit
	return r
}

// Products to skip, for pagination.
func (r ApiStorefrontProductsListRequest) Offset(offset int32) ApiStorefrontProductsListRequest {
	r.offset = &offset
	return r
}

// Only products in this product collection (id or URL handle).
func (r ApiStorefrontProductsListRequest) Collection(collection string) ApiStorefrontProductsListRequest {
	r.collection = &collection
	return r
}

// Only products carrying this tag.
func (r ApiStorefrontProductsListRequest) Tag(tag string) ApiStorefrontProductsListRequest {
	r.tag = &tag
	return r
}

// Case-insensitive match on product name or description.
func (r ApiStorefrontProductsListRequest) Search(search string) ApiStorefrontProductsListRequest {
	r.search = &search
	return r
}

// featured (collection order, else newest), newest, name, price_asc or price_desc.
func (r ApiStorefrontProductsListRequest) Sort(sort string) ApiStorefrontProductsListRequest {
	r.sort = &sort
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontProductsListRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontProductsListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontProductsListRequest) Execute() (*StorefrontProductListResponse, *http.Response, error) {
	var value *StorefrontProductListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontProductsListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontProductsList List storefront products

List active, cart-eligible products with browser-safe display data, active prices, and derived availability. Only products listed on the Storefront are returned (store_visibility visible, or scheduled and due). Intended for publishable-key storefront and CMS-picker use; never returns product metadata or internal fields.

---

**Related endpoints**

- `GET /storefront/products/{productId}` — Get a storefront product
- `GET /storefront/origins` — List storefront origins
- `POST /storefront/origins` — Register a storefront origin
- `DELETE /storefront/origins/{originId}` — Remove a storefront origin
- `GET /storefront/status` — Get storefront integration status

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorefrontProductsListRequest
*/
func (a *StorefrontAPIService) StorefrontProductsList(ctx context.Context) ApiStorefrontProductsListRequest {
	return ApiStorefrontProductsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StorefrontProductListResponse
func (a *StorefrontAPIService) StorefrontProductsListExecute(r ApiStorefrontProductsListRequest) (*StorefrontProductListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontProductListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontProductsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.collection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "collection", r.collection, "form", "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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
			var v StorefrontProductError
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
			var v StorefrontProductError
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

type ApiStorefrontStatusGetRequest struct {
	ctx            context.Context
	ApiService     StorefrontAPI
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiStorefrontStatusGetRequest) WithRequestOptions(opts ...RequestOption) ApiStorefrontStatusGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiStorefrontStatusGetRequest) Execute() (*StorefrontStatusResponse, *http.Response, error) {
	var value *StorefrontStatusResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.StorefrontStatusGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
StorefrontStatusGet Get storefront integration status

Readiness report for the Cart/headless-storefront setup: activation, managed keys, registered origins, product-read readiness, webhook health, and availability tracking. CART_DISABLED / KEYS_MISSING / ORIGIN_MISSING / PRODUCT_READ_UNAVAILABLE are fail-level; WEBHOOK_MISSING / WEBHOOK_UNREACHABLE are warn-level. Never returns key material.

---

**Related endpoints**

- `GET /storefront/products` — List storefront products
- `GET /storefront/products/{productId}` — Get a storefront product
- `GET /storefront/origins` — List storefront origins
- `POST /storefront/origins` — Register a storefront origin
- `DELETE /storefront/origins/{originId}` — Remove a storefront origin

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStorefrontStatusGetRequest
*/
func (a *StorefrontAPIService) StorefrontStatusGet(ctx context.Context) ApiStorefrontStatusGetRequest {
	return ApiStorefrontStatusGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StorefrontStatusResponse
func (a *StorefrontAPIService) StorefrontStatusGetExecute(r ApiStorefrontStatusGetRequest) (*StorefrontStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StorefrontStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorefrontAPIService.StorefrontStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/storefront/status"

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
			var v StorefrontStatusError
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
			var v StorefrontStatusError
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
