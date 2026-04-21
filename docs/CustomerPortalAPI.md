# \CustomerPortalAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerPortalCustomerGet**](CustomerPortalAPI.md#CustomerPortalCustomerGet) | **Get** /customer-portal/customer | Retrieve the authenticated customer
[**CustomerPortalInvoicesGet**](CustomerPortalAPI.md#CustomerPortalInvoicesGet) | **Get** /customer-portal/invoices/{id} | Retrieve an invoice
[**CustomerPortalInvoicesList**](CustomerPortalAPI.md#CustomerPortalInvoicesList) | **Get** /customer-portal/invoices | List the authenticated customer&#39;s invoices
[**CustomerPortalSessionsCreate**](CustomerPortalAPI.md#CustomerPortalSessionsCreate) | **Post** /customer-portal/sessions | Create a customer-portal session
[**CustomerPortalSubscriptionsCancel**](CustomerPortalAPI.md#CustomerPortalSubscriptionsCancel) | **Post** /customer-portal/subscriptions/{id}/cancel | Cancel a subscription
[**CustomerPortalSubscriptionsGet**](CustomerPortalAPI.md#CustomerPortalSubscriptionsGet) | **Get** /customer-portal/subscriptions/{id} | Retrieve a subscription
[**CustomerPortalSubscriptionsList**](CustomerPortalAPI.md#CustomerPortalSubscriptionsList) | **Get** /customer-portal/subscriptions | List the authenticated customer&#39;s subscriptions



## CustomerPortalCustomerGet

> PortalCustomerResponse CustomerPortalCustomerGet(ctx).Execute()

Retrieve the authenticated customer



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
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalCustomerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalCustomerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalCustomerGet`: PortalCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalCustomerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalCustomerGetRequest struct via the builder pattern


### Return type

[**PortalCustomerResponse**](PortalCustomerResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalInvoicesGet

> PortalInvoiceResponse CustomerPortalInvoicesGet(ctx, id).Execute()

Retrieve an invoice



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
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalInvoicesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalInvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalInvoicesGet`: PortalInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalInvoicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalInvoiceResponse**](PortalInvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalInvoicesList

> PortalInvoiceList CustomerPortalInvoicesList(ctx).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the authenticated customer's invoices



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
	limit := int32(56) // int32 |  (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalInvoicesList(context.Background()).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalInvoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalInvoicesList`: PortalInvoiceList
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalInvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **startingAfter** | **string** |  | 
 **endingBefore** | **string** |  | 

### Return type

[**PortalInvoiceList**](PortalInvoiceList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSessionsCreate

> CustomerPortalSessionCreateResponse CustomerPortalSessionsCreate(ctx).CreateCustomerPortalSessionRequest(createCustomerPortalSessionRequest).Execute()

Create a customer-portal session



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
	createCustomerPortalSessionRequest := *openapiclient.NewCreateCustomerPortalSessionRequest("cus_a1b2c3d4e5f6") // CreateCustomerPortalSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSessionsCreate(context.Background()).CreateCustomerPortalSessionRequest(createCustomerPortalSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSessionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSessionsCreate`: CustomerPortalSessionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSessionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSessionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerPortalSessionRequest** | [**CreateCustomerPortalSessionRequest**](CreateCustomerPortalSessionRequest.md) |  | 

### Return type

[**CustomerPortalSessionCreateResponse**](CustomerPortalSessionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsCancel

> PortalSubscriptionCancelResponse CustomerPortalSubscriptionsCancel(ctx, id).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()

Cancel a subscription



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
	cancelSubscriptionRequest := *openapiclient.NewCancelSubscriptionRequest() // CancelSubscriptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsCancel(context.Background(), id).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsCancel`: PortalSubscriptionCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelSubscriptionRequest** | [**CancelSubscriptionRequest**](CancelSubscriptionRequest.md) |  | 

### Return type

[**PortalSubscriptionCancelResponse**](PortalSubscriptionCancelResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsGet

> PortalSubscriptionResponse CustomerPortalSubscriptionsGet(ctx, id).Execute()

Retrieve a subscription



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
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsGet`: PortalSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalSubscriptionResponse**](PortalSubscriptionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsList

> PortalSubscriptionList CustomerPortalSubscriptionsList(ctx).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the authenticated customer's subscriptions



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
	limit := int32(56) // int32 |  (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsList(context.Background()).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsList`: PortalSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **startingAfter** | **string** |  | 
 **endingBefore** | **string** |  | 

### Return type

[**PortalSubscriptionList**](PortalSubscriptionList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

