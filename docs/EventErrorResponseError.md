# EventErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Error type | 
**Code** | **string** | Error code | 
**Message** | **string** | Human-readable error message | 
**Param** | Pointer to **string** | Parameter that caused the error | [optional] 

## Methods

### NewEventErrorResponseError

`func NewEventErrorResponseError(type_ string, code string, message string, ) *EventErrorResponseError`

NewEventErrorResponseError instantiates a new EventErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventErrorResponseErrorWithDefaults

`func NewEventErrorResponseErrorWithDefaults() *EventErrorResponseError`

NewEventErrorResponseErrorWithDefaults instantiates a new EventErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventErrorResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventErrorResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventErrorResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *EventErrorResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EventErrorResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EventErrorResponseError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *EventErrorResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventErrorResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventErrorResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *EventErrorResponseError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *EventErrorResponseError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *EventErrorResponseError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *EventErrorResponseError) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


