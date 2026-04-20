# PortalInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**InvoiceNumber** | **NullableString** |  | 
**Status** | **string** |  | 
**Currency** | **NullableString** |  | 
**Subtotal** | **NullableInt32** |  | 
**TaxAmount** | **NullableInt32** |  | 
**DiscountAmount** | **NullableInt32** |  | 
**Total** | **int32** |  | 
**AmountPaid** | **NullableInt32** |  | 
**AmountDue** | **NullableInt32** |  | 
**InvoiceDate** | **NullableTime** |  | 
**DueDate** | **NullableTime** |  | 
**PaidAt** | **NullableTime** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPortalInvoice

`func NewPortalInvoice(id string, object string, invoiceNumber NullableString, status string, currency NullableString, subtotal NullableInt32, taxAmount NullableInt32, discountAmount NullableInt32, total int32, amountPaid NullableInt32, amountDue NullableInt32, invoiceDate NullableTime, dueDate NullableTime, paidAt NullableTime, createdAt time.Time, ) *PortalInvoice`

NewPortalInvoice instantiates a new PortalInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalInvoiceWithDefaults

`func NewPortalInvoiceWithDefaults() *PortalInvoice`

NewPortalInvoiceWithDefaults instantiates a new PortalInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalInvoice) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PortalInvoice) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalInvoice) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalInvoice) SetObject(v string)`

SetObject sets Object field to given value.


### GetInvoiceNumber

`func (o *PortalInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *PortalInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *PortalInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### SetInvoiceNumberNil

`func (o *PortalInvoice) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *PortalInvoice) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetStatus

`func (o *PortalInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortalInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortalInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *PortalInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PortalInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PortalInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *PortalInvoice) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PortalInvoice) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetSubtotal

`func (o *PortalInvoice) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *PortalInvoice) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *PortalInvoice) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### SetSubtotalNil

`func (o *PortalInvoice) SetSubtotalNil(b bool)`

 SetSubtotalNil sets the value for Subtotal to be an explicit nil

### UnsetSubtotal
`func (o *PortalInvoice) UnsetSubtotal()`

UnsetSubtotal ensures that no value is present for Subtotal, not even an explicit nil
### GetTaxAmount

`func (o *PortalInvoice) GetTaxAmount() int32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PortalInvoice) GetTaxAmountOk() (*int32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PortalInvoice) SetTaxAmount(v int32)`

SetTaxAmount sets TaxAmount field to given value.


### SetTaxAmountNil

`func (o *PortalInvoice) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *PortalInvoice) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetDiscountAmount

`func (o *PortalInvoice) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *PortalInvoice) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *PortalInvoice) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.


### SetDiscountAmountNil

`func (o *PortalInvoice) SetDiscountAmountNil(b bool)`

 SetDiscountAmountNil sets the value for DiscountAmount to be an explicit nil

### UnsetDiscountAmount
`func (o *PortalInvoice) UnsetDiscountAmount()`

UnsetDiscountAmount ensures that no value is present for DiscountAmount, not even an explicit nil
### GetTotal

`func (o *PortalInvoice) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PortalInvoice) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PortalInvoice) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetAmountPaid

`func (o *PortalInvoice) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *PortalInvoice) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *PortalInvoice) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### SetAmountPaidNil

`func (o *PortalInvoice) SetAmountPaidNil(b bool)`

 SetAmountPaidNil sets the value for AmountPaid to be an explicit nil

### UnsetAmountPaid
`func (o *PortalInvoice) UnsetAmountPaid()`

UnsetAmountPaid ensures that no value is present for AmountPaid, not even an explicit nil
### GetAmountDue

`func (o *PortalInvoice) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *PortalInvoice) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *PortalInvoice) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### SetAmountDueNil

`func (o *PortalInvoice) SetAmountDueNil(b bool)`

 SetAmountDueNil sets the value for AmountDue to be an explicit nil

### UnsetAmountDue
`func (o *PortalInvoice) UnsetAmountDue()`

UnsetAmountDue ensures that no value is present for AmountDue, not even an explicit nil
### GetInvoiceDate

`func (o *PortalInvoice) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *PortalInvoice) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *PortalInvoice) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.


### SetInvoiceDateNil

`func (o *PortalInvoice) SetInvoiceDateNil(b bool)`

 SetInvoiceDateNil sets the value for InvoiceDate to be an explicit nil

### UnsetInvoiceDate
`func (o *PortalInvoice) UnsetInvoiceDate()`

UnsetInvoiceDate ensures that no value is present for InvoiceDate, not even an explicit nil
### GetDueDate

`func (o *PortalInvoice) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PortalInvoice) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PortalInvoice) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.


### SetDueDateNil

`func (o *PortalInvoice) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *PortalInvoice) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaidAt

`func (o *PortalInvoice) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *PortalInvoice) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *PortalInvoice) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.


### SetPaidAtNil

`func (o *PortalInvoice) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *PortalInvoice) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetCreatedAt

`func (o *PortalInvoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalInvoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalInvoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


