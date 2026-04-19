# \EntitlementsAPI

All URIs are relative to *https://sandbox-api.revkeen.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsCheck**](EntitlementsAPI.md#EntitlementsCheck) | **Get** /entitlements/check | Check entitlement access
[**EntitlementsList**](EntitlementsAPI.md#EntitlementsList) | **Get** /entitlements | List entitlements



## EntitlementsCheck

> EntitlementCheckResponse EntitlementsCheck(ctx).CustomerId(customerId).BenefitKey(benefitKey).Execute()

Check entitlement access



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
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID (required)
	benefitKey := "benefitKey_example" // string | Benefit key to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsCheck(context.Background()).CustomerId(customerId).BenefitKey(benefitKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsCheck`: EntitlementCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Customer UUID (required) | 
 **benefitKey** | **string** | Benefit key to check | 

### Return type

[**EntitlementCheckResponse**](EntitlementCheckResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsList

> EntitlementListResponse EntitlementsList(ctx).CustomerId(customerId).IncludeExpired(includeExpired).BenefitType(benefitType).Category(category).Limit(limit).Offset(offset).Execute()

List entitlements



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
	customerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Customer UUID (required)
	includeExpired := true // bool | Include expired entitlements (optional) (default to false)
	benefitType := "benefitType_example" // string | Filter by benefit type (optional)
	category := "category_example" // string | Filter by category (optional)
	limit := int32(56) // int32 | Maximum results (1-100) (optional) (default to 50)
	offset := int32(56) // int32 | Results to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsList(context.Background()).CustomerId(customerId).IncludeExpired(includeExpired).BenefitType(benefitType).Category(category).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsList`: EntitlementListResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** | Customer UUID (required) | 
 **includeExpired** | **bool** | Include expired entitlements | [default to false]
 **benefitType** | **string** | Filter by benefit type | 
 **category** | **string** | Filter by category | 
 **limit** | **int32** | Maximum results (1-100) | [default to 50]
 **offset** | **int32** | Results to skip | [default to 0]

### Return type

[**EntitlementListResponse**](EntitlementListResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

