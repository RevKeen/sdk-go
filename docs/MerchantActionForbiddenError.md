# MerchantActionForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  |
**Success** | **bool** |  |
**Action** | **string** |  |
**ResourceType** | **string** |  |
**Reason** | Pointer to **string** |  | [optional]

## Methods

### NewMerchantActionForbiddenError

`func NewMerchantActionForbiddenError(error_ string, success bool, action string, resourceType string, ) *MerchantActionForbiddenError`

NewMerchantActionForbiddenError instantiates a new MerchantActionForbiddenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantActionForbiddenErrorWithDefaults

`func NewMerchantActionForbiddenErrorWithDefaults() *MerchantActionForbiddenError`

NewMerchantActionForbiddenErrorWithDefaults instantiates a new MerchantActionForbiddenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MerchantActionForbiddenError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MerchantActionForbiddenError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MerchantActionForbiddenError) SetError(v string)`

SetError sets Error field to given value.


### GetSuccess

`func (o *MerchantActionForbiddenError) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MerchantActionForbiddenError) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MerchantActionForbiddenError) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetAction

`func (o *MerchantActionForbiddenError) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MerchantActionForbiddenError) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MerchantActionForbiddenError) SetAction(v string)`

SetAction sets Action field to given value.


### GetResourceType

`func (o *MerchantActionForbiddenError) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MerchantActionForbiddenError) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MerchantActionForbiddenError) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetReason

`func (o *MerchantActionForbiddenError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantActionForbiddenError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantActionForbiddenError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantActionForbiddenError) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
