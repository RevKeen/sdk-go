# CreateSetupIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | ID of the customer to attach the payment method to | [optional] 
**PaymentMethodTypes** | Pointer to **[]string** | Allowed payment method types | [optional] [default to {"card"}]
**Usage** | Pointer to **string** | How the payment method will be used | [optional] [default to "off_session"]
**Description** | Pointer to **string** | Merchant description for reference | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata to attach | [optional] 

## Methods

### NewCreateSetupIntentRequest

`func NewCreateSetupIntentRequest() *CreateSetupIntentRequest`

NewCreateSetupIntentRequest instantiates a new CreateSetupIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSetupIntentRequestWithDefaults

`func NewCreateSetupIntentRequestWithDefaults() *CreateSetupIntentRequest`

NewCreateSetupIntentRequestWithDefaults instantiates a new CreateSetupIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateSetupIntentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateSetupIntentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateSetupIntentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreateSetupIntentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetPaymentMethodTypes

`func (o *CreateSetupIntentRequest) GetPaymentMethodTypes() []string`

GetPaymentMethodTypes returns the PaymentMethodTypes field if non-nil, zero value otherwise.

### GetPaymentMethodTypesOk

`func (o *CreateSetupIntentRequest) GetPaymentMethodTypesOk() (*[]string, bool)`

GetPaymentMethodTypesOk returns a tuple with the PaymentMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTypes

`func (o *CreateSetupIntentRequest) SetPaymentMethodTypes(v []string)`

SetPaymentMethodTypes sets PaymentMethodTypes field to given value.

### HasPaymentMethodTypes

`func (o *CreateSetupIntentRequest) HasPaymentMethodTypes() bool`

HasPaymentMethodTypes returns a boolean if a field has been set.

### GetUsage

`func (o *CreateSetupIntentRequest) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateSetupIntentRequest) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateSetupIntentRequest) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CreateSetupIntentRequest) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetDescription

`func (o *CreateSetupIntentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSetupIntentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSetupIntentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSetupIntentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateSetupIntentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateSetupIntentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateSetupIntentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateSetupIntentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


