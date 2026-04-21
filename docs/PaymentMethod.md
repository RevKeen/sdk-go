# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the payment method | 
**Object** | **string** | Object type, always &#39;payment_method&#39; | 
**PublicId** | **NullableString** | Public ID visible in API responses (pm_xxx format) | 
**Type** | **string** | The type of payment method | 
**Status** | **string** | The status of the payment method | 
**CustomerId** | **string** | ID of the customer this payment method belongs to | 
**IsDefault** | **bool** | Whether this is the customer&#39;s default payment method | 
**Card** | [**CardDetails**](CardDetails.md) |  | 
**UsBankAccount** | [**BankAccountDetails**](BankAccountDetails.md) |  | 
**BillingDetails** | [**BillingDetails**](BillingDetails.md) |  | 
**Metadata** | **map[string]interface{}** | Custom metadata attached to the payment method | 
**CreatedAt** | **time.Time** | When the payment method was created | 
**UpdatedAt** | **time.Time** | When the payment method was last updated | 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(id string, object string, publicId NullableString, type_ string, status string, customerId string, isDefault bool, card CardDetails, usBankAccount BankAccountDetails, billingDetails BillingDetails, metadata map[string]interface{}, createdAt time.Time, updatedAt time.Time, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.


### GetPublicId

`func (o *PaymentMethod) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *PaymentMethod) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *PaymentMethod) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### SetPublicIdNil

`func (o *PaymentMethod) SetPublicIdNil(b bool)`

 SetPublicIdNil sets the value for PublicId to be an explicit nil

### UnsetPublicId
`func (o *PaymentMethod) UnsetPublicId()`

UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil
### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.


### GetStatus

`func (o *PaymentMethod) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentMethod) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentMethod) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCustomerId

`func (o *PaymentMethod) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *PaymentMethod) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *PaymentMethod) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetIsDefault

`func (o *PaymentMethod) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethod) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCard

`func (o *PaymentMethod) GetCard() CardDetails`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethod) GetCardOk() (*CardDetails, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethod) SetCard(v CardDetails)`

SetCard sets Card field to given value.


### GetUsBankAccount

`func (o *PaymentMethod) GetUsBankAccount() BankAccountDetails`

GetUsBankAccount returns the UsBankAccount field if non-nil, zero value otherwise.

### GetUsBankAccountOk

`func (o *PaymentMethod) GetUsBankAccountOk() (*BankAccountDetails, bool)`

GetUsBankAccountOk returns a tuple with the UsBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsBankAccount

`func (o *PaymentMethod) SetUsBankAccount(v BankAccountDetails)`

SetUsBankAccount sets UsBankAccount field to given value.


### GetBillingDetails

`func (o *PaymentMethod) GetBillingDetails() BillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *PaymentMethod) GetBillingDetailsOk() (*BillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *PaymentMethod) SetBillingDetails(v BillingDetails)`

SetBillingDetails sets BillingDetails field to given value.


### GetMetadata

`func (o *PaymentMethod) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PaymentMethod) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PaymentMethod) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetCreatedAt

`func (o *PaymentMethod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentMethod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PaymentMethod) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentMethod) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


