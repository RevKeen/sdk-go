# SubscriptionPlanChangeProration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAmountMinor** | **int64** | Credit for the unused time on the current plan, in minor units. |
**ChargeAmountMinor** | **int64** | Charge for the remaining time on the new plan, in minor units. |
**NetAmountMinor** | **int64** | &#x60;charge_amount_minor&#x60; minus &#x60;credit_amount_minor&#x60;. Positive is added to the next renewal invoice; negative is credited there. |
**IsUpgrade** | **bool** | Whether the new plan costs more than the current one. |
**DaysRemaining** | **int32** | Days left in the current period at the effective date. |
**DaysInPeriod** | **int32** | Days in the current period. |
**Currency** | **string** | Three-letter ISO currency code of the amounts. |

## Methods

### NewSubscriptionPlanChangeProration

`func NewSubscriptionPlanChangeProration(creditAmountMinor int64, chargeAmountMinor int64, netAmountMinor int64, isUpgrade bool, daysRemaining int32, daysInPeriod int32, currency string, ) *SubscriptionPlanChangeProration`

NewSubscriptionPlanChangeProration instantiates a new SubscriptionPlanChangeProration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanChangeProrationWithDefaults

`func NewSubscriptionPlanChangeProrationWithDefaults() *SubscriptionPlanChangeProration`

NewSubscriptionPlanChangeProrationWithDefaults instantiates a new SubscriptionPlanChangeProration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAmountMinor

`func (o *SubscriptionPlanChangeProration) GetCreditAmountMinor() int64`

GetCreditAmountMinor returns the CreditAmountMinor field if non-nil, zero value otherwise.

### GetCreditAmountMinorOk

`func (o *SubscriptionPlanChangeProration) GetCreditAmountMinorOk() (*int64, bool)`

GetCreditAmountMinorOk returns a tuple with the CreditAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountMinor

`func (o *SubscriptionPlanChangeProration) SetCreditAmountMinor(v int64)`

SetCreditAmountMinor sets CreditAmountMinor field to given value.


### GetChargeAmountMinor

`func (o *SubscriptionPlanChangeProration) GetChargeAmountMinor() int64`

GetChargeAmountMinor returns the ChargeAmountMinor field if non-nil, zero value otherwise.

### GetChargeAmountMinorOk

`func (o *SubscriptionPlanChangeProration) GetChargeAmountMinorOk() (*int64, bool)`

GetChargeAmountMinorOk returns a tuple with the ChargeAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmountMinor

`func (o *SubscriptionPlanChangeProration) SetChargeAmountMinor(v int64)`

SetChargeAmountMinor sets ChargeAmountMinor field to given value.


### GetNetAmountMinor

`func (o *SubscriptionPlanChangeProration) GetNetAmountMinor() int64`

GetNetAmountMinor returns the NetAmountMinor field if non-nil, zero value otherwise.

### GetNetAmountMinorOk

`func (o *SubscriptionPlanChangeProration) GetNetAmountMinorOk() (*int64, bool)`

GetNetAmountMinorOk returns a tuple with the NetAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmountMinor

`func (o *SubscriptionPlanChangeProration) SetNetAmountMinor(v int64)`

SetNetAmountMinor sets NetAmountMinor field to given value.


### GetIsUpgrade

`func (o *SubscriptionPlanChangeProration) GetIsUpgrade() bool`

GetIsUpgrade returns the IsUpgrade field if non-nil, zero value otherwise.

### GetIsUpgradeOk

`func (o *SubscriptionPlanChangeProration) GetIsUpgradeOk() (*bool, bool)`

GetIsUpgradeOk returns a tuple with the IsUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpgrade

`func (o *SubscriptionPlanChangeProration) SetIsUpgrade(v bool)`

SetIsUpgrade sets IsUpgrade field to given value.


### GetDaysRemaining

`func (o *SubscriptionPlanChangeProration) GetDaysRemaining() int32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *SubscriptionPlanChangeProration) GetDaysRemainingOk() (*int32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *SubscriptionPlanChangeProration) SetDaysRemaining(v int32)`

SetDaysRemaining sets DaysRemaining field to given value.


### GetDaysInPeriod

`func (o *SubscriptionPlanChangeProration) GetDaysInPeriod() int32`

GetDaysInPeriod returns the DaysInPeriod field if non-nil, zero value otherwise.

### GetDaysInPeriodOk

`func (o *SubscriptionPlanChangeProration) GetDaysInPeriodOk() (*int32, bool)`

GetDaysInPeriodOk returns a tuple with the DaysInPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInPeriod

`func (o *SubscriptionPlanChangeProration) SetDaysInPeriod(v int32)`

SetDaysInPeriod sets DaysInPeriod field to given value.


### GetCurrency

`func (o *SubscriptionPlanChangeProration) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionPlanChangeProration) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionPlanChangeProration) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
