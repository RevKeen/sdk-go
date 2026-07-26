# \StorefrontAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorefrontOriginsCreate**](StorefrontAPI.md#StorefrontOriginsCreate) | **Post** /storefront/origins | Register a storefront origin
[**StorefrontOriginsDelete**](StorefrontAPI.md#StorefrontOriginsDelete) | **Delete** /storefront/origins/{originId} | Remove a storefront origin
[**StorefrontOriginsList**](StorefrontAPI.md#StorefrontOriginsList) | **Get** /storefront/origins | List storefront origins
[**StorefrontProductsGet**](StorefrontAPI.md#StorefrontProductsGet) | **Get** /storefront/products/{productId} | Get a storefront product
[**StorefrontProductsList**](StorefrontAPI.md#StorefrontProductsList) | **Get** /storefront/products | List storefront products
[**StorefrontStatusGet**](StorefrontAPI.md#StorefrontStatusGet) | **Get** /storefront/status | Get storefront integration status



## StorefrontOriginsCreate

> StorefrontOriginCreateResponse StorefrontOriginsCreate(ctx).StorefrontOriginCreateRequest(storefrontOriginCreateRequest).Execute()

Register a storefront origin



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
	storefrontOriginCreateRequest := *openapiclient.NewStorefrontOriginCreateRequest("https://shop.example.com") // StorefrontOriginCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorefrontAPI.StorefrontOriginsCreate(context.Background()).StorefrontOriginCreateRequest(storefrontOriginCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontOriginsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontOriginsCreate`: StorefrontOriginCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontOriginsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontOriginsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storefrontOriginCreateRequest** | [**StorefrontOriginCreateRequest**](StorefrontOriginCreateRequest.md) |  | 

### Return type

[**StorefrontOriginCreateResponse**](StorefrontOriginCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorefrontOriginsDelete

> StorefrontOriginDeleteResponse StorefrontOriginsDelete(ctx, originId).Execute()

Remove a storefront origin



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
	originId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Storefront origin id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorefrontAPI.StorefrontOriginsDelete(context.Background(), originId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontOriginsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontOriginsDelete`: StorefrontOriginDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontOriginsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**originId** | **string** | Storefront origin id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontOriginsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorefrontOriginDeleteResponse**](StorefrontOriginDeleteResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorefrontOriginsList

> StorefrontOriginListResponse StorefrontOriginsList(ctx).Execute()

List storefront origins



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
	resp, r, err := apiClient.StorefrontAPI.StorefrontOriginsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontOriginsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontOriginsList`: StorefrontOriginListResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontOriginsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontOriginsListRequest struct via the builder pattern


### Return type

[**StorefrontOriginListResponse**](StorefrontOriginListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorefrontProductsGet

> StorefrontProductResponse StorefrontProductsGet(ctx, productId).Execute()

Get a storefront product



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
	productId := "productId_example" // string | Product UUID, merchant product reference, or slug.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorefrontAPI.StorefrontProductsGet(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontProductsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontProductsGet`: StorefrontProductResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontProductsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product UUID, merchant product reference, or slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorefrontProductResponse**](StorefrontProductResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorefrontProductsList

> StorefrontProductListResponse StorefrontProductsList(ctx).Limit(limit).Execute()

List storefront products



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
	limit := int32(56) // int32 | Maximum products to return (default 50, max 100). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorefrontAPI.StorefrontProductsList(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontProductsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontProductsList`: StorefrontProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontProductsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontProductsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum products to return (default 50, max 100). | 

### Return type

[**StorefrontProductListResponse**](StorefrontProductListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorefrontStatusGet

> StorefrontStatusResponse StorefrontStatusGet(ctx).Execute()

Get storefront integration status



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
	resp, r, err := apiClient.StorefrontAPI.StorefrontStatusGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontAPI.StorefrontStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorefrontStatusGet`: StorefrontStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `StorefrontAPI.StorefrontStatusGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStorefrontStatusGetRequest struct via the builder pattern


### Return type

[**StorefrontStatusResponse**](StorefrontStatusResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

