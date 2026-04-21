# DryRunResultDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **float32** |  | 
**Status** | **string** |  | 
**ValidationDetails** | [**DryRunResultDataInnerValidationDetails**](DryRunResultDataInnerValidationDetails.md) |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewDryRunResultDataInner

`func NewDryRunResultDataInner(index float32, status string, validationDetails DryRunResultDataInnerValidationDetails, ) *DryRunResultDataInner`

NewDryRunResultDataInner instantiates a new DryRunResultDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDryRunResultDataInnerWithDefaults

`func NewDryRunResultDataInnerWithDefaults() *DryRunResultDataInner`

NewDryRunResultDataInnerWithDefaults instantiates a new DryRunResultDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *DryRunResultDataInner) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DryRunResultDataInner) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DryRunResultDataInner) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetStatus

`func (o *DryRunResultDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DryRunResultDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DryRunResultDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetValidationDetails

`func (o *DryRunResultDataInner) GetValidationDetails() DryRunResultDataInnerValidationDetails`

GetValidationDetails returns the ValidationDetails field if non-nil, zero value otherwise.

### GetValidationDetailsOk

`func (o *DryRunResultDataInner) GetValidationDetailsOk() (*DryRunResultDataInnerValidationDetails, bool)`

GetValidationDetailsOk returns a tuple with the ValidationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationDetails

`func (o *DryRunResultDataInner) SetValidationDetails(v DryRunResultDataInnerValidationDetails)`

SetValidationDetails sets ValidationDetails field to given value.


### GetReason

`func (o *DryRunResultDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DryRunResultDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DryRunResultDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DryRunResultDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


