# \WebhookEndpointsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookEndpointsCreate**](WebhookEndpointsAPI.md#WebhookEndpointsCreate) | **Post** /webhook-endpoints | 
[**WebhookEndpointsDelete**](WebhookEndpointsAPI.md#WebhookEndpointsDelete) | **Delete** /webhook-endpoints/{id} | 
[**WebhookEndpointsGet**](WebhookEndpointsAPI.md#WebhookEndpointsGet) | **Get** /webhook-endpoints/{id} | 
[**WebhookEndpointsList**](WebhookEndpointsAPI.md#WebhookEndpointsList) | **Get** /webhook-endpoints | 
[**WebhookEndpointsRotateSecret**](WebhookEndpointsAPI.md#WebhookEndpointsRotateSecret) | **Post** /webhook-endpoints/{id}/rotate-secret | 
[**WebhookEndpointsUpdate**](WebhookEndpointsAPI.md#WebhookEndpointsUpdate) | **Patch** /webhook-endpoints/{id} | 



## WebhookEndpointsCreate

> WebhookEndpointsCreate201Response WebhookEndpointsCreate(ctx).WebhookEndpointsCreateRequest(webhookEndpointsCreateRequest).Execute()



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
	webhookEndpointsCreateRequest := *openapiclient.NewWebhookEndpointsCreateRequest("Url_example", []string{"EnabledEvents_example"}) // WebhookEndpointsCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsCreate(context.Background()).WebhookEndpointsCreateRequest(webhookEndpointsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsCreate`: WebhookEndpointsCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookEndpointsCreateRequest** | [**WebhookEndpointsCreateRequest**](WebhookEndpointsCreateRequest.md) |  | 

### Return type

[**WebhookEndpointsCreate201Response**](WebhookEndpointsCreate201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointsDelete

> WebhookEndpointsDelete200Response WebhookEndpointsDelete(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsDelete`: WebhookEndpointsDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpointsDelete200Response**](WebhookEndpointsDelete200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointsGet

> WebhookEndpointsGet200Response WebhookEndpointsGet(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsGet`: WebhookEndpointsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpointsGet200Response**](WebhookEndpointsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointsList

> WebhookEndpointsList200Response WebhookEndpointsList(ctx).Status(status).Limit(limit).Offset(offset).Execute()



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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsList(context.Background()).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsList`: WebhookEndpointsList200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**WebhookEndpointsList200Response**](WebhookEndpointsList200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointsRotateSecret

> WebhookEndpointsCreate201Response WebhookEndpointsRotateSecret(ctx, id).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsRotateSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsRotateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsRotateSecret`: WebhookEndpointsCreate201Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsRotateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsRotateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebhookEndpointsCreate201Response**](WebhookEndpointsCreate201Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookEndpointsUpdate

> WebhookEndpointsGet200Response WebhookEndpointsUpdate(ctx, id).WebhookEndpointsUpdateRequest(webhookEndpointsUpdateRequest).Execute()



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	webhookEndpointsUpdateRequest := *openapiclient.NewWebhookEndpointsUpdateRequest() // WebhookEndpointsUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookEndpointsAPI.WebhookEndpointsUpdate(context.Background(), id).WebhookEndpointsUpdateRequest(webhookEndpointsUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookEndpointsAPI.WebhookEndpointsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhookEndpointsUpdate`: WebhookEndpointsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookEndpointsAPI.WebhookEndpointsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookEndpointsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookEndpointsUpdateRequest** | [**WebhookEndpointsUpdateRequest**](WebhookEndpointsUpdateRequest.md) |  | 

### Return type

[**WebhookEndpointsGet200Response**](WebhookEndpointsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

