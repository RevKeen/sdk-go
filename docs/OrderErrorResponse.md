# OrderErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**PriceErrorResponseError**](PriceErrorResponseError.md) |  | 

## Methods

### NewOrderErrorResponse

`func NewOrderErrorResponse(error_ PriceErrorResponseError, ) *OrderErrorResponse`

NewOrderErrorResponse instantiates a new OrderErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderErrorResponseWithDefaults

`func NewOrderErrorResponseWithDefaults() *OrderErrorResponse`

NewOrderErrorResponseWithDefaults instantiates a new OrderErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OrderErrorResponse) GetError() PriceErrorResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OrderErrorResponse) GetErrorOk() (*PriceErrorResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OrderErrorResponse) SetError(v PriceErrorResponseError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


