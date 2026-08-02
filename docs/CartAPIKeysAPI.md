# \CartAPIKeysAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CartApiKeysEnsure**](CartAPIKeysAPI.md#CartApiKeysEnsure) | **Post** /cart-api-keys/ensure | Ensure managed Cart API keys exist
[**CartApiKeysRotate**](CartAPIKeysAPI.md#CartApiKeysRotate) | **Post** /cart-api-keys/{kind}/rotate | Rotate a managed Cart API key
[**CartApiKeysStatus**](CartAPIKeysAPI.md#CartApiKeysStatus) | **Get** /cart-api-keys/status | Get managed Cart API key status



## CartApiKeysEnsure

> CartApiKeysStatus200Response CartApiKeysEnsure(ctx).Execute()

Ensure managed Cart API keys exist



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPIKeysAPI.CartApiKeysEnsure(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPIKeysAPI.CartApiKeysEnsure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartApiKeysEnsure`: CartApiKeysStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPIKeysAPI.CartApiKeysEnsure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCartApiKeysEnsureRequest struct via the builder pattern


### Return type

[**CartApiKeysStatus200Response**](CartApiKeysStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartApiKeysRotate

> CartApiKeysRotate200Response CartApiKeysRotate(ctx, kind).CartApiKeysRotateRequest(cartApiKeysRotateRequest).Execute()

Rotate a managed Cart API key



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
	kind := "publishable" // string | 
	cartApiKeysRotateRequest := *openapiclient.NewCartApiKeysRotateRequest() // CartApiKeysRotateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPIKeysAPI.CartApiKeysRotate(context.Background(), kind).CartApiKeysRotateRequest(cartApiKeysRotateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPIKeysAPI.CartApiKeysRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartApiKeysRotate`: CartApiKeysRotate200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPIKeysAPI.CartApiKeysRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartApiKeysRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cartApiKeysRotateRequest** | [**CartApiKeysRotateRequest**](CartApiKeysRotateRequest.md) |  | 

### Return type

[**CartApiKeysRotate200Response**](CartApiKeysRotate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartApiKeysStatus

> CartApiKeysStatus200Response CartApiKeysStatus(ctx).Execute()

Get managed Cart API key status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartAPIKeysAPI.CartApiKeysStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartAPIKeysAPI.CartApiKeysStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartApiKeysStatus`: CartApiKeysStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CartAPIKeysAPI.CartApiKeysStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCartApiKeysStatusRequest struct via the builder pattern


### Return type

[**CartApiKeysStatus200Response**](CartApiKeysStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

