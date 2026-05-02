# CreditNoteLineList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]CreditNoteLine**](CreditNoteLine.md) |  | 
**HasMore** | **bool** |  | 
**Url** | **string** |  | 

## Methods

### NewCreditNoteLineList

`func NewCreditNoteLineList(object string, data []CreditNoteLine, hasMore bool, url string, ) *CreditNoteLineList`

NewCreditNoteLineList instantiates a new CreditNoteLineList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteLineListWithDefaults

`func NewCreditNoteLineListWithDefaults() *CreditNoteLineList`

NewCreditNoteLineListWithDefaults instantiates a new CreditNoteLineList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CreditNoteLineList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreditNoteLineList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreditNoteLineList) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *CreditNoteLineList) GetData() []CreditNoteLine`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditNoteLineList) GetDataOk() (*[]CreditNoteLine, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditNoteLineList) SetData(v []CreditNoteLine)`

SetData sets Data field to given value.


### GetHasMore

`func (o *CreditNoteLineList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CreditNoteLineList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CreditNoteLineList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *CreditNoteLineList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreditNoteLineList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreditNoteLineList) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


