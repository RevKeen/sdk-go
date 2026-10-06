# RefundTerminalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMinor** | Pointer to **int64** | Amount to refund in minor units. Omit for a full refund. | [optional]
**Reason** | Pointer to **string** | Reason for the refund | [optional]
**RestockLineItems** | Pointer to [**[]CreateRefundInputRestockLineItemsInner**](CreateRefundInputRestockLineItemsInner.md) | Units to put back on sale once the refund is approved, by invoice line. Omit to restock nothing. | [optional]

## Methods

### NewRefundTerminalPaymentRequest

`func NewRefundTerminalPaymentRequest() *RefundTerminalPaymentRequest`

NewRefundTerminalPaymentRequest instantiates a new RefundTerminalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundTerminalPaymentRequestWithDefaults

`func NewRefundTerminalPaymentRequestWithDefaults() *RefundTerminalPaymentRequest`

NewRefundTerminalPaymentRequestWithDefaults instantiates a new RefundTerminalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountMinor

`func (o *RefundTerminalPaymentRequest) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *RefundTerminalPaymentRequest) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *RefundTerminalPaymentRequest) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *RefundTerminalPaymentRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetReason

`func (o *RefundTerminalPaymentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RefundTerminalPaymentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RefundTerminalPaymentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RefundTerminalPaymentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRestockLineItems

`func (o *RefundTerminalPaymentRequest) GetRestockLineItems() []CreateRefundInputRestockLineItemsInner`

GetRestockLineItems returns the RestockLineItems field if non-nil, zero value otherwise.

### GetRestockLineItemsOk

`func (o *RefundTerminalPaymentRequest) GetRestockLineItemsOk() (*[]CreateRefundInputRestockLineItemsInner, bool)`

GetRestockLineItemsOk returns a tuple with the RestockLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockLineItems

`func (o *RefundTerminalPaymentRequest) SetRestockLineItems(v []CreateRefundInputRestockLineItemsInner)`

SetRestockLineItems sets RestockLineItems field to given value.

### HasRestockLineItems

`func (o *RefundTerminalPaymentRequest) HasRestockLineItems() bool`

HasRestockLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
