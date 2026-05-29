/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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


type CartSessionsAPI interface {

	/*
	CartSessionsAddLineItem Add a line item to a cart session

	Append a line item to an open cart session. Subtotals are recomputed atomically; the updated cart is returned.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCartSessionsAddLineItemRequest
	*/
	CartSessionsAddLineItem(ctx context.Context, id string) ApiCartSessionsAddLineItemRequest

	// CartSessionsAddLineItemExecute executes the request
	//  @return CartSessionResponse
	CartSessionsAddLineItemExecute(r ApiCartSessionsAddLineItemRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsApplyDiscountCode Set or clear a cart discount code

	Set the cart's discount code. Pass `null` to clear. The code is stored only — discount pricing math has not shipped yet. Convert blocks with `CART_SESSION_DISCOUNT_PENDING` while a code is set.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCartSessionsApplyDiscountCodeRequest
	*/
	CartSessionsApplyDiscountCode(ctx context.Context, id string) ApiCartSessionsApplyDiscountCodeRequest

	// CartSessionsApplyDiscountCodeExecute executes the request
	//  @return CartSessionResponse
	CartSessionsApplyDiscountCodeExecute(r ApiCartSessionsApplyDiscountCodeRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsConvert Convert a cart session into a checkout session

	Atomically materialise a pending checkout session from the cart snapshot, flip the cart to `converted`, and emit `commerce.cart.converted` + `commerce.checkout.started` through the outbox.

Idempotent on re-call (returns the existing checkout session). Concurrency-safe via a compare-and-swap lock on cart status.

Validation that runs inside the lock and rolls back on failure:
- `CART_SESSION_EMPTY` — the cart has no line items.
- `CART_SESSION_DISCOUNT_PENDING` — the cart has a discount code set; clear it first.
- `CART_SESSION_NOT_FOUND` — no cart for this id under the calling merchant.
- `CART_SESSION_CLOSED` — the cart is already `abandoned` or `expired`.

On rollback the cart returns to `open` and the customer can retry after fixing the cause.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCartSessionsConvertRequest
	*/
	CartSessionsConvert(ctx context.Context, id string) ApiCartSessionsConvertRequest

	// CartSessionsConvertExecute executes the request
	//  @return CartConversionResponse
	CartSessionsConvertExecute(r ApiCartSessionsConvertRequest) (*CartConversionResponse, *http.Response, error)

	/*
	CartSessionsCreate Create a cart session

	Create an empty (or pre-populated) cart session. The returned cart is `open` and mutable until it is converted into a checkout session, or until it is swept to `abandoned` / `expired`.

---

**Related endpoints**

- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/convert` — Convert to a checkout session

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

---

**Related endpoints**

- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCartSessionsCreateRequest
	*/
	CartSessionsCreate(ctx context.Context) ApiCartSessionsCreateRequest

	// CartSessionsCreateExecute executes the request
	//  @return CartSessionResponse
	CartSessionsCreateExecute(r ApiCartSessionsCreateRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsGet Retrieve a cart session

	Fetch the current state of a cart session by id. Returns the same shape as the create response.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCartSessionsGetRequest
	*/
	CartSessionsGet(ctx context.Context, id string) ApiCartSessionsGetRequest

	// CartSessionsGetExecute executes the request
	//  @return CartSessionResponse
	CartSessionsGetExecute(r ApiCartSessionsGetRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsRemoveLineItem Remove a line item from a cart session

	Remove a line item from an open cart. Subtotals are recomputed; the updated cart is returned.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@param lineId Cart line item id (the `id` field returned on `cart_session.line_items[]`).
	@return ApiCartSessionsRemoveLineItemRequest
	*/
	CartSessionsRemoveLineItem(ctx context.Context, id string, lineId string) ApiCartSessionsRemoveLineItemRequest

	// CartSessionsRemoveLineItemExecute executes the request
	//  @return CartSessionResponse
	CartSessionsRemoveLineItemExecute(r ApiCartSessionsRemoveLineItemRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsToggleAddOn Toggle an add-on on a cart session

	Add or remove an add-on product from `add_ons_selected`. The call is idempotent: sending the same desired state for an add-on already in that state is a no-op and emits no event.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ApiCartSessionsToggleAddOnRequest
	*/
	CartSessionsToggleAddOn(ctx context.Context, id string) ApiCartSessionsToggleAddOnRequest

	// CartSessionsToggleAddOnExecute executes the request
	//  @return CartSessionResponse
	CartSessionsToggleAddOnExecute(r ApiCartSessionsToggleAddOnRequest) (*CartSessionResponse, *http.Response, error)

	/*
	CartSessionsUpdateLineItem Update a line item's quantity

	Change the quantity of a single line item on an open cart. Other fields on the line item are immutable; recreate the item to change them.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@param lineId Cart line item id (the `id` field returned on `cart_session.line_items[]`).
	@return ApiCartSessionsUpdateLineItemRequest
	*/
	CartSessionsUpdateLineItem(ctx context.Context, id string, lineId string) ApiCartSessionsUpdateLineItemRequest

	// CartSessionsUpdateLineItemExecute executes the request
	//  @return CartSessionResponse
	CartSessionsUpdateLineItemExecute(r ApiCartSessionsUpdateLineItemRequest) (*CartSessionResponse, *http.Response, error)
}

// CartSessionsAPIService CartSessionsAPI service
type CartSessionsAPIService service

type ApiCartSessionsAddLineItemRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
	cartLineItemInput *CartLineItemInput
}

func (r ApiCartSessionsAddLineItemRequest) CartLineItemInput(cartLineItemInput CartLineItemInput) ApiCartSessionsAddLineItemRequest {
	r.cartLineItemInput = &cartLineItemInput
	return r
}

func (r ApiCartSessionsAddLineItemRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsAddLineItemExecute(r)
}

/*
CartSessionsAddLineItem Add a line item to a cart session

Append a line item to an open cart session. Subtotals are recomputed atomically; the updated cart is returned.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCartSessionsAddLineItemRequest
*/
func (a *CartSessionsAPIService) CartSessionsAddLineItem(ctx context.Context, id string) ApiCartSessionsAddLineItemRequest {
	return ApiCartSessionsAddLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsAddLineItemExecute(r ApiCartSessionsAddLineItemRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsAddLineItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/line-items"
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
	localVarPostBody = r.cartLineItemInput
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsApplyDiscountCodeRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
	applyCartDiscountCodeInput *ApplyCartDiscountCodeInput
}

func (r ApiCartSessionsApplyDiscountCodeRequest) ApplyCartDiscountCodeInput(applyCartDiscountCodeInput ApplyCartDiscountCodeInput) ApiCartSessionsApplyDiscountCodeRequest {
	r.applyCartDiscountCodeInput = &applyCartDiscountCodeInput
	return r
}

func (r ApiCartSessionsApplyDiscountCodeRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsApplyDiscountCodeExecute(r)
}

/*
CartSessionsApplyDiscountCode Set or clear a cart discount code

Set the cart's discount code. Pass `null` to clear. The code is stored only — discount pricing math has not shipped yet. Convert blocks with `CART_SESSION_DISCOUNT_PENDING` while a code is set.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCartSessionsApplyDiscountCodeRequest
*/
func (a *CartSessionsAPIService) CartSessionsApplyDiscountCode(ctx context.Context, id string) ApiCartSessionsApplyDiscountCodeRequest {
	return ApiCartSessionsApplyDiscountCodeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsApplyDiscountCodeExecute(r ApiCartSessionsApplyDiscountCodeRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsApplyDiscountCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/discount-code"
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
	localVarPostBody = r.applyCartDiscountCodeInput
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsConvertRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
}

func (r ApiCartSessionsConvertRequest) Execute() (*CartConversionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsConvertExecute(r)
}

/*
CartSessionsConvert Convert a cart session into a checkout session

Atomically materialise a pending checkout session from the cart snapshot, flip the cart to `converted`, and emit `commerce.cart.converted` + `commerce.checkout.started` through the outbox.

Idempotent on re-call (returns the existing checkout session). Concurrency-safe via a compare-and-swap lock on cart status.

Validation that runs inside the lock and rolls back on failure:
- `CART_SESSION_EMPTY` — the cart has no line items.
- `CART_SESSION_DISCOUNT_PENDING` — the cart has a discount code set; clear it first.
- `CART_SESSION_NOT_FOUND` — no cart for this id under the calling merchant.
- `CART_SESSION_CLOSED` — the cart is already `abandoned` or `expired`.

On rollback the cart returns to `open` and the customer can retry after fixing the cause.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCartSessionsConvertRequest
*/
func (a *CartSessionsAPIService) CartSessionsConvert(ctx context.Context, id string) ApiCartSessionsConvertRequest {
	return ApiCartSessionsConvertRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CartConversionResponse
func (a *CartSessionsAPIService) CartSessionsConvertExecute(r ApiCartSessionsConvertRequest) (*CartConversionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartConversionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsConvert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/convert"
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsCreateRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	createCartSessionInput *CreateCartSessionInput
}

func (r ApiCartSessionsCreateRequest) CreateCartSessionInput(createCartSessionInput CreateCartSessionInput) ApiCartSessionsCreateRequest {
	r.createCartSessionInput = &createCartSessionInput
	return r
}

func (r ApiCartSessionsCreateRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsCreateExecute(r)
}

/*
CartSessionsCreate Create a cart session

Create an empty (or pre-populated) cart session. The returned cart is `open` and mutable until it is converted into a checkout session, or until it is swept to `abandoned` / `expired`.

---

**Related endpoints**

- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/convert` — Convert to a checkout session

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

---

**Related endpoints**

- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCartSessionsCreateRequest
*/
func (a *CartSessionsAPIService) CartSessionsCreate(ctx context.Context) ApiCartSessionsCreateRequest {
	return ApiCartSessionsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsCreateExecute(r ApiCartSessionsCreateRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions"

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
	localVarPostBody = r.createCartSessionInput
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsGetRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
}

func (r ApiCartSessionsGetRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsGetExecute(r)
}

/*
CartSessionsGet Retrieve a cart session

Fetch the current state of a cart session by id. Returns the same shape as the create response.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCartSessionsGetRequest
*/
func (a *CartSessionsAPIService) CartSessionsGet(ctx context.Context, id string) ApiCartSessionsGetRequest {
	return ApiCartSessionsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsGetExecute(r ApiCartSessionsGetRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}"
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsRemoveLineItemRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
	lineId string
}

func (r ApiCartSessionsRemoveLineItemRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsRemoveLineItemExecute(r)
}

/*
CartSessionsRemoveLineItem Remove a line item from a cart session

Remove a line item from an open cart. Subtotals are recomputed; the updated cart is returned.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param lineId Cart line item id (the `id` field returned on `cart_session.line_items[]`).
 @return ApiCartSessionsRemoveLineItemRequest
*/
func (a *CartSessionsAPIService) CartSessionsRemoveLineItem(ctx context.Context, id string, lineId string) ApiCartSessionsRemoveLineItemRequest {
	return ApiCartSessionsRemoveLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		lineId: lineId,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsRemoveLineItemExecute(r ApiCartSessionsRemoveLineItemRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsRemoveLineItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/line-items/{lineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lineId"+"}", url.PathEscape(parameterValueToString(r.lineId, "lineId")), -1)

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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsToggleAddOnRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
	toggleCartAddOnInput *ToggleCartAddOnInput
}

func (r ApiCartSessionsToggleAddOnRequest) ToggleCartAddOnInput(toggleCartAddOnInput ToggleCartAddOnInput) ApiCartSessionsToggleAddOnRequest {
	r.toggleCartAddOnInput = &toggleCartAddOnInput
	return r
}

func (r ApiCartSessionsToggleAddOnRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsToggleAddOnExecute(r)
}

/*
CartSessionsToggleAddOn Toggle an add-on on a cart session

Add or remove an add-on product from `add_ons_selected`. The call is idempotent: sending the same desired state for an add-on already in that state is a no-op and emits no event.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `PATCH /cart-sessions/{id}/line-items/{lineId}` — Update a line item's quantity
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiCartSessionsToggleAddOnRequest
*/
func (a *CartSessionsAPIService) CartSessionsToggleAddOn(ctx context.Context, id string) ApiCartSessionsToggleAddOnRequest {
	return ApiCartSessionsToggleAddOnRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsToggleAddOnExecute(r ApiCartSessionsToggleAddOnRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsToggleAddOn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/add-ons"
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
	localVarPostBody = r.toggleCartAddOnInput
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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

type ApiCartSessionsUpdateLineItemRequest struct {
	ctx context.Context
	ApiService CartSessionsAPI
	id string
	lineId string
	updateCartLineItemInput *UpdateCartLineItemInput
}

func (r ApiCartSessionsUpdateLineItemRequest) UpdateCartLineItemInput(updateCartLineItemInput UpdateCartLineItemInput) ApiCartSessionsUpdateLineItemRequest {
	r.updateCartLineItemInput = &updateCartLineItemInput
	return r
}

func (r ApiCartSessionsUpdateLineItemRequest) Execute() (*CartSessionResponse, *http.Response, error) {
	return r.ApiService.CartSessionsUpdateLineItemExecute(r)
}

/*
CartSessionsUpdateLineItem Update a line item's quantity

Change the quantity of a single line item on an open cart. Other fields on the line item are immutable; recreate the item to change them.

---

**Related endpoints**

- `POST /cart-sessions` — Create a cart session
- `GET /cart-sessions/{id}` — Retrieve a cart session
- `POST /cart-sessions/{id}/line-items` — Add a line item to a cart session
- `DELETE /cart-sessions/{id}/line-items/{lineId}` — Remove a line item from a cart session
- `POST /cart-sessions/{id}/add-ons` — Toggle an add-on on a cart session
- `POST /cart-sessions/{id}/discount-code` — Set or clear a cart discount code
- `POST /cart-sessions/{id}/convert` — Convert a cart session into a checkout session

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.
- `409 conflict` — Idempotency-Key collision with a different body, or a concurrent state-transition conflict.

**Idempotency**

Pass an `Idempotency-Key` header (UUID v4 recommended) to make retries safe. Keys are valid for 24 hours; see [the idempotency guide](/docs/fundamentals/idempotency).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @param lineId Cart line item id (the `id` field returned on `cart_session.line_items[]`).
 @return ApiCartSessionsUpdateLineItemRequest
*/
func (a *CartSessionsAPIService) CartSessionsUpdateLineItem(ctx context.Context, id string, lineId string) ApiCartSessionsUpdateLineItemRequest {
	return ApiCartSessionsUpdateLineItemRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		lineId: lineId,
	}
}

// Execute executes the request
//  @return CartSessionResponse
func (a *CartSessionsAPIService) CartSessionsUpdateLineItemExecute(r ApiCartSessionsUpdateLineItemRequest) (*CartSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CartSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CartSessionsAPIService.CartSessionsUpdateLineItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cart-sessions/{id}/line-items/{lineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lineId"+"}", url.PathEscape(parameterValueToString(r.lineId, "lineId")), -1)

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
	localVarPostBody = r.updateCartLineItemInput
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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
			var v CartSessionErrorResponse
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
