# \AnalyticsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsCollectionsGetDso**](AnalyticsAPI.md#AnalyticsCollectionsGetDso) | **Get** /analytics/collections/dso | Days Sales Outstanding
[**AnalyticsCustomersGetLtv**](AnalyticsAPI.md#AnalyticsCustomersGetLtv) | **Get** /analytics/customers/ltv | Customer LTV
[**AnalyticsInvoicesGetArAging**](AnalyticsAPI.md#AnalyticsInvoicesGetArAging) | **Get** /analytics/invoices/ar-aging | A/R Aging Report
[**AnalyticsRevenueGetMrrSummary**](AnalyticsAPI.md#AnalyticsRevenueGetMrrSummary) | **Get** /analytics/revenue/mrr-summary | MRR Summary
[**AnalyticsRevenueGetTimeSeries**](AnalyticsAPI.md#AnalyticsRevenueGetTimeSeries) | **Get** /analytics/revenue/time-series | Revenue Time Series



## AnalyticsCollectionsGetDso

> AnalyticsCollectionsGetDso200Response AnalyticsCollectionsGetDso(ctx).StartDate(startDate).EndDate(endDate).Execute()

Days Sales Outstanding



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
	startDate := time.Now() // time.Time | Start date (ISO 8601) (optional)
	endDate := time.Now() // time.Time | End date (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsCollectionsGetDso(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsCollectionsGetDso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCollectionsGetDso`: AnalyticsCollectionsGetDso200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsCollectionsGetDso`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCollectionsGetDsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date (ISO 8601) | 
 **endDate** | **time.Time** | End date (ISO 8601) | 

### Return type

[**AnalyticsCollectionsGetDso200Response**](AnalyticsCollectionsGetDso200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsCustomersGetLtv

> AnalyticsCustomersGetLtv200Response AnalyticsCustomersGetLtv(ctx).TopN(topN).Execute()

Customer LTV



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
	topN := float32(10) // float32 | Number of top customers to return (1-100) (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsCustomersGetLtv(context.Background()).TopN(topN).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsCustomersGetLtv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsCustomersGetLtv`: AnalyticsCustomersGetLtv200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsCustomersGetLtv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsCustomersGetLtvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topN** | **float32** | Number of top customers to return (1-100) | [default to 10]

### Return type

[**AnalyticsCustomersGetLtv200Response**](AnalyticsCustomersGetLtv200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsInvoicesGetArAging

> AnalyticsInvoicesGetArAging200Response AnalyticsInvoicesGetArAging(ctx).AsOfDate(asOfDate).Execute()

A/R Aging Report



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
	asOfDate := time.Now() // time.Time | Calculate aging as of this date (ISO 8601) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsInvoicesGetArAging(context.Background()).AsOfDate(asOfDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsInvoicesGetArAging``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsInvoicesGetArAging`: AnalyticsInvoicesGetArAging200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsInvoicesGetArAging`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsInvoicesGetArAgingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asOfDate** | **time.Time** | Calculate aging as of this date (ISO 8601) | 

### Return type

[**AnalyticsInvoicesGetArAging200Response**](AnalyticsInvoicesGetArAging200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRevenueGetMrrSummary

> AnalyticsRevenueGetMrrSummary200Response AnalyticsRevenueGetMrrSummary(ctx).AsOfDate(asOfDate).Execute()

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
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsRevenueGetMrrSummary(context.Background()).AsOfDate(asOfDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsRevenueGetMrrSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRevenueGetMrrSummary`: AnalyticsRevenueGetMrrSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsRevenueGetMrrSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRevenueGetMrrSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asOfDate** | **time.Time** | Calculate MRR as of this date (ISO 8601) | 

### Return type

[**AnalyticsRevenueGetMrrSummary200Response**](AnalyticsRevenueGetMrrSummary200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsRevenueGetTimeSeries

> AnalyticsRevenueGetTimeSeries200Response AnalyticsRevenueGetTimeSeries(ctx).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()

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
	resp, r, err := apiClient.AnalyticsAPI.AnalyticsRevenueGetTimeSeries(context.Background()).StartDate(startDate).EndDate(endDate).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsRevenueGetTimeSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsRevenueGetTimeSeries`: AnalyticsRevenueGetTimeSeries200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.AnalyticsRevenueGetTimeSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsRevenueGetTimeSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **time.Time** | Start date (ISO 8601) | 
 **endDate** | **time.Time** | End date (ISO 8601) | 
 **granularity** | **string** | Time granularity for data points | [default to &quot;month&quot;]

### Return type

[**AnalyticsRevenueGetTimeSeries200Response**](AnalyticsRevenueGetTimeSeries200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

