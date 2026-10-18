# InvoicePdfResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**InvoicePdfResponseData**](InvoicePdfResponseData.md) |  |
**RequestId** | **string** |  |

## Methods

### NewInvoicePdfResponse

`func NewInvoicePdfResponse(data InvoicePdfResponseData, requestId string, ) *InvoicePdfResponse`

NewInvoicePdfResponse instantiates a new InvoicePdfResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePdfResponseWithDefaults

`func NewInvoicePdfResponseWithDefaults() *InvoicePdfResponse`

NewInvoicePdfResponseWithDefaults instantiates a new InvoicePdfResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InvoicePdfResponse) GetData() InvoicePdfResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoicePdfResponse) GetDataOk() (*InvoicePdfResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoicePdfResponse) SetData(v InvoicePdfResponseData)`

SetData sets Data field to given value.


### GetRequestId

`func (o *InvoicePdfResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InvoicePdfResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InvoicePdfResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
