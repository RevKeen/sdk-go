# PortalMandate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | **string** |  | 
**Status** | **string** |  | 
**AccountHolderName** | **NullableString** |  | 
**AccountNumberLast4** | **NullableString** |  | 
**BankName** | **NullableString** |  | 
**MandateReference** | **NullableString** |  | 
**NextCollectionDate** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 

## Methods

### NewPortalMandate

`func NewPortalMandate(id string, object string, status string, accountHolderName NullableString, accountNumberLast4 NullableString, bankName NullableString, mandateReference NullableString, nextCollectionDate NullableString, createdAt NullableTime, ) *PortalMandate`

NewPortalMandate instantiates a new PortalMandate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalMandateWithDefaults

`func NewPortalMandateWithDefaults() *PortalMandate`

NewPortalMandateWithDefaults instantiates a new PortalMandate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalMandate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalMandate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalMandate) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *PortalMandate) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PortalMandate) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PortalMandate) SetObject(v string)`

SetObject sets Object field to given value.


### GetStatus

`func (o *PortalMandate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortalMandate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortalMandate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountHolderName

`func (o *PortalMandate) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *PortalMandate) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *PortalMandate) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### SetAccountHolderNameNil

`func (o *PortalMandate) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *PortalMandate) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetAccountNumberLast4

`func (o *PortalMandate) GetAccountNumberLast4() string`

GetAccountNumberLast4 returns the AccountNumberLast4 field if non-nil, zero value otherwise.

### GetAccountNumberLast4Ok

`func (o *PortalMandate) GetAccountNumberLast4Ok() (*string, bool)`

GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLast4

`func (o *PortalMandate) SetAccountNumberLast4(v string)`

SetAccountNumberLast4 sets AccountNumberLast4 field to given value.


### SetAccountNumberLast4Nil

`func (o *PortalMandate) SetAccountNumberLast4Nil(b bool)`

 SetAccountNumberLast4Nil sets the value for AccountNumberLast4 to be an explicit nil

### UnsetAccountNumberLast4
`func (o *PortalMandate) UnsetAccountNumberLast4()`

UnsetAccountNumberLast4 ensures that no value is present for AccountNumberLast4, not even an explicit nil
### GetBankName

`func (o *PortalMandate) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PortalMandate) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PortalMandate) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *PortalMandate) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *PortalMandate) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetMandateReference

`func (o *PortalMandate) GetMandateReference() string`

GetMandateReference returns the MandateReference field if non-nil, zero value otherwise.

### GetMandateReferenceOk

`func (o *PortalMandate) GetMandateReferenceOk() (*string, bool)`

GetMandateReferenceOk returns a tuple with the MandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateReference

`func (o *PortalMandate) SetMandateReference(v string)`

SetMandateReference sets MandateReference field to given value.


### SetMandateReferenceNil

`func (o *PortalMandate) SetMandateReferenceNil(b bool)`

 SetMandateReferenceNil sets the value for MandateReference to be an explicit nil

### UnsetMandateReference
`func (o *PortalMandate) UnsetMandateReference()`

UnsetMandateReference ensures that no value is present for MandateReference, not even an explicit nil
### GetNextCollectionDate

`func (o *PortalMandate) GetNextCollectionDate() string`

GetNextCollectionDate returns the NextCollectionDate field if non-nil, zero value otherwise.

### GetNextCollectionDateOk

`func (o *PortalMandate) GetNextCollectionDateOk() (*string, bool)`

GetNextCollectionDateOk returns a tuple with the NextCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCollectionDate

`func (o *PortalMandate) SetNextCollectionDate(v string)`

SetNextCollectionDate sets NextCollectionDate field to given value.


### SetNextCollectionDateNil

`func (o *PortalMandate) SetNextCollectionDateNil(b bool)`

 SetNextCollectionDateNil sets the value for NextCollectionDate to be an explicit nil

### UnsetNextCollectionDate
`func (o *PortalMandate) UnsetNextCollectionDate()`

UnsetNextCollectionDate ensures that no value is present for NextCollectionDate, not even an explicit nil
### GetCreatedAt

`func (o *PortalMandate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortalMandate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortalMandate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PortalMandate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PortalMandate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


