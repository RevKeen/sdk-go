# ProductCollectionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** |  |
**Code** | Pointer to **string** |  | [optional]
**ProductIds** | Pointer to **[]string** |  | [optional]

## Methods

### NewProductCollectionError

`func NewProductCollectionError(error_ string, ) *ProductCollectionError`

NewProductCollectionError instantiates a new ProductCollectionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCollectionErrorWithDefaults

`func NewProductCollectionErrorWithDefaults() *ProductCollectionError`

NewProductCollectionErrorWithDefaults instantiates a new ProductCollectionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ProductCollectionError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProductCollectionError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProductCollectionError) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *ProductCollectionError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProductCollectionError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProductCollectionError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProductCollectionError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetProductIds

`func (o *ProductCollectionError) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *ProductCollectionError) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *ProductCollectionError) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.

### HasProductIds

`func (o *ProductCollectionError) HasProductIds() bool`

HasProductIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
