/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
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


type PaymentLinksAPI interface {

	/*
	PaymentLinksActivate Activate a payment link (deprecated)

	**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "active"` instead.** This convenience endpoint will be removed in a future API version. Enable a payment link to accept payments. Cannot be used on archived links.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksActivateRequest

	Deprecated
	*/
	PaymentLinksActivate(ctx context.Context, id string) ApiPaymentLinksActivateRequest

	// PaymentLinksActivateExecute executes the request
	//  @return PaymentLinksUpdate200Response
	// Deprecated
	PaymentLinksActivateExecute(r ApiPaymentLinksActivateRequest) (*PaymentLinksUpdate200Response, *http.Response, error)

	/*
	PaymentLinksArchive Archive a payment link (deprecated)

	**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "archived"` instead.** This convenience endpoint will be removed in a future API version. Permanently disable a payment link. This action cannot be undone.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksArchiveRequest

	Deprecated
	*/
	PaymentLinksArchive(ctx context.Context, id string) ApiPaymentLinksArchiveRequest

	// PaymentLinksArchiveExecute executes the request
	//  @return PaymentLinksUpdate200Response
	// Deprecated
	PaymentLinksArchiveExecute(r ApiPaymentLinksArchiveRequest) (*PaymentLinksUpdate200Response, *http.Response, error)

	/*
	PaymentLinksCreate Create a new payment link

	Create a new payment link for invoices, subscriptions, or custom amounts

---

**Related endpoints**

- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPaymentLinksCreateRequest
	*/
	PaymentLinksCreate(ctx context.Context) ApiPaymentLinksCreateRequest

	// PaymentLinksCreateExecute executes the request
	//  @return PaymentLinksList200ResponseDataInner
	PaymentLinksCreateExecute(r ApiPaymentLinksCreateRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error)

	/*
	PaymentLinksDeactivate Deactivate a payment link (deprecated)

	**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "inactive"` instead.** This convenience endpoint will be removed in a future API version. Temporarily disable a payment link. It can be reactivated later.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksDeactivateRequest

	Deprecated
	*/
	PaymentLinksDeactivate(ctx context.Context, id string) ApiPaymentLinksDeactivateRequest

	// PaymentLinksDeactivateExecute executes the request
	//  @return PaymentLinksUpdate200Response
	// Deprecated
	PaymentLinksDeactivateExecute(r ApiPaymentLinksDeactivateRequest) (*PaymentLinksUpdate200Response, *http.Response, error)

	/*
	PaymentLinksExpire Expire a payment link (deprecated)

	**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "expired"` instead.** This convenience endpoint will be removed in a future API version. Mark a payment link as expired, preventing further use.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksExpireRequest

	Deprecated
	*/
	PaymentLinksExpire(ctx context.Context, id string) ApiPaymentLinksExpireRequest

	// PaymentLinksExpireExecute executes the request
	//  @return PaymentLinksList200ResponseDataInner
	// Deprecated
	PaymentLinksExpireExecute(r ApiPaymentLinksExpireRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error)

	/*
	PaymentLinksGet Get payment link by ID

	Retrieve a payment link by its UUID or public_id

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksGetRequest
	*/
	PaymentLinksGet(ctx context.Context, id string) ApiPaymentLinksGetRequest

	// PaymentLinksGetExecute executes the request
	//  @return PaymentLinksList200ResponseDataInner
	PaymentLinksGetExecute(r ApiPaymentLinksGetRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error)

	/*
	PaymentLinksList List payment links

	List payment links for the authenticated merchant with pagination and filters

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPaymentLinksListRequest
	*/
	PaymentLinksList(ctx context.Context) ApiPaymentLinksListRequest

	// PaymentLinksListExecute executes the request
	//  @return PaymentLinksList200Response
	PaymentLinksListExecute(r ApiPaymentLinksListRequest) (*PaymentLinksList200Response, *http.Response, error)

	/*
	PaymentLinksUpdate Update payment link status

	Change the status of a payment link. Active links accept payments, inactive links are temporarily disabled, and archived links are permanently disabled and cannot be reactivated.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment link UUID or public_id
	@return ApiPaymentLinksUpdateRequest
	*/
	PaymentLinksUpdate(ctx context.Context, id string) ApiPaymentLinksUpdateRequest

	// PaymentLinksUpdateExecute executes the request
	//  @return PaymentLinksUpdate200Response
	PaymentLinksUpdateExecute(r ApiPaymentLinksUpdateRequest) (*PaymentLinksUpdate200Response, *http.Response, error)
}

// PaymentLinksAPIService PaymentLinksAPI service
type PaymentLinksAPIService service

type ApiPaymentLinksActivateRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
}

func (r ApiPaymentLinksActivateRequest) Execute() (*PaymentLinksUpdate200Response, *http.Response, error) {
	return r.ApiService.PaymentLinksActivateExecute(r)
}

/*
PaymentLinksActivate Activate a payment link (deprecated)

**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "active"` instead.** This convenience endpoint will be removed in a future API version. Enable a payment link to accept payments. Cannot be used on archived links.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksActivateRequest

Deprecated
*/
func (a *PaymentLinksAPIService) PaymentLinksActivate(ctx context.Context, id string) ApiPaymentLinksActivateRequest {
	return ApiPaymentLinksActivateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksUpdate200Response
// Deprecated
func (a *PaymentLinksAPIService) PaymentLinksActivateExecute(r ApiPaymentLinksActivateRequest) (*PaymentLinksUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksActivate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
			var v PaymentLinksList400Response
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
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksArchiveRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
}

func (r ApiPaymentLinksArchiveRequest) Execute() (*PaymentLinksUpdate200Response, *http.Response, error) {
	return r.ApiService.PaymentLinksArchiveExecute(r)
}

/*
PaymentLinksArchive Archive a payment link (deprecated)

**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "archived"` instead.** This convenience endpoint will be removed in a future API version. Permanently disable a payment link. This action cannot be undone.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksArchiveRequest

Deprecated
*/
func (a *PaymentLinksAPIService) PaymentLinksArchive(ctx context.Context, id string) ApiPaymentLinksArchiveRequest {
	return ApiPaymentLinksArchiveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksUpdate200Response
// Deprecated
func (a *PaymentLinksAPIService) PaymentLinksArchiveExecute(r ApiPaymentLinksArchiveRequest) (*PaymentLinksUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}/archive"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksCreateRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	paymentLinksCreateRequest *PaymentLinksCreateRequest
}

func (r ApiPaymentLinksCreateRequest) PaymentLinksCreateRequest(paymentLinksCreateRequest PaymentLinksCreateRequest) ApiPaymentLinksCreateRequest {
	r.paymentLinksCreateRequest = &paymentLinksCreateRequest
	return r
}

func (r ApiPaymentLinksCreateRequest) Execute() (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PaymentLinksCreateExecute(r)
}

/*
PaymentLinksCreate Create a new payment link

Create a new payment link for invoices, subscriptions, or custom amounts

---

**Related endpoints**

- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaymentLinksCreateRequest
*/
func (a *PaymentLinksAPIService) PaymentLinksCreate(ctx context.Context) ApiPaymentLinksCreateRequest {
	return ApiPaymentLinksCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentLinksList200ResponseDataInner
func (a *PaymentLinksAPIService) PaymentLinksCreateExecute(r ApiPaymentLinksCreateRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksList200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links"

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
	localVarPostBody = r.paymentLinksCreateRequest
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
			var v PaymentLinksList400Response
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
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksDeactivateRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
}

func (r ApiPaymentLinksDeactivateRequest) Execute() (*PaymentLinksUpdate200Response, *http.Response, error) {
	return r.ApiService.PaymentLinksDeactivateExecute(r)
}

/*
PaymentLinksDeactivate Deactivate a payment link (deprecated)

**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "inactive"` instead.** This convenience endpoint will be removed in a future API version. Temporarily disable a payment link. It can be reactivated later.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksDeactivateRequest

Deprecated
*/
func (a *PaymentLinksAPIService) PaymentLinksDeactivate(ctx context.Context, id string) ApiPaymentLinksDeactivateRequest {
	return ApiPaymentLinksDeactivateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksUpdate200Response
// Deprecated
func (a *PaymentLinksAPIService) PaymentLinksDeactivateExecute(r ApiPaymentLinksDeactivateRequest) (*PaymentLinksUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksDeactivate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
			var v PaymentLinksList400Response
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
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksExpireRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
	paymentLinksExpireRequest *PaymentLinksExpireRequest
}

func (r ApiPaymentLinksExpireRequest) PaymentLinksExpireRequest(paymentLinksExpireRequest PaymentLinksExpireRequest) ApiPaymentLinksExpireRequest {
	r.paymentLinksExpireRequest = &paymentLinksExpireRequest
	return r
}

func (r ApiPaymentLinksExpireRequest) Execute() (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PaymentLinksExpireExecute(r)
}

/*
PaymentLinksExpire Expire a payment link (deprecated)

**Deprecated — use `PATCH /v2/payment-links/{id}/status` with `status: "expired"` instead.** This convenience endpoint will be removed in a future API version. Mark a payment link as expired, preventing further use.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksExpireRequest

Deprecated
*/
func (a *PaymentLinksAPIService) PaymentLinksExpire(ctx context.Context, id string) ApiPaymentLinksExpireRequest {
	return ApiPaymentLinksExpireRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksList200ResponseDataInner
// Deprecated
func (a *PaymentLinksAPIService) PaymentLinksExpireExecute(r ApiPaymentLinksExpireRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksList200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksExpire")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}/expire"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
	localVarPostBody = r.paymentLinksExpireRequest
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
			var v PaymentLinksList400Response
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
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksGetRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
}

func (r ApiPaymentLinksGetRequest) Execute() (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	return r.ApiService.PaymentLinksGetExecute(r)
}

/*
PaymentLinksGet Get payment link by ID

Retrieve a payment link by its UUID or public_id

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksGetRequest
*/
func (a *PaymentLinksAPIService) PaymentLinksGet(ctx context.Context, id string) ApiPaymentLinksGetRequest {
	return ApiPaymentLinksGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksList200ResponseDataInner
func (a *PaymentLinksAPIService) PaymentLinksGetExecute(r ApiPaymentLinksGetRequest) (*PaymentLinksList200ResponseDataInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksList200ResponseDataInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPaymentLinksListRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	limit *int32
	offset *int32
	status *string
	mode *string
	search *string
}

func (r ApiPaymentLinksListRequest) Limit(limit int32) ApiPaymentLinksListRequest {
	r.limit = &limit
	return r
}

func (r ApiPaymentLinksListRequest) Offset(offset int32) ApiPaymentLinksListRequest {
	r.offset = &offset
	return r
}

func (r ApiPaymentLinksListRequest) Status(status string) ApiPaymentLinksListRequest {
	r.status = &status
	return r
}

func (r ApiPaymentLinksListRequest) Mode(mode string) ApiPaymentLinksListRequest {
	r.mode = &mode
	return r
}

func (r ApiPaymentLinksListRequest) Search(search string) ApiPaymentLinksListRequest {
	r.search = &search
	return r
}

func (r ApiPaymentLinksListRequest) Execute() (*PaymentLinksList200Response, *http.Response, error) {
	return r.ApiService.PaymentLinksListExecute(r)
}

/*
PaymentLinksList List payment links

List payment links for the authenticated merchant with pagination and filters

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `PATCH /payment-links/{id}/status` — Update payment link status
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaymentLinksListRequest
*/
func (a *PaymentLinksAPIService) PaymentLinksList(ctx context.Context) ApiPaymentLinksListRequest {
	return ApiPaymentLinksListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentLinksList200Response
func (a *PaymentLinksAPIService) PaymentLinksListExecute(r ApiPaymentLinksListRequest) (*PaymentLinksList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
        var defaultValue int32 = 50
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
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "form", "")
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
			var v PaymentLinksList400Response
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

type ApiPaymentLinksUpdateRequest struct {
	ctx context.Context
	ApiService PaymentLinksAPI
	id string
	paymentLinksUpdateRequest *PaymentLinksUpdateRequest
}

func (r ApiPaymentLinksUpdateRequest) PaymentLinksUpdateRequest(paymentLinksUpdateRequest PaymentLinksUpdateRequest) ApiPaymentLinksUpdateRequest {
	r.paymentLinksUpdateRequest = &paymentLinksUpdateRequest
	return r
}

func (r ApiPaymentLinksUpdateRequest) Execute() (*PaymentLinksUpdate200Response, *http.Response, error) {
	return r.ApiService.PaymentLinksUpdateExecute(r)
}

/*
PaymentLinksUpdate Update payment link status

Change the status of a payment link. Active links accept payments, inactive links are temporarily disabled, and archived links are permanently disabled and cannot be reactivated.

---

**Related endpoints**

- `POST /payment-links` — Create a new payment link
- `GET /payment-links` — List payment links
- `GET /payment-links/{id}` — Get payment link by ID
- `POST /payment-links/{id}/expire` — Expire a payment link (deprecated)
- `POST /payment-links/{id}/deactivate` — Deactivate a payment link (deprecated)
- `POST /payment-links/{id}/activate` — Activate a payment link (deprecated)
- `POST /payment-links/{id}/archive` — Archive a payment link (deprecated)

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment link UUID or public_id
 @return ApiPaymentLinksUpdateRequest
*/
func (a *PaymentLinksAPIService) PaymentLinksUpdate(ctx context.Context, id string) ApiPaymentLinksUpdateRequest {
	return ApiPaymentLinksUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentLinksUpdate200Response
func (a *PaymentLinksAPIService) PaymentLinksUpdateExecute(r ApiPaymentLinksUpdateRequest) (*PaymentLinksUpdate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentLinksUpdate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentLinksAPIService.PaymentLinksUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-links/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 1 {
		return localVarReturnValue, nil, reportError("id must have at least 1 elements")
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
	localVarPostBody = r.paymentLinksUpdateRequest
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
			var v PaymentLinksList400Response
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
			var v PaymentLinksList400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
