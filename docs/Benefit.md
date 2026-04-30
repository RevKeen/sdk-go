# Benefit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**BenefitType** | **string** |  | 
**BenefitKey** | **string** |  | 
**Category** | **NullableString** |  | 
**IconUrl** | **NullableString** |  | 
**DisplayOrder** | **NullableString** |  | 
**IsActive** | **NullableBool** |  | 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**Config** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewBenefit

`func NewBenefit(id string, name string, description NullableString, benefitType string, benefitKey string, category NullableString, iconUrl NullableString, displayOrder NullableString, isActive NullableBool, ) *Benefit`

NewBenefit instantiates a new Benefit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenefitWithDefaults

`func NewBenefitWithDefaults() *Benefit`

NewBenefitWithDefaults instantiates a new Benefit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Benefit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Benefit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Benefit) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Benefit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Benefit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Benefit) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Benefit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Benefit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Benefit) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Benefit) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Benefit) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBenefitType

`func (o *Benefit) GetBenefitType() string`

GetBenefitType returns the BenefitType field if non-nil, zero value otherwise.

### GetBenefitTypeOk

`func (o *Benefit) GetBenefitTypeOk() (*string, bool)`

GetBenefitTypeOk returns a tuple with the BenefitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitType

`func (o *Benefit) SetBenefitType(v string)`

SetBenefitType sets BenefitType field to given value.


### GetBenefitKey

`func (o *Benefit) GetBenefitKey() string`

GetBenefitKey returns the BenefitKey field if non-nil, zero value otherwise.

### GetBenefitKeyOk

`func (o *Benefit) GetBenefitKeyOk() (*string, bool)`

GetBenefitKeyOk returns a tuple with the BenefitKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenefitKey

`func (o *Benefit) SetBenefitKey(v string)`

SetBenefitKey sets BenefitKey field to given value.


### GetCategory

`func (o *Benefit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Benefit) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Benefit) SetCategory(v string)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *Benefit) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Benefit) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIconUrl

`func (o *Benefit) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Benefit) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Benefit) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *Benefit) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *Benefit) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetDisplayOrder

`func (o *Benefit) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *Benefit) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *Benefit) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.


### SetDisplayOrderNil

`func (o *Benefit) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *Benefit) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetIsActive

`func (o *Benefit) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Benefit) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Benefit) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### SetIsActiveNil

`func (o *Benefit) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *Benefit) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetDefaultValue

`func (o *Benefit) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Benefit) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Benefit) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *Benefit) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *Benefit) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *Benefit) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetConfig

`func (o *Benefit) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Benefit) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Benefit) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Benefit) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *Benefit) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Benefit) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


