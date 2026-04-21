# WebhookEndpointsList200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the webhook endpoint | 
**Url** | **string** | The URL that receives webhook POST requests | 
**Description** | **NullableString** | Optional human-readable description | 
**EnabledEvents** | **[]string** | Event types this endpoint subscribes to | 
**Status** | **string** | Whether the endpoint is currently receiving events | 
**CircuitBreakerState** | **string** | Circuit breaker status. Opens after repeated delivery failures. | 
**TotalDeliveries** | **float32** | Total number of delivery attempts | 
**SuccessfulDeliveries** | **float32** | Number of successful deliveries | 
**FailedDeliveries** | **float32** | Number of failed deliveries | 
**LastDeliveryAt** | **NullableString** | Timestamp of the most recent delivery attempt | 
**CreatedAt** | **string** | When the endpoint was created | 
**UpdatedAt** | **string** | When the endpoint was last modified | 

## Methods

### NewWebhookEndpointsList200ResponseDataInner

`func NewWebhookEndpointsList200ResponseDataInner(id string, url string, description NullableString, enabledEvents []string, status string, circuitBreakerState string, totalDeliveries float32, successfulDeliveries float32, failedDeliveries float32, lastDeliveryAt NullableString, createdAt string, updatedAt string, ) *WebhookEndpointsList200ResponseDataInner`

NewWebhookEndpointsList200ResponseDataInner instantiates a new WebhookEndpointsList200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointsList200ResponseDataInnerWithDefaults

`func NewWebhookEndpointsList200ResponseDataInnerWithDefaults() *WebhookEndpointsList200ResponseDataInner`

NewWebhookEndpointsList200ResponseDataInnerWithDefaults instantiates a new WebhookEndpointsList200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEndpointsList200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEndpointsList200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebhookEndpointsList200ResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpointsList200ResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *WebhookEndpointsList200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpointsList200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookEndpointsList200ResponseDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookEndpointsList200ResponseDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabledEvents

`func (o *WebhookEndpointsList200ResponseDataInner) GetEnabledEvents() []string`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetEnabledEventsOk() (*[]string, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpointsList200ResponseDataInner) SetEnabledEvents(v []string)`

SetEnabledEvents sets EnabledEvents field to given value.


### GetStatus

`func (o *WebhookEndpointsList200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpointsList200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCircuitBreakerState

`func (o *WebhookEndpointsList200ResponseDataInner) GetCircuitBreakerState() string`

GetCircuitBreakerState returns the CircuitBreakerState field if non-nil, zero value otherwise.

### GetCircuitBreakerStateOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetCircuitBreakerStateOk() (*string, bool)`

GetCircuitBreakerStateOk returns a tuple with the CircuitBreakerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerState

`func (o *WebhookEndpointsList200ResponseDataInner) SetCircuitBreakerState(v string)`

SetCircuitBreakerState sets CircuitBreakerState field to given value.


### GetTotalDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) GetTotalDeliveries() float32`

GetTotalDeliveries returns the TotalDeliveries field if non-nil, zero value otherwise.

### GetTotalDeliveriesOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetTotalDeliveriesOk() (*float32, bool)`

GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) SetTotalDeliveries(v float32)`

SetTotalDeliveries sets TotalDeliveries field to given value.


### GetSuccessfulDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) GetSuccessfulDeliveries() float32`

GetSuccessfulDeliveries returns the SuccessfulDeliveries field if non-nil, zero value otherwise.

### GetSuccessfulDeliveriesOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetSuccessfulDeliveriesOk() (*float32, bool)`

GetSuccessfulDeliveriesOk returns a tuple with the SuccessfulDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) SetSuccessfulDeliveries(v float32)`

SetSuccessfulDeliveries sets SuccessfulDeliveries field to given value.


### GetFailedDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) GetFailedDeliveries() float32`

GetFailedDeliveries returns the FailedDeliveries field if non-nil, zero value otherwise.

### GetFailedDeliveriesOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetFailedDeliveriesOk() (*float32, bool)`

GetFailedDeliveriesOk returns a tuple with the FailedDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeliveries

`func (o *WebhookEndpointsList200ResponseDataInner) SetFailedDeliveries(v float32)`

SetFailedDeliveries sets FailedDeliveries field to given value.


### GetLastDeliveryAt

`func (o *WebhookEndpointsList200ResponseDataInner) GetLastDeliveryAt() string`

GetLastDeliveryAt returns the LastDeliveryAt field if non-nil, zero value otherwise.

### GetLastDeliveryAtOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetLastDeliveryAtOk() (*string, bool)`

GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveryAt

`func (o *WebhookEndpointsList200ResponseDataInner) SetLastDeliveryAt(v string)`

SetLastDeliveryAt sets LastDeliveryAt field to given value.


### SetLastDeliveryAtNil

`func (o *WebhookEndpointsList200ResponseDataInner) SetLastDeliveryAtNil(b bool)`

 SetLastDeliveryAtNil sets the value for LastDeliveryAt to be an explicit nil

### UnsetLastDeliveryAt
`func (o *WebhookEndpointsList200ResponseDataInner) UnsetLastDeliveryAt()`

UnsetLastDeliveryAt ensures that no value is present for LastDeliveryAt, not even an explicit nil
### GetCreatedAt

`func (o *WebhookEndpointsList200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEndpointsList200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookEndpointsList200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookEndpointsList200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookEndpointsList200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


