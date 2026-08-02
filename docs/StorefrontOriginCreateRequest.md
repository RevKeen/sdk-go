# StorefrontOriginCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | **string** | Exact origin, e.g. https://shop.example.com or http://localhost:3000. | 

## Methods

### NewStorefrontOriginCreateRequest

`func NewStorefrontOriginCreateRequest(origin string, ) *StorefrontOriginCreateRequest`

NewStorefrontOriginCreateRequest instantiates a new StorefrontOriginCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontOriginCreateRequestWithDefaults

`func NewStorefrontOriginCreateRequestWithDefaults() *StorefrontOriginCreateRequest`

NewStorefrontOriginCreateRequestWithDefaults instantiates a new StorefrontOriginCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *StorefrontOriginCreateRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *StorefrontOriginCreateRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *StorefrontOriginCreateRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


