# \WebhookDeliveriesAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookDeliveriesGet**](WebhookDeliveriesAPI.md#WebhookDeliveriesGet) | **Get** /webhook-deliveries/{id} | Retrieve a webhook delivery
[**WebhookDeliveriesList**](WebhookDeliveriesAPI.md#WebhookDeliveriesList) | **Get** /webhook-deliveries | List webhook deliveries
[**WebhookDeliveriesRetry**](WebhookDeliveriesAPI.md#WebhookDeliveriesRetry) | **Post** /webhook-deliveries/{id}/retry | Retry a webhook delivery



## WebhookDeliveriesGet

> WebhookDeliveryResponse WebhookDeliveriesGet(ctx, id).Execute()

Retrieve a webhook delivery



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook delivery ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookDeliveriesAPI.WebhookDeliveriesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveriesAPI.WebhookDeliveriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeliveriesGet`: WebhookDeliveryResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveriesAPI.WebhookDeliveriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook delivery ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookDeliveryResponse**](WebhookDeliveryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveriesList

> WebhookDeliveryListResponse WebhookDeliveriesList(ctx).EndpointId(endpointId).EventId(eventId).Status(status).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List webhook deliveries



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
	endpointId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by webhook endpoint ID (optional)
	eventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by source event ID (optional)
	status := "status_example" // string | Filter by delivery status (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return deliveries created before the row with this ID. (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return deliveries created after the row with this ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookDeliveriesAPI.WebhookDeliveriesList(context.Background()).EndpointId(endpointId).EventId(eventId).Status(status).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveriesAPI.WebhookDeliveriesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeliveriesList`: WebhookDeliveryListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveriesAPI.WebhookDeliveriesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveriesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointId** | **string** | Filter by webhook endpoint ID | 
 **eventId** | **string** | Filter by source event ID | 
 **status** | **string** | Filter by delivery status | 
 **limit** | **int32** |  | [default to 20]
 **startingAfter** | **string** | Cursor — return deliveries created before the row with this ID. | 
 **endingBefore** | **string** | Cursor — return deliveries created after the row with this ID. | 

### Return type

[**WebhookDeliveryListResponse**](WebhookDeliveryListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookDeliveriesRetry

> WebhookDeliveryRetryResponse WebhookDeliveriesRetry(ctx, id).Execute()

Retry a webhook delivery



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook delivery ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookDeliveriesAPI.WebhookDeliveriesRetry(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookDeliveriesAPI.WebhookDeliveriesRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookDeliveriesRetry`: WebhookDeliveryRetryResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookDeliveriesAPI.WebhookDeliveriesRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Webhook delivery ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookDeliveriesRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookDeliveryRetryResponse**](WebhookDeliveryRetryResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

