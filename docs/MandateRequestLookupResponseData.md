# MandateRequestLookupResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mandate request ID | 
**MerchantId** | **string** | Merchant ID | 
**MerchantSlug** | **string** | Merchant slug | 
**CustomerId** | **string** | Customer ID | 
**InvoiceId** | **NullableString** | Invoice ID, if the request is tied to an invoice | 
**ExpiresAt** | **NullableString** | Expiry timestamp (ISO 8601) | 
**Status** | **string** | Mandate request status | 
**Merchant** | [**MandateRequestLookupResponseDataMerchant**](MandateRequestLookupResponseDataMerchant.md) |  | 
**Customer** | [**MandateRequestLookupResponseDataCustomer**](MandateRequestLookupResponseDataCustomer.md) |  | 
**Invoice** | [**MandateRequestLookupResponseDataInvoice**](MandateRequestLookupResponseDataInvoice.md) |  | 

## Methods

### NewMandateRequestLookupResponseData

`func NewMandateRequestLookupResponseData(id string, merchantId string, merchantSlug string, customerId string, invoiceId NullableString, expiresAt NullableString, status string, merchant MandateRequestLookupResponseDataMerchant, customer MandateRequestLookupResponseDataCustomer, invoice MandateRequestLookupResponseDataInvoice, ) *MandateRequestLookupResponseData`

NewMandateRequestLookupResponseData instantiates a new MandateRequestLookupResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateRequestLookupResponseDataWithDefaults

`func NewMandateRequestLookupResponseDataWithDefaults() *MandateRequestLookupResponseData`

NewMandateRequestLookupResponseDataWithDefaults instantiates a new MandateRequestLookupResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MandateRequestLookupResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MandateRequestLookupResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MandateRequestLookupResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *MandateRequestLookupResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *MandateRequestLookupResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *MandateRequestLookupResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantSlug

`func (o *MandateRequestLookupResponseData) GetMerchantSlug() string`

GetMerchantSlug returns the MerchantSlug field if non-nil, zero value otherwise.

### GetMerchantSlugOk

`func (o *MandateRequestLookupResponseData) GetMerchantSlugOk() (*string, bool)`

GetMerchantSlugOk returns a tuple with the MerchantSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantSlug

`func (o *MandateRequestLookupResponseData) SetMerchantSlug(v string)`

SetMerchantSlug sets MerchantSlug field to given value.


### GetCustomerId

`func (o *MandateRequestLookupResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MandateRequestLookupResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MandateRequestLookupResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *MandateRequestLookupResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *MandateRequestLookupResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *MandateRequestLookupResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *MandateRequestLookupResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *MandateRequestLookupResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetExpiresAt

`func (o *MandateRequestLookupResponseData) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MandateRequestLookupResponseData) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MandateRequestLookupResponseData) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *MandateRequestLookupResponseData) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *MandateRequestLookupResponseData) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetStatus

`func (o *MandateRequestLookupResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MandateRequestLookupResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MandateRequestLookupResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMerchant

`func (o *MandateRequestLookupResponseData) GetMerchant() MandateRequestLookupResponseDataMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MandateRequestLookupResponseData) GetMerchantOk() (*MandateRequestLookupResponseDataMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MandateRequestLookupResponseData) SetMerchant(v MandateRequestLookupResponseDataMerchant)`

SetMerchant sets Merchant field to given value.


### GetCustomer

`func (o *MandateRequestLookupResponseData) GetCustomer() MandateRequestLookupResponseDataCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *MandateRequestLookupResponseData) GetCustomerOk() (*MandateRequestLookupResponseDataCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *MandateRequestLookupResponseData) SetCustomer(v MandateRequestLookupResponseDataCustomer)`

SetCustomer sets Customer field to given value.


### GetInvoice

`func (o *MandateRequestLookupResponseData) GetInvoice() MandateRequestLookupResponseDataInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MandateRequestLookupResponseData) GetInvoiceOk() (*MandateRequestLookupResponseDataInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MandateRequestLookupResponseData) SetInvoice(v MandateRequestLookupResponseDataInvoice)`

SetInvoice sets Invoice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


