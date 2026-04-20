# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier (UUID) | 
**Object** | **string** | Object type | 
**ProductId** | **string** | User-facing product identifier | 
**Name** | **string** | Product display name | 
**Description** | Pointer to **NullableString** | Product description | [optional] 
**Kind** | **string** | Product kind | 
**PricingModel** | **string** | How the product is priced | 
**AmountMinor** | **NullableInt32** | Price in minor units (cents/pence) | 
**Currency** | **string** | Three-letter ISO currency code | 
**Interval** | Pointer to **NullableString** | Billing interval (day, week, month, year) | [optional] 
**IntervalCount** | Pointer to **NullableInt32** | Number of intervals between billings | [optional] 
**TrialDays** | **int32** | Free trial period in days | 
**FulfillmentType** | **string** | Fulfillment type | 
**BillingAnchorRule** | Pointer to **NullableString** | Billing date calculation rule | [optional] 
**BillingAnchorDay** | Pointer to **NullableInt32** | Day of month for billing (1-31) | [optional] 
**FirstChargeBehavior** | Pointer to **NullableString** | First payment timing | [optional] 
**EndBehavior** | Pointer to **NullableString** | Subscription end behavior | [optional] 
**MaxPayments** | Pointer to **NullableInt32** | Max billing cycles for fixed-payment subscriptions | [optional] 
**UsageMeterId** | Pointer to **NullableString** | Associated usage meter ID | [optional] 
**Slug** | Pointer to **NullableString** | URL-friendly slug | [optional] 
**IsActive** | **bool** | Whether the product is active | 
**IsArchived** | **bool** | Whether the product is archived | 
**ImageUrl** | Pointer to **NullableString** | Product image URL | [optional] 
**TaxBehavior** | Pointer to **NullableString** | Tax behavior (exclusive, inclusive, location) | [optional] 
**TaxCode** | Pointer to **NullableString** | Tax code for tax calculation | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp (ISO 8601) | 
**UpdatedAt** | **time.Time** | Last update timestamp (ISO 8601) | 

## Methods

### NewProduct

`func NewProduct(id string, object string, productId string, name string, kind string, pricingModel string, amountMinor NullableInt32, currency string, trialDays int32, fulfillmentType string, isActive bool, isArchived bool, createdAt time.Time, updatedAt time.Time, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Product) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Product) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Product) SetObject(v string)`

SetObject sets Object field to given value.


### GetProductId

`func (o *Product) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Product) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Product) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *Product) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Product) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Product) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Product) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Product) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *Product) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Product) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Product) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPricingModel

`func (o *Product) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *Product) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *Product) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetAmountMinor

`func (o *Product) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *Product) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *Product) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *Product) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *Product) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *Product) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Product) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Product) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetInterval

`func (o *Product) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *Product) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *Product) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *Product) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *Product) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *Product) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *Product) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *Product) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *Product) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *Product) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *Product) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *Product) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialDays

`func (o *Product) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *Product) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *Product) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.


### GetFulfillmentType

`func (o *Product) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *Product) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *Product) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.


### GetBillingAnchorRule

`func (o *Product) GetBillingAnchorRule() string`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *Product) GetBillingAnchorRuleOk() (*string, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *Product) SetBillingAnchorRule(v string)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *Product) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### SetBillingAnchorRuleNil

`func (o *Product) SetBillingAnchorRuleNil(b bool)`

 SetBillingAnchorRuleNil sets the value for BillingAnchorRule to be an explicit nil

### UnsetBillingAnchorRule
`func (o *Product) UnsetBillingAnchorRule()`

UnsetBillingAnchorRule ensures that no value is present for BillingAnchorRule, not even an explicit nil
### GetBillingAnchorDay

`func (o *Product) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *Product) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *Product) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *Product) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *Product) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *Product) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetFirstChargeBehavior

`func (o *Product) GetFirstChargeBehavior() string`

GetFirstChargeBehavior returns the FirstChargeBehavior field if non-nil, zero value otherwise.

### GetFirstChargeBehaviorOk

`func (o *Product) GetFirstChargeBehaviorOk() (*string, bool)`

GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeBehavior

`func (o *Product) SetFirstChargeBehavior(v string)`

SetFirstChargeBehavior sets FirstChargeBehavior field to given value.

### HasFirstChargeBehavior

`func (o *Product) HasFirstChargeBehavior() bool`

HasFirstChargeBehavior returns a boolean if a field has been set.

### SetFirstChargeBehaviorNil

`func (o *Product) SetFirstChargeBehaviorNil(b bool)`

 SetFirstChargeBehaviorNil sets the value for FirstChargeBehavior to be an explicit nil

### UnsetFirstChargeBehavior
`func (o *Product) UnsetFirstChargeBehavior()`

UnsetFirstChargeBehavior ensures that no value is present for FirstChargeBehavior, not even an explicit nil
### GetEndBehavior

`func (o *Product) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *Product) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *Product) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *Product) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### SetEndBehaviorNil

`func (o *Product) SetEndBehaviorNil(b bool)`

 SetEndBehaviorNil sets the value for EndBehavior to be an explicit nil

### UnsetEndBehavior
`func (o *Product) UnsetEndBehavior()`

UnsetEndBehavior ensures that no value is present for EndBehavior, not even an explicit nil
### GetMaxPayments

`func (o *Product) GetMaxPayments() int32`

GetMaxPayments returns the MaxPayments field if non-nil, zero value otherwise.

### GetMaxPaymentsOk

`func (o *Product) GetMaxPaymentsOk() (*int32, bool)`

GetMaxPaymentsOk returns a tuple with the MaxPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayments

`func (o *Product) SetMaxPayments(v int32)`

SetMaxPayments sets MaxPayments field to given value.

### HasMaxPayments

`func (o *Product) HasMaxPayments() bool`

HasMaxPayments returns a boolean if a field has been set.

### SetMaxPaymentsNil

`func (o *Product) SetMaxPaymentsNil(b bool)`

 SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil

### UnsetMaxPayments
`func (o *Product) UnsetMaxPayments()`

UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
### GetUsageMeterId

`func (o *Product) GetUsageMeterId() string`

GetUsageMeterId returns the UsageMeterId field if non-nil, zero value otherwise.

### GetUsageMeterIdOk

`func (o *Product) GetUsageMeterIdOk() (*string, bool)`

GetUsageMeterIdOk returns a tuple with the UsageMeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterId

`func (o *Product) SetUsageMeterId(v string)`

SetUsageMeterId sets UsageMeterId field to given value.

### HasUsageMeterId

`func (o *Product) HasUsageMeterId() bool`

HasUsageMeterId returns a boolean if a field has been set.

### SetUsageMeterIdNil

`func (o *Product) SetUsageMeterIdNil(b bool)`

 SetUsageMeterIdNil sets the value for UsageMeterId to be an explicit nil

### UnsetUsageMeterId
`func (o *Product) UnsetUsageMeterId()`

UnsetUsageMeterId ensures that no value is present for UsageMeterId, not even an explicit nil
### GetSlug

`func (o *Product) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Product) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Product) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Product) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *Product) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *Product) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetIsActive

`func (o *Product) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Product) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Product) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsArchived

`func (o *Product) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Product) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Product) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetImageUrl

`func (o *Product) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Product) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Product) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Product) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *Product) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *Product) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetTaxBehavior

`func (o *Product) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *Product) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *Product) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *Product) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### SetTaxBehaviorNil

`func (o *Product) SetTaxBehaviorNil(b bool)`

 SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil

### UnsetTaxBehavior
`func (o *Product) UnsetTaxBehavior()`

UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil
### GetTaxCode

`func (o *Product) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Product) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Product) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Product) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *Product) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *Product) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetMetadata

`func (o *Product) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Product) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Product) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Product) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Product) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Product) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


