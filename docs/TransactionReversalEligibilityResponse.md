# TransactionReversalEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanReverse** | **bool** |  | 
**TransactionId** | **string** |  | 
**Gateway** | **string** |  | 
**PaymentMethod** | **NullableString** |  | 
**AvailableOperations** | **[]string** |  | 
**Constraints** | [**TransactionReversalEligibilityResponseConstraints**](TransactionReversalEligibilityResponseConstraints.md) |  | 
**CardBrand** | **NullableString** |  | 
**CardLast4** | **NullableString** |  | 
**TerminalSerial** | **NullableString** |  | 
**TerminalUti** | **NullableString** |  | 

## Methods

### NewTransactionReversalEligibilityResponse

`func NewTransactionReversalEligibilityResponse(canReverse bool, transactionId string, gateway string, paymentMethod NullableString, availableOperations []string, constraints TransactionReversalEligibilityResponseConstraints, cardBrand NullableString, cardLast4 NullableString, terminalSerial NullableString, terminalUti NullableString, ) *TransactionReversalEligibilityResponse`

NewTransactionReversalEligibilityResponse instantiates a new TransactionReversalEligibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReversalEligibilityResponseWithDefaults

`func NewTransactionReversalEligibilityResponseWithDefaults() *TransactionReversalEligibilityResponse`

NewTransactionReversalEligibilityResponseWithDefaults instantiates a new TransactionReversalEligibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanReverse

`func (o *TransactionReversalEligibilityResponse) GetCanReverse() bool`

GetCanReverse returns the CanReverse field if non-nil, zero value otherwise.

### GetCanReverseOk

`func (o *TransactionReversalEligibilityResponse) GetCanReverseOk() (*bool, bool)`

GetCanReverseOk returns a tuple with the CanReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReverse

`func (o *TransactionReversalEligibilityResponse) SetCanReverse(v bool)`

SetCanReverse sets CanReverse field to given value.


### GetTransactionId

`func (o *TransactionReversalEligibilityResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionReversalEligibilityResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionReversalEligibilityResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetGateway

`func (o *TransactionReversalEligibilityResponse) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *TransactionReversalEligibilityResponse) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *TransactionReversalEligibilityResponse) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetPaymentMethod

`func (o *TransactionReversalEligibilityResponse) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *TransactionReversalEligibilityResponse) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *TransactionReversalEligibilityResponse) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### SetPaymentMethodNil

`func (o *TransactionReversalEligibilityResponse) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *TransactionReversalEligibilityResponse) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetAvailableOperations

`func (o *TransactionReversalEligibilityResponse) GetAvailableOperations() []string`

GetAvailableOperations returns the AvailableOperations field if non-nil, zero value otherwise.

### GetAvailableOperationsOk

`func (o *TransactionReversalEligibilityResponse) GetAvailableOperationsOk() (*[]string, bool)`

GetAvailableOperationsOk returns a tuple with the AvailableOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOperations

`func (o *TransactionReversalEligibilityResponse) SetAvailableOperations(v []string)`

SetAvailableOperations sets AvailableOperations field to given value.


### GetConstraints

`func (o *TransactionReversalEligibilityResponse) GetConstraints() TransactionReversalEligibilityResponseConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *TransactionReversalEligibilityResponse) GetConstraintsOk() (*TransactionReversalEligibilityResponseConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *TransactionReversalEligibilityResponse) SetConstraints(v TransactionReversalEligibilityResponseConstraints)`

SetConstraints sets Constraints field to given value.


### GetCardBrand

`func (o *TransactionReversalEligibilityResponse) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *TransactionReversalEligibilityResponse) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *TransactionReversalEligibilityResponse) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.


### SetCardBrandNil

`func (o *TransactionReversalEligibilityResponse) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *TransactionReversalEligibilityResponse) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLast4

`func (o *TransactionReversalEligibilityResponse) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *TransactionReversalEligibilityResponse) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *TransactionReversalEligibilityResponse) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.


### SetCardLast4Nil

`func (o *TransactionReversalEligibilityResponse) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *TransactionReversalEligibilityResponse) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetTerminalSerial

`func (o *TransactionReversalEligibilityResponse) GetTerminalSerial() string`

GetTerminalSerial returns the TerminalSerial field if non-nil, zero value otherwise.

### GetTerminalSerialOk

`func (o *TransactionReversalEligibilityResponse) GetTerminalSerialOk() (*string, bool)`

GetTerminalSerialOk returns a tuple with the TerminalSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSerial

`func (o *TransactionReversalEligibilityResponse) SetTerminalSerial(v string)`

SetTerminalSerial sets TerminalSerial field to given value.


### SetTerminalSerialNil

`func (o *TransactionReversalEligibilityResponse) SetTerminalSerialNil(b bool)`

 SetTerminalSerialNil sets the value for TerminalSerial to be an explicit nil

### UnsetTerminalSerial
`func (o *TransactionReversalEligibilityResponse) UnsetTerminalSerial()`

UnsetTerminalSerial ensures that no value is present for TerminalSerial, not even an explicit nil
### GetTerminalUti

`func (o *TransactionReversalEligibilityResponse) GetTerminalUti() string`

GetTerminalUti returns the TerminalUti field if non-nil, zero value otherwise.

### GetTerminalUtiOk

`func (o *TransactionReversalEligibilityResponse) GetTerminalUtiOk() (*string, bool)`

GetTerminalUtiOk returns a tuple with the TerminalUti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalUti

`func (o *TransactionReversalEligibilityResponse) SetTerminalUti(v string)`

SetTerminalUti sets TerminalUti field to given value.


### SetTerminalUtiNil

`func (o *TransactionReversalEligibilityResponse) SetTerminalUtiNil(b bool)`

 SetTerminalUtiNil sets the value for TerminalUti to be an explicit nil

### UnsetTerminalUti
`func (o *TransactionReversalEligibilityResponse) UnsetTerminalUti()`

UnsetTerminalUti ensures that no value is present for TerminalUti, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


