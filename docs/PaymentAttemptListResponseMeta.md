# PaymentAttemptListResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Number of results returned | 
**Limit** | Pointer to **int32** | Limit used in query | [optional] 

## Methods

### NewPaymentAttemptListResponseMeta

`func NewPaymentAttemptListResponseMeta(count int32, ) *PaymentAttemptListResponseMeta`

NewPaymentAttemptListResponseMeta instantiates a new PaymentAttemptListResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAttemptListResponseMetaWithDefaults

`func NewPaymentAttemptListResponseMetaWithDefaults() *PaymentAttemptListResponseMeta`

NewPaymentAttemptListResponseMetaWithDefaults instantiates a new PaymentAttemptListResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaymentAttemptListResponseMeta) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentAttemptListResponseMeta) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentAttemptListResponseMeta) SetCount(v int32)`

SetCount sets Count field to given value.


### GetLimit

`func (o *PaymentAttemptListResponseMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaymentAttemptListResponseMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaymentAttemptListResponseMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaymentAttemptListResponseMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


