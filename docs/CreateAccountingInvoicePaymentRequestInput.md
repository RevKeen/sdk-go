# CreateAccountingInvoicePaymentRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**ConnectionId** | **string** |  | 
**ProviderAccountId** | **string** |  | 
**ExternalInvoice** | [**CreateAccountingInvoicePaymentRequestInputExternalInvoice**](CreateAccountingInvoicePaymentRequestInputExternalInvoice.md) |  | 
**ExternalCustomer** | Pointer to [**CreateAccountingInvoicePaymentRequestInputExternalCustomer**](CreateAccountingInvoicePaymentRequestInputExternalCustomer.md) |  | [optional] 
**AmountDueMinor** | **int32** |  | 
**TotalAmountMinor** | Pointer to **NullableInt32** |  | [optional] 
**AmountPaidMinor** | Pointer to **NullableInt32** |  | [optional] 
**Currency** | **string** |  | 
**Checkout** | Pointer to [**CreateAccountingInvoicePaymentRequestInputCheckout**](CreateAccountingInvoicePaymentRequestInputCheckout.md) |  | [optional] 
**PayloadFingerprint** | Pointer to **NullableString** |  | [optional] 
**ProviderMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**SafeProviderInvoiceSnapshot** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateAccountingInvoicePaymentRequestInput

`func NewCreateAccountingInvoicePaymentRequestInput(provider string, connectionId string, providerAccountId string, externalInvoice CreateAccountingInvoicePaymentRequestInputExternalInvoice, amountDueMinor int32, currency string, ) *CreateAccountingInvoicePaymentRequestInput`

NewCreateAccountingInvoicePaymentRequestInput instantiates a new CreateAccountingInvoicePaymentRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountingInvoicePaymentRequestInputWithDefaults

`func NewCreateAccountingInvoicePaymentRequestInputWithDefaults() *CreateAccountingInvoicePaymentRequestInput`

NewCreateAccountingInvoicePaymentRequestInputWithDefaults instantiates a new CreateAccountingInvoicePaymentRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateAccountingInvoicePaymentRequestInput) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetConnectionId

`func (o *CreateAccountingInvoicePaymentRequestInput) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateAccountingInvoicePaymentRequestInput) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetProviderAccountId

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *CreateAccountingInvoicePaymentRequestInput) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetExternalInvoice

`func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalInvoice() CreateAccountingInvoicePaymentRequestInputExternalInvoice`

GetExternalInvoice returns the ExternalInvoice field if non-nil, zero value otherwise.

### GetExternalInvoiceOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalInvoiceOk() (*CreateAccountingInvoicePaymentRequestInputExternalInvoice, bool)`

GetExternalInvoiceOk returns a tuple with the ExternalInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalInvoice

`func (o *CreateAccountingInvoicePaymentRequestInput) SetExternalInvoice(v CreateAccountingInvoicePaymentRequestInputExternalInvoice)`

SetExternalInvoice sets ExternalInvoice field to given value.


### GetExternalCustomer

`func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalCustomer() CreateAccountingInvoicePaymentRequestInputExternalCustomer`

GetExternalCustomer returns the ExternalCustomer field if non-nil, zero value otherwise.

### GetExternalCustomerOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetExternalCustomerOk() (*CreateAccountingInvoicePaymentRequestInputExternalCustomer, bool)`

GetExternalCustomerOk returns a tuple with the ExternalCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomer

`func (o *CreateAccountingInvoicePaymentRequestInput) SetExternalCustomer(v CreateAccountingInvoicePaymentRequestInputExternalCustomer)`

SetExternalCustomer sets ExternalCustomer field to given value.

### HasExternalCustomer

`func (o *CreateAccountingInvoicePaymentRequestInput) HasExternalCustomer() bool`

HasExternalCustomer returns a boolean if a field has been set.

### GetAmountDueMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountDueMinor() int32`

GetAmountDueMinor returns the AmountDueMinor field if non-nil, zero value otherwise.

### GetAmountDueMinorOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountDueMinorOk() (*int32, bool)`

GetAmountDueMinorOk returns a tuple with the AmountDueMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDueMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountDueMinor(v int32)`

SetAmountDueMinor sets AmountDueMinor field to given value.


### GetTotalAmountMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) GetTotalAmountMinor() int32`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetTotalAmountMinorOk() (*int32, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) SetTotalAmountMinor(v int32)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.

### HasTotalAmountMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) HasTotalAmountMinor() bool`

HasTotalAmountMinor returns a boolean if a field has been set.

### SetTotalAmountMinorNil

`func (o *CreateAccountingInvoicePaymentRequestInput) SetTotalAmountMinorNil(b bool)`

 SetTotalAmountMinorNil sets the value for TotalAmountMinor to be an explicit nil

### UnsetTotalAmountMinor
`func (o *CreateAccountingInvoicePaymentRequestInput) UnsetTotalAmountMinor()`

UnsetTotalAmountMinor ensures that no value is present for TotalAmountMinor, not even an explicit nil
### GetAmountPaidMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountPaidMinor() int32`

GetAmountPaidMinor returns the AmountPaidMinor field if non-nil, zero value otherwise.

### GetAmountPaidMinorOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetAmountPaidMinorOk() (*int32, bool)`

GetAmountPaidMinorOk returns a tuple with the AmountPaidMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaidMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountPaidMinor(v int32)`

SetAmountPaidMinor sets AmountPaidMinor field to given value.

### HasAmountPaidMinor

`func (o *CreateAccountingInvoicePaymentRequestInput) HasAmountPaidMinor() bool`

HasAmountPaidMinor returns a boolean if a field has been set.

### SetAmountPaidMinorNil

`func (o *CreateAccountingInvoicePaymentRequestInput) SetAmountPaidMinorNil(b bool)`

 SetAmountPaidMinorNil sets the value for AmountPaidMinor to be an explicit nil

### UnsetAmountPaidMinor
`func (o *CreateAccountingInvoicePaymentRequestInput) UnsetAmountPaidMinor()`

UnsetAmountPaidMinor ensures that no value is present for AmountPaidMinor, not even an explicit nil
### GetCurrency

`func (o *CreateAccountingInvoicePaymentRequestInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateAccountingInvoicePaymentRequestInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCheckout

`func (o *CreateAccountingInvoicePaymentRequestInput) GetCheckout() CreateAccountingInvoicePaymentRequestInputCheckout`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetCheckoutOk() (*CreateAccountingInvoicePaymentRequestInputCheckout, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *CreateAccountingInvoicePaymentRequestInput) SetCheckout(v CreateAccountingInvoicePaymentRequestInputCheckout)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *CreateAccountingInvoicePaymentRequestInput) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetPayloadFingerprint

`func (o *CreateAccountingInvoicePaymentRequestInput) GetPayloadFingerprint() string`

GetPayloadFingerprint returns the PayloadFingerprint field if non-nil, zero value otherwise.

### GetPayloadFingerprintOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetPayloadFingerprintOk() (*string, bool)`

GetPayloadFingerprintOk returns a tuple with the PayloadFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFingerprint

`func (o *CreateAccountingInvoicePaymentRequestInput) SetPayloadFingerprint(v string)`

SetPayloadFingerprint sets PayloadFingerprint field to given value.

### HasPayloadFingerprint

`func (o *CreateAccountingInvoicePaymentRequestInput) HasPayloadFingerprint() bool`

HasPayloadFingerprint returns a boolean if a field has been set.

### SetPayloadFingerprintNil

`func (o *CreateAccountingInvoicePaymentRequestInput) SetPayloadFingerprintNil(b bool)`

 SetPayloadFingerprintNil sets the value for PayloadFingerprint to be an explicit nil

### UnsetPayloadFingerprint
`func (o *CreateAccountingInvoicePaymentRequestInput) UnsetPayloadFingerprint()`

UnsetPayloadFingerprint ensures that no value is present for PayloadFingerprint, not even an explicit nil
### GetProviderMetadata

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderMetadata() map[string]interface{}`

GetProviderMetadata returns the ProviderMetadata field if non-nil, zero value otherwise.

### GetProviderMetadataOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetProviderMetadataOk() (*map[string]interface{}, bool)`

GetProviderMetadataOk returns a tuple with the ProviderMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMetadata

`func (o *CreateAccountingInvoicePaymentRequestInput) SetProviderMetadata(v map[string]interface{})`

SetProviderMetadata sets ProviderMetadata field to given value.

### HasProviderMetadata

`func (o *CreateAccountingInvoicePaymentRequestInput) HasProviderMetadata() bool`

HasProviderMetadata returns a boolean if a field has been set.

### GetSafeProviderInvoiceSnapshot

`func (o *CreateAccountingInvoicePaymentRequestInput) GetSafeProviderInvoiceSnapshot() map[string]interface{}`

GetSafeProviderInvoiceSnapshot returns the SafeProviderInvoiceSnapshot field if non-nil, zero value otherwise.

### GetSafeProviderInvoiceSnapshotOk

`func (o *CreateAccountingInvoicePaymentRequestInput) GetSafeProviderInvoiceSnapshotOk() (*map[string]interface{}, bool)`

GetSafeProviderInvoiceSnapshotOk returns a tuple with the SafeProviderInvoiceSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeProviderInvoiceSnapshot

`func (o *CreateAccountingInvoicePaymentRequestInput) SetSafeProviderInvoiceSnapshot(v map[string]interface{})`

SetSafeProviderInvoiceSnapshot sets SafeProviderInvoiceSnapshot field to given value.

### HasSafeProviderInvoiceSnapshot

`func (o *CreateAccountingInvoicePaymentRequestInput) HasSafeProviderInvoiceSnapshot() bool`

HasSafeProviderInvoiceSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


