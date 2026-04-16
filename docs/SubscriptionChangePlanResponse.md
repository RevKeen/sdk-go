# SubscriptionChangePlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**SubscriptionCreateResponseData**](SubscriptionCreateResponseData.md) |  | 
**Proration** | Pointer to **map[string]interface{}** | Proration details | [optional] 
**PreviousPlanId** | Pointer to **string** |  | [optional] 
**InvoiceCreated** | **bool** |  | 

## Methods

### NewSubscriptionChangePlanResponse

`func NewSubscriptionChangePlanResponse(subscription SubscriptionCreateResponseData, invoiceCreated bool, ) *SubscriptionChangePlanResponse`

NewSubscriptionChangePlanResponse instantiates a new SubscriptionChangePlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionChangePlanResponseWithDefaults

`func NewSubscriptionChangePlanResponseWithDefaults() *SubscriptionChangePlanResponse`

NewSubscriptionChangePlanResponseWithDefaults instantiates a new SubscriptionChangePlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *SubscriptionChangePlanResponse) GetSubscription() SubscriptionCreateResponseData`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionChangePlanResponse) GetSubscriptionOk() (*SubscriptionCreateResponseData, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionChangePlanResponse) SetSubscription(v SubscriptionCreateResponseData)`

SetSubscription sets Subscription field to given value.


### GetProration

`func (o *SubscriptionChangePlanResponse) GetProration() map[string]interface{}`

GetProration returns the Proration field if non-nil, zero value otherwise.

### GetProrationOk

`func (o *SubscriptionChangePlanResponse) GetProrationOk() (*map[string]interface{}, bool)`

GetProrationOk returns a tuple with the Proration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProration

`func (o *SubscriptionChangePlanResponse) SetProration(v map[string]interface{})`

SetProration sets Proration field to given value.

### HasProration

`func (o *SubscriptionChangePlanResponse) HasProration() bool`

HasProration returns a boolean if a field has been set.

### GetPreviousPlanId

`func (o *SubscriptionChangePlanResponse) GetPreviousPlanId() string`

GetPreviousPlanId returns the PreviousPlanId field if non-nil, zero value otherwise.

### GetPreviousPlanIdOk

`func (o *SubscriptionChangePlanResponse) GetPreviousPlanIdOk() (*string, bool)`

GetPreviousPlanIdOk returns a tuple with the PreviousPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPlanId

`func (o *SubscriptionChangePlanResponse) SetPreviousPlanId(v string)`

SetPreviousPlanId sets PreviousPlanId field to given value.

### HasPreviousPlanId

`func (o *SubscriptionChangePlanResponse) HasPreviousPlanId() bool`

HasPreviousPlanId returns a boolean if a field has been set.

### GetInvoiceCreated

`func (o *SubscriptionChangePlanResponse) GetInvoiceCreated() bool`

GetInvoiceCreated returns the InvoiceCreated field if non-nil, zero value otherwise.

### GetInvoiceCreatedOk

`func (o *SubscriptionChangePlanResponse) GetInvoiceCreatedOk() (*bool, bool)`

GetInvoiceCreatedOk returns a tuple with the InvoiceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCreated

`func (o *SubscriptionChangePlanResponse) SetInvoiceCreated(v bool)`

SetInvoiceCreated sets InvoiceCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


