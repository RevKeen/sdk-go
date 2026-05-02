# CapturePaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountToCapture** | Pointer to **int32** | Amount to capture in cents. Defaults to full authorized amount. | [optional] 

## Methods

### NewCapturePaymentIntentRequest

`func NewCapturePaymentIntentRequest() *CapturePaymentIntentRequest`

NewCapturePaymentIntentRequest instantiates a new CapturePaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapturePaymentIntentRequestWithDefaults

`func NewCapturePaymentIntentRequestWithDefaults() *CapturePaymentIntentRequest`

NewCapturePaymentIntentRequestWithDefaults instantiates a new CapturePaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountToCapture

`func (o *CapturePaymentIntentRequest) GetAmountToCapture() int32`

GetAmountToCapture returns the AmountToCapture field if non-nil, zero value otherwise.

### GetAmountToCaptureOk

`func (o *CapturePaymentIntentRequest) GetAmountToCaptureOk() (*int32, bool)`

GetAmountToCaptureOk returns a tuple with the AmountToCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToCapture

`func (o *CapturePaymentIntentRequest) SetAmountToCapture(v int32)`

SetAmountToCapture sets AmountToCapture field to given value.

### HasAmountToCapture

`func (o *CapturePaymentIntentRequest) HasAmountToCapture() bool`

HasAmountToCapture returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


