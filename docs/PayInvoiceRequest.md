# PayInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Collect the invoice&#39;s outstanding amount from the customer. The collection is scheduled, not completed: the invoice stays unpaid until the money settles. |
**PaymentMethod** | **string** | How to collect. Required. Only &#x60;direct_debit&#x60; is available; card collection is not available through this operation. |
**Reference** | **string** | Your reference for the payment (cheque number, bank transfer id, receipt number). Required. It is stored on the payment record. |
**AmountMinor** | **int64** | Amount received, in minor units of the invoice currency (for example pence). Required, a positive integer, never defaulted. It cannot exceed the invoice total less what is already paid. A smaller amount records a partial payment and moves the invoice to partially_paid. |
**Notes** | Pointer to **string** | Free-text note about the payment. Stored on the payment record. | [optional]
**PaidAt** | Pointer to **time.Time** | When the money was received (ISO 8601). Defaults to now. It cannot be in the future. | [optional]

## Methods

### NewPayInvoiceRequest

`func NewPayInvoiceRequest(mode string, paymentMethod string, reference string, amountMinor int64, ) *PayInvoiceRequest`

NewPayInvoiceRequest instantiates a new PayInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayInvoiceRequestWithDefaults

`func NewPayInvoiceRequestWithDefaults() *PayInvoiceRequest`

NewPayInvoiceRequestWithDefaults instantiates a new PayInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *PayInvoiceRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PayInvoiceRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PayInvoiceRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPaymentMethod

`func (o *PayInvoiceRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PayInvoiceRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PayInvoiceRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetReference

`func (o *PayInvoiceRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PayInvoiceRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PayInvoiceRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAmountMinor

`func (o *PayInvoiceRequest) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *PayInvoiceRequest) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *PayInvoiceRequest) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetNotes

`func (o *PayInvoiceRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PayInvoiceRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PayInvoiceRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PayInvoiceRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPaidAt

`func (o *PayInvoiceRequest) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *PayInvoiceRequest) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *PayInvoiceRequest) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *PayInvoiceRequest) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
