# InvoicesFinalizeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAdvance** | Pointer to **bool** | If true, automatically transitions to &#39;approved&#39; status after finalization. | [optional] 

## Methods

### NewInvoicesFinalizeRequest

`func NewInvoicesFinalizeRequest() *InvoicesFinalizeRequest`

NewInvoicesFinalizeRequest instantiates a new InvoicesFinalizeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesFinalizeRequestWithDefaults

`func NewInvoicesFinalizeRequestWithDefaults() *InvoicesFinalizeRequest`

NewInvoicesFinalizeRequestWithDefaults instantiates a new InvoicesFinalizeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAdvance

`func (o *InvoicesFinalizeRequest) GetAutoAdvance() bool`

GetAutoAdvance returns the AutoAdvance field if non-nil, zero value otherwise.

### GetAutoAdvanceOk

`func (o *InvoicesFinalizeRequest) GetAutoAdvanceOk() (*bool, bool)`

GetAutoAdvanceOk returns a tuple with the AutoAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAdvance

`func (o *InvoicesFinalizeRequest) SetAutoAdvance(v bool)`

SetAutoAdvance sets AutoAdvance field to given value.

### HasAutoAdvance

`func (o *InvoicesFinalizeRequest) HasAutoAdvance() bool`

HasAutoAdvance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


