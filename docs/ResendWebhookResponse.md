# ResendWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the webhook was queued for resending | 
**WebhookEndpointId** | Pointer to **string** | The webhook endpoint ID the event was sent to | [optional] 
**Message** | Pointer to **string** | Additional information | [optional] 

## Methods

### NewResendWebhookResponse

`func NewResendWebhookResponse(success bool, ) *ResendWebhookResponse`

NewResendWebhookResponse instantiates a new ResendWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResendWebhookResponseWithDefaults

`func NewResendWebhookResponseWithDefaults() *ResendWebhookResponse`

NewResendWebhookResponseWithDefaults instantiates a new ResendWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResendWebhookResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResendWebhookResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResendWebhookResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetWebhookEndpointId

`func (o *ResendWebhookResponse) GetWebhookEndpointId() string`

GetWebhookEndpointId returns the WebhookEndpointId field if non-nil, zero value otherwise.

### GetWebhookEndpointIdOk

`func (o *ResendWebhookResponse) GetWebhookEndpointIdOk() (*string, bool)`

GetWebhookEndpointIdOk returns a tuple with the WebhookEndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointId

`func (o *ResendWebhookResponse) SetWebhookEndpointId(v string)`

SetWebhookEndpointId sets WebhookEndpointId field to given value.

### HasWebhookEndpointId

`func (o *ResendWebhookResponse) HasWebhookEndpointId() bool`

HasWebhookEndpointId returns a boolean if a field has been set.

### GetMessage

`func (o *ResendWebhookResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResendWebhookResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResendWebhookResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResendWebhookResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


