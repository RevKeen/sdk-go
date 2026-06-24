# CreateTerminalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | The ID of the terminal device to send the payment to. Use List Devices to discover available device IDs. Even merchants with a single terminal must pass the device_id explicitly — there is no auto-routing fallback. | 
**AmountMinor** | **int32** | Amount in minor units (e.g., pence for GBP, cents for USD) | 
**Currency** | **string** | ISO 4217 currency code | 
**InvoiceId** | Pointer to **string** | The invoice to charge. Omit for walk-in or ad-hoc payments where no invoice exists. When omitted, the terminal payment is recorded without an invoice association. | [optional] 
**Reference** | Pointer to **string** | Custom reference for the payment. Auto-generated if not provided. | [optional] 

## Methods

### NewCreateTerminalPaymentRequest

`func NewCreateTerminalPaymentRequest(deviceId string, amountMinor int32, currency string, ) *CreateTerminalPaymentRequest`

NewCreateTerminalPaymentRequest instantiates a new CreateTerminalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTerminalPaymentRequestWithDefaults

`func NewCreateTerminalPaymentRequestWithDefaults() *CreateTerminalPaymentRequest`

NewCreateTerminalPaymentRequestWithDefaults instantiates a new CreateTerminalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *CreateTerminalPaymentRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CreateTerminalPaymentRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CreateTerminalPaymentRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetAmountMinor

`func (o *CreateTerminalPaymentRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateTerminalPaymentRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateTerminalPaymentRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *CreateTerminalPaymentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateTerminalPaymentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateTerminalPaymentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInvoiceId

`func (o *CreateTerminalPaymentRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateTerminalPaymentRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateTerminalPaymentRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreateTerminalPaymentRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetReference

`func (o *CreateTerminalPaymentRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *CreateTerminalPaymentRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *CreateTerminalPaymentRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *CreateTerminalPaymentRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


