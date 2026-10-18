# CreditNoteListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CreditNote**](CreditNote.md) |  |
**Meta** | [**DiscountListResponseMeta**](DiscountListResponseMeta.md) |  |

## Methods

### NewCreditNoteListResponse

`func NewCreditNoteListResponse(data []CreditNote, meta DiscountListResponseMeta, ) *CreditNoteListResponse`

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


### GetMeta

`func (o *CreditNoteListResponse) GetMeta() DiscountListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CreditNoteListResponse) GetMetaOk() (*DiscountListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CreditNoteListResponse) SetMeta(v DiscountListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
