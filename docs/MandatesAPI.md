# \MandatesAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MandatesCancel**](MandatesAPI.md#MandatesCancel) | **Post** /mandates/{id}/cancel | Cancel a Direct Debit mandate
[**MandatesGet**](MandatesAPI.md#MandatesGet) | **Get** /mandates/{id} | Retrieve a Direct Debit mandate
[**MandatesList**](MandatesAPI.md#MandatesList) | **Get** /mandates | List Direct Debit mandates
[**MandatesReinstate**](MandatesAPI.md#MandatesReinstate) | **Post** /mandates/{id}/reinstate | Reinstate a suspended Direct Debit mandate
[**MandatesScheduleCollection**](MandatesAPI.md#MandatesScheduleCollection) | **Post** /mandates/{id}/collections | Schedule a one-off Direct Debit collection
[**MandatesSuspend**](MandatesAPI.md#MandatesSuspend) | **Post** /mandates/{id}/suspend | Suspend a Direct Debit mandate



## MandatesCancel

> MandateActionResponse MandatesCancel(ctx, id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()

Cancel a Direct Debit mandate



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesCancel(context.Background(), id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesCancel`: MandateActionResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **mandateActionRequest** | [**MandateActionRequest**](MandateActionRequest.md) |  |

### Return type

[**MandateActionResponse**](MandateActionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MandatesGet

> MandateResponse MandatesGet(ctx, id).Execute()

Retrieve a Direct Debit mandate



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesGet`: MandateResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MandateResponse**](MandateResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MandatesList

> MandateListResponse MandatesList(ctx).CustomerId(customerId).Status(status).Execute()

List Direct Debit mandates



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
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter mandates by customer (UUID) (optional)
	status := "status_example" // string | Filter mandates by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesList(context.Background()).CustomerId(customerId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesList`: MandateListResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMandatesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Filter mandates by customer (UUID) |
 **status** | **string** | Filter mandates by status |

### Return type

[**MandateListResponse**](MandateListResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MandatesReinstate

> MandateActionResponse MandatesReinstate(ctx, id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()

Reinstate a suspended Direct Debit mandate



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesReinstate(context.Background(), id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesReinstate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesReinstate`: MandateActionResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesReinstate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesReinstateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **mandateActionRequest** | [**MandateActionRequest**](MandateActionRequest.md) |  |

### Return type

[**MandateActionResponse**](MandateActionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MandatesScheduleCollection

> ScheduleCollectionResponse MandatesScheduleCollection(ctx, id).IdempotencyKey(idempotencyKey).ScheduleCollectionRequest(scheduleCollectionRequest).Execute()

Schedule a one-off Direct Debit collection



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	scheduleCollectionRequest := *openapiclient.NewScheduleCollectionRequest(int64(2500), "GBP", "SourceType_example", "SourceId_example") // ScheduleCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesScheduleCollection(context.Background(), id).IdempotencyKey(idempotencyKey).ScheduleCollectionRequest(scheduleCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesScheduleCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesScheduleCollection`: ScheduleCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesScheduleCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesScheduleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **scheduleCollectionRequest** | [**ScheduleCollectionRequest**](ScheduleCollectionRequest.md) |  |

### Return type

[**ScheduleCollectionResponse**](ScheduleCollectionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MandatesSuspend

> MandateActionResponse MandatesSuspend(ctx, id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()

Suspend a Direct Debit mandate



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesSuspend(context.Background(), id).IdempotencyKey(idempotencyKey).MandateActionRequest(mandateActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesSuspend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesSuspend`: MandateActionResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesSuspend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **mandateActionRequest** | [**MandateActionRequest**](MandateActionRequest.md) |  |

### Return type

[**MandateActionResponse**](MandateActionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
