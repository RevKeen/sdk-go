# \CreditNotesAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditNotesCreate**](CreditNotesAPI.md#CreditNotesCreate) | **Post** /credit_notes | Create a credit note
[**CreditNotesGet**](CreditNotesAPI.md#CreditNotesGet) | **Get** /credit_notes/{id} | Get credit note by ID
[**CreditNotesList**](CreditNotesAPI.md#CreditNotesList) | **Get** /credit_notes | List credit notes
[**CreditNotesListLines**](CreditNotesAPI.md#CreditNotesListLines) | **Get** /credit_notes/{id}/lines | List line items on a credit note
[**CreditNotesPreview**](CreditNotesAPI.md#CreditNotesPreview) | **Post** /credit_notes/preview | Preview a credit note without creating it
[**CreditNotesVoid**](CreditNotesAPI.md#CreditNotesVoid) | **Post** /credit_notes/{id}/void | Void a credit note



## CreditNotesCreate

> CreditNoteCreateResponse CreditNotesCreate(ctx).CreateCreditNoteInput(createCreditNoteInput).Execute()

Create a credit note



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
	createCreditNoteInput := *openapiclient.NewCreateCreditNoteInput("660e8400-e29b-41d4-a716-446655440000", int32(5000)) // CreateCreditNoteInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesCreate(context.Background()).CreateCreditNoteInput(createCreditNoteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesCreate`: CreditNoteCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreditNoteInput** | [**CreateCreditNoteInput**](CreateCreditNoteInput.md) |  | 

### Return type

[**CreditNoteCreateResponse**](CreditNoteCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNotesGet

> CreditNoteRetrieveResponse CreditNotesGet(ctx, id).Execute()

Get credit note by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit note UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesGet`: CreditNoteRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNoteRetrieveResponse**](CreditNoteRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNotesList

> CreditNoteListResponse CreditNotesList(ctx).Status(status).InvoiceId(invoiceId).CustomerId(customerId).CreditMethod(creditMethod).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).Offset(offset).Execute()

List credit notes



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
	status := "status_example" // string | Filter by credit note status (optional)
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by invoice ID (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer ID (optional)
	creditMethod := "creditMethod_example" // string | Filter by credit method (optional)
	createdAfter := "2024-01-01T00:00:00Z" // string | ISO 8601 date - only credit notes created after this date (optional)
	createdBefore := "2024-12-31T23:59:59Z" // string | ISO 8601 date - only credit notes created before this date (optional)
	limit := float32(20) // float32 | Number of results to return (1-100) (optional) (default to 20)
	offset := float32(0) // float32 | Number of results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesList(context.Background()).Status(status).InvoiceId(invoiceId).CustomerId(customerId).CreditMethod(creditMethod).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesList`: CreditNoteListResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by credit note status | 
 **invoiceId** | **string** | Filter by invoice ID | 
 **customerId** | **string** | Filter by customer ID | 
 **creditMethod** | **string** | Filter by credit method | 
 **createdAfter** | **string** | ISO 8601 date - only credit notes created after this date | 
 **createdBefore** | **string** | ISO 8601 date - only credit notes created before this date | 
 **limit** | **float32** | Number of results to return (1-100) | [default to 20]
 **offset** | **float32** | Number of results to skip | [default to 0]

### Return type

[**CreditNoteListResponse**](CreditNoteListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNotesListLines

> CreditNoteLineList CreditNotesListLines(ctx, id).Limit(limit).Execute()

List line items on a credit note



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit note UUID
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesListLines(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesListLines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesListLines`: CreditNoteLineList
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesListLines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesListLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]

### Return type

[**CreditNoteLineList**](CreditNoteLineList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNotesPreview

> CreditNotePreviewResponse CreditNotesPreview(ctx).PreviewCreditNoteInput(previewCreditNoteInput).Execute()

Preview a credit note without creating it



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
	previewCreditNoteInput := *openapiclient.NewPreviewCreditNoteInput("660e8400-e29b-41d4-a716-446655440000", int32(5000)) // PreviewCreditNoteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesPreview(context.Background()).PreviewCreditNoteInput(previewCreditNoteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesPreview`: CreditNotePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewCreditNoteInput** | [**PreviewCreditNoteInput**](PreviewCreditNoteInput.md) |  | 

### Return type

[**CreditNotePreviewResponse**](CreditNotePreviewResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNotesVoid

> CreditNoteVoidResponse CreditNotesVoid(ctx, id).Execute()

Void a credit note



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Credit note UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditNotesVoid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditNotesVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNotesVoid`: CreditNoteVoidResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditNotesVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditNotesVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreditNoteVoidResponse**](CreditNoteVoidResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

