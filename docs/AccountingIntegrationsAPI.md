# \AccountingIntegrationsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountingInvoicePaymentRequestsCreate**](AccountingIntegrationsAPI.md#AccountingInvoicePaymentRequestsCreate) | **Post** /integrations/accounting/invoice-payment-requests | Create an accounting-led invoice payment request
[**AccountingInvoicePaymentRequestsGet**](AccountingIntegrationsAPI.md#AccountingInvoicePaymentRequestsGet) | **Get** /integrations/accounting/invoice-payment-requests/{id} | Get an accounting-led invoice payment request



## AccountingInvoicePaymentRequestsCreate

> AccountingInvoicePaymentRequestsCreate200Response AccountingInvoicePaymentRequestsCreate(ctx).CreateAccountingInvoicePaymentRequestInput(createAccountingInvoicePaymentRequestInput).Execute()

Create an accounting-led invoice payment request



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
	createAccountingInvoicePaymentRequestInput := *openapiclient.NewCreateAccountingInvoicePaymentRequestInput("Provider_example", "ConnectionId_example", "ProviderAccountId_example", *openapiclient.NewCreateAccountingInvoicePaymentRequestInputExternalInvoice("Id_example"), int32(123), "Currency_example") // CreateAccountingInvoicePaymentRequestInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsCreate(context.Background()).CreateAccountingInvoicePaymentRequestInput(createAccountingInvoicePaymentRequestInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountingInvoicePaymentRequestsCreate`: AccountingInvoicePaymentRequestsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountingInvoicePaymentRequestsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountingInvoicePaymentRequestInput** | [**CreateAccountingInvoicePaymentRequestInput**](CreateAccountingInvoicePaymentRequestInput.md) |  | 

### Return type

[**AccountingInvoicePaymentRequestsCreate200Response**](AccountingInvoicePaymentRequestsCreate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountingInvoicePaymentRequestsGet

> AccountingInvoicePaymentRequestsCreate200Response AccountingInvoicePaymentRequestsGet(ctx, id).Execute()

Get an accounting-led invoice payment request



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
	resp, r, err := apiClient.AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountingInvoicePaymentRequestsGet`: AccountingInvoicePaymentRequestsCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountingIntegrationsAPI.AccountingInvoicePaymentRequestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountingInvoicePaymentRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountingInvoicePaymentRequestsCreate200Response**](AccountingInvoicePaymentRequestsCreate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

