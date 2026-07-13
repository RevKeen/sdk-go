# VoidListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Void**](Void.md) |  | 
**Meta** | [**RefundListResponseMeta**](RefundListResponseMeta.md) |  | 

## Methods

### NewVoidListResponse

`func NewVoidListResponse(data []Void, meta RefundListResponseMeta, ) *VoidListResponse`

NewVoidListResponse instantiates a new VoidListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidListResponseWithDefaults

`func NewVoidListResponseWithDefaults() *VoidListResponse`

NewVoidListResponseWithDefaults instantiates a new VoidListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VoidListResponse) GetData() []Void`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VoidListResponse) GetDataOk() (*[]Void, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VoidListResponse) SetData(v []Void)`

SetData sets Data field to given value.


### GetMeta

`func (o *VoidListResponse) GetMeta() RefundListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VoidListResponse) GetMetaOk() (*RefundListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VoidListResponse) SetMeta(v RefundListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


