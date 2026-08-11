# StorefrontPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProductId** | **string** |  | 
**Currency** | **string** |  | 
**UnitAmount** | **NullableInt32** |  | 
**UnitAmountDecimal** | **NullableString** |  | 
**Type** | **NullableString** |  | 
**Interval** | **NullableString** |  | 
**IntervalCount** | **NullableInt32** |  | 
**BillingScheme** | **string** |  | 
**UsageType** | **NullableString** |  | 
**PackageSize** | **NullableInt32** |  | 
**TrialPeriodDays** | **NullableInt32** |  | 

## Methods

### NewStorefrontPrice

`func NewStorefrontPrice(id string, productId string, currency string, unitAmount NullableInt32, unitAmountDecimal NullableString, type_ NullableString, interval NullableString, intervalCount NullableInt32, billingScheme string, usageType NullableString, packageSize NullableInt32, trialPeriodDays NullableInt32, ) *StorefrontPrice`

NewStorefrontPrice instantiates a new StorefrontPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontPriceWithDefaults

`func NewStorefrontPriceWithDefaults() *StorefrontPrice`

NewStorefrontPriceWithDefaults instantiates a new StorefrontPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorefrontPrice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorefrontPrice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorefrontPrice) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *StorefrontPrice) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *StorefrontPrice) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *StorefrontPrice) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetCurrency

`func (o *StorefrontPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StorefrontPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StorefrontPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitAmount

`func (o *StorefrontPrice) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *StorefrontPrice) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *StorefrontPrice) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### SetUnitAmountNil

`func (o *StorefrontPrice) SetUnitAmountNil(b bool)`

 SetUnitAmountNil sets the value for UnitAmount to be an explicit nil

### UnsetUnitAmount
`func (o *StorefrontPrice) UnsetUnitAmount()`

UnsetUnitAmount ensures that no value is present for UnitAmount, not even an explicit nil
### GetUnitAmountDecimal

`func (o *StorefrontPrice) GetUnitAmountDecimal() string`

GetUnitAmountDecimal returns the UnitAmountDecimal field if non-nil, zero value otherwise.

### GetUnitAmountDecimalOk

`func (o *StorefrontPrice) GetUnitAmountDecimalOk() (*string, bool)`

GetUnitAmountDecimalOk returns a tuple with the UnitAmountDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountDecimal

`func (o *StorefrontPrice) SetUnitAmountDecimal(v string)`

SetUnitAmountDecimal sets UnitAmountDecimal field to given value.


### SetUnitAmountDecimalNil

`func (o *StorefrontPrice) SetUnitAmountDecimalNil(b bool)`

 SetUnitAmountDecimalNil sets the value for UnitAmountDecimal to be an explicit nil

### UnsetUnitAmountDecimal
`func (o *StorefrontPrice) UnsetUnitAmountDecimal()`

UnsetUnitAmountDecimal ensures that no value is present for UnitAmountDecimal, not even an explicit nil
### GetType

`func (o *StorefrontPrice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorefrontPrice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorefrontPrice) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StorefrontPrice) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StorefrontPrice) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetInterval

`func (o *StorefrontPrice) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *StorefrontPrice) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *StorefrontPrice) SetInterval(v string)`

SetInterval sets Interval field to given value.


### SetIntervalNil

`func (o *StorefrontPrice) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *StorefrontPrice) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *StorefrontPrice) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *StorefrontPrice) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *StorefrontPrice) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.


### SetIntervalCountNil

`func (o *StorefrontPrice) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *StorefrontPrice) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetBillingScheme

`func (o *StorefrontPrice) GetBillingScheme() string`

GetBillingScheme returns the BillingScheme field if non-nil, zero value otherwise.

### GetBillingSchemeOk

`func (o *StorefrontPrice) GetBillingSchemeOk() (*string, bool)`

GetBillingSchemeOk returns a tuple with the BillingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingScheme

`func (o *StorefrontPrice) SetBillingScheme(v string)`

SetBillingScheme sets BillingScheme field to given value.


### GetUsageType

`func (o *StorefrontPrice) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *StorefrontPrice) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *StorefrontPrice) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.


### SetUsageTypeNil

`func (o *StorefrontPrice) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *StorefrontPrice) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetPackageSize

`func (o *StorefrontPrice) GetPackageSize() int32`

GetPackageSize returns the PackageSize field if non-nil, zero value otherwise.

### GetPackageSizeOk

`func (o *StorefrontPrice) GetPackageSizeOk() (*int32, bool)`

GetPackageSizeOk returns a tuple with the PackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSize

`func (o *StorefrontPrice) SetPackageSize(v int32)`

SetPackageSize sets PackageSize field to given value.


### SetPackageSizeNil

`func (o *StorefrontPrice) SetPackageSizeNil(b bool)`

 SetPackageSizeNil sets the value for PackageSize to be an explicit nil

### UnsetPackageSize
`func (o *StorefrontPrice) UnsetPackageSize()`

UnsetPackageSize ensures that no value is present for PackageSize, not even an explicit nil
### GetTrialPeriodDays

`func (o *StorefrontPrice) GetTrialPeriodDays() int32`

GetTrialPeriodDays returns the TrialPeriodDays field if non-nil, zero value otherwise.

### GetTrialPeriodDaysOk

`func (o *StorefrontPrice) GetTrialPeriodDaysOk() (*int32, bool)`

GetTrialPeriodDaysOk returns a tuple with the TrialPeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPeriodDays

`func (o *StorefrontPrice) SetTrialPeriodDays(v int32)`

SetTrialPeriodDays sets TrialPeriodDays field to given value.


### SetTrialPeriodDaysNil

`func (o *StorefrontPrice) SetTrialPeriodDaysNil(b bool)`

 SetTrialPeriodDaysNil sets the value for TrialPeriodDays to be an explicit nil

### UnsetTrialPeriodDays
`func (o *StorefrontPrice) UnsetTrialPeriodDays()`

UnsetTrialPeriodDays ensures that no value is present for TrialPeriodDays, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


