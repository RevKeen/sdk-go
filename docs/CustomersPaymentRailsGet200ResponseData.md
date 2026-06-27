# CustomersPaymentRailsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**InvoiceId** | Pointer to **NullableString** |  | [optional] 
**Rails** | [**[]CustomersPaymentRailsGet200ResponseDataRailsInner**](CustomersPaymentRailsGet200ResponseDataRailsInner.md) |  | 

## Methods

### NewCustomersPaymentRailsGet200ResponseData

`func NewCustomersPaymentRailsGet200ResponseData(customerId string, rails []CustomersPaymentRailsGet200ResponseDataRailsInner, ) *CustomersPaymentRailsGet200ResponseData`

NewCustomersPaymentRailsGet200ResponseData instantiates a new CustomersPaymentRailsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersPaymentRailsGet200ResponseDataWithDefaults

`func NewCustomersPaymentRailsGet200ResponseDataWithDefaults() *CustomersPaymentRailsGet200ResponseData`

NewCustomersPaymentRailsGet200ResponseDataWithDefaults instantiates a new CustomersPaymentRailsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomersPaymentRailsGet200ResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomersPaymentRailsGet200ResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomersPaymentRailsGet200ResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *CustomersPaymentRailsGet200ResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CustomersPaymentRailsGet200ResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CustomersPaymentRailsGet200ResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CustomersPaymentRailsGet200ResponseData) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### SetInvoiceIdNil

`func (o *CustomersPaymentRailsGet200ResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *CustomersPaymentRailsGet200ResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetRails

`func (o *CustomersPaymentRailsGet200ResponseData) GetRails() []CustomersPaymentRailsGet200ResponseDataRailsInner`

GetRails returns the Rails field if non-nil, zero value otherwise.

### GetRailsOk

`func (o *CustomersPaymentRailsGet200ResponseData) GetRailsOk() (*[]CustomersPaymentRailsGet200ResponseDataRailsInner, bool)`

GetRailsOk returns a tuple with the Rails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRails

`func (o *CustomersPaymentRailsGet200ResponseData) SetRails(v []CustomersPaymentRailsGet200ResponseDataRailsInner)`

SetRails sets Rails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


