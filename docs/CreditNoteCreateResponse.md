# CreditNoteCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreditNote**](CreditNote.md) | A credit note formally reduces the amount owed on an invoice. It is the accounting counterpart of a refund — the credit note is the document, the refund is the money movement. | 

## Methods

### NewCreditNoteCreateResponse

`func NewCreditNoteCreateResponse(data CreditNote, ) *CreditNoteCreateResponse`

NewCreditNoteCreateResponse instantiates a new CreditNoteCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteCreateResponseWithDefaults

`func NewCreditNoteCreateResponseWithDefaults() *CreditNoteCreateResponse`

NewCreditNoteCreateResponseWithDefaults instantiates a new CreditNoteCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreditNoteCreateResponse) GetData() CreditNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditNoteCreateResponse) GetDataOk() (*CreditNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditNoteCreateResponse) SetData(v CreditNote)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


