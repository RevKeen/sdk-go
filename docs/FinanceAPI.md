# \FinanceAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FinanceGetIncome**](FinanceAPI.md#FinanceGetIncome) | **Get** /finance/income | Get income report
[**FinanceGetSummary**](FinanceAPI.md#FinanceGetSummary) | **Get** /finance/summary | Get finance summary



## FinanceGetIncome

> FinanceGetIncome200Response FinanceGetIncome(ctx).StartDate(startDate).EndDate(endDate).GroupBy(groupBy).SubscriptionId(subscriptionId).Status(status).Execute()

Get income report



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
	startDate := "2024-01-01" // string | Start date
	endDate := "2024-12-31" // string | End date
	groupBy := "groupBy_example" // string | Group results by time period (optional) (default to "day")
	subscriptionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by subscription UUID (optional)
	status := "status_example" // string | Filter by status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FinanceAPI.FinanceGetIncome(context.Background()).StartDate(startDate).EndDate(endDate).GroupBy(groupBy).SubscriptionId(subscriptionId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.FinanceGetIncome``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinanceGetIncome`: FinanceGetIncome200Response
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.FinanceGetIncome`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFinanceGetIncomeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Start date | 
 **endDate** | **string** | End date | 
 **groupBy** | **string** | Group results by time period | [default to &quot;day&quot;]
 **subscriptionId** | **string** | Filter by subscription UUID | 
 **status** | **string** | Filter by status | 

### Return type

[**FinanceGetIncome200Response**](FinanceGetIncome200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinanceGetSummary

> FinanceGetSummary200Response FinanceGetSummary(ctx).Execute()

Get finance summary



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
	resp, r, err := apiClient.FinanceAPI.FinanceGetSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FinanceAPI.FinanceGetSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinanceGetSummary`: FinanceGetSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `FinanceAPI.FinanceGetSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFinanceGetSummaryRequest struct via the builder pattern


### Return type

[**FinanceGetSummary200Response**](FinanceGetSummary200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

