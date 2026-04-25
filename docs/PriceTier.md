# PriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpTo** | **NullableInt32** | Upper bound of this tier (1-based inclusive). &#x60;null&#x60; &#x3D; open-ended catch-all. Only the FINAL tier may set &#x60;up_to: null&#x60;. | 
**UnitAmountMinor** | **NullableInt32** | Per-unit charge in minor currency units (cents). May be null if the tier charges only a flat fee. | 
**FlatAmountMinor** | **NullableInt32** | Flat fee charged once when quantity enters this tier. May be null if the tier charges only a per-unit amount. | 

## Methods

### NewPriceTier

`func NewPriceTier(upTo NullableInt32, unitAmountMinor NullableInt32, flatAmountMinor NullableInt32, ) *PriceTier`

NewPriceTier instantiates a new PriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTierWithDefaults

`func NewPriceTierWithDefaults() *PriceTier`

NewPriceTierWithDefaults instantiates a new PriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpTo

`func (o *PriceTier) GetUpTo() int32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *PriceTier) GetUpToOk() (*int32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *PriceTier) SetUpTo(v int32)`

SetUpTo sets UpTo field to given value.


### SetUpToNil

`func (o *PriceTier) SetUpToNil(b bool)`

 SetUpToNil sets the value for UpTo to be an explicit nil

### UnsetUpTo
`func (o *PriceTier) UnsetUpTo()`

UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
### GetUnitAmountMinor

`func (o *PriceTier) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *PriceTier) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *PriceTier) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### SetUnitAmountMinorNil

`func (o *PriceTier) SetUnitAmountMinorNil(b bool)`

 SetUnitAmountMinorNil sets the value for UnitAmountMinor to be an explicit nil

### UnsetUnitAmountMinor
`func (o *PriceTier) UnsetUnitAmountMinor()`

UnsetUnitAmountMinor ensures that no value is present for UnitAmountMinor, not even an explicit nil
### GetFlatAmountMinor

`func (o *PriceTier) GetFlatAmountMinor() int32`

GetFlatAmountMinor returns the FlatAmountMinor field if non-nil, zero value otherwise.

### GetFlatAmountMinorOk

`func (o *PriceTier) GetFlatAmountMinorOk() (*int32, bool)`

GetFlatAmountMinorOk returns a tuple with the FlatAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountMinor

`func (o *PriceTier) SetFlatAmountMinor(v int32)`

SetFlatAmountMinor sets FlatAmountMinor field to given value.


### SetFlatAmountMinorNil

`func (o *PriceTier) SetFlatAmountMinorNil(b bool)`

 SetFlatAmountMinorNil sets the value for FlatAmountMinor to be an explicit nil

### UnsetFlatAmountMinor
`func (o *PriceTier) UnsetFlatAmountMinor()`

UnsetFlatAmountMinor ensures that no value is present for FlatAmountMinor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


