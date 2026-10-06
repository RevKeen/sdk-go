# \ProductsAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCollectionsCreate**](ProductsAPI.md#ProductCollectionsCreate) | **Post** /product-collections | Create a product collection
[**ProductCollectionsDelete**](ProductsAPI.md#ProductCollectionsDelete) | **Delete** /product-collections/{id} | Delete a product collection
[**ProductCollectionsGet**](ProductsAPI.md#ProductCollectionsGet) | **Get** /product-collections/{id} | Get a product collection
[**ProductCollectionsList**](ProductsAPI.md#ProductCollectionsList) | **Get** /product-collections | List product collections
[**ProductCollectionsUpdate**](ProductsAPI.md#ProductCollectionsUpdate) | **Patch** /product-collections/{id} | Update a product collection
[**ProductsCreate**](ProductsAPI.md#ProductsCreate) | **Post** /products | Create product
[**ProductsGet**](ProductsAPI.md#ProductsGet) | **Get** /products/{id} | Get product by ID
[**ProductsList**](ProductsAPI.md#ProductsList) | **Get** /products | List products
[**ProductsUpdate**](ProductsAPI.md#ProductsUpdate) | **Patch** /products/{id} | Update product



## ProductCollectionsCreate

> ProductCollectionCreateResponse ProductCollectionsCreate(ctx).ProductCollectionCreateRequest(productCollectionCreateRequest).Execute()

Create a product collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	productCollectionCreateRequest := *openapiclient.NewProductCollectionCreateRequest("Name_example", "Slug_example") // ProductCollectionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCollectionsCreate(context.Background()).ProductCollectionCreateRequest(productCollectionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCollectionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCollectionsCreate`: ProductCollectionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCollectionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCollectionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productCollectionCreateRequest** | [**ProductCollectionCreateRequest**](ProductCollectionCreateRequest.md) |  |

### Return type

[**ProductCollectionCreateResponse**](ProductCollectionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCollectionsDelete

> ProductCollectionDeletedResponse ProductCollectionsDelete(ctx, id).Execute()

Delete a product collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCollectionsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCollectionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCollectionsDelete`: ProductCollectionDeletedResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCollectionsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiProductCollectionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductCollectionDeletedResponse**](ProductCollectionDeletedResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCollectionsGet

> ProductCollectionResponse ProductCollectionsGet(ctx, id).Execute()

Get a product collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "id_example" // string | Collection id or URL handle.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCollectionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCollectionsGet`: ProductCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCollectionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Collection id or URL handle. |

### Other Parameters

Other parameters are passed through a pointer to a apiProductCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductCollectionResponse**](ProductCollectionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCollectionsList

> ProductCollectionListResponse ProductCollectionsList(ctx).Execute()

List product collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCollectionsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCollectionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCollectionsList`: ProductCollectionListResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCollectionsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductCollectionsListRequest struct via the builder pattern


### Return type

[**ProductCollectionListResponse**](ProductCollectionListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductCollectionsUpdate

> ProductCollectionUpdateResponse ProductCollectionsUpdate(ctx, id).ProductCollectionUpdateRequest(productCollectionUpdateRequest).Execute()

Update a product collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	productCollectionUpdateRequest := *openapiclient.NewProductCollectionUpdateRequest() // ProductCollectionUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ProductCollectionsUpdate(context.Background(), id).ProductCollectionUpdateRequest(productCollectionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ProductCollectionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCollectionsUpdate`: ProductCollectionUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ProductCollectionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiProductCollectionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productCollectionUpdateRequest** | [**ProductCollectionUpdateRequest**](ProductCollectionUpdateRequest.md) |  |

### Return type

[**ProductCollectionUpdateResponse**](ProductCollectionUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	createProductRequest := *openapiclient.NewCreateProductRequest("ProductId_example", "Name_example", "Kind_example", "PricingModel_example", int64(123)) // CreateProductRequest | Product creation details

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
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

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
	openapiclient "github.com/revkeen/sdk-go"
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
	openapiclient "github.com/revkeen/sdk-go"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
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
