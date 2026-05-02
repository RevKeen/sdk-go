# Business

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier (UUID) | 
**Object** | **string** | Object type | 
**CustomerId** | Pointer to **NullableString** | Associated customer ID | [optional] 
**Name** | **string** | Business name | 
**CompanyNumber** | Pointer to **NullableString** | Company registration number | [optional] 
**TaxIdentifier** | Pointer to **NullableString** | Tax ID (VAT, EIN, etc.) | [optional] 
**BillingEmail** | Pointer to **NullableString** | Billing email address | [optional] 
**BillingPhone** | Pointer to **NullableString** | Billing phone number | [optional] 
**BillingAddressLine1** | Pointer to **NullableString** | Billing address line 1 | [optional] 
**BillingAddressLine2** | Pointer to **NullableString** | Billing address line 2 | [optional] 
**BillingCity** | Pointer to **NullableString** | Billing city | [optional] 
**BillingPostcode** | Pointer to **NullableString** | Billing postcode/ZIP | [optional] 
**BillingCountry** | Pointer to **NullableString** | Billing country (ISO 3166-1 alpha-2) | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp (ISO 8601) | 
**UpdatedAt** | **time.Time** | Last update timestamp (ISO 8601) | 

## Methods

### NewBusiness

`func NewBusiness(id string, object string, name string, createdAt time.Time, updatedAt time.Time, ) *Business`

NewBusiness instantiates a new Business object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessWithDefaults

`func NewBusinessWithDefaults() *Business`

NewBusinessWithDefaults instantiates a new Business object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Business) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Business) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Business) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Business) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Business) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Business) SetObject(v string)`

SetObject sets Object field to given value.


### GetCustomerId

`func (o *Business) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Business) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Business) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Business) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Business) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Business) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetName

`func (o *Business) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Business) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Business) SetName(v string)`

SetName sets Name field to given value.


### GetCompanyNumber

`func (o *Business) GetCompanyNumber() string`

GetCompanyNumber returns the CompanyNumber field if non-nil, zero value otherwise.

### GetCompanyNumberOk

`func (o *Business) GetCompanyNumberOk() (*string, bool)`

GetCompanyNumberOk returns a tuple with the CompanyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNumber

`func (o *Business) SetCompanyNumber(v string)`

SetCompanyNumber sets CompanyNumber field to given value.

### HasCompanyNumber

`func (o *Business) HasCompanyNumber() bool`

HasCompanyNumber returns a boolean if a field has been set.

### SetCompanyNumberNil

`func (o *Business) SetCompanyNumberNil(b bool)`

 SetCompanyNumberNil sets the value for CompanyNumber to be an explicit nil

### UnsetCompanyNumber
`func (o *Business) UnsetCompanyNumber()`

UnsetCompanyNumber ensures that no value is present for CompanyNumber, not even an explicit nil
### GetTaxIdentifier

`func (o *Business) GetTaxIdentifier() string`

GetTaxIdentifier returns the TaxIdentifier field if non-nil, zero value otherwise.

### GetTaxIdentifierOk

`func (o *Business) GetTaxIdentifierOk() (*string, bool)`

GetTaxIdentifierOk returns a tuple with the TaxIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdentifier

`func (o *Business) SetTaxIdentifier(v string)`

SetTaxIdentifier sets TaxIdentifier field to given value.

### HasTaxIdentifier

`func (o *Business) HasTaxIdentifier() bool`

HasTaxIdentifier returns a boolean if a field has been set.

### SetTaxIdentifierNil

`func (o *Business) SetTaxIdentifierNil(b bool)`

 SetTaxIdentifierNil sets the value for TaxIdentifier to be an explicit nil

### UnsetTaxIdentifier
`func (o *Business) UnsetTaxIdentifier()`

UnsetTaxIdentifier ensures that no value is present for TaxIdentifier, not even an explicit nil
### GetBillingEmail

`func (o *Business) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *Business) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *Business) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *Business) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### SetBillingEmailNil

`func (o *Business) SetBillingEmailNil(b bool)`

 SetBillingEmailNil sets the value for BillingEmail to be an explicit nil

### UnsetBillingEmail
`func (o *Business) UnsetBillingEmail()`

UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
### GetBillingPhone

`func (o *Business) GetBillingPhone() string`

GetBillingPhone returns the BillingPhone field if non-nil, zero value otherwise.

### GetBillingPhoneOk

`func (o *Business) GetBillingPhoneOk() (*string, bool)`

GetBillingPhoneOk returns a tuple with the BillingPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPhone

`func (o *Business) SetBillingPhone(v string)`

SetBillingPhone sets BillingPhone field to given value.

### HasBillingPhone

`func (o *Business) HasBillingPhone() bool`

HasBillingPhone returns a boolean if a field has been set.

### SetBillingPhoneNil

`func (o *Business) SetBillingPhoneNil(b bool)`

 SetBillingPhoneNil sets the value for BillingPhone to be an explicit nil

### UnsetBillingPhone
`func (o *Business) UnsetBillingPhone()`

UnsetBillingPhone ensures that no value is present for BillingPhone, not even an explicit nil
### GetBillingAddressLine1

`func (o *Business) GetBillingAddressLine1() string`

GetBillingAddressLine1 returns the BillingAddressLine1 field if non-nil, zero value otherwise.

### GetBillingAddressLine1Ok

`func (o *Business) GetBillingAddressLine1Ok() (*string, bool)`

GetBillingAddressLine1Ok returns a tuple with the BillingAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressLine1

`func (o *Business) SetBillingAddressLine1(v string)`

SetBillingAddressLine1 sets BillingAddressLine1 field to given value.

### HasBillingAddressLine1

`func (o *Business) HasBillingAddressLine1() bool`

HasBillingAddressLine1 returns a boolean if a field has been set.

### SetBillingAddressLine1Nil

`func (o *Business) SetBillingAddressLine1Nil(b bool)`

 SetBillingAddressLine1Nil sets the value for BillingAddressLine1 to be an explicit nil

### UnsetBillingAddressLine1
`func (o *Business) UnsetBillingAddressLine1()`

UnsetBillingAddressLine1 ensures that no value is present for BillingAddressLine1, not even an explicit nil
### GetBillingAddressLine2

`func (o *Business) GetBillingAddressLine2() string`

GetBillingAddressLine2 returns the BillingAddressLine2 field if non-nil, zero value otherwise.

### GetBillingAddressLine2Ok

`func (o *Business) GetBillingAddressLine2Ok() (*string, bool)`

GetBillingAddressLine2Ok returns a tuple with the BillingAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddressLine2

`func (o *Business) SetBillingAddressLine2(v string)`

SetBillingAddressLine2 sets BillingAddressLine2 field to given value.

### HasBillingAddressLine2

`func (o *Business) HasBillingAddressLine2() bool`

HasBillingAddressLine2 returns a boolean if a field has been set.

### SetBillingAddressLine2Nil

`func (o *Business) SetBillingAddressLine2Nil(b bool)`

 SetBillingAddressLine2Nil sets the value for BillingAddressLine2 to be an explicit nil

### UnsetBillingAddressLine2
`func (o *Business) UnsetBillingAddressLine2()`

UnsetBillingAddressLine2 ensures that no value is present for BillingAddressLine2, not even an explicit nil
### GetBillingCity

`func (o *Business) GetBillingCity() string`

GetBillingCity returns the BillingCity field if non-nil, zero value otherwise.

### GetBillingCityOk

`func (o *Business) GetBillingCityOk() (*string, bool)`

GetBillingCityOk returns a tuple with the BillingCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCity

`func (o *Business) SetBillingCity(v string)`

SetBillingCity sets BillingCity field to given value.

### HasBillingCity

`func (o *Business) HasBillingCity() bool`

HasBillingCity returns a boolean if a field has been set.

### SetBillingCityNil

`func (o *Business) SetBillingCityNil(b bool)`

 SetBillingCityNil sets the value for BillingCity to be an explicit nil

### UnsetBillingCity
`func (o *Business) UnsetBillingCity()`

UnsetBillingCity ensures that no value is present for BillingCity, not even an explicit nil
### GetBillingPostcode

`func (o *Business) GetBillingPostcode() string`

GetBillingPostcode returns the BillingPostcode field if non-nil, zero value otherwise.

### GetBillingPostcodeOk

`func (o *Business) GetBillingPostcodeOk() (*string, bool)`

GetBillingPostcodeOk returns a tuple with the BillingPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPostcode

`func (o *Business) SetBillingPostcode(v string)`

SetBillingPostcode sets BillingPostcode field to given value.

### HasBillingPostcode

`func (o *Business) HasBillingPostcode() bool`

HasBillingPostcode returns a boolean if a field has been set.

### SetBillingPostcodeNil

`func (o *Business) SetBillingPostcodeNil(b bool)`

 SetBillingPostcodeNil sets the value for BillingPostcode to be an explicit nil

### UnsetBillingPostcode
`func (o *Business) UnsetBillingPostcode()`

UnsetBillingPostcode ensures that no value is present for BillingPostcode, not even an explicit nil
### GetBillingCountry

`func (o *Business) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *Business) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *Business) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.

### HasBillingCountry

`func (o *Business) HasBillingCountry() bool`

HasBillingCountry returns a boolean if a field has been set.

### SetBillingCountryNil

`func (o *Business) SetBillingCountryNil(b bool)`

 SetBillingCountryNil sets the value for BillingCountry to be an explicit nil

### UnsetBillingCountry
`func (o *Business) UnsetBillingCountry()`

UnsetBillingCountry ensures that no value is present for BillingCountry, not even an explicit nil
### GetMetadata

`func (o *Business) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Business) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Business) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Business) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Business) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Business) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Business) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Business) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Business) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Business) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


