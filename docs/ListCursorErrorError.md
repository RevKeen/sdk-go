# ListCursorErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Code** | **string** | &#x60;parameter_invalid&#x60; when both cursors are supplied; &#x60;resource_missing&#x60; when the cursor names no object of the caller&#39;s merchant. |
**Message** | **string** |  |
**Param** | **string** | The cursor parameter the error is about. |
**RequestId** | Pointer to **string** |  | [optional]

## Methods

### NewListCursorErrorError

`func NewListCursorErrorError(type_ string, code string, message string, param string, ) *ListCursorErrorError`

NewListCursorErrorError instantiates a new ListCursorErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCursorErrorErrorWithDefaults

`func NewListCursorErrorErrorWithDefaults() *ListCursorErrorError`

NewListCursorErrorErrorWithDefaults instantiates a new ListCursorErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListCursorErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCursorErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCursorErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ListCursorErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCursorErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCursorErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ListCursorErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListCursorErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListCursorErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *ListCursorErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ListCursorErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ListCursorErrorError) SetParam(v string)`

SetParam sets Param field to given value.


### GetRequestId

`func (o *ListCursorErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListCursorErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListCursorErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListCursorErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
