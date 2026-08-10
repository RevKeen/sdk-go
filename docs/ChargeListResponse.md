# ChargeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ChargeCreateResponseData**](ChargeCreateResponseData.md) |  | 
**Pagination** | [**ChargeListResponsePagination**](ChargeListResponsePagination.md) |  | 

## Methods

### NewChargeListResponse

`func NewChargeListResponse(data []ChargeCreateResponseData, pagination ChargeListResponsePagination, ) *ChargeListResponse`

NewChargeListResponse instantiates a new ChargeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeListResponseWithDefaults

`func NewChargeListResponseWithDefaults() *ChargeListResponse`

NewChargeListResponseWithDefaults instantiates a new ChargeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ChargeListResponse) GetData() []ChargeCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ChargeListResponse) GetDataOk() (*[]ChargeCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ChargeListResponse) SetData(v []ChargeCreateResponseData)`

SetData sets Data field to given value.


### GetPagination

`func (o *ChargeListResponse) GetPagination() ChargeListResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ChargeListResponse) GetPaginationOk() (*ChargeListResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ChargeListResponse) SetPagination(v ChargeListResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


