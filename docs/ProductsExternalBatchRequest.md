# ProductsExternalBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Integration source (e.g., practicehub, wodify) | 
**Products** | [**[]ProductsExternalBatchRequestProductsInner**](ProductsExternalBatchRequestProductsInner.md) | Array of products to upsert (max 100) | 

## Methods

### NewProductsExternalBatchRequest

`func NewProductsExternalBatchRequest(source string, products []ProductsExternalBatchRequestProductsInner, ) *ProductsExternalBatchRequest`

NewProductsExternalBatchRequest instantiates a new ProductsExternalBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsExternalBatchRequestWithDefaults

`func NewProductsExternalBatchRequestWithDefaults() *ProductsExternalBatchRequest`

NewProductsExternalBatchRequestWithDefaults instantiates a new ProductsExternalBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ProductsExternalBatchRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProductsExternalBatchRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProductsExternalBatchRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetProducts

`func (o *ProductsExternalBatchRequest) GetProducts() []ProductsExternalBatchRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductsExternalBatchRequest) GetProductsOk() (*[]ProductsExternalBatchRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductsExternalBatchRequest) SetProducts(v []ProductsExternalBatchRequestProductsInner)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


