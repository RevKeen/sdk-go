# CheckoutSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Status** | **string** |  | 
**Mode** | **NullableString** |  | 
**AmountTotal** | **NullableFloat32** |  | 
**Currency** | **NullableString** |  | 
**CustomerId** | **NullableString** |  | 
**PaymentIntentId** | **NullableString** |  | 
**InvoiceId** | **NullableString** |  | 
**SubscriptionId** | **NullableString** |  | 
**Url** | **NullableString** |  | 
**SuccessUrl** | **NullableString** |  | 
**CancelUrl** | **NullableString** |  | 
**ExpiresAt** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**AllowedMethods** | **[]string** |  | 
**SelectedMethod** | **NullableString** |  | 

## Methods

### NewCheckoutSession

`func NewCheckoutSession(id string, object string, status string, mode NullableString, amountTotal NullableFloat32, currency NullableString, customerId NullableString, paymentIntentId NullableString, invoiceId NullableString, subscriptionId NullableString, url NullableString, successUrl NullableString, cancelUrl NullableString, expiresAt NullableString, createdAt NullableString, allowedMethods []string, selectedMethod NullableString, ) *CheckoutSession`

NewCheckoutSession instantiates a new CheckoutSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutSessionWithDefaults

`func NewCheckoutSessionWithDefaults() *CheckoutSession`

NewCheckoutSessionWithDefaults instantiates a new CheckoutSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckoutSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutSession) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *CheckoutSession) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CheckoutSession) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CheckoutSession) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *CheckoutSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckoutSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckoutSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMode

`func (o *CheckoutSession) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CheckoutSession) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CheckoutSession) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *CheckoutSession) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *CheckoutSession) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetAmountTotal

`func (o *CheckoutSession) GetAmountTotal() float32`

GetAmountTotal returns the AmountTotal field if non-nil, zero value otherwise.

### GetAmountTotalOk

`func (o *CheckoutSession) GetAmountTotalOk() (*float32, bool)`

GetAmountTotalOk returns a tuple with the AmountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotal

`func (o *CheckoutSession) SetAmountTotal(v float32)`

SetAmountTotal sets AmountTotal field to given value.


### SetAmountTotalNil

`func (o *CheckoutSession) SetAmountTotalNil(b bool)`

 SetAmountTotalNil sets the value for AmountTotal to be an explicit nil

### UnsetAmountTotal
`func (o *CheckoutSession) UnsetAmountTotal()`

UnsetAmountTotal ensures that no value is present for AmountTotal, not even an explicit nil
### GetCurrency

`func (o *CheckoutSession) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutSession) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutSession) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CheckoutSession) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CheckoutSession) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCustomerId

`func (o *CheckoutSession) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CheckoutSession) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CheckoutSession) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *CheckoutSession) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *CheckoutSession) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPaymentIntentId

`func (o *CheckoutSession) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *CheckoutSession) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *CheckoutSession) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.


### SetPaymentIntentIdNil

`func (o *CheckoutSession) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *CheckoutSession) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetInvoiceId

`func (o *CheckoutSession) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CheckoutSession) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CheckoutSession) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *CheckoutSession) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *CheckoutSession) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionId

`func (o *CheckoutSession) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CheckoutSession) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CheckoutSession) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *CheckoutSession) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CheckoutSession) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetUrl

`func (o *CheckoutSession) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckoutSession) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckoutSession) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *CheckoutSession) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CheckoutSession) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetSuccessUrl

`func (o *CheckoutSession) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutSession) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutSession) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.


### SetSuccessUrlNil

`func (o *CheckoutSession) SetSuccessUrlNil(b bool)`

 SetSuccessUrlNil sets the value for SuccessUrl to be an explicit nil

### UnsetSuccessUrl
`func (o *CheckoutSession) UnsetSuccessUrl()`

UnsetSuccessUrl ensures that no value is present for SuccessUrl, not even an explicit nil
### GetCancelUrl

`func (o *CheckoutSession) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *CheckoutSession) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *CheckoutSession) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### SetCancelUrlNil

`func (o *CheckoutSession) SetCancelUrlNil(b bool)`

 SetCancelUrlNil sets the value for CancelUrl to be an explicit nil

### UnsetCancelUrl
`func (o *CheckoutSession) UnsetCancelUrl()`

UnsetCancelUrl ensures that no value is present for CancelUrl, not even an explicit nil
### GetExpiresAt

`func (o *CheckoutSession) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutSession) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutSession) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *CheckoutSession) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CheckoutSession) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *CheckoutSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckoutSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckoutSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CheckoutSession) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CheckoutSession) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetAllowedMethods

`func (o *CheckoutSession) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *CheckoutSession) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *CheckoutSession) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.


### GetSelectedMethod

`func (o *CheckoutSession) GetSelectedMethod() string`

GetSelectedMethod returns the SelectedMethod field if non-nil, zero value otherwise.

### GetSelectedMethodOk

`func (o *CheckoutSession) GetSelectedMethodOk() (*string, bool)`

GetSelectedMethodOk returns a tuple with the SelectedMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedMethod

`func (o *CheckoutSession) SetSelectedMethod(v string)`

SetSelectedMethod sets SelectedMethod field to given value.


### SetSelectedMethodNil

`func (o *CheckoutSession) SetSelectedMethodNil(b bool)`

 SetSelectedMethodNil sets the value for SelectedMethod to be an explicit nil

### UnsetSelectedMethod
`func (o *CheckoutSession) UnsetSelectedMethod()`

UnsetSelectedMethod ensures that no value is present for SelectedMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


