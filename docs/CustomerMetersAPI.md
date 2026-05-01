# \CustomerMetersAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerMetersGet**](CustomerMetersAPI.md#CustomerMetersGet) | **Get** /customer-meters/{customer_id}/{meter_id} | Retrieve a customer-meter aggregate
[**CustomerMetersList**](CustomerMetersAPI.md#CustomerMetersList) | **Get** /customer-meters | List a customer&#39;s meter usage



## CustomerMetersGet

> CustomerMeterResponse CustomerMetersGet(ctx, customerId, meterId).Execute()

Retrieve a customer-meter aggregate



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
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID
	meterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Meter UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerMetersAPI.CustomerMetersGet(context.Background(), customerId, meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerMetersAPI.CustomerMetersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerMetersGet`: CustomerMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerMetersAPI.CustomerMetersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Customer UUID | 
**meterId** | **string** | Meter UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerMetersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomerMeterResponse**](CustomerMeterResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerMetersList

> CustomerMeterList CustomerMetersList(ctx).CustomerId(customerId).MeterId(meterId).Execute()

List a customer's meter usage



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
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID. Required — customer-meters are not queryable merchant-wide via this endpoint. Use /v2/meters for merchant-level meter definitions.
	meterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional single-meter filter. When provided, the response contains at most one entry. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerMetersAPI.CustomerMetersList(context.Background()).CustomerId(customerId).MeterId(meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerMetersAPI.CustomerMetersList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerMetersList`: CustomerMeterList
	fmt.Fprintf(os.Stdout, "Response from `CustomerMetersAPI.CustomerMetersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerMetersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Customer UUID. Required — customer-meters are not queryable merchant-wide via this endpoint. Use /v2/meters for merchant-level meter definitions. | 
 **meterId** | **string** | Optional single-meter filter. When provided, the response contains at most one entry. | 

### Return type

[**CustomerMeterList**](CustomerMeterList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

