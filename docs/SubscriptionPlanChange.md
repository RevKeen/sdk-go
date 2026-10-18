# SubscriptionPlanChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  |
**CurrentPriceId** | **string** | The price the subscription is on before the change. |
**CurrentPlanName** | **NullableString** | The current plan&#39;s product name. |
**NewPriceId** | **string** | The price the subscription moves to. A product id in the request resolves to its default price. |
**NewPlanName** | **NullableString** | The new plan&#39;s product name. |
**NewProductId** | **NullableString** | The product of the new price. |
**BillingInterval** | **NullableString** | The new price&#39;s interval. Always the subscription&#39;s own: changing interval is refused. |
**BillingIntervalCount** | **NullableInt32** | The new price&#39;s interval count. |
**Quantity** | **int32** | The subscribed quantity, which a plan change keeps. |
**CurrentAmountMinor** | **NullableInt64** | The renewal amount before the change, in minor units. |
**NewAmountMinor** | **int64** | The renewal amount after the change: the new unit amount times &#x60;quantity&#x60;. Always above zero. |
**ProrationBehavior** | **string** | How the difference is billed. &#x60;next_invoice&#x60; (also what &#x60;invoice_immediately&#x60; means here) adds it to the next renewal invoice; &#x60;none&#x60; changes the plan with no proration. |
**EffectiveDate** | **time.Time** | When the change takes effect for proration (ISO 8601). |
**Proration** | Pointer to [**SubscriptionPlanChangeProration**](SubscriptionPlanChangeProration.md) |  | [optional]
**PendingProrationNetAmountMinor** | **NullableInt64** | What the next renewal invoice adds (positive) or credits (negative) after this change, including anything already pending from earlier changes. Null when nothing is pending. |

## Methods

### NewSubscriptionPlanChange

`func NewSubscriptionPlanChange(subscriptionId string, currentPriceId string, currentPlanName NullableString, newPriceId string, newPlanName NullableString, newProductId NullableString, billingInterval NullableString, billingIntervalCount NullableInt32, quantity int32, currentAmountMinor NullableInt64, newAmountMinor int64, prorationBehavior string, effectiveDate time.Time, pendingProrationNetAmountMinor NullableInt64, ) *SubscriptionPlanChange`

NewSubscriptionPlanChange instantiates a new SubscriptionPlanChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanChangeWithDefaults

`func NewSubscriptionPlanChangeWithDefaults() *SubscriptionPlanChange`

NewSubscriptionPlanChangeWithDefaults instantiates a new SubscriptionPlanChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SubscriptionPlanChange) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionPlanChange) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionPlanChange) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCurrentPriceId

`func (o *SubscriptionPlanChange) GetCurrentPriceId() string`

GetCurrentPriceId returns the CurrentPriceId field if non-nil, zero value otherwise.

### GetCurrentPriceIdOk

`func (o *SubscriptionPlanChange) GetCurrentPriceIdOk() (*string, bool)`

GetCurrentPriceIdOk returns a tuple with the CurrentPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPriceId

`func (o *SubscriptionPlanChange) SetCurrentPriceId(v string)`

SetCurrentPriceId sets CurrentPriceId field to given value.


### GetCurrentPlanName

`func (o *SubscriptionPlanChange) GetCurrentPlanName() string`

GetCurrentPlanName returns the CurrentPlanName field if non-nil, zero value otherwise.

### GetCurrentPlanNameOk

`func (o *SubscriptionPlanChange) GetCurrentPlanNameOk() (*string, bool)`

GetCurrentPlanNameOk returns a tuple with the CurrentPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlanName

`func (o *SubscriptionPlanChange) SetCurrentPlanName(v string)`

SetCurrentPlanName sets CurrentPlanName field to given value.


### SetCurrentPlanNameNil

`func (o *SubscriptionPlanChange) SetCurrentPlanNameNil(b bool)`

 SetCurrentPlanNameNil sets the value for CurrentPlanName to be an explicit nil

### UnsetCurrentPlanName
`func (o *SubscriptionPlanChange) UnsetCurrentPlanName()`

UnsetCurrentPlanName ensures that no value is present for CurrentPlanName, not even an explicit nil
### GetNewPriceId

`func (o *SubscriptionPlanChange) GetNewPriceId() string`

GetNewPriceId returns the NewPriceId field if non-nil, zero value otherwise.

### GetNewPriceIdOk

`func (o *SubscriptionPlanChange) GetNewPriceIdOk() (*string, bool)`

GetNewPriceIdOk returns a tuple with the NewPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPriceId

`func (o *SubscriptionPlanChange) SetNewPriceId(v string)`

SetNewPriceId sets NewPriceId field to given value.


### GetNewPlanName

`func (o *SubscriptionPlanChange) GetNewPlanName() string`

GetNewPlanName returns the NewPlanName field if non-nil, zero value otherwise.

### GetNewPlanNameOk

`func (o *SubscriptionPlanChange) GetNewPlanNameOk() (*string, bool)`

GetNewPlanNameOk returns a tuple with the NewPlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanName

`func (o *SubscriptionPlanChange) SetNewPlanName(v string)`

SetNewPlanName sets NewPlanName field to given value.


### SetNewPlanNameNil

`func (o *SubscriptionPlanChange) SetNewPlanNameNil(b bool)`

 SetNewPlanNameNil sets the value for NewPlanName to be an explicit nil

### UnsetNewPlanName
`func (o *SubscriptionPlanChange) UnsetNewPlanName()`

UnsetNewPlanName ensures that no value is present for NewPlanName, not even an explicit nil
### GetNewProductId

`func (o *SubscriptionPlanChange) GetNewProductId() string`

GetNewProductId returns the NewProductId field if non-nil, zero value otherwise.

### GetNewProductIdOk

`func (o *SubscriptionPlanChange) GetNewProductIdOk() (*string, bool)`

GetNewProductIdOk returns a tuple with the NewProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewProductId

`func (o *SubscriptionPlanChange) SetNewProductId(v string)`

SetNewProductId sets NewProductId field to given value.


### SetNewProductIdNil

`func (o *SubscriptionPlanChange) SetNewProductIdNil(b bool)`

 SetNewProductIdNil sets the value for NewProductId to be an explicit nil

### UnsetNewProductId
`func (o *SubscriptionPlanChange) UnsetNewProductId()`

UnsetNewProductId ensures that no value is present for NewProductId, not even an explicit nil
### GetBillingInterval

`func (o *SubscriptionPlanChange) GetBillingInterval() string`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *SubscriptionPlanChange) GetBillingIntervalOk() (*string, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *SubscriptionPlanChange) SetBillingInterval(v string)`

SetBillingInterval sets BillingInterval field to given value.


### SetBillingIntervalNil

`func (o *SubscriptionPlanChange) SetBillingIntervalNil(b bool)`

 SetBillingIntervalNil sets the value for BillingInterval to be an explicit nil

### UnsetBillingInterval
`func (o *SubscriptionPlanChange) UnsetBillingInterval()`

UnsetBillingInterval ensures that no value is present for BillingInterval, not even an explicit nil
### GetBillingIntervalCount

`func (o *SubscriptionPlanChange) GetBillingIntervalCount() int32`

GetBillingIntervalCount returns the BillingIntervalCount field if non-nil, zero value otherwise.

### GetBillingIntervalCountOk

`func (o *SubscriptionPlanChange) GetBillingIntervalCountOk() (*int32, bool)`

GetBillingIntervalCountOk returns a tuple with the BillingIntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalCount

`func (o *SubscriptionPlanChange) SetBillingIntervalCount(v int32)`

SetBillingIntervalCount sets BillingIntervalCount field to given value.


### SetBillingIntervalCountNil

`func (o *SubscriptionPlanChange) SetBillingIntervalCountNil(b bool)`

 SetBillingIntervalCountNil sets the value for BillingIntervalCount to be an explicit nil

### UnsetBillingIntervalCount
`func (o *SubscriptionPlanChange) UnsetBillingIntervalCount()`

UnsetBillingIntervalCount ensures that no value is present for BillingIntervalCount, not even an explicit nil
### GetQuantity

`func (o *SubscriptionPlanChange) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionPlanChange) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionPlanChange) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCurrentAmountMinor

`func (o *SubscriptionPlanChange) GetCurrentAmountMinor() int64`

GetCurrentAmountMinor returns the CurrentAmountMinor field if non-nil, zero value otherwise.

### GetCurrentAmountMinorOk

`func (o *SubscriptionPlanChange) GetCurrentAmountMinorOk() (*int64, bool)`

GetCurrentAmountMinorOk returns a tuple with the CurrentAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAmountMinor

`func (o *SubscriptionPlanChange) SetCurrentAmountMinor(v int64)`

SetCurrentAmountMinor sets CurrentAmountMinor field to given value.


### SetCurrentAmountMinorNil

`func (o *SubscriptionPlanChange) SetCurrentAmountMinorNil(b bool)`

 SetCurrentAmountMinorNil sets the value for CurrentAmountMinor to be an explicit nil

### UnsetCurrentAmountMinor
`func (o *SubscriptionPlanChange) UnsetCurrentAmountMinor()`

UnsetCurrentAmountMinor ensures that no value is present for CurrentAmountMinor, not even an explicit nil
### GetNewAmountMinor

`func (o *SubscriptionPlanChange) GetNewAmountMinor() int64`

GetNewAmountMinor returns the NewAmountMinor field if non-nil, zero value otherwise.

### GetNewAmountMinorOk

`func (o *SubscriptionPlanChange) GetNewAmountMinorOk() (*int64, bool)`

GetNewAmountMinorOk returns a tuple with the NewAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmountMinor

`func (o *SubscriptionPlanChange) SetNewAmountMinor(v int64)`

SetNewAmountMinor sets NewAmountMinor field to given value.


### GetProrationBehavior

`func (o *SubscriptionPlanChange) GetProrationBehavior() string`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *SubscriptionPlanChange) GetProrationBehaviorOk() (*string, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *SubscriptionPlanChange) SetProrationBehavior(v string)`

SetProrationBehavior sets ProrationBehavior field to given value.


### GetEffectiveDate

`func (o *SubscriptionPlanChange) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SubscriptionPlanChange) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SubscriptionPlanChange) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetProration

`func (o *SubscriptionPlanChange) GetProration() SubscriptionPlanChangeProration`

GetProration returns the Proration field if non-nil, zero value otherwise.

### GetProrationOk

`func (o *SubscriptionPlanChange) GetProrationOk() (*SubscriptionPlanChangeProration, bool)`

GetProrationOk returns a tuple with the Proration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProration

`func (o *SubscriptionPlanChange) SetProration(v SubscriptionPlanChangeProration)`

SetProration sets Proration field to given value.

### HasProration

`func (o *SubscriptionPlanChange) HasProration() bool`

HasProration returns a boolean if a field has been set.

### GetPendingProrationNetAmountMinor

`func (o *SubscriptionPlanChange) GetPendingProrationNetAmountMinor() int64`

GetPendingProrationNetAmountMinor returns the PendingProrationNetAmountMinor field if non-nil, zero value otherwise.

### GetPendingProrationNetAmountMinorOk

`func (o *SubscriptionPlanChange) GetPendingProrationNetAmountMinorOk() (*int64, bool)`

GetPendingProrationNetAmountMinorOk returns a tuple with the PendingProrationNetAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingProrationNetAmountMinor

`func (o *SubscriptionPlanChange) SetPendingProrationNetAmountMinor(v int64)`

SetPendingProrationNetAmountMinor sets PendingProrationNetAmountMinor field to given value.


### SetPendingProrationNetAmountMinorNil

`func (o *SubscriptionPlanChange) SetPendingProrationNetAmountMinorNil(b bool)`

 SetPendingProrationNetAmountMinorNil sets the value for PendingProrationNetAmountMinor to be an explicit nil

### UnsetPendingProrationNetAmountMinor
`func (o *SubscriptionPlanChange) UnsetPendingProrationNetAmountMinor()`

UnsetPendingProrationNetAmountMinor ensures that no value is present for PendingProrationNetAmountMinor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
