# ProductImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Storage key returned by the upload (product-images/&lt;merchant&gt;/...) or an https URL. |
**AssetId** | Pointer to **NullableString** | Media library asset ID, when the image came from an upload. | [optional]
**Alt** | Pointer to **NullableString** | Alt text for shoppers using screen readers. | [optional]
**Source** | Pointer to **string** | Where the image came from. Defaults to merchant. Pexels images must keep photographer and attribution_url. | [optional]
**Photographer** | Pointer to **NullableString** | Photographer credit (required for Pexels). | [optional]
**AttributionUrl** | Pointer to **NullableString** | https link for the credit (required for Pexels). | [optional]

## Methods

### NewProductImage

`func NewProductImage(key string, ) *ProductImage`

NewProductImage instantiates a new ProductImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductImageWithDefaults

`func NewProductImageWithDefaults() *ProductImage`

NewProductImageWithDefaults instantiates a new ProductImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ProductImage) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProductImage) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProductImage) SetKey(v string)`

SetKey sets Key field to given value.


### GetAssetId

`func (o *ProductImage) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *ProductImage) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *ProductImage) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *ProductImage) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### SetAssetIdNil

`func (o *ProductImage) SetAssetIdNil(b bool)`

 SetAssetIdNil sets the value for AssetId to be an explicit nil

### UnsetAssetId
`func (o *ProductImage) UnsetAssetId()`

UnsetAssetId ensures that no value is present for AssetId, not even an explicit nil
### GetAlt

`func (o *ProductImage) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *ProductImage) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *ProductImage) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *ProductImage) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### SetAltNil

`func (o *ProductImage) SetAltNil(b bool)`

 SetAltNil sets the value for Alt to be an explicit nil

### UnsetAlt
`func (o *ProductImage) UnsetAlt()`

UnsetAlt ensures that no value is present for Alt, not even an explicit nil
### GetSource

`func (o *ProductImage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductImage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductImage) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProductImage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPhotographer

`func (o *ProductImage) GetPhotographer() string`

GetPhotographer returns the Photographer field if non-nil, zero value otherwise.

### GetPhotographerOk

`func (o *ProductImage) GetPhotographerOk() (*string, bool)`

GetPhotographerOk returns a tuple with the Photographer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotographer

`func (o *ProductImage) SetPhotographer(v string)`

SetPhotographer sets Photographer field to given value.

### HasPhotographer

`func (o *ProductImage) HasPhotographer() bool`

HasPhotographer returns a boolean if a field has been set.

### SetPhotographerNil

`func (o *ProductImage) SetPhotographerNil(b bool)`

 SetPhotographerNil sets the value for Photographer to be an explicit nil

### UnsetPhotographer
`func (o *ProductImage) UnsetPhotographer()`

UnsetPhotographer ensures that no value is present for Photographer, not even an explicit nil
### GetAttributionUrl

`func (o *ProductImage) GetAttributionUrl() string`

GetAttributionUrl returns the AttributionUrl field if non-nil, zero value otherwise.

### GetAttributionUrlOk

`func (o *ProductImage) GetAttributionUrlOk() (*string, bool)`

GetAttributionUrlOk returns a tuple with the AttributionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionUrl

`func (o *ProductImage) SetAttributionUrl(v string)`

SetAttributionUrl sets AttributionUrl field to given value.

### HasAttributionUrl

`func (o *ProductImage) HasAttributionUrl() bool`

HasAttributionUrl returns a boolean if a field has been set.

### SetAttributionUrlNil

`func (o *ProductImage) SetAttributionUrlNil(b bool)`

 SetAttributionUrlNil sets the value for AttributionUrl to be an explicit nil

### UnsetAttributionUrl
`func (o *ProductImage) UnsetAttributionUrl()`

UnsetAttributionUrl ensures that no value is present for AttributionUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
