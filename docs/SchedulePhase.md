# SchedulePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PhaseItem**](PhaseItem.md) | Products/prices included in this phase | 
**StartDate** | **string** | When this phase starts (ISO 8601 date) | 
**EndDate** | Pointer to **NullableString** | When this phase ends (null for indefinite) | [optional] 
**ProrationBehavior** | Pointer to **string** | How to handle prorations when entering this phase | [optional] 
**BillingCycleAnchor** | Pointer to **NullableString** | Override billing cycle anchor for this phase | [optional] 
**DefaultPaymentMethod** | Pointer to **NullableString** | Payment method to use for this phase | [optional] 
**CollectionMethod** | Pointer to **string** | How to collect payment for this phase | [optional] 
**Coupon** | Pointer to **NullableString** | Coupon code to apply during this phase | [optional] 
**TrialEnd** | Pointer to **NullableString** | End of trial period (if applicable) | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Phase-specific metadata | [optional] 

## Methods

### NewSchedulePhase

`func NewSchedulePhase(items []PhaseItem, startDate string, ) *SchedulePhase`

NewSchedulePhase instantiates a new SchedulePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePhaseWithDefaults

`func NewSchedulePhaseWithDefaults() *SchedulePhase`

NewSchedulePhaseWithDefaults instantiates a new SchedulePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SchedulePhase) GetItems() []PhaseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SchedulePhase) GetItemsOk() (*[]PhaseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SchedulePhase) SetItems(v []PhaseItem)`

SetItems sets Items field to given value.


### GetStartDate

`func (o *SchedulePhase) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SchedulePhase) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SchedulePhase) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *SchedulePhase) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SchedulePhase) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SchedulePhase) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SchedulePhase) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *SchedulePhase) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *SchedulePhase) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetProrationBehavior

`func (o *SchedulePhase) GetProrationBehavior() string`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *SchedulePhase) GetProrationBehaviorOk() (*string, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *SchedulePhase) SetProrationBehavior(v string)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *SchedulePhase) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *SchedulePhase) GetBillingCycleAnchor() string`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *SchedulePhase) GetBillingCycleAnchorOk() (*string, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *SchedulePhase) SetBillingCycleAnchor(v string)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *SchedulePhase) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### SetBillingCycleAnchorNil

`func (o *SchedulePhase) SetBillingCycleAnchorNil(b bool)`

 SetBillingCycleAnchorNil sets the value for BillingCycleAnchor to be an explicit nil

### UnsetBillingCycleAnchor
`func (o *SchedulePhase) UnsetBillingCycleAnchor()`

UnsetBillingCycleAnchor ensures that no value is present for BillingCycleAnchor, not even an explicit nil
### GetDefaultPaymentMethod

`func (o *SchedulePhase) GetDefaultPaymentMethod() string`

GetDefaultPaymentMethod returns the DefaultPaymentMethod field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodOk

`func (o *SchedulePhase) GetDefaultPaymentMethodOk() (*string, bool)`

GetDefaultPaymentMethodOk returns a tuple with the DefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethod

`func (o *SchedulePhase) SetDefaultPaymentMethod(v string)`

SetDefaultPaymentMethod sets DefaultPaymentMethod field to given value.

### HasDefaultPaymentMethod

`func (o *SchedulePhase) HasDefaultPaymentMethod() bool`

HasDefaultPaymentMethod returns a boolean if a field has been set.

### SetDefaultPaymentMethodNil

`func (o *SchedulePhase) SetDefaultPaymentMethodNil(b bool)`

 SetDefaultPaymentMethodNil sets the value for DefaultPaymentMethod to be an explicit nil

### UnsetDefaultPaymentMethod
`func (o *SchedulePhase) UnsetDefaultPaymentMethod()`

UnsetDefaultPaymentMethod ensures that no value is present for DefaultPaymentMethod, not even an explicit nil
### GetCollectionMethod

`func (o *SchedulePhase) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *SchedulePhase) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *SchedulePhase) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.

### HasCollectionMethod

`func (o *SchedulePhase) HasCollectionMethod() bool`

HasCollectionMethod returns a boolean if a field has been set.

### GetCoupon

`func (o *SchedulePhase) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *SchedulePhase) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *SchedulePhase) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *SchedulePhase) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### SetCouponNil

`func (o *SchedulePhase) SetCouponNil(b bool)`

 SetCouponNil sets the value for Coupon to be an explicit nil

### UnsetCoupon
`func (o *SchedulePhase) UnsetCoupon()`

UnsetCoupon ensures that no value is present for Coupon, not even an explicit nil
### GetTrialEnd

`func (o *SchedulePhase) GetTrialEnd() string`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SchedulePhase) GetTrialEndOk() (*string, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SchedulePhase) SetTrialEnd(v string)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SchedulePhase) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *SchedulePhase) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *SchedulePhase) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetMetadata

`func (o *SchedulePhase) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchedulePhase) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchedulePhase) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchedulePhase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


