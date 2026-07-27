# StorefrontStatusKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Publishable** | [**StorefrontKeyStatus**](StorefrontKeyStatus.md) |  | 
**Secret** | [**StorefrontKeyStatus**](StorefrontKeyStatus.md) |  | 

## Methods

### NewStorefrontStatusKeys

`func NewStorefrontStatusKeys(publishable StorefrontKeyStatus, secret StorefrontKeyStatus, ) *StorefrontStatusKeys`

NewStorefrontStatusKeys instantiates a new StorefrontStatusKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontStatusKeysWithDefaults

`func NewStorefrontStatusKeysWithDefaults() *StorefrontStatusKeys`

NewStorefrontStatusKeysWithDefaults instantiates a new StorefrontStatusKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishable

`func (o *StorefrontStatusKeys) GetPublishable() StorefrontKeyStatus`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *StorefrontStatusKeys) GetPublishableOk() (*StorefrontKeyStatus, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *StorefrontStatusKeys) SetPublishable(v StorefrontKeyStatus)`

SetPublishable sets Publishable field to given value.


### GetSecret

`func (o *StorefrontStatusKeys) GetSecret() StorefrontKeyStatus`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *StorefrontStatusKeys) GetSecretOk() (*StorefrontKeyStatus, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *StorefrontStatusKeys) SetSecret(v StorefrontKeyStatus)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


