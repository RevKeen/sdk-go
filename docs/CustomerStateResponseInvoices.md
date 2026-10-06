# CustomerStateResponseInvoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overdue** | **[]map[string]*interface{}** |  |
**Recent** | **[]map[string]*interface{}** |  |

## Methods

### NewCustomerStateResponseInvoices

`func NewCustomerStateResponseInvoices(overdue []map[string]*interface{}, recent []map[string]*interface{}, ) *CustomerStateResponseInvoices`

NewCustomerStateResponseInvoices instantiates a new CustomerStateResponseInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerStateResponseInvoicesWithDefaults

`func NewCustomerStateResponseInvoicesWithDefaults() *CustomerStateResponseInvoices`

NewCustomerStateResponseInvoicesWithDefaults instantiates a new CustomerStateResponseInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverdue

`func (o *CustomerStateResponseInvoices) GetOverdue() []map[string]*interface{}`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *CustomerStateResponseInvoices) GetOverdueOk() (*[]map[string]*interface{}, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *CustomerStateResponseInvoices) SetOverdue(v []map[string]*interface{})`

SetOverdue sets Overdue field to given value.


### GetRecent

`func (o *CustomerStateResponseInvoices) GetRecent() []map[string]*interface{}`

GetRecent returns the Recent field if non-nil, zero value otherwise.

### GetRecentOk

`func (o *CustomerStateResponseInvoices) GetRecentOk() (*[]map[string]*interface{}, bool)`

GetRecentOk returns a tuple with the Recent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecent

`func (o *CustomerStateResponseInvoices) SetRecent(v []map[string]*interface{})`

SetRecent sets Recent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
