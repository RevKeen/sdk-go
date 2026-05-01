# EntitlementCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BenefitKey** | **string** |  | 
**HasAccess** | **bool** |  | 
**AccessLevel** | **string** |  | 
**Status** | **NullableString** |  | 
**Benefit** | [**Benefit**](Benefit.md) |  | 
**Reason** | **NullableString** |  | 

## Methods

### NewEntitlementCheck

`func NewEntitlementCheck(benefitKey string, hasAccess bool, accessLevel string, status NullableString, benefit Benefit, reason NullableString, ) *EntitlementCheck`

NewEntitlementCheck instantiates a new EntitlementCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementCheckWithDefaults

`func NewEntitlementCheckWithDefaults() *EntitlementCheck`

NewEntitlementCheckWithDefaults instantiates a new EntitlementCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenefitKey

`func (o *EntitlementCheck) GetBenefitKey() string`

GetBenefitKey returns the BenefitKey field if non-nil, zero value otherwise.

### GetBenefitKeyOk

`func (o *EntitlementCheck) GetBenefitKeyOk() (*string, bool)`

GetBenefitKeyOk returns a tuple with the BenefitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitKey

`func (o *EntitlementCheck) SetBenefitKey(v string)`

SetBenefitKey sets BenefitKey field to given value.


### GetHasAccess

`func (o *EntitlementCheck) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *EntitlementCheck) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *EntitlementCheck) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.


### GetAccessLevel

`func (o *EntitlementCheck) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *EntitlementCheck) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *EntitlementCheck) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetStatus

`func (o *EntitlementCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitlementCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitlementCheck) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *EntitlementCheck) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EntitlementCheck) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetBenefit

`func (o *EntitlementCheck) GetBenefit() Benefit`

GetBenefit returns the Benefit field if non-nil, zero value otherwise.

### GetBenefitOk

`func (o *EntitlementCheck) GetBenefitOk() (*Benefit, bool)`

GetBenefitOk returns a tuple with the Benefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefit

`func (o *EntitlementCheck) SetBenefit(v Benefit)`

SetBenefit sets Benefit field to given value.


### GetReason

`func (o *EntitlementCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EntitlementCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EntitlementCheck) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *EntitlementCheck) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *EntitlementCheck) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


