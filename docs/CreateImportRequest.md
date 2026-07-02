# CreateImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of resource to import. | 
**FileUrl** | Pointer to **string** | URL to the CSV file to import (if not using multipart upload) | [optional] 
**ColumnMapping** | Pointer to **map[string]interface{}** | Map CSV column names to resource fields (e.g., { &#39;Email Address&#39;: &#39;email&#39;, &#39;Full Name&#39;: &#39;name&#39; }) | [optional] 
**OnDuplicate** | Pointer to **string** | Behavior when a duplicate record is found. skip: ignore the row. update: merge new data. error: fail the row. | [optional] [default to "skip"]

## Methods

### NewCreateImportRequest

`func NewCreateImportRequest(resourceType string, ) *CreateImportRequest`

NewCreateImportRequest instantiates a new CreateImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImportRequestWithDefaults

`func NewCreateImportRequestWithDefaults() *CreateImportRequest`

NewCreateImportRequestWithDefaults instantiates a new CreateImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateImportRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateImportRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateImportRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetFileUrl

`func (o *CreateImportRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *CreateImportRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *CreateImportRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *CreateImportRequest) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetColumnMapping

`func (o *CreateImportRequest) GetColumnMapping() map[string]interface{}`

GetColumnMapping returns the ColumnMapping field if non-nil, zero value otherwise.

### GetColumnMappingOk

`func (o *CreateImportRequest) GetColumnMappingOk() (*map[string]interface{}, bool)`

GetColumnMappingOk returns a tuple with the ColumnMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnMapping

`func (o *CreateImportRequest) SetColumnMapping(v map[string]interface{})`

SetColumnMapping sets ColumnMapping field to given value.

### HasColumnMapping

`func (o *CreateImportRequest) HasColumnMapping() bool`

HasColumnMapping returns a boolean if a field has been set.

### GetOnDuplicate

`func (o *CreateImportRequest) GetOnDuplicate() string`

GetOnDuplicate returns the OnDuplicate field if non-nil, zero value otherwise.

### GetOnDuplicateOk

`func (o *CreateImportRequest) GetOnDuplicateOk() (*string, bool)`

GetOnDuplicateOk returns a tuple with the OnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDuplicate

`func (o *CreateImportRequest) SetOnDuplicate(v string)`

SetOnDuplicate sets OnDuplicate field to given value.

### HasOnDuplicate

`func (o *CreateImportRequest) HasOnDuplicate() bool`

HasOnDuplicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


