# PaymentMethodList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Object type, always &#39;list&#39; | 
**Data** | [**[]PaymentMethod**](PaymentMethod.md) | Array of payment methods | 
**HasMore** | **bool** | Whether there are more results available | 
**Url** | **string** | URL for this list resource | 

## Methods

### NewPaymentMethodList

`func NewPaymentMethodList(object string, data []PaymentMethod, hasMore bool, url string, ) *PaymentMethodList`

NewPaymentMethodList instantiates a new PaymentMethodList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodListWithDefaults

`func NewPaymentMethodListWithDefaults() *PaymentMethodList`

NewPaymentMethodListWithDefaults instantiates a new PaymentMethodList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *PaymentMethodList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodList) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *PaymentMethodList) GetData() []PaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentMethodList) GetDataOk() (*[]PaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentMethodList) SetData(v []PaymentMethod)`

SetData sets Data field to given value.


### GetHasMore

`func (o *PaymentMethodList) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PaymentMethodList) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PaymentMethodList) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetUrl

`func (o *PaymentMethodList) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaymentMethodList) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaymentMethodList) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


