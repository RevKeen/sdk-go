# PhaseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceId** | **string** | Price ID for this item | 
**Quantity** | Pointer to **int32** | Quantity of this item | [optional] [default to 1]

## Methods

### NewPhaseItem

`func NewPhaseItem(priceId string, ) *PhaseItem`

NewPhaseItem instantiates a new PhaseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseItemWithDefaults

`func NewPhaseItemWithDefaults() *PhaseItem`

NewPhaseItemWithDefaults instantiates a new PhaseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceId

`func (o *PhaseItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *PhaseItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *PhaseItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetQuantity

`func (o *PhaseItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PhaseItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PhaseItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PhaseItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


