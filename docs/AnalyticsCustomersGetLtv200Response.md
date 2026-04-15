# AnalyticsCustomersGetLtv200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**[]AnalyticsCustomersGetLtv200ResponseCustomersInner**](AnalyticsCustomersGetLtv200ResponseCustomersInner.md) |  | 
**AvgLtv** | **float32** |  | 
**MedianLtv** | **float32** |  | 

## Methods

### NewAnalyticsCustomersGetLtv200Response

`func NewAnalyticsCustomersGetLtv200Response(customers []AnalyticsCustomersGetLtv200ResponseCustomersInner, avgLtv float32, medianLtv float32, ) *AnalyticsCustomersGetLtv200Response`

NewAnalyticsCustomersGetLtv200Response instantiates a new AnalyticsCustomersGetLtv200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsCustomersGetLtv200ResponseWithDefaults

`func NewAnalyticsCustomersGetLtv200ResponseWithDefaults() *AnalyticsCustomersGetLtv200Response`

NewAnalyticsCustomersGetLtv200ResponseWithDefaults instantiates a new AnalyticsCustomersGetLtv200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *AnalyticsCustomersGetLtv200Response) GetCustomers() []AnalyticsCustomersGetLtv200ResponseCustomersInner`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *AnalyticsCustomersGetLtv200Response) GetCustomersOk() (*[]AnalyticsCustomersGetLtv200ResponseCustomersInner, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *AnalyticsCustomersGetLtv200Response) SetCustomers(v []AnalyticsCustomersGetLtv200ResponseCustomersInner)`

SetCustomers sets Customers field to given value.


### GetAvgLtv

`func (o *AnalyticsCustomersGetLtv200Response) GetAvgLtv() float32`

GetAvgLtv returns the AvgLtv field if non-nil, zero value otherwise.

### GetAvgLtvOk

`func (o *AnalyticsCustomersGetLtv200Response) GetAvgLtvOk() (*float32, bool)`

GetAvgLtvOk returns a tuple with the AvgLtv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgLtv

`func (o *AnalyticsCustomersGetLtv200Response) SetAvgLtv(v float32)`

SetAvgLtv sets AvgLtv field to given value.


### GetMedianLtv

`func (o *AnalyticsCustomersGetLtv200Response) GetMedianLtv() float32`

GetMedianLtv returns the MedianLtv field if non-nil, zero value otherwise.

### GetMedianLtvOk

`func (o *AnalyticsCustomersGetLtv200Response) GetMedianLtvOk() (*float32, bool)`

GetMedianLtvOk returns a tuple with the MedianLtv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianLtv

`func (o *AnalyticsCustomersGetLtv200Response) SetMedianLtv(v float32)`

SetMedianLtv sets MedianLtv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


