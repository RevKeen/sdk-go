# PriceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]Price**](Price.md) |  | 
**HasMore** | **bool** |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPriceListResponse

`func NewPriceListResponse(object string, data []Price, hasMore bool, ) *PriceListResponse`

NewPriceListResponse instantiates a new PriceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceListResponseWithDefaults

`func NewPriceListResponseWithDefaults() *PriceListResponse`

NewPriceListResponseWithDefaults instantiates a new PriceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PriceListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PriceListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PriceListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PriceListResponse) GetData() []Price`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PriceListResponse) GetDataOk() (*[]Price, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PriceListResponse) SetData(v []Price)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PriceListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PriceListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PriceListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *PriceListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PriceListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PriceListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PriceListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


