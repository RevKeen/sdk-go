# SubscriptionSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Public schedule ID (sub_sched_xxx) | 
**Object** | **string** | Object type | 
**Status** | **string** | Schedule status. not_started: Scheduled for future. active: Currently executing phases. completed: All phases finished. canceled: Manually canceled. released: Detached from subscription. | 
**Customer** | **NullableString** | Customer ID | 
**Subscription** | **NullableString** | Linked subscription ID | 
**CurrentPhase** | [**SubscriptionScheduleCurrentPhase**](SubscriptionScheduleCurrentPhase.md) |  | 
**Phases** | [**[]SchedulePhase**](SchedulePhase.md) | All phases in the schedule | 
**EndBehavior** | **string** | What happens when the schedule completes. cancel: Cancel the subscription. release: Make subscription standalone. | 
**ReleasedAt** | Pointer to **NullableInt32** | When released (Unix timestamp) | [optional] 
**ReleasedSubscription** | Pointer to **NullableString** | Subscription ID at release time | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata | [optional] 
**Created** | **int32** | Unix timestamp of creation | 
**Livemode** | **bool** | Whether in live mode | 

## Methods

### NewSubscriptionSchedule

`func NewSubscriptionSchedule(id string, object string, status string, customer NullableString, subscription NullableString, currentPhase SubscriptionScheduleCurrentPhase, phases []SchedulePhase, endBehavior string, created int32, livemode bool, ) *SubscriptionSchedule`

NewSubscriptionSchedule instantiates a new SubscriptionSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionScheduleWithDefaults

`func NewSubscriptionScheduleWithDefaults() *SubscriptionSchedule`

NewSubscriptionScheduleWithDefaults instantiates a new SubscriptionSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionSchedule) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *SubscriptionSchedule) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionSchedule) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionSchedule) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *SubscriptionSchedule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionSchedule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionSchedule) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomer

`func (o *SubscriptionSchedule) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *SubscriptionSchedule) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *SubscriptionSchedule) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### SetCustomerNil

`func (o *SubscriptionSchedule) SetCustomerNil(b bool)`

 SetCustomerNil sets the value for Customer to be an explicit nil

### UnsetCustomer
`func (o *SubscriptionSchedule) UnsetCustomer()`

UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
### GetSubscription

`func (o *SubscriptionSchedule) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionSchedule) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionSchedule) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### SetSubscriptionNil

`func (o *SubscriptionSchedule) SetSubscriptionNil(b bool)`

 SetSubscriptionNil sets the value for Subscription to be an explicit nil

### UnsetSubscription
`func (o *SubscriptionSchedule) UnsetSubscription()`

UnsetSubscription ensures that no value is present for Subscription, not even an explicit nil
### GetCurrentPhase

`func (o *SubscriptionSchedule) GetCurrentPhase() SubscriptionScheduleCurrentPhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *SubscriptionSchedule) GetCurrentPhaseOk() (*SubscriptionScheduleCurrentPhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *SubscriptionSchedule) SetCurrentPhase(v SubscriptionScheduleCurrentPhase)`

SetCurrentPhase sets CurrentPhase field to given value.


### GetPhases

`func (o *SubscriptionSchedule) GetPhases() []SchedulePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *SubscriptionSchedule) GetPhasesOk() (*[]SchedulePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *SubscriptionSchedule) SetPhases(v []SchedulePhase)`

SetPhases sets Phases field to given value.


### GetEndBehavior

`func (o *SubscriptionSchedule) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *SubscriptionSchedule) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *SubscriptionSchedule) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.


### GetReleasedAt

`func (o *SubscriptionSchedule) GetReleasedAt() int32`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *SubscriptionSchedule) GetReleasedAtOk() (*int32, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *SubscriptionSchedule) SetReleasedAt(v int32)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *SubscriptionSchedule) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### SetReleasedAtNil

`func (o *SubscriptionSchedule) SetReleasedAtNil(b bool)`

 SetReleasedAtNil sets the value for ReleasedAt to be an explicit nil

### UnsetReleasedAt
`func (o *SubscriptionSchedule) UnsetReleasedAt()`

UnsetReleasedAt ensures that no value is present for ReleasedAt, not even an explicit nil
### GetReleasedSubscription

`func (o *SubscriptionSchedule) GetReleasedSubscription() string`

GetReleasedSubscription returns the ReleasedSubscription field if non-nil, zero value otherwise.

### GetReleasedSubscriptionOk

`func (o *SubscriptionSchedule) GetReleasedSubscriptionOk() (*string, bool)`

GetReleasedSubscriptionOk returns a tuple with the ReleasedSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSubscription

`func (o *SubscriptionSchedule) SetReleasedSubscription(v string)`

SetReleasedSubscription sets ReleasedSubscription field to given value.

### HasReleasedSubscription

`func (o *SubscriptionSchedule) HasReleasedSubscription() bool`

HasReleasedSubscription returns a boolean if a field has been set.

### SetReleasedSubscriptionNil

`func (o *SubscriptionSchedule) SetReleasedSubscriptionNil(b bool)`

 SetReleasedSubscriptionNil sets the value for ReleasedSubscription to be an explicit nil

### UnsetReleasedSubscription
`func (o *SubscriptionSchedule) UnsetReleasedSubscription()`

UnsetReleasedSubscription ensures that no value is present for ReleasedSubscription, not even an explicit nil
### GetMetadata

`func (o *SubscriptionSchedule) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionSchedule) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionSchedule) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionSchedule) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreated

`func (o *SubscriptionSchedule) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubscriptionSchedule) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubscriptionSchedule) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetLivemode

`func (o *SubscriptionSchedule) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *SubscriptionSchedule) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *SubscriptionSchedule) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


