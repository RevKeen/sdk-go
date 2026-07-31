# StorefrontProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**ProductId** | **NullableString** |  | 
**Slug** | **NullableString** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Kind** | **string** |  | 
**PricingModel** | **string** |  | 
**Currency** | **string** |  | 
**ImageUrl** | **NullableString** |  | 
**DefaultPriceId** | **NullableString** |  | 
**Prices** | [**[]StorefrontPrice**](StorefrontPrice.md) |  | 
**TrialDays** | **int32** |  | 
**UsageMeterId** | **NullableString** |  | 
**TaxBehavior** | **NullableString** |  | 
**TaxCode** | **NullableString** |  | 
**Availability** | [**StorefrontAvailability**](StorefrontAvailability.md) |  | 

## Methods

### NewStorefrontProduct

`func NewStorefrontProduct(id string, object string, productId NullableString, slug NullableString, name string, description NullableString, kind string, pricingModel string, currency string, imageUrl NullableString, defaultPriceId NullableString, prices []StorefrontPrice, trialDays int32, usageMeterId NullableString, taxBehavior NullableString, taxCode NullableString, availability StorefrontAvailability, ) *StorefrontProduct`

NewStorefrontProduct instantiates a new StorefrontProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontProductWithDefaults

`func NewStorefrontProductWithDefaults() *StorefrontProduct`

NewStorefrontProductWithDefaults instantiates a new StorefrontProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorefrontProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorefrontProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorefrontProduct) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *StorefrontProduct) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StorefrontProduct) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StorefrontProduct) SetObject(v string)`

SetObject sets Object field to given value.


### GetProductId

`func (o *StorefrontProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *StorefrontProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *StorefrontProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### SetProductIdNil

`func (o *StorefrontProduct) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *StorefrontProduct) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetSlug

`func (o *StorefrontProduct) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StorefrontProduct) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StorefrontProduct) SetSlug(v string)`

SetSlug sets Slug field to given value.


### SetSlugNil

`func (o *StorefrontProduct) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *StorefrontProduct) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *StorefrontProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorefrontProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorefrontProduct) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StorefrontProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorefrontProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorefrontProduct) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *StorefrontProduct) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StorefrontProduct) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *StorefrontProduct) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StorefrontProduct) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StorefrontProduct) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPricingModel

`func (o *StorefrontProduct) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *StorefrontProduct) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *StorefrontProduct) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetCurrency

`func (o *StorefrontProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StorefrontProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StorefrontProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetImageUrl

`func (o *StorefrontProduct) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *StorefrontProduct) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *StorefrontProduct) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### SetImageUrlNil

`func (o *StorefrontProduct) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *StorefrontProduct) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDefaultPriceId

`func (o *StorefrontProduct) GetDefaultPriceId() string`

GetDefaultPriceId returns the DefaultPriceId field if non-nil, zero value otherwise.

### GetDefaultPriceIdOk

`func (o *StorefrontProduct) GetDefaultPriceIdOk() (*string, bool)`

GetDefaultPriceIdOk returns a tuple with the DefaultPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriceId

`func (o *StorefrontProduct) SetDefaultPriceId(v string)`

SetDefaultPriceId sets DefaultPriceId field to given value.


### SetDefaultPriceIdNil

`func (o *StorefrontProduct) SetDefaultPriceIdNil(b bool)`

 SetDefaultPriceIdNil sets the value for DefaultPriceId to be an explicit nil

### UnsetDefaultPriceId
`func (o *StorefrontProduct) UnsetDefaultPriceId()`

UnsetDefaultPriceId ensures that no value is present for DefaultPriceId, not even an explicit nil
### GetPrices

`func (o *StorefrontProduct) GetPrices() []StorefrontPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *StorefrontProduct) GetPricesOk() (*[]StorefrontPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *StorefrontProduct) SetPrices(v []StorefrontPrice)`

SetPrices sets Prices field to given value.


### GetTrialDays

`func (o *StorefrontProduct) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *StorefrontProduct) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *StorefrontProduct) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.


### GetUsageMeterId

`func (o *StorefrontProduct) GetUsageMeterId() string`

GetUsageMeterId returns the UsageMeterId field if non-nil, zero value otherwise.

### GetUsageMeterIdOk

`func (o *StorefrontProduct) GetUsageMeterIdOk() (*string, bool)`

GetUsageMeterIdOk returns a tuple with the UsageMeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterId

`func (o *StorefrontProduct) SetUsageMeterId(v string)`

SetUsageMeterId sets UsageMeterId field to given value.


### SetUsageMeterIdNil

`func (o *StorefrontProduct) SetUsageMeterIdNil(b bool)`

 SetUsageMeterIdNil sets the value for UsageMeterId to be an explicit nil

### UnsetUsageMeterId
`func (o *StorefrontProduct) UnsetUsageMeterId()`

UnsetUsageMeterId ensures that no value is present for UsageMeterId, not even an explicit nil
### GetTaxBehavior

`func (o *StorefrontProduct) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *StorefrontProduct) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *StorefrontProduct) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.


### SetTaxBehaviorNil

`func (o *StorefrontProduct) SetTaxBehaviorNil(b bool)`

 SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil

### UnsetTaxBehavior
`func (o *StorefrontProduct) UnsetTaxBehavior()`

UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil
### GetTaxCode

`func (o *StorefrontProduct) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *StorefrontProduct) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *StorefrontProduct) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.


### SetTaxCodeNil

`func (o *StorefrontProduct) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *StorefrontProduct) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetAvailability

`func (o *StorefrontProduct) GetAvailability() StorefrontAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *StorefrontProduct) GetAvailabilityOk() (*StorefrontAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *StorefrontProduct) SetAvailability(v StorefrontAvailability)`

SetAvailability sets Availability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


