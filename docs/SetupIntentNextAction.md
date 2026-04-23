# SetupIntentNextAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of action to take | 
**RedirectToUrl** | Pointer to [**SetupIntentNextActionRedirectToUrl**](SetupIntentNextActionRedirectToUrl.md) |  | [optional] 
**UseStripeSdk** | Pointer to **map[string]interface{}** | SDK-specific data for client-side handling | [optional] 

## Methods

### NewSetupIntentNextAction

`func NewSetupIntentNextAction(type_ string, ) *SetupIntentNextAction`

NewSetupIntentNextAction instantiates a new SetupIntentNextAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetupIntentNextActionWithDefaults

`func NewSetupIntentNextActionWithDefaults() *SetupIntentNextAction`

NewSetupIntentNextActionWithDefaults instantiates a new SetupIntentNextAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetupIntentNextAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetupIntentNextAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetupIntentNextAction) SetType(v string)`

SetType sets Type field to given value.


### GetRedirectToUrl

`func (o *SetupIntentNextAction) GetRedirectToUrl() SetupIntentNextActionRedirectToUrl`

GetRedirectToUrl returns the RedirectToUrl field if non-nil, zero value otherwise.

### GetRedirectToUrlOk

`func (o *SetupIntentNextAction) GetRedirectToUrlOk() (*SetupIntentNextActionRedirectToUrl, bool)`

GetRedirectToUrlOk returns a tuple with the RedirectToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToUrl

`func (o *SetupIntentNextAction) SetRedirectToUrl(v SetupIntentNextActionRedirectToUrl)`

SetRedirectToUrl sets RedirectToUrl field to given value.

### HasRedirectToUrl

`func (o *SetupIntentNextAction) HasRedirectToUrl() bool`

HasRedirectToUrl returns a boolean if a field has been set.

### GetUseStripeSdk

`func (o *SetupIntentNextAction) GetUseStripeSdk() map[string]interface{}`

GetUseStripeSdk returns the UseStripeSdk field if non-nil, zero value otherwise.

### GetUseStripeSdkOk

`func (o *SetupIntentNextAction) GetUseStripeSdkOk() (*map[string]interface{}, bool)`

GetUseStripeSdkOk returns a tuple with the UseStripeSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStripeSdk

`func (o *SetupIntentNextAction) SetUseStripeSdk(v map[string]interface{})`

SetUseStripeSdk sets UseStripeSdk field to given value.

### HasUseStripeSdk

`func (o *SetupIntentNextAction) HasUseStripeSdk() bool`

HasUseStripeSdk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


