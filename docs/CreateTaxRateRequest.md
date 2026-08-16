# CreateTaxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name for the tax rate | 
**Description** | Pointer to **string** | Description of the tax rate | [optional] 
**Jurisdiction** | Pointer to **string** | Tax jurisdiction (e.g., &#39;US-CA&#39;, &#39;DE&#39;, &#39;GB&#39;) | [optional] 
**Percentage** | **float32** | Tax percentage (e.g., 8.25 for 8.25%) | 
**Inclusive** | Pointer to **bool** | Whether tax is included in displayed price | [optional] [default to false]
**TaxType** | Pointer to **string** | Type of tax | [optional] [default to "vat"]
**Country** | Pointer to **string** | ISO 2-letter country code | [optional] 
**State** | Pointer to **string** | State/province code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 

## Methods

### NewCreateTaxRateRequest

`func NewCreateTaxRateRequest(displayName string, percentage float32, ) *CreateTaxRateRequest`

NewCreateTaxRateRequest instantiates a new CreateTaxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaxRateRequestWithDefaults

`func NewCreateTaxRateRequestWithDefaults() *CreateTaxRateRequest`

NewCreateTaxRateRequestWithDefaults instantiates a new CreateTaxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *CreateTaxRateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateTaxRateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateTaxRateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *CreateTaxRateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTaxRateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTaxRateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTaxRateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJurisdiction

`func (o *CreateTaxRateRequest) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *CreateTaxRateRequest) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *CreateTaxRateRequest) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *CreateTaxRateRequest) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetPercentage

`func (o *CreateTaxRateRequest) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *CreateTaxRateRequest) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *CreateTaxRateRequest) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetInclusive

`func (o *CreateTaxRateRequest) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *CreateTaxRateRequest) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *CreateTaxRateRequest) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *CreateTaxRateRequest) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.

### GetTaxType

`func (o *CreateTaxRateRequest) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *CreateTaxRateRequest) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *CreateTaxRateRequest) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *CreateTaxRateRequest) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.

### GetCountry

`func (o *CreateTaxRateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateTaxRateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateTaxRateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateTaxRateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *CreateTaxRateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateTaxRateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateTaxRateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateTaxRateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateTaxRateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateTaxRateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateTaxRateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateTaxRateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


