# CreateMeterPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingModel** | **string** | Pricing model. &#39;percentage&#39; charges a basis-points percentage of the metered amount (REV-3465); other models charge per unit, graduated tiers, volume tiers, or packages. | 
**Currency** | **string** | ISO 4217 currency code. A usage-based subscription has one currency; metered events attributed to it are assumed to be in this currency and are not FX-converted by RevKeen. | 
**UnitAmountMinor** | Pointer to **int32** | Price per unit in minor units (for per_unit). For a sub-penny rate use &#x60;unit_amount_decimal&#x60;. | [optional] 
**UnitAmountDecimal** | Pointer to **string** | Exact per-unit rate in minor units as a decimal string (sub-penny safe, e.g. \&quot;0.02\&quot; &#x3D; 0.02 pence). Takes precedence over &#x60;unit_amount_minor&#x60;. Mirrors Stripe&#39;s &#x60;unit_amount_decimal&#x60;. | [optional] 
**FlatFeeMinor** | Pointer to **int32** | Base charge per period in minor units | [optional] 
**PackageSize** | Pointer to **int32** | Units per package (for package model) | [optional] 
**CostPerUnitMinor** | Pointer to **int32** | Merchant&#39;s cost per unit for margin intelligence | [optional] 
**Name** | Pointer to **string** | Customer-facing price name | [optional] 
**Nickname** | Pointer to **string** | Internal price nickname | [optional] 
**Description** | Pointer to **string** | Description of this price option | [optional] 
**Tiers** | Pointer to [**[]PriceTier**](PriceTier.md) | Tier configuration (for graduated and volume) | [optional] 
**FilterConditions** | Pointer to **[]map[string]interface{}** | Optional event property predicates for dimensional pricing. Empty or omitted &#x3D; base/fallthrough price. | [optional] 
**FilterPriority** | Pointer to **int32** | Tie-breaker when multiple matching filters have the same specificity. | [optional] 
**PayInAdvance** | Pointer to **bool** | When true, matching usage events debit the customer&#39;s prepaid wallet at ingest time. | [optional] 
**SpendingMinimumCents** | Pointer to **int32** | Minimum minor-unit amount applied to each pay-in-advance usage drawdown. | [optional] 
**PercentageRateBps** | Pointer to **int32** | Rate in basis points. 70 &#x3D; 0.7%. Required when pricing_model&#x3D;&#39;percentage&#39;. Range 0–10000 (0%–100%). | [optional] 

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

### GetUnitAmountDecimal

`func (o *CreateMeterPriceRequest) GetUnitAmountDecimal() string`

GetUnitAmountDecimal returns the UnitAmountDecimal field if non-nil, zero value otherwise.

### GetUnitAmountDecimalOk

`func (o *CreateMeterPriceRequest) GetUnitAmountDecimalOk() (*string, bool)`

GetUnitAmountDecimalOk returns a tuple with the UnitAmountDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountDecimal

`func (o *CreateMeterPriceRequest) SetUnitAmountDecimal(v string)`

SetUnitAmountDecimal sets UnitAmountDecimal field to given value.

### HasUnitAmountDecimal

`func (o *CreateMeterPriceRequest) HasUnitAmountDecimal() bool`

HasUnitAmountDecimal returns a boolean if a field has been set.

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

### GetName

`func (o *CreateMeterPriceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMeterPriceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMeterPriceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateMeterPriceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *CreateMeterPriceRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreateMeterPriceRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreateMeterPriceRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CreateMeterPriceRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMeterPriceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMeterPriceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMeterPriceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMeterPriceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTiers

`func (o *CreateMeterPriceRequest) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CreateMeterPriceRequest) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CreateMeterPriceRequest) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *CreateMeterPriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetFilterConditions

`func (o *CreateMeterPriceRequest) GetFilterConditions() []map[string]interface{}`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *CreateMeterPriceRequest) GetFilterConditionsOk() (*[]map[string]interface{}, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *CreateMeterPriceRequest) SetFilterConditions(v []map[string]interface{})`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *CreateMeterPriceRequest) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### GetFilterPriority

`func (o *CreateMeterPriceRequest) GetFilterPriority() int32`

GetFilterPriority returns the FilterPriority field if non-nil, zero value otherwise.

### GetFilterPriorityOk

`func (o *CreateMeterPriceRequest) GetFilterPriorityOk() (*int32, bool)`

GetFilterPriorityOk returns a tuple with the FilterPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPriority

`func (o *CreateMeterPriceRequest) SetFilterPriority(v int32)`

SetFilterPriority sets FilterPriority field to given value.

### HasFilterPriority

`func (o *CreateMeterPriceRequest) HasFilterPriority() bool`

HasFilterPriority returns a boolean if a field has been set.

### GetPayInAdvance

`func (o *CreateMeterPriceRequest) GetPayInAdvance() bool`

GetPayInAdvance returns the PayInAdvance field if non-nil, zero value otherwise.

### GetPayInAdvanceOk

`func (o *CreateMeterPriceRequest) GetPayInAdvanceOk() (*bool, bool)`

GetPayInAdvanceOk returns a tuple with the PayInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayInAdvance

`func (o *CreateMeterPriceRequest) SetPayInAdvance(v bool)`

SetPayInAdvance sets PayInAdvance field to given value.

### HasPayInAdvance

`func (o *CreateMeterPriceRequest) HasPayInAdvance() bool`

HasPayInAdvance returns a boolean if a field has been set.

### GetSpendingMinimumCents

`func (o *CreateMeterPriceRequest) GetSpendingMinimumCents() int32`

GetSpendingMinimumCents returns the SpendingMinimumCents field if non-nil, zero value otherwise.

### GetSpendingMinimumCentsOk

`func (o *CreateMeterPriceRequest) GetSpendingMinimumCentsOk() (*int32, bool)`

GetSpendingMinimumCentsOk returns a tuple with the SpendingMinimumCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendingMinimumCents

`func (o *CreateMeterPriceRequest) SetSpendingMinimumCents(v int32)`

SetSpendingMinimumCents sets SpendingMinimumCents field to given value.

### HasSpendingMinimumCents

`func (o *CreateMeterPriceRequest) HasSpendingMinimumCents() bool`

HasSpendingMinimumCents returns a boolean if a field has been set.

### GetPercentageRateBps

`func (o *CreateMeterPriceRequest) GetPercentageRateBps() int32`

GetPercentageRateBps returns the PercentageRateBps field if non-nil, zero value otherwise.

### GetPercentageRateBpsOk

`func (o *CreateMeterPriceRequest) GetPercentageRateBpsOk() (*int32, bool)`

GetPercentageRateBpsOk returns a tuple with the PercentageRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageRateBps

`func (o *CreateMeterPriceRequest) SetPercentageRateBps(v int32)`

SetPercentageRateBps sets PercentageRateBps field to given value.

### HasPercentageRateBps

`func (o *CreateMeterPriceRequest) HasPercentageRateBps() bool`

HasPercentageRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


