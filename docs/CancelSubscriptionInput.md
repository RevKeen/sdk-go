# CancelSubscriptionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Cancel mode: \&quot;immediately\&quot; revokes access now, \&quot;period_end\&quot; schedules cancellation at billing period end | 
**Reason** | Pointer to **string** | Optional reason for cancellation | [optional] 

## Methods

### NewCancelSubscriptionInput

`func NewCancelSubscriptionInput(mode string, ) *CancelSubscriptionInput`

NewCancelSubscriptionInput instantiates a new CancelSubscriptionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSubscriptionInputWithDefaults

`func NewCancelSubscriptionInputWithDefaults() *CancelSubscriptionInput`

NewCancelSubscriptionInputWithDefaults instantiates a new CancelSubscriptionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *CancelSubscriptionInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CancelSubscriptionInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CancelSubscriptionInput) SetMode(v string)`

SetMode sets Mode field to given value.


### GetReason

`func (o *CancelSubscriptionInput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelSubscriptionInput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelSubscriptionInput) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelSubscriptionInput) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


