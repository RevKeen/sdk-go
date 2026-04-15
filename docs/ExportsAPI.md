# \ExportsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportsCreate**](ExportsAPI.md#ExportsCreate) | **Post** /exports | Create an export
[**ExportsRetrieve**](ExportsAPI.md#ExportsRetrieve) | **Get** /exports/{id} | Get export status



## ExportsCreate

> ExportCreateResponse ExportsCreate(ctx).CreateExportRequest(createExportRequest).Execute()

Create an export



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
	createExportRequest := *openapiclient.NewCreateExportRequest("ResourceType_example") // CreateExportRequest | Export configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsCreate(context.Background()).CreateExportRequest(createExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsCreate`: ExportCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExportRequest** | [**CreateExportRequest**](CreateExportRequest.md) | Export configuration | 

### Return type

[**ExportCreateResponse**](ExportCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsRetrieve

> ExportRetrieveResponse ExportsRetrieve(ctx, id).Execute()

Get export status



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Export job UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsRetrieve`: ExportRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Export job UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportRetrieveResponse**](ExportRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

