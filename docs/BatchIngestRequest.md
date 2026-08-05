# BatchIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]UsageEvent**](UsageEvent.md) | Array of usage events (1–1000) | 

## Methods

### NewBatchIngestRequest

`func NewBatchIngestRequest(events []UsageEvent, ) *BatchIngestRequest`

NewBatchIngestRequest instantiates a new BatchIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchIngestRequestWithDefaults

`func NewBatchIngestRequestWithDefaults() *BatchIngestRequest`

NewBatchIngestRequestWithDefaults instantiates a new BatchIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *BatchIngestRequest) GetEvents() []UsageEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BatchIngestRequest) GetEventsOk() (*[]UsageEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BatchIngestRequest) SetEvents(v []UsageEvent)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


