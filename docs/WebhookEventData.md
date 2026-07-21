# WebhookEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **map[string]interface{}** | The object that triggered the event (e.g., Payment, Invoice, Subscription) | [optional] 
**PreviousAttributes** | Pointer to **map[string]interface{}** | Previous values of attributes that changed (for update events) | [optional] 

## Methods

### NewWebhookEventData

`func NewWebhookEventData() *WebhookEventData`

NewWebhookEventData instantiates a new WebhookEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventDataWithDefaults

`func NewWebhookEventDataWithDefaults() *WebhookEventData`

NewWebhookEventDataWithDefaults instantiates a new WebhookEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *WebhookEventData) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookEventData) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookEventData) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookEventData) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPreviousAttributes

`func (o *WebhookEventData) GetPreviousAttributes() map[string]interface{}`

GetPreviousAttributes returns the PreviousAttributes field if non-nil, zero value otherwise.

### GetPreviousAttributesOk

`func (o *WebhookEventData) GetPreviousAttributesOk() (*map[string]interface{}, bool)`

GetPreviousAttributesOk returns a tuple with the PreviousAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAttributes

`func (o *WebhookEventData) SetPreviousAttributes(v map[string]interface{})`

SetPreviousAttributes sets PreviousAttributes field to given value.

### HasPreviousAttributes

`func (o *WebhookEventData) HasPreviousAttributes() bool`

HasPreviousAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


