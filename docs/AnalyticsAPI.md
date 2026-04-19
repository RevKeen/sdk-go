# \AnalyticsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsRevenueMrrSummary**](AnalyticsAPI.md#AnalyticsRevenueMrrSummary) | **Get** /analytics/revenue/mrr-summary | MRR Summary
[**AnalyticsRevenueTimeSeries**](AnalyticsAPI.md#AnalyticsRevenueTimeSeries) | **Get** /analytics/revenue/time-series | Revenue Time Series



## AnalyticsRevenueMrrSummary

> AnalyticsRevenueMrrSummary200Response AnalyticsRevenueMrrSummary(ctx).AsOfDate(asOfDate).Execute()

MRR Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	asOfDate := time.Now() // time.Time | Calculate MRR as of this date (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsRevenueMrrSummary(context.Background()).AsOfDate(asOfDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsRevenueMrrSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRevenueMrrSummary`: AnalyticsRevenueMrrSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsRevenueMrrSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRevenueMrrSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asOfDate** | **time.Time** | Calculate MRR as of this date (ISO 8601) | 

### Return type

[**AnalyticsRevenueMrrSummary200Response**](AnalyticsRevenueMrrSummary200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRevenueTimeSeries

> AnalyticsRevenueTimeSeries200Response AnalyticsRevenueTimeSeries(ctx).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()

Revenue Time Series



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/revkeen/revkeen-go"
)

func main() {
	startDate := time.Now() // time.Time | Start date (ISO 8601)
	endDate := time.Now() // time.Time | End date (ISO 8601)
	granularity := "granularity_example" // string | Time granularity for data points (optional) (default to "month")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsRevenueTimeSeries(context.Background()).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsRevenueTimeSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRevenueTimeSeries`: AnalyticsRevenueTimeSeries200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsRevenueTimeSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRevenueTimeSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date (ISO 8601) | 
 **endDate** | **time.Time** | End date (ISO 8601) | 
 **granularity** | **string** | Time granularity for data points | [default to &quot;month&quot;]

### Return type

[**AnalyticsRevenueTimeSeries200Response**](AnalyticsRevenueTimeSeries200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

