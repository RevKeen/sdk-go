# UpdatePaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Amount in cents | [optional] 
**Currency** | Pointer to **string** | Three-letter ISO currency code | [optional] 
**PaymentMethod** | Pointer to **string** | Payment method ID | [optional] 
**Description** | Pointer to **string** | Description for merchant reference | [optional] 
**StatementDescriptor** | Pointer to **string** | Statement descriptor | [optional] 
**StatementDescriptorSuffix** | Pointer to **string** | Statement descriptor suffix | [optional] 
**ReceiptEmail** | Pointer to **string** | Email to send receipt to | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewUpdatePaymentIntentRequest

`func NewUpdatePaymentIntentRequest() *UpdatePaymentIntentRequest`

NewUpdatePaymentIntentRequest instantiates a new UpdatePaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentIntentRequestWithDefaults

`func NewUpdatePaymentIntentRequestWithDefaults() *UpdatePaymentIntentRequest`

NewUpdatePaymentIntentRequestWithDefaults instantiates a new UpdatePaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UpdatePaymentIntentRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdatePaymentIntentRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdatePaymentIntentRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdatePaymentIntentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdatePaymentIntentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdatePaymentIntentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdatePaymentIntentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdatePaymentIntentRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *UpdatePaymentIntentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *UpdatePaymentIntentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *UpdatePaymentIntentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *UpdatePaymentIntentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePaymentIntentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePaymentIntentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePaymentIntentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePaymentIntentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *UpdatePaymentIntentRequest) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *UpdatePaymentIntentRequest) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *UpdatePaymentIntentRequest) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *UpdatePaymentIntentRequest) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetStatementDescriptorSuffix

`func (o *UpdatePaymentIntentRequest) GetStatementDescriptorSuffix() string`

GetStatementDescriptorSuffix returns the StatementDescriptorSuffix field if non-nil, zero value otherwise.

### GetStatementDescriptorSuffixOk

`func (o *UpdatePaymentIntentRequest) GetStatementDescriptorSuffixOk() (*string, bool)`

GetStatementDescriptorSuffixOk returns a tuple with the StatementDescriptorSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptorSuffix

`func (o *UpdatePaymentIntentRequest) SetStatementDescriptorSuffix(v string)`

SetStatementDescriptorSuffix sets StatementDescriptorSuffix field to given value.

### HasStatementDescriptorSuffix

`func (o *UpdatePaymentIntentRequest) HasStatementDescriptorSuffix() bool`

HasStatementDescriptorSuffix returns a boolean if a field has been set.

### GetReceiptEmail

`func (o *UpdatePaymentIntentRequest) GetReceiptEmail() string`

GetReceiptEmail returns the ReceiptEmail field if non-nil, zero value otherwise.

### GetReceiptEmailOk

`func (o *UpdatePaymentIntentRequest) GetReceiptEmailOk() (*string, bool)`

GetReceiptEmailOk returns a tuple with the ReceiptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptEmail

`func (o *UpdatePaymentIntentRequest) SetReceiptEmail(v string)`

SetReceiptEmail sets ReceiptEmail field to given value.

### HasReceiptEmail

`func (o *UpdatePaymentIntentRequest) HasReceiptEmail() bool`

HasReceiptEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePaymentIntentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePaymentIntentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePaymentIntentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePaymentIntentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


