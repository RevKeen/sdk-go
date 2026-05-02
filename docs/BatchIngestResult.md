# BatchIngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Summary** | [**BatchIngestResultSummary**](BatchIngestResultSummary.md) |  | 
**Data** | [**[]BatchIngestResultDataInner**](BatchIngestResultDataInner.md) |  | 

## Methods

### NewBatchIngestResult

`func NewBatchIngestResult(object string, summary BatchIngestResultSummary, data []BatchIngestResultDataInner, ) *BatchIngestResult`

NewBatchIngestResult instantiates a new BatchIngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchIngestResultWithDefaults

`func NewBatchIngestResultWithDefaults() *BatchIngestResult`

NewBatchIngestResultWithDefaults instantiates a new BatchIngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *BatchIngestResult) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *BatchIngestResult) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *BatchIngestResult) SetObject(v string)`

SetObject sets Object field to given value.


### GetSummary

`func (o *BatchIngestResult) GetSummary() BatchIngestResultSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BatchIngestResult) GetSummaryOk() (*BatchIngestResultSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BatchIngestResult) SetSummary(v BatchIngestResultSummary)`

SetSummary sets Summary field to given value.


### GetData

`func (o *BatchIngestResult) GetData() []BatchIngestResultDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BatchIngestResult) GetDataOk() (*[]BatchIngestResultDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BatchIngestResult) SetData(v []BatchIngestResultDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


