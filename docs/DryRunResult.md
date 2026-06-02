# DryRunResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Summary** | [**DryRunResultSummary**](DryRunResultSummary.md) |  | 
**Data** | [**[]DryRunResultDataInner**](DryRunResultDataInner.md) |  | 

## Methods

### NewDryRunResult

`func NewDryRunResult(object string, summary DryRunResultSummary, data []DryRunResultDataInner, ) *DryRunResult`

NewDryRunResult instantiates a new DryRunResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunResultWithDefaults

`func NewDryRunResultWithDefaults() *DryRunResult`

NewDryRunResultWithDefaults instantiates a new DryRunResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *DryRunResult) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DryRunResult) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DryRunResult) SetObject(v string)`

SetObject sets Object field to given value.


### GetSummary

`func (o *DryRunResult) GetSummary() DryRunResultSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DryRunResult) GetSummaryOk() (*DryRunResultSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DryRunResult) SetSummary(v DryRunResultSummary)`

SetSummary sets Summary field to given value.


### GetData

`func (o *DryRunResult) GetData() []DryRunResultDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DryRunResult) GetDataOk() (*[]DryRunResultDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DryRunResult) SetData(v []DryRunResultDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


