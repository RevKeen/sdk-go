# CreditEligibilityResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanCredit** | **bool** | Whether a credit note can be issued for this invoice |
**MaxCreditableMinor** | **int64** | Maximum amount that can be credited in minor units |
**TotalCreditedMinor** | **int64** | Total amount already credited against this invoice |
**InvoiceTotalMinor** | **int64** | Original invoice total in minor units |
**Reason** | Pointer to **string** | Reason why the invoice cannot be credited, if applicable | [optional]

## Methods

### NewCreditEligibilityResponseData

`func NewCreditEligibilityResponseData(canCredit bool, maxCreditableMinor int64, totalCreditedMinor int64, invoiceTotalMinor int64, ) *CreditEligibilityResponseData`

NewCreditEligibilityResponseData instantiates a new CreditEligibilityResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditEligibilityResponseDataWithDefaults

`func NewCreditEligibilityResponseDataWithDefaults() *CreditEligibilityResponseData`

NewCreditEligibilityResponseDataWithDefaults instantiates a new CreditEligibilityResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanCredit

`func (o *CreditEligibilityResponseData) GetCanCredit() bool`

GetCanCredit returns the CanCredit field if non-nil, zero value otherwise.

### GetCanCreditOk

`func (o *CreditEligibilityResponseData) GetCanCreditOk() (*bool, bool)`

GetCanCreditOk returns a tuple with the CanCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanCredit

`func (o *CreditEligibilityResponseData) SetCanCredit(v bool)`

SetCanCredit sets CanCredit field to given value.


### GetMaxCreditableMinor

`func (o *CreditEligibilityResponseData) GetMaxCreditableMinor() int64`

GetMaxCreditableMinor returns the MaxCreditableMinor field if non-nil, zero value otherwise.

### GetMaxCreditableMinorOk

`func (o *CreditEligibilityResponseData) GetMaxCreditableMinorOk() (*int64, bool)`

GetMaxCreditableMinorOk returns a tuple with the MaxCreditableMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCreditableMinor

`func (o *CreditEligibilityResponseData) SetMaxCreditableMinor(v int64)`

SetMaxCreditableMinor sets MaxCreditableMinor field to given value.


### GetTotalCreditedMinor

`func (o *CreditEligibilityResponseData) GetTotalCreditedMinor() int64`

GetTotalCreditedMinor returns the TotalCreditedMinor field if non-nil, zero value otherwise.

### GetTotalCreditedMinorOk

`func (o *CreditEligibilityResponseData) GetTotalCreditedMinorOk() (*int64, bool)`

GetTotalCreditedMinorOk returns a tuple with the TotalCreditedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCreditedMinor

`func (o *CreditEligibilityResponseData) SetTotalCreditedMinor(v int64)`

SetTotalCreditedMinor sets TotalCreditedMinor field to given value.


### GetInvoiceTotalMinor

`func (o *CreditEligibilityResponseData) GetInvoiceTotalMinor() int64`

GetInvoiceTotalMinor returns the InvoiceTotalMinor field if non-nil, zero value otherwise.

### GetInvoiceTotalMinorOk

`func (o *CreditEligibilityResponseData) GetInvoiceTotalMinorOk() (*int64, bool)`

GetInvoiceTotalMinorOk returns a tuple with the InvoiceTotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceTotalMinor

`func (o *CreditEligibilityResponseData) SetInvoiceTotalMinor(v int64)`

SetInvoiceTotalMinor sets InvoiceTotalMinor field to given value.


### GetReason

`func (o *CreditEligibilityResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreditEligibilityResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreditEligibilityResponseData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreditEligibilityResponseData) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
