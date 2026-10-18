# UpdateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional]
**Description** | Pointer to **NullableString** |  | [optional]
**AmountMinor** | Pointer to **int64** |  | [optional]
**Currency** | Pointer to **string** |  | [optional]
**Interval** | Pointer to **NullableString** |  | [optional]
**IntervalCount** | Pointer to **NullableInt32** |  | [optional]
**TrialDays** | Pointer to **int32** |  | [optional]
**IsActive** | Pointer to **bool** |  | [optional]
**IsArchived** | Pointer to **bool** |  | [optional]
**Slug** | Pointer to **NullableString** | URL handle, unique among this merchant&#39;s products. Pass null to clear it. | [optional]
**Images** | Pointer to [**[]ProductImage**](ProductImage.md) | Ordered extra images (at most 20). Replaces the list. The main image stays in image_url. | [optional]
**SeoTitle** | Pointer to **NullableString** | Search and social title (at most 70 characters). Defaults to the name. | [optional]
**SeoDescription** | Pointer to **NullableString** | Meta description (at most 320 characters). | [optional]
**StoreVisibility** | Pointer to **string** | Whether the product is listed on the Storefront. hidden: only payment links and invoices. scheduled: listed from store_visible_from. | [optional]
**StoreVisibleFrom** | Pointer to **NullableTime** | When a scheduled product appears on the Storefront (ISO 8601). Required when store_visibility is scheduled. | [optional]
**FulfillmentType** | Pointer to **string** | Fulfillment type for the product. | [optional]
**TrackInventory** | Pointer to **bool** | Turn stock tracking on (requires stock_quantity) or off. One-time products only. Stock is stored in the product&#39;s availability record. | [optional]
**StockQuantity** | Pointer to **NullableInt32** | Deprecated. Sets the units available to sell in the product&#39;s availability record. Required (non-negative) when track_inventory is true. | [optional]
**Tags** | Pointer to **[]string** | Tags for Storefront filters and search. Replaces the product&#39;s tags. At most 50, each up to 50 characters. | [optional]
**BillingAnchorRule** | Pointer to **string** | How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month. | [optional]
**BillingAnchorDay** | Pointer to **NullableInt32** |  | [optional]
**FirstChargeBehavior** | Pointer to **string** | When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date. | [optional]
**EndBehavior** | Pointer to **string** | How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles. | [optional]
**MaxPayments** | Pointer to **NullableInt32** |  | [optional]
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata for the product | [optional]
**ImageUrl** | Pointer to **NullableString** | Product image S3 key or public URL. Pass null to remove. | [optional]

## Methods

### NewUpdateProductRequest

`func NewUpdateProductRequest() *UpdateProductRequest`

NewUpdateProductRequest instantiates a new UpdateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductRequestWithDefaults

`func NewUpdateProductRequestWithDefaults() *UpdateProductRequest`

NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProductRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAmountMinor

`func (o *UpdateProductRequest) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *UpdateProductRequest) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *UpdateProductRequest) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *UpdateProductRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateProductRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateProductRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateProductRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *UpdateProductRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateProductRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateProductRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateProductRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *UpdateProductRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *UpdateProductRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *UpdateProductRequest) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UpdateProductRequest) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UpdateProductRequest) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UpdateProductRequest) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *UpdateProductRequest) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *UpdateProductRequest) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialDays

`func (o *UpdateProductRequest) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *UpdateProductRequest) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *UpdateProductRequest) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *UpdateProductRequest) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateProductRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateProductRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateProductRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateProductRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsArchived

`func (o *UpdateProductRequest) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UpdateProductRequest) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UpdateProductRequest) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *UpdateProductRequest) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetSlug

`func (o *UpdateProductRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *UpdateProductRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *UpdateProductRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *UpdateProductRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *UpdateProductRequest) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *UpdateProductRequest) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetImages

`func (o *UpdateProductRequest) GetImages() []ProductImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *UpdateProductRequest) GetImagesOk() (*[]ProductImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *UpdateProductRequest) SetImages(v []ProductImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *UpdateProductRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetSeoTitle

`func (o *UpdateProductRequest) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *UpdateProductRequest) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *UpdateProductRequest) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *UpdateProductRequest) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *UpdateProductRequest) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *UpdateProductRequest) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoDescription

`func (o *UpdateProductRequest) GetSeoDescription() string`

GetSeoDescription returns the SeoDescription field if non-nil, zero value otherwise.

### GetSeoDescriptionOk

`func (o *UpdateProductRequest) GetSeoDescriptionOk() (*string, bool)`

GetSeoDescriptionOk returns a tuple with the SeoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoDescription

`func (o *UpdateProductRequest) SetSeoDescription(v string)`

SetSeoDescription sets SeoDescription field to given value.

### HasSeoDescription

`func (o *UpdateProductRequest) HasSeoDescription() bool`

HasSeoDescription returns a boolean if a field has been set.

### SetSeoDescriptionNil

`func (o *UpdateProductRequest) SetSeoDescriptionNil(b bool)`

 SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil

### UnsetSeoDescription
`func (o *UpdateProductRequest) UnsetSeoDescription()`

UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
### GetStoreVisibility

`func (o *UpdateProductRequest) GetStoreVisibility() string`

GetStoreVisibility returns the StoreVisibility field if non-nil, zero value otherwise.

### GetStoreVisibilityOk

`func (o *UpdateProductRequest) GetStoreVisibilityOk() (*string, bool)`

GetStoreVisibilityOk returns a tuple with the StoreVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreVisibility

`func (o *UpdateProductRequest) SetStoreVisibility(v string)`

SetStoreVisibility sets StoreVisibility field to given value.

### HasStoreVisibility

`func (o *UpdateProductRequest) HasStoreVisibility() bool`

HasStoreVisibility returns a boolean if a field has been set.

### GetStoreVisibleFrom

`func (o *UpdateProductRequest) GetStoreVisibleFrom() time.Time`

GetStoreVisibleFrom returns the StoreVisibleFrom field if non-nil, zero value otherwise.

### GetStoreVisibleFromOk

`func (o *UpdateProductRequest) GetStoreVisibleFromOk() (*time.Time, bool)`

GetStoreVisibleFromOk returns a tuple with the StoreVisibleFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreVisibleFrom

`func (o *UpdateProductRequest) SetStoreVisibleFrom(v time.Time)`

SetStoreVisibleFrom sets StoreVisibleFrom field to given value.

### HasStoreVisibleFrom

`func (o *UpdateProductRequest) HasStoreVisibleFrom() bool`

HasStoreVisibleFrom returns a boolean if a field has been set.

### SetStoreVisibleFromNil

`func (o *UpdateProductRequest) SetStoreVisibleFromNil(b bool)`

 SetStoreVisibleFromNil sets the value for StoreVisibleFrom to be an explicit nil

### UnsetStoreVisibleFrom
`func (o *UpdateProductRequest) UnsetStoreVisibleFrom()`

UnsetStoreVisibleFrom ensures that no value is present for StoreVisibleFrom, not even an explicit nil
### GetFulfillmentType

`func (o *UpdateProductRequest) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *UpdateProductRequest) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *UpdateProductRequest) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *UpdateProductRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetTrackInventory

`func (o *UpdateProductRequest) GetTrackInventory() bool`

GetTrackInventory returns the TrackInventory field if non-nil, zero value otherwise.

### GetTrackInventoryOk

`func (o *UpdateProductRequest) GetTrackInventoryOk() (*bool, bool)`

GetTrackInventoryOk returns a tuple with the TrackInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackInventory

`func (o *UpdateProductRequest) SetTrackInventory(v bool)`

SetTrackInventory sets TrackInventory field to given value.

### HasTrackInventory

`func (o *UpdateProductRequest) HasTrackInventory() bool`

HasTrackInventory returns a boolean if a field has been set.

### GetStockQuantity

`func (o *UpdateProductRequest) GetStockQuantity() int32`

GetStockQuantity returns the StockQuantity field if non-nil, zero value otherwise.

### GetStockQuantityOk

`func (o *UpdateProductRequest) GetStockQuantityOk() (*int32, bool)`

GetStockQuantityOk returns a tuple with the StockQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockQuantity

`func (o *UpdateProductRequest) SetStockQuantity(v int32)`

SetStockQuantity sets StockQuantity field to given value.

### HasStockQuantity

`func (o *UpdateProductRequest) HasStockQuantity() bool`

HasStockQuantity returns a boolean if a field has been set.

### SetStockQuantityNil

`func (o *UpdateProductRequest) SetStockQuantityNil(b bool)`

 SetStockQuantityNil sets the value for StockQuantity to be an explicit nil

### UnsetStockQuantity
`func (o *UpdateProductRequest) UnsetStockQuantity()`

UnsetStockQuantity ensures that no value is present for StockQuantity, not even an explicit nil
### GetTags

`func (o *UpdateProductRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateProductRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateProductRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateProductRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBillingAnchorRule

`func (o *UpdateProductRequest) GetBillingAnchorRule() string`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *UpdateProductRequest) GetBillingAnchorRuleOk() (*string, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *UpdateProductRequest) SetBillingAnchorRule(v string)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *UpdateProductRequest) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *UpdateProductRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *UpdateProductRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *UpdateProductRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *UpdateProductRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *UpdateProductRequest) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *UpdateProductRequest) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetFirstChargeBehavior

`func (o *UpdateProductRequest) GetFirstChargeBehavior() string`

GetFirstChargeBehavior returns the FirstChargeBehavior field if non-nil, zero value otherwise.

### GetFirstChargeBehaviorOk

`func (o *UpdateProductRequest) GetFirstChargeBehaviorOk() (*string, bool)`

GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeBehavior

`func (o *UpdateProductRequest) SetFirstChargeBehavior(v string)`

SetFirstChargeBehavior sets FirstChargeBehavior field to given value.

### HasFirstChargeBehavior

`func (o *UpdateProductRequest) HasFirstChargeBehavior() bool`

HasFirstChargeBehavior returns a boolean if a field has been set.

### GetEndBehavior

`func (o *UpdateProductRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *UpdateProductRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *UpdateProductRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *UpdateProductRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetMaxPayments

`func (o *UpdateProductRequest) GetMaxPayments() int32`

GetMaxPayments returns the MaxPayments field if non-nil, zero value otherwise.

### GetMaxPaymentsOk

`func (o *UpdateProductRequest) GetMaxPaymentsOk() (*int32, bool)`

GetMaxPaymentsOk returns a tuple with the MaxPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayments

`func (o *UpdateProductRequest) SetMaxPayments(v int32)`

SetMaxPayments sets MaxPayments field to given value.

### HasMaxPayments

`func (o *UpdateProductRequest) HasMaxPayments() bool`

HasMaxPayments returns a boolean if a field has been set.

### SetMaxPaymentsNil

`func (o *UpdateProductRequest) SetMaxPaymentsNil(b bool)`

 SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil

### UnsetMaxPayments
`func (o *UpdateProductRequest) UnsetMaxPayments()`

UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
### GetMetadata

`func (o *UpdateProductRequest) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateProductRequest) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateProductRequest) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateProductRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImageUrl

`func (o *UpdateProductRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UpdateProductRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UpdateProductRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UpdateProductRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *UpdateProductRequest) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *UpdateProductRequest) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
