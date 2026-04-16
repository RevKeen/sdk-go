# \ProductsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsCreate**](ProductsAPI.md#ProductsCreate) | **Post** /products | Create product
[**ProductsExternalBatch**](ProductsAPI.md#ProductsExternalBatch) | **Put** /products/external/batch | Batch upsert products by external ID
[**ProductsExternalUpsert**](ProductsAPI.md#ProductsExternalUpsert) | **Put** /products/external/{source}/{externalId} | Upsert product by external ID
[**ProductsList**](ProductsAPI.md#ProductsList) | **Get** /products | List products
[**ProductsRetrieve**](ProductsAPI.md#ProductsRetrieve) | **Get** /products/{id} | Get product by ID
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


## ProductsExternalBatch

> ProductsExternalBatch200Response ProductsExternalBatch(ctx).ProductsExternalBatchRequest(productsExternalBatchRequest).Execute()

Batch upsert products by external ID



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
	productsExternalBatchRequest := *openapiclient.NewProductsExternalBatchRequest("practicehub", []openapiclient.ProductsExternalBatchRequestProductsInner{*openapiclient.NewProductsExternalBatchRequestProductsInner("prod_12345", "Monthly Membership", int32(9900))}) // ProductsExternalBatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsExternalBatch(context.Background()).ProductsExternalBatchRequest(productsExternalBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsExternalBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsExternalBatch`: ProductsExternalBatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsExternalBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsExternalBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productsExternalBatchRequest** | [**ProductsExternalBatchRequest**](ProductsExternalBatchRequest.md) |  | 

### Return type

[**ProductsExternalBatch200Response**](ProductsExternalBatch200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsExternalUpsert

> ProductsExternalUpsert200Response ProductsExternalUpsert(ctx, source, externalId).ProductsExternalUpsertRequest(productsExternalUpsertRequest).Execute()

Upsert product by external ID



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
	source := "practicehub" // string | Integration source (e.g., practicehub, wodify)
	externalId := "prod_12345" // string | External system's product ID
	productsExternalUpsertRequest := *openapiclient.NewProductsExternalUpsertRequest("Monthly Membership", int32(9900)) // ProductsExternalUpsertRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductsExternalUpsert(context.Background(), source, externalId).ProductsExternalUpsertRequest(productsExternalUpsertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsExternalUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsExternalUpsert`: ProductsExternalUpsert200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsExternalUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | Integration source (e.g., practicehub, wodify) | 
**externalId** | **string** | External system&#39;s product ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsExternalUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productsExternalUpsertRequest** | [**ProductsExternalUpsertRequest**](ProductsExternalUpsertRequest.md) |  | 

### Return type

[**ProductsExternalUpsert200Response**](ProductsExternalUpsert200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
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


## ProductsRetrieve

> ProductRetrieveResponse ProductsRetrieve(ctx, id).Execute()

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
	resp, r, err := apiClient.ProductsAPI.ProductsRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductsRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductsRetrieve`: ProductRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsRetrieveRequest struct via the builder pattern


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

