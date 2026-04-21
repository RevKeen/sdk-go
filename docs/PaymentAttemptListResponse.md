# PaymentAttemptListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]PaymentAttempt**](PaymentAttempt.md) |  | 
**Meta** | [**PaymentAttemptListResponseMeta**](PaymentAttemptListResponseMeta.md) |  | 

## Methods

### NewPaymentAttemptListResponse

`func NewPaymentAttemptListResponse(data []PaymentAttempt, meta PaymentAttemptListResponseMeta, ) *PaymentAttemptListResponse`

NewPaymentAttemptListResponse instantiates a new PaymentAttemptListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAttemptListResponseWithDefaults

`func NewPaymentAttemptListResponseWithDefaults() *PaymentAttemptListResponse`

NewPaymentAttemptListResponseWithDefaults instantiates a new PaymentAttemptListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaymentAttemptListResponse) GetData() []PaymentAttempt`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentAttemptListResponse) GetDataOk() (*[]PaymentAttempt, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentAttemptListResponse) SetData(v []PaymentAttempt)`

SetData sets Data field to given value.


### GetMeta

`func (o *PaymentAttemptListResponse) GetMeta() PaymentAttemptListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PaymentAttemptListResponse) GetMetaOk() (*PaymentAttemptListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PaymentAttemptListResponse) SetMeta(v PaymentAttemptListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


