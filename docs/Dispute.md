# Dispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the dispute | 
**PublicId** | **string** | Public-facing dispute identifier | 
**PaymentId** | **NullableString** | The original payment ID this dispute is for (alias for parent_transaction_id) | 
**ParentTransactionId** | **NullableString** | Parent transaction ID in unified transaction model. Same as payment_id for disputes. | 
**Gateway** | **string** | Name of the payment processor that handled the original transaction | 
**GatewayDisputeId** | **string** | Dispute identifier assigned by the payment gateway | 
**GatewayTransactionId** | **NullableString** | Original transaction identifier from the payment gateway | 
**AmountMinor** | **float32** | Disputed amount in cents | 
**Currency** | **string** | Three-letter ISO 4217 currency code | 
**ReasonCode** | **NullableString** | Card network reason code | 
**Reason** | **NullableString** | Human-readable reason for the dispute | 
**NetworkReasonCode** | **NullableString** | Card network-specific reason code | 
**NetworkReasonDescription** | **NullableString** | Card network-specific reason description | 
**CustomerName** | **NullableString** | Name of the customer who filed the dispute | 
**CardBrand** | **NullableString** | Card brand used for the original payment | 
**CardLast4** | **NullableString** | Last four digits of the card used for the original payment | 
**EvidenceDueBy** | **NullableString** | Deadline for submitting evidence to contest the dispute | 
**EvidenceSubmitted** | **bool** | Whether evidence has been submitted for this dispute | 
**EvidenceSubmittedAt** | **NullableString** | Timestamp when evidence was submitted | 
**Status** | **string** | Current status of the dispute lifecycle | 
**Resolution** | **NullableString** | Final resolution outcome of the dispute | 
**ResolvedAt** | **NullableString** | Timestamp when the dispute was resolved | 
**DisputedAt** | **string** | Timestamp when the dispute was filed | 
**CreatedAt** | **string** | Timestamp when the dispute record was created | 

## Methods

### NewDispute

`func NewDispute(id string, publicId string, paymentId NullableString, parentTransactionId NullableString, gateway string, gatewayDisputeId string, gatewayTransactionId NullableString, amountMinor float32, currency string, reasonCode NullableString, reason NullableString, networkReasonCode NullableString, networkReasonDescription NullableString, customerName NullableString, cardBrand NullableString, cardLast4 NullableString, evidenceDueBy NullableString, evidenceSubmitted bool, evidenceSubmittedAt NullableString, status string, resolution NullableString, resolvedAt NullableString, disputedAt string, createdAt string, ) *Dispute`

NewDispute instantiates a new Dispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisputeWithDefaults

`func NewDisputeWithDefaults() *Dispute`

NewDisputeWithDefaults instantiates a new Dispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Dispute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dispute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dispute) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *Dispute) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Dispute) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Dispute) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetPaymentId

`func (o *Dispute) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Dispute) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Dispute) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### SetPaymentIdNil

`func (o *Dispute) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *Dispute) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetParentTransactionId

`func (o *Dispute) GetParentTransactionId() string`

GetParentTransactionId returns the ParentTransactionId field if non-nil, zero value otherwise.

### GetParentTransactionIdOk

`func (o *Dispute) GetParentTransactionIdOk() (*string, bool)`

GetParentTransactionIdOk returns a tuple with the ParentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTransactionId

`func (o *Dispute) SetParentTransactionId(v string)`

SetParentTransactionId sets ParentTransactionId field to given value.


### SetParentTransactionIdNil

`func (o *Dispute) SetParentTransactionIdNil(b bool)`

 SetParentTransactionIdNil sets the value for ParentTransactionId to be an explicit nil

### UnsetParentTransactionId
`func (o *Dispute) UnsetParentTransactionId()`

UnsetParentTransactionId ensures that no value is present for ParentTransactionId, not even an explicit nil
### GetGateway

`func (o *Dispute) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Dispute) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Dispute) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayDisputeId

`func (o *Dispute) GetGatewayDisputeId() string`

GetGatewayDisputeId returns the GatewayDisputeId field if non-nil, zero value otherwise.

### GetGatewayDisputeIdOk

`func (o *Dispute) GetGatewayDisputeIdOk() (*string, bool)`

GetGatewayDisputeIdOk returns a tuple with the GatewayDisputeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDisputeId

`func (o *Dispute) SetGatewayDisputeId(v string)`

SetGatewayDisputeId sets GatewayDisputeId field to given value.


### GetGatewayTransactionId

`func (o *Dispute) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *Dispute) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *Dispute) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.


### SetGatewayTransactionIdNil

`func (o *Dispute) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *Dispute) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetAmountMinor

`func (o *Dispute) GetAmountMinor() float32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *Dispute) GetAmountMinorOk() (*float32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *Dispute) SetAmountMinor(v float32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *Dispute) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Dispute) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Dispute) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetReasonCode

`func (o *Dispute) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *Dispute) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *Dispute) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### SetReasonCodeNil

`func (o *Dispute) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *Dispute) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetReason

`func (o *Dispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Dispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Dispute) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *Dispute) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *Dispute) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetNetworkReasonCode

`func (o *Dispute) GetNetworkReasonCode() string`

GetNetworkReasonCode returns the NetworkReasonCode field if non-nil, zero value otherwise.

### GetNetworkReasonCodeOk

`func (o *Dispute) GetNetworkReasonCodeOk() (*string, bool)`

GetNetworkReasonCodeOk returns a tuple with the NetworkReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReasonCode

`func (o *Dispute) SetNetworkReasonCode(v string)`

SetNetworkReasonCode sets NetworkReasonCode field to given value.


### SetNetworkReasonCodeNil

`func (o *Dispute) SetNetworkReasonCodeNil(b bool)`

 SetNetworkReasonCodeNil sets the value for NetworkReasonCode to be an explicit nil

### UnsetNetworkReasonCode
`func (o *Dispute) UnsetNetworkReasonCode()`

UnsetNetworkReasonCode ensures that no value is present for NetworkReasonCode, not even an explicit nil
### GetNetworkReasonDescription

`func (o *Dispute) GetNetworkReasonDescription() string`

GetNetworkReasonDescription returns the NetworkReasonDescription field if non-nil, zero value otherwise.

### GetNetworkReasonDescriptionOk

`func (o *Dispute) GetNetworkReasonDescriptionOk() (*string, bool)`

GetNetworkReasonDescriptionOk returns a tuple with the NetworkReasonDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkReasonDescription

`func (o *Dispute) SetNetworkReasonDescription(v string)`

SetNetworkReasonDescription sets NetworkReasonDescription field to given value.


### SetNetworkReasonDescriptionNil

`func (o *Dispute) SetNetworkReasonDescriptionNil(b bool)`

 SetNetworkReasonDescriptionNil sets the value for NetworkReasonDescription to be an explicit nil

### UnsetNetworkReasonDescription
`func (o *Dispute) UnsetNetworkReasonDescription()`

UnsetNetworkReasonDescription ensures that no value is present for NetworkReasonDescription, not even an explicit nil
### GetCustomerName

`func (o *Dispute) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *Dispute) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *Dispute) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### SetCustomerNameNil

`func (o *Dispute) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *Dispute) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetCardBrand

`func (o *Dispute) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *Dispute) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *Dispute) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.


### SetCardBrandNil

`func (o *Dispute) SetCardBrandNil(b bool)`

 SetCardBrandNil sets the value for CardBrand to be an explicit nil

### UnsetCardBrand
`func (o *Dispute) UnsetCardBrand()`

UnsetCardBrand ensures that no value is present for CardBrand, not even an explicit nil
### GetCardLast4

`func (o *Dispute) GetCardLast4() string`

GetCardLast4 returns the CardLast4 field if non-nil, zero value otherwise.

### GetCardLast4Ok

`func (o *Dispute) GetCardLast4Ok() (*string, bool)`

GetCardLast4Ok returns a tuple with the CardLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLast4

`func (o *Dispute) SetCardLast4(v string)`

SetCardLast4 sets CardLast4 field to given value.


### SetCardLast4Nil

`func (o *Dispute) SetCardLast4Nil(b bool)`

 SetCardLast4Nil sets the value for CardLast4 to be an explicit nil

### UnsetCardLast4
`func (o *Dispute) UnsetCardLast4()`

UnsetCardLast4 ensures that no value is present for CardLast4, not even an explicit nil
### GetEvidenceDueBy

`func (o *Dispute) GetEvidenceDueBy() string`

GetEvidenceDueBy returns the EvidenceDueBy field if non-nil, zero value otherwise.

### GetEvidenceDueByOk

`func (o *Dispute) GetEvidenceDueByOk() (*string, bool)`

GetEvidenceDueByOk returns a tuple with the EvidenceDueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDueBy

`func (o *Dispute) SetEvidenceDueBy(v string)`

SetEvidenceDueBy sets EvidenceDueBy field to given value.


### SetEvidenceDueByNil

`func (o *Dispute) SetEvidenceDueByNil(b bool)`

 SetEvidenceDueByNil sets the value for EvidenceDueBy to be an explicit nil

### UnsetEvidenceDueBy
`func (o *Dispute) UnsetEvidenceDueBy()`

UnsetEvidenceDueBy ensures that no value is present for EvidenceDueBy, not even an explicit nil
### GetEvidenceSubmitted

`func (o *Dispute) GetEvidenceSubmitted() bool`

GetEvidenceSubmitted returns the EvidenceSubmitted field if non-nil, zero value otherwise.

### GetEvidenceSubmittedOk

`func (o *Dispute) GetEvidenceSubmittedOk() (*bool, bool)`

GetEvidenceSubmittedOk returns a tuple with the EvidenceSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSubmitted

`func (o *Dispute) SetEvidenceSubmitted(v bool)`

SetEvidenceSubmitted sets EvidenceSubmitted field to given value.


### GetEvidenceSubmittedAt

`func (o *Dispute) GetEvidenceSubmittedAt() string`

GetEvidenceSubmittedAt returns the EvidenceSubmittedAt field if non-nil, zero value otherwise.

### GetEvidenceSubmittedAtOk

`func (o *Dispute) GetEvidenceSubmittedAtOk() (*string, bool)`

GetEvidenceSubmittedAtOk returns a tuple with the EvidenceSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceSubmittedAt

`func (o *Dispute) SetEvidenceSubmittedAt(v string)`

SetEvidenceSubmittedAt sets EvidenceSubmittedAt field to given value.


### SetEvidenceSubmittedAtNil

`func (o *Dispute) SetEvidenceSubmittedAtNil(b bool)`

 SetEvidenceSubmittedAtNil sets the value for EvidenceSubmittedAt to be an explicit nil

### UnsetEvidenceSubmittedAt
`func (o *Dispute) UnsetEvidenceSubmittedAt()`

UnsetEvidenceSubmittedAt ensures that no value is present for EvidenceSubmittedAt, not even an explicit nil
### GetStatus

`func (o *Dispute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Dispute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Dispute) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetResolution

`func (o *Dispute) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *Dispute) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *Dispute) SetResolution(v string)`

SetResolution sets Resolution field to given value.


### SetResolutionNil

`func (o *Dispute) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *Dispute) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetResolvedAt

`func (o *Dispute) GetResolvedAt() string`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Dispute) GetResolvedAtOk() (*string, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Dispute) SetResolvedAt(v string)`

SetResolvedAt sets ResolvedAt field to given value.


### SetResolvedAtNil

`func (o *Dispute) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *Dispute) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetDisputedAt

`func (o *Dispute) GetDisputedAt() string`

GetDisputedAt returns the DisputedAt field if non-nil, zero value otherwise.

### GetDisputedAtOk

`func (o *Dispute) GetDisputedAtOk() (*string, bool)`

GetDisputedAtOk returns a tuple with the DisputedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedAt

`func (o *Dispute) SetDisputedAt(v string)`

SetDisputedAt sets DisputedAt field to given value.


### GetCreatedAt

`func (o *Dispute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Dispute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Dispute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


