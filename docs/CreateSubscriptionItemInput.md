# CreateSubscriptionItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **string** | Product UUID to link to this item | [optional] 
**PriceId** | Pointer to **string** | Price UUID to link to this item | [optional] 
**Description** | Pointer to **string** | Item description | [optional] 
**Quantity** | Pointer to **int32** | Item quantity | [optional] [default to 1]
**Currency** | **string** | ISO 4217 currency code | 
**UnitAmountMinor** | **int32** | Unit price in cents | 
**FulfillmentType** | Pointer to **string** | Determines if orders are created on renewal | [optional] [default to "none"]
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewCreateSubscriptionItemInput

`func NewCreateSubscriptionItemInput(currency string, unitAmountMinor int32, ) *CreateSubscriptionItemInput`

NewCreateSubscriptionItemInput instantiates a new CreateSubscriptionItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionItemInputWithDefaults

`func NewCreateSubscriptionItemInputWithDefaults() *CreateSubscriptionItemInput`

NewCreateSubscriptionItemInputWithDefaults instantiates a new CreateSubscriptionItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateSubscriptionItemInput) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateSubscriptionItemInput) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateSubscriptionItemInput) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CreateSubscriptionItemInput) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPriceId

`func (o *CreateSubscriptionItemInput) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *CreateSubscriptionItemInput) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *CreateSubscriptionItemInput) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *CreateSubscriptionItemInput) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSubscriptionItemInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSubscriptionItemInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSubscriptionItemInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSubscriptionItemInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *CreateSubscriptionItemInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateSubscriptionItemInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateSubscriptionItemInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateSubscriptionItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateSubscriptionItemInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateSubscriptionItemInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateSubscriptionItemInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUnitAmountMinor

`func (o *CreateSubscriptionItemInput) GetUnitAmountMinor() int32`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *CreateSubscriptionItemInput) GetUnitAmountMinorOk() (*int32, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *CreateSubscriptionItemInput) SetUnitAmountMinor(v int32)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetFulfillmentType

`func (o *CreateSubscriptionItemInput) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *CreateSubscriptionItemInput) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *CreateSubscriptionItemInput) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *CreateSubscriptionItemInput) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSubscriptionItemInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSubscriptionItemInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSubscriptionItemInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSubscriptionItemInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


