# StorefrontKeyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Present** | **bool** |  | 
**Active** | **bool** |  | 
**CreatedAt** | **NullableString** |  | 
**LastUsedAt** | **NullableString** |  | 

## Methods

### NewStorefrontKeyStatus

`func NewStorefrontKeyStatus(present bool, active bool, createdAt NullableString, lastUsedAt NullableString, ) *StorefrontKeyStatus`

NewStorefrontKeyStatus instantiates a new StorefrontKeyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontKeyStatusWithDefaults

`func NewStorefrontKeyStatusWithDefaults() *StorefrontKeyStatus`

NewStorefrontKeyStatusWithDefaults instantiates a new StorefrontKeyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresent

`func (o *StorefrontKeyStatus) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *StorefrontKeyStatus) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *StorefrontKeyStatus) SetPresent(v bool)`

SetPresent sets Present field to given value.


### GetActive

`func (o *StorefrontKeyStatus) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StorefrontKeyStatus) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StorefrontKeyStatus) SetActive(v bool)`

SetActive sets Active field to given value.


### GetCreatedAt

`func (o *StorefrontKeyStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorefrontKeyStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorefrontKeyStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StorefrontKeyStatus) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StorefrontKeyStatus) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastUsedAt

`func (o *StorefrontKeyStatus) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *StorefrontKeyStatus) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *StorefrontKeyStatus) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *StorefrontKeyStatus) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *StorefrontKeyStatus) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


