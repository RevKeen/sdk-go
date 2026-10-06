# InvoiceLineUsageEventProvenanceItemEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**EventName** | **NullableString** |  |
**EventTime** | **NullableTime** |  |
**Quantity** | **string** |  |
**IdempotencyKey** | **NullableString** |  |

## Methods

### NewInvoiceLineUsageEventProvenanceItemEvent

`func NewInvoiceLineUsageEventProvenanceItemEvent(id string, eventName NullableString, eventTime NullableTime, quantity string, idempotencyKey NullableString, ) *InvoiceLineUsageEventProvenanceItemEvent`

NewInvoiceLineUsageEventProvenanceItemEvent instantiates a new InvoiceLineUsageEventProvenanceItemEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageEventProvenanceItemEventWithDefaults

`func NewInvoiceLineUsageEventProvenanceItemEventWithDefaults() *InvoiceLineUsageEventProvenanceItemEvent`

NewInvoiceLineUsageEventProvenanceItemEventWithDefaults instantiates a new InvoiceLineUsageEventProvenanceItemEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetId(v string)`

SetId sets Id field to given value.


### GetEventName

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### SetEventNameNil

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetEventNameNil(b bool)`

 SetEventNameNil sets the value for EventName to be an explicit nil

### UnsetEventName
`func (o *InvoiceLineUsageEventProvenanceItemEvent) UnsetEventName()`

UnsetEventName ensures that no value is present for EventName, not even an explicit nil
### GetEventTime

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### SetEventTimeNil

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetEventTimeNil(b bool)`

 SetEventTimeNil sets the value for EventTime to be an explicit nil

### UnsetEventTime
`func (o *InvoiceLineUsageEventProvenanceItemEvent) UnsetEventTime()`

UnsetEventTime ensures that no value is present for EventTime, not even an explicit nil
### GetQuantity

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetIdempotencyKey

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *InvoiceLineUsageEventProvenanceItemEvent) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### SetIdempotencyKeyNil

`func (o *InvoiceLineUsageEventProvenanceItemEvent) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *InvoiceLineUsageEventProvenanceItemEvent) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
