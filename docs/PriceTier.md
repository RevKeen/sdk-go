# PriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpTo** | **NullableInt32** | Upper bound of this tier (1-based inclusive). &#x60;null&#x60; &#x3D; open-ended catch-all. Only the FINAL tier may set &#x60;up_to: null&#x60;. | 
**UnitAmountMinor** | **NullableInt32** | Per-unit charge in minor currency units (cents). May be null if the tier charges only a flat fee, or (REV-6249) when the rate is sub-penny and carried by &#x60;unit_amount_decimal&#x60;. | 
**UnitAmountDecimal** | Pointer to **NullableString** | Exact per-unit rate for this tier in minor currency units, as a decimal string (sub-penny safe, e.g. \&quot;0.0000024\&quot;). Takes precedence over &#x60;unit_amount_minor&#x60;, which is null whenever the rate is not a whole number of minor units. Must be &gt; 0 with at most 12 decimal places. Tier BOUNDARIES remain integer. | [optional] 
**FlatAmountMinor** | **NullableInt32** | Flat fee charged once when quantity enters this tier. May be null if the tier charges only a per-unit amount. Always a whole number of minor units — only the per-unit RATE supports sub-penny precision. | 

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
### GetUnitAmountDecimal

`func (o *PriceTier) GetUnitAmountDecimal() string`

GetUnitAmountDecimal returns the UnitAmountDecimal field if non-nil, zero value otherwise.

### GetUnitAmountDecimalOk

`func (o *PriceTier) GetUnitAmountDecimalOk() (*string, bool)`

GetUnitAmountDecimalOk returns a tuple with the UnitAmountDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountDecimal

`func (o *PriceTier) SetUnitAmountDecimal(v string)`

SetUnitAmountDecimal sets UnitAmountDecimal field to given value.

### HasUnitAmountDecimal

`func (o *PriceTier) HasUnitAmountDecimal() bool`

HasUnitAmountDecimal returns a boolean if a field has been set.

### SetUnitAmountDecimalNil

`func (o *PriceTier) SetUnitAmountDecimalNil(b bool)`

 SetUnitAmountDecimalNil sets the value for UnitAmountDecimal to be an explicit nil

### UnsetUnitAmountDecimal
`func (o *PriceTier) UnsetUnitAmountDecimal()`

UnsetUnitAmountDecimal ensures that no value is present for UnitAmountDecimal, not even an explicit nil
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


