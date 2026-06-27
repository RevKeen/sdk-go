# UpdateDiscountInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DiscountValue** | Pointer to **float32** |  | [optional] 
**AppliesTo** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**ProductIds** | Pointer to **[]string** |  | [optional] 
**MaxRedemptions** | Pointer to **NullableInt32** |  | [optional] 
**MaxRedemptionsPerUser** | Pointer to **int32** |  | [optional] 
**ValidFrom** | Pointer to **NullableString** |  | [optional] 
**ValidUntil** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**RecurringType** | Pointer to **NullableString** |  | [optional] 
**RecurringCycles** | Pointer to **NullableInt32** |  | [optional] 
**FirstTimeCustomer** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateDiscountInput

`func NewUpdateDiscountInput() *UpdateDiscountInput`

NewUpdateDiscountInput instantiates a new UpdateDiscountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDiscountInputWithDefaults

`func NewUpdateDiscountInputWithDefaults() *UpdateDiscountInput`

NewUpdateDiscountInputWithDefaults instantiates a new UpdateDiscountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDiscountInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDiscountInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDiscountInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDiscountInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDiscountInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDiscountInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDiscountInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDiscountInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateDiscountInput) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateDiscountInput) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiscountValue

`func (o *UpdateDiscountInput) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *UpdateDiscountInput) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *UpdateDiscountInput) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.

### HasDiscountValue

`func (o *UpdateDiscountInput) HasDiscountValue() bool`

HasDiscountValue returns a boolean if a field has been set.

### GetAppliesTo

`func (o *UpdateDiscountInput) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *UpdateDiscountInput) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *UpdateDiscountInput) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *UpdateDiscountInput) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### GetScope

`func (o *UpdateDiscountInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateDiscountInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateDiscountInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateDiscountInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetProductIds

`func (o *UpdateDiscountInput) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *UpdateDiscountInput) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *UpdateDiscountInput) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *UpdateDiscountInput) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### SetProductIdsNil

`func (o *UpdateDiscountInput) SetProductIdsNil(b bool)`

 SetProductIdsNil sets the value for ProductIds to be an explicit nil

### UnsetProductIds
`func (o *UpdateDiscountInput) UnsetProductIds()`

UnsetProductIds ensures that no value is present for ProductIds, not even an explicit nil
### GetMaxRedemptions

`func (o *UpdateDiscountInput) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *UpdateDiscountInput) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *UpdateDiscountInput) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *UpdateDiscountInput) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### SetMaxRedemptionsNil

`func (o *UpdateDiscountInput) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *UpdateDiscountInput) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetMaxRedemptionsPerUser

`func (o *UpdateDiscountInput) GetMaxRedemptionsPerUser() int32`

GetMaxRedemptionsPerUser returns the MaxRedemptionsPerUser field if non-nil, zero value otherwise.

### GetMaxRedemptionsPerUserOk

`func (o *UpdateDiscountInput) GetMaxRedemptionsPerUserOk() (*int32, bool)`

GetMaxRedemptionsPerUserOk returns a tuple with the MaxRedemptionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionsPerUser

`func (o *UpdateDiscountInput) SetMaxRedemptionsPerUser(v int32)`

SetMaxRedemptionsPerUser sets MaxRedemptionsPerUser field to given value.

### HasMaxRedemptionsPerUser

`func (o *UpdateDiscountInput) HasMaxRedemptionsPerUser() bool`

HasMaxRedemptionsPerUser returns a boolean if a field has been set.

### GetValidFrom

`func (o *UpdateDiscountInput) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *UpdateDiscountInput) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *UpdateDiscountInput) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *UpdateDiscountInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFromNil

`func (o *UpdateDiscountInput) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *UpdateDiscountInput) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetValidUntil

`func (o *UpdateDiscountInput) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *UpdateDiscountInput) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *UpdateDiscountInput) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *UpdateDiscountInput) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### SetValidUntilNil

`func (o *UpdateDiscountInput) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *UpdateDiscountInput) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetIsActive

`func (o *UpdateDiscountInput) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateDiscountInput) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateDiscountInput) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateDiscountInput) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsArchived

`func (o *UpdateDiscountInput) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UpdateDiscountInput) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UpdateDiscountInput) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *UpdateDiscountInput) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetRecurringType

`func (o *UpdateDiscountInput) GetRecurringType() string`

GetRecurringType returns the RecurringType field if non-nil, zero value otherwise.

### GetRecurringTypeOk

`func (o *UpdateDiscountInput) GetRecurringTypeOk() (*string, bool)`

GetRecurringTypeOk returns a tuple with the RecurringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringType

`func (o *UpdateDiscountInput) SetRecurringType(v string)`

SetRecurringType sets RecurringType field to given value.

### HasRecurringType

`func (o *UpdateDiscountInput) HasRecurringType() bool`

HasRecurringType returns a boolean if a field has been set.

### SetRecurringTypeNil

`func (o *UpdateDiscountInput) SetRecurringTypeNil(b bool)`

 SetRecurringTypeNil sets the value for RecurringType to be an explicit nil

### UnsetRecurringType
`func (o *UpdateDiscountInput) UnsetRecurringType()`

UnsetRecurringType ensures that no value is present for RecurringType, not even an explicit nil
### GetRecurringCycles

`func (o *UpdateDiscountInput) GetRecurringCycles() int32`

GetRecurringCycles returns the RecurringCycles field if non-nil, zero value otherwise.

### GetRecurringCyclesOk

`func (o *UpdateDiscountInput) GetRecurringCyclesOk() (*int32, bool)`

GetRecurringCyclesOk returns a tuple with the RecurringCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCycles

`func (o *UpdateDiscountInput) SetRecurringCycles(v int32)`

SetRecurringCycles sets RecurringCycles field to given value.

### HasRecurringCycles

`func (o *UpdateDiscountInput) HasRecurringCycles() bool`

HasRecurringCycles returns a boolean if a field has been set.

### SetRecurringCyclesNil

`func (o *UpdateDiscountInput) SetRecurringCyclesNil(b bool)`

 SetRecurringCyclesNil sets the value for RecurringCycles to be an explicit nil

### UnsetRecurringCycles
`func (o *UpdateDiscountInput) UnsetRecurringCycles()`

UnsetRecurringCycles ensures that no value is present for RecurringCycles, not even an explicit nil
### GetFirstTimeCustomer

`func (o *UpdateDiscountInput) GetFirstTimeCustomer() bool`

GetFirstTimeCustomer returns the FirstTimeCustomer field if non-nil, zero value otherwise.

### GetFirstTimeCustomerOk

`func (o *UpdateDiscountInput) GetFirstTimeCustomerOk() (*bool, bool)`

GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeCustomer

`func (o *UpdateDiscountInput) SetFirstTimeCustomer(v bool)`

SetFirstTimeCustomer sets FirstTimeCustomer field to given value.

### HasFirstTimeCustomer

`func (o *UpdateDiscountInput) HasFirstTimeCustomer() bool`

HasFirstTimeCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


