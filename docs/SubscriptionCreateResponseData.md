# SubscriptionCreateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MerchantId** | **string** |  | 
**CustomerId** | **string** |  | 
**ProductId** | Pointer to **string** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Quantity** | **int32** |  | 
**AmountMinor** | **int32** |  | 
**Currency** | **string** |  | 
**BillingInterval** | **string** |  | 
**CurrentPeriodStart** | **time.Time** |  | 
**CurrentPeriodEnd** | **time.Time** |  | 
**TrialEnd** | Pointer to **time.Time** |  | [optional] 
**CanceledAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Dunning** | Pointer to [**SubscriptionCreateResponseDataDunning**](SubscriptionCreateResponseDataDunning.md) |  | [optional] 

## Methods

### NewSubscriptionCreateResponseData

`func NewSubscriptionCreateResponseData(id string, merchantId string, customerId string, status string, quantity int32, amountMinor int32, currency string, billingInterval string, currentPeriodStart time.Time, currentPeriodEnd time.Time, createdAt time.Time, updatedAt time.Time, ) *SubscriptionCreateResponseData`

NewSubscriptionCreateResponseData instantiates a new SubscriptionCreateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateResponseDataWithDefaults

`func NewSubscriptionCreateResponseDataWithDefaults() *SubscriptionCreateResponseData`

NewSubscriptionCreateResponseDataWithDefaults instantiates a new SubscriptionCreateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionCreateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionCreateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionCreateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetMerchantId

`func (o *SubscriptionCreateResponseData) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *SubscriptionCreateResponseData) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *SubscriptionCreateResponseData) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetCustomerId

`func (o *SubscriptionCreateResponseData) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionCreateResponseData) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionCreateResponseData) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetProductId

`func (o *SubscriptionCreateResponseData) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionCreateResponseData) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionCreateResponseData) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SubscriptionCreateResponseData) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPriceId

`func (o *SubscriptionCreateResponseData) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionCreateResponseData) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionCreateResponseData) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *SubscriptionCreateResponseData) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionCreateResponseData) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionCreateResponseData) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionCreateResponseData) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionCreateResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionCreateResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionCreateResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionCreateResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuantity

`func (o *SubscriptionCreateResponseData) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionCreateResponseData) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionCreateResponseData) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetAmountMinor

`func (o *SubscriptionCreateResponseData) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *SubscriptionCreateResponseData) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *SubscriptionCreateResponseData) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.


### GetCurrency

`func (o *SubscriptionCreateResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionCreateResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionCreateResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetBillingInterval

`func (o *SubscriptionCreateResponseData) GetBillingInterval() string`

GetBillingInterval returns the BillingInterval field if non-nil, zero value otherwise.

### GetBillingIntervalOk

`func (o *SubscriptionCreateResponseData) GetBillingIntervalOk() (*string, bool)`

GetBillingIntervalOk returns a tuple with the BillingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInterval

`func (o *SubscriptionCreateResponseData) SetBillingInterval(v string)`

SetBillingInterval sets BillingInterval field to given value.


### GetCurrentPeriodStart

`func (o *SubscriptionCreateResponseData) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *SubscriptionCreateResponseData) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *SubscriptionCreateResponseData) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### GetCurrentPeriodEnd

`func (o *SubscriptionCreateResponseData) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionCreateResponseData) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionCreateResponseData) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetTrialEnd

`func (o *SubscriptionCreateResponseData) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionCreateResponseData) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionCreateResponseData) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionCreateResponseData) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetCanceledAt

`func (o *SubscriptionCreateResponseData) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *SubscriptionCreateResponseData) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *SubscriptionCreateResponseData) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *SubscriptionCreateResponseData) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriptionCreateResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriptionCreateResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriptionCreateResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SubscriptionCreateResponseData) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriptionCreateResponseData) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriptionCreateResponseData) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDunning

`func (o *SubscriptionCreateResponseData) GetDunning() SubscriptionCreateResponseDataDunning`

GetDunning returns the Dunning field if non-nil, zero value otherwise.

### GetDunningOk

`func (o *SubscriptionCreateResponseData) GetDunningOk() (*SubscriptionCreateResponseDataDunning, bool)`

GetDunningOk returns a tuple with the Dunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunning

`func (o *SubscriptionCreateResponseData) SetDunning(v SubscriptionCreateResponseDataDunning)`

SetDunning sets Dunning field to given value.

### HasDunning

`func (o *SubscriptionCreateResponseData) HasDunning() bool`

HasDunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


