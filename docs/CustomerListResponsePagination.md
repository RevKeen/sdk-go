# CustomerListResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **float32** |  | 
**Offset** | **float32** |  | 
**Total** | Pointer to **float32** |  | [optional] 

## Methods

### NewCustomerListResponsePagination

`func NewCustomerListResponsePagination(limit float32, offset float32, ) *CustomerListResponsePagination`

NewCustomerListResponsePagination instantiates a new CustomerListResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerListResponsePaginationWithDefaults

`func NewCustomerListResponsePaginationWithDefaults() *CustomerListResponsePagination`

NewCustomerListResponsePaginationWithDefaults instantiates a new CustomerListResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *CustomerListResponsePagination) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CustomerListResponsePagination) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CustomerListResponsePagination) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *CustomerListResponsePagination) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CustomerListResponsePagination) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CustomerListResponsePagination) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *CustomerListResponsePagination) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CustomerListResponsePagination) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CustomerListResponsePagination) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CustomerListResponsePagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


