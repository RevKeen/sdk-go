# SubscriptionCancelSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Subscription**](Subscription.md) |  |
**DdCollectionStopped** | Pointer to [**DdCollectionStopped**](DdCollectionStopped.md) |  | [optional]

## Methods

### NewSubscriptionCancelSubscriptionResponse

`func NewSubscriptionCancelSubscriptionResponse(data Subscription, ) *SubscriptionCancelSubscriptionResponse`

NewSubscriptionCancelSubscriptionResponse instantiates a new SubscriptionCancelSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCancelSubscriptionResponseWithDefaults

`func NewSubscriptionCancelSubscriptionResponseWithDefaults() *SubscriptionCancelSubscriptionResponse`

NewSubscriptionCancelSubscriptionResponseWithDefaults instantiates a new SubscriptionCancelSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionCancelSubscriptionResponse) GetData() Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionCancelSubscriptionResponse) GetDataOk() (*Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionCancelSubscriptionResponse) SetData(v Subscription)`

SetData sets Data field to given value.


### GetDdCollectionStopped

`func (o *SubscriptionCancelSubscriptionResponse) GetDdCollectionStopped() DdCollectionStopped`

GetDdCollectionStopped returns the DdCollectionStopped field if non-nil, zero value otherwise.

### GetDdCollectionStoppedOk

`func (o *SubscriptionCancelSubscriptionResponse) GetDdCollectionStoppedOk() (*DdCollectionStopped, bool)`

GetDdCollectionStoppedOk returns a tuple with the DdCollectionStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdCollectionStopped

`func (o *SubscriptionCancelSubscriptionResponse) SetDdCollectionStopped(v DdCollectionStopped)`

SetDdCollectionStopped sets DdCollectionStopped field to given value.

### HasDdCollectionStopped

`func (o *SubscriptionCancelSubscriptionResponse) HasDdCollectionStopped() bool`

HasDdCollectionStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
