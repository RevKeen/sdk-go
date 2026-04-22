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
)


type EntitlementsAPI interface {

	/*
	EntitlementsCheck Check entitlement access

	Check if a customer has access to a specific benefit by key. This is the primary endpoint for feature gating and licensing checks.

---

**Related endpoints**

- `GET /customers/{customerId}/entitlements` — List customer entitlements
- `POST /customers/{customerId}/entitlements` — Grant entitlement to customer
- `DELETE /customers/{customerId}/entitlements` — Revoke entitlement by benefit key
- `GET /customers/{customerId}/entitlements/check` — Check customer entitlement
- `DELETE /customers/{customerId}/entitlements/{entitlementId}` — Revoke entitlement by ID
- `GET /entitlements` — List entitlements

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEntitlementsCheckRequest
	*/
	EntitlementsCheck(ctx context.Context) ApiEntitlementsCheckRequest

	// EntitlementsCheckExecute executes the request
	//  @return EntitlementCheckResponse
	EntitlementsCheckExecute(r ApiEntitlementsCheckRequest) (*EntitlementCheckResponse, *http.Response, error)

	/*
	EntitlementsList List entitlements

	Retrieve all entitlements for a customer. Pass `customer_id` as a query parameter. Includes computed access status based on subscription state.

---

**Related endpoints**

- `GET /customers/{customerId}/entitlements` — List customer entitlements
- `POST /customers/{customerId}/entitlements` — Grant entitlement to customer
- `DELETE /customers/{customerId}/entitlements` — Revoke entitlement by benefit key
- `GET /customers/{customerId}/entitlements/check` — Check customer entitlement
- `DELETE /customers/{customerId}/entitlements/{entitlementId}` — Revoke entitlement by ID
- `GET /entitlements/check` — Check entitlement access

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiEntitlementsListRequest
	*/
	EntitlementsList(ctx context.Context) ApiEntitlementsListRequest

	// EntitlementsListExecute executes the request
	//  @return EntitlementListResponse
	EntitlementsListExecute(r ApiEntitlementsListRequest) (*EntitlementListResponse, *http.Response, error)
}

// EntitlementsAPIService EntitlementsAPI service
type EntitlementsAPIService service

type ApiEntitlementsCheckRequest struct {
	ctx context.Context
	ApiService EntitlementsAPI
	customerId *string
	benefitKey *string
}

// Customer UUID (required)
func (r ApiEntitlementsCheckRequest) CustomerId(customerId string) ApiEntitlementsCheckRequest {
	r.customerId = &customerId
	return r
}

// Benefit key to check
func (r ApiEntitlementsCheckRequest) BenefitKey(benefitKey string) ApiEntitlementsCheckRequest {
	r.benefitKey = &benefitKey
	return r
}

func (r ApiEntitlementsCheckRequest) Execute() (*EntitlementCheckResponse, *http.Response, error) {
	return r.ApiService.EntitlementsCheckExecute(r)
}

/*
EntitlementsCheck Check entitlement access

Check if a customer has access to a specific benefit by key. This is the primary endpoint for feature gating and licensing checks.

---

**Related endpoints**

- `GET /customers/{customerId}/entitlements` — List customer entitlements
- `POST /customers/{customerId}/entitlements` — Grant entitlement to customer
- `DELETE /customers/{customerId}/entitlements` — Revoke entitlement by benefit key
- `GET /customers/{customerId}/entitlements/check` — Check customer entitlement
- `DELETE /customers/{customerId}/entitlements/{entitlementId}` — Revoke entitlement by ID
- `GET /entitlements` — List entitlements

**Common errors**

- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntitlementsCheckRequest
*/
func (a *EntitlementsAPIService) EntitlementsCheck(ctx context.Context) ApiEntitlementsCheckRequest {
	return ApiEntitlementsCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntitlementCheckResponse
func (a *EntitlementsAPIService) EntitlementsCheckExecute(r ApiEntitlementsCheckRequest) (*EntitlementCheckResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntitlementCheckResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntitlementsAPIService.EntitlementsCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entitlements/check"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerId == nil {
		return localVarReturnValue, nil, reportError("customerId is required and must be specified")
	}
	if r.benefitKey == nil {
		return localVarReturnValue, nil, reportError("benefitKey is required and must be specified")
	}
	if strlen(*r.benefitKey) < 1 {
		return localVarReturnValue, nil, reportError("benefitKey must have at least 1 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "benefit_key", r.benefitKey, "form", "")
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

type ApiEntitlementsListRequest struct {
	ctx context.Context
	ApiService EntitlementsAPI
	customerId *string
	includeExpired *bool
	benefitType *string
	category *string
	limit *int32
	offset *int32
}

// Customer UUID (required)
func (r ApiEntitlementsListRequest) CustomerId(customerId string) ApiEntitlementsListRequest {
	r.customerId = &customerId
	return r
}

// Include expired entitlements
func (r ApiEntitlementsListRequest) IncludeExpired(includeExpired bool) ApiEntitlementsListRequest {
	r.includeExpired = &includeExpired
	return r
}

// Filter by benefit type
func (r ApiEntitlementsListRequest) BenefitType(benefitType string) ApiEntitlementsListRequest {
	r.benefitType = &benefitType
	return r
}

// Filter by category
func (r ApiEntitlementsListRequest) Category(category string) ApiEntitlementsListRequest {
	r.category = &category
	return r
}

// Maximum results (1-100)
func (r ApiEntitlementsListRequest) Limit(limit int32) ApiEntitlementsListRequest {
	r.limit = &limit
	return r
}

// Results to skip
func (r ApiEntitlementsListRequest) Offset(offset int32) ApiEntitlementsListRequest {
	r.offset = &offset
	return r
}

func (r ApiEntitlementsListRequest) Execute() (*EntitlementListResponse, *http.Response, error) {
	return r.ApiService.EntitlementsListExecute(r)
}

/*
EntitlementsList List entitlements

Retrieve all entitlements for a customer. Pass `customer_id` as a query parameter. Includes computed access status based on subscription state.

---

**Related endpoints**

- `GET /customers/{customerId}/entitlements` — List customer entitlements
- `POST /customers/{customerId}/entitlements` — Grant entitlement to customer
- `DELETE /customers/{customerId}/entitlements` — Revoke entitlement by benefit key
- `GET /customers/{customerId}/entitlements/check` — Check customer entitlement
- `DELETE /customers/{customerId}/entitlements/{entitlementId}` — Revoke entitlement by ID
- `GET /entitlements/check` — Check entitlement access

**Common errors**

- `400 invalid_request` — malformed payload or failed validation.
- `401 unauthenticated` — missing, malformed, or revoked API key.
- `404 resource_missing` — the referenced resource does not exist or is not visible to your key.

**Pagination**

Offset-based with `limit` (default 25, max 100) and `offset`. The response `pagination` block includes `total` and `hasMore`. See [the pagination guide](/docs/fundamentals/pagination) for SDK auto-paging helpers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntitlementsListRequest
*/
func (a *EntitlementsAPIService) EntitlementsList(ctx context.Context) ApiEntitlementsListRequest {
	return ApiEntitlementsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntitlementListResponse
func (a *EntitlementsAPIService) EntitlementsListExecute(r ApiEntitlementsListRequest) (*EntitlementListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntitlementListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntitlementsAPIService.EntitlementsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entitlements"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerId == nil {
		return localVarReturnValue, nil, reportError("customerId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customer_id", r.customerId, "form", "")
	if r.includeExpired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_expired", r.includeExpired, "form", "")
	} else {
        var defaultValue bool = false
        parameterAddToHeaderOrQuery(localVarQueryParams, "include_expired", defaultValue, "form", "")
        r.includeExpired = &defaultValue
	}
	if r.benefitType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "benefit_type", r.benefitType, "form", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "form", "")
	}
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
