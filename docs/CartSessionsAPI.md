# \CartSessionsAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CartSessionsAddLineItem**](CartSessionsAPI.md#CartSessionsAddLineItem) | **Post** /cart-sessions/{id}/line-items | Add a line item to a cart session
[**CartSessionsApplyDiscountCode**](CartSessionsAPI.md#CartSessionsApplyDiscountCode) | **Post** /cart-sessions/{id}/discount-code | Set or clear a cart discount code
[**CartSessionsConvert**](CartSessionsAPI.md#CartSessionsConvert) | **Post** /cart-sessions/{id}/convert | Convert a cart session into a checkout session
[**CartSessionsCreate**](CartSessionsAPI.md#CartSessionsCreate) | **Post** /cart-sessions | Create a cart session
[**CartSessionsGet**](CartSessionsAPI.md#CartSessionsGet) | **Get** /cart-sessions/{id} | Retrieve a cart session
[**CartSessionsRemoveLineItem**](CartSessionsAPI.md#CartSessionsRemoveLineItem) | **Delete** /cart-sessions/{id}/line-items/{lineId} | Remove a line item from a cart session
[**CartSessionsSetContact**](CartSessionsAPI.md#CartSessionsSetContact) | **Post** /cart-sessions/{id}/contact | Capture customer email + consent on a cart
[**CartSessionsToggleAddOn**](CartSessionsAPI.md#CartSessionsToggleAddOn) | **Post** /cart-sessions/{id}/add-ons | Toggle an add-on on a cart session
[**CartSessionsUpdateLineItem**](CartSessionsAPI.md#CartSessionsUpdateLineItem) | **Patch** /cart-sessions/{id}/line-items/{lineId} | Update a line item&#39;s quantity



## CartSessionsAddLineItem

> CartSessionResponse CartSessionsAddLineItem(ctx, id).CartLineItemInput(cartLineItemInput).Execute()

Add a line item to a cart session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	cartLineItemInput := *openapiclient.NewCartLineItemInput("11111111-1111-1111-1111-111111111111", "Annual plan", int32(1), int32(9900), "GBP") // CartLineItemInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsAddLineItem(context.Background(), id).CartLineItemInput(cartLineItemInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsAddLineItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsAddLineItem`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsAddLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsAddLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cartLineItemInput** | [**CartLineItemInput**](CartLineItemInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsApplyDiscountCode

> CartSessionResponse CartSessionsApplyDiscountCode(ctx, id).ApplyCartDiscountCodeInput(applyCartDiscountCodeInput).Execute()

Set or clear a cart discount code



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	applyCartDiscountCodeInput := *openapiclient.NewApplyCartDiscountCodeInput("SAVE10") // ApplyCartDiscountCodeInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsApplyDiscountCode(context.Background(), id).ApplyCartDiscountCodeInput(applyCartDiscountCodeInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsApplyDiscountCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsApplyDiscountCode`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsApplyDiscountCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsApplyDiscountCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyCartDiscountCodeInput** | [**ApplyCartDiscountCodeInput**](ApplyCartDiscountCodeInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsConvert

> CartConversionResponse CartSessionsConvert(ctx, id).Execute()

Convert a cart session into a checkout session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsConvert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsConvert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsConvert`: CartConversionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsConvert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CartConversionResponse**](CartConversionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsCreate

> CartSessionResponse CartSessionsCreate(ctx).CreateCartSessionInput(createCartSessionInput).Execute()

Create a cart session



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
	createCartSessionInput := *openapiclient.NewCreateCartSessionInput("GBP") // CreateCartSessionInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsCreate(context.Background()).CreateCartSessionInput(createCartSessionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsCreate`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCartSessionInput** | [**CreateCartSessionInput**](CreateCartSessionInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsGet

> CartSessionResponse CartSessionsGet(ctx, id).Execute()

Retrieve a cart session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsGet`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsRemoveLineItem

> CartSessionResponse CartSessionsRemoveLineItem(ctx, id, lineId).Execute()

Remove a line item from a cart session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	lineId := "cli_550e8400-e29b-41d4-a716-446655440000" // string | Cart line item id (the `id` field returned on `cart_session.line_items[]`).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsRemoveLineItem(context.Background(), id, lineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsRemoveLineItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsRemoveLineItem`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsRemoveLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**lineId** | **string** | Cart line item id (the &#x60;id&#x60; field returned on &#x60;cart_session.line_items[]&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsRemoveLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsSetContact

> CartSessionResponse CartSessionsSetContact(ctx, id).SetCartContactInput(setCartContactInput).Execute()

Capture customer email + consent on a cart



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	setCartContactInput := *openapiclient.NewSetCartContactInput() // SetCartContactInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsSetContact(context.Background(), id).SetCartContactInput(setCartContactInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsSetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsSetContact`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsSetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsSetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setCartContactInput** | [**SetCartContactInput**](SetCartContactInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsToggleAddOn

> CartSessionResponse CartSessionsToggleAddOn(ctx, id).ToggleCartAddOnInput(toggleCartAddOnInput).Execute()

Toggle an add-on on a cart session



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	toggleCartAddOnInput := *openapiclient.NewToggleCartAddOnInput("22222222-2222-2222-2222-222222222222", true) // ToggleCartAddOnInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsToggleAddOn(context.Background(), id).ToggleCartAddOnInput(toggleCartAddOnInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsToggleAddOn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsToggleAddOn`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsToggleAddOn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsToggleAddOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toggleCartAddOnInput** | [**ToggleCartAddOnInput**](ToggleCartAddOnInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CartSessionsUpdateLineItem

> CartSessionResponse CartSessionsUpdateLineItem(ctx, id, lineId).UpdateCartLineItemInput(updateCartLineItemInput).Execute()

Update a line item's quantity



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
	id := "550e8400-e29b-41d4-a716-446655440000" // string | 
	lineId := "cli_550e8400-e29b-41d4-a716-446655440000" // string | Cart line item id (the `id` field returned on `cart_session.line_items[]`).
	updateCartLineItemInput := *openapiclient.NewUpdateCartLineItemInput(int32(2)) // UpdateCartLineItemInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartSessionsAPI.CartSessionsUpdateLineItem(context.Background(), id, lineId).UpdateCartLineItemInput(updateCartLineItemInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartSessionsAPI.CartSessionsUpdateLineItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CartSessionsUpdateLineItem`: CartSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CartSessionsAPI.CartSessionsUpdateLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**lineId** | **string** | Cart line item id (the &#x60;id&#x60; field returned on &#x60;cart_session.line_items[]&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCartSessionsUpdateLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCartLineItemInput** | [**UpdateCartLineItemInput**](UpdateCartLineItemInput.md) |  | 

### Return type

[**CartSessionResponse**](CartSessionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

