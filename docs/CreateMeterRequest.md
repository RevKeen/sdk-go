# CreateMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Human-readable meter name | 
**EventName** | **string** | Event name to match against | 
**Aggregation** | Pointer to **string** | Aggregation method | [optional] [default to "count"]
**Slug** | Pointer to **string** | URL-safe identifier | [optional] 
**Description** | Pointer to **string** | Meter description | [optional] 
**ValueKey** | Pointer to **string** | Property key for sum/max/last aggregations | [optional] 
**FilterConditions** | Pointer to **[]map[string]interface{}** | Filter conditions for matching events | [optional] 
**UniqueCountKey** | Pointer to **string** | Property key for count_unique aggregation | [optional] 
**UnitName** | Pointer to **string** | Display unit name | [optional] 
**CarryForward** | Pointer to **bool** | Carry forward last value across periods | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata | [optional] 
**AlertThresholds** | Pointer to **[]float32** | Threshold percentages that trigger usage.threshold.reached webhooks. | [optional] 

## Methods

### NewCreateMeterRequest

`func NewCreateMeterRequest(name string, eventName string, ) *CreateMeterRequest`

NewCreateMeterRequest instantiates a new CreateMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMeterRequestWithDefaults

`func NewCreateMeterRequestWithDefaults() *CreateMeterRequest`

NewCreateMeterRequestWithDefaults instantiates a new CreateMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMeterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMeterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMeterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEventName

`func (o *CreateMeterRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateMeterRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateMeterRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetAggregation

`func (o *CreateMeterRequest) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *CreateMeterRequest) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *CreateMeterRequest) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *CreateMeterRequest) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetSlug

`func (o *CreateMeterRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateMeterRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateMeterRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateMeterRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *CreateMeterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMeterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMeterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMeterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueKey

`func (o *CreateMeterRequest) GetValueKey() string`

GetValueKey returns the ValueKey field if non-nil, zero value otherwise.

### GetValueKeyOk

`func (o *CreateMeterRequest) GetValueKeyOk() (*string, bool)`

GetValueKeyOk returns a tuple with the ValueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueKey

`func (o *CreateMeterRequest) SetValueKey(v string)`

SetValueKey sets ValueKey field to given value.

### HasValueKey

`func (o *CreateMeterRequest) HasValueKey() bool`

HasValueKey returns a boolean if a field has been set.

### GetFilterConditions

`func (o *CreateMeterRequest) GetFilterConditions() []map[string]interface{}`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *CreateMeterRequest) GetFilterConditionsOk() (*[]map[string]interface{}, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *CreateMeterRequest) SetFilterConditions(v []map[string]interface{})`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *CreateMeterRequest) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### GetUniqueCountKey

`func (o *CreateMeterRequest) GetUniqueCountKey() string`

GetUniqueCountKey returns the UniqueCountKey field if non-nil, zero value otherwise.

### GetUniqueCountKeyOk

`func (o *CreateMeterRequest) GetUniqueCountKeyOk() (*string, bool)`

GetUniqueCountKeyOk returns a tuple with the UniqueCountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCountKey

`func (o *CreateMeterRequest) SetUniqueCountKey(v string)`

SetUniqueCountKey sets UniqueCountKey field to given value.

### HasUniqueCountKey

`func (o *CreateMeterRequest) HasUniqueCountKey() bool`

HasUniqueCountKey returns a boolean if a field has been set.

### GetUnitName

`func (o *CreateMeterRequest) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *CreateMeterRequest) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *CreateMeterRequest) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *CreateMeterRequest) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetCarryForward

`func (o *CreateMeterRequest) GetCarryForward() bool`

GetCarryForward returns the CarryForward field if non-nil, zero value otherwise.

### GetCarryForwardOk

`func (o *CreateMeterRequest) GetCarryForwardOk() (*bool, bool)`

GetCarryForwardOk returns a tuple with the CarryForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryForward

`func (o *CreateMeterRequest) SetCarryForward(v bool)`

SetCarryForward sets CarryForward field to given value.

### HasCarryForward

`func (o *CreateMeterRequest) HasCarryForward() bool`

HasCarryForward returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateMeterRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateMeterRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateMeterRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateMeterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAlertThresholds

`func (o *CreateMeterRequest) GetAlertThresholds() []float32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *CreateMeterRequest) GetAlertThresholdsOk() (*[]float32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *CreateMeterRequest) SetAlertThresholds(v []float32)`

SetAlertThresholds sets AlertThresholds field to given value.

### HasAlertThresholds

`func (o *CreateMeterRequest) HasAlertThresholds() bool`

HasAlertThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


