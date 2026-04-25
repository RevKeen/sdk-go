# OpenDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the dispute | 
**PublicId** | **string** | Public-facing dispute identifier | 
**PaymentId** | **NullableString** | The original payment ID this dispute is for | 
**ParentTransactionId** | **NullableString** | Parent transaction ID in unified transaction model | 
**Gateway** | **string** | Payment gateway that processed the original transaction | 
**AmountMinor** | **float32** | Disputed amount in cents | 
**Currency** | **string** | Three-letter ISO 4217 currency code | 
**Reason** | **NullableString** | Human-readable reason for the dispute | 
**CustomerName** | **NullableString** | Name of the customer who filed the dispute | 
**CardBrand** | **NullableString** | Card brand used for the original payment | 
**CardLast4** | **NullableString** | Last four digits of the card used for the original payment | 
**EvidenceDueBy** | **NullableString** | Deadline for submitting evidence to contest the dispute | 
**DaysUntilDue** | **NullableFloat32** | Number of days remaining until the evidence deadline | 
**IsOverdue** | **bool** | Whether the evidence submission deadline has passed | 
**DisputedAt** | **string** | Timestamp when the dispute was filed | 

## Methods

### NewOpenDispute

`func NewOpenDispute(id string, publicId string, paymentId NullableString, parentTransactionId NullableString, gateway string, amountMinor float32, currency string, reason NullableString, customerName NullableString, cardBrand NullableString, cardLast4 NullableString, evidenceDueBy NullableString, daysUntilDue NullableFloat32, isOverdue bool, disputedAt string, ) *OpenDispute`

NewOpenDispute instantiates a new OpenDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenDisputeWithDefaults

`func NewOpenDisputeWithDefaults() *OpenDispute`

NewOpenDisputeWithDefaults instantiates a new OpenDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenDispute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenDispute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenDispute) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *OpenDispute) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *OpenDispute) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *OpenDispute) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetPaymentId

`func (o *OpenDispute) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *OpenDispute) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *OpenDispute) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### SetPaymentIdNil

`func (o *OpenDispute) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *OpenDispute) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetParentTransactionId

`func (o *OpenDispute) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *OpenDispute) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *OpenDispute) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### SetParentTransactionIdNil

`func (o *OpenDispute) SetParentTransactionIdNil(b bool)`

 SetParentTransactionIdNil sets the value for ParentTransactionId to be an explicit nil

### UnsetParentTransactionId
`func (o *OpenDispute) UnsetParentTransactionId()`

UnsetParentTransactionId ensures that no value is present for ParentTransactionId, not even an explicit nil
### GetGateway

`func (o *OpenDispute) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OpenDispute) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OpenDispute) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetAmountMinor

`func (o *OpenDispute) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *OpenDispute) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *OpenDispute) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *OpenDispute) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OpenDispute) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OpenDispute) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReason

`func (o *OpenDispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OpenDispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OpenDispute) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *OpenDispute) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *OpenDispute) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetCustomerName

`func (o *OpenDispute) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *OpenDispute) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *OpenDispute) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### SetCustomerNameNil

`func (o *OpenDispute) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *OpenDispute) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetCardBrand

`func (o *OpenDispute) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *OpenDispute) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *OpenDispute) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.


### SetCardBrandNil

`func (o *OpenDispute) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *OpenDispute) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLast4

`func (o *OpenDispute) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *OpenDispute) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *OpenDispute) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.


### SetCardLast4Nil

`func (o *OpenDispute) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *OpenDispute) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetEvidenceDueBy

`func (o *OpenDispute) GetEvidenceDueBy() string`

GetEvidenceDueBy returns the EvidenceDueBy field if non-nil, zero value otherwise.

### GetEvidenceDueByOk

`func (o *OpenDispute) GetEvidenceDueByOk() (*string, bool)`

GetEvidenceDueByOk returns a tuple with the EvidenceDueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDueBy

`func (o *OpenDispute) SetEvidenceDueBy(v string)`

SetEvidenceDueBy sets EvidenceDueBy field to given value.


### SetEvidenceDueByNil

`func (o *OpenDispute) SetEvidenceDueByNil(b bool)`

 SetEvidenceDueByNil sets the value for EvidenceDueBy to be an explicit nil

### UnsetEvidenceDueBy
`func (o *OpenDispute) UnsetEvidenceDueBy()`

UnsetEvidenceDueBy ensures that no value is present for EvidenceDueBy, not even an explicit nil
### GetDaysUntilDue

`func (o *OpenDispute) GetDaysUntilDue() float32`

GetDaysUntilDue returns the DaysUntilDue field if non-nil, zero value otherwise.

### GetDaysUntilDueOk

`func (o *OpenDispute) GetDaysUntilDueOk() (*float32, bool)`

GetDaysUntilDueOk returns a tuple with the DaysUntilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilDue

`func (o *OpenDispute) SetDaysUntilDue(v float32)`

SetDaysUntilDue sets DaysUntilDue field to given value.


### SetDaysUntilDueNil

`func (o *OpenDispute) SetDaysUntilDueNil(b bool)`

 SetDaysUntilDueNil sets the value for DaysUntilDue to be an explicit nil

### UnsetDaysUntilDue
`func (o *OpenDispute) UnsetDaysUntilDue()`

UnsetDaysUntilDue ensures that no value is present for DaysUntilDue, not even an explicit nil
### GetIsOverdue

`func (o *OpenDispute) GetIsOverdue() bool`

GetIsOverdue returns the IsOverdue field if non-nil, zero value otherwise.

### GetIsOverdueOk

`func (o *OpenDispute) GetIsOverdueOk() (*bool, bool)`

GetIsOverdueOk returns a tuple with the IsOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdue

`func (o *OpenDispute) SetIsOverdue(v bool)`

SetIsOverdue sets IsOverdue field to given value.


### GetDisputedAt

`func (o *OpenDispute) GetDisputedAt() string`

GetDisputedAt returns the DisputedAt field if non-nil, zero value otherwise.

### GetDisputedAtOk

`func (o *OpenDispute) GetDisputedAtOk() (*string, bool)`

GetDisputedAtOk returns a tuple with the DisputedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAt

`func (o *OpenDispute) SetDisputedAt(v string)`

SetDisputedAt sets DisputedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


