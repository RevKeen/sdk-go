# UsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Event name matching a meter&#39;s event_name | 
**CustomerId** | Pointer to **string** | RevKeen customer ID | [optional] 
**ExternalCustomerId** | Pointer to **string** | Your external customer identifier | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription ID to attribute usage to | [optional] 
**MeterId** | Pointer to **string** | Direct meter ID (alternative to event name matching) | [optional] 
**Quantity** | Pointer to **float32** | Event quantity (default: 1) | [optional] [default to 1]
**Timestamp** | Pointer to **string** | ISO 8601 timestamp (default: now) | [optional] 
**IdempotencyKey** | Pointer to **string** | Unique key for deduplication | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary properties for filtering and aggregation | [optional] 

## Methods

### NewUsageEvent

`func NewUsageEvent(name string, ) *UsageEvent`

NewUsageEvent instantiates a new UsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEventWithDefaults

`func NewUsageEventWithDefaults() *UsageEvent`

NewUsageEventWithDefaults instantiates a new UsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsageEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageEvent) SetName(v string)`

SetName sets Name field to given value.


### GetCustomerId

`func (o *UsageEvent) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UsageEvent) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UsageEvent) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UsageEvent) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *UsageEvent) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *UsageEvent) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *UsageEvent) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *UsageEvent) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UsageEvent) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UsageEvent) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UsageEvent) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UsageEvent) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetMeterId

`func (o *UsageEvent) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageEvent) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageEvent) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *UsageEvent) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetQuantity

`func (o *UsageEvent) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageEvent) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageEvent) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageEvent) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTimestamp

`func (o *UsageEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsageEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsageEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UsageEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *UsageEvent) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UsageEvent) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UsageEvent) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *UsageEvent) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *UsageEvent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageEvent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageEvent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageEvent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


