# UpdateMeterPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitAmountMinor** | Pointer to **int32** |  | [optional] 
**FlatFeeMinor** | Pointer to **int32** |  | [optional] 
**PackageSize** | Pointer to **int32** |  | [optional] 
**CostPerUnitMinor** | Pointer to **int32** |  | [optional] 
**Tiers** | Pointer to [**[]CreateMeterPriceRequestTiersInner**](CreateMeterPriceRequestTiersInner.md) |  | [optional] 

## Methods

### NewUpdateMeterPriceRequest

`func NewUpdateMeterPriceRequest() *UpdateMeterPriceRequest`

NewUpdateMeterPriceRequest instantiates a new UpdateMeterPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeterPriceRequestWithDefaults

`func NewUpdateMeterPriceRequestWithDefaults() *UpdateMeterPriceRequest`

NewUpdateMeterPriceRequestWithDefaults instantiates a new UpdateMeterPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitAmountMinor

`func (o *UpdateMeterPriceRequest) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *UpdateMeterPriceRequest) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *UpdateMeterPriceRequest) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.

### HasUnitAmountMinor

`func (o *UpdateMeterPriceRequest) HasUnitAmountMinor() bool`

HasUnitAmountMinor returns a boolean if a field has been set.

### GetFlatFeeMinor

`func (o *UpdateMeterPriceRequest) GetFlatFeeMinor() int32`

GetFlatFeeMinor returns the FlatFeeMinor field if non-nil, zero value otherwise.

### GetFlatFeeMinorOk

`func (o *UpdateMeterPriceRequest) GetFlatFeeMinorOk() (*int32, bool)`

GetFlatFeeMinorOk returns a tuple with the FlatFeeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatFeeMinor

`func (o *UpdateMeterPriceRequest) SetFlatFeeMinor(v int32)`

SetFlatFeeMinor sets FlatFeeMinor field to given value.

### HasFlatFeeMinor

`func (o *UpdateMeterPriceRequest) HasFlatFeeMinor() bool`

HasFlatFeeMinor returns a boolean if a field has been set.

### GetPackageSize

`func (o *UpdateMeterPriceRequest) GetPackageSize() int32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *UpdateMeterPriceRequest) GetPackageSizeOk() (*int32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *UpdateMeterPriceRequest) SetPackageSize(v int32)`

SetPackageSize sets PackageSize field to given value.

### HasPackageSize

`func (o *UpdateMeterPriceRequest) HasPackageSize() bool`

HasPackageSize returns a boolean if a field has been set.

### GetCostPerUnitMinor

`func (o *UpdateMeterPriceRequest) GetCostPerUnitMinor() int32`

GetCostPerUnitMinor returns the CostPerUnitMinor field if non-nil, zero value otherwise.

### GetCostPerUnitMinorOk

`func (o *UpdateMeterPriceRequest) GetCostPerUnitMinorOk() (*int32, bool)`

GetCostPerUnitMinorOk returns a tuple with the CostPerUnitMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerUnitMinor

`func (o *UpdateMeterPriceRequest) SetCostPerUnitMinor(v int32)`

SetCostPerUnitMinor sets CostPerUnitMinor field to given value.

### HasCostPerUnitMinor

`func (o *UpdateMeterPriceRequest) HasCostPerUnitMinor() bool`

HasCostPerUnitMinor returns a boolean if a field has been set.

### GetTiers

`func (o *UpdateMeterPriceRequest) GetTiers() []CreateMeterPriceRequestTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *UpdateMeterPriceRequest) GetTiersOk() (*[]CreateMeterPriceRequestTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *UpdateMeterPriceRequest) SetTiers(v []CreateMeterPriceRequestTiersInner)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *UpdateMeterPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


