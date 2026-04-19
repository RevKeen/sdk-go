# \EventsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsGet**](EventsAPI.md#EventsGet) | **Get** /events/{id} | Retrieve an event
[**EventsList**](EventsAPI.md#EventsList) | **Get** /events | List events
[**EventsResend**](EventsAPI.md#EventsResend) | **Post** /events/{id}/resend | Resend webhook for an event



## EventsGet

> Event EventsGet(ctx, id).Execute()

Retrieve an event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsGet`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsList

> EventListResponse EventsList(ctx).Type_(type_).Types(types).CustomerId(customerId).InvoiceId(invoiceId).SubscriptionId(subscriptionId).OrderId(orderId).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List events



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
	type_ := "invoice.paid" // string | Filter by event type (e.g., invoice.paid) (optional)
	types := "invoice.paid,payment.succeeded" // string | Filter by multiple event types (comma-separated) (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer ID (optional)
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by invoice ID (optional)
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by subscription ID (optional)
	orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by order ID (optional)
	createdGte := float32(8.14) // float32 | Filter by created_at >= (Unix timestamp) (optional)
	createdLte := float32(8.14) // float32 | Filter by created_at <= (Unix timestamp) (optional)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination - return results after this event ID (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination - return results before this event ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsList(context.Background()).Type_(type_).Types(types).CustomerId(customerId).InvoiceId(invoiceId).SubscriptionId(subscriptionId).OrderId(orderId).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsList`: EventListResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Filter by event type (e.g., invoice.paid) | 
 **types** | **string** | Filter by multiple event types (comma-separated) | 
 **customerId** | **string** | Filter by customer ID | 
 **invoiceId** | **string** | Filter by invoice ID | 
 **subscriptionId** | **string** | Filter by subscription ID | 
 **orderId** | **string** | Filter by order ID | 
 **createdGte** | **float32** | Filter by created_at &gt;&#x3D; (Unix timestamp) | 
 **createdLte** | **float32** | Filter by created_at &lt;&#x3D; (Unix timestamp) | 
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **startingAfter** | **string** | Cursor for pagination - return results after this event ID | 
 **endingBefore** | **string** | Cursor for pagination - return results before this event ID | 

### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsResend

> ResendWebhookResponse EventsResend(ctx, id).Execute()

Resend webhook for an event



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsResend(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsResend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsResend`: ResendWebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsResend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Event ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsResendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResendWebhookResponse**](ResendWebhookResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

