# BillingPreviewErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Field** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingPreviewErrorResponse

`func NewBillingPreviewErrorResponse(error_ string, ) *BillingPreviewErrorResponse`

NewBillingPreviewErrorResponse instantiates a new BillingPreviewErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPreviewErrorResponseWithDefaults

`func NewBillingPreviewErrorResponseWithDefaults() *BillingPreviewErrorResponse`

NewBillingPreviewErrorResponseWithDefaults instantiates a new BillingPreviewErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BillingPreviewErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BillingPreviewErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BillingPreviewErrorResponse) SetError(v string)`

SetError sets Error field to given value.


### GetField

`func (o *BillingPreviewErrorResponse) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BillingPreviewErrorResponse) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BillingPreviewErrorResponse) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *BillingPreviewErrorResponse) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


