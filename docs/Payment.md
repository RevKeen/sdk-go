# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the payment | 
**Object** | **string** | Object type | 
**InvoiceId** | Pointer to **NullableString** | Associated invoice ID | [optional] 
**CustomerId** | Pointer to **NullableString** | Customer ID | [optional] 
**Status** | **string** | Payment status. pending: Processing. authorized: Funds reserved. captured/succeeded: Funds collected. failed: Payment failed. voided: Canceled before capture. refunded/partially_refunded: Returned to customer. | 
**Currency** | **string** | Three-letter ISO currency code | 
**Amount** | **int32** | Payment amount in cents | 
**AmountCaptured** | Pointer to **NullableInt32** | Amount captured in cents (for auth-capture flow) | [optional] 
**AmountRefunded** | Pointer to **int32** | Amount refunded in cents | [optional] [default to 0]
**Gateway** | Pointer to **NullableString** | Name of the payment processor that handled this payment | [optional] 
**GatewayTransactionId** | Pointer to **NullableString** | Processor&#39;s transaction reference | [optional] 
**GatewayResponseCode** | Pointer to **NullableString** | Normalized response code from the processor | [optional] 
**GatewayResponseText** | Pointer to **NullableString** | Human-readable response message from the processor | [optional] 
**PaymentMethodType** | Pointer to **NullableString** | Payment method type (card, ach, wallet) | [optional] 
**CardBrand** | Pointer to **NullableString** | Card brand (visa, mastercard, etc.) | [optional] 
**CardLastFour** | Pointer to **NullableString** | Last 4 digits of card | [optional] 
**PaymentChannel** | Pointer to **NullableString** | Payment channel. card_present: terminal/POS payment. card_not_present: online or recurring payment. bank_transfer: ACH/direct debit. manual: manually recorded. | [optional] 
**EntryMode** | Pointer to **NullableString** | Card entry mode for card-present (terminal) payments. null for online payments. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AuthorizedAt** | Pointer to **NullableTime** |  | [optional] 
**CapturedAt** | Pointer to **NullableTime** |  | [optional] 
**VoidedAt** | Pointer to **NullableTime** |  | [optional] 
**RefundedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewPayment

`func NewPayment(id string, object string, status string, currency string, amount int32, createdAt time.Time, updatedAt time.Time, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Payment) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Payment) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Payment) SetObject(v string)`

SetObject sets Object field to given value.


### GetInvoiceId

`func (o *Payment) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Payment) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Payment) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Payment) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *Payment) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *Payment) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetCustomerId

`func (o *Payment) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Payment) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Payment) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Payment) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Payment) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Payment) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCurrency

`func (o *Payment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *Payment) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountCaptured

`func (o *Payment) GetAmountCaptured() int32`

GetAmountCaptured returns the AmountCaptured field if non-nil, zero value otherwise.

### GetAmountCapturedOk

`func (o *Payment) GetAmountCapturedOk() (*int32, bool)`

GetAmountCapturedOk returns a tuple with the AmountCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCaptured

`func (o *Payment) SetAmountCaptured(v int32)`

SetAmountCaptured sets AmountCaptured field to given value.

### HasAmountCaptured

`func (o *Payment) HasAmountCaptured() bool`

HasAmountCaptured returns a boolean if a field has been set.

### SetAmountCapturedNil

`func (o *Payment) SetAmountCapturedNil(b bool)`

 SetAmountCapturedNil sets the value for AmountCaptured to be an explicit nil

### UnsetAmountCaptured
`func (o *Payment) UnsetAmountCaptured()`

UnsetAmountCaptured ensures that no value is present for AmountCaptured, not even an explicit nil
### GetAmountRefunded

`func (o *Payment) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *Payment) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *Payment) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.

### HasAmountRefunded

`func (o *Payment) HasAmountRefunded() bool`

HasAmountRefunded returns a boolean if a field has been set.

### GetGateway

`func (o *Payment) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Payment) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Payment) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Payment) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Payment) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Payment) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGatewayTransactionId

`func (o *Payment) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *Payment) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *Payment) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *Payment) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *Payment) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *Payment) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetGatewayResponseCode

`func (o *Payment) GetGatewayResponseCode() string`

GetGatewayResponseCode returns the GatewayResponseCode field if non-nil, zero value otherwise.

### GetGatewayResponseCodeOk

`func (o *Payment) GetGatewayResponseCodeOk() (*string, bool)`

GetGatewayResponseCodeOk returns a tuple with the GatewayResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseCode

`func (o *Payment) SetGatewayResponseCode(v string)`

SetGatewayResponseCode sets GatewayResponseCode field to given value.

### HasGatewayResponseCode

`func (o *Payment) HasGatewayResponseCode() bool`

HasGatewayResponseCode returns a boolean if a field has been set.

### SetGatewayResponseCodeNil

`func (o *Payment) SetGatewayResponseCodeNil(b bool)`

 SetGatewayResponseCodeNil sets the value for GatewayResponseCode to be an explicit nil

### UnsetGatewayResponseCode
`func (o *Payment) UnsetGatewayResponseCode()`

UnsetGatewayResponseCode ensures that no value is present for GatewayResponseCode, not even an explicit nil
### GetGatewayResponseText

`func (o *Payment) GetGatewayResponseText() string`

GetGatewayResponseText returns the GatewayResponseText field if non-nil, zero value otherwise.

### GetGatewayResponseTextOk

`func (o *Payment) GetGatewayResponseTextOk() (*string, bool)`

GetGatewayResponseTextOk returns a tuple with the GatewayResponseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseText

`func (o *Payment) SetGatewayResponseText(v string)`

SetGatewayResponseText sets GatewayResponseText field to given value.

### HasGatewayResponseText

`func (o *Payment) HasGatewayResponseText() bool`

HasGatewayResponseText returns a boolean if a field has been set.

### SetGatewayResponseTextNil

`func (o *Payment) SetGatewayResponseTextNil(b bool)`

 SetGatewayResponseTextNil sets the value for GatewayResponseText to be an explicit nil

### UnsetGatewayResponseText
`func (o *Payment) UnsetGatewayResponseText()`

UnsetGatewayResponseText ensures that no value is present for GatewayResponseText, not even an explicit nil
### GetPaymentMethodType

`func (o *Payment) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *Payment) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *Payment) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *Payment) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### SetPaymentMethodTypeNil

`func (o *Payment) SetPaymentMethodTypeNil(b bool)`

 SetPaymentMethodTypeNil sets the value for PaymentMethodType to be an explicit nil

### UnsetPaymentMethodType
`func (o *Payment) UnsetPaymentMethodType()`

UnsetPaymentMethodType ensures that no value is present for PaymentMethodType, not even an explicit nil
### GetCardBrand

`func (o *Payment) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *Payment) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *Payment) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *Payment) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *Payment) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *Payment) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLastFour

`func (o *Payment) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *Payment) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *Payment) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *Payment) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *Payment) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *Payment) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetPaymentChannel

`func (o *Payment) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *Payment) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *Payment) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *Payment) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### SetPaymentChannelNil

`func (o *Payment) SetPaymentChannelNil(b bool)`

 SetPaymentChannelNil sets the value for PaymentChannel to be an explicit nil

### UnsetPaymentChannel
`func (o *Payment) UnsetPaymentChannel()`

UnsetPaymentChannel ensures that no value is present for PaymentChannel, not even an explicit nil
### GetEntryMode

`func (o *Payment) GetEntryMode() string`

GetEntryMode returns the EntryMode field if non-nil, zero value otherwise.

### GetEntryModeOk

`func (o *Payment) GetEntryModeOk() (*string, bool)`

GetEntryModeOk returns a tuple with the EntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMode

`func (o *Payment) SetEntryMode(v string)`

SetEntryMode sets EntryMode field to given value.

### HasEntryMode

`func (o *Payment) HasEntryMode() bool`

HasEntryMode returns a boolean if a field has been set.

### SetEntryModeNil

`func (o *Payment) SetEntryModeNil(b bool)`

 SetEntryModeNil sets the value for EntryMode to be an explicit nil

### UnsetEntryMode
`func (o *Payment) UnsetEntryMode()`

UnsetEntryMode ensures that no value is present for EntryMode, not even an explicit nil
### GetMetadata

`func (o *Payment) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Payment) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Payment) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Payment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Payment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Payment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Payment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Payment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Payment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Payment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAuthorizedAt

`func (o *Payment) GetAuthorizedAt() time.Time`

GetAuthorizedAt returns the AuthorizedAt field if non-nil, zero value otherwise.

### GetAuthorizedAtOk

`func (o *Payment) GetAuthorizedAtOk() (*time.Time, bool)`

GetAuthorizedAtOk returns a tuple with the AuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAt

`func (o *Payment) SetAuthorizedAt(v time.Time)`

SetAuthorizedAt sets AuthorizedAt field to given value.

### HasAuthorizedAt

`func (o *Payment) HasAuthorizedAt() bool`

HasAuthorizedAt returns a boolean if a field has been set.

### SetAuthorizedAtNil

`func (o *Payment) SetAuthorizedAtNil(b bool)`

 SetAuthorizedAtNil sets the value for AuthorizedAt to be an explicit nil

### UnsetAuthorizedAt
`func (o *Payment) UnsetAuthorizedAt()`

UnsetAuthorizedAt ensures that no value is present for AuthorizedAt, not even an explicit nil
### GetCapturedAt

`func (o *Payment) GetCapturedAt() time.Time`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *Payment) GetCapturedAtOk() (*time.Time, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *Payment) SetCapturedAt(v time.Time)`

SetCapturedAt sets CapturedAt field to given value.

### HasCapturedAt

`func (o *Payment) HasCapturedAt() bool`

HasCapturedAt returns a boolean if a field has been set.

### SetCapturedAtNil

`func (o *Payment) SetCapturedAtNil(b bool)`

 SetCapturedAtNil sets the value for CapturedAt to be an explicit nil

### UnsetCapturedAt
`func (o *Payment) UnsetCapturedAt()`

UnsetCapturedAt ensures that no value is present for CapturedAt, not even an explicit nil
### GetVoidedAt

`func (o *Payment) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *Payment) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *Payment) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *Payment) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### SetVoidedAtNil

`func (o *Payment) SetVoidedAtNil(b bool)`

 SetVoidedAtNil sets the value for VoidedAt to be an explicit nil

### UnsetVoidedAt
`func (o *Payment) UnsetVoidedAt()`

UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil
### GetRefundedAt

`func (o *Payment) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *Payment) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *Payment) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *Payment) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *Payment) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *Payment) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


