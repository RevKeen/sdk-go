# InvoicesVoidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason for voiding the invoice | [optional] 

## Methods

### NewInvoicesVoidRequest

`func NewInvoicesVoidRequest() *InvoicesVoidRequest`

NewInvoicesVoidRequest instantiates a new InvoicesVoidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesVoidRequestWithDefaults

`func NewInvoicesVoidRequestWithDefaults() *InvoicesVoidRequest`

NewInvoicesVoidRequestWithDefaults instantiates a new InvoicesVoidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *InvoicesVoidRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoicesVoidRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoicesVoidRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoicesVoidRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


