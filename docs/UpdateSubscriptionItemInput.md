# UpdateSubscriptionItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | Item quantity | [optional] 
**Description** | Pointer to **string** | Item description | [optional] 
**UnitAmountMinor** | Pointer to **int32** | Unit price in cents | [optional] 
**FulfillmentType** | Pointer to **string** | Determines if orders are created on renewal | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewUpdateSubscriptionItemInput

`func NewUpdateSubscriptionItemInput() *UpdateSubscriptionItemInput`

NewUpdateSubscriptionItemInput instantiates a new UpdateSubscriptionItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionItemInputWithDefaults

`func NewUpdateSubscriptionItemInputWithDefaults() *UpdateSubscriptionItemInput`

NewUpdateSubscriptionItemInputWithDefaults instantiates a new UpdateSubscriptionItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *UpdateSubscriptionItemInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateSubscriptionItemInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateSubscriptionItemInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpdateSubscriptionItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSubscriptionItemInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSubscriptionItemInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSubscriptionItemInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSubscriptionItemInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnitAmountMinor

`func (o *UpdateSubscriptionItemInput) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *UpdateSubscriptionItemInput) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *UpdateSubscriptionItemInput) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.

### HasUnitAmountMinor

`func (o *UpdateSubscriptionItemInput) HasUnitAmountMinor() bool`

HasUnitAmountMinor returns a boolean if a field has been set.

### GetFulfillmentType

`func (o *UpdateSubscriptionItemInput) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *UpdateSubscriptionItemInput) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *UpdateSubscriptionItemInput) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *UpdateSubscriptionItemInput) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSubscriptionItemInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSubscriptionItemInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSubscriptionItemInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSubscriptionItemInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


