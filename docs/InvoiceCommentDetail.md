# InvoiceCommentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InvoiceId** | **string** |  | 
**UserId** | **NullableString** |  | 
**Content** | **string** |  | 
**IsInternal** | **bool** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewInvoiceCommentDetail

`func NewInvoiceCommentDetail(id string, invoiceId string, userId NullableString, content string, isInternal bool, createdAt string, updatedAt NullableString, ) *InvoiceCommentDetail`

NewInvoiceCommentDetail instantiates a new InvoiceCommentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCommentDetailWithDefaults

`func NewInvoiceCommentDetailWithDefaults() *InvoiceCommentDetail`

NewInvoiceCommentDetailWithDefaults instantiates a new InvoiceCommentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceCommentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceCommentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceCommentDetail) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *InvoiceCommentDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceCommentDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceCommentDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetUserId

`func (o *InvoiceCommentDetail) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InvoiceCommentDetail) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InvoiceCommentDetail) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *InvoiceCommentDetail) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InvoiceCommentDetail) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetContent

`func (o *InvoiceCommentDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InvoiceCommentDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InvoiceCommentDetail) SetContent(v string)`

SetContent sets Content field to given value.


### GetIsInternal

`func (o *InvoiceCommentDetail) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *InvoiceCommentDetail) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *InvoiceCommentDetail) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.


### GetCreatedAt

`func (o *InvoiceCommentDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvoiceCommentDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvoiceCommentDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InvoiceCommentDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvoiceCommentDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvoiceCommentDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *InvoiceCommentDetail) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *InvoiceCommentDetail) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


