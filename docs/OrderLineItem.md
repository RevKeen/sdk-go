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
**FulfillmentType** | Pointer to **string** | How this line is fulfilled | [optional]
**FulfillmentStatus** | Pointer to **string** | Fulfilment state of this line; processing means partly shipped | [optional]
**FulfilledQuantity** | Pointer to **int32** | Units of this line shipped so far | [optional]
**TrackingNumber** | Pointer to **NullableString** |  | [optional]
**TrackingUrl** | Pointer to **NullableString** |  | [optional]
**Carrier** | Pointer to **NullableString** |  | [optional]
**ShippedAt** | Pointer to **NullableTime** |  | [optional]
**DeliveredAt** | Pointer to **NullableTime** |  | [optional]

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

### GetFulfillmentType

`func (o *OrderLineItem) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *OrderLineItem) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *OrderLineItem) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *OrderLineItem) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *OrderLineItem) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *OrderLineItem) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *OrderLineItem) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *OrderLineItem) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetFulfilledQuantity

`func (o *OrderLineItem) GetFulfilledQuantity() int32`

GetFulfilledQuantity returns the FulfilledQuantity field if non-nil, zero value otherwise.

### GetFulfilledQuantityOk

`func (o *OrderLineItem) GetFulfilledQuantityOk() (*int32, bool)`

GetFulfilledQuantityOk returns a tuple with the FulfilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledQuantity

`func (o *OrderLineItem) SetFulfilledQuantity(v int32)`

SetFulfilledQuantity sets FulfilledQuantity field to given value.

### HasFulfilledQuantity

`func (o *OrderLineItem) HasFulfilledQuantity() bool`

HasFulfilledQuantity returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *OrderLineItem) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrderLineItem) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrderLineItem) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *OrderLineItem) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *OrderLineItem) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *OrderLineItem) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingUrl

`func (o *OrderLineItem) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *OrderLineItem) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *OrderLineItem) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *OrderLineItem) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### SetTrackingUrlNil

`func (o *OrderLineItem) SetTrackingUrlNil(b bool)`

 SetTrackingUrlNil sets the value for TrackingUrl to be an explicit nil

### UnsetTrackingUrl
`func (o *OrderLineItem) UnsetTrackingUrl()`

UnsetTrackingUrl ensures that no value is present for TrackingUrl, not even an explicit nil
### GetCarrier

`func (o *OrderLineItem) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *OrderLineItem) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *OrderLineItem) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *OrderLineItem) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *OrderLineItem) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *OrderLineItem) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetShippedAt

`func (o *OrderLineItem) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *OrderLineItem) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *OrderLineItem) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *OrderLineItem) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *OrderLineItem) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *OrderLineItem) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetDeliveredAt

`func (o *OrderLineItem) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *OrderLineItem) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *OrderLineItem) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *OrderLineItem) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *OrderLineItem) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *OrderLineItem) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
