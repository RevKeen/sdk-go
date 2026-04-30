# UsageEventListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | **[]map[string]interface{}** | Array of usage event objects | 
**HasMore** | **bool** |  | 

## Methods

### NewUsageEventListResponse

`func NewUsageEventListResponse(object string, data []map[string]interface{}, hasMore bool, ) *UsageEventListResponse`

NewUsageEventListResponse instantiates a new UsageEventListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEventListResponseWithDefaults

`func NewUsageEventListResponseWithDefaults() *UsageEventListResponse`

NewUsageEventListResponseWithDefaults instantiates a new UsageEventListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *UsageEventListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsageEventListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsageEventListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *UsageEventListResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageEventListResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageEventListResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.


### GetHasMore

`func (o *UsageEventListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UsageEventListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UsageEventListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


