# UsageBalanceMeter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeterId** | **string** | Meter ID | 
**MeterName** | **string** | Meter display name | 
**UnitName** | **NullableString** | Unit display name | 
**CurrentValue** | **float32** | Aggregated usage in current period | 
**IncludedQuantity** | **float32** | Units included in plan (0 if pure metered) | 
**UsedQuantity** | **float32** | Usage consumed (same as current_value) | 
**RemainingIncluded** | **float32** | Remaining included units (max 0) | 
**OverageQuantity** | **float32** | Units over allowance (max 0) | 
**EstimatedAmountMinor** | **float32** | Estimated charge for overage in minor units | 
**TotalCostMinor** | **NullableFloat32** | Sum of event cost annotations (null if no cost data) | 
**MarginMinor** | **NullableFloat32** | Estimated margin: revenue - cost (null if no cost data) | 
**Currency** | **string** | Currency code | 
**PeriodStart** | **string** | Current billing period start (ISO 8601) | 
**PeriodEnd** | **string** | Current billing period end (ISO 8601) | 

## Methods

### NewUsageBalanceMeter

`func NewUsageBalanceMeter(meterId string, meterName string, unitName NullableString, currentValue float32, includedQuantity float32, usedQuantity float32, remainingIncluded float32, overageQuantity float32, estimatedAmountMinor float32, totalCostMinor NullableFloat32, marginMinor NullableFloat32, currency string, periodStart string, periodEnd string, ) *UsageBalanceMeter`

NewUsageBalanceMeter instantiates a new UsageBalanceMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBalanceMeterWithDefaults

`func NewUsageBalanceMeterWithDefaults() *UsageBalanceMeter`

NewUsageBalanceMeterWithDefaults instantiates a new UsageBalanceMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeterId

`func (o *UsageBalanceMeter) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageBalanceMeter) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageBalanceMeter) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetMeterName

`func (o *UsageBalanceMeter) GetMeterName() string`

GetMeterName returns the MeterName field if non-nil, zero value otherwise.

### GetMeterNameOk

`func (o *UsageBalanceMeter) GetMeterNameOk() (*string, bool)`

GetMeterNameOk returns a tuple with the MeterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterName

`func (o *UsageBalanceMeter) SetMeterName(v string)`

SetMeterName sets MeterName field to given value.


### GetUnitName

`func (o *UsageBalanceMeter) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *UsageBalanceMeter) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *UsageBalanceMeter) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### SetUnitNameNil

`func (o *UsageBalanceMeter) SetUnitNameNil(b bool)`

 SetUnitNameNil sets the value for UnitName to be an explicit nil

### UnsetUnitName
`func (o *UsageBalanceMeter) UnsetUnitName()`

UnsetUnitName ensures that no value is present for UnitName, not even an explicit nil
### GetCurrentValue

`func (o *UsageBalanceMeter) GetCurrentValue() float32`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *UsageBalanceMeter) GetCurrentValueOk() (*float32, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *UsageBalanceMeter) SetCurrentValue(v float32)`

SetCurrentValue sets CurrentValue field to given value.


### GetIncludedQuantity

`func (o *UsageBalanceMeter) GetIncludedQuantity() float32`

GetIncludedQuantity returns the IncludedQuantity field if non-nil, zero value otherwise.

### GetIncludedQuantityOk

`func (o *UsageBalanceMeter) GetIncludedQuantityOk() (*float32, bool)`

GetIncludedQuantityOk returns a tuple with the IncludedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedQuantity

`func (o *UsageBalanceMeter) SetIncludedQuantity(v float32)`

SetIncludedQuantity sets IncludedQuantity field to given value.


### GetUsedQuantity

`func (o *UsageBalanceMeter) GetUsedQuantity() float32`

GetUsedQuantity returns the UsedQuantity field if non-nil, zero value otherwise.

### GetUsedQuantityOk

`func (o *UsageBalanceMeter) GetUsedQuantityOk() (*float32, bool)`

GetUsedQuantityOk returns a tuple with the UsedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedQuantity

`func (o *UsageBalanceMeter) SetUsedQuantity(v float32)`

SetUsedQuantity sets UsedQuantity field to given value.


### GetRemainingIncluded

`func (o *UsageBalanceMeter) GetRemainingIncluded() float32`

GetRemainingIncluded returns the RemainingIncluded field if non-nil, zero value otherwise.

### GetRemainingIncludedOk

`func (o *UsageBalanceMeter) GetRemainingIncludedOk() (*float32, bool)`

GetRemainingIncludedOk returns a tuple with the RemainingIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingIncluded

`func (o *UsageBalanceMeter) SetRemainingIncluded(v float32)`

SetRemainingIncluded sets RemainingIncluded field to given value.


### GetOverageQuantity

`func (o *UsageBalanceMeter) GetOverageQuantity() float32`

GetOverageQuantity returns the OverageQuantity field if non-nil, zero value otherwise.

### GetOverageQuantityOk

`func (o *UsageBalanceMeter) GetOverageQuantityOk() (*float32, bool)`

GetOverageQuantityOk returns a tuple with the OverageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageQuantity

`func (o *UsageBalanceMeter) SetOverageQuantity(v float32)`

SetOverageQuantity sets OverageQuantity field to given value.


### GetEstimatedAmountMinor

`func (o *UsageBalanceMeter) GetEstimatedAmountMinor() float32`

GetEstimatedAmountMinor returns the EstimatedAmountMinor field if non-nil, zero value otherwise.

### GetEstimatedAmountMinorOk

`func (o *UsageBalanceMeter) GetEstimatedAmountMinorOk() (*float32, bool)`

GetEstimatedAmountMinorOk returns a tuple with the EstimatedAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAmountMinor

`func (o *UsageBalanceMeter) SetEstimatedAmountMinor(v float32)`

SetEstimatedAmountMinor sets EstimatedAmountMinor field to given value.


### GetTotalCostMinor

`func (o *UsageBalanceMeter) GetTotalCostMinor() float32`

GetTotalCostMinor returns the TotalCostMinor field if non-nil, zero value otherwise.

### GetTotalCostMinorOk

`func (o *UsageBalanceMeter) GetTotalCostMinorOk() (*float32, bool)`

GetTotalCostMinorOk returns a tuple with the TotalCostMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostMinor

`func (o *UsageBalanceMeter) SetTotalCostMinor(v float32)`

SetTotalCostMinor sets TotalCostMinor field to given value.


### SetTotalCostMinorNil

`func (o *UsageBalanceMeter) SetTotalCostMinorNil(b bool)`

 SetTotalCostMinorNil sets the value for TotalCostMinor to be an explicit nil

### UnsetTotalCostMinor
`func (o *UsageBalanceMeter) UnsetTotalCostMinor()`

UnsetTotalCostMinor ensures that no value is present for TotalCostMinor, not even an explicit nil
### GetMarginMinor

`func (o *UsageBalanceMeter) GetMarginMinor() float32`

GetMarginMinor returns the MarginMinor field if non-nil, zero value otherwise.

### GetMarginMinorOk

`func (o *UsageBalanceMeter) GetMarginMinorOk() (*float32, bool)`

GetMarginMinorOk returns a tuple with the MarginMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginMinor

`func (o *UsageBalanceMeter) SetMarginMinor(v float32)`

SetMarginMinor sets MarginMinor field to given value.


### SetMarginMinorNil

`func (o *UsageBalanceMeter) SetMarginMinorNil(b bool)`

 SetMarginMinorNil sets the value for MarginMinor to be an explicit nil

### UnsetMarginMinor
`func (o *UsageBalanceMeter) UnsetMarginMinor()`

UnsetMarginMinor ensures that no value is present for MarginMinor, not even an explicit nil
### GetCurrency

`func (o *UsageBalanceMeter) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageBalanceMeter) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageBalanceMeter) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPeriodStart

`func (o *UsageBalanceMeter) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UsageBalanceMeter) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UsageBalanceMeter) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPeriodEnd

`func (o *UsageBalanceMeter) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UsageBalanceMeter) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UsageBalanceMeter) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


