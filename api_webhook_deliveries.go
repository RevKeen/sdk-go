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


type WebhookDeliveriesAPI interface {

	/*
	WebhookDeliveriesGet Retrieve a webhook delivery

	Retrieve the details of a single webhook delivery including its last attempt outcome.

---

**Related endpoints**

- `GET /webhook-deliveries` — List webhook deliveries
- `POST /webhook-deliveries/{id}/retry` — Retry a webhook delivery

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook delivery ID
	@return ApiWebhookDeliveriesGetRequest
	*/
	WebhookDeliveriesGet(ctx context.Context, id string) ApiWebhookDeliveriesGetRequest

	// WebhookDeliveriesGetExecute executes the request
	//  @return WebhookDeliveryResponse
	WebhookDeliveriesGetExecute(r ApiWebhookDeliveriesGetRequest) (*WebhookDeliveryResponse, *http.Response, error)

	/*
	WebhookDeliveriesList List webhook deliveries

	List individual webhook delivery attempts across all endpoints for the authenticated merchant. Use filters to scope to a specific endpoint, a specific event, or a specific delivery status. Results are returned in reverse chronological order.

---

**Related endpoints**

- `GET /webhook-deliveries/{id}` — Retrieve a webhook delivery
- `POST /webhook-deliveries/{id}/retry` — Retry a webhook delivery

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhookDeliveriesListRequest
	*/
	WebhookDeliveriesList(ctx context.Context) ApiWebhookDeliveriesListRequest

	// WebhookDeliveriesListExecute executes the request
	//  @return WebhookDeliveryListResponse
	WebhookDeliveriesListExecute(r ApiWebhookDeliveriesListRequest) (*WebhookDeliveryListResponse, *http.Response, error)

	/*
	WebhookDeliveriesRetry Retry a webhook delivery

	Queue an immediate retry of this delivery. The delivery is marked `pending` and `next_retry_at` is set to now; the dispatcher picks it up on its next tick. Calling retry on an already-pending delivery just advances its retry time. Dead-lettered deliveries can be retried once; succeeded deliveries cannot.

---

**Related endpoints**

- `GET /webhook-deliveries` — List webhook deliveries
- `GET /webhook-deliveries/{id}` — Retrieve a webhook delivery

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Webhook delivery ID
	@return ApiWebhookDeliveriesRetryRequest
	*/
	WebhookDeliveriesRetry(ctx context.Context, id string) ApiWebhookDeliveriesRetryRequest

	// WebhookDeliveriesRetryExecute executes the request
	//  @return WebhookDeliveryRetryResponse
	WebhookDeliveriesRetryExecute(r ApiWebhookDeliveriesRetryRequest) (*WebhookDeliveryRetryResponse, *http.Response, error)
}

// WebhookDeliveriesAPIService WebhookDeliveriesAPI service
type WebhookDeliveriesAPIService service

type ApiWebhookDeliveriesGetRequest struct {
	ctx context.Context
	ApiService WebhookDeliveriesAPI
	id string
}

func (r ApiWebhookDeliveriesGetRequest) Execute() (*WebhookDeliveryResponse, *http.Response, error) {
	return r.ApiService.WebhookDeliveriesGetExecute(r)
}

/*
WebhookDeliveriesGet Retrieve a webhook delivery

Retrieve the details of a single webhook delivery including its last attempt outcome.

---

**Related endpoints**

- `GET /webhook-deliveries` — List webhook deliveries
- `POST /webhook-deliveries/{id}/retry` — Retry a webhook delivery

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook delivery ID
 @return ApiWebhookDeliveriesGetRequest
*/
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesGet(ctx context.Context, id string) ApiWebhookDeliveriesGetRequest {
	return ApiWebhookDeliveriesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookDeliveryResponse
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesGetExecute(r ApiWebhookDeliveriesGetRequest) (*WebhookDeliveryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookDeliveryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveriesAPIService.WebhookDeliveriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-deliveries/{id}"
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v WebhookDeliveryErrorResponse
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

type ApiWebhookDeliveriesListRequest struct {
	ctx context.Context
	ApiService WebhookDeliveriesAPI
	endpointId *string
	eventId *string
	status *string
	limit *int32
	startingAfter *string
	endingBefore *string
}

// Filter by webhook endpoint ID
func (r ApiWebhookDeliveriesListRequest) EndpointId(endpointId string) ApiWebhookDeliveriesListRequest {
	r.endpointId = &endpointId
	return r
}

// Filter by source event ID
func (r ApiWebhookDeliveriesListRequest) EventId(eventId string) ApiWebhookDeliveriesListRequest {
	r.eventId = &eventId
	return r
}

// Filter by delivery status
func (r ApiWebhookDeliveriesListRequest) Status(status string) ApiWebhookDeliveriesListRequest {
	r.status = &status
	return r
}

func (r ApiWebhookDeliveriesListRequest) Limit(limit int32) ApiWebhookDeliveriesListRequest {
	r.limit = &limit
	return r
}

// Cursor — return deliveries created before the row with this ID.
func (r ApiWebhookDeliveriesListRequest) StartingAfter(startingAfter string) ApiWebhookDeliveriesListRequest {
	r.startingAfter = &startingAfter
	return r
}

// Cursor — return deliveries created after the row with this ID.
func (r ApiWebhookDeliveriesListRequest) EndingBefore(endingBefore string) ApiWebhookDeliveriesListRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r ApiWebhookDeliveriesListRequest) Execute() (*WebhookDeliveryListResponse, *http.Response, error) {
	return r.ApiService.WebhookDeliveriesListExecute(r)
}

/*
WebhookDeliveriesList List webhook deliveries

List individual webhook delivery attempts across all endpoints for the authenticated merchant. Use filters to scope to a specific endpoint, a specific event, or a specific delivery status. Results are returned in reverse chronological order.

---

**Related endpoints**

- `GET /webhook-deliveries/{id}` — Retrieve a webhook delivery
- `POST /webhook-deliveries/{id}/retry` — Retry a webhook delivery

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhookDeliveriesListRequest
*/
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesList(ctx context.Context) ApiWebhookDeliveriesListRequest {
	return ApiWebhookDeliveriesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WebhookDeliveryListResponse
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesListExecute(r ApiWebhookDeliveriesListRequest) (*WebhookDeliveryListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookDeliveryListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveriesAPIService.WebhookDeliveriesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-deliveries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.endpointId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endpoint_id", r.endpointId, "form", "")
	}
	if r.eventId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "event_id", r.eventId, "form", "")
	}
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
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "starting_after", r.startingAfter, "form", "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ending_before", r.endingBefore, "form", "")
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
			var v WebhookDeliveryErrorResponse
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

type ApiWebhookDeliveriesRetryRequest struct {
	ctx context.Context
	ApiService WebhookDeliveriesAPI
	id string
}

func (r ApiWebhookDeliveriesRetryRequest) Execute() (*WebhookDeliveryRetryResponse, *http.Response, error) {
	return r.ApiService.WebhookDeliveriesRetryExecute(r)
}

/*
WebhookDeliveriesRetry Retry a webhook delivery

Queue an immediate retry of this delivery. The delivery is marked `pending` and `next_retry_at` is set to now; the dispatcher picks it up on its next tick. Calling retry on an already-pending delivery just advances its retry time. Dead-lettered deliveries can be retried once; succeeded deliveries cannot.

---

**Related endpoints**

- `GET /webhook-deliveries` — List webhook deliveries
- `GET /webhook-deliveries/{id}` — Retrieve a webhook delivery

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Webhook delivery ID
 @return ApiWebhookDeliveriesRetryRequest
*/
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesRetry(ctx context.Context, id string) ApiWebhookDeliveriesRetryRequest {
	return ApiWebhookDeliveriesRetryRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return WebhookDeliveryRetryResponse
func (a *WebhookDeliveriesAPIService) WebhookDeliveriesRetryExecute(r ApiWebhookDeliveriesRetryRequest) (*WebhookDeliveryRetryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WebhookDeliveryRetryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhookDeliveriesAPIService.WebhookDeliveriesRetry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhook-deliveries/{id}/retry"
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v WebhookDeliveryErrorResponse
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
			var v WebhookDeliveryErrorResponse
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
