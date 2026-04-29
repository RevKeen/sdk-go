# SetupIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the setup intent | 
**Object** | **string** | Object type, always &#39;setup_intent&#39; | 
**PublicId** | **string** | Public ID visible in API responses (seti_xxx format) | 
**Status** | **string** | The status of the setup intent | 
**CustomerId** | **NullableString** | ID of the customer this setup intent is for | 
**PaymentMethodId** | **NullableString** | ID of the payment method being set up | 
**PaymentMethodTypes** | **[]string** | Allowed payment method types for this setup | 
**Usage** | **string** | Indicates how the payment method will be used | 
**NextAction** | [**SetupIntentNextAction**](SetupIntentNextAction.md) |  | 
**ClientSecret** | **string** | Client secret for frontend confirmation | 
**Gateway** | **string** | Name of the payment processor that handled this setup intent | 
**LastError** | [**SetupIntentError**](SetupIntentError.md) |  | 
**CancellationReason** | **NullableString** | Reason for cancellation if canceled | 
**Description** | **NullableString** | Merchant description for reference | 
**Metadata** | **map[string]interface{}** | Custom metadata attached to the setup intent | 
**ConfirmedAt** | **NullableTime** | When the setup intent was confirmed | 
**CanceledAt** | **NullableTime** | When the setup intent was canceled | 
**CreatedAt** | **time.Time** | When the setup intent was created | 
**UpdatedAt** | **time.Time** | When the setup intent was last updated | 

## Methods

### NewSetupIntent

`func NewSetupIntent(id string, object string, publicId string, status string, customerId NullableString, paymentMethodId NullableString, paymentMethodTypes []string, usage string, nextAction SetupIntentNextAction, clientSecret string, gateway string, lastError SetupIntentError, cancellationReason NullableString, description NullableString, metadata map[string]interface{}, confirmedAt NullableTime, canceledAt NullableTime, createdAt time.Time, updatedAt time.Time, ) *SetupIntent`

NewSetupIntent instantiates a new SetupIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupIntentWithDefaults

`func NewSetupIntentWithDefaults() *SetupIntent`

NewSetupIntentWithDefaults instantiates a new SetupIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetupIntent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetupIntent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetupIntent) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *SetupIntent) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SetupIntent) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SetupIntent) SetObject(v string)`

SetObject sets Object field to given value.


### GetPublicId

`func (o *SetupIntent) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SetupIntent) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SetupIntent) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetStatus

`func (o *SetupIntent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetupIntent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetupIntent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomerId

`func (o *SetupIntent) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SetupIntent) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SetupIntent) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *SetupIntent) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *SetupIntent) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetPaymentMethodId

`func (o *SetupIntent) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *SetupIntent) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *SetupIntent) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.


### SetPaymentMethodIdNil

`func (o *SetupIntent) SetPaymentMethodIdNil(b bool)`

 SetPaymentMethodIdNil sets the value for PaymentMethodId to be an explicit nil

### UnsetPaymentMethodId
`func (o *SetupIntent) UnsetPaymentMethodId()`

UnsetPaymentMethodId ensures that no value is present for PaymentMethodId, not even an explicit nil
### GetPaymentMethodTypes

`func (o *SetupIntent) GetPaymentMethodTypes() []string`

GetPaymentMethodTypes returns the PaymentMethodTypes field if non-nil, zero value otherwise.

### GetPaymentMethodTypesOk

`func (o *SetupIntent) GetPaymentMethodTypesOk() (*[]string, bool)`

GetPaymentMethodTypesOk returns a tuple with the PaymentMethodTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTypes

`func (o *SetupIntent) SetPaymentMethodTypes(v []string)`

SetPaymentMethodTypes sets PaymentMethodTypes field to given value.


### GetUsage

`func (o *SetupIntent) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SetupIntent) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SetupIntent) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetNextAction

`func (o *SetupIntent) GetNextAction() SetupIntentNextAction`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *SetupIntent) GetNextActionOk() (*SetupIntentNextAction, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *SetupIntent) SetNextAction(v SetupIntentNextAction)`

SetNextAction sets NextAction field to given value.


### GetClientSecret

`func (o *SetupIntent) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SetupIntent) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SetupIntent) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGateway

`func (o *SetupIntent) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *SetupIntent) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *SetupIntent) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetLastError

`func (o *SetupIntent) GetLastError() SetupIntentError`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *SetupIntent) GetLastErrorOk() (*SetupIntentError, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *SetupIntent) SetLastError(v SetupIntentError)`

SetLastError sets LastError field to given value.


### GetCancellationReason

`func (o *SetupIntent) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *SetupIntent) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *SetupIntent) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.


### SetCancellationReasonNil

`func (o *SetupIntent) SetCancellationReasonNil(b bool)`

 SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil

### UnsetCancellationReason
`func (o *SetupIntent) UnsetCancellationReason()`

UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
### GetDescription

`func (o *SetupIntent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetupIntent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetupIntent) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *SetupIntent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SetupIntent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *SetupIntent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SetupIntent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SetupIntent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetConfirmedAt

`func (o *SetupIntent) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *SetupIntent) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *SetupIntent) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.


### SetConfirmedAtNil

`func (o *SetupIntent) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *SetupIntent) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetCanceledAt

`func (o *SetupIntent) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SetupIntent) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SetupIntent) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.


### SetCanceledAtNil

`func (o *SetupIntent) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *SetupIntent) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetCreatedAt

`func (o *SetupIntent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SetupIntent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SetupIntent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SetupIntent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SetupIntent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SetupIntent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


