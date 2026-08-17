# CollectionSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Collection schedule ID | 
**MandateId** | **string** | Mandate the collection is scheduled against | 
**LzScheduleId** | **string** | Bureau schedule reference | 
**AmountMinor** | **int32** | Collection amount in minor units (pence) | 
**Currency** | **string** | ISO 4217 currency | 
**SourceType** | **string** | invoice | payment_link | 
**SourceId** | **string** | Invoice or payment-link ID the collection settles | 
**Status** | **string** | Schedule status | 
**CollectionDate** | **NullableString** | Collection date (YYYY-MM-DD) | 
**NoticeDate** | **NullableString** | Advance-notice date (YYYY-MM-DD) | 
**SubmissionDate** | **NullableString** | Bacs submission date (YYYY-MM-DD) | 
**AlreadyScheduled** | **bool** | True when an existing schedule for this source was returned instead of creating a duplicate | 

## Methods

### NewCollectionSchedule

`func NewCollectionSchedule(id NullableString, mandateId string, lzScheduleId string, amountMinor int32, currency string, sourceType string, sourceId string, status string, collectionDate NullableString, noticeDate NullableString, submissionDate NullableString, alreadyScheduled bool, ) *CollectionSchedule`

NewCollectionSchedule instantiates a new CollectionSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionScheduleWithDefaults

`func NewCollectionScheduleWithDefaults() *CollectionSchedule`

NewCollectionScheduleWithDefaults instantiates a new CollectionSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionSchedule) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CollectionSchedule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CollectionSchedule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMandateId

`func (o *CollectionSchedule) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *CollectionSchedule) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *CollectionSchedule) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetLzScheduleId

`func (o *CollectionSchedule) GetLzScheduleId() string`

GetLzScheduleId returns the LzScheduleId field if non-nil, zero value otherwise.

### GetLzScheduleIdOk

`func (o *CollectionSchedule) GetLzScheduleIdOk() (*string, bool)`

GetLzScheduleIdOk returns a tuple with the LzScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLzScheduleId

`func (o *CollectionSchedule) SetLzScheduleId(v string)`

SetLzScheduleId sets LzScheduleId field to given value.


### GetAmountMinor

`func (o *CollectionSchedule) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CollectionSchedule) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CollectionSchedule) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *CollectionSchedule) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CollectionSchedule) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CollectionSchedule) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSourceType

`func (o *CollectionSchedule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CollectionSchedule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CollectionSchedule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *CollectionSchedule) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CollectionSchedule) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CollectionSchedule) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetStatus

`func (o *CollectionSchedule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CollectionSchedule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CollectionSchedule) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCollectionDate

`func (o *CollectionSchedule) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *CollectionSchedule) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *CollectionSchedule) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.


### SetCollectionDateNil

`func (o *CollectionSchedule) SetCollectionDateNil(b bool)`

 SetCollectionDateNil sets the value for CollectionDate to be an explicit nil

### UnsetCollectionDate
`func (o *CollectionSchedule) UnsetCollectionDate()`

UnsetCollectionDate ensures that no value is present for CollectionDate, not even an explicit nil
### GetNoticeDate

`func (o *CollectionSchedule) GetNoticeDate() string`

GetNoticeDate returns the NoticeDate field if non-nil, zero value otherwise.

### GetNoticeDateOk

`func (o *CollectionSchedule) GetNoticeDateOk() (*string, bool)`

GetNoticeDateOk returns a tuple with the NoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDate

`func (o *CollectionSchedule) SetNoticeDate(v string)`

SetNoticeDate sets NoticeDate field to given value.


### SetNoticeDateNil

`func (o *CollectionSchedule) SetNoticeDateNil(b bool)`

 SetNoticeDateNil sets the value for NoticeDate to be an explicit nil

### UnsetNoticeDate
`func (o *CollectionSchedule) UnsetNoticeDate()`

UnsetNoticeDate ensures that no value is present for NoticeDate, not even an explicit nil
### GetSubmissionDate

`func (o *CollectionSchedule) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *CollectionSchedule) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *CollectionSchedule) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.


### SetSubmissionDateNil

`func (o *CollectionSchedule) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *CollectionSchedule) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil
### GetAlreadyScheduled

`func (o *CollectionSchedule) GetAlreadyScheduled() bool`

GetAlreadyScheduled returns the AlreadyScheduled field if non-nil, zero value otherwise.

### GetAlreadyScheduledOk

`func (o *CollectionSchedule) GetAlreadyScheduledOk() (*bool, bool)`

GetAlreadyScheduledOk returns a tuple with the AlreadyScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyScheduled

`func (o *CollectionSchedule) SetAlreadyScheduled(v bool)`

SetAlreadyScheduled sets AlreadyScheduled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


