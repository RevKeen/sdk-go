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


type PaymentIntentsAPI interface {

	/*
	PaymentIntentsCancel Cancel a payment intent

	Cancels a payment intent. Cannot cancel if already succeeded or canceled.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment intent ID (pi_xxx)
	@return ApiPaymentIntentsCancelRequest
	*/
	PaymentIntentsCancel(ctx context.Context, id string) ApiPaymentIntentsCancelRequest

	// PaymentIntentsCancelExecute executes the request
	//  @return PaymentIntent
	PaymentIntentsCancelExecute(r ApiPaymentIntentsCancelRequest) (*PaymentIntent, *http.Response, error)

	/*
	PaymentIntentsCapture Capture a payment intent

	Captures a payment intent that was created with capture_method=manual.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment intent ID (pi_xxx)
	@return ApiPaymentIntentsCaptureRequest
	*/
	PaymentIntentsCapture(ctx context.Context, id string) ApiPaymentIntentsCaptureRequest

	// PaymentIntentsCaptureExecute executes the request
	//  @return PaymentIntent
	PaymentIntentsCaptureExecute(r ApiPaymentIntentsCaptureRequest) (*PaymentIntent, *http.Response, error)

	/*
	PaymentIntentsConfirm Confirm a payment intent

	Confirms the payment intent. May return requires_action if 3DS authentication is needed.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment intent ID (pi_xxx)
	@return ApiPaymentIntentsConfirmRequest
	*/
	PaymentIntentsConfirm(ctx context.Context, id string) ApiPaymentIntentsConfirmRequest

	// PaymentIntentsConfirmExecute executes the request
	//  @return PaymentIntent
	PaymentIntentsConfirmExecute(r ApiPaymentIntentsConfirmRequest) (*PaymentIntent, *http.Response, error)

	/*
	PaymentIntentsCreate Create a payment intent

	Creates a payment intent to orchestrate payment collection with support for 3DS/SCA authentication.

---

**Related endpoints**

- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPaymentIntentsCreateRequest
	*/
	PaymentIntentsCreate(ctx context.Context) ApiPaymentIntentsCreateRequest

	// PaymentIntentsCreateExecute executes the request
	//  @return PaymentIntent
	PaymentIntentsCreateExecute(r ApiPaymentIntentsCreateRequest) (*PaymentIntent, *http.Response, error)

	/*
	PaymentIntentsGet Retrieve a payment intent

	Retrieves a payment intent by its ID (pi_xxx).

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Payment intent ID (pi_xxx)
	@return ApiPaymentIntentsGetRequest
	*/
	PaymentIntentsGet(ctx context.Context, id string) ApiPaymentIntentsGetRequest

	// PaymentIntentsGetExecute executes the request
	//  @return PaymentIntent
	PaymentIntentsGetExecute(r ApiPaymentIntentsGetRequest) (*PaymentIntent, *http.Response, error)

	/*
	PaymentIntentsList List payment intents

	Returns a list of payment intents with optional filtering.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPaymentIntentsListRequest
	*/
	PaymentIntentsList(ctx context.Context) ApiPaymentIntentsListRequest

	// PaymentIntentsListExecute executes the request
	//  @return PaymentIntentListResponse
	PaymentIntentsListExecute(r ApiPaymentIntentsListRequest) (*PaymentIntentListResponse, *http.Response, error)
}

// PaymentIntentsAPIService PaymentIntentsAPI service
type PaymentIntentsAPIService service

type ApiPaymentIntentsCancelRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	id string
	cancelPaymentIntentRequest *CancelPaymentIntentRequest
}

func (r ApiPaymentIntentsCancelRequest) CancelPaymentIntentRequest(cancelPaymentIntentRequest CancelPaymentIntentRequest) ApiPaymentIntentsCancelRequest {
	r.cancelPaymentIntentRequest = &cancelPaymentIntentRequest
	return r
}

func (r ApiPaymentIntentsCancelRequest) Execute() (*PaymentIntent, *http.Response, error) {
	return r.ApiService.PaymentIntentsCancelExecute(r)
}

/*
PaymentIntentsCancel Cancel a payment intent

Cancels a payment intent. Cannot cancel if already succeeded or canceled.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment intent ID (pi_xxx)
 @return ApiPaymentIntentsCancelRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsCancel(ctx context.Context, id string) ApiPaymentIntentsCancelRequest {
	return ApiPaymentIntentsCancelRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentIntent
func (a *PaymentIntentsAPIService) PaymentIntentsCancelExecute(r ApiPaymentIntentsCancelRequest) (*PaymentIntent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{id}/cancel"
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
	localVarPostBody = r.cancelPaymentIntentRequest
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
			var v PaymentIntentErrorResponse
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
			var v PaymentIntentErrorResponse
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

type ApiPaymentIntentsCaptureRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	id string
	capturePaymentIntentRequest *CapturePaymentIntentRequest
}

func (r ApiPaymentIntentsCaptureRequest) CapturePaymentIntentRequest(capturePaymentIntentRequest CapturePaymentIntentRequest) ApiPaymentIntentsCaptureRequest {
	r.capturePaymentIntentRequest = &capturePaymentIntentRequest
	return r
}

func (r ApiPaymentIntentsCaptureRequest) Execute() (*PaymentIntent, *http.Response, error) {
	return r.ApiService.PaymentIntentsCaptureExecute(r)
}

/*
PaymentIntentsCapture Capture a payment intent

Captures a payment intent that was created with capture_method=manual.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment intent ID (pi_xxx)
 @return ApiPaymentIntentsCaptureRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsCapture(ctx context.Context, id string) ApiPaymentIntentsCaptureRequest {
	return ApiPaymentIntentsCaptureRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentIntent
func (a *PaymentIntentsAPIService) PaymentIntentsCaptureExecute(r ApiPaymentIntentsCaptureRequest) (*PaymentIntent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsCapture")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{id}/capture"
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
	localVarPostBody = r.capturePaymentIntentRequest
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
			var v PaymentIntentErrorResponse
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
			var v PaymentIntentErrorResponse
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

type ApiPaymentIntentsConfirmRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	id string
	confirmPaymentIntentRequest *ConfirmPaymentIntentRequest
}

func (r ApiPaymentIntentsConfirmRequest) ConfirmPaymentIntentRequest(confirmPaymentIntentRequest ConfirmPaymentIntentRequest) ApiPaymentIntentsConfirmRequest {
	r.confirmPaymentIntentRequest = &confirmPaymentIntentRequest
	return r
}

func (r ApiPaymentIntentsConfirmRequest) Execute() (*PaymentIntent, *http.Response, error) {
	return r.ApiService.PaymentIntentsConfirmExecute(r)
}

/*
PaymentIntentsConfirm Confirm a payment intent

Confirms the payment intent. May return requires_action if 3DS authentication is needed.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment intent ID (pi_xxx)
 @return ApiPaymentIntentsConfirmRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsConfirm(ctx context.Context, id string) ApiPaymentIntentsConfirmRequest {
	return ApiPaymentIntentsConfirmRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentIntent
func (a *PaymentIntentsAPIService) PaymentIntentsConfirmExecute(r ApiPaymentIntentsConfirmRequest) (*PaymentIntent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsConfirm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{id}/confirm"
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
	localVarPostBody = r.confirmPaymentIntentRequest
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
			var v PaymentIntentErrorResponse
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
			var v PaymentIntentErrorResponse
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

type ApiPaymentIntentsCreateRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	createPaymentIntentRequest *CreatePaymentIntentRequest
}

func (r ApiPaymentIntentsCreateRequest) CreatePaymentIntentRequest(createPaymentIntentRequest CreatePaymentIntentRequest) ApiPaymentIntentsCreateRequest {
	r.createPaymentIntentRequest = &createPaymentIntentRequest
	return r
}

func (r ApiPaymentIntentsCreateRequest) Execute() (*PaymentIntent, *http.Response, error) {
	return r.ApiService.PaymentIntentsCreateExecute(r)
}

/*
PaymentIntentsCreate Create a payment intent

Creates a payment intent to orchestrate payment collection with support for 3DS/SCA authentication.

---

**Related endpoints**

- `GET /payment-intents` — List payment intents
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaymentIntentsCreateRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsCreate(ctx context.Context) ApiPaymentIntentsCreateRequest {
	return ApiPaymentIntentsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentIntent
func (a *PaymentIntentsAPIService) PaymentIntentsCreateExecute(r ApiPaymentIntentsCreateRequest) (*PaymentIntent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents"

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
	localVarPostBody = r.createPaymentIntentRequest
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
			var v PaymentIntentErrorResponse
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
			var v PaymentIntentErrorResponse
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

type ApiPaymentIntentsGetRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	id string
}

func (r ApiPaymentIntentsGetRequest) Execute() (*PaymentIntent, *http.Response, error) {
	return r.ApiService.PaymentIntentsGetExecute(r)
}

/*
PaymentIntentsGet Retrieve a payment intent

Retrieves a payment intent by its ID (pi_xxx).

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents` — List payment intents
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Payment intent ID (pi_xxx)
 @return ApiPaymentIntentsGetRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsGet(ctx context.Context, id string) ApiPaymentIntentsGetRequest {
	return ApiPaymentIntentsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PaymentIntent
func (a *PaymentIntentsAPIService) PaymentIntentsGetExecute(r ApiPaymentIntentsGetRequest) (*PaymentIntent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{id}"
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
			var v PaymentIntentErrorResponse
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

type ApiPaymentIntentsListRequest struct {
	ctx context.Context
	ApiService PaymentIntentsAPI
	customer *string
	status *string
	createdGte *float32
	createdLte *float32
	limit *int32
	startingAfter *string
	endingBefore *string
}

// Filter by customer ID
func (r ApiPaymentIntentsListRequest) Customer(customer string) ApiPaymentIntentsListRequest {
	r.customer = &customer
	return r
}

// Filter by status
func (r ApiPaymentIntentsListRequest) Status(status string) ApiPaymentIntentsListRequest {
	r.status = &status
	return r
}

// Filter by created_at &gt;&#x3D; (Unix timestamp)
func (r ApiPaymentIntentsListRequest) CreatedGte(createdGte float32) ApiPaymentIntentsListRequest {
	r.createdGte = &createdGte
	return r
}

// Filter by created_at &lt;&#x3D; (Unix timestamp)
func (r ApiPaymentIntentsListRequest) CreatedLte(createdLte float32) ApiPaymentIntentsListRequest {
	r.createdLte = &createdLte
	return r
}

// Maximum number of results (1-100)
func (r ApiPaymentIntentsListRequest) Limit(limit int32) ApiPaymentIntentsListRequest {
	r.limit = &limit
	return r
}

// Cursor for pagination - return results after this ID (pi_xxx)
func (r ApiPaymentIntentsListRequest) StartingAfter(startingAfter string) ApiPaymentIntentsListRequest {
	r.startingAfter = &startingAfter
	return r
}

// Cursor for pagination - return results before this ID (pi_xxx)
func (r ApiPaymentIntentsListRequest) EndingBefore(endingBefore string) ApiPaymentIntentsListRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r ApiPaymentIntentsListRequest) Execute() (*PaymentIntentListResponse, *http.Response, error) {
	return r.ApiService.PaymentIntentsListExecute(r)
}

/*
PaymentIntentsList List payment intents

Returns a list of payment intents with optional filtering.

---

**Related endpoints**

- `POST /payment-intents` — Create a payment intent
- `GET /payment-intents/{id}` — Retrieve a payment intent
- `POST /payment-intents/{id}` — Update a payment intent
- `POST /payment-intents/{id}/confirm` — Confirm a payment intent
- `POST /payment-intents/{id}/capture` — Capture a payment intent
- `POST /payment-intents/{id}/cancel` — Cancel a payment intent

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPaymentIntentsListRequest
*/
func (a *PaymentIntentsAPIService) PaymentIntentsList(ctx context.Context) ApiPaymentIntentsListRequest {
	return ApiPaymentIntentsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentIntentListResponse
func (a *PaymentIntentsAPIService) PaymentIntentsListExecute(r ApiPaymentIntentsListRequest) (*PaymentIntentListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntentListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.PaymentIntentsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.customer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customer", r.customer, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.createdGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_gte", r.createdGte, "form", "")
	}
	if r.createdLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_lte", r.createdLte, "form", "")
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
			var v PaymentIntentErrorResponse
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
