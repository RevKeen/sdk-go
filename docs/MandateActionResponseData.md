# MandateActionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mandate ID | 
**Status** | **string** | Resulting mandate status | 
**MandateReference** | Pointer to **NullableString** | Mandate reference (reinstate) | [optional] 

## Methods

### NewMandateActionResponseData

`func NewMandateActionResponseData(id string, status string, ) *MandateActionResponseData`

NewMandateActionResponseData instantiates a new MandateActionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateActionResponseDataWithDefaults

`func NewMandateActionResponseDataWithDefaults() *MandateActionResponseData`

NewMandateActionResponseDataWithDefaults instantiates a new MandateActionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MandateActionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MandateActionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MandateActionResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *MandateActionResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MandateActionResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MandateActionResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMandateReference

`func (o *MandateActionResponseData) GetMandateReference() string`

GetMandateReference returns the MandateReference field if non-nil, zero value otherwise.

### GetMandateReferenceOk

`func (o *MandateActionResponseData) GetMandateReferenceOk() (*string, bool)`

GetMandateReferenceOk returns a tuple with the MandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateReference

`func (o *MandateActionResponseData) SetMandateReference(v string)`

SetMandateReference sets MandateReference field to given value.

### HasMandateReference

`func (o *MandateActionResponseData) HasMandateReference() bool`

HasMandateReference returns a boolean if a field has been set.

### SetMandateReferenceNil

`func (o *MandateActionResponseData) SetMandateReferenceNil(b bool)`

 SetMandateReferenceNil sets the value for MandateReference to be an explicit nil

### UnsetMandateReference
`func (o *MandateActionResponseData) UnsetMandateReference()`

UnsetMandateReference ensures that no value is present for MandateReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


