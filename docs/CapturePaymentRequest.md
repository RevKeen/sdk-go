# CapturePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount to capture in cents. Defaults to full authorized amount. Must be &lt;&#x3D; authorized amount. | [optional] 

## Methods

### NewCapturePaymentRequest

`func NewCapturePaymentRequest() *CapturePaymentRequest`

NewCapturePaymentRequest instantiates a new CapturePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapturePaymentRequestWithDefaults

`func NewCapturePaymentRequestWithDefaults() *CapturePaymentRequest`

NewCapturePaymentRequestWithDefaults instantiates a new CapturePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CapturePaymentRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CapturePaymentRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CapturePaymentRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CapturePaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


