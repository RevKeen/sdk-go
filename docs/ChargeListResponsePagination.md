# ChargeListResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **float32** | Current page number | 
**Limit** | **float32** | Maximum results per page | 
**Total** | **float32** | Total number of results matching the query | 
**TotalPages** | **float32** | Total number of pages available | 

## Methods

### NewChargeListResponsePagination

`func NewChargeListResponsePagination(page float32, limit float32, total float32, totalPages float32, ) *ChargeListResponsePagination`

NewChargeListResponsePagination instantiates a new ChargeListResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeListResponsePaginationWithDefaults

`func NewChargeListResponsePaginationWithDefaults() *ChargeListResponsePagination`

NewChargeListResponsePaginationWithDefaults instantiates a new ChargeListResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ChargeListResponsePagination) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ChargeListResponsePagination) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ChargeListResponsePagination) SetPage(v float32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *ChargeListResponsePagination) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ChargeListResponsePagination) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ChargeListResponsePagination) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *ChargeListResponsePagination) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ChargeListResponsePagination) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ChargeListResponsePagination) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *ChargeListResponsePagination) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *ChargeListResponsePagination) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *ChargeListResponsePagination) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


