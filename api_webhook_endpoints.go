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


type WebhookEndpointsAPI interface {

	/*
	WebhookEndpointsCreate Create webhook endpoint

	Create a webhook endpoint to receive event notifications at the given URL. Returns the endpoint with its signing secret — store the secret securely, it is only shown once.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhookEndpointsCreateRequest
	*/
	WebhookEndpointsCreate(ctx context.Context) ApiWebhookEndpointsCreateRequest

	// WebhookEndpointsCreateExecute executes the request
	//  @return WebhookEndpointsCreate201Response
	WebhookEndpointsCreateExecute(r ApiWebhookEndpointsCreateRequest) (*WebhookEndpointsCreate201Response, *http.Response, error)

	/*
	WebhookEndpointsDelete Delete webhook endpoint

	Permanently delete a webhook endpoint and all its event subscriptions.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook endpoint ID
	@return ApiWebhookEndpointsDeleteRequest
	*/
	WebhookEndpointsDelete(ctx context.Context, id string) ApiWebhookEndpointsDeleteRequest

	// WebhookEndpointsDeleteExecute executes the request
	//  @return WebhookEndpointsDelete200Response
	WebhookEndpointsDeleteExecute(r ApiWebhookEndpointsDeleteRequest) (*WebhookEndpointsDelete200Response, *http.Response, error)

	/*
	WebhookEndpointsGet Get webhook endpoint

	Retrieve a single webhook endpoint by ID, including delivery statistics.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook endpoint ID
	@return ApiWebhookEndpointsGetRequest
	*/
	WebhookEndpointsGet(ctx context.Context, id string) ApiWebhookEndpointsGetRequest

	// WebhookEndpointsGetExecute executes the request
	//  @return WebhookEndpointsGet200Response
	WebhookEndpointsGetExecute(r ApiWebhookEndpointsGetRequest) (*WebhookEndpointsGet200Response, *http.Response, error)

	/*
	WebhookEndpointsList List webhook endpoints

	List all webhook endpoints for the authenticated merchant. Supports filtering by status and pagination.

---

**Related endpoints**

- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhookEndpointsListRequest
	*/
	WebhookEndpointsList(ctx context.Context) ApiWebhookEndpointsListRequest

	// WebhookEndpointsListExecute executes the request
	//  @return WebhookEndpointsList200Response
	WebhookEndpointsListExecute(r ApiWebhookEndpointsListRequest) (*WebhookEndpointsList200Response, *http.Response, error)

	/*
	WebhookEndpointsRotateSecret Rotate signing secret

	Generate a new signing secret for a webhook endpoint. The old secret is immediately invalidated. Store the new secret securely — it is only returned in this response.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook endpoint ID
	@return ApiWebhookEndpointsRotateSecretRequest
	*/
	WebhookEndpointsRotateSecret(ctx context.Context, id string) ApiWebhookEndpointsRotateSecretRequest

	// WebhookEndpointsRotateSecretExecute executes the request
	//  @return WebhookEndpointsCreate201Response
	WebhookEndpointsRotateSecretExecute(r ApiWebhookEndpointsRotateSecretRequest) (*WebhookEndpointsCreate201Response, *http.Response, error)

	/*
	WebhookEndpointsUpdate Update webhook endpoint

	Update a webhook endpoint's URL, subscribed events, description, or enabled status.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook endpoint ID
	@return ApiWebhookEndpointsUpdateRequest
	*/
	WebhookEndpointsUpdate(ctx context.Context, id string) ApiWebhookEndpointsUpdateRequest

	// WebhookEndpointsUpdateExecute executes the request
	//  @return WebhookEndpointsGet200Response
	WebhookEndpointsUpdateExecute(r ApiWebhookEndpointsUpdateRequest) (*WebhookEndpointsGet200Response, *http.Response, error)
}

// WebhookEndpointsAPIService WebhookEndpointsAPI service
type WebhookEndpointsAPIService service

type ApiWebhookEndpointsCreateRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	webhookEndpointsCreateRequest *WebhookEndpointsCreateRequest
}

func (r ApiWebhookEndpointsCreateRequest) WebhookEndpointsCreateRequest(webhookEndpointsCreateRequest WebhookEndpointsCreateRequest) ApiWebhookEndpointsCreateRequest {
	r.webhookEndpointsCreateRequest = &webhookEndpointsCreateRequest
	return r
}

func (r ApiWebhookEndpointsCreateRequest) Execute() (*WebhookEndpointsCreate201Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsCreateExecute(r)
}

/*
WebhookEndpointsCreate Create webhook endpoint

Create a webhook endpoint to receive event notifications at the given URL. Returns the endpoint with its signing secret — store the secret securely, it is only shown once.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookEndpointsCreateRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsCreate(ctx context.Context) ApiWebhookEndpointsCreateRequest {
	return ApiWebhookEndpointsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookEndpointsCreate201Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsCreateExecute(r ApiWebhookEndpointsCreateRequest) (*WebhookEndpointsCreate201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsCreate201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints"

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
	localVarPostBody = r.webhookEndpointsCreateRequest
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

type ApiWebhookEndpointsDeleteRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	id string
}

func (r ApiWebhookEndpointsDeleteRequest) Execute() (*WebhookEndpointsDelete200Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsDeleteExecute(r)
}

/*
WebhookEndpointsDelete Delete webhook endpoint

Permanently delete a webhook endpoint and all its event subscriptions.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook endpoint ID
 @return ApiWebhookEndpointsDeleteRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsDelete(ctx context.Context, id string) ApiWebhookEndpointsDeleteRequest {
	return ApiWebhookEndpointsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookEndpointsDelete200Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsDeleteExecute(r ApiWebhookEndpointsDeleteRequest) (*WebhookEndpointsDelete200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsDelete200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints/{id}"
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

type ApiWebhookEndpointsGetRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	id string
}

func (r ApiWebhookEndpointsGetRequest) Execute() (*WebhookEndpointsGet200Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsGetExecute(r)
}

/*
WebhookEndpointsGet Get webhook endpoint

Retrieve a single webhook endpoint by ID, including delivery statistics.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook endpoint ID
 @return ApiWebhookEndpointsGetRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsGet(ctx context.Context, id string) ApiWebhookEndpointsGetRequest {
	return ApiWebhookEndpointsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookEndpointsGet200Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsGetExecute(r ApiWebhookEndpointsGetRequest) (*WebhookEndpointsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints/{id}"
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

type ApiWebhookEndpointsListRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	status *string
	limit *int32
	offset *int32
}

// Filter by endpoint status
func (r ApiWebhookEndpointsListRequest) Status(status string) ApiWebhookEndpointsListRequest {
	r.status = &status
	return r
}

// Maximum number of results to return (1-100, default 20)
func (r ApiWebhookEndpointsListRequest) Limit(limit int32) ApiWebhookEndpointsListRequest {
	r.limit = &limit
	return r
}

// Number of results to skip for pagination
func (r ApiWebhookEndpointsListRequest) Offset(offset int32) ApiWebhookEndpointsListRequest {
	r.offset = &offset
	return r
}

func (r ApiWebhookEndpointsListRequest) Execute() (*WebhookEndpointsList200Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsListExecute(r)
}

/*
WebhookEndpointsList List webhook endpoints

List all webhook endpoints for the authenticated merchant. Supports filtering by status and pagination.

---

**Related endpoints**

- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookEndpointsListRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsList(ctx context.Context) ApiWebhookEndpointsListRequest {
	return ApiWebhookEndpointsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookEndpointsList200Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsListExecute(r ApiWebhookEndpointsListRequest) (*WebhookEndpointsList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
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

type ApiWebhookEndpointsRotateSecretRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	id string
}

func (r ApiWebhookEndpointsRotateSecretRequest) Execute() (*WebhookEndpointsCreate201Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsRotateSecretExecute(r)
}

/*
WebhookEndpointsRotateSecret Rotate signing secret

Generate a new signing secret for a webhook endpoint. The old secret is immediately invalidated. Store the new secret securely — it is only returned in this response.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `PATCH /webhook-endpoints/{id}` — Update webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook endpoint ID
 @return ApiWebhookEndpointsRotateSecretRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsRotateSecret(ctx context.Context, id string) ApiWebhookEndpointsRotateSecretRequest {
	return ApiWebhookEndpointsRotateSecretRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookEndpointsCreate201Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsRotateSecretExecute(r ApiWebhookEndpointsRotateSecretRequest) (*WebhookEndpointsCreate201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsCreate201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsRotateSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints/{id}/rotate-secret"
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

type ApiWebhookEndpointsUpdateRequest struct {
	ctx context.Context
	ApiService WebhookEndpointsAPI
	id string
	webhookEndpointsUpdateRequest *WebhookEndpointsUpdateRequest
}

func (r ApiWebhookEndpointsUpdateRequest) WebhookEndpointsUpdateRequest(webhookEndpointsUpdateRequest WebhookEndpointsUpdateRequest) ApiWebhookEndpointsUpdateRequest {
	r.webhookEndpointsUpdateRequest = &webhookEndpointsUpdateRequest
	return r
}

func (r ApiWebhookEndpointsUpdateRequest) Execute() (*WebhookEndpointsGet200Response, *http.Response, error) {
	return r.ApiService.WebhookEndpointsUpdateExecute(r)
}

/*
WebhookEndpointsUpdate Update webhook endpoint

Update a webhook endpoint's URL, subscribed events, description, or enabled status.

---

**Related endpoints**

- `GET /webhook-endpoints` — List webhook endpoints
- `POST /webhook-endpoints` — Create webhook endpoint
- `GET /webhook-endpoints/{id}` — Get webhook endpoint
- `DELETE /webhook-endpoints/{id}` — Delete webhook endpoint
- `POST /webhook-endpoints/{id}/rotate-secret` — Rotate signing secret

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook endpoint ID
 @return ApiWebhookEndpointsUpdateRequest
*/
func (a *WebhookEndpointsAPIService) WebhookEndpointsUpdate(ctx context.Context, id string) ApiWebhookEndpointsUpdateRequest {
	return ApiWebhookEndpointsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookEndpointsGet200Response
func (a *WebhookEndpointsAPIService) WebhookEndpointsUpdateExecute(r ApiWebhookEndpointsUpdateRequest) (*WebhookEndpointsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookEndpointsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookEndpointsAPIService.WebhookEndpointsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-endpoints/{id}"
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
	localVarPostBody = r.webhookEndpointsUpdateRequest
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
