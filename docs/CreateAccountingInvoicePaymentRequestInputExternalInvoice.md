# CreateAccountingInvoicePaymentRequestInputExternalInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Number** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**IssuedAt** | Pointer to **NullableTime** |  | [optional] 
**DueAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCreateAccountingInvoicePaymentRequestInputExternalInvoice

`func NewCreateAccountingInvoicePaymentRequestInputExternalInvoice(id string, ) *CreateAccountingInvoicePaymentRequestInputExternalInvoice`

NewCreateAccountingInvoicePaymentRequestInputExternalInvoice instantiates a new CreateAccountingInvoicePaymentRequestInputExternalInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountingInvoicePaymentRequestInputExternalInvoiceWithDefaults

`func NewCreateAccountingInvoicePaymentRequestInputExternalInvoiceWithDefaults() *CreateAccountingInvoicePaymentRequestInputExternalInvoice`

NewCreateAccountingInvoicePaymentRequestInputExternalInvoiceWithDefaults instantiates a new CreateAccountingInvoicePaymentRequestInputExternalInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetStatus

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIssuedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### SetIssuedAtNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetIssuedAtNil(b bool)`

 SetIssuedAtNil sets the value for IssuedAt to be an explicit nil

### UnsetIssuedAt
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetIssuedAt()`

UnsetIssuedAt ensures that no value is present for IssuedAt, not even an explicit nil
### GetDueAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetDueAt() time.Time`

GetDueAt returns the DueAt field if non-nil, zero value otherwise.

### GetDueAtOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetDueAtOk() (*time.Time, bool)`

GetDueAtOk returns a tuple with the DueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetDueAt(v time.Time)`

SetDueAt sets DueAt field to given value.

### HasDueAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasDueAt() bool`

HasDueAt returns a boolean if a field has been set.

### SetDueAtNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetDueAtNil(b bool)`

 SetDueAtNil sets the value for DueAt to be an explicit nil

### UnsetDueAt
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetDueAt()`

UnsetDueAt ensures that no value is present for DueAt, not even an explicit nil
### GetUpdatedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CreateAccountingInvoicePaymentRequestInputExternalInvoice) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


