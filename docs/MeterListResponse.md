# MeterListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]map[string]interface{}** | Array of meter objects | 
**Pagination** | [**MeterListResponsePagination**](MeterListResponsePagination.md) |  | 

## Methods

### NewMeterListResponse

`func NewMeterListResponse(data []map[string]interface{}, pagination MeterListResponsePagination, ) *MeterListResponse`

NewMeterListResponse instantiates a new MeterListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterListResponseWithDefaults

`func NewMeterListResponseWithDefaults() *MeterListResponse`

NewMeterListResponseWithDefaults instantiates a new MeterListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MeterListResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MeterListResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MeterListResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.


### GetPagination

`func (o *MeterListResponse) GetPagination() MeterListResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MeterListResponse) GetPaginationOk() (*MeterListResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MeterListResponse) SetPagination(v MeterListResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


