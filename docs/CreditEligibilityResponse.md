# CreditEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**CreditEligibilityResponseData**](CreditEligibilityResponseData.md) |  | 
**RequestId** | **string** |  | 

## Methods

### NewCreditEligibilityResponse

`func NewCreditEligibilityResponse(data CreditEligibilityResponseData, requestId string, ) *CreditEligibilityResponse`

NewCreditEligibilityResponse instantiates a new CreditEligibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditEligibilityResponseWithDefaults

`func NewCreditEligibilityResponseWithDefaults() *CreditEligibilityResponse`

NewCreditEligibilityResponseWithDefaults instantiates a new CreditEligibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreditEligibilityResponse) GetData() CreditEligibilityResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreditEligibilityResponse) GetDataOk() (*CreditEligibilityResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreditEligibilityResponse) SetData(v CreditEligibilityResponseData)`

SetData sets Data field to given value.


### GetRequestId

`func (o *CreditEligibilityResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreditEligibilityResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreditEligibilityResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


