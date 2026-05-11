# CreateCheckoutSessionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**AmountMinor** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**SuccessUrl** | Pointer to **string** |  | [optional] 
**CancelUrl** | Pointer to **string** |  | [optional] 
**AllowedMethods** | Pointer to **[]string** | Payment methods to offer. Intersected with merchant capabilities. Defaults to merchant config. | [optional] 
**CompanionDeviceId** | Pointer to **string** | Target a registered companion device. Session is pushed via SSE to the device. | [optional] 

## Methods

### NewCreateCheckoutSessionInput

`func NewCreateCheckoutSessionInput() *CreateCheckoutSessionInput`

NewCreateCheckoutSessionInput instantiates a new CreateCheckoutSessionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCheckoutSessionInputWithDefaults

`func NewCreateCheckoutSessionInputWithDefaults() *CreateCheckoutSessionInput`

NewCreateCheckoutSessionInputWithDefaults instantiates a new CreateCheckoutSessionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreateCheckoutSessionInput) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateCheckoutSessionInput) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateCheckoutSessionInput) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreateCheckoutSessionInput) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetProductId

`func (o *CreateCheckoutSessionInput) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateCheckoutSessionInput) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateCheckoutSessionInput) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreateCheckoutSessionInput) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetAmountMinor

`func (o *CreateCheckoutSessionInput) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateCheckoutSessionInput) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateCheckoutSessionInput) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *CreateCheckoutSessionInput) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateCheckoutSessionInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCheckoutSessionInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCheckoutSessionInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCheckoutSessionInput) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CreateCheckoutSessionInput) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreateCheckoutSessionInput) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreateCheckoutSessionInput) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CreateCheckoutSessionInput) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *CreateCheckoutSessionInput) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreateCheckoutSessionInput) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreateCheckoutSessionInput) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *CreateCheckoutSessionInput) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *CreateCheckoutSessionInput) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *CreateCheckoutSessionInput) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *CreateCheckoutSessionInput) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *CreateCheckoutSessionInput) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### GetCompanionDeviceId

`func (o *CreateCheckoutSessionInput) GetCompanionDeviceId() string`

GetCompanionDeviceId returns the CompanionDeviceId field if non-nil, zero value otherwise.

### GetCompanionDeviceIdOk

`func (o *CreateCheckoutSessionInput) GetCompanionDeviceIdOk() (*string, bool)`

GetCompanionDeviceIdOk returns a tuple with the CompanionDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanionDeviceId

`func (o *CreateCheckoutSessionInput) SetCompanionDeviceId(v string)`

SetCompanionDeviceId sets CompanionDeviceId field to given value.

### HasCompanionDeviceId

`func (o *CreateCheckoutSessionInput) HasCompanionDeviceId() bool`

HasCompanionDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


