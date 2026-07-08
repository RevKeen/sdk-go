# DdMandateRequestLifecycleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**SubscriptionId** | **NullableString** |  | 
**MandateId** | **NullableString** |  | 
**Token** | **string** | Internal raw request token. Do not expose directly to customers. | 
**SignableUrl** | **string** | Customer-facing tracked URL: pay.revkeen.com/dd_&lt;token&gt; | 
**CreatedAt** | **NullableString** |  | 
**ExpiresAt** | **NullableString** |  | 
**ConsumedAt** | **NullableString** |  | 
**CancelledAt** | **NullableString** |  | 
**ExpiredAt** | **NullableString** |  | 
**Status** | [**DdMandateRequestLifecycleStatus**](DdMandateRequestLifecycleStatus.md) |  | 

## Methods

### NewDdMandateRequestLifecycleItem

`func NewDdMandateRequestLifecycleItem(id string, merchantId string, customerId string, invoiceId NullableString, subscriptionId NullableString, mandateId NullableString, token string, signableUrl string, createdAt NullableString, expiresAt NullableString, consumedAt NullableString, cancelledAt NullableString, expiredAt NullableString, status DdMandateRequestLifecycleStatus, ) *DdMandateRequestLifecycleItem`

NewDdMandateRequestLifecycleItem instantiates a new DdMandateRequestLifecycleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdMandateRequestLifecycleItemWithDefaults

`func NewDdMandateRequestLifecycleItemWithDefaults() *DdMandateRequestLifecycleItem`

NewDdMandateRequestLifecycleItemWithDefaults instantiates a new DdMandateRequestLifecycleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdMandateRequestLifecycleItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdMandateRequestLifecycleItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdMandateRequestLifecycleItem) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *DdMandateRequestLifecycleItem) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DdMandateRequestLifecycleItem) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DdMandateRequestLifecycleItem) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *DdMandateRequestLifecycleItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdMandateRequestLifecycleItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdMandateRequestLifecycleItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *DdMandateRequestLifecycleItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DdMandateRequestLifecycleItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DdMandateRequestLifecycleItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *DdMandateRequestLifecycleItem) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *DdMandateRequestLifecycleItem) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionId

`func (o *DdMandateRequestLifecycleItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DdMandateRequestLifecycleItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DdMandateRequestLifecycleItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *DdMandateRequestLifecycleItem) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *DdMandateRequestLifecycleItem) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetMandateId

`func (o *DdMandateRequestLifecycleItem) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DdMandateRequestLifecycleItem) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DdMandateRequestLifecycleItem) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### SetMandateIdNil

`func (o *DdMandateRequestLifecycleItem) SetMandateIdNil(b bool)`

 SetMandateIdNil sets the value for MandateId to be an explicit nil

### UnsetMandateId
`func (o *DdMandateRequestLifecycleItem) UnsetMandateId()`

UnsetMandateId ensures that no value is present for MandateId, not even an explicit nil
### GetToken

`func (o *DdMandateRequestLifecycleItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DdMandateRequestLifecycleItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DdMandateRequestLifecycleItem) SetToken(v string)`

SetToken sets Token field to given value.


### GetSignableUrl

`func (o *DdMandateRequestLifecycleItem) GetSignableUrl() string`

GetSignableUrl returns the SignableUrl field if non-nil, zero value otherwise.

### GetSignableUrlOk

`func (o *DdMandateRequestLifecycleItem) GetSignableUrlOk() (*string, bool)`

GetSignableUrlOk returns a tuple with the SignableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableUrl

`func (o *DdMandateRequestLifecycleItem) SetSignableUrl(v string)`

SetSignableUrl sets SignableUrl field to given value.


### GetCreatedAt

`func (o *DdMandateRequestLifecycleItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdMandateRequestLifecycleItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdMandateRequestLifecycleItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DdMandateRequestLifecycleItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdMandateRequestLifecycleItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *DdMandateRequestLifecycleItem) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DdMandateRequestLifecycleItem) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DdMandateRequestLifecycleItem) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *DdMandateRequestLifecycleItem) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DdMandateRequestLifecycleItem) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetConsumedAt

`func (o *DdMandateRequestLifecycleItem) GetConsumedAt() string`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *DdMandateRequestLifecycleItem) GetConsumedAtOk() (*string, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *DdMandateRequestLifecycleItem) SetConsumedAt(v string)`

SetConsumedAt sets ConsumedAt field to given value.


### SetConsumedAtNil

`func (o *DdMandateRequestLifecycleItem) SetConsumedAtNil(b bool)`

 SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil

### UnsetConsumedAt
`func (o *DdMandateRequestLifecycleItem) UnsetConsumedAt()`

UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
### GetCancelledAt

`func (o *DdMandateRequestLifecycleItem) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DdMandateRequestLifecycleItem) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DdMandateRequestLifecycleItem) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.


### SetCancelledAtNil

`func (o *DdMandateRequestLifecycleItem) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *DdMandateRequestLifecycleItem) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetExpiredAt

`func (o *DdMandateRequestLifecycleItem) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *DdMandateRequestLifecycleItem) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *DdMandateRequestLifecycleItem) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *DdMandateRequestLifecycleItem) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *DdMandateRequestLifecycleItem) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetStatus

`func (o *DdMandateRequestLifecycleItem) GetStatus() DdMandateRequestLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdMandateRequestLifecycleItem) GetStatusOk() (*DdMandateRequestLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdMandateRequestLifecycleItem) SetStatus(v DdMandateRequestLifecycleStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


