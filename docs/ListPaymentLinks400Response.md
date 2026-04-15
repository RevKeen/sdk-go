# ListPaymentLinks400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Details** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewListPaymentLinks400Response

`func NewListPaymentLinks400Response(error_ string, ) *ListPaymentLinks400Response`

NewListPaymentLinks400Response instantiates a new ListPaymentLinks400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentLinks400ResponseWithDefaults

`func NewListPaymentLinks400ResponseWithDefaults() *ListPaymentLinks400Response`

NewListPaymentLinks400ResponseWithDefaults instantiates a new ListPaymentLinks400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ListPaymentLinks400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListPaymentLinks400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListPaymentLinks400Response) SetError(v string)`

SetError sets Error field to given value.


### GetDetails

`func (o *ListPaymentLinks400Response) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ListPaymentLinks400Response) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ListPaymentLinks400Response) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ListPaymentLinks400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ListPaymentLinks400Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ListPaymentLinks400Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


