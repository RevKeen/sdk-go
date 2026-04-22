# PauseSubscriptionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Behavior** | Pointer to **string** | How to handle invoices during the pause period | [optional] [default to "keep_as_draft"]
**ResumesAt** | Pointer to **time.Time** | ISO 8601 date when the subscription should automatically resume. If not provided, subscription stays paused until manually resumed. | [optional] 

## Methods

### NewPauseSubscriptionInput

`func NewPauseSubscriptionInput() *PauseSubscriptionInput`

NewPauseSubscriptionInput instantiates a new PauseSubscriptionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseSubscriptionInputWithDefaults

`func NewPauseSubscriptionInputWithDefaults() *PauseSubscriptionInput`

NewPauseSubscriptionInputWithDefaults instantiates a new PauseSubscriptionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBehavior

`func (o *PauseSubscriptionInput) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *PauseSubscriptionInput) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *PauseSubscriptionInput) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *PauseSubscriptionInput) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetResumesAt

`func (o *PauseSubscriptionInput) GetResumesAt() time.Time`

GetResumesAt returns the ResumesAt field if non-nil, zero value otherwise.

### GetResumesAtOk

`func (o *PauseSubscriptionInput) GetResumesAtOk() (*time.Time, bool)`

GetResumesAtOk returns a tuple with the ResumesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesAt

`func (o *PauseSubscriptionInput) SetResumesAt(v time.Time)`

SetResumesAt sets ResumesAt field to given value.

### HasResumesAt

`func (o *PauseSubscriptionInput) HasResumesAt() bool`

HasResumesAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


