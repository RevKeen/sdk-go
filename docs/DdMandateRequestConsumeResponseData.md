# DdMandateRequestConsumeResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**MandateId** | **NullableString** |  | 
**ConsumedAt** | **NullableString** |  | 
**Status** | **string** |  | 
**Idempotent** | Pointer to **bool** |  | [optional] 
**PdfDocuments** | Pointer to [**[]DdMandatePdfReference**](DdMandatePdfReference.md) |  | [optional] 

## Methods

### NewDdMandateRequestConsumeResponseData

`func NewDdMandateRequestConsumeResponseData(id string, customerId string, invoiceId NullableString, mandateId NullableString, consumedAt NullableString, status string, ) *DdMandateRequestConsumeResponseData`

NewDdMandateRequestConsumeResponseData instantiates a new DdMandateRequestConsumeResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdMandateRequestConsumeResponseDataWithDefaults

`func NewDdMandateRequestConsumeResponseDataWithDefaults() *DdMandateRequestConsumeResponseData`

NewDdMandateRequestConsumeResponseDataWithDefaults instantiates a new DdMandateRequestConsumeResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdMandateRequestConsumeResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdMandateRequestConsumeResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdMandateRequestConsumeResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *DdMandateRequestConsumeResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdMandateRequestConsumeResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdMandateRequestConsumeResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *DdMandateRequestConsumeResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DdMandateRequestConsumeResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DdMandateRequestConsumeResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *DdMandateRequestConsumeResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *DdMandateRequestConsumeResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetMandateId

`func (o *DdMandateRequestConsumeResponseData) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DdMandateRequestConsumeResponseData) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DdMandateRequestConsumeResponseData) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### SetMandateIdNil

`func (o *DdMandateRequestConsumeResponseData) SetMandateIdNil(b bool)`

 SetMandateIdNil sets the value for MandateId to be an explicit nil

### UnsetMandateId
`func (o *DdMandateRequestConsumeResponseData) UnsetMandateId()`

UnsetMandateId ensures that no value is present for MandateId, not even an explicit nil
### GetConsumedAt

`func (o *DdMandateRequestConsumeResponseData) GetConsumedAt() string`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *DdMandateRequestConsumeResponseData) GetConsumedAtOk() (*string, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *DdMandateRequestConsumeResponseData) SetConsumedAt(v string)`

SetConsumedAt sets ConsumedAt field to given value.


### SetConsumedAtNil

`func (o *DdMandateRequestConsumeResponseData) SetConsumedAtNil(b bool)`

 SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil

### UnsetConsumedAt
`func (o *DdMandateRequestConsumeResponseData) UnsetConsumedAt()`

UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
### GetStatus

`func (o *DdMandateRequestConsumeResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdMandateRequestConsumeResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdMandateRequestConsumeResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdempotent

`func (o *DdMandateRequestConsumeResponseData) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *DdMandateRequestConsumeResponseData) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *DdMandateRequestConsumeResponseData) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *DdMandateRequestConsumeResponseData) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.

### GetPdfDocuments

`func (o *DdMandateRequestConsumeResponseData) GetPdfDocuments() []DdMandatePdfReference`

GetPdfDocuments returns the PdfDocuments field if non-nil, zero value otherwise.

### GetPdfDocumentsOk

`func (o *DdMandateRequestConsumeResponseData) GetPdfDocumentsOk() (*[]DdMandatePdfReference, bool)`

GetPdfDocumentsOk returns a tuple with the PdfDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfDocuments

`func (o *DdMandateRequestConsumeResponseData) SetPdfDocuments(v []DdMandatePdfReference)`

SetPdfDocuments sets PdfDocuments field to given value.

### HasPdfDocuments

`func (o *DdMandateRequestConsumeResponseData) HasPdfDocuments() bool`

HasPdfDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


