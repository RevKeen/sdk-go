# AttachPaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | ID of the customer to attach the payment method to | 

## Methods

### NewAttachPaymentMethodRequest

`func NewAttachPaymentMethodRequest(customerId string, ) *AttachPaymentMethodRequest`

NewAttachPaymentMethodRequest instantiates a new AttachPaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPaymentMethodRequestWithDefaults

`func NewAttachPaymentMethodRequestWithDefaults() *AttachPaymentMethodRequest`

NewAttachPaymentMethodRequestWithDefaults instantiates a new AttachPaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *AttachPaymentMethodRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *AttachPaymentMethodRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *AttachPaymentMethodRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


