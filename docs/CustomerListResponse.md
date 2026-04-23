# CustomerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CustomerCreateResponseData**](CustomerCreateResponseData.md) |  | 
**Pagination** | [**CustomerListResponsePagination**](CustomerListResponsePagination.md) |  | 

## Methods

### NewCustomerListResponse

`func NewCustomerListResponse(data []CustomerCreateResponseData, pagination CustomerListResponsePagination, ) *CustomerListResponse`

NewCustomerListResponse instantiates a new CustomerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListResponseWithDefaults

`func NewCustomerListResponseWithDefaults() *CustomerListResponse`

NewCustomerListResponseWithDefaults instantiates a new CustomerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomerListResponse) GetData() []CustomerCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerListResponse) GetDataOk() (*[]CustomerCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerListResponse) SetData(v []CustomerCreateResponseData)`

SetData sets Data field to given value.


### GetPagination

`func (o *CustomerListResponse) GetPagination() CustomerListResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CustomerListResponse) GetPaginationOk() (*CustomerListResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CustomerListResponse) SetPagination(v CustomerListResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


