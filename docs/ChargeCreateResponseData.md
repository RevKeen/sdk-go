# ChargeCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**PaymentMethodId** | **NullableString** |  | 
**AmountMinor** | **int32** |  | 
**AmountCapturedMinor** | **int32** |  | 
**AmountRefundedMinor** | **int32** |  | 
**Currency** | **string** |  | 
**Status** | **string** |  | 
**Description** | **string** |  | 
**StatementDescriptor** | **NullableString** |  | 
**FailureCode** | **NullableString** |  | 
**FailureMessage** | **NullableString** |  | 
**GatewayTransactionId** | **NullableString** |  | 
**ReceiptUrl** | **NullableString** |  | 
**Captured** | **bool** |  | 
**Metadata** | **map[string]interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewChargeCreateResponseData

`func NewChargeCreateResponseData(id string, merchantId string, customerId string, invoiceId NullableString, paymentMethodId NullableString, amountMinor int32, amountCapturedMinor int32, amountRefundedMinor int32, currency string, status string, description string, statementDescriptor NullableString, failureCode NullableString, failureMessage NullableString, gatewayTransactionId NullableString, receiptUrl NullableString, captured bool, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *ChargeCreateResponseData`

NewChargeCreateResponseData instantiates a new ChargeCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeCreateResponseDataWithDefaults

`func NewChargeCreateResponseDataWithDefaults() *ChargeCreateResponseData`

NewChargeCreateResponseDataWithDefaults instantiates a new ChargeCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *ChargeCreateResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ChargeCreateResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ChargeCreateResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *ChargeCreateResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargeCreateResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargeCreateResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *ChargeCreateResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ChargeCreateResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ChargeCreateResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *ChargeCreateResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ChargeCreateResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetPaymentMethodId

`func (o *ChargeCreateResponseData) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ChargeCreateResponseData) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ChargeCreateResponseData) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### SetPaymentMethodIdNil

`func (o *ChargeCreateResponseData) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *ChargeCreateResponseData) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetAmountMinor

`func (o *ChargeCreateResponseData) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *ChargeCreateResponseData) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *ChargeCreateResponseData) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetAmountCapturedMinor

`func (o *ChargeCreateResponseData) GetAmountCapturedMinor() int32`

GetAmountCapturedMinor returns the AmountCapturedMinor field if non-nil, zero value otherwise.

### GetAmountCapturedMinorOk

`func (o *ChargeCreateResponseData) GetAmountCapturedMinorOk() (*int32, bool)`

GetAmountCapturedMinorOk returns a tuple with the AmountCapturedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCapturedMinor

`func (o *ChargeCreateResponseData) SetAmountCapturedMinor(v int32)`

SetAmountCapturedMinor sets AmountCapturedMinor field to given value.


### GetAmountRefundedMinor

`func (o *ChargeCreateResponseData) GetAmountRefundedMinor() int32`

GetAmountRefundedMinor returns the AmountRefundedMinor field if non-nil, zero value otherwise.

### GetAmountRefundedMinorOk

`func (o *ChargeCreateResponseData) GetAmountRefundedMinorOk() (*int32, bool)`

GetAmountRefundedMinorOk returns a tuple with the AmountRefundedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefundedMinor

`func (o *ChargeCreateResponseData) SetAmountRefundedMinor(v int32)`

SetAmountRefundedMinor sets AmountRefundedMinor field to given value.


### GetCurrency

`func (o *ChargeCreateResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargeCreateResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargeCreateResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *ChargeCreateResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeCreateResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeCreateResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ChargeCreateResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeCreateResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeCreateResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatementDescriptor

`func (o *ChargeCreateResponseData) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *ChargeCreateResponseData) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *ChargeCreateResponseData) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.


### SetStatementDescriptorNil

`func (o *ChargeCreateResponseData) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *ChargeCreateResponseData) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetFailureCode

`func (o *ChargeCreateResponseData) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargeCreateResponseData) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargeCreateResponseData) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *ChargeCreateResponseData) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *ChargeCreateResponseData) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureMessage

`func (o *ChargeCreateResponseData) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargeCreateResponseData) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargeCreateResponseData) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.


### SetFailureMessageNil

`func (o *ChargeCreateResponseData) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *ChargeCreateResponseData) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetGatewayTransactionId

`func (o *ChargeCreateResponseData) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *ChargeCreateResponseData) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *ChargeCreateResponseData) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.


### SetGatewayTransactionIdNil

`func (o *ChargeCreateResponseData) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *ChargeCreateResponseData) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetReceiptUrl

`func (o *ChargeCreateResponseData) GetReceiptUrl() string`

GetReceiptUrl returns the ReceiptUrl field if non-nil, zero value otherwise.

### GetReceiptUrlOk

`func (o *ChargeCreateResponseData) GetReceiptUrlOk() (*string, bool)`

GetReceiptUrlOk returns a tuple with the ReceiptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptUrl

`func (o *ChargeCreateResponseData) SetReceiptUrl(v string)`

SetReceiptUrl sets ReceiptUrl field to given value.


### SetReceiptUrlNil

`func (o *ChargeCreateResponseData) SetReceiptUrlNil(b bool)`

 SetReceiptUrlNil sets the value for ReceiptUrl to be an explicit nil

### UnsetReceiptUrl
`func (o *ChargeCreateResponseData) UnsetReceiptUrl()`

UnsetReceiptUrl ensures that no value is present for ReceiptUrl, not even an explicit nil
### GetCaptured

`func (o *ChargeCreateResponseData) GetCaptured() bool`

GetCaptured returns the Captured field if non-nil, zero value otherwise.

### GetCapturedOk

`func (o *ChargeCreateResponseData) GetCapturedOk() (*bool, bool)`

GetCapturedOk returns a tuple with the Captured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptured

`func (o *ChargeCreateResponseData) SetCaptured(v bool)`

SetCaptured sets Captured field to given value.


### GetMetadata

`func (o *ChargeCreateResponseData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChargeCreateResponseData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChargeCreateResponseData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *ChargeCreateResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeCreateResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeCreateResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChargeCreateResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChargeCreateResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChargeCreateResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


