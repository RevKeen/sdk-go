# CreditNoteEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  |
**TotalAmountMinor** | **int64** |  |
**TotalPaidMinor** | **int64** |  |
**TotalCreditedMinor** | **int64** |  |
**MaxCreditableMinor** | **int64** |  |
**Eligible** | **bool** |  |
**Payments** | [**[]CreditNoteEligibilityResponsePaymentsInner**](CreditNoteEligibilityResponsePaymentsInner.md) |  |

## Methods

### NewCreditNoteEligibilityResponse

`func NewCreditNoteEligibilityResponse(invoiceId string, totalAmountMinor int64, totalPaidMinor int64, totalCreditedMinor int64, maxCreditableMinor int64, eligible bool, payments []CreditNoteEligibilityResponsePaymentsInner, ) *CreditNoteEligibilityResponse`

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

`func (o *CreditNoteEligibilityResponse) GetTotalAmountMinor() int64`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalAmountMinorOk() (*int64, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *CreditNoteEligibilityResponse) SetTotalAmountMinor(v int64)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetTotalPaidMinor

`func (o *CreditNoteEligibilityResponse) GetTotalPaidMinor() int64`

GetTotalPaidMinor returns the TotalPaidMinor field if non-nil, zero value otherwise.

### GetTotalPaidMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalPaidMinorOk() (*int64, bool)`

GetTotalPaidMinorOk returns a tuple with the TotalPaidMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPaidMinor

`func (o *CreditNoteEligibilityResponse) SetTotalPaidMinor(v int64)`

SetTotalPaidMinor sets TotalPaidMinor field to given value.


### GetTotalCreditedMinor

`func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinor() int64`

GetTotalCreditedMinor returns the TotalCreditedMinor field if non-nil, zero value otherwise.

### GetTotalCreditedMinorOk

`func (o *CreditNoteEligibilityResponse) GetTotalCreditedMinorOk() (*int64, bool)`

GetTotalCreditedMinorOk returns a tuple with the TotalCreditedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditedMinor

`func (o *CreditNoteEligibilityResponse) SetTotalCreditedMinor(v int64)`

SetTotalCreditedMinor sets TotalCreditedMinor field to given value.


### GetMaxCreditableMinor

`func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinor() int64`

GetMaxCreditableMinor returns the MaxCreditableMinor field if non-nil, zero value otherwise.

### GetMaxCreditableMinorOk

`func (o *CreditNoteEligibilityResponse) GetMaxCreditableMinorOk() (*int64, bool)`

GetMaxCreditableMinorOk returns a tuple with the MaxCreditableMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreditableMinor

`func (o *CreditNoteEligibilityResponse) SetMaxCreditableMinor(v int64)`

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
