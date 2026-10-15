# ApiErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of error returned |
**Code** | **string** | Machine-readable error code for programmatic handling |
**Message** | **string** | Human-readable error message |
**Param** | Pointer to **string** | The parameter that caused the error | [optional]
**StatusCode** | Pointer to **int32** | HTTP status code | [optional]
**RequestId** | Pointer to **string** | Unique request identifier for debugging | [optional]
**Details** | Pointer to **map[string]interface{}** | Additional context about the error | [optional]
**DeclineCode** | Pointer to **string** | Card decline code from the payment processor | [optional]
**DocUrl** | Pointer to **string** | URL to documentation about this error | [optional]

## Methods

### NewApiErrorError

`func NewApiErrorError(type_ string, code string, message string, ) *ApiErrorError`

NewApiErrorError instantiates a new ApiErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorErrorWithDefaults

`func NewApiErrorErrorWithDefaults() *ApiErrorError`

NewApiErrorErrorWithDefaults instantiates a new ApiErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *ApiErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ApiErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *ApiErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ApiErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ApiErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ApiErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetStatusCode

`func (o *ApiErrorError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApiErrorError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApiErrorError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApiErrorError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetRequestId

`func (o *ApiErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApiErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApiErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ApiErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDetails

`func (o *ApiErrorError) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiErrorError) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ApiErrorError) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ApiErrorError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDeclineCode

`func (o *ApiErrorError) GetDeclineCode() string`

GetDeclineCode returns the DeclineCode field if non-nil, zero value otherwise.

### GetDeclineCodeOk

`func (o *ApiErrorError) GetDeclineCodeOk() (*string, bool)`

GetDeclineCodeOk returns a tuple with the DeclineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineCode

`func (o *ApiErrorError) SetDeclineCode(v string)`

SetDeclineCode sets DeclineCode field to given value.

### HasDeclineCode

`func (o *ApiErrorError) HasDeclineCode() bool`

HasDeclineCode returns a boolean if a field has been set.

### GetDocUrl

`func (o *ApiErrorError) GetDocUrl() string`

GetDocUrl returns the DocUrl field if non-nil, zero value otherwise.

### GetDocUrlOk

`func (o *ApiErrorError) GetDocUrlOk() (*string, bool)`

GetDocUrlOk returns a tuple with the DocUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocUrl

`func (o *ApiErrorError) SetDocUrl(v string)`

SetDocUrl sets DocUrl field to given value.

### HasDocUrl

`func (o *ApiErrorError) HasDocUrl() bool`

HasDocUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
