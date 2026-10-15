# InvoiceLineUsageEventsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional]
**Data** | [**[]InvoiceLineUsageEventProvenanceItem**](InvoiceLineUsageEventProvenanceItem.md) |  |
**HasMore** | **bool** |  |
**Url** | Pointer to **string** |  | [optional]
**Line** | [**InvoiceLineUsageEventsListLine**](InvoiceLineUsageEventsListLine.md) |  |

## Methods

### NewInvoiceLineUsageEventsList

`func NewInvoiceLineUsageEventsList(data []InvoiceLineUsageEventProvenanceItem, hasMore bool, line InvoiceLineUsageEventsListLine, ) *InvoiceLineUsageEventsList`

NewInvoiceLineUsageEventsList instantiates a new InvoiceLineUsageEventsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageEventsListWithDefaults

`func NewInvoiceLineUsageEventsListWithDefaults() *InvoiceLineUsageEventsList`

NewInvoiceLineUsageEventsListWithDefaults instantiates a new InvoiceLineUsageEventsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *InvoiceLineUsageEventsList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InvoiceLineUsageEventsList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InvoiceLineUsageEventsList) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *InvoiceLineUsageEventsList) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetData

`func (o *InvoiceLineUsageEventsList) GetData() []InvoiceLineUsageEventProvenanceItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoiceLineUsageEventsList) GetDataOk() (*[]InvoiceLineUsageEventProvenanceItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoiceLineUsageEventsList) SetData(v []InvoiceLineUsageEventProvenanceItem)`

SetData sets Data field to given value.


### GetHasMore

`func (o *InvoiceLineUsageEventsList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *InvoiceLineUsageEventsList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *InvoiceLineUsageEventsList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *InvoiceLineUsageEventsList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InvoiceLineUsageEventsList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InvoiceLineUsageEventsList) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InvoiceLineUsageEventsList) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLine

`func (o *InvoiceLineUsageEventsList) GetLine() InvoiceLineUsageEventsListLine`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *InvoiceLineUsageEventsList) GetLineOk() (*InvoiceLineUsageEventsListLine, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *InvoiceLineUsageEventsList) SetLine(v InvoiceLineUsageEventsListLine)`

SetLine sets Line field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
