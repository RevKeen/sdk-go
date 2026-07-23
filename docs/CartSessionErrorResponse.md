# CartSessionErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewCartSessionErrorResponse

`func NewCartSessionErrorResponse(error_ string, ) *CartSessionErrorResponse`

NewCartSessionErrorResponse instantiates a new CartSessionErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartSessionErrorResponseWithDefaults

`func NewCartSessionErrorResponseWithDefaults() *CartSessionErrorResponse`

NewCartSessionErrorResponseWithDefaults instantiates a new CartSessionErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CartSessionErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CartSessionErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CartSessionErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *CartSessionErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CartSessionErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CartSessionErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CartSessionErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


