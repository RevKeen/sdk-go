# CartApiKeysStatus200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publishable** | [**CartApiKeysStatus200ResponseDataPublishable**](CartApiKeysStatus200ResponseDataPublishable.md) |  | 
**Secret** | [**CartApiKeysStatus200ResponseDataPublishable**](CartApiKeysStatus200ResponseDataPublishable.md) |  | 
**Ready** | **bool** |  | 
**Created** | Pointer to **[]string** |  | [optional] 
**PublishableApiKey** | Pointer to **string** |  | [optional] 
**SecretApiKey** | Pointer to **string** |  | [optional] 

## Methods

### NewCartApiKeysStatus200ResponseData

`func NewCartApiKeysStatus200ResponseData(publishable CartApiKeysStatus200ResponseDataPublishable, secret CartApiKeysStatus200ResponseDataPublishable, ready bool, ) *CartApiKeysStatus200ResponseData`

NewCartApiKeysStatus200ResponseData instantiates a new CartApiKeysStatus200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartApiKeysStatus200ResponseDataWithDefaults

`func NewCartApiKeysStatus200ResponseDataWithDefaults() *CartApiKeysStatus200ResponseData`

NewCartApiKeysStatus200ResponseDataWithDefaults instantiates a new CartApiKeysStatus200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishable

`func (o *CartApiKeysStatus200ResponseData) GetPublishable() CartApiKeysStatus200ResponseDataPublishable`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *CartApiKeysStatus200ResponseData) GetPublishableOk() (*CartApiKeysStatus200ResponseDataPublishable, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *CartApiKeysStatus200ResponseData) SetPublishable(v CartApiKeysStatus200ResponseDataPublishable)`

SetPublishable sets Publishable field to given value.


### GetSecret

`func (o *CartApiKeysStatus200ResponseData) GetSecret() CartApiKeysStatus200ResponseDataPublishable`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CartApiKeysStatus200ResponseData) GetSecretOk() (*CartApiKeysStatus200ResponseDataPublishable, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CartApiKeysStatus200ResponseData) SetSecret(v CartApiKeysStatus200ResponseDataPublishable)`

SetSecret sets Secret field to given value.


### GetReady

`func (o *CartApiKeysStatus200ResponseData) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *CartApiKeysStatus200ResponseData) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *CartApiKeysStatus200ResponseData) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetCreated

`func (o *CartApiKeysStatus200ResponseData) GetCreated() []string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CartApiKeysStatus200ResponseData) GetCreatedOk() (*[]string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CartApiKeysStatus200ResponseData) SetCreated(v []string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CartApiKeysStatus200ResponseData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPublishableApiKey

`func (o *CartApiKeysStatus200ResponseData) GetPublishableApiKey() string`

GetPublishableApiKey returns the PublishableApiKey field if non-nil, zero value otherwise.

### GetPublishableApiKeyOk

`func (o *CartApiKeysStatus200ResponseData) GetPublishableApiKeyOk() (*string, bool)`

GetPublishableApiKeyOk returns a tuple with the PublishableApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableApiKey

`func (o *CartApiKeysStatus200ResponseData) SetPublishableApiKey(v string)`

SetPublishableApiKey sets PublishableApiKey field to given value.

### HasPublishableApiKey

`func (o *CartApiKeysStatus200ResponseData) HasPublishableApiKey() bool`

HasPublishableApiKey returns a boolean if a field has been set.

### GetSecretApiKey

`func (o *CartApiKeysStatus200ResponseData) GetSecretApiKey() string`

GetSecretApiKey returns the SecretApiKey field if non-nil, zero value otherwise.

### GetSecretApiKeyOk

`func (o *CartApiKeysStatus200ResponseData) GetSecretApiKeyOk() (*string, bool)`

GetSecretApiKeyOk returns a tuple with the SecretApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretApiKey

`func (o *CartApiKeysStatus200ResponseData) SetSecretApiKey(v string)`

SetSecretApiKey sets SecretApiKey field to given value.

### HasSecretApiKey

`func (o *CartApiKeysStatus200ResponseData) HasSecretApiKey() bool`

HasSecretApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


