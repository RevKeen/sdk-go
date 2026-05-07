# SubscriptionCreateResponseDataDunning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsInDunning** | **bool** |  | 
**Phase** | **NullableFloat32** |  | 
**PhaseLabel** | **NullableString** |  | 
**PhaseSeverity** | **NullableString** |  | 
**RetryCount** | **float32** |  | 
**TotalPossibleRetries** | **float32** |  | 
**NextRetryAt** | **NullableTime** |  | 
**DaysInDunning** | **float32** |  | 
**AccessRestricted** | **bool** |  | 

## Methods

### NewSubscriptionCreateResponseDataDunning

`func NewSubscriptionCreateResponseDataDunning(isInDunning bool, phase NullableFloat32, phaseLabel NullableString, phaseSeverity NullableString, retryCount float32, totalPossibleRetries float32, nextRetryAt NullableTime, daysInDunning float32, accessRestricted bool, ) *SubscriptionCreateResponseDataDunning`

NewSubscriptionCreateResponseDataDunning instantiates a new SubscriptionCreateResponseDataDunning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateResponseDataDunningWithDefaults

`func NewSubscriptionCreateResponseDataDunningWithDefaults() *SubscriptionCreateResponseDataDunning`

NewSubscriptionCreateResponseDataDunningWithDefaults instantiates a new SubscriptionCreateResponseDataDunning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsInDunning

`func (o *SubscriptionCreateResponseDataDunning) GetIsInDunning() bool`

GetIsInDunning returns the IsInDunning field if non-nil, zero value otherwise.

### GetIsInDunningOk

`func (o *SubscriptionCreateResponseDataDunning) GetIsInDunningOk() (*bool, bool)`

GetIsInDunningOk returns a tuple with the IsInDunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInDunning

`func (o *SubscriptionCreateResponseDataDunning) SetIsInDunning(v bool)`

SetIsInDunning sets IsInDunning field to given value.


### GetPhase

`func (o *SubscriptionCreateResponseDataDunning) GetPhase() float32`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SubscriptionCreateResponseDataDunning) GetPhaseOk() (*float32, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SubscriptionCreateResponseDataDunning) SetPhase(v float32)`

SetPhase sets Phase field to given value.


### SetPhaseNil

`func (o *SubscriptionCreateResponseDataDunning) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *SubscriptionCreateResponseDataDunning) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetPhaseLabel

`func (o *SubscriptionCreateResponseDataDunning) GetPhaseLabel() string`

GetPhaseLabel returns the PhaseLabel field if non-nil, zero value otherwise.

### GetPhaseLabelOk

`func (o *SubscriptionCreateResponseDataDunning) GetPhaseLabelOk() (*string, bool)`

GetPhaseLabelOk returns a tuple with the PhaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseLabel

`func (o *SubscriptionCreateResponseDataDunning) SetPhaseLabel(v string)`

SetPhaseLabel sets PhaseLabel field to given value.


### SetPhaseLabelNil

`func (o *SubscriptionCreateResponseDataDunning) SetPhaseLabelNil(b bool)`

 SetPhaseLabelNil sets the value for PhaseLabel to be an explicit nil

### UnsetPhaseLabel
`func (o *SubscriptionCreateResponseDataDunning) UnsetPhaseLabel()`

UnsetPhaseLabel ensures that no value is present for PhaseLabel, not even an explicit nil
### GetPhaseSeverity

`func (o *SubscriptionCreateResponseDataDunning) GetPhaseSeverity() string`

GetPhaseSeverity returns the PhaseSeverity field if non-nil, zero value otherwise.

### GetPhaseSeverityOk

`func (o *SubscriptionCreateResponseDataDunning) GetPhaseSeverityOk() (*string, bool)`

GetPhaseSeverityOk returns a tuple with the PhaseSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseSeverity

`func (o *SubscriptionCreateResponseDataDunning) SetPhaseSeverity(v string)`

SetPhaseSeverity sets PhaseSeverity field to given value.


### SetPhaseSeverityNil

`func (o *SubscriptionCreateResponseDataDunning) SetPhaseSeverityNil(b bool)`

 SetPhaseSeverityNil sets the value for PhaseSeverity to be an explicit nil

### UnsetPhaseSeverity
`func (o *SubscriptionCreateResponseDataDunning) UnsetPhaseSeverity()`

UnsetPhaseSeverity ensures that no value is present for PhaseSeverity, not even an explicit nil
### GetRetryCount

`func (o *SubscriptionCreateResponseDataDunning) GetRetryCount() float32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *SubscriptionCreateResponseDataDunning) GetRetryCountOk() (*float32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *SubscriptionCreateResponseDataDunning) SetRetryCount(v float32)`

SetRetryCount sets RetryCount field to given value.


### GetTotalPossibleRetries

`func (o *SubscriptionCreateResponseDataDunning) GetTotalPossibleRetries() float32`

GetTotalPossibleRetries returns the TotalPossibleRetries field if non-nil, zero value otherwise.

### GetTotalPossibleRetriesOk

`func (o *SubscriptionCreateResponseDataDunning) GetTotalPossibleRetriesOk() (*float32, bool)`

GetTotalPossibleRetriesOk returns a tuple with the TotalPossibleRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPossibleRetries

`func (o *SubscriptionCreateResponseDataDunning) SetTotalPossibleRetries(v float32)`

SetTotalPossibleRetries sets TotalPossibleRetries field to given value.


### GetNextRetryAt

`func (o *SubscriptionCreateResponseDataDunning) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *SubscriptionCreateResponseDataDunning) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *SubscriptionCreateResponseDataDunning) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### SetNextRetryAtNil

`func (o *SubscriptionCreateResponseDataDunning) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *SubscriptionCreateResponseDataDunning) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil
### GetDaysInDunning

`func (o *SubscriptionCreateResponseDataDunning) GetDaysInDunning() float32`

GetDaysInDunning returns the DaysInDunning field if non-nil, zero value otherwise.

### GetDaysInDunningOk

`func (o *SubscriptionCreateResponseDataDunning) GetDaysInDunningOk() (*float32, bool)`

GetDaysInDunningOk returns a tuple with the DaysInDunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInDunning

`func (o *SubscriptionCreateResponseDataDunning) SetDaysInDunning(v float32)`

SetDaysInDunning sets DaysInDunning field to given value.


### GetAccessRestricted

`func (o *SubscriptionCreateResponseDataDunning) GetAccessRestricted() bool`

GetAccessRestricted returns the AccessRestricted field if non-nil, zero value otherwise.

### GetAccessRestrictedOk

`func (o *SubscriptionCreateResponseDataDunning) GetAccessRestrictedOk() (*bool, bool)`

GetAccessRestrictedOk returns a tuple with the AccessRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRestricted

`func (o *SubscriptionCreateResponseDataDunning) SetAccessRestricted(v bool)`

SetAccessRestricted sets AccessRestricted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


