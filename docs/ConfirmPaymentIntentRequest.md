# ConfirmPaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethod** | Pointer to **string** | Payment method ID. Required if not already set on the intent. | [optional] 
**ReturnUrl** | Pointer to **string** | URL to redirect to after 3DS authentication | [optional] 

## Methods

### NewConfirmPaymentIntentRequest

`func NewConfirmPaymentIntentRequest() *ConfirmPaymentIntentRequest`

NewConfirmPaymentIntentRequest instantiates a new ConfirmPaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmPaymentIntentRequestWithDefaults

`func NewConfirmPaymentIntentRequestWithDefaults() *ConfirmPaymentIntentRequest`

NewConfirmPaymentIntentRequestWithDefaults instantiates a new ConfirmPaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethod

`func (o *ConfirmPaymentIntentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *ConfirmPaymentIntentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *ConfirmPaymentIntentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *ConfirmPaymentIntentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetReturnUrl

`func (o *ConfirmPaymentIntentRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *ConfirmPaymentIntentRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *ConfirmPaymentIntentRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *ConfirmPaymentIntentRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


