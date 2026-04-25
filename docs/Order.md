# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the order | 
**Object** | **string** | Object type | 
**PublicId** | **string** | Public-facing order ID | 
**CustomerId** | Pointer to **NullableString** | ID of the customer | [optional] 
**Status** | **string** | Order status. draft: Being built. pending: Awaiting payment. paid: Fully paid. partially_paid: Partial payment received. refunded: Refunded. canceled: Canceled. fulfilled: Shipped/delivered. | 
**BillingType** | **string** | Billing type. one_time: Single purchase. recurring: Subscription-based. | 
**FulfillmentStatus** | Pointer to **NullableString** | Fulfillment status for physical goods. | [optional] 
**Currency** | **string** | Three-letter ISO currency code | 
**Subtotal** | **int32** | Subtotal in cents | 
**Discount** | **int32** | Total discount in cents | 
**Tax** | **int32** | Total tax in cents | 
**Total** | **int32** | Total amount in cents | 
**AmountPaid** | **int32** | Amount paid in cents | 
**AmountRefunded** | **int32** | Amount refunded in cents | 
**AmountDue** | **int32** | Amount still due in cents | 
**LineItems** | Pointer to [**[]OrderLineItem**](OrderLineItem.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**InvoiceId** | Pointer to **NullableString** | ID of the linked invoice (Commercial Truth) | [optional] 
**SubscriptionId** | Pointer to **NullableString** | ID of the linked subscription (for recurring fulfillment) | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**PaidAt** | Pointer to **NullableTime** |  | [optional] 
**CanceledAt** | Pointer to **NullableTime** |  | [optional] 
**FulfilledAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewOrder

`func NewOrder(id string, object string, publicId string, status string, billingType string, currency string, subtotal int32, discount int32, tax int32, total int32, amountPaid int32, amountRefunded int32, amountDue int32, createdAt time.Time, updatedAt time.Time, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Order) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Order) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Order) SetObject(v string)`

SetObject sets Object field to given value.


### GetPublicId

`func (o *Order) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Order) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Order) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetCustomerId

`func (o *Order) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Order) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Order) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Order) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Order) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Order) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBillingType

`func (o *Order) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *Order) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *Order) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.


### GetFulfillmentStatus

`func (o *Order) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *Order) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *Order) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *Order) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### SetFulfillmentStatusNil

`func (o *Order) SetFulfillmentStatusNil(b bool)`

 SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil

### UnsetFulfillmentStatus
`func (o *Order) UnsetFulfillmentStatus()`

UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
### GetCurrency

`func (o *Order) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSubtotal

`func (o *Order) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *Order) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *Order) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetDiscount

`func (o *Order) GetDiscount() int32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Order) GetDiscountOk() (*int32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Order) SetDiscount(v int32)`

SetDiscount sets Discount field to given value.


### GetTax

`func (o *Order) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *Order) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *Order) SetTax(v int32)`

SetTax sets Tax field to given value.


### GetTotal

`func (o *Order) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Order) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Order) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetAmountPaid

`func (o *Order) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *Order) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *Order) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.


### GetAmountRefunded

`func (o *Order) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *Order) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *Order) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.


### GetAmountDue

`func (o *Order) GetAmountDue() int32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *Order) GetAmountDueOk() (*int32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *Order) SetAmountDue(v int32)`

SetAmountDue sets AmountDue field to given value.


### GetLineItems

`func (o *Order) GetLineItems() []OrderLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Order) GetLineItemsOk() (*[]OrderLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Order) SetLineItems(v []OrderLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Order) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMetadata

`func (o *Order) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Order) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Order) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Order) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *Order) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Order) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Order) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Order) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Order) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Order) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInvoiceId

`func (o *Order) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Order) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Order) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Order) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *Order) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *Order) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionId

`func (o *Order) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Order) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Order) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Order) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *Order) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *Order) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPaidAt

`func (o *Order) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *Order) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *Order) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *Order) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *Order) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *Order) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetCanceledAt

`func (o *Order) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Order) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Order) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Order) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *Order) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *Order) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetFulfilledAt

`func (o *Order) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *Order) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *Order) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.

### HasFulfilledAt

`func (o *Order) HasFulfilledAt() bool`

HasFulfilledAt returns a boolean if a field has been set.

### SetFulfilledAtNil

`func (o *Order) SetFulfilledAtNil(b bool)`

 SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil

### UnsetFulfilledAt
`func (o *Order) UnsetFulfilledAt()`

UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil
### GetExpiresAt

`func (o *Order) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Order) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Order) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Order) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *Order) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Order) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


