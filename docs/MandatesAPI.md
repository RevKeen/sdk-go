# \MandatesAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MandatesCancel**](MandatesAPI.md#MandatesCancel) | **Post** /mandates/{id}/cancel | Cancel a Direct Debit mandate
[**MandatesCreate**](MandatesAPI.md#MandatesCreate) | **Post** /mandates | Create a Direct Debit mandate
[**MandatesGet**](MandatesAPI.md#MandatesGet) | **Get** /mandates/{id} | Retrieve a Direct Debit mandate
[**MandatesList**](MandatesAPI.md#MandatesList) | **Get** /mandates | List Direct Debit mandates
[**MandatesReinstate**](MandatesAPI.md#MandatesReinstate) | **Post** /mandates/{id}/reinstate | Reinstate a suspended Direct Debit mandate
[**MandatesScheduleCollection**](MandatesAPI.md#MandatesScheduleCollection) | **Post** /mandates/{id}/collections | Schedule a one-off Direct Debit collection
[**MandatesSuspend**](MandatesAPI.md#MandatesSuspend) | **Post** /mandates/{id}/suspend | Suspend a Direct Debit mandate



## MandatesCancel

> MandateActionResponse MandatesCancel(ctx, id).MandateActionRequest(mandateActionRequest).Execute()

Cancel a Direct Debit mandate



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
	id := "id_example" // string | Mandate ID
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesCancel(context.Background(), id).MandateActionRequest(mandateActionRequest).Execute()
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
**id** | **string** | Mandate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## MandatesCreate

> MandateResponse MandatesCreate(ctx).CreateMandateRequest(createMandateRequest).Execute()

Create a Direct Debit mandate



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
	createMandateRequest := *openapiclient.NewCreateMandateRequest("CustomerId_example", "Jane Smith", "20-00-00", "55779911") // CreateMandateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesCreate(context.Background()).CreateMandateRequest(createMandateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatesAPI.MandatesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MandatesCreate`: MandateResponse
	fmt.Fprintf(os.Stdout, "Response from `MandatesAPI.MandatesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMandatesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMandateRequest** | [**CreateMandateRequest**](CreateMandateRequest.md) |  | 

### Return type

[**MandateResponse**](MandateResponse.md)

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
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	id := "id_example" // string | Mandate ID

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
**id** | **string** | Mandate ID | 

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
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	customerId := "customerId_example" // string | Filter mandates by customer (optional)
	status := "status_example" // string | Filter mandates by status (e.g. active, suspended, cancelled) (optional)

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
 **customerId** | **string** | Filter mandates by customer | 
 **status** | **string** | Filter mandates by status (e.g. active, suspended, cancelled) | 

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

> MandateActionResponse MandatesReinstate(ctx, id).MandateActionRequest(mandateActionRequest).Execute()

Reinstate a suspended Direct Debit mandate



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
	id := "id_example" // string | Mandate ID
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesReinstate(context.Background(), id).MandateActionRequest(mandateActionRequest).Execute()
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
**id** | **string** | Mandate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesReinstateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> ScheduleCollectionResponse MandatesScheduleCollection(ctx, id).ScheduleCollectionRequest(scheduleCollectionRequest).Execute()

Schedule a one-off Direct Debit collection



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
	id := "id_example" // string | Mandate ID
	scheduleCollectionRequest := *openapiclient.NewScheduleCollectionRequest(int32(2500), "GBP", "SourceType_example", "SourceId_example") // ScheduleCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesScheduleCollection(context.Background(), id).ScheduleCollectionRequest(scheduleCollectionRequest).Execute()
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
**id** | **string** | Mandate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesScheduleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

> MandateActionResponse MandatesSuspend(ctx, id).MandateActionRequest(mandateActionRequest).Execute()

Suspend a Direct Debit mandate



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
	id := "id_example" // string | Mandate ID
	mandateActionRequest := *openapiclient.NewMandateActionRequest() // MandateActionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatesAPI.MandatesSuspend(context.Background(), id).MandateActionRequest(mandateActionRequest).Execute()
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
**id** | **string** | Mandate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMandatesSuspendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

