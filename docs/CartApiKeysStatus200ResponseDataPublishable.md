# CartApiKeysStatus200ResponseDataPublishable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  |
**Present** | **bool** |  |
**Active** | **bool** |  |
**Id** | **NullableString** |  |
**Scopes** | **[]string** |  |
**CreatedAt** | **NullableTime** |  |
**LastUsedAt** | **NullableTime** |  |
**RevokedAt** | **NullableTime** |  |

## Methods

### NewCartApiKeysStatus200ResponseDataPublishable

`func NewCartApiKeysStatus200ResponseDataPublishable(kind string, present bool, active bool, id NullableString, scopes []string, createdAt NullableTime, lastUsedAt NullableTime, revokedAt NullableTime, ) *CartApiKeysStatus200ResponseDataPublishable`

NewCartApiKeysStatus200ResponseDataPublishable instantiates a new CartApiKeysStatus200ResponseDataPublishable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartApiKeysStatus200ResponseDataPublishableWithDefaults

`func NewCartApiKeysStatus200ResponseDataPublishableWithDefaults() *CartApiKeysStatus200ResponseDataPublishable`

NewCartApiKeysStatus200ResponseDataPublishableWithDefaults instantiates a new CartApiKeysStatus200ResponseDataPublishable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPresent

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetPresent(v bool)`

SetPresent sets Present field to given value.


### GetActive

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetActive(v bool)`

SetActive sets Active field to given value.


### GetId

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CartApiKeysStatus200ResponseDataPublishable) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetScopes

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetCreatedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CartApiKeysStatus200ResponseDataPublishable) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLastUsedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *CartApiKeysStatus200ResponseDataPublishable) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetRevokedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *CartApiKeysStatus200ResponseDataPublishable) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.


### SetRevokedAtNil

`func (o *CartApiKeysStatus200ResponseDataPublishable) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *CartApiKeysStatus200ResponseDataPublishable) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
