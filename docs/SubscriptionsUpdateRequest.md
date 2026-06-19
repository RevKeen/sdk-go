# SubscriptionsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**AmountMinor** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**NextBillingDate** | Pointer to **string** |  | [optional] 
**BillingAnchorDay** | Pointer to **int32** |  | [optional] 
**PauseAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriptionsUpdateRequest

`func NewSubscriptionsUpdateRequest() *SubscriptionsUpdateRequest`

NewSubscriptionsUpdateRequest instantiates a new SubscriptionsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsUpdateRequestWithDefaults

`func NewSubscriptionsUpdateRequestWithDefaults() *SubscriptionsUpdateRequest`

NewSubscriptionsUpdateRequestWithDefaults instantiates a new SubscriptionsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SubscriptionsUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionsUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionsUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionsUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAmountMinor

`func (o *SubscriptionsUpdateRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *SubscriptionsUpdateRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *SubscriptionsUpdateRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *SubscriptionsUpdateRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionsUpdateRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionsUpdateRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionsUpdateRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionsUpdateRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNextBillingDate

`func (o *SubscriptionsUpdateRequest) GetNextBillingDate() string`

GetNextBillingDate returns the NextBillingDate field if non-nil, zero value otherwise.

### GetNextBillingDateOk

`func (o *SubscriptionsUpdateRequest) GetNextBillingDateOk() (*string, bool)`

GetNextBillingDateOk returns a tuple with the NextBillingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBillingDate

`func (o *SubscriptionsUpdateRequest) SetNextBillingDate(v string)`

SetNextBillingDate sets NextBillingDate field to given value.

### HasNextBillingDate

`func (o *SubscriptionsUpdateRequest) HasNextBillingDate() bool`

HasNextBillingDate returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *SubscriptionsUpdateRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *SubscriptionsUpdateRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *SubscriptionsUpdateRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *SubscriptionsUpdateRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### GetPauseAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) GetPauseAtPeriodEnd() bool`

GetPauseAtPeriodEnd returns the PauseAtPeriodEnd field if non-nil, zero value otherwise.

### GetPauseAtPeriodEndOk

`func (o *SubscriptionsUpdateRequest) GetPauseAtPeriodEndOk() (*bool, bool)`

GetPauseAtPeriodEndOk returns a tuple with the PauseAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) SetPauseAtPeriodEnd(v bool)`

SetPauseAtPeriodEnd sets PauseAtPeriodEnd field to given value.

### HasPauseAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) HasPauseAtPeriodEnd() bool`

HasPauseAtPeriodEnd returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *SubscriptionsUpdateRequest) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *SubscriptionsUpdateRequest) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionsUpdateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionsUpdateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionsUpdateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionsUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


