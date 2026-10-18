# SubscriptionResumeConflictErrorErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MandateId** | Pointer to **string** | The Direct Debit mandate the subscription collects against. | [optional]
**MandateUsabilityCode** | Pointer to **string** | Why the mandate cannot collect, for example &#x60;local_cancelled&#x60;, &#x60;local_suspended&#x60;, &#x60;provider_status_pending&#x60; or &#x60;mandate_missing&#x60;. | [optional]

## Methods

### NewSubscriptionResumeConflictErrorErrorDetails

`func NewSubscriptionResumeConflictErrorErrorDetails() *SubscriptionResumeConflictErrorErrorDetails`

NewSubscriptionResumeConflictErrorErrorDetails instantiates a new SubscriptionResumeConflictErrorErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResumeConflictErrorErrorDetailsWithDefaults

`func NewSubscriptionResumeConflictErrorErrorDetailsWithDefaults() *SubscriptionResumeConflictErrorErrorDetails`

NewSubscriptionResumeConflictErrorErrorDetailsWithDefaults instantiates a new SubscriptionResumeConflictErrorErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMandateId

`func (o *SubscriptionResumeConflictErrorErrorDetails) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *SubscriptionResumeConflictErrorErrorDetails) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *SubscriptionResumeConflictErrorErrorDetails) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.

### HasMandateId

`func (o *SubscriptionResumeConflictErrorErrorDetails) HasMandateId() bool`

HasMandateId returns a boolean if a field has been set.

### GetMandateUsabilityCode

`func (o *SubscriptionResumeConflictErrorErrorDetails) GetMandateUsabilityCode() string`

GetMandateUsabilityCode returns the MandateUsabilityCode field if non-nil, zero value otherwise.

### GetMandateUsabilityCodeOk

`func (o *SubscriptionResumeConflictErrorErrorDetails) GetMandateUsabilityCodeOk() (*string, bool)`

GetMandateUsabilityCodeOk returns a tuple with the MandateUsabilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateUsabilityCode

`func (o *SubscriptionResumeConflictErrorErrorDetails) SetMandateUsabilityCode(v string)`

SetMandateUsabilityCode sets MandateUsabilityCode field to given value.

### HasMandateUsabilityCode

`func (o *SubscriptionResumeConflictErrorErrorDetails) HasMandateUsabilityCode() bool`

HasMandateUsabilityCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
