# WebhookEndpointsRedeliverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The event to send again: the &#x60;event_id&#x60; of one of its webhook deliveries (see &#x60;GET /webhook-deliveries&#x60;). |

## Methods

### NewWebhookEndpointsRedeliverRequest

`func NewWebhookEndpointsRedeliverRequest(eventId string, ) *WebhookEndpointsRedeliverRequest`

NewWebhookEndpointsRedeliverRequest instantiates a new WebhookEndpointsRedeliverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointsRedeliverRequestWithDefaults

`func NewWebhookEndpointsRedeliverRequestWithDefaults() *WebhookEndpointsRedeliverRequest`

NewWebhookEndpointsRedeliverRequestWithDefaults instantiates a new WebhookEndpointsRedeliverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *WebhookEndpointsRedeliverRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookEndpointsRedeliverRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookEndpointsRedeliverRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
