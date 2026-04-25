# CreditNoteEligibilityResponsePaymentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** |  | 
**Gateway** | **string** |  | 
**PaymentMethod** | **NullableString** |  | 
**AmountMinor** | **float32** |  | 
**AmountRefundedMinor** | **float32** |  | 
**AmountAvailableMinor** | **float32** |  | 
**CardBrand** | **NullableString** |  | 
**CardLast4** | **NullableString** |  | 
**TerminalSerial** | **NullableString** |  | 
**AvailableOperations** | **[]string** |  | 
**Constraints** | [**CreditNoteEligibilityResponsePaymentsInnerConstraints**](CreditNoteEligibilityResponsePaymentsInnerConstraints.md) |  | 

## Methods

### NewCreditNoteEligibilityResponsePaymentsInner

`func NewCreditNoteEligibilityResponsePaymentsInner(transactionId string, gateway string, paymentMethod NullableString, amountMinor float32, amountRefundedMinor float32, amountAvailableMinor float32, cardBrand NullableString, cardLast4 NullableString, terminalSerial NullableString, availableOperations []string, constraints CreditNoteEligibilityResponsePaymentsInnerConstraints, ) *CreditNoteEligibilityResponsePaymentsInner`

NewCreditNoteEligibilityResponsePaymentsInner instantiates a new CreditNoteEligibilityResponsePaymentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEligibilityResponsePaymentsInnerWithDefaults

`func NewCreditNoteEligibilityResponsePaymentsInnerWithDefaults() *CreditNoteEligibilityResponsePaymentsInner`

NewCreditNoteEligibilityResponsePaymentsInnerWithDefaults instantiates a new CreditNoteEligibilityResponsePaymentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetGateway

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetPaymentMethod

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### SetPaymentMethodNil

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *CreditNoteEligibilityResponsePaymentsInner) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetAmountMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetAmountRefundedMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountRefundedMinor() float32`

GetAmountRefundedMinor returns the AmountRefundedMinor field if non-nil, zero value otherwise.

### GetAmountRefundedMinorOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountRefundedMinorOk() (*float32, bool)`

GetAmountRefundedMinorOk returns a tuple with the AmountRefundedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefundedMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountRefundedMinor(v float32)`

SetAmountRefundedMinor sets AmountRefundedMinor field to given value.


### GetAmountAvailableMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountAvailableMinor() float32`

GetAmountAvailableMinor returns the AmountAvailableMinor field if non-nil, zero value otherwise.

### GetAmountAvailableMinorOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAmountAvailableMinorOk() (*float32, bool)`

GetAmountAvailableMinorOk returns a tuple with the AmountAvailableMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAvailableMinor

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetAmountAvailableMinor(v float32)`

SetAmountAvailableMinor sets AmountAvailableMinor field to given value.


### GetCardBrand

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.


### SetCardBrandNil

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *CreditNoteEligibilityResponsePaymentsInner) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLast4

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.


### SetCardLast4Nil

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *CreditNoteEligibilityResponsePaymentsInner) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetTerminalSerial

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetTerminalSerial() string`

GetTerminalSerial returns the TerminalSerial field if non-nil, zero value otherwise.

### GetTerminalSerialOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetTerminalSerialOk() (*string, bool)`

GetTerminalSerialOk returns a tuple with the TerminalSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSerial

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetTerminalSerial(v string)`

SetTerminalSerial sets TerminalSerial field to given value.


### SetTerminalSerialNil

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetTerminalSerialNil(b bool)`

 SetTerminalSerialNil sets the value for TerminalSerial to be an explicit nil

### UnsetTerminalSerial
`func (o *CreditNoteEligibilityResponsePaymentsInner) UnsetTerminalSerial()`

UnsetTerminalSerial ensures that no value is present for TerminalSerial, not even an explicit nil
### GetAvailableOperations

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAvailableOperations() []string`

GetAvailableOperations returns the AvailableOperations field if non-nil, zero value otherwise.

### GetAvailableOperationsOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetAvailableOperationsOk() (*[]string, bool)`

GetAvailableOperationsOk returns a tuple with the AvailableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOperations

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetAvailableOperations(v []string)`

SetAvailableOperations sets AvailableOperations field to given value.


### GetConstraints

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetConstraints() CreditNoteEligibilityResponsePaymentsInnerConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *CreditNoteEligibilityResponsePaymentsInner) GetConstraintsOk() (*CreditNoteEligibilityResponsePaymentsInnerConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *CreditNoteEligibilityResponsePaymentsInner) SetConstraints(v CreditNoteEligibilityResponsePaymentsInnerConstraints)`

SetConstraints sets Constraints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


