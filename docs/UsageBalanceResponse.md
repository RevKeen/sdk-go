# UsageBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Meters** | [**[]UsageBalanceMeter**](UsageBalanceMeter.md) |  | 
**TotalEstimatedAmountMinor** | **float32** | Total estimated charge across all meters | 
**TotalCostMinor** | **NullableFloat32** | Total cost across all meters (null if no cost data) | 
**Currency** | **string** | Currency code | 

## Methods

### NewUsageBalanceResponse

`func NewUsageBalanceResponse(object string, meters []UsageBalanceMeter, totalEstimatedAmountMinor float32, totalCostMinor NullableFloat32, currency string, ) *UsageBalanceResponse`

NewUsageBalanceResponse instantiates a new UsageBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBalanceResponseWithDefaults

`func NewUsageBalanceResponseWithDefaults() *UsageBalanceResponse`

NewUsageBalanceResponseWithDefaults instantiates a new UsageBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *UsageBalanceResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsageBalanceResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsageBalanceResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetMeters

`func (o *UsageBalanceResponse) GetMeters() []UsageBalanceMeter`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *UsageBalanceResponse) GetMetersOk() (*[]UsageBalanceMeter, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *UsageBalanceResponse) SetMeters(v []UsageBalanceMeter)`

SetMeters sets Meters field to given value.


### GetTotalEstimatedAmountMinor

`func (o *UsageBalanceResponse) GetTotalEstimatedAmountMinor() float32`

GetTotalEstimatedAmountMinor returns the TotalEstimatedAmountMinor field if non-nil, zero value otherwise.

### GetTotalEstimatedAmountMinorOk

`func (o *UsageBalanceResponse) GetTotalEstimatedAmountMinorOk() (*float32, bool)`

GetTotalEstimatedAmountMinorOk returns a tuple with the TotalEstimatedAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEstimatedAmountMinor

`func (o *UsageBalanceResponse) SetTotalEstimatedAmountMinor(v float32)`

SetTotalEstimatedAmountMinor sets TotalEstimatedAmountMinor field to given value.


### GetTotalCostMinor

`func (o *UsageBalanceResponse) GetTotalCostMinor() float32`

GetTotalCostMinor returns the TotalCostMinor field if non-nil, zero value otherwise.

### GetTotalCostMinorOk

`func (o *UsageBalanceResponse) GetTotalCostMinorOk() (*float32, bool)`

GetTotalCostMinorOk returns a tuple with the TotalCostMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostMinor

`func (o *UsageBalanceResponse) SetTotalCostMinor(v float32)`

SetTotalCostMinor sets TotalCostMinor field to given value.


### SetTotalCostMinorNil

`func (o *UsageBalanceResponse) SetTotalCostMinorNil(b bool)`

 SetTotalCostMinorNil sets the value for TotalCostMinor to be an explicit nil

### UnsetTotalCostMinor
`func (o *UsageBalanceResponse) UnsetTotalCostMinor()`

UnsetTotalCostMinor ensures that no value is present for TotalCostMinor, not even an explicit nil
### GetCurrency

`func (o *UsageBalanceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UsageBalanceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UsageBalanceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


