# OrderLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ProductId** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitPrice** | **int32** | Unit price in cents | 
**Subtotal** | **int32** | Subtotal in cents (unit_price * quantity) | 
**Discount** | Pointer to **int32** | Discount in cents | [optional] [default to 0]
**Tax** | Pointer to **int32** | Tax in cents | [optional] [default to 0]
**Total** | **int32** | Total in cents (subtotal - discount + tax) | 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**Position** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewOrderLineItem

`func NewOrderLineItem(id string, description string, quantity int32, unitPrice int32, subtotal int32, total int32, ) *OrderLineItem`

NewOrderLineItem instantiates a new OrderLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLineItemWithDefaults

`func NewOrderLineItemWithDefaults() *OrderLineItem`

NewOrderLineItemWithDefaults instantiates a new OrderLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderLineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderLineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderLineItem) SetId(v string)`

SetId sets Id field to given value.


### GetProductId

`func (o *OrderLineItem) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *OrderLineItem) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *OrderLineItem) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *OrderLineItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *OrderLineItem) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *OrderLineItem) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetDescription

`func (o *OrderLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetQuantity

`func (o *OrderLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitPrice

`func (o *OrderLineItem) GetUnitPrice() int32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderLineItem) GetUnitPriceOk() (*int32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderLineItem) SetUnitPrice(v int32)`

SetUnitPrice sets UnitPrice field to given value.


### GetSubtotal

`func (o *OrderLineItem) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *OrderLineItem) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *OrderLineItem) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetDiscount

`func (o *OrderLineItem) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderLineItem) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderLineItem) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderLineItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetTax

`func (o *OrderLineItem) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *OrderLineItem) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *OrderLineItem) SetTax(v int32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *OrderLineItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTotal

`func (o *OrderLineItem) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrderLineItem) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrderLineItem) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMetadata

`func (o *OrderLineItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderLineItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderLineItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderLineItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPosition

`func (o *OrderLineItem) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *OrderLineItem) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *OrderLineItem) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *OrderLineItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


