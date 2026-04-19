# CreateCustomerPortalSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | UUID of the customer this session will authenticate as. | 
**TtlMinutes** | Pointer to **int32** | Lifetime of the session token in minutes. Between 5 and 240; default 60. Shorter TTLs reduce blast radius if a token leaks. | [optional] 

## Methods

### NewCreateCustomerPortalSessionRequest

`func NewCreateCustomerPortalSessionRequest(customerId string, ) *CreateCustomerPortalSessionRequest`

NewCreateCustomerPortalSessionRequest instantiates a new CreateCustomerPortalSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerPortalSessionRequestWithDefaults

`func NewCreateCustomerPortalSessionRequestWithDefaults() *CreateCustomerPortalSessionRequest`

NewCreateCustomerPortalSessionRequestWithDefaults instantiates a new CreateCustomerPortalSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateCustomerPortalSessionRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateCustomerPortalSessionRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateCustomerPortalSessionRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetTtlMinutes

`func (o *CreateCustomerPortalSessionRequest) GetTtlMinutes() int32`

GetTtlMinutes returns the TtlMinutes field if non-nil, zero value otherwise.

### GetTtlMinutesOk

`func (o *CreateCustomerPortalSessionRequest) GetTtlMinutesOk() (*int32, bool)`

GetTtlMinutesOk returns a tuple with the TtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlMinutes

`func (o *CreateCustomerPortalSessionRequest) SetTtlMinutes(v int32)`

SetTtlMinutes sets TtlMinutes field to given value.

### HasTtlMinutes

`func (o *CreateCustomerPortalSessionRequest) HasTtlMinutes() bool`

HasTtlMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


