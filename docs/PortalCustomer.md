# PortalCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Email** | **NullableString** |  | 
**Name** | **NullableString** |  | 
**Phone** | **NullableString** |  | 
**CompanyName** | **NullableString** |  | 
**DefaultCurrency** | **NullableString** |  | 
**Locale** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPortalCustomer

`func NewPortalCustomer(id string, object string, email NullableString, name NullableString, phone NullableString, companyName NullableString, defaultCurrency NullableString, locale NullableString, createdAt time.Time, ) *PortalCustomer`

NewPortalCustomer instantiates a new PortalCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCustomerWithDefaults

`func NewPortalCustomerWithDefaults() *PortalCustomer`

NewPortalCustomerWithDefaults instantiates a new PortalCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalCustomer) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PortalCustomer) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalCustomer) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalCustomer) SetObject(v string)`

SetObject sets Object field to given value.


### GetEmail

`func (o *PortalCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PortalCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PortalCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *PortalCustomer) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PortalCustomer) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *PortalCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalCustomer) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PortalCustomer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PortalCustomer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *PortalCustomer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PortalCustomer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PortalCustomer) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *PortalCustomer) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PortalCustomer) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetCompanyName

`func (o *PortalCustomer) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PortalCustomer) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PortalCustomer) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### SetCompanyNameNil

`func (o *PortalCustomer) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *PortalCustomer) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetDefaultCurrency

`func (o *PortalCustomer) GetDefaultCurrency() string`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *PortalCustomer) GetDefaultCurrencyOk() (*string, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *PortalCustomer) SetDefaultCurrency(v string)`

SetDefaultCurrency sets DefaultCurrency field to given value.


### SetDefaultCurrencyNil

`func (o *PortalCustomer) SetDefaultCurrencyNil(b bool)`

 SetDefaultCurrencyNil sets the value for DefaultCurrency to be an explicit nil

### UnsetDefaultCurrency
`func (o *PortalCustomer) UnsetDefaultCurrency()`

UnsetDefaultCurrency ensures that no value is present for DefaultCurrency, not even an explicit nil
### GetLocale

`func (o *PortalCustomer) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *PortalCustomer) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *PortalCustomer) SetLocale(v string)`

SetLocale sets Locale field to given value.


### SetLocaleNil

`func (o *PortalCustomer) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *PortalCustomer) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetCreatedAt

`func (o *PortalCustomer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalCustomer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalCustomer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


