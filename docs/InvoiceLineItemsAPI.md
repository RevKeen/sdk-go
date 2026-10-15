# \InvoiceLineItemsAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceLineItemsListUsageEvents**](InvoiceLineItemsAPI.md#InvoiceLineItemsListUsageEvents) | **Get** /invoice-line-items/{id}/usage-events | List usage events for an invoice line



## InvoiceLineItemsListUsageEvents

> InvoiceLineUsageEventsList InvoiceLineItemsListUsageEvents(ctx, id).StartingAfter(startingAfter).Limit(limit).IncludeExcluded(includeExcluded).Execute()

List usage events for an invoice line



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
	id := "11111111-1111-4111-8111-111111111111" // string | Invoice line item ID
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor: return entries with id greater than this value (optional)
	limit := "50" // string | Max billed entries to return (default 50, max 100) (optional)
	includeExcluded := "false" // string | When true, also return excluded-but-considered decisions for the line period (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceLineItemsAPI.InvoiceLineItemsListUsageEvents(context.Background(), id).StartingAfter(startingAfter).Limit(limit).IncludeExcluded(includeExcluded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceLineItemsAPI.InvoiceLineItemsListUsageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceLineItemsListUsageEvents`: InvoiceLineUsageEventsList
	fmt.Fprintf(os.Stdout, "Response from `InvoiceLineItemsAPI.InvoiceLineItemsListUsageEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice line item ID |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceLineItemsListUsageEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startingAfter** | **string** | Cursor: return entries with id greater than this value |
 **limit** | **string** | Max billed entries to return (default 50, max 100) |
 **includeExcluded** | **string** | When true, also return excluded-but-considered decisions for the line period |

### Return type

[**InvoiceLineUsageEventsList**](InvoiceLineUsageEventsList.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
