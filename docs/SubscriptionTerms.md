# SubscriptionTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionMethod** | **string** | How to collect payment for recurring invoices | 
**StartMode** | **string** | When the subscription billing cycle starts | 
**StartDate** | Pointer to **string** | Start date (ISO 8601) — required when start_mode is &#39;specific_date&#39; | [optional] 
**DurationType** | **string** | How the subscription duration is determined | 
**DurationCount** | Pointer to **int32** | Number of billing cycles — required when duration_type is &#39;fixed_cycles&#39; | [optional] 
**EndDate** | Pointer to **string** | End date (ISO 8601) — required when duration_type is &#39;end_date&#39; | [optional] 
**FirstPaymentBehavior** | **string** | Whether to charge the first cycle immediately or defer to the start date | 

## Methods

### NewSubscriptionTerms

`func NewSubscriptionTerms(collectionMethod string, startMode string, durationType string, firstPaymentBehavior string, ) *SubscriptionTerms`

NewSubscriptionTerms instantiates a new SubscriptionTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTermsWithDefaults

`func NewSubscriptionTermsWithDefaults() *SubscriptionTerms`

NewSubscriptionTermsWithDefaults instantiates a new SubscriptionTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionMethod

`func (o *SubscriptionTerms) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *SubscriptionTerms) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *SubscriptionTerms) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetStartMode

`func (o *SubscriptionTerms) GetStartMode() string`

GetStartMode returns the StartMode field if non-nil, zero value otherwise.

### GetStartModeOk

`func (o *SubscriptionTerms) GetStartModeOk() (*string, bool)`

GetStartModeOk returns a tuple with the StartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMode

`func (o *SubscriptionTerms) SetStartMode(v string)`

SetStartMode sets StartMode field to given value.


### GetStartDate

`func (o *SubscriptionTerms) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionTerms) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionTerms) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionTerms) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDurationType

`func (o *SubscriptionTerms) GetDurationType() string`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *SubscriptionTerms) GetDurationTypeOk() (*string, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *SubscriptionTerms) SetDurationType(v string)`

SetDurationType sets DurationType field to given value.


### GetDurationCount

`func (o *SubscriptionTerms) GetDurationCount() int32`

GetDurationCount returns the DurationCount field if non-nil, zero value otherwise.

### GetDurationCountOk

`func (o *SubscriptionTerms) GetDurationCountOk() (*int32, bool)`

GetDurationCountOk returns a tuple with the DurationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCount

`func (o *SubscriptionTerms) SetDurationCount(v int32)`

SetDurationCount sets DurationCount field to given value.

### HasDurationCount

`func (o *SubscriptionTerms) HasDurationCount() bool`

HasDurationCount returns a boolean if a field has been set.

### GetEndDate

`func (o *SubscriptionTerms) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionTerms) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionTerms) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SubscriptionTerms) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFirstPaymentBehavior

`func (o *SubscriptionTerms) GetFirstPaymentBehavior() string`

GetFirstPaymentBehavior returns the FirstPaymentBehavior field if non-nil, zero value otherwise.

### GetFirstPaymentBehaviorOk

`func (o *SubscriptionTerms) GetFirstPaymentBehaviorOk() (*string, bool)`

GetFirstPaymentBehaviorOk returns a tuple with the FirstPaymentBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaymentBehavior

`func (o *SubscriptionTerms) SetFirstPaymentBehavior(v string)`

SetFirstPaymentBehavior sets FirstPaymentBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


