# Refund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PublicId** | **string** |  | 
**PaymentId** | **string** | The original payment ID this refund is for (alias for parent_transaction_id) | 
**ParentTransactionId** | **string** | Parent transaction ID in unified transaction model. Same as payment_id for refunds. | 
**OrderId** | **NullableString** |  | 
**Gateway** | **string** | Name of the payment processor that handled this refund | 
**GatewayRefundId** | **NullableString** |  | 
**AmountMinor** | **float32** | Refund amount in cents | 
**Currency** | **string** |  | 
**Reason** | **NullableString** |  | 
**ReasonDetails** | **NullableString** |  | 
**Status** | **string** |  | 
**FailureReason** | **NullableString** |  | 
**FailureCode** | **NullableString** |  | 
**ProcessedAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewRefund

`func NewRefund(id string, publicId string, paymentId string, parentTransactionId string, orderId NullableString, gateway string, gatewayRefundId NullableString, amountMinor float32, currency string, reason NullableString, reasonDetails NullableString, status string, failureReason NullableString, failureCode NullableString, processedAt NullableString, createdAt string, ) *Refund`

NewRefund instantiates a new Refund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundWithDefaults

`func NewRefundWithDefaults() *Refund`

NewRefundWithDefaults instantiates a new Refund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Refund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Refund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Refund) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *Refund) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Refund) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Refund) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetPaymentId

`func (o *Refund) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Refund) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Refund) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetParentTransactionId

`func (o *Refund) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *Refund) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *Refund) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### GetOrderId

`func (o *Refund) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Refund) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Refund) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Refund) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Refund) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetGateway

`func (o *Refund) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Refund) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Refund) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayRefundId

`func (o *Refund) GetGatewayRefundId() string`

GetGatewayRefundId returns the GatewayRefundId field if non-nil, zero value otherwise.

### GetGatewayRefundIdOk

`func (o *Refund) GetGatewayRefundIdOk() (*string, bool)`

GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRefundId

`func (o *Refund) SetGatewayRefundId(v string)`

SetGatewayRefundId sets GatewayRefundId field to given value.


### SetGatewayRefundIdNil

`func (o *Refund) SetGatewayRefundIdNil(b bool)`

 SetGatewayRefundIdNil sets the value for GatewayRefundId to be an explicit nil

### UnsetGatewayRefundId
`func (o *Refund) UnsetGatewayRefundId()`

UnsetGatewayRefundId ensures that no value is present for GatewayRefundId, not even an explicit nil
### GetAmountMinor

`func (o *Refund) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *Refund) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *Refund) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *Refund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Refund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Refund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReason

`func (o *Refund) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Refund) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Refund) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *Refund) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Refund) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonDetails

`func (o *Refund) GetReasonDetails() string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *Refund) GetReasonDetailsOk() (*string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *Refund) SetReasonDetails(v string)`

SetReasonDetails sets ReasonDetails field to given value.


### SetReasonDetailsNil

`func (o *Refund) SetReasonDetailsNil(b bool)`

 SetReasonDetailsNil sets the value for ReasonDetails to be an explicit nil

### UnsetReasonDetails
`func (o *Refund) UnsetReasonDetails()`

UnsetReasonDetails ensures that no value is present for ReasonDetails, not even an explicit nil
### GetStatus

`func (o *Refund) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Refund) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Refund) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFailureReason

`func (o *Refund) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *Refund) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *Refund) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *Refund) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *Refund) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetFailureCode

`func (o *Refund) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Refund) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Refund) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *Refund) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *Refund) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetProcessedAt

`func (o *Refund) GetProcessedAt() string`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *Refund) GetProcessedAtOk() (*string, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *Refund) SetProcessedAt(v string)`

SetProcessedAt sets ProcessedAt field to given value.


### SetProcessedAtNil

`func (o *Refund) SetProcessedAtNil(b bool)`

 SetProcessedAtNil sets the value for ProcessedAt to be an explicit nil

### UnsetProcessedAt
`func (o *Refund) UnsetProcessedAt()`

UnsetProcessedAt ensures that no value is present for ProcessedAt, not even an explicit nil
### GetCreatedAt

`func (o *Refund) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Refund) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Refund) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


