# CreditNoteLineAllocationInputV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceLineItemId** | **string** | Invoice line item receiving the credit |
**Quantity** | **int32** | Credited quantity from the source invoice line |
**UnitAmountMinor** | **int64** | Credited unit amount in minor units |
**TotalAmountMinor** | **int64** | Pre-tax credited subtotal in minor units |
**TaxAmountMinor** | **int64** | Credited tax amount in minor units |

## Methods

### NewCreditNoteLineAllocationInputV1

`func NewCreditNoteLineAllocationInputV1(invoiceLineItemId string, quantity int32, unitAmountMinor int64, totalAmountMinor int64, taxAmountMinor int64, ) *CreditNoteLineAllocationInputV1`

NewCreditNoteLineAllocationInputV1 instantiates a new CreditNoteLineAllocationInputV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteLineAllocationInputV1WithDefaults

`func NewCreditNoteLineAllocationInputV1WithDefaults() *CreditNoteLineAllocationInputV1`

NewCreditNoteLineAllocationInputV1WithDefaults instantiates a new CreditNoteLineAllocationInputV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceLineItemId

`func (o *CreditNoteLineAllocationInputV1) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *CreditNoteLineAllocationInputV1) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *CreditNoteLineAllocationInputV1) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.


### GetQuantity

`func (o *CreditNoteLineAllocationInputV1) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreditNoteLineAllocationInputV1) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreditNoteLineAllocationInputV1) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmountMinor

`func (o *CreditNoteLineAllocationInputV1) GetUnitAmountMinor() int64`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *CreditNoteLineAllocationInputV1) GetUnitAmountMinorOk() (*int64, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *CreditNoteLineAllocationInputV1) SetUnitAmountMinor(v int64)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetTotalAmountMinor

`func (o *CreditNoteLineAllocationInputV1) GetTotalAmountMinor() int64`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *CreditNoteLineAllocationInputV1) GetTotalAmountMinorOk() (*int64, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *CreditNoteLineAllocationInputV1) SetTotalAmountMinor(v int64)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetTaxAmountMinor

`func (o *CreditNoteLineAllocationInputV1) GetTaxAmountMinor() int64`

GetTaxAmountMinor returns the TaxAmountMinor field if non-nil, zero value otherwise.

### GetTaxAmountMinorOk

`func (o *CreditNoteLineAllocationInputV1) GetTaxAmountMinorOk() (*int64, bool)`

GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountMinor

`func (o *CreditNoteLineAllocationInputV1) SetTaxAmountMinor(v int64)`

SetTaxAmountMinor sets TaxAmountMinor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
