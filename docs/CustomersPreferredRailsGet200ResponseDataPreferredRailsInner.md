# CustomersPreferredRailsGet200ResponseDataPreferredRailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rail** | **string** |  | 
**Available** | **bool** |  | 
**ReasonIfNot** | **NullableString** |  | 
**LastUsedAt** | **NullableTime** |  | 
**MandateId** | Pointer to **string** |  | [optional] 
**Rank** | **int32** |  | 
**Score** | **float32** |  | 
**Reason** | **string** |  | 
**Reasons** | **[]string** |  | 
**EstimatedFeeMinor** | Pointer to **int32** |  | [optional] 
**NetAfterFeesMinor** | Pointer to **int32** |  | [optional] 
**Mandate** | Pointer to [**CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate**](CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate.md) |  | [optional] 

## Methods

### NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInner

`func NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInner(rail string, available bool, reasonIfNot NullableString, lastUsedAt NullableTime, rank int32, score float32, reason string, reasons []string, ) *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner`

NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInner instantiates a new CustomersPreferredRailsGet200ResponseDataPreferredRailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInnerWithDefaults

`func NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInnerWithDefaults() *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner`

NewCustomersPreferredRailsGet200ResponseDataPreferredRailsInnerWithDefaults instantiates a new CustomersPreferredRailsGet200ResponseDataPreferredRailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRail

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetRail(v string)`

SetRail sets Rail field to given value.


### GetAvailable

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetReasonIfNot

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonIfNot() string`

GetReasonIfNot returns the ReasonIfNot field if non-nil, zero value otherwise.

### GetReasonIfNotOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonIfNotOk() (*string, bool)`

GetReasonIfNotOk returns a tuple with the ReasonIfNot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonIfNot

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReasonIfNot(v string)`

SetReasonIfNot sets ReasonIfNot field to given value.


### SetReasonIfNotNil

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReasonIfNotNil(b bool)`

 SetReasonIfNotNil sets the value for ReasonIfNot to be an explicit nil

### UnsetReasonIfNot
`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) UnsetReasonIfNot()`

UnsetReasonIfNot ensures that no value is present for ReasonIfNot, not even an explicit nil
### GetLastUsedAt

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetMandateId

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.

### HasMandateId

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasMandateId() bool`

HasMandateId returns a boolean if a field has been set.

### GetRank

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetScore

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetScore(v float32)`

SetScore sets Score field to given value.


### GetReason

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReasons

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetReasons(v []string)`

SetReasons sets Reasons field to given value.


### GetEstimatedFeeMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetEstimatedFeeMinor() int32`

GetEstimatedFeeMinor returns the EstimatedFeeMinor field if non-nil, zero value otherwise.

### GetEstimatedFeeMinorOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetEstimatedFeeMinorOk() (*int32, bool)`

GetEstimatedFeeMinorOk returns a tuple with the EstimatedFeeMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetEstimatedFeeMinor(v int32)`

SetEstimatedFeeMinor sets EstimatedFeeMinor field to given value.

### HasEstimatedFeeMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasEstimatedFeeMinor() bool`

HasEstimatedFeeMinor returns a boolean if a field has been set.

### GetNetAfterFeesMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetNetAfterFeesMinor() int32`

GetNetAfterFeesMinor returns the NetAfterFeesMinor field if non-nil, zero value otherwise.

### GetNetAfterFeesMinorOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetNetAfterFeesMinorOk() (*int32, bool)`

GetNetAfterFeesMinorOk returns a tuple with the NetAfterFeesMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAfterFeesMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetNetAfterFeesMinor(v int32)`

SetNetAfterFeesMinor sets NetAfterFeesMinor field to given value.

### HasNetAfterFeesMinor

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasNetAfterFeesMinor() bool`

HasNetAfterFeesMinor returns a boolean if a field has been set.

### GetMandate

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandate() CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate`

GetMandate returns the Mandate field if non-nil, zero value otherwise.

### GetMandateOk

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) GetMandateOk() (*CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate, bool)`

GetMandateOk returns a tuple with the Mandate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandate

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) SetMandate(v CustomersPreferredRailsGet200ResponseDataPreferredRailsInnerMandate)`

SetMandate sets Mandate field to given value.

### HasMandate

`func (o *CustomersPreferredRailsGet200ResponseDataPreferredRailsInner) HasMandate() bool`

HasMandate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


