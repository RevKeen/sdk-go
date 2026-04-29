# CustomerCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique customer identifier | 
**MerchantId** | **string** | ID of the merchant this customer belongs to | 
**MerchantRefId** | Pointer to **string** | Merchant-assigned reference ID for external system mapping | [optional] 
**Email** | **string** | Customer&#39;s email address | 
**Name** | Pointer to **string** | Customer&#39;s full name | [optional] 
**Phone** | Pointer to **string** | Customer&#39;s phone number | [optional] 
**CompanyName** | Pointer to **string** | Customer&#39;s company name | [optional] 
**AddressLine1** | Pointer to **string** | Billing address line 1 | [optional] 
**AddressLine2** | Pointer to **string** | Billing address line 2 | [optional] 
**City** | Pointer to **string** | Billing address city | [optional] 
**State** | Pointer to **string** | Billing address state or county | [optional] 
**PostalCode** | Pointer to **string** | Billing address postal code | [optional] 
**Country** | Pointer to **string** | Billing address country code | [optional] 
**AuthUserId** | Pointer to **string** | Linked Better Auth user ID for portal access | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata attached to this customer | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Merchant-defined custom fields displayed as columns in the dashboard | [optional] 
**CreatedAt** | **time.Time** | When the customer was created (ISO 8601) | 
**UpdatedAt** | **time.Time** | When the customer was last updated (ISO 8601) | 

## Methods

### NewCustomerCreateResponseData

`func NewCustomerCreateResponseData(id string, merchantId string, email string, createdAt time.Time, updatedAt time.Time, ) *CustomerCreateResponseData`

NewCustomerCreateResponseData instantiates a new CustomerCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerCreateResponseDataWithDefaults

`func NewCustomerCreateResponseDataWithDefaults() *CustomerCreateResponseData`

NewCustomerCreateResponseDataWithDefaults instantiates a new CustomerCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *CustomerCreateResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CustomerCreateResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CustomerCreateResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetMerchantRefId

`func (o *CustomerCreateResponseData) GetMerchantRefId() string`

GetMerchantRefId returns the MerchantRefId field if non-nil, zero value otherwise.

### GetMerchantRefIdOk

`func (o *CustomerCreateResponseData) GetMerchantRefIdOk() (*string, bool)`

GetMerchantRefIdOk returns a tuple with the MerchantRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantRefId

`func (o *CustomerCreateResponseData) SetMerchantRefId(v string)`

SetMerchantRefId sets MerchantRefId field to given value.

### HasMerchantRefId

`func (o *CustomerCreateResponseData) HasMerchantRefId() bool`

HasMerchantRefId returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerCreateResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerCreateResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerCreateResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CustomerCreateResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerCreateResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerCreateResponseData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerCreateResponseData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerCreateResponseData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerCreateResponseData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerCreateResponseData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerCreateResponseData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomerCreateResponseData) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomerCreateResponseData) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomerCreateResponseData) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomerCreateResponseData) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CustomerCreateResponseData) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomerCreateResponseData) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomerCreateResponseData) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomerCreateResponseData) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CustomerCreateResponseData) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomerCreateResponseData) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomerCreateResponseData) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomerCreateResponseData) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *CustomerCreateResponseData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerCreateResponseData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerCreateResponseData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomerCreateResponseData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *CustomerCreateResponseData) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerCreateResponseData) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerCreateResponseData) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerCreateResponseData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerCreateResponseData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerCreateResponseData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerCreateResponseData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomerCreateResponseData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerCreateResponseData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerCreateResponseData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerCreateResponseData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerCreateResponseData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAuthUserId

`func (o *CustomerCreateResponseData) GetAuthUserId() string`

GetAuthUserId returns the AuthUserId field if non-nil, zero value otherwise.

### GetAuthUserIdOk

`func (o *CustomerCreateResponseData) GetAuthUserIdOk() (*string, bool)`

GetAuthUserIdOk returns a tuple with the AuthUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUserId

`func (o *CustomerCreateResponseData) SetAuthUserId(v string)`

SetAuthUserId sets AuthUserId field to given value.

### HasAuthUserId

`func (o *CustomerCreateResponseData) HasAuthUserId() bool`

HasAuthUserId returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerCreateResponseData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerCreateResponseData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerCreateResponseData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerCreateResponseData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomerCreateResponseData) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomerCreateResponseData) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomerCreateResponseData) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomerCreateResponseData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CustomerCreateResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerCreateResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerCreateResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomerCreateResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerCreateResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerCreateResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


