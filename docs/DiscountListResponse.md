# DiscountListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Discount**](Discount.md) |  | 
**Meta** | [**DiscountListResponseMeta**](DiscountListResponseMeta.md) |  | 

## Methods

### NewDiscountListResponse

`func NewDiscountListResponse(data []Discount, meta DiscountListResponseMeta, ) *DiscountListResponse`

NewDiscountListResponse instantiates a new DiscountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountListResponseWithDefaults

`func NewDiscountListResponseWithDefaults() *DiscountListResponse`

NewDiscountListResponseWithDefaults instantiates a new DiscountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DiscountListResponse) GetData() []Discount`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DiscountListResponse) GetDataOk() (*[]Discount, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DiscountListResponse) SetData(v []Discount)`

SetData sets Data field to given value.


### GetMeta

`func (o *DiscountListResponse) GetMeta() DiscountListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiscountListResponse) GetMetaOk() (*DiscountListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiscountListResponse) SetMeta(v DiscountListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


