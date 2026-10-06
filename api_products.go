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

type ProductsAPI interface {

	/*
			ProductCollectionsCreate Create a product collection

			Create a product collection. The URL handle must be unique for the merchant, and every product must belong to the merchant.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
		- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiProductCollectionsCreateRequest
	*/
	ProductCollectionsCreate(ctx context.Context) ApiProductCollectionsCreateRequest

	// ProductCollectionsCreateExecute executes the request
	//  @return ProductCollectionCreateResponse
	ProductCollectionsCreateExecute(r ApiProductCollectionsCreateRequest) (*ProductCollectionCreateResponse, *http.Response, error)

	/*
			ProductCollectionsDelete Delete a product collection

			Delete a product collection. Its products are not changed.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection

		**Common errors**

		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiProductCollectionsDeleteRequest
	*/
	ProductCollectionsDelete(ctx context.Context, id string) ApiProductCollectionsDeleteRequest

	// ProductCollectionsDeleteExecute executes the request
	//  @return ProductCollectionDeletedResponse
	ProductCollectionsDeleteExecute(r ApiProductCollectionsDeleteRequest) (*ProductCollectionDeletedResponse, *http.Response, error)

	/*
			ProductCollectionsGet Get a product collection

			Get one product collection by id or URL handle.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection

		**Common errors**

		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Collection id or URL handle.
			@return ApiProductCollectionsGetRequest
	*/
	ProductCollectionsGet(ctx context.Context, id string) ApiProductCollectionsGetRequest

	// ProductCollectionsGetExecute executes the request
	//  @return ProductCollectionResponse
	ProductCollectionsGetExecute(r ApiProductCollectionsGetRequest) (*ProductCollectionResponse, *http.Response, error)

	/*
			ProductCollectionsList List product collections

			List the merchant's product collections, with member product ids in display order.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `POST /product-collections` — Create a product collection
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiProductCollectionsListRequest
	*/
	ProductCollectionsList(ctx context.Context) ApiProductCollectionsListRequest

	// ProductCollectionsListExecute executes the request
	//  @return ProductCollectionListResponse
	ProductCollectionsListExecute(r ApiProductCollectionsListRequest) (*ProductCollectionListResponse, *http.Response, error)

	/*
			ProductCollectionsUpdate Update a product collection

			Update a product collection. `product_ids` replaces the membership in the given order. Send `expected_version` to refuse the change if the collection was edited since you read it.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection

		**Common errors**

		- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
		- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiProductCollectionsUpdateRequest
	*/
	ProductCollectionsUpdate(ctx context.Context, id string) ApiProductCollectionsUpdateRequest

	// ProductCollectionsUpdateExecute executes the request
	//  @return ProductCollectionUpdateResponse
	ProductCollectionsUpdateExecute(r ApiProductCollectionsUpdateRequest) (*ProductCollectionUpdateResponse, *http.Response, error)

	/*
			ProductsCreate Create product

			Create a new product. Requires 'products:write' scope.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiProductsCreateRequest
	*/
	ProductsCreate(ctx context.Context) ApiProductsCreateRequest

	// ProductsCreateExecute executes the request
	//  @return ProductCreateResponse
	ProductsCreateExecute(r ApiProductsCreateRequest) (*ProductCreateResponse, *http.Response, error)

	/*
			ProductsGet Get product by ID

			Get a product by its ID. Requires 'products:read' scope.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiProductsGetRequest
	*/
	ProductsGet(ctx context.Context, id string) ApiProductsGetRequest

	// ProductsGetExecute executes the request
	//  @return ProductRetrieveResponse
	ProductsGetExecute(r ApiProductsGetRequest) (*ProductRetrieveResponse, *http.Response, error)

	/*
			ProductsList List products

			List products with pagination and optional search.

		---

		**Related endpoints**

		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PATCH /products/{id}` — Update product
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiProductsListRequest
	*/
	ProductsList(ctx context.Context) ApiProductsListRequest

	// ProductsListExecute executes the request
	//  @return ProductListResponse
	ProductsListExecute(r ApiProductsListRequest) (*ProductListResponse, *http.Response, error)

	/*
			ProductsUpdate Update product

			Update an existing product. Requires 'products:write' scope.

		---

		**Related endpoints**

		- `GET /products` — List products
		- `POST /products` — Create product
		- `GET /products/{id}` — Get product by ID
		- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
		- `PUT /products/external/batch` — Batch upsert products by external ID
		- `GET /product-collections` — List product collections
		- `POST /product-collections` — Create a product collection
		- `GET /product-collections/{id}` — Get a product collection

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiProductsUpdateRequest
	*/
	ProductsUpdate(ctx context.Context, id string) ApiProductsUpdateRequest

	// ProductsUpdateExecute executes the request
	//  @return ProductUpdateResponse
	ProductsUpdateExecute(r ApiProductsUpdateRequest) (*ProductUpdateResponse, *http.Response, error)
}

// ProductsAPIService ProductsAPI service
type ProductsAPIService service

type ApiProductCollectionsCreateRequest struct {
	ctx                            context.Context
	ApiService                     ProductsAPI
	productCollectionCreateRequest *ProductCollectionCreateRequest
	runtimeOptions                 []RequestOption
}

func (r ApiProductCollectionsCreateRequest) ProductCollectionCreateRequest(productCollectionCreateRequest ProductCollectionCreateRequest) ApiProductCollectionsCreateRequest {
	r.productCollectionCreateRequest = &productCollectionCreateRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductCollectionsCreateRequest) WithRequestOptions(opts ...RequestOption) ApiProductCollectionsCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductCollectionsCreateRequest) Execute() (*ProductCollectionCreateResponse, *http.Response, error) {
	var value *ProductCollectionCreateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductCollectionsCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductCollectionsCreate Create a product collection

Create a product collection. The URL handle must be unique for the merchant, and every product must belong to the merchant.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductCollectionsCreateRequest
*/
func (a *ProductsAPIService) ProductCollectionsCreate(ctx context.Context) ApiProductCollectionsCreateRequest {
	return ApiProductCollectionsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductCollectionCreateResponse
func (a *ProductsAPIService) ProductCollectionsCreateExecute(r ApiProductCollectionsCreateRequest) (*ProductCollectionCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCollectionCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductCollectionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product-collections"

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
	localVarPostBody = r.productCollectionCreateRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProductCollectionError
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
			var v ProductCollectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ProductCollectionError
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

type ApiProductCollectionsDeleteRequest struct {
	ctx            context.Context
	ApiService     ProductsAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductCollectionsDeleteRequest) WithRequestOptions(opts ...RequestOption) ApiProductCollectionsDeleteRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductCollectionsDeleteRequest) Execute() (*ProductCollectionDeletedResponse, *http.Response, error) {
	var value *ProductCollectionDeletedResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductCollectionsDeleteExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductCollectionsDelete Delete a product collection

Delete a product collection. Its products are not changed.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection

**Common errors**

- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiProductCollectionsDeleteRequest
*/
func (a *ProductsAPIService) ProductCollectionsDelete(ctx context.Context, id string) ApiProductCollectionsDeleteRequest {
	return ApiProductCollectionsDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProductCollectionDeletedResponse
func (a *ProductsAPIService) ProductCollectionsDeleteExecute(r ApiProductCollectionsDeleteRequest) (*ProductCollectionDeletedResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCollectionDeletedResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductCollectionsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product-collections/{id}"
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProductCollectionError
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
			var v ProductCollectionError
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

type ApiProductCollectionsGetRequest struct {
	ctx            context.Context
	ApiService     ProductsAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductCollectionsGetRequest) WithRequestOptions(opts ...RequestOption) ApiProductCollectionsGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductCollectionsGetRequest) Execute() (*ProductCollectionResponse, *http.Response, error) {
	var value *ProductCollectionResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductCollectionsGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductCollectionsGet Get a product collection

Get one product collection by id or URL handle.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection

**Common errors**

- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Collection id or URL handle.
	@return ApiProductCollectionsGetRequest
*/
func (a *ProductsAPIService) ProductCollectionsGet(ctx context.Context, id string) ApiProductCollectionsGetRequest {
	return ApiProductCollectionsGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProductCollectionResponse
func (a *ProductsAPIService) ProductCollectionsGetExecute(r ApiProductCollectionsGetRequest) (*ProductCollectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCollectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductCollectionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product-collections/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
	}
	if strlen(r.id) > 80 {
		return localVarReturnValue, nil, reportError("id must have less than 80 elements")
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProductCollectionError
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
			var v ProductCollectionError
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

type ApiProductCollectionsListRequest struct {
	ctx            context.Context
	ApiService     ProductsAPI
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductCollectionsListRequest) WithRequestOptions(opts ...RequestOption) ApiProductCollectionsListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductCollectionsListRequest) Execute() (*ProductCollectionListResponse, *http.Response, error) {
	var value *ProductCollectionListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductCollectionsListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductCollectionsList List product collections

List the merchant's product collections, with member product ids in display order.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `POST /product-collections` — Create a product collection
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductCollectionsListRequest
*/
func (a *ProductsAPIService) ProductCollectionsList(ctx context.Context) ApiProductCollectionsListRequest {
	return ApiProductCollectionsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductCollectionListResponse
func (a *ProductsAPIService) ProductCollectionsListExecute(r ApiProductCollectionsListRequest) (*ProductCollectionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCollectionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductCollectionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product-collections"

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProductCollectionError
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

type ApiProductCollectionsUpdateRequest struct {
	ctx                            context.Context
	ApiService                     ProductsAPI
	id                             string
	productCollectionUpdateRequest *ProductCollectionUpdateRequest
	runtimeOptions                 []RequestOption
}

func (r ApiProductCollectionsUpdateRequest) ProductCollectionUpdateRequest(productCollectionUpdateRequest ProductCollectionUpdateRequest) ApiProductCollectionsUpdateRequest {
	r.productCollectionUpdateRequest = &productCollectionUpdateRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductCollectionsUpdateRequest) WithRequestOptions(opts ...RequestOption) ApiProductCollectionsUpdateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductCollectionsUpdateRequest) Execute() (*ProductCollectionUpdateResponse, *http.Response, error) {
	var value *ProductCollectionUpdateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductCollectionsUpdateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductCollectionsUpdate Update a product collection

Update a product collection. `product_ids` replaces the membership in the given order. Send `expected_version` to refuse the change if the collection was edited since you read it.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection

**Common errors**

- `403 authorization_error` — the caller is identified but the action is denied. Codes include `insufficient_permissions`, `ip_not_allowed`, `merchant_mismatch`, and `origin_not_allowed`.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.
- `422 unprocessable_entity` — business-rule failure (for example, refunding more than the original charge).

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiProductCollectionsUpdateRequest
*/
func (a *ProductsAPIService) ProductCollectionsUpdate(ctx context.Context, id string) ApiProductCollectionsUpdateRequest {
	return ApiProductCollectionsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProductCollectionUpdateResponse
func (a *ProductsAPIService) ProductCollectionsUpdateExecute(r ApiProductCollectionsUpdateRequest) (*ProductCollectionUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCollectionUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductCollectionsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/product-collections/{id}"
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
	// body params
	localVarPostBody = r.productCollectionUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProductCollectionError
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
			var v ProductCollectionError
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
			var v ProductCollectionError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ProductCollectionError
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

type ApiProductsCreateRequest struct {
	ctx                  context.Context
	ApiService           ProductsAPI
	createProductRequest *CreateProductRequest
	runtimeOptions       []RequestOption
}

// Product creation details
func (r ApiProductsCreateRequest) CreateProductRequest(createProductRequest CreateProductRequest) ApiProductsCreateRequest {
	r.createProductRequest = &createProductRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductsCreateRequest) WithRequestOptions(opts ...RequestOption) ApiProductsCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductsCreateRequest) Execute() (*ProductCreateResponse, *http.Response, error) {
	var value *ProductCreateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductsCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductsCreate Create product

Create a new product. Requires 'products:write' scope.

---

**Related endpoints**

- `GET /products` — List products
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductsCreateRequest
*/
func (a *ProductsAPIService) ProductsCreate(ctx context.Context) ApiProductsCreateRequest {
	return ApiProductsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductCreateResponse
func (a *ProductsAPIService) ProductsCreateExecute(r ApiProductsCreateRequest) (*ProductCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createProductRequest == nil {
		return localVarReturnValue, nil, reportError("createProductRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.createProductRequest
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
			var v ProductsList401Response
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
			var v ProductsList401Response
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

type ApiProductsGetRequest struct {
	ctx            context.Context
	ApiService     ProductsAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductsGetRequest) WithRequestOptions(opts ...RequestOption) ApiProductsGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductsGetRequest) Execute() (*ProductRetrieveResponse, *http.Response, error) {
	var value *ProductRetrieveResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductsGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductsGet Get product by ID

Get a product by its ID. Requires 'products:read' scope.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiProductsGetRequest
*/
func (a *ProductsAPIService) ProductsGet(ctx context.Context, id string) ApiProductsGetRequest {
	return ApiProductsGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProductRetrieveResponse
func (a *ProductsAPIService) ProductsGetExecute(r ApiProductsGetRequest) (*ProductRetrieveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductRetrieveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/{id}"
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
			var v ProductsList401Response
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
			var v ProductsList401Response
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

type ApiProductsListRequest struct {
	ctx            context.Context
	ApiService     ProductsAPI
	page           *int32
	limit          *int32
	search         *string
	runtimeOptions []RequestOption
}

// Page number
func (r ApiProductsListRequest) Page(page int32) ApiProductsListRequest {
	r.page = &page
	return r
}

// Maximum number of results (1-100)
func (r ApiProductsListRequest) Limit(limit int32) ApiProductsListRequest {
	r.limit = &limit
	return r
}

// Search term to filter products
func (r ApiProductsListRequest) Search(search string) ApiProductsListRequest {
	r.search = &search
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductsListRequest) WithRequestOptions(opts ...RequestOption) ApiProductsListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductsListRequest) Execute() (*ProductListResponse, *http.Response, error) {
	var value *ProductListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductsListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductsList List products

List products with pagination and optional search.

---

**Related endpoints**

- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PATCH /products/{id}` — Update product
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiProductsListRequest
*/
func (a *ProductsAPIService) ProductsList(ctx context.Context) ApiProductsListRequest {
	return ApiProductsListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProductListResponse
func (a *ProductsAPIService) ProductsListExecute(r ApiProductsListRequest) (*ProductListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "form", "")
		r.limit = &defaultValue
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProductErrorResponse
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
			var v ProductsList401Response
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

type ApiProductsUpdateRequest struct {
	ctx                  context.Context
	ApiService           ProductsAPI
	id                   string
	updateProductRequest *UpdateProductRequest
	runtimeOptions       []RequestOption
}

// Product update details
func (r ApiProductsUpdateRequest) UpdateProductRequest(updateProductRequest UpdateProductRequest) ApiProductsUpdateRequest {
	r.updateProductRequest = &updateProductRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiProductsUpdateRequest) WithRequestOptions(opts ...RequestOption) ApiProductsUpdateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiProductsUpdateRequest) Execute() (*ProductUpdateResponse, *http.Response, error) {
	var value *ProductUpdateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.ProductsUpdateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
ProductsUpdate Update product

Update an existing product. Requires 'products:write' scope.

---

**Related endpoints**

- `GET /products` — List products
- `POST /products` — Create product
- `GET /products/{id}` — Get product by ID
- `PUT /products/external/{source}/{externalId}` — Upsert product by external ID
- `PUT /products/external/batch` — Batch upsert products by external ID
- `GET /product-collections` — List product collections
- `POST /product-collections` — Create a product collection
- `GET /product-collections/{id}` — Get a product collection

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiProductsUpdateRequest
*/
func (a *ProductsAPIService) ProductsUpdate(ctx context.Context, id string) ApiProductsUpdateRequest {
	return ApiProductsUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProductUpdateResponse
func (a *ProductsAPIService) ProductsUpdateExecute(r ApiProductsUpdateRequest) (*ProductUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductsAPIService.ProductsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateProductRequest == nil {
		return localVarReturnValue, nil, reportError("updateProductRequest is required and must be specified")
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
	// body params
	localVarPostBody = r.updateProductRequest
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
			var v ProductsList401Response
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
			var v ProductsList401Response
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
			var v ProductsList401Response
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
