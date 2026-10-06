# \BillingAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingOverviewGet**](BillingAPI.md#BillingOverviewGet) | **Get** /billing/overview | Get merchant billing overview



## BillingOverviewGet

> BillingOverviewGet200Response BillingOverviewGet(ctx).Execute()

Get merchant billing overview



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
	resp, r, err := apiClient.BillingAPI.BillingOverviewGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingOverviewGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingOverviewGet`: BillingOverviewGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingOverviewGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingOverviewGetRequest struct via the builder pattern


### Return type

[**BillingOverviewGet200Response**](BillingOverviewGet200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
