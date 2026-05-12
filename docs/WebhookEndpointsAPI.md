# \WebhookEndpointsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookEndpointsCreate**](WebhookEndpointsAPI.md#WebhookEndpointsCreate) | **Post** /webhook-endpoints | Create webhook endpoint
[**WebhookEndpointsDelete**](WebhookEndpointsAPI.md#WebhookEndpointsDelete) | **Delete** /webhook-endpoints/{id} | Delete webhook endpoint
[**WebhookEndpointsGet**](WebhookEndpointsAPI.md#WebhookEndpointsGet) | **Get** /webhook-endpoints/{id} | Get webhook endpoint
[**WebhookEndpointsList**](WebhookEndpointsAPI.md#WebhookEndpointsList) | **Get** /webhook-endpoints | List webhook endpoints
[**WebhookEndpointsRotateSecret**](WebhookEndpointsAPI.md#WebhookEndpointsRotateSecret) | **Post** /webhook-endpoints/{id}/rotate-secret | Rotate signing secret
[**WebhookEndpointsUpdate**](WebhookEndpointsAPI.md#WebhookEndpointsUpdate) | **Patch** /webhook-endpoints/{id} | Update webhook endpoint



## WebhookEndpointsCreate

> WebhookEndpointsCreate201Response WebhookEndpointsCreate(ctx).WebhookEndpointsCreateRequest(webhookEndpointsCreateRequest).Execute()

Create webhook endpoint



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

Delete webhook endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook endpoint ID

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
**id** | **string** | Webhook endpoint ID | 

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

Get webhook endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook endpoint ID

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
**id** | **string** | Webhook endpoint ID | 

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

List webhook endpoints



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
	status := "status_example" // string | Filter by endpoint status (optional)
	limit := int32(56) // int32 | Maximum number of results to return (1-100, default 20) (optional) (default to 20)
	offset := int32(56) // int32 | Number of results to skip for pagination (optional) (default to 0)

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
 **status** | **string** | Filter by endpoint status | 
 **limit** | **int32** | Maximum number of results to return (1-100, default 20) | [default to 20]
 **offset** | **int32** | Number of results to skip for pagination | [default to 0]

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

Rotate signing secret



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook endpoint ID

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
**id** | **string** | Webhook endpoint ID | 

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

Update webhook endpoint



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Webhook endpoint ID
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
**id** | **string** | Webhook endpoint ID | 

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

