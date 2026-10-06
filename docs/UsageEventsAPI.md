# \UsageEventsAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageEventsGet**](UsageEventsAPI.md#UsageEventsGet) | **Get** /usage-events/{id} | Get a usage event
[**UsageEventsIngest**](UsageEventsAPI.md#UsageEventsIngest) | **Post** /usage-events | Ingest usage events
[**UsageEventsList**](UsageEventsAPI.md#UsageEventsList) | **Get** /usage-events | Query usage events



## UsageEventsGet

> UsageEventGetResponse UsageEventsGet(ctx, id).Execute()

Get a usage event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Usage event ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageEventsAPI.UsageEventsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageEventsAPI.UsageEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageEventsGet`: UsageEventGetResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageEventsAPI.UsageEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Usage event ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUsageEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsageEventGetResponse**](UsageEventGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageEventsIngest

> BatchIngestResult UsageEventsIngest(ctx).IdempotencyKey(idempotencyKey).UsageIngestRequest(usageIngestRequest).Execute()

Ingest usage events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	usageIngestRequest := *openapiclient.NewUsageIngestRequest("api_call", "evt_unique_123", []openapiclient.UsageEvent{*openapiclient.NewUsageEvent("api_call", "evt_unique_123")}) // UsageIngestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageEventsAPI.UsageEventsIngest(context.Background()).IdempotencyKey(idempotencyKey).UsageIngestRequest(usageIngestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageEventsAPI.UsageEventsIngest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageEventsIngest`: BatchIngestResult
	fmt.Fprintf(os.Stdout, "Response from `UsageEventsAPI.UsageEventsIngest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageEventsIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **usageIngestRequest** | [**UsageIngestRequest**](UsageIngestRequest.md) |  |

### Return type

[**BatchIngestResult**](BatchIngestResult.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageEventsList

> UsageEventListResponse UsageEventsList(ctx).MeterId(meterId).CustomerId(customerId).ExternalCustomerId(externalCustomerId).SubscriptionId(subscriptionId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()

Query usage events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	meterId := "meterId_example" // string | Filter by meter ID (optional)
	customerId := "customerId_example" // string | Filter by customer ID (optional)
	externalCustomerId := "externalCustomerId_example" // string | Filter by external customer ID (optional)
	subscriptionId := "subscriptionId_example" // string | Filter by subscription ID (optional)
	startTime := time.Now() // time.Time | Start of time range (ISO 8601) (optional)
	endTime := time.Now() // time.Time | End of time range (ISO 8601) (optional)
	limit := "100" // string | Max results (default 100, max 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageEventsAPI.UsageEventsList(context.Background()).MeterId(meterId).CustomerId(customerId).ExternalCustomerId(externalCustomerId).SubscriptionId(subscriptionId).StartTime(startTime).EndTime(endTime).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageEventsAPI.UsageEventsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageEventsList`: UsageEventListResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageEventsAPI.UsageEventsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meterId** | **string** | Filter by meter ID |
 **customerId** | **string** | Filter by customer ID |
 **externalCustomerId** | **string** | Filter by external customer ID |
 **subscriptionId** | **string** | Filter by subscription ID |
 **startTime** | **time.Time** | Start of time range (ISO 8601) |
 **endTime** | **time.Time** | End of time range (ISO 8601) |
 **limit** | **string** | Max results (default 100, max 100) |

### Return type

[**UsageEventListResponse**](UsageEventListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
