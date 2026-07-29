# StorefrontStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Ready** | **bool** |  | 
**Checks** | [**[]StorefrontStatusCheck**](StorefrontStatusCheck.md) |  | 
**Cart** | [**StorefrontStatusCart**](StorefrontStatusCart.md) |  | 
**Keys** | [**StorefrontStatusKeys**](StorefrontStatusKeys.md) |  | 
**Origins** | [**StorefrontStatusOrigins**](StorefrontStatusOrigins.md) |  | 
**ProductRead** | [**StorefrontStatusProductRead**](StorefrontStatusProductRead.md) |  | 
**Webhooks** | [**StorefrontStatusWebhooks**](StorefrontStatusWebhooks.md) |  | 
**Availability** | [**StorefrontStatusAvailability**](StorefrontStatusAvailability.md) |  | 

## Methods

### NewStorefrontStatus

`func NewStorefrontStatus(object string, ready bool, checks []StorefrontStatusCheck, cart StorefrontStatusCart, keys StorefrontStatusKeys, origins StorefrontStatusOrigins, productRead StorefrontStatusProductRead, webhooks StorefrontStatusWebhooks, availability StorefrontStatusAvailability, ) *StorefrontStatus`

NewStorefrontStatus instantiates a new StorefrontStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontStatusWithDefaults

`func NewStorefrontStatusWithDefaults() *StorefrontStatus`

NewStorefrontStatusWithDefaults instantiates a new StorefrontStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *StorefrontStatus) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StorefrontStatus) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StorefrontStatus) SetObject(v string)`

SetObject sets Object field to given value.


### GetReady

`func (o *StorefrontStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *StorefrontStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *StorefrontStatus) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetChecks

`func (o *StorefrontStatus) GetChecks() []StorefrontStatusCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *StorefrontStatus) GetChecksOk() (*[]StorefrontStatusCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *StorefrontStatus) SetChecks(v []StorefrontStatusCheck)`

SetChecks sets Checks field to given value.


### GetCart

`func (o *StorefrontStatus) GetCart() StorefrontStatusCart`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *StorefrontStatus) GetCartOk() (*StorefrontStatusCart, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *StorefrontStatus) SetCart(v StorefrontStatusCart)`

SetCart sets Cart field to given value.


### GetKeys

`func (o *StorefrontStatus) GetKeys() StorefrontStatusKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *StorefrontStatus) GetKeysOk() (*StorefrontStatusKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *StorefrontStatus) SetKeys(v StorefrontStatusKeys)`

SetKeys sets Keys field to given value.


### GetOrigins

`func (o *StorefrontStatus) GetOrigins() StorefrontStatusOrigins`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *StorefrontStatus) GetOriginsOk() (*StorefrontStatusOrigins, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *StorefrontStatus) SetOrigins(v StorefrontStatusOrigins)`

SetOrigins sets Origins field to given value.


### GetProductRead

`func (o *StorefrontStatus) GetProductRead() StorefrontStatusProductRead`

GetProductRead returns the ProductRead field if non-nil, zero value otherwise.

### GetProductReadOk

`func (o *StorefrontStatus) GetProductReadOk() (*StorefrontStatusProductRead, bool)`

GetProductReadOk returns a tuple with the ProductRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRead

`func (o *StorefrontStatus) SetProductRead(v StorefrontStatusProductRead)`

SetProductRead sets ProductRead field to given value.


### GetWebhooks

`func (o *StorefrontStatus) GetWebhooks() StorefrontStatusWebhooks`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *StorefrontStatus) GetWebhooksOk() (*StorefrontStatusWebhooks, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *StorefrontStatus) SetWebhooks(v StorefrontStatusWebhooks)`

SetWebhooks sets Webhooks field to given value.


### GetAvailability

`func (o *StorefrontStatus) GetAvailability() StorefrontStatusAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *StorefrontStatus) GetAvailabilityOk() (*StorefrontStatusAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *StorefrontStatus) SetAvailability(v StorefrontStatusAvailability)`

SetAvailability sets Availability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


