/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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


type InvoicesAPI interface {

	/*
	InvoicesCreate Create invoice

	Create a new invoice.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

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
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

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
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

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
	InvoicesSend Send an invoice

	Sends an invoice to the customer via the specified channel (email, SMS, or WhatsApp). Invoice must be approved first.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

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
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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

type ApiInvoicesCreateRequest struct {
	ctx context.Context
	ApiService InvoicesAPI
	invoicesCreateRequest *InvoicesCreateRequest
}

// Invoice creation details
func (r ApiInvoicesCreateRequest) InvoicesCreateRequest(invoicesCreateRequest InvoicesCreateRequest) ApiInvoicesCreateRequest {
	r.invoicesCreateRequest = &invoicesCreateRequest
	return r
}

func (r ApiInvoicesCreateRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesCreateExecute(r)
}

/*
InvoicesCreate Create invoice

Create a new invoice.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInvoicesCreateRequest
*/
func (a *InvoicesAPIService) InvoicesCreate(ctx context.Context) ApiInvoicesCreateRequest {
	return ApiInvoicesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesCreateExecute(r ApiInvoicesCreateRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	id string
	invoicesFinalizeRequest *InvoicesFinalizeRequest
}

// Finalization options
func (r ApiInvoicesFinalizeRequest) InvoicesFinalizeRequest(invoicesFinalizeRequest InvoicesFinalizeRequest) ApiInvoicesFinalizeRequest {
	r.invoicesFinalizeRequest = &invoicesFinalizeRequest
	return r
}

func (r ApiInvoicesFinalizeRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesFinalizeExecute(r)
}

/*
InvoicesFinalize Finalize an invoice

Finalizes a draft invoice, locking it for payment. Assigns invoice number and generates public token. After finalization, financial fields become immutable.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesFinalizeExecute(r ApiInvoicesFinalizeRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	id string
}

func (r ApiInvoicesGetRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesGetExecute(r)
}

/*
InvoicesGet Get invoice

Get a single invoice by ID.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiInvoicesGetRequest
*/
func (a *InvoicesAPIService) InvoicesGet(ctx context.Context, id string) ApiInvoicesGetRequest {
	return ApiInvoicesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesGetExecute(r ApiInvoicesGetRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	status *string
	customerId *string
	limit *int32
	offset *int32
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

func (r ApiInvoicesListRequest) Execute() (*InvoiceListResponse, *http.Response, error) {
	return r.ApiService.InvoicesListExecute(r)
}

/*
InvoicesList List invoices

List invoices with pagination and filtering.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInvoicesListRequest
*/
func (a *InvoicesAPIService) InvoicesList(ctx context.Context) ApiInvoicesListRequest {
	return ApiInvoicesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvoiceListResponse
func (a *InvoicesAPIService) InvoicesListExecute(r ApiInvoicesListRequest) (*InvoiceListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceListResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	id string
	invoicesSendRequest *InvoicesSendRequest
}

// Send options
func (r ApiInvoicesSendRequest) InvoicesSendRequest(invoicesSendRequest InvoicesSendRequest) ApiInvoicesSendRequest {
	r.invoicesSendRequest = &invoicesSendRequest
	return r
}

func (r ApiInvoicesSendRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesSendExecute(r)
}

/*
InvoicesSend Send an invoice

Sends an invoice to the customer via the specified channel (email, SMS, or WhatsApp). Invoice must be approved first.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesSendExecute(r ApiInvoicesSendRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	id string
	invoicesUpdateRequest *InvoicesUpdateRequest
}

// Invoice update details
func (r ApiInvoicesUpdateRequest) InvoicesUpdateRequest(invoicesUpdateRequest InvoicesUpdateRequest) ApiInvoicesUpdateRequest {
	r.invoicesUpdateRequest = &invoicesUpdateRequest
	return r
}

func (r ApiInvoicesUpdateRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesUpdateExecute(r)
}

/*
InvoicesUpdate Update invoice

Update an existing invoice.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice
- `GET /invoices/{id}/comments` — List invoice comments

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiInvoicesUpdateRequest
*/
func (a *InvoicesAPIService) InvoicesUpdate(ctx context.Context, id string) ApiInvoicesUpdateRequest {
	return ApiInvoicesUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesUpdateExecute(r ApiInvoicesUpdateRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
	ctx context.Context
	ApiService InvoicesAPI
	id string
	invoicesVoidRequest *InvoicesVoidRequest
}

// Void details
func (r ApiInvoicesVoidRequest) InvoicesVoidRequest(invoicesVoidRequest InvoicesVoidRequest) ApiInvoicesVoidRequest {
	r.invoicesVoidRequest = &invoicesVoidRequest
	return r
}

func (r ApiInvoicesVoidRequest) Execute() (*InvoiceResponse, *http.Response, error) {
	return r.ApiService.InvoicesVoidExecute(r)
}

/*
InvoicesVoid Void an invoice

Voids an invoice. Only invoices without recorded payments can be voided. Use refund instead for paid invoices.

---

**Related endpoints**

- `PUT /invoices/external/batch` — Batch upsert invoices by external ID
- `GET /invoices` — List invoices
- `POST /invoices` — Create invoice
- `GET /invoices/{id}` — Get invoice
- `PATCH /invoices/{id}` — Update invoice
- `DELETE /invoices/{id}` — Delete invoice
- `POST /invoices/{id}/refund` — Refund invoice
- `POST /invoices/{id}/reject` — Reject invoice

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
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InvoiceResponse
func (a *InvoicesAPIService) InvoicesVoidExecute(r ApiInvoicesVoidRequest) (*InvoiceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceResponse
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
