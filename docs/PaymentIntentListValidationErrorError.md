# PaymentIntentListValidationErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Code** | **string** |  |
**Message** | **string** |  |
**RequestId** | Pointer to **string** |  | [optional]
**Param** | **string** | The query parameter that was rejected. |

## Methods

### NewPaymentIntentListValidationErrorError

`func NewPaymentIntentListValidationErrorError(type_ string, code string, message string, param string, ) *PaymentIntentListValidationErrorError`

NewPaymentIntentListValidationErrorError instantiates a new PaymentIntentListValidationErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentListValidationErrorErrorWithDefaults

`func NewPaymentIntentListValidationErrorErrorWithDefaults() *PaymentIntentListValidationErrorError`

NewPaymentIntentListValidationErrorErrorWithDefaults instantiates a new PaymentIntentListValidationErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentIntentListValidationErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentIntentListValidationErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentIntentListValidationErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *PaymentIntentListValidationErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentIntentListValidationErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentIntentListValidationErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *PaymentIntentListValidationErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentIntentListValidationErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentIntentListValidationErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRequestId

`func (o *PaymentIntentListValidationErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentIntentListValidationErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentIntentListValidationErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentIntentListValidationErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetParam

`func (o *PaymentIntentListValidationErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *PaymentIntentListValidationErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *PaymentIntentListValidationErrorError) SetParam(v string)`

SetParam sets Param field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
