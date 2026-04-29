# \CustomersAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersCreate**](CustomersAPI.md#CustomersCreate) | **Post** /customers | Create a new customer
[**CustomersGet**](CustomersAPI.md#CustomersGet) | **Get** /customers/{id} | Get customer by ID
[**CustomersList**](CustomersAPI.md#CustomersList) | **Get** /customers | List customers
[**CustomersPaymentMethodsList**](CustomersAPI.md#CustomersPaymentMethodsList) | **Get** /customers/{id}/payment-methods | Get customer payment methods
[**CustomersPaymentRailsGet**](CustomersAPI.md#CustomersPaymentRailsGet) | **Get** /customers/{id}/payment-rails | Get customer payment rail availability
[**CustomersPreferredRailsGet**](CustomersAPI.md#CustomersPreferredRailsGet) | **Get** /customers/{id}/preferred-rails | Get preferred payment rails
[**CustomersUpdate**](CustomersAPI.md#CustomersUpdate) | **Patch** /customers/{id} | Update customer details



## CustomersCreate

> CustomerCreateResponse CustomersCreate(ctx).CustomersCreateRequest(customersCreateRequest).Execute()

Create a new customer



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
	customersCreateRequest := *openapiclient.NewCustomersCreateRequest("Email_example") // CustomersCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersCreate(context.Background()).CustomersCreateRequest(customersCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersCreate`: CustomerCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customersCreateRequest** | [**CustomersCreateRequest**](CustomersCreateRequest.md) |  | 

### Return type

[**CustomerCreateResponse**](CustomerCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersGet

> CustomerRetrieveResponse CustomersGet(ctx, id).Execute()

Get customer by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersGet`: CustomerRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerRetrieveResponse**](CustomerRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersList

> CustomerListResponse CustomersList(ctx).Limit(limit).Offset(offset).Search(search).Execute()

List customers



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
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	offset := int32(56) // int32 | Number of results to skip (optional) (default to 0)
	search := "search_example" // string | Search term to filter customers (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersList(context.Background()).Limit(limit).Offset(offset).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersList`: CustomerListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **offset** | **int32** | Number of results to skip | [default to 0]
 **search** | **string** | Search term to filter customers | 

### Return type

[**CustomerListResponse**](CustomerListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersPaymentMethodsList

> CustomerPaymentMethodsListResponse CustomersPaymentMethodsList(ctx, id).Execute()

Get customer payment methods



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersPaymentMethodsList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersPaymentMethodsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersPaymentMethodsList`: CustomerPaymentMethodsListResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersPaymentMethodsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersPaymentMethodsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerPaymentMethodsListResponse**](CustomerPaymentMethodsListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersPaymentRailsGet

> CustomersPaymentRailsGet200Response CustomersPaymentRailsGet(ctx, id).InvoiceId(invoiceId).Execute()

Get customer payment rail availability



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional invoice context used to apply invoice-level rail restrictions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersPaymentRailsGet(context.Background(), id).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersPaymentRailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersPaymentRailsGet`: CustomersPaymentRailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersPaymentRailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersPaymentRailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceId** | **string** | Optional invoice context used to apply invoice-level rail restrictions. | 

### Return type

[**CustomersPaymentRailsGet200Response**](CustomersPaymentRailsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersPreferredRailsGet

> CustomersPreferredRailsGet200Response CustomersPreferredRailsGet(ctx, id).InvoiceId(invoiceId).Amount(amount).AmountMinor(amountMinor).Execute()

Get preferred payment rails



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	invoiceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional invoice context used to apply invoice-level rail restrictions and amount. (optional)
	amount := int32(56) // int32 | Optional amount in minor units. Use amount_minor for the explicit alias. (optional)
	amountMinor := int32(56) // int32 | Optional amount in minor units. Overrides invoice total for margin-aware ranking. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersPreferredRailsGet(context.Background(), id).InvoiceId(invoiceId).Amount(amount).AmountMinor(amountMinor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersPreferredRailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersPreferredRailsGet`: CustomersPreferredRailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersPreferredRailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersPreferredRailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoiceId** | **string** | Optional invoice context used to apply invoice-level rail restrictions and amount. | 
 **amount** | **int32** | Optional amount in minor units. Use amount_minor for the explicit alias. | 
 **amountMinor** | **int32** | Optional amount in minor units. Overrides invoice total for margin-aware ranking. | 

### Return type

[**CustomersPreferredRailsGet200Response**](CustomersPreferredRailsGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersUpdate

> CustomerUpdateResponse CustomersUpdate(ctx, id).CustomersUpdateRequest(customersUpdateRequest).Execute()

Update customer details



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	customersUpdateRequest := *openapiclient.NewCustomersUpdateRequest() // CustomersUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersUpdate(context.Background(), id).CustomersUpdateRequest(customersUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersUpdate`: CustomerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customersUpdateRequest** | [**CustomersUpdateRequest**](CustomersUpdateRequest.md) |  | 

### Return type

[**CustomerUpdateResponse**](CustomerUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

