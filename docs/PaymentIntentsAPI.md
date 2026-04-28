# \PaymentIntentsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentIntentsCancel**](PaymentIntentsAPI.md#PaymentIntentsCancel) | **Post** /payment-intents/{id}/cancel | Cancel a payment intent
[**PaymentIntentsCapture**](PaymentIntentsAPI.md#PaymentIntentsCapture) | **Post** /payment-intents/{id}/capture | Capture a payment intent
[**PaymentIntentsConfirm**](PaymentIntentsAPI.md#PaymentIntentsConfirm) | **Post** /payment-intents/{id}/confirm | Confirm a payment intent
[**PaymentIntentsCreate**](PaymentIntentsAPI.md#PaymentIntentsCreate) | **Post** /payment-intents | Create a payment intent
[**PaymentIntentsGet**](PaymentIntentsAPI.md#PaymentIntentsGet) | **Get** /payment-intents/{id} | Retrieve a payment intent
[**PaymentIntentsList**](PaymentIntentsAPI.md#PaymentIntentsList) | **Get** /payment-intents | List payment intents



## PaymentIntentsCancel

> PaymentIntent PaymentIntentsCancel(ctx, id).CancelPaymentIntentRequest(cancelPaymentIntentRequest).Execute()

Cancel a payment intent



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
	id := "id_example" // string | Payment intent ID (pi_xxx)
	cancelPaymentIntentRequest := *openapiclient.NewCancelPaymentIntentRequest() // CancelPaymentIntentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsCancel(context.Background(), id).CancelPaymentIntentRequest(cancelPaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsCancel`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment intent ID (pi_xxx) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelPaymentIntentRequest** | [**CancelPaymentIntentRequest**](CancelPaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentIntentsCapture

> PaymentIntent PaymentIntentsCapture(ctx, id).CapturePaymentIntentRequest(capturePaymentIntentRequest).Execute()

Capture a payment intent



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
	id := "id_example" // string | Payment intent ID (pi_xxx)
	capturePaymentIntentRequest := *openapiclient.NewCapturePaymentIntentRequest() // CapturePaymentIntentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsCapture(context.Background(), id).CapturePaymentIntentRequest(capturePaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsCapture`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment intent ID (pi_xxx) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **capturePaymentIntentRequest** | [**CapturePaymentIntentRequest**](CapturePaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentIntentsConfirm

> PaymentIntent PaymentIntentsConfirm(ctx, id).ConfirmPaymentIntentRequest(confirmPaymentIntentRequest).Execute()

Confirm a payment intent



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
	id := "id_example" // string | Payment intent ID (pi_xxx)
	confirmPaymentIntentRequest := *openapiclient.NewConfirmPaymentIntentRequest() // ConfirmPaymentIntentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsConfirm(context.Background(), id).ConfirmPaymentIntentRequest(confirmPaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsConfirm`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment intent ID (pi_xxx) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmPaymentIntentRequest** | [**ConfirmPaymentIntentRequest**](ConfirmPaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentIntentsCreate

> PaymentIntent PaymentIntentsCreate(ctx).CreatePaymentIntentRequest(createPaymentIntentRequest).Execute()

Create a payment intent



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
	createPaymentIntentRequest := *openapiclient.NewCreatePaymentIntentRequest(int32(5000), "Customer_example") // CreatePaymentIntentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsCreate(context.Background()).CreatePaymentIntentRequest(createPaymentIntentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsCreate`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentIntentRequest** | [**CreatePaymentIntentRequest**](CreatePaymentIntentRequest.md) |  | 

### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentIntentsGet

> PaymentIntent PaymentIntentsGet(ctx, id).Execute()

Retrieve a payment intent



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
	id := "id_example" // string | Payment intent ID (pi_xxx)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsGet`: PaymentIntent
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment intent ID (pi_xxx) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentIntent**](PaymentIntent.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentIntentsList

> PaymentIntentListResponse PaymentIntentsList(ctx).Customer(customer).Status(status).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List payment intents



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
	customer := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer ID (optional)
	status := "status_example" // string | Filter by status (optional)
	createdGte := float32(8.14) // float32 | Filter by created_at >= (Unix timestamp) (optional)
	createdLte := float32(8.14) // float32 | Filter by created_at <= (Unix timestamp) (optional)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	startingAfter := "startingAfter_example" // string | Cursor for pagination - return results after this ID (pi_xxx) (optional)
	endingBefore := "endingBefore_example" // string | Cursor for pagination - return results before this ID (pi_xxx) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentIntentsAPI.PaymentIntentsList(context.Background()).Customer(customer).Status(status).CreatedGte(createdGte).CreatedLte(createdLte).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentIntentsAPI.PaymentIntentsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentIntentsList`: PaymentIntentListResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentIntentsAPI.PaymentIntentsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentIntentsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | **string** | Filter by customer ID | 
 **status** | **string** | Filter by status | 
 **createdGte** | **float32** | Filter by created_at &gt;&#x3D; (Unix timestamp) | 
 **createdLte** | **float32** | Filter by created_at &lt;&#x3D; (Unix timestamp) | 
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **startingAfter** | **string** | Cursor for pagination - return results after this ID (pi_xxx) | 
 **endingBefore** | **string** | Cursor for pagination - return results before this ID (pi_xxx) | 

### Return type

[**PaymentIntentListResponse**](PaymentIntentListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

