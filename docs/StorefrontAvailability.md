# StorefrontAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Remaining** | **NullableInt32** |  | 
**DisplayMode** | **string** |  | 
**LowStockThreshold** | **NullableInt32** |  | 

## Methods

### NewStorefrontAvailability

`func NewStorefrontAvailability(status string, remaining NullableInt32, displayMode string, lowStockThreshold NullableInt32, ) *StorefrontAvailability`

NewStorefrontAvailability instantiates a new StorefrontAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontAvailabilityWithDefaults

`func NewStorefrontAvailabilityWithDefaults() *StorefrontAvailability`

NewStorefrontAvailabilityWithDefaults instantiates a new StorefrontAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StorefrontAvailability) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorefrontAvailability) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorefrontAvailability) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRemaining

`func (o *StorefrontAvailability) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *StorefrontAvailability) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *StorefrontAvailability) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### SetRemainingNil

`func (o *StorefrontAvailability) SetRemainingNil(b bool)`

 SetRemainingNil sets the value for Remaining to be an explicit nil

### UnsetRemaining
`func (o *StorefrontAvailability) UnsetRemaining()`

UnsetRemaining ensures that no value is present for Remaining, not even an explicit nil
### GetDisplayMode

`func (o *StorefrontAvailability) GetDisplayMode() string`

GetDisplayMode returns the DisplayMode field if non-nil, zero value otherwise.

### GetDisplayModeOk

`func (o *StorefrontAvailability) GetDisplayModeOk() (*string, bool)`

GetDisplayModeOk returns a tuple with the DisplayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMode

`func (o *StorefrontAvailability) SetDisplayMode(v string)`

SetDisplayMode sets DisplayMode field to given value.


### GetLowStockThreshold

`func (o *StorefrontAvailability) GetLowStockThreshold() int32`

GetLowStockThreshold returns the LowStockThreshold field if non-nil, zero value otherwise.

### GetLowStockThresholdOk

`func (o *StorefrontAvailability) GetLowStockThresholdOk() (*int32, bool)`

GetLowStockThresholdOk returns a tuple with the LowStockThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowStockThreshold

`func (o *StorefrontAvailability) SetLowStockThreshold(v int32)`

SetLowStockThreshold sets LowStockThreshold field to given value.


### SetLowStockThresholdNil

`func (o *StorefrontAvailability) SetLowStockThresholdNil(b bool)`

 SetLowStockThresholdNil sets the value for LowStockThreshold to be an explicit nil

### UnsetLowStockThreshold
`func (o *StorefrontAvailability) UnsetLowStockThreshold()`

UnsetLowStockThreshold ensures that no value is present for LowStockThreshold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


