# FulfillOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | Pointer to **string** | Shipping tracking number | [optional]
**Carrier** | Pointer to **string** | Shipping carrier (e.g., &#39;ups&#39;, &#39;fedex&#39;, &#39;usps&#39;) | [optional]
**LineItemIds** | Pointer to **[]string** | Specific line items to fulfill, each in full (defaults to every unfulfilled line) | [optional]
**TrackingUrl** | Pointer to **string** | Carrier tracking link | [optional]
**Status** | Pointer to **string** | Mark the lines shipped (default) or delivered | [optional] [default to "shipped"]
**NotifyCustomer** | Pointer to **bool** | Send fulfillment notification to customer | [optional] [default to true]

## Methods

### NewFulfillOrderRequest

`func NewFulfillOrderRequest() *FulfillOrderRequest`

NewFulfillOrderRequest instantiates a new FulfillOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillOrderRequestWithDefaults

`func NewFulfillOrderRequestWithDefaults() *FulfillOrderRequest`

NewFulfillOrderRequestWithDefaults instantiates a new FulfillOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *FulfillOrderRequest) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *FulfillOrderRequest) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *FulfillOrderRequest) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *FulfillOrderRequest) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetCarrier

`func (o *FulfillOrderRequest) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *FulfillOrderRequest) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *FulfillOrderRequest) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *FulfillOrderRequest) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetLineItemIds

`func (o *FulfillOrderRequest) GetLineItemIds() []string`

GetLineItemIds returns the LineItemIds field if non-nil, zero value otherwise.

### GetLineItemIdsOk

`func (o *FulfillOrderRequest) GetLineItemIdsOk() (*[]string, bool)`

GetLineItemIdsOk returns a tuple with the LineItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemIds

`func (o *FulfillOrderRequest) SetLineItemIds(v []string)`

SetLineItemIds sets LineItemIds field to given value.

### HasLineItemIds

`func (o *FulfillOrderRequest) HasLineItemIds() bool`

HasLineItemIds returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *FulfillOrderRequest) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *FulfillOrderRequest) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *FulfillOrderRequest) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *FulfillOrderRequest) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetStatus

`func (o *FulfillOrderRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FulfillOrderRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FulfillOrderRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FulfillOrderRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNotifyCustomer

`func (o *FulfillOrderRequest) GetNotifyCustomer() bool`

GetNotifyCustomer returns the NotifyCustomer field if non-nil, zero value otherwise.

### GetNotifyCustomerOk

`func (o *FulfillOrderRequest) GetNotifyCustomerOk() (*bool, bool)`

GetNotifyCustomerOk returns a tuple with the NotifyCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCustomer

`func (o *FulfillOrderRequest) SetNotifyCustomer(v bool)`

SetNotifyCustomer sets NotifyCustomer field to given value.

### HasNotifyCustomer

`func (o *FulfillOrderRequest) HasNotifyCustomer() bool`

HasNotifyCustomer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
