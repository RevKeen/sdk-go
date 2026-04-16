# Void

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PublicId** | **string** |  | 
**PaymentId** | **string** | The original payment ID this void is for (alias for parent_transaction_id) | 
**ParentTransactionId** | **string** | Parent transaction ID in unified transaction model. | 
**OrderId** | **NullableString** |  | 
**Gateway** | **string** |  | 
**GatewayVoidId** | **NullableString** |  | 
**AmountMinor** | **float32** | Void amount in cents (always full amount of original transaction) | 
**Currency** | **string** |  | 
**Reason** | **NullableString** |  | 
**ReasonCode** | **NullableString** |  | 
**Status** | **string** |  | 
**FailureReason** | **NullableString** |  | 
**FailureCode** | **NullableString** |  | 
**VoidedAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewVoid

`func NewVoid(id string, publicId string, paymentId string, parentTransactionId string, orderId NullableString, gateway string, gatewayVoidId NullableString, amountMinor float32, currency string, reason NullableString, reasonCode NullableString, status string, failureReason NullableString, failureCode NullableString, voidedAt NullableString, createdAt string, ) *Void`

NewVoid instantiates a new Void object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidWithDefaults

`func NewVoidWithDefaults() *Void`

NewVoidWithDefaults instantiates a new Void object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Void) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Void) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Void) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *Void) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Void) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Void) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetPaymentId

`func (o *Void) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Void) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Void) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetParentTransactionId

`func (o *Void) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *Void) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *Void) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### GetOrderId

`func (o *Void) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Void) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Void) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Void) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Void) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetGateway

`func (o *Void) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Void) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Void) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayVoidId

`func (o *Void) GetGatewayVoidId() string`

GetGatewayVoidId returns the GatewayVoidId field if non-nil, zero value otherwise.

### GetGatewayVoidIdOk

`func (o *Void) GetGatewayVoidIdOk() (*string, bool)`

GetGatewayVoidIdOk returns a tuple with the GatewayVoidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVoidId

`func (o *Void) SetGatewayVoidId(v string)`

SetGatewayVoidId sets GatewayVoidId field to given value.


### SetGatewayVoidIdNil

`func (o *Void) SetGatewayVoidIdNil(b bool)`

 SetGatewayVoidIdNil sets the value for GatewayVoidId to be an explicit nil

### UnsetGatewayVoidId
`func (o *Void) UnsetGatewayVoidId()`

UnsetGatewayVoidId ensures that no value is present for GatewayVoidId, not even an explicit nil
### GetAmountMinor

`func (o *Void) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *Void) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *Void) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *Void) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Void) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Void) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReason

`func (o *Void) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Void) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Void) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *Void) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Void) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetReasonCode

`func (o *Void) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *Void) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *Void) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### SetReasonCodeNil

`func (o *Void) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *Void) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetStatus

`func (o *Void) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Void) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Void) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetFailureReason

`func (o *Void) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *Void) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *Void) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *Void) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *Void) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetFailureCode

`func (o *Void) GetFailureCode() string`

GetFailureCode returns the FailureCode field if non-nil, zero value otherwise.

### GetFailureCodeOk

`func (o *Void) GetFailureCodeOk() (*string, bool)`

GetFailureCodeOk returns a tuple with the FailureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCode

`func (o *Void) SetFailureCode(v string)`

SetFailureCode sets FailureCode field to given value.


### SetFailureCodeNil

`func (o *Void) SetFailureCodeNil(b bool)`

 SetFailureCodeNil sets the value for FailureCode to be an explicit nil

### UnsetFailureCode
`func (o *Void) UnsetFailureCode()`

UnsetFailureCode ensures that no value is present for FailureCode, not even an explicit nil
### GetVoidedAt

`func (o *Void) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *Void) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *Void) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.


### SetVoidedAtNil

`func (o *Void) SetVoidedAtNil(b bool)`

 SetVoidedAtNil sets the value for VoidedAt to be an explicit nil

### UnsetVoidedAt
`func (o *Void) UnsetVoidedAt()`

UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil
### GetCreatedAt

`func (o *Void) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Void) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Void) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


