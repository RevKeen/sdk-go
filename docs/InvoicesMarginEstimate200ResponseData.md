# InvoicesMarginEstimate200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** |  | 
**AmountMinor** | **int32** |  | 
**Currency** | **string** |  | 
**Estimates** | [**[]InvoicesMarginEstimate200ResponseDataEstimatesInner**](InvoicesMarginEstimate200ResponseDataEstimatesInner.md) |  | 

## Methods

### NewInvoicesMarginEstimate200ResponseData

`func NewInvoicesMarginEstimate200ResponseData(invoiceId string, amountMinor int32, currency string, estimates []InvoicesMarginEstimate200ResponseDataEstimatesInner, ) *InvoicesMarginEstimate200ResponseData`

NewInvoicesMarginEstimate200ResponseData instantiates a new InvoicesMarginEstimate200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesMarginEstimate200ResponseDataWithDefaults

`func NewInvoicesMarginEstimate200ResponseDataWithDefaults() *InvoicesMarginEstimate200ResponseData`

NewInvoicesMarginEstimate200ResponseDataWithDefaults instantiates a new InvoicesMarginEstimate200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *InvoicesMarginEstimate200ResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoicesMarginEstimate200ResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoicesMarginEstimate200ResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMinor

`func (o *InvoicesMarginEstimate200ResponseData) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoicesMarginEstimate200ResponseData) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoicesMarginEstimate200ResponseData) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *InvoicesMarginEstimate200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicesMarginEstimate200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicesMarginEstimate200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEstimates

`func (o *InvoicesMarginEstimate200ResponseData) GetEstimates() []InvoicesMarginEstimate200ResponseDataEstimatesInner`

GetEstimates returns the Estimates field if non-nil, zero value otherwise.

### GetEstimatesOk

`func (o *InvoicesMarginEstimate200ResponseData) GetEstimatesOk() (*[]InvoicesMarginEstimate200ResponseDataEstimatesInner, bool)`

GetEstimatesOk returns a tuple with the Estimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimates

`func (o *InvoicesMarginEstimate200ResponseData) SetEstimates(v []InvoicesMarginEstimate200ResponseDataEstimatesInner)`

SetEstimates sets Estimates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


