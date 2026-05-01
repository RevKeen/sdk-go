# PayOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **string** | Saved payment method ID | [optional] 
**PaymentToken** | Pointer to **string** | One-time payment token | [optional] 
**Amount** | Pointer to **int32** | Partial payment amount in cents (defaults to full amount_due) | [optional] 

## Methods

### NewPayOrderRequest

`func NewPayOrderRequest() *PayOrderRequest`

NewPayOrderRequest instantiates a new PayOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOrderRequestWithDefaults

`func NewPayOrderRequestWithDefaults() *PayOrderRequest`

NewPayOrderRequestWithDefaults instantiates a new PayOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *PayOrderRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *PayOrderRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *PayOrderRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *PayOrderRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentToken

`func (o *PayOrderRequest) GetPaymentToken() string`

GetPaymentToken returns the PaymentToken field if non-nil, zero value otherwise.

### GetPaymentTokenOk

`func (o *PayOrderRequest) GetPaymentTokenOk() (*string, bool)`

GetPaymentTokenOk returns a tuple with the PaymentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentToken

`func (o *PayOrderRequest) SetPaymentToken(v string)`

SetPaymentToken sets PaymentToken field to given value.

### HasPaymentToken

`func (o *PayOrderRequest) HasPaymentToken() bool`

HasPaymentToken returns a boolean if a field has been set.

### GetAmount

`func (o *PayOrderRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PayOrderRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PayOrderRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PayOrderRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


