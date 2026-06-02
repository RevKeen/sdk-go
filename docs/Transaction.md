# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the transaction | 
**Object** | **string** | Object type | 
**Type** | **string** | Transaction type. sale: original payment. refund: money returned. void: cancel before settlement. capture: capture a previous auth. dispute: chargeback. adjustment: manual correction. | 
**Status** | **string** | Transaction status. pending: processing. succeeded: completed. failed: declined or errored. voided: canceled. | 
**InvoiceId** | **string** | Associated invoice ID | 
**CustomerId** | Pointer to **NullableString** | Customer ID | [optional] 
**ParentTransactionId** | Pointer to **NullableString** | Parent transaction ID (required for refunds, voids, disputes, adjustments) | [optional] 
**Currency** | **string** | Three-letter ISO currency code | 
**Amount** | **int32** | Transaction amount in minor units (cents) | 
**AmountRefunded** | Pointer to **int32** | Amount refunded in minor units | [optional] [default to 0]
**Gateway** | Pointer to **NullableString** | Payment gateway (e.g., &#39;nmi&#39;) | [optional] 
**GatewayTransactionId** | Pointer to **NullableString** | Gateway&#39;s transaction reference | [optional] 
**GatewayResponseCode** | Pointer to **NullableString** | Gateway response code | [optional] 
**GatewayResponseText** | Pointer to **NullableString** | Gateway response message | [optional] 
**PaymentMethodType** | Pointer to **NullableString** | Payment method type (card, ach, wallet) | [optional] 
**CardBrand** | Pointer to **NullableString** | Card brand (visa, mastercard, etc.) | [optional] 
**CardLastFour** | Pointer to **NullableString** | Last 4 digits of card | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp (ISO 8601) | 
**UpdatedAt** | **time.Time** | Last update timestamp (ISO 8601) | 

## Methods

### NewTransaction

`func NewTransaction(id string, object string, type_ string, status string, invoiceId string, currency string, amount int32, createdAt time.Time, updatedAt time.Time, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Transaction) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Transaction) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Transaction) SetObject(v string)`

SetObject sets Object field to given value.


### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInvoiceId

`func (o *Transaction) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Transaction) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Transaction) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetCustomerId

`func (o *Transaction) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Transaction) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Transaction) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Transaction) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Transaction) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Transaction) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetParentTransactionId

`func (o *Transaction) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *Transaction) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *Transaction) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.

### HasParentTransactionId

`func (o *Transaction) HasParentTransactionId() bool`

HasParentTransactionId returns a boolean if a field has been set.

### SetParentTransactionIdNil

`func (o *Transaction) SetParentTransactionIdNil(b bool)`

 SetParentTransactionIdNil sets the value for ParentTransactionId to be an explicit nil

### UnsetParentTransactionId
`func (o *Transaction) UnsetParentTransactionId()`

UnsetParentTransactionId ensures that no value is present for ParentTransactionId, not even an explicit nil
### GetCurrency

`func (o *Transaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *Transaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountRefunded

`func (o *Transaction) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *Transaction) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *Transaction) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.

### HasAmountRefunded

`func (o *Transaction) HasAmountRefunded() bool`

HasAmountRefunded returns a boolean if a field has been set.

### GetGateway

`func (o *Transaction) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Transaction) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Transaction) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *Transaction) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *Transaction) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *Transaction) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetGatewayTransactionId

`func (o *Transaction) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *Transaction) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *Transaction) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.

### HasGatewayTransactionId

`func (o *Transaction) HasGatewayTransactionId() bool`

HasGatewayTransactionId returns a boolean if a field has been set.

### SetGatewayTransactionIdNil

`func (o *Transaction) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *Transaction) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetGatewayResponseCode

`func (o *Transaction) GetGatewayResponseCode() string`

GetGatewayResponseCode returns the GatewayResponseCode field if non-nil, zero value otherwise.

### GetGatewayResponseCodeOk

`func (o *Transaction) GetGatewayResponseCodeOk() (*string, bool)`

GetGatewayResponseCodeOk returns a tuple with the GatewayResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseCode

`func (o *Transaction) SetGatewayResponseCode(v string)`

SetGatewayResponseCode sets GatewayResponseCode field to given value.

### HasGatewayResponseCode

`func (o *Transaction) HasGatewayResponseCode() bool`

HasGatewayResponseCode returns a boolean if a field has been set.

### SetGatewayResponseCodeNil

`func (o *Transaction) SetGatewayResponseCodeNil(b bool)`

 SetGatewayResponseCodeNil sets the value for GatewayResponseCode to be an explicit nil

### UnsetGatewayResponseCode
`func (o *Transaction) UnsetGatewayResponseCode()`

UnsetGatewayResponseCode ensures that no value is present for GatewayResponseCode, not even an explicit nil
### GetGatewayResponseText

`func (o *Transaction) GetGatewayResponseText() string`

GetGatewayResponseText returns the GatewayResponseText field if non-nil, zero value otherwise.

### GetGatewayResponseTextOk

`func (o *Transaction) GetGatewayResponseTextOk() (*string, bool)`

GetGatewayResponseTextOk returns a tuple with the GatewayResponseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayResponseText

`func (o *Transaction) SetGatewayResponseText(v string)`

SetGatewayResponseText sets GatewayResponseText field to given value.

### HasGatewayResponseText

`func (o *Transaction) HasGatewayResponseText() bool`

HasGatewayResponseText returns a boolean if a field has been set.

### SetGatewayResponseTextNil

`func (o *Transaction) SetGatewayResponseTextNil(b bool)`

 SetGatewayResponseTextNil sets the value for GatewayResponseText to be an explicit nil

### UnsetGatewayResponseText
`func (o *Transaction) UnsetGatewayResponseText()`

UnsetGatewayResponseText ensures that no value is present for GatewayResponseText, not even an explicit nil
### GetPaymentMethodType

`func (o *Transaction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *Transaction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *Transaction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *Transaction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### SetPaymentMethodTypeNil

`func (o *Transaction) SetPaymentMethodTypeNil(b bool)`

 SetPaymentMethodTypeNil sets the value for PaymentMethodType to be an explicit nil

### UnsetPaymentMethodType
`func (o *Transaction) UnsetPaymentMethodType()`

UnsetPaymentMethodType ensures that no value is present for PaymentMethodType, not even an explicit nil
### GetCardBrand

`func (o *Transaction) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *Transaction) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *Transaction) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *Transaction) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### SetCardBrandNil

`func (o *Transaction) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *Transaction) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLastFour

`func (o *Transaction) GetCardLastFour() string`

GetCardLastFour returns the CardLastFour field if non-nil, zero value otherwise.

### GetCardLastFourOk

`func (o *Transaction) GetCardLastFourOk() (*string, bool)`

GetCardLastFourOk returns a tuple with the CardLastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLastFour

`func (o *Transaction) SetCardLastFour(v string)`

SetCardLastFour sets CardLastFour field to given value.

### HasCardLastFour

`func (o *Transaction) HasCardLastFour() bool`

HasCardLastFour returns a boolean if a field has been set.

### SetCardLastFourNil

`func (o *Transaction) SetCardLastFourNil(b bool)`

 SetCardLastFourNil sets the value for CardLastFour to be an explicit nil

### UnsetCardLastFour
`func (o *Transaction) UnsetCardLastFour()`

UnsetCardLastFour ensures that no value is present for CardLastFour, not even an explicit nil
### GetMetadata

`func (o *Transaction) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Transaction) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Transaction) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Transaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


