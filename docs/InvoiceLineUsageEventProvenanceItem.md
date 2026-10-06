# InvoiceLineUsageEventProvenanceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | **string** |  |
**ExclusionReason** | Pointer to **NullableString** |  | [optional]
**Entry** | Pointer to [**NullableInvoiceLineUsageEventProvenanceItemEntry**](InvoiceLineUsageEventProvenanceItemEntry.md) |  | [optional]
**Match** | Pointer to [**NullableInvoiceLineUsageEventProvenanceItemMatch**](InvoiceLineUsageEventProvenanceItemMatch.md) |  | [optional]
**Event** | [**InvoiceLineUsageEventProvenanceItemEvent**](InvoiceLineUsageEventProvenanceItemEvent.md) |  |

## Methods

### NewInvoiceLineUsageEventProvenanceItem

`func NewInvoiceLineUsageEventProvenanceItem(outcome string, event InvoiceLineUsageEventProvenanceItemEvent, ) *InvoiceLineUsageEventProvenanceItem`

NewInvoiceLineUsageEventProvenanceItem instantiates a new InvoiceLineUsageEventProvenanceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageEventProvenanceItemWithDefaults

`func NewInvoiceLineUsageEventProvenanceItemWithDefaults() *InvoiceLineUsageEventProvenanceItem`

NewInvoiceLineUsageEventProvenanceItemWithDefaults instantiates a new InvoiceLineUsageEventProvenanceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *InvoiceLineUsageEventProvenanceItem) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *InvoiceLineUsageEventProvenanceItem) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *InvoiceLineUsageEventProvenanceItem) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetExclusionReason

`func (o *InvoiceLineUsageEventProvenanceItem) GetExclusionReason() string`

GetExclusionReason returns the ExclusionReason field if non-nil, zero value otherwise.

### GetExclusionReasonOk

`func (o *InvoiceLineUsageEventProvenanceItem) GetExclusionReasonOk() (*string, bool)`

GetExclusionReasonOk returns a tuple with the ExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionReason

`func (o *InvoiceLineUsageEventProvenanceItem) SetExclusionReason(v string)`

SetExclusionReason sets ExclusionReason field to given value.

### HasExclusionReason

`func (o *InvoiceLineUsageEventProvenanceItem) HasExclusionReason() bool`

HasExclusionReason returns a boolean if a field has been set.

### SetExclusionReasonNil

`func (o *InvoiceLineUsageEventProvenanceItem) SetExclusionReasonNil(b bool)`

 SetExclusionReasonNil sets the value for ExclusionReason to be an explicit nil

### UnsetExclusionReason
`func (o *InvoiceLineUsageEventProvenanceItem) UnsetExclusionReason()`

UnsetExclusionReason ensures that no value is present for ExclusionReason, not even an explicit nil
### GetEntry

`func (o *InvoiceLineUsageEventProvenanceItem) GetEntry() InvoiceLineUsageEventProvenanceItemEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *InvoiceLineUsageEventProvenanceItem) GetEntryOk() (*InvoiceLineUsageEventProvenanceItemEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *InvoiceLineUsageEventProvenanceItem) SetEntry(v InvoiceLineUsageEventProvenanceItemEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *InvoiceLineUsageEventProvenanceItem) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### SetEntryNil

`func (o *InvoiceLineUsageEventProvenanceItem) SetEntryNil(b bool)`

 SetEntryNil sets the value for Entry to be an explicit nil

### UnsetEntry
`func (o *InvoiceLineUsageEventProvenanceItem) UnsetEntry()`

UnsetEntry ensures that no value is present for Entry, not even an explicit nil
### GetMatch

`func (o *InvoiceLineUsageEventProvenanceItem) GetMatch() InvoiceLineUsageEventProvenanceItemMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *InvoiceLineUsageEventProvenanceItem) GetMatchOk() (*InvoiceLineUsageEventProvenanceItemMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *InvoiceLineUsageEventProvenanceItem) SetMatch(v InvoiceLineUsageEventProvenanceItemMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *InvoiceLineUsageEventProvenanceItem) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *InvoiceLineUsageEventProvenanceItem) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *InvoiceLineUsageEventProvenanceItem) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil
### GetEvent

`func (o *InvoiceLineUsageEventProvenanceItem) GetEvent() InvoiceLineUsageEventProvenanceItemEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *InvoiceLineUsageEventProvenanceItem) GetEventOk() (*InvoiceLineUsageEventProvenanceItemEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *InvoiceLineUsageEventProvenanceItem) SetEvent(v InvoiceLineUsageEventProvenanceItemEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
