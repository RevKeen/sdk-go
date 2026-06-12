# \WebhooksAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksCheckoutSessionCompleted**](WebhooksAPI.md#WebhooksCheckoutSessionCompleted) | **Post** /checkout.session.completed | Checkout session completed
[**WebhooksCheckoutSessionExpired**](WebhooksAPI.md#WebhooksCheckoutSessionExpired) | **Post** /checkout.session.expired | Checkout session expired
[**WebhooksCollectionFailed**](WebhooksAPI.md#WebhooksCollectionFailed) | **Post** /collection.failed | Collection failed
[**WebhooksCollectionIndemnityClaimed**](WebhooksAPI.md#WebhooksCollectionIndemnityClaimed) | **Post** /collection.indemnity_claimed | Indemnity claimed
[**WebhooksCollectionNoticeSent**](WebhooksAPI.md#WebhooksCollectionNoticeSent) | **Post** /collection.notice_sent | Advance notice sent
[**WebhooksCollectionScheduled**](WebhooksAPI.md#WebhooksCollectionScheduled) | **Post** /collection.scheduled | Collection scheduled
[**WebhooksCollectionSucceeded**](WebhooksAPI.md#WebhooksCollectionSucceeded) | **Post** /collection.succeeded | Collection succeeded
[**WebhooksCreditNoteCreated**](WebhooksAPI.md#WebhooksCreditNoteCreated) | **Post** /credit_note.created | Credit note created
[**WebhooksCreditNoteIssued**](WebhooksAPI.md#WebhooksCreditNoteIssued) | **Post** /credit_note.issued | Credit note issued
[**WebhooksCreditNoteVoided**](WebhooksAPI.md#WebhooksCreditNoteVoided) | **Post** /credit_note.voided | Credit note voided
[**WebhooksCustomerCreated**](WebhooksAPI.md#WebhooksCustomerCreated) | **Post** /customer.created | Customer created
[**WebhooksCustomerUpdated**](WebhooksAPI.md#WebhooksCustomerUpdated) | **Post** /customer.updated | Customer updated
[**WebhooksInvoiceCreated**](WebhooksAPI.md#WebhooksInvoiceCreated) | **Post** /invoice.created | Invoice created
[**WebhooksInvoiceOverdue**](WebhooksAPI.md#WebhooksInvoiceOverdue) | **Post** /invoice.overdue | Invoice overdue
[**WebhooksInvoicePaid**](WebhooksAPI.md#WebhooksInvoicePaid) | **Post** /invoice.paid | Invoice paid
[**WebhooksMandateActivated**](WebhooksAPI.md#WebhooksMandateActivated) | **Post** /mandate.activated | Mandate activated
[**WebhooksMandateAuddisRejected**](WebhooksAPI.md#WebhooksMandateAuddisRejected) | **Post** /mandate.auddis_rejected | Mandate rejected (AUDDIS)
[**WebhooksMandateCancelled**](WebhooksAPI.md#WebhooksMandateCancelled) | **Post** /mandate.cancelled | Mandate cancelled
[**WebhooksMandateCreated**](WebhooksAPI.md#WebhooksMandateCreated) | **Post** /mandate.created | Mandate created
[**WebhooksMandateSuspended**](WebhooksAPI.md#WebhooksMandateSuspended) | **Post** /mandate.suspended | Mandate suspended
[**WebhooksOrderCreated**](WebhooksAPI.md#WebhooksOrderCreated) | **Post** /order.created | Order created
[**WebhooksOrderFulfilled**](WebhooksAPI.md#WebhooksOrderFulfilled) | **Post** /order.fulfilled | Order fulfilled
[**WebhooksOrderPaid**](WebhooksAPI.md#WebhooksOrderPaid) | **Post** /order.paid | Order paid
[**WebhooksPaymentFailed**](WebhooksAPI.md#WebhooksPaymentFailed) | **Post** /payment.failed | Payment failed
[**WebhooksPaymentSucceeded**](WebhooksAPI.md#WebhooksPaymentSucceeded) | **Post** /payment.succeeded | Payment succeeded
[**WebhooksRefundCreated**](WebhooksAPI.md#WebhooksRefundCreated) | **Post** /refund.created | Refund created
[**WebhooksRefundSucceeded**](WebhooksAPI.md#WebhooksRefundSucceeded) | **Post** /refund.succeeded | Refund succeeded
[**WebhooksSettlementCreated**](WebhooksAPI.md#WebhooksSettlementCreated) | **Post** /settlement.created | Settlement created
[**WebhooksSubscriptionActivated**](WebhooksAPI.md#WebhooksSubscriptionActivated) | **Post** /subscription.activated | Subscription activated
[**WebhooksSubscriptionCanceled**](WebhooksAPI.md#WebhooksSubscriptionCanceled) | **Post** /subscription.canceled | Subscription canceled
[**WebhooksSubscriptionCreated**](WebhooksAPI.md#WebhooksSubscriptionCreated) | **Post** /subscription.created | Subscription created
[**WebhooksSubscriptionRenewed**](WebhooksAPI.md#WebhooksSubscriptionRenewed) | **Post** /subscription.renewed | Subscription renewed
[**WebhooksTerminalPaymentCancelled**](WebhooksAPI.md#WebhooksTerminalPaymentCancelled) | **Post** /terminal_payment.cancelled | Terminal payment cancelled
[**WebhooksTerminalPaymentDeclined**](WebhooksAPI.md#WebhooksTerminalPaymentDeclined) | **Post** /terminal_payment.declined | Terminal payment declined
[**WebhooksTerminalPaymentError**](WebhooksAPI.md#WebhooksTerminalPaymentError) | **Post** /terminal_payment.error | Terminal payment error
[**WebhooksTerminalPaymentRequested**](WebhooksAPI.md#WebhooksTerminalPaymentRequested) | **Post** /terminal_payment.requested | Terminal payment requested
[**WebhooksTerminalPaymentSucceeded**](WebhooksAPI.md#WebhooksTerminalPaymentSucceeded) | **Post** /terminal_payment.succeeded | Terminal payment succeeded
[**WebhooksTerminalRefundSucceeded**](WebhooksAPI.md#WebhooksTerminalRefundSucceeded) | **Post** /terminal_refund.succeeded | Terminal refund succeeded
[**WebhooksTerminalVoidSucceeded**](WebhooksAPI.md#WebhooksTerminalVoidSucceeded) | **Post** /terminal_void.succeeded | Terminal void succeeded
[**WebhooksUsageEventRejected**](WebhooksAPI.md#WebhooksUsageEventRejected) | **Post** /usage.event.rejected | Usage event rejected
[**WebhooksUsagePeriodFinalized**](WebhooksAPI.md#WebhooksUsagePeriodFinalized) | **Post** /usage.period_finalized | Usage period finalized
[**WebhooksUsageThresholdReached**](WebhooksAPI.md#WebhooksUsageThresholdReached) | **Post** /usage.threshold.reached | Usage threshold reached
[**WebhooksVoidCreated**](WebhooksAPI.md#WebhooksVoidCreated) | **Post** /void.created | Void created
[**WebhooksVoidFailed**](WebhooksAPI.md#WebhooksVoidFailed) | **Post** /void.failed | Void failed
[**WebhooksVoidSucceeded**](WebhooksAPI.md#WebhooksVoidSucceeded) | **Post** /void.succeeded | Void succeeded



## WebhooksCheckoutSessionCompleted

> WebhooksCheckoutSessionCompleted(ctx).WebhookEvent(webhookEvent).Execute()

Checkout session completed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCheckoutSessionCompleted(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCheckoutSessionCompleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCheckoutSessionCompletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCheckoutSessionExpired

> WebhooksCheckoutSessionExpired(ctx).WebhookEvent(webhookEvent).Execute()

Checkout session expired



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCheckoutSessionExpired(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCheckoutSessionExpired``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCheckoutSessionExpiredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCollectionFailed

> WebhooksCollectionFailed(ctx).WebhookEvent(webhookEvent).Execute()

Collection failed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCollectionFailed(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCollectionFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCollectionFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCollectionIndemnityClaimed

> WebhooksCollectionIndemnityClaimed(ctx).WebhookEvent(webhookEvent).Execute()

Indemnity claimed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCollectionIndemnityClaimed(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCollectionIndemnityClaimed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCollectionIndemnityClaimedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCollectionNoticeSent

> WebhooksCollectionNoticeSent(ctx).WebhookEvent(webhookEvent).Execute()

Advance notice sent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCollectionNoticeSent(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCollectionNoticeSent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCollectionNoticeSentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCollectionScheduled

> WebhooksCollectionScheduled(ctx).WebhookEvent(webhookEvent).Execute()

Collection scheduled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCollectionScheduled(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCollectionScheduled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCollectionScheduledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCollectionSucceeded

> WebhooksCollectionSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Collection succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCollectionSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCollectionSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCollectionSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCreditNoteCreated

> WebhooksCreditNoteCreated(ctx).WebhookEvent(webhookEvent).Execute()

Credit note created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCreditNoteCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCreditNoteCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreditNoteCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCreditNoteIssued

> WebhooksCreditNoteIssued(ctx).WebhookEvent(webhookEvent).Execute()

Credit note issued



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCreditNoteIssued(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCreditNoteIssued``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreditNoteIssuedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCreditNoteVoided

> WebhooksCreditNoteVoided(ctx).WebhookEvent(webhookEvent).Execute()

Credit note voided



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCreditNoteVoided(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCreditNoteVoided``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCreditNoteVoidedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCustomerCreated

> WebhooksCustomerCreated(ctx).WebhookEvent(webhookEvent).Execute()

Customer created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCustomerCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCustomerCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCustomerCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksCustomerUpdated

> WebhooksCustomerUpdated(ctx).WebhookEvent(webhookEvent).Execute()

Customer updated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksCustomerUpdated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksCustomerUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksCustomerUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksInvoiceCreated

> WebhooksInvoiceCreated(ctx).WebhookEvent(webhookEvent).Execute()

Invoice created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksInvoiceCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksInvoiceCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksInvoiceCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksInvoiceOverdue

> WebhooksInvoiceOverdue(ctx).WebhookEvent(webhookEvent).Execute()

Invoice overdue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksInvoiceOverdue(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksInvoiceOverdue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksInvoiceOverdueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksInvoicePaid

> WebhooksInvoicePaid(ctx).WebhookEvent(webhookEvent).Execute()

Invoice paid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksInvoicePaid(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksInvoicePaid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksInvoicePaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksMandateActivated

> WebhooksMandateActivated(ctx).WebhookEvent(webhookEvent).Execute()

Mandate activated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksMandateActivated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksMandateActivated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksMandateActivatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksMandateAuddisRejected

> WebhooksMandateAuddisRejected(ctx).WebhookEvent(webhookEvent).Execute()

Mandate rejected (AUDDIS)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksMandateAuddisRejected(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksMandateAuddisRejected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksMandateAuddisRejectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksMandateCancelled

> WebhooksMandateCancelled(ctx).WebhookEvent(webhookEvent).Execute()

Mandate cancelled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksMandateCancelled(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksMandateCancelled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksMandateCancelledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksMandateCreated

> WebhooksMandateCreated(ctx).WebhookEvent(webhookEvent).Execute()

Mandate created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksMandateCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksMandateCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksMandateCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksMandateSuspended

> WebhooksMandateSuspended(ctx).WebhookEvent(webhookEvent).Execute()

Mandate suspended



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksMandateSuspended(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksMandateSuspended``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksMandateSuspendedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksOrderCreated

> WebhooksOrderCreated(ctx).WebhookEvent(webhookEvent).Execute()

Order created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksOrderCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksOrderCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksOrderCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksOrderFulfilled

> WebhooksOrderFulfilled(ctx).WebhookEvent(webhookEvent).Execute()

Order fulfilled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksOrderFulfilled(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksOrderFulfilled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksOrderFulfilledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksOrderPaid

> WebhooksOrderPaid(ctx).WebhookEvent(webhookEvent).Execute()

Order paid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksOrderPaid(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksOrderPaid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksOrderPaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPaymentFailed

> WebhooksPaymentFailed(ctx).WebhookEvent(webhookEvent).Execute()

Payment failed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksPaymentFailed(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksPaymentFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPaymentFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksPaymentSucceeded

> WebhooksPaymentSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Payment succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksPaymentSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksPaymentSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksPaymentSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksRefundCreated

> WebhooksRefundCreated(ctx).WebhookEvent(webhookEvent).Execute()

Refund created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksRefundCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksRefundCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksRefundCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksRefundSucceeded

> WebhooksRefundSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Refund succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksRefundSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksRefundSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksRefundSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSettlementCreated

> WebhooksSettlementCreated(ctx).WebhookEvent(webhookEvent).Execute()

Settlement created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksSettlementCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksSettlementCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSettlementCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSubscriptionActivated

> WebhooksSubscriptionActivated(ctx).WebhookEvent(webhookEvent).Execute()

Subscription activated



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksSubscriptionActivated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksSubscriptionActivated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSubscriptionActivatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSubscriptionCanceled

> WebhooksSubscriptionCanceled(ctx).WebhookEvent(webhookEvent).Execute()

Subscription canceled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksSubscriptionCanceled(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksSubscriptionCanceled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSubscriptionCanceledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSubscriptionCreated

> WebhooksSubscriptionCreated(ctx).WebhookEvent(webhookEvent).Execute()

Subscription created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksSubscriptionCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksSubscriptionCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSubscriptionCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksSubscriptionRenewed

> WebhooksSubscriptionRenewed(ctx).WebhookEvent(webhookEvent).Execute()

Subscription renewed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksSubscriptionRenewed(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksSubscriptionRenewed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksSubscriptionRenewedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalPaymentCancelled

> WebhooksTerminalPaymentCancelled(ctx).WebhookEvent(webhookEvent).Execute()

Terminal payment cancelled



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalPaymentCancelled(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalPaymentCancelled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalPaymentCancelledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalPaymentDeclined

> WebhooksTerminalPaymentDeclined(ctx).WebhookEvent(webhookEvent).Execute()

Terminal payment declined



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalPaymentDeclined(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalPaymentDeclined``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalPaymentDeclinedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalPaymentError

> WebhooksTerminalPaymentError(ctx).WebhookEvent(webhookEvent).Execute()

Terminal payment error



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalPaymentError(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalPaymentError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalPaymentErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalPaymentRequested

> WebhooksTerminalPaymentRequested(ctx).WebhookEvent(webhookEvent).Execute()

Terminal payment requested



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalPaymentRequested(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalPaymentRequested``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalPaymentRequestedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalPaymentSucceeded

> WebhooksTerminalPaymentSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Terminal payment succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalPaymentSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalPaymentSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalPaymentSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalRefundSucceeded

> WebhooksTerminalRefundSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Terminal refund succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalRefundSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalRefundSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalRefundSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksTerminalVoidSucceeded

> WebhooksTerminalVoidSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Terminal void succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksTerminalVoidSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksTerminalVoidSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksTerminalVoidSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksUsageEventRejected

> WebhooksUsageEventRejected(ctx).WebhookEvent(webhookEvent).Execute()

Usage event rejected



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksUsageEventRejected(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksUsageEventRejected``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksUsageEventRejectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksUsagePeriodFinalized

> WebhooksUsagePeriodFinalized(ctx).WebhookEvent(webhookEvent).Execute()

Usage period finalized



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksUsagePeriodFinalized(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksUsagePeriodFinalized``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksUsagePeriodFinalizedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksUsageThresholdReached

> WebhooksUsageThresholdReached(ctx).WebhookEvent(webhookEvent).Execute()

Usage threshold reached



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksUsageThresholdReached(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksUsageThresholdReached``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksUsageThresholdReachedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksVoidCreated

> WebhooksVoidCreated(ctx).WebhookEvent(webhookEvent).Execute()

Void created



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksVoidCreated(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksVoidCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksVoidCreatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksVoidFailed

> WebhooksVoidFailed(ctx).WebhookEvent(webhookEvent).Execute()

Void failed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksVoidFailed(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksVoidFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksVoidFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksVoidSucceeded

> WebhooksVoidSucceeded(ctx).WebhookEvent(webhookEvent).Execute()

Void succeeded



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	webhookEvent :=  // WebhookEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.WebhooksVoidSucceeded(context.Background()).WebhookEvent(webhookEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksVoidSucceeded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksVoidSucceededRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEvent** | [**WebhookEvent**](WebhookEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

