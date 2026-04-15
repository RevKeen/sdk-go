# \PricesAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricesArchive**](PricesAPI.md#PricesArchive) | **Delete** /prices/{id} | Archive a price
[**PricesCreate**](PricesAPI.md#PricesCreate) | **Post** /prices | Create a price
[**PricesList**](PricesAPI.md#PricesList) | **Get** /prices | List prices
[**PricesRetrieve**](PricesAPI.md#PricesRetrieve) | **Get** /prices/{id} | Retrieve a price
[**PricesUpdate**](PricesAPI.md#PricesUpdate) | **Patch** /prices/{id} | Update a price



## PricesArchive

> Price PricesArchive(ctx, id).Execute()

Archive a price



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Price ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesArchive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesArchive`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Price**](Price.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesCreate

> Price PricesCreate(ctx).CreatePriceRequest(createPriceRequest).Execute()

Create a price



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
	createPriceRequest := *openapiclient.NewCreatePriceRequest("ProductId_example") // CreatePriceRequest | Price creation details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesCreate(context.Background()).CreatePriceRequest(createPriceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesCreate`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPriceRequest** | [**CreatePriceRequest**](CreatePriceRequest.md) | Price creation details | 

### Return type

[**Price**](Price.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesList

> PriceListResponse PricesList(ctx).ProductId(productId).Type_(type_).PricingModel(pricingModel).Active(active).Currency(currency).LookupKey(lookupKey).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List prices



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
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by product ID (optional)
	type_ := "type__example" // string | Filter by price type (optional)
	pricingModel := "pricingModel_example" // string | Filter by pricing model (optional)
	active := "active_example" // string | Filter by active status (optional)
	currency := "currency_example" // string | Filter by currency (optional)
	lookupKey := "lookupKey_example" // string | Filter by lookup key (optional)
	limit := int32(56) // int32 | Number of results (1-100) (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination (ID to start after) (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor for pagination (ID to end before) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesList(context.Background()).ProductId(productId).Type_(type_).PricingModel(pricingModel).Active(active).Currency(currency).LookupKey(lookupKey).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesList`: PriceListResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | Filter by product ID | 
 **type_** | **string** | Filter by price type | 
 **pricingModel** | **string** | Filter by pricing model | 
 **active** | **string** | Filter by active status | 
 **currency** | **string** | Filter by currency | 
 **lookupKey** | **string** | Filter by lookup key | 
 **limit** | **int32** | Number of results (1-100) | [default to 20]
 **startingAfter** | **string** | Cursor for pagination (ID to start after) | 
 **endingBefore** | **string** | Cursor for pagination (ID to end before) | 

### Return type

[**PriceListResponse**](PriceListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesRetrieve

> Price PricesRetrieve(ctx, id).Execute()

Retrieve a price



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Price ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesRetrieve`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Price**](Price.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUpdate

> Price PricesUpdate(ctx, id).UpdatePriceRequest(updatePriceRequest).Execute()

Update a price



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Price ID
	updatePriceRequest := *openapiclient.NewUpdatePriceRequest() // UpdatePriceRequest | Price update details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesUpdate(context.Background(), id).UpdatePriceRequest(updatePriceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUpdate`: Price
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePriceRequest** | [**UpdatePriceRequest**](UpdatePriceRequest.md) | Price update details | 

### Return type

[**Price**](Price.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

