# CreatePaymentLinkRequestLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **string** |  | [optional] 
**PriceData** | Pointer to [**CreatePaymentLinkRequestLineItemsInnerPriceData**](CreatePaymentLinkRequestLineItemsInnerPriceData.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**MerchantReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreatePaymentLinkRequestLineItemsInner

`func NewCreatePaymentLinkRequestLineItemsInner() *CreatePaymentLinkRequestLineItemsInner`

NewCreatePaymentLinkRequestLineItemsInner instantiates a new CreatePaymentLinkRequestLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentLinkRequestLineItemsInnerWithDefaults

`func NewCreatePaymentLinkRequestLineItemsInnerWithDefaults() *CreatePaymentLinkRequestLineItemsInner`

NewCreatePaymentLinkRequestLineItemsInnerWithDefaults instantiates a new CreatePaymentLinkRequestLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *CreatePaymentLinkRequestLineItemsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreatePaymentLinkRequestLineItemsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreatePaymentLinkRequestLineItemsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreatePaymentLinkRequestLineItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceData

`func (o *CreatePaymentLinkRequestLineItemsInner) GetPriceData() CreatePaymentLinkRequestLineItemsInnerPriceData`

GetPriceData returns the PriceData field if non-nil, zero value otherwise.

### GetPriceDataOk

`func (o *CreatePaymentLinkRequestLineItemsInner) GetPriceDataOk() (*CreatePaymentLinkRequestLineItemsInnerPriceData, bool)`

GetPriceDataOk returns a tuple with the PriceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceData

`func (o *CreatePaymentLinkRequestLineItemsInner) SetPriceData(v CreatePaymentLinkRequestLineItemsInnerPriceData)`

SetPriceData sets PriceData field to given value.

### HasPriceData

`func (o *CreatePaymentLinkRequestLineItemsInner) HasPriceData() bool`

HasPriceData returns a boolean if a field has been set.

### GetQuantity

`func (o *CreatePaymentLinkRequestLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreatePaymentLinkRequestLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreatePaymentLinkRequestLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreatePaymentLinkRequestLineItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMerchantReference

`func (o *CreatePaymentLinkRequestLineItemsInner) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *CreatePaymentLinkRequestLineItemsInner) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *CreatePaymentLinkRequestLineItemsInner) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *CreatePaymentLinkRequestLineItemsInner) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### SetMerchantReferenceNil

`func (o *CreatePaymentLinkRequestLineItemsInner) SetMerchantReferenceNil(b bool)`

 SetMerchantReferenceNil sets the value for MerchantReference to be an explicit nil

### UnsetMerchantReference
`func (o *CreatePaymentLinkRequestLineItemsInner) UnsetMerchantReference()`

UnsetMerchantReference ensures that no value is present for MerchantReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


