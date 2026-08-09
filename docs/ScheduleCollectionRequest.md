# ScheduleCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMinor** | **int32** | Collection amount in minor units (pence) | 
**Currency** | **string** | ISO 4217 currency. Bacs collections are GBP-only. | 
**SourceType** | **string** | What the collection settles. The settlement webhook resolves invoices by this reference. | 
**SourceId** | **string** | Invoice or payment-link ID the collection is for | 
**RequestedCollectionDate** | Pointer to **string** | Earliest acceptable collection date (ISO date). The backend resolves the Bacs working-day chain; must be on/after the bureau&#39;s earliest allowed date. | [optional] 

## Methods

### NewScheduleCollectionRequest

`func NewScheduleCollectionRequest(amountMinor int32, currency string, sourceType string, sourceId string, ) *ScheduleCollectionRequest`

NewScheduleCollectionRequest instantiates a new ScheduleCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCollectionRequestWithDefaults

`func NewScheduleCollectionRequestWithDefaults() *ScheduleCollectionRequest`

NewScheduleCollectionRequestWithDefaults instantiates a new ScheduleCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountMinor

`func (o *ScheduleCollectionRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *ScheduleCollectionRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *ScheduleCollectionRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *ScheduleCollectionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ScheduleCollectionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ScheduleCollectionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSourceType

`func (o *ScheduleCollectionRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ScheduleCollectionRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ScheduleCollectionRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *ScheduleCollectionRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ScheduleCollectionRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ScheduleCollectionRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetRequestedCollectionDate

`func (o *ScheduleCollectionRequest) GetRequestedCollectionDate() string`

GetRequestedCollectionDate returns the RequestedCollectionDate field if non-nil, zero value otherwise.

### GetRequestedCollectionDateOk

`func (o *ScheduleCollectionRequest) GetRequestedCollectionDateOk() (*string, bool)`

GetRequestedCollectionDateOk returns a tuple with the RequestedCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCollectionDate

`func (o *ScheduleCollectionRequest) SetRequestedCollectionDate(v string)`

SetRequestedCollectionDate sets RequestedCollectionDate field to given value.

### HasRequestedCollectionDate

`func (o *ScheduleCollectionRequest) HasRequestedCollectionDate() bool`

HasRequestedCollectionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


