# UpdatePaymentLinkStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**StatusChangedAt** | **time.Time** |  | 
**IsActive** | **bool** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUpdatePaymentLinkStatus200Response

`func NewUpdatePaymentLinkStatus200Response(id string, status string, statusChangedAt time.Time, isActive bool, updatedAt time.Time, ) *UpdatePaymentLinkStatus200Response`

NewUpdatePaymentLinkStatus200Response instantiates a new UpdatePaymentLinkStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentLinkStatus200ResponseWithDefaults

`func NewUpdatePaymentLinkStatus200ResponseWithDefaults() *UpdatePaymentLinkStatus200Response`

NewUpdatePaymentLinkStatus200ResponseWithDefaults instantiates a new UpdatePaymentLinkStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePaymentLinkStatus200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePaymentLinkStatus200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePaymentLinkStatus200Response) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *UpdatePaymentLinkStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePaymentLinkStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePaymentLinkStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusChangedAt

`func (o *UpdatePaymentLinkStatus200Response) GetStatusChangedAt() time.Time`

GetStatusChangedAt returns the StatusChangedAt field if non-nil, zero value otherwise.

### GetStatusChangedAtOk

`func (o *UpdatePaymentLinkStatus200Response) GetStatusChangedAtOk() (*time.Time, bool)`

GetStatusChangedAtOk returns a tuple with the StatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedAt

`func (o *UpdatePaymentLinkStatus200Response) SetStatusChangedAt(v time.Time)`

SetStatusChangedAt sets StatusChangedAt field to given value.


### GetIsActive

`func (o *UpdatePaymentLinkStatus200Response) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdatePaymentLinkStatus200Response) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdatePaymentLinkStatus200Response) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUpdatedAt

`func (o *UpdatePaymentLinkStatus200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatePaymentLinkStatus200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatePaymentLinkStatus200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


