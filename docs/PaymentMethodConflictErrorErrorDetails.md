# PaymentMethodConflictErrorErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Why a detach was refused: the method is the customer&#39;s default, or an active, trialing or past_due subscription bills on it. | [optional]
**SubscriptionIds** | Pointer to **[]string** | The subscriptions that bill on this method (reason &#x60;active_subscription&#x60;). | [optional]

## Methods

### NewPaymentMethodConflictErrorErrorDetails

`func NewPaymentMethodConflictErrorErrorDetails() *PaymentMethodConflictErrorErrorDetails`

NewPaymentMethodConflictErrorErrorDetails instantiates a new PaymentMethodConflictErrorErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodConflictErrorErrorDetailsWithDefaults

`func NewPaymentMethodConflictErrorErrorDetailsWithDefaults() *PaymentMethodConflictErrorErrorDetails`

NewPaymentMethodConflictErrorErrorDetailsWithDefaults instantiates a new PaymentMethodConflictErrorErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PaymentMethodConflictErrorErrorDetails) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PaymentMethodConflictErrorErrorDetails) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PaymentMethodConflictErrorErrorDetails) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PaymentMethodConflictErrorErrorDetails) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *PaymentMethodConflictErrorErrorDetails) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *PaymentMethodConflictErrorErrorDetails) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *PaymentMethodConflictErrorErrorDetails) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *PaymentMethodConflictErrorErrorDetails) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
