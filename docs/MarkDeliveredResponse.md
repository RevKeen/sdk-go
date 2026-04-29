# MarkDeliveredResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**MarkedDelivered** | **int32** |  | 
**ReminderEligibleNow** | **int32** | Count of invoices that would match the default overdue ladder (day 1+) at the next hourly reminder cron. With anchor &#x3D; NOW() on mark, this is typically 0 — the first reminder fires tomorrow. | 
**ReminderEligibleWithin7Days** | **int32** | Count of invoices that will receive at least one reminder within 7 days of this attestation, assuming the default ladder [1, 7, 14, 30]. | 
**Skipped** | [**[]MarkDeliveredResponseSkippedInner**](MarkDeliveredResponseSkippedInner.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewMarkDeliveredResponse

`func NewMarkDeliveredResponse(success bool, markedDelivered int32, reminderEligibleNow int32, reminderEligibleWithin7Days int32, skipped []MarkDeliveredResponseSkippedInner, requestId string, ) *MarkDeliveredResponse`

NewMarkDeliveredResponse instantiates a new MarkDeliveredResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkDeliveredResponseWithDefaults

`func NewMarkDeliveredResponseWithDefaults() *MarkDeliveredResponse`

NewMarkDeliveredResponseWithDefaults instantiates a new MarkDeliveredResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MarkDeliveredResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MarkDeliveredResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MarkDeliveredResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetMarkedDelivered

`func (o *MarkDeliveredResponse) GetMarkedDelivered() int32`

GetMarkedDelivered returns the MarkedDelivered field if non-nil, zero value otherwise.

### GetMarkedDeliveredOk

`func (o *MarkDeliveredResponse) GetMarkedDeliveredOk() (*int32, bool)`

GetMarkedDeliveredOk returns a tuple with the MarkedDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkedDelivered

`func (o *MarkDeliveredResponse) SetMarkedDelivered(v int32)`

SetMarkedDelivered sets MarkedDelivered field to given value.


### GetReminderEligibleNow

`func (o *MarkDeliveredResponse) GetReminderEligibleNow() int32`

GetReminderEligibleNow returns the ReminderEligibleNow field if non-nil, zero value otherwise.

### GetReminderEligibleNowOk

`func (o *MarkDeliveredResponse) GetReminderEligibleNowOk() (*int32, bool)`

GetReminderEligibleNowOk returns a tuple with the ReminderEligibleNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderEligibleNow

`func (o *MarkDeliveredResponse) SetReminderEligibleNow(v int32)`

SetReminderEligibleNow sets ReminderEligibleNow field to given value.


### GetReminderEligibleWithin7Days

`func (o *MarkDeliveredResponse) GetReminderEligibleWithin7Days() int32`

GetReminderEligibleWithin7Days returns the ReminderEligibleWithin7Days field if non-nil, zero value otherwise.

### GetReminderEligibleWithin7DaysOk

`func (o *MarkDeliveredResponse) GetReminderEligibleWithin7DaysOk() (*int32, bool)`

GetReminderEligibleWithin7DaysOk returns a tuple with the ReminderEligibleWithin7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminderEligibleWithin7Days

`func (o *MarkDeliveredResponse) SetReminderEligibleWithin7Days(v int32)`

SetReminderEligibleWithin7Days sets ReminderEligibleWithin7Days field to given value.


### GetSkipped

`func (o *MarkDeliveredResponse) GetSkipped() []MarkDeliveredResponseSkippedInner`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *MarkDeliveredResponse) GetSkippedOk() (*[]MarkDeliveredResponseSkippedInner, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *MarkDeliveredResponse) SetSkipped(v []MarkDeliveredResponseSkippedInner)`

SetSkipped sets Skipped field to given value.


### GetRequestId

`func (o *MarkDeliveredResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *MarkDeliveredResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *MarkDeliveredResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


