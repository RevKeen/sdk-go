# CustomersUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Updated email address | [optional] 
**Name** | Pointer to **string** | Updated full name | [optional] 
**Phone** | Pointer to **string** | Updated phone number | [optional] 
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


