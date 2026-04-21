# CreateOrderRequestLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Product ID (optional for ad-hoc items) | [optional] 
**Description** | **string** |  | 
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**UnitPrice** | **int32** | Unit price in cents | 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewCreateOrderRequestLineItemsInner

`func NewCreateOrderRequestLineItemsInner(description string, unitPrice int32, ) *CreateOrderRequestLineItemsInner`

NewCreateOrderRequestLineItemsInner instantiates a new CreateOrderRequestLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestLineItemsInnerWithDefaults

`func NewCreateOrderRequestLineItemsInnerWithDefaults() *CreateOrderRequestLineItemsInner`

NewCreateOrderRequestLineItemsInnerWithDefaults instantiates a new CreateOrderRequestLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateOrderRequestLineItemsInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateOrderRequestLineItemsInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateOrderRequestLineItemsInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreateOrderRequestLineItemsInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrderRequestLineItemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrderRequestLineItemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrderRequestLineItemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *CreateOrderRequestLineItemsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateOrderRequestLineItemsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateOrderRequestLineItemsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateOrderRequestLineItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *CreateOrderRequestLineItemsInner) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CreateOrderRequestLineItemsInner) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CreateOrderRequestLineItemsInner) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetMetadata

`func (o *CreateOrderRequestLineItemsInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateOrderRequestLineItemsInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateOrderRequestLineItemsInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateOrderRequestLineItemsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


