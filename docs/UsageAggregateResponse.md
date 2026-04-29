# UsageAggregateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**MeterId** | **string** |  | 
**StartTime** | **string** |  | 
**EndTime** | **string** |  | 
**Value** | **float32** | Aggregated usage value | 
**EventCount** | **float32** | Number of events in range | 

## Methods

### NewUsageAggregateResponse

`func NewUsageAggregateResponse(object string, meterId string, startTime string, endTime string, value float32, eventCount float32, ) *UsageAggregateResponse`

NewUsageAggregateResponse instantiates a new UsageAggregateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAggregateResponseWithDefaults

`func NewUsageAggregateResponseWithDefaults() *UsageAggregateResponse`

NewUsageAggregateResponseWithDefaults instantiates a new UsageAggregateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *UsageAggregateResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *UsageAggregateResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *UsageAggregateResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetMeterId

`func (o *UsageAggregateResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UsageAggregateResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UsageAggregateResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetStartTime

`func (o *UsageAggregateResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UsageAggregateResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UsageAggregateResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *UsageAggregateResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UsageAggregateResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UsageAggregateResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetValue

`func (o *UsageAggregateResponse) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageAggregateResponse) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageAggregateResponse) SetValue(v float32)`

SetValue sets Value field to given value.


### GetEventCount

`func (o *UsageAggregateResponse) GetEventCount() float32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *UsageAggregateResponse) GetEventCountOk() (*float32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *UsageAggregateResponse) SetEventCount(v float32)`

SetEventCount sets EventCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


