# PortalSubscriptionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]PortalSubscription**](PortalSubscription.md) |  | 
**HasMore** | **bool** |  | 
**Url** | **string** |  | 

## Methods

### NewPortalSubscriptionList

`func NewPortalSubscriptionList(object string, data []PortalSubscription, hasMore bool, url string, ) *PortalSubscriptionList`

NewPortalSubscriptionList instantiates a new PortalSubscriptionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSubscriptionListWithDefaults

`func NewPortalSubscriptionListWithDefaults() *PortalSubscriptionList`

NewPortalSubscriptionListWithDefaults instantiates a new PortalSubscriptionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PortalSubscriptionList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalSubscriptionList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalSubscriptionList) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PortalSubscriptionList) GetData() []PortalSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PortalSubscriptionList) GetDataOk() (*[]PortalSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PortalSubscriptionList) SetData(v []PortalSubscription)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PortalSubscriptionList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PortalSubscriptionList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PortalSubscriptionList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *PortalSubscriptionList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PortalSubscriptionList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PortalSubscriptionList) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


