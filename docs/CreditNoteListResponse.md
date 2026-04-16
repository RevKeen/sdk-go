# CreditNoteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreditNote**](CreditNote.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewCreditNoteListResponse

`func NewCreditNoteListResponse(data []CreditNote, requestId string, ) *CreditNoteListResponse`

NewCreditNoteListResponse instantiates a new CreditNoteListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteListResponseWithDefaults

`func NewCreditNoteListResponseWithDefaults() *CreditNoteListResponse`

NewCreditNoteListResponseWithDefaults instantiates a new CreditNoteListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreditNoteListResponse) GetData() []CreditNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditNoteListResponse) GetDataOk() (*[]CreditNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditNoteListResponse) SetData(v []CreditNote)`

SetData sets Data field to given value.


### GetRequestId

`func (o *CreditNoteListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreditNoteListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreditNoteListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


