# CreateExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The type of resource to export. | 
**Format** | Pointer to **string** | Output file format. csv: Comma-separated values. xlsx: Excel workbook. | [optional] [default to "csv"]
**Filters** | Pointer to **map[string]interface{}** | Filters to apply (e.g., { status: &#39;paid&#39;, created_gte: &#39;2024-01-01&#39; }) | [optional] 
**ColumnSet** | Pointer to **string** | Column set to include. standard: common fields. full: all fields. minimal: IDs and key fields only. | [optional] [default to "standard"]

## Methods

### NewCreateExportRequest

`func NewCreateExportRequest(resourceType string, ) *CreateExportRequest`

NewCreateExportRequest instantiates a new CreateExportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExportRequestWithDefaults

`func NewCreateExportRequestWithDefaults() *CreateExportRequest`

NewCreateExportRequestWithDefaults instantiates a new CreateExportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateExportRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateExportRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateExportRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetFormat

`func (o *CreateExportRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateExportRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateExportRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateExportRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFilters

`func (o *CreateExportRequest) GetFilters() map[string]interface{}`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CreateExportRequest) GetFiltersOk() (*map[string]interface{}, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CreateExportRequest) SetFilters(v map[string]interface{})`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CreateExportRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetColumnSet

`func (o *CreateExportRequest) GetColumnSet() string`

GetColumnSet returns the ColumnSet field if non-nil, zero value otherwise.

### GetColumnSetOk

`func (o *CreateExportRequest) GetColumnSetOk() (*string, bool)`

GetColumnSetOk returns a tuple with the ColumnSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnSet

`func (o *CreateExportRequest) SetColumnSet(v string)`

SetColumnSet sets ColumnSet field to given value.

### HasColumnSet

`func (o *CreateExportRequest) HasColumnSet() bool`

HasColumnSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


