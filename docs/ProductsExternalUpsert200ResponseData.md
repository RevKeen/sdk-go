# ProductsExternalUpsert200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Kind** | **NullableString** |  | 
**AmountCents** | **NullableFloat32** |  | 
**Currency** | **NullableString** |  | 
**IsActive** | **NullableBool** |  | 
**ExternalId** | **NullableString** |  | 
**ExternalSystem** | **NullableString** |  | 
**ExternalRef** | **NullableString** |  | 
**Interval** | **NullableString** |  | 
**IntervalCount** | **NullableFloat32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewProductsExternalUpsert200ResponseData

`func NewProductsExternalUpsert200ResponseData(id string, merchantId string, name string, description NullableString, kind NullableString, amountCents NullableFloat32, currency NullableString, isActive NullableBool, externalId NullableString, externalSystem NullableString, externalRef NullableString, interval NullableString, intervalCount NullableFloat32, createdAt time.Time, updatedAt time.Time, ) *ProductsExternalUpsert200ResponseData`

NewProductsExternalUpsert200ResponseData instantiates a new ProductsExternalUpsert200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsExternalUpsert200ResponseDataWithDefaults

`func NewProductsExternalUpsert200ResponseDataWithDefaults() *ProductsExternalUpsert200ResponseData`

NewProductsExternalUpsert200ResponseDataWithDefaults instantiates a new ProductsExternalUpsert200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductsExternalUpsert200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductsExternalUpsert200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductsExternalUpsert200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *ProductsExternalUpsert200ResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ProductsExternalUpsert200ResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ProductsExternalUpsert200ResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetName

`func (o *ProductsExternalUpsert200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsExternalUpsert200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsExternalUpsert200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductsExternalUpsert200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductsExternalUpsert200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductsExternalUpsert200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ProductsExternalUpsert200ResponseData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ProductsExternalUpsert200ResponseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *ProductsExternalUpsert200ResponseData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProductsExternalUpsert200ResponseData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProductsExternalUpsert200ResponseData) SetKind(v string)`

SetKind sets Kind field to given value.


### SetKindNil

`func (o *ProductsExternalUpsert200ResponseData) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *ProductsExternalUpsert200ResponseData) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetAmountCents

`func (o *ProductsExternalUpsert200ResponseData) GetAmountCents() float32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *ProductsExternalUpsert200ResponseData) GetAmountCentsOk() (*float32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *ProductsExternalUpsert200ResponseData) SetAmountCents(v float32)`

SetAmountCents sets AmountCents field to given value.


### SetAmountCentsNil

`func (o *ProductsExternalUpsert200ResponseData) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *ProductsExternalUpsert200ResponseData) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetCurrency

`func (o *ProductsExternalUpsert200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductsExternalUpsert200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductsExternalUpsert200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *ProductsExternalUpsert200ResponseData) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ProductsExternalUpsert200ResponseData) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetIsActive

`func (o *ProductsExternalUpsert200ResponseData) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductsExternalUpsert200ResponseData) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductsExternalUpsert200ResponseData) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### SetIsActiveNil

`func (o *ProductsExternalUpsert200ResponseData) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *ProductsExternalUpsert200ResponseData) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetExternalId

`func (o *ProductsExternalUpsert200ResponseData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProductsExternalUpsert200ResponseData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProductsExternalUpsert200ResponseData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *ProductsExternalUpsert200ResponseData) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ProductsExternalUpsert200ResponseData) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalSystem

`func (o *ProductsExternalUpsert200ResponseData) GetExternalSystem() string`

GetExternalSystem returns the ExternalSystem field if non-nil, zero value otherwise.

### GetExternalSystemOk

`func (o *ProductsExternalUpsert200ResponseData) GetExternalSystemOk() (*string, bool)`

GetExternalSystemOk returns a tuple with the ExternalSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSystem

`func (o *ProductsExternalUpsert200ResponseData) SetExternalSystem(v string)`

SetExternalSystem sets ExternalSystem field to given value.


### SetExternalSystemNil

`func (o *ProductsExternalUpsert200ResponseData) SetExternalSystemNil(b bool)`

 SetExternalSystemNil sets the value for ExternalSystem to be an explicit nil

### UnsetExternalSystem
`func (o *ProductsExternalUpsert200ResponseData) UnsetExternalSystem()`

UnsetExternalSystem ensures that no value is present for ExternalSystem, not even an explicit nil
### GetExternalRef

`func (o *ProductsExternalUpsert200ResponseData) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *ProductsExternalUpsert200ResponseData) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *ProductsExternalUpsert200ResponseData) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.


### SetExternalRefNil

`func (o *ProductsExternalUpsert200ResponseData) SetExternalRefNil(b bool)`

 SetExternalRefNil sets the value for ExternalRef to be an explicit nil

### UnsetExternalRef
`func (o *ProductsExternalUpsert200ResponseData) UnsetExternalRef()`

UnsetExternalRef ensures that no value is present for ExternalRef, not even an explicit nil
### GetInterval

`func (o *ProductsExternalUpsert200ResponseData) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ProductsExternalUpsert200ResponseData) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ProductsExternalUpsert200ResponseData) SetInterval(v string)`

SetInterval sets Interval field to given value.


### SetIntervalNil

`func (o *ProductsExternalUpsert200ResponseData) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *ProductsExternalUpsert200ResponseData) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *ProductsExternalUpsert200ResponseData) GetIntervalCount() float32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *ProductsExternalUpsert200ResponseData) GetIntervalCountOk() (*float32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *ProductsExternalUpsert200ResponseData) SetIntervalCount(v float32)`

SetIntervalCount sets IntervalCount field to given value.


### SetIntervalCountNil

`func (o *ProductsExternalUpsert200ResponseData) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *ProductsExternalUpsert200ResponseData) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetCreatedAt

`func (o *ProductsExternalUpsert200ResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductsExternalUpsert200ResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductsExternalUpsert200ResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProductsExternalUpsert200ResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProductsExternalUpsert200ResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProductsExternalUpsert200ResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


