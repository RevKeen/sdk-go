# CancelSubscriptionScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNow** | Pointer to **bool** | Generate a final invoice immediately | [optional] 
**Prorate** | Pointer to **bool** | Prorate final invoice | [optional] 

## Methods

### NewCancelSubscriptionScheduleRequest

`func NewCancelSubscriptionScheduleRequest() *CancelSubscriptionScheduleRequest`

NewCancelSubscriptionScheduleRequest instantiates a new CancelSubscriptionScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionScheduleRequestWithDefaults

`func NewCancelSubscriptionScheduleRequestWithDefaults() *CancelSubscriptionScheduleRequest`

NewCancelSubscriptionScheduleRequestWithDefaults instantiates a new CancelSubscriptionScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNow

`func (o *CancelSubscriptionScheduleRequest) GetInvoiceNow() bool`

GetInvoiceNow returns the InvoiceNow field if non-nil, zero value otherwise.

### GetInvoiceNowOk

`func (o *CancelSubscriptionScheduleRequest) GetInvoiceNowOk() (*bool, bool)`

GetInvoiceNowOk returns a tuple with the InvoiceNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNow

`func (o *CancelSubscriptionScheduleRequest) SetInvoiceNow(v bool)`

SetInvoiceNow sets InvoiceNow field to given value.

### HasInvoiceNow

`func (o *CancelSubscriptionScheduleRequest) HasInvoiceNow() bool`

HasInvoiceNow returns a boolean if a field has been set.

### GetProrate

`func (o *CancelSubscriptionScheduleRequest) GetProrate() bool`

GetProrate returns the Prorate field if non-nil, zero value otherwise.

### GetProrateOk

`func (o *CancelSubscriptionScheduleRequest) GetProrateOk() (*bool, bool)`

GetProrateOk returns a tuple with the Prorate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrate

`func (o *CancelSubscriptionScheduleRequest) SetProrate(v bool)`

SetProrate sets Prorate field to given value.

### HasProrate

`func (o *CancelSubscriptionScheduleRequest) HasProrate() bool`

HasProrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


