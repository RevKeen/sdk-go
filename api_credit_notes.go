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

type CreditNotesAPI interface {

	/*
			CreditNotesCreate Create a credit note

			Issue a credit note for a paid or partially paid invoice. The credit can be applied via refund to payment method, customer balance, or marked as external.

		---

		**Related endpoints**

		- `GET /credit_notes` — List credit notes
		- `GET /credit_notes/{id}` — Get credit note by ID
		- `POST /credit_notes/{id}/void` — Void a credit note
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `GET /credit_notes/{id}/lines` — List line items on a credit note
		- `POST /credit_notes/preview` — Preview a credit note without creating it

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCreditNotesCreateRequest
	*/
	CreditNotesCreate(ctx context.Context) ApiCreditNotesCreateRequest

	// CreditNotesCreateExecute executes the request
	//  @return CreditNoteCreateResponse
	CreditNotesCreateExecute(r ApiCreditNotesCreateRequest) (*CreditNoteCreateResponse, *http.Response, error)

	/*
			CreditNotesGet Get credit note by ID

			Retrieve a single credit note by its UUID.

		---

		**Related endpoints**

		- `GET /credit_notes` — List credit notes
		- `POST /credit_notes` — Create a credit note
		- `POST /credit_notes/{id}/void` — Void a credit note
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `GET /credit_notes/{id}/lines` — List line items on a credit note
		- `POST /credit_notes/preview` — Preview a credit note without creating it

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Credit note UUID
			@return ApiCreditNotesGetRequest
	*/
	CreditNotesGet(ctx context.Context, id string) ApiCreditNotesGetRequest

	// CreditNotesGetExecute executes the request
	//  @return CreditNoteRetrieveResponse
	CreditNotesGetExecute(r ApiCreditNotesGetRequest) (*CreditNoteRetrieveResponse, *http.Response, error)

	/*
			CreditNotesList List credit notes

			Retrieve a paginated list of credit notes with optional filters. Results are ordered by creation date (newest first).

		---

		**Related endpoints**

		- `POST /credit_notes` — Create a credit note
		- `GET /credit_notes/{id}` — Get credit note by ID
		- `POST /credit_notes/{id}/void` — Void a credit note
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `GET /credit_notes/{id}/lines` — List line items on a credit note
		- `POST /credit_notes/preview` — Preview a credit note without creating it

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCreditNotesListRequest
	*/
	CreditNotesList(ctx context.Context) ApiCreditNotesListRequest

	// CreditNotesListExecute executes the request
	//  @return CreditNoteListResponse
	CreditNotesListExecute(r ApiCreditNotesListRequest) (*CreditNoteListResponse, *http.Response, error)

	/*
			CreditNotesListLines List line items on a credit note

			Returns the line items attached to a credit note. Each line describes a portion of the invoice that the credit applies to — either a whole invoice line or a prorated slice. Use this to reconstruct the credit's effect per invoice line (for example, when displaying the credit breakdown in your own UI).

		---

		**Related endpoints**

		- `GET /credit_notes` — List credit notes
		- `POST /credit_notes` — Create a credit note
		- `GET /credit_notes/{id}` — Get credit note by ID
		- `POST /credit_notes/{id}/void` — Void a credit note
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `POST /credit_notes/preview` — Preview a credit note without creating it

		**Common errors**

		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Pagination**

		Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Credit note UUID
			@return ApiCreditNotesListLinesRequest
	*/
	CreditNotesListLines(ctx context.Context, id string) ApiCreditNotesListLinesRequest

	// CreditNotesListLinesExecute executes the request
	//  @return CreditNoteLineList
	CreditNotesListLinesExecute(r ApiCreditNotesListLinesRequest) (*CreditNoteLineList, *http.Response, error)

	/*
			CreditNotesPreview Preview a credit note without creating it

			Compute what a credit note would look like against a specific invoice without persisting anything. Use this to validate amounts and show a "here's what will happen" UI before the merchant commits. No side effects — no DB writes, no events emitted, no refunds initiated. Mirrors Stripe's `POST /v1/credit_notes/preview`.

		---

		**Related endpoints**

		- `GET /credit_notes` — List credit notes
		- `POST /credit_notes` — Create a credit note
		- `GET /credit_notes/{id}` — Get credit note by ID
		- `POST /credit_notes/{id}/void` — Void a credit note
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `GET /credit_notes/{id}/lines` — List line items on a credit note

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCreditNotesPreviewRequest
	*/
	CreditNotesPreview(ctx context.Context) ApiCreditNotesPreviewRequest

	// CreditNotesPreviewExecute executes the request
	//  @return CreditNotePreviewResponse
	CreditNotesPreviewExecute(r ApiCreditNotesPreviewRequest) (*CreditNotePreviewResponse, *http.Response, error)

	/*
			CreditNotesVoid Void a credit note

			Void a credit note that has been issued but not yet fully applied. This is an accounting void — it reverses the credit without creating a new financial transaction.

		---

		**Related endpoints**

		- `GET /credit_notes` — List credit notes
		- `POST /credit_notes` — Create a credit note
		- `GET /credit_notes/{id}` — Get credit note by ID
		- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
		- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
		- `GET /credit_notes/{id}/lines` — List line items on a credit note
		- `POST /credit_notes/preview` — Preview a credit note without creating it

		**Common errors**

		- `400 invalid_request` — malformed payload or failed validation.
		- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
		- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

		**Idempotency**

		Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id Credit note UUID
			@return ApiCreditNotesVoidRequest
	*/
	CreditNotesVoid(ctx context.Context, id string) ApiCreditNotesVoidRequest

	// CreditNotesVoidExecute executes the request
	//  @return CreditNoteVoidResponse
	CreditNotesVoidExecute(r ApiCreditNotesVoidRequest) (*CreditNoteVoidResponse, *http.Response, error)
}

// CreditNotesAPIService CreditNotesAPI service
type CreditNotesAPIService service

type ApiCreditNotesCreateRequest struct {
	ctx                   context.Context
	ApiService            CreditNotesAPI
	idempotencyKey        *string
	createCreditNoteInput *CreateCreditNoteInput
	runtimeOptions        []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiCreditNotesCreateRequest) IdempotencyKey(idempotencyKey string) ApiCreditNotesCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiCreditNotesCreateRequest) CreateCreditNoteInput(createCreditNoteInput CreateCreditNoteInput) ApiCreditNotesCreateRequest {
	r.createCreditNoteInput = &createCreditNoteInput
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesCreateRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesCreateRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesCreateRequest) Execute() (*CreditNoteCreateResponse, *http.Response, error) {
	var value *CreditNoteCreateResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesCreateExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesCreate Create a credit note

Issue a credit note for a paid or partially paid invoice. The credit can be applied via refund to payment method, customer balance, or marked as external.

---

**Related endpoints**

- `GET /credit_notes` — List credit notes
- `GET /credit_notes/{id}` — Get credit note by ID
- `POST /credit_notes/{id}/void` — Void a credit note
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `GET /credit_notes/{id}/lines` — List line items on a credit note
- `POST /credit_notes/preview` — Preview a credit note without creating it

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreditNotesCreateRequest
*/
func (a *CreditNotesAPIService) CreditNotesCreate(ctx context.Context) ApiCreditNotesCreateRequest {
	return ApiCreditNotesCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreditNoteCreateResponse
func (a *CreditNotesAPIService) CreditNotesCreateExecute(r ApiCreditNotesCreateRequest) (*CreditNoteCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNoteCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes"

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
	localVarPostBody = r.createCreditNoteInput
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

type ApiCreditNotesGetRequest struct {
	ctx            context.Context
	ApiService     CreditNotesAPI
	id             string
	runtimeOptions []RequestOption
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesGetRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesGetRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesGetRequest) Execute() (*CreditNoteRetrieveResponse, *http.Response, error) {
	var value *CreditNoteRetrieveResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesGetExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesGet Get credit note by ID

Retrieve a single credit note by its UUID.

---

**Related endpoints**

- `GET /credit_notes` — List credit notes
- `POST /credit_notes` — Create a credit note
- `POST /credit_notes/{id}/void` — Void a credit note
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `GET /credit_notes/{id}/lines` — List line items on a credit note
- `POST /credit_notes/preview` — Preview a credit note without creating it

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Credit note UUID
	@return ApiCreditNotesGetRequest
*/
func (a *CreditNotesAPIService) CreditNotesGet(ctx context.Context, id string) ApiCreditNotesGetRequest {
	return ApiCreditNotesGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CreditNoteRetrieveResponse
func (a *CreditNotesAPIService) CreditNotesGetExecute(r ApiCreditNotesGetRequest) (*CreditNoteRetrieveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNoteRetrieveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes/{id}"
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

type ApiCreditNotesListRequest struct {
	ctx            context.Context
	ApiService     CreditNotesAPI
	status         *string
	invoiceId      *string
	customerId     *string
	creditMethod   *string
	createdAfter   *string
	createdBefore  *string
	limit          *float32
	offset         *float32
	runtimeOptions []RequestOption
}

// Filter by credit note status
func (r ApiCreditNotesListRequest) Status(status string) ApiCreditNotesListRequest {
	r.status = &status
	return r
}

// Filter by invoice ID
func (r ApiCreditNotesListRequest) InvoiceId(invoiceId string) ApiCreditNotesListRequest {
	r.invoiceId = &invoiceId
	return r
}

// Filter by customer ID
func (r ApiCreditNotesListRequest) CustomerId(customerId string) ApiCreditNotesListRequest {
	r.customerId = &customerId
	return r
}

// Filter by credit method
func (r ApiCreditNotesListRequest) CreditMethod(creditMethod string) ApiCreditNotesListRequest {
	r.creditMethod = &creditMethod
	return r
}

// ISO 8601 date - only credit notes created after this date
func (r ApiCreditNotesListRequest) CreatedAfter(createdAfter string) ApiCreditNotesListRequest {
	r.createdAfter = &createdAfter
	return r
}

// ISO 8601 date - only credit notes created before this date
func (r ApiCreditNotesListRequest) CreatedBefore(createdBefore string) ApiCreditNotesListRequest {
	r.createdBefore = &createdBefore
	return r
}

// Number of results to return (1-100)
func (r ApiCreditNotesListRequest) Limit(limit float32) ApiCreditNotesListRequest {
	r.limit = &limit
	return r
}

// Number of results to skip
func (r ApiCreditNotesListRequest) Offset(offset float32) ApiCreditNotesListRequest {
	r.offset = &offset
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesListRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesListRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesListRequest) Execute() (*CreditNoteListResponse, *http.Response, error) {
	var value *CreditNoteListResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesListExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesList List credit notes

Retrieve a paginated list of credit notes with optional filters. Results are ordered by creation date (newest first).

---

**Related endpoints**

- `POST /credit_notes` — Create a credit note
- `GET /credit_notes/{id}` — Get credit note by ID
- `POST /credit_notes/{id}/void` — Void a credit note
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `GET /credit_notes/{id}/lines` — List line items on a credit note
- `POST /credit_notes/preview` — Preview a credit note without creating it

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreditNotesListRequest
*/
func (a *CreditNotesAPIService) CreditNotesList(ctx context.Context) ApiCreditNotesListRequest {
	return ApiCreditNotesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreditNoteListResponse
func (a *CreditNotesAPIService) CreditNotesListExecute(r ApiCreditNotesListRequest) (*CreditNoteListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNoteListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.invoiceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoice_id", r.invoiceId, "form", "")
	}
	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "form", "")
	}
	if r.creditMethod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "credit_method", r.creditMethod, "form", "")
	}
	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_after", r.createdAfter, "form", "")
	}
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_before", r.createdBefore, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue float32 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "form", "")
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	} else {
		var defaultValue float32 = 0
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

type ApiCreditNotesListLinesRequest struct {
	ctx            context.Context
	ApiService     CreditNotesAPI
	id             string
	limit          *int32
	runtimeOptions []RequestOption
}

// Maximum number of line items to return (1-100, default 100).
func (r ApiCreditNotesListLinesRequest) Limit(limit int32) ApiCreditNotesListLinesRequest {
	r.limit = &limit
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesListLinesRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesListLinesRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesListLinesRequest) Execute() (*CreditNoteLineList, *http.Response, error) {
	var value *CreditNoteLineList
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, false)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesListLinesExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesListLines List line items on a credit note

Returns the line items attached to a credit note. Each line describes a portion of the invoice that the credit applies to — either a whole invoice line or a prorated slice. Use this to reconstruct the credit's effect per invoice line (for example, when displaying the credit breakdown in your own UI).

---

**Related endpoints**

- `GET /credit_notes` — List credit notes
- `POST /credit_notes` — Create a credit note
- `GET /credit_notes/{id}` — Get credit note by ID
- `POST /credit_notes/{id}/void` — Void a credit note
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `POST /credit_notes/preview` — Preview a credit note without creating it

**Common errors**

- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Credit note UUID
	@return ApiCreditNotesListLinesRequest
*/
func (a *CreditNotesAPIService) CreditNotesListLines(ctx context.Context, id string) ApiCreditNotesListLinesRequest {
	return ApiCreditNotesListLinesRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CreditNoteLineList
func (a *CreditNotesAPIService) CreditNotesListLinesExecute(r ApiCreditNotesListLinesRequest) (*CreditNoteLineList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNoteLineList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesListLines")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes/{id}/lines"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", defaultValue, "form", "")
		r.limit = &defaultValue
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

type ApiCreditNotesPreviewRequest struct {
	ctx                    context.Context
	ApiService             CreditNotesAPI
	previewCreditNoteInput *PreviewCreditNoteInput
	idempotencyKey         *string
	runtimeOptions         []RequestOption
}

func (r ApiCreditNotesPreviewRequest) PreviewCreditNoteInput(previewCreditNoteInput PreviewCreditNoteInput) ApiCreditNotesPreviewRequest {
	r.previewCreditNoteInput = &previewCreditNoteInput
	return r
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiCreditNotesPreviewRequest) IdempotencyKey(idempotencyKey string) ApiCreditNotesPreviewRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesPreviewRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesPreviewRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesPreviewRequest) Execute() (*CreditNotePreviewResponse, *http.Response, error) {
	var value *CreditNotePreviewResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesPreviewExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesPreview Preview a credit note without creating it

Compute what a credit note would look like against a specific invoice without persisting anything. Use this to validate amounts and show a "here's what will happen" UI before the merchant commits. No side effects — no DB writes, no events emitted, no refunds initiated. Mirrors Stripe's `POST /v1/credit_notes/preview`.

---

**Related endpoints**

- `GET /credit_notes` — List credit notes
- `POST /credit_notes` — Create a credit note
- `GET /credit_notes/{id}` — Get credit note by ID
- `POST /credit_notes/{id}/void` — Void a credit note
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `GET /credit_notes/{id}/lines` — List line items on a credit note

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreditNotesPreviewRequest
*/
func (a *CreditNotesAPIService) CreditNotesPreview(ctx context.Context) ApiCreditNotesPreviewRequest {
	return ApiCreditNotesPreviewRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreditNotePreviewResponse
func (a *CreditNotesAPIService) CreditNotesPreviewExecute(r ApiCreditNotesPreviewRequest) (*CreditNotePreviewResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNotePreviewResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesPreview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes/preview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.previewCreditNoteInput == nil {
		return localVarReturnValue, nil, reportError("previewCreditNoteInput is required and must be specified")
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
	localVarPostBody = r.previewCreditNoteInput
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

type ApiCreditNotesVoidRequest struct {
	ctx            context.Context
	ApiService     CreditNotesAPI
	id             string
	idempotencyKey *string
	runtimeOptions []RequestOption
}

// Stable caller-owned key for this logical mutation. Reuse only with an identical request.
func (r ApiCreditNotesVoidRequest) IdempotencyKey(idempotencyKey string) ApiCreditNotesVoidRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

// WithRequestOptions applies per-call IDs and timeout overrides to this request.
func (r ApiCreditNotesVoidRequest) WithRequestOptions(opts ...RequestOption) ApiCreditNotesVoidRequest {
	r.runtimeOptions = append(r.runtimeOptions, opts...)
	return r
}

func (r ApiCreditNotesVoidRequest) Execute() (*CreditNoteVoidResponse, *http.Response, error) {
	var value *CreditNoteVoidResponse
	preparedContext, optionErr := prepareRuntimeRequest(r.ctx, r.runtimeOptions, true)
	r.ctx = preparedContext
	if optionErr != nil {
		return value, nil, wrapRequestError(r.ctx, nil, optionErr)
	}
	value, response, err := r.ApiService.CreditNotesVoidExecute(r)
	return value, response, wrapRequestError(r.ctx, response, err)
}

/*
CreditNotesVoid Void a credit note

Void a credit note that has been issued but not yet fully applied. This is an accounting void — it reverses the credit without creating a new financial transaction.

---

**Related endpoints**

- `GET /credit_notes` — List credit notes
- `POST /credit_notes` — Create a credit note
- `GET /credit_notes/{id}` — Get credit note by ID
- `GET /credit_notes/invoice/{invoice_id}/eligibility` — Check credit note eligibility for an invoice
- `GET /credit_notes/transaction/{transaction_id}/reversal-eligibility` — Check reversal eligibility for a transaction
- `GET /credit_notes/{id}/lines` — List line items on a credit note
- `POST /credit_notes/preview` — Preview a credit note without creating it

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 authentication_error` — missing, invalid, expired, or revoked credential. Codes: `authentication_failed`, `invalid_api_key`, `expired_api_key`, `api_key_revoked`, `session_invalid`, `merchant_required`. Carries a `WWW-Authenticate: Bearer` challenge.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Credit note UUID
	@return ApiCreditNotesVoidRequest
*/
func (a *CreditNotesAPIService) CreditNotesVoid(ctx context.Context, id string) ApiCreditNotesVoidRequest {
	return ApiCreditNotesVoidRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CreditNoteVoidResponse
func (a *CreditNotesAPIService) CreditNotesVoidExecute(r ApiCreditNotesVoidRequest) (*CreditNoteVoidResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreditNoteVoidResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CreditNotesAPIService.CreditNotesVoid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_notes/{id}/void"
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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
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
