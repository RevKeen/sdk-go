# CreateMandateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Customer the mandate is created for | 
**AccountHolderName** | **string** | Name on the bank account | 
**SortCode** | **string** | UK sort code (6 digits, dashes optional) | 
**AccountNumber** | **string** | UK account number (8 digits) | 
**BackupPaymentMethodId** | Pointer to **string** | Stored card used as recovery fallback if a collection fails | [optional] 
**RequestToken** | Pointer to **string** | Signed mandate-request token to consume on creation | [optional] 
**Acceptance** | Pointer to [**CreateMandateRequestAcceptance**](CreateMandateRequestAcceptance.md) |  | [optional] 

## Methods

### NewCreateMandateRequest

`func NewCreateMandateRequest(customerId string, accountHolderName string, sortCode string, accountNumber string, ) *CreateMandateRequest`

NewCreateMandateRequest instantiates a new CreateMandateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMandateRequestWithDefaults

`func NewCreateMandateRequestWithDefaults() *CreateMandateRequest`

NewCreateMandateRequestWithDefaults instantiates a new CreateMandateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *CreateMandateRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *CreateMandateRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *CreateMandateRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetAccountHolderName

`func (o *CreateMandateRequest) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *CreateMandateRequest) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *CreateMandateRequest) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### GetSortCode

`func (o *CreateMandateRequest) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *CreateMandateRequest) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *CreateMandateRequest) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetAccountNumber

`func (o *CreateMandateRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateMandateRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateMandateRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBackupPaymentMethodId

`func (o *CreateMandateRequest) GetBackupPaymentMethodId() string`

GetBackupPaymentMethodId returns the BackupPaymentMethodId field if non-nil, zero value otherwise.

### GetBackupPaymentMethodIdOk

`func (o *CreateMandateRequest) GetBackupPaymentMethodIdOk() (*string, bool)`

GetBackupPaymentMethodIdOk returns a tuple with the BackupPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPaymentMethodId

`func (o *CreateMandateRequest) SetBackupPaymentMethodId(v string)`

SetBackupPaymentMethodId sets BackupPaymentMethodId field to given value.

### HasBackupPaymentMethodId

`func (o *CreateMandateRequest) HasBackupPaymentMethodId() bool`

HasBackupPaymentMethodId returns a boolean if a field has been set.

### GetRequestToken

`func (o *CreateMandateRequest) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *CreateMandateRequest) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *CreateMandateRequest) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *CreateMandateRequest) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetAcceptance

`func (o *CreateMandateRequest) GetAcceptance() CreateMandateRequestAcceptance`

GetAcceptance returns the Acceptance field if non-nil, zero value otherwise.

### GetAcceptanceOk

`func (o *CreateMandateRequest) GetAcceptanceOk() (*CreateMandateRequestAcceptance, bool)`

GetAcceptanceOk returns a tuple with the Acceptance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptance

`func (o *CreateMandateRequest) SetAcceptance(v CreateMandateRequestAcceptance)`

SetAcceptance sets Acceptance field to given value.

### HasAcceptance

`func (o *CreateMandateRequest) HasAcceptance() bool`

HasAcceptance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


