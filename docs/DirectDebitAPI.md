# \DirectDebitAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DdMandateRequestsGet**](DirectDebitAPI.md#DdMandateRequestsGet) | **Get** /dd/mandate-requests/{token} | Look up a pending signed mandate request
[**DdPreview**](DirectDebitAPI.md#DdPreview) | **Post** /dd/preview | Preview the Direct Debit collection date chain
[**DdValidate**](DirectDebitAPI.md#DdValidate) | **Post** /dd/validate | Validate a UK bank account for Direct Debit



## DdMandateRequestsGet

> MandateRequestLookupResponse DdMandateRequestsGet(ctx, token).Execute()

Look up a pending signed mandate request



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
	token := "token_example" // string | Signed mandate-request token from the emailed link

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectDebitAPI.DdMandateRequestsGet(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectDebitAPI.DdMandateRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdMandateRequestsGet`: MandateRequestLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectDebitAPI.DdMandateRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Signed mandate-request token from the emailed link | 

### Other Parameters

Other parameters are passed through a pointer to a apiDdMandateRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MandateRequestLookupResponse**](MandateRequestLookupResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdPreview

> PreviewCollectionResponse DdPreview(ctx).PreviewCollectionRequest(previewCollectionRequest).Execute()

Preview the Direct Debit collection date chain



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
	previewCollectionRequest := *openapiclient.NewPreviewCollectionRequest() // PreviewCollectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectDebitAPI.DdPreview(context.Background()).PreviewCollectionRequest(previewCollectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectDebitAPI.DdPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdPreview`: PreviewCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectDebitAPI.DdPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDdPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewCollectionRequest** | [**PreviewCollectionRequest**](PreviewCollectionRequest.md) |  | 

### Return type

[**PreviewCollectionResponse**](PreviewCollectionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdValidate

> ValidateBankAccountResponse DdValidate(ctx).ValidateBankAccountRequest(validateBankAccountRequest).Execute()

Validate a UK bank account for Direct Debit



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
	validateBankAccountRequest := *openapiclient.NewValidateBankAccountRequest("20-00-00", "55779911") // ValidateBankAccountRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DirectDebitAPI.DdValidate(context.Background()).ValidateBankAccountRequest(validateBankAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DirectDebitAPI.DdValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdValidate`: ValidateBankAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `DirectDebitAPI.DdValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDdValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateBankAccountRequest** | [**ValidateBankAccountRequest**](ValidateBankAccountRequest.md) |  | 

### Return type

[**ValidateBankAccountResponse**](ValidateBankAccountResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

