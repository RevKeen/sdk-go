# SubscriptionCancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  |
**DdCollectionStopped** | Pointer to [**DdCollectionStopped**](DdCollectionStopped.md) |  | [optional]

## Methods

### NewSubscriptionCancelResponse

`func NewSubscriptionCancelResponse(success bool, ) *SubscriptionCancelResponse`

NewSubscriptionCancelResponse instantiates a new SubscriptionCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCancelResponseWithDefaults

`func NewSubscriptionCancelResponseWithDefaults() *SubscriptionCancelResponse`

NewSubscriptionCancelResponseWithDefaults instantiates a new SubscriptionCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SubscriptionCancelResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SubscriptionCancelResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SubscriptionCancelResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetDdCollectionStopped

`func (o *SubscriptionCancelResponse) GetDdCollectionStopped() DdCollectionStopped`

GetDdCollectionStopped returns the DdCollectionStopped field if non-nil, zero value otherwise.

### GetDdCollectionStoppedOk

`func (o *SubscriptionCancelResponse) GetDdCollectionStoppedOk() (*DdCollectionStopped, bool)`

GetDdCollectionStoppedOk returns a tuple with the DdCollectionStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdCollectionStopped

`func (o *SubscriptionCancelResponse) SetDdCollectionStopped(v DdCollectionStopped)`

SetDdCollectionStopped sets DdCollectionStopped field to given value.

### HasDdCollectionStopped

`func (o *SubscriptionCancelResponse) HasDdCollectionStopped() bool`

HasDdCollectionStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
