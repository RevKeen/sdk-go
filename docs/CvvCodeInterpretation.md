# CvvCodeInterpretation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Raw CVV response code from gateway | 
**MatchType** | **string** | Type of CVV match | 
**Description** | **string** | Human-readable description of the CVV result | 
**RiskLevel** | **string** | Risk level based on CVV result. low: match, medium: not processed, high: no match | 

## Methods

### NewCvvCodeInterpretation

`func NewCvvCodeInterpretation(code string, matchType string, description string, riskLevel string, ) *CvvCodeInterpretation`

NewCvvCodeInterpretation instantiates a new CvvCodeInterpretation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvvCodeInterpretationWithDefaults

`func NewCvvCodeInterpretationWithDefaults() *CvvCodeInterpretation`

NewCvvCodeInterpretationWithDefaults instantiates a new CvvCodeInterpretation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CvvCodeInterpretation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CvvCodeInterpretation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CvvCodeInterpretation) SetCode(v string)`

SetCode sets Code field to given value.


### GetMatchType

`func (o *CvvCodeInterpretation) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *CvvCodeInterpretation) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *CvvCodeInterpretation) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetDescription

`func (o *CvvCodeInterpretation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CvvCodeInterpretation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CvvCodeInterpretation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRiskLevel

`func (o *CvvCodeInterpretation) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *CvvCodeInterpretation) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *CvvCodeInterpretation) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


