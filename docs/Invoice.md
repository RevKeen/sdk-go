# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique invoice identifier | 
**CustomerUuid** | **string** | ID of the customer this invoice belongs to | 
**InvoiceNumber** | Pointer to **string** | Merchant-assigned invoice number for reference | [optional] 
**TotalMinor** | **int32** | Total amount in minor units (cents) | 
**Currency** | **string** | Three-letter ISO 4217 currency code | 
**AllowedMethods** | Pointer to **[]string** | Allowed payment methods for checkout sessions against this invoice | [optional] 
**Status** | **string** | Invoice lifecycle status: draft, open, finalized, paid, void, or uncollectible | 
**DueDate** | Pointer to **NullableString** | Payment due date in ISO 8601 format | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Key-value pairs for custom fields | [optional] 
**ExternalSource** | Pointer to **NullableString** | External integration source (e.g., practicehub, wodify) | [optional] 
**ExternalType** | Pointer to **NullableString** | Type within external system (e.g., appointment, membership) | [optional] 
**ExternalId** | Pointer to **NullableString** | ID from external system | [optional] 
**SubscriptionTerms** | Pointer to [**InvoiceSubscriptionTerms**](InvoiceSubscriptionTerms.md) |  | [optional] 
**CreatedAt** | **string** | When the invoice was created (ISO 8601) | 
**UpdatedAt** | **string** | When the invoice was last updated (ISO 8601) | 

## Methods

### NewInvoice

`func NewInvoice(id string, customerUuid string, totalMinor int32, currency string, status string, createdAt string, updatedAt string, ) *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.


### GetCustomerUuid

`func (o *Invoice) GetCustomerUuid() string`

GetCustomerUuid returns the CustomerUuid field if non-nil, zero value otherwise.

### GetCustomerUuidOk

`func (o *Invoice) GetCustomerUuidOk() (*string, bool)`

GetCustomerUuidOk returns a tuple with the CustomerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUuid

`func (o *Invoice) SetCustomerUuid(v string)`

SetCustomerUuid sets CustomerUuid field to given value.


### GetInvoiceNumber

`func (o *Invoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Invoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Invoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *Invoice) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTotalMinor

`func (o *Invoice) GetTotalMinor() int32`

GetTotalMinor returns the TotalMinor field if non-nil, zero value otherwise.

### GetTotalMinorOk

`func (o *Invoice) GetTotalMinorOk() (*int32, bool)`

GetTotalMinorOk returns a tuple with the TotalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMinor

`func (o *Invoice) SetTotalMinor(v int32)`

SetTotalMinor sets TotalMinor field to given value.


### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAllowedMethods

`func (o *Invoice) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *Invoice) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *Invoice) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *Invoice) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.

### SetAllowedMethodsNil

`func (o *Invoice) SetAllowedMethodsNil(b bool)`

 SetAllowedMethodsNil sets the value for AllowedMethods to be an explicit nil

### UnsetAllowedMethods
`func (o *Invoice) UnsetAllowedMethods()`

UnsetAllowedMethods ensures that no value is present for AllowedMethods, not even an explicit nil
### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDueDate

`func (o *Invoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *Invoice) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Invoice) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetCustomFields

`func (o *Invoice) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Invoice) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Invoice) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Invoice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetExternalSource

`func (o *Invoice) GetExternalSource() string`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *Invoice) GetExternalSourceOk() (*string, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *Invoice) SetExternalSource(v string)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *Invoice) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *Invoice) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *Invoice) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalType

`func (o *Invoice) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *Invoice) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *Invoice) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *Invoice) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *Invoice) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *Invoice) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetExternalId

`func (o *Invoice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Invoice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Invoice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Invoice) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Invoice) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Invoice) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSubscriptionTerms

`func (o *Invoice) GetSubscriptionTerms() InvoiceSubscriptionTerms`

GetSubscriptionTerms returns the SubscriptionTerms field if non-nil, zero value otherwise.

### GetSubscriptionTermsOk

`func (o *Invoice) GetSubscriptionTermsOk() (*InvoiceSubscriptionTerms, bool)`

GetSubscriptionTermsOk returns a tuple with the SubscriptionTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTerms

`func (o *Invoice) SetSubscriptionTerms(v InvoiceSubscriptionTerms)`

SetSubscriptionTerms sets SubscriptionTerms field to given value.

### HasSubscriptionTerms

`func (o *Invoice) HasSubscriptionTerms() bool`

HasSubscriptionTerms returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Invoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Invoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Invoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Invoice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invoice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invoice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


