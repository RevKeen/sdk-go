# RecordInvoicePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Record a payment the customer made outside RevKeen (bank transfer, cash, cheque). RevKeen does not collect any money in this mode. |
**PaymentMethod** | **string** | How the customer paid. Required. |
**Reference** | **string** | Your reference for the payment (cheque number, bank transfer id, receipt number). Required. It is stored on the payment record. |
**AmountMinor** | **int64** | Amount received, in minor units of the invoice currency (for example pence). Required, a positive integer, never defaulted. It cannot exceed the invoice total less what is already paid. A smaller amount records a partial payment and moves the invoice to partially_paid. |
**Notes** | Pointer to **string** | Free-text note about the payment. Stored on the payment record. | [optional]
**PaidAt** | Pointer to **time.Time** | When the money was received (ISO 8601). Defaults to now. It cannot be in the future. | [optional]

## Methods

### NewRecordInvoicePaymentRequest

`func NewRecordInvoicePaymentRequest(mode string, paymentMethod string, reference string, amountMinor int64, ) *RecordInvoicePaymentRequest`

NewRecordInvoicePaymentRequest instantiates a new RecordInvoicePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordInvoicePaymentRequestWithDefaults

`func NewRecordInvoicePaymentRequestWithDefaults() *RecordInvoicePaymentRequest`

NewRecordInvoicePaymentRequestWithDefaults instantiates a new RecordInvoicePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RecordInvoicePaymentRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RecordInvoicePaymentRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RecordInvoicePaymentRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPaymentMethod

`func (o *RecordInvoicePaymentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *RecordInvoicePaymentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *RecordInvoicePaymentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetReference

`func (o *RecordInvoicePaymentRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *RecordInvoicePaymentRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *RecordInvoicePaymentRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAmountMinor

`func (o *RecordInvoicePaymentRequest) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *RecordInvoicePaymentRequest) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *RecordInvoicePaymentRequest) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetNotes

`func (o *RecordInvoicePaymentRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RecordInvoicePaymentRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RecordInvoicePaymentRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RecordInvoicePaymentRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPaidAt

`func (o *RecordInvoicePaymentRequest) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *RecordInvoicePaymentRequest) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *RecordInvoicePaymentRequest) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *RecordInvoicePaymentRequest) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
