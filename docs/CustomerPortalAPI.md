# \CustomerPortalAPI

All URIs are relative to *https://staging-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerPortalCustomerGet**](CustomerPortalAPI.md#CustomerPortalCustomerGet) | **Get** /customer-portal/customer | Retrieve the authenticated customer
[**CustomerPortalInvoicesGet**](CustomerPortalAPI.md#CustomerPortalInvoicesGet) | **Get** /customer-portal/invoices/{id} | Retrieve an invoice
[**CustomerPortalInvoicesList**](CustomerPortalAPI.md#CustomerPortalInvoicesList) | **Get** /customer-portal/invoices | List the authenticated customer&#39;s invoices
[**CustomerPortalMandatesCancel**](CustomerPortalAPI.md#CustomerPortalMandatesCancel) | **Post** /customer-portal/mandates/{id}/cancel | Cancel a Direct Debit mandate
[**CustomerPortalMandatesList**](CustomerPortalAPI.md#CustomerPortalMandatesList) | **Get** /customer-portal/mandates | List the authenticated customer&#39;s Direct Debit mandates
[**CustomerPortalMandatesReauthorize**](CustomerPortalAPI.md#CustomerPortalMandatesReauthorize) | **Post** /customer-portal/mandates/{id}/re-authorize | Re-authorise a Direct Debit mandate with new bank details
[**CustomerPortalMandatesReauthorizeSendOtp**](CustomerPortalAPI.md#CustomerPortalMandatesReauthorizeSendOtp) | **Post** /customer-portal/mandates/{id}/re-authorize/send-otp | Send a step-up OTP for a Direct Debit bank-detail change
[**CustomerPortalMandatesReauthorizeVerifyOtp**](CustomerPortalAPI.md#CustomerPortalMandatesReauthorizeVerifyOtp) | **Post** /customer-portal/mandates/{id}/re-authorize/verify-otp | Verify a step-up OTP and receive a change-of-bank token
[**CustomerPortalSessionsCreate**](CustomerPortalAPI.md#CustomerPortalSessionsCreate) | **Post** /customer-portal/sessions | Create a customer-portal session
[**CustomerPortalSubscriptionsCancel**](CustomerPortalAPI.md#CustomerPortalSubscriptionsCancel) | **Post** /customer-portal/subscriptions/{id}/cancel | Cancel a subscription
[**CustomerPortalSubscriptionsGet**](CustomerPortalAPI.md#CustomerPortalSubscriptionsGet) | **Get** /customer-portal/subscriptions/{id} | Retrieve a subscription
[**CustomerPortalSubscriptionsList**](CustomerPortalAPI.md#CustomerPortalSubscriptionsList) | **Get** /customer-portal/subscriptions | List the authenticated customer&#39;s subscriptions



## CustomerPortalCustomerGet

> PortalCustomerResponse CustomerPortalCustomerGet(ctx).Execute()

Retrieve the authenticated customer



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
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalCustomerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalCustomerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalCustomerGet`: PortalCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalCustomerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalCustomerGetRequest struct via the builder pattern


### Return type

[**PortalCustomerResponse**](PortalCustomerResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalInvoicesGet

> PortalInvoiceResponse CustomerPortalInvoicesGet(ctx, id).Execute()

Retrieve an invoice



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Invoice ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalInvoicesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalInvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalInvoicesGet`: PortalInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalInvoicesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalInvoiceResponse**](PortalInvoiceResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalInvoicesList

> PortalInvoiceList CustomerPortalInvoicesList(ctx).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the authenticated customer's invoices



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
	limit := int32(56) // int32 | Maximum number of results to return (1-100, default 20). (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return results created before the row with this ID. (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return results created after the row with this ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalInvoicesList(context.Background()).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalInvoicesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalInvoicesList`: PortalInvoiceList
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalInvoicesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalInvoicesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results to return (1-100, default 20). | [default to 20]
 **startingAfter** | **string** | Cursor — return results created before the row with this ID. | 
 **endingBefore** | **string** | Cursor — return results created after the row with this ID. | 

### Return type

[**PortalInvoiceList**](PortalInvoiceList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalMandatesCancel

> CustomerPortalMandatesCancel200Response CustomerPortalMandatesCancel(ctx, id).PaymentLinksExpireRequest(paymentLinksExpireRequest).Execute()

Cancel a Direct Debit mandate



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mandate ID
	paymentLinksExpireRequest := *openapiclient.NewPaymentLinksExpireRequest() // PaymentLinksExpireRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalMandatesCancel(context.Background(), id).PaymentLinksExpireRequest(paymentLinksExpireRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalMandatesCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalMandatesCancel`: CustomerPortalMandatesCancel200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalMandatesCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mandate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalMandatesCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paymentLinksExpireRequest** | [**PaymentLinksExpireRequest**](PaymentLinksExpireRequest.md) |  | 

### Return type

[**CustomerPortalMandatesCancel200Response**](CustomerPortalMandatesCancel200Response.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalMandatesList

> PortalMandateList CustomerPortalMandatesList(ctx).Execute()

List the authenticated customer's Direct Debit mandates



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
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalMandatesList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalMandatesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalMandatesList`: PortalMandateList
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalMandatesList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalMandatesListRequest struct via the builder pattern


### Return type

[**PortalMandateList**](PortalMandateList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalMandatesReauthorize

> PortalReauthorizeMandateResponse CustomerPortalMandatesReauthorize(ctx, id).PortalReauthorizeMandateRequest(portalReauthorizeMandateRequest).Execute()

Re-authorise a Direct Debit mandate with new bank details



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | ID of the mandate being amended.
	portalReauthorizeMandateRequest := *openapiclient.NewPortalReauthorizeMandateRequest("AccountHolderName_example", "SortCode_example", "AccountNumber_example", "VerificationToken_example") // PortalReauthorizeMandateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalMandatesReauthorize(context.Background(), id).PortalReauthorizeMandateRequest(portalReauthorizeMandateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalMandatesReauthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalMandatesReauthorize`: PortalReauthorizeMandateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalMandatesReauthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the mandate being amended. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalMandatesReauthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalReauthorizeMandateRequest** | [**PortalReauthorizeMandateRequest**](PortalReauthorizeMandateRequest.md) |  | 

### Return type

[**PortalReauthorizeMandateResponse**](PortalReauthorizeMandateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalMandatesReauthorizeSendOtp

> PortalReauthorizeSendOtpResponse CustomerPortalMandatesReauthorizeSendOtp(ctx, id).Execute()

Send a step-up OTP for a Direct Debit bank-detail change



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalMandatesReauthorizeSendOtp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalMandatesReauthorizeSendOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalMandatesReauthorizeSendOtp`: PortalReauthorizeSendOtpResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalMandatesReauthorizeSendOtp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalMandatesReauthorizeSendOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalReauthorizeSendOtpResponse**](PortalReauthorizeSendOtpResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalMandatesReauthorizeVerifyOtp

> PortalReauthorizeVerifyOtpResponse CustomerPortalMandatesReauthorizeVerifyOtp(ctx, id).PortalReauthorizeVerifyOtpRequest(portalReauthorizeVerifyOtpRequest).Execute()

Verify a step-up OTP and receive a change-of-bank token



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	portalReauthorizeVerifyOtpRequest := *openapiclient.NewPortalReauthorizeVerifyOtpRequest("Code_example") // PortalReauthorizeVerifyOtpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalMandatesReauthorizeVerifyOtp(context.Background(), id).PortalReauthorizeVerifyOtpRequest(portalReauthorizeVerifyOtpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalMandatesReauthorizeVerifyOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalMandatesReauthorizeVerifyOtp`: PortalReauthorizeVerifyOtpResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalMandatesReauthorizeVerifyOtp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalMandatesReauthorizeVerifyOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalReauthorizeVerifyOtpRequest** | [**PortalReauthorizeVerifyOtpRequest**](PortalReauthorizeVerifyOtpRequest.md) |  | 

### Return type

[**PortalReauthorizeVerifyOtpResponse**](PortalReauthorizeVerifyOtpResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSessionsCreate

> CustomerPortalSessionCreateResponse CustomerPortalSessionsCreate(ctx).CreateCustomerPortalSessionRequest(createCustomerPortalSessionRequest).Execute()

Create a customer-portal session



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
	createCustomerPortalSessionRequest := *openapiclient.NewCreateCustomerPortalSessionRequest("cus_a1b2c3d4e5f6") // CreateCustomerPortalSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSessionsCreate(context.Background()).CreateCustomerPortalSessionRequest(createCustomerPortalSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSessionsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSessionsCreate`: CustomerPortalSessionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSessionsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSessionsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomerPortalSessionRequest** | [**CreateCustomerPortalSessionRequest**](CreateCustomerPortalSessionRequest.md) |  | 

### Return type

[**CustomerPortalSessionCreateResponse**](CustomerPortalSessionCreateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsCancel

> PortalSubscriptionCancelResponse CustomerPortalSubscriptionsCancel(ctx, id).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()

Cancel a subscription



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription ID to cancel.
	cancelSubscriptionRequest := *openapiclient.NewCancelSubscriptionRequest() // CancelSubscriptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsCancel(context.Background(), id).CancelSubscriptionRequest(cancelSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsCancel`: PortalSubscriptionCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelSubscriptionRequest** | [**CancelSubscriptionRequest**](CancelSubscriptionRequest.md) |  | 

### Return type

[**PortalSubscriptionCancelResponse**](PortalSubscriptionCancelResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsGet

> PortalSubscriptionResponse CustomerPortalSubscriptionsGet(ctx, id).Execute()

Retrieve a subscription



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Subscription ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsGet`: PortalSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortalSubscriptionResponse**](PortalSubscriptionResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerPortalSubscriptionsList

> PortalSubscriptionList CustomerPortalSubscriptionsList(ctx).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the authenticated customer's subscriptions



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
	limit := int32(56) // int32 | Maximum number of results to return (1-100, default 20). (optional) (default to 20)
	startingAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return results created before the row with this ID. (optional)
	endingBefore := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Cursor — return results created after the row with this ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.CustomerPortalSubscriptionsList(context.Background()).Limit(limit).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.CustomerPortalSubscriptionsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerPortalSubscriptionsList`: PortalSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.CustomerPortalSubscriptionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerPortalSubscriptionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results to return (1-100, default 20). | [default to 20]
 **startingAfter** | **string** | Cursor — return results created before the row with this ID. | 
 **endingBefore** | **string** | Cursor — return results created after the row with this ID. | 

### Return type

[**PortalSubscriptionList**](PortalSubscriptionList.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

