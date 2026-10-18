# CreditNotePreviewAfterCredit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAmountDueMinor** | **int64** |  |
**NewAmountCreditedMinor** | **int64** |  |
**WouldLeaveOutstanding** | **bool** | True if applying this credit would leave a non-zero balance on the invoice. |

## Methods

### NewCreditNotePreviewAfterCredit

`func NewCreditNotePreviewAfterCredit(newAmountDueMinor int64, newAmountCreditedMinor int64, wouldLeaveOutstanding bool, ) *CreditNotePreviewAfterCredit`

NewCreditNotePreviewAfterCredit instantiates a new CreditNotePreviewAfterCredit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNotePreviewAfterCreditWithDefaults

`func NewCreditNotePreviewAfterCreditWithDefaults() *CreditNotePreviewAfterCredit`

NewCreditNotePreviewAfterCreditWithDefaults instantiates a new CreditNotePreviewAfterCredit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAmountDueMinor

`func (o *CreditNotePreviewAfterCredit) GetNewAmountDueMinor() int64`

GetNewAmountDueMinor returns the NewAmountDueMinor field if non-nil, zero value otherwise.

### GetNewAmountDueMinorOk

`func (o *CreditNotePreviewAfterCredit) GetNewAmountDueMinorOk() (*int64, bool)`

GetNewAmountDueMinorOk returns a tuple with the NewAmountDueMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmountDueMinor

`func (o *CreditNotePreviewAfterCredit) SetNewAmountDueMinor(v int64)`

SetNewAmountDueMinor sets NewAmountDueMinor field to given value.


### GetNewAmountCreditedMinor

`func (o *CreditNotePreviewAfterCredit) GetNewAmountCreditedMinor() int64`

GetNewAmountCreditedMinor returns the NewAmountCreditedMinor field if non-nil, zero value otherwise.

### GetNewAmountCreditedMinorOk

`func (o *CreditNotePreviewAfterCredit) GetNewAmountCreditedMinorOk() (*int64, bool)`

GetNewAmountCreditedMinorOk returns a tuple with the NewAmountCreditedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmountCreditedMinor

`func (o *CreditNotePreviewAfterCredit) SetNewAmountCreditedMinor(v int64)`

SetNewAmountCreditedMinor sets NewAmountCreditedMinor field to given value.


### GetWouldLeaveOutstanding

`func (o *CreditNotePreviewAfterCredit) GetWouldLeaveOutstanding() bool`

GetWouldLeaveOutstanding returns the WouldLeaveOutstanding field if non-nil, zero value otherwise.

### GetWouldLeaveOutstandingOk

`func (o *CreditNotePreviewAfterCredit) GetWouldLeaveOutstandingOk() (*bool, bool)`

GetWouldLeaveOutstandingOk returns a tuple with the WouldLeaveOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWouldLeaveOutstanding

`func (o *CreditNotePreviewAfterCredit) SetWouldLeaveOutstanding(v bool)`

SetWouldLeaveOutstanding sets WouldLeaveOutstanding field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
