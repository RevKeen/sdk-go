# PaymentIntentLastPaymentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code | 
**Message** | **string** | Human-readable error message | 
**DeclineCode** | Pointer to **string** | Card decline code if applicable | [optional] 
**PaymentMethod** | Pointer to [**PaymentIntentLastPaymentErrorPaymentMethod**](PaymentIntentLastPaymentErrorPaymentMethod.md) |  | [optional] 

## Methods

### NewPaymentIntentLastPaymentError

`func NewPaymentIntentLastPaymentError(code string, message string, ) *PaymentIntentLastPaymentError`

NewPaymentIntentLastPaymentError instantiates a new PaymentIntentLastPaymentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentLastPaymentErrorWithDefaults

`func NewPaymentIntentLastPaymentErrorWithDefaults() *PaymentIntentLastPaymentError`

NewPaymentIntentLastPaymentErrorWithDefaults instantiates a new PaymentIntentLastPaymentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PaymentIntentLastPaymentError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PaymentIntentLastPaymentError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PaymentIntentLastPaymentError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *PaymentIntentLastPaymentError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentIntentLastPaymentError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentIntentLastPaymentError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDeclineCode

`func (o *PaymentIntentLastPaymentError) GetDeclineCode() string`

GetDeclineCode returns the DeclineCode field if non-nil, zero value otherwise.

### GetDeclineCodeOk

`func (o *PaymentIntentLastPaymentError) GetDeclineCodeOk() (*string, bool)`

GetDeclineCodeOk returns a tuple with the DeclineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineCode

`func (o *PaymentIntentLastPaymentError) SetDeclineCode(v string)`

SetDeclineCode sets DeclineCode field to given value.

### HasDeclineCode

`func (o *PaymentIntentLastPaymentError) HasDeclineCode() bool`

HasDeclineCode returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentIntentLastPaymentError) GetPaymentMethod() PaymentIntentLastPaymentErrorPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentIntentLastPaymentError) GetPaymentMethodOk() (*PaymentIntentLastPaymentErrorPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentIntentLastPaymentError) SetPaymentMethod(v PaymentIntentLastPaymentErrorPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentIntentLastPaymentError) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


