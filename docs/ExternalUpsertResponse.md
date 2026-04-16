# ExternalUpsertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Invoice**](Invoice.md) |  | 
**Created** | **bool** |  | 
**Warnings** | Pointer to **[]string** |  | [optional] 
**RequestId** | **string** |  | 

## Methods

### NewExternalUpsertResponse

`func NewExternalUpsertResponse(data Invoice, created bool, requestId string, ) *ExternalUpsertResponse`

NewExternalUpsertResponse instantiates a new ExternalUpsertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUpsertResponseWithDefaults

`func NewExternalUpsertResponseWithDefaults() *ExternalUpsertResponse`

NewExternalUpsertResponseWithDefaults instantiates a new ExternalUpsertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ExternalUpsertResponse) GetData() Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExternalUpsertResponse) GetDataOk() (*Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExternalUpsertResponse) SetData(v Invoice)`

SetData sets Data field to given value.


### GetCreated

`func (o *ExternalUpsertResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExternalUpsertResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExternalUpsertResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.


### GetWarnings

`func (o *ExternalUpsertResponse) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ExternalUpsertResponse) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ExternalUpsertResponse) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ExternalUpsertResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetRequestId

`func (o *ExternalUpsertResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ExternalUpsertResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ExternalUpsertResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


