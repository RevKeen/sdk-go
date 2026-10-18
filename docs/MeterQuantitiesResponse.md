# MeterQuantitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  |
**MeterId** | **string** |  |
**StartTimestamp** | **time.Time** |  |
**EndTimestamp** | **time.Time** |  |
**Interval** | **string** |  |
**Quantities** | [**[]MeterQuantitiesResponseQuantitiesInner**](MeterQuantitiesResponseQuantitiesInner.md) |  |

## Methods

### NewMeterQuantitiesResponse

`func NewMeterQuantitiesResponse(object string, meterId string, startTimestamp time.Time, endTimestamp time.Time, interval string, quantities []MeterQuantitiesResponseQuantitiesInner, ) *MeterQuantitiesResponse`

NewMeterQuantitiesResponse instantiates a new MeterQuantitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterQuantitiesResponseWithDefaults

`func NewMeterQuantitiesResponseWithDefaults() *MeterQuantitiesResponse`

NewMeterQuantitiesResponseWithDefaults instantiates a new MeterQuantitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MeterQuantitiesResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MeterQuantitiesResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MeterQuantitiesResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetMeterId

`func (o *MeterQuantitiesResponse) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *MeterQuantitiesResponse) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *MeterQuantitiesResponse) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetStartTimestamp

`func (o *MeterQuantitiesResponse) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *MeterQuantitiesResponse) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *MeterQuantitiesResponse) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.


### GetEndTimestamp

`func (o *MeterQuantitiesResponse) GetEndTimestamp() time.Time`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *MeterQuantitiesResponse) GetEndTimestampOk() (*time.Time, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *MeterQuantitiesResponse) SetEndTimestamp(v time.Time)`

SetEndTimestamp sets EndTimestamp field to given value.


### GetInterval

`func (o *MeterQuantitiesResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *MeterQuantitiesResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *MeterQuantitiesResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetQuantities

`func (o *MeterQuantitiesResponse) GetQuantities() []MeterQuantitiesResponseQuantitiesInner`

GetQuantities returns the Quantities field if non-nil, zero value otherwise.

### GetQuantitiesOk

`func (o *MeterQuantitiesResponse) GetQuantitiesOk() (*[]MeterQuantitiesResponseQuantitiesInner, bool)`

GetQuantitiesOk returns a tuple with the Quantities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantities

`func (o *MeterQuantitiesResponse) SetQuantities(v []MeterQuantitiesResponseQuantitiesInner)`

SetQuantities sets Quantities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
