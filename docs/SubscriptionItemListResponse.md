# SubscriptionItemListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionItem**](SubscriptionItem.md) |  | 
**Meta** | [**SubscriptionItemListResponseMeta**](SubscriptionItemListResponseMeta.md) |  | 

## Methods

### NewSubscriptionItemListResponse

`func NewSubscriptionItemListResponse(data []SubscriptionItem, meta SubscriptionItemListResponseMeta, ) *SubscriptionItemListResponse`

NewSubscriptionItemListResponse instantiates a new SubscriptionItemListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemListResponseWithDefaults

`func NewSubscriptionItemListResponseWithDefaults() *SubscriptionItemListResponse`

NewSubscriptionItemListResponseWithDefaults instantiates a new SubscriptionItemListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionItemListResponse) GetData() []SubscriptionItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionItemListResponse) GetDataOk() (*[]SubscriptionItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionItemListResponse) SetData(v []SubscriptionItem)`

SetData sets Data field to given value.


### GetMeta

`func (o *SubscriptionItemListResponse) GetMeta() SubscriptionItemListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SubscriptionItemListResponse) GetMetaOk() (*SubscriptionItemListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SubscriptionItemListResponse) SetMeta(v SubscriptionItemListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


