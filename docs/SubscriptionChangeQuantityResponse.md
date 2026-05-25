# SubscriptionChangeQuantityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**SubscriptionCreateResponseData**](SubscriptionCreateResponseData.md) |  | 
**Proration** | Pointer to **map[string]interface{}** | Proration details | [optional] 
**PreviousQuantity** | Pointer to **float32** |  | [optional] 
**InvoiceCreated** | **bool** |  | 

## Methods

### NewSubscriptionChangeQuantityResponse

`func NewSubscriptionChangeQuantityResponse(subscription SubscriptionCreateResponseData, invoiceCreated bool, ) *SubscriptionChangeQuantityResponse`

NewSubscriptionChangeQuantityResponse instantiates a new SubscriptionChangeQuantityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionChangeQuantityResponseWithDefaults

`func NewSubscriptionChangeQuantityResponseWithDefaults() *SubscriptionChangeQuantityResponse`

NewSubscriptionChangeQuantityResponseWithDefaults instantiates a new SubscriptionChangeQuantityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *SubscriptionChangeQuantityResponse) GetSubscription() SubscriptionCreateResponseData`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionChangeQuantityResponse) GetSubscriptionOk() (*SubscriptionCreateResponseData, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionChangeQuantityResponse) SetSubscription(v SubscriptionCreateResponseData)`

SetSubscription sets Subscription field to given value.


### GetProration

`func (o *SubscriptionChangeQuantityResponse) GetProration() map[string]interface{}`

GetProration returns the Proration field if non-nil, zero value otherwise.

### GetProrationOk

`func (o *SubscriptionChangeQuantityResponse) GetProrationOk() (*map[string]interface{}, bool)`

GetProrationOk returns a tuple with the Proration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProration

`func (o *SubscriptionChangeQuantityResponse) SetProration(v map[string]interface{})`

SetProration sets Proration field to given value.

### HasProration

`func (o *SubscriptionChangeQuantityResponse) HasProration() bool`

HasProration returns a boolean if a field has been set.

### GetPreviousQuantity

`func (o *SubscriptionChangeQuantityResponse) GetPreviousQuantity() float32`

GetPreviousQuantity returns the PreviousQuantity field if non-nil, zero value otherwise.

### GetPreviousQuantityOk

`func (o *SubscriptionChangeQuantityResponse) GetPreviousQuantityOk() (*float32, bool)`

GetPreviousQuantityOk returns a tuple with the PreviousQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousQuantity

`func (o *SubscriptionChangeQuantityResponse) SetPreviousQuantity(v float32)`

SetPreviousQuantity sets PreviousQuantity field to given value.

### HasPreviousQuantity

`func (o *SubscriptionChangeQuantityResponse) HasPreviousQuantity() bool`

HasPreviousQuantity returns a boolean if a field has been set.

### GetInvoiceCreated

`func (o *SubscriptionChangeQuantityResponse) GetInvoiceCreated() bool`

GetInvoiceCreated returns the InvoiceCreated field if non-nil, zero value otherwise.

### GetInvoiceCreatedOk

`func (o *SubscriptionChangeQuantityResponse) GetInvoiceCreatedOk() (*bool, bool)`

GetInvoiceCreatedOk returns a tuple with the InvoiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCreated

`func (o *SubscriptionChangeQuantityResponse) SetInvoiceCreated(v bool)`

SetInvoiceCreated sets InvoiceCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


