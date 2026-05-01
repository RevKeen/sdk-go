# PaymentLinksCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductIds** | Pointer to **[]string** |  | [optional] 
**LineItems** | Pointer to [**[]PaymentLinksCreateRequestLineItemsInner**](PaymentLinksCreateRequestLineItemsInner.md) |  | [optional] 
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
**SelectionMode** | Pointer to **string** | Customer-facing checkout shape. bundle&#x3D;all products together (default); single_select&#x3D;radio picker; multi_select&#x3D;cart with quantity per line. Companion FKs (default_product_id, highlighted_product_id) are only valid when selection_mode&#x3D;&#39;single_select&#39;. | [optional] [default to "bundle"]
**DefaultProductId** | Pointer to **NullableString** | For single_select only — option pre-selected when checkout loads. Must reference a product attached to this payment link. | [optional] 
**HighlightedProductId** | Pointer to **NullableString** | For single_select only — option shown with a &#39;Popular&#39; badge in the picker. | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewPaymentLinksCreateRequest

`func NewPaymentLinksCreateRequest() *PaymentLinksCreateRequest`

NewPaymentLinksCreateRequest instantiates a new PaymentLinksCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinksCreateRequestWithDefaults

`func NewPaymentLinksCreateRequestWithDefaults() *PaymentLinksCreateRequest`

NewPaymentLinksCreateRequestWithDefaults instantiates a new PaymentLinksCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductIds

`func (o *PaymentLinksCreateRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *PaymentLinksCreateRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *PaymentLinksCreateRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *PaymentLinksCreateRequest) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetLineItems

`func (o *PaymentLinksCreateRequest) GetLineItems() []PaymentLinksCreateRequestLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PaymentLinksCreateRequest) GetLineItemsOk() (*[]PaymentLinksCreateRequestLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PaymentLinksCreateRequest) SetLineItems(v []PaymentLinksCreateRequestLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PaymentLinksCreateRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetName

`func (o *PaymentLinksCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentLinksCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentLinksCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentLinksCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *PaymentLinksCreateRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PaymentLinksCreateRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PaymentLinksCreateRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PaymentLinksCreateRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMode

`func (o *PaymentLinksCreateRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaymentLinksCreateRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaymentLinksCreateRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaymentLinksCreateRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetAmountMinor

`func (o *PaymentLinksCreateRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *PaymentLinksCreateRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *PaymentLinksCreateRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *PaymentLinksCreateRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentLinksCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLinksCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLinksCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentLinksCreateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceId

`func (o *PaymentLinksCreateRequest) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentLinksCreateRequest) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentLinksCreateRequest) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PaymentLinksCreateRequest) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PaymentLinksCreateRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentLinksCreateRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentLinksCreateRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *PaymentLinksCreateRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *PaymentLinksCreateRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *PaymentLinksCreateRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *PaymentLinksCreateRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *PaymentLinksCreateRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetCancelUrl

`func (o *PaymentLinksCreateRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PaymentLinksCreateRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PaymentLinksCreateRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *PaymentLinksCreateRequest) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *PaymentLinksCreateRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PaymentLinksCreateRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PaymentLinksCreateRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *PaymentLinksCreateRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetCollectShipping

`func (o *PaymentLinksCreateRequest) GetCollectShipping() bool`

GetCollectShipping returns the CollectShipping field if non-nil, zero value otherwise.

### GetCollectShippingOk

`func (o *PaymentLinksCreateRequest) GetCollectShippingOk() (*bool, bool)`

GetCollectShippingOk returns a tuple with the CollectShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectShipping

`func (o *PaymentLinksCreateRequest) SetCollectShipping(v bool)`

SetCollectShipping sets CollectShipping field to given value.

### HasCollectShipping

`func (o *PaymentLinksCreateRequest) HasCollectShipping() bool`

HasCollectShipping returns a boolean if a field has been set.

### GetAllowPromotionCodes

`func (o *PaymentLinksCreateRequest) GetAllowPromotionCodes() bool`

GetAllowPromotionCodes returns the AllowPromotionCodes field if non-nil, zero value otherwise.

### GetAllowPromotionCodesOk

`func (o *PaymentLinksCreateRequest) GetAllowPromotionCodesOk() (*bool, bool)`

GetAllowPromotionCodesOk returns a tuple with the AllowPromotionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPromotionCodes

`func (o *PaymentLinksCreateRequest) SetAllowPromotionCodes(v bool)`

SetAllowPromotionCodes sets AllowPromotionCodes field to given value.

### HasAllowPromotionCodes

`func (o *PaymentLinksCreateRequest) HasAllowPromotionCodes() bool`

HasAllowPromotionCodes returns a boolean if a field has been set.

### GetTaxBehavior

`func (o *PaymentLinksCreateRequest) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *PaymentLinksCreateRequest) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *PaymentLinksCreateRequest) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *PaymentLinksCreateRequest) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentLinksCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentLinksCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentLinksCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentLinksCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PaymentLinksCreateRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PaymentLinksCreateRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PaymentLinksCreateRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PaymentLinksCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSelectionMode

`func (o *PaymentLinksCreateRequest) GetSelectionMode() string`

GetSelectionMode returns the SelectionMode field if non-nil, zero value otherwise.

### GetSelectionModeOk

`func (o *PaymentLinksCreateRequest) GetSelectionModeOk() (*string, bool)`

GetSelectionModeOk returns a tuple with the SelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionMode

`func (o *PaymentLinksCreateRequest) SetSelectionMode(v string)`

SetSelectionMode sets SelectionMode field to given value.

### HasSelectionMode

`func (o *PaymentLinksCreateRequest) HasSelectionMode() bool`

HasSelectionMode returns a boolean if a field has been set.

### GetDefaultProductId

`func (o *PaymentLinksCreateRequest) GetDefaultProductId() string`

GetDefaultProductId returns the DefaultProductId field if non-nil, zero value otherwise.

### GetDefaultProductIdOk

`func (o *PaymentLinksCreateRequest) GetDefaultProductIdOk() (*string, bool)`

GetDefaultProductIdOk returns a tuple with the DefaultProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProductId

`func (o *PaymentLinksCreateRequest) SetDefaultProductId(v string)`

SetDefaultProductId sets DefaultProductId field to given value.

### HasDefaultProductId

`func (o *PaymentLinksCreateRequest) HasDefaultProductId() bool`

HasDefaultProductId returns a boolean if a field has been set.

### SetDefaultProductIdNil

`func (o *PaymentLinksCreateRequest) SetDefaultProductIdNil(b bool)`

 SetDefaultProductIdNil sets the value for DefaultProductId to be an explicit nil

### UnsetDefaultProductId
`func (o *PaymentLinksCreateRequest) UnsetDefaultProductId()`

UnsetDefaultProductId ensures that no value is present for DefaultProductId, not even an explicit nil
### GetHighlightedProductId

`func (o *PaymentLinksCreateRequest) GetHighlightedProductId() string`

GetHighlightedProductId returns the HighlightedProductId field if non-nil, zero value otherwise.

### GetHighlightedProductIdOk

`func (o *PaymentLinksCreateRequest) GetHighlightedProductIdOk() (*string, bool)`

GetHighlightedProductIdOk returns a tuple with the HighlightedProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightedProductId

`func (o *PaymentLinksCreateRequest) SetHighlightedProductId(v string)`

SetHighlightedProductId sets HighlightedProductId field to given value.

### HasHighlightedProductId

`func (o *PaymentLinksCreateRequest) HasHighlightedProductId() bool`

HasHighlightedProductId returns a boolean if a field has been set.

### SetHighlightedProductIdNil

`func (o *PaymentLinksCreateRequest) SetHighlightedProductIdNil(b bool)`

 SetHighlightedProductIdNil sets the value for HighlightedProductId to be an explicit nil

### UnsetHighlightedProductId
`func (o *PaymentLinksCreateRequest) UnsetHighlightedProductId()`

UnsetHighlightedProductId ensures that no value is present for HighlightedProductId, not even an explicit nil
### GetIdempotencyKey

`func (o *PaymentLinksCreateRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *PaymentLinksCreateRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *PaymentLinksCreateRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *PaymentLinksCreateRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


