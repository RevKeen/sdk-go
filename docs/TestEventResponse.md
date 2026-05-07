# TestEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the test event was created | 
**Event** | Pointer to [**Event**](Event.md) | The created test event | [optional] 
**Message** | Pointer to **string** | Additional information | [optional] 

## Methods

### NewTestEventResponse

`func NewTestEventResponse(success bool, ) *TestEventResponse`

NewTestEventResponse instantiates a new TestEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEventResponseWithDefaults

`func NewTestEventResponseWithDefaults() *TestEventResponse`

NewTestEventResponseWithDefaults instantiates a new TestEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *TestEventResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TestEventResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TestEventResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetEvent

`func (o *TestEventResponse) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TestEventResponse) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TestEventResponse) SetEvent(v Event)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TestEventResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetMessage

`func (o *TestEventResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestEventResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestEventResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestEventResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


