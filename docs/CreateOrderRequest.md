# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** | Customer ID (optional for guest orders) | [optional] 
**Currency** | Pointer to **string** | Three-letter ISO currency code | [optional] [default to "usd"]
**BillingType** | Pointer to **string** | Billing type. one_time: Single purchase. recurring: Subscription-based. | [optional] [default to "one_time"]
**LineItems** | [**[]CreateOrderRequestLineItemsInner**](CreateOrderRequestLineItemsInner.md) | Order line items (at least one required) | 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Expiration timestamp for the order | [optional] 
**InvoiceId** | Pointer to **string** | ID of the linked invoice (Commercial Truth) | [optional] 
**SubscriptionId** | Pointer to **string** | ID of the linked subscription (for recurring fulfillment) | [optional] 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(lineItems []CreateOrderRequestLineItemsInner, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateOrderRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateOrderRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateOrderRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreateOrderRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCurrency

`func (o *CreateOrderRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateOrderRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateOrderRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateOrderRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetBillingType

`func (o *CreateOrderRequest) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *CreateOrderRequest) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *CreateOrderRequest) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *CreateOrderRequest) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetLineItems

`func (o *CreateOrderRequest) GetLineItems() []CreateOrderRequestLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreateOrderRequest) GetLineItemsOk() (*[]CreateOrderRequestLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreateOrderRequest) SetLineItems(v []CreateOrderRequestLineItemsInner)`

SetLineItems sets LineItems field to given value.


### GetMetadata

`func (o *CreateOrderRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateOrderRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateOrderRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateOrderRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotes

`func (o *CreateOrderRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateOrderRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateOrderRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateOrderRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateOrderRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateOrderRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateOrderRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateOrderRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreateOrderRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateOrderRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateOrderRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreateOrderRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreateOrderRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateOrderRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateOrderRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateOrderRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


