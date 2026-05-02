# UpdateTaxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name for the tax rate | [optional] 
**Description** | Pointer to **NullableString** | Description of the tax rate | [optional] 
**Jurisdiction** | Pointer to **NullableString** | Tax jurisdiction | [optional] 
**Country** | Pointer to **NullableString** | ISO 2-letter country code | [optional] 
**State** | Pointer to **NullableString** | State/province code | [optional] 
**Active** | Pointer to **bool** | Whether the tax rate is active | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewUpdateTaxRateRequest

`func NewUpdateTaxRateRequest() *UpdateTaxRateRequest`

NewUpdateTaxRateRequest instantiates a new UpdateTaxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaxRateRequestWithDefaults

`func NewUpdateTaxRateRequestWithDefaults() *UpdateTaxRateRequest`

NewUpdateTaxRateRequestWithDefaults instantiates a new UpdateTaxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateTaxRateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateTaxRateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateTaxRateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateTaxRateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTaxRateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTaxRateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTaxRateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTaxRateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTaxRateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTaxRateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJurisdiction

`func (o *UpdateTaxRateRequest) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *UpdateTaxRateRequest) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *UpdateTaxRateRequest) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *UpdateTaxRateRequest) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### SetJurisdictionNil

`func (o *UpdateTaxRateRequest) SetJurisdictionNil(b bool)`

 SetJurisdictionNil sets the value for Jurisdiction to be an explicit nil

### UnsetJurisdiction
`func (o *UpdateTaxRateRequest) UnsetJurisdiction()`

UnsetJurisdiction ensures that no value is present for Jurisdiction, not even an explicit nil
### GetCountry

`func (o *UpdateTaxRateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateTaxRateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateTaxRateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateTaxRateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *UpdateTaxRateRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *UpdateTaxRateRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *UpdateTaxRateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateTaxRateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateTaxRateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateTaxRateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UpdateTaxRateRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UpdateTaxRateRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetActive

`func (o *UpdateTaxRateRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateTaxRateRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateTaxRateRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateTaxRateRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateTaxRateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateTaxRateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateTaxRateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateTaxRateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


