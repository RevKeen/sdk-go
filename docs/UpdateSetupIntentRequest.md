# UpdateSetupIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | ID of the customer to attach the payment method to | [optional] 
**PaymentMethodTypes** | Pointer to **[]string** | Allowed payment method types | [optional] 
**Description** | Pointer to **string** | Merchant description for reference | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata to attach | [optional] 

## Methods

### NewUpdateSetupIntentRequest

`func NewUpdateSetupIntentRequest() *UpdateSetupIntentRequest`

NewUpdateSetupIntentRequest instantiates a new UpdateSetupIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSetupIntentRequestWithDefaults

`func NewUpdateSetupIntentRequestWithDefaults() *UpdateSetupIntentRequest`

NewUpdateSetupIntentRequestWithDefaults instantiates a new UpdateSetupIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *UpdateSetupIntentRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UpdateSetupIntentRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UpdateSetupIntentRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UpdateSetupIntentRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetPaymentMethodTypes

`func (o *UpdateSetupIntentRequest) GetPaymentMethodTypes() []string`

GetPaymentMethodTypes returns the PaymentMethodTypes field if non-nil, zero value otherwise.

### GetPaymentMethodTypesOk

`func (o *UpdateSetupIntentRequest) GetPaymentMethodTypesOk() (*[]string, bool)`

GetPaymentMethodTypesOk returns a tuple with the PaymentMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTypes

`func (o *UpdateSetupIntentRequest) SetPaymentMethodTypes(v []string)`

SetPaymentMethodTypes sets PaymentMethodTypes field to given value.

### HasPaymentMethodTypes

`func (o *UpdateSetupIntentRequest) HasPaymentMethodTypes() bool`

HasPaymentMethodTypes returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSetupIntentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSetupIntentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSetupIntentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSetupIntentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateSetupIntentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateSetupIntentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateSetupIntentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateSetupIntentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


