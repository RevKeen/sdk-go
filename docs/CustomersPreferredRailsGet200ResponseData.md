# CustomersPreferredRailsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**InvoiceId** | **NullableString** |  | 
**AmountMinor** | **NullableInt32** |  | 
**Currency** | **NullableString** |  | 
**PreferredRails** | [**[]CustomersPreferredRailsGet200ResponseDataPreferredRailsInner**](CustomersPreferredRailsGet200ResponseDataPreferredRailsInner.md) |  | 

## Methods

### NewCustomersPreferredRailsGet200ResponseData

`func NewCustomersPreferredRailsGet200ResponseData(customerId string, invoiceId NullableString, amountMinor NullableInt32, currency NullableString, preferredRails []CustomersPreferredRailsGet200ResponseDataPreferredRailsInner, ) *CustomersPreferredRailsGet200ResponseData`

NewCustomersPreferredRailsGet200ResponseData instantiates a new CustomersPreferredRailsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomersPreferredRailsGet200ResponseDataWithDefaults

`func NewCustomersPreferredRailsGet200ResponseDataWithDefaults() *CustomersPreferredRailsGet200ResponseData`

NewCustomersPreferredRailsGet200ResponseDataWithDefaults instantiates a new CustomersPreferredRailsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CustomersPreferredRailsGet200ResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CustomersPreferredRailsGet200ResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CustomersPreferredRailsGet200ResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetInvoiceId

`func (o *CustomersPreferredRailsGet200ResponseData) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CustomersPreferredRailsGet200ResponseData) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CustomersPreferredRailsGet200ResponseData) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *CustomersPreferredRailsGet200ResponseData) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *CustomersPreferredRailsGet200ResponseData) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetAmountMinor

`func (o *CustomersPreferredRailsGet200ResponseData) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CustomersPreferredRailsGet200ResponseData) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CustomersPreferredRailsGet200ResponseData) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *CustomersPreferredRailsGet200ResponseData) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *CustomersPreferredRailsGet200ResponseData) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *CustomersPreferredRailsGet200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CustomersPreferredRailsGet200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CustomersPreferredRailsGet200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CustomersPreferredRailsGet200ResponseData) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CustomersPreferredRailsGet200ResponseData) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPreferredRails

`func (o *CustomersPreferredRailsGet200ResponseData) GetPreferredRails() []CustomersPreferredRailsGet200ResponseDataPreferredRailsInner`

GetPreferredRails returns the PreferredRails field if non-nil, zero value otherwise.

### GetPreferredRailsOk

`func (o *CustomersPreferredRailsGet200ResponseData) GetPreferredRailsOk() (*[]CustomersPreferredRailsGet200ResponseDataPreferredRailsInner, bool)`

GetPreferredRailsOk returns a tuple with the PreferredRails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredRails

`func (o *CustomersPreferredRailsGet200ResponseData) SetPreferredRails(v []CustomersPreferredRailsGet200ResponseDataPreferredRailsInner)`

SetPreferredRails sets PreferredRails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


