# CartCheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **NullableString** |  | 
**SessionToken** | **NullableString** |  | 
**Status** | **string** |  | 
**Mode** | **NullableString** |  | 
**AmountMinor** | **NullableInt32** |  | 
**Currency** | **NullableString** |  | 
**LineItems** | [**[]CartLineItem**](CartLineItem.md) |  | 
**Metadata** | **map[string]interface{}** |  | 
**ExpiresAt** | **NullableTime** |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 

## Methods

### NewCartCheckoutSession

`func NewCartCheckoutSession(id string, object string, merchantId string, customerId NullableString, sessionToken NullableString, status string, mode NullableString, amountMinor NullableInt32, currency NullableString, lineItems []CartLineItem, metadata map[string]interface{}, expiresAt NullableTime, createdAt NullableTime, updatedAt NullableTime, ) *CartCheckoutSession`

NewCartCheckoutSession instantiates a new CartCheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartCheckoutSessionWithDefaults

`func NewCartCheckoutSessionWithDefaults() *CartCheckoutSession`

NewCartCheckoutSessionWithDefaults instantiates a new CartCheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartCheckoutSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartCheckoutSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartCheckoutSession) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CartCheckoutSession) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CartCheckoutSession) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CartCheckoutSession) SetObject(v string)`

SetObject sets Object field to given value.


### GetMerchantId

`func (o *CartCheckoutSession) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CartCheckoutSession) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CartCheckoutSession) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *CartCheckoutSession) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CartCheckoutSession) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CartCheckoutSession) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *CartCheckoutSession) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CartCheckoutSession) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetSessionToken

`func (o *CartCheckoutSession) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *CartCheckoutSession) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *CartCheckoutSession) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.


### SetSessionTokenNil

`func (o *CartCheckoutSession) SetSessionTokenNil(b bool)`

 SetSessionTokenNil sets the value for SessionToken to be an explicit nil

### UnsetSessionToken
`func (o *CartCheckoutSession) UnsetSessionToken()`

UnsetSessionToken ensures that no value is present for SessionToken, not even an explicit nil
### GetStatus

`func (o *CartCheckoutSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CartCheckoutSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CartCheckoutSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMode

`func (o *CartCheckoutSession) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CartCheckoutSession) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CartCheckoutSession) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *CartCheckoutSession) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *CartCheckoutSession) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetAmountMinor

`func (o *CartCheckoutSession) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CartCheckoutSession) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CartCheckoutSession) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *CartCheckoutSession) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *CartCheckoutSession) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *CartCheckoutSession) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartCheckoutSession) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartCheckoutSession) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CartCheckoutSession) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CartCheckoutSession) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetLineItems

`func (o *CartCheckoutSession) GetLineItems() []CartLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CartCheckoutSession) GetLineItemsOk() (*[]CartLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CartCheckoutSession) SetLineItems(v []CartLineItem)`

SetLineItems sets LineItems field to given value.


### GetMetadata

`func (o *CartCheckoutSession) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CartCheckoutSession) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CartCheckoutSession) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetExpiresAt

`func (o *CartCheckoutSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CartCheckoutSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CartCheckoutSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CartCheckoutSession) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CartCheckoutSession) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *CartCheckoutSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CartCheckoutSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CartCheckoutSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CartCheckoutSession) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CartCheckoutSession) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CartCheckoutSession) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CartCheckoutSession) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CartCheckoutSession) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *CartCheckoutSession) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CartCheckoutSession) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


