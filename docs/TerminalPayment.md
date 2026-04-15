# TerminalPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this terminal payment attempt | 
**InvoiceId** | **NullableString** | Associated invoice ID, or null for walk-in/ad-hoc payments | 
**DeviceId** | **NullableString** | The terminal device that processed (or is processing) this payment | 
**Type** | **string** | Transaction type. sale: original charge. refund: money returned. void: pre-settlement cancellation. | 
**Status** | **string** | Payment lifecycle status. requested: command sent, awaiting card. in_progress: terminal processing. approved: payment succeeded. declined: issuer declined. cancelled: merchant cancelled. error: terminal error. timed_out: no response within 3 minutes. | 
**AmountMinor** | **float32** | Amount in minor units (e.g., pence for GBP, cents for USD) | 
**Currency** | **string** | ISO 4217 currency code | 
**Reference** | **NullableString** | Payment reference (invoice number or custom reference) | 
**TerminalSerial** | **NullableString** | Serial number of the PAX terminal that processed this payment | 
**Uti** | **NullableString** | Unique Transaction Identifier from the terminal | 
**AuthCode** | **NullableString** | Authorization code from the payment processor. Present when approved. | 
**ResponseCode** | **NullableString** | Terminal response code. &#39;00&#39; indicates approval. | 
**Rrn** | **NullableString** | Retrieval Reference Number for settlement reconciliation | 
**CardScheme** | **NullableString** | Card network (e.g., VISA, MASTERCARD, AMEX) | 
**MaskedPan** | **NullableString** | Masked card number — only the last 4 digits are visible | 
**EntryMode** | **NullableString** | How the card was read: contactless (NFC tap), emv_chip (chip insert), magnetic_stripe (swipe), manual_entry (keyed), or fallback (chip-to-swipe) | 
**ErrorMessage** | **NullableString** | Human-readable error message when status is error or timed_out | 
**CreatedAt** | **string** | ISO 8601 timestamp when the payment was initiated | 
**CompletedAt** | **NullableString** | ISO 8601 timestamp when the terminal returned a result, or null if still in progress | 

## Methods

### NewTerminalPayment

`func NewTerminalPayment(id string, invoiceId NullableString, deviceId NullableString, type_ string, status string, amountMinor float32, currency string, reference NullableString, terminalSerial NullableString, uti NullableString, authCode NullableString, responseCode NullableString, rrn NullableString, cardScheme NullableString, maskedPan NullableString, entryMode NullableString, errorMessage NullableString, createdAt string, completedAt NullableString, ) *TerminalPayment`

NewTerminalPayment instantiates a new TerminalPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalPaymentWithDefaults

`func NewTerminalPaymentWithDefaults() *TerminalPayment`

NewTerminalPaymentWithDefaults instantiates a new TerminalPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerminalPayment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerminalPayment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerminalPayment) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *TerminalPayment) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *TerminalPayment) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *TerminalPayment) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *TerminalPayment) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *TerminalPayment) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetDeviceId

`func (o *TerminalPayment) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TerminalPayment) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TerminalPayment) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### SetDeviceIdNil

`func (o *TerminalPayment) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *TerminalPayment) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetType

`func (o *TerminalPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerminalPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerminalPayment) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *TerminalPayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerminalPayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerminalPayment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountMinor

`func (o *TerminalPayment) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *TerminalPayment) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *TerminalPayment) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *TerminalPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TerminalPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TerminalPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReference

`func (o *TerminalPayment) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TerminalPayment) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TerminalPayment) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *TerminalPayment) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *TerminalPayment) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetTerminalSerial

`func (o *TerminalPayment) GetTerminalSerial() string`

GetTerminalSerial returns the TerminalSerial field if non-nil, zero value otherwise.

### GetTerminalSerialOk

`func (o *TerminalPayment) GetTerminalSerialOk() (*string, bool)`

GetTerminalSerialOk returns a tuple with the TerminalSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSerial

`func (o *TerminalPayment) SetTerminalSerial(v string)`

SetTerminalSerial sets TerminalSerial field to given value.


### SetTerminalSerialNil

`func (o *TerminalPayment) SetTerminalSerialNil(b bool)`

 SetTerminalSerialNil sets the value for TerminalSerial to be an explicit nil

### UnsetTerminalSerial
`func (o *TerminalPayment) UnsetTerminalSerial()`

UnsetTerminalSerial ensures that no value is present for TerminalSerial, not even an explicit nil
### GetUti

`func (o *TerminalPayment) GetUti() string`

GetUti returns the Uti field if non-nil, zero value otherwise.

### GetUtiOk

`func (o *TerminalPayment) GetUtiOk() (*string, bool)`

GetUtiOk returns a tuple with the Uti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUti

`func (o *TerminalPayment) SetUti(v string)`

SetUti sets Uti field to given value.


### SetUtiNil

`func (o *TerminalPayment) SetUtiNil(b bool)`

 SetUtiNil sets the value for Uti to be an explicit nil

### UnsetUti
`func (o *TerminalPayment) UnsetUti()`

UnsetUti ensures that no value is present for Uti, not even an explicit nil
### GetAuthCode

`func (o *TerminalPayment) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *TerminalPayment) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *TerminalPayment) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### SetAuthCodeNil

`func (o *TerminalPayment) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *TerminalPayment) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetResponseCode

`func (o *TerminalPayment) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *TerminalPayment) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *TerminalPayment) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### SetResponseCodeNil

`func (o *TerminalPayment) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *TerminalPayment) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetRrn

`func (o *TerminalPayment) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *TerminalPayment) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *TerminalPayment) SetRrn(v string)`

SetRrn sets Rrn field to given value.


### SetRrnNil

`func (o *TerminalPayment) SetRrnNil(b bool)`

 SetRrnNil sets the value for Rrn to be an explicit nil

### UnsetRrn
`func (o *TerminalPayment) UnsetRrn()`

UnsetRrn ensures that no value is present for Rrn, not even an explicit nil
### GetCardScheme

`func (o *TerminalPayment) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *TerminalPayment) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *TerminalPayment) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.


### SetCardSchemeNil

`func (o *TerminalPayment) SetCardSchemeNil(b bool)`

 SetCardSchemeNil sets the value for CardScheme to be an explicit nil

### UnsetCardScheme
`func (o *TerminalPayment) UnsetCardScheme()`

UnsetCardScheme ensures that no value is present for CardScheme, not even an explicit nil
### GetMaskedPan

`func (o *TerminalPayment) GetMaskedPan() string`

GetMaskedPan returns the MaskedPan field if non-nil, zero value otherwise.

### GetMaskedPanOk

`func (o *TerminalPayment) GetMaskedPanOk() (*string, bool)`

GetMaskedPanOk returns a tuple with the MaskedPan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedPan

`func (o *TerminalPayment) SetMaskedPan(v string)`

SetMaskedPan sets MaskedPan field to given value.


### SetMaskedPanNil

`func (o *TerminalPayment) SetMaskedPanNil(b bool)`

 SetMaskedPanNil sets the value for MaskedPan to be an explicit nil

### UnsetMaskedPan
`func (o *TerminalPayment) UnsetMaskedPan()`

UnsetMaskedPan ensures that no value is present for MaskedPan, not even an explicit nil
### GetEntryMode

`func (o *TerminalPayment) GetEntryMode() string`

GetEntryMode returns the EntryMode field if non-nil, zero value otherwise.

### GetEntryModeOk

`func (o *TerminalPayment) GetEntryModeOk() (*string, bool)`

GetEntryModeOk returns a tuple with the EntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMode

`func (o *TerminalPayment) SetEntryMode(v string)`

SetEntryMode sets EntryMode field to given value.


### SetEntryModeNil

`func (o *TerminalPayment) SetEntryModeNil(b bool)`

 SetEntryModeNil sets the value for EntryMode to be an explicit nil

### UnsetEntryMode
`func (o *TerminalPayment) UnsetEntryMode()`

UnsetEntryMode ensures that no value is present for EntryMode, not even an explicit nil
### GetErrorMessage

`func (o *TerminalPayment) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TerminalPayment) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TerminalPayment) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *TerminalPayment) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TerminalPayment) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetCreatedAt

`func (o *TerminalPayment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TerminalPayment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TerminalPayment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *TerminalPayment) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *TerminalPayment) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *TerminalPayment) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *TerminalPayment) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *TerminalPayment) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


