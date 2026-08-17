# MeterPriceTierInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpTo** | Pointer to **NullableInt32** | Upper bound of this tier (1-based inclusive). null &#x3D; open-ended final tier. New clients should use this canonical shape; legacy first_unit/last_unit is accepted during the deprecation window. | [optional] 
**UnitAmountMinor** | **int32** | Per-unit price in this tier (minor units) | 
**FlatAmountMinor** | Pointer to **int32** | Flat fee for entering this tier (minor units) | [optional] 

## Methods

### NewMeterPriceTierInput

`func NewMeterPriceTierInput(unitAmountMinor int32, ) *MeterPriceTierInput`

NewMeterPriceTierInput instantiates a new MeterPriceTierInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterPriceTierInputWithDefaults

`func NewMeterPriceTierInputWithDefaults() *MeterPriceTierInput`

NewMeterPriceTierInputWithDefaults instantiates a new MeterPriceTierInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpTo

`func (o *MeterPriceTierInput) GetUpTo() int32`

GetUpTo returns the UpTo field if non-nil, zero value otherwise.

### GetUpToOk

`func (o *MeterPriceTierInput) GetUpToOk() (*int32, bool)`

GetUpToOk returns a tuple with the UpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTo

`func (o *MeterPriceTierInput) SetUpTo(v int32)`

SetUpTo sets UpTo field to given value.

### HasUpTo

`func (o *MeterPriceTierInput) HasUpTo() bool`

HasUpTo returns a boolean if a field has been set.

### SetUpToNil

`func (o *MeterPriceTierInput) SetUpToNil(b bool)`

 SetUpToNil sets the value for UpTo to be an explicit nil

### UnsetUpTo
`func (o *MeterPriceTierInput) UnsetUpTo()`

UnsetUpTo ensures that no value is present for UpTo, not even an explicit nil
### GetUnitAmountMinor

`func (o *MeterPriceTierInput) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *MeterPriceTierInput) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *MeterPriceTierInput) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetFlatAmountMinor

`func (o *MeterPriceTierInput) GetFlatAmountMinor() int32`

GetFlatAmountMinor returns the FlatAmountMinor field if non-nil, zero value otherwise.

### GetFlatAmountMinorOk

`func (o *MeterPriceTierInput) GetFlatAmountMinorOk() (*int32, bool)`

GetFlatAmountMinorOk returns a tuple with the FlatAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountMinor

`func (o *MeterPriceTierInput) SetFlatAmountMinor(v int32)`

SetFlatAmountMinor sets FlatAmountMinor field to given value.

### HasFlatAmountMinor

`func (o *MeterPriceTierInput) HasFlatAmountMinor() bool`

HasFlatAmountMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


