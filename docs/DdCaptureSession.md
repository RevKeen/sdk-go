# DdCaptureSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionToken** | Pointer to **string** | Plaintext capture-session token. Only present in the response that issues it (create / token re-issue); never returned on subsequent reads. | [optional] 
**MerchantId** | **string** |  | 
**CustomerId** | **NullableString** |  | 
**CheckoutSessionId** | **NullableString** |  | 
**MandateRequestId** | **NullableString** |  | 
**Source** | **string** |  | 
**Status** | **string** |  | 
**AccountHolderName** | **NullableString** |  | 
**SortCodeLast2** | **NullableString** | Last 2 digits of the sort code — never the full value | 
**AccountNumberLast4** | **NullableString** | Last 4 digits of the account number | 
**BankName** | **NullableString** |  | 
**NoticeDate** | **NullableString** | Advance-notice date (YYYY-MM-DD) | 
**SubmissionDate** | **NullableString** | Bacs submission date (YYYY-MM-DD) | 
**CollectionDate** | **NullableString** | Collection date (YYYY-MM-DD) | 
**SettlementDate** | **NullableString** | Settlement date (YYYY-MM-DD) | 
**NoticeDays** | **NullableInt32** |  | 
**MandateId** | **NullableString** | Mandate created from this capture session, once completed | 
**ExpiresAt** | **NullableString** |  | 
**CompletedAt** | **NullableString** |  | 
**AbandonedAt** | **NullableString** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewDdCaptureSession

`func NewDdCaptureSession(id string, merchantId string, customerId NullableString, checkoutSessionId NullableString, mandateRequestId NullableString, source string, status string, accountHolderName NullableString, sortCodeLast2 NullableString, accountNumberLast4 NullableString, bankName NullableString, noticeDate NullableString, submissionDate NullableString, collectionDate NullableString, settlementDate NullableString, noticeDays NullableInt32, mandateId NullableString, expiresAt NullableString, completedAt NullableString, abandonedAt NullableString, createdAt NullableString, updatedAt NullableString, ) *DdCaptureSession`

NewDdCaptureSession instantiates a new DdCaptureSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdCaptureSessionWithDefaults

`func NewDdCaptureSessionWithDefaults() *DdCaptureSession`

NewDdCaptureSessionWithDefaults instantiates a new DdCaptureSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DdCaptureSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DdCaptureSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DdCaptureSession) SetId(v string)`

SetId sets Id field to given value.


### GetSessionToken

`func (o *DdCaptureSession) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *DdCaptureSession) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *DdCaptureSession) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *DdCaptureSession) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetMerchantId

`func (o *DdCaptureSession) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DdCaptureSession) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DdCaptureSession) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *DdCaptureSession) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DdCaptureSession) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DdCaptureSession) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### SetCustomerIdNil

`func (o *DdCaptureSession) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *DdCaptureSession) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCheckoutSessionId

`func (o *DdCaptureSession) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *DdCaptureSession) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *DdCaptureSession) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.


### SetCheckoutSessionIdNil

`func (o *DdCaptureSession) SetCheckoutSessionIdNil(b bool)`

 SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil

### UnsetCheckoutSessionId
`func (o *DdCaptureSession) UnsetCheckoutSessionId()`

UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
### GetMandateRequestId

`func (o *DdCaptureSession) GetMandateRequestId() string`

GetMandateRequestId returns the MandateRequestId field if non-nil, zero value otherwise.

### GetMandateRequestIdOk

`func (o *DdCaptureSession) GetMandateRequestIdOk() (*string, bool)`

GetMandateRequestIdOk returns a tuple with the MandateRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateRequestId

`func (o *DdCaptureSession) SetMandateRequestId(v string)`

SetMandateRequestId sets MandateRequestId field to given value.


### SetMandateRequestIdNil

`func (o *DdCaptureSession) SetMandateRequestIdNil(b bool)`

 SetMandateRequestIdNil sets the value for MandateRequestId to be an explicit nil

### UnsetMandateRequestId
`func (o *DdCaptureSession) UnsetMandateRequestId()`

UnsetMandateRequestId ensures that no value is present for MandateRequestId, not even an explicit nil
### GetSource

`func (o *DdCaptureSession) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DdCaptureSession) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DdCaptureSession) SetSource(v string)`

SetSource sets Source field to given value.


### GetStatus

`func (o *DdCaptureSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DdCaptureSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DdCaptureSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountHolderName

`func (o *DdCaptureSession) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *DdCaptureSession) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *DdCaptureSession) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### SetAccountHolderNameNil

`func (o *DdCaptureSession) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *DdCaptureSession) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetSortCodeLast2

`func (o *DdCaptureSession) GetSortCodeLast2() string`

GetSortCodeLast2 returns the SortCodeLast2 field if non-nil, zero value otherwise.

### GetSortCodeLast2Ok

`func (o *DdCaptureSession) GetSortCodeLast2Ok() (*string, bool)`

GetSortCodeLast2Ok returns a tuple with the SortCodeLast2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCodeLast2

`func (o *DdCaptureSession) SetSortCodeLast2(v string)`

SetSortCodeLast2 sets SortCodeLast2 field to given value.


### SetSortCodeLast2Nil

`func (o *DdCaptureSession) SetSortCodeLast2Nil(b bool)`

 SetSortCodeLast2Nil sets the value for SortCodeLast2 to be an explicit nil

### UnsetSortCodeLast2
`func (o *DdCaptureSession) UnsetSortCodeLast2()`

UnsetSortCodeLast2 ensures that no value is present for SortCodeLast2, not even an explicit nil
### GetAccountNumberLast4

`func (o *DdCaptureSession) GetAccountNumberLast4() string`

GetAccountNumberLast4 returns the AccountNumberLast4 field if non-nil, zero value otherwise.

### GetAccountNumberLast4Ok

`func (o *DdCaptureSession) GetAccountNumberLast4Ok() (*string, bool)`

GetAccountNumberLast4Ok returns a tuple with the AccountNumberLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumberLast4

`func (o *DdCaptureSession) SetAccountNumberLast4(v string)`

SetAccountNumberLast4 sets AccountNumberLast4 field to given value.


### SetAccountNumberLast4Nil

`func (o *DdCaptureSession) SetAccountNumberLast4Nil(b bool)`

 SetAccountNumberLast4Nil sets the value for AccountNumberLast4 to be an explicit nil

### UnsetAccountNumberLast4
`func (o *DdCaptureSession) UnsetAccountNumberLast4()`

UnsetAccountNumberLast4 ensures that no value is present for AccountNumberLast4, not even an explicit nil
### GetBankName

`func (o *DdCaptureSession) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *DdCaptureSession) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *DdCaptureSession) SetBankName(v string)`

SetBankName sets BankName field to given value.


### SetBankNameNil

`func (o *DdCaptureSession) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *DdCaptureSession) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetNoticeDate

`func (o *DdCaptureSession) GetNoticeDate() string`

GetNoticeDate returns the NoticeDate field if non-nil, zero value otherwise.

### GetNoticeDateOk

`func (o *DdCaptureSession) GetNoticeDateOk() (*string, bool)`

GetNoticeDateOk returns a tuple with the NoticeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDate

`func (o *DdCaptureSession) SetNoticeDate(v string)`

SetNoticeDate sets NoticeDate field to given value.


### SetNoticeDateNil

`func (o *DdCaptureSession) SetNoticeDateNil(b bool)`

 SetNoticeDateNil sets the value for NoticeDate to be an explicit nil

### UnsetNoticeDate
`func (o *DdCaptureSession) UnsetNoticeDate()`

UnsetNoticeDate ensures that no value is present for NoticeDate, not even an explicit nil
### GetSubmissionDate

`func (o *DdCaptureSession) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *DdCaptureSession) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *DdCaptureSession) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.


### SetSubmissionDateNil

`func (o *DdCaptureSession) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *DdCaptureSession) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil
### GetCollectionDate

`func (o *DdCaptureSession) GetCollectionDate() string`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *DdCaptureSession) GetCollectionDateOk() (*string, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *DdCaptureSession) SetCollectionDate(v string)`

SetCollectionDate sets CollectionDate field to given value.


### SetCollectionDateNil

`func (o *DdCaptureSession) SetCollectionDateNil(b bool)`

 SetCollectionDateNil sets the value for CollectionDate to be an explicit nil

### UnsetCollectionDate
`func (o *DdCaptureSession) UnsetCollectionDate()`

UnsetCollectionDate ensures that no value is present for CollectionDate, not even an explicit nil
### GetSettlementDate

`func (o *DdCaptureSession) GetSettlementDate() string`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *DdCaptureSession) GetSettlementDateOk() (*string, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *DdCaptureSession) SetSettlementDate(v string)`

SetSettlementDate sets SettlementDate field to given value.


### SetSettlementDateNil

`func (o *DdCaptureSession) SetSettlementDateNil(b bool)`

 SetSettlementDateNil sets the value for SettlementDate to be an explicit nil

### UnsetSettlementDate
`func (o *DdCaptureSession) UnsetSettlementDate()`

UnsetSettlementDate ensures that no value is present for SettlementDate, not even an explicit nil
### GetNoticeDays

`func (o *DdCaptureSession) GetNoticeDays() int32`

GetNoticeDays returns the NoticeDays field if non-nil, zero value otherwise.

### GetNoticeDaysOk

`func (o *DdCaptureSession) GetNoticeDaysOk() (*int32, bool)`

GetNoticeDaysOk returns a tuple with the NoticeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeDays

`func (o *DdCaptureSession) SetNoticeDays(v int32)`

SetNoticeDays sets NoticeDays field to given value.


### SetNoticeDaysNil

`func (o *DdCaptureSession) SetNoticeDaysNil(b bool)`

 SetNoticeDaysNil sets the value for NoticeDays to be an explicit nil

### UnsetNoticeDays
`func (o *DdCaptureSession) UnsetNoticeDays()`

UnsetNoticeDays ensures that no value is present for NoticeDays, not even an explicit nil
### GetMandateId

`func (o *DdCaptureSession) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *DdCaptureSession) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *DdCaptureSession) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### SetMandateIdNil

`func (o *DdCaptureSession) SetMandateIdNil(b bool)`

 SetMandateIdNil sets the value for MandateId to be an explicit nil

### UnsetMandateId
`func (o *DdCaptureSession) UnsetMandateId()`

UnsetMandateId ensures that no value is present for MandateId, not even an explicit nil
### GetExpiresAt

`func (o *DdCaptureSession) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DdCaptureSession) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DdCaptureSession) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *DdCaptureSession) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *DdCaptureSession) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetCompletedAt

`func (o *DdCaptureSession) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DdCaptureSession) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DdCaptureSession) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *DdCaptureSession) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *DdCaptureSession) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetAbandonedAt

`func (o *DdCaptureSession) GetAbandonedAt() string`

GetAbandonedAt returns the AbandonedAt field if non-nil, zero value otherwise.

### GetAbandonedAtOk

`func (o *DdCaptureSession) GetAbandonedAtOk() (*string, bool)`

GetAbandonedAtOk returns a tuple with the AbandonedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandonedAt

`func (o *DdCaptureSession) SetAbandonedAt(v string)`

SetAbandonedAt sets AbandonedAt field to given value.


### SetAbandonedAtNil

`func (o *DdCaptureSession) SetAbandonedAtNil(b bool)`

 SetAbandonedAtNil sets the value for AbandonedAt to be an explicit nil

### UnsetAbandonedAt
`func (o *DdCaptureSession) UnsetAbandonedAt()`

UnsetAbandonedAt ensures that no value is present for AbandonedAt, not even an explicit nil
### GetCreatedAt

`func (o *DdCaptureSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DdCaptureSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DdCaptureSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *DdCaptureSession) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *DdCaptureSession) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *DdCaptureSession) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DdCaptureSession) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DdCaptureSession) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *DdCaptureSession) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *DdCaptureSession) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


