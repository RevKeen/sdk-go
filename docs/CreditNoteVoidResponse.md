# CreditNoteVoidResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreditNote**](CreditNote.md) | A credit note formally reduces the amount owed on an invoice. It is the accounting counterpart of a refund — the credit note is the document, the refund is the money movement. | 

## Methods

### NewCreditNoteVoidResponse

`func NewCreditNoteVoidResponse(data CreditNote, ) *CreditNoteVoidResponse`

NewCreditNoteVoidResponse instantiates a new CreditNoteVoidResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteVoidResponseWithDefaults

`func NewCreditNoteVoidResponseWithDefaults() *CreditNoteVoidResponse`

NewCreditNoteVoidResponseWithDefaults instantiates a new CreditNoteVoidResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreditNoteVoidResponse) GetData() CreditNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditNoteVoidResponse) GetDataOk() (*CreditNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditNoteVoidResponse) SetData(v CreditNote)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


