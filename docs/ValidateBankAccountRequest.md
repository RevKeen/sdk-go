# ValidateBankAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortCode** | **string** | UK sort code (6 digits, dashes optional) | 
**AccountNumber** | **string** | UK account number (8 digits) | 

## Methods

### NewValidateBankAccountRequest

`func NewValidateBankAccountRequest(sortCode string, accountNumber string, ) *ValidateBankAccountRequest`

NewValidateBankAccountRequest instantiates a new ValidateBankAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateBankAccountRequestWithDefaults

`func NewValidateBankAccountRequestWithDefaults() *ValidateBankAccountRequest`

NewValidateBankAccountRequestWithDefaults instantiates a new ValidateBankAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortCode

`func (o *ValidateBankAccountRequest) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *ValidateBankAccountRequest) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *ValidateBankAccountRequest) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetAccountNumber

`func (o *ValidateBankAccountRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *ValidateBankAccountRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *ValidateBankAccountRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


