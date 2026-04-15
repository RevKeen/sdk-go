# ListPaymentLinks200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**PublicId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Mode** | **string** |  | 
**Status** | **string** |  | 
**AmountMinor** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**ProductIds** | Pointer to **[]string** |  | [optional] 
**LineItems** | Pointer to [**[]ListPaymentLinks200ResponseDataInnerLineItemsInner**](ListPaymentLinks200ResponseDataInnerLineItemsInner.md) |  | [optional] 
**SuccessUrl** | Pointer to **NullableString** | Custom URL to redirect after successful payment. Supports placeholders: {CHECKOUT_ID}, {TRANSACTION_ID}, {PAYMENT_ID}, {CUSTOMER_ID}, {AMOUNT}, {CURRENCY}, {PORTAL_URL}. If not set, redirects to RevKeen&#39;s default success page. | [optional] 
**CancelUrl** | Pointer to **NullableString** | Custom URL to redirect if customer cancels the checkout. If not set, no cancel redirect is configured. | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**CollectShipping** | Pointer to **bool** |  | [optional] [default to false]
**AllowPromotionCodes** | Pointer to **bool** |  | [optional] [default to true]
**TaxBehavior** | Pointer to **string** |  | [optional] [default to "unspecified"]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewListPaymentLinks200ResponseDataInner

`func NewListPaymentLinks200ResponseDataInner(id string, merchantId string, publicId string, mode string, status string, createdAt time.Time, updatedAt time.Time, ) *ListPaymentLinks200ResponseDataInner`

NewListPaymentLinks200ResponseDataInner instantiates a new ListPaymentLinks200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentLinks200ResponseDataInnerWithDefaults

`func NewListPaymentLinks200ResponseDataInnerWithDefaults() *ListPaymentLinks200ResponseDataInner`

NewListPaymentLinks200ResponseDataInnerWithDefaults instantiates a new ListPaymentLinks200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPaymentLinks200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPaymentLinks200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPaymentLinks200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *ListPaymentLinks200ResponseDataInner) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ListPaymentLinks200ResponseDataInner) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ListPaymentLinks200ResponseDataInner) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetPublicId

`func (o *ListPaymentLinks200ResponseDataInner) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *ListPaymentLinks200ResponseDataInner) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *ListPaymentLinks200ResponseDataInner) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetName

`func (o *ListPaymentLinks200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPaymentLinks200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPaymentLinks200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPaymentLinks200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *ListPaymentLinks200ResponseDataInner) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListPaymentLinks200ResponseDataInner) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListPaymentLinks200ResponseDataInner) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListPaymentLinks200ResponseDataInner) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *ListPaymentLinks200ResponseDataInner) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *ListPaymentLinks200ResponseDataInner) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetMode

`func (o *ListPaymentLinks200ResponseDataInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ListPaymentLinks200ResponseDataInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ListPaymentLinks200ResponseDataInner) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStatus

`func (o *ListPaymentLinks200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPaymentLinks200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPaymentLinks200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountMinor

`func (o *ListPaymentLinks200ResponseDataInner) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *ListPaymentLinks200ResponseDataInner) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *ListPaymentLinks200ResponseDataInner) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *ListPaymentLinks200ResponseDataInner) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### SetAmountMinorNil

`func (o *ListPaymentLinks200ResponseDataInner) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *ListPaymentLinks200ResponseDataInner) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *ListPaymentLinks200ResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListPaymentLinks200ResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListPaymentLinks200ResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListPaymentLinks200ResponseDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInvoiceId

`func (o *ListPaymentLinks200ResponseDataInner) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *ListPaymentLinks200ResponseDataInner) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *ListPaymentLinks200ResponseDataInner) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *ListPaymentLinks200ResponseDataInner) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *ListPaymentLinks200ResponseDataInner) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *ListPaymentLinks200ResponseDataInner) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionId

`func (o *ListPaymentLinks200ResponseDataInner) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ListPaymentLinks200ResponseDataInner) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ListPaymentLinks200ResponseDataInner) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ListPaymentLinks200ResponseDataInner) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ListPaymentLinks200ResponseDataInner) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ListPaymentLinks200ResponseDataInner) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetProductIds

`func (o *ListPaymentLinks200ResponseDataInner) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *ListPaymentLinks200ResponseDataInner) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *ListPaymentLinks200ResponseDataInner) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *ListPaymentLinks200ResponseDataInner) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.

### GetLineItems

`func (o *ListPaymentLinks200ResponseDataInner) GetLineItems() []ListPaymentLinks200ResponseDataInnerLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *ListPaymentLinks200ResponseDataInner) GetLineItemsOk() (*[]ListPaymentLinks200ResponseDataInnerLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *ListPaymentLinks200ResponseDataInner) SetLineItems(v []ListPaymentLinks200ResponseDataInnerLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *ListPaymentLinks200ResponseDataInner) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *ListPaymentLinks200ResponseDataInner) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *ListPaymentLinks200ResponseDataInner) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *ListPaymentLinks200ResponseDataInner) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *ListPaymentLinks200ResponseDataInner) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### SetSuccessUrlNil

`func (o *ListPaymentLinks200ResponseDataInner) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *ListPaymentLinks200ResponseDataInner) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetCancelUrl

`func (o *ListPaymentLinks200ResponseDataInner) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *ListPaymentLinks200ResponseDataInner) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *ListPaymentLinks200ResponseDataInner) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *ListPaymentLinks200ResponseDataInner) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### SetCancelUrlNil

`func (o *ListPaymentLinks200ResponseDataInner) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *ListPaymentLinks200ResponseDataInner) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetRedirectUrl

`func (o *ListPaymentLinks200ResponseDataInner) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ListPaymentLinks200ResponseDataInner) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ListPaymentLinks200ResponseDataInner) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ListPaymentLinks200ResponseDataInner) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *ListPaymentLinks200ResponseDataInner) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *ListPaymentLinks200ResponseDataInner) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetCollectShipping

`func (o *ListPaymentLinks200ResponseDataInner) GetCollectShipping() bool`

GetCollectShipping returns the CollectShipping field if non-nil, zero value otherwise.

### GetCollectShippingOk

`func (o *ListPaymentLinks200ResponseDataInner) GetCollectShippingOk() (*bool, bool)`

GetCollectShippingOk returns a tuple with the CollectShipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectShipping

`func (o *ListPaymentLinks200ResponseDataInner) SetCollectShipping(v bool)`

SetCollectShipping sets CollectShipping field to given value.

### HasCollectShipping

`func (o *ListPaymentLinks200ResponseDataInner) HasCollectShipping() bool`

HasCollectShipping returns a boolean if a field has been set.

### GetAllowPromotionCodes

`func (o *ListPaymentLinks200ResponseDataInner) GetAllowPromotionCodes() bool`

GetAllowPromotionCodes returns the AllowPromotionCodes field if non-nil, zero value otherwise.

### GetAllowPromotionCodesOk

`func (o *ListPaymentLinks200ResponseDataInner) GetAllowPromotionCodesOk() (*bool, bool)`

GetAllowPromotionCodesOk returns a tuple with the AllowPromotionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPromotionCodes

`func (o *ListPaymentLinks200ResponseDataInner) SetAllowPromotionCodes(v bool)`

SetAllowPromotionCodes sets AllowPromotionCodes field to given value.

### HasAllowPromotionCodes

`func (o *ListPaymentLinks200ResponseDataInner) HasAllowPromotionCodes() bool`

HasAllowPromotionCodes returns a boolean if a field has been set.

### GetTaxBehavior

`func (o *ListPaymentLinks200ResponseDataInner) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *ListPaymentLinks200ResponseDataInner) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *ListPaymentLinks200ResponseDataInner) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *ListPaymentLinks200ResponseDataInner) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### GetMetadata

`func (o *ListPaymentLinks200ResponseDataInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListPaymentLinks200ResponseDataInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListPaymentLinks200ResponseDataInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListPaymentLinks200ResponseDataInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ListPaymentLinks200ResponseDataInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListPaymentLinks200ResponseDataInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListPaymentLinks200ResponseDataInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ListPaymentLinks200ResponseDataInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ListPaymentLinks200ResponseDataInner) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ListPaymentLinks200ResponseDataInner) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *ListPaymentLinks200ResponseDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListPaymentLinks200ResponseDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListPaymentLinks200ResponseDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListPaymentLinks200ResponseDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListPaymentLinks200ResponseDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListPaymentLinks200ResponseDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


