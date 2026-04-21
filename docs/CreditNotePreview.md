# CreditNotePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**InvoiceId** | **string** |  | 
**ProposedAmountMinor** | **int32** |  | 
**ProposedTaxAmountMinor** | **NullableInt32** |  | 
**Currency** | **string** |  | 
**Invoice** | [**CreditNotePreviewInvoice**](CreditNotePreviewInvoice.md) |  | 
**AfterCredit** | [**CreditNotePreviewAfterCredit**](CreditNotePreviewAfterCredit.md) |  | 
**ExceedsMaxCreditable** | **bool** | True if &#x60;amount_minor&#x60; exceeds the invoice&#39;s remaining creditable amount. A real create call would be rejected. | 
**CreditMethod** | **string** |  | 

## Methods

### NewCreditNotePreview

`func NewCreditNotePreview(object string, invoiceId string, proposedAmountMinor int32, proposedTaxAmountMinor NullableInt32, currency string, invoice CreditNotePreviewInvoice, afterCredit CreditNotePreviewAfterCredit, exceedsMaxCreditable bool, creditMethod string, ) *CreditNotePreview`

NewCreditNotePreview instantiates a new CreditNotePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNotePreviewWithDefaults

`func NewCreditNotePreviewWithDefaults() *CreditNotePreview`

NewCreditNotePreviewWithDefaults instantiates a new CreditNotePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreditNotePreview) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreditNotePreview) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreditNotePreview) SetObject(v string)`

SetObject sets Object field to given value.


### GetInvoiceId

`func (o *CreditNotePreview) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNotePreview) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNotePreview) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetProposedAmountMinor

`func (o *CreditNotePreview) GetProposedAmountMinor() int32`

GetProposedAmountMinor returns the ProposedAmountMinor field if non-nil, zero value otherwise.

### GetProposedAmountMinorOk

`func (o *CreditNotePreview) GetProposedAmountMinorOk() (*int32, bool)`

GetProposedAmountMinorOk returns a tuple with the ProposedAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedAmountMinor

`func (o *CreditNotePreview) SetProposedAmountMinor(v int32)`

SetProposedAmountMinor sets ProposedAmountMinor field to given value.


### GetProposedTaxAmountMinor

`func (o *CreditNotePreview) GetProposedTaxAmountMinor() int32`

GetProposedTaxAmountMinor returns the ProposedTaxAmountMinor field if non-nil, zero value otherwise.

### GetProposedTaxAmountMinorOk

`func (o *CreditNotePreview) GetProposedTaxAmountMinorOk() (*int32, bool)`

GetProposedTaxAmountMinorOk returns a tuple with the ProposedTaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedTaxAmountMinor

`func (o *CreditNotePreview) SetProposedTaxAmountMinor(v int32)`

SetProposedTaxAmountMinor sets ProposedTaxAmountMinor field to given value.


### SetProposedTaxAmountMinorNil

`func (o *CreditNotePreview) SetProposedTaxAmountMinorNil(b bool)`

 SetProposedTaxAmountMinorNil sets the value for ProposedTaxAmountMinor to be an explicit nil

### UnsetProposedTaxAmountMinor
`func (o *CreditNotePreview) UnsetProposedTaxAmountMinor()`

UnsetProposedTaxAmountMinor ensures that no value is present for ProposedTaxAmountMinor, not even an explicit nil
### GetCurrency

`func (o *CreditNotePreview) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNotePreview) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNotePreview) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInvoice

`func (o *CreditNotePreview) GetInvoice() CreditNotePreviewInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *CreditNotePreview) GetInvoiceOk() (*CreditNotePreviewInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *CreditNotePreview) SetInvoice(v CreditNotePreviewInvoice)`

SetInvoice sets Invoice field to given value.


### GetAfterCredit

`func (o *CreditNotePreview) GetAfterCredit() CreditNotePreviewAfterCredit`

GetAfterCredit returns the AfterCredit field if non-nil, zero value otherwise.

### GetAfterCreditOk

`func (o *CreditNotePreview) GetAfterCreditOk() (*CreditNotePreviewAfterCredit, bool)`

GetAfterCreditOk returns a tuple with the AfterCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCredit

`func (o *CreditNotePreview) SetAfterCredit(v CreditNotePreviewAfterCredit)`

SetAfterCredit sets AfterCredit field to given value.


### GetExceedsMaxCreditable

`func (o *CreditNotePreview) GetExceedsMaxCreditable() bool`

GetExceedsMaxCreditable returns the ExceedsMaxCreditable field if non-nil, zero value otherwise.

### GetExceedsMaxCreditableOk

`func (o *CreditNotePreview) GetExceedsMaxCreditableOk() (*bool, bool)`

GetExceedsMaxCreditableOk returns a tuple with the ExceedsMaxCreditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceedsMaxCreditable

`func (o *CreditNotePreview) SetExceedsMaxCreditable(v bool)`

SetExceedsMaxCreditable sets ExceedsMaxCreditable field to given value.


### GetCreditMethod

`func (o *CreditNotePreview) GetCreditMethod() string`

GetCreditMethod returns the CreditMethod field if non-nil, zero value otherwise.

### GetCreditMethodOk

`func (o *CreditNotePreview) GetCreditMethodOk() (*string, bool)`

GetCreditMethodOk returns a tuple with the CreditMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMethod

`func (o *CreditNotePreview) SetCreditMethod(v string)`

SetCreditMethod sets CreditMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


