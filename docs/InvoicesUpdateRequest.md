# InvoicesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TotalMinor** | Pointer to **int32** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**SubscriptionTerms** | Pointer to [**InvoicesUpdateRequestSubscriptionTerms**](InvoicesUpdateRequestSubscriptionTerms.md) |  | [optional] 

## Methods

### NewInvoicesUpdateRequest

`func NewInvoicesUpdateRequest() *InvoicesUpdateRequest`

NewInvoicesUpdateRequest instantiates a new InvoicesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesUpdateRequestWithDefaults

`func NewInvoicesUpdateRequestWithDefaults() *InvoicesUpdateRequest`

NewInvoicesUpdateRequestWithDefaults instantiates a new InvoicesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InvoicesUpdateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoicesUpdateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoicesUpdateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoicesUpdateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalMinor

`func (o *InvoicesUpdateRequest) GetTotalMinor() int32`

GetTotalMinor returns the TotalMinor field if non-nil, zero value otherwise.

### GetTotalMinorOk

`func (o *InvoicesUpdateRequest) GetTotalMinorOk() (*int32, bool)`

GetTotalMinorOk returns a tuple with the TotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinor

`func (o *InvoicesUpdateRequest) SetTotalMinor(v int32)`

SetTotalMinor sets TotalMinor field to given value.

### HasTotalMinor

`func (o *InvoicesUpdateRequest) HasTotalMinor() bool`

HasTotalMinor returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoicesUpdateRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoicesUpdateRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoicesUpdateRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoicesUpdateRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSubscriptionTerms

`func (o *InvoicesUpdateRequest) GetSubscriptionTerms() InvoicesUpdateRequestSubscriptionTerms`

GetSubscriptionTerms returns the SubscriptionTerms field if non-nil, zero value otherwise.

### GetSubscriptionTermsOk

`func (o *InvoicesUpdateRequest) GetSubscriptionTermsOk() (*InvoicesUpdateRequestSubscriptionTerms, bool)`

GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTerms

`func (o *InvoicesUpdateRequest) SetSubscriptionTerms(v InvoicesUpdateRequestSubscriptionTerms)`

SetSubscriptionTerms sets SubscriptionTerms field to given value.

### HasSubscriptionTerms

`func (o *InvoicesUpdateRequest) HasSubscriptionTerms() bool`

HasSubscriptionTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


