# WebhookDeliveryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]WebhookDelivery**](WebhookDelivery.md) |  | 
**HasMore** | **bool** |  | 
**Url** | **string** |  | 

## Methods

### NewWebhookDeliveryListResponse

`func NewWebhookDeliveryListResponse(object string, data []WebhookDelivery, hasMore bool, url string, ) *WebhookDeliveryListResponse`

NewWebhookDeliveryListResponse instantiates a new WebhookDeliveryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryListResponseWithDefaults

`func NewWebhookDeliveryListResponseWithDefaults() *WebhookDeliveryListResponse`

NewWebhookDeliveryListResponseWithDefaults instantiates a new WebhookDeliveryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *WebhookDeliveryListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookDeliveryListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookDeliveryListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *WebhookDeliveryListResponse) GetData() []WebhookDelivery`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookDeliveryListResponse) GetDataOk() (*[]WebhookDelivery, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookDeliveryListResponse) SetData(v []WebhookDelivery)`

SetData sets Data field to given value.


### GetHasMore

`func (o *WebhookDeliveryListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *WebhookDeliveryListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *WebhookDeliveryListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *WebhookDeliveryListResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookDeliveryListResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookDeliveryListResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


