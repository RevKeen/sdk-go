# BacsDirectDebitDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortCode** | **NullableString** | Masked sort code — never the full value | 
**AccountNumberLast4** | **NullableString** | Last 4 digits of the account number | 
**BankName** | **NullableString** | Resolved bank name | 
**MandateId** | **string** | The Direct Debit mandate backing this payment method | 
**MandateRef** | **NullableString** | RevKeen mandate reference (RK-XXXXXX) | 
**MandateStatus** | **string** | Mandate status | 

## Methods

### NewBacsDirectDebitDetails

`func NewBacsDirectDebitDetails(sortCode NullableString, accountNumberLast4 NullableString, bankName NullableString, mandateId string, mandateRef NullableString, mandateStatus string, ) *BacsDirectDebitDetails`

NewBacsDirectDebitDetails instantiates a new BacsDirectDebitDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBacsDirectDebitDetailsWithDefaults

`func NewBacsDirectDebitDetailsWithDefaults() *BacsDirectDebitDetails`

NewBacsDirectDebitDetailsWithDefaults instantiates a new BacsDirectDebitDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortCode

`func (o *BacsDirectDebitDetails) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *BacsDirectDebitDetails) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *BacsDirectDebitDetails) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### SetSortCodeNil

`func (o *BacsDirectDebitDetails) SetSortCodeNil(b bool)`

 SetSortCodeNil sets the value for SortCode to be an explicit nil

### UnsetSortCode
`func (o *BacsDirectDebitDetails) UnsetSortCode()`

UnsetSortCode ensures that no value is present for SortCode, not even an explicit nil
### GetAccountNumberLast4

`func (o *BacsDirectDebitDetails) GetAccountNumberLast4() string`

GetAccountNumberLast4 returns the AccountNumberLast4 field if non-nil, zero value otherwise.

### GetAccountNumberLast4Ok

`func (o *BacsDirectDebitDetails) GetAccountNumberLast4Ok() (*string, bool)`

GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLast4

`func (o *BacsDirectDebitDetails) SetAccountNumberLast4(v string)`

SetAccountNumberLast4 sets AccountNumberLast4 field to given value.


### SetAccountNumberLast4Nil

`func (o *BacsDirectDebitDetails) SetAccountNumberLast4Nil(b bool)`

 SetAccountNumberLast4Nil sets the value for AccountNumberLast4 to be an explicit nil

### UnsetAccountNumberLast4
`func (o *BacsDirectDebitDetails) UnsetAccountNumberLast4()`

UnsetAccountNumberLast4 ensures that no value is present for AccountNumberLast4, not even an explicit nil
### GetBankName

`func (o *BacsDirectDebitDetails) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *BacsDirectDebitDetails) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *BacsDirectDebitDetails) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *BacsDirectDebitDetails) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *BacsDirectDebitDetails) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetMandateId

`func (o *BacsDirectDebitDetails) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *BacsDirectDebitDetails) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *BacsDirectDebitDetails) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetMandateRef

`func (o *BacsDirectDebitDetails) GetMandateRef() string`

GetMandateRef returns the MandateRef field if non-nil, zero value otherwise.

### GetMandateRefOk

`func (o *BacsDirectDebitDetails) GetMandateRefOk() (*string, bool)`

GetMandateRefOk returns a tuple with the MandateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRef

`func (o *BacsDirectDebitDetails) SetMandateRef(v string)`

SetMandateRef sets MandateRef field to given value.


### SetMandateRefNil

`func (o *BacsDirectDebitDetails) SetMandateRefNil(b bool)`

 SetMandateRefNil sets the value for MandateRef to be an explicit nil

### UnsetMandateRef
`func (o *BacsDirectDebitDetails) UnsetMandateRef()`

UnsetMandateRef ensures that no value is present for MandateRef, not even an explicit nil
### GetMandateStatus

`func (o *BacsDirectDebitDetails) GetMandateStatus() string`

GetMandateStatus returns the MandateStatus field if non-nil, zero value otherwise.

### GetMandateStatusOk

`func (o *BacsDirectDebitDetails) GetMandateStatusOk() (*string, bool)`

GetMandateStatusOk returns a tuple with the MandateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateStatus

`func (o *BacsDirectDebitDetails) SetMandateStatus(v string)`

SetMandateStatus sets MandateStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


