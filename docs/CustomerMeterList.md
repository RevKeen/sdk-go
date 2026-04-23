# CustomerMeterList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]CustomerMeter**](CustomerMeter.md) |  | 
**HasMore** | **bool** | Customer-meter list responses are not paginated — every meter the merchant has is returned for the specified customer. | 
**Url** | **string** |  | 

## Methods

### NewCustomerMeterList

`func NewCustomerMeterList(object string, data []CustomerMeter, hasMore bool, url string, ) *CustomerMeterList`

NewCustomerMeterList instantiates a new CustomerMeterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerMeterListWithDefaults

`func NewCustomerMeterListWithDefaults() *CustomerMeterList`

NewCustomerMeterListWithDefaults instantiates a new CustomerMeterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *CustomerMeterList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CustomerMeterList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CustomerMeterList) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *CustomerMeterList) GetData() []CustomerMeter`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerMeterList) GetDataOk() (*[]CustomerMeter, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerMeterList) SetData(v []CustomerMeter)`

SetData sets Data field to given value.


### GetHasMore

`func (o *CustomerMeterList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CustomerMeterList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CustomerMeterList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *CustomerMeterList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CustomerMeterList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CustomerMeterList) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


