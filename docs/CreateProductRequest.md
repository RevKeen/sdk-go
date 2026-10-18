# CreateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  |
**Name** | **string** |  |
**Description** | Pointer to **NullableString** |  | [optional]
**Kind** | **string** |  |
**PricingModel** | **string** |  |
**AmountMinor** | **int64** |  |
**Currency** | Pointer to **string** | Three-letter ISO currency code. Defaults to the merchant&#39;s default currency. | [optional]
**Interval** | Pointer to **NullableString** |  | [optional]
**IntervalCount** | Pointer to **NullableInt32** |  | [optional]
**TrialDays** | Pointer to **int32** |  | [optional]
**UsageMeterId** | Pointer to **NullableString** |  | [optional]
**Slug** | Pointer to **string** | URL handle, unique among this merchant&#39;s products. Normalised to lower-case letters, digits and hyphens. Derived from product_id when omitted. | [optional]
**Images** | Pointer to [**[]ProductImage**](ProductImage.md) | Ordered extra images (at most 20). Replaces the list. The main image stays in image_url. | [optional]
**SeoTitle** | Pointer to **NullableString** | Search and social title (at most 70 characters). Defaults to the name. | [optional]
**SeoDescription** | Pointer to **NullableString** | Meta description (at most 320 characters). | [optional]
**StoreVisibility** | Pointer to **string** | Whether the product is listed on the Storefront. hidden: only payment links and invoices. scheduled: listed from store_visible_from. | [optional]
**StoreVisibleFrom** | Pointer to **NullableTime** | When a scheduled product appears on the Storefront (ISO 8601). Required when store_visibility is scheduled. | [optional]
**FulfillmentType** | Pointer to **string** | Fulfillment type for the product. Defaults to &#39;none&#39; (service). | [optional] [default to "none"]
**TrackInventory** | Pointer to **bool** | Track stock for this product (one-time products only; requires stock_quantity). Stock is stored in the product&#39;s availability record, not on the product. | [optional] [default to false]
**StockQuantity** | Pointer to **NullableInt32** | Deprecated. Sets the units available to sell in the product&#39;s availability record. Required (non-negative) when track_inventory is true. | [optional]
**Tags** | Pointer to **[]string** | Tags for Storefront filters and search. Replaces the product&#39;s tags. At most 50, each up to 50 characters. | [optional]
**BillingAnchorRule** | Pointer to **string** | How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month. | [optional]
**BillingAnchorDay** | Pointer to **NullableInt32** | Day of month (1-31) when billing_anchor_rule is &#39;day_of_month&#39; | [optional]
**FirstChargeBehavior** | Pointer to **string** | When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date. | [optional]
**EndBehavior** | Pointer to **string** | How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles. | [optional]
**MaxPayments** | Pointer to **NullableInt32** | Max billing cycles when end_behavior is &#39;fixed_payments&#39; | [optional]
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata for the product | [optional]
**ImageUrl** | Pointer to **NullableString** | Product image S3 key or public URL (upload via /internal/storage first) | [optional]

## Methods

### NewCreateProductRequest

`func NewCreateProductRequest(productId string, name string, kind string, pricingModel string, amountMinor int64, ) *CreateProductRequest`

NewCreateProductRequest instantiates a new CreateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductRequestWithDefaults

`func NewCreateProductRequestWithDefaults() *CreateProductRequest`

NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateProductRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateProductRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateProductRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *CreateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *CreateProductRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateProductRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateProductRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPricingModel

`func (o *CreateProductRequest) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreateProductRequest) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreateProductRequest) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetAmountMinor

`func (o *CreateProductRequest) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateProductRequest) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateProductRequest) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *CreateProductRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateProductRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateProductRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *CreateProductRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateProductRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateProductRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreateProductRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *CreateProductRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *CreateProductRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *CreateProductRequest) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *CreateProductRequest) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *CreateProductRequest) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *CreateProductRequest) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *CreateProductRequest) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *CreateProductRequest) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialDays

`func (o *CreateProductRequest) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *CreateProductRequest) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *CreateProductRequest) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *CreateProductRequest) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetUsageMeterId

`func (o *CreateProductRequest) GetUsageMeterId() string`

GetUsageMeterId returns the UsageMeterId field if non-nil, zero value otherwise.

### GetUsageMeterIdOk

`func (o *CreateProductRequest) GetUsageMeterIdOk() (*string, bool)`

GetUsageMeterIdOk returns a tuple with the UsageMeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterId

`func (o *CreateProductRequest) SetUsageMeterId(v string)`

SetUsageMeterId sets UsageMeterId field to given value.

### HasUsageMeterId

`func (o *CreateProductRequest) HasUsageMeterId() bool`

HasUsageMeterId returns a boolean if a field has been set.

### SetUsageMeterIdNil

`func (o *CreateProductRequest) SetUsageMeterIdNil(b bool)`

 SetUsageMeterIdNil sets the value for UsageMeterId to be an explicit nil

### UnsetUsageMeterId
`func (o *CreateProductRequest) UnsetUsageMeterId()`

UnsetUsageMeterId ensures that no value is present for UsageMeterId, not even an explicit nil
### GetSlug

`func (o *CreateProductRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateProductRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateProductRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateProductRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetImages

`func (o *CreateProductRequest) GetImages() []ProductImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreateProductRequest) GetImagesOk() (*[]ProductImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreateProductRequest) SetImages(v []ProductImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *CreateProductRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetSeoTitle

`func (o *CreateProductRequest) GetSeoTitle() string`

GetSeoTitle returns the SeoTitle field if non-nil, zero value otherwise.

### GetSeoTitleOk

`func (o *CreateProductRequest) GetSeoTitleOk() (*string, bool)`

GetSeoTitleOk returns a tuple with the SeoTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoTitle

`func (o *CreateProductRequest) SetSeoTitle(v string)`

SetSeoTitle sets SeoTitle field to given value.

### HasSeoTitle

`func (o *CreateProductRequest) HasSeoTitle() bool`

HasSeoTitle returns a boolean if a field has been set.

### SetSeoTitleNil

`func (o *CreateProductRequest) SetSeoTitleNil(b bool)`

 SetSeoTitleNil sets the value for SeoTitle to be an explicit nil

### UnsetSeoTitle
`func (o *CreateProductRequest) UnsetSeoTitle()`

UnsetSeoTitle ensures that no value is present for SeoTitle, not even an explicit nil
### GetSeoDescription

`func (o *CreateProductRequest) GetSeoDescription() string`

GetSeoDescription returns the SeoDescription field if non-nil, zero value otherwise.

### GetSeoDescriptionOk

`func (o *CreateProductRequest) GetSeoDescriptionOk() (*string, bool)`

GetSeoDescriptionOk returns a tuple with the SeoDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeoDescription

`func (o *CreateProductRequest) SetSeoDescription(v string)`

SetSeoDescription sets SeoDescription field to given value.

### HasSeoDescription

`func (o *CreateProductRequest) HasSeoDescription() bool`

HasSeoDescription returns a boolean if a field has been set.

### SetSeoDescriptionNil

`func (o *CreateProductRequest) SetSeoDescriptionNil(b bool)`

 SetSeoDescriptionNil sets the value for SeoDescription to be an explicit nil

### UnsetSeoDescription
`func (o *CreateProductRequest) UnsetSeoDescription()`

UnsetSeoDescription ensures that no value is present for SeoDescription, not even an explicit nil
### GetStoreVisibility

`func (o *CreateProductRequest) GetStoreVisibility() string`

GetStoreVisibility returns the StoreVisibility field if non-nil, zero value otherwise.

### GetStoreVisibilityOk

`func (o *CreateProductRequest) GetStoreVisibilityOk() (*string, bool)`

GetStoreVisibilityOk returns a tuple with the StoreVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreVisibility

`func (o *CreateProductRequest) SetStoreVisibility(v string)`

SetStoreVisibility sets StoreVisibility field to given value.

### HasStoreVisibility

`func (o *CreateProductRequest) HasStoreVisibility() bool`

HasStoreVisibility returns a boolean if a field has been set.

### GetStoreVisibleFrom

`func (o *CreateProductRequest) GetStoreVisibleFrom() time.Time`

GetStoreVisibleFrom returns the StoreVisibleFrom field if non-nil, zero value otherwise.

### GetStoreVisibleFromOk

`func (o *CreateProductRequest) GetStoreVisibleFromOk() (*time.Time, bool)`

GetStoreVisibleFromOk returns a tuple with the StoreVisibleFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreVisibleFrom

`func (o *CreateProductRequest) SetStoreVisibleFrom(v time.Time)`

SetStoreVisibleFrom sets StoreVisibleFrom field to given value.

### HasStoreVisibleFrom

`func (o *CreateProductRequest) HasStoreVisibleFrom() bool`

HasStoreVisibleFrom returns a boolean if a field has been set.

### SetStoreVisibleFromNil

`func (o *CreateProductRequest) SetStoreVisibleFromNil(b bool)`

 SetStoreVisibleFromNil sets the value for StoreVisibleFrom to be an explicit nil

### UnsetStoreVisibleFrom
`func (o *CreateProductRequest) UnsetStoreVisibleFrom()`

UnsetStoreVisibleFrom ensures that no value is present for StoreVisibleFrom, not even an explicit nil
### GetFulfillmentType

`func (o *CreateProductRequest) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *CreateProductRequest) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *CreateProductRequest) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *CreateProductRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetTrackInventory

`func (o *CreateProductRequest) GetTrackInventory() bool`

GetTrackInventory returns the TrackInventory field if non-nil, zero value otherwise.

### GetTrackInventoryOk

`func (o *CreateProductRequest) GetTrackInventoryOk() (*bool, bool)`

GetTrackInventoryOk returns a tuple with the TrackInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackInventory

`func (o *CreateProductRequest) SetTrackInventory(v bool)`

SetTrackInventory sets TrackInventory field to given value.

### HasTrackInventory

`func (o *CreateProductRequest) HasTrackInventory() bool`

HasTrackInventory returns a boolean if a field has been set.

### GetStockQuantity

`func (o *CreateProductRequest) GetStockQuantity() int32`

GetStockQuantity returns the StockQuantity field if non-nil, zero value otherwise.

### GetStockQuantityOk

`func (o *CreateProductRequest) GetStockQuantityOk() (*int32, bool)`

GetStockQuantityOk returns a tuple with the StockQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockQuantity

`func (o *CreateProductRequest) SetStockQuantity(v int32)`

SetStockQuantity sets StockQuantity field to given value.

### HasStockQuantity

`func (o *CreateProductRequest) HasStockQuantity() bool`

HasStockQuantity returns a boolean if a field has been set.

### SetStockQuantityNil

`func (o *CreateProductRequest) SetStockQuantityNil(b bool)`

 SetStockQuantityNil sets the value for StockQuantity to be an explicit nil

### UnsetStockQuantity
`func (o *CreateProductRequest) UnsetStockQuantity()`

UnsetStockQuantity ensures that no value is present for StockQuantity, not even an explicit nil
### GetTags

`func (o *CreateProductRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateProductRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateProductRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateProductRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBillingAnchorRule

`func (o *CreateProductRequest) GetBillingAnchorRule() string`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *CreateProductRequest) GetBillingAnchorRuleOk() (*string, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *CreateProductRequest) SetBillingAnchorRule(v string)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *CreateProductRequest) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *CreateProductRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *CreateProductRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *CreateProductRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *CreateProductRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *CreateProductRequest) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *CreateProductRequest) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetFirstChargeBehavior

`func (o *CreateProductRequest) GetFirstChargeBehavior() string`

GetFirstChargeBehavior returns the FirstChargeBehavior field if non-nil, zero value otherwise.

### GetFirstChargeBehaviorOk

`func (o *CreateProductRequest) GetFirstChargeBehaviorOk() (*string, bool)`

GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeBehavior

`func (o *CreateProductRequest) SetFirstChargeBehavior(v string)`

SetFirstChargeBehavior sets FirstChargeBehavior field to given value.

### HasFirstChargeBehavior

`func (o *CreateProductRequest) HasFirstChargeBehavior() bool`

HasFirstChargeBehavior returns a boolean if a field has been set.

### GetEndBehavior

`func (o *CreateProductRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *CreateProductRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *CreateProductRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *CreateProductRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetMaxPayments

`func (o *CreateProductRequest) GetMaxPayments() int32`

GetMaxPayments returns the MaxPayments field if non-nil, zero value otherwise.

### GetMaxPaymentsOk

`func (o *CreateProductRequest) GetMaxPaymentsOk() (*int32, bool)`

GetMaxPaymentsOk returns a tuple with the MaxPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayments

`func (o *CreateProductRequest) SetMaxPayments(v int32)`

SetMaxPayments sets MaxPayments field to given value.

### HasMaxPayments

`func (o *CreateProductRequest) HasMaxPayments() bool`

HasMaxPayments returns a boolean if a field has been set.

### SetMaxPaymentsNil

`func (o *CreateProductRequest) SetMaxPaymentsNil(b bool)`

 SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil

### UnsetMaxPayments
`func (o *CreateProductRequest) UnsetMaxPayments()`

UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
### GetMetadata

`func (o *CreateProductRequest) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateProductRequest) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateProductRequest) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateProductRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateProductRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateProductRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateProductRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateProductRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateProductRequest) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateProductRequest) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
