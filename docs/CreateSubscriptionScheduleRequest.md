# CreateSubscriptionScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to **string** | Customer ID (required if no subscription) | [optional] 
**Subscription** | Pointer to **string** | Existing subscription to attach schedule to | [optional] 
**StartDate** | Pointer to [**CreateSubscriptionScheduleRequestStartDate**](CreateSubscriptionScheduleRequestStartDate.md) |  | [optional] 
**Phases** | [**[]SchedulePhase**](SchedulePhase.md) | Phases defining the schedule (at least one required) | 
**EndBehavior** | Pointer to **string** | What happens when all phases complete | [optional] [default to "cancel"]
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata | [optional] 

## Methods

### NewCreateSubscriptionScheduleRequest

`func NewCreateSubscriptionScheduleRequest(phases []SchedulePhase, ) *CreateSubscriptionScheduleRequest`

NewCreateSubscriptionScheduleRequest instantiates a new CreateSubscriptionScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionScheduleRequestWithDefaults

`func NewCreateSubscriptionScheduleRequestWithDefaults() *CreateSubscriptionScheduleRequest`

NewCreateSubscriptionScheduleRequestWithDefaults instantiates a new CreateSubscriptionScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CreateSubscriptionScheduleRequest) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreateSubscriptionScheduleRequest) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreateSubscriptionScheduleRequest) SetCustomer(v string)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CreateSubscriptionScheduleRequest) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSubscription

`func (o *CreateSubscriptionScheduleRequest) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CreateSubscriptionScheduleRequest) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CreateSubscriptionScheduleRequest) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *CreateSubscriptionScheduleRequest) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateSubscriptionScheduleRequest) GetStartDate() CreateSubscriptionScheduleRequestStartDate`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateSubscriptionScheduleRequest) GetStartDateOk() (*CreateSubscriptionScheduleRequestStartDate, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateSubscriptionScheduleRequest) SetStartDate(v CreateSubscriptionScheduleRequestStartDate)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateSubscriptionScheduleRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetPhases

`func (o *CreateSubscriptionScheduleRequest) GetPhases() []SchedulePhase`

GetPhases returns the Phases field if non-nil, zero value otherwise.

### GetPhasesOk

`func (o *CreateSubscriptionScheduleRequest) GetPhasesOk() (*[]SchedulePhase, bool)`

GetPhasesOk returns a tuple with the Phases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhases

`func (o *CreateSubscriptionScheduleRequest) SetPhases(v []SchedulePhase)`

SetPhases sets Phases field to given value.


### GetEndBehavior

`func (o *CreateSubscriptionScheduleRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *CreateSubscriptionScheduleRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *CreateSubscriptionScheduleRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *CreateSubscriptionScheduleRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSubscriptionScheduleRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubscriptionScheduleRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubscriptionScheduleRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSubscriptionScheduleRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


