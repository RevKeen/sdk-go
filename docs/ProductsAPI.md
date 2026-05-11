# \ProductsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsCreate**](ProductsAPI.md#ProductsCreate) | **Post** /products | Create product
[**ProductsGet**](ProductsAPI.md#ProductsGet) | **Get** /products/{id} | Get product by ID
[**ProductsList**](ProductsAPI.md#ProductsList) | **Get** /products | List products
[**ProductsUpdate**](ProductsAPI.md#ProductsUpdate) | **Patch** /products/{id} | Update product



## ProductsCreate

> ProductCreateResponse ProductsCreate(ctx).CreateProductRequest(createProductRequest).Execute()

Create product



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
	createProductRequest := *openapiclient.NewCreateProductRequest("ProductId_example", "Name_example", "Kind_example", "PricingModel_example", int32(123)) // CreateProductRequest | Product creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsCreate(context.Background()).CreateProductRequest(createProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsCreate`: ProductCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) | Product creation details | 

### Return type

[**ProductCreateResponse**](ProductCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsGet

> ProductRetrieveResponse ProductsGet(ctx, id).Execute()

Get product by ID



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
	id := "prod_123" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsGet`: ProductRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductRetrieveResponse**](ProductRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsList

> ProductListResponse ProductsList(ctx).Page(page).Limit(limit).Search(search).Execute()

List products



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
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	search := "search_example" // string | Search term to filter products (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsList(context.Background()).Page(page).Limit(limit).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsList`: ProductListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **search** | **string** | Search term to filter products | 

### Return type

[**ProductListResponse**](ProductListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsUpdate

> ProductUpdateResponse ProductsUpdate(ctx, id).UpdateProductRequest(updateProductRequest).Execute()

Update product



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
	id := "prod_123" // string | 
	updateProductRequest := *openapiclient.NewUpdateProductRequest() // UpdateProductRequest | Product update details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsUpdate(context.Background(), id).UpdateProductRequest(updateProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsUpdate`: ProductUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProductRequest** | [**UpdateProductRequest**](UpdateProductRequest.md) | Product update details | 

### Return type

[**ProductUpdateResponse**](ProductUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

