# CustomerPortalSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Server-generated session identifier. Embedded in the token&#39;s &#x60;jti&#x60; claim. |
**Token** | **string** | Short-lived bearer token. The customer (or client app acting on their behalf) sends this as &#x60;Authorization: Bearer &lt;token&gt;&#x60; on every subsequent /v2/customer-portal/_* call. Treat as a credential — never log it and never persist it in a browser-accessible store. |
**CustomerId** | **string** |  |
**ExpiresAt** | **time.Time** | ISO-8601 timestamp when this token expires. Default TTL is 60 minutes. |
**CreatedAt** | **time.Time** |  |

## Methods

### NewCustomerPortalSession

`func NewCustomerPortalSession(id string, token string, customerId string, expiresAt time.Time, createdAt time.Time, ) *CustomerPortalSession`

NewCustomerPortalSession instantiates a new CustomerPortalSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPortalSessionWithDefaults

`func NewCustomerPortalSessionWithDefaults() *CustomerPortalSession`

NewCustomerPortalSessionWithDefaults instantiates a new CustomerPortalSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerPortalSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPortalSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPortalSession) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *CustomerPortalSession) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CustomerPortalSession) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CustomerPortalSession) SetToken(v string)`

SetToken sets Token field to given value.


### GetCustomerId

`func (o *CustomerPortalSession) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomerPortalSession) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomerPortalSession) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetExpiresAt

`func (o *CustomerPortalSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CustomerPortalSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CustomerPortalSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *CustomerPortalSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CustomerPortalSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CustomerPortalSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
