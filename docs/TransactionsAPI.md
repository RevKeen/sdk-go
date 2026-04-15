# \TransactionsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsList**](TransactionsAPI.md#TransactionsList) | **Get** /transactions | List transactions
[**TransactionsRetrieve**](TransactionsAPI.md#TransactionsRetrieve) | **Get** /transactions/{id} | Get transaction by ID



## TransactionsList

> TransactionListResponse TransactionsList(ctx).Type_(type_).Status(status).CustomerId(customerId).InvoiceId(invoiceId).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List transactions



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
	type_ := "type__example" // string | Filter by transaction type (optional)
	status := "status_example" // string | Filter by status (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer ID (optional)
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by invoice ID (optional)
	createdGte := float32(8.14) // float32 | Filter by created_at >= (Unix timestamp) (optional)
	createdLte := float32(8.14) // float32 | Filter by created_at <= (Unix timestamp) (optional)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination - return results after this ID (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination - return results before this ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.TransactionsList(context.Background()).Type_(type_).Status(status).CustomerId(customerId).InvoiceId(invoiceId).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsList`: TransactionListResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Filter by transaction type | 
 **status** | **string** | Filter by status | 
 **customerId** | **string** | Filter by customer ID | 
 **invoiceId** | **string** | Filter by invoice ID | 
 **createdGte** | **float32** | Filter by created_at &gt;&#x3D; (Unix timestamp) | 
 **createdLte** | **float32** | Filter by created_at &lt;&#x3D; (Unix timestamp) | 
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **startingAfter** | **string** | Cursor for pagination - return results after this ID | 
 **endingBefore** | **string** | Cursor for pagination - return results before this ID | 

### Return type

[**TransactionListResponse**](TransactionListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsRetrieve

> TransactionRetrieveResponse TransactionsRetrieve(ctx, id).Execute()

Get transaction by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transaction UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.TransactionsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsRetrieve`: TransactionRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionRetrieveResponse**](TransactionRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

