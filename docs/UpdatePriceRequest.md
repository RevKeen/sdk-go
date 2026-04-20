# UpdatePriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the price is active | [optional] 
**Nickname** | Pointer to **NullableString** | Display name | [optional] 
**LookupKey** | Pointer to **NullableString** | Stable key for API lookups | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewUpdatePriceRequest

`func NewUpdatePriceRequest() *UpdatePriceRequest`

NewUpdatePriceRequest instantiates a new UpdatePriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceRequestWithDefaults

`func NewUpdatePriceRequestWithDefaults() *UpdatePriceRequest`

NewUpdatePriceRequestWithDefaults instantiates a new UpdatePriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdatePriceRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdatePriceRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdatePriceRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdatePriceRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetNickname

`func (o *UpdatePriceRequest) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UpdatePriceRequest) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UpdatePriceRequest) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UpdatePriceRequest) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### SetNicknameNil

`func (o *UpdatePriceRequest) SetNicknameNil(b bool)`

 SetNicknameNil sets the value for Nickname to be an explicit nil

### UnsetNickname
`func (o *UpdatePriceRequest) UnsetNickname()`

UnsetNickname ensures that no value is present for Nickname, not even an explicit nil
### GetLookupKey

`func (o *UpdatePriceRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *UpdatePriceRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *UpdatePriceRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *UpdatePriceRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### SetLookupKeyNil

`func (o *UpdatePriceRequest) SetLookupKeyNil(b bool)`

 SetLookupKeyNil sets the value for LookupKey to be an explicit nil

### UnsetLookupKey
`func (o *UpdatePriceRequest) UnsetLookupKey()`

UnsetLookupKey ensures that no value is present for LookupKey, not even an explicit nil
### GetMetadata

`func (o *UpdatePriceRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePriceRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePriceRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePriceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


