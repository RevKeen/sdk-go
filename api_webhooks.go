/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

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


type WebhooksAPI interface {

	/*
	WebhooksCheckoutSessionCompleted Checkout session completed

	Sent when a checkout session is successfully completed. **Action required:** Fulfill the customer's order.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCheckoutSessionCompletedRequest
	*/
	WebhooksCheckoutSessionCompleted(ctx context.Context) ApiWebhooksCheckoutSessionCompletedRequest

	// WebhooksCheckoutSessionCompletedExecute executes the request
	WebhooksCheckoutSessionCompletedExecute(r ApiWebhooksCheckoutSessionCompletedRequest) (*http.Response, error)

	/*
	WebhooksCheckoutSessionExpired Checkout session expired

	Sent when a checkout session expires before completion.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCheckoutSessionExpiredRequest
	*/
	WebhooksCheckoutSessionExpired(ctx context.Context) ApiWebhooksCheckoutSessionExpiredRequest

	// WebhooksCheckoutSessionExpiredExecute executes the request
	WebhooksCheckoutSessionExpiredExecute(r ApiWebhooksCheckoutSessionExpiredRequest) (*http.Response, error)

	/*
	WebhooksCollectionFailed Collection failed

	Sent when a Direct Debit collection fails (e.g. insufficient funds). Contains the Bacs reason code. **Action required:** RevKeen Recovery may retry or fall back to a stored card per your policy.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCollectionFailedRequest
	*/
	WebhooksCollectionFailed(ctx context.Context) ApiWebhooksCollectionFailedRequest

	// WebhooksCollectionFailedExecute executes the request
	WebhooksCollectionFailedExecute(r ApiWebhooksCollectionFailedRequest) (*http.Response, error)

	/*
	WebhooksCollectionIndemnityClaimed Indemnity claimed

	Sent when a customer raises a Direct Debit indemnity claim (the Bacs chargeback equivalent) and the collection is reversed. **Action required:** Review the disputed collection.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCollectionIndemnityClaimedRequest
	*/
	WebhooksCollectionIndemnityClaimed(ctx context.Context) ApiWebhooksCollectionIndemnityClaimedRequest

	// WebhooksCollectionIndemnityClaimedExecute executes the request
	WebhooksCollectionIndemnityClaimedExecute(r ApiWebhooksCollectionIndemnityClaimedRequest) (*http.Response, error)

	/*
	WebhooksCollectionNoticeSent Advance notice sent

	Sent when the mandatory advance notice for an upcoming collection has been issued to the customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCollectionNoticeSentRequest
	*/
	WebhooksCollectionNoticeSent(ctx context.Context) ApiWebhooksCollectionNoticeSentRequest

	// WebhooksCollectionNoticeSentExecute executes the request
	WebhooksCollectionNoticeSentExecute(r ApiWebhooksCollectionNoticeSentRequest) (*http.Response, error)

	/*
	WebhooksCollectionScheduled Collection scheduled

	Sent when a Direct Debit collection is scheduled against an active mandate. Includes the working-day collection date.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCollectionScheduledRequest
	*/
	WebhooksCollectionScheduled(ctx context.Context) ApiWebhooksCollectionScheduledRequest

	// WebhooksCollectionScheduledExecute executes the request
	WebhooksCollectionScheduledExecute(r ApiWebhooksCollectionScheduledRequest) (*http.Response, error)

	/*
	WebhooksCollectionSucceeded Collection succeeded

	Sent when a Direct Debit collection is successfully taken. **Action required:** Provision goods or services to your customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCollectionSucceededRequest
	*/
	WebhooksCollectionSucceeded(ctx context.Context) ApiWebhooksCollectionSucceededRequest

	// WebhooksCollectionSucceededExecute executes the request
	WebhooksCollectionSucceededExecute(r ApiWebhooksCollectionSucceededRequest) (*http.Response, error)

	/*
	WebhooksCreditNoteCreated Credit note created

	Sent when a credit note is created for an invoice.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCreditNoteCreatedRequest
	*/
	WebhooksCreditNoteCreated(ctx context.Context) ApiWebhooksCreditNoteCreatedRequest

	// WebhooksCreditNoteCreatedExecute executes the request
	WebhooksCreditNoteCreatedExecute(r ApiWebhooksCreditNoteCreatedRequest) (*http.Response, error)

	/*
	WebhooksCreditNoteIssued Credit note issued

	Sent when a credit note is issued (finalized) and its PDF is generated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCreditNoteIssuedRequest
	*/
	WebhooksCreditNoteIssued(ctx context.Context) ApiWebhooksCreditNoteIssuedRequest

	// WebhooksCreditNoteIssuedExecute executes the request
	WebhooksCreditNoteIssuedExecute(r ApiWebhooksCreditNoteIssuedRequest) (*http.Response, error)

	/*
	WebhooksCreditNoteVoided Credit note voided

	Sent when a credit note is voided (accounting reversal).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCreditNoteVoidedRequest
	*/
	WebhooksCreditNoteVoided(ctx context.Context) ApiWebhooksCreditNoteVoidedRequest

	// WebhooksCreditNoteVoidedExecute executes the request
	WebhooksCreditNoteVoidedExecute(r ApiWebhooksCreditNoteVoidedRequest) (*http.Response, error)

	/*
	WebhooksCustomerCreated Customer created

	Sent when a new customer is created.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCustomerCreatedRequest
	*/
	WebhooksCustomerCreated(ctx context.Context) ApiWebhooksCustomerCreatedRequest

	// WebhooksCustomerCreatedExecute executes the request
	WebhooksCustomerCreatedExecute(r ApiWebhooksCustomerCreatedRequest) (*http.Response, error)

	/*
	WebhooksCustomerUpdated Customer updated

	Sent when customer information is updated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksCustomerUpdatedRequest
	*/
	WebhooksCustomerUpdated(ctx context.Context) ApiWebhooksCustomerUpdatedRequest

	// WebhooksCustomerUpdatedExecute executes the request
	WebhooksCustomerUpdatedExecute(r ApiWebhooksCustomerUpdatedRequest) (*http.Response, error)

	/*
	WebhooksInvoiceCreated Invoice created

	Sent when a new invoice is created, either manually or from a subscription renewal.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksInvoiceCreatedRequest
	*/
	WebhooksInvoiceCreated(ctx context.Context) ApiWebhooksInvoiceCreatedRequest

	// WebhooksInvoiceCreatedExecute executes the request
	WebhooksInvoiceCreatedExecute(r ApiWebhooksInvoiceCreatedRequest) (*http.Response, error)

	/*
	WebhooksInvoiceOverdue Invoice overdue

	Sent when an invoice becomes overdue. Consider sending payment reminders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksInvoiceOverdueRequest
	*/
	WebhooksInvoiceOverdue(ctx context.Context) ApiWebhooksInvoiceOverdueRequest

	// WebhooksInvoiceOverdueExecute executes the request
	WebhooksInvoiceOverdueExecute(r ApiWebhooksInvoiceOverdueRequest) (*http.Response, error)

	/*
	WebhooksInvoicePaid Invoice paid

	Sent when an invoice is fully paid. **Action required:** Fulfill the order or activate the subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksInvoicePaidRequest
	*/
	WebhooksInvoicePaid(ctx context.Context) ApiWebhooksInvoicePaidRequest

	// WebhooksInvoicePaidExecute executes the request
	WebhooksInvoicePaidExecute(r ApiWebhooksInvoicePaidRequest) (*http.Response, error)

	/*
	WebhooksMandateActivated Mandate activated

	Sent when a mandate is lodged and active. **Action required:** You can now collect from this customer via Direct Debit.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMandateActivatedRequest
	*/
	WebhooksMandateActivated(ctx context.Context) ApiWebhooksMandateActivatedRequest

	// WebhooksMandateActivatedExecute executes the request
	WebhooksMandateActivatedExecute(r ApiWebhooksMandateActivatedRequest) (*http.Response, error)

	/*
	WebhooksMandateAuddisRejected Mandate rejected (AUDDIS)

	Sent when the customer's bank rejects the mandate setup (AUDDIS). **Action required:** The mandate cannot be used — re-collect bank details or use another payment method.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMandateAuddisRejectedRequest
	*/
	WebhooksMandateAuddisRejected(ctx context.Context) ApiWebhooksMandateAuddisRejectedRequest

	// WebhooksMandateAuddisRejectedExecute executes the request
	WebhooksMandateAuddisRejectedExecute(r ApiWebhooksMandateAuddisRejectedRequest) (*http.Response, error)

	/*
	WebhooksMandateCancelled Mandate cancelled

	Sent when a mandate is cancelled (by you, the customer, or the bank). It can no longer be collected against.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMandateCancelledRequest
	*/
	WebhooksMandateCancelled(ctx context.Context) ApiWebhooksMandateCancelledRequest

	// WebhooksMandateCancelledExecute executes the request
	WebhooksMandateCancelledExecute(r ApiWebhooksMandateCancelledRequest) (*http.Response, error)

	/*
	WebhooksMandateCreated Mandate created

	Sent when a Direct Debit mandate is created and submitted to Bacs for lodgement. The mandate is not yet collectable — wait for `mandate.activated`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMandateCreatedRequest
	*/
	WebhooksMandateCreated(ctx context.Context) ApiWebhooksMandateCreatedRequest

	// WebhooksMandateCreatedExecute executes the request
	WebhooksMandateCreatedExecute(r ApiWebhooksMandateCreatedRequest) (*http.Response, error)

	/*
	WebhooksMandateSuspended Mandate suspended

	Sent when a mandate is suspended and can no longer be collected against until reinstated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMandateSuspendedRequest
	*/
	WebhooksMandateSuspended(ctx context.Context) ApiWebhooksMandateSuspendedRequest

	// WebhooksMandateSuspendedExecute executes the request
	WebhooksMandateSuspendedExecute(r ApiWebhooksMandateSuspendedRequest) (*http.Response, error)

	/*
	WebhooksMeterArchived Meter archived

	Sent when a usage meter is archived and can no longer be attached to new pricing.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterArchivedRequest
	*/
	WebhooksMeterArchived(ctx context.Context) ApiWebhooksMeterArchivedRequest

	// WebhooksMeterArchivedExecute executes the request
	WebhooksMeterArchivedExecute(r ApiWebhooksMeterArchivedRequest) (*http.Response, error)

	/*
	WebhooksMeterCreated Meter created

	Sent when a usage meter is created.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterCreatedRequest
	*/
	WebhooksMeterCreated(ctx context.Context) ApiWebhooksMeterCreatedRequest

	// WebhooksMeterCreatedExecute executes the request
	WebhooksMeterCreatedExecute(r ApiWebhooksMeterCreatedRequest) (*http.Response, error)

	/*
	WebhooksMeterPriceCreated Meter price created

	Sent when a meter price is created for a product price.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterPriceCreatedRequest
	*/
	WebhooksMeterPriceCreated(ctx context.Context) ApiWebhooksMeterPriceCreatedRequest

	// WebhooksMeterPriceCreatedExecute executes the request
	WebhooksMeterPriceCreatedExecute(r ApiWebhooksMeterPriceCreatedRequest) (*http.Response, error)

	/*
	WebhooksMeterPriceDeactivated Meter price deactivated

	Sent when a meter price is deactivated and no longer available for new subscriptions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterPriceDeactivatedRequest
	*/
	WebhooksMeterPriceDeactivated(ctx context.Context) ApiWebhooksMeterPriceDeactivatedRequest

	// WebhooksMeterPriceDeactivatedExecute executes the request
	WebhooksMeterPriceDeactivatedExecute(r ApiWebhooksMeterPriceDeactivatedRequest) (*http.Response, error)

	/*
	WebhooksMeterPriceUpdated Meter price updated

	Sent when a meter price is updated. The payload includes changed attributes when available.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterPriceUpdatedRequest
	*/
	WebhooksMeterPriceUpdated(ctx context.Context) ApiWebhooksMeterPriceUpdatedRequest

	// WebhooksMeterPriceUpdatedExecute executes the request
	WebhooksMeterPriceUpdatedExecute(r ApiWebhooksMeterPriceUpdatedRequest) (*http.Response, error)

	/*
	WebhooksMeterUpdated Meter updated

	Sent when a usage meter is updated. The payload includes changed attributes when available.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksMeterUpdatedRequest
	*/
	WebhooksMeterUpdated(ctx context.Context) ApiWebhooksMeterUpdatedRequest

	// WebhooksMeterUpdatedExecute executes the request
	WebhooksMeterUpdatedExecute(r ApiWebhooksMeterUpdatedRequest) (*http.Response, error)

	/*
	WebhooksOrderCreated Order created

	Sent when a new order is created.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksOrderCreatedRequest
	*/
	WebhooksOrderCreated(ctx context.Context) ApiWebhooksOrderCreatedRequest

	// WebhooksOrderCreatedExecute executes the request
	WebhooksOrderCreatedExecute(r ApiWebhooksOrderCreatedRequest) (*http.Response, error)

	/*
	WebhooksOrderFulfilled Order fulfilled

	Sent when an order is marked as fulfilled.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksOrderFulfilledRequest
	*/
	WebhooksOrderFulfilled(ctx context.Context) ApiWebhooksOrderFulfilledRequest

	// WebhooksOrderFulfilledExecute executes the request
	WebhooksOrderFulfilledExecute(r ApiWebhooksOrderFulfilledRequest) (*http.Response, error)

	/*
	WebhooksOrderPaid Order paid

	Sent when an order is fully paid. **Action required:** Begin fulfillment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksOrderPaidRequest
	*/
	WebhooksOrderPaid(ctx context.Context) ApiWebhooksOrderPaidRequest

	// WebhooksOrderPaidExecute executes the request
	WebhooksOrderPaidExecute(r ApiWebhooksOrderPaidRequest) (*http.Response, error)

	/*
	WebhooksPaymentFailed Payment failed

	Sent when a payment attempt fails. Contains failure reason and suggested next steps.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksPaymentFailedRequest
	*/
	WebhooksPaymentFailed(ctx context.Context) ApiWebhooksPaymentFailedRequest

	// WebhooksPaymentFailedExecute executes the request
	WebhooksPaymentFailedExecute(r ApiWebhooksPaymentFailedRequest) (*http.Response, error)

	/*
	WebhooksPaymentSucceeded Payment succeeded

	Sent when a payment is successfully captured. **Action required:** Provision goods or services to your customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksPaymentSucceededRequest
	*/
	WebhooksPaymentSucceeded(ctx context.Context) ApiWebhooksPaymentSucceededRequest

	// WebhooksPaymentSucceededExecute executes the request
	WebhooksPaymentSucceededExecute(r ApiWebhooksPaymentSucceededRequest) (*http.Response, error)

	/*
	WebhooksRefundCreated Refund created

	Sent when a refund is initiated.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksRefundCreatedRequest
	*/
	WebhooksRefundCreated(ctx context.Context) ApiWebhooksRefundCreatedRequest

	// WebhooksRefundCreatedExecute executes the request
	WebhooksRefundCreatedExecute(r ApiWebhooksRefundCreatedRequest) (*http.Response, error)

	/*
	WebhooksRefundSucceeded Refund succeeded

	Sent when a refund is successfully processed by the payment gateway.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksRefundSucceededRequest
	*/
	WebhooksRefundSucceeded(ctx context.Context) ApiWebhooksRefundSucceededRequest

	// WebhooksRefundSucceededExecute executes the request
	WebhooksRefundSucceededExecute(r ApiWebhooksRefundSucceededRequest) (*http.Response, error)

	/*
	WebhooksSettlementCreated Settlement created

	Sent when collected Direct Debit funds are settled and a payout is created. Gross settlement — the bureau fee is reported separately and never netted from the payout.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksSettlementCreatedRequest
	*/
	WebhooksSettlementCreated(ctx context.Context) ApiWebhooksSettlementCreatedRequest

	// WebhooksSettlementCreatedExecute executes the request
	WebhooksSettlementCreatedExecute(r ApiWebhooksSettlementCreatedRequest) (*http.Response, error)

	/*
	WebhooksSubscriptionActivated Subscription activated

	Sent when a subscription becomes active after successful payment. **Action required:** Grant access to your service.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksSubscriptionActivatedRequest
	*/
	WebhooksSubscriptionActivated(ctx context.Context) ApiWebhooksSubscriptionActivatedRequest

	// WebhooksSubscriptionActivatedExecute executes the request
	WebhooksSubscriptionActivatedExecute(r ApiWebhooksSubscriptionActivatedRequest) (*http.Response, error)

	/*
	WebhooksSubscriptionCanceled Subscription canceled

	Sent when a subscription is canceled. Access continues until the end of the billing period.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksSubscriptionCanceledRequest
	*/
	WebhooksSubscriptionCanceled(ctx context.Context) ApiWebhooksSubscriptionCanceledRequest

	// WebhooksSubscriptionCanceledExecute executes the request
	WebhooksSubscriptionCanceledExecute(r ApiWebhooksSubscriptionCanceledRequest) (*http.Response, error)

	/*
	WebhooksSubscriptionCreated Subscription created

	Sent when a new subscription is created (before activation).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksSubscriptionCreatedRequest
	*/
	WebhooksSubscriptionCreated(ctx context.Context) ApiWebhooksSubscriptionCreatedRequest

	// WebhooksSubscriptionCreatedExecute executes the request
	WebhooksSubscriptionCreatedExecute(r ApiWebhooksSubscriptionCreatedRequest) (*http.Response, error)

	/*
	WebhooksSubscriptionRenewed Subscription renewed

	Sent when a subscription successfully renews. A new invoice has been created and paid.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksSubscriptionRenewedRequest
	*/
	WebhooksSubscriptionRenewed(ctx context.Context) ApiWebhooksSubscriptionRenewedRequest

	// WebhooksSubscriptionRenewedExecute executes the request
	WebhooksSubscriptionRenewedExecute(r ApiWebhooksSubscriptionRenewedRequest) (*http.Response, error)

	/*
	WebhooksTerminalPaymentCancelled Terminal payment cancelled

	Sent when a card-present payment is cancelled by the merchant before terminal response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalPaymentCancelledRequest
	*/
	WebhooksTerminalPaymentCancelled(ctx context.Context) ApiWebhooksTerminalPaymentCancelledRequest

	// WebhooksTerminalPaymentCancelledExecute executes the request
	WebhooksTerminalPaymentCancelledExecute(r ApiWebhooksTerminalPaymentCancelledRequest) (*http.Response, error)

	/*
	WebhooksTerminalPaymentDeclined Terminal payment declined

	Sent when a card-present payment is declined by the terminal or card issuer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalPaymentDeclinedRequest
	*/
	WebhooksTerminalPaymentDeclined(ctx context.Context) ApiWebhooksTerminalPaymentDeclinedRequest

	// WebhooksTerminalPaymentDeclinedExecute executes the request
	WebhooksTerminalPaymentDeclinedExecute(r ApiWebhooksTerminalPaymentDeclinedRequest) (*http.Response, error)

	/*
	WebhooksTerminalPaymentError Terminal payment error

	Sent when a card-present payment fails due to timeout, terminal error, or connection loss. Check `failure_reason` field: `timeout`, `terminal_error`, or `connection_lost`.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalPaymentErrorRequest
	*/
	WebhooksTerminalPaymentError(ctx context.Context) ApiWebhooksTerminalPaymentErrorRequest

	// WebhooksTerminalPaymentErrorExecute executes the request
	WebhooksTerminalPaymentErrorExecute(r ApiWebhooksTerminalPaymentErrorRequest) (*http.Response, error)

	/*
	WebhooksTerminalPaymentRequested Terminal payment requested

	Sent when a card-present payment is dispatched to a terminal device. The payment is in-progress and awaiting terminal response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalPaymentRequestedRequest
	*/
	WebhooksTerminalPaymentRequested(ctx context.Context) ApiWebhooksTerminalPaymentRequestedRequest

	// WebhooksTerminalPaymentRequestedExecute executes the request
	WebhooksTerminalPaymentRequestedExecute(r ApiWebhooksTerminalPaymentRequestedRequest) (*http.Response, error)

	/*
	WebhooksTerminalPaymentSucceeded Terminal payment succeeded

	Sent when a card-present payment is approved by the terminal. **Action required:** Provide goods or services to the customer.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalPaymentSucceededRequest
	*/
	WebhooksTerminalPaymentSucceeded(ctx context.Context) ApiWebhooksTerminalPaymentSucceededRequest

	// WebhooksTerminalPaymentSucceededExecute executes the request
	WebhooksTerminalPaymentSucceededExecute(r ApiWebhooksTerminalPaymentSucceededRequest) (*http.Response, error)

	/*
	WebhooksTerminalRefundSucceeded Terminal refund succeeded

	Sent when a card-present refund is approved by the terminal.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalRefundSucceededRequest
	*/
	WebhooksTerminalRefundSucceeded(ctx context.Context) ApiWebhooksTerminalRefundSucceededRequest

	// WebhooksTerminalRefundSucceededExecute executes the request
	WebhooksTerminalRefundSucceededExecute(r ApiWebhooksTerminalRefundSucceededRequest) (*http.Response, error)

	/*
	WebhooksTerminalVoidSucceeded Terminal void succeeded

	Sent when a card-present void is approved by the terminal. The original transaction has been reversed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksTerminalVoidSucceededRequest
	*/
	WebhooksTerminalVoidSucceeded(ctx context.Context) ApiWebhooksTerminalVoidSucceededRequest

	// WebhooksTerminalVoidSucceededExecute executes the request
	WebhooksTerminalVoidSucceededExecute(r ApiWebhooksTerminalVoidSucceededRequest) (*http.Response, error)

	/*
	WebhooksUsageEventIngested Usage event ingested

	Sent after a usage event is accepted by ingestion. This is a high-volume event and should be subscribed to explicitly with narrow endpoint filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksUsageEventIngestedRequest
	*/
	WebhooksUsageEventIngested(ctx context.Context) ApiWebhooksUsageEventIngestedRequest

	// WebhooksUsageEventIngestedExecute executes the request
	WebhooksUsageEventIngestedExecute(r ApiWebhooksUsageEventIngestedRequest) (*http.Response, error)

	/*
	WebhooksUsageEventRejected Usage event rejected

	Sent when a usage event is rejected during ingestion (validation failure) or asynchronously by ClickHouse (malformed message, ~60s delay).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksUsageEventRejectedRequest
	*/
	WebhooksUsageEventRejected(ctx context.Context) ApiWebhooksUsageEventRejectedRequest

	// WebhooksUsageEventRejectedExecute executes the request
	WebhooksUsageEventRejectedExecute(r ApiWebhooksUsageEventRejectedRequest) (*http.Response, error)

	/*
	WebhooksUsageInvoiceCreated Usage invoice created

	Sent when finalized usage records are converted into an invoice for a customer subscription.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksUsageInvoiceCreatedRequest
	*/
	WebhooksUsageInvoiceCreated(ctx context.Context) ApiWebhooksUsageInvoiceCreatedRequest

	// WebhooksUsageInvoiceCreatedExecute executes the request
	WebhooksUsageInvoiceCreatedExecute(r ApiWebhooksUsageInvoiceCreatedRequest) (*http.Response, error)

	/*
	WebhooksUsagePeriodFinalized Usage period finalized

	Sent when a billing period is finalized with final usage quantities. An invoice will be generated shortly after.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksUsagePeriodFinalizedRequest
	*/
	WebhooksUsagePeriodFinalized(ctx context.Context) ApiWebhooksUsagePeriodFinalizedRequest

	// WebhooksUsagePeriodFinalizedExecute executes the request
	WebhooksUsagePeriodFinalizedExecute(r ApiWebhooksUsagePeriodFinalizedRequest) (*http.Response, error)

	/*
	WebhooksUsageThresholdReached Usage threshold reached

	Sent when a customer's usage crosses a configured threshold percentage on a meter. Use this to alert customers approaching or exceeding their included allowance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksUsageThresholdReachedRequest
	*/
	WebhooksUsageThresholdReached(ctx context.Context) ApiWebhooksUsageThresholdReachedRequest

	// WebhooksUsageThresholdReachedExecute executes the request
	WebhooksUsageThresholdReachedExecute(r ApiWebhooksUsageThresholdReachedRequest) (*http.Response, error)

	/*
	WebhooksVoidCreated Void created

	Sent when a void is initiated for an unsettled transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksVoidCreatedRequest
	*/
	WebhooksVoidCreated(ctx context.Context) ApiWebhooksVoidCreatedRequest

	// WebhooksVoidCreatedExecute executes the request
	WebhooksVoidCreatedExecute(r ApiWebhooksVoidCreatedRequest) (*http.Response, error)

	/*
	WebhooksVoidFailed Void failed

	Sent when a void fails (e.g., transaction already settled). Consider using a refund instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksVoidFailedRequest
	*/
	WebhooksVoidFailed(ctx context.Context) ApiWebhooksVoidFailedRequest

	// WebhooksVoidFailedExecute executes the request
	WebhooksVoidFailedExecute(r ApiWebhooksVoidFailedRequest) (*http.Response, error)

	/*
	WebhooksVoidSucceeded Void succeeded

	Sent when a void is successfully processed by the payment gateway. The original transaction has been canceled.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiWebhooksVoidSucceededRequest
	*/
	WebhooksVoidSucceeded(ctx context.Context) ApiWebhooksVoidSucceededRequest

	// WebhooksVoidSucceededExecute executes the request
	WebhooksVoidSucceededExecute(r ApiWebhooksVoidSucceededRequest) (*http.Response, error)
}

// WebhooksAPIService WebhooksAPI service
type WebhooksAPIService service

type ApiWebhooksCheckoutSessionCompletedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCheckoutSessionCompletedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCheckoutSessionCompletedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCheckoutSessionCompletedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCheckoutSessionCompletedExecute(r)
}

/*
WebhooksCheckoutSessionCompleted Checkout session completed

Sent when a checkout session is successfully completed. **Action required:** Fulfill the customer's order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCheckoutSessionCompletedRequest
*/
func (a *WebhooksAPIService) WebhooksCheckoutSessionCompleted(ctx context.Context) ApiWebhooksCheckoutSessionCompletedRequest {
	return ApiWebhooksCheckoutSessionCompletedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCheckoutSessionCompletedExecute(r ApiWebhooksCheckoutSessionCompletedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCheckoutSessionCompleted")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout.session.completed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCheckoutSessionExpiredRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCheckoutSessionExpiredRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCheckoutSessionExpiredRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCheckoutSessionExpiredRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCheckoutSessionExpiredExecute(r)
}

/*
WebhooksCheckoutSessionExpired Checkout session expired

Sent when a checkout session expires before completion.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCheckoutSessionExpiredRequest
*/
func (a *WebhooksAPIService) WebhooksCheckoutSessionExpired(ctx context.Context) ApiWebhooksCheckoutSessionExpiredRequest {
	return ApiWebhooksCheckoutSessionExpiredRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCheckoutSessionExpiredExecute(r ApiWebhooksCheckoutSessionExpiredRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCheckoutSessionExpired")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/checkout.session.expired"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCollectionFailedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCollectionFailedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCollectionFailedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCollectionFailedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCollectionFailedExecute(r)
}

/*
WebhooksCollectionFailed Collection failed

Sent when a Direct Debit collection fails (e.g. insufficient funds). Contains the Bacs reason code. **Action required:** RevKeen Recovery may retry or fall back to a stored card per your policy.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCollectionFailedRequest
*/
func (a *WebhooksAPIService) WebhooksCollectionFailed(ctx context.Context) ApiWebhooksCollectionFailedRequest {
	return ApiWebhooksCollectionFailedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCollectionFailedExecute(r ApiWebhooksCollectionFailedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCollectionFailed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collection.failed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCollectionIndemnityClaimedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCollectionIndemnityClaimedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCollectionIndemnityClaimedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCollectionIndemnityClaimedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCollectionIndemnityClaimedExecute(r)
}

/*
WebhooksCollectionIndemnityClaimed Indemnity claimed

Sent when a customer raises a Direct Debit indemnity claim (the Bacs chargeback equivalent) and the collection is reversed. **Action required:** Review the disputed collection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCollectionIndemnityClaimedRequest
*/
func (a *WebhooksAPIService) WebhooksCollectionIndemnityClaimed(ctx context.Context) ApiWebhooksCollectionIndemnityClaimedRequest {
	return ApiWebhooksCollectionIndemnityClaimedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCollectionIndemnityClaimedExecute(r ApiWebhooksCollectionIndemnityClaimedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCollectionIndemnityClaimed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collection.indemnity_claimed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCollectionNoticeSentRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCollectionNoticeSentRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCollectionNoticeSentRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCollectionNoticeSentRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCollectionNoticeSentExecute(r)
}

/*
WebhooksCollectionNoticeSent Advance notice sent

Sent when the mandatory advance notice for an upcoming collection has been issued to the customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCollectionNoticeSentRequest
*/
func (a *WebhooksAPIService) WebhooksCollectionNoticeSent(ctx context.Context) ApiWebhooksCollectionNoticeSentRequest {
	return ApiWebhooksCollectionNoticeSentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCollectionNoticeSentExecute(r ApiWebhooksCollectionNoticeSentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCollectionNoticeSent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collection.notice_sent"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCollectionScheduledRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCollectionScheduledRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCollectionScheduledRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCollectionScheduledRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCollectionScheduledExecute(r)
}

/*
WebhooksCollectionScheduled Collection scheduled

Sent when a Direct Debit collection is scheduled against an active mandate. Includes the working-day collection date.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCollectionScheduledRequest
*/
func (a *WebhooksAPIService) WebhooksCollectionScheduled(ctx context.Context) ApiWebhooksCollectionScheduledRequest {
	return ApiWebhooksCollectionScheduledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCollectionScheduledExecute(r ApiWebhooksCollectionScheduledRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCollectionScheduled")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collection.scheduled"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCollectionSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCollectionSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCollectionSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCollectionSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCollectionSucceededExecute(r)
}

/*
WebhooksCollectionSucceeded Collection succeeded

Sent when a Direct Debit collection is successfully taken. **Action required:** Provision goods or services to your customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCollectionSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksCollectionSucceeded(ctx context.Context) ApiWebhooksCollectionSucceededRequest {
	return ApiWebhooksCollectionSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCollectionSucceededExecute(r ApiWebhooksCollectionSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCollectionSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/collection.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCreditNoteCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCreditNoteCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCreditNoteCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCreditNoteCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCreditNoteCreatedExecute(r)
}

/*
WebhooksCreditNoteCreated Credit note created

Sent when a credit note is created for an invoice.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCreditNoteCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksCreditNoteCreated(ctx context.Context) ApiWebhooksCreditNoteCreatedRequest {
	return ApiWebhooksCreditNoteCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCreditNoteCreatedExecute(r ApiWebhooksCreditNoteCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCreditNoteCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_note.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCreditNoteIssuedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCreditNoteIssuedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCreditNoteIssuedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCreditNoteIssuedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCreditNoteIssuedExecute(r)
}

/*
WebhooksCreditNoteIssued Credit note issued

Sent when a credit note is issued (finalized) and its PDF is generated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCreditNoteIssuedRequest
*/
func (a *WebhooksAPIService) WebhooksCreditNoteIssued(ctx context.Context) ApiWebhooksCreditNoteIssuedRequest {
	return ApiWebhooksCreditNoteIssuedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCreditNoteIssuedExecute(r ApiWebhooksCreditNoteIssuedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCreditNoteIssued")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_note.issued"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCreditNoteVoidedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCreditNoteVoidedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCreditNoteVoidedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCreditNoteVoidedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCreditNoteVoidedExecute(r)
}

/*
WebhooksCreditNoteVoided Credit note voided

Sent when a credit note is voided (accounting reversal).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCreditNoteVoidedRequest
*/
func (a *WebhooksAPIService) WebhooksCreditNoteVoided(ctx context.Context) ApiWebhooksCreditNoteVoidedRequest {
	return ApiWebhooksCreditNoteVoidedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCreditNoteVoidedExecute(r ApiWebhooksCreditNoteVoidedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCreditNoteVoided")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/credit_note.voided"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCustomerCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCustomerCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCustomerCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCustomerCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCustomerCreatedExecute(r)
}

/*
WebhooksCustomerCreated Customer created

Sent when a new customer is created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCustomerCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksCustomerCreated(ctx context.Context) ApiWebhooksCustomerCreatedRequest {
	return ApiWebhooksCustomerCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCustomerCreatedExecute(r ApiWebhooksCustomerCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCustomerCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksCustomerUpdatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksCustomerUpdatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksCustomerUpdatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksCustomerUpdatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksCustomerUpdatedExecute(r)
}

/*
WebhooksCustomerUpdated Customer updated

Sent when customer information is updated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksCustomerUpdatedRequest
*/
func (a *WebhooksAPIService) WebhooksCustomerUpdated(ctx context.Context) ApiWebhooksCustomerUpdatedRequest {
	return ApiWebhooksCustomerUpdatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksCustomerUpdatedExecute(r ApiWebhooksCustomerUpdatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksCustomerUpdated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer.updated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksInvoiceCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksInvoiceCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksInvoiceCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksInvoiceCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksInvoiceCreatedExecute(r)
}

/*
WebhooksInvoiceCreated Invoice created

Sent when a new invoice is created, either manually or from a subscription renewal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksInvoiceCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksInvoiceCreated(ctx context.Context) ApiWebhooksInvoiceCreatedRequest {
	return ApiWebhooksInvoiceCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksInvoiceCreatedExecute(r ApiWebhooksInvoiceCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksInvoiceCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoice.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksInvoiceOverdueRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksInvoiceOverdueRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksInvoiceOverdueRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksInvoiceOverdueRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksInvoiceOverdueExecute(r)
}

/*
WebhooksInvoiceOverdue Invoice overdue

Sent when an invoice becomes overdue. Consider sending payment reminders.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksInvoiceOverdueRequest
*/
func (a *WebhooksAPIService) WebhooksInvoiceOverdue(ctx context.Context) ApiWebhooksInvoiceOverdueRequest {
	return ApiWebhooksInvoiceOverdueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksInvoiceOverdueExecute(r ApiWebhooksInvoiceOverdueRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksInvoiceOverdue")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoice.overdue"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksInvoicePaidRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksInvoicePaidRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksInvoicePaidRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksInvoicePaidRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksInvoicePaidExecute(r)
}

/*
WebhooksInvoicePaid Invoice paid

Sent when an invoice is fully paid. **Action required:** Fulfill the order or activate the subscription.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksInvoicePaidRequest
*/
func (a *WebhooksAPIService) WebhooksInvoicePaid(ctx context.Context) ApiWebhooksInvoicePaidRequest {
	return ApiWebhooksInvoicePaidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksInvoicePaidExecute(r ApiWebhooksInvoicePaidRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksInvoicePaid")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoice.paid"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMandateActivatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMandateActivatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMandateActivatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMandateActivatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMandateActivatedExecute(r)
}

/*
WebhooksMandateActivated Mandate activated

Sent when a mandate is lodged and active. **Action required:** You can now collect from this customer via Direct Debit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMandateActivatedRequest
*/
func (a *WebhooksAPIService) WebhooksMandateActivated(ctx context.Context) ApiWebhooksMandateActivatedRequest {
	return ApiWebhooksMandateActivatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMandateActivatedExecute(r ApiWebhooksMandateActivatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMandateActivated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandate.activated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMandateAuddisRejectedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMandateAuddisRejectedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMandateAuddisRejectedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMandateAuddisRejectedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMandateAuddisRejectedExecute(r)
}

/*
WebhooksMandateAuddisRejected Mandate rejected (AUDDIS)

Sent when the customer's bank rejects the mandate setup (AUDDIS). **Action required:** The mandate cannot be used — re-collect bank details or use another payment method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMandateAuddisRejectedRequest
*/
func (a *WebhooksAPIService) WebhooksMandateAuddisRejected(ctx context.Context) ApiWebhooksMandateAuddisRejectedRequest {
	return ApiWebhooksMandateAuddisRejectedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMandateAuddisRejectedExecute(r ApiWebhooksMandateAuddisRejectedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMandateAuddisRejected")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandate.auddis_rejected"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMandateCancelledRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMandateCancelledRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMandateCancelledRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMandateCancelledRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMandateCancelledExecute(r)
}

/*
WebhooksMandateCancelled Mandate cancelled

Sent when a mandate is cancelled (by you, the customer, or the bank). It can no longer be collected against.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMandateCancelledRequest
*/
func (a *WebhooksAPIService) WebhooksMandateCancelled(ctx context.Context) ApiWebhooksMandateCancelledRequest {
	return ApiWebhooksMandateCancelledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMandateCancelledExecute(r ApiWebhooksMandateCancelledRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMandateCancelled")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandate.cancelled"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMandateCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMandateCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMandateCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMandateCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMandateCreatedExecute(r)
}

/*
WebhooksMandateCreated Mandate created

Sent when a Direct Debit mandate is created and submitted to Bacs for lodgement. The mandate is not yet collectable — wait for `mandate.activated`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMandateCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksMandateCreated(ctx context.Context) ApiWebhooksMandateCreatedRequest {
	return ApiWebhooksMandateCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMandateCreatedExecute(r ApiWebhooksMandateCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMandateCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandate.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMandateSuspendedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMandateSuspendedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMandateSuspendedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMandateSuspendedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMandateSuspendedExecute(r)
}

/*
WebhooksMandateSuspended Mandate suspended

Sent when a mandate is suspended and can no longer be collected against until reinstated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMandateSuspendedRequest
*/
func (a *WebhooksAPIService) WebhooksMandateSuspended(ctx context.Context) ApiWebhooksMandateSuspendedRequest {
	return ApiWebhooksMandateSuspendedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMandateSuspendedExecute(r ApiWebhooksMandateSuspendedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMandateSuspended")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mandate.suspended"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterArchivedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterArchivedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterArchivedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterArchivedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterArchivedExecute(r)
}

/*
WebhooksMeterArchived Meter archived

Sent when a usage meter is archived and can no longer be attached to new pricing.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterArchivedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterArchived(ctx context.Context) ApiWebhooksMeterArchivedRequest {
	return ApiWebhooksMeterArchivedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterArchivedExecute(r ApiWebhooksMeterArchivedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterArchived")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter.archived"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterCreatedExecute(r)
}

/*
WebhooksMeterCreated Meter created

Sent when a usage meter is created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterCreated(ctx context.Context) ApiWebhooksMeterCreatedRequest {
	return ApiWebhooksMeterCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterCreatedExecute(r ApiWebhooksMeterCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterPriceCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterPriceCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterPriceCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterPriceCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterPriceCreatedExecute(r)
}

/*
WebhooksMeterPriceCreated Meter price created

Sent when a meter price is created for a product price.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterPriceCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterPriceCreated(ctx context.Context) ApiWebhooksMeterPriceCreatedRequest {
	return ApiWebhooksMeterPriceCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterPriceCreatedExecute(r ApiWebhooksMeterPriceCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterPriceCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter_price.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterPriceDeactivatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterPriceDeactivatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterPriceDeactivatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterPriceDeactivatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterPriceDeactivatedExecute(r)
}

/*
WebhooksMeterPriceDeactivated Meter price deactivated

Sent when a meter price is deactivated and no longer available for new subscriptions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterPriceDeactivatedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterPriceDeactivated(ctx context.Context) ApiWebhooksMeterPriceDeactivatedRequest {
	return ApiWebhooksMeterPriceDeactivatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterPriceDeactivatedExecute(r ApiWebhooksMeterPriceDeactivatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterPriceDeactivated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter_price.deactivated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterPriceUpdatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterPriceUpdatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterPriceUpdatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterPriceUpdatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterPriceUpdatedExecute(r)
}

/*
WebhooksMeterPriceUpdated Meter price updated

Sent when a meter price is updated. The payload includes changed attributes when available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterPriceUpdatedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterPriceUpdated(ctx context.Context) ApiWebhooksMeterPriceUpdatedRequest {
	return ApiWebhooksMeterPriceUpdatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterPriceUpdatedExecute(r ApiWebhooksMeterPriceUpdatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterPriceUpdated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter_price.updated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksMeterUpdatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksMeterUpdatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksMeterUpdatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksMeterUpdatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksMeterUpdatedExecute(r)
}

/*
WebhooksMeterUpdated Meter updated

Sent when a usage meter is updated. The payload includes changed attributes when available.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksMeterUpdatedRequest
*/
func (a *WebhooksAPIService) WebhooksMeterUpdated(ctx context.Context) ApiWebhooksMeterUpdatedRequest {
	return ApiWebhooksMeterUpdatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksMeterUpdatedExecute(r ApiWebhooksMeterUpdatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksMeterUpdated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/meter.updated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksOrderCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksOrderCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksOrderCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksOrderCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksOrderCreatedExecute(r)
}

/*
WebhooksOrderCreated Order created

Sent when a new order is created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksOrderCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksOrderCreated(ctx context.Context) ApiWebhooksOrderCreatedRequest {
	return ApiWebhooksOrderCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksOrderCreatedExecute(r ApiWebhooksOrderCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksOrderCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksOrderFulfilledRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksOrderFulfilledRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksOrderFulfilledRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksOrderFulfilledRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksOrderFulfilledExecute(r)
}

/*
WebhooksOrderFulfilled Order fulfilled

Sent when an order is marked as fulfilled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksOrderFulfilledRequest
*/
func (a *WebhooksAPIService) WebhooksOrderFulfilled(ctx context.Context) ApiWebhooksOrderFulfilledRequest {
	return ApiWebhooksOrderFulfilledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksOrderFulfilledExecute(r ApiWebhooksOrderFulfilledRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksOrderFulfilled")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order.fulfilled"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksOrderPaidRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksOrderPaidRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksOrderPaidRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksOrderPaidRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksOrderPaidExecute(r)
}

/*
WebhooksOrderPaid Order paid

Sent when an order is fully paid. **Action required:** Begin fulfillment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksOrderPaidRequest
*/
func (a *WebhooksAPIService) WebhooksOrderPaid(ctx context.Context) ApiWebhooksOrderPaidRequest {
	return ApiWebhooksOrderPaidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksOrderPaidExecute(r ApiWebhooksOrderPaidRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksOrderPaid")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/order.paid"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksPaymentFailedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksPaymentFailedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksPaymentFailedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksPaymentFailedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksPaymentFailedExecute(r)
}

/*
WebhooksPaymentFailed Payment failed

Sent when a payment attempt fails. Contains failure reason and suggested next steps.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksPaymentFailedRequest
*/
func (a *WebhooksAPIService) WebhooksPaymentFailed(ctx context.Context) ApiWebhooksPaymentFailedRequest {
	return ApiWebhooksPaymentFailedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksPaymentFailedExecute(r ApiWebhooksPaymentFailedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksPaymentFailed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment.failed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksPaymentSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksPaymentSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksPaymentSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksPaymentSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksPaymentSucceededExecute(r)
}

/*
WebhooksPaymentSucceeded Payment succeeded

Sent when a payment is successfully captured. **Action required:** Provision goods or services to your customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksPaymentSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksPaymentSucceeded(ctx context.Context) ApiWebhooksPaymentSucceededRequest {
	return ApiWebhooksPaymentSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksPaymentSucceededExecute(r ApiWebhooksPaymentSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksPaymentSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksRefundCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksRefundCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksRefundCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksRefundCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksRefundCreatedExecute(r)
}

/*
WebhooksRefundCreated Refund created

Sent when a refund is initiated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksRefundCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksRefundCreated(ctx context.Context) ApiWebhooksRefundCreatedRequest {
	return ApiWebhooksRefundCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksRefundCreatedExecute(r ApiWebhooksRefundCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksRefundCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/refund.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksRefundSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksRefundSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksRefundSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksRefundSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksRefundSucceededExecute(r)
}

/*
WebhooksRefundSucceeded Refund succeeded

Sent when a refund is successfully processed by the payment gateway.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksRefundSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksRefundSucceeded(ctx context.Context) ApiWebhooksRefundSucceededRequest {
	return ApiWebhooksRefundSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksRefundSucceededExecute(r ApiWebhooksRefundSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksRefundSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/refund.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksSettlementCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksSettlementCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksSettlementCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksSettlementCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksSettlementCreatedExecute(r)
}

/*
WebhooksSettlementCreated Settlement created

Sent when collected Direct Debit funds are settled and a payout is created. Gross settlement — the bureau fee is reported separately and never netted from the payout.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksSettlementCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksSettlementCreated(ctx context.Context) ApiWebhooksSettlementCreatedRequest {
	return ApiWebhooksSettlementCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksSettlementCreatedExecute(r ApiWebhooksSettlementCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksSettlementCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/settlement.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksSubscriptionActivatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksSubscriptionActivatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksSubscriptionActivatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksSubscriptionActivatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksSubscriptionActivatedExecute(r)
}

/*
WebhooksSubscriptionActivated Subscription activated

Sent when a subscription becomes active after successful payment. **Action required:** Grant access to your service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksSubscriptionActivatedRequest
*/
func (a *WebhooksAPIService) WebhooksSubscriptionActivated(ctx context.Context) ApiWebhooksSubscriptionActivatedRequest {
	return ApiWebhooksSubscriptionActivatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksSubscriptionActivatedExecute(r ApiWebhooksSubscriptionActivatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksSubscriptionActivated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription.activated"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksSubscriptionCanceledRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksSubscriptionCanceledRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksSubscriptionCanceledRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksSubscriptionCanceledRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksSubscriptionCanceledExecute(r)
}

/*
WebhooksSubscriptionCanceled Subscription canceled

Sent when a subscription is canceled. Access continues until the end of the billing period.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksSubscriptionCanceledRequest
*/
func (a *WebhooksAPIService) WebhooksSubscriptionCanceled(ctx context.Context) ApiWebhooksSubscriptionCanceledRequest {
	return ApiWebhooksSubscriptionCanceledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksSubscriptionCanceledExecute(r ApiWebhooksSubscriptionCanceledRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksSubscriptionCanceled")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription.canceled"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksSubscriptionCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksSubscriptionCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksSubscriptionCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksSubscriptionCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksSubscriptionCreatedExecute(r)
}

/*
WebhooksSubscriptionCreated Subscription created

Sent when a new subscription is created (before activation).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksSubscriptionCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksSubscriptionCreated(ctx context.Context) ApiWebhooksSubscriptionCreatedRequest {
	return ApiWebhooksSubscriptionCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksSubscriptionCreatedExecute(r ApiWebhooksSubscriptionCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksSubscriptionCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksSubscriptionRenewedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksSubscriptionRenewedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksSubscriptionRenewedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksSubscriptionRenewedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksSubscriptionRenewedExecute(r)
}

/*
WebhooksSubscriptionRenewed Subscription renewed

Sent when a subscription successfully renews. A new invoice has been created and paid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksSubscriptionRenewedRequest
*/
func (a *WebhooksAPIService) WebhooksSubscriptionRenewed(ctx context.Context) ApiWebhooksSubscriptionRenewedRequest {
	return ApiWebhooksSubscriptionRenewedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksSubscriptionRenewedExecute(r ApiWebhooksSubscriptionRenewedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksSubscriptionRenewed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription.renewed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalPaymentCancelledRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalPaymentCancelledRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalPaymentCancelledRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalPaymentCancelledRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalPaymentCancelledExecute(r)
}

/*
WebhooksTerminalPaymentCancelled Terminal payment cancelled

Sent when a card-present payment is cancelled by the merchant before terminal response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalPaymentCancelledRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalPaymentCancelled(ctx context.Context) ApiWebhooksTerminalPaymentCancelledRequest {
	return ApiWebhooksTerminalPaymentCancelledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalPaymentCancelledExecute(r ApiWebhooksTerminalPaymentCancelledRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalPaymentCancelled")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_payment.cancelled"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalPaymentDeclinedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalPaymentDeclinedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalPaymentDeclinedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalPaymentDeclinedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalPaymentDeclinedExecute(r)
}

/*
WebhooksTerminalPaymentDeclined Terminal payment declined

Sent when a card-present payment is declined by the terminal or card issuer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalPaymentDeclinedRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalPaymentDeclined(ctx context.Context) ApiWebhooksTerminalPaymentDeclinedRequest {
	return ApiWebhooksTerminalPaymentDeclinedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalPaymentDeclinedExecute(r ApiWebhooksTerminalPaymentDeclinedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalPaymentDeclined")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_payment.declined"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalPaymentErrorRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalPaymentErrorRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalPaymentErrorRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalPaymentErrorRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalPaymentErrorExecute(r)
}

/*
WebhooksTerminalPaymentError Terminal payment error

Sent when a card-present payment fails due to timeout, terminal error, or connection loss. Check `failure_reason` field: `timeout`, `terminal_error`, or `connection_lost`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalPaymentErrorRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalPaymentError(ctx context.Context) ApiWebhooksTerminalPaymentErrorRequest {
	return ApiWebhooksTerminalPaymentErrorRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalPaymentErrorExecute(r ApiWebhooksTerminalPaymentErrorRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalPaymentError")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_payment.error"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalPaymentRequestedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalPaymentRequestedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalPaymentRequestedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalPaymentRequestedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalPaymentRequestedExecute(r)
}

/*
WebhooksTerminalPaymentRequested Terminal payment requested

Sent when a card-present payment is dispatched to a terminal device. The payment is in-progress and awaiting terminal response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalPaymentRequestedRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalPaymentRequested(ctx context.Context) ApiWebhooksTerminalPaymentRequestedRequest {
	return ApiWebhooksTerminalPaymentRequestedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalPaymentRequestedExecute(r ApiWebhooksTerminalPaymentRequestedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalPaymentRequested")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_payment.requested"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalPaymentSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalPaymentSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalPaymentSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalPaymentSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalPaymentSucceededExecute(r)
}

/*
WebhooksTerminalPaymentSucceeded Terminal payment succeeded

Sent when a card-present payment is approved by the terminal. **Action required:** Provide goods or services to the customer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalPaymentSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalPaymentSucceeded(ctx context.Context) ApiWebhooksTerminalPaymentSucceededRequest {
	return ApiWebhooksTerminalPaymentSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalPaymentSucceededExecute(r ApiWebhooksTerminalPaymentSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalPaymentSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_payment.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalRefundSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalRefundSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalRefundSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalRefundSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalRefundSucceededExecute(r)
}

/*
WebhooksTerminalRefundSucceeded Terminal refund succeeded

Sent when a card-present refund is approved by the terminal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalRefundSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalRefundSucceeded(ctx context.Context) ApiWebhooksTerminalRefundSucceededRequest {
	return ApiWebhooksTerminalRefundSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalRefundSucceededExecute(r ApiWebhooksTerminalRefundSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalRefundSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_refund.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksTerminalVoidSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksTerminalVoidSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksTerminalVoidSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksTerminalVoidSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksTerminalVoidSucceededExecute(r)
}

/*
WebhooksTerminalVoidSucceeded Terminal void succeeded

Sent when a card-present void is approved by the terminal. The original transaction has been reversed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksTerminalVoidSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksTerminalVoidSucceeded(ctx context.Context) ApiWebhooksTerminalVoidSucceededRequest {
	return ApiWebhooksTerminalVoidSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksTerminalVoidSucceededExecute(r ApiWebhooksTerminalVoidSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksTerminalVoidSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/terminal_void.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksUsageEventIngestedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksUsageEventIngestedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksUsageEventIngestedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksUsageEventIngestedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksUsageEventIngestedExecute(r)
}

/*
WebhooksUsageEventIngested Usage event ingested

Sent after a usage event is accepted by ingestion. This is a high-volume event and should be subscribed to explicitly with narrow endpoint filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksUsageEventIngestedRequest
*/
func (a *WebhooksAPIService) WebhooksUsageEventIngested(ctx context.Context) ApiWebhooksUsageEventIngestedRequest {
	return ApiWebhooksUsageEventIngestedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksUsageEventIngestedExecute(r ApiWebhooksUsageEventIngestedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksUsageEventIngested")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage.event.ingested"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksUsageEventRejectedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksUsageEventRejectedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksUsageEventRejectedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksUsageEventRejectedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksUsageEventRejectedExecute(r)
}

/*
WebhooksUsageEventRejected Usage event rejected

Sent when a usage event is rejected during ingestion (validation failure) or asynchronously by ClickHouse (malformed message, ~60s delay).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksUsageEventRejectedRequest
*/
func (a *WebhooksAPIService) WebhooksUsageEventRejected(ctx context.Context) ApiWebhooksUsageEventRejectedRequest {
	return ApiWebhooksUsageEventRejectedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksUsageEventRejectedExecute(r ApiWebhooksUsageEventRejectedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksUsageEventRejected")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage.event.rejected"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksUsageInvoiceCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksUsageInvoiceCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksUsageInvoiceCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksUsageInvoiceCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksUsageInvoiceCreatedExecute(r)
}

/*
WebhooksUsageInvoiceCreated Usage invoice created

Sent when finalized usage records are converted into an invoice for a customer subscription.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksUsageInvoiceCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksUsageInvoiceCreated(ctx context.Context) ApiWebhooksUsageInvoiceCreatedRequest {
	return ApiWebhooksUsageInvoiceCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksUsageInvoiceCreatedExecute(r ApiWebhooksUsageInvoiceCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksUsageInvoiceCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage.invoice.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksUsagePeriodFinalizedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksUsagePeriodFinalizedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksUsagePeriodFinalizedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksUsagePeriodFinalizedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksUsagePeriodFinalizedExecute(r)
}

/*
WebhooksUsagePeriodFinalized Usage period finalized

Sent when a billing period is finalized with final usage quantities. An invoice will be generated shortly after.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksUsagePeriodFinalizedRequest
*/
func (a *WebhooksAPIService) WebhooksUsagePeriodFinalized(ctx context.Context) ApiWebhooksUsagePeriodFinalizedRequest {
	return ApiWebhooksUsagePeriodFinalizedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksUsagePeriodFinalizedExecute(r ApiWebhooksUsagePeriodFinalizedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksUsagePeriodFinalized")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage.period_finalized"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksUsageThresholdReachedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksUsageThresholdReachedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksUsageThresholdReachedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksUsageThresholdReachedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksUsageThresholdReachedExecute(r)
}

/*
WebhooksUsageThresholdReached Usage threshold reached

Sent when a customer's usage crosses a configured threshold percentage on a meter. Use this to alert customers approaching or exceeding their included allowance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksUsageThresholdReachedRequest
*/
func (a *WebhooksAPIService) WebhooksUsageThresholdReached(ctx context.Context) ApiWebhooksUsageThresholdReachedRequest {
	return ApiWebhooksUsageThresholdReachedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksUsageThresholdReachedExecute(r ApiWebhooksUsageThresholdReachedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksUsageThresholdReached")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage.threshold.reached"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksVoidCreatedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksVoidCreatedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksVoidCreatedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksVoidCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksVoidCreatedExecute(r)
}

/*
WebhooksVoidCreated Void created

Sent when a void is initiated for an unsettled transaction.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksVoidCreatedRequest
*/
func (a *WebhooksAPIService) WebhooksVoidCreated(ctx context.Context) ApiWebhooksVoidCreatedRequest {
	return ApiWebhooksVoidCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksVoidCreatedExecute(r ApiWebhooksVoidCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksVoidCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/void.created"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksVoidFailedRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksVoidFailedRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksVoidFailedRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksVoidFailedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksVoidFailedExecute(r)
}

/*
WebhooksVoidFailed Void failed

Sent when a void fails (e.g., transaction already settled). Consider using a refund instead.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksVoidFailedRequest
*/
func (a *WebhooksAPIService) WebhooksVoidFailed(ctx context.Context) ApiWebhooksVoidFailedRequest {
	return ApiWebhooksVoidFailedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksVoidFailedExecute(r ApiWebhooksVoidFailedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksVoidFailed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/void.failed"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksVoidSucceededRequest struct {
	ctx context.Context
	ApiService WebhooksAPI
	webhookEvent *WebhookEvent
}

func (r ApiWebhooksVoidSucceededRequest) WebhookEvent(webhookEvent WebhookEvent) ApiWebhooksVoidSucceededRequest {
	r.webhookEvent = &webhookEvent
	return r
}

func (r ApiWebhooksVoidSucceededRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksVoidSucceededExecute(r)
}

/*
WebhooksVoidSucceeded Void succeeded

Sent when a void is successfully processed by the payment gateway. The original transaction has been canceled.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksVoidSucceededRequest
*/
func (a *WebhooksAPIService) WebhooksVoidSucceeded(ctx context.Context) ApiWebhooksVoidSucceededRequest {
	return ApiWebhooksVoidSucceededRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *WebhooksAPIService) WebhooksVoidSucceededExecute(r ApiWebhooksVoidSucceededRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksAPIService.WebhooksVoidSucceeded")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/void.succeeded"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.webhookEvent
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
