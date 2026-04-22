# WebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the webhook endpoint | 
**Url** | **string** | The URL that receives webhook POST requests | 
**Description** | **NullableString** | Optional human-readable description | 
**EnabledEvents** | **[]string** | Event types this endpoint subscribes to | 
**Status** | **string** | Whether the endpoint is currently receiving events | 
**Secret** | Pointer to **string** | Signing secret. Returned only on create or secret rotation. | [optional] 
**CircuitBreakerState** | **string** | Circuit breaker status. Opens after repeated delivery failures. | 
**TotalDeliveries** | **float32** | Total number of delivery attempts | 
**SuccessfulDeliveries** | **float32** | Number of successful deliveries | 
**FailedDeliveries** | **float32** | Number of failed deliveries | 
**LastDeliveryAt** | **NullableString** | Timestamp of the most recent delivery attempt | 
**CreatedAt** | **string** | When the endpoint was created | 
**UpdatedAt** | **string** | When the endpoint was last modified | 

## Methods

### NewWebhookEndpoint

`func NewWebhookEndpoint(id string, url string, description NullableString, enabledEvents []string, status string, circuitBreakerState string, totalDeliveries float32, successfulDeliveries float32, failedDeliveries float32, lastDeliveryAt NullableString, createdAt string, updatedAt string, ) *WebhookEndpoint`

NewWebhookEndpoint instantiates a new WebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEndpointWithDefaults

`func NewWebhookEndpointWithDefaults() *WebhookEndpoint`

NewWebhookEndpointWithDefaults instantiates a new WebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *WebhookEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *WebhookEndpoint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookEndpoint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookEndpoint) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookEndpoint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookEndpoint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabledEvents

`func (o *WebhookEndpoint) GetEnabledEvents() []string`

GetEnabledEvents returns the EnabledEvents field if non-nil, zero value otherwise.

### GetEnabledEventsOk

`func (o *WebhookEndpoint) GetEnabledEventsOk() (*[]string, bool)`

GetEnabledEventsOk returns a tuple with the EnabledEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEvents

`func (o *WebhookEndpoint) SetEnabledEvents(v []string)`

SetEnabledEvents sets EnabledEvents field to given value.


### GetStatus

`func (o *WebhookEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSecret

`func (o *WebhookEndpoint) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookEndpoint) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookEndpoint) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookEndpoint) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCircuitBreakerState

`func (o *WebhookEndpoint) GetCircuitBreakerState() string`

GetCircuitBreakerState returns the CircuitBreakerState field if non-nil, zero value otherwise.

### GetCircuitBreakerStateOk

`func (o *WebhookEndpoint) GetCircuitBreakerStateOk() (*string, bool)`

GetCircuitBreakerStateOk returns a tuple with the CircuitBreakerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerState

`func (o *WebhookEndpoint) SetCircuitBreakerState(v string)`

SetCircuitBreakerState sets CircuitBreakerState field to given value.


### GetTotalDeliveries

`func (o *WebhookEndpoint) GetTotalDeliveries() float32`

GetTotalDeliveries returns the TotalDeliveries field if non-nil, zero value otherwise.

### GetTotalDeliveriesOk

`func (o *WebhookEndpoint) GetTotalDeliveriesOk() (*float32, bool)`

GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeliveries

`func (o *WebhookEndpoint) SetTotalDeliveries(v float32)`

SetTotalDeliveries sets TotalDeliveries field to given value.


### GetSuccessfulDeliveries

`func (o *WebhookEndpoint) GetSuccessfulDeliveries() float32`

GetSuccessfulDeliveries returns the SuccessfulDeliveries field if non-nil, zero value otherwise.

### GetSuccessfulDeliveriesOk

`func (o *WebhookEndpoint) GetSuccessfulDeliveriesOk() (*float32, bool)`

GetSuccessfulDeliveriesOk returns a tuple with the SuccessfulDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulDeliveries

`func (o *WebhookEndpoint) SetSuccessfulDeliveries(v float32)`

SetSuccessfulDeliveries sets SuccessfulDeliveries field to given value.


### GetFailedDeliveries

`func (o *WebhookEndpoint) GetFailedDeliveries() float32`

GetFailedDeliveries returns the FailedDeliveries field if non-nil, zero value otherwise.

### GetFailedDeliveriesOk

`func (o *WebhookEndpoint) GetFailedDeliveriesOk() (*float32, bool)`

GetFailedDeliveriesOk returns a tuple with the FailedDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDeliveries

`func (o *WebhookEndpoint) SetFailedDeliveries(v float32)`

SetFailedDeliveries sets FailedDeliveries field to given value.


### GetLastDeliveryAt

`func (o *WebhookEndpoint) GetLastDeliveryAt() string`

GetLastDeliveryAt returns the LastDeliveryAt field if non-nil, zero value otherwise.

### GetLastDeliveryAtOk

`func (o *WebhookEndpoint) GetLastDeliveryAtOk() (*string, bool)`

GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveryAt

`func (o *WebhookEndpoint) SetLastDeliveryAt(v string)`

SetLastDeliveryAt sets LastDeliveryAt field to given value.


### SetLastDeliveryAtNil

`func (o *WebhookEndpoint) SetLastDeliveryAtNil(b bool)`

 SetLastDeliveryAtNil sets the value for LastDeliveryAt to be an explicit nil

### UnsetLastDeliveryAt
`func (o *WebhookEndpoint) UnsetLastDeliveryAt()`

UnsetLastDeliveryAt ensures that no value is present for LastDeliveryAt, not even an explicit nil
### GetCreatedAt

`func (o *WebhookEndpoint) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookEndpoint) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookEndpoint) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookEndpoint) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookEndpoint) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookEndpoint) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


