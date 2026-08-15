# Entitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CustomerId** | **string** |  | 
**BenefitId** | **string** |  | 
**Benefit** | [**Benefit**](Benefit.md) |  | 
**GrantedAt** | **time.Time** |  | 
**ExpiresAt** | **NullableTime** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Status** | **string** |  | 
**HasAccess** | **bool** |  | 
**AccessLevel** | **string** |  | 
**SubscriptionId** | **NullableString** |  | 
**SubscriptionStatus** | **NullableString** |  | 

## Methods

### NewEntitlement

`func NewEntitlement(id string, customerId string, benefitId string, benefit Benefit, grantedAt time.Time, expiresAt NullableTime, metadata map[string]interface{}, status string, hasAccess bool, accessLevel string, subscriptionId NullableString, subscriptionStatus NullableString, ) *Entitlement`

NewEntitlement instantiates a new Entitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithDefaults

`func NewEntitlementWithDefaults() *Entitlement`

NewEntitlementWithDefaults instantiates a new Entitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlement) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *Entitlement) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Entitlement) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Entitlement) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBenefitId

`func (o *Entitlement) GetBenefitId() string`

GetBenefitId returns the BenefitId field if non-nil, zero value otherwise.

### GetBenefitIdOk

`func (o *Entitlement) GetBenefitIdOk() (*string, bool)`

GetBenefitIdOk returns a tuple with the BenefitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitId

`func (o *Entitlement) SetBenefitId(v string)`

SetBenefitId sets BenefitId field to given value.


### GetBenefit

`func (o *Entitlement) GetBenefit() Benefit`

GetBenefit returns the Benefit field if non-nil, zero value otherwise.

### GetBenefitOk

`func (o *Entitlement) GetBenefitOk() (*Benefit, bool)`

GetBenefitOk returns a tuple with the Benefit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefit

`func (o *Entitlement) SetBenefit(v Benefit)`

SetBenefit sets Benefit field to given value.


### GetGrantedAt

`func (o *Entitlement) GetGrantedAt() time.Time`

GetGrantedAt returns the GrantedAt field if non-nil, zero value otherwise.

### GetGrantedAtOk

`func (o *Entitlement) GetGrantedAtOk() (*time.Time, bool)`

GetGrantedAtOk returns a tuple with the GrantedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedAt

`func (o *Entitlement) SetGrantedAt(v time.Time)`

SetGrantedAt sets GrantedAt field to given value.


### GetExpiresAt

`func (o *Entitlement) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Entitlement) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Entitlement) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *Entitlement) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Entitlement) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetMetadata

`func (o *Entitlement) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Entitlement) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Entitlement) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *Entitlement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Entitlement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Entitlement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHasAccess

`func (o *Entitlement) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *Entitlement) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *Entitlement) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.


### GetAccessLevel

`func (o *Entitlement) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *Entitlement) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *Entitlement) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.


### GetSubscriptionId

`func (o *Entitlement) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Entitlement) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Entitlement) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *Entitlement) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *Entitlement) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetSubscriptionStatus

`func (o *Entitlement) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *Entitlement) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *Entitlement) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.


### SetSubscriptionStatusNil

`func (o *Entitlement) SetSubscriptionStatusNil(b bool)`

 SetSubscriptionStatusNil sets the value for SubscriptionStatus to be an explicit nil

### UnsetSubscriptionStatus
`func (o *Entitlement) UnsetSubscriptionStatus()`

UnsetSubscriptionStatus ensures that no value is present for SubscriptionStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


