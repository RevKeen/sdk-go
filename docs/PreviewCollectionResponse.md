# PreviewCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoticeDate** | **string** | Advance-notice date (YYYY-MM-DD) | 
**SubmissionDate** | **string** | Bacs submission date (YYYY-MM-DD) | 
**CollectionDate** | **string** | Collection date (YYYY-MM-DD) | 
**SettlementDate** | **string** | Settlement date (YYYY-MM-DD) | 
**Eligible** | **bool** | Whether the requested collection is eligible | 
**IneligibleReason** | **NullableString** | Reason the collection is ineligible, if applicable | 
**NoticeDays** | **int32** | Advance-notice days applied | 

## Methods

### NewPreviewCollectionResponse

`func NewPreviewCollectionResponse(noticeDate string, submissionDate string, collectionDate string, settlementDate string, eligible bool, ineligibleReason NullableString, noticeDays int32, ) *PreviewCollectionResponse`

NewPreviewCollectionResponse instantiates a new PreviewCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewCollectionResponseWithDefaults

`func NewPreviewCollectionResponseWithDefaults() *PreviewCollectionResponse`

NewPreviewCollectionResponseWithDefaults instantiates a new PreviewCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoticeDate

`func (o *PreviewCollectionResponse) GetNoticeDate() string`

GetNoticeDate returns the NoticeDate field if non-nil, zero value otherwise.

### GetNoticeDateOk

`func (o *PreviewCollectionResponse) GetNoticeDateOk() (*string, bool)`

GetNoticeDateOk returns a tuple with the NoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDate

`func (o *PreviewCollectionResponse) SetNoticeDate(v string)`

SetNoticeDate sets NoticeDate field to given value.


### GetSubmissionDate

`func (o *PreviewCollectionResponse) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *PreviewCollectionResponse) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *PreviewCollectionResponse) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.


### GetCollectionDate

`func (o *PreviewCollectionResponse) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *PreviewCollectionResponse) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *PreviewCollectionResponse) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.


### GetSettlementDate

`func (o *PreviewCollectionResponse) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *PreviewCollectionResponse) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *PreviewCollectionResponse) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### GetEligible

`func (o *PreviewCollectionResponse) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *PreviewCollectionResponse) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *PreviewCollectionResponse) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetIneligibleReason

`func (o *PreviewCollectionResponse) GetIneligibleReason() string`

GetIneligibleReason returns the IneligibleReason field if non-nil, zero value otherwise.

### GetIneligibleReasonOk

`func (o *PreviewCollectionResponse) GetIneligibleReasonOk() (*string, bool)`

GetIneligibleReasonOk returns a tuple with the IneligibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIneligibleReason

`func (o *PreviewCollectionResponse) SetIneligibleReason(v string)`

SetIneligibleReason sets IneligibleReason field to given value.


### SetIneligibleReasonNil

`func (o *PreviewCollectionResponse) SetIneligibleReasonNil(b bool)`

 SetIneligibleReasonNil sets the value for IneligibleReason to be an explicit nil

### UnsetIneligibleReason
`func (o *PreviewCollectionResponse) UnsetIneligibleReason()`

UnsetIneligibleReason ensures that no value is present for IneligibleReason, not even an explicit nil
### GetNoticeDays

`func (o *PreviewCollectionResponse) GetNoticeDays() int32`

GetNoticeDays returns the NoticeDays field if non-nil, zero value otherwise.

### GetNoticeDaysOk

`func (o *PreviewCollectionResponse) GetNoticeDaysOk() (*int32, bool)`

GetNoticeDaysOk returns a tuple with the NoticeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDays

`func (o *PreviewCollectionResponse) SetNoticeDays(v int32)`

SetNoticeDays sets NoticeDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


