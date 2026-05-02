# Import

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique import job ID | 
**Object** | **string** | Object type | 
**Status** | **string** | Import job status. pending: queued. validating: checking data. processing: creating records. completed: all rows imported. completed_with_errors: some rows failed. failed: import aborted. | 
**ResourceType** | **string** | The type of resource to import. | 
**TotalRows** | Pointer to **NullableInt32** | Total rows in the uploaded file | [optional] 
**ProcessedRows** | Pointer to **NullableInt32** | Number of rows processed so far | [optional] 
**SuccessCount** | Pointer to **NullableInt32** | Number of rows successfully imported | [optional] 
**ErrorCount** | Pointer to **NullableInt32** | Number of rows that failed | [optional] 
**Errors** | Pointer to [**[]ImportErrorsInner**](ImportErrorsInner.md) | Detailed error information for failed rows | [optional] 
**CreatedAt** | **time.Time** | When the import was created (ISO 8601) | 
**CompletedAt** | Pointer to **NullableTime** | When the import completed (ISO 8601) | [optional] 

## Methods

### NewImport

`func NewImport(id string, object string, status string, resourceType string, createdAt time.Time, ) *Import`

NewImport instantiates a new Import object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportWithDefaults

`func NewImportWithDefaults() *Import`

NewImportWithDefaults instantiates a new Import object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Import) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Import) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Import) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Import) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Import) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Import) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *Import) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Import) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Import) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResourceType

`func (o *Import) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Import) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Import) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetTotalRows

`func (o *Import) GetTotalRows() int32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *Import) GetTotalRowsOk() (*int32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *Import) SetTotalRows(v int32)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *Import) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### SetTotalRowsNil

`func (o *Import) SetTotalRowsNil(b bool)`

 SetTotalRowsNil sets the value for TotalRows to be an explicit nil

### UnsetTotalRows
`func (o *Import) UnsetTotalRows()`

UnsetTotalRows ensures that no value is present for TotalRows, not even an explicit nil
### GetProcessedRows

`func (o *Import) GetProcessedRows() int32`

GetProcessedRows returns the ProcessedRows field if non-nil, zero value otherwise.

### GetProcessedRowsOk

`func (o *Import) GetProcessedRowsOk() (*int32, bool)`

GetProcessedRowsOk returns a tuple with the ProcessedRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRows

`func (o *Import) SetProcessedRows(v int32)`

SetProcessedRows sets ProcessedRows field to given value.

### HasProcessedRows

`func (o *Import) HasProcessedRows() bool`

HasProcessedRows returns a boolean if a field has been set.

### SetProcessedRowsNil

`func (o *Import) SetProcessedRowsNil(b bool)`

 SetProcessedRowsNil sets the value for ProcessedRows to be an explicit nil

### UnsetProcessedRows
`func (o *Import) UnsetProcessedRows()`

UnsetProcessedRows ensures that no value is present for ProcessedRows, not even an explicit nil
### GetSuccessCount

`func (o *Import) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *Import) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *Import) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *Import) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### SetSuccessCountNil

`func (o *Import) SetSuccessCountNil(b bool)`

 SetSuccessCountNil sets the value for SuccessCount to be an explicit nil

### UnsetSuccessCount
`func (o *Import) UnsetSuccessCount()`

UnsetSuccessCount ensures that no value is present for SuccessCount, not even an explicit nil
### GetErrorCount

`func (o *Import) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *Import) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *Import) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *Import) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### SetErrorCountNil

`func (o *Import) SetErrorCountNil(b bool)`

 SetErrorCountNil sets the value for ErrorCount to be an explicit nil

### UnsetErrorCount
`func (o *Import) UnsetErrorCount()`

UnsetErrorCount ensures that no value is present for ErrorCount, not even an explicit nil
### GetErrors

`func (o *Import) GetErrors() []ImportErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Import) GetErrorsOk() (*[]ImportErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Import) SetErrors(v []ImportErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Import) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *Import) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *Import) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil
### GetCreatedAt

`func (o *Import) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Import) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Import) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *Import) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Import) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Import) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Import) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Import) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Import) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


