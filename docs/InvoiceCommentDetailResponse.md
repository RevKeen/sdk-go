# InvoiceCommentDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**InvoiceCommentDetail**](InvoiceCommentDetail.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewInvoiceCommentDetailResponse

`func NewInvoiceCommentDetailResponse(data InvoiceCommentDetail, requestId string, ) *InvoiceCommentDetailResponse`

NewInvoiceCommentDetailResponse instantiates a new InvoiceCommentDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCommentDetailResponseWithDefaults

`func NewInvoiceCommentDetailResponseWithDefaults() *InvoiceCommentDetailResponse`

NewInvoiceCommentDetailResponseWithDefaults instantiates a new InvoiceCommentDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InvoiceCommentDetailResponse) GetData() InvoiceCommentDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoiceCommentDetailResponse) GetDataOk() (*InvoiceCommentDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoiceCommentDetailResponse) SetData(v InvoiceCommentDetail)`

SetData sets Data field to given value.


### GetRequestId

`func (o *InvoiceCommentDetailResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InvoiceCommentDetailResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InvoiceCommentDetailResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


