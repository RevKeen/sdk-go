# UpdateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AmountMinor** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to **NullableString** |  | [optional] 
**IntervalCount** | Pointer to **NullableInt32** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**FulfillmentType** | Pointer to **string** | Fulfillment type for the product. | [optional] 
**BillingAnchorRule** | Pointer to **string** | How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month. | [optional] 
**BillingAnchorDay** | Pointer to **NullableInt32** |  | [optional] 
**FirstChargeBehavior** | Pointer to **string** | When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date. | [optional] 
**EndBehavior** | Pointer to **string** | How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles. | [optional] 
**MaxPayments** | Pointer to **NullableInt32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata for the product | [optional] 

## Methods

### NewUpdateProductRequest

`func NewUpdateProductRequest() *UpdateProductRequest`

NewUpdateProductRequest instantiates a new UpdateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductRequestWithDefaults

`func NewUpdateProductRequestWithDefaults() *UpdateProductRequest`

NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateProductRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateProductRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAmountMinor

`func (o *UpdateProductRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *UpdateProductRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *UpdateProductRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *UpdateProductRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateProductRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateProductRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateProductRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *UpdateProductRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateProductRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateProductRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateProductRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *UpdateProductRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *UpdateProductRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *UpdateProductRequest) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UpdateProductRequest) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UpdateProductRequest) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UpdateProductRequest) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *UpdateProductRequest) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *UpdateProductRequest) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialDays

`func (o *UpdateProductRequest) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *UpdateProductRequest) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *UpdateProductRequest) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *UpdateProductRequest) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateProductRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateProductRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateProductRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateProductRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsArchived

`func (o *UpdateProductRequest) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UpdateProductRequest) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UpdateProductRequest) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *UpdateProductRequest) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetFulfillmentType

`func (o *UpdateProductRequest) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *UpdateProductRequest) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *UpdateProductRequest) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *UpdateProductRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetBillingAnchorRule

`func (o *UpdateProductRequest) GetBillingAnchorRule() string`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *UpdateProductRequest) GetBillingAnchorRuleOk() (*string, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *UpdateProductRequest) SetBillingAnchorRule(v string)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *UpdateProductRequest) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *UpdateProductRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *UpdateProductRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *UpdateProductRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *UpdateProductRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *UpdateProductRequest) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *UpdateProductRequest) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetFirstChargeBehavior

`func (o *UpdateProductRequest) GetFirstChargeBehavior() string`

GetFirstChargeBehavior returns the FirstChargeBehavior field if non-nil, zero value otherwise.

### GetFirstChargeBehaviorOk

`func (o *UpdateProductRequest) GetFirstChargeBehaviorOk() (*string, bool)`

GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeBehavior

`func (o *UpdateProductRequest) SetFirstChargeBehavior(v string)`

SetFirstChargeBehavior sets FirstChargeBehavior field to given value.

### HasFirstChargeBehavior

`func (o *UpdateProductRequest) HasFirstChargeBehavior() bool`

HasFirstChargeBehavior returns a boolean if a field has been set.

### GetEndBehavior

`func (o *UpdateProductRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *UpdateProductRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *UpdateProductRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *UpdateProductRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetMaxPayments

`func (o *UpdateProductRequest) GetMaxPayments() int32`

GetMaxPayments returns the MaxPayments field if non-nil, zero value otherwise.

### GetMaxPaymentsOk

`func (o *UpdateProductRequest) GetMaxPaymentsOk() (*int32, bool)`

GetMaxPaymentsOk returns a tuple with the MaxPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayments

`func (o *UpdateProductRequest) SetMaxPayments(v int32)`

SetMaxPayments sets MaxPayments field to given value.

### HasMaxPayments

`func (o *UpdateProductRequest) HasMaxPayments() bool`

HasMaxPayments returns a boolean if a field has been set.

### SetMaxPaymentsNil

`func (o *UpdateProductRequest) SetMaxPaymentsNil(b bool)`

 SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil

### UnsetMaxPayments
`func (o *UpdateProductRequest) UnsetMaxPayments()`

UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
### GetMetadata

`func (o *UpdateProductRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateProductRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateProductRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateProductRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


