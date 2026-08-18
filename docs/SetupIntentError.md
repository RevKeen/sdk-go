# SetupIntentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code | 
**Message** | **string** | Human-readable error message | 
**DeclineCode** | Pointer to **string** | Decline code from the payment gateway | [optional] 
**PaymentMethod** | Pointer to [**SetupIntentErrorPaymentMethod**](SetupIntentErrorPaymentMethod.md) |  | [optional] 

## Methods

### NewSetupIntentError

`func NewSetupIntentError(code string, message string, ) *SetupIntentError`

NewSetupIntentError instantiates a new SetupIntentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupIntentErrorWithDefaults

`func NewSetupIntentErrorWithDefaults() *SetupIntentError`

NewSetupIntentErrorWithDefaults instantiates a new SetupIntentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SetupIntentError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SetupIntentError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SetupIntentError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *SetupIntentError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SetupIntentError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SetupIntentError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDeclineCode

`func (o *SetupIntentError) GetDeclineCode() string`

GetDeclineCode returns the DeclineCode field if non-nil, zero value otherwise.

### GetDeclineCodeOk

`func (o *SetupIntentError) GetDeclineCodeOk() (*string, bool)`

GetDeclineCodeOk returns a tuple with the DeclineCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineCode

`func (o *SetupIntentError) SetDeclineCode(v string)`

SetDeclineCode sets DeclineCode field to given value.

### HasDeclineCode

`func (o *SetupIntentError) HasDeclineCode() bool`

HasDeclineCode returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *SetupIntentError) GetPaymentMethod() SetupIntentErrorPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *SetupIntentError) GetPaymentMethodOk() (*SetupIntentErrorPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *SetupIntentError) SetPaymentMethod(v SetupIntentErrorPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *SetupIntentError) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


