# RateLimitErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Code** | **string** |  |
**Message** | **string** |  |
**RequestId** | Pointer to **string** |  | [optional]
**Details** | Pointer to [**RateLimitErrorErrorDetails**](RateLimitErrorErrorDetails.md) |  | [optional]

## Methods

### NewRateLimitErrorError

`func NewRateLimitErrorError(type_ string, code string, message string, ) *RateLimitErrorError`

NewRateLimitErrorError instantiates a new RateLimitErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitErrorErrorWithDefaults

`func NewRateLimitErrorErrorWithDefaults() *RateLimitErrorError`

NewRateLimitErrorErrorWithDefaults instantiates a new RateLimitErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RateLimitErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RateLimitErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RateLimitErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *RateLimitErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RateLimitErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RateLimitErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *RateLimitErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RateLimitErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RateLimitErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestId

`func (o *RateLimitErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RateLimitErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RateLimitErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RateLimitErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDetails

`func (o *RateLimitErrorError) GetDetails() RateLimitErrorErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RateLimitErrorError) GetDetailsOk() (*RateLimitErrorErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RateLimitErrorError) SetDetails(v RateLimitErrorErrorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RateLimitErrorError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
