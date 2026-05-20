# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the price | 
**Object** | **string** | Object type | 
**ProductId** | **string** | ID of the product this price belongs to | 
**Active** | **bool** | Whether the price is active | 
**Currency** | **string** | Three-letter ISO currency code (lowercase) | 
**UnitAmount** | **NullableInt32** | Price in minor units (cents) | 
**Type** | **string** | Price type | 
**PricingModel** | **string** | Pricing model | 
**Interval** | Pointer to **NullableString** | Billing interval (recurring only) | [optional] 
**IntervalCount** | Pointer to **NullableInt32** | Number of intervals between billings | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** | Trial period in days (recurring only) | [optional] 
**MinimumAmount** | Pointer to **NullableInt32** | PWYW: minimum amount in cents | [optional] 
**MaximumAmount** | Pointer to **NullableInt32** | PWYW: maximum amount in cents | [optional] 
**SuggestedAmount** | Pointer to **NullableInt32** | PWYW: suggested amount in cents | [optional] 
**PresetAmounts** | Pointer to **[]int32** | PWYW: quick-select amounts | [optional] 
**Nickname** | Pointer to **NullableString** | Display name (e.g., &#39;Monthly&#39;, &#39;Annual - Save 17%&#39;) | [optional] 
**LookupKey** | Pointer to **NullableString** | Stable key for API lookups | [optional] 
**BillingScheme** | **string** | Billing scheme. Defaults to &#x60;per_unit&#x60;. Set to &#x60;tiered&#x60; along with &#x60;tiers_mode&#x60; + &#x60;tiers&#x60; to use graduated or volume pricing. | 
**TiersMode** | Pointer to **NullableString** | Tiered pricing mode. Required when &#x60;billing_scheme&#x60; is &#x60;tiered&#x60;, must be null otherwise. | [optional] 
**Tiers** | Pointer to [**[]PriceTier**](PriceTier.md) | Price tiers (ordered by &#x60;up_to&#x60;). Present only when &#x60;billing_scheme&#x60; is &#x60;tiered&#x60;. At least 2 tiers, final tier MUST have &#x60;up_to: null&#x60;. Immutable after creation. | [optional] 
**TransformQuantity** | Pointer to [**PriceTransformQuantity**](PriceTransformQuantity.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp | 
**UpdatedAt** | **time.Time** | Last update timestamp | 

## Methods

### NewPrice

`func NewPrice(id string, object string, productId string, active bool, currency string, unitAmount NullableInt32, type_ string, pricingModel string, billingScheme string, createdAt time.Time, updatedAt time.Time, ) *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Price) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Price) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Price) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Price) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Price) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Price) SetObject(v string)`

SetObject sets Object field to given value.


### GetProductId

`func (o *Price) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Price) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Price) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetActive

`func (o *Price) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Price) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Price) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitAmount

`func (o *Price) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *Price) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *Price) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### SetUnitAmountNil

`func (o *Price) SetUnitAmountNil(b bool)`

 SetUnitAmountNil sets the value for UnitAmount to be an explicit nil

### UnsetUnitAmount
`func (o *Price) UnsetUnitAmount()`

UnsetUnitAmount ensures that no value is present for UnitAmount, not even an explicit nil
### GetType

`func (o *Price) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Price) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Price) SetType(v string)`

SetType sets Type field to given value.


### GetPricingModel

`func (o *Price) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *Price) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *Price) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetInterval

`func (o *Price) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Price) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Price) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Price) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *Price) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *Price) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *Price) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Price) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Price) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Price) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *Price) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *Price) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialPeriodDays

`func (o *Price) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *Price) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *Price) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *Price) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *Price) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *Price) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetMinimumAmount

`func (o *Price) GetMinimumAmount() int32`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *Price) GetMinimumAmountOk() (*int32, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *Price) SetMinimumAmount(v int32)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *Price) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### SetMinimumAmountNil

`func (o *Price) SetMinimumAmountNil(b bool)`

 SetMinimumAmountNil sets the value for MinimumAmount to be an explicit nil

### UnsetMinimumAmount
`func (o *Price) UnsetMinimumAmount()`

UnsetMinimumAmount ensures that no value is present for MinimumAmount, not even an explicit nil
### GetMaximumAmount

`func (o *Price) GetMaximumAmount() int32`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *Price) GetMaximumAmountOk() (*int32, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *Price) SetMaximumAmount(v int32)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *Price) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.

### SetMaximumAmountNil

`func (o *Price) SetMaximumAmountNil(b bool)`

 SetMaximumAmountNil sets the value for MaximumAmount to be an explicit nil

### UnsetMaximumAmount
`func (o *Price) UnsetMaximumAmount()`

UnsetMaximumAmount ensures that no value is present for MaximumAmount, not even an explicit nil
### GetSuggestedAmount

`func (o *Price) GetSuggestedAmount() int32`

GetSuggestedAmount returns the SuggestedAmount field if non-nil, zero value otherwise.

### GetSuggestedAmountOk

`func (o *Price) GetSuggestedAmountOk() (*int32, bool)`

GetSuggestedAmountOk returns a tuple with the SuggestedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAmount

`func (o *Price) SetSuggestedAmount(v int32)`

SetSuggestedAmount sets SuggestedAmount field to given value.

### HasSuggestedAmount

`func (o *Price) HasSuggestedAmount() bool`

HasSuggestedAmount returns a boolean if a field has been set.

### SetSuggestedAmountNil

`func (o *Price) SetSuggestedAmountNil(b bool)`

 SetSuggestedAmountNil sets the value for SuggestedAmount to be an explicit nil

### UnsetSuggestedAmount
`func (o *Price) UnsetSuggestedAmount()`

UnsetSuggestedAmount ensures that no value is present for SuggestedAmount, not even an explicit nil
### GetPresetAmounts

`func (o *Price) GetPresetAmounts() []int32`

GetPresetAmounts returns the PresetAmounts field if non-nil, zero value otherwise.

### GetPresetAmountsOk

`func (o *Price) GetPresetAmountsOk() (*[]int32, bool)`

GetPresetAmountsOk returns a tuple with the PresetAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetAmounts

`func (o *Price) SetPresetAmounts(v []int32)`

SetPresetAmounts sets PresetAmounts field to given value.

### HasPresetAmounts

`func (o *Price) HasPresetAmounts() bool`

HasPresetAmounts returns a boolean if a field has been set.

### SetPresetAmountsNil

`func (o *Price) SetPresetAmountsNil(b bool)`

 SetPresetAmountsNil sets the value for PresetAmounts to be an explicit nil

### UnsetPresetAmounts
`func (o *Price) UnsetPresetAmounts()`

UnsetPresetAmounts ensures that no value is present for PresetAmounts, not even an explicit nil
### GetNickname

`func (o *Price) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Price) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Price) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Price) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *Price) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *Price) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetLookupKey

`func (o *Price) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *Price) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *Price) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *Price) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### SetLookupKeyNil

`func (o *Price) SetLookupKeyNil(b bool)`

 SetLookupKeyNil sets the value for LookupKey to be an explicit nil

### UnsetLookupKey
`func (o *Price) UnsetLookupKey()`

UnsetLookupKey ensures that no value is present for LookupKey, not even an explicit nil
### GetBillingScheme

`func (o *Price) GetBillingScheme() string`

GetBillingScheme returns the BillingScheme field if non-nil, zero value otherwise.

### GetBillingSchemeOk

`func (o *Price) GetBillingSchemeOk() (*string, bool)`

GetBillingSchemeOk returns a tuple with the BillingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingScheme

`func (o *Price) SetBillingScheme(v string)`

SetBillingScheme sets BillingScheme field to given value.


### GetTiersMode

`func (o *Price) GetTiersMode() string`

GetTiersMode returns the TiersMode field if non-nil, zero value otherwise.

### GetTiersModeOk

`func (o *Price) GetTiersModeOk() (*string, bool)`

GetTiersModeOk returns a tuple with the TiersMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiersMode

`func (o *Price) SetTiersMode(v string)`

SetTiersMode sets TiersMode field to given value.

### HasTiersMode

`func (o *Price) HasTiersMode() bool`

HasTiersMode returns a boolean if a field has been set.

### SetTiersModeNil

`func (o *Price) SetTiersModeNil(b bool)`

 SetTiersModeNil sets the value for TiersMode to be an explicit nil

### UnsetTiersMode
`func (o *Price) UnsetTiersMode()`

UnsetTiersMode ensures that no value is present for TiersMode, not even an explicit nil
### GetTiers

`func (o *Price) GetTiers() []PriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *Price) GetTiersOk() (*[]PriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *Price) SetTiers(v []PriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *Price) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiersNil

`func (o *Price) SetTiersNil(b bool)`

 SetTiersNil sets the value for Tiers to be an explicit nil

### UnsetTiers
`func (o *Price) UnsetTiers()`

UnsetTiers ensures that no value is present for Tiers, not even an explicit nil
### GetTransformQuantity

`func (o *Price) GetTransformQuantity() PriceTransformQuantity`

GetTransformQuantity returns the TransformQuantity field if non-nil, zero value otherwise.

### GetTransformQuantityOk

`func (o *Price) GetTransformQuantityOk() (*PriceTransformQuantity, bool)`

GetTransformQuantityOk returns a tuple with the TransformQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformQuantity

`func (o *Price) SetTransformQuantity(v PriceTransformQuantity)`

SetTransformQuantity sets TransformQuantity field to given value.

### HasTransformQuantity

`func (o *Price) HasTransformQuantity() bool`

HasTransformQuantity returns a boolean if a field has been set.

### GetMetadata

`func (o *Price) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Price) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Price) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Price) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Price) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Price) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Price) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Price) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Price) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Price) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


