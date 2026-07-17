# TerminalPaymentErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewTerminalPaymentErrorResponse

`func NewTerminalPaymentErrorResponse(error_ string, ) *TerminalPaymentErrorResponse`

NewTerminalPaymentErrorResponse instantiates a new TerminalPaymentErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalPaymentErrorResponseWithDefaults

`func NewTerminalPaymentErrorResponseWithDefaults() *TerminalPaymentErrorResponse`

NewTerminalPaymentErrorResponseWithDefaults instantiates a new TerminalPaymentErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *TerminalPaymentErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TerminalPaymentErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TerminalPaymentErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetMessage

`func (o *TerminalPaymentErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TerminalPaymentErrorResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TerminalPaymentErrorResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TerminalPaymentErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *TerminalPaymentErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TerminalPaymentErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TerminalPaymentErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TerminalPaymentErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


