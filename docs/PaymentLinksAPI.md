# \PaymentLinksAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentLinksActivate**](PaymentLinksAPI.md#PaymentLinksActivate) | **Post** /payment-links/{id}/activate | Activate a payment link (deprecated)
[**PaymentLinksArchive**](PaymentLinksAPI.md#PaymentLinksArchive) | **Post** /payment-links/{id}/archive | Archive a payment link (deprecated)
[**PaymentLinksCreate**](PaymentLinksAPI.md#PaymentLinksCreate) | **Post** /payment-links | Create a new payment link
[**PaymentLinksDeactivate**](PaymentLinksAPI.md#PaymentLinksDeactivate) | **Post** /payment-links/{id}/deactivate | Deactivate a payment link (deprecated)
[**PaymentLinksExpire**](PaymentLinksAPI.md#PaymentLinksExpire) | **Post** /payment-links/{id}/expire | Expire a payment link (deprecated)
[**PaymentLinksGet**](PaymentLinksAPI.md#PaymentLinksGet) | **Get** /payment-links/{id} | Get payment link by ID
[**PaymentLinksList**](PaymentLinksAPI.md#PaymentLinksList) | **Get** /payment-links | List payment links
[**PaymentLinksUpdate**](PaymentLinksAPI.md#PaymentLinksUpdate) | **Patch** /payment-links/{id}/status | Update payment link status



## PaymentLinksActivate

> PaymentLinksUpdate200Response PaymentLinksActivate(ctx, id).Execute()

Activate a payment link (deprecated)



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
	id := "id_example" // string | Payment link UUID or public_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksActivate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksActivate`: PaymentLinksUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinksUpdate200Response**](PaymentLinksUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksArchive

> PaymentLinksUpdate200Response PaymentLinksArchive(ctx, id).Execute()

Archive a payment link (deprecated)



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
	id := "id_example" // string | Payment link UUID or public_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksArchive(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksArchive`: PaymentLinksUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinksUpdate200Response**](PaymentLinksUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksCreate

> PaymentLinksList200ResponseDataInner PaymentLinksCreate(ctx).PaymentLinksCreateRequest(paymentLinksCreateRequest).Execute()

Create a new payment link



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
	paymentLinksCreateRequest := *openapiclient.NewPaymentLinksCreateRequest() // PaymentLinksCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksCreate(context.Background()).PaymentLinksCreateRequest(paymentLinksCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksCreate`: PaymentLinksList200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentLinksCreateRequest** | [**PaymentLinksCreateRequest**](PaymentLinksCreateRequest.md) |  | 

### Return type

[**PaymentLinksList200ResponseDataInner**](PaymentLinksList200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksDeactivate

> PaymentLinksUpdate200Response PaymentLinksDeactivate(ctx, id).Execute()

Deactivate a payment link (deprecated)



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
	id := "id_example" // string | Payment link UUID or public_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksDeactivate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksDeactivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksDeactivate`: PaymentLinksUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinksUpdate200Response**](PaymentLinksUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksExpire

> PaymentLinksList200ResponseDataInner PaymentLinksExpire(ctx, id).PaymentLinksExpireRequest(paymentLinksExpireRequest).Execute()

Expire a payment link (deprecated)



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
	id := "id_example" // string | Payment link UUID or public_id
	paymentLinksExpireRequest := *openapiclient.NewPaymentLinksExpireRequest() // PaymentLinksExpireRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksExpire(context.Background(), id).PaymentLinksExpireRequest(paymentLinksExpireRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksExpire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksExpire`: PaymentLinksList200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksExpire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksExpireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentLinksExpireRequest** | [**PaymentLinksExpireRequest**](PaymentLinksExpireRequest.md) |  | 

### Return type

[**PaymentLinksList200ResponseDataInner**](PaymentLinksList200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksGet

> PaymentLinksList200ResponseDataInner PaymentLinksGet(ctx, id).Execute()

Get payment link by ID



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
	id := "id_example" // string | Payment link UUID or public_id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksGet`: PaymentLinksList200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentLinksList200ResponseDataInner**](PaymentLinksList200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksList

> PaymentLinksList200Response PaymentLinksList(ctx).Limit(limit).Offset(offset).Status(status).Mode(mode).Search(search).Execute()

List payment links



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
	limit := int32(56) // int32 |  (optional) (default to 50)
	offset := int32(56) // int32 |  (optional) (default to 0)
	status := "status_example" // string |  (optional)
	mode := "mode_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksList(context.Background()).Limit(limit).Offset(offset).Status(status).Mode(mode).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksList`: PaymentLinksList200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **status** | **string** |  | 
 **mode** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**PaymentLinksList200Response**](PaymentLinksList200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentLinksUpdate

> PaymentLinksUpdate200Response PaymentLinksUpdate(ctx, id).PaymentLinksUpdateRequest(paymentLinksUpdateRequest).Execute()

Update payment link status



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
	id := "id_example" // string | Payment link UUID or public_id
	paymentLinksUpdateRequest := *openapiclient.NewPaymentLinksUpdateRequest("Status_example") // PaymentLinksUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.PaymentLinksUpdate(context.Background(), id).PaymentLinksUpdateRequest(paymentLinksUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.PaymentLinksUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentLinksUpdate`: PaymentLinksUpdate200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.PaymentLinksUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentLinksUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentLinksUpdateRequest** | [**PaymentLinksUpdateRequest**](PaymentLinksUpdateRequest.md) |  | 

### Return type

[**PaymentLinksUpdate200Response**](PaymentLinksUpdate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

