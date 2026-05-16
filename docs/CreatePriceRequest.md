# CreatePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | ID of the product this price belongs to | 
**Currency** | Pointer to **string** | Three-letter ISO currency code | [optional] [default to "usd"]
**UnitAmount** | Pointer to **int32** | Price in minor units (cents). Required for fixed pricing. | [optional] 
**Type** | Pointer to **string** | Price type | [optional] [default to "one_time"]
**PricingModel** | Pointer to **string** | Pricing model | [optional] [default to "fixed"]
**Interval** | Pointer to **string** | Billing interval (required for recurring) | [optional] 
**IntervalCount** | Pointer to **int32** | Number of intervals between billings (1-12) | [optional] 
**TrialPeriodDays** | Pointer to **int32** | Trial period in days (0-730) | [optional] 
**MinimumAmount** | Pointer to **int32** | PWYW: minimum amount in cents | [optional] 
**MaximumAmount** | Pointer to **int32** | PWYW: maximum amount in cents | [optional] 
**SuggestedAmount** | Pointer to **int32** | PWYW: suggested amount in cents | [optional] 
**PresetAmounts** | Pointer to **[]int32** | PWYW: quick-select amounts (max 10) | [optional] 
**Nickname** | Pointer to **string** | Display name (e.g., &#39;Monthly&#39;, &#39;Annual&#39;) | [optional] 
**LookupKey** | Pointer to **string** | Stable key for API lookups | [optional] 
**BillingScheme** | Pointer to **string** | Billing scheme. Defaults to &#x60;per_unit&#x60;. Set to &#x60;tiered&#x60; with &#x60;tiers_mode&#x60; + &#x60;tiers&#x60; for graduated or volume pricing. | [optional] 
**TiersMode** | Pointer to **string** | Tiered pricing mode. Required when &#x60;billing_scheme&#x60; is &#x60;tiered&#x60;, must be omitted otherwise. | [optional] 
**Tiers** | Pointer to [**[]PriceTier**](PriceTier.md) | Price tiers (ordered by &#x60;up_to&#x60;, min 2, max 50). Required when &#x60;billing_scheme&#x60; is &#x60;tiered&#x60;. Final tier MUST have &#x60;up_to: null&#x60;. | [optional] 
**TransformQuantity** | Pointer to [**TransformQuantity**](TransformQuantity.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewCreatePriceRequest

`func NewCreatePriceRequest(productId string, ) *CreatePriceRequest`

NewCreatePriceRequest instantiates a new CreatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePriceRequestWithDefaults

`func NewCreatePriceRequestWithDefaults() *CreatePriceRequest`

NewCreatePriceRequestWithDefaults instantiates a new CreatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreatePriceRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreatePriceRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreatePriceRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetCurrency

`func (o *CreatePriceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePriceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePriceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePriceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnitAmount

`func (o *CreatePriceRequest) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *CreatePriceRequest) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *CreatePriceRequest) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *CreatePriceRequest) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetType

`func (o *CreatePriceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePriceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePriceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreatePriceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPricingModel

`func (o *CreatePriceRequest) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreatePriceRequest) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreatePriceRequest) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *CreatePriceRequest) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### GetInterval

`func (o *CreatePriceRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreatePriceRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreatePriceRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreatePriceRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalCount

`func (o *CreatePriceRequest) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *CreatePriceRequest) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *CreatePriceRequest) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *CreatePriceRequest) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetTrialPeriodDays

`func (o *CreatePriceRequest) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CreatePriceRequest) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CreatePriceRequest) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CreatePriceRequest) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *CreatePriceRequest) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *CreatePriceRequest) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *CreatePriceRequest) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *CreatePriceRequest) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *CreatePriceRequest) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *CreatePriceRequest) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *CreatePriceRequest) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *CreatePriceRequest) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### GetSuggestedAmount

`func (o *CreatePriceRequest) GetSuggestedAmount() int32`

GetSuggestedAmount returns the SuggestedAmount field if non-nil, zero value otherwise.

### GetSuggestedAmountOk

`func (o *CreatePriceRequest) GetSuggestedAmountOk() (*int32, bool)`

GetSuggestedAmountOk returns a tuple with the SuggestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAmount

`func (o *CreatePriceRequest) SetSuggestedAmount(v int32)`

SetSuggestedAmount sets SuggestedAmount field to given value.

### HasSuggestedAmount

`func (o *CreatePriceRequest) HasSuggestedAmount() bool`

HasSuggestedAmount returns a boolean if a field has been set.

### GetPresetAmounts

`func (o *CreatePriceRequest) GetPresetAmounts() []int32`

GetPresetAmounts returns the PresetAmounts field if non-nil, zero value otherwise.

### GetPresetAmountsOk

`func (o *CreatePriceRequest) GetPresetAmountsOk() (*[]int32, bool)`

GetPresetAmountsOk returns a tuple with the PresetAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetAmounts

`func (o *CreatePriceRequest) SetPresetAmounts(v []int32)`

SetPresetAmounts sets PresetAmounts field to given value.

### HasPresetAmounts

`func (o *CreatePriceRequest) HasPresetAmounts() bool`

HasPresetAmounts returns a boolean if a field has been set.

### GetNickname

`func (o *CreatePriceRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *CreatePriceRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *CreatePriceRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *CreatePriceRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetLookupKey

`func (o *CreatePriceRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *CreatePriceRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *CreatePriceRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *CreatePriceRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetBillingScheme

`func (o *CreatePriceRequest) GetBillingScheme() string`

GetBillingScheme returns the BillingScheme field if non-nil, zero value otherwise.

### GetBillingSchemeOk

`func (o *CreatePriceRequest) GetBillingSchemeOk() (*string, bool)`

GetBillingSchemeOk returns a tuple with the BillingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingScheme

`func (o *CreatePriceRequest) SetBillingScheme(v string)`

SetBillingScheme sets BillingScheme field to given value.

### HasBillingScheme

`func (o *CreatePriceRequest) HasBillingScheme() bool`

HasBillingScheme returns a boolean if a field has been set.

### GetTiersMode

`func (o *CreatePriceRequest) GetTiersMode() string`

GetTiersMode returns the TiersMode field if non-nil, zero value otherwise.

### GetTiersModeOk

`func (o *CreatePriceRequest) GetTiersModeOk() (*string, bool)`

GetTiersModeOk returns a tuple with the TiersMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersMode

`func (o *CreatePriceRequest) SetTiersMode(v string)`

SetTiersMode sets TiersMode field to given value.

### HasTiersMode

`func (o *CreatePriceRequest) HasTiersMode() bool`

HasTiersMode returns a boolean if a field has been set.

### GetTiers

`func (o *CreatePriceRequest) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CreatePriceRequest) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CreatePriceRequest) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *CreatePriceRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTransformQuantity

`func (o *CreatePriceRequest) GetTransformQuantity() TransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *CreatePriceRequest) GetTransformQuantityOk() (*TransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *CreatePriceRequest) SetTransformQuantity(v TransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *CreatePriceRequest) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePriceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePriceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePriceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePriceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


