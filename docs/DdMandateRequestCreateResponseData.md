# DdMandateRequestCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**Token** | **string** | Signed mandate-request token for the emailed link | 
**SignableUrl** | **string** | Customer-facing URL that opens the sign flow | 
**CreatedAt** | **NullableString** |  | 
**ExpiresAt** | **NullableString** |  | 
**Status** | **string** |  | 
**Idempotent** | **bool** | True when an existing live request was returned instead of creating a new one | 

## Methods

### NewDdMandateRequestCreateResponseData

`func NewDdMandateRequestCreateResponseData(id string, merchantId string, customerId string, invoiceId NullableString, token string, signableUrl string, createdAt NullableString, expiresAt NullableString, status string, idempotent bool, ) *DdMandateRequestCreateResponseData`

NewDdMandateRequestCreateResponseData instantiates a new DdMandateRequestCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdMandateRequestCreateResponseDataWithDefaults

`func NewDdMandateRequestCreateResponseDataWithDefaults() *DdMandateRequestCreateResponseData`

NewDdMandateRequestCreateResponseDataWithDefaults instantiates a new DdMandateRequestCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdMandateRequestCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdMandateRequestCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdMandateRequestCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *DdMandateRequestCreateResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DdMandateRequestCreateResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DdMandateRequestCreateResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *DdMandateRequestCreateResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdMandateRequestCreateResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdMandateRequestCreateResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *DdMandateRequestCreateResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DdMandateRequestCreateResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DdMandateRequestCreateResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *DdMandateRequestCreateResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *DdMandateRequestCreateResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetToken

`func (o *DdMandateRequestCreateResponseData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DdMandateRequestCreateResponseData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DdMandateRequestCreateResponseData) SetToken(v string)`

SetToken sets Token field to given value.


### GetSignableUrl

`func (o *DdMandateRequestCreateResponseData) GetSignableUrl() string`

GetSignableUrl returns the SignableUrl field if non-nil, zero value otherwise.

### GetSignableUrlOk

`func (o *DdMandateRequestCreateResponseData) GetSignableUrlOk() (*string, bool)`

GetSignableUrlOk returns a tuple with the SignableUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableUrl

`func (o *DdMandateRequestCreateResponseData) SetSignableUrl(v string)`

SetSignableUrl sets SignableUrl field to given value.


### GetCreatedAt

`func (o *DdMandateRequestCreateResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdMandateRequestCreateResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdMandateRequestCreateResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DdMandateRequestCreateResponseData) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdMandateRequestCreateResponseData) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *DdMandateRequestCreateResponseData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DdMandateRequestCreateResponseData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DdMandateRequestCreateResponseData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *DdMandateRequestCreateResponseData) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DdMandateRequestCreateResponseData) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetStatus

`func (o *DdMandateRequestCreateResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdMandateRequestCreateResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdMandateRequestCreateResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdempotent

`func (o *DdMandateRequestCreateResponseData) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *DdMandateRequestCreateResponseData) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *DdMandateRequestCreateResponseData) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


