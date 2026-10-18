# \PaymentMethodsAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentMethodsDetach**](PaymentMethodsAPI.md#PaymentMethodsDetach) | **Post** /payment-methods/{id}/detach | Detach a payment method from its customer
[**PaymentMethodsSetDefault**](PaymentMethodsAPI.md#PaymentMethodsSetDefault) | **Post** /payment-methods/{id}/set-default | Set a payment method as its customer&#39;s default



## PaymentMethodsDetach

> PaymentMethod PaymentMethodsDetach(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Detach a payment method from its customer



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | Payment method UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.PaymentMethodsDetach(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.PaymentMethodsDetach``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodsDetach`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.PaymentMethodsDetach`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment method UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsDetachRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentMethodsSetDefault

> PaymentMethod PaymentMethodsSetDefault(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Set a payment method as its customer's default



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | Payment method UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.PaymentMethodsSetDefault(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.PaymentMethodsSetDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentMethodsSetDefault`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.PaymentMethodsSetDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment method UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentMethodsSetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
