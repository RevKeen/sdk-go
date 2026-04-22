# CreateCreditNoteInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The ID of the invoice to issue a credit note for | 
**AmountMinor** | **int32** | Amount to credit in minor units (cents). Capped at 2,000,000,000 to stay below the 32-bit integer ceiling of the underlying DB column. | 
**TaxAmountMinor** | Pointer to **int32** | Tax portion of the credit in minor units (cents). For UK/EU VAT compliance. Capped at 2,000,000,000. | [optional] 
**CreditMethod** | Pointer to **string** | How the credit should be applied | [optional] [default to "refund_to_payment_method"]
**Reason** | Pointer to **string** | Reason for the credit note | [optional] 
**ReasonCode** | Pointer to **string** | Standardized reason code | [optional] 
**CancelSubscription** | Pointer to **bool** | Whether to cancel the associated subscription after issuing | [optional] [default to false]
**IsProrated** | Pointer to **bool** | Whether this credit note represents a prorated amount | [optional] 
**ProrationDaysTotal** | Pointer to **int32** | Total days in the billing period (for prorated credits). Constrained to 1–366. | [optional] 
**ProrationDaysUnused** | Pointer to **int32** | Unused days in the billing period (for prorated credits). Must be &lt;&#x3D; proration_days_total. Constrained to 0–366. | [optional] 
**IdempotencyKey** | Pointer to **string** | Idempotency key to prevent duplicate credit notes | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata | [optional] 
**AutoRoute** | Pointer to **bool** | When true, automatically trigger the process-credit-note task for multi-gateway reversal routing. Response will include gateway_operations array. Always combine with &#x60;idempotency_key&#x60; to avoid duplicate gateway operations. | [optional] [default to false]

## Methods

### NewCreateCreditNoteInput

`func NewCreateCreditNoteInput(invoiceId string, amountMinor int32, ) *CreateCreditNoteInput`

NewCreateCreditNoteInput instantiates a new CreateCreditNoteInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCreditNoteInputWithDefaults

`func NewCreateCreditNoteInputWithDefaults() *CreateCreditNoteInput`

NewCreateCreditNoteInputWithDefaults instantiates a new CreateCreditNoteInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *CreateCreditNoteInput) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreateCreditNoteInput) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreateCreditNoteInput) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetAmountMinor

`func (o *CreateCreditNoteInput) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateCreditNoteInput) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateCreditNoteInput) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetTaxAmountMinor

`func (o *CreateCreditNoteInput) GetTaxAmountMinor() int32`

GetTaxAmountMinor returns the TaxAmountMinor field if non-nil, zero value otherwise.

### GetTaxAmountMinorOk

`func (o *CreateCreditNoteInput) GetTaxAmountMinorOk() (*int32, bool)`

GetTaxAmountMinorOk returns a tuple with the TaxAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountMinor

`func (o *CreateCreditNoteInput) SetTaxAmountMinor(v int32)`

SetTaxAmountMinor sets TaxAmountMinor field to given value.

### HasTaxAmountMinor

`func (o *CreateCreditNoteInput) HasTaxAmountMinor() bool`

HasTaxAmountMinor returns a boolean if a field has been set.

### GetCreditMethod

`func (o *CreateCreditNoteInput) GetCreditMethod() string`

GetCreditMethod returns the CreditMethod field if non-nil, zero value otherwise.

### GetCreditMethodOk

`func (o *CreateCreditNoteInput) GetCreditMethodOk() (*string, bool)`

GetCreditMethodOk returns a tuple with the CreditMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMethod

`func (o *CreateCreditNoteInput) SetCreditMethod(v string)`

SetCreditMethod sets CreditMethod field to given value.

### HasCreditMethod

`func (o *CreateCreditNoteInput) HasCreditMethod() bool`

HasCreditMethod returns a boolean if a field has been set.

### GetReason

`func (o *CreateCreditNoteInput) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCreditNoteInput) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCreditNoteInput) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCreditNoteInput) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonCode

`func (o *CreateCreditNoteInput) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *CreateCreditNoteInput) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *CreateCreditNoteInput) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *CreateCreditNoteInput) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### GetCancelSubscription

`func (o *CreateCreditNoteInput) GetCancelSubscription() bool`

GetCancelSubscription returns the CancelSubscription field if non-nil, zero value otherwise.

### GetCancelSubscriptionOk

`func (o *CreateCreditNoteInput) GetCancelSubscriptionOk() (*bool, bool)`

GetCancelSubscriptionOk returns a tuple with the CancelSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSubscription

`func (o *CreateCreditNoteInput) SetCancelSubscription(v bool)`

SetCancelSubscription sets CancelSubscription field to given value.

### HasCancelSubscription

`func (o *CreateCreditNoteInput) HasCancelSubscription() bool`

HasCancelSubscription returns a boolean if a field has been set.

### GetIsProrated

`func (o *CreateCreditNoteInput) GetIsProrated() bool`

GetIsProrated returns the IsProrated field if non-nil, zero value otherwise.

### GetIsProratedOk

`func (o *CreateCreditNoteInput) GetIsProratedOk() (*bool, bool)`

GetIsProratedOk returns a tuple with the IsProrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProrated

`func (o *CreateCreditNoteInput) SetIsProrated(v bool)`

SetIsProrated sets IsProrated field to given value.

### HasIsProrated

`func (o *CreateCreditNoteInput) HasIsProrated() bool`

HasIsProrated returns a boolean if a field has been set.

### GetProrationDaysTotal

`func (o *CreateCreditNoteInput) GetProrationDaysTotal() int32`

GetProrationDaysTotal returns the ProrationDaysTotal field if non-nil, zero value otherwise.

### GetProrationDaysTotalOk

`func (o *CreateCreditNoteInput) GetProrationDaysTotalOk() (*int32, bool)`

GetProrationDaysTotalOk returns a tuple with the ProrationDaysTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDaysTotal

`func (o *CreateCreditNoteInput) SetProrationDaysTotal(v int32)`

SetProrationDaysTotal sets ProrationDaysTotal field to given value.

### HasProrationDaysTotal

`func (o *CreateCreditNoteInput) HasProrationDaysTotal() bool`

HasProrationDaysTotal returns a boolean if a field has been set.

### GetProrationDaysUnused

`func (o *CreateCreditNoteInput) GetProrationDaysUnused() int32`

GetProrationDaysUnused returns the ProrationDaysUnused field if non-nil, zero value otherwise.

### GetProrationDaysUnusedOk

`func (o *CreateCreditNoteInput) GetProrationDaysUnusedOk() (*int32, bool)`

GetProrationDaysUnusedOk returns a tuple with the ProrationDaysUnused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDaysUnused

`func (o *CreateCreditNoteInput) SetProrationDaysUnused(v int32)`

SetProrationDaysUnused sets ProrationDaysUnused field to given value.

### HasProrationDaysUnused

`func (o *CreateCreditNoteInput) HasProrationDaysUnused() bool`

HasProrationDaysUnused returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CreateCreditNoteInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CreateCreditNoteInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CreateCreditNoteInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CreateCreditNoteInput) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateCreditNoteInput) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateCreditNoteInput) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateCreditNoteInput) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateCreditNoteInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAutoRoute

`func (o *CreateCreditNoteInput) GetAutoRoute() bool`

GetAutoRoute returns the AutoRoute field if non-nil, zero value otherwise.

### GetAutoRouteOk

`func (o *CreateCreditNoteInput) GetAutoRouteOk() (*bool, bool)`

GetAutoRouteOk returns a tuple with the AutoRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRoute

`func (o *CreateCreditNoteInput) SetAutoRoute(v bool)`

SetAutoRoute sets AutoRoute field to given value.

### HasAutoRoute

`func (o *CreateCreditNoteInput) HasAutoRoute() bool`

HasAutoRoute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


