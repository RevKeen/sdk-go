# PaymentIntentErrorResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Error type (api_error, invalid_request_error, etc.) | 
**Code** | **string** | Error code | 
**Message** | **string** | Human-readable error message | 
**Param** | Pointer to **string** | Parameter that caused the error | [optional] 

## Methods

### NewPaymentIntentErrorResponseError

`func NewPaymentIntentErrorResponseError(type_ string, code string, message string, ) *PaymentIntentErrorResponseError`

NewPaymentIntentErrorResponseError instantiates a new PaymentIntentErrorResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentErrorResponseErrorWithDefaults

`func NewPaymentIntentErrorResponseErrorWithDefaults() *PaymentIntentErrorResponseError`

NewPaymentIntentErrorResponseErrorWithDefaults instantiates a new PaymentIntentErrorResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentIntentErrorResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentIntentErrorResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentIntentErrorResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *PaymentIntentErrorResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentIntentErrorResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentIntentErrorResponseError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *PaymentIntentErrorResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentIntentErrorResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentIntentErrorResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *PaymentIntentErrorResponseError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *PaymentIntentErrorResponseError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *PaymentIntentErrorResponseError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *PaymentIntentErrorResponseError) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


