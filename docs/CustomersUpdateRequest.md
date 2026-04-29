# CustomersUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Updated email address | [optional] 
**Name** | Pointer to **string** | Updated full name | [optional] 
**Phone** | Pointer to **string** | Updated phone number | [optional] 
**CompanyName** | Pointer to **string** | Updated company name | [optional] 
**AddressLine1** | Pointer to **string** | Updated billing address line 1 | [optional] 
**AddressLine2** | Pointer to **string** | Updated billing address line 2 | [optional] 
**City** | Pointer to **string** | Updated billing address city | [optional] 
**State** | Pointer to **string** | Updated billing address state or county | [optional] 
**PostalCode** | Pointer to **string** | Updated billing address postal code | [optional] 
**Country** | Pointer to **string** | Updated billing address country code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata to merge with existing values (set a key to null to remove it) | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Merchant-defined custom fields displayed as columns in the dashboard | [optional] 

## Methods

### NewCustomersUpdateRequest

`func NewCustomersUpdateRequest() *CustomersUpdateRequest`

NewCustomersUpdateRequest instantiates a new CustomersUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersUpdateRequestWithDefaults

`func NewCustomersUpdateRequestWithDefaults() *CustomersUpdateRequest`

NewCustomersUpdateRequestWithDefaults instantiates a new CustomersUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomersUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomersUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomersUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomersUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CustomersUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomersUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomersUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomersUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *CustomersUpdateRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomersUpdateRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomersUpdateRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomersUpdateRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetCompanyName

`func (o *CustomersUpdateRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CustomersUpdateRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CustomersUpdateRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CustomersUpdateRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *CustomersUpdateRequest) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CustomersUpdateRequest) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CustomersUpdateRequest) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CustomersUpdateRequest) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CustomersUpdateRequest) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CustomersUpdateRequest) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CustomersUpdateRequest) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CustomersUpdateRequest) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *CustomersUpdateRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomersUpdateRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomersUpdateRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CustomersUpdateRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *CustomersUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomersUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomersUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomersUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomersUpdateRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomersUpdateRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomersUpdateRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CustomersUpdateRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *CustomersUpdateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomersUpdateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomersUpdateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomersUpdateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomersUpdateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomersUpdateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomersUpdateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomersUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomersUpdateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomersUpdateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomersUpdateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomersUpdateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


