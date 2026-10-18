# ChargeListResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**CustomerId** | **string** |  |
**InvoiceId** | **NullableString** |  |
**PaymentMethodId** | **NullableString** |  |
**AmountMinor** | **int64** |  |
**AmountCapturedMinor** | **int64** |  |
**AmountRefundedMinor** | **int64** |  |
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

### NewChargeListResponseDataInner

`func NewChargeListResponseDataInner(id string, customerId string, invoiceId NullableString, paymentMethodId NullableString, amountMinor int64, amountCapturedMinor int64, amountRefundedMinor int64, currency string, status string, description string, statementDescriptor NullableString, failureCode NullableString, failureMessage NullableString, gatewayTransactionId NullableString, receiptUrl NullableString, captured bool, metadata map[string]*interface{}, createdAt time.Time, updatedAt time.Time, ) *ChargeListResponseDataInner`

NewChargeListResponseDataInner instantiates a new ChargeListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeListResponseDataInnerWithDefaults

`func NewChargeListResponseDataInnerWithDefaults() *ChargeListResponseDataInner`

NewChargeListResponseDataInnerWithDefaults instantiates a new ChargeListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChargeListResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChargeListResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChargeListResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *ChargeListResponseDataInner) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ChargeListResponseDataInner) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ChargeListResponseDataInner) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *ChargeListResponseDataInner) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ChargeListResponseDataInner) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ChargeListResponseDataInner) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *ChargeListResponseDataInner) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ChargeListResponseDataInner) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetPaymentMethodId

`func (o *ChargeListResponseDataInner) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ChargeListResponseDataInner) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ChargeListResponseDataInner) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### SetPaymentMethodIdNil

`func (o *ChargeListResponseDataInner) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *ChargeListResponseDataInner) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetAmountMinor

`func (o *ChargeListResponseDataInner) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *ChargeListResponseDataInner) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *ChargeListResponseDataInner) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetAmountCapturedMinor

`func (o *ChargeListResponseDataInner) GetAmountCapturedMinor() int64`

GetAmountCapturedMinor returns the AmountCapturedMinor field if non-nil, zero value otherwise.

### GetAmountCapturedMinorOk

`func (o *ChargeListResponseDataInner) GetAmountCapturedMinorOk() (*int64, bool)`

GetAmountCapturedMinorOk returns a tuple with the AmountCapturedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCapturedMinor

`func (o *ChargeListResponseDataInner) SetAmountCapturedMinor(v int64)`

SetAmountCapturedMinor sets AmountCapturedMinor field to given value.


### GetAmountRefundedMinor

`func (o *ChargeListResponseDataInner) GetAmountRefundedMinor() int64`

GetAmountRefundedMinor returns the AmountRefundedMinor field if non-nil, zero value otherwise.

### GetAmountRefundedMinorOk

`func (o *ChargeListResponseDataInner) GetAmountRefundedMinorOk() (*int64, bool)`

GetAmountRefundedMinorOk returns a tuple with the AmountRefundedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefundedMinor

`func (o *ChargeListResponseDataInner) SetAmountRefundedMinor(v int64)`

SetAmountRefundedMinor sets AmountRefundedMinor field to given value.


### GetCurrency

`func (o *ChargeListResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargeListResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargeListResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *ChargeListResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChargeListResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChargeListResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ChargeListResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChargeListResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChargeListResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatementDescriptor

`func (o *ChargeListResponseDataInner) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *ChargeListResponseDataInner) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *ChargeListResponseDataInner) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.


### SetStatementDescriptorNil

`func (o *ChargeListResponseDataInner) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *ChargeListResponseDataInner) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetFailureCode

`func (o *ChargeListResponseDataInner) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *ChargeListResponseDataInner) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *ChargeListResponseDataInner) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *ChargeListResponseDataInner) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *ChargeListResponseDataInner) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetFailureMessage

`func (o *ChargeListResponseDataInner) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ChargeListResponseDataInner) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ChargeListResponseDataInner) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.


### SetFailureMessageNil

`func (o *ChargeListResponseDataInner) SetFailureMessageNil(b bool)`

 SetFailureMessageNil sets the value for FailureMessage to be an explicit nil

### UnsetFailureMessage
`func (o *ChargeListResponseDataInner) UnsetFailureMessage()`

UnsetFailureMessage ensures that no value is present for FailureMessage, not even an explicit nil
### GetGatewayTransactionId

`func (o *ChargeListResponseDataInner) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *ChargeListResponseDataInner) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *ChargeListResponseDataInner) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.


### SetGatewayTransactionIdNil

`func (o *ChargeListResponseDataInner) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *ChargeListResponseDataInner) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetReceiptUrl

`func (o *ChargeListResponseDataInner) GetReceiptUrl() string`

GetReceiptUrl returns the ReceiptUrl field if non-nil, zero value otherwise.

### GetReceiptUrlOk

`func (o *ChargeListResponseDataInner) GetReceiptUrlOk() (*string, bool)`

GetReceiptUrlOk returns a tuple with the ReceiptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptUrl

`func (o *ChargeListResponseDataInner) SetReceiptUrl(v string)`

SetReceiptUrl sets ReceiptUrl field to given value.


### SetReceiptUrlNil

`func (o *ChargeListResponseDataInner) SetReceiptUrlNil(b bool)`

 SetReceiptUrlNil sets the value for ReceiptUrl to be an explicit nil

### UnsetReceiptUrl
`func (o *ChargeListResponseDataInner) UnsetReceiptUrl()`

UnsetReceiptUrl ensures that no value is present for ReceiptUrl, not even an explicit nil
### GetCaptured

`func (o *ChargeListResponseDataInner) GetCaptured() bool`

GetCaptured returns the Captured field if non-nil, zero value otherwise.

### GetCapturedOk

`func (o *ChargeListResponseDataInner) GetCapturedOk() (*bool, bool)`

GetCapturedOk returns a tuple with the Captured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptured

`func (o *ChargeListResponseDataInner) SetCaptured(v bool)`

SetCaptured sets Captured field to given value.


### GetMetadata

`func (o *ChargeListResponseDataInner) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChargeListResponseDataInner) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChargeListResponseDataInner) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *ChargeListResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChargeListResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChargeListResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ChargeListResponseDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChargeListResponseDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChargeListResponseDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
