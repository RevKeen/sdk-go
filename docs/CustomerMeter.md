# CustomerMeter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**CustomerId** | **string** |  | 
**MeterId** | **string** |  | 
**MeterName** | **string** |  | 
**MeterEventName** | **string** | The &#x60;event_name&#x60; the merchant&#39;s code posts to /v2/usage-events for this meter. | 
**Aggregation** | **string** | How quantities are aggregated. Typical values: &#x60;sum&#x60;, &#x60;count&#x60;, &#x60;max&#x60;, &#x60;unique&#x60;. | 
**UnitName** | **NullableString** | Human-readable unit for display (e.g. &#x60;tokens&#x60;, &#x60;requests&#x60;, &#x60;GB&#x60;). Merchant-configurable. | 
**TotalQuantity** | **float32** | Aggregate across all usage events for this (customer, meter) pair. The aggregation function applied matches the meter&#39;s &#x60;aggregation&#x60; field. | 
**EventCount** | **int32** | Number of &#x60;usage_events&#x60; rows that contributed to the aggregate. Diagnostic — not the aggregate itself. | 
**LastEventAt** | **NullableTime** | Timestamp of the most recent usage event. Null if no events have been recorded. | 

## Methods

### NewCustomerMeter

`func NewCustomerMeter(object string, customerId string, meterId string, meterName string, meterEventName string, aggregation string, unitName NullableString, totalQuantity float32, eventCount int32, lastEventAt NullableTime, ) *CustomerMeter`

NewCustomerMeter instantiates a new CustomerMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMeterWithDefaults

`func NewCustomerMeterWithDefaults() *CustomerMeter`

NewCustomerMeterWithDefaults instantiates a new CustomerMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CustomerMeter) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerMeter) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerMeter) SetObject(v string)`

SetObject sets Object field to given value.


### GetCustomerId

`func (o *CustomerMeter) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerMeter) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerMeter) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetMeterId

`func (o *CustomerMeter) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *CustomerMeter) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *CustomerMeter) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetMeterName

`func (o *CustomerMeter) GetMeterName() string`

GetMeterName returns the MeterName field if non-nil, zero value otherwise.

### GetMeterNameOk

`func (o *CustomerMeter) GetMeterNameOk() (*string, bool)`

GetMeterNameOk returns a tuple with the MeterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterName

`func (o *CustomerMeter) SetMeterName(v string)`

SetMeterName sets MeterName field to given value.


### GetMeterEventName

`func (o *CustomerMeter) GetMeterEventName() string`

GetMeterEventName returns the MeterEventName field if non-nil, zero value otherwise.

### GetMeterEventNameOk

`func (o *CustomerMeter) GetMeterEventNameOk() (*string, bool)`

GetMeterEventNameOk returns a tuple with the MeterEventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterEventName

`func (o *CustomerMeter) SetMeterEventName(v string)`

SetMeterEventName sets MeterEventName field to given value.


### GetAggregation

`func (o *CustomerMeter) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *CustomerMeter) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *CustomerMeter) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.


### GetUnitName

`func (o *CustomerMeter) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *CustomerMeter) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *CustomerMeter) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### SetUnitNameNil

`func (o *CustomerMeter) SetUnitNameNil(b bool)`

 SetUnitNameNil sets the value for UnitName to be an explicit nil

### UnsetUnitName
`func (o *CustomerMeter) UnsetUnitName()`

UnsetUnitName ensures that no value is present for UnitName, not even an explicit nil
### GetTotalQuantity

`func (o *CustomerMeter) GetTotalQuantity() float32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *CustomerMeter) GetTotalQuantityOk() (*float32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *CustomerMeter) SetTotalQuantity(v float32)`

SetTotalQuantity sets TotalQuantity field to given value.


### GetEventCount

`func (o *CustomerMeter) GetEventCount() int32`

GetEventCount returns the EventCount field if non-nil, zero value otherwise.

### GetEventCountOk

`func (o *CustomerMeter) GetEventCountOk() (*int32, bool)`

GetEventCountOk returns a tuple with the EventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCount

`func (o *CustomerMeter) SetEventCount(v int32)`

SetEventCount sets EventCount field to given value.


### GetLastEventAt

`func (o *CustomerMeter) GetLastEventAt() time.Time`

GetLastEventAt returns the LastEventAt field if non-nil, zero value otherwise.

### GetLastEventAtOk

`func (o *CustomerMeter) GetLastEventAtOk() (*time.Time, bool)`

GetLastEventAtOk returns a tuple with the LastEventAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventAt

`func (o *CustomerMeter) SetLastEventAt(v time.Time)`

SetLastEventAt sets LastEventAt field to given value.


### SetLastEventAtNil

`func (o *CustomerMeter) SetLastEventAtNil(b bool)`

 SetLastEventAtNil sets the value for LastEventAt to be an explicit nil

### UnsetLastEventAt
`func (o *CustomerMeter) UnsetLastEventAt()`

UnsetLastEventAt ensures that no value is present for LastEventAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


