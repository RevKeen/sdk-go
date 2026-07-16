# DdMandateRequestCancelResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**Status** | **string** |  | 
**CancelledAt** | Pointer to **NullableString** |  | [optional] 
**CancelledBy** | Pointer to **NullableString** |  | [optional] 
**ExpiredAt** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Idempotent** | Pointer to **bool** |  | [optional] 

## Methods

### NewDdMandateRequestCancelResponseData

`func NewDdMandateRequestCancelResponseData(id string, customerId string, invoiceId NullableString, status string, ) *DdMandateRequestCancelResponseData`

NewDdMandateRequestCancelResponseData instantiates a new DdMandateRequestCancelResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdMandateRequestCancelResponseDataWithDefaults

`func NewDdMandateRequestCancelResponseDataWithDefaults() *DdMandateRequestCancelResponseData`

NewDdMandateRequestCancelResponseDataWithDefaults instantiates a new DdMandateRequestCancelResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdMandateRequestCancelResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdMandateRequestCancelResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdMandateRequestCancelResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerId

`func (o *DdMandateRequestCancelResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdMandateRequestCancelResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdMandateRequestCancelResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *DdMandateRequestCancelResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DdMandateRequestCancelResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DdMandateRequestCancelResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *DdMandateRequestCancelResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *DdMandateRequestCancelResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetStatus

`func (o *DdMandateRequestCancelResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdMandateRequestCancelResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdMandateRequestCancelResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCancelledAt

`func (o *DdMandateRequestCancelResponseData) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DdMandateRequestCancelResponseData) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DdMandateRequestCancelResponseData) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *DdMandateRequestCancelResponseData) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *DdMandateRequestCancelResponseData) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *DdMandateRequestCancelResponseData) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetCancelledBy

`func (o *DdMandateRequestCancelResponseData) GetCancelledBy() string`

GetCancelledBy returns the CancelledBy field if non-nil, zero value otherwise.

### GetCancelledByOk

`func (o *DdMandateRequestCancelResponseData) GetCancelledByOk() (*string, bool)`

GetCancelledByOk returns a tuple with the CancelledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledBy

`func (o *DdMandateRequestCancelResponseData) SetCancelledBy(v string)`

SetCancelledBy sets CancelledBy field to given value.

### HasCancelledBy

`func (o *DdMandateRequestCancelResponseData) HasCancelledBy() bool`

HasCancelledBy returns a boolean if a field has been set.

### SetCancelledByNil

`func (o *DdMandateRequestCancelResponseData) SetCancelledByNil(b bool)`

 SetCancelledByNil sets the value for CancelledBy to be an explicit nil

### UnsetCancelledBy
`func (o *DdMandateRequestCancelResponseData) UnsetCancelledBy()`

UnsetCancelledBy ensures that no value is present for CancelledBy, not even an explicit nil
### GetExpiredAt

`func (o *DdMandateRequestCancelResponseData) GetExpiredAt() string`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *DdMandateRequestCancelResponseData) GetExpiredAtOk() (*string, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *DdMandateRequestCancelResponseData) SetExpiredAt(v string)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *DdMandateRequestCancelResponseData) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *DdMandateRequestCancelResponseData) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *DdMandateRequestCancelResponseData) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetReason

`func (o *DdMandateRequestCancelResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DdMandateRequestCancelResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DdMandateRequestCancelResponseData) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DdMandateRequestCancelResponseData) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *DdMandateRequestCancelResponseData) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *DdMandateRequestCancelResponseData) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetIdempotent

`func (o *DdMandateRequestCancelResponseData) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *DdMandateRequestCancelResponseData) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *DdMandateRequestCancelResponseData) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *DdMandateRequestCancelResponseData) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


