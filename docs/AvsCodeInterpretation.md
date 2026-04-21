# AvsCodeInterpretation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Raw AVS response code from gateway | 
**MatchType** | **string** | Type of address match | 
**AddressMatch** | **NullableBool** | Whether the street address matched | 
**ZipMatch** | **NullableBool** | Whether the ZIP/postal code matched | 
**Description** | **string** | Human-readable description of the AVS result | 
**RiskLevel** | **string** | Risk level based on AVS result. low: full match, medium: partial match, high: no match | 

## Methods

### NewAvsCodeInterpretation

`func NewAvsCodeInterpretation(code string, matchType string, addressMatch NullableBool, zipMatch NullableBool, description string, riskLevel string, ) *AvsCodeInterpretation`

NewAvsCodeInterpretation instantiates a new AvsCodeInterpretation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvsCodeInterpretationWithDefaults

`func NewAvsCodeInterpretationWithDefaults() *AvsCodeInterpretation`

NewAvsCodeInterpretationWithDefaults instantiates a new AvsCodeInterpretation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AvsCodeInterpretation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AvsCodeInterpretation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AvsCodeInterpretation) SetCode(v string)`

SetCode sets Code field to given value.


### GetMatchType

`func (o *AvsCodeInterpretation) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *AvsCodeInterpretation) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *AvsCodeInterpretation) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetAddressMatch

`func (o *AvsCodeInterpretation) GetAddressMatch() bool`

GetAddressMatch returns the AddressMatch field if non-nil, zero value otherwise.

### GetAddressMatchOk

`func (o *AvsCodeInterpretation) GetAddressMatchOk() (*bool, bool)`

GetAddressMatchOk returns a tuple with the AddressMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMatch

`func (o *AvsCodeInterpretation) SetAddressMatch(v bool)`

SetAddressMatch sets AddressMatch field to given value.


### SetAddressMatchNil

`func (o *AvsCodeInterpretation) SetAddressMatchNil(b bool)`

 SetAddressMatchNil sets the value for AddressMatch to be an explicit nil

### UnsetAddressMatch
`func (o *AvsCodeInterpretation) UnsetAddressMatch()`

UnsetAddressMatch ensures that no value is present for AddressMatch, not even an explicit nil
### GetZipMatch

`func (o *AvsCodeInterpretation) GetZipMatch() bool`

GetZipMatch returns the ZipMatch field if non-nil, zero value otherwise.

### GetZipMatchOk

`func (o *AvsCodeInterpretation) GetZipMatchOk() (*bool, bool)`

GetZipMatchOk returns a tuple with the ZipMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipMatch

`func (o *AvsCodeInterpretation) SetZipMatch(v bool)`

SetZipMatch sets ZipMatch field to given value.


### SetZipMatchNil

`func (o *AvsCodeInterpretation) SetZipMatchNil(b bool)`

 SetZipMatchNil sets the value for ZipMatch to be an explicit nil

### UnsetZipMatch
`func (o *AvsCodeInterpretation) UnsetZipMatch()`

UnsetZipMatch ensures that no value is present for ZipMatch, not even an explicit nil
### GetDescription

`func (o *AvsCodeInterpretation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AvsCodeInterpretation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AvsCodeInterpretation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRiskLevel

`func (o *AvsCodeInterpretation) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *AvsCodeInterpretation) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *AvsCodeInterpretation) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


