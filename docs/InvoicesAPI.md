# \InvoicesAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceUsageExplain**](InvoicesAPI.md#InvoiceUsageExplain) | **Get** /invoices/{id}/usage-explain | Explain invoice usage charges
[**InvoicesCreate**](InvoicesAPI.md#InvoicesCreate) | **Post** /invoices | Create invoice
[**InvoicesFinalize**](InvoicesAPI.md#InvoicesFinalize) | **Post** /invoices/{id}/finalize | Finalize an invoice
[**InvoicesGet**](InvoicesAPI.md#InvoicesGet) | **Get** /invoices/{id} | Get invoice
[**InvoicesGetPdf**](InvoicesAPI.md#InvoicesGetPdf) | **Get** /invoices/{id}/pdf | Get invoice PDF link
[**InvoicesList**](InvoicesAPI.md#InvoicesList) | **Get** /invoices | List invoices
[**InvoicesMarginEstimate**](InvoicesAPI.md#InvoicesMarginEstimate) | **Post** /invoices/{id}/margin-estimate | Estimate invoice net by payment rail
[**InvoicesPay**](InvoicesAPI.md#InvoicesPay) | **Post** /invoices/{id}/pay | Record or collect a payment
[**InvoicesSend**](InvoicesAPI.md#InvoicesSend) | **Post** /invoices/{id}/send | Send an invoice
[**InvoicesUpdate**](InvoicesAPI.md#InvoicesUpdate) | **Patch** /invoices/{id} | Update invoice
[**InvoicesVoid**](InvoicesAPI.md#InvoicesVoid) | **Post** /invoices/{id}/void | Void an invoice



## InvoiceUsageExplain

> InvoiceUsageExplain200Response InvoiceUsageExplain(ctx, id).Execute()

Explain invoice usage charges



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invoice UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoiceUsageExplain(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoiceUsageExplain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceUsageExplain`: InvoiceUsageExplain200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoiceUsageExplain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceUsageExplainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceUsageExplain200Response**](InvoiceUsageExplain200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesCreate

> InvoiceResponse InvoicesCreate(ctx).InvoicesCreateRequest(invoicesCreateRequest).IdempotencyKey(idempotencyKey).Execute()

Create invoice



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
	invoicesCreateRequest := *openapiclient.NewInvoicesCreateRequest("CustomerUuid_example", int64(123), "Currency_example") // InvoicesCreateRequest | Invoice creation details
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesCreate(context.Background()).InvoicesCreateRequest(invoicesCreateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesCreate`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoicesCreateRequest** | [**InvoicesCreateRequest**](InvoicesCreateRequest.md) | Invoice creation details |
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesFinalize

> InvoiceResponse InvoicesFinalize(ctx, id).IdempotencyKey(idempotencyKey).InvoicesFinalizeRequest(invoicesFinalizeRequest).Execute()

Finalize an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	invoicesFinalizeRequest := *openapiclient.NewInvoicesFinalizeRequest() // InvoicesFinalizeRequest | Finalization options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesFinalize(context.Background(), id).IdempotencyKey(idempotencyKey).InvoicesFinalizeRequest(invoicesFinalizeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesFinalize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesFinalize`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesFinalize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesFinalizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **invoicesFinalizeRequest** | [**InvoicesFinalizeRequest**](InvoicesFinalizeRequest.md) | Finalization options |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesGet

> InvoiceResponse InvoicesGet(ctx, id).Execute()

Get invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesGet`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesGetPdf

> InvoicePdfResponse InvoicesGetPdf(ctx, id).Execute()

Get invoice PDF link



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesGetPdf(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesGetPdf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesGetPdf`: InvoicePdfResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesGetPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesGetPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvoicePdfResponse**](InvoicePdfResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesList

> InvoiceListResponse InvoicesList(ctx).Status(status).CustomerId(customerId).Limit(limit).Offset(offset).Execute()

List invoices



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
	status := "status_example" // string | Filter by invoice status (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer ID (UUID) (optional)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	offset := int32(56) // int32 | Number of results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesList(context.Background()).Status(status).CustomerId(customerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesList`: InvoiceListResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by invoice status |
 **customerId** | **string** | Filter by customer ID (UUID) |
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **offset** | **int32** | Number of results to skip | [default to 0]

### Return type

[**InvoiceListResponse**](InvoiceListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesMarginEstimate

> InvoicesMarginEstimate200Response InvoicesMarginEstimate(ctx, id).InvoicesMarginEstimateRequest(invoicesMarginEstimateRequest).IdempotencyKey(idempotencyKey).Execute()

Estimate invoice net by payment rail



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invoice ID.
	invoicesMarginEstimateRequest := *openapiclient.NewInvoicesMarginEstimateRequest() // InvoicesMarginEstimateRequest | Margin estimate inputs
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesMarginEstimate(context.Background(), id).InvoicesMarginEstimateRequest(invoicesMarginEstimateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesMarginEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesMarginEstimate`: InvoicesMarginEstimate200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesMarginEstimate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID. |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesMarginEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesMarginEstimateRequest** | [**InvoicesMarginEstimateRequest**](InvoicesMarginEstimateRequest.md) | Margin estimate inputs |
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**InvoicesMarginEstimate200Response**](InvoicesMarginEstimate200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesPay

> InvoiceResponse InvoicesPay(ctx, id).PayInvoiceRequest(payInvoiceRequest).IdempotencyKey(idempotencyKey).Execute()

Record or collect a payment



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	payInvoiceRequest := openapiclient.PayInvoiceRequest{CollectInvoicePaymentRequest: openapiclient.NewCollectInvoicePaymentRequest("collect", "direct_debit")} // PayInvoiceRequest | Payment details
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesPay(context.Background(), id).PayInvoiceRequest(payInvoiceRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesPay`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payInvoiceRequest** | [**PayInvoiceRequest**](PayInvoiceRequest.md) | Payment details |
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesSend

> InvoiceResponse InvoicesSend(ctx, id).IdempotencyKey(idempotencyKey).InvoicesSendRequest(invoicesSendRequest).Execute()

Send an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	invoicesSendRequest := *openapiclient.NewInvoicesSendRequest() // InvoicesSendRequest | Send options (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesSend(context.Background(), id).IdempotencyKey(idempotencyKey).InvoicesSendRequest(invoicesSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesSend`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **invoicesSendRequest** | [**InvoicesSendRequest**](InvoicesSendRequest.md) | Send options |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesUpdate

> InvoiceResponse InvoicesUpdate(ctx, id).InvoicesUpdateRequest(invoicesUpdateRequest).IdempotencyKey(idempotencyKey).Execute()

Update invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	invoicesUpdateRequest := *openapiclient.NewInvoicesUpdateRequest() // InvoicesUpdateRequest | Invoice update details
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesUpdate(context.Background(), id).InvoicesUpdateRequest(invoicesUpdateRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesUpdate`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invoicesUpdateRequest** | [**InvoicesUpdateRequest**](InvoicesUpdateRequest.md) | Invoice update details |
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesVoid

> InvoiceResponse InvoicesVoid(ctx, id).IdempotencyKey(idempotencyKey).InvoicesVoidRequest(invoicesVoidRequest).Execute()

Void an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	invoicesVoidRequest := *openapiclient.NewInvoicesVoidRequest() // InvoicesVoidRequest | Void details (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesVoid(context.Background(), id).IdempotencyKey(idempotencyKey).InvoicesVoidRequest(invoicesVoidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesVoid`: InvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **invoicesVoidRequest** | [**InvoicesVoidRequest**](InvoicesVoidRequest.md) | Void details |

### Return type

[**InvoiceResponse**](InvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
