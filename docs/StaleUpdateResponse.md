# StaleUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ignored** | **bool** |  | 
**Reason** | **string** |  | 
**StoredTimestamp** | **string** |  | 
**IncomingTimestamp** | **string** |  | 
**RequestId** | **string** |  | 

## Methods

### NewStaleUpdateResponse

`func NewStaleUpdateResponse(ignored bool, reason string, storedTimestamp string, incomingTimestamp string, requestId string, ) *StaleUpdateResponse`

NewStaleUpdateResponse instantiates a new StaleUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaleUpdateResponseWithDefaults

`func NewStaleUpdateResponseWithDefaults() *StaleUpdateResponse`

NewStaleUpdateResponseWithDefaults instantiates a new StaleUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnored

`func (o *StaleUpdateResponse) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *StaleUpdateResponse) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *StaleUpdateResponse) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.


### GetReason

`func (o *StaleUpdateResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StaleUpdateResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StaleUpdateResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStoredTimestamp

`func (o *StaleUpdateResponse) GetStoredTimestamp() string`

GetStoredTimestamp returns the StoredTimestamp field if non-nil, zero value otherwise.

### GetStoredTimestampOk

`func (o *StaleUpdateResponse) GetStoredTimestampOk() (*string, bool)`

GetStoredTimestampOk returns a tuple with the StoredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredTimestamp

`func (o *StaleUpdateResponse) SetStoredTimestamp(v string)`

SetStoredTimestamp sets StoredTimestamp field to given value.


### GetIncomingTimestamp

`func (o *StaleUpdateResponse) GetIncomingTimestamp() string`

GetIncomingTimestamp returns the IncomingTimestamp field if non-nil, zero value otherwise.

### GetIncomingTimestampOk

`func (o *StaleUpdateResponse) GetIncomingTimestampOk() (*string, bool)`

GetIncomingTimestampOk returns a tuple with the IncomingTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingTimestamp

`func (o *StaleUpdateResponse) SetIncomingTimestamp(v string)`

SetIncomingTimestamp sets IncomingTimestamp field to given value.


### GetRequestId

`func (o *StaleUpdateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *StaleUpdateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *StaleUpdateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


