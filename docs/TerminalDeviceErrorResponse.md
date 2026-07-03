# TerminalDeviceErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewTerminalDeviceErrorResponse

`func NewTerminalDeviceErrorResponse(error_ string, ) *TerminalDeviceErrorResponse`

NewTerminalDeviceErrorResponse instantiates a new TerminalDeviceErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalDeviceErrorResponseWithDefaults

`func NewTerminalDeviceErrorResponseWithDefaults() *TerminalDeviceErrorResponse`

NewTerminalDeviceErrorResponseWithDefaults instantiates a new TerminalDeviceErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TerminalDeviceErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TerminalDeviceErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TerminalDeviceErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *TerminalDeviceErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TerminalDeviceErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TerminalDeviceErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TerminalDeviceErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


