# CartLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProductId** | **string** |  | 
**Name** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitPriceMinor** | **int32** |  | 
**Currency** | **string** |  | 
**Recurring** | [**CartLineItemRecurring**](CartLineItemRecurring.md) |  | 
**BillingMaxCycles** | Pointer to **NullableInt32** |  | [optional] 
**TrialPeriodDays** | Pointer to **NullableInt32** |  | [optional] 
**StartRule** | Pointer to [**NullableCartStartRule**](CartStartRule.md) |  | [optional] 
**BillingAnchorRule** | Pointer to [**NullableCartBillingAnchorRule**](CartBillingAnchorRule.md) |  | [optional] 
**BillingAnchorDay** | Pointer to **NullableInt32** |  | [optional] 
**DueTodayMinor** | Pointer to **NullableInt32** |  | [optional] 
**FirstChargeMinor** | Pointer to **NullableInt32** |  | [optional] 
**FirstRenewalAt** | Pointer to **NullableTime** |  | [optional] 
**EffectiveStartRule** | Pointer to [**NullableCartStartRule**](CartStartRule.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartLineItem

`func NewCartLineItem(id string, productId string, name string, quantity int32, unitPriceMinor int32, currency string, recurring CartLineItemRecurring, ) *CartLineItem`

NewCartLineItem instantiates a new CartLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartLineItemWithDefaults

`func NewCartLineItemWithDefaults() *CartLineItem`

NewCartLineItemWithDefaults instantiates a new CartLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartLineItem) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *CartLineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CartLineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CartLineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *CartLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *CartLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPriceMinor

`func (o *CartLineItem) GetUnitPriceMinor() int32`

GetUnitPriceMinor returns the UnitPriceMinor field if non-nil, zero value otherwise.

### GetUnitPriceMinorOk

`func (o *CartLineItem) GetUnitPriceMinorOk() (*int32, bool)`

GetUnitPriceMinorOk returns a tuple with the UnitPriceMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceMinor

`func (o *CartLineItem) SetUnitPriceMinor(v int32)`

SetUnitPriceMinor sets UnitPriceMinor field to given value.


### GetCurrency

`func (o *CartLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRecurring

`func (o *CartLineItem) GetRecurring() CartLineItemRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CartLineItem) GetRecurringOk() (*CartLineItemRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CartLineItem) SetRecurring(v CartLineItemRecurring)`

SetRecurring sets Recurring field to given value.


### GetBillingMaxCycles

`func (o *CartLineItem) GetBillingMaxCycles() int32`

GetBillingMaxCycles returns the BillingMaxCycles field if non-nil, zero value otherwise.

### GetBillingMaxCyclesOk

`func (o *CartLineItem) GetBillingMaxCyclesOk() (*int32, bool)`

GetBillingMaxCyclesOk returns a tuple with the BillingMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMaxCycles

`func (o *CartLineItem) SetBillingMaxCycles(v int32)`

SetBillingMaxCycles sets BillingMaxCycles field to given value.

### HasBillingMaxCycles

`func (o *CartLineItem) HasBillingMaxCycles() bool`

HasBillingMaxCycles returns a boolean if a field has been set.

### SetBillingMaxCyclesNil

`func (o *CartLineItem) SetBillingMaxCyclesNil(b bool)`

 SetBillingMaxCyclesNil sets the value for BillingMaxCycles to be an explicit nil

### UnsetBillingMaxCycles
`func (o *CartLineItem) UnsetBillingMaxCycles()`

UnsetBillingMaxCycles ensures that no value is present for BillingMaxCycles, not even an explicit nil
### GetTrialPeriodDays

`func (o *CartLineItem) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *CartLineItem) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *CartLineItem) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.

### HasTrialPeriodDays

`func (o *CartLineItem) HasTrialPeriodDays() bool`

HasTrialPeriodDays returns a boolean if a field has been set.

### SetTrialPeriodDaysNil

`func (o *CartLineItem) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *CartLineItem) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil
### GetStartRule

`func (o *CartLineItem) GetStartRule() CartStartRule`

GetStartRule returns the StartRule field if non-nil, zero value otherwise.

### GetStartRuleOk

`func (o *CartLineItem) GetStartRuleOk() (*CartStartRule, bool)`

GetStartRuleOk returns a tuple with the StartRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartRule

`func (o *CartLineItem) SetStartRule(v CartStartRule)`

SetStartRule sets StartRule field to given value.

### HasStartRule

`func (o *CartLineItem) HasStartRule() bool`

HasStartRule returns a boolean if a field has been set.

### SetStartRuleNil

`func (o *CartLineItem) SetStartRuleNil(b bool)`

 SetStartRuleNil sets the value for StartRule to be an explicit nil

### UnsetStartRule
`func (o *CartLineItem) UnsetStartRule()`

UnsetStartRule ensures that no value is present for StartRule, not even an explicit nil
### GetBillingAnchorRule

`func (o *CartLineItem) GetBillingAnchorRule() CartBillingAnchorRule`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *CartLineItem) GetBillingAnchorRuleOk() (*CartBillingAnchorRule, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *CartLineItem) SetBillingAnchorRule(v CartBillingAnchorRule)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *CartLineItem) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### SetBillingAnchorRuleNil

`func (o *CartLineItem) SetBillingAnchorRuleNil(b bool)`

 SetBillingAnchorRuleNil sets the value for BillingAnchorRule to be an explicit nil

### UnsetBillingAnchorRule
`func (o *CartLineItem) UnsetBillingAnchorRule()`

UnsetBillingAnchorRule ensures that no value is present for BillingAnchorRule, not even an explicit nil
### GetBillingAnchorDay

`func (o *CartLineItem) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *CartLineItem) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *CartLineItem) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *CartLineItem) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *CartLineItem) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *CartLineItem) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetDueTodayMinor

`func (o *CartLineItem) GetDueTodayMinor() int32`

GetDueTodayMinor returns the DueTodayMinor field if non-nil, zero value otherwise.

### GetDueTodayMinorOk

`func (o *CartLineItem) GetDueTodayMinorOk() (*int32, bool)`

GetDueTodayMinorOk returns a tuple with the DueTodayMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTodayMinor

`func (o *CartLineItem) SetDueTodayMinor(v int32)`

SetDueTodayMinor sets DueTodayMinor field to given value.

### HasDueTodayMinor

`func (o *CartLineItem) HasDueTodayMinor() bool`

HasDueTodayMinor returns a boolean if a field has been set.

### SetDueTodayMinorNil

`func (o *CartLineItem) SetDueTodayMinorNil(b bool)`

 SetDueTodayMinorNil sets the value for DueTodayMinor to be an explicit nil

### UnsetDueTodayMinor
`func (o *CartLineItem) UnsetDueTodayMinor()`

UnsetDueTodayMinor ensures that no value is present for DueTodayMinor, not even an explicit nil
### GetFirstChargeMinor

`func (o *CartLineItem) GetFirstChargeMinor() int32`

GetFirstChargeMinor returns the FirstChargeMinor field if non-nil, zero value otherwise.

### GetFirstChargeMinorOk

`func (o *CartLineItem) GetFirstChargeMinorOk() (*int32, bool)`

GetFirstChargeMinorOk returns a tuple with the FirstChargeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeMinor

`func (o *CartLineItem) SetFirstChargeMinor(v int32)`

SetFirstChargeMinor sets FirstChargeMinor field to given value.

### HasFirstChargeMinor

`func (o *CartLineItem) HasFirstChargeMinor() bool`

HasFirstChargeMinor returns a boolean if a field has been set.

### SetFirstChargeMinorNil

`func (o *CartLineItem) SetFirstChargeMinorNil(b bool)`

 SetFirstChargeMinorNil sets the value for FirstChargeMinor to be an explicit nil

### UnsetFirstChargeMinor
`func (o *CartLineItem) UnsetFirstChargeMinor()`

UnsetFirstChargeMinor ensures that no value is present for FirstChargeMinor, not even an explicit nil
### GetFirstRenewalAt

`func (o *CartLineItem) GetFirstRenewalAt() time.Time`

GetFirstRenewalAt returns the FirstRenewalAt field if non-nil, zero value otherwise.

### GetFirstRenewalAtOk

`func (o *CartLineItem) GetFirstRenewalAtOk() (*time.Time, bool)`

GetFirstRenewalAtOk returns a tuple with the FirstRenewalAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRenewalAt

`func (o *CartLineItem) SetFirstRenewalAt(v time.Time)`

SetFirstRenewalAt sets FirstRenewalAt field to given value.

### HasFirstRenewalAt

`func (o *CartLineItem) HasFirstRenewalAt() bool`

HasFirstRenewalAt returns a boolean if a field has been set.

### SetFirstRenewalAtNil

`func (o *CartLineItem) SetFirstRenewalAtNil(b bool)`

 SetFirstRenewalAtNil sets the value for FirstRenewalAt to be an explicit nil

### UnsetFirstRenewalAt
`func (o *CartLineItem) UnsetFirstRenewalAt()`

UnsetFirstRenewalAt ensures that no value is present for FirstRenewalAt, not even an explicit nil
### GetEffectiveStartRule

`func (o *CartLineItem) GetEffectiveStartRule() CartStartRule`

GetEffectiveStartRule returns the EffectiveStartRule field if non-nil, zero value otherwise.

### GetEffectiveStartRuleOk

`func (o *CartLineItem) GetEffectiveStartRuleOk() (*CartStartRule, bool)`

GetEffectiveStartRuleOk returns a tuple with the EffectiveStartRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStartRule

`func (o *CartLineItem) SetEffectiveStartRule(v CartStartRule)`

SetEffectiveStartRule sets EffectiveStartRule field to given value.

### HasEffectiveStartRule

`func (o *CartLineItem) HasEffectiveStartRule() bool`

HasEffectiveStartRule returns a boolean if a field has been set.

### SetEffectiveStartRuleNil

`func (o *CartLineItem) SetEffectiveStartRuleNil(b bool)`

 SetEffectiveStartRuleNil sets the value for EffectiveStartRule to be an explicit nil

### UnsetEffectiveStartRule
`func (o *CartLineItem) UnsetEffectiveStartRule()`

UnsetEffectiveStartRule ensures that no value is present for EffectiveStartRule, not even an explicit nil
### GetMetadata

`func (o *CartLineItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CartLineItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CartLineItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CartLineItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


