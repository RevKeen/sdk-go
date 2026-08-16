# RefundListResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | Total number of results matching the query | 
**Limit** | **float32** | Maximum results per page | 
**Offset** | **float32** | Number of results skipped | 
**HasMore** | **bool** | Whether more results exist beyond this page | 

## Methods

### NewRefundListResponseMeta

`func NewRefundListResponseMeta(total float32, limit float32, offset float32, hasMore bool, ) *RefundListResponseMeta`

NewRefundListResponseMeta instantiates a new RefundListResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundListResponseMetaWithDefaults

`func NewRefundListResponseMetaWithDefaults() *RefundListResponseMeta`

NewRefundListResponseMetaWithDefaults instantiates a new RefundListResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *RefundListResponseMeta) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RefundListResponseMeta) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RefundListResponseMeta) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *RefundListResponseMeta) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RefundListResponseMeta) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RefundListResponseMeta) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *RefundListResponseMeta) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RefundListResponseMeta) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RefundListResponseMeta) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetHasMore

`func (o *RefundListResponseMeta) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *RefundListResponseMeta) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *RefundListResponseMeta) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


