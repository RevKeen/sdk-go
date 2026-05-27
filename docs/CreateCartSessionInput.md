# CreateCartSessionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** |  | 
**Mode** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**LineItems** | Pointer to [**[]CartLineItemInput**](CartLineItemInput.md) |  | [optional] 
**AddOnsOffered** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateCartSessionInput

`func NewCreateCartSessionInput(currency string, ) *CreateCartSessionInput`

NewCreateCartSessionInput instantiates a new CreateCartSessionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCartSessionInputWithDefaults

`func NewCreateCartSessionInputWithDefaults() *CreateCartSessionInput`

NewCreateCartSessionInputWithDefaults instantiates a new CreateCartSessionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CreateCartSessionInput) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCartSessionInput) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCartSessionInput) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMode

`func (o *CreateCartSessionInput) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CreateCartSessionInput) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CreateCartSessionInput) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CreateCartSessionInput) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomerId

`func (o *CreateCartSessionInput) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCartSessionInput) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCartSessionInput) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *CreateCartSessionInput) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetLineItems

`func (o *CreateCartSessionInput) GetLineItems() []CartLineItemInput`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CreateCartSessionInput) GetLineItemsOk() (*[]CartLineItemInput, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CreateCartSessionInput) SetLineItems(v []CartLineItemInput)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *CreateCartSessionInput) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetAddOnsOffered

`func (o *CreateCartSessionInput) GetAddOnsOffered() []string`

GetAddOnsOffered returns the AddOnsOffered field if non-nil, zero value otherwise.

### GetAddOnsOfferedOk

`func (o *CreateCartSessionInput) GetAddOnsOfferedOk() (*[]string, bool)`

GetAddOnsOfferedOk returns a tuple with the AddOnsOffered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOnsOffered

`func (o *CreateCartSessionInput) SetAddOnsOffered(v []string)`

SetAddOnsOffered sets AddOnsOffered field to given value.

### HasAddOnsOffered

`func (o *CreateCartSessionInput) HasAddOnsOffered() bool`

HasAddOnsOffered returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCartSessionInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCartSessionInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCartSessionInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCartSessionInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


