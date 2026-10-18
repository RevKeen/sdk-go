# \SubscriptionsAPI

All URIs are relative to *https://api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsCancel**](SubscriptionsAPI.md#SubscriptionsCancel) | **Post** /subscriptions/{id}/cancel | Cancel subscription
[**SubscriptionsChangePlan**](SubscriptionsAPI.md#SubscriptionsChangePlan) | **Post** /subscriptions/{id}/change-plan | Change subscription plan
[**SubscriptionsCreate**](SubscriptionsAPI.md#SubscriptionsCreate) | **Post** /subscriptions | Create a new subscription
[**SubscriptionsGet**](SubscriptionsAPI.md#SubscriptionsGet) | **Get** /subscriptions/{id} | Get subscription by ID
[**SubscriptionsList**](SubscriptionsAPI.md#SubscriptionsList) | **Get** /subscriptions | List subscriptions
[**SubscriptionsPause**](SubscriptionsAPI.md#SubscriptionsPause) | **Post** /subscriptions/{id}/pause | Pause subscription
[**SubscriptionsPreviewChangePlan**](SubscriptionsAPI.md#SubscriptionsPreviewChangePlan) | **Post** /subscriptions/{id}/preview-change-plan | Preview a plan change
[**SubscriptionsResume**](SubscriptionsAPI.md#SubscriptionsResume) | **Post** /subscriptions/{id}/resume | Resume subscription
[**SubscriptionsUpdate**](SubscriptionsAPI.md#SubscriptionsUpdate) | **Patch** /subscriptions/{id} | Update subscription details



## SubscriptionsCancel

> SubscriptionCancelSubscriptionResponse SubscriptionsCancel(ctx, id).IdempotencyKey(idempotencyKey).CancelSubscriptionInput(cancelSubscriptionInput).Execute()

Cancel subscription



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	cancelSubscriptionInput := *openapiclient.NewCancelSubscriptionInput("Mode_example") // CancelSubscriptionInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsCancel(context.Background(), id).IdempotencyKey(idempotencyKey).CancelSubscriptionInput(cancelSubscriptionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsCancel`: SubscriptionCancelSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **cancelSubscriptionInput** | [**CancelSubscriptionInput**](CancelSubscriptionInput.md) |  |

### Return type

[**SubscriptionCancelSubscriptionResponse**](SubscriptionCancelSubscriptionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsChangePlan

> SubscriptionChangePlanResponse SubscriptionsChangePlan(ctx, id).IdempotencyKey(idempotencyKey).SubscriptionsChangePlanRequest(subscriptionsChangePlanRequest).Execute()

Change subscription plan



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	subscriptionsChangePlanRequest := *openapiclient.NewSubscriptionsChangePlanRequest("NewPlanId_example") // SubscriptionsChangePlanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsChangePlan(context.Background(), id).IdempotencyKey(idempotencyKey).SubscriptionsChangePlanRequest(subscriptionsChangePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsChangePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsChangePlan`: SubscriptionChangePlanResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsChangePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsChangePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **subscriptionsChangePlanRequest** | [**SubscriptionsChangePlanRequest**](SubscriptionsChangePlanRequest.md) |  |

### Return type

[**SubscriptionChangePlanResponse**](SubscriptionChangePlanResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsCreate

> SubscriptionCreateResponse SubscriptionsCreate(ctx).IdempotencyKey(idempotencyKey).SubscriptionsCreateRequest(subscriptionsCreateRequest).Execute()

Create a new subscription



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
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	subscriptionsCreateRequest := *openapiclient.NewSubscriptionsCreateRequest("CustomerId_example") // SubscriptionsCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsCreate(context.Background()).IdempotencyKey(idempotencyKey).SubscriptionsCreateRequest(subscriptionsCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsCreate`: SubscriptionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **subscriptionsCreateRequest** | [**SubscriptionsCreateRequest**](SubscriptionsCreateRequest.md) |  |

### Return type

[**SubscriptionCreateResponse**](SubscriptionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGet

> SubscriptionRetrieveResponse SubscriptionsGet(ctx, id).Execute()

Get subscription by ID



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGet`: SubscriptionRetrieveResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionRetrieveResponse**](SubscriptionRetrieveResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsList

> SubscriptionListResponse SubscriptionsList(ctx).Page(page).Limit(limit).Status(status).CustomerId(customerId).Execute()

List subscriptions



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
	page := int32(56) // int32 | Page number (optional) (default to 1)
	limit := int32(56) // int32 | Maximum number of results (1-100) (optional) (default to 20)
	status := "status_example" // string | Filter by subscription status (optional)
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by customer UUID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsList(context.Background()).Page(page).Limit(limit).Status(status).CustomerId(customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsList`: SubscriptionListResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **limit** | **int32** | Maximum number of results (1-100) | [default to 20]
 **status** | **string** | Filter by subscription status |
 **customerId** | **string** | Filter by customer UUID |

### Return type

[**SubscriptionListResponse**](SubscriptionListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPause

> SubscriptionPauseResponse SubscriptionsPause(ctx, id).IdempotencyKey(idempotencyKey).Body(body).Execute()

Pause subscription



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPause(context.Background(), id).IdempotencyKey(idempotencyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPause`: SubscriptionPauseResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **body** | **map[string]interface{}** |  |

### Return type

[**SubscriptionPauseResponse**](SubscriptionPauseResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPreviewChangePlan

> SubscriptionChangePlanPreviewResponse SubscriptionsPreviewChangePlan(ctx, id).IdempotencyKey(idempotencyKey).PreviewSubscriptionPlanChangeInput(previewSubscriptionPlanChangeInput).Execute()

Preview a plan change



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	previewSubscriptionPlanChangeInput := *openapiclient.NewPreviewSubscriptionPlanChangeInput("b5d8e2f1-3a6c-4d9b-8e7f-1c4a2b5d6e73") // PreviewSubscriptionPlanChangeInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPreviewChangePlan(context.Background(), id).IdempotencyKey(idempotencyKey).PreviewSubscriptionPlanChangeInput(previewSubscriptionPlanChangeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPreviewChangePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPreviewChangePlan`: SubscriptionChangePlanPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPreviewChangePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPreviewChangePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **previewSubscriptionPlanChangeInput** | [**PreviewSubscriptionPlanChangeInput**](PreviewSubscriptionPlanChangeInput.md) |  |

### Return type

[**SubscriptionChangePlanPreviewResponse**](SubscriptionChangePlanPreviewResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsResume

> SubscriptionResumeResponse SubscriptionsResume(ctx, id).IdempotencyKey(idempotencyKey).Execute()

Resume subscription



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsResume(context.Background(), id).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsResume`: SubscriptionResumeResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |

### Return type

[**SubscriptionResumeResponse**](SubscriptionResumeResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdate

> SubscriptionUpdateResponse SubscriptionsUpdate(ctx, id).IdempotencyKey(idempotencyKey).SubscriptionUpdateInput(subscriptionUpdateInput).Execute()

Update subscription details



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription UUID
	idempotencyKey := "idempotencyKey_example" // string | Stable caller-owned key for this logical mutation. Reuse only with an identical request. (optional)
	subscriptionUpdateInput := *openapiclient.NewSubscriptionUpdateInput() // SubscriptionUpdateInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsUpdate(context.Background(), id).IdempotencyKey(idempotencyKey).SubscriptionUpdateInput(subscriptionUpdateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsUpdate`: SubscriptionUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Stable caller-owned key for this logical mutation. Reuse only with an identical request. |
 **subscriptionUpdateInput** | [**SubscriptionUpdateInput**](SubscriptionUpdateInput.md) |  |

### Return type

[**SubscriptionUpdateResponse**](SubscriptionUpdateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
