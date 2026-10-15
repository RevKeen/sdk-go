# SubscriptionsChangePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPlanId** | **string** |  |
**ProrationBehavior** | Pointer to **string** | How the price difference is billed. &#x60;next_invoice&#x60; (and &#x60;invoice_immediately&#x60;, accepted for compatibility and applied the same way) adds it to the next renewal invoice as a line; a downgrade credit is carried there, never refunded, and a credit larger than the renewal carries to the following one. &#x60;none&#x60; changes without proration. | [optional]
**EffectiveDate** | Pointer to **time.Time** |  | [optional]
**Reason** | Pointer to **string** |  | [optional]

## Methods

### NewSubscriptionsChangePlanRequest

`func NewSubscriptionsChangePlanRequest(newPlanId string, ) *SubscriptionsChangePlanRequest`

NewSubscriptionsChangePlanRequest instantiates a new SubscriptionsChangePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsChangePlanRequestWithDefaults

`func NewSubscriptionsChangePlanRequestWithDefaults() *SubscriptionsChangePlanRequest`

NewSubscriptionsChangePlanRequestWithDefaults instantiates a new SubscriptionsChangePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPlanId

`func (o *SubscriptionsChangePlanRequest) GetNewPlanId() string`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *SubscriptionsChangePlanRequest) GetNewPlanIdOk() (*string, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *SubscriptionsChangePlanRequest) SetNewPlanId(v string)`

SetNewPlanId sets NewPlanId field to given value.


### GetProrationBehavior

`func (o *SubscriptionsChangePlanRequest) GetProrationBehavior() string`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *SubscriptionsChangePlanRequest) GetProrationBehaviorOk() (*string, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *SubscriptionsChangePlanRequest) SetProrationBehavior(v string)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *SubscriptionsChangePlanRequest) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *SubscriptionsChangePlanRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *SubscriptionsChangePlanRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *SubscriptionsChangePlanRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *SubscriptionsChangePlanRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetReason

`func (o *SubscriptionsChangePlanRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubscriptionsChangePlanRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubscriptionsChangePlanRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubscriptionsChangePlanRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
