# ChargeErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeErrorResponse

`func NewChargeErrorResponse(error_ string, ) *ChargeErrorResponse`

NewChargeErrorResponse instantiates a new ChargeErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeErrorResponseWithDefaults

`func NewChargeErrorResponseWithDefaults() *ChargeErrorResponse`

NewChargeErrorResponseWithDefaults instantiates a new ChargeErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ChargeErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChargeErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChargeErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *ChargeErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ChargeErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ChargeErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ChargeErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


