# BankAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | **NullableString** | The name of the bank | 
**Last4** | **NullableString** | The last 4 digits of the account number | 
**RoutingNumberLast4** | **NullableString** | The last 4 digits of the routing number | 
**AccountType** | **NullableString** | The type of bank account | 

## Methods

### NewBankAccountDetails

`func NewBankAccountDetails(bankName NullableString, last4 NullableString, routingNumberLast4 NullableString, accountType NullableString, ) *BankAccountDetails`

NewBankAccountDetails instantiates a new BankAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountDetailsWithDefaults

`func NewBankAccountDetailsWithDefaults() *BankAccountDetails`

NewBankAccountDetailsWithDefaults instantiates a new BankAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *BankAccountDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BankAccountDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BankAccountDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *BankAccountDetails) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *BankAccountDetails) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetLast4

`func (o *BankAccountDetails) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *BankAccountDetails) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *BankAccountDetails) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### SetLast4Nil

`func (o *BankAccountDetails) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *BankAccountDetails) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetRoutingNumberLast4

`func (o *BankAccountDetails) GetRoutingNumberLast4() string`

GetRoutingNumberLast4 returns the RoutingNumberLast4 field if non-nil, zero value otherwise.

### GetRoutingNumberLast4Ok

`func (o *BankAccountDetails) GetRoutingNumberLast4Ok() (*string, bool)`

GetRoutingNumberLast4Ok returns a tuple with the RoutingNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumberLast4

`func (o *BankAccountDetails) SetRoutingNumberLast4(v string)`

SetRoutingNumberLast4 sets RoutingNumberLast4 field to given value.


### SetRoutingNumberLast4Nil

`func (o *BankAccountDetails) SetRoutingNumberLast4Nil(b bool)`

 SetRoutingNumberLast4Nil sets the value for RoutingNumberLast4 to be an explicit nil

### UnsetRoutingNumberLast4
`func (o *BankAccountDetails) UnsetRoutingNumberLast4()`

UnsetRoutingNumberLast4 ensures that no value is present for RoutingNumberLast4, not even an explicit nil
### GetAccountType

`func (o *BankAccountDetails) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankAccountDetails) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankAccountDetails) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### SetAccountTypeNil

`func (o *BankAccountDetails) SetAccountTypeNil(b bool)`

 SetAccountTypeNil sets the value for AccountType to be an explicit nil

### UnsetAccountType
`func (o *BankAccountDetails) UnsetAccountType()`

UnsetAccountType ensures that no value is present for AccountType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


