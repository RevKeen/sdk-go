# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique event identifier | 
**Object** | **string** | Object type | 
**Type** | **string** | Event type (e.g., invoice.paid, payment.succeeded) | 
**Data** | [**EventData**](EventData.md) |  | 
**Request** | [**EventRequest**](EventRequest.md) |  | 
**PendingWebhooks** | **int32** | Number of webhooks still pending delivery for this event | 
**ApiVersion** | **NullableString** | API version used to render this event | 
**Livemode** | **bool** | Whether this event was produced in live mode | 
**Created** | **int32** | Unix timestamp when the event was created | 

## Methods

### NewEvent

`func NewEvent(id string, object string, type_ string, data EventData, request EventRequest, pendingWebhooks int32, apiVersion NullableString, livemode bool, created int32, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Event) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Event) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Event) SetObject(v string)`

SetObject sets Object field to given value.


### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *Event) GetData() EventData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Event) GetDataOk() (*EventData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Event) SetData(v EventData)`

SetData sets Data field to given value.


### GetRequest

`func (o *Event) GetRequest() EventRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Event) GetRequestOk() (*EventRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Event) SetRequest(v EventRequest)`

SetRequest sets Request field to given value.


### GetPendingWebhooks

`func (o *Event) GetPendingWebhooks() int32`

GetPendingWebhooks returns the PendingWebhooks field if non-nil, zero value otherwise.

### GetPendingWebhooksOk

`func (o *Event) GetPendingWebhooksOk() (*int32, bool)`

GetPendingWebhooksOk returns a tuple with the PendingWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWebhooks

`func (o *Event) SetPendingWebhooks(v int32)`

SetPendingWebhooks sets PendingWebhooks field to given value.


### GetApiVersion

`func (o *Event) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Event) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Event) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### SetApiVersionNil

`func (o *Event) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *Event) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetLivemode

`func (o *Event) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *Event) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *Event) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.


### GetCreated

`func (o *Event) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Event) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Event) SetCreated(v int32)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


