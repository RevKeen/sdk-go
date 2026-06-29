# CustomerOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PublicId** | **string** |  | 
**CustomerId** | **string** |  | 
**Status** | **string** |  | 
**TotalMinor** | **int32** |  | 
**Currency** | **string** |  | 
**FulfilledAt** | **NullableString** |  | 
**CanceledAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewCustomerOrder

`func NewCustomerOrder(id string, publicId string, customerId string, status string, totalMinor int32, currency string, fulfilledAt NullableString, canceledAt NullableString, createdAt string, updatedAt string, ) *CustomerOrder`

NewCustomerOrder instantiates a new CustomerOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOrderWithDefaults

`func NewCustomerOrderWithDefaults() *CustomerOrder`

NewCustomerOrderWithDefaults instantiates a new CustomerOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerOrder) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *CustomerOrder) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *CustomerOrder) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *CustomerOrder) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetCustomerId

`func (o *CustomerOrder) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerOrder) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerOrder) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStatus

`func (o *CustomerOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerOrder) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalMinor

`func (o *CustomerOrder) GetTotalMinor() int32`

GetTotalMinor returns the TotalMinor field if non-nil, zero value otherwise.

### GetTotalMinorOk

`func (o *CustomerOrder) GetTotalMinorOk() (*int32, bool)`

GetTotalMinorOk returns a tuple with the TotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinor

`func (o *CustomerOrder) SetTotalMinor(v int32)`

SetTotalMinor sets TotalMinor field to given value.


### GetCurrency

`func (o *CustomerOrder) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomerOrder) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomerOrder) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFulfilledAt

`func (o *CustomerOrder) GetFulfilledAt() string`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *CustomerOrder) GetFulfilledAtOk() (*string, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *CustomerOrder) SetFulfilledAt(v string)`

SetFulfilledAt sets FulfilledAt field to given value.


### SetFulfilledAtNil

`func (o *CustomerOrder) SetFulfilledAtNil(b bool)`

 SetFulfilledAtNil sets the value for FulfilledAt to be an explicit nil

### UnsetFulfilledAt
`func (o *CustomerOrder) UnsetFulfilledAt()`

UnsetFulfilledAt ensures that no value is present for FulfilledAt, not even an explicit nil
### GetCanceledAt

`func (o *CustomerOrder) GetCanceledAt() string`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *CustomerOrder) GetCanceledAtOk() (*string, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *CustomerOrder) SetCanceledAt(v string)`

SetCanceledAt sets CanceledAt field to given value.


### SetCanceledAtNil

`func (o *CustomerOrder) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *CustomerOrder) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetCreatedAt

`func (o *CustomerOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CustomerOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


