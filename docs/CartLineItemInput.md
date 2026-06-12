# CartLineItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  | 
**Name** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitPriceMinor** | **int32** | Unit price in minor units (pence, cents). | 
**Currency** | **string** |  | 
**Recurring** | Pointer to [**CartLineItemRecurring**](CartLineItemRecurring.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCartLineItemInput

`func NewCartLineItemInput(productId string, name string, quantity int32, unitPriceMinor int32, currency string, ) *CartLineItemInput`

NewCartLineItemInput instantiates a new CartLineItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartLineItemInputWithDefaults

`func NewCartLineItemInputWithDefaults() *CartLineItemInput`

NewCartLineItemInputWithDefaults instantiates a new CartLineItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CartLineItemInput) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CartLineItemInput) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CartLineItemInput) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *CartLineItemInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CartLineItemInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CartLineItemInput) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *CartLineItemInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CartLineItemInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CartLineItemInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPriceMinor

`func (o *CartLineItemInput) GetUnitPriceMinor() int32`

GetUnitPriceMinor returns the UnitPriceMinor field if non-nil, zero value otherwise.

### GetUnitPriceMinorOk

`func (o *CartLineItemInput) GetUnitPriceMinorOk() (*int32, bool)`

GetUnitPriceMinorOk returns a tuple with the UnitPriceMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceMinor

`func (o *CartLineItemInput) SetUnitPriceMinor(v int32)`

SetUnitPriceMinor sets UnitPriceMinor field to given value.


### GetCurrency

`func (o *CartLineItemInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CartLineItemInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CartLineItemInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRecurring

`func (o *CartLineItemInput) GetRecurring() CartLineItemRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CartLineItemInput) GetRecurringOk() (*CartLineItemRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CartLineItemInput) SetRecurring(v CartLineItemRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *CartLineItemInput) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetMetadata

`func (o *CartLineItemInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CartLineItemInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CartLineItemInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CartLineItemInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


