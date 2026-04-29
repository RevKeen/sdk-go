# BillingPreviewResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner**](SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner.md) |  | 
**Summary** | **string** |  | 
**TotalAmountMinor** | **int32** |  | 
**FirstChargeDate** | **time.Time** |  | 
**TrialEndDate** | **NullableTime** |  | 
**ScheduleEndDate** | **NullableTime** |  | 
**Timezone** | **string** |  | 

## Methods

### NewBillingPreviewResponseData

`func NewBillingPreviewResponseData(items []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, summary string, totalAmountMinor int32, firstChargeDate time.Time, trialEndDate NullableTime, scheduleEndDate NullableTime, timezone string, ) *BillingPreviewResponseData`

NewBillingPreviewResponseData instantiates a new BillingPreviewResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPreviewResponseDataWithDefaults

`func NewBillingPreviewResponseDataWithDefaults() *BillingPreviewResponseData`

NewBillingPreviewResponseDataWithDefaults instantiates a new BillingPreviewResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BillingPreviewResponseData) GetItems() []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BillingPreviewResponseData) GetItemsOk() (*[]SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BillingPreviewResponseData) SetItems(v []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner)`

SetItems sets Items field to given value.


### GetSummary

`func (o *BillingPreviewResponseData) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BillingPreviewResponseData) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BillingPreviewResponseData) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTotalAmountMinor

`func (o *BillingPreviewResponseData) GetTotalAmountMinor() int32`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *BillingPreviewResponseData) GetTotalAmountMinorOk() (*int32, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *BillingPreviewResponseData) SetTotalAmountMinor(v int32)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetFirstChargeDate

`func (o *BillingPreviewResponseData) GetFirstChargeDate() time.Time`

GetFirstChargeDate returns the FirstChargeDate field if non-nil, zero value otherwise.

### GetFirstChargeDateOk

`func (o *BillingPreviewResponseData) GetFirstChargeDateOk() (*time.Time, bool)`

GetFirstChargeDateOk returns a tuple with the FirstChargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeDate

`func (o *BillingPreviewResponseData) SetFirstChargeDate(v time.Time)`

SetFirstChargeDate sets FirstChargeDate field to given value.


### GetTrialEndDate

`func (o *BillingPreviewResponseData) GetTrialEndDate() time.Time`

GetTrialEndDate returns the TrialEndDate field if non-nil, zero value otherwise.

### GetTrialEndDateOk

`func (o *BillingPreviewResponseData) GetTrialEndDateOk() (*time.Time, bool)`

GetTrialEndDateOk returns a tuple with the TrialEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEndDate

`func (o *BillingPreviewResponseData) SetTrialEndDate(v time.Time)`

SetTrialEndDate sets TrialEndDate field to given value.


### SetTrialEndDateNil

`func (o *BillingPreviewResponseData) SetTrialEndDateNil(b bool)`

 SetTrialEndDateNil sets the value for TrialEndDate to be an explicit nil

### UnsetTrialEndDate
`func (o *BillingPreviewResponseData) UnsetTrialEndDate()`

UnsetTrialEndDate ensures that no value is present for TrialEndDate, not even an explicit nil
### GetScheduleEndDate

`func (o *BillingPreviewResponseData) GetScheduleEndDate() time.Time`

GetScheduleEndDate returns the ScheduleEndDate field if non-nil, zero value otherwise.

### GetScheduleEndDateOk

`func (o *BillingPreviewResponseData) GetScheduleEndDateOk() (*time.Time, bool)`

GetScheduleEndDateOk returns a tuple with the ScheduleEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEndDate

`func (o *BillingPreviewResponseData) SetScheduleEndDate(v time.Time)`

SetScheduleEndDate sets ScheduleEndDate field to given value.


### SetScheduleEndDateNil

`func (o *BillingPreviewResponseData) SetScheduleEndDateNil(b bool)`

 SetScheduleEndDateNil sets the value for ScheduleEndDate to be an explicit nil

### UnsetScheduleEndDate
`func (o *BillingPreviewResponseData) UnsetScheduleEndDate()`

UnsetScheduleEndDate ensures that no value is present for ScheduleEndDate, not even an explicit nil
### GetTimezone

`func (o *BillingPreviewResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BillingPreviewResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BillingPreviewResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


