# MerchantAuthorizationDeniedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  |
**Error** | **string** |  |
**Action** | **string** |  |
**ResourceType** | **string** |  |
**Reason** | Pointer to **string** |  | [optional]

## Methods

### NewMerchantAuthorizationDeniedError

`func NewMerchantAuthorizationDeniedError(success bool, error_ string, action string, resourceType string, ) *MerchantAuthorizationDeniedError`

NewMerchantAuthorizationDeniedError instantiates a new MerchantAuthorizationDeniedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAuthorizationDeniedErrorWithDefaults

`func NewMerchantAuthorizationDeniedErrorWithDefaults() *MerchantAuthorizationDeniedError`

NewMerchantAuthorizationDeniedErrorWithDefaults instantiates a new MerchantAuthorizationDeniedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MerchantAuthorizationDeniedError) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MerchantAuthorizationDeniedError) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MerchantAuthorizationDeniedError) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetError

`func (o *MerchantAuthorizationDeniedError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MerchantAuthorizationDeniedError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MerchantAuthorizationDeniedError) SetError(v string)`

SetError sets Error field to given value.


### GetAction

`func (o *MerchantAuthorizationDeniedError) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MerchantAuthorizationDeniedError) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MerchantAuthorizationDeniedError) SetAction(v string)`

SetAction sets Action field to given value.


### GetResourceType

`func (o *MerchantAuthorizationDeniedError) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MerchantAuthorizationDeniedError) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MerchantAuthorizationDeniedError) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetReason

`func (o *MerchantAuthorizationDeniedError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantAuthorizationDeniedError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantAuthorizationDeniedError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantAuthorizationDeniedError) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
