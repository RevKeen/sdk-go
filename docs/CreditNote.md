# CreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique credit note identifier | 
**CreditNoteNumber** | **NullableString** | Sequential credit note number for accounting reference | 
**InvoiceId** | **string** | ID of the invoice this credit note applies to | 
**CustomerId** | **NullableString** | ID of the customer who received the credit | 
**AmountMinor** | **int32** | Credit amount in minor units (cents) | 
**TaxAmountMinor** | **NullableInt32** | Tax portion of the credit in minor units | 
**Currency** | **string** | Three-letter ISO 4217 currency code | 
**Status** | **string** | Credit note status: draft, issued, or void | 
**Reason** | **NullableString** | Human-readable reason for the credit | 
**ReasonCode** | **NullableString** | Machine-readable reason code (e.g., billing_error, customer_request) | 
**CreditMethod** | **NullableString** | How the credit is applied: refund_to_payment_method, customer_balance, or external | 
**PdfUrl** | **NullableString** | URL to the credit note PDF document | 
**IssuedAt** | **NullableString** | When the credit note was issued (ISO 8601) | 
**CreatedAt** | **string** | When the credit note was created (ISO 8601) | 
**UpdatedAt** | **string** | When the credit note was last updated (ISO 8601) | 

## Methods

### NewCreditNote

`func NewCreditNote(id string, creditNoteNumber NullableString, invoiceId string, customerId NullableString, amountMinor int32, taxAmountMinor NullableInt32, currency string, status string, reason NullableString, reasonCode NullableString, creditMethod NullableString, pdfUrl NullableString, issuedAt NullableString, createdAt string, updatedAt string, ) *CreditNote`

NewCreditNote instantiates a new CreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteWithDefaults

`func NewCreditNoteWithDefaults() *CreditNote`

NewCreditNoteWithDefaults instantiates a new CreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditNote) SetId(v string)`

SetId sets Id field to given value.


### GetCreditNoteNumber

`func (o *CreditNote) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *CreditNote) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *CreditNote) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.


### SetCreditNoteNumberNil

`func (o *CreditNote) SetCreditNoteNumberNil(b bool)`

 SetCreditNoteNumberNil sets the value for CreditNoteNumber to be an explicit nil

### UnsetCreditNoteNumber
`func (o *CreditNote) UnsetCreditNoteNumber()`

UnsetCreditNoteNumber ensures that no value is present for CreditNoteNumber, not even an explicit nil
### GetInvoiceId

`func (o *CreditNote) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNote) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNote) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetCustomerId

`func (o *CreditNote) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreditNote) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreditNote) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *CreditNote) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CreditNote) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetAmountMinor

`func (o *CreditNote) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreditNote) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreditNote) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetTaxAmountMinor

`func (o *CreditNote) GetTaxAmountMinor() int32`

GetTaxAmountMinor returns the TaxAmountMinor field if non-nil, zero value otherwise.

### GetTaxAmountMinorOk

`func (o *CreditNote) GetTaxAmountMinorOk() (*int32, bool)`

GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountMinor

`func (o *CreditNote) SetTaxAmountMinor(v int32)`

SetTaxAmountMinor sets TaxAmountMinor field to given value.


### SetTaxAmountMinorNil

`func (o *CreditNote) SetTaxAmountMinorNil(b bool)`

 SetTaxAmountMinorNil sets the value for TaxAmountMinor to be an explicit nil

### UnsetTaxAmountMinor
`func (o *CreditNote) UnsetTaxAmountMinor()`

UnsetTaxAmountMinor ensures that no value is present for TaxAmountMinor, not even an explicit nil
### GetCurrency

`func (o *CreditNote) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreditNote) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreditNote) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *CreditNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreditNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreditNote) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *CreditNote) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditNote) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditNote) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *CreditNote) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreditNote) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *CreditNote) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CreditNote) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CreditNote) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### SetReasonCodeNil

`func (o *CreditNote) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *CreditNote) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetCreditMethod

`func (o *CreditNote) GetCreditMethod() string`

GetCreditMethod returns the CreditMethod field if non-nil, zero value otherwise.

### GetCreditMethodOk

`func (o *CreditNote) GetCreditMethodOk() (*string, bool)`

GetCreditMethodOk returns a tuple with the CreditMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMethod

`func (o *CreditNote) SetCreditMethod(v string)`

SetCreditMethod sets CreditMethod field to given value.


### SetCreditMethodNil

`func (o *CreditNote) SetCreditMethodNil(b bool)`

 SetCreditMethodNil sets the value for CreditMethod to be an explicit nil

### UnsetCreditMethod
`func (o *CreditNote) UnsetCreditMethod()`

UnsetCreditMethod ensures that no value is present for CreditMethod, not even an explicit nil
### GetPdfUrl

`func (o *CreditNote) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *CreditNote) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *CreditNote) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.


### SetPdfUrlNil

`func (o *CreditNote) SetPdfUrlNil(b bool)`

 SetPdfUrlNil sets the value for PdfUrl to be an explicit nil

### UnsetPdfUrl
`func (o *CreditNote) UnsetPdfUrl()`

UnsetPdfUrl ensures that no value is present for PdfUrl, not even an explicit nil
### GetIssuedAt

`func (o *CreditNote) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *CreditNote) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *CreditNote) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.


### SetIssuedAtNil

`func (o *CreditNote) SetIssuedAtNil(b bool)`

 SetIssuedAtNil sets the value for IssuedAt to be an explicit nil

### UnsetIssuedAt
`func (o *CreditNote) UnsetIssuedAt()`

UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
### GetCreatedAt

`func (o *CreditNote) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNote) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNote) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreditNote) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreditNote) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreditNote) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


