# \PaymentLinksAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePaymentLink**](PaymentLinksAPI.md#ActivatePaymentLink) | **Post** /payment-links/{id}/activate | Activate a payment link
[**ArchivePaymentLink**](PaymentLinksAPI.md#ArchivePaymentLink) | **Post** /payment-links/{id}/archive | Archive a payment link
[**CreatePaymentLink**](PaymentLinksAPI.md#CreatePaymentLink) | **Post** /payment-links | Create a new payment link
[**DeactivatePaymentLink**](PaymentLinksAPI.md#DeactivatePaymentLink) | **Post** /payment-links/{id}/deactivate | Deactivate a payment link
[**ExpirePaymentLink**](PaymentLinksAPI.md#ExpirePaymentLink) | **Post** /payment-links/{id}/expire | Expire a payment link
[**GetPaymentLink**](PaymentLinksAPI.md#GetPaymentLink) | **Get** /payment-links/{id} | Get payment link by ID
[**ListPaymentLinks**](PaymentLinksAPI.md#ListPaymentLinks) | **Get** /payment-links | List payment links
[**UpdatePaymentLinkStatus**](PaymentLinksAPI.md#UpdatePaymentLinkStatus) | **Patch** /payment-links/{id}/status | Update payment link status



## ActivatePaymentLink

> UpdatePaymentLinkStatus200Response ActivatePaymentLink(ctx, id).Execute()

Activate a payment link



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
	resp, r, err := apiClient.PaymentLinksAPI.ActivatePaymentLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.ActivatePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivatePaymentLink`: UpdatePaymentLinkStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.ActivatePaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdatePaymentLinkStatus200Response**](UpdatePaymentLinkStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchivePaymentLink

> UpdatePaymentLinkStatus200Response ArchivePaymentLink(ctx, id).Execute()

Archive a payment link



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
	resp, r, err := apiClient.PaymentLinksAPI.ArchivePaymentLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.ArchivePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchivePaymentLink`: UpdatePaymentLinkStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.ArchivePaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdatePaymentLinkStatus200Response**](UpdatePaymentLinkStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentLink

> ListPaymentLinks200ResponseDataInner CreatePaymentLink(ctx).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()

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
	createPaymentLinkRequest := *openapiclient.NewCreatePaymentLinkRequest() // CreatePaymentLinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.CreatePaymentLink(context.Background()).CreatePaymentLinkRequest(createPaymentLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.CreatePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentLink`: ListPaymentLinks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.CreatePaymentLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPaymentLinkRequest** | [**CreatePaymentLinkRequest**](CreatePaymentLinkRequest.md) |  | 

### Return type

[**ListPaymentLinks200ResponseDataInner**](ListPaymentLinks200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePaymentLink

> UpdatePaymentLinkStatus200Response DeactivatePaymentLink(ctx, id).Execute()

Deactivate a payment link



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
	resp, r, err := apiClient.PaymentLinksAPI.DeactivatePaymentLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.DeactivatePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivatePaymentLink`: UpdatePaymentLinkStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.DeactivatePaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdatePaymentLinkStatus200Response**](UpdatePaymentLinkStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpirePaymentLink

> ListPaymentLinks200ResponseDataInner ExpirePaymentLink(ctx, id).ExpirePaymentLinkRequest(expirePaymentLinkRequest).Execute()

Expire a payment link



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
	expirePaymentLinkRequest := *openapiclient.NewExpirePaymentLinkRequest() // ExpirePaymentLinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.ExpirePaymentLink(context.Background(), id).ExpirePaymentLinkRequest(expirePaymentLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.ExpirePaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpirePaymentLink`: ListPaymentLinks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.ExpirePaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpirePaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expirePaymentLinkRequest** | [**ExpirePaymentLinkRequest**](ExpirePaymentLinkRequest.md) |  | 

### Return type

[**ListPaymentLinks200ResponseDataInner**](ListPaymentLinks200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentLink

> ListPaymentLinks200ResponseDataInner GetPaymentLink(ctx, id).Execute()

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
	resp, r, err := apiClient.PaymentLinksAPI.GetPaymentLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.GetPaymentLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentLink`: ListPaymentLinks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.GetPaymentLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPaymentLinks200ResponseDataInner**](ListPaymentLinks200ResponseDataInner.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentLinks

> ListPaymentLinks200Response ListPaymentLinks(ctx).Limit(limit).Offset(offset).Status(status).Mode(mode).Search(search).Execute()

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
	resp, r, err := apiClient.PaymentLinksAPI.ListPaymentLinks(context.Background()).Limit(limit).Offset(offset).Status(status).Mode(mode).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.ListPaymentLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentLinks`: ListPaymentLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.ListPaymentLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** |  | [default to 0]
 **status** | **string** |  | 
 **mode** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**ListPaymentLinks200Response**](ListPaymentLinks200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePaymentLinkStatus

> UpdatePaymentLinkStatus200Response UpdatePaymentLinkStatus(ctx, id).UpdatePaymentLinkStatusRequest(updatePaymentLinkStatusRequest).Execute()

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
	updatePaymentLinkStatusRequest := *openapiclient.NewUpdatePaymentLinkStatusRequest("Status_example") // UpdatePaymentLinkStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentLinksAPI.UpdatePaymentLinkStatus(context.Background(), id).UpdatePaymentLinkStatusRequest(updatePaymentLinkStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentLinksAPI.UpdatePaymentLinkStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePaymentLinkStatus`: UpdatePaymentLinkStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentLinksAPI.UpdatePaymentLinkStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment link UUID or public_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentLinkStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePaymentLinkStatusRequest** | [**UpdatePaymentLinkStatusRequest**](UpdatePaymentLinkStatusRequest.md) |  | 

### Return type

[**UpdatePaymentLinkStatus200Response**](UpdatePaymentLinkStatus200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

