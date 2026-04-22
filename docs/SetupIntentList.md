# SetupIntentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type, always &#39;list&#39; | 
**Data** | [**[]SetupIntent**](SetupIntent.md) | Array of setup intents | 
**HasMore** | **bool** | Whether there are more results available | 
**TotalCount** | **int32** | Total count of matching setup intents | 
**Url** | **string** | URL for this list resource | 

## Methods

### NewSetupIntentList

`func NewSetupIntentList(object string, data []SetupIntent, hasMore bool, totalCount int32, url string, ) *SetupIntentList`

NewSetupIntentList instantiates a new SetupIntentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupIntentListWithDefaults

`func NewSetupIntentListWithDefaults() *SetupIntentList`

NewSetupIntentListWithDefaults instantiates a new SetupIntentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SetupIntentList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SetupIntentList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SetupIntentList) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *SetupIntentList) GetData() []SetupIntent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SetupIntentList) GetDataOk() (*[]SetupIntent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SetupIntentList) SetData(v []SetupIntent)`

SetData sets Data field to given value.


### GetHasMore

`func (o *SetupIntentList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SetupIntentList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SetupIntentList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *SetupIntentList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SetupIntentList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SetupIntentList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetUrl

`func (o *SetupIntentList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SetupIntentList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SetupIntentList) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


