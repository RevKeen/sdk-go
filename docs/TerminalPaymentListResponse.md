# TerminalPaymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TerminalPayment**](TerminalPayment.md) |  | 
**Meta** | [**TerminalPaymentListResponseMeta**](TerminalPaymentListResponseMeta.md) |  | 

## Methods

### NewTerminalPaymentListResponse

`func NewTerminalPaymentListResponse(data []TerminalPayment, meta TerminalPaymentListResponseMeta, ) *TerminalPaymentListResponse`

NewTerminalPaymentListResponse instantiates a new TerminalPaymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalPaymentListResponseWithDefaults

`func NewTerminalPaymentListResponseWithDefaults() *TerminalPaymentListResponse`

NewTerminalPaymentListResponseWithDefaults instantiates a new TerminalPaymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TerminalPaymentListResponse) GetData() []TerminalPayment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TerminalPaymentListResponse) GetDataOk() (*[]TerminalPayment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TerminalPaymentListResponse) SetData(v []TerminalPayment)`

SetData sets Data field to given value.


### GetMeta

`func (o *TerminalPaymentListResponse) GetMeta() TerminalPaymentListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TerminalPaymentListResponse) GetMetaOk() (*TerminalPaymentListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TerminalPaymentListResponse) SetMeta(v TerminalPaymentListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


