# CustomerStateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | **map[string]interface{}** |  |
**Subscriptions** | [**CustomerStateResponseSubscriptions**](CustomerStateResponseSubscriptions.md) |  |
**Invoices** | [**CustomerStateResponseInvoices**](CustomerStateResponseInvoices.md) |  |

## Methods

### NewCustomerStateResponse

`func NewCustomerStateResponse(customer map[string]*interface{}, subscriptions CustomerStateResponseSubscriptions, invoices CustomerStateResponseInvoices, ) *CustomerStateResponse`

NewCustomerStateResponse instantiates a new CustomerStateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerStateResponseWithDefaults

`func NewCustomerStateResponseWithDefaults() *CustomerStateResponse`

NewCustomerStateResponseWithDefaults instantiates a new CustomerStateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerStateResponse) GetCustomer() map[string]*interface{}`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerStateResponse) GetCustomerOk() (*map[string]*interface{}, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerStateResponse) SetCustomer(v map[string]*interface{})`

SetCustomer sets Customer field to given value.


### GetSubscriptions

`func (o *CustomerStateResponse) GetSubscriptions() CustomerStateResponseSubscriptions`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CustomerStateResponse) GetSubscriptionsOk() (*CustomerStateResponseSubscriptions, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CustomerStateResponse) SetSubscriptions(v CustomerStateResponseSubscriptions)`

SetSubscriptions sets Subscriptions field to given value.


### GetInvoices

`func (o *CustomerStateResponse) GetInvoices() CustomerStateResponseInvoices`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *CustomerStateResponse) GetInvoicesOk() (*CustomerStateResponseInvoices, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *CustomerStateResponse) SetInvoices(v CustomerStateResponseInvoices)`

SetInvoices sets Invoices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
