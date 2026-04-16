# ProductsExternalBatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **int32** | Number of products created | 
**Updated** | **int32** | Number of products updated | 
**Skipped** | **int32** | Number of products skipped due to stale data | 
**Failed** | [**[]ProductsExternalBatch200ResponseFailedInner**](ProductsExternalBatch200ResponseFailedInner.md) | List of failed products with error messages | 

## Methods

### NewProductsExternalBatch200Response

`func NewProductsExternalBatch200Response(created int32, updated int32, skipped int32, failed []ProductsExternalBatch200ResponseFailedInner, ) *ProductsExternalBatch200Response`

NewProductsExternalBatch200Response instantiates a new ProductsExternalBatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductsExternalBatch200ResponseWithDefaults

`func NewProductsExternalBatch200ResponseWithDefaults() *ProductsExternalBatch200Response`

NewProductsExternalBatch200ResponseWithDefaults instantiates a new ProductsExternalBatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ProductsExternalBatch200Response) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ProductsExternalBatch200Response) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ProductsExternalBatch200Response) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *ProductsExternalBatch200Response) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ProductsExternalBatch200Response) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ProductsExternalBatch200Response) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.


### GetSkipped

`func (o *ProductsExternalBatch200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ProductsExternalBatch200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ProductsExternalBatch200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.


### GetFailed

`func (o *ProductsExternalBatch200Response) GetFailed() []ProductsExternalBatch200ResponseFailedInner`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ProductsExternalBatch200Response) GetFailedOk() (*[]ProductsExternalBatch200ResponseFailedInner, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ProductsExternalBatch200Response) SetFailed(v []ProductsExternalBatch200ResponseFailedInner)`

SetFailed sets Failed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


