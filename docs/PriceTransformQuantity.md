# PriceTransformQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DivideBy** | **int32** | Package size — quantity is divided by this value before pricing. Must be &gt;&#x3D; 1. Example: &#x60;divide_by: 10&#x60; with &#x60;round: up&#x60; charges per pack of 10. | 
**Round** | **string** | Rounding mode when quantity is not an exact multiple of &#x60;divide_by&#x60;. | 

## Methods

### NewPriceTransformQuantity

`func NewPriceTransformQuantity(divideBy int32, round string, ) *PriceTransformQuantity`

NewPriceTransformQuantity instantiates a new PriceTransformQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTransformQuantityWithDefaults

`func NewPriceTransformQuantityWithDefaults() *PriceTransformQuantity`

NewPriceTransformQuantityWithDefaults instantiates a new PriceTransformQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDivideBy

`func (o *PriceTransformQuantity) GetDivideBy() int32`

GetDivideBy returns the DivideBy field if non-nil, zero value otherwise.

### GetDivideByOk

`func (o *PriceTransformQuantity) GetDivideByOk() (*int32, bool)`

GetDivideByOk returns a tuple with the DivideBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideBy

`func (o *PriceTransformQuantity) SetDivideBy(v int32)`

SetDivideBy sets DivideBy field to given value.


### GetRound

`func (o *PriceTransformQuantity) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *PriceTransformQuantity) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *PriceTransformQuantity) SetRound(v string)`

SetRound sets Round field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


