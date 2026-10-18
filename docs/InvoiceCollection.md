# InvoiceCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | &#x60;scheduled&#x60;: the collection is with the Direct Debit bureau for &#x60;collection_date&#x60;. &#x60;awaiting_advance_notice&#x60;: the customer&#39;s advance notice has been sent and the collection is submitted once its delivery is confirmed; the dates are not yet fixed. |
**PaymentMethod** | **string** |  |
**Id** | **NullableString** | The collection schedule id, once scheduled. |
**MandateId** | **NullableString** |  |
**AmountMinor** | **NullableInt64** | The amount being collected, in minor units, once scheduled. |
**Currency** | **NullableString** |  |
**CollectionDate** | **NullableString** | The date the money is collected (YYYY-MM-DD), once scheduled. |
**NoticeDate** | **NullableString** |  |
**SubmissionDate** | **NullableString** |  |

## Methods

### NewInvoiceCollection

`func NewInvoiceCollection(status string, paymentMethod string, id NullableString, mandateId NullableString, amountMinor NullableInt64, currency NullableString, collectionDate NullableString, noticeDate NullableString, submissionDate NullableString, ) *InvoiceCollection`

NewInvoiceCollection instantiates a new InvoiceCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCollectionWithDefaults

`func NewInvoiceCollectionWithDefaults() *InvoiceCollection`

NewInvoiceCollectionWithDefaults instantiates a new InvoiceCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InvoiceCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceCollection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentMethod

`func (o *InvoiceCollection) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *InvoiceCollection) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *InvoiceCollection) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetId

`func (o *InvoiceCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceCollection) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *InvoiceCollection) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceCollection) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMandateId

`func (o *InvoiceCollection) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *InvoiceCollection) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *InvoiceCollection) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### SetMandateIdNil

`func (o *InvoiceCollection) SetMandateIdNil(b bool)`

 SetMandateIdNil sets the value for MandateId to be an explicit nil

### UnsetMandateId
`func (o *InvoiceCollection) UnsetMandateId()`

UnsetMandateId ensures that no value is present for MandateId, not even an explicit nil
### GetAmountMinor

`func (o *InvoiceCollection) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoiceCollection) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoiceCollection) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *InvoiceCollection) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *InvoiceCollection) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *InvoiceCollection) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceCollection) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceCollection) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *InvoiceCollection) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InvoiceCollection) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCollectionDate

`func (o *InvoiceCollection) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *InvoiceCollection) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *InvoiceCollection) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.


### SetCollectionDateNil

`func (o *InvoiceCollection) SetCollectionDateNil(b bool)`

 SetCollectionDateNil sets the value for CollectionDate to be an explicit nil

### UnsetCollectionDate
`func (o *InvoiceCollection) UnsetCollectionDate()`

UnsetCollectionDate ensures that no value is present for CollectionDate, not even an explicit nil
### GetNoticeDate

`func (o *InvoiceCollection) GetNoticeDate() string`

GetNoticeDate returns the NoticeDate field if non-nil, zero value otherwise.

### GetNoticeDateOk

`func (o *InvoiceCollection) GetNoticeDateOk() (*string, bool)`

GetNoticeDateOk returns a tuple with the NoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDate

`func (o *InvoiceCollection) SetNoticeDate(v string)`

SetNoticeDate sets NoticeDate field to given value.


### SetNoticeDateNil

`func (o *InvoiceCollection) SetNoticeDateNil(b bool)`

 SetNoticeDateNil sets the value for NoticeDate to be an explicit nil

### UnsetNoticeDate
`func (o *InvoiceCollection) UnsetNoticeDate()`

UnsetNoticeDate ensures that no value is present for NoticeDate, not even an explicit nil
### GetSubmissionDate

`func (o *InvoiceCollection) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *InvoiceCollection) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *InvoiceCollection) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.


### SetSubmissionDateNil

`func (o *InvoiceCollection) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *InvoiceCollection) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
