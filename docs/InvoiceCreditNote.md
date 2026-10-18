# InvoiceCreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique credit note identifier |
**CreditNoteNumber** | **NullableString** | Sequential credit note number for accounting reference |
**InvoiceId** | **string** | ID of the invoice this credit note applies to |
**CustomerId** | **NullableString** | ID of the customer who received the credit |
**AmountMinor** | **int64** | Credit amount in minor units (cents) |
**TaxAmountMinor** | **NullableInt64** | Tax portion of the credit in minor units |
**Currency** | **string** | Three-letter ISO 4217 currency code |
**Status** | **string** | Credit note status: draft, issued, or void |
**Reason** | **NullableString** | Human-readable reason for the credit |
**ReasonCode** | **NullableString** | Machine-readable reason code (e.g., billing_error, customer_request) |
**CreditMethod** | **NullableString** | How the credit is applied: refund_to_payment_method, customer_balance, or external |
**PdfUrl** | **NullableString** | URL to the credit note PDF document |
**IssuedAt** | **NullableTime** | When the credit note was issued (ISO 8601) |
**CreatedAt** | **time.Time** | When the credit note was created (ISO 8601) |
**UpdatedAt** | **time.Time** | When the credit note was last updated (ISO 8601) |

## Methods

### NewInvoiceCreditNote

`func NewInvoiceCreditNote(id string, creditNoteNumber NullableString, invoiceId string, customerId NullableString, amountMinor int64, taxAmountMinor NullableInt64, currency string, status string, reason NullableString, reasonCode NullableString, creditMethod NullableString, pdfUrl NullableString, issuedAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *InvoiceCreditNote`

NewInvoiceCreditNote instantiates a new InvoiceCreditNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreditNoteWithDefaults

`func NewInvoiceCreditNoteWithDefaults() *InvoiceCreditNote`

NewInvoiceCreditNoteWithDefaults instantiates a new InvoiceCreditNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceCreditNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceCreditNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceCreditNote) SetId(v string)`

SetId sets Id field to given value.


### GetCreditNoteNumber

`func (o *InvoiceCreditNote) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *InvoiceCreditNote) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *InvoiceCreditNote) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.


### SetCreditNoteNumberNil

`func (o *InvoiceCreditNote) SetCreditNoteNumberNil(b bool)`

 SetCreditNoteNumberNil sets the value for CreditNoteNumber to be an explicit nil

### UnsetCreditNoteNumber
`func (o *InvoiceCreditNote) UnsetCreditNoteNumber()`

UnsetCreditNoteNumber ensures that no value is present for CreditNoteNumber, not even an explicit nil
### GetInvoiceId

`func (o *InvoiceCreditNote) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceCreditNote) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceCreditNote) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetCustomerId

`func (o *InvoiceCreditNote) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *InvoiceCreditNote) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *InvoiceCreditNote) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *InvoiceCreditNote) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *InvoiceCreditNote) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetAmountMinor

`func (o *InvoiceCreditNote) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoiceCreditNote) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoiceCreditNote) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetTaxAmountMinor

`func (o *InvoiceCreditNote) GetTaxAmountMinor() int64`

GetTaxAmountMinor returns the TaxAmountMinor field if non-nil, zero value otherwise.

### GetTaxAmountMinorOk

`func (o *InvoiceCreditNote) GetTaxAmountMinorOk() (*int64, bool)`

GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountMinor

`func (o *InvoiceCreditNote) SetTaxAmountMinor(v int64)`

SetTaxAmountMinor sets TaxAmountMinor field to given value.


### SetTaxAmountMinorNil

`func (o *InvoiceCreditNote) SetTaxAmountMinorNil(b bool)`

 SetTaxAmountMinorNil sets the value for TaxAmountMinor to be an explicit nil

### UnsetTaxAmountMinor
`func (o *InvoiceCreditNote) UnsetTaxAmountMinor()`

UnsetTaxAmountMinor ensures that no value is present for TaxAmountMinor, not even an explicit nil
### GetCurrency

`func (o *InvoiceCreditNote) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceCreditNote) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceCreditNote) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *InvoiceCreditNote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceCreditNote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceCreditNote) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *InvoiceCreditNote) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceCreditNote) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceCreditNote) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *InvoiceCreditNote) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *InvoiceCreditNote) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *InvoiceCreditNote) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *InvoiceCreditNote) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *InvoiceCreditNote) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### SetReasonCodeNil

`func (o *InvoiceCreditNote) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *InvoiceCreditNote) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetCreditMethod

`func (o *InvoiceCreditNote) GetCreditMethod() string`

GetCreditMethod returns the CreditMethod field if non-nil, zero value otherwise.

### GetCreditMethodOk

`func (o *InvoiceCreditNote) GetCreditMethodOk() (*string, bool)`

GetCreditMethodOk returns a tuple with the CreditMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMethod

`func (o *InvoiceCreditNote) SetCreditMethod(v string)`

SetCreditMethod sets CreditMethod field to given value.


### SetCreditMethodNil

`func (o *InvoiceCreditNote) SetCreditMethodNil(b bool)`

 SetCreditMethodNil sets the value for CreditMethod to be an explicit nil

### UnsetCreditMethod
`func (o *InvoiceCreditNote) UnsetCreditMethod()`

UnsetCreditMethod ensures that no value is present for CreditMethod, not even an explicit nil
### GetPdfUrl

`func (o *InvoiceCreditNote) GetPdfUrl() string`

GetPdfUrl returns the PdfUrl field if non-nil, zero value otherwise.

### GetPdfUrlOk

`func (o *InvoiceCreditNote) GetPdfUrlOk() (*string, bool)`

GetPdfUrlOk returns a tuple with the PdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfUrl

`func (o *InvoiceCreditNote) SetPdfUrl(v string)`

SetPdfUrl sets PdfUrl field to given value.


### SetPdfUrlNil

`func (o *InvoiceCreditNote) SetPdfUrlNil(b bool)`

 SetPdfUrlNil sets the value for PdfUrl to be an explicit nil

### UnsetPdfUrl
`func (o *InvoiceCreditNote) UnsetPdfUrl()`

UnsetPdfUrl ensures that no value is present for PdfUrl, not even an explicit nil
### GetIssuedAt

`func (o *InvoiceCreditNote) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *InvoiceCreditNote) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *InvoiceCreditNote) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### SetIssuedAtNil

`func (o *InvoiceCreditNote) SetIssuedAtNil(b bool)`

 SetIssuedAtNil sets the value for IssuedAt to be an explicit nil

### UnsetIssuedAt
`func (o *InvoiceCreditNote) UnsetIssuedAt()`

UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
### GetCreatedAt

`func (o *InvoiceCreditNote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceCreditNote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceCreditNote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceCreditNote) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceCreditNote) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceCreditNote) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
