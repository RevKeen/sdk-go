# PaymentIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Public payment intent ID (pi_xxx) | 
**Object** | **string** | Object type | 
**Amount** | **int32** | Amount in cents | 
**AmountCapturable** | **int32** | Amount that can be captured (for manual capture) | 
**AmountReceived** | **int32** | Amount actually received | 
**Currency** | **string** | Three-letter ISO currency code | 
**Customer** | **string** | Customer ID | 
**Description** | Pointer to **NullableString** | Description for merchant reference | [optional] 
**LastPaymentError** | Pointer to [**PaymentIntentLastPaymentError**](PaymentIntentLastPaymentError.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**NextAction** | Pointer to [**PaymentIntentNextAction**](PaymentIntentNextAction.md) |  | [optional] 
**PaymentMethod** | Pointer to **NullableString** | Payment method ID | [optional] 
**ReceiptEmail** | Pointer to **NullableString** | Email for receipt | [optional] 
**StatementDescriptor** | Pointer to **NullableString** | Statement descriptor | [optional] 
**StatementDescriptorSuffix** | Pointer to **NullableString** | Statement descriptor suffix | [optional] 
**Status** | **string** | Payment intent status. requires_payment_method: Needs payment method. requires_confirmation: Ready to confirm. requires_action: Requires customer action (3DS). processing: Being processed. succeeded: Payment complete. canceled: Canceled. | 
**CaptureMethod** | **string** | Capture method. automatic: Capture immediately on confirmation. manual: Authorize then capture separately. | 
**ClientSecret** | **string** | Client secret for frontend confirmation | 
**CanceledAt** | Pointer to **NullableTime** | When the intent was canceled | [optional] 
**CancellationReason** | Pointer to **NullableString** | Why the intent was canceled | [optional] 
**Created** | **int32** | Unix timestamp of creation | 
**Livemode** | **bool** | Whether in live mode | 

## Methods

### NewPaymentIntent

`func NewPaymentIntent(id string, object string, amount int32, amountCapturable int32, amountReceived int32, currency string, customer string, status string, captureMethod string, clientSecret string, created int32, livemode bool, ) *PaymentIntent`

NewPaymentIntent instantiates a new PaymentIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentWithDefaults

`func NewPaymentIntentWithDefaults() *PaymentIntent`

NewPaymentIntentWithDefaults instantiates a new PaymentIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentIntent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentIntent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentIntent) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentIntent) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentIntent) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentIntent) SetObject(v string)`

SetObject sets Object field to given value.


### GetAmount

`func (o *PaymentIntent) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentIntent) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentIntent) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountCapturable

`func (o *PaymentIntent) GetAmountCapturable() int32`

GetAmountCapturable returns the AmountCapturable field if non-nil, zero value otherwise.

### GetAmountCapturableOk

`func (o *PaymentIntent) GetAmountCapturableOk() (*int32, bool)`

GetAmountCapturableOk returns a tuple with the AmountCapturable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCapturable

`func (o *PaymentIntent) SetAmountCapturable(v int32)`

SetAmountCapturable sets AmountCapturable field to given value.


### GetAmountReceived

`func (o *PaymentIntent) GetAmountReceived() int32`

GetAmountReceived returns the AmountReceived field if non-nil, zero value otherwise.

### GetAmountReceivedOk

`func (o *PaymentIntent) GetAmountReceivedOk() (*int32, bool)`

GetAmountReceivedOk returns a tuple with the AmountReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountReceived

`func (o *PaymentIntent) SetAmountReceived(v int32)`

SetAmountReceived sets AmountReceived field to given value.


### GetCurrency

`func (o *PaymentIntent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomer

`func (o *PaymentIntent) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *PaymentIntent) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *PaymentIntent) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetDescription

`func (o *PaymentIntent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentIntent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentIntent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentIntent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentIntent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentIntent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLastPaymentError

`func (o *PaymentIntent) GetLastPaymentError() PaymentIntentLastPaymentError`

GetLastPaymentError returns the LastPaymentError field if non-nil, zero value otherwise.

### GetLastPaymentErrorOk

`func (o *PaymentIntent) GetLastPaymentErrorOk() (*PaymentIntentLastPaymentError, bool)`

GetLastPaymentErrorOk returns a tuple with the LastPaymentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentError

`func (o *PaymentIntent) SetLastPaymentError(v PaymentIntentLastPaymentError)`

SetLastPaymentError sets LastPaymentError field to given value.

### HasLastPaymentError

`func (o *PaymentIntent) HasLastPaymentError() bool`

HasLastPaymentError returns a boolean if a field has been set.

### GetMetadata

`func (o *PaymentIntent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentIntent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentIntent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PaymentIntent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNextAction

`func (o *PaymentIntent) GetNextAction() PaymentIntentNextAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *PaymentIntent) GetNextActionOk() (*PaymentIntentNextAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *PaymentIntent) SetNextAction(v PaymentIntentNextAction)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *PaymentIntent) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *PaymentIntent) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentIntent) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentIntent) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentIntent) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *PaymentIntent) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *PaymentIntent) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetReceiptEmail

`func (o *PaymentIntent) GetReceiptEmail() string`

GetReceiptEmail returns the ReceiptEmail field if non-nil, zero value otherwise.

### GetReceiptEmailOk

`func (o *PaymentIntent) GetReceiptEmailOk() (*string, bool)`

GetReceiptEmailOk returns a tuple with the ReceiptEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptEmail

`func (o *PaymentIntent) SetReceiptEmail(v string)`

SetReceiptEmail sets ReceiptEmail field to given value.

### HasReceiptEmail

`func (o *PaymentIntent) HasReceiptEmail() bool`

HasReceiptEmail returns a boolean if a field has been set.

### SetReceiptEmailNil

`func (o *PaymentIntent) SetReceiptEmailNil(b bool)`

 SetReceiptEmailNil sets the value for ReceiptEmail to be an explicit nil

### UnsetReceiptEmail
`func (o *PaymentIntent) UnsetReceiptEmail()`

UnsetReceiptEmail ensures that no value is present for ReceiptEmail, not even an explicit nil
### GetStatementDescriptor

`func (o *PaymentIntent) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *PaymentIntent) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *PaymentIntent) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *PaymentIntent) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### SetStatementDescriptorNil

`func (o *PaymentIntent) SetStatementDescriptorNil(b bool)`

 SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil

### UnsetStatementDescriptor
`func (o *PaymentIntent) UnsetStatementDescriptor()`

UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
### GetStatementDescriptorSuffix

`func (o *PaymentIntent) GetStatementDescriptorSuffix() string`

GetStatementDescriptorSuffix returns the StatementDescriptorSuffix field if non-nil, zero value otherwise.

### GetStatementDescriptorSuffixOk

`func (o *PaymentIntent) GetStatementDescriptorSuffixOk() (*string, bool)`

GetStatementDescriptorSuffixOk returns a tuple with the StatementDescriptorSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptorSuffix

`func (o *PaymentIntent) SetStatementDescriptorSuffix(v string)`

SetStatementDescriptorSuffix sets StatementDescriptorSuffix field to given value.

### HasStatementDescriptorSuffix

`func (o *PaymentIntent) HasStatementDescriptorSuffix() bool`

HasStatementDescriptorSuffix returns a boolean if a field has been set.

### SetStatementDescriptorSuffixNil

`func (o *PaymentIntent) SetStatementDescriptorSuffixNil(b bool)`

 SetStatementDescriptorSuffixNil sets the value for StatementDescriptorSuffix to be an explicit nil

### UnsetStatementDescriptorSuffix
`func (o *PaymentIntent) UnsetStatementDescriptorSuffix()`

UnsetStatementDescriptorSuffix ensures that no value is present for StatementDescriptorSuffix, not even an explicit nil
### GetStatus

`func (o *PaymentIntent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentIntent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentIntent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCaptureMethod

`func (o *PaymentIntent) GetCaptureMethod() string`

GetCaptureMethod returns the CaptureMethod field if non-nil, zero value otherwise.

### GetCaptureMethodOk

`func (o *PaymentIntent) GetCaptureMethodOk() (*string, bool)`

GetCaptureMethodOk returns a tuple with the CaptureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMethod

`func (o *PaymentIntent) SetCaptureMethod(v string)`

SetCaptureMethod sets CaptureMethod field to given value.


### GetClientSecret

`func (o *PaymentIntent) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PaymentIntent) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PaymentIntent) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCanceledAt

`func (o *PaymentIntent) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *PaymentIntent) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *PaymentIntent) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *PaymentIntent) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *PaymentIntent) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *PaymentIntent) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetCancellationReason

`func (o *PaymentIntent) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *PaymentIntent) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *PaymentIntent) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *PaymentIntent) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReasonNil

`func (o *PaymentIntent) SetCancellationReasonNil(b bool)`

 SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil

### UnsetCancellationReason
`func (o *PaymentIntent) UnsetCancellationReason()`

UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
### GetCreated

`func (o *PaymentIntent) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentIntent) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentIntent) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetLivemode

`func (o *PaymentIntent) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *PaymentIntent) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *PaymentIntent) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


