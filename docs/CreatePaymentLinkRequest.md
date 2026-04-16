# CreatePaymentLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductIds** | Pointer to **[]string** |  | [optional] 
**LineItems** | Pointer to [**[]CreatePaymentLinkRequestLineItemsInner**](CreatePaymentLinkRequestLineItemsInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] [default to "custom"]
**AmountMinor** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**InvoiceId** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**SuccessUrl** | Pointer to **string** | Custom URL to redirect after successful payment. Supports placeholders: {CHECKOUT_ID}, {TRANSACTION_ID}, {PAYMENT_ID}, {CUSTOMER_ID}, {AMOUNT}, {CURRENCY}, {PORTAL_URL}. If not set, redirects to RevKeen&#39;s default success page. | [optional] 
**CancelUrl** | Pointer to **string** | Custom URL to redirect if customer cancels the checkout. If not set, no cancel redirect is configured. | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**CollectShipping** | Pointer to **bool** |  | [optional] [default to false]
**AllowPromotionCodes** | Pointer to **bool** |  | [optional] [default to true]
**TaxBehavior** | Pointer to **string** |  | [optional] [default to "unspecified"]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCreatePaymentLinkRequest

`func NewCreatePaymentLinkRequest() *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequest instantiates a new CreatePaymentLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentLinkRequestWithDefaults

`func NewCreatePaymentLinkRequestWithDefaults() *CreatePaymentLinkRequest`

NewCreatePaymentLinkRequestWithDefaults instantiates a new CreatePaymentLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductIds

`func (o *CreatePaymentLinkRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *CreatePaymentLinkRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *CreatePaymentLinkRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *CreatePaymentLinkRequest) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetLineItems

`func (o *CreatePaymentLinkRequest) GetLineItems() []CreatePaymentLinkRequestLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreatePaymentLinkRequest) GetLineItemsOk() (*[]CreatePaymentLinkRequestLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreatePaymentLinkRequest) SetLineItems(v []CreatePaymentLinkRequestLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreatePaymentLinkRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetName

`func (o *CreatePaymentLinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePaymentLinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePaymentLinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePaymentLinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *CreatePaymentLinkRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreatePaymentLinkRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreatePaymentLinkRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreatePaymentLinkRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMode

`func (o *CreatePaymentLinkRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreatePaymentLinkRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreatePaymentLinkRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreatePaymentLinkRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAmountMinor

`func (o *CreatePaymentLinkRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreatePaymentLinkRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreatePaymentLinkRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *CreatePaymentLinkRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePaymentLinkRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentLinkRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentLinkRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentLinkRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceId

`func (o *CreatePaymentLinkRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreatePaymentLinkRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreatePaymentLinkRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreatePaymentLinkRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *CreatePaymentLinkRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreatePaymentLinkRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreatePaymentLinkRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreatePaymentLinkRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CreatePaymentLinkRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CreatePaymentLinkRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CreatePaymentLinkRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CreatePaymentLinkRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *CreatePaymentLinkRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CreatePaymentLinkRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CreatePaymentLinkRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *CreatePaymentLinkRequest) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CreatePaymentLinkRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreatePaymentLinkRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreatePaymentLinkRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreatePaymentLinkRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetCollectShipping

`func (o *CreatePaymentLinkRequest) GetCollectShipping() bool`

GetCollectShipping returns the CollectShipping field if non-nil, zero value otherwise.

### GetCollectShippingOk

`func (o *CreatePaymentLinkRequest) GetCollectShippingOk() (*bool, bool)`

GetCollectShippingOk returns a tuple with the CollectShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectShipping

`func (o *CreatePaymentLinkRequest) SetCollectShipping(v bool)`

SetCollectShipping sets CollectShipping field to given value.

### HasCollectShipping

`func (o *CreatePaymentLinkRequest) HasCollectShipping() bool`

HasCollectShipping returns a boolean if a field has been set.

### GetAllowPromotionCodes

`func (o *CreatePaymentLinkRequest) GetAllowPromotionCodes() bool`

GetAllowPromotionCodes returns the AllowPromotionCodes field if non-nil, zero value otherwise.

### GetAllowPromotionCodesOk

`func (o *CreatePaymentLinkRequest) GetAllowPromotionCodesOk() (*bool, bool)`

GetAllowPromotionCodesOk returns a tuple with the AllowPromotionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPromotionCodes

`func (o *CreatePaymentLinkRequest) SetAllowPromotionCodes(v bool)`

SetAllowPromotionCodes sets AllowPromotionCodes field to given value.

### HasAllowPromotionCodes

`func (o *CreatePaymentLinkRequest) HasAllowPromotionCodes() bool`

HasAllowPromotionCodes returns a boolean if a field has been set.

### GetTaxBehavior

`func (o *CreatePaymentLinkRequest) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *CreatePaymentLinkRequest) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *CreatePaymentLinkRequest) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *CreatePaymentLinkRequest) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePaymentLinkRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePaymentLinkRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePaymentLinkRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePaymentLinkRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreatePaymentLinkRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreatePaymentLinkRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreatePaymentLinkRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreatePaymentLinkRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CreatePaymentLinkRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CreatePaymentLinkRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CreatePaymentLinkRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CreatePaymentLinkRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


