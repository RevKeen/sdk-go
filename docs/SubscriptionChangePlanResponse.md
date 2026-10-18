# SubscriptionChangePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Subscription**](Subscription.md) |  |
**PlanChange** | [**SubscriptionPlanChange**](SubscriptionPlanChange.md) |  |

## Methods

### NewSubscriptionChangePlanResponse

`func NewSubscriptionChangePlanResponse(data Subscription, planChange SubscriptionPlanChange, ) *SubscriptionChangePlanResponse`

NewSubscriptionChangePlanResponse instantiates a new SubscriptionChangePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionChangePlanResponseWithDefaults

`func NewSubscriptionChangePlanResponseWithDefaults() *SubscriptionChangePlanResponse`

NewSubscriptionChangePlanResponseWithDefaults instantiates a new SubscriptionChangePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionChangePlanResponse) GetData() Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionChangePlanResponse) GetDataOk() (*Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionChangePlanResponse) SetData(v Subscription)`

SetData sets Data field to given value.


### GetPlanChange

`func (o *SubscriptionChangePlanResponse) GetPlanChange() SubscriptionPlanChange`

GetPlanChange returns the PlanChange field if non-nil, zero value otherwise.

### GetPlanChangeOk

`func (o *SubscriptionChangePlanResponse) GetPlanChangeOk() (*SubscriptionPlanChange, bool)`

GetPlanChangeOk returns a tuple with the PlanChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanChange

`func (o *SubscriptionChangePlanResponse) SetPlanChange(v SubscriptionPlanChange)`

SetPlanChange sets PlanChange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
