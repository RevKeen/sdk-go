# InvoiceCollectionAcceptedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Invoice**](Invoice.md) |  |
**Collection** | [**InvoiceCollection**](InvoiceCollection.md) |  |
**RequestId** | **string** |  |

## Methods

### NewInvoiceCollectionAcceptedResponse

`func NewInvoiceCollectionAcceptedResponse(data Invoice, collection InvoiceCollection, requestId string, ) *InvoiceCollectionAcceptedResponse`

NewInvoiceCollectionAcceptedResponse instantiates a new InvoiceCollectionAcceptedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCollectionAcceptedResponseWithDefaults

`func NewInvoiceCollectionAcceptedResponseWithDefaults() *InvoiceCollectionAcceptedResponse`

NewInvoiceCollectionAcceptedResponseWithDefaults instantiates a new InvoiceCollectionAcceptedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InvoiceCollectionAcceptedResponse) GetData() Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoiceCollectionAcceptedResponse) GetDataOk() (*Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoiceCollectionAcceptedResponse) SetData(v Invoice)`

SetData sets Data field to given value.


### GetCollection

`func (o *InvoiceCollectionAcceptedResponse) GetCollection() InvoiceCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *InvoiceCollectionAcceptedResponse) GetCollectionOk() (*InvoiceCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *InvoiceCollectionAcceptedResponse) SetCollection(v InvoiceCollection)`

SetCollection sets Collection field to given value.


### GetRequestId

`func (o *InvoiceCollectionAcceptedResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InvoiceCollectionAcceptedResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InvoiceCollectionAcceptedResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
