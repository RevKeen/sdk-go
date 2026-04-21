# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique event identifier | 
**Type** | **string** | Event type using dot notation (e.g., &#x60;payment.succeeded&#x60;) | 
**Created** | **int32** | Unix timestamp when the event was created | 
**Livemode** | Pointer to **bool** | Whether this event occurred in live mode | [optional] 
**Data** | [**WebhookEventData**](WebhookEventData.md) |  | 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(id string, type_ string, created int32, data WebhookEventData, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEvent) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WebhookEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookEvent) SetType(v string)`

SetType sets Type field to given value.


### GetCreated

`func (o *WebhookEvent) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookEvent) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookEvent) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetLivemode

`func (o *WebhookEvent) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *WebhookEvent) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *WebhookEvent) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *WebhookEvent) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetData

`func (o *WebhookEvent) GetData() WebhookEventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookEvent) GetDataOk() (*WebhookEventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookEvent) SetData(v WebhookEventData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


