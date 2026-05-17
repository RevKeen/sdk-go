# \CheckoutSessionsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutSessionsCreate**](CheckoutSessionsAPI.md#CheckoutSessionsCreate) | **Post** /checkout-sessions | Create a checkout session
[**CheckoutSessionsExpire**](CheckoutSessionsAPI.md#CheckoutSessionsExpire) | **Post** /checkout-sessions/{id}/expire | Expire a checkout session
[**CheckoutSessionsGet**](CheckoutSessionsAPI.md#CheckoutSessionsGet) | **Get** /checkout-sessions/{id} | Retrieve a checkout session



## CheckoutSessionsCreate

> CheckoutSessionCreateResponse CheckoutSessionsCreate(ctx).CreateCheckoutSessionInput(createCheckoutSessionInput).Execute()

Create a checkout session



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
	createCheckoutSessionInput := *openapiclient.NewCreateCheckoutSessionInput() // CreateCheckoutSessionInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.CheckoutSessionsCreate(context.Background()).CreateCheckoutSessionInput(createCheckoutSessionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.CheckoutSessionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutSessionsCreate`: CheckoutSessionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.CheckoutSessionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutSessionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckoutSessionInput** | [**CreateCheckoutSessionInput**](CreateCheckoutSessionInput.md) |  | 

### Return type

[**CheckoutSessionCreateResponse**](CheckoutSessionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutSessionsExpire

> CheckoutSessionExpireResponse CheckoutSessionsExpire(ctx, id).Execute()

Expire a checkout session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.CheckoutSessionsExpire(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.CheckoutSessionsExpire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutSessionsExpire`: CheckoutSessionExpireResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.CheckoutSessionsExpire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutSessionsExpireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSessionExpireResponse**](CheckoutSessionExpireResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutSessionsGet

> CheckoutSessionRetrieveResponse CheckoutSessionsGet(ctx, id).Execute()

Retrieve a checkout session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.CheckoutSessionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.CheckoutSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutSessionsGet`: CheckoutSessionRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.CheckoutSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSessionRetrieveResponse**](CheckoutSessionRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

