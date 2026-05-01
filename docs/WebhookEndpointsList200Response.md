# WebhookEndpointsList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]WebhookEndpointsList200ResponseDataInner**](WebhookEndpointsList200ResponseDataInner.md) |  | 
**Meta** | [**WebhookEndpointsList200ResponseMeta**](WebhookEndpointsList200ResponseMeta.md) |  | 

## Methods

### NewWebhookEndpointsList200Response

`func NewWebhookEndpointsList200Response(data []WebhookEndpointsList200ResponseDataInner, meta WebhookEndpointsList200ResponseMeta, ) *WebhookEndpointsList200Response`

NewWebhookEndpointsList200Response instantiates a new WebhookEndpointsList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointsList200ResponseWithDefaults

`func NewWebhookEndpointsList200ResponseWithDefaults() *WebhookEndpointsList200Response`

NewWebhookEndpointsList200ResponseWithDefaults instantiates a new WebhookEndpointsList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebhookEndpointsList200Response) GetData() []WebhookEndpointsList200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEndpointsList200Response) GetDataOk() (*[]WebhookEndpointsList200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEndpointsList200Response) SetData(v []WebhookEndpointsList200ResponseDataInner)`

SetData sets Data field to given value.


### GetMeta

`func (o *WebhookEndpointsList200Response) GetMeta() WebhookEndpointsList200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WebhookEndpointsList200Response) GetMetaOk() (*WebhookEndpointsList200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WebhookEndpointsList200Response) SetMeta(v WebhookEndpointsList200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


