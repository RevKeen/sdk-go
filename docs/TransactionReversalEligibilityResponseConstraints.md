# TransactionReversalEligibilityResponseConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAmount** | **float32** |  | 
**PartialSupported** | **bool** |  | 
**RequiresCardPresent** | **bool** |  | 
**RequiresTerminalOnline** | **bool** |  | 
**RequiresCustomerIban** | **bool** |  | 
**Reason** | **NullableString** |  | 

## Methods

### NewTransactionReversalEligibilityResponseConstraints

`func NewTransactionReversalEligibilityResponseConstraints(maxAmount float32, partialSupported bool, requiresCardPresent bool, requiresTerminalOnline bool, requiresCustomerIban bool, reason NullableString, ) *TransactionReversalEligibilityResponseConstraints`

NewTransactionReversalEligibilityResponseConstraints instantiates a new TransactionReversalEligibilityResponseConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReversalEligibilityResponseConstraintsWithDefaults

`func NewTransactionReversalEligibilityResponseConstraintsWithDefaults() *TransactionReversalEligibilityResponseConstraints`

NewTransactionReversalEligibilityResponseConstraintsWithDefaults instantiates a new TransactionReversalEligibilityResponseConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAmount

`func (o *TransactionReversalEligibilityResponseConstraints) GetMaxAmount() float32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetMaxAmountOk() (*float32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *TransactionReversalEligibilityResponseConstraints) SetMaxAmount(v float32)`

SetMaxAmount sets MaxAmount field to given value.


### GetPartialSupported

`func (o *TransactionReversalEligibilityResponseConstraints) GetPartialSupported() bool`

GetPartialSupported returns the PartialSupported field if non-nil, zero value otherwise.

### GetPartialSupportedOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetPartialSupportedOk() (*bool, bool)`

GetPartialSupportedOk returns a tuple with the PartialSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialSupported

`func (o *TransactionReversalEligibilityResponseConstraints) SetPartialSupported(v bool)`

SetPartialSupported sets PartialSupported field to given value.


### GetRequiresCardPresent

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresCardPresent() bool`

GetRequiresCardPresent returns the RequiresCardPresent field if non-nil, zero value otherwise.

### GetRequiresCardPresentOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresCardPresentOk() (*bool, bool)`

GetRequiresCardPresentOk returns a tuple with the RequiresCardPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCardPresent

`func (o *TransactionReversalEligibilityResponseConstraints) SetRequiresCardPresent(v bool)`

SetRequiresCardPresent sets RequiresCardPresent field to given value.


### GetRequiresTerminalOnline

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresTerminalOnline() bool`

GetRequiresTerminalOnline returns the RequiresTerminalOnline field if non-nil, zero value otherwise.

### GetRequiresTerminalOnlineOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresTerminalOnlineOk() (*bool, bool)`

GetRequiresTerminalOnlineOk returns a tuple with the RequiresTerminalOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTerminalOnline

`func (o *TransactionReversalEligibilityResponseConstraints) SetRequiresTerminalOnline(v bool)`

SetRequiresTerminalOnline sets RequiresTerminalOnline field to given value.


### GetRequiresCustomerIban

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresCustomerIban() bool`

GetRequiresCustomerIban returns the RequiresCustomerIban field if non-nil, zero value otherwise.

### GetRequiresCustomerIbanOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetRequiresCustomerIbanOk() (*bool, bool)`

GetRequiresCustomerIbanOk returns a tuple with the RequiresCustomerIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCustomerIban

`func (o *TransactionReversalEligibilityResponseConstraints) SetRequiresCustomerIban(v bool)`

SetRequiresCustomerIban sets RequiresCustomerIban field to given value.


### GetReason

`func (o *TransactionReversalEligibilityResponseConstraints) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TransactionReversalEligibilityResponseConstraints) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TransactionReversalEligibilityResponseConstraints) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *TransactionReversalEligibilityResponseConstraints) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *TransactionReversalEligibilityResponseConstraints) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


