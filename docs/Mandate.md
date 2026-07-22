# Mandate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mandate ID | 
**MandateRef** | **string** | RevKeen mandate reference (Bacs DDI reference) | 
**MandateReference** | **NullableString** | Human-facing mandate reference | 
**CustomerId** | **string** | Customer the mandate belongs to | 
**Status** | **string** | Mandate status | 
**AccountHolderName** | **NullableString** | Name on the bank account | 
**SortCode** | **string** | Masked sort code — never the full value | 
**AccountNumberLast4** | **NullableString** | Last 4 digits of the account number | 
**BankName** | **NullableString** | Resolved bank name | 
**StatementName** | Pointer to **NullableString** | Bacs statement descriptor (max 18 chars) | [optional] 
**NoticeDays** | **NullableInt32** | Advance-notice days applied | 
**FirstCollectionDate** | Pointer to **NullableString** | Earliest collection date (YYYY-MM-DD) | [optional] 
**NextCollectionDate** | Pointer to **NullableString** | Next scheduled collection date (YYYY-MM-DD) | [optional] 
**BackupPaymentMethodId** | Pointer to **NullableString** | Recovery fallback card | [optional] 
**ActivatedAt** | Pointer to **NullableString** | When the mandate became active (ISO 8601) | [optional] 
**SuspendedAt** | Pointer to **NullableString** | When the mandate was suspended (ISO 8601) | [optional] 
**CancelledAt** | Pointer to **NullableString** | When the mandate was cancelled (ISO 8601) | [optional] 
**FailureReason** | Pointer to **NullableString** | Most recent failure reason | [optional] 
**MandateRequestId** | Pointer to **NullableString** | Mandate-request consumed on creation, if any | [optional] 
**InvoiceId** | Pointer to **NullableString** | Invoice linked via the mandate-request, if any | [optional] 
**PdfDocuments** | Pointer to [**[]DdMandatePdfReference**](DdMandatePdfReference.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableString** | Creation timestamp (ISO 8601) | [optional] 
**UpdatedAt** | Pointer to **NullableString** | Last-updated timestamp (ISO 8601) | [optional] 

## Methods

### NewMandate

`func NewMandate(id string, mandateRef string, mandateReference NullableString, customerId string, status string, accountHolderName NullableString, sortCode string, accountNumberLast4 NullableString, bankName NullableString, noticeDays NullableInt32, ) *Mandate`

NewMandate instantiates a new Mandate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateWithDefaults

`func NewMandateWithDefaults() *Mandate`

NewMandateWithDefaults instantiates a new Mandate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Mandate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mandate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mandate) SetId(v string)`

SetId sets Id field to given value.


### GetMandateRef

`func (o *Mandate) GetMandateRef() string`

GetMandateRef returns the MandateRef field if non-nil, zero value otherwise.

### GetMandateRefOk

`func (o *Mandate) GetMandateRefOk() (*string, bool)`

GetMandateRefOk returns a tuple with the MandateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRef

`func (o *Mandate) SetMandateRef(v string)`

SetMandateRef sets MandateRef field to given value.


### GetMandateReference

`func (o *Mandate) GetMandateReference() string`

GetMandateReference returns the MandateReference field if non-nil, zero value otherwise.

### GetMandateReferenceOk

`func (o *Mandate) GetMandateReferenceOk() (*string, bool)`

GetMandateReferenceOk returns a tuple with the MandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateReference

`func (o *Mandate) SetMandateReference(v string)`

SetMandateReference sets MandateReference field to given value.


### SetMandateReferenceNil

`func (o *Mandate) SetMandateReferenceNil(b bool)`

 SetMandateReferenceNil sets the value for MandateReference to be an explicit nil

### UnsetMandateReference
`func (o *Mandate) UnsetMandateReference()`

UnsetMandateReference ensures that no value is present for MandateReference, not even an explicit nil
### GetCustomerId

`func (o *Mandate) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Mandate) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Mandate) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStatus

`func (o *Mandate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Mandate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Mandate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountHolderName

`func (o *Mandate) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *Mandate) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *Mandate) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### SetAccountHolderNameNil

`func (o *Mandate) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *Mandate) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSortCode

`func (o *Mandate) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *Mandate) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *Mandate) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetAccountNumberLast4

`func (o *Mandate) GetAccountNumberLast4() string`

GetAccountNumberLast4 returns the AccountNumberLast4 field if non-nil, zero value otherwise.

### GetAccountNumberLast4Ok

`func (o *Mandate) GetAccountNumberLast4Ok() (*string, bool)`

GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLast4

`func (o *Mandate) SetAccountNumberLast4(v string)`

SetAccountNumberLast4 sets AccountNumberLast4 field to given value.


### SetAccountNumberLast4Nil

`func (o *Mandate) SetAccountNumberLast4Nil(b bool)`

 SetAccountNumberLast4Nil sets the value for AccountNumberLast4 to be an explicit nil

### UnsetAccountNumberLast4
`func (o *Mandate) UnsetAccountNumberLast4()`

UnsetAccountNumberLast4 ensures that no value is present for AccountNumberLast4, not even an explicit nil
### GetBankName

`func (o *Mandate) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *Mandate) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *Mandate) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *Mandate) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *Mandate) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetStatementName

`func (o *Mandate) GetStatementName() string`

GetStatementName returns the StatementName field if non-nil, zero value otherwise.

### GetStatementNameOk

`func (o *Mandate) GetStatementNameOk() (*string, bool)`

GetStatementNameOk returns a tuple with the StatementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementName

`func (o *Mandate) SetStatementName(v string)`

SetStatementName sets StatementName field to given value.

### HasStatementName

`func (o *Mandate) HasStatementName() bool`

HasStatementName returns a boolean if a field has been set.

### SetStatementNameNil

`func (o *Mandate) SetStatementNameNil(b bool)`

 SetStatementNameNil sets the value for StatementName to be an explicit nil

### UnsetStatementName
`func (o *Mandate) UnsetStatementName()`

UnsetStatementName ensures that no value is present for StatementName, not even an explicit nil
### GetNoticeDays

`func (o *Mandate) GetNoticeDays() int32`

GetNoticeDays returns the NoticeDays field if non-nil, zero value otherwise.

### GetNoticeDaysOk

`func (o *Mandate) GetNoticeDaysOk() (*int32, bool)`

GetNoticeDaysOk returns a tuple with the NoticeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDays

`func (o *Mandate) SetNoticeDays(v int32)`

SetNoticeDays sets NoticeDays field to given value.


### SetNoticeDaysNil

`func (o *Mandate) SetNoticeDaysNil(b bool)`

 SetNoticeDaysNil sets the value for NoticeDays to be an explicit nil

### UnsetNoticeDays
`func (o *Mandate) UnsetNoticeDays()`

UnsetNoticeDays ensures that no value is present for NoticeDays, not even an explicit nil
### GetFirstCollectionDate

`func (o *Mandate) GetFirstCollectionDate() string`

GetFirstCollectionDate returns the FirstCollectionDate field if non-nil, zero value otherwise.

### GetFirstCollectionDateOk

`func (o *Mandate) GetFirstCollectionDateOk() (*string, bool)`

GetFirstCollectionDateOk returns a tuple with the FirstCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCollectionDate

`func (o *Mandate) SetFirstCollectionDate(v string)`

SetFirstCollectionDate sets FirstCollectionDate field to given value.

### HasFirstCollectionDate

`func (o *Mandate) HasFirstCollectionDate() bool`

HasFirstCollectionDate returns a boolean if a field has been set.

### SetFirstCollectionDateNil

`func (o *Mandate) SetFirstCollectionDateNil(b bool)`

 SetFirstCollectionDateNil sets the value for FirstCollectionDate to be an explicit nil

### UnsetFirstCollectionDate
`func (o *Mandate) UnsetFirstCollectionDate()`

UnsetFirstCollectionDate ensures that no value is present for FirstCollectionDate, not even an explicit nil
### GetNextCollectionDate

`func (o *Mandate) GetNextCollectionDate() string`

GetNextCollectionDate returns the NextCollectionDate field if non-nil, zero value otherwise.

### GetNextCollectionDateOk

`func (o *Mandate) GetNextCollectionDateOk() (*string, bool)`

GetNextCollectionDateOk returns a tuple with the NextCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCollectionDate

`func (o *Mandate) SetNextCollectionDate(v string)`

SetNextCollectionDate sets NextCollectionDate field to given value.

### HasNextCollectionDate

`func (o *Mandate) HasNextCollectionDate() bool`

HasNextCollectionDate returns a boolean if a field has been set.

### SetNextCollectionDateNil

`func (o *Mandate) SetNextCollectionDateNil(b bool)`

 SetNextCollectionDateNil sets the value for NextCollectionDate to be an explicit nil

### UnsetNextCollectionDate
`func (o *Mandate) UnsetNextCollectionDate()`

UnsetNextCollectionDate ensures that no value is present for NextCollectionDate, not even an explicit nil
### GetBackupPaymentMethodId

`func (o *Mandate) GetBackupPaymentMethodId() string`

GetBackupPaymentMethodId returns the BackupPaymentMethodId field if non-nil, zero value otherwise.

### GetBackupPaymentMethodIdOk

`func (o *Mandate) GetBackupPaymentMethodIdOk() (*string, bool)`

GetBackupPaymentMethodIdOk returns a tuple with the BackupPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPaymentMethodId

`func (o *Mandate) SetBackupPaymentMethodId(v string)`

SetBackupPaymentMethodId sets BackupPaymentMethodId field to given value.

### HasBackupPaymentMethodId

`func (o *Mandate) HasBackupPaymentMethodId() bool`

HasBackupPaymentMethodId returns a boolean if a field has been set.

### SetBackupPaymentMethodIdNil

`func (o *Mandate) SetBackupPaymentMethodIdNil(b bool)`

 SetBackupPaymentMethodIdNil sets the value for BackupPaymentMethodId to be an explicit nil

### UnsetBackupPaymentMethodId
`func (o *Mandate) UnsetBackupPaymentMethodId()`

UnsetBackupPaymentMethodId ensures that no value is present for BackupPaymentMethodId, not even an explicit nil
### GetActivatedAt

`func (o *Mandate) GetActivatedAt() string`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Mandate) GetActivatedAtOk() (*string, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *Mandate) SetActivatedAt(v string)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *Mandate) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *Mandate) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *Mandate) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetSuspendedAt

`func (o *Mandate) GetSuspendedAt() string`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *Mandate) GetSuspendedAtOk() (*string, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *Mandate) SetSuspendedAt(v string)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *Mandate) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.

### SetSuspendedAtNil

`func (o *Mandate) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *Mandate) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetCancelledAt

`func (o *Mandate) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *Mandate) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *Mandate) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *Mandate) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *Mandate) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *Mandate) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetFailureReason

`func (o *Mandate) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *Mandate) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *Mandate) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *Mandate) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *Mandate) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *Mandate) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetMandateRequestId

`func (o *Mandate) GetMandateRequestId() string`

GetMandateRequestId returns the MandateRequestId field if non-nil, zero value otherwise.

### GetMandateRequestIdOk

`func (o *Mandate) GetMandateRequestIdOk() (*string, bool)`

GetMandateRequestIdOk returns a tuple with the MandateRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRequestId

`func (o *Mandate) SetMandateRequestId(v string)`

SetMandateRequestId sets MandateRequestId field to given value.

### HasMandateRequestId

`func (o *Mandate) HasMandateRequestId() bool`

HasMandateRequestId returns a boolean if a field has been set.

### SetMandateRequestIdNil

`func (o *Mandate) SetMandateRequestIdNil(b bool)`

 SetMandateRequestIdNil sets the value for MandateRequestId to be an explicit nil

### UnsetMandateRequestId
`func (o *Mandate) UnsetMandateRequestId()`

UnsetMandateRequestId ensures that no value is present for MandateRequestId, not even an explicit nil
### GetInvoiceId

`func (o *Mandate) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Mandate) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Mandate) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Mandate) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *Mandate) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *Mandate) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetPdfDocuments

`func (o *Mandate) GetPdfDocuments() []DdMandatePdfReference`

GetPdfDocuments returns the PdfDocuments field if non-nil, zero value otherwise.

### GetPdfDocumentsOk

`func (o *Mandate) GetPdfDocumentsOk() (*[]DdMandatePdfReference, bool)`

GetPdfDocumentsOk returns a tuple with the PdfDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfDocuments

`func (o *Mandate) SetPdfDocuments(v []DdMandatePdfReference)`

SetPdfDocuments sets PdfDocuments field to given value.

### HasPdfDocuments

`func (o *Mandate) HasPdfDocuments() bool`

HasPdfDocuments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Mandate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Mandate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Mandate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Mandate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Mandate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Mandate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Mandate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Mandate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Mandate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Mandate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Mandate) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Mandate) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


