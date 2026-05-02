# UpdateMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Updated meter name | [optional] 
**Slug** | Pointer to **string** | Updated slug | [optional] 
**Description** | Pointer to **string** | Updated description | [optional] 
**Status** | Pointer to **string** | Meter status | [optional] 
**ValueKey** | Pointer to **string** | Updated value key | [optional] 
**FilterConditions** | Pointer to **[]map[string]interface{}** | Updated filter conditions | [optional] 
**UniqueCountKey** | Pointer to **string** | Updated unique count key | [optional] 
**UnitName** | Pointer to **string** | Updated unit name | [optional] 
**CarryForward** | Pointer to **bool** | Updated carry forward setting | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Updated metadata | [optional] 
**AlertThresholds** | Pointer to **[]float32** | Updated threshold percentages for usage.threshold.reached webhooks. | [optional] 

## Methods

### NewUpdateMeterRequest

`func NewUpdateMeterRequest() *UpdateMeterRequest`

NewUpdateMeterRequest instantiates a new UpdateMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeterRequestWithDefaults

`func NewUpdateMeterRequestWithDefaults() *UpdateMeterRequest`

NewUpdateMeterRequestWithDefaults instantiates a new UpdateMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMeterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMeterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMeterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMeterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *UpdateMeterRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UpdateMeterRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UpdateMeterRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UpdateMeterRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateMeterRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMeterRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMeterRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMeterRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateMeterRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateMeterRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateMeterRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateMeterRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValueKey

`func (o *UpdateMeterRequest) GetValueKey() string`

GetValueKey returns the ValueKey field if non-nil, zero value otherwise.

### GetValueKeyOk

`func (o *UpdateMeterRequest) GetValueKeyOk() (*string, bool)`

GetValueKeyOk returns a tuple with the ValueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueKey

`func (o *UpdateMeterRequest) SetValueKey(v string)`

SetValueKey sets ValueKey field to given value.

### HasValueKey

`func (o *UpdateMeterRequest) HasValueKey() bool`

HasValueKey returns a boolean if a field has been set.

### GetFilterConditions

`func (o *UpdateMeterRequest) GetFilterConditions() []map[string]interface{}`

GetFilterConditions returns the FilterConditions field if non-nil, zero value otherwise.

### GetFilterConditionsOk

`func (o *UpdateMeterRequest) GetFilterConditionsOk() (*[]map[string]interface{}, bool)`

GetFilterConditionsOk returns a tuple with the FilterConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterConditions

`func (o *UpdateMeterRequest) SetFilterConditions(v []map[string]interface{})`

SetFilterConditions sets FilterConditions field to given value.

### HasFilterConditions

`func (o *UpdateMeterRequest) HasFilterConditions() bool`

HasFilterConditions returns a boolean if a field has been set.

### GetUniqueCountKey

`func (o *UpdateMeterRequest) GetUniqueCountKey() string`

GetUniqueCountKey returns the UniqueCountKey field if non-nil, zero value otherwise.

### GetUniqueCountKeyOk

`func (o *UpdateMeterRequest) GetUniqueCountKeyOk() (*string, bool)`

GetUniqueCountKeyOk returns a tuple with the UniqueCountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCountKey

`func (o *UpdateMeterRequest) SetUniqueCountKey(v string)`

SetUniqueCountKey sets UniqueCountKey field to given value.

### HasUniqueCountKey

`func (o *UpdateMeterRequest) HasUniqueCountKey() bool`

HasUniqueCountKey returns a boolean if a field has been set.

### GetUnitName

`func (o *UpdateMeterRequest) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *UpdateMeterRequest) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *UpdateMeterRequest) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *UpdateMeterRequest) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### GetCarryForward

`func (o *UpdateMeterRequest) GetCarryForward() bool`

GetCarryForward returns the CarryForward field if non-nil, zero value otherwise.

### GetCarryForwardOk

`func (o *UpdateMeterRequest) GetCarryForwardOk() (*bool, bool)`

GetCarryForwardOk returns a tuple with the CarryForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryForward

`func (o *UpdateMeterRequest) SetCarryForward(v bool)`

SetCarryForward sets CarryForward field to given value.

### HasCarryForward

`func (o *UpdateMeterRequest) HasCarryForward() bool`

HasCarryForward returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateMeterRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateMeterRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateMeterRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateMeterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAlertThresholds

`func (o *UpdateMeterRequest) GetAlertThresholds() []float32`

GetAlertThresholds returns the AlertThresholds field if non-nil, zero value otherwise.

### GetAlertThresholdsOk

`func (o *UpdateMeterRequest) GetAlertThresholdsOk() (*[]float32, bool)`

GetAlertThresholdsOk returns a tuple with the AlertThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThresholds

`func (o *UpdateMeterRequest) SetAlertThresholds(v []float32)`

SetAlertThresholds sets AlertThresholds field to given value.

### HasAlertThresholds

`func (o *UpdateMeterRequest) HasAlertThresholds() bool`

HasAlertThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


