# UsageBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Meters** | [**[]UsageBalanceMeter**](UsageBalanceMeter.md) |  | 
**TotalEstimatedAmountMinor** | **float32** | Total estimated charge across all meters | 
**TotalCostMinor** | **NullableFloat32** | Total cost across all meters (null if no cost data) | 
**TotalPaymentProcessingCostMinor** | **NullableFloat32** | Payment-processing cost from transaction_costs for matching paid transactions. Null when no matching settled/paid cost evidence is available. | 
**TotalPlatformFeeMinor** | **NullableFloat32** | RevKeen/platform fee from transaction_costs for matching paid transactions. Null when no matching settled/paid cost evidence is available. | 
**TrueNetMarginMinor** | **NullableFloat32** | Total estimated charge minus AI provider COGS, payment-processing cost, and platform fee. Null until all cost inputs are available. | 
**TrueNetMarginPercent** | **NullableFloat32** | true_net_margin_minor as a percentage of total_estimated_amount_minor. Null until true-net margin is available. | 
**Currency** | **string** | Currency code | 

## Methods

### NewUsageBalanceResponse

`func NewUsageBalanceResponse(object string, meters []UsageBalanceMeter, totalEstimatedAmountMinor float32, totalCostMinor NullableFloat32, totalPaymentProcessingCostMinor NullableFloat32, totalPlatformFeeMinor NullableFloat32, trueNetMarginMinor NullableFloat32, trueNetMarginPercent NullableFloat32, currency string, ) *UsageBalanceResponse`

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
### GetTotalPaymentProcessingCostMinor

`func (o *UsageBalanceResponse) GetTotalPaymentProcessingCostMinor() float32`

GetTotalPaymentProcessingCostMinor returns the TotalPaymentProcessingCostMinor field if non-nil, zero value otherwise.

### GetTotalPaymentProcessingCostMinorOk

`func (o *UsageBalanceResponse) GetTotalPaymentProcessingCostMinorOk() (*float32, bool)`

GetTotalPaymentProcessingCostMinorOk returns a tuple with the TotalPaymentProcessingCostMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaymentProcessingCostMinor

`func (o *UsageBalanceResponse) SetTotalPaymentProcessingCostMinor(v float32)`

SetTotalPaymentProcessingCostMinor sets TotalPaymentProcessingCostMinor field to given value.


### SetTotalPaymentProcessingCostMinorNil

`func (o *UsageBalanceResponse) SetTotalPaymentProcessingCostMinorNil(b bool)`

 SetTotalPaymentProcessingCostMinorNil sets the value for TotalPaymentProcessingCostMinor to be an explicit nil

### UnsetTotalPaymentProcessingCostMinor
`func (o *UsageBalanceResponse) UnsetTotalPaymentProcessingCostMinor()`

UnsetTotalPaymentProcessingCostMinor ensures that no value is present for TotalPaymentProcessingCostMinor, not even an explicit nil
### GetTotalPlatformFeeMinor

`func (o *UsageBalanceResponse) GetTotalPlatformFeeMinor() float32`

GetTotalPlatformFeeMinor returns the TotalPlatformFeeMinor field if non-nil, zero value otherwise.

### GetTotalPlatformFeeMinorOk

`func (o *UsageBalanceResponse) GetTotalPlatformFeeMinorOk() (*float32, bool)`

GetTotalPlatformFeeMinorOk returns a tuple with the TotalPlatformFeeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPlatformFeeMinor

`func (o *UsageBalanceResponse) SetTotalPlatformFeeMinor(v float32)`

SetTotalPlatformFeeMinor sets TotalPlatformFeeMinor field to given value.


### SetTotalPlatformFeeMinorNil

`func (o *UsageBalanceResponse) SetTotalPlatformFeeMinorNil(b bool)`

 SetTotalPlatformFeeMinorNil sets the value for TotalPlatformFeeMinor to be an explicit nil

### UnsetTotalPlatformFeeMinor
`func (o *UsageBalanceResponse) UnsetTotalPlatformFeeMinor()`

UnsetTotalPlatformFeeMinor ensures that no value is present for TotalPlatformFeeMinor, not even an explicit nil
### GetTrueNetMarginMinor

`func (o *UsageBalanceResponse) GetTrueNetMarginMinor() float32`

GetTrueNetMarginMinor returns the TrueNetMarginMinor field if non-nil, zero value otherwise.

### GetTrueNetMarginMinorOk

`func (o *UsageBalanceResponse) GetTrueNetMarginMinorOk() (*float32, bool)`

GetTrueNetMarginMinorOk returns a tuple with the TrueNetMarginMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueNetMarginMinor

`func (o *UsageBalanceResponse) SetTrueNetMarginMinor(v float32)`

SetTrueNetMarginMinor sets TrueNetMarginMinor field to given value.


### SetTrueNetMarginMinorNil

`func (o *UsageBalanceResponse) SetTrueNetMarginMinorNil(b bool)`

 SetTrueNetMarginMinorNil sets the value for TrueNetMarginMinor to be an explicit nil

### UnsetTrueNetMarginMinor
`func (o *UsageBalanceResponse) UnsetTrueNetMarginMinor()`

UnsetTrueNetMarginMinor ensures that no value is present for TrueNetMarginMinor, not even an explicit nil
### GetTrueNetMarginPercent

`func (o *UsageBalanceResponse) GetTrueNetMarginPercent() float32`

GetTrueNetMarginPercent returns the TrueNetMarginPercent field if non-nil, zero value otherwise.

### GetTrueNetMarginPercentOk

`func (o *UsageBalanceResponse) GetTrueNetMarginPercentOk() (*float32, bool)`

GetTrueNetMarginPercentOk returns a tuple with the TrueNetMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueNetMarginPercent

`func (o *UsageBalanceResponse) SetTrueNetMarginPercent(v float32)`

SetTrueNetMarginPercent sets TrueNetMarginPercent field to given value.


### SetTrueNetMarginPercentNil

`func (o *UsageBalanceResponse) SetTrueNetMarginPercentNil(b bool)`

 SetTrueNetMarginPercentNil sets the value for TrueNetMarginPercent to be an explicit nil

### UnsetTrueNetMarginPercent
`func (o *UsageBalanceResponse) UnsetTrueNetMarginPercent()`

UnsetTrueNetMarginPercent ensures that no value is present for TrueNetMarginPercent, not even an explicit nil
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


