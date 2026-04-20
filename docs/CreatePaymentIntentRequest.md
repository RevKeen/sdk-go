# CreatePaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount in cents. Must be greater than 0. | 
**Currency** | Pointer to **string** | Three-letter ISO currency code. Defaults to USD. | [optional] [default to "USD"]
**Customer** | **string** | Customer ID to charge | 
**PaymentMethod** | Pointer to **string** | Payment method ID. If not provided, status will be requires_payment_method. | [optional] 
**CaptureMethod** | Pointer to **string** | Capture method. Defaults to automatic. | [optional] [default to "automatic"]
**Description** | Pointer to **string** | Description for merchant reference | [optional] 
**StatementDescriptor** | Pointer to **string** | Statement descriptor shown on customer&#39;s statement | [optional] 
**StatementDescriptorSuffix** | Pointer to **string** | Statement descriptor suffix | [optional] 
**ReceiptEmail** | Pointer to **string** | Email to send receipt to | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**GatewayMerchantId** | Pointer to **string** | Specific gateway merchant ID for multi-MID setups (NMI) | [optional] 

## Methods

### NewCreatePaymentIntentRequest

`func NewCreatePaymentIntentRequest(amount int32, customer string, ) *CreatePaymentIntentRequest`

NewCreatePaymentIntentRequest instantiates a new CreatePaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentIntentRequestWithDefaults

`func NewCreatePaymentIntentRequestWithDefaults() *CreatePaymentIntentRequest`

NewCreatePaymentIntentRequestWithDefaults instantiates a new CreatePaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreatePaymentIntentRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentIntentRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentIntentRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *CreatePaymentIntentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentIntentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentIntentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentIntentRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *CreatePaymentIntentRequest) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CreatePaymentIntentRequest) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CreatePaymentIntentRequest) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetPaymentMethod

`func (o *CreatePaymentIntentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CreatePaymentIntentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CreatePaymentIntentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CreatePaymentIntentRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCaptureMethod

`func (o *CreatePaymentIntentRequest) GetCaptureMethod() string`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *CreatePaymentIntentRequest) GetCaptureMethodOk() (*string, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *CreatePaymentIntentRequest) SetCaptureMethod(v string)`

SetCaptureMethod sets CaptureMethod field to given value.

### HasCaptureMethod

`func (o *CreatePaymentIntentRequest) HasCaptureMethod() bool`

HasCaptureMethod returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePaymentIntentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePaymentIntentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePaymentIntentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePaymentIntentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *CreatePaymentIntentRequest) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *CreatePaymentIntentRequest) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *CreatePaymentIntentRequest) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *CreatePaymentIntentRequest) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetStatementDescriptorSuffix

`func (o *CreatePaymentIntentRequest) GetStatementDescriptorSuffix() string`

GetStatementDescriptorSuffix returns the StatementDescriptorSuffix field if non-nil, zero value otherwise.

### GetStatementDescriptorSuffixOk

`func (o *CreatePaymentIntentRequest) GetStatementDescriptorSuffixOk() (*string, bool)`

GetStatementDescriptorSuffixOk returns a tuple with the StatementDescriptorSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptorSuffix

`func (o *CreatePaymentIntentRequest) SetStatementDescriptorSuffix(v string)`

SetStatementDescriptorSuffix sets StatementDescriptorSuffix field to given value.

### HasStatementDescriptorSuffix

`func (o *CreatePaymentIntentRequest) HasStatementDescriptorSuffix() bool`

HasStatementDescriptorSuffix returns a boolean if a field has been set.

### GetReceiptEmail

`func (o *CreatePaymentIntentRequest) GetReceiptEmail() string`

GetReceiptEmail returns the ReceiptEmail field if non-nil, zero value otherwise.

### GetReceiptEmailOk

`func (o *CreatePaymentIntentRequest) GetReceiptEmailOk() (*string, bool)`

GetReceiptEmailOk returns a tuple with the ReceiptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptEmail

`func (o *CreatePaymentIntentRequest) SetReceiptEmail(v string)`

SetReceiptEmail sets ReceiptEmail field to given value.

### HasReceiptEmail

`func (o *CreatePaymentIntentRequest) HasReceiptEmail() bool`

HasReceiptEmail returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePaymentIntentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePaymentIntentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePaymentIntentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePaymentIntentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetGatewayMerchantId

`func (o *CreatePaymentIntentRequest) GetGatewayMerchantId() string`

GetGatewayMerchantId returns the GatewayMerchantId field if non-nil, zero value otherwise.

### GetGatewayMerchantIdOk

`func (o *CreatePaymentIntentRequest) GetGatewayMerchantIdOk() (*string, bool)`

GetGatewayMerchantIdOk returns a tuple with the GatewayMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMerchantId

`func (o *CreatePaymentIntentRequest) SetGatewayMerchantId(v string)`

SetGatewayMerchantId sets GatewayMerchantId field to given value.

### HasGatewayMerchantId

`func (o *CreatePaymentIntentRequest) HasGatewayMerchantId() bool`

HasGatewayMerchantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


