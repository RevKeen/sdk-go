# CheckoutSessionCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**PublicToken** | **string** |  | 
**AmountMinor** | Pointer to **float32** |  | [optional] 
**Currency** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 
**ExpiresAt** | **string** |  | 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 
**SelectedMethod** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCheckoutSessionCreateResponseData

`func NewCheckoutSessionCreateResponseData(id string, url string, publicToken string, currency string, expiresAt string, ) *CheckoutSessionCreateResponseData`

NewCheckoutSessionCreateResponseData instantiates a new CheckoutSessionCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionCreateResponseDataWithDefaults

`func NewCheckoutSessionCreateResponseDataWithDefaults() *CheckoutSessionCreateResponseData`

NewCheckoutSessionCreateResponseDataWithDefaults instantiates a new CheckoutSessionCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckoutSessionCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSessionCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSessionCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *CheckoutSessionCreateResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutSessionCreateResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutSessionCreateResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublicToken

`func (o *CheckoutSessionCreateResponseData) GetPublicToken() string`

GetPublicToken returns the PublicToken field if non-nil, zero value otherwise.

### GetPublicTokenOk

`func (o *CheckoutSessionCreateResponseData) GetPublicTokenOk() (*string, bool)`

GetPublicTokenOk returns a tuple with the PublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicToken

`func (o *CheckoutSessionCreateResponseData) SetPublicToken(v string)`

SetPublicToken sets PublicToken field to given value.


### GetAmountMinor

`func (o *CheckoutSessionCreateResponseData) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CheckoutSessionCreateResponseData) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CheckoutSessionCreateResponseData) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *CheckoutSessionCreateResponseData) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *CheckoutSessionCreateResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutSessionCreateResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutSessionCreateResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *CheckoutSessionCreateResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CheckoutSessionCreateResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CheckoutSessionCreateResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CheckoutSessionCreateResponseData) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CheckoutSessionCreateResponseData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutSessionCreateResponseData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutSessionCreateResponseData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetAllowedMethods

`func (o *CheckoutSessionCreateResponseData) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *CheckoutSessionCreateResponseData) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *CheckoutSessionCreateResponseData) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *CheckoutSessionCreateResponseData) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetSelectedMethod

`func (o *CheckoutSessionCreateResponseData) GetSelectedMethod() string`

GetSelectedMethod returns the SelectedMethod field if non-nil, zero value otherwise.

### GetSelectedMethodOk

`func (o *CheckoutSessionCreateResponseData) GetSelectedMethodOk() (*string, bool)`

GetSelectedMethodOk returns a tuple with the SelectedMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedMethod

`func (o *CheckoutSessionCreateResponseData) SetSelectedMethod(v string)`

SetSelectedMethod sets SelectedMethod field to given value.

### HasSelectedMethod

`func (o *CheckoutSessionCreateResponseData) HasSelectedMethod() bool`

HasSelectedMethod returns a boolean if a field has been set.

### SetSelectedMethodNil

`func (o *CheckoutSessionCreateResponseData) SetSelectedMethodNil(b bool)`

 SetSelectedMethodNil sets the value for SelectedMethod to be an explicit nil

### UnsetSelectedMethod
`func (o *CheckoutSessionCreateResponseData) UnsetSelectedMethod()`

UnsetSelectedMethod ensures that no value is present for SelectedMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


