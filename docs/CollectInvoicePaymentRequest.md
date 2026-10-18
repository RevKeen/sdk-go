# CollectInvoicePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Collect the invoice&#39;s outstanding amount from the customer. The collection is scheduled, not completed: the invoice stays unpaid until the money settles. |
**PaymentMethod** | **string** | How to collect. Required. Only &#x60;direct_debit&#x60; is available; card collection is not available through this operation. |

## Methods

### NewCollectInvoicePaymentRequest

`func NewCollectInvoicePaymentRequest(mode string, paymentMethod string, ) *CollectInvoicePaymentRequest`

NewCollectInvoicePaymentRequest instantiates a new CollectInvoicePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectInvoicePaymentRequestWithDefaults

`func NewCollectInvoicePaymentRequestWithDefaults() *CollectInvoicePaymentRequest`

NewCollectInvoicePaymentRequestWithDefaults instantiates a new CollectInvoicePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *CollectInvoicePaymentRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CollectInvoicePaymentRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CollectInvoicePaymentRequest) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPaymentMethod

`func (o *CollectInvoicePaymentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CollectInvoicePaymentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CollectInvoicePaymentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
