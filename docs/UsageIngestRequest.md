# UsageIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Event name matching a meter&#39;s event_name | 
**CustomerId** | Pointer to **string** | RevKeen customer ID. Prefer external_customer_id unless your system already stores RevKeen customer UUIDs. | [optional] 
**ExternalCustomerId** | Pointer to **string** | Your external customer identifier. Recommended attribution field; unknown values fail rather than creating orphan usage. | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription ID to attribute usage to. Usage is assumed to share the subscription currency; RevKeen does not perform FX conversion in metering. | [optional] 
**MeterId** | Pointer to **string** | Direct meter ID (alternative to event name matching) | [optional] 
**Quantity** | Pointer to **float32** | Event quantity (default: 1). For percentage-priced meters, send the signed transaction amount in the subscription currency&#39;s minor units. | [optional] [default to 1]
**Timestamp** | Pointer to **string** | ISO 8601 timestamp (default: now) | [optional] 
**IdempotencyKey** | Pointer to **string** | Merchant-scoped deduplication key. Reusing the same key returns duplicate status and does not create another billable event. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary properties for filtering and aggregation | [optional] 
**Events** | [**[]UsageEvent**](UsageEvent.md) | Array of usage events (1–1000) | 

## Methods

### NewUsageIngestRequest

`func NewUsageIngestRequest(name string, events []UsageEvent, ) *UsageIngestRequest`

NewUsageIngestRequest instantiates a new UsageIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageIngestRequestWithDefaults

`func NewUsageIngestRequestWithDefaults() *UsageIngestRequest`

NewUsageIngestRequestWithDefaults instantiates a new UsageIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsageIngestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageIngestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageIngestRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCustomerId

`func (o *UsageIngestRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UsageIngestRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UsageIngestRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UsageIngestRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *UsageIngestRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *UsageIngestRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *UsageIngestRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *UsageIngestRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UsageIngestRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UsageIngestRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UsageIngestRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UsageIngestRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetMeterId

`func (o *UsageIngestRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageIngestRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageIngestRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *UsageIngestRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetQuantity

`func (o *UsageIngestRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageIngestRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageIngestRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageIngestRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTimestamp

`func (o *UsageIngestRequest) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsageIngestRequest) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsageIngestRequest) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UsageIngestRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *UsageIngestRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UsageIngestRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UsageIngestRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *UsageIngestRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *UsageIngestRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageIngestRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageIngestRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UsageIngestRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEvents

`func (o *UsageIngestRequest) GetEvents() []UsageEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UsageIngestRequest) GetEventsOk() (*[]UsageEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UsageIngestRequest) SetEvents(v []UsageEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


