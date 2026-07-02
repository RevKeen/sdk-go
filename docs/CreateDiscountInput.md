# CreateDiscountInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Unique discount code (will be uppercased) | 
**Name** | **string** | Human-readable name | 
**Description** | Pointer to **string** | Optional description | [optional] 
**DiscountType** | **string** | Type of discount | 
**DiscountValue** | **float32** | Discount value - percentage (1-100) or amount in cents | 
**AppliesTo** | Pointer to **string** | Which products the discount applies to | [optional] [default to "all"]
**Scope** | Pointer to **string** | Scope of where discount can be applied | [optional] 
**ProductIds** | Pointer to **[]string** | Product IDs if applies_to is &#39;specific_products&#39; | [optional] 
**MaxRedemptions** | Pointer to **int32** | Maximum total redemptions allowed | [optional] 
**MaxRedemptionsPerUser** | Pointer to **int32** | Maximum redemptions per customer (0 &#x3D; unlimited) | [optional] [default to 0]
**ValidFrom** | Pointer to **string** | Start date (ISO 8601) | [optional] 
**ValidUntil** | Pointer to **string** | End date (ISO 8601) | [optional] 
**RecurringType** | Pointer to **string** | How the discount applies to recurring payments | [optional] 
**RecurringCycles** | Pointer to **int32** | Number of billing cycles (only if recurring_type is &#39;repeating&#39;) | [optional] 
**FirstTimeCustomer** | Pointer to **bool** | Only available to first-time customers | [optional] [default to false]

## Methods

### NewCreateDiscountInput

`func NewCreateDiscountInput(code string, name string, discountType string, discountValue float32, ) *CreateDiscountInput`

NewCreateDiscountInput instantiates a new CreateDiscountInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscountInputWithDefaults

`func NewCreateDiscountInputWithDefaults() *CreateDiscountInput`

NewCreateDiscountInputWithDefaults instantiates a new CreateDiscountInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateDiscountInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateDiscountInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateDiscountInput) SetCode(v string)`

SetCode sets Code field to given value.


### GetName

`func (o *CreateDiscountInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDiscountInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDiscountInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateDiscountInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDiscountInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDiscountInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDiscountInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountType

`func (o *CreateDiscountInput) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *CreateDiscountInput) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *CreateDiscountInput) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetDiscountValue

`func (o *CreateDiscountInput) GetDiscountValue() float32`

GetDiscountValue returns the DiscountValue field if non-nil, zero value otherwise.

### GetDiscountValueOk

`func (o *CreateDiscountInput) GetDiscountValueOk() (*float32, bool)`

GetDiscountValueOk returns a tuple with the DiscountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountValue

`func (o *CreateDiscountInput) SetDiscountValue(v float32)`

SetDiscountValue sets DiscountValue field to given value.


### GetAppliesTo

`func (o *CreateDiscountInput) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *CreateDiscountInput) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *CreateDiscountInput) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *CreateDiscountInput) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### GetScope

`func (o *CreateDiscountInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateDiscountInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateDiscountInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateDiscountInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetProductIds

`func (o *CreateDiscountInput) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *CreateDiscountInput) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *CreateDiscountInput) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *CreateDiscountInput) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetMaxRedemptions

`func (o *CreateDiscountInput) GetMaxRedemptions() int32`

GetMaxRedemptions returns the MaxRedemptions field if non-nil, zero value otherwise.

### GetMaxRedemptionsOk

`func (o *CreateDiscountInput) GetMaxRedemptionsOk() (*int32, bool)`

GetMaxRedemptionsOk returns a tuple with the MaxRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptions

`func (o *CreateDiscountInput) SetMaxRedemptions(v int32)`

SetMaxRedemptions sets MaxRedemptions field to given value.

### HasMaxRedemptions

`func (o *CreateDiscountInput) HasMaxRedemptions() bool`

HasMaxRedemptions returns a boolean if a field has been set.

### GetMaxRedemptionsPerUser

`func (o *CreateDiscountInput) GetMaxRedemptionsPerUser() int32`

GetMaxRedemptionsPerUser returns the MaxRedemptionsPerUser field if non-nil, zero value otherwise.

### GetMaxRedemptionsPerUserOk

`func (o *CreateDiscountInput) GetMaxRedemptionsPerUserOk() (*int32, bool)`

GetMaxRedemptionsPerUserOk returns a tuple with the MaxRedemptionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionsPerUser

`func (o *CreateDiscountInput) SetMaxRedemptionsPerUser(v int32)`

SetMaxRedemptionsPerUser sets MaxRedemptionsPerUser field to given value.

### HasMaxRedemptionsPerUser

`func (o *CreateDiscountInput) HasMaxRedemptionsPerUser() bool`

HasMaxRedemptionsPerUser returns a boolean if a field has been set.

### GetValidFrom

`func (o *CreateDiscountInput) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CreateDiscountInput) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CreateDiscountInput) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *CreateDiscountInput) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidUntil

`func (o *CreateDiscountInput) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *CreateDiscountInput) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *CreateDiscountInput) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *CreateDiscountInput) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetRecurringType

`func (o *CreateDiscountInput) GetRecurringType() string`

GetRecurringType returns the RecurringType field if non-nil, zero value otherwise.

### GetRecurringTypeOk

`func (o *CreateDiscountInput) GetRecurringTypeOk() (*string, bool)`

GetRecurringTypeOk returns a tuple with the RecurringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringType

`func (o *CreateDiscountInput) SetRecurringType(v string)`

SetRecurringType sets RecurringType field to given value.

### HasRecurringType

`func (o *CreateDiscountInput) HasRecurringType() bool`

HasRecurringType returns a boolean if a field has been set.

### GetRecurringCycles

`func (o *CreateDiscountInput) GetRecurringCycles() int32`

GetRecurringCycles returns the RecurringCycles field if non-nil, zero value otherwise.

### GetRecurringCyclesOk

`func (o *CreateDiscountInput) GetRecurringCyclesOk() (*int32, bool)`

GetRecurringCyclesOk returns a tuple with the RecurringCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCycles

`func (o *CreateDiscountInput) SetRecurringCycles(v int32)`

SetRecurringCycles sets RecurringCycles field to given value.

### HasRecurringCycles

`func (o *CreateDiscountInput) HasRecurringCycles() bool`

HasRecurringCycles returns a boolean if a field has been set.

### GetFirstTimeCustomer

`func (o *CreateDiscountInput) GetFirstTimeCustomer() bool`

GetFirstTimeCustomer returns the FirstTimeCustomer field if non-nil, zero value otherwise.

### GetFirstTimeCustomerOk

`func (o *CreateDiscountInput) GetFirstTimeCustomerOk() (*bool, bool)`

GetFirstTimeCustomerOk returns a tuple with the FirstTimeCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTimeCustomer

`func (o *CreateDiscountInput) SetFirstTimeCustomer(v bool)`

SetFirstTimeCustomer sets FirstTimeCustomer field to given value.

### HasFirstTimeCustomer

`func (o *CreateDiscountInput) HasFirstTimeCustomer() bool`

HasFirstTimeCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


