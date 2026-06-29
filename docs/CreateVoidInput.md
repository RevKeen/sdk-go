# CreateVoidInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The ID of the payment to void. Must be unsettled. | 
**Reason** | Pointer to **string** | Reason for the void | [optional] 
**ReasonCode** | Pointer to **string** | Standardized reason code for the void | [optional] 

## Methods

### NewCreateVoidInput

`func NewCreateVoidInput(paymentId string, ) *CreateVoidInput`

NewCreateVoidInput instantiates a new CreateVoidInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVoidInputWithDefaults

`func NewCreateVoidInputWithDefaults() *CreateVoidInput`

NewCreateVoidInputWithDefaults instantiates a new CreateVoidInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *CreateVoidInput) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateVoidInput) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateVoidInput) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetReason

`func (o *CreateVoidInput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateVoidInput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateVoidInput) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateVoidInput) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *CreateVoidInput) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CreateVoidInput) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CreateVoidInput) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CreateVoidInput) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


