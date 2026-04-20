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


type SubscriptionsAPI interface {

	/*
	SubscriptionsCancel Cancel subscription

	Cancel a subscription either immediately or at the end of the current billing period.

**Modes:**
- `immediately`: Status set to "canceled", access revoked now
- `period_end`: cancelAtPeriodEnd flag set, access continues until period end

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Subscription UUID
	@return ApiSubscriptionsCancelRequest
	*/
	SubscriptionsCancel(ctx context.Context, id string) ApiSubscriptionsCancelRequest

	// SubscriptionsCancelExecute executes the request
	//  @return SubscriptionCancelSubscriptionResponse
	SubscriptionsCancelExecute(r ApiSubscriptionsCancelRequest) (*SubscriptionCancelSubscriptionResponse, *http.Response, error)

	/*
	SubscriptionsCreate Create a new subscription

	Create a new subscription for a customer

---

**Related endpoints**

- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubscriptionsCreateRequest
	*/
	SubscriptionsCreate(ctx context.Context) ApiSubscriptionsCreateRequest

	// SubscriptionsCreateExecute executes the request
	//  @return SubscriptionCreateResponse
	SubscriptionsCreateExecute(r ApiSubscriptionsCreateRequest) (*SubscriptionCreateResponse, *http.Response, error)

	/*
	SubscriptionsGet Get subscription by ID

	Retrieve detailed information about a specific subscription

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Subscription UUID
	@return ApiSubscriptionsGetRequest
	*/
	SubscriptionsGet(ctx context.Context, id string) ApiSubscriptionsGetRequest

	// SubscriptionsGetExecute executes the request
	//  @return SubscriptionRetrieveResponse
	SubscriptionsGetExecute(r ApiSubscriptionsGetRequest) (*SubscriptionRetrieveResponse, *http.Response, error)

	/*
	SubscriptionsList List subscriptions

	Retrieve a paginated list of subscriptions with optional filters

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubscriptionsListRequest
	*/
	SubscriptionsList(ctx context.Context) ApiSubscriptionsListRequest

	// SubscriptionsListExecute executes the request
	//  @return SubscriptionListResponse
	SubscriptionsListExecute(r ApiSubscriptionsListRequest) (*SubscriptionListResponse, *http.Response, error)

	/*
	SubscriptionsUpdate Update subscription details

	Update an existing subscription's properties

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Subscription UUID
	@return ApiSubscriptionsUpdateRequest
	*/
	SubscriptionsUpdate(ctx context.Context, id string) ApiSubscriptionsUpdateRequest

	// SubscriptionsUpdateExecute executes the request
	//  @return SubscriptionUpdateResponse
	SubscriptionsUpdateExecute(r ApiSubscriptionsUpdateRequest) (*SubscriptionUpdateResponse, *http.Response, error)
}

// SubscriptionsAPIService SubscriptionsAPI service
type SubscriptionsAPIService service

type ApiSubscriptionsCancelRequest struct {
	ctx context.Context
	ApiService SubscriptionsAPI
	id string
	cancelSubscriptionInput *CancelSubscriptionInput
}

func (r ApiSubscriptionsCancelRequest) CancelSubscriptionInput(cancelSubscriptionInput CancelSubscriptionInput) ApiSubscriptionsCancelRequest {
	r.cancelSubscriptionInput = &cancelSubscriptionInput
	return r
}

func (r ApiSubscriptionsCancelRequest) Execute() (*SubscriptionCancelSubscriptionResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsCancelExecute(r)
}

/*
SubscriptionsCancel Cancel subscription

Cancel a subscription either immediately or at the end of the current billing period.

**Modes:**
- `immediately`: Status set to "canceled", access revoked now
- `period_end`: cancelAtPeriodEnd flag set, access continues until period end

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Subscription UUID
 @return ApiSubscriptionsCancelRequest
*/
func (a *SubscriptionsAPIService) SubscriptionsCancel(ctx context.Context, id string) ApiSubscriptionsCancelRequest {
	return ApiSubscriptionsCancelRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionCancelSubscriptionResponse
func (a *SubscriptionsAPIService) SubscriptionsCancelExecute(r ApiSubscriptionsCancelRequest) (*SubscriptionCancelSubscriptionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionCancelSubscriptionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.SubscriptionsCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{id}/cancel"
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
	localVarPostBody = r.cancelSubscriptionInput
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

type ApiSubscriptionsCreateRequest struct {
	ctx context.Context
	ApiService SubscriptionsAPI
	subscriptionsCreateRequest *SubscriptionsCreateRequest
}

func (r ApiSubscriptionsCreateRequest) SubscriptionsCreateRequest(subscriptionsCreateRequest SubscriptionsCreateRequest) ApiSubscriptionsCreateRequest {
	r.subscriptionsCreateRequest = &subscriptionsCreateRequest
	return r
}

func (r ApiSubscriptionsCreateRequest) Execute() (*SubscriptionCreateResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsCreateExecute(r)
}

/*
SubscriptionsCreate Create a new subscription

Create a new subscription for a customer

---

**Related endpoints**

- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscriptionsCreateRequest
*/
func (a *SubscriptionsAPIService) SubscriptionsCreate(ctx context.Context) ApiSubscriptionsCreateRequest {
	return ApiSubscriptionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriptionCreateResponse
func (a *SubscriptionsAPIService) SubscriptionsCreateExecute(r ApiSubscriptionsCreateRequest) (*SubscriptionCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.SubscriptionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions"

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
	localVarPostBody = r.subscriptionsCreateRequest
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

type ApiSubscriptionsGetRequest struct {
	ctx context.Context
	ApiService SubscriptionsAPI
	id string
}

func (r ApiSubscriptionsGetRequest) Execute() (*SubscriptionRetrieveResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsGetExecute(r)
}

/*
SubscriptionsGet Get subscription by ID

Retrieve detailed information about a specific subscription

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Subscription UUID
 @return ApiSubscriptionsGetRequest
*/
func (a *SubscriptionsAPIService) SubscriptionsGet(ctx context.Context, id string) ApiSubscriptionsGetRequest {
	return ApiSubscriptionsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionRetrieveResponse
func (a *SubscriptionsAPIService) SubscriptionsGetExecute(r ApiSubscriptionsGetRequest) (*SubscriptionRetrieveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionRetrieveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.SubscriptionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{id}"
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

type ApiSubscriptionsListRequest struct {
	ctx context.Context
	ApiService SubscriptionsAPI
	page *int32
	limit *int32
	status *string
	customerId *string
}

// Page number
func (r ApiSubscriptionsListRequest) Page(page int32) ApiSubscriptionsListRequest {
	r.page = &page
	return r
}

// Maximum number of results (1-100)
func (r ApiSubscriptionsListRequest) Limit(limit int32) ApiSubscriptionsListRequest {
	r.limit = &limit
	return r
}

// Filter by subscription status
func (r ApiSubscriptionsListRequest) Status(status string) ApiSubscriptionsListRequest {
	r.status = &status
	return r
}

// Filter by customer UUID
func (r ApiSubscriptionsListRequest) CustomerId(customerId string) ApiSubscriptionsListRequest {
	r.customerId = &customerId
	return r
}

func (r ApiSubscriptionsListRequest) Execute() (*SubscriptionListResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsListExecute(r)
}

/*
SubscriptionsList List subscriptions

Retrieve a paginated list of subscriptions with optional filters

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions/{id}` — Get subscription by ID
- `PATCH /subscriptions/{id}` — Update subscription details
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscriptionsListRequest
*/
func (a *SubscriptionsAPIService) SubscriptionsList(ctx context.Context) ApiSubscriptionsListRequest {
	return ApiSubscriptionsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriptionListResponse
func (a *SubscriptionsAPIService) SubscriptionsListExecute(r ApiSubscriptionsListRequest) (*SubscriptionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.SubscriptionsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions"

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
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.customerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerId", r.customerId, "form", "")
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

type ApiSubscriptionsUpdateRequest struct {
	ctx context.Context
	ApiService SubscriptionsAPI
	id string
	subscriptionsUpdateRequest *SubscriptionsUpdateRequest
}

func (r ApiSubscriptionsUpdateRequest) SubscriptionsUpdateRequest(subscriptionsUpdateRequest SubscriptionsUpdateRequest) ApiSubscriptionsUpdateRequest {
	r.subscriptionsUpdateRequest = &subscriptionsUpdateRequest
	return r
}

func (r ApiSubscriptionsUpdateRequest) Execute() (*SubscriptionUpdateResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsUpdateExecute(r)
}

/*
SubscriptionsUpdate Update subscription details

Update an existing subscription's properties

---

**Related endpoints**

- `POST /subscriptions` — Create a new subscription
- `GET /subscriptions` — List subscriptions
- `GET /subscriptions/{id}` — Get subscription by ID
- `DELETE /subscriptions/{id}` — Delete subscription
- `POST /subscriptions/{id}/change-plan` — Change subscription plan
- `POST /subscriptions/{id}/change-quantity` — Change subscription quantity
- `POST /subscriptions/{id}/preview-renewal` — Preview subscription renewal
- `POST /subscriptions/{id}/pause` — Pause subscription

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Subscription UUID
 @return ApiSubscriptionsUpdateRequest
*/
func (a *SubscriptionsAPIService) SubscriptionsUpdate(ctx context.Context, id string) ApiSubscriptionsUpdateRequest {
	return ApiSubscriptionsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionUpdateResponse
func (a *SubscriptionsAPIService) SubscriptionsUpdateExecute(r ApiSubscriptionsUpdateRequest) (*SubscriptionUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsAPIService.SubscriptionsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscriptions/{id}"
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
	localVarPostBody = r.subscriptionsUpdateRequest
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
