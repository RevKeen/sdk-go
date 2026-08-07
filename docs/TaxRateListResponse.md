# TaxRateListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]TaxRate**](TaxRate.md) |  | 
**HasMore** | **bool** |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewTaxRateListResponse

`func NewTaxRateListResponse(object string, data []TaxRate, hasMore bool, ) *TaxRateListResponse`

NewTaxRateListResponse instantiates a new TaxRateListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateListResponseWithDefaults

`func NewTaxRateListResponseWithDefaults() *TaxRateListResponse`

NewTaxRateListResponseWithDefaults instantiates a new TaxRateListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *TaxRateListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TaxRateListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TaxRateListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *TaxRateListResponse) GetData() []TaxRate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TaxRateListResponse) GetDataOk() (*[]TaxRate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TaxRateListResponse) SetData(v []TaxRate)`

SetData sets Data field to given value.


### GetHasMore

`func (o *TaxRateListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TaxRateListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TaxRateListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *TaxRateListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TaxRateListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TaxRateListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TaxRateListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


