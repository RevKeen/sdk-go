# MandateListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MandateRef** | **string** |  | 
**MandateReference** | **NullableString** |  | 
**CustomerId** | **string** |  | 
**Status** | **string** |  | 
**AccountHolderName** | **NullableString** |  | 
**SortCode** | **string** |  | 
**AccountNumberLast4** | **NullableString** |  | 
**BankName** | **NullableString** |  | 
**NoticeDays** | **NullableInt32** |  | 
**ActivatedAt** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 

## Methods

### NewMandateListItem

`func NewMandateListItem(id string, mandateRef string, mandateReference NullableString, customerId string, status string, accountHolderName NullableString, sortCode string, accountNumberLast4 NullableString, bankName NullableString, noticeDays NullableInt32, activatedAt NullableString, createdAt NullableString, ) *MandateListItem`

NewMandateListItem instantiates a new MandateListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandateListItemWithDefaults

`func NewMandateListItemWithDefaults() *MandateListItem`

NewMandateListItemWithDefaults instantiates a new MandateListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MandateListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MandateListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MandateListItem) SetId(v string)`

SetId sets Id field to given value.


### GetMandateRef

`func (o *MandateListItem) GetMandateRef() string`

GetMandateRef returns the MandateRef field if non-nil, zero value otherwise.

### GetMandateRefOk

`func (o *MandateListItem) GetMandateRefOk() (*string, bool)`

GetMandateRefOk returns a tuple with the MandateRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRef

`func (o *MandateListItem) SetMandateRef(v string)`

SetMandateRef sets MandateRef field to given value.


### GetMandateReference

`func (o *MandateListItem) GetMandateReference() string`

GetMandateReference returns the MandateReference field if non-nil, zero value otherwise.

### GetMandateReferenceOk

`func (o *MandateListItem) GetMandateReferenceOk() (*string, bool)`

GetMandateReferenceOk returns a tuple with the MandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateReference

`func (o *MandateListItem) SetMandateReference(v string)`

SetMandateReference sets MandateReference field to given value.


### SetMandateReferenceNil

`func (o *MandateListItem) SetMandateReferenceNil(b bool)`

 SetMandateReferenceNil sets the value for MandateReference to be an explicit nil

### UnsetMandateReference
`func (o *MandateListItem) UnsetMandateReference()`

UnsetMandateReference ensures that no value is present for MandateReference, not even an explicit nil
### GetCustomerId

`func (o *MandateListItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MandateListItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MandateListItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetStatus

`func (o *MandateListItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MandateListItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MandateListItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountHolderName

`func (o *MandateListItem) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *MandateListItem) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *MandateListItem) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### SetAccountHolderNameNil

`func (o *MandateListItem) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *MandateListItem) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSortCode

`func (o *MandateListItem) GetSortCode() string`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *MandateListItem) GetSortCodeOk() (*string, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *MandateListItem) SetSortCode(v string)`

SetSortCode sets SortCode field to given value.


### GetAccountNumberLast4

`func (o *MandateListItem) GetAccountNumberLast4() string`

GetAccountNumberLast4 returns the AccountNumberLast4 field if non-nil, zero value otherwise.

### GetAccountNumberLast4Ok

`func (o *MandateListItem) GetAccountNumberLast4Ok() (*string, bool)`

GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLast4

`func (o *MandateListItem) SetAccountNumberLast4(v string)`

SetAccountNumberLast4 sets AccountNumberLast4 field to given value.


### SetAccountNumberLast4Nil

`func (o *MandateListItem) SetAccountNumberLast4Nil(b bool)`

 SetAccountNumberLast4Nil sets the value for AccountNumberLast4 to be an explicit nil

### UnsetAccountNumberLast4
`func (o *MandateListItem) UnsetAccountNumberLast4()`

UnsetAccountNumberLast4 ensures that no value is present for AccountNumberLast4, not even an explicit nil
### GetBankName

`func (o *MandateListItem) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *MandateListItem) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *MandateListItem) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *MandateListItem) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *MandateListItem) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetNoticeDays

`func (o *MandateListItem) GetNoticeDays() int32`

GetNoticeDays returns the NoticeDays field if non-nil, zero value otherwise.

### GetNoticeDaysOk

`func (o *MandateListItem) GetNoticeDaysOk() (*int32, bool)`

GetNoticeDaysOk returns a tuple with the NoticeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDays

`func (o *MandateListItem) SetNoticeDays(v int32)`

SetNoticeDays sets NoticeDays field to given value.


### SetNoticeDaysNil

`func (o *MandateListItem) SetNoticeDaysNil(b bool)`

 SetNoticeDaysNil sets the value for NoticeDays to be an explicit nil

### UnsetNoticeDays
`func (o *MandateListItem) UnsetNoticeDays()`

UnsetNoticeDays ensures that no value is present for NoticeDays, not even an explicit nil
### GetActivatedAt

`func (o *MandateListItem) GetActivatedAt() string`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *MandateListItem) GetActivatedAtOk() (*string, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *MandateListItem) SetActivatedAt(v string)`

SetActivatedAt sets ActivatedAt field to given value.


### SetActivatedAtNil

`func (o *MandateListItem) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *MandateListItem) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetCreatedAt

`func (o *MandateListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MandateListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MandateListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *MandateListItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MandateListItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


