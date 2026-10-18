# AnalyticsRevenueMrrSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentMrr** | **int64** | Current Monthly Recurring Revenue in minor units |
**PreviousMrr** | **int64** | Previous period MRR in minor units |
**Growth** | **int64** | Net MRR growth in minor units |
**NewMrr** | **int64** | MRR from new subscriptions |
**ChurnedMrr** | **int64** | MRR lost to churn |
**ExpansionMrr** | **int64** | MRR gained from upgrades |
**ContractionMrr** | **int64** | MRR lost to downgrades |
**AsOfDate** | **string** | Date the MRR was calculated as of |

## Methods

### NewAnalyticsRevenueMrrSummary200Response

`func NewAnalyticsRevenueMrrSummary200Response(currentMrr int64, previousMrr int64, growth int64, newMrr int64, churnedMrr int64, expansionMrr int64, contractionMrr int64, asOfDate string, ) *AnalyticsRevenueMrrSummary200Response`

NewAnalyticsRevenueMrrSummary200Response instantiates a new AnalyticsRevenueMrrSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsRevenueMrrSummary200ResponseWithDefaults

`func NewAnalyticsRevenueMrrSummary200ResponseWithDefaults() *AnalyticsRevenueMrrSummary200Response`

NewAnalyticsRevenueMrrSummary200ResponseWithDefaults instantiates a new AnalyticsRevenueMrrSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetCurrentMrr() int64`

GetCurrentMrr returns the CurrentMrr field if non-nil, zero value otherwise.

### GetCurrentMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetCurrentMrrOk() (*int64, bool)`

GetCurrentMrrOk returns a tuple with the CurrentMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetCurrentMrr(v int64)`

SetCurrentMrr sets CurrentMrr field to given value.


### GetPreviousMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetPreviousMrr() int64`

GetPreviousMrr returns the PreviousMrr field if non-nil, zero value otherwise.

### GetPreviousMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetPreviousMrrOk() (*int64, bool)`

GetPreviousMrrOk returns a tuple with the PreviousMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetPreviousMrr(v int64)`

SetPreviousMrr sets PreviousMrr field to given value.


### GetGrowth

`func (o *AnalyticsRevenueMrrSummary200Response) GetGrowth() int64`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetGrowthOk() (*int64, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *AnalyticsRevenueMrrSummary200Response) SetGrowth(v int64)`

SetGrowth sets Growth field to given value.


### GetNewMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetNewMrr() int64`

GetNewMrr returns the NewMrr field if non-nil, zero value otherwise.

### GetNewMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetNewMrrOk() (*int64, bool)`

GetNewMrrOk returns a tuple with the NewMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetNewMrr(v int64)`

SetNewMrr sets NewMrr field to given value.


### GetChurnedMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetChurnedMrr() int64`

GetChurnedMrr returns the ChurnedMrr field if non-nil, zero value otherwise.

### GetChurnedMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetChurnedMrrOk() (*int64, bool)`

GetChurnedMrrOk returns a tuple with the ChurnedMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnedMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetChurnedMrr(v int64)`

SetChurnedMrr sets ChurnedMrr field to given value.


### GetExpansionMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetExpansionMrr() int64`

GetExpansionMrr returns the ExpansionMrr field if non-nil, zero value otherwise.

### GetExpansionMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetExpansionMrrOk() (*int64, bool)`

GetExpansionMrrOk returns a tuple with the ExpansionMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansionMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetExpansionMrr(v int64)`

SetExpansionMrr sets ExpansionMrr field to given value.


### GetContractionMrr

`func (o *AnalyticsRevenueMrrSummary200Response) GetContractionMrr() int64`

GetContractionMrr returns the ContractionMrr field if non-nil, zero value otherwise.

### GetContractionMrrOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetContractionMrrOk() (*int64, bool)`

GetContractionMrrOk returns a tuple with the ContractionMrr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractionMrr

`func (o *AnalyticsRevenueMrrSummary200Response) SetContractionMrr(v int64)`

SetContractionMrr sets ContractionMrr field to given value.


### GetAsOfDate

`func (o *AnalyticsRevenueMrrSummary200Response) GetAsOfDate() string`

GetAsOfDate returns the AsOfDate field if non-nil, zero value otherwise.

### GetAsOfDateOk

`func (o *AnalyticsRevenueMrrSummary200Response) GetAsOfDateOk() (*string, bool)`

GetAsOfDateOk returns a tuple with the AsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfDate

`func (o *AnalyticsRevenueMrrSummary200Response) SetAsOfDate(v string)`

SetAsOfDate sets AsOfDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
