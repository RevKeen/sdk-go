# \RefundsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefundsCreate**](RefundsAPI.md#RefundsCreate) | **Post** /refunds | Create a refund
[**RefundsGet**](RefundsAPI.md#RefundsGet) | **Get** /refunds/{id} | Get refund by ID
[**RefundsList**](RefundsAPI.md#RefundsList) | **Get** /refunds | List refunds



## RefundsCreate

> RefundCreateResponse RefundsCreate(ctx).CreateRefundInput(createRefundInput).Execute()

Create a refund



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
	createRefundInput := *openapiclient.NewCreateRefundInput("660e8400-e29b-41d4-a716-446655440000") // CreateRefundInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundsCreate(context.Background()).CreateRefundInput(createRefundInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundsCreate`: RefundCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRefundInput** | [**CreateRefundInput**](CreateRefundInput.md) |  | 

### Return type

[**RefundCreateResponse**](RefundCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundsGet

> RefundRetrieveResponse RefundsGet(ctx, id).Execute()

Get refund by ID



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
	id := "id_example" // string | Refund public ID (ref_xxx) or UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundsGet`: RefundRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Refund public ID (ref_xxx) or UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefundRetrieveResponse**](RefundRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundsList

> RefundListResponse RefundsList(ctx).Status(status).PaymentId(paymentId).Gateway(gateway).Reason(reason).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).Offset(offset).Execute()

List refunds



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
	status := "status_example" // string | Filter by refund status (optional)
	paymentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by original payment ID (optional)
	gateway := "gateway_example" // string | Filter by payment gateway (nmi, stripe, etc.) (optional)
	reason := "reason_example" // string | Filter by refund reason (optional)
	createdAfter := "2024-01-01T00:00:00Z" // string | ISO 8601 date - only refunds created after this date (optional)
	createdBefore := "2024-12-31T23:59:59Z" // string | ISO 8601 date - only refunds created before this date (optional)
	limit := float32(20) // float32 | Number of results to return (1-100) (optional) (default to 20)
	offset := float32(0) // float32 | Number of results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefundsAPI.RefundsList(context.Background()).Status(status).PaymentId(paymentId).Gateway(gateway).Reason(reason).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefundsAPI.RefundsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundsList`: RefundListResponse
	fmt.Fprintf(os.Stdout, "Response from `RefundsAPI.RefundsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefundsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by refund status | 
 **paymentId** | **string** | Filter by original payment ID | 
 **gateway** | **string** | Filter by payment gateway (nmi, stripe, etc.) | 
 **reason** | **string** | Filter by refund reason | 
 **createdAfter** | **string** | ISO 8601 date - only refunds created after this date | 
 **createdBefore** | **string** | ISO 8601 date - only refunds created before this date | 
 **limit** | **float32** | Number of results to return (1-100) | [default to 20]
 **offset** | **float32** | Number of results to skip | [default to 0]

### Return type

[**RefundListResponse**](RefundListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

