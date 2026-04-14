# FinanceGetIncome200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**FinanceGetIncome200ResponseSummary**](FinanceGetIncome200ResponseSummary.md) |  | 
**Rows** | **[]map[string]interface{}** | Income report rows | 

## Methods

### NewFinanceGetIncome200Response

`func NewFinanceGetIncome200Response(summary FinanceGetIncome200ResponseSummary, rows []map[string]interface{}, ) *FinanceGetIncome200Response`

NewFinanceGetIncome200Response instantiates a new FinanceGetIncome200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceGetIncome200ResponseWithDefaults

`func NewFinanceGetIncome200ResponseWithDefaults() *FinanceGetIncome200Response`

NewFinanceGetIncome200ResponseWithDefaults instantiates a new FinanceGetIncome200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *FinanceGetIncome200Response) GetSummary() FinanceGetIncome200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *FinanceGetIncome200Response) GetSummaryOk() (*FinanceGetIncome200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *FinanceGetIncome200Response) SetSummary(v FinanceGetIncome200ResponseSummary)`

SetSummary sets Summary field to given value.


### GetRows

`func (o *FinanceGetIncome200Response) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *FinanceGetIncome200Response) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *FinanceGetIncome200Response) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


