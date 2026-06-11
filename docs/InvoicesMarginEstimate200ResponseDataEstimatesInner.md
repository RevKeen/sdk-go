# InvoicesMarginEstimate200ResponseDataEstimatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rail** | **string** |  | 
**CardScheme** | **NullableString** |  | 
**GrossMinor** | **int32** |  | 
**EstimatedFeeMinor** | **int32** |  | 
**NetAfterFeesMinor** | **int32** |  | 
**FeeBreakdown** | [**InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown**](InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown.md) |  | 
**PricingSource** | **string** |  | 
**MatrixVersion** | **string** |  | 

## Methods

### NewInvoicesMarginEstimate200ResponseDataEstimatesInner

`func NewInvoicesMarginEstimate200ResponseDataEstimatesInner(rail string, cardScheme NullableString, grossMinor int32, estimatedFeeMinor int32, netAfterFeesMinor int32, feeBreakdown InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown, pricingSource string, matrixVersion string, ) *InvoicesMarginEstimate200ResponseDataEstimatesInner`

NewInvoicesMarginEstimate200ResponseDataEstimatesInner instantiates a new InvoicesMarginEstimate200ResponseDataEstimatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesMarginEstimate200ResponseDataEstimatesInnerWithDefaults

`func NewInvoicesMarginEstimate200ResponseDataEstimatesInnerWithDefaults() *InvoicesMarginEstimate200ResponseDataEstimatesInner`

NewInvoicesMarginEstimate200ResponseDataEstimatesInnerWithDefaults instantiates a new InvoicesMarginEstimate200ResponseDataEstimatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRail

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetRail(v string)`

SetRail sets Rail field to given value.


### GetCardScheme

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.


### SetCardSchemeNil

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetCardSchemeNil(b bool)`

 SetCardSchemeNil sets the value for CardScheme to be an explicit nil

### UnsetCardScheme
`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) UnsetCardScheme()`

UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
### GetGrossMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetGrossMinor() int32`

GetGrossMinor returns the GrossMinor field if non-nil, zero value otherwise.

### GetGrossMinorOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetGrossMinorOk() (*int32, bool)`

GetGrossMinorOk returns a tuple with the GrossMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetGrossMinor(v int32)`

SetGrossMinor sets GrossMinor field to given value.


### GetEstimatedFeeMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetEstimatedFeeMinor() int32`

GetEstimatedFeeMinor returns the EstimatedFeeMinor field if non-nil, zero value otherwise.

### GetEstimatedFeeMinorOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetEstimatedFeeMinorOk() (*int32, bool)`

GetEstimatedFeeMinorOk returns a tuple with the EstimatedFeeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetEstimatedFeeMinor(v int32)`

SetEstimatedFeeMinor sets EstimatedFeeMinor field to given value.


### GetNetAfterFeesMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetNetAfterFeesMinor() int32`

GetNetAfterFeesMinor returns the NetAfterFeesMinor field if non-nil, zero value otherwise.

### GetNetAfterFeesMinorOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetNetAfterFeesMinorOk() (*int32, bool)`

GetNetAfterFeesMinorOk returns a tuple with the NetAfterFeesMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAfterFeesMinor

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetNetAfterFeesMinor(v int32)`

SetNetAfterFeesMinor sets NetAfterFeesMinor field to given value.


### GetFeeBreakdown

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetFeeBreakdown() InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown`

GetFeeBreakdown returns the FeeBreakdown field if non-nil, zero value otherwise.

### GetFeeBreakdownOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetFeeBreakdownOk() (*InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown, bool)`

GetFeeBreakdownOk returns a tuple with the FeeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeBreakdown

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetFeeBreakdown(v InvoicesMarginEstimate200ResponseDataEstimatesInnerFeeBreakdown)`

SetFeeBreakdown sets FeeBreakdown field to given value.


### GetPricingSource

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetPricingSource() string`

GetPricingSource returns the PricingSource field if non-nil, zero value otherwise.

### GetPricingSourceOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetPricingSourceOk() (*string, bool)`

GetPricingSourceOk returns a tuple with the PricingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingSource

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetPricingSource(v string)`

SetPricingSource sets PricingSource field to given value.


### GetMatrixVersion

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetMatrixVersion() string`

GetMatrixVersion returns the MatrixVersion field if non-nil, zero value otherwise.

### GetMatrixVersionOk

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) GetMatrixVersionOk() (*string, bool)`

GetMatrixVersionOk returns a tuple with the MatrixVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrixVersion

`func (o *InvoicesMarginEstimate200ResponseDataEstimatesInner) SetMatrixVersion(v string)`

SetMatrixVersion sets MatrixVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


