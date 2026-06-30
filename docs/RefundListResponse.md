# RefundListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Refund**](Refund.md) |  | 
**Meta** | [**RefundListResponseMeta**](RefundListResponseMeta.md) |  | 

## Methods

### NewRefundListResponse

`func NewRefundListResponse(data []Refund, meta RefundListResponseMeta, ) *RefundListResponse`

NewRefundListResponse instantiates a new RefundListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundListResponseWithDefaults

`func NewRefundListResponseWithDefaults() *RefundListResponse`

NewRefundListResponseWithDefaults instantiates a new RefundListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RefundListResponse) GetData() []Refund`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RefundListResponse) GetDataOk() (*[]Refund, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RefundListResponse) SetData(v []Refund)`

SetData sets Data field to given value.


### GetMeta

`func (o *RefundListResponse) GetMeta() RefundListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RefundListResponse) GetMetaOk() (*RefundListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RefundListResponse) SetMeta(v RefundListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


