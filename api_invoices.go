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

type InvoicesAPI interface {

	/*
			InvoiceUsageExplain Explain invoice usage charges

			Return why this invoice billed the amount it did: line-item attribution plus usage provenance counts for usage lines. Use this instead of chaining invoices_get + invoice_line_items_list_usage_events.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
		- `POST /invoices/{id}/reject` — Reject invoice

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Invoice UUID
			@return ApiInvoiceUsageExplainRequest
	*/
	InvoiceUsageExplain(ctx context.Context, id string) ApiInvoiceUsageExplainRequest

	// InvoiceUsageExplainExecute executes the request
	//  @return InvoiceUsageExplain200Response
	InvoiceUsageExplainExecute(r ApiInvoiceUsageExplainRequest) (*InvoiceUsageExplain200Response, *http.Response, error)

	/*
			InvoicesCreate Create invoice

			Create a new invoice.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
		- `POST /invoices/{id}/reject` — Reject invoice

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiInvoicesCreateRequest
	*/
	InvoicesCreate(ctx context.Context) ApiInvoicesCreateRequest

	// InvoicesCreateExecute executes the request
	//  @return InvoiceResponse
	InvoicesCreateExecute(r ApiInvoicesCreateRequest) (*InvoiceResponse, *http.Response, error)

	/*
			InvoicesFinalize Finalize an invoice

			Finalizes a draft invoice, locking it for payment. Assigns invoice number and generates public token. After finalization, financial fields become immutable.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiInvoicesFinalizeRequest
	*/
	InvoicesFinalize(ctx context.Context, id string) ApiInvoicesFinalizeRequest

	// InvoicesFinalizeExecute executes the request
	//  @return InvoiceResponse
	InvoicesFinalizeExecute(r ApiInvoicesFinalizeRequest) (*InvoiceResponse, *http.Response, error)

	/*
			InvoicesGet Get invoice

			Get a single invoice by ID.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
		- `POST /invoices/{id}/reject` — Reject invoice

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiInvoicesGetRequest
	*/
	InvoicesGet(ctx context.Context, id string) ApiInvoicesGetRequest

	// InvoicesGetExecute executes the request
	//  @return InvoiceResponse
	InvoicesGetExecute(r ApiInvoicesGetRequest) (*InvoiceResponse, *http.Response, error)

	/*
			InvoicesList List invoices

			List invoices with pagination and filtering.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
		- `POST /invoices/{id}/reject` — Reject invoice

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiInvoicesListRequest
	*/
	InvoicesList(ctx context.Context) ApiInvoicesListRequest

	// InvoicesListExecute executes the request
	//  @return InvoiceListResponse
	InvoicesListExecute(r ApiInvoicesListRequest) (*InvoiceListResponse, *http.Response, error)

	/*
			InvoicesMarginEstimate Estimate invoice net by payment rail

			Estimate per-rail processing fees and net-after-fees for an invoice amount.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/reject` — Reject invoice

		**Common errors**

		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Invoice ID.
			@return ApiInvoicesMarginEstimateRequest
	*/
	InvoicesMarginEstimate(ctx context.Context, id string) ApiInvoicesMarginEstimateRequest

	// InvoicesMarginEstimateExecute executes the request
	//  @return InvoicesMarginEstimate200Response
	InvoicesMarginEstimateExecute(r ApiInvoicesMarginEstimateRequest) (*InvoicesMarginEstimate200Response, *http.Response, error)

	/*
			InvoicesSend Send an invoice

			Sends an invoice to the customer via the specified channel (email, SMS, or WhatsApp). Invoice must be approved first.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiInvoicesSendRequest
	*/
	InvoicesSend(ctx context.Context, id string) ApiInvoicesSendRequest

	// InvoicesSendExecute executes the request
	//  @return InvoiceResponse
	InvoicesSendExecute(r ApiInvoicesSendRequest) (*InvoiceResponse, *http.Response, error)

	/*
			InvoicesUpdate Update invoice

			Update an existing invoice.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
		- `POST /invoices/{id}/reject` — Reject invoice

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiInvoicesUpdateRequest
	*/
	InvoicesUpdate(ctx context.Context, id string) ApiInvoicesUpdateRequest

	// InvoicesUpdateExecute executes the request
	//  @return InvoiceResponse
	InvoicesUpdateExecute(r ApiInvoicesUpdateRequest) (*InvoiceResponse, *http.Response, error)

	/*
			InvoicesVoid Void an invoice

			Voids an invoice. Only invoices without recorded payments can be voided. Use refund instead for paid invoices.

		---

		**Related endpoints**

		- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
		- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
		- `GET /invoices` — List invoices
		- `POST /invoices` — Create invoice
		- `GET /invoices/{id}` — Get invoice
		- `PATCH /invoices/{id}` — Update invoice
		- `DELETE /invoices/{id}` — Delete invoice
		- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id
			@return ApiInvoicesVoidRequest
	*/
	InvoicesVoid(ctx context.Context, id string) ApiInvoicesVoidRequest

	// InvoicesVoidExecute executes the request
	//  @return InvoiceResponse
	InvoicesVoidExecute(r ApiInvoicesVoidRequest) (*InvoiceResponse, *http.Response, error)
}

// InvoicesAPIService InvoicesAPI service
type InvoicesAPIService service

type ApiInvoiceUsageExplainRequest struct {
	ctx            context.Context
	ApiService     InvoicesAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoiceUsageExplainRequest) WithRequestOptions(opts ...RequestOption) ApiInvoiceUsageExplainRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoiceUsageExplainRequest) Execute() (*InvoiceUsageExplain200Response, *http.Response, error) {
	var value *InvoiceUsageExplain200Response
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoiceUsageExplainExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoiceUsageExplain Explain invoice usage charges

Return why this invoice billed the amount it did: line-item attribution plus usage provenance counts for usage lines. Use this instead of chaining invoices_get + invoice_line_items_list_usage_events.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
- `POST /invoices/{id}/reject` — Reject invoice

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Invoice UUID
	@return ApiInvoiceUsageExplainRequest
*/
func (a *InvoicesAPIService) InvoiceUsageExplain(ctx context.Context, id string) ApiInvoiceUsageExplainRequest {
	return ApiInvoiceUsageExplainRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceUsageExplain200Response
func (a *InvoicesAPIService) InvoiceUsageExplainExecute(r ApiInvoiceUsageExplainRequest) (*InvoiceUsageExplain200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceUsageExplain200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoiceUsageExplain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}/usage-explain"
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

type ApiInvoicesCreateRequest struct {
	ctx                   context.Context
	ApiService            InvoicesAPI
	invoicesCreateRequest *InvoicesCreateRequest
	idempotencyKey        *string
	runtimeOptions        []RequestOption
}

// Invoice creation details
func (r ApiInvoicesCreateRequest) InvoicesCreateRequest(invoicesCreateRequest InvoicesCreateRequest) ApiInvoicesCreateRequest {
	r.invoicesCreateRequest = &invoicesCreateRequest
	return r
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesCreateRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesCreateRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesCreateRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesCreate Create invoice

Create a new invoice.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
- `POST /invoices/{id}/reject` — Reject invoice

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInvoicesCreateRequest
*/
func (a *InvoicesAPIService) InvoicesCreate(ctx context.Context) ApiInvoicesCreateRequest {
	return ApiInvoicesCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesCreateExecute(r ApiInvoicesCreateRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoicesCreateRequest == nil {
		return localVarReturnValue, nil, reportError("invoicesCreateRequest is required and must be specified")
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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.invoicesCreateRequest
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
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiInvoicesFinalizeRequest struct {
	ctx                     context.Context
	ApiService              InvoicesAPI
	id                      string
	idempotencyKey          *string
	invoicesFinalizeRequest *InvoicesFinalizeRequest
	runtimeOptions          []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesFinalizeRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesFinalizeRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// Finalization options
func (r ApiInvoicesFinalizeRequest) InvoicesFinalizeRequest(invoicesFinalizeRequest InvoicesFinalizeRequest) ApiInvoicesFinalizeRequest {
	r.invoicesFinalizeRequest = &invoicesFinalizeRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesFinalizeRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesFinalizeRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesFinalizeRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesFinalizeExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesFinalize Finalize an invoice

Finalizes a draft invoice, locking it for payment. Assigns invoice number and generates public token. After finalization, financial fields become immutable.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiInvoicesFinalizeRequest
*/
func (a *InvoicesAPIService) InvoicesFinalize(ctx context.Context, id string) ApiInvoicesFinalizeRequest {
	return ApiInvoicesFinalizeRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesFinalizeExecute(r ApiInvoicesFinalizeRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesFinalize")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}/finalize"
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
	localVarPostBody = r.invoicesFinalizeRequest
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

type ApiInvoicesGetRequest struct {
	ctx            context.Context
	ApiService     InvoicesAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesGetRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesGetRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesGet Get invoice

Get a single invoice by ID.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
- `POST /invoices/{id}/reject` — Reject invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiInvoicesGetRequest
*/
func (a *InvoicesAPIService) InvoicesGet(ctx context.Context, id string) ApiInvoicesGetRequest {
	return ApiInvoicesGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesGetExecute(r ApiInvoicesGetRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}"
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
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiInvoicesListRequest struct {
	ctx            context.Context
	ApiService     InvoicesAPI
	status         *string
	customerId     *string
	limit          *int32
	offset         *int32
	runtimeOptions []RequestOption
}

// Filter by invoice status
func (r ApiInvoicesListRequest) Status(status string) ApiInvoicesListRequest {
	r.status = &status
	return r
}

// Filter by customer ID
func (r ApiInvoicesListRequest) CustomerId(customerId string) ApiInvoicesListRequest {
	r.customerId = &customerId
	return r
}

// Maximum number of results (1-100)
func (r ApiInvoicesListRequest) Limit(limit int32) ApiInvoicesListRequest {
	r.limit = &limit
	return r
}

// Number of results to skip
func (r ApiInvoicesListRequest) Offset(offset int32) ApiInvoicesListRequest {
	r.offset = &offset
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesListRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesListRequest) Execute() (*InvoiceListResponse, *http.Response, error) {
	var value *InvoiceListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesList List invoices

List invoices with pagination and filtering.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
- `POST /invoices/{id}/reject` — Reject invoice

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInvoicesListRequest
*/
func (a *InvoicesAPIService) InvoicesList(ctx context.Context) ApiInvoicesListRequest {
	return ApiInvoicesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InvoiceListResponse
func (a *InvoicesAPIService) InvoicesListExecute(r ApiInvoicesListRequest) (*InvoiceListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerId", r.customerId, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "form", "")
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	} else {
		var defaultValue int32 = 0
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", defaultValue, "form", "")
		r.offset = &defaultValue
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
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiInvoicesMarginEstimateRequest struct {
	ctx                           context.Context
	ApiService                    InvoicesAPI
	id                            string
	invoicesMarginEstimateRequest *InvoicesMarginEstimateRequest
	idempotencyKey                *string
	runtimeOptions                []RequestOption
}

// Margin estimate inputs
func (r ApiInvoicesMarginEstimateRequest) InvoicesMarginEstimateRequest(invoicesMarginEstimateRequest InvoicesMarginEstimateRequest) ApiInvoicesMarginEstimateRequest {
	r.invoicesMarginEstimateRequest = &invoicesMarginEstimateRequest
	return r
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesMarginEstimateRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesMarginEstimateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesMarginEstimateRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesMarginEstimateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesMarginEstimateRequest) Execute() (*InvoicesMarginEstimate200Response, *http.Response, error) {
	var value *InvoicesMarginEstimate200Response
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesMarginEstimateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesMarginEstimate Estimate invoice net by payment rail

Estimate per-rail processing fees and net-after-fees for an invoice amount.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/reject` — Reject invoice

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Invoice ID.
	@return ApiInvoicesMarginEstimateRequest
*/
func (a *InvoicesAPIService) InvoicesMarginEstimate(ctx context.Context, id string) ApiInvoicesMarginEstimateRequest {
	return ApiInvoicesMarginEstimateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoicesMarginEstimate200Response
func (a *InvoicesAPIService) InvoicesMarginEstimateExecute(r ApiInvoicesMarginEstimateRequest) (*InvoicesMarginEstimate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoicesMarginEstimate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesMarginEstimate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}/margin-estimate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoicesMarginEstimateRequest == nil {
		return localVarReturnValue, nil, reportError("invoicesMarginEstimateRequest is required and must be specified")
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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.invoicesMarginEstimateRequest
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

type ApiInvoicesSendRequest struct {
	ctx                 context.Context
	ApiService          InvoicesAPI
	id                  string
	idempotencyKey      *string
	invoicesSendRequest *InvoicesSendRequest
	runtimeOptions      []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesSendRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesSendRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// Send options
func (r ApiInvoicesSendRequest) InvoicesSendRequest(invoicesSendRequest InvoicesSendRequest) ApiInvoicesSendRequest {
	r.invoicesSendRequest = &invoicesSendRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesSendRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesSendRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesSendRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesSendExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesSend Send an invoice

Sends an invoice to the customer via the specified channel (email, SMS, or WhatsApp). Invoice must be approved first.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiInvoicesSendRequest
*/
func (a *InvoicesAPIService) InvoicesSend(ctx context.Context, id string) ApiInvoicesSendRequest {
	return ApiInvoicesSendRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesSendExecute(r ApiInvoicesSendRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesSend")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}/send"
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
	localVarPostBody = r.invoicesSendRequest
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

type ApiInvoicesUpdateRequest struct {
	ctx                   context.Context
	ApiService            InvoicesAPI
	id                    string
	invoicesUpdateRequest *InvoicesUpdateRequest
	idempotencyKey        *string
	runtimeOptions        []RequestOption
}

// Invoice update details
func (r ApiInvoicesUpdateRequest) InvoicesUpdateRequest(invoicesUpdateRequest InvoicesUpdateRequest) ApiInvoicesUpdateRequest {
	r.invoicesUpdateRequest = &invoicesUpdateRequest
	return r
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesUpdateRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesUpdateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesUpdateRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesUpdateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesUpdateRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesUpdateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesUpdate Update invoice

Update an existing invoice.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail
- `POST /invoices/{id}/reject` — Reject invoice

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiInvoicesUpdateRequest
*/
func (a *InvoicesAPIService) InvoicesUpdate(ctx context.Context, id string) ApiInvoicesUpdateRequest {
	return ApiInvoicesUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesUpdateExecute(r ApiInvoicesUpdateRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.invoicesUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("invoicesUpdateRequest is required and must be specified")
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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.invoicesUpdateRequest
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
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiInvoicesVoidRequest struct {
	ctx                 context.Context
	ApiService          InvoicesAPI
	id                  string
	idempotencyKey      *string
	invoicesVoidRequest *InvoicesVoidRequest
	runtimeOptions      []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiInvoicesVoidRequest) IdempotencyKey(idempotencyKey string) ApiInvoicesVoidRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// Void details
func (r ApiInvoicesVoidRequest) InvoicesVoidRequest(invoicesVoidRequest InvoicesVoidRequest) ApiInvoicesVoidRequest {
	r.invoicesVoidRequest = &invoicesVoidRequest
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiInvoicesVoidRequest) WithRequestOptions(opts ...RequestOption) ApiInvoicesVoidRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiInvoicesVoidRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	var value *InvoiceResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.InvoicesVoidExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
InvoicesVoid Void an invoice

Voids an invoice. Only invoices without recorded payments can be voided. Use refund instead for paid invoices.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices/{id}/usage-explain` — Explain invoice usage charges
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/margin-estimate` — Estimate invoice net by payment rail

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiInvoicesVoidRequest
*/
func (a *InvoicesAPIService) InvoicesVoid(ctx context.Context, id string) ApiInvoicesVoidRequest {
	return ApiInvoicesVoidRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return InvoiceResponse
func (a *InvoicesAPIService) InvoicesVoidExecute(r ApiInvoicesVoidRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *InvoiceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesAPIService.InvoicesVoid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/{id}/void"
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
	localVarPostBody = r.invoicesVoidRequest
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
