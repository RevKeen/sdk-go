# WebhookEndpointsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The URL to receive webhook POST requests. Must be HTTPS in production. | 
**EnabledEvents** | **[]string** | Event types to subscribe to. At least one event is required. | 
**Description** | Pointer to **string** | Optional human-readable description (max 500 characters) | [optional] 

## Methods

### NewWebhookEndpointsCreateRequest

`func NewWebhookEndpointsCreateRequest(url string, enabledEvents []string, ) *WebhookEndpointsCreateRequest`

NewWebhookEndpointsCreateRequest instantiates a new WebhookEndpointsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointsCreateRequestWithDefaults

`func NewWebhookEndpointsCreateRequestWithDefaults() *WebhookEndpointsCreateRequest`

NewWebhookEndpointsCreateRequestWithDefaults instantiates a new WebhookEndpointsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookEndpointsCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpointsCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpointsCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabledEvents

`func (o *WebhookEndpointsCreateRequest) GetEnabledEvents() []string`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpointsCreateRequest) GetEnabledEventsOk() (*[]string, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpointsCreateRequest) SetEnabledEvents(v []string)`

SetEnabledEvents sets EnabledEvents field to given value.


### GetDescription

`func (o *WebhookEndpointsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpointsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpointsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookEndpointsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


