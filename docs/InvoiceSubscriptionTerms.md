# InvoiceSubscriptionTerms

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

### NewInvoiceSubscriptionTerms

`func NewInvoiceSubscriptionTerms(collectionMethod string, startMode string, durationType string, firstPaymentBehavior string, ) *InvoiceSubscriptionTerms`

NewInvoiceSubscriptionTerms instantiates a new InvoiceSubscriptionTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceSubscriptionTermsWithDefaults

`func NewInvoiceSubscriptionTermsWithDefaults() *InvoiceSubscriptionTerms`

NewInvoiceSubscriptionTermsWithDefaults instantiates a new InvoiceSubscriptionTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionMethod

`func (o *InvoiceSubscriptionTerms) GetCollectionMethod() string`

GetCollectionMethod returns the CollectionMethod field if non-nil, zero value otherwise.

### GetCollectionMethodOk

`func (o *InvoiceSubscriptionTerms) GetCollectionMethodOk() (*string, bool)`

GetCollectionMethodOk returns a tuple with the CollectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMethod

`func (o *InvoiceSubscriptionTerms) SetCollectionMethod(v string)`

SetCollectionMethod sets CollectionMethod field to given value.


### GetStartMode

`func (o *InvoiceSubscriptionTerms) GetStartMode() string`

GetStartMode returns the StartMode field if non-nil, zero value otherwise.

### GetStartModeOk

`func (o *InvoiceSubscriptionTerms) GetStartModeOk() (*string, bool)`

GetStartModeOk returns a tuple with the StartMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMode

`func (o *InvoiceSubscriptionTerms) SetStartMode(v string)`

SetStartMode sets StartMode field to given value.


### GetStartDate

`func (o *InvoiceSubscriptionTerms) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvoiceSubscriptionTerms) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvoiceSubscriptionTerms) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InvoiceSubscriptionTerms) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetDurationType

`func (o *InvoiceSubscriptionTerms) GetDurationType() string`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *InvoiceSubscriptionTerms) GetDurationTypeOk() (*string, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *InvoiceSubscriptionTerms) SetDurationType(v string)`

SetDurationType sets DurationType field to given value.


### GetDurationCount

`func (o *InvoiceSubscriptionTerms) GetDurationCount() int32`

GetDurationCount returns the DurationCount field if non-nil, zero value otherwise.

### GetDurationCountOk

`func (o *InvoiceSubscriptionTerms) GetDurationCountOk() (*int32, bool)`

GetDurationCountOk returns a tuple with the DurationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCount

`func (o *InvoiceSubscriptionTerms) SetDurationCount(v int32)`

SetDurationCount sets DurationCount field to given value.

### HasDurationCount

`func (o *InvoiceSubscriptionTerms) HasDurationCount() bool`

HasDurationCount returns a boolean if a field has been set.

### GetEndDate

`func (o *InvoiceSubscriptionTerms) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvoiceSubscriptionTerms) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvoiceSubscriptionTerms) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InvoiceSubscriptionTerms) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetFirstPaymentBehavior

`func (o *InvoiceSubscriptionTerms) GetFirstPaymentBehavior() string`

GetFirstPaymentBehavior returns the FirstPaymentBehavior field if non-nil, zero value otherwise.

### GetFirstPaymentBehaviorOk

`func (o *InvoiceSubscriptionTerms) GetFirstPaymentBehaviorOk() (*string, bool)`

GetFirstPaymentBehaviorOk returns a tuple with the FirstPaymentBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaymentBehavior

`func (o *InvoiceSubscriptionTerms) SetFirstPaymentBehavior(v string)`

SetFirstPaymentBehavior sets FirstPaymentBehavior field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


