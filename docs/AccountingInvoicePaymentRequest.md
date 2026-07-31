# AccountingInvoicePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**MerchantId** | **string** |  | 
**Provider** | **string** |  | 
**ConnectionId** | **string** |  | 
**ProviderAccountId** | **string** |  | 
**ExternalInvoiceId** | **string** |  | 
**ExternalInvoiceNumber** | Pointer to **NullableString** |  | [optional] 
**ExternalInvoiceStatus** | Pointer to **NullableString** |  | [optional] 
**ExternalInvoiceUrl** | Pointer to **NullableString** |  | [optional] 
**ExternalInvoiceIssuedAt** | Pointer to **NullableString** |  | [optional] 
**ExternalInvoiceDueAt** | Pointer to **NullableString** |  | [optional] 
**ExternalInvoiceUpdatedAt** | Pointer to **NullableString** |  | [optional] 
**ExternalCustomerId** | Pointer to **NullableString** |  | [optional] 
**ExternalCustomerReference** | Pointer to **NullableString** |  | [optional] 
**ExternalCustomerName** | Pointer to **NullableString** |  | [optional] 
**ExternalCustomerEmail** | Pointer to **NullableString** |  | [optional] 
**AmountDueMinor** | **int32** |  | 
**TotalAmountMinor** | Pointer to **NullableInt32** |  | [optional] 
**AmountPaidMinor** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | **string** |  | 
**CheckoutSessionId** | Pointer to **NullableString** |  | [optional] 
**CheckoutUrl** | Pointer to **NullableString** |  | [optional] 
**CheckoutExpiresAt** | Pointer to **NullableString** |  | [optional] 
**CheckoutSuccessUrl** | Pointer to **NullableString** |  | [optional] 
**CheckoutCancelUrl** | Pointer to **NullableString** |  | [optional] 
**CheckoutAllowedMethods** | Pointer to **[]string** |  | [optional] 
**Status** | **string** |  | 
**SyncStatus** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**PayloadFingerprint** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewAccountingInvoicePaymentRequest

`func NewAccountingInvoicePaymentRequest(id string, object string, merchantId string, provider string, connectionId string, providerAccountId string, externalInvoiceId string, amountDueMinor int32, currency string, status string, syncStatus string, idempotencyKey string, payloadFingerprint string, createdAt string, updatedAt string, ) *AccountingInvoicePaymentRequest`

NewAccountingInvoicePaymentRequest instantiates a new AccountingInvoicePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingInvoicePaymentRequestWithDefaults

`func NewAccountingInvoicePaymentRequestWithDefaults() *AccountingInvoicePaymentRequest`

NewAccountingInvoicePaymentRequestWithDefaults instantiates a new AccountingInvoicePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountingInvoicePaymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountingInvoicePaymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountingInvoicePaymentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AccountingInvoicePaymentRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountingInvoicePaymentRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountingInvoicePaymentRequest) SetObject(v string)`

SetObject sets Object field to given value.


### GetMerchantId

`func (o *AccountingInvoicePaymentRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *AccountingInvoicePaymentRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *AccountingInvoicePaymentRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetProvider

`func (o *AccountingInvoicePaymentRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AccountingInvoicePaymentRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AccountingInvoicePaymentRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetConnectionId

`func (o *AccountingInvoicePaymentRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *AccountingInvoicePaymentRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *AccountingInvoicePaymentRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetProviderAccountId

`func (o *AccountingInvoicePaymentRequest) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *AccountingInvoicePaymentRequest) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *AccountingInvoicePaymentRequest) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetExternalInvoiceId

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceId() string`

GetExternalInvoiceId returns the ExternalInvoiceId field if non-nil, zero value otherwise.

### GetExternalInvoiceIdOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIdOk() (*string, bool)`

GetExternalInvoiceIdOk returns a tuple with the ExternalInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceId

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceId(v string)`

SetExternalInvoiceId sets ExternalInvoiceId field to given value.


### GetExternalInvoiceNumber

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceNumber() string`

GetExternalInvoiceNumber returns the ExternalInvoiceNumber field if non-nil, zero value otherwise.

### GetExternalInvoiceNumberOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceNumberOk() (*string, bool)`

GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceNumber

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceNumber(v string)`

SetExternalInvoiceNumber sets ExternalInvoiceNumber field to given value.

### HasExternalInvoiceNumber

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceNumber() bool`

HasExternalInvoiceNumber returns a boolean if a field has been set.

### SetExternalInvoiceNumberNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceNumberNil(b bool)`

 SetExternalInvoiceNumberNil sets the value for ExternalInvoiceNumber to be an explicit nil

### UnsetExternalInvoiceNumber
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceNumber()`

UnsetExternalInvoiceNumber ensures that no value is present for ExternalInvoiceNumber, not even an explicit nil
### GetExternalInvoiceStatus

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceStatus() string`

GetExternalInvoiceStatus returns the ExternalInvoiceStatus field if non-nil, zero value otherwise.

### GetExternalInvoiceStatusOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceStatusOk() (*string, bool)`

GetExternalInvoiceStatusOk returns a tuple with the ExternalInvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceStatus

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceStatus(v string)`

SetExternalInvoiceStatus sets ExternalInvoiceStatus field to given value.

### HasExternalInvoiceStatus

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceStatus() bool`

HasExternalInvoiceStatus returns a boolean if a field has been set.

### SetExternalInvoiceStatusNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceStatusNil(b bool)`

 SetExternalInvoiceStatusNil sets the value for ExternalInvoiceStatus to be an explicit nil

### UnsetExternalInvoiceStatus
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceStatus()`

UnsetExternalInvoiceStatus ensures that no value is present for ExternalInvoiceStatus, not even an explicit nil
### GetExternalInvoiceUrl

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUrl() string`

GetExternalInvoiceUrl returns the ExternalInvoiceUrl field if non-nil, zero value otherwise.

### GetExternalInvoiceUrlOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUrlOk() (*string, bool)`

GetExternalInvoiceUrlOk returns a tuple with the ExternalInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceUrl

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUrl(v string)`

SetExternalInvoiceUrl sets ExternalInvoiceUrl field to given value.

### HasExternalInvoiceUrl

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceUrl() bool`

HasExternalInvoiceUrl returns a boolean if a field has been set.

### SetExternalInvoiceUrlNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUrlNil(b bool)`

 SetExternalInvoiceUrlNil sets the value for ExternalInvoiceUrl to be an explicit nil

### UnsetExternalInvoiceUrl
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceUrl()`

UnsetExternalInvoiceUrl ensures that no value is present for ExternalInvoiceUrl, not even an explicit nil
### GetExternalInvoiceIssuedAt

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIssuedAt() string`

GetExternalInvoiceIssuedAt returns the ExternalInvoiceIssuedAt field if non-nil, zero value otherwise.

### GetExternalInvoiceIssuedAtOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceIssuedAtOk() (*string, bool)`

GetExternalInvoiceIssuedAtOk returns a tuple with the ExternalInvoiceIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceIssuedAt

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceIssuedAt(v string)`

SetExternalInvoiceIssuedAt sets ExternalInvoiceIssuedAt field to given value.

### HasExternalInvoiceIssuedAt

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceIssuedAt() bool`

HasExternalInvoiceIssuedAt returns a boolean if a field has been set.

### SetExternalInvoiceIssuedAtNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceIssuedAtNil(b bool)`

 SetExternalInvoiceIssuedAtNil sets the value for ExternalInvoiceIssuedAt to be an explicit nil

### UnsetExternalInvoiceIssuedAt
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceIssuedAt()`

UnsetExternalInvoiceIssuedAt ensures that no value is present for ExternalInvoiceIssuedAt, not even an explicit nil
### GetExternalInvoiceDueAt

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceDueAt() string`

GetExternalInvoiceDueAt returns the ExternalInvoiceDueAt field if non-nil, zero value otherwise.

### GetExternalInvoiceDueAtOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceDueAtOk() (*string, bool)`

GetExternalInvoiceDueAtOk returns a tuple with the ExternalInvoiceDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceDueAt

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceDueAt(v string)`

SetExternalInvoiceDueAt sets ExternalInvoiceDueAt field to given value.

### HasExternalInvoiceDueAt

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceDueAt() bool`

HasExternalInvoiceDueAt returns a boolean if a field has been set.

### SetExternalInvoiceDueAtNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceDueAtNil(b bool)`

 SetExternalInvoiceDueAtNil sets the value for ExternalInvoiceDueAt to be an explicit nil

### UnsetExternalInvoiceDueAt
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceDueAt()`

UnsetExternalInvoiceDueAt ensures that no value is present for ExternalInvoiceDueAt, not even an explicit nil
### GetExternalInvoiceUpdatedAt

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUpdatedAt() string`

GetExternalInvoiceUpdatedAt returns the ExternalInvoiceUpdatedAt field if non-nil, zero value otherwise.

### GetExternalInvoiceUpdatedAtOk

`func (o *AccountingInvoicePaymentRequest) GetExternalInvoiceUpdatedAtOk() (*string, bool)`

GetExternalInvoiceUpdatedAtOk returns a tuple with the ExternalInvoiceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoiceUpdatedAt

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUpdatedAt(v string)`

SetExternalInvoiceUpdatedAt sets ExternalInvoiceUpdatedAt field to given value.

### HasExternalInvoiceUpdatedAt

`func (o *AccountingInvoicePaymentRequest) HasExternalInvoiceUpdatedAt() bool`

HasExternalInvoiceUpdatedAt returns a boolean if a field has been set.

### SetExternalInvoiceUpdatedAtNil

`func (o *AccountingInvoicePaymentRequest) SetExternalInvoiceUpdatedAtNil(b bool)`

 SetExternalInvoiceUpdatedAtNil sets the value for ExternalInvoiceUpdatedAt to be an explicit nil

### UnsetExternalInvoiceUpdatedAt
`func (o *AccountingInvoicePaymentRequest) UnsetExternalInvoiceUpdatedAt()`

UnsetExternalInvoiceUpdatedAt ensures that no value is present for ExternalInvoiceUpdatedAt, not even an explicit nil
### GetExternalCustomerId

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *AccountingInvoicePaymentRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### SetExternalCustomerIdNil

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerIdNil(b bool)`

 SetExternalCustomerIdNil sets the value for ExternalCustomerId to be an explicit nil

### UnsetExternalCustomerId
`func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerId()`

UnsetExternalCustomerId ensures that no value is present for ExternalCustomerId, not even an explicit nil
### GetExternalCustomerReference

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerReference() string`

GetExternalCustomerReference returns the ExternalCustomerReference field if non-nil, zero value otherwise.

### GetExternalCustomerReferenceOk

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerReferenceOk() (*string, bool)`

GetExternalCustomerReferenceOk returns a tuple with the ExternalCustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerReference

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerReference(v string)`

SetExternalCustomerReference sets ExternalCustomerReference field to given value.

### HasExternalCustomerReference

`func (o *AccountingInvoicePaymentRequest) HasExternalCustomerReference() bool`

HasExternalCustomerReference returns a boolean if a field has been set.

### SetExternalCustomerReferenceNil

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerReferenceNil(b bool)`

 SetExternalCustomerReferenceNil sets the value for ExternalCustomerReference to be an explicit nil

### UnsetExternalCustomerReference
`func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerReference()`

UnsetExternalCustomerReference ensures that no value is present for ExternalCustomerReference, not even an explicit nil
### GetExternalCustomerName

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerName() string`

GetExternalCustomerName returns the ExternalCustomerName field if non-nil, zero value otherwise.

### GetExternalCustomerNameOk

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerNameOk() (*string, bool)`

GetExternalCustomerNameOk returns a tuple with the ExternalCustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerName

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerName(v string)`

SetExternalCustomerName sets ExternalCustomerName field to given value.

### HasExternalCustomerName

`func (o *AccountingInvoicePaymentRequest) HasExternalCustomerName() bool`

HasExternalCustomerName returns a boolean if a field has been set.

### SetExternalCustomerNameNil

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerNameNil(b bool)`

 SetExternalCustomerNameNil sets the value for ExternalCustomerName to be an explicit nil

### UnsetExternalCustomerName
`func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerName()`

UnsetExternalCustomerName ensures that no value is present for ExternalCustomerName, not even an explicit nil
### GetExternalCustomerEmail

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerEmail() string`

GetExternalCustomerEmail returns the ExternalCustomerEmail field if non-nil, zero value otherwise.

### GetExternalCustomerEmailOk

`func (o *AccountingInvoicePaymentRequest) GetExternalCustomerEmailOk() (*string, bool)`

GetExternalCustomerEmailOk returns a tuple with the ExternalCustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerEmail

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerEmail(v string)`

SetExternalCustomerEmail sets ExternalCustomerEmail field to given value.

### HasExternalCustomerEmail

`func (o *AccountingInvoicePaymentRequest) HasExternalCustomerEmail() bool`

HasExternalCustomerEmail returns a boolean if a field has been set.

### SetExternalCustomerEmailNil

`func (o *AccountingInvoicePaymentRequest) SetExternalCustomerEmailNil(b bool)`

 SetExternalCustomerEmailNil sets the value for ExternalCustomerEmail to be an explicit nil

### UnsetExternalCustomerEmail
`func (o *AccountingInvoicePaymentRequest) UnsetExternalCustomerEmail()`

UnsetExternalCustomerEmail ensures that no value is present for ExternalCustomerEmail, not even an explicit nil
### GetAmountDueMinor

`func (o *AccountingInvoicePaymentRequest) GetAmountDueMinor() int32`

GetAmountDueMinor returns the AmountDueMinor field if non-nil, zero value otherwise.

### GetAmountDueMinorOk

`func (o *AccountingInvoicePaymentRequest) GetAmountDueMinorOk() (*int32, bool)`

GetAmountDueMinorOk returns a tuple with the AmountDueMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueMinor

`func (o *AccountingInvoicePaymentRequest) SetAmountDueMinor(v int32)`

SetAmountDueMinor sets AmountDueMinor field to given value.


### GetTotalAmountMinor

`func (o *AccountingInvoicePaymentRequest) GetTotalAmountMinor() int32`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *AccountingInvoicePaymentRequest) GetTotalAmountMinorOk() (*int32, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *AccountingInvoicePaymentRequest) SetTotalAmountMinor(v int32)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.

### HasTotalAmountMinor

`func (o *AccountingInvoicePaymentRequest) HasTotalAmountMinor() bool`

HasTotalAmountMinor returns a boolean if a field has been set.

### SetTotalAmountMinorNil

`func (o *AccountingInvoicePaymentRequest) SetTotalAmountMinorNil(b bool)`

 SetTotalAmountMinorNil sets the value for TotalAmountMinor to be an explicit nil

### UnsetTotalAmountMinor
`func (o *AccountingInvoicePaymentRequest) UnsetTotalAmountMinor()`

UnsetTotalAmountMinor ensures that no value is present for TotalAmountMinor, not even an explicit nil
### GetAmountPaidMinor

`func (o *AccountingInvoicePaymentRequest) GetAmountPaidMinor() int32`

GetAmountPaidMinor returns the AmountPaidMinor field if non-nil, zero value otherwise.

### GetAmountPaidMinorOk

`func (o *AccountingInvoicePaymentRequest) GetAmountPaidMinorOk() (*int32, bool)`

GetAmountPaidMinorOk returns a tuple with the AmountPaidMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidMinor

`func (o *AccountingInvoicePaymentRequest) SetAmountPaidMinor(v int32)`

SetAmountPaidMinor sets AmountPaidMinor field to given value.

### HasAmountPaidMinor

`func (o *AccountingInvoicePaymentRequest) HasAmountPaidMinor() bool`

HasAmountPaidMinor returns a boolean if a field has been set.

### SetAmountPaidMinorNil

`func (o *AccountingInvoicePaymentRequest) SetAmountPaidMinorNil(b bool)`

 SetAmountPaidMinorNil sets the value for AmountPaidMinor to be an explicit nil

### UnsetAmountPaidMinor
`func (o *AccountingInvoicePaymentRequest) UnsetAmountPaidMinor()`

UnsetAmountPaidMinor ensures that no value is present for AmountPaidMinor, not even an explicit nil
### GetCurrency

`func (o *AccountingInvoicePaymentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AccountingInvoicePaymentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AccountingInvoicePaymentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCheckoutSessionId

`func (o *AccountingInvoicePaymentRequest) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *AccountingInvoicePaymentRequest) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.

### HasCheckoutSessionId

`func (o *AccountingInvoicePaymentRequest) HasCheckoutSessionId() bool`

HasCheckoutSessionId returns a boolean if a field has been set.

### SetCheckoutSessionIdNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutSessionIdNil(b bool)`

 SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil

### UnsetCheckoutSessionId
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutSessionId()`

UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
### GetCheckoutUrl

`func (o *AccountingInvoicePaymentRequest) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *AccountingInvoicePaymentRequest) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *AccountingInvoicePaymentRequest) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetCheckoutExpiresAt

`func (o *AccountingInvoicePaymentRequest) GetCheckoutExpiresAt() string`

GetCheckoutExpiresAt returns the CheckoutExpiresAt field if non-nil, zero value otherwise.

### GetCheckoutExpiresAtOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutExpiresAtOk() (*string, bool)`

GetCheckoutExpiresAtOk returns a tuple with the CheckoutExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutExpiresAt

`func (o *AccountingInvoicePaymentRequest) SetCheckoutExpiresAt(v string)`

SetCheckoutExpiresAt sets CheckoutExpiresAt field to given value.

### HasCheckoutExpiresAt

`func (o *AccountingInvoicePaymentRequest) HasCheckoutExpiresAt() bool`

HasCheckoutExpiresAt returns a boolean if a field has been set.

### SetCheckoutExpiresAtNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutExpiresAtNil(b bool)`

 SetCheckoutExpiresAtNil sets the value for CheckoutExpiresAt to be an explicit nil

### UnsetCheckoutExpiresAt
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutExpiresAt()`

UnsetCheckoutExpiresAt ensures that no value is present for CheckoutExpiresAt, not even an explicit nil
### GetCheckoutSuccessUrl

`func (o *AccountingInvoicePaymentRequest) GetCheckoutSuccessUrl() string`

GetCheckoutSuccessUrl returns the CheckoutSuccessUrl field if non-nil, zero value otherwise.

### GetCheckoutSuccessUrlOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutSuccessUrlOk() (*string, bool)`

GetCheckoutSuccessUrlOk returns a tuple with the CheckoutSuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSuccessUrl

`func (o *AccountingInvoicePaymentRequest) SetCheckoutSuccessUrl(v string)`

SetCheckoutSuccessUrl sets CheckoutSuccessUrl field to given value.

### HasCheckoutSuccessUrl

`func (o *AccountingInvoicePaymentRequest) HasCheckoutSuccessUrl() bool`

HasCheckoutSuccessUrl returns a boolean if a field has been set.

### SetCheckoutSuccessUrlNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutSuccessUrlNil(b bool)`

 SetCheckoutSuccessUrlNil sets the value for CheckoutSuccessUrl to be an explicit nil

### UnsetCheckoutSuccessUrl
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutSuccessUrl()`

UnsetCheckoutSuccessUrl ensures that no value is present for CheckoutSuccessUrl, not even an explicit nil
### GetCheckoutCancelUrl

`func (o *AccountingInvoicePaymentRequest) GetCheckoutCancelUrl() string`

GetCheckoutCancelUrl returns the CheckoutCancelUrl field if non-nil, zero value otherwise.

### GetCheckoutCancelUrlOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutCancelUrlOk() (*string, bool)`

GetCheckoutCancelUrlOk returns a tuple with the CheckoutCancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutCancelUrl

`func (o *AccountingInvoicePaymentRequest) SetCheckoutCancelUrl(v string)`

SetCheckoutCancelUrl sets CheckoutCancelUrl field to given value.

### HasCheckoutCancelUrl

`func (o *AccountingInvoicePaymentRequest) HasCheckoutCancelUrl() bool`

HasCheckoutCancelUrl returns a boolean if a field has been set.

### SetCheckoutCancelUrlNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutCancelUrlNil(b bool)`

 SetCheckoutCancelUrlNil sets the value for CheckoutCancelUrl to be an explicit nil

### UnsetCheckoutCancelUrl
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutCancelUrl()`

UnsetCheckoutCancelUrl ensures that no value is present for CheckoutCancelUrl, not even an explicit nil
### GetCheckoutAllowedMethods

`func (o *AccountingInvoicePaymentRequest) GetCheckoutAllowedMethods() []string`

GetCheckoutAllowedMethods returns the CheckoutAllowedMethods field if non-nil, zero value otherwise.

### GetCheckoutAllowedMethodsOk

`func (o *AccountingInvoicePaymentRequest) GetCheckoutAllowedMethodsOk() (*[]string, bool)`

GetCheckoutAllowedMethodsOk returns a tuple with the CheckoutAllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutAllowedMethods

`func (o *AccountingInvoicePaymentRequest) SetCheckoutAllowedMethods(v []string)`

SetCheckoutAllowedMethods sets CheckoutAllowedMethods field to given value.

### HasCheckoutAllowedMethods

`func (o *AccountingInvoicePaymentRequest) HasCheckoutAllowedMethods() bool`

HasCheckoutAllowedMethods returns a boolean if a field has been set.

### SetCheckoutAllowedMethodsNil

`func (o *AccountingInvoicePaymentRequest) SetCheckoutAllowedMethodsNil(b bool)`

 SetCheckoutAllowedMethodsNil sets the value for CheckoutAllowedMethods to be an explicit nil

### UnsetCheckoutAllowedMethods
`func (o *AccountingInvoicePaymentRequest) UnsetCheckoutAllowedMethods()`

UnsetCheckoutAllowedMethods ensures that no value is present for CheckoutAllowedMethods, not even an explicit nil
### GetStatus

`func (o *AccountingInvoicePaymentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountingInvoicePaymentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountingInvoicePaymentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSyncStatus

`func (o *AccountingInvoicePaymentRequest) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *AccountingInvoicePaymentRequest) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *AccountingInvoicePaymentRequest) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.


### GetIdempotencyKey

`func (o *AccountingInvoicePaymentRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *AccountingInvoicePaymentRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *AccountingInvoicePaymentRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetPayloadFingerprint

`func (o *AccountingInvoicePaymentRequest) GetPayloadFingerprint() string`

GetPayloadFingerprint returns the PayloadFingerprint field if non-nil, zero value otherwise.

### GetPayloadFingerprintOk

`func (o *AccountingInvoicePaymentRequest) GetPayloadFingerprintOk() (*string, bool)`

GetPayloadFingerprintOk returns a tuple with the PayloadFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFingerprint

`func (o *AccountingInvoicePaymentRequest) SetPayloadFingerprint(v string)`

SetPayloadFingerprint sets PayloadFingerprint field to given value.


### GetCreatedAt

`func (o *AccountingInvoicePaymentRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountingInvoicePaymentRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountingInvoicePaymentRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccountingInvoicePaymentRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountingInvoicePaymentRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountingInvoicePaymentRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


