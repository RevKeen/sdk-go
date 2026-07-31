# PaymentLinksCreateRequestLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **string** |  | [optional] 
**PriceData** | Pointer to [**PaymentLinksCreateRequestLineItemsInnerPriceData**](PaymentLinksCreateRequestLineItemsInnerPriceData.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**MerchantReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentLinksCreateRequestLineItemsInner

`func NewPaymentLinksCreateRequestLineItemsInner() *PaymentLinksCreateRequestLineItemsInner`

NewPaymentLinksCreateRequestLineItemsInner instantiates a new PaymentLinksCreateRequestLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinksCreateRequestLineItemsInnerWithDefaults

`func NewPaymentLinksCreateRequestLineItemsInnerWithDefaults() *PaymentLinksCreateRequestLineItemsInner`

NewPaymentLinksCreateRequestLineItemsInnerWithDefaults instantiates a new PaymentLinksCreateRequestLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PaymentLinksCreateRequestLineItemsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PaymentLinksCreateRequestLineItemsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PaymentLinksCreateRequestLineItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceData

`func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceData() PaymentLinksCreateRequestLineItemsInnerPriceData`

GetPriceData returns the PriceData field if non-nil, zero value otherwise.

### GetPriceDataOk

`func (o *PaymentLinksCreateRequestLineItemsInner) GetPriceDataOk() (*PaymentLinksCreateRequestLineItemsInnerPriceData, bool)`

GetPriceDataOk returns a tuple with the PriceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceData

`func (o *PaymentLinksCreateRequestLineItemsInner) SetPriceData(v PaymentLinksCreateRequestLineItemsInnerPriceData)`

SetPriceData sets PriceData field to given value.

### HasPriceData

`func (o *PaymentLinksCreateRequestLineItemsInner) HasPriceData() bool`

HasPriceData returns a boolean if a field has been set.

### GetQuantity

`func (o *PaymentLinksCreateRequestLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PaymentLinksCreateRequestLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PaymentLinksCreateRequestLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PaymentLinksCreateRequestLineItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMerchantReference

`func (o *PaymentLinksCreateRequestLineItemsInner) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *PaymentLinksCreateRequestLineItemsInner) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *PaymentLinksCreateRequestLineItemsInner) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *PaymentLinksCreateRequestLineItemsInner) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### SetMerchantReferenceNil

`func (o *PaymentLinksCreateRequestLineItemsInner) SetMerchantReferenceNil(b bool)`

 SetMerchantReferenceNil sets the value for MerchantReference to be an explicit nil

### UnsetMerchantReference
`func (o *PaymentLinksCreateRequestLineItemsInner) UnsetMerchantReference()`

UnsetMerchantReference ensures that no value is present for MerchantReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


