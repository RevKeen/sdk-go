# CreateMeterPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingModel** | **string** | Pricing model | 
**Currency** | **string** | ISO 4217 currency code | 
**UnitAmountMinor** | Pointer to **int32** | Price per unit in minor units (for per_unit) | [optional] 
**FlatFeeMinor** | Pointer to **int32** | Base charge per period in minor units | [optional] 
**PackageSize** | Pointer to **int32** | Units per package (for package model) | [optional] 
**CostPerUnitMinor** | Pointer to **int32** | Merchant&#39;s cost per unit for margin intelligence | [optional] 
**Tiers** | Pointer to [**[]CreateMeterPriceRequestTiersInner**](CreateMeterPriceRequestTiersInner.md) | Tier configuration (for graduated and volume) | [optional] 

## Methods

### NewCreateMeterPriceRequest

`func NewCreateMeterPriceRequest(pricingModel string, currency string, ) *CreateMeterPriceRequest`

NewCreateMeterPriceRequest instantiates a new CreateMeterPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMeterPriceRequestWithDefaults

`func NewCreateMeterPriceRequestWithDefaults() *CreateMeterPriceRequest`

NewCreateMeterPriceRequestWithDefaults instantiates a new CreateMeterPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingModel

`func (o *CreateMeterPriceRequest) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreateMeterPriceRequest) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreateMeterPriceRequest) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetCurrency

`func (o *CreateMeterPriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateMeterPriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateMeterPriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitAmountMinor

`func (o *CreateMeterPriceRequest) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *CreateMeterPriceRequest) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *CreateMeterPriceRequest) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.

### HasUnitAmountMinor

`func (o *CreateMeterPriceRequest) HasUnitAmountMinor() bool`

HasUnitAmountMinor returns a boolean if a field has been set.

### GetFlatFeeMinor

`func (o *CreateMeterPriceRequest) GetFlatFeeMinor() int32`

GetFlatFeeMinor returns the FlatFeeMinor field if non-nil, zero value otherwise.

### GetFlatFeeMinorOk

`func (o *CreateMeterPriceRequest) GetFlatFeeMinorOk() (*int32, bool)`

GetFlatFeeMinorOk returns a tuple with the FlatFeeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatFeeMinor

`func (o *CreateMeterPriceRequest) SetFlatFeeMinor(v int32)`

SetFlatFeeMinor sets FlatFeeMinor field to given value.

### HasFlatFeeMinor

`func (o *CreateMeterPriceRequest) HasFlatFeeMinor() bool`

HasFlatFeeMinor returns a boolean if a field has been set.

### GetPackageSize

`func (o *CreateMeterPriceRequest) GetPackageSize() int32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *CreateMeterPriceRequest) GetPackageSizeOk() (*int32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *CreateMeterPriceRequest) SetPackageSize(v int32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *CreateMeterPriceRequest) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.

### GetCostPerUnitMinor

`func (o *CreateMeterPriceRequest) GetCostPerUnitMinor() int32`

GetCostPerUnitMinor returns the CostPerUnitMinor field if non-nil, zero value otherwise.

### GetCostPerUnitMinorOk

`func (o *CreateMeterPriceRequest) GetCostPerUnitMinorOk() (*int32, bool)`

GetCostPerUnitMinorOk returns a tuple with the CostPerUnitMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerUnitMinor

`func (o *CreateMeterPriceRequest) SetCostPerUnitMinor(v int32)`

SetCostPerUnitMinor sets CostPerUnitMinor field to given value.

### HasCostPerUnitMinor

`func (o *CreateMeterPriceRequest) HasCostPerUnitMinor() bool`

HasCostPerUnitMinor returns a boolean if a field has been set.

### GetTiers

`func (o *CreateMeterPriceRequest) GetTiers() []CreateMeterPriceRequestTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CreateMeterPriceRequest) GetTiersOk() (*[]CreateMeterPriceRequestTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CreateMeterPriceRequest) SetTiers(v []CreateMeterPriceRequestTiersInner)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *CreateMeterPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


