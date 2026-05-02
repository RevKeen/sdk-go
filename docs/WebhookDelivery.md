# WebhookDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**EndpointId** | **string** |  | 
**EventId** | **string** | The event this delivery attempted to send. | 
**EventType** | **NullableString** |  | 
**Status** | **string** | Delivery state. pending: queued or retrying. succeeded: 2xx response received. failed: last attempt failed, may retry based on attempts/max_attempts. dead_lettered: max attempts exhausted, will not retry automatically. | 
**Attempts** | **int32** |  | 
**MaxAttempts** | **int32** |  | 
**LastStatusCode** | **NullableInt32** | HTTP status code returned on the last attempt. | 
**LastError** | **NullableString** | Truncated error message from the last failed attempt. | 
**LastErrorCode** | **NullableString** |  | 
**LastDurationMs** | **NullableInt32** |  | 
**LastAttemptAt** | **NullableTime** |  | 
**NextRetryAt** | **NullableTime** |  | 
**DeliveredAt** | **NullableTime** |  | 
**DeadLetteredAt** | **NullableTime** |  | 
**DeadLetterReason** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewWebhookDelivery

`func NewWebhookDelivery(id string, object string, endpointId string, eventId string, eventType NullableString, status string, attempts int32, maxAttempts int32, lastStatusCode NullableInt32, lastError NullableString, lastErrorCode NullableString, lastDurationMs NullableInt32, lastAttemptAt NullableTime, nextRetryAt NullableTime, deliveredAt NullableTime, deadLetteredAt NullableTime, deadLetterReason NullableString, createdAt time.Time, updatedAt time.Time, ) *WebhookDelivery`

NewWebhookDelivery instantiates a new WebhookDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryWithDefaults

`func NewWebhookDeliveryWithDefaults() *WebhookDelivery`

NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDelivery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDelivery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDelivery) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *WebhookDelivery) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookDelivery) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookDelivery) SetObject(v string)`

SetObject sets Object field to given value.


### GetEndpointId

`func (o *WebhookDelivery) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *WebhookDelivery) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *WebhookDelivery) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.


### GetEventId

`func (o *WebhookDelivery) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhookDelivery) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhookDelivery) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventType

`func (o *WebhookDelivery) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookDelivery) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookDelivery) SetEventType(v string)`

SetEventType sets EventType field to given value.


### SetEventTypeNil

`func (o *WebhookDelivery) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *WebhookDelivery) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetStatus

`func (o *WebhookDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAttempts

`func (o *WebhookDelivery) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *WebhookDelivery) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *WebhookDelivery) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.


### GetMaxAttempts

`func (o *WebhookDelivery) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *WebhookDelivery) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *WebhookDelivery) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.


### GetLastStatusCode

`func (o *WebhookDelivery) GetLastStatusCode() int32`

GetLastStatusCode returns the LastStatusCode field if non-nil, zero value otherwise.

### GetLastStatusCodeOk

`func (o *WebhookDelivery) GetLastStatusCodeOk() (*int32, bool)`

GetLastStatusCodeOk returns a tuple with the LastStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusCode

`func (o *WebhookDelivery) SetLastStatusCode(v int32)`

SetLastStatusCode sets LastStatusCode field to given value.


### SetLastStatusCodeNil

`func (o *WebhookDelivery) SetLastStatusCodeNil(b bool)`

 SetLastStatusCodeNil sets the value for LastStatusCode to be an explicit nil

### UnsetLastStatusCode
`func (o *WebhookDelivery) UnsetLastStatusCode()`

UnsetLastStatusCode ensures that no value is present for LastStatusCode, not even an explicit nil
### GetLastError

`func (o *WebhookDelivery) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *WebhookDelivery) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *WebhookDelivery) SetLastError(v string)`

SetLastError sets LastError field to given value.


### SetLastErrorNil

`func (o *WebhookDelivery) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *WebhookDelivery) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetLastErrorCode

`func (o *WebhookDelivery) GetLastErrorCode() string`

GetLastErrorCode returns the LastErrorCode field if non-nil, zero value otherwise.

### GetLastErrorCodeOk

`func (o *WebhookDelivery) GetLastErrorCodeOk() (*string, bool)`

GetLastErrorCodeOk returns a tuple with the LastErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorCode

`func (o *WebhookDelivery) SetLastErrorCode(v string)`

SetLastErrorCode sets LastErrorCode field to given value.


### SetLastErrorCodeNil

`func (o *WebhookDelivery) SetLastErrorCodeNil(b bool)`

 SetLastErrorCodeNil sets the value for LastErrorCode to be an explicit nil

### UnsetLastErrorCode
`func (o *WebhookDelivery) UnsetLastErrorCode()`

UnsetLastErrorCode ensures that no value is present for LastErrorCode, not even an explicit nil
### GetLastDurationMs

`func (o *WebhookDelivery) GetLastDurationMs() int32`

GetLastDurationMs returns the LastDurationMs field if non-nil, zero value otherwise.

### GetLastDurationMsOk

`func (o *WebhookDelivery) GetLastDurationMsOk() (*int32, bool)`

GetLastDurationMsOk returns a tuple with the LastDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDurationMs

`func (o *WebhookDelivery) SetLastDurationMs(v int32)`

SetLastDurationMs sets LastDurationMs field to given value.


### SetLastDurationMsNil

`func (o *WebhookDelivery) SetLastDurationMsNil(b bool)`

 SetLastDurationMsNil sets the value for LastDurationMs to be an explicit nil

### UnsetLastDurationMs
`func (o *WebhookDelivery) UnsetLastDurationMs()`

UnsetLastDurationMs ensures that no value is present for LastDurationMs, not even an explicit nil
### GetLastAttemptAt

`func (o *WebhookDelivery) GetLastAttemptAt() time.Time`

GetLastAttemptAt returns the LastAttemptAt field if non-nil, zero value otherwise.

### GetLastAttemptAtOk

`func (o *WebhookDelivery) GetLastAttemptAtOk() (*time.Time, bool)`

GetLastAttemptAtOk returns a tuple with the LastAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttemptAt

`func (o *WebhookDelivery) SetLastAttemptAt(v time.Time)`

SetLastAttemptAt sets LastAttemptAt field to given value.


### SetLastAttemptAtNil

`func (o *WebhookDelivery) SetLastAttemptAtNil(b bool)`

 SetLastAttemptAtNil sets the value for LastAttemptAt to be an explicit nil

### UnsetLastAttemptAt
`func (o *WebhookDelivery) UnsetLastAttemptAt()`

UnsetLastAttemptAt ensures that no value is present for LastAttemptAt, not even an explicit nil
### GetNextRetryAt

`func (o *WebhookDelivery) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *WebhookDelivery) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *WebhookDelivery) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### SetNextRetryAtNil

`func (o *WebhookDelivery) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *WebhookDelivery) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil
### GetDeliveredAt

`func (o *WebhookDelivery) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *WebhookDelivery) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *WebhookDelivery) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.


### SetDeliveredAtNil

`func (o *WebhookDelivery) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *WebhookDelivery) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetDeadLetteredAt

`func (o *WebhookDelivery) GetDeadLetteredAt() time.Time`

GetDeadLetteredAt returns the DeadLetteredAt field if non-nil, zero value otherwise.

### GetDeadLetteredAtOk

`func (o *WebhookDelivery) GetDeadLetteredAtOk() (*time.Time, bool)`

GetDeadLetteredAtOk returns a tuple with the DeadLetteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetteredAt

`func (o *WebhookDelivery) SetDeadLetteredAt(v time.Time)`

SetDeadLetteredAt sets DeadLetteredAt field to given value.


### SetDeadLetteredAtNil

`func (o *WebhookDelivery) SetDeadLetteredAtNil(b bool)`

 SetDeadLetteredAtNil sets the value for DeadLetteredAt to be an explicit nil

### UnsetDeadLetteredAt
`func (o *WebhookDelivery) UnsetDeadLetteredAt()`

UnsetDeadLetteredAt ensures that no value is present for DeadLetteredAt, not even an explicit nil
### GetDeadLetterReason

`func (o *WebhookDelivery) GetDeadLetterReason() string`

GetDeadLetterReason returns the DeadLetterReason field if non-nil, zero value otherwise.

### GetDeadLetterReasonOk

`func (o *WebhookDelivery) GetDeadLetterReasonOk() (*string, bool)`

GetDeadLetterReasonOk returns a tuple with the DeadLetterReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadLetterReason

`func (o *WebhookDelivery) SetDeadLetterReason(v string)`

SetDeadLetterReason sets DeadLetterReason field to given value.


### SetDeadLetterReasonNil

`func (o *WebhookDelivery) SetDeadLetterReasonNil(b bool)`

 SetDeadLetterReasonNil sets the value for DeadLetterReason to be an explicit nil

### UnsetDeadLetterReason
`func (o *WebhookDelivery) UnsetDeadLetterReason()`

UnsetDeadLetterReason ensures that no value is present for DeadLetterReason, not even an explicit nil
### GetCreatedAt

`func (o *WebhookDelivery) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookDelivery) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookDelivery) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WebhookDelivery) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookDelivery) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookDelivery) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


