# EntitlementListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Entitlement**](Entitlement.md) |  | 
**Pagination** | [**ListPaymentLinks200ResponsePagination**](ListPaymentLinks200ResponsePagination.md) |  | 
**Customer** | [**EntitlementListResponseCustomer**](EntitlementListResponseCustomer.md) |  | 

## Methods

### NewEntitlementListResponse

`func NewEntitlementListResponse(data []Entitlement, pagination ListPaymentLinks200ResponsePagination, customer EntitlementListResponseCustomer, ) *EntitlementListResponse`

NewEntitlementListResponse instantiates a new EntitlementListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementListResponseWithDefaults

`func NewEntitlementListResponseWithDefaults() *EntitlementListResponse`

NewEntitlementListResponseWithDefaults instantiates a new EntitlementListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EntitlementListResponse) GetData() []Entitlement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EntitlementListResponse) GetDataOk() (*[]Entitlement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EntitlementListResponse) SetData(v []Entitlement)`

SetData sets Data field to given value.


### GetPagination

`func (o *EntitlementListResponse) GetPagination() ListPaymentLinks200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *EntitlementListResponse) GetPaginationOk() (*ListPaymentLinks200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *EntitlementListResponse) SetPagination(v ListPaymentLinks200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetCustomer

`func (o *EntitlementListResponse) GetCustomer() EntitlementListResponseCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *EntitlementListResponse) GetCustomerOk() (*EntitlementListResponseCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *EntitlementListResponse) SetCustomer(v EntitlementListResponseCustomer)`

SetCustomer sets Customer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


