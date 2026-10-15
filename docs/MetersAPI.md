# \MetersAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetersCreate**](MetersAPI.md#MetersCreate) | **Post** /meters | Create a meter
[**MetersGet**](MetersAPI.md#MetersGet) | **Get** /meters/{id} | Get a meter
[**MetersList**](MetersAPI.md#MetersList) | **Get** /meters | List meters
[**MetersQuantities**](MetersAPI.md#MetersQuantities) | **Get** /meters/{id}/quantities | Get meter quantities
[**MetersUpdate**](MetersAPI.md#MetersUpdate) | **Patch** /meters/{id} | Update a meter



## MetersCreate

> CreateMeterResponse MetersCreate(ctx).CreateMeterRequest(createMeterRequest).Execute()

Create a meter



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
	createMeterRequest := *openapiclient.NewCreateMeterRequest("API Calls", "api_call") // CreateMeterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersCreate(context.Background()).CreateMeterRequest(createMeterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersCreate`: CreateMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMeterRequest** | [**CreateMeterRequest**](CreateMeterRequest.md) |  |

### Return type

[**CreateMeterResponse**](CreateMeterResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersGet

> MeterGetResponse MetersGet(ctx, id).Execute()

Get a meter



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
	id := "e3e70682-c209-1cac-a29f-6fbed82c07cd" // string | Meter UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersGet`: MeterGetResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter UUID. |

### Other Parameters

Other parameters are passed through a pointer to a apiMetersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MeterGetResponse**](MeterGetResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersList

> MeterListResponse MetersList(ctx).IncludeArchived(includeArchived).Limit(limit).Offset(offset).Execute()

List meters



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
	includeArchived := "false" // string | Include archived meters (optional)
	limit := "50" // string | Max results (default 50, max 100) (optional)
	offset := "0" // string | Pagination offset (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersList(context.Background()).IncludeArchived(includeArchived).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersList`: MeterListResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeArchived** | **string** | Include archived meters |
 **limit** | **string** | Max results (default 50, max 100) |
 **offset** | **string** | Pagination offset |

### Return type

[**MeterListResponse**](MeterListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersQuantities

> MeterQuantitiesResponse MetersQuantities(ctx, id).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Interval(interval).CustomerId(customerId).ExternalCustomerId(externalCustomerId).SubscriptionId(subscriptionId).Timezone(timezone).Execute()

Get meter quantities



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
	id := "e3e70682-c209-1cac-a29f-6fbed82c07cd" // string | Meter UUID.
	startTimestamp := time.Now() // time.Time | Start timestamp (ISO 8601, UTC)
	endTimestamp := time.Now() // time.Time | End timestamp (ISO 8601, UTC)
	interval := "day" // string | Bucket width. UTC date_trunc semantics; week is ISO (Monday).
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer UUID (optional)
	externalCustomerId := "externalCustomerId_example" // string | Filter by the merchant's external customer ID (optional)
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by subscription UUID (optional)
	timezone := "UTC" // string | Only UTC is supported. Omit or pass UTC. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersQuantities(context.Background(), id).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).Interval(interval).CustomerId(customerId).ExternalCustomerId(externalCustomerId).SubscriptionId(subscriptionId).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersQuantities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersQuantities`: MeterQuantitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersQuantities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter UUID. |

### Other Parameters

Other parameters are passed through a pointer to a apiMetersQuantitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTimestamp** | **time.Time** | Start timestamp (ISO 8601, UTC) |
 **endTimestamp** | **time.Time** | End timestamp (ISO 8601, UTC) |
 **interval** | **string** | Bucket width. UTC date_trunc semantics; week is ISO (Monday). |
 **customerId** | **string** | Filter by customer UUID |
 **externalCustomerId** | **string** | Filter by the merchant&#39;s external customer ID |
 **subscriptionId** | **string** | Filter by subscription UUID |
 **timezone** | **string** | Only UTC is supported. Omit or pass UTC. |

### Return type

[**MeterQuantitiesResponse**](MeterQuantitiesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersUpdate

> UpdateMeterResponse MetersUpdate(ctx, id).UpdateMeterRequest(updateMeterRequest).Execute()

Update a meter



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
	id := "e3e70682-c209-1cac-a29f-6fbed82c07cd" // string | Meter UUID.
	updateMeterRequest := *openapiclient.NewUpdateMeterRequest() // UpdateMeterRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersUpdate(context.Background(), id).UpdateMeterRequest(updateMeterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersUpdate`: UpdateMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter UUID. |

### Other Parameters

Other parameters are passed through a pointer to a apiMetersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMeterRequest** | [**UpdateMeterRequest**](UpdateMeterRequest.md) |  |

### Return type

[**UpdateMeterResponse**](UpdateMeterResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
