# CreditNoteEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**TotalAmountMinor** | **float32** |  | 
**TotalPaidMinor** | **float32** |  | 
**TotalCreditedMinor** | **float32** |  | 
**MaxCreditableMinor** | **float32** |  | 
**Eligible** | **bool** |  | 
**Payments** | [**[]CreditNoteEligibilityResponsePaymentsInner**](CreditNoteEligibilityResponsePaymentsInner.md) |  | 

## Methods

### NewCreditNoteEligibilityResponse

`func NewCreditNoteEligibilityResponse(invoiceId string, totalAmountMinor float32, totalPaidMinor float32, totalCreditedMinor float32, maxCreditableMinor float32, eligible bool, payments []CreditNoteEligibilityResponsePaymentsInner, ) *CreditNoteEligibilityResponse`

NewCreditNoteEligibilityResponse instantiates a new CreditNoteEligibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteEligibilityResponseWithDefaults

`func NewCreditNoteEligibilityResponseWithDefaults() *CreditNoteEligibilityResponse`

NewCreditNoteEligibilityResponseWithDefaults instantiates a new CreditNoteEligibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreditNoteEligibilityResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditNoteEligibilityResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditNoteEligibilityResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetTotalAmountMinor

`func (o *CreditNoteEligibilityResponse) GetTotalAmountMinor() float32`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalAmountMinorOk() (*float32, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *CreditNoteEligibilityResponse) SetTotalAmountMinor(v float32)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetTotalPaidMinor

`func (o *CreditNoteEligibilityResponse) GetTotalPaidMinor() float32`

GetTotalPaidMinor returns the TotalPaidMinor field if non-nil, zero value otherwise.

### GetTotalPaidMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalPaidMinorOk() (*float32, bool)`

GetTotalPaidMinorOk returns a tuple with the TotalPaidMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidMinor

`func (o *CreditNoteEligibilityResponse) SetTotalPaidMinor(v float32)`

SetTotalPaidMinor sets TotalPaidMinor field to given value.


### GetTotalCreditedMinor

`func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinor() float32`

GetTotalCreditedMinor returns the TotalCreditedMinor field if non-nil, zero value otherwise.

### GetTotalCreditedMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinorOk() (*float32, bool)`

GetTotalCreditedMinorOk returns a tuple with the TotalCreditedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditedMinor

`func (o *CreditNoteEligibilityResponse) SetTotalCreditedMinor(v float32)`

SetTotalCreditedMinor sets TotalCreditedMinor field to given value.


### GetMaxCreditableMinor

`func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinor() float32`

GetMaxCreditableMinor returns the MaxCreditableMinor field if non-nil, zero value otherwise.

### GetMaxCreditableMinorOk

`func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinorOk() (*float32, bool)`

GetMaxCreditableMinorOk returns a tuple with the MaxCreditableMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreditableMinor

`func (o *CreditNoteEligibilityResponse) SetMaxCreditableMinor(v float32)`

SetMaxCreditableMinor sets MaxCreditableMinor field to given value.


### GetEligible

`func (o *CreditNoteEligibilityResponse) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *CreditNoteEligibilityResponse) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *CreditNoteEligibilityResponse) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetPayments

`func (o *CreditNoteEligibilityResponse) GetPayments() []CreditNoteEligibilityResponsePaymentsInner`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *CreditNoteEligibilityResponse) GetPaymentsOk() (*[]CreditNoteEligibilityResponsePaymentsInner, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *CreditNoteEligibilityResponse) SetPayments(v []CreditNoteEligibilityResponsePaymentsInner)`

SetPayments sets Payments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


