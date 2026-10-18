# ProductCollectionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**Slug** | Pointer to **string** |  | [optional]
**Description** | Pointer to **NullableString** |  | [optional]
**ImageAssetId** | Pointer to **NullableString** |  | [optional]
**Position** | Pointer to **int32** |  | [optional]
**ProductIds** | Pointer to **[]string** | Member product ids, in display order. Replaces the membership when present. | [optional]
**ExpectedVersion** | Pointer to **int32** |  | [optional]

## Methods

### NewProductCollectionUpdateRequest

`func NewProductCollectionUpdateRequest() *ProductCollectionUpdateRequest`

NewProductCollectionUpdateRequest instantiates a new ProductCollectionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCollectionUpdateRequestWithDefaults

`func NewProductCollectionUpdateRequestWithDefaults() *ProductCollectionUpdateRequest`

NewProductCollectionUpdateRequestWithDefaults instantiates a new ProductCollectionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductCollectionUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCollectionUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCollectionUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProductCollectionUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ProductCollectionUpdateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProductCollectionUpdateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProductCollectionUpdateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ProductCollectionUpdateRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *ProductCollectionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCollectionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCollectionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductCollectionUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductCollectionUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductCollectionUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageAssetId

`func (o *ProductCollectionUpdateRequest) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *ProductCollectionUpdateRequest) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *ProductCollectionUpdateRequest) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *ProductCollectionUpdateRequest) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *ProductCollectionUpdateRequest) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *ProductCollectionUpdateRequest) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetPosition

`func (o *ProductCollectionUpdateRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductCollectionUpdateRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductCollectionUpdateRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProductCollectionUpdateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProductIds

`func (o *ProductCollectionUpdateRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *ProductCollectionUpdateRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *ProductCollectionUpdateRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *ProductCollectionUpdateRequest) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetExpectedVersion

`func (o *ProductCollectionUpdateRequest) GetExpectedVersion() int32`

GetExpectedVersion returns the ExpectedVersion field if non-nil, zero value otherwise.

### GetExpectedVersionOk

`func (o *ProductCollectionUpdateRequest) GetExpectedVersionOk() (*int32, bool)`

GetExpectedVersionOk returns a tuple with the ExpectedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedVersion

`func (o *ProductCollectionUpdateRequest) SetExpectedVersion(v int32)`

SetExpectedVersion sets ExpectedVersion field to given value.

### HasExpectedVersion

`func (o *ProductCollectionUpdateRequest) HasExpectedVersion() bool`

HasExpectedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
