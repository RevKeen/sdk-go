# ListPaymentLinks200ResponseDataInnerLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  | 
**PriceId** | Pointer to **NullableString** |  | [optional] 
**Quantity** | **int32** |  | 
**MerchantReference** | Pointer to **NullableString** |  | [optional] 
**IsAdhoc** | **bool** |  | 
**ProductDetails** | Pointer to [**ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetails**](ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetails.md) |  | [optional] 

## Methods

### NewListPaymentLinks200ResponseDataInnerLineItemsInner

`func NewListPaymentLinks200ResponseDataInnerLineItemsInner(productId string, quantity int32, isAdhoc bool, ) *ListPaymentLinks200ResponseDataInnerLineItemsInner`

NewListPaymentLinks200ResponseDataInnerLineItemsInner instantiates a new ListPaymentLinks200ResponseDataInnerLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentLinks200ResponseDataInnerLineItemsInnerWithDefaults

`func NewListPaymentLinks200ResponseDataInnerLineItemsInnerWithDefaults() *ListPaymentLinks200ResponseDataInnerLineItemsInner`

NewListPaymentLinks200ResponseDataInnerLineItemsInnerWithDefaults instantiates a new ListPaymentLinks200ResponseDataInnerLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetPriceId

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### SetPriceIdNil

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetQuantity

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetMerchantReference

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetMerchantReference() string`

GetMerchantReference returns the MerchantReference field if non-nil, zero value otherwise.

### GetMerchantReferenceOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetMerchantReferenceOk() (*string, bool)`

GetMerchantReferenceOk returns a tuple with the MerchantReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantReference

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetMerchantReference(v string)`

SetMerchantReference sets MerchantReference field to given value.

### HasMerchantReference

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) HasMerchantReference() bool`

HasMerchantReference returns a boolean if a field has been set.

### SetMerchantReferenceNil

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetMerchantReferenceNil(b bool)`

 SetMerchantReferenceNil sets the value for MerchantReference to be an explicit nil

### UnsetMerchantReference
`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) UnsetMerchantReference()`

UnsetMerchantReference ensures that no value is present for MerchantReference, not even an explicit nil
### GetIsAdhoc

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetIsAdhoc() bool`

GetIsAdhoc returns the IsAdhoc field if non-nil, zero value otherwise.

### GetIsAdhocOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetIsAdhocOk() (*bool, bool)`

GetIsAdhocOk returns a tuple with the IsAdhoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdhoc

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetIsAdhoc(v bool)`

SetIsAdhoc sets IsAdhoc field to given value.


### GetProductDetails

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetProductDetails() ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetails`

GetProductDetails returns the ProductDetails field if non-nil, zero value otherwise.

### GetProductDetailsOk

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) GetProductDetailsOk() (*ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetails, bool)`

GetProductDetailsOk returns a tuple with the ProductDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDetails

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) SetProductDetails(v ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetails)`

SetProductDetails sets ProductDetails field to given value.

### HasProductDetails

`func (o *ListPaymentLinks200ResponseDataInnerLineItemsInner) HasProductDetails() bool`

HasProductDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


