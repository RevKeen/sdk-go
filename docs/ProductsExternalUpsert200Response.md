# ProductsExternalUpsert200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**ProductsExternalUpsert200ResponseData**](ProductsExternalUpsert200ResponseData.md) |  | 
**Created** | **bool** |  | 
**Skipped** | Pointer to **bool** |  | [optional] 

## Methods

### NewProductsExternalUpsert200Response

`func NewProductsExternalUpsert200Response(data ProductsExternalUpsert200ResponseData, created bool, ) *ProductsExternalUpsert200Response`

NewProductsExternalUpsert200Response instantiates a new ProductsExternalUpsert200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsExternalUpsert200ResponseWithDefaults

`func NewProductsExternalUpsert200ResponseWithDefaults() *ProductsExternalUpsert200Response`

NewProductsExternalUpsert200ResponseWithDefaults instantiates a new ProductsExternalUpsert200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProductsExternalUpsert200Response) GetData() ProductsExternalUpsert200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductsExternalUpsert200Response) GetDataOk() (*ProductsExternalUpsert200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductsExternalUpsert200Response) SetData(v ProductsExternalUpsert200ResponseData)`

SetData sets Data field to given value.


### GetCreated

`func (o *ProductsExternalUpsert200Response) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProductsExternalUpsert200Response) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProductsExternalUpsert200Response) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetSkipped

`func (o *ProductsExternalUpsert200Response) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ProductsExternalUpsert200Response) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ProductsExternalUpsert200Response) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *ProductsExternalUpsert200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


