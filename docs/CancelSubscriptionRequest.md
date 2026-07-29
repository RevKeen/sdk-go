# CancelSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAtPeriodEnd** | Pointer to **bool** | If &#x60;true&#x60; (default), cancel at the end of the current period. If &#x60;false&#x60;, cancel immediately. | [optional] [default to true]
**Reason** | Pointer to **string** | Merchant-visible note recording why the customer canceled. | [optional] 

## Methods

### NewCancelSubscriptionRequest

`func NewCancelSubscriptionRequest() *CancelSubscriptionRequest`

NewCancelSubscriptionRequest instantiates a new CancelSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionRequestWithDefaults

`func NewCancelSubscriptionRequestWithDefaults() *CancelSubscriptionRequest`

NewCancelSubscriptionRequestWithDefaults instantiates a new CancelSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAtPeriodEnd

`func (o *CancelSubscriptionRequest) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *CancelSubscriptionRequest) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *CancelSubscriptionRequest) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *CancelSubscriptionRequest) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetReason

`func (o *CancelSubscriptionRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelSubscriptionRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelSubscriptionRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelSubscriptionRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


