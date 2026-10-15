# SubscriptionDeleteConflictErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** |  |
**UnbilledUsageRecordCount** | Pointer to **int32** |  | [optional]
**UsageMeterDecisionCount** | Pointer to **int32** |  | [optional]

## Methods

### NewSubscriptionDeleteConflictErrorDetails

`func NewSubscriptionDeleteConflictErrorDetails(reason string, ) *SubscriptionDeleteConflictErrorDetails`

NewSubscriptionDeleteConflictErrorDetails instantiates a new SubscriptionDeleteConflictErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDeleteConflictErrorDetailsWithDefaults

`func NewSubscriptionDeleteConflictErrorDetailsWithDefaults() *SubscriptionDeleteConflictErrorDetails`

NewSubscriptionDeleteConflictErrorDetailsWithDefaults instantiates a new SubscriptionDeleteConflictErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *SubscriptionDeleteConflictErrorDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubscriptionDeleteConflictErrorDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubscriptionDeleteConflictErrorDetails) SetReason(v string)`

SetReason sets Reason field to given value.


### GetUnbilledUsageRecordCount

`func (o *SubscriptionDeleteConflictErrorDetails) GetUnbilledUsageRecordCount() int32`

GetUnbilledUsageRecordCount returns the UnbilledUsageRecordCount field if non-nil, zero value otherwise.

### GetUnbilledUsageRecordCountOk

`func (o *SubscriptionDeleteConflictErrorDetails) GetUnbilledUsageRecordCountOk() (*int32, bool)`

GetUnbilledUsageRecordCountOk returns a tuple with the UnbilledUsageRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledUsageRecordCount

`func (o *SubscriptionDeleteConflictErrorDetails) SetUnbilledUsageRecordCount(v int32)`

SetUnbilledUsageRecordCount sets UnbilledUsageRecordCount field to given value.

### HasUnbilledUsageRecordCount

`func (o *SubscriptionDeleteConflictErrorDetails) HasUnbilledUsageRecordCount() bool`

HasUnbilledUsageRecordCount returns a boolean if a field has been set.

### GetUsageMeterDecisionCount

`func (o *SubscriptionDeleteConflictErrorDetails) GetUsageMeterDecisionCount() int32`

GetUsageMeterDecisionCount returns the UsageMeterDecisionCount field if non-nil, zero value otherwise.

### GetUsageMeterDecisionCountOk

`func (o *SubscriptionDeleteConflictErrorDetails) GetUsageMeterDecisionCountOk() (*int32, bool)`

GetUsageMeterDecisionCountOk returns a tuple with the UsageMeterDecisionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterDecisionCount

`func (o *SubscriptionDeleteConflictErrorDetails) SetUsageMeterDecisionCount(v int32)`

SetUsageMeterDecisionCount sets UsageMeterDecisionCount field to given value.

### HasUsageMeterDecisionCount

`func (o *SubscriptionDeleteConflictErrorDetails) HasUsageMeterDecisionCount() bool`

HasUsageMeterDecisionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
