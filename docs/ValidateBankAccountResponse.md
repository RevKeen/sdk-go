# ValidateBankAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | Whether the account details passed modulus/sort-code validation | 
**BankName** | **NullableString** | Resolved bank name, if known | 
**DdEnabled** | **bool** | Whether the account accepts Direct Debit instructions | 
**Errors** | Pointer to **[]string** | Validation error messages, if any | [optional] 

## Methods

### NewValidateBankAccountResponse

`func NewValidateBankAccountResponse(valid bool, bankName NullableString, ddEnabled bool, ) *ValidateBankAccountResponse`

NewValidateBankAccountResponse instantiates a new ValidateBankAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateBankAccountResponseWithDefaults

`func NewValidateBankAccountResponseWithDefaults() *ValidateBankAccountResponse`

NewValidateBankAccountResponseWithDefaults instantiates a new ValidateBankAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidateBankAccountResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidateBankAccountResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidateBankAccountResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetBankName

`func (o *ValidateBankAccountResponse) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *ValidateBankAccountResponse) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *ValidateBankAccountResponse) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *ValidateBankAccountResponse) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *ValidateBankAccountResponse) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetDdEnabled

`func (o *ValidateBankAccountResponse) GetDdEnabled() bool`

GetDdEnabled returns the DdEnabled field if non-nil, zero value otherwise.

### GetDdEnabledOk

`func (o *ValidateBankAccountResponse) GetDdEnabledOk() (*bool, bool)`

GetDdEnabledOk returns a tuple with the DdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdEnabled

`func (o *ValidateBankAccountResponse) SetDdEnabled(v bool)`

SetDdEnabled sets DdEnabled field to given value.


### GetErrors

`func (o *ValidateBankAccountResponse) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidateBankAccountResponse) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidateBankAccountResponse) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidateBankAccountResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ValidateBankAccountResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ValidateBankAccountResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


