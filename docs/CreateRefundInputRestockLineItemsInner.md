# CreateRefundInputRestockLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceLineItemId** | **string** | An invoice line of the refunded payment |
**Quantity** | **int32** | Units of that line to put back on sale |

## Methods

### NewCreateRefundInputRestockLineItemsInner

`func NewCreateRefundInputRestockLineItemsInner(invoiceLineItemId string, quantity int32, ) *CreateRefundInputRestockLineItemsInner`

NewCreateRefundInputRestockLineItemsInner instantiates a new CreateRefundInputRestockLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRefundInputRestockLineItemsInnerWithDefaults

`func NewCreateRefundInputRestockLineItemsInnerWithDefaults() *CreateRefundInputRestockLineItemsInner`

NewCreateRefundInputRestockLineItemsInnerWithDefaults instantiates a new CreateRefundInputRestockLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceLineItemId

`func (o *CreateRefundInputRestockLineItemsInner) GetInvoiceLineItemId() string`

GetInvoiceLineItemId returns the InvoiceLineItemId field if non-nil, zero value otherwise.

### GetInvoiceLineItemIdOk

`func (o *CreateRefundInputRestockLineItemsInner) GetInvoiceLineItemIdOk() (*string, bool)`

GetInvoiceLineItemIdOk returns a tuple with the InvoiceLineItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLineItemId

`func (o *CreateRefundInputRestockLineItemsInner) SetInvoiceLineItemId(v string)`

SetInvoiceLineItemId sets InvoiceLineItemId field to given value.


### GetQuantity

`func (o *CreateRefundInputRestockLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateRefundInputRestockLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateRefundInputRestockLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
