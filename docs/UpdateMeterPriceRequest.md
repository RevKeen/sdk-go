# UpdateMeterPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitAmountMinor** | Pointer to **int32** |  | [optional] 
**FlatFeeMinor** | Pointer to **int32** |  | [optional] 
**PackageSize** | Pointer to **int32** |  | [optional] 
**CostPerUnitMinor** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Nickname** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Tiers** | Pointer to [**[]PriceTier**](PriceTier.md) |  | [optional] 
**PercentageRateBps** | Pointer to **int32** | Rate in basis points. 70 &#x3D; 0.7%. Only meaningful when the price&#39;s pricing_model is &#39;percentage&#39;. | [optional] 

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

### GetName

`func (o *UpdateMeterPriceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMeterPriceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMeterPriceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMeterPriceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateMeterPriceRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateMeterPriceRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNickname

`func (o *UpdateMeterPriceRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdateMeterPriceRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdateMeterPriceRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdateMeterPriceRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *UpdateMeterPriceRequest) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *UpdateMeterPriceRequest) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetDescription

`func (o *UpdateMeterPriceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMeterPriceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMeterPriceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMeterPriceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMeterPriceRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMeterPriceRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTiers

`func (o *UpdateMeterPriceRequest) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *UpdateMeterPriceRequest) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *UpdateMeterPriceRequest) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *UpdateMeterPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetPercentageRateBps

`func (o *UpdateMeterPriceRequest) GetPercentageRateBps() int32`

GetPercentageRateBps returns the PercentageRateBps field if non-nil, zero value otherwise.

### GetPercentageRateBpsOk

`func (o *UpdateMeterPriceRequest) GetPercentageRateBpsOk() (*int32, bool)`

GetPercentageRateBpsOk returns a tuple with the PercentageRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRateBps

`func (o *UpdateMeterPriceRequest) SetPercentageRateBps(v int32)`

SetPercentageRateBps sets PercentageRateBps field to given value.

### HasPercentageRateBps

`func (o *UpdateMeterPriceRequest) HasPercentageRateBps() bool`

HasPercentageRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


