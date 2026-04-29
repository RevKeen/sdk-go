# CreateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Kind** | **string** |  | 
**PricingModel** | **string** |  | 
**AmountMinor** | **int32** |  | 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Interval** | Pointer to **NullableString** |  | [optional] 
**IntervalCount** | Pointer to **NullableInt32** |  | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**UsageMeterId** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**FulfillmentType** | Pointer to **string** | Fulfillment type for the product. Defaults to &#39;none&#39; (service). | [optional] [default to "none"]
**BillingAnchorRule** | Pointer to **string** | How billing dates are calculated. same_day: Bill on same day as start. day_of_month: Bill on specific day (1-31). last_day: Bill on last day of month. | [optional] 
**BillingAnchorDay** | Pointer to **NullableInt32** | Day of month (1-31) when billing_anchor_rule is &#39;day_of_month&#39; | [optional] 
**FirstChargeBehavior** | Pointer to **string** | When first payment is collected. immediate: Charge on start. next_anchor: Charge on first scheduled date. prorate: Prorate until first date. | [optional] 
**EndBehavior** | Pointer to **string** | How subscription ends. until_canceled: Runs forever. fixed_payments: Ends after N billing cycles. | [optional] 
**MaxPayments** | Pointer to **NullableInt32** | Max billing cycles when end_behavior is &#39;fixed_payments&#39; | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary key-value metadata for the product | [optional] 

## Methods

### NewCreateProductRequest

`func NewCreateProductRequest(productId string, name string, kind string, pricingModel string, amountMinor int32, ) *CreateProductRequest`

NewCreateProductRequest instantiates a new CreateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductRequestWithDefaults

`func NewCreateProductRequestWithDefaults() *CreateProductRequest`

NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CreateProductRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CreateProductRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CreateProductRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetName

`func (o *CreateProductRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProductRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProductRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateProductRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProductRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProductRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProductRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateProductRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateProductRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *CreateProductRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateProductRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateProductRequest) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPricingModel

`func (o *CreateProductRequest) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *CreateProductRequest) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *CreateProductRequest) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetAmountMinor

`func (o *CreateProductRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *CreateProductRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *CreateProductRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *CreateProductRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateProductRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateProductRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateProductRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInterval

`func (o *CreateProductRequest) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateProductRequest) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateProductRequest) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CreateProductRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *CreateProductRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *CreateProductRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetIntervalCount

`func (o *CreateProductRequest) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *CreateProductRequest) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *CreateProductRequest) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *CreateProductRequest) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### SetIntervalCountNil

`func (o *CreateProductRequest) SetIntervalCountNil(b bool)`

 SetIntervalCountNil sets the value for IntervalCount to be an explicit nil

### UnsetIntervalCount
`func (o *CreateProductRequest) UnsetIntervalCount()`

UnsetIntervalCount ensures that no value is present for IntervalCount, not even an explicit nil
### GetTrialDays

`func (o *CreateProductRequest) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *CreateProductRequest) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *CreateProductRequest) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *CreateProductRequest) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetUsageMeterId

`func (o *CreateProductRequest) GetUsageMeterId() string`

GetUsageMeterId returns the UsageMeterId field if non-nil, zero value otherwise.

### GetUsageMeterIdOk

`func (o *CreateProductRequest) GetUsageMeterIdOk() (*string, bool)`

GetUsageMeterIdOk returns a tuple with the UsageMeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeterId

`func (o *CreateProductRequest) SetUsageMeterId(v string)`

SetUsageMeterId sets UsageMeterId field to given value.

### HasUsageMeterId

`func (o *CreateProductRequest) HasUsageMeterId() bool`

HasUsageMeterId returns a boolean if a field has been set.

### SetUsageMeterIdNil

`func (o *CreateProductRequest) SetUsageMeterIdNil(b bool)`

 SetUsageMeterIdNil sets the value for UsageMeterId to be an explicit nil

### UnsetUsageMeterId
`func (o *CreateProductRequest) UnsetUsageMeterId()`

UnsetUsageMeterId ensures that no value is present for UsageMeterId, not even an explicit nil
### GetSlug

`func (o *CreateProductRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateProductRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateProductRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateProductRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetFulfillmentType

`func (o *CreateProductRequest) GetFulfillmentType() string`

GetFulfillmentType returns the FulfillmentType field if non-nil, zero value otherwise.

### GetFulfillmentTypeOk

`func (o *CreateProductRequest) GetFulfillmentTypeOk() (*string, bool)`

GetFulfillmentTypeOk returns a tuple with the FulfillmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentType

`func (o *CreateProductRequest) SetFulfillmentType(v string)`

SetFulfillmentType sets FulfillmentType field to given value.

### HasFulfillmentType

`func (o *CreateProductRequest) HasFulfillmentType() bool`

HasFulfillmentType returns a boolean if a field has been set.

### GetBillingAnchorRule

`func (o *CreateProductRequest) GetBillingAnchorRule() string`

GetBillingAnchorRule returns the BillingAnchorRule field if non-nil, zero value otherwise.

### GetBillingAnchorRuleOk

`func (o *CreateProductRequest) GetBillingAnchorRuleOk() (*string, bool)`

GetBillingAnchorRuleOk returns a tuple with the BillingAnchorRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorRule

`func (o *CreateProductRequest) SetBillingAnchorRule(v string)`

SetBillingAnchorRule sets BillingAnchorRule field to given value.

### HasBillingAnchorRule

`func (o *CreateProductRequest) HasBillingAnchorRule() bool`

HasBillingAnchorRule returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *CreateProductRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *CreateProductRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *CreateProductRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *CreateProductRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### SetBillingAnchorDayNil

`func (o *CreateProductRequest) SetBillingAnchorDayNil(b bool)`

 SetBillingAnchorDayNil sets the value for BillingAnchorDay to be an explicit nil

### UnsetBillingAnchorDay
`func (o *CreateProductRequest) UnsetBillingAnchorDay()`

UnsetBillingAnchorDay ensures that no value is present for BillingAnchorDay, not even an explicit nil
### GetFirstChargeBehavior

`func (o *CreateProductRequest) GetFirstChargeBehavior() string`

GetFirstChargeBehavior returns the FirstChargeBehavior field if non-nil, zero value otherwise.

### GetFirstChargeBehaviorOk

`func (o *CreateProductRequest) GetFirstChargeBehaviorOk() (*string, bool)`

GetFirstChargeBehaviorOk returns a tuple with the FirstChargeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstChargeBehavior

`func (o *CreateProductRequest) SetFirstChargeBehavior(v string)`

SetFirstChargeBehavior sets FirstChargeBehavior field to given value.

### HasFirstChargeBehavior

`func (o *CreateProductRequest) HasFirstChargeBehavior() bool`

HasFirstChargeBehavior returns a boolean if a field has been set.

### GetEndBehavior

`func (o *CreateProductRequest) GetEndBehavior() string`

GetEndBehavior returns the EndBehavior field if non-nil, zero value otherwise.

### GetEndBehaviorOk

`func (o *CreateProductRequest) GetEndBehaviorOk() (*string, bool)`

GetEndBehaviorOk returns a tuple with the EndBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBehavior

`func (o *CreateProductRequest) SetEndBehavior(v string)`

SetEndBehavior sets EndBehavior field to given value.

### HasEndBehavior

`func (o *CreateProductRequest) HasEndBehavior() bool`

HasEndBehavior returns a boolean if a field has been set.

### GetMaxPayments

`func (o *CreateProductRequest) GetMaxPayments() int32`

GetMaxPayments returns the MaxPayments field if non-nil, zero value otherwise.

### GetMaxPaymentsOk

`func (o *CreateProductRequest) GetMaxPaymentsOk() (*int32, bool)`

GetMaxPaymentsOk returns a tuple with the MaxPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayments

`func (o *CreateProductRequest) SetMaxPayments(v int32)`

SetMaxPayments sets MaxPayments field to given value.

### HasMaxPayments

`func (o *CreateProductRequest) HasMaxPayments() bool`

HasMaxPayments returns a boolean if a field has been set.

### SetMaxPaymentsNil

`func (o *CreateProductRequest) SetMaxPaymentsNil(b bool)`

 SetMaxPaymentsNil sets the value for MaxPayments to be an explicit nil

### UnsetMaxPayments
`func (o *CreateProductRequest) UnsetMaxPayments()`

UnsetMaxPayments ensures that no value is present for MaxPayments, not even an explicit nil
### GetMetadata

`func (o *CreateProductRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateProductRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateProductRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateProductRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


