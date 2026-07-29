# DeclineAnalyticsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalAttempts** | **int32** |  | 
**ApprovedCount** | **int32** |  | 
**DeclinedCount** | **int32** |  | 
**ApprovalRate** | **float32** |  | 
**SoftDeclineCount** | **int32** |  | 
**HardDeclineCount** | **int32** |  | 
**TopDeclineCategories** | [**[]DeclineAnalyticsResponseDataTopDeclineCategoriesInner**](DeclineAnalyticsResponseDataTopDeclineCategoriesInner.md) |  | 
**ByGateway** | [**map[string]DeclineAnalyticsResponseDataByGatewayValue**](DeclineAnalyticsResponseDataByGatewayValue.md) | Breakdown by gateway | 

## Methods

### NewDeclineAnalyticsResponseData

`func NewDeclineAnalyticsResponseData(totalAttempts int32, approvedCount int32, declinedCount int32, approvalRate float32, softDeclineCount int32, hardDeclineCount int32, topDeclineCategories []DeclineAnalyticsResponseDataTopDeclineCategoriesInner, byGateway map[string]DeclineAnalyticsResponseDataByGatewayValue, ) *DeclineAnalyticsResponseData`

NewDeclineAnalyticsResponseData instantiates a new DeclineAnalyticsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeclineAnalyticsResponseDataWithDefaults

`func NewDeclineAnalyticsResponseDataWithDefaults() *DeclineAnalyticsResponseData`

NewDeclineAnalyticsResponseDataWithDefaults instantiates a new DeclineAnalyticsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalAttempts

`func (o *DeclineAnalyticsResponseData) GetTotalAttempts() int32`

GetTotalAttempts returns the TotalAttempts field if non-nil, zero value otherwise.

### GetTotalAttemptsOk

`func (o *DeclineAnalyticsResponseData) GetTotalAttemptsOk() (*int32, bool)`

GetTotalAttemptsOk returns a tuple with the TotalAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAttempts

`func (o *DeclineAnalyticsResponseData) SetTotalAttempts(v int32)`

SetTotalAttempts sets TotalAttempts field to given value.


### GetApprovedCount

`func (o *DeclineAnalyticsResponseData) GetApprovedCount() int32`

GetApprovedCount returns the ApprovedCount field if non-nil, zero value otherwise.

### GetApprovedCountOk

`func (o *DeclineAnalyticsResponseData) GetApprovedCountOk() (*int32, bool)`

GetApprovedCountOk returns a tuple with the ApprovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedCount

`func (o *DeclineAnalyticsResponseData) SetApprovedCount(v int32)`

SetApprovedCount sets ApprovedCount field to given value.


### GetDeclinedCount

`func (o *DeclineAnalyticsResponseData) GetDeclinedCount() int32`

GetDeclinedCount returns the DeclinedCount field if non-nil, zero value otherwise.

### GetDeclinedCountOk

`func (o *DeclineAnalyticsResponseData) GetDeclinedCountOk() (*int32, bool)`

GetDeclinedCountOk returns a tuple with the DeclinedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedCount

`func (o *DeclineAnalyticsResponseData) SetDeclinedCount(v int32)`

SetDeclinedCount sets DeclinedCount field to given value.


### GetApprovalRate

`func (o *DeclineAnalyticsResponseData) GetApprovalRate() float32`

GetApprovalRate returns the ApprovalRate field if non-nil, zero value otherwise.

### GetApprovalRateOk

`func (o *DeclineAnalyticsResponseData) GetApprovalRateOk() (*float32, bool)`

GetApprovalRateOk returns a tuple with the ApprovalRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRate

`func (o *DeclineAnalyticsResponseData) SetApprovalRate(v float32)`

SetApprovalRate sets ApprovalRate field to given value.


### GetSoftDeclineCount

`func (o *DeclineAnalyticsResponseData) GetSoftDeclineCount() int32`

GetSoftDeclineCount returns the SoftDeclineCount field if non-nil, zero value otherwise.

### GetSoftDeclineCountOk

`func (o *DeclineAnalyticsResponseData) GetSoftDeclineCountOk() (*int32, bool)`

GetSoftDeclineCountOk returns a tuple with the SoftDeclineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeclineCount

`func (o *DeclineAnalyticsResponseData) SetSoftDeclineCount(v int32)`

SetSoftDeclineCount sets SoftDeclineCount field to given value.


### GetHardDeclineCount

`func (o *DeclineAnalyticsResponseData) GetHardDeclineCount() int32`

GetHardDeclineCount returns the HardDeclineCount field if non-nil, zero value otherwise.

### GetHardDeclineCountOk

`func (o *DeclineAnalyticsResponseData) GetHardDeclineCountOk() (*int32, bool)`

GetHardDeclineCountOk returns a tuple with the HardDeclineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeclineCount

`func (o *DeclineAnalyticsResponseData) SetHardDeclineCount(v int32)`

SetHardDeclineCount sets HardDeclineCount field to given value.


### GetTopDeclineCategories

`func (o *DeclineAnalyticsResponseData) GetTopDeclineCategories() []DeclineAnalyticsResponseDataTopDeclineCategoriesInner`

GetTopDeclineCategories returns the TopDeclineCategories field if non-nil, zero value otherwise.

### GetTopDeclineCategoriesOk

`func (o *DeclineAnalyticsResponseData) GetTopDeclineCategoriesOk() (*[]DeclineAnalyticsResponseDataTopDeclineCategoriesInner, bool)`

GetTopDeclineCategoriesOk returns a tuple with the TopDeclineCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopDeclineCategories

`func (o *DeclineAnalyticsResponseData) SetTopDeclineCategories(v []DeclineAnalyticsResponseDataTopDeclineCategoriesInner)`

SetTopDeclineCategories sets TopDeclineCategories field to given value.


### GetByGateway

`func (o *DeclineAnalyticsResponseData) GetByGateway() map[string]DeclineAnalyticsResponseDataByGatewayValue`

GetByGateway returns the ByGateway field if non-nil, zero value otherwise.

### GetByGatewayOk

`func (o *DeclineAnalyticsResponseData) GetByGatewayOk() (*map[string]DeclineAnalyticsResponseDataByGatewayValue, bool)`

GetByGatewayOk returns a tuple with the ByGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByGateway

`func (o *DeclineAnalyticsResponseData) SetByGateway(v map[string]DeclineAnalyticsResponseDataByGatewayValue)`

SetByGateway sets ByGateway field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


