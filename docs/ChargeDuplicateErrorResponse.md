# ChargeDuplicateErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**ExistingChargeId** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeDuplicateErrorResponse

`func NewChargeDuplicateErrorResponse(error_ string, ) *ChargeDuplicateErrorResponse`

NewChargeDuplicateErrorResponse instantiates a new ChargeDuplicateErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeDuplicateErrorResponseWithDefaults

`func NewChargeDuplicateErrorResponseWithDefaults() *ChargeDuplicateErrorResponse`

NewChargeDuplicateErrorResponseWithDefaults instantiates a new ChargeDuplicateErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ChargeDuplicateErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChargeDuplicateErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChargeDuplicateErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetExistingChargeId

`func (o *ChargeDuplicateErrorResponse) GetExistingChargeId() string`

GetExistingChargeId returns the ExistingChargeId field if non-nil, zero value otherwise.

### GetExistingChargeIdOk

`func (o *ChargeDuplicateErrorResponse) GetExistingChargeIdOk() (*string, bool)`

GetExistingChargeIdOk returns a tuple with the ExistingChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingChargeId

`func (o *ChargeDuplicateErrorResponse) SetExistingChargeId(v string)`

SetExistingChargeId sets ExistingChargeId field to given value.

### HasExistingChargeId

`func (o *ChargeDuplicateErrorResponse) HasExistingChargeId() bool`

HasExistingChargeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


