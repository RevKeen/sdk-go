# \InvoicesAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicesCreate**](InvoicesAPI.md#InvoicesCreate) | **Post** /invoices | Create invoice
[**InvoicesFinalize**](InvoicesAPI.md#InvoicesFinalize) | **Post** /invoices/{id}/finalize | Finalize an invoice
[**InvoicesGet**](InvoicesAPI.md#InvoicesGet) | **Get** /invoices/{id} | Get invoice
[**InvoicesList**](InvoicesAPI.md#InvoicesList) | **Get** /invoices | List invoices
[**InvoicesMarginEstimate**](InvoicesAPI.md#InvoicesMarginEstimate) | **Post** /invoices/{id}/margin-estimate | Estimate invoice net by payment rail
[**InvoicesSend**](InvoicesAPI.md#InvoicesSend) | **Post** /invoices/{id}/send | Send an invoice
[**InvoicesUpdate**](InvoicesAPI.md#InvoicesUpdate) | **Patch** /invoices/{id} | Update invoice
[**InvoicesVoid**](InvoicesAPI.md#InvoicesVoid) | **Post** /invoices/{id}/void | Void an invoice



## InvoicesCreate

> InvoiceResponse InvoicesCreate(ctx).InvoicesCreateRequest(invoicesCreateRequest).Execute()

Create invoice



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
	invoicesCreateRequest := *openapiclient.NewInvoicesCreateRequest("CustomerUuid_example", int32(123)) // InvoicesCreateRequest | Invoice creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesCreate(context.Background()).InvoicesCreateRequest(invoicesCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesCreate`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoicesCreateRequest** | [**InvoicesCreateRequest**](InvoicesCreateRequest.md) | Invoice creation details | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesFinalize

> InvoiceResponse InvoicesFinalize(ctx, id).InvoicesFinalizeRequest(invoicesFinalizeRequest).Execute()

Finalize an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoicesFinalizeRequest := *openapiclient.NewInvoicesFinalizeRequest() // InvoicesFinalizeRequest | Finalization options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesFinalize(context.Background(), id).InvoicesFinalizeRequest(invoicesFinalizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesFinalize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesFinalize`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesFinalize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesFinalizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesFinalizeRequest** | [**InvoicesFinalizeRequest**](InvoicesFinalizeRequest.md) | Finalization options | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesGet

> InvoiceResponse InvoicesGet(ctx, id).Execute()

Get invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesGet`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesList

> InvoiceListResponse InvoicesList(ctx).Status(status).CustomerId(customerId).Limit(limit).Offset(offset).Execute()

List invoices



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
	status := "status_example" // string | Filter by invoice status (optional)
	customerId := "customerId_example" // string | Filter by customer ID (optional)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	offset := int32(56) // int32 | Number of results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesList(context.Background()).Status(status).CustomerId(customerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesList`: InvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by invoice status | 
 **customerId** | **string** | Filter by customer ID | 
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **offset** | **int32** | Number of results to skip | [default to 0]

### Return type

[**InvoiceListResponse**](InvoiceListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesMarginEstimate

> InvoicesMarginEstimate200Response InvoicesMarginEstimate(ctx, id).InvoicesMarginEstimateRequest(invoicesMarginEstimateRequest).Execute()

Estimate invoice net by payment rail



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoicesMarginEstimateRequest := *openapiclient.NewInvoicesMarginEstimateRequest() // InvoicesMarginEstimateRequest | Margin estimate inputs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesMarginEstimate(context.Background(), id).InvoicesMarginEstimateRequest(invoicesMarginEstimateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesMarginEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesMarginEstimate`: InvoicesMarginEstimate200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesMarginEstimate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesMarginEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesMarginEstimateRequest** | [**InvoicesMarginEstimateRequest**](InvoicesMarginEstimateRequest.md) | Margin estimate inputs | 

### Return type

[**InvoicesMarginEstimate200Response**](InvoicesMarginEstimate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesSend

> InvoiceResponse InvoicesSend(ctx, id).InvoicesSendRequest(invoicesSendRequest).Execute()

Send an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoicesSendRequest := *openapiclient.NewInvoicesSendRequest() // InvoicesSendRequest | Send options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesSend(context.Background(), id).InvoicesSendRequest(invoicesSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesSend`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesSendRequest** | [**InvoicesSendRequest**](InvoicesSendRequest.md) | Send options | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesUpdate

> InvoiceResponse InvoicesUpdate(ctx, id).InvoicesUpdateRequest(invoicesUpdateRequest).Execute()

Update invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoicesUpdateRequest := *openapiclient.NewInvoicesUpdateRequest() // InvoicesUpdateRequest | Invoice update details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesUpdate(context.Background(), id).InvoicesUpdateRequest(invoicesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesUpdate`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesUpdateRequest** | [**InvoicesUpdateRequest**](InvoicesUpdateRequest.md) | Invoice update details | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesVoid

> InvoiceResponse InvoicesVoid(ctx, id).InvoicesVoidRequest(invoicesVoidRequest).Execute()

Void an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invoicesVoidRequest := *openapiclient.NewInvoicesVoidRequest() // InvoicesVoidRequest | Void details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesVoid(context.Background(), id).InvoicesVoidRequest(invoicesVoidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesVoid`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesVoidRequest** | [**InvoicesVoidRequest**](InvoicesVoidRequest.md) | Void details | 

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

