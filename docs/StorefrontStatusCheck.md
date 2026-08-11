# StorefrontStatusCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 
**Message** | **string** |  | 
**NextAction** | Pointer to **string** |  | [optional] 

## Methods

### NewStorefrontStatusCheck

`func NewStorefrontStatusCheck(id string, status string, message string, ) *StorefrontStatusCheck`

NewStorefrontStatusCheck instantiates a new StorefrontStatusCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorefrontStatusCheckWithDefaults

`func NewStorefrontStatusCheckWithDefaults() *StorefrontStatusCheck`

NewStorefrontStatusCheckWithDefaults instantiates a new StorefrontStatusCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorefrontStatusCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorefrontStatusCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorefrontStatusCheck) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *StorefrontStatusCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorefrontStatusCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorefrontStatusCheck) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCode

`func (o *StorefrontStatusCheck) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StorefrontStatusCheck) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StorefrontStatusCheck) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StorefrontStatusCheck) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *StorefrontStatusCheck) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StorefrontStatusCheck) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StorefrontStatusCheck) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNextAction

`func (o *StorefrontStatusCheck) GetNextAction() string`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *StorefrontStatusCheck) GetNextActionOk() (*string, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *StorefrontStatusCheck) SetNextAction(v string)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *StorefrontStatusCheck) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


