# DdMandateRequestActiveItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**MerchantId** | **string** |  |
**CustomerId** | **string** |  |
**InvoiceId** | **string** |  |
**SubscriptionId** | **string** |  |
**MandateId** | **string** |  |
**Token** | **string** | Internal raw request token. Do not expose directly to customers. |
**SignableUrl** | **string** | Customer-facing tracked URL: pay.revkeen.com/dd_&lt;token&gt; |
**CreatedAt** | **time.Time** |  |
**ExpiresAt** | **time.Time** |  |
**ConsumedAt** | **time.Time** |  |
**CancelledAt** | **time.Time** |  |
**ExpiredAt** | **time.Time** |  |
**Status** | [**DdMandateRequestLifecycleStatus**](DdMandateRequestLifecycleStatus.md) |  |
**NotificationDeliveryEventType** | **string** |  |
**NotificationDeliveryStatus** | **string** |  |
**NotificationDeliveryError** | **string** |  |
**NotificationDeliveryUpdatedAt** | **time.Time** |  |
**PdfDocuments** | [**[]DdMandatePdfReference**](DdMandatePdfReference.md) |  |
**OpenedAt** | **NullableTime** | When the customer first opened the request link (its first capture session). Null until opened. |

## Methods

### NewDdMandateRequestActiveItem

`func NewDdMandateRequestActiveItem(id string, merchantId string, customerId string, invoiceId string, subscriptionId string, mandateId string, token string, signableUrl string, createdAt time.Time, expiresAt time.Time, consumedAt time.Time, cancelledAt time.Time, expiredAt time.Time, status DdMandateRequestLifecycleStatus, notificationDeliveryEventType string, notificationDeliveryStatus string, notificationDeliveryError string, notificationDeliveryUpdatedAt time.Time, pdfDocuments []DdMandatePdfReference, openedAt NullableTime, ) *DdMandateRequestActiveItem`

NewDdMandateRequestActiveItem instantiates a new DdMandateRequestActiveItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdMandateRequestActiveItemWithDefaults

`func NewDdMandateRequestActiveItemWithDefaults() *DdMandateRequestActiveItem`

NewDdMandateRequestActiveItemWithDefaults instantiates a new DdMandateRequestActiveItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdMandateRequestActiveItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdMandateRequestActiveItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdMandateRequestActiveItem) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *DdMandateRequestActiveItem) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DdMandateRequestActiveItem) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DdMandateRequestActiveItem) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *DdMandateRequestActiveItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdMandateRequestActiveItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdMandateRequestActiveItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *DdMandateRequestActiveItem) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DdMandateRequestActiveItem) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DdMandateRequestActiveItem) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSubscriptionId

`func (o *DdMandateRequestActiveItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DdMandateRequestActiveItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DdMandateRequestActiveItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetMandateId

`func (o *DdMandateRequestActiveItem) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DdMandateRequestActiveItem) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DdMandateRequestActiveItem) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetToken

`func (o *DdMandateRequestActiveItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DdMandateRequestActiveItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DdMandateRequestActiveItem) SetToken(v string)`

SetToken sets Token field to given value.


### GetSignableUrl

`func (o *DdMandateRequestActiveItem) GetSignableUrl() string`

GetSignableUrl returns the SignableUrl field if non-nil, zero value otherwise.

### GetSignableUrlOk

`func (o *DdMandateRequestActiveItem) GetSignableUrlOk() (*string, bool)`

GetSignableUrlOk returns a tuple with the SignableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableUrl

`func (o *DdMandateRequestActiveItem) SetSignableUrl(v string)`

SetSignableUrl sets SignableUrl field to given value.


### GetCreatedAt

`func (o *DdMandateRequestActiveItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdMandateRequestActiveItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdMandateRequestActiveItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *DdMandateRequestActiveItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DdMandateRequestActiveItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DdMandateRequestActiveItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetConsumedAt

`func (o *DdMandateRequestActiveItem) GetConsumedAt() time.Time`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *DdMandateRequestActiveItem) GetConsumedAtOk() (*time.Time, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *DdMandateRequestActiveItem) SetConsumedAt(v time.Time)`

SetConsumedAt sets ConsumedAt field to given value.


### GetCancelledAt

`func (o *DdMandateRequestActiveItem) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DdMandateRequestActiveItem) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DdMandateRequestActiveItem) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.


### GetExpiredAt

`func (o *DdMandateRequestActiveItem) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *DdMandateRequestActiveItem) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *DdMandateRequestActiveItem) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.


### GetStatus

`func (o *DdMandateRequestActiveItem) GetStatus() DdMandateRequestLifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdMandateRequestActiveItem) GetStatusOk() (*DdMandateRequestLifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdMandateRequestActiveItem) SetStatus(v DdMandateRequestLifecycleStatus)`

SetStatus sets Status field to given value.


### GetNotificationDeliveryEventType

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryEventType() string`

GetNotificationDeliveryEventType returns the NotificationDeliveryEventType field if non-nil, zero value otherwise.

### GetNotificationDeliveryEventTypeOk

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryEventTypeOk() (*string, bool)`

GetNotificationDeliveryEventTypeOk returns a tuple with the NotificationDeliveryEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDeliveryEventType

`func (o *DdMandateRequestActiveItem) SetNotificationDeliveryEventType(v string)`

SetNotificationDeliveryEventType sets NotificationDeliveryEventType field to given value.


### GetNotificationDeliveryStatus

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryStatus() string`

GetNotificationDeliveryStatus returns the NotificationDeliveryStatus field if non-nil, zero value otherwise.

### GetNotificationDeliveryStatusOk

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryStatusOk() (*string, bool)`

GetNotificationDeliveryStatusOk returns a tuple with the NotificationDeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDeliveryStatus

`func (o *DdMandateRequestActiveItem) SetNotificationDeliveryStatus(v string)`

SetNotificationDeliveryStatus sets NotificationDeliveryStatus field to given value.


### GetNotificationDeliveryError

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryError() string`

GetNotificationDeliveryError returns the NotificationDeliveryError field if non-nil, zero value otherwise.

### GetNotificationDeliveryErrorOk

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryErrorOk() (*string, bool)`

GetNotificationDeliveryErrorOk returns a tuple with the NotificationDeliveryError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDeliveryError

`func (o *DdMandateRequestActiveItem) SetNotificationDeliveryError(v string)`

SetNotificationDeliveryError sets NotificationDeliveryError field to given value.


### GetNotificationDeliveryUpdatedAt

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryUpdatedAt() time.Time`

GetNotificationDeliveryUpdatedAt returns the NotificationDeliveryUpdatedAt field if non-nil, zero value otherwise.

### GetNotificationDeliveryUpdatedAtOk

`func (o *DdMandateRequestActiveItem) GetNotificationDeliveryUpdatedAtOk() (*time.Time, bool)`

GetNotificationDeliveryUpdatedAtOk returns a tuple with the NotificationDeliveryUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDeliveryUpdatedAt

`func (o *DdMandateRequestActiveItem) SetNotificationDeliveryUpdatedAt(v time.Time)`

SetNotificationDeliveryUpdatedAt sets NotificationDeliveryUpdatedAt field to given value.


### GetPdfDocuments

`func (o *DdMandateRequestActiveItem) GetPdfDocuments() []DdMandatePdfReference`

GetPdfDocuments returns the PdfDocuments field if non-nil, zero value otherwise.

### GetPdfDocumentsOk

`func (o *DdMandateRequestActiveItem) GetPdfDocumentsOk() (*[]DdMandatePdfReference, bool)`

GetPdfDocumentsOk returns a tuple with the PdfDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfDocuments

`func (o *DdMandateRequestActiveItem) SetPdfDocuments(v []DdMandatePdfReference)`

SetPdfDocuments sets PdfDocuments field to given value.


### GetOpenedAt

`func (o *DdMandateRequestActiveItem) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *DdMandateRequestActiveItem) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *DdMandateRequestActiveItem) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.


### SetOpenedAtNil

`func (o *DdMandateRequestActiveItem) SetOpenedAtNil(b bool)`

 SetOpenedAtNil sets the value for OpenedAt to be an explicit nil

### UnsetOpenedAt
`func (o *DdMandateRequestActiveItem) UnsetOpenedAt()`

UnsetOpenedAt ensures that no value is present for OpenedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
