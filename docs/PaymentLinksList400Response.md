# PaymentLinksList400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  | 
**Details** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPaymentLinksList400Response

`func NewPaymentLinksList400Response(error_ string, ) *PaymentLinksList400Response`

NewPaymentLinksList400Response instantiates a new PaymentLinksList400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinksList400ResponseWithDefaults

`func NewPaymentLinksList400ResponseWithDefaults() *PaymentLinksList400Response`

NewPaymentLinksList400ResponseWithDefaults instantiates a new PaymentLinksList400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PaymentLinksList400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PaymentLinksList400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PaymentLinksList400Response) SetError(v string)`

SetError sets Error field to given value.


### GetDetails

`func (o *PaymentLinksList400Response) GetDetails() interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PaymentLinksList400Response) GetDetailsOk() (*interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PaymentLinksList400Response) SetDetails(v interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PaymentLinksList400Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *PaymentLinksList400Response) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *PaymentLinksList400Response) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


