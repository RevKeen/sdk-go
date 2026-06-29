# UpdatePaymentMethodRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingDetails** | Pointer to [**BillingDetails**](BillingDetails.md) | Updated billing details | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata to attach | [optional] 

## Methods

### NewUpdatePaymentMethodRequest

`func NewUpdatePaymentMethodRequest() *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequest instantiates a new UpdatePaymentMethodRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePaymentMethodRequestWithDefaults

`func NewUpdatePaymentMethodRequestWithDefaults() *UpdatePaymentMethodRequest`

NewUpdatePaymentMethodRequestWithDefaults instantiates a new UpdatePaymentMethodRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingDetails

`func (o *UpdatePaymentMethodRequest) GetBillingDetails() BillingDetails`

GetBillingDetails returns the BillingDetails field if non-nil, zero value otherwise.

### GetBillingDetailsOk

`func (o *UpdatePaymentMethodRequest) GetBillingDetailsOk() (*BillingDetails, bool)`

GetBillingDetailsOk returns a tuple with the BillingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDetails

`func (o *UpdatePaymentMethodRequest) SetBillingDetails(v BillingDetails)`

SetBillingDetails sets BillingDetails field to given value.

### HasBillingDetails

`func (o *UpdatePaymentMethodRequest) HasBillingDetails() bool`

HasBillingDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePaymentMethodRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePaymentMethodRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePaymentMethodRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePaymentMethodRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


