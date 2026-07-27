# UsageEventRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Usage event ID | 
**MerchantId** | **string** | Owning merchant ID | 
**CustomerId** | **NullableString** | Customer ID | 
**SubscriptionId** | **NullableString** | Subscription ID | 
**MeterId** | **NullableString** | Meter ID | 
**Quantity** | **float32** | Event quantity | 
**EventTime** | **time.Time** | Event timestamp (ISO 8601) | 
**IdempotencyKey** | **NullableString** | Deduplication key | 
**ExternalId** | **NullableString** | External event identifier | 
**Source** | **string** | Ingestion source | 
**Metadata** | **map[string]interface{}** | Arbitrary event metadata | 
**Properties** | **map[string]interface{}** | Arbitrary event properties used for filtering/aggregation | 
**IngestionTimestamp** | **NullableTime** | When the event was ingested (ISO 8601) | 
**CreatedAt** | **time.Time** | Creation timestamp (ISO 8601) | 

## Methods

### NewUsageEventRecord

`func NewUsageEventRecord(id string, merchantId string, customerId NullableString, subscriptionId NullableString, meterId NullableString, quantity float32, eventTime time.Time, idempotencyKey NullableString, externalId NullableString, source string, metadata map[string]interface{}, properties map[string]interface{}, ingestionTimestamp NullableTime, createdAt time.Time, ) *UsageEventRecord`

NewUsageEventRecord instantiates a new UsageEventRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEventRecordWithDefaults

`func NewUsageEventRecordWithDefaults() *UsageEventRecord`

NewUsageEventRecordWithDefaults instantiates a new UsageEventRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsageEventRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageEventRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageEventRecord) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *UsageEventRecord) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UsageEventRecord) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UsageEventRecord) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *UsageEventRecord) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UsageEventRecord) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UsageEventRecord) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *UsageEventRecord) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *UsageEventRecord) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetSubscriptionId

`func (o *UsageEventRecord) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UsageEventRecord) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UsageEventRecord) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *UsageEventRecord) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *UsageEventRecord) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetMeterId

`func (o *UsageEventRecord) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageEventRecord) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageEventRecord) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### SetMeterIdNil

`func (o *UsageEventRecord) SetMeterIdNil(b bool)`

 SetMeterIdNil sets the value for MeterId to be an explicit nil

### UnsetMeterId
`func (o *UsageEventRecord) UnsetMeterId()`

UnsetMeterId ensures that no value is present for MeterId, not even an explicit nil
### GetQuantity

`func (o *UsageEventRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageEventRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageEventRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetEventTime

`func (o *UsageEventRecord) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *UsageEventRecord) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *UsageEventRecord) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetIdempotencyKey

`func (o *UsageEventRecord) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UsageEventRecord) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UsageEventRecord) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### SetIdempotencyKeyNil

`func (o *UsageEventRecord) SetIdempotencyKeyNil(b bool)`

 SetIdempotencyKeyNil sets the value for IdempotencyKey to be an explicit nil

### UnsetIdempotencyKey
`func (o *UsageEventRecord) UnsetIdempotencyKey()`

UnsetIdempotencyKey ensures that no value is present for IdempotencyKey, not even an explicit nil
### GetExternalId

`func (o *UsageEventRecord) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UsageEventRecord) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UsageEventRecord) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *UsageEventRecord) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UsageEventRecord) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSource

`func (o *UsageEventRecord) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageEventRecord) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageEventRecord) SetSource(v string)`

SetSource sets Source field to given value.


### GetMetadata

`func (o *UsageEventRecord) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UsageEventRecord) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UsageEventRecord) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *UsageEventRecord) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UsageEventRecord) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UsageEventRecord) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetIngestionTimestamp

`func (o *UsageEventRecord) GetIngestionTimestamp() time.Time`

GetIngestionTimestamp returns the IngestionTimestamp field if non-nil, zero value otherwise.

### GetIngestionTimestampOk

`func (o *UsageEventRecord) GetIngestionTimestampOk() (*time.Time, bool)`

GetIngestionTimestampOk returns a tuple with the IngestionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestionTimestamp

`func (o *UsageEventRecord) SetIngestionTimestamp(v time.Time)`

SetIngestionTimestamp sets IngestionTimestamp field to given value.


### SetIngestionTimestampNil

`func (o *UsageEventRecord) SetIngestionTimestampNil(b bool)`

 SetIngestionTimestampNil sets the value for IngestionTimestamp to be an explicit nil

### UnsetIngestionTimestamp
`func (o *UsageEventRecord) UnsetIngestionTimestamp()`

UnsetIngestionTimestamp ensures that no value is present for IngestionTimestamp, not even an explicit nil
### GetCreatedAt

`func (o *UsageEventRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsageEventRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsageEventRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


