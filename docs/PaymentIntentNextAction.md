# PaymentIntentNextAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of action required | 
**RedirectToUrl** | Pointer to [**PaymentIntentNextActionRedirectToUrl**](PaymentIntentNextActionRedirectToUrl.md) |  | [optional] 

## Methods

### NewPaymentIntentNextAction

`func NewPaymentIntentNextAction(type_ string, ) *PaymentIntentNextAction`

NewPaymentIntentNextAction instantiates a new PaymentIntentNextAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentNextActionWithDefaults

`func NewPaymentIntentNextActionWithDefaults() *PaymentIntentNextAction`

NewPaymentIntentNextActionWithDefaults instantiates a new PaymentIntentNextAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentIntentNextAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentIntentNextAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentIntentNextAction) SetType(v string)`

SetType sets Type field to given value.


### GetRedirectToUrl

`func (o *PaymentIntentNextAction) GetRedirectToUrl() PaymentIntentNextActionRedirectToUrl`

GetRedirectToUrl returns the RedirectToUrl field if non-nil, zero value otherwise.

### GetRedirectToUrlOk

`func (o *PaymentIntentNextAction) GetRedirectToUrlOk() (*PaymentIntentNextActionRedirectToUrl, bool)`

GetRedirectToUrlOk returns a tuple with the RedirectToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectToUrl

`func (o *PaymentIntentNextAction) SetRedirectToUrl(v PaymentIntentNextActionRedirectToUrl)`

SetRedirectToUrl sets RedirectToUrl field to given value.

### HasRedirectToUrl

`func (o *PaymentIntentNextAction) HasRedirectToUrl() bool`

HasRedirectToUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


