# ProductCollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  |
**Slug** | **string** |  |
**Description** | Pointer to **NullableString** |  | [optional]
**ImageAssetId** | Pointer to **NullableString** |  | [optional]
**Position** | Pointer to **int32** |  | [optional]
**ProductIds** | Pointer to **[]string** | Member product ids, in display order. Replaces the membership when present. | [optional]

## Methods

### NewProductCollectionCreateRequest

`func NewProductCollectionCreateRequest(name string, slug string, ) *ProductCollectionCreateRequest`

NewProductCollectionCreateRequest instantiates a new ProductCollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCollectionCreateRequestWithDefaults

`func NewProductCollectionCreateRequestWithDefaults() *ProductCollectionCreateRequest`

NewProductCollectionCreateRequestWithDefaults instantiates a new ProductCollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductCollectionCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductCollectionCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductCollectionCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *ProductCollectionCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ProductCollectionCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ProductCollectionCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *ProductCollectionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductCollectionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductCollectionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductCollectionCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ProductCollectionCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductCollectionCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetImageAssetId

`func (o *ProductCollectionCreateRequest) GetImageAssetId() string`

GetImageAssetId returns the ImageAssetId field if non-nil, zero value otherwise.

### GetImageAssetIdOk

`func (o *ProductCollectionCreateRequest) GetImageAssetIdOk() (*string, bool)`

GetImageAssetIdOk returns a tuple with the ImageAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAssetId

`func (o *ProductCollectionCreateRequest) SetImageAssetId(v string)`

SetImageAssetId sets ImageAssetId field to given value.

### HasImageAssetId

`func (o *ProductCollectionCreateRequest) HasImageAssetId() bool`

HasImageAssetId returns a boolean if a field has been set.

### SetImageAssetIdNil

`func (o *ProductCollectionCreateRequest) SetImageAssetIdNil(b bool)`

 SetImageAssetIdNil sets the value for ImageAssetId to be an explicit nil

### UnsetImageAssetId
`func (o *ProductCollectionCreateRequest) UnsetImageAssetId()`

UnsetImageAssetId ensures that no value is present for ImageAssetId, not even an explicit nil
### GetPosition

`func (o *ProductCollectionCreateRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ProductCollectionCreateRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ProductCollectionCreateRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ProductCollectionCreateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProductIds

`func (o *ProductCollectionCreateRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *ProductCollectionCreateRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *ProductCollectionCreateRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *ProductCollectionCreateRequest) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
