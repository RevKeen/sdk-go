# TaxRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the tax rate | 
**Object** | **string** | Object type | 
**DisplayName** | **string** | Display name for the tax rate | 
**Description** | Pointer to **NullableString** | Description of the tax rate | [optional] 
**Jurisdiction** | Pointer to **NullableString** | Tax jurisdiction (e.g., &#39;US-CA&#39;, &#39;DE&#39;, &#39;GB&#39;) | [optional] 
**Percentage** | **float32** | Tax percentage (e.g., 8.25 for 8.25%) | 
**Inclusive** | **bool** | Whether tax is included in displayed price (EU style) | 
**TaxType** | **string** | Type of tax | 
**Country** | Pointer to **NullableString** | ISO country code | [optional] 
**State** | Pointer to **NullableString** | State/province code | [optional] 
**Active** | **bool** | Whether the tax rate is active | 
**Metadata** | Pointer to **map[string]interface{}** | Custom key-value metadata | [optional] 
**CreatedAt** | **time.Time** | Creation timestamp | 
**UpdatedAt** | **time.Time** | Last update timestamp | 

## Methods

### NewTaxRate

`func NewTaxRate(id string, object string, displayName string, percentage float32, inclusive bool, taxType string, active bool, createdAt time.Time, updatedAt time.Time, ) *TaxRate`

NewTaxRate instantiates a new TaxRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateWithDefaults

`func NewTaxRateWithDefaults() *TaxRate`

NewTaxRateWithDefaults instantiates a new TaxRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRate) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *TaxRate) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TaxRate) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TaxRate) SetObject(v string)`

SetObject sets Object field to given value.


### GetDisplayName

`func (o *TaxRate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TaxRate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TaxRate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *TaxRate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxRate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxRate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxRate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaxRate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaxRate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJurisdiction

`func (o *TaxRate) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *TaxRate) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *TaxRate) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *TaxRate) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### SetJurisdictionNil

`func (o *TaxRate) SetJurisdictionNil(b bool)`

 SetJurisdictionNil sets the value for Jurisdiction to be an explicit nil

### UnsetJurisdiction
`func (o *TaxRate) UnsetJurisdiction()`

UnsetJurisdiction ensures that no value is present for Jurisdiction, not even an explicit nil
### GetPercentage

`func (o *TaxRate) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *TaxRate) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *TaxRate) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetInclusive

`func (o *TaxRate) GetInclusive() bool`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *TaxRate) GetInclusiveOk() (*bool, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *TaxRate) SetInclusive(v bool)`

SetInclusive sets Inclusive field to given value.


### GetTaxType

`func (o *TaxRate) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *TaxRate) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *TaxRate) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.


### GetCountry

`func (o *TaxRate) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxRate) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxRate) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxRate) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *TaxRate) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *TaxRate) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetState

`func (o *TaxRate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaxRate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaxRate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TaxRate) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *TaxRate) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *TaxRate) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetActive

`func (o *TaxRate) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *TaxRate) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *TaxRate) SetActive(v bool)`

SetActive sets Active field to given value.


### GetMetadata

`func (o *TaxRate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TaxRate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TaxRate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TaxRate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaxRate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaxRate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaxRate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TaxRate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaxRate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaxRate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


