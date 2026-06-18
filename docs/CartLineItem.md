# CartLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProductId** | **string** |  | 
**Name** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitPriceMinor** | **int32** |  | 
**Currency** | **string** |  | 
**Recurring** | [**CartLineItemRecurring**](CartLineItemRecurring.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartLineItem

`func NewCartLineItem(id string, productId string, name string, quantity int32, unitPriceMinor int32, currency string, recurring CartLineItemRecurring, ) *CartLineItem`

NewCartLineItem instantiates a new CartLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartLineItemWithDefaults

`func NewCartLineItemWithDefaults() *CartLineItem`

NewCartLineItemWithDefaults instantiates a new CartLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CartLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartLineItem) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *CartLineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CartLineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CartLineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *CartLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartLineItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *CartLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPriceMinor

`func (o *CartLineItem) GetUnitPriceMinor() int32`

GetUnitPriceMinor returns the UnitPriceMinor field if non-nil, zero value otherwise.

### GetUnitPriceMinorOk

`func (o *CartLineItem) GetUnitPriceMinorOk() (*int32, bool)`

GetUnitPriceMinorOk returns a tuple with the UnitPriceMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceMinor

`func (o *CartLineItem) SetUnitPriceMinor(v int32)`

SetUnitPriceMinor sets UnitPriceMinor field to given value.


### GetCurrency

`func (o *CartLineItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartLineItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartLineItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRecurring

`func (o *CartLineItem) GetRecurring() CartLineItemRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CartLineItem) GetRecurringOk() (*CartLineItemRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CartLineItem) SetRecurring(v CartLineItemRecurring)`

SetRecurring sets Recurring field to given value.


### GetMetadata

`func (o *CartLineItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CartLineItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CartLineItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CartLineItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


