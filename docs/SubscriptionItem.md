# SubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubscriptionId** | **string** |  | 
**ProductId** | **NullableString** |  | 
**PriceId** | **NullableString** |  | 
**Description** | **NullableString** |  | 
**Quantity** | **int32** | Item quantity | 
**Currency** | **string** |  | 
**UnitAmountMinor** | **int32** | Unit price in cents | 
**FulfillmentType** | **string** | none: service (no order), physical: creates shipping order, digital: creates download order | 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewSubscriptionItem

`func NewSubscriptionItem(id string, subscriptionId string, productId NullableString, priceId NullableString, description NullableString, quantity int32, currency string, unitAmountMinor int32, fulfillmentType string, createdAt string, updatedAt string, ) *SubscriptionItem`

NewSubscriptionItem instantiates a new SubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionItemWithDefaults

`func NewSubscriptionItemWithDefaults() *SubscriptionItem`

NewSubscriptionItemWithDefaults instantiates a new SubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionItem) SetId(v string)`

SetId sets Id field to given value.


### GetSubscriptionId

`func (o *SubscriptionItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetProductId

`func (o *SubscriptionItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### SetProductIdNil

`func (o *SubscriptionItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *SubscriptionItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetPriceId

`func (o *SubscriptionItem) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionItem) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionItem) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### SetPriceIdNil

`func (o *SubscriptionItem) SetPriceIdNil(b bool)`

 SetPriceIdNil sets the value for PriceId to be an explicit nil

### UnsetPriceId
`func (o *SubscriptionItem) UnsetPriceId()`

UnsetPriceId ensures that no value is present for PriceId, not even an explicit nil
### GetDescription

`func (o *SubscriptionItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SubscriptionItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubscriptionItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *SubscriptionItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetCurrency

`func (o *SubscriptionItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitAmountMinor

`func (o *SubscriptionItem) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *SubscriptionItem) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *SubscriptionItem) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetFulfillmentType

`func (o *SubscriptionItem) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *SubscriptionItem) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *SubscriptionItem) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.


### GetMetadata

`func (o *SubscriptionItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


