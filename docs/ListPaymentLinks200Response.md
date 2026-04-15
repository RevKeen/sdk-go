# ListPaymentLinks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ListPaymentLinks200ResponseDataInner**](ListPaymentLinks200ResponseDataInner.md) |  | 
**Pagination** | [**ListPaymentLinks200ResponsePagination**](ListPaymentLinks200ResponsePagination.md) |  | 

## Methods

### NewListPaymentLinks200Response

`func NewListPaymentLinks200Response(data []ListPaymentLinks200ResponseDataInner, pagination ListPaymentLinks200ResponsePagination, ) *ListPaymentLinks200Response`

NewListPaymentLinks200Response instantiates a new ListPaymentLinks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentLinks200ResponseWithDefaults

`func NewListPaymentLinks200ResponseWithDefaults() *ListPaymentLinks200Response`

NewListPaymentLinks200ResponseWithDefaults instantiates a new ListPaymentLinks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListPaymentLinks200Response) GetData() []ListPaymentLinks200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListPaymentLinks200Response) GetDataOk() (*[]ListPaymentLinks200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListPaymentLinks200Response) SetData(v []ListPaymentLinks200ResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *ListPaymentLinks200Response) GetPagination() ListPaymentLinks200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPaymentLinks200Response) GetPaginationOk() (*ListPaymentLinks200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPaymentLinks200Response) SetPagination(v ListPaymentLinks200ResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


