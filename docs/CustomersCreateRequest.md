# CustomersCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Customer&#39;s email address (must be unique per merchant) | 
**Name** | Pointer to **string** | Customer&#39;s full name | [optional] 
**Phone** | Pointer to **string** | Customer&#39;s phone number | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata to attach to this customer | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Merchant-defined custom fields displayed as columns in the dashboard | [optional] 

## Methods

### NewCustomersCreateRequest

`func NewCustomersCreateRequest(email string, ) *CustomersCreateRequest`

NewCustomersCreateRequest instantiates a new CustomersCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersCreateRequestWithDefaults

`func NewCustomersCreateRequestWithDefaults() *CustomersCreateRequest`

NewCustomersCreateRequestWithDefaults instantiates a new CustomersCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CustomersCreateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomersCreateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomersCreateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *CustomersCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomersCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomersCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomersCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *CustomersCreateRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomersCreateRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomersCreateRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomersCreateRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomersCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomersCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomersCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomersCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCustomFields

`func (o *CustomersCreateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CustomersCreateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CustomersCreateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CustomersCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


