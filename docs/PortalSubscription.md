# PortalSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Status** | **string** |  | 
**Currency** | **NullableString** |  | 
**AmountMinor** | **NullableInt32** |  | 
**Interval** | **NullableString** |  | 
**IntervalCount** | **NullableInt32** |  | 
**CurrentPeriodStart** | **NullableTime** |  | 
**CurrentPeriodEnd** | **NullableTime** |  | 
**TrialEnd** | **NullableTime** |  | 
**CancelAt** | **NullableTime** |  | 
**CanceledAt** | **NullableTime** |  | 
**StartedAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPortalSubscription

`func NewPortalSubscription(id string, object string, status string, currency NullableString, amountMinor NullableInt32, interval NullableString, intervalCount NullableInt32, currentPeriodStart NullableTime, currentPeriodEnd NullableTime, trialEnd NullableTime, cancelAt NullableTime, canceledAt NullableTime, startedAt NullableTime, createdAt time.Time, ) *PortalSubscription`

NewPortalSubscription instantiates a new PortalSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSubscriptionWithDefaults

`func NewPortalSubscriptionWithDefaults() *PortalSubscription`

NewPortalSubscriptionWithDefaults instantiates a new PortalSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalSubscription) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PortalSubscription) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalSubscription) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalSubscription) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *PortalSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortalSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortalSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *PortalSubscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PortalSubscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PortalSubscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *PortalSubscription) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PortalSubscription) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAmountMinor

`func (o *PortalSubscription) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *PortalSubscription) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *PortalSubscription) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *PortalSubscription) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *PortalSubscription) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetInterval

`func (o *PortalSubscription) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PortalSubscription) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PortalSubscription) SetInterval(v string)`

SetInterval sets Interval field to given value.


### SetIntervalNil

`func (o *PortalSubscription) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *PortalSubscription) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *PortalSubscription) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *PortalSubscription) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *PortalSubscription) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.


### SetIntervalCountNil

`func (o *PortalSubscription) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *PortalSubscription) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetCurrentPeriodStart

`func (o *PortalSubscription) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *PortalSubscription) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *PortalSubscription) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### SetCurrentPeriodStartNil

`func (o *PortalSubscription) SetCurrentPeriodStartNil(b bool)`

 SetCurrentPeriodStartNil sets the value for CurrentPeriodStart to be an explicit nil

### UnsetCurrentPeriodStart
`func (o *PortalSubscription) UnsetCurrentPeriodStart()`

UnsetCurrentPeriodStart ensures that no value is present for CurrentPeriodStart, not even an explicit nil
### GetCurrentPeriodEnd

`func (o *PortalSubscription) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *PortalSubscription) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *PortalSubscription) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### SetCurrentPeriodEndNil

`func (o *PortalSubscription) SetCurrentPeriodEndNil(b bool)`

 SetCurrentPeriodEndNil sets the value for CurrentPeriodEnd to be an explicit nil

### UnsetCurrentPeriodEnd
`func (o *PortalSubscription) UnsetCurrentPeriodEnd()`

UnsetCurrentPeriodEnd ensures that no value is present for CurrentPeriodEnd, not even an explicit nil
### GetTrialEnd

`func (o *PortalSubscription) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *PortalSubscription) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *PortalSubscription) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.


### SetTrialEndNil

`func (o *PortalSubscription) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *PortalSubscription) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetCancelAt

`func (o *PortalSubscription) GetCancelAt() time.Time`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *PortalSubscription) GetCancelAtOk() (*time.Time, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *PortalSubscription) SetCancelAt(v time.Time)`

SetCancelAt sets CancelAt field to given value.


### SetCancelAtNil

`func (o *PortalSubscription) SetCancelAtNil(b bool)`

 SetCancelAtNil sets the value for CancelAt to be an explicit nil

### UnsetCancelAt
`func (o *PortalSubscription) UnsetCancelAt()`

UnsetCancelAt ensures that no value is present for CancelAt, not even an explicit nil
### GetCanceledAt

`func (o *PortalSubscription) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *PortalSubscription) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *PortalSubscription) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.


### SetCanceledAtNil

`func (o *PortalSubscription) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *PortalSubscription) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetStartedAt

`func (o *PortalSubscription) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *PortalSubscription) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *PortalSubscription) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *PortalSubscription) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *PortalSubscription) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCreatedAt

`func (o *PortalSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


