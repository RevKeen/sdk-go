# SubscriptionScheduleCurrentPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **int32** | Phase start (Unix timestamp) | 
**EndDate** | **NullableInt32** | Phase end (Unix timestamp) | 

## Methods

### NewSubscriptionScheduleCurrentPhase

`func NewSubscriptionScheduleCurrentPhase(startDate int32, endDate NullableInt32, ) *SubscriptionScheduleCurrentPhase`

NewSubscriptionScheduleCurrentPhase instantiates a new SubscriptionScheduleCurrentPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionScheduleCurrentPhaseWithDefaults

`func NewSubscriptionScheduleCurrentPhaseWithDefaults() *SubscriptionScheduleCurrentPhase`

NewSubscriptionScheduleCurrentPhaseWithDefaults instantiates a new SubscriptionScheduleCurrentPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *SubscriptionScheduleCurrentPhase) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionScheduleCurrentPhase) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionScheduleCurrentPhase) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *SubscriptionScheduleCurrentPhase) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionScheduleCurrentPhase) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionScheduleCurrentPhase) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *SubscriptionScheduleCurrentPhase) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *SubscriptionScheduleCurrentPhase) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


