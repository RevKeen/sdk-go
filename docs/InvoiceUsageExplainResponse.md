# InvoiceUsageExplainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [**InvoiceUsageExplainResponseInvoice**](InvoiceUsageExplainResponseInvoice.md) |  |
**Lines** | [**[]InvoiceUsageExplainResponseLinesInner**](InvoiceUsageExplainResponseLinesInner.md) |  |
**Summary** | **string** |  |

## Methods

### NewInvoiceUsageExplainResponse

`func NewInvoiceUsageExplainResponse(invoice InvoiceUsageExplainResponseInvoice, lines []InvoiceUsageExplainResponseLinesInner, summary string, ) *InvoiceUsageExplainResponse`

NewInvoiceUsageExplainResponse instantiates a new InvoiceUsageExplainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUsageExplainResponseWithDefaults

`func NewInvoiceUsageExplainResponseWithDefaults() *InvoiceUsageExplainResponse`

NewInvoiceUsageExplainResponseWithDefaults instantiates a new InvoiceUsageExplainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *InvoiceUsageExplainResponse) GetInvoice() InvoiceUsageExplainResponseInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceUsageExplainResponse) GetInvoiceOk() (*InvoiceUsageExplainResponseInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceUsageExplainResponse) SetInvoice(v InvoiceUsageExplainResponseInvoice)`

SetInvoice sets Invoice field to given value.


### GetLines

`func (o *InvoiceUsageExplainResponse) GetLines() []InvoiceUsageExplainResponseLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *InvoiceUsageExplainResponse) GetLinesOk() (*[]InvoiceUsageExplainResponseLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *InvoiceUsageExplainResponse) SetLines(v []InvoiceUsageExplainResponseLinesInner)`

SetLines sets Lines field to given value.


### GetSummary

`func (o *InvoiceUsageExplainResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *InvoiceUsageExplainResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *InvoiceUsageExplainResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
