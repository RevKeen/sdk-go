# \UsageAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageBalanceGet**](UsageAPI.md#UsageBalanceGet) | **Get** /usage/balance | Get current usage balance



## UsageBalanceGet

> UsageBalanceResponse UsageBalanceGet(ctx).SubscriptionId(subscriptionId).CustomerId(customerId).ExternalCustomerId(externalCustomerId).MeterId(meterId).Execute()

Get current usage balance



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
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter to a specific subscription (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter to a specific customer by RevKeen ID (optional)
	externalCustomerId := "usr_123" // string | Filter to a customer by your external identifier (optional)
	meterId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter to a single meter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageBalanceGet(context.Background()).SubscriptionId(subscriptionId).CustomerId(customerId).ExternalCustomerId(externalCustomerId).MeterId(meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageBalanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageBalanceGet`: UsageBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | Filter to a specific subscription |
 **customerId** | **string** | Filter to a specific customer by RevKeen ID |
 **externalCustomerId** | **string** | Filter to a customer by your external identifier |
 **meterId** | **string** | Filter to a single meter |

### Return type

[**UsageBalanceResponse**](UsageBalanceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
