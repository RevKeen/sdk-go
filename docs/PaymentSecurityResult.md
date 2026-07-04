# PaymentSecurityResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avs** | [**AvsCodeInterpretation**](AvsCodeInterpretation.md) |  | 
**Cvv** | [**CvvCodeInterpretation**](CvvCodeInterpretation.md) |  | 
**OverallRiskLevel** | **string** | Combined risk level from AVS + CVV checks | 
**Recommendation** | **string** | Recommended action based on security checks | 

## Methods

### NewPaymentSecurityResult

`func NewPaymentSecurityResult(avs AvsCodeInterpretation, cvv CvvCodeInterpretation, overallRiskLevel string, recommendation string, ) *PaymentSecurityResult`

NewPaymentSecurityResult instantiates a new PaymentSecurityResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSecurityResultWithDefaults

`func NewPaymentSecurityResultWithDefaults() *PaymentSecurityResult`

NewPaymentSecurityResultWithDefaults instantiates a new PaymentSecurityResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvs

`func (o *PaymentSecurityResult) GetAvs() AvsCodeInterpretation`

GetAvs returns the Avs field if non-nil, zero value otherwise.

### GetAvsOk

`func (o *PaymentSecurityResult) GetAvsOk() (*AvsCodeInterpretation, bool)`

GetAvsOk returns a tuple with the Avs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvs

`func (o *PaymentSecurityResult) SetAvs(v AvsCodeInterpretation)`

SetAvs sets Avs field to given value.


### GetCvv

`func (o *PaymentSecurityResult) GetCvv() CvvCodeInterpretation`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *PaymentSecurityResult) GetCvvOk() (*CvvCodeInterpretation, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *PaymentSecurityResult) SetCvv(v CvvCodeInterpretation)`

SetCvv sets Cvv field to given value.


### GetOverallRiskLevel

`func (o *PaymentSecurityResult) GetOverallRiskLevel() string`

GetOverallRiskLevel returns the OverallRiskLevel field if non-nil, zero value otherwise.

### GetOverallRiskLevelOk

`func (o *PaymentSecurityResult) GetOverallRiskLevelOk() (*string, bool)`

GetOverallRiskLevelOk returns a tuple with the OverallRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallRiskLevel

`func (o *PaymentSecurityResult) SetOverallRiskLevel(v string)`

SetOverallRiskLevel sets OverallRiskLevel field to given value.


### GetRecommendation

`func (o *PaymentSecurityResult) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *PaymentSecurityResult) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *PaymentSecurityResult) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


