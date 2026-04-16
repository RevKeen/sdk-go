# AnalyticsInvoicesGetArAging200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOfDate** | **string** |  | 
**Buckets** | [**[]AnalyticsInvoicesGetArAging200ResponseBucketsInner**](AnalyticsInvoicesGetArAging200ResponseBucketsInner.md) |  | 
**TotalOutstanding** | **float32** |  | 
**TotalInvoices** | **float32** |  | 

## Methods

### NewAnalyticsInvoicesGetArAging200Response

`func NewAnalyticsInvoicesGetArAging200Response(asOfDate string, buckets []AnalyticsInvoicesGetArAging200ResponseBucketsInner, totalOutstanding float32, totalInvoices float32, ) *AnalyticsInvoicesGetArAging200Response`

NewAnalyticsInvoicesGetArAging200Response instantiates a new AnalyticsInvoicesGetArAging200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsInvoicesGetArAging200ResponseWithDefaults

`func NewAnalyticsInvoicesGetArAging200ResponseWithDefaults() *AnalyticsInvoicesGetArAging200Response`

NewAnalyticsInvoicesGetArAging200ResponseWithDefaults instantiates a new AnalyticsInvoicesGetArAging200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOfDate

`func (o *AnalyticsInvoicesGetArAging200Response) GetAsOfDate() string`

GetAsOfDate returns the AsOfDate field if non-nil, zero value otherwise.

### GetAsOfDateOk

`func (o *AnalyticsInvoicesGetArAging200Response) GetAsOfDateOk() (*string, bool)`

GetAsOfDateOk returns a tuple with the AsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOfDate

`func (o *AnalyticsInvoicesGetArAging200Response) SetAsOfDate(v string)`

SetAsOfDate sets AsOfDate field to given value.


### GetBuckets

`func (o *AnalyticsInvoicesGetArAging200Response) GetBuckets() []AnalyticsInvoicesGetArAging200ResponseBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AnalyticsInvoicesGetArAging200Response) GetBucketsOk() (*[]AnalyticsInvoicesGetArAging200ResponseBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AnalyticsInvoicesGetArAging200Response) SetBuckets(v []AnalyticsInvoicesGetArAging200ResponseBucketsInner)`

SetBuckets sets Buckets field to given value.


### GetTotalOutstanding

`func (o *AnalyticsInvoicesGetArAging200Response) GetTotalOutstanding() float32`

GetTotalOutstanding returns the TotalOutstanding field if non-nil, zero value otherwise.

### GetTotalOutstandingOk

`func (o *AnalyticsInvoicesGetArAging200Response) GetTotalOutstandingOk() (*float32, bool)`

GetTotalOutstandingOk returns a tuple with the TotalOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutstanding

`func (o *AnalyticsInvoicesGetArAging200Response) SetTotalOutstanding(v float32)`

SetTotalOutstanding sets TotalOutstanding field to given value.


### GetTotalInvoices

`func (o *AnalyticsInvoicesGetArAging200Response) GetTotalInvoices() float32`

GetTotalInvoices returns the TotalInvoices field if non-nil, zero value otherwise.

### GetTotalInvoicesOk

`func (o *AnalyticsInvoicesGetArAging200Response) GetTotalInvoicesOk() (*float32, bool)`

GetTotalInvoicesOk returns a tuple with the TotalInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvoices

`func (o *AnalyticsInvoicesGetArAging200Response) SetTotalInvoices(v float32)`

SetTotalInvoices sets TotalInvoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


