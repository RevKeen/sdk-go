# CreateRefundInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The ID of the payment to refund | 
**AmountMinor** | Pointer to **float32** | Amount to refund in cents. If not provided, refunds the full remaining amount. | [optional] 
**Reason** | Pointer to **string** | Reason for the refund | [optional] 
**ReasonDetails** | Pointer to **string** | Additional details about the refund reason | [optional] 

## Methods

### NewCreateRefundInput

`func NewCreateRefundInput(paymentId string, ) *CreateRefundInput`

NewCreateRefundInput instantiates a new CreateRefundInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundInputWithDefaults

`func NewCreateRefundInputWithDefaults() *CreateRefundInput`

NewCreateRefundInputWithDefaults instantiates a new CreateRefundInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *CreateRefundInput) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateRefundInput) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateRefundInput) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetAmountMinor

`func (o *CreateRefundInput) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateRefundInput) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateRefundInput) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *CreateRefundInput) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetReason

`func (o *CreateRefundInput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateRefundInput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateRefundInput) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateRefundInput) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonDetails

`func (o *CreateRefundInput) GetReasonDetails() string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *CreateRefundInput) GetReasonDetailsOk() (*string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *CreateRefundInput) SetReasonDetails(v string)`

SetReasonDetails sets ReasonDetails field to given value.

### HasReasonDetails

`func (o *CreateRefundInput) HasReasonDetails() bool`

HasReasonDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


