# SubscriptionScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]SubscriptionSchedule**](SubscriptionSchedule.md) |  | 
**HasMore** | **bool** |  | 
**Url** | **string** | The URL for accessing this list | 

## Methods

### NewSubscriptionScheduleListResponse

`func NewSubscriptionScheduleListResponse(object string, data []SubscriptionSchedule, hasMore bool, url string, ) *SubscriptionScheduleListResponse`

NewSubscriptionScheduleListResponse instantiates a new SubscriptionScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionScheduleListResponseWithDefaults

`func NewSubscriptionScheduleListResponseWithDefaults() *SubscriptionScheduleListResponse`

NewSubscriptionScheduleListResponseWithDefaults instantiates a new SubscriptionScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SubscriptionScheduleListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubscriptionScheduleListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubscriptionScheduleListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *SubscriptionScheduleListResponse) GetData() []SubscriptionSchedule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionScheduleListResponse) GetDataOk() (*[]SubscriptionSchedule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionScheduleListResponse) SetData(v []SubscriptionSchedule)`

SetData sets Data field to given value.


### GetHasMore

`func (o *SubscriptionScheduleListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SubscriptionScheduleListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SubscriptionScheduleListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *SubscriptionScheduleListResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SubscriptionScheduleListResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SubscriptionScheduleListResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


