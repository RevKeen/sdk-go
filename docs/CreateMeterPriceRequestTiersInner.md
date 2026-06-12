# CreateMeterPriceRequestTiersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstUnit** | **int32** | Start of tier range | 
**LastUnit** | Pointer to **NullableInt32** | End of tier range (null &#x3D; unbounded) | [optional] 
**UnitAmountMinor** | **int32** | Per-unit price in this tier (minor units) | 
**FlatAmountMinor** | Pointer to **int32** | Flat fee for entering this tier (minor units) | [optional] 

## Methods

### NewCreateMeterPriceRequestTiersInner

`func NewCreateMeterPriceRequestTiersInner(firstUnit int32, unitAmountMinor int32, ) *CreateMeterPriceRequestTiersInner`

NewCreateMeterPriceRequestTiersInner instantiates a new CreateMeterPriceRequestTiersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMeterPriceRequestTiersInnerWithDefaults

`func NewCreateMeterPriceRequestTiersInnerWithDefaults() *CreateMeterPriceRequestTiersInner`

NewCreateMeterPriceRequestTiersInnerWithDefaults instantiates a new CreateMeterPriceRequestTiersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstUnit

`func (o *CreateMeterPriceRequestTiersInner) GetFirstUnit() int32`

GetFirstUnit returns the FirstUnit field if non-nil, zero value otherwise.

### GetFirstUnitOk

`func (o *CreateMeterPriceRequestTiersInner) GetFirstUnitOk() (*int32, bool)`

GetFirstUnitOk returns a tuple with the FirstUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUnit

`func (o *CreateMeterPriceRequestTiersInner) SetFirstUnit(v int32)`

SetFirstUnit sets FirstUnit field to given value.


### GetLastUnit

`func (o *CreateMeterPriceRequestTiersInner) GetLastUnit() int32`

GetLastUnit returns the LastUnit field if non-nil, zero value otherwise.

### GetLastUnitOk

`func (o *CreateMeterPriceRequestTiersInner) GetLastUnitOk() (*int32, bool)`

GetLastUnitOk returns a tuple with the LastUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUnit

`func (o *CreateMeterPriceRequestTiersInner) SetLastUnit(v int32)`

SetLastUnit sets LastUnit field to given value.

### HasLastUnit

`func (o *CreateMeterPriceRequestTiersInner) HasLastUnit() bool`

HasLastUnit returns a boolean if a field has been set.

### SetLastUnitNil

`func (o *CreateMeterPriceRequestTiersInner) SetLastUnitNil(b bool)`

 SetLastUnitNil sets the value for LastUnit to be an explicit nil

### UnsetLastUnit
`func (o *CreateMeterPriceRequestTiersInner) UnsetLastUnit()`

UnsetLastUnit ensures that no value is present for LastUnit, not even an explicit nil
### GetUnitAmountMinor

`func (o *CreateMeterPriceRequestTiersInner) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *CreateMeterPriceRequestTiersInner) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *CreateMeterPriceRequestTiersInner) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetFlatAmountMinor

`func (o *CreateMeterPriceRequestTiersInner) GetFlatAmountMinor() int32`

GetFlatAmountMinor returns the FlatAmountMinor field if non-nil, zero value otherwise.

### GetFlatAmountMinorOk

`func (o *CreateMeterPriceRequestTiersInner) GetFlatAmountMinorOk() (*int32, bool)`

GetFlatAmountMinorOk returns a tuple with the FlatAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatAmountMinor

`func (o *CreateMeterPriceRequestTiersInner) SetFlatAmountMinor(v int32)`

SetFlatAmountMinor sets FlatAmountMinor field to given value.

### HasFlatAmountMinor

`func (o *CreateMeterPriceRequestTiersInner) HasFlatAmountMinor() bool`

HasFlatAmountMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


