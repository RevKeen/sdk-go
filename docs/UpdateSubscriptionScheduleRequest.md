# UpdateSubscriptionScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phases** | Pointer to [**[]SchedulePhase**](SchedulePhase.md) | Replace all phases | [optional] 
**EndBehavior** | Pointer to **string** | What happens when all phases complete | [optional] 
**ProrationBehavior** | Pointer to **string** | Proration behavior when updating phases | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata | [optional] 

## Methods

### NewUpdateSubscriptionScheduleRequest

`func NewUpdateSubscriptionScheduleRequest() *UpdateSubscriptionScheduleRequest`

NewUpdateSubscriptionScheduleRequest instantiates a new UpdateSubscriptionScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionScheduleRequestWithDefaults

`func NewUpdateSubscriptionScheduleRequestWithDefaults() *UpdateSubscriptionScheduleRequest`

NewUpdateSubscriptionScheduleRequestWithDefaults instantiates a new UpdateSubscriptionScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhases

`func (o *UpdateSubscriptionScheduleRequest) GetPhases() []SchedulePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *UpdateSubscriptionScheduleRequest) GetPhasesOk() (*[]SchedulePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *UpdateSubscriptionScheduleRequest) SetPhases(v []SchedulePhase)`

SetPhases sets Phases field to given value.

### HasPhases

`func (o *UpdateSubscriptionScheduleRequest) HasPhases() bool`

HasPhases returns a boolean if a field has been set.

### GetEndBehavior

`func (o *UpdateSubscriptionScheduleRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *UpdateSubscriptionScheduleRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *UpdateSubscriptionScheduleRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *UpdateSubscriptionScheduleRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetProrationBehavior

`func (o *UpdateSubscriptionScheduleRequest) GetProrationBehavior() string`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *UpdateSubscriptionScheduleRequest) GetProrationBehaviorOk() (*string, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *UpdateSubscriptionScheduleRequest) SetProrationBehavior(v string)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *UpdateSubscriptionScheduleRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSubscriptionScheduleRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSubscriptionScheduleRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSubscriptionScheduleRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSubscriptionScheduleRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


