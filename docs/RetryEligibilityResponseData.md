# RetryEligibilityResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldRetry** | **bool** | Whether the subscription should be retried | 
**Reason** | **string** | Explanation for the retry decision | 
**LastAttempt** | [**RetryEligibilityResponseDataLastAttempt**](RetryEligibilityResponseDataLastAttempt.md) |  | 

## Methods

### NewRetryEligibilityResponseData

`func NewRetryEligibilityResponseData(shouldRetry bool, reason string, lastAttempt RetryEligibilityResponseDataLastAttempt, ) *RetryEligibilityResponseData`

NewRetryEligibilityResponseData instantiates a new RetryEligibilityResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryEligibilityResponseDataWithDefaults

`func NewRetryEligibilityResponseDataWithDefaults() *RetryEligibilityResponseData`

NewRetryEligibilityResponseDataWithDefaults instantiates a new RetryEligibilityResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldRetry

`func (o *RetryEligibilityResponseData) GetShouldRetry() bool`

GetShouldRetry returns the ShouldRetry field if non-nil, zero value otherwise.

### GetShouldRetryOk

`func (o *RetryEligibilityResponseData) GetShouldRetryOk() (*bool, bool)`

GetShouldRetryOk returns a tuple with the ShouldRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRetry

`func (o *RetryEligibilityResponseData) SetShouldRetry(v bool)`

SetShouldRetry sets ShouldRetry field to given value.


### GetReason

`func (o *RetryEligibilityResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RetryEligibilityResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RetryEligibilityResponseData) SetReason(v string)`

SetReason sets Reason field to given value.


### GetLastAttempt

`func (o *RetryEligibilityResponseData) GetLastAttempt() RetryEligibilityResponseDataLastAttempt`

GetLastAttempt returns the LastAttempt field if non-nil, zero value otherwise.

### GetLastAttemptOk

`func (o *RetryEligibilityResponseData) GetLastAttemptOk() (*RetryEligibilityResponseDataLastAttempt, bool)`

GetLastAttemptOk returns a tuple with the LastAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttempt

`func (o *RetryEligibilityResponseData) SetLastAttempt(v RetryEligibilityResponseDataLastAttempt)`

SetLastAttempt sets LastAttempt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


