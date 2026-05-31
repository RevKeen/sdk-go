# SubscriptionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionCreateResponseData**](SubscriptionCreateResponseData.md) |  | 
**Pagination** | [**SubscriptionListResponsePagination**](SubscriptionListResponsePagination.md) |  | 

## Methods

### NewSubscriptionListResponse

`func NewSubscriptionListResponse(data []SubscriptionCreateResponseData, pagination SubscriptionListResponsePagination, ) *SubscriptionListResponse`

NewSubscriptionListResponse instantiates a new SubscriptionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListResponseWithDefaults

`func NewSubscriptionListResponseWithDefaults() *SubscriptionListResponse`

NewSubscriptionListResponseWithDefaults instantiates a new SubscriptionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionListResponse) GetData() []SubscriptionCreateResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionListResponse) GetDataOk() (*[]SubscriptionCreateResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionListResponse) SetData(v []SubscriptionCreateResponseData)`

SetData sets Data field to given value.


### GetPagination

`func (o *SubscriptionListResponse) GetPagination() SubscriptionListResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SubscriptionListResponse) GetPaginationOk() (*SubscriptionListResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SubscriptionListResponse) SetPagination(v SubscriptionListResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


