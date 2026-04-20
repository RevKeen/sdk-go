# CreditNoteLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreditNoteId** | **string** |  | 
**InvoiceLineItemId** | **NullableString** | The invoice line item this credit line applies to, when the credit was scoped to a specific line. Null for invoice-level credits. | 
**Description** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitAmountMinor** | **int32** | Per-unit credit amount in cents | 
**TotalAmountMinor** | **int32** | Total credit for this line in cents | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewCreditNoteLine

`func NewCreditNoteLine(id string, object string, creditNoteId string, invoiceLineItemId NullableString, description string, quantity int32, unitAmountMinor int32, totalAmountMinor int32, metadata map[string]interface{}, createdAt time.Time, ) *CreditNoteLine`

NewCreditNoteLine instantiates a new CreditNoteLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteLineWithDefaults

`func NewCreditNoteLineWithDefaults() *CreditNoteLine`

NewCreditNoteLineWithDefaults instantiates a new CreditNoteLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreditNoteLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreditNoteLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreditNoteLine) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CreditNoteLine) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreditNoteLine) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreditNoteLine) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreditNoteId

`func (o *CreditNoteLine) GetCreditNoteId() string`

GetCreditNoteId returns the CreditNoteId field if non-nil, zero value otherwise.

### GetCreditNoteIdOk

`func (o *CreditNoteLine) GetCreditNoteIdOk() (*string, bool)`

GetCreditNoteIdOk returns a tuple with the CreditNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteId

`func (o *CreditNoteLine) SetCreditNoteId(v string)`

SetCreditNoteId sets CreditNoteId field to given value.


### GetInvoiceLineItemId

`func (o *CreditNoteLine) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *CreditNoteLine) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *CreditNoteLine) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.


### SetInvoiceLineItemIdNil

`func (o *CreditNoteLine) SetInvoiceLineItemIdNil(b bool)`

 SetInvoiceLineItemIdNil sets the value for InvoiceLineItemId to be an explicit nil

### UnsetInvoiceLineItemId
`func (o *CreditNoteLine) UnsetInvoiceLineItemId()`

UnsetInvoiceLineItemId ensures that no value is present for InvoiceLineItemId, not even an explicit nil
### GetDescription

`func (o *CreditNoteLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditNoteLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditNoteLine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *CreditNoteLine) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreditNoteLine) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreditNoteLine) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmountMinor

`func (o *CreditNoteLine) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *CreditNoteLine) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *CreditNoteLine) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetTotalAmountMinor

`func (o *CreditNoteLine) GetTotalAmountMinor() int32`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *CreditNoteLine) GetTotalAmountMinorOk() (*int32, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *CreditNoteLine) SetTotalAmountMinor(v int32)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetMetadata

`func (o *CreditNoteLine) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreditNoteLine) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreditNoteLine) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *CreditNoteLine) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreditNoteLine) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreditNoteLine) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


