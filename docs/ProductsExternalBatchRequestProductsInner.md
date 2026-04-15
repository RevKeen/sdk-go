# ProductsExternalBatchRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | External system&#39;s product ID | 
**Name** | **string** | Product name | 
**Description** | Pointer to **string** | Product description | [optional] 
**Kind** | Pointer to **string** | Product kind - subscription or one-time | [optional] 
**AmountCents** | **int32** | Price in minor units (cents) | 
**Currency** | Pointer to **string** | ISO 4217 currency code | [optional] [default to "USD"]
**IsActive** | Pointer to **bool** | Whether the product is active | [optional] [default to true]
**Interval** | Pointer to **string** | Billing interval for subscriptions | [optional] 
**IntervalCount** | Pointer to **int32** | Number of intervals between billings | [optional] 
**ExternalUpdatedAt** | Pointer to **time.Time** | External system&#39;s last update timestamp for stale update protection | [optional] 
**ExternalRef** | Pointer to **string** | Additional external reference (e.g., membership ID) | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional metadata from external system | [optional] 

## Methods

### NewProductsExternalBatchRequestProductsInner

`func NewProductsExternalBatchRequestProductsInner(externalId string, name string, amountCents int32, ) *ProductsExternalBatchRequestProductsInner`

NewProductsExternalBatchRequestProductsInner instantiates a new ProductsExternalBatchRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsExternalBatchRequestProductsInnerWithDefaults

`func NewProductsExternalBatchRequestProductsInnerWithDefaults() *ProductsExternalBatchRequestProductsInner`

NewProductsExternalBatchRequestProductsInnerWithDefaults instantiates a new ProductsExternalBatchRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProductsExternalBatchRequestProductsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *ProductsExternalBatchRequestProductsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductsExternalBatchRequestProductsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductsExternalBatchRequestProductsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ProductsExternalBatchRequestProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductsExternalBatchRequestProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductsExternalBatchRequestProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductsExternalBatchRequestProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *ProductsExternalBatchRequestProductsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProductsExternalBatchRequestProductsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProductsExternalBatchRequestProductsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProductsExternalBatchRequestProductsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetAmountCents

`func (o *ProductsExternalBatchRequestProductsInner) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *ProductsExternalBatchRequestProductsInner) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *ProductsExternalBatchRequestProductsInner) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCurrency

`func (o *ProductsExternalBatchRequestProductsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductsExternalBatchRequestProductsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductsExternalBatchRequestProductsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductsExternalBatchRequestProductsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIsActive

`func (o *ProductsExternalBatchRequestProductsInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ProductsExternalBatchRequestProductsInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ProductsExternalBatchRequestProductsInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ProductsExternalBatchRequestProductsInner) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetInterval

`func (o *ProductsExternalBatchRequestProductsInner) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ProductsExternalBatchRequestProductsInner) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ProductsExternalBatchRequestProductsInner) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ProductsExternalBatchRequestProductsInner) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetIntervalCount

`func (o *ProductsExternalBatchRequestProductsInner) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *ProductsExternalBatchRequestProductsInner) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *ProductsExternalBatchRequestProductsInner) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *ProductsExternalBatchRequestProductsInner) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetExternalUpdatedAt

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalUpdatedAt() time.Time`

GetExternalUpdatedAt returns the ExternalUpdatedAt field if non-nil, zero value otherwise.

### GetExternalUpdatedAtOk

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalUpdatedAtOk() (*time.Time, bool)`

GetExternalUpdatedAtOk returns a tuple with the ExternalUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUpdatedAt

`func (o *ProductsExternalBatchRequestProductsInner) SetExternalUpdatedAt(v time.Time)`

SetExternalUpdatedAt sets ExternalUpdatedAt field to given value.

### HasExternalUpdatedAt

`func (o *ProductsExternalBatchRequestProductsInner) HasExternalUpdatedAt() bool`

HasExternalUpdatedAt returns a boolean if a field has been set.

### GetExternalRef

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *ProductsExternalBatchRequestProductsInner) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *ProductsExternalBatchRequestProductsInner) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *ProductsExternalBatchRequestProductsInner) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetMetadata

`func (o *ProductsExternalBatchRequestProductsInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProductsExternalBatchRequestProductsInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProductsExternalBatchRequestProductsInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProductsExternalBatchRequestProductsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


