# SubscriptionPauseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Subscription**](Subscription.md) |  |
**DdCollectionStopped** | Pointer to [**DdCollectionStopped**](DdCollectionStopped.md) |  | [optional]

## Methods

### NewSubscriptionPauseResponse

`func NewSubscriptionPauseResponse(data Subscription, ) *SubscriptionPauseResponse`

NewSubscriptionPauseResponse instantiates a new SubscriptionPauseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPauseResponseWithDefaults

`func NewSubscriptionPauseResponseWithDefaults() *SubscriptionPauseResponse`

NewSubscriptionPauseResponseWithDefaults instantiates a new SubscriptionPauseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionPauseResponse) GetData() Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionPauseResponse) GetDataOk() (*Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionPauseResponse) SetData(v Subscription)`

SetData sets Data field to given value.


### GetDdCollectionStopped

`func (o *SubscriptionPauseResponse) GetDdCollectionStopped() DdCollectionStopped`

GetDdCollectionStopped returns the DdCollectionStopped field if non-nil, zero value otherwise.

### GetDdCollectionStoppedOk

`func (o *SubscriptionPauseResponse) GetDdCollectionStoppedOk() (*DdCollectionStopped, bool)`

GetDdCollectionStoppedOk returns a tuple with the DdCollectionStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdCollectionStopped

`func (o *SubscriptionPauseResponse) SetDdCollectionStopped(v DdCollectionStopped)`

SetDdCollectionStopped sets DdCollectionStopped field to given value.

### HasDdCollectionStopped

`func (o *SubscriptionPauseResponse) HasDdCollectionStopped() bool`

HasDdCollectionStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
