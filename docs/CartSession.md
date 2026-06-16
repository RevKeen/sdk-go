# CartSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **NullableString** |  | 
**Currency** | **string** |  | 
**Mode** | **string** |  | 
**Status** | [**CartSessionStatus**](CartSessionStatus.md) |  | 
**LineItems** | [**[]CartLineItem**](CartLineItem.md) |  | 
**AddOnsOffered** | **[]string** |  | 
**AddOnsSelected** | **[]string** |  | 
**DiscountCode** | **NullableString** |  | 
**SubtotalMinor** | **int32** |  | 
**TotalMinor** | **int32** |  | 
**Metadata** | **map[string]interface{}** |  | 
**ConvertedToCheckoutSessionId** | **NullableString** |  | 
**PublicToken** | **string** | Anon-access token. Use to build /c/[cart_session_id]?token&#x3D;... URLs for customer-facing cart checkout. Becomes useless once the cart leaves status&#x3D;&#39;open&#39;. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewCartSession

`func NewCartSession(id string, object string, merchantId string, customerId NullableString, currency string, mode string, status CartSessionStatus, lineItems []CartLineItem, addOnsOffered []string, addOnsSelected []string, discountCode NullableString, subtotalMinor int32, totalMinor int32, metadata map[string]interface{}, convertedToCheckoutSessionId NullableString, publicToken string, createdAt time.Time, updatedAt time.Time, expiresAt time.Time, ) *CartSession`

NewCartSession instantiates a new CartSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartSessionWithDefaults

`func NewCartSessionWithDefaults() *CartSession`

NewCartSessionWithDefaults instantiates a new CartSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartSession) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CartSession) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CartSession) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CartSession) SetObject(v string)`

SetObject sets Object field to given value.


### GetMerchantId

`func (o *CartSession) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CartSession) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CartSession) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *CartSession) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CartSession) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CartSession) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *CartSession) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CartSession) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCurrency

`func (o *CartSession) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartSession) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartSession) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMode

`func (o *CartSession) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CartSession) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CartSession) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStatus

`func (o *CartSession) GetStatus() CartSessionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CartSession) GetStatusOk() (*CartSessionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CartSession) SetStatus(v CartSessionStatus)`

SetStatus sets Status field to given value.


### GetLineItems

`func (o *CartSession) GetLineItems() []CartLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CartSession) GetLineItemsOk() (*[]CartLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CartSession) SetLineItems(v []CartLineItem)`

SetLineItems sets LineItems field to given value.


### GetAddOnsOffered

`func (o *CartSession) GetAddOnsOffered() []string`

GetAddOnsOffered returns the AddOnsOffered field if non-nil, zero value otherwise.

### GetAddOnsOfferedOk

`func (o *CartSession) GetAddOnsOfferedOk() (*[]string, bool)`

GetAddOnsOfferedOk returns a tuple with the AddOnsOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnsOffered

`func (o *CartSession) SetAddOnsOffered(v []string)`

SetAddOnsOffered sets AddOnsOffered field to given value.


### GetAddOnsSelected

`func (o *CartSession) GetAddOnsSelected() []string`

GetAddOnsSelected returns the AddOnsSelected field if non-nil, zero value otherwise.

### GetAddOnsSelectedOk

`func (o *CartSession) GetAddOnsSelectedOk() (*[]string, bool)`

GetAddOnsSelectedOk returns a tuple with the AddOnsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnsSelected

`func (o *CartSession) SetAddOnsSelected(v []string)`

SetAddOnsSelected sets AddOnsSelected field to given value.


### GetDiscountCode

`func (o *CartSession) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *CartSession) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *CartSession) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.


### SetDiscountCodeNil

`func (o *CartSession) SetDiscountCodeNil(b bool)`

 SetDiscountCodeNil sets the value for DiscountCode to be an explicit nil

### UnsetDiscountCode
`func (o *CartSession) UnsetDiscountCode()`

UnsetDiscountCode ensures that no value is present for DiscountCode, not even an explicit nil
### GetSubtotalMinor

`func (o *CartSession) GetSubtotalMinor() int32`

GetSubtotalMinor returns the SubtotalMinor field if non-nil, zero value otherwise.

### GetSubtotalMinorOk

`func (o *CartSession) GetSubtotalMinorOk() (*int32, bool)`

GetSubtotalMinorOk returns a tuple with the SubtotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotalMinor

`func (o *CartSession) SetSubtotalMinor(v int32)`

SetSubtotalMinor sets SubtotalMinor field to given value.


### GetTotalMinor

`func (o *CartSession) GetTotalMinor() int32`

GetTotalMinor returns the TotalMinor field if non-nil, zero value otherwise.

### GetTotalMinorOk

`func (o *CartSession) GetTotalMinorOk() (*int32, bool)`

GetTotalMinorOk returns a tuple with the TotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinor

`func (o *CartSession) SetTotalMinor(v int32)`

SetTotalMinor sets TotalMinor field to given value.


### GetMetadata

`func (o *CartSession) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CartSession) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CartSession) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetConvertedToCheckoutSessionId

`func (o *CartSession) GetConvertedToCheckoutSessionId() string`

GetConvertedToCheckoutSessionId returns the ConvertedToCheckoutSessionId field if non-nil, zero value otherwise.

### GetConvertedToCheckoutSessionIdOk

`func (o *CartSession) GetConvertedToCheckoutSessionIdOk() (*string, bool)`

GetConvertedToCheckoutSessionIdOk returns a tuple with the ConvertedToCheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedToCheckoutSessionId

`func (o *CartSession) SetConvertedToCheckoutSessionId(v string)`

SetConvertedToCheckoutSessionId sets ConvertedToCheckoutSessionId field to given value.


### SetConvertedToCheckoutSessionIdNil

`func (o *CartSession) SetConvertedToCheckoutSessionIdNil(b bool)`

 SetConvertedToCheckoutSessionIdNil sets the value for ConvertedToCheckoutSessionId to be an explicit nil

### UnsetConvertedToCheckoutSessionId
`func (o *CartSession) UnsetConvertedToCheckoutSessionId()`

UnsetConvertedToCheckoutSessionId ensures that no value is present for ConvertedToCheckoutSessionId, not even an explicit nil
### GetPublicToken

`func (o *CartSession) GetPublicToken() string`

GetPublicToken returns the PublicToken field if non-nil, zero value otherwise.

### GetPublicTokenOk

`func (o *CartSession) GetPublicTokenOk() (*string, bool)`

GetPublicTokenOk returns a tuple with the PublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicToken

`func (o *CartSession) SetPublicToken(v string)`

SetPublicToken sets PublicToken field to given value.


### GetCreatedAt

`func (o *CartSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CartSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CartSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CartSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CartSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CartSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetExpiresAt

`func (o *CartSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CartSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CartSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


