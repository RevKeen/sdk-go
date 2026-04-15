# InvoicesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerUuid** | **string** |  | 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**TotalMinor** | **int32** |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**DueDate** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Key-value pairs for custom fields | [optional] 
**SubscriptionTerms** | Pointer to [**SubscriptionTerms**](SubscriptionTerms.md) |  | [optional] 
**AllowedMethods** | Pointer to **[]string** | Restrict checkout payment methods for this invoice. When set, narrows available rails for checkout sessions created against this invoice. Omit to use merchant default. | [optional] 

## Methods

### NewInvoicesCreateRequest

`func NewInvoicesCreateRequest(customerUuid string, totalMinor int32, ) *InvoicesCreateRequest`

NewInvoicesCreateRequest instantiates a new InvoicesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesCreateRequestWithDefaults

`func NewInvoicesCreateRequestWithDefaults() *InvoicesCreateRequest`

NewInvoicesCreateRequestWithDefaults instantiates a new InvoicesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerUuid

`func (o *InvoicesCreateRequest) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *InvoicesCreateRequest) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *InvoicesCreateRequest) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.


### GetInvoiceNumber

`func (o *InvoicesCreateRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoicesCreateRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoicesCreateRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoicesCreateRequest) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTotalMinor

`func (o *InvoicesCreateRequest) GetTotalMinor() int32`

GetTotalMinor returns the TotalMinor field if non-nil, zero value otherwise.

### GetTotalMinorOk

`func (o *InvoicesCreateRequest) GetTotalMinorOk() (*int32, bool)`

GetTotalMinorOk returns a tuple with the TotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinor

`func (o *InvoicesCreateRequest) SetTotalMinor(v int32)`

SetTotalMinor sets TotalMinor field to given value.


### GetCurrency

`func (o *InvoicesCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoicesCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoicesCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoicesCreateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoicesCreateRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoicesCreateRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoicesCreateRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoicesCreateRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetStatus

`func (o *InvoicesCreateRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoicesCreateRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoicesCreateRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoicesCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *InvoicesCreateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InvoicesCreateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InvoicesCreateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InvoicesCreateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetSubscriptionTerms

`func (o *InvoicesCreateRequest) GetSubscriptionTerms() SubscriptionTerms`

GetSubscriptionTerms returns the SubscriptionTerms field if non-nil, zero value otherwise.

### GetSubscriptionTermsOk

`func (o *InvoicesCreateRequest) GetSubscriptionTermsOk() (*SubscriptionTerms, bool)`

GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTerms

`func (o *InvoicesCreateRequest) SetSubscriptionTerms(v SubscriptionTerms)`

SetSubscriptionTerms sets SubscriptionTerms field to given value.

### HasSubscriptionTerms

`func (o *InvoicesCreateRequest) HasSubscriptionTerms() bool`

HasSubscriptionTerms returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *InvoicesCreateRequest) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *InvoicesCreateRequest) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *InvoicesCreateRequest) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *InvoicesCreateRequest) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


