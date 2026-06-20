# SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequence** | **int32** |  | 
**Date** | **time.Time** |  | 
**FormattedDate** | **string** |  | 
**DayName** | **string** |  | 
**AmountMinor** | **int32** |  | 
**FormattedAmount** | **string** |  | 
**IsTrialEnd** | Pointer to **bool** |  | [optional] 
**IsFirstCharge** | Pointer to **bool** |  | [optional] 
**IsFinalPayment** | Pointer to **bool** |  | [optional] 

## Methods

### NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner

`func NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner(sequence int32, date time.Time, formattedDate string, dayName string, amountMinor int32, formattedAmount string, ) *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner`

NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner instantiates a new SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInnerWithDefaults

`func NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInnerWithDefaults() *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner`

NewSubscriptionPreviewRenewalResponseDataUpcomingRenewalsInnerWithDefaults instantiates a new SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequence

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetDate

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetFormattedDate

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetFormattedDate() string`

GetFormattedDate returns the FormattedDate field if non-nil, zero value otherwise.

### GetFormattedDateOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetFormattedDateOk() (*string, bool)`

GetFormattedDateOk returns a tuple with the FormattedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDate

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetFormattedDate(v string)`

SetFormattedDate sets FormattedDate field to given value.


### GetDayName

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetDayName() string`

GetDayName returns the DayName field if non-nil, zero value otherwise.

### GetDayNameOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetDayNameOk() (*string, bool)`

GetDayNameOk returns a tuple with the DayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayName

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetDayName(v string)`

SetDayName sets DayName field to given value.


### GetAmountMinor

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetFormattedAmount

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetFormattedAmount() string`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetFormattedAmountOk() (*string, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetFormattedAmount(v string)`

SetFormattedAmount sets FormattedAmount field to given value.


### GetIsTrialEnd

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsTrialEnd() bool`

GetIsTrialEnd returns the IsTrialEnd field if non-nil, zero value otherwise.

### GetIsTrialEndOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsTrialEndOk() (*bool, bool)`

GetIsTrialEndOk returns a tuple with the IsTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrialEnd

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetIsTrialEnd(v bool)`

SetIsTrialEnd sets IsTrialEnd field to given value.

### HasIsTrialEnd

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) HasIsTrialEnd() bool`

HasIsTrialEnd returns a boolean if a field has been set.

### GetIsFirstCharge

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsFirstCharge() bool`

GetIsFirstCharge returns the IsFirstCharge field if non-nil, zero value otherwise.

### GetIsFirstChargeOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsFirstChargeOk() (*bool, bool)`

GetIsFirstChargeOk returns a tuple with the IsFirstCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstCharge

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetIsFirstCharge(v bool)`

SetIsFirstCharge sets IsFirstCharge field to given value.

### HasIsFirstCharge

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) HasIsFirstCharge() bool`

HasIsFirstCharge returns a boolean if a field has been set.

### GetIsFinalPayment

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsFinalPayment() bool`

GetIsFinalPayment returns the IsFinalPayment field if non-nil, zero value otherwise.

### GetIsFinalPaymentOk

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) GetIsFinalPaymentOk() (*bool, bool)`

GetIsFinalPaymentOk returns a tuple with the IsFinalPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinalPayment

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) SetIsFinalPayment(v bool)`

SetIsFinalPayment sets IsFinalPayment field to given value.

### HasIsFinalPayment

`func (o *SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner) HasIsFinalPayment() bool`

HasIsFinalPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


