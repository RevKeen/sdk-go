# CustomersPaymentRailsGet200ResponseDataRailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rail** | **string** |  | 
**Available** | **bool** |  | 
**ReasonIfNot** | **NullableString** |  | 
**LastUsedAt** | **NullableTime** |  | 
**MandateId** | Pointer to **string** |  | [optional] 
**MandateRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCustomersPaymentRailsGet200ResponseDataRailsInner

`func NewCustomersPaymentRailsGet200ResponseDataRailsInner(rail string, available bool, reasonIfNot NullableString, lastUsedAt NullableTime, ) *CustomersPaymentRailsGet200ResponseDataRailsInner`

NewCustomersPaymentRailsGet200ResponseDataRailsInner instantiates a new CustomersPaymentRailsGet200ResponseDataRailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersPaymentRailsGet200ResponseDataRailsInnerWithDefaults

`func NewCustomersPaymentRailsGet200ResponseDataRailsInnerWithDefaults() *CustomersPaymentRailsGet200ResponseDataRailsInner`

NewCustomersPaymentRailsGet200ResponseDataRailsInnerWithDefaults instantiates a new CustomersPaymentRailsGet200ResponseDataRailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRail

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetRail(v string)`

SetRail sets Rail field to given value.


### GetAvailable

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetReasonIfNot

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetReasonIfNot() string`

GetReasonIfNot returns the ReasonIfNot field if non-nil, zero value otherwise.

### GetReasonIfNotOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetReasonIfNotOk() (*string, bool)`

GetReasonIfNotOk returns a tuple with the ReasonIfNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonIfNot

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetReasonIfNot(v string)`

SetReasonIfNot sets ReasonIfNot field to given value.


### SetReasonIfNotNil

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetReasonIfNotNil(b bool)`

 SetReasonIfNotNil sets the value for ReasonIfNot to be an explicit nil

### UnsetReasonIfNot
`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) UnsetReasonIfNot()`

UnsetReasonIfNot ensures that no value is present for ReasonIfNot, not even an explicit nil
### GetLastUsedAt

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetMandateId

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.

### HasMandateId

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) HasMandateId() bool`

HasMandateId returns a boolean if a field has been set.

### GetMandateRef

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetMandateRef() string`

GetMandateRef returns the MandateRef field if non-nil, zero value otherwise.

### GetMandateRefOk

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) GetMandateRefOk() (*string, bool)`

GetMandateRefOk returns a tuple with the MandateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRef

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetMandateRef(v string)`

SetMandateRef sets MandateRef field to given value.

### HasMandateRef

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) HasMandateRef() bool`

HasMandateRef returns a boolean if a field has been set.

### SetMandateRefNil

`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) SetMandateRefNil(b bool)`

 SetMandateRefNil sets the value for MandateRef to be an explicit nil

### UnsetMandateRef
`func (o *CustomersPaymentRailsGet200ResponseDataRailsInner) UnsetMandateRef()`

UnsetMandateRef ensures that no value is present for MandateRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


