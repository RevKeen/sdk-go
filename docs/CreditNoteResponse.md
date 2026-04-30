# CreditNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreditNote**](CreditNote.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewCreditNoteResponse

`func NewCreditNoteResponse(data CreditNote, requestId string, ) *CreditNoteResponse`

NewCreditNoteResponse instantiates a new CreditNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteResponseWithDefaults

`func NewCreditNoteResponseWithDefaults() *CreditNoteResponse`

NewCreditNoteResponseWithDefaults instantiates a new CreditNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreditNoteResponse) GetData() CreditNote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditNoteResponse) GetDataOk() (*CreditNote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditNoteResponse) SetData(v CreditNote)`

SetData sets Data field to given value.


### GetRequestId

`func (o *CreditNoteResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreditNoteResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreditNoteResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


