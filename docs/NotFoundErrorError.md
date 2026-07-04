# NotFoundErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Code** | **string** |  | 
**Message** | **string** |  | 
**Param** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotFoundErrorError

`func NewNotFoundErrorError(type_ string, code string, message string, ) *NotFoundErrorError`

NewNotFoundErrorError instantiates a new NotFoundErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorErrorWithDefaults

`func NewNotFoundErrorErrorWithDefaults() *NotFoundErrorError`

NewNotFoundErrorErrorWithDefaults instantiates a new NotFoundErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotFoundErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotFoundErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotFoundErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *NotFoundErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *NotFoundErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotFoundErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotFoundErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *NotFoundErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *NotFoundErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *NotFoundErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *NotFoundErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetRequestId

`func (o *NotFoundErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *NotFoundErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *NotFoundErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *NotFoundErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


