# ConflictErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Code** | **string** |  |
**Message** | **string** |  |
**Param** | Pointer to **string** |  | [optional]
**RequestId** | Pointer to **string** |  | [optional]
**Details** | Pointer to **map[string]interface{}** | Additional conflict details | [optional]

## Methods

### NewConflictErrorError

`func NewConflictErrorError(type_ string, code string, message string, ) *ConflictErrorError`

NewConflictErrorError instantiates a new ConflictErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictErrorErrorWithDefaults

`func NewConflictErrorErrorWithDefaults() *ConflictErrorError`

NewConflictErrorErrorWithDefaults instantiates a new ConflictErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConflictErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConflictErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConflictErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ConflictErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConflictErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConflictErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ConflictErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConflictErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConflictErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *ConflictErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ConflictErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ConflictErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ConflictErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetRequestId

`func (o *ConflictErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ConflictErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ConflictErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ConflictErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDetails

`func (o *ConflictErrorError) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ConflictErrorError) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ConflictErrorError) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ConflictErrorError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
