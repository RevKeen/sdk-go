# Discount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Code** | **string** |  | 
**Name** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**DiscountType** | **string** |  | 
**DiscountValue** | **float32** | Amount off (cents) or percentage | 
**AppliesTo** | **NullableString** |  | 
**ProductIds** | **[]string** |  | 
**MaxRedemptions** | **NullableFloat32** |  | 
**CurrentRedemptions** | **NullableFloat32** |  | 
**ValidFrom** | **NullableString** |  | 
**ValidUntil** | **NullableString** |  | 
**RecurringType** | **NullableString** |  | 
**RecurringCycles** | **NullableFloat32** |  | 
**FirstTimeCustomer** | **NullableBool** |  | 
**IsActive** | **NullableBool** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewDiscount

`func NewDiscount(id string, code string, name NullableString, description NullableString, discountType string, discountValue float32, appliesTo NullableString, productIds []string, maxRedemptions NullableFloat32, currentRedemptions NullableFloat32, validFrom NullableString, validUntil NullableString, recurringType NullableString, recurringCycles NullableFloat32, firstTimeCustomer NullableBool, isActive NullableBool, createdAt NullableString, updatedAt NullableString, ) *Discount`

NewDiscount instantiates a new Discount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountWithDefaults

`func NewDiscountWithDefaults() *Discount`

NewDiscountWithDefaults instantiates a new Discount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Discount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Discount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Discount) SetId(v string)`

SetId sets Id field to given value.


### GetCode

`func (o *Discount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Discount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Discount) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *Discount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Discount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Discount) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Discount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Discount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Discount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Discount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Discount) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Discount) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Discount) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDiscountType

`func (o *Discount) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *Discount) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *Discount) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetDiscountValue

`func (o *Discount) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *Discount) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *Discount) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.


### GetAppliesTo

`func (o *Discount) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *Discount) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *Discount) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.


### SetAppliesToNil

`func (o *Discount) SetAppliesToNil(b bool)`

 SetAppliesToNil sets the value for AppliesTo to be an explicit nil

### UnsetAppliesTo
`func (o *Discount) UnsetAppliesTo()`

UnsetAppliesTo ensures that no value is present for AppliesTo, not even an explicit nil
### GetProductIds

`func (o *Discount) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *Discount) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *Discount) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.


### SetProductIdsNil

`func (o *Discount) SetProductIdsNil(b bool)`

 SetProductIdsNil sets the value for ProductIds to be an explicit nil

### UnsetProductIds
`func (o *Discount) UnsetProductIds()`

UnsetProductIds ensures that no value is present for ProductIds, not even an explicit nil
### GetMaxRedemptions

`func (o *Discount) GetMaxRedemptions() float32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *Discount) GetMaxRedemptionsOk() (*float32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *Discount) SetMaxRedemptions(v float32)`

SetMaxRedemptions sets MaxRedemptions field to given value.


### SetMaxRedemptionsNil

`func (o *Discount) SetMaxRedemptionsNil(b bool)`

 SetMaxRedemptionsNil sets the value for MaxRedemptions to be an explicit nil

### UnsetMaxRedemptions
`func (o *Discount) UnsetMaxRedemptions()`

UnsetMaxRedemptions ensures that no value is present for MaxRedemptions, not even an explicit nil
### GetCurrentRedemptions

`func (o *Discount) GetCurrentRedemptions() float32`

GetCurrentRedemptions returns the CurrentRedemptions field if non-nil, zero value otherwise.

### GetCurrentRedemptionsOk

`func (o *Discount) GetCurrentRedemptionsOk() (*float32, bool)`

GetCurrentRedemptionsOk returns a tuple with the CurrentRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRedemptions

`func (o *Discount) SetCurrentRedemptions(v float32)`

SetCurrentRedemptions sets CurrentRedemptions field to given value.


### SetCurrentRedemptionsNil

`func (o *Discount) SetCurrentRedemptionsNil(b bool)`

 SetCurrentRedemptionsNil sets the value for CurrentRedemptions to be an explicit nil

### UnsetCurrentRedemptions
`func (o *Discount) UnsetCurrentRedemptions()`

UnsetCurrentRedemptions ensures that no value is present for CurrentRedemptions, not even an explicit nil
### GetValidFrom

`func (o *Discount) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *Discount) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *Discount) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### SetValidFromNil

`func (o *Discount) SetValidFromNil(b bool)`

 SetValidFromNil sets the value for ValidFrom to be an explicit nil

### UnsetValidFrom
`func (o *Discount) UnsetValidFrom()`

UnsetValidFrom ensures that no value is present for ValidFrom, not even an explicit nil
### GetValidUntil

`func (o *Discount) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *Discount) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *Discount) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.


### SetValidUntilNil

`func (o *Discount) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *Discount) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil
### GetRecurringType

`func (o *Discount) GetRecurringType() string`

GetRecurringType returns the RecurringType field if non-nil, zero value otherwise.

### GetRecurringTypeOk

`func (o *Discount) GetRecurringTypeOk() (*string, bool)`

GetRecurringTypeOk returns a tuple with the RecurringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringType

`func (o *Discount) SetRecurringType(v string)`

SetRecurringType sets RecurringType field to given value.


### SetRecurringTypeNil

`func (o *Discount) SetRecurringTypeNil(b bool)`

 SetRecurringTypeNil sets the value for RecurringType to be an explicit nil

### UnsetRecurringType
`func (o *Discount) UnsetRecurringType()`

UnsetRecurringType ensures that no value is present for RecurringType, not even an explicit nil
### GetRecurringCycles

`func (o *Discount) GetRecurringCycles() float32`

GetRecurringCycles returns the RecurringCycles field if non-nil, zero value otherwise.

### GetRecurringCyclesOk

`func (o *Discount) GetRecurringCyclesOk() (*float32, bool)`

GetRecurringCyclesOk returns a tuple with the RecurringCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCycles

`func (o *Discount) SetRecurringCycles(v float32)`

SetRecurringCycles sets RecurringCycles field to given value.


### SetRecurringCyclesNil

`func (o *Discount) SetRecurringCyclesNil(b bool)`

 SetRecurringCyclesNil sets the value for RecurringCycles to be an explicit nil

### UnsetRecurringCycles
`func (o *Discount) UnsetRecurringCycles()`

UnsetRecurringCycles ensures that no value is present for RecurringCycles, not even an explicit nil
### GetFirstTimeCustomer

`func (o *Discount) GetFirstTimeCustomer() bool`

GetFirstTimeCustomer returns the FirstTimeCustomer field if non-nil, zero value otherwise.

### GetFirstTimeCustomerOk

`func (o *Discount) GetFirstTimeCustomerOk() (*bool, bool)`

GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeCustomer

`func (o *Discount) SetFirstTimeCustomer(v bool)`

SetFirstTimeCustomer sets FirstTimeCustomer field to given value.


### SetFirstTimeCustomerNil

`func (o *Discount) SetFirstTimeCustomerNil(b bool)`

 SetFirstTimeCustomerNil sets the value for FirstTimeCustomer to be an explicit nil

### UnsetFirstTimeCustomer
`func (o *Discount) UnsetFirstTimeCustomer()`

UnsetFirstTimeCustomer ensures that no value is present for FirstTimeCustomer, not even an explicit nil
### GetIsActive

`func (o *Discount) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Discount) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Discount) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### SetIsActiveNil

`func (o *Discount) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *Discount) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetCreatedAt

`func (o *Discount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Discount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Discount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Discount) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Discount) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *Discount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Discount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Discount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Discount) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Discount) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


