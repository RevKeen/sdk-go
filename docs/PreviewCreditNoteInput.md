# PreviewCreditNoteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | Invoice to preview a credit note against. | 
**AmountMinor** | **int32** | Proposed credit amount in cents. The preview validates that this does not exceed the invoice&#39;s outstanding creditable amount. | 
**TaxAmountMinor** | Pointer to **int32** | Proposed tax portion of the credit in cents. | [optional] 
**CreditMethod** | Pointer to **string** | Planned credit method. Defaults to refund_to_payment_method. | [optional] 
**ReasonCode** | Pointer to **string** |  | [optional] 

## Methods

### NewPreviewCreditNoteInput

`func NewPreviewCreditNoteInput(invoiceId string, amountMinor int32, ) *PreviewCreditNoteInput`

NewPreviewCreditNoteInput instantiates a new PreviewCreditNoteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewCreditNoteInputWithDefaults

`func NewPreviewCreditNoteInputWithDefaults() *PreviewCreditNoteInput`

NewPreviewCreditNoteInputWithDefaults instantiates a new PreviewCreditNoteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *PreviewCreditNoteInput) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PreviewCreditNoteInput) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PreviewCreditNoteInput) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMinor

`func (o *PreviewCreditNoteInput) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *PreviewCreditNoteInput) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *PreviewCreditNoteInput) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetTaxAmountMinor

`func (o *PreviewCreditNoteInput) GetTaxAmountMinor() int32`

GetTaxAmountMinor returns the TaxAmountMinor field if non-nil, zero value otherwise.

### GetTaxAmountMinorOk

`func (o *PreviewCreditNoteInput) GetTaxAmountMinorOk() (*int32, bool)`

GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountMinor

`func (o *PreviewCreditNoteInput) SetTaxAmountMinor(v int32)`

SetTaxAmountMinor sets TaxAmountMinor field to given value.

### HasTaxAmountMinor

`func (o *PreviewCreditNoteInput) HasTaxAmountMinor() bool`

HasTaxAmountMinor returns a boolean if a field has been set.

### GetCreditMethod

`func (o *PreviewCreditNoteInput) GetCreditMethod() string`

GetCreditMethod returns the CreditMethod field if non-nil, zero value otherwise.

### GetCreditMethodOk

`func (o *PreviewCreditNoteInput) GetCreditMethodOk() (*string, bool)`

GetCreditMethodOk returns a tuple with the CreditMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMethod

`func (o *PreviewCreditNoteInput) SetCreditMethod(v string)`

SetCreditMethod sets CreditMethod field to given value.

### HasCreditMethod

`func (o *PreviewCreditNoteInput) HasCreditMethod() bool`

HasCreditMethod returns a boolean if a field has been set.

### GetReasonCode

`func (o *PreviewCreditNoteInput) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *PreviewCreditNoteInput) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *PreviewCreditNoteInput) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *PreviewCreditNoteInput) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


