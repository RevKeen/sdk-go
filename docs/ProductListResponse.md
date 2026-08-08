# ProductListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Product**](Product.md) |  | 
**Pagination** | [**Pagination**](Pagination.md) |  | 

## Methods

### NewProductListResponse

`func NewProductListResponse(data []Product, pagination Pagination, ) *ProductListResponse`

NewProductListResponse instantiates a new ProductListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductListResponseWithDefaults

`func NewProductListResponseWithDefaults() *ProductListResponse`

NewProductListResponseWithDefaults instantiates a new ProductListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProductListResponse) GetData() []Product`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductListResponse) GetDataOk() (*[]Product, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductListResponse) SetData(v []Product)`

SetData sets Data field to given value.


### GetPagination

`func (o *ProductListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ProductListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ProductListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


