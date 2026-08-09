# InvoicesSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Communication channel to send the invoice through | [optional] [default to "email"]
**TemplateId** | Pointer to **string** | Optional custom template ID to use for the notification | [optional] 

## Methods

### NewInvoicesSendRequest

`func NewInvoicesSendRequest() *InvoicesSendRequest`

NewInvoicesSendRequest instantiates a new InvoicesSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesSendRequestWithDefaults

`func NewInvoicesSendRequestWithDefaults() *InvoicesSendRequest`

NewInvoicesSendRequestWithDefaults instantiates a new InvoicesSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InvoicesSendRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InvoicesSendRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InvoicesSendRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InvoicesSendRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTemplateId

`func (o *InvoicesSendRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *InvoicesSendRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *InvoicesSendRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *InvoicesSendRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


