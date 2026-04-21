# SubscriptionsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** |  | 
**BusinessId** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PriceId** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**Quantity** | Pointer to **int32** |  | [optional] [default to 1]
**TrialEnd** | Pointer to **time.Time** |  | [optional] 
**BillingAnchorDay** | Pointer to **int32** |  | [optional] 
**PauseAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **bool** |  | [optional] 
**BillingStartsOn** | Pointer to **time.Time** |  | [optional] 
**BillingEndStrategy** | Pointer to **string** |  | [optional] 
**BillingEndsOn** | Pointer to **time.Time** |  | [optional] 
**BillingMaxCycles** | Pointer to **int32** |  | [optional] 
**ProrationMode** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSubscriptionsCreateRequest

`func NewSubscriptionsCreateRequest(customerId string, ) *SubscriptionsCreateRequest`

NewSubscriptionsCreateRequest instantiates a new SubscriptionsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsCreateRequestWithDefaults

`func NewSubscriptionsCreateRequestWithDefaults() *SubscriptionsCreateRequest`

NewSubscriptionsCreateRequestWithDefaults instantiates a new SubscriptionsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *SubscriptionsCreateRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *SubscriptionsCreateRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *SubscriptionsCreateRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetBusinessId

`func (o *SubscriptionsCreateRequest) GetBusinessId() string`

GetBusinessId returns the BusinessId field if non-nil, zero value otherwise.

### GetBusinessIdOk

`func (o *SubscriptionsCreateRequest) GetBusinessIdOk() (*string, bool)`

GetBusinessIdOk returns a tuple with the BusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessId

`func (o *SubscriptionsCreateRequest) SetBusinessId(v string)`

SetBusinessId sets BusinessId field to given value.

### HasBusinessId

`func (o *SubscriptionsCreateRequest) HasBusinessId() bool`

HasBusinessId returns a boolean if a field has been set.

### GetProductId

`func (o *SubscriptionsCreateRequest) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SubscriptionsCreateRequest) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SubscriptionsCreateRequest) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SubscriptionsCreateRequest) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionsCreateRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionsCreateRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionsCreateRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionsCreateRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPriceId

`func (o *SubscriptionsCreateRequest) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *SubscriptionsCreateRequest) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *SubscriptionsCreateRequest) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.

### HasPriceId

`func (o *SubscriptionsCreateRequest) HasPriceId() bool`

HasPriceId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *SubscriptionsCreateRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *SubscriptionsCreateRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *SubscriptionsCreateRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *SubscriptionsCreateRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetStartDate

`func (o *SubscriptionsCreateRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionsCreateRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionsCreateRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionsCreateRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionsCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionsCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionsCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionsCreateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetQuantity

`func (o *SubscriptionsCreateRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SubscriptionsCreateRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SubscriptionsCreateRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SubscriptionsCreateRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTrialEnd

`func (o *SubscriptionsCreateRequest) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *SubscriptionsCreateRequest) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *SubscriptionsCreateRequest) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *SubscriptionsCreateRequest) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *SubscriptionsCreateRequest) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *SubscriptionsCreateRequest) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *SubscriptionsCreateRequest) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *SubscriptionsCreateRequest) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### GetPauseAtPeriodEnd

`func (o *SubscriptionsCreateRequest) GetPauseAtPeriodEnd() bool`

GetPauseAtPeriodEnd returns the PauseAtPeriodEnd field if non-nil, zero value otherwise.

### GetPauseAtPeriodEndOk

`func (o *SubscriptionsCreateRequest) GetPauseAtPeriodEndOk() (*bool, bool)`

GetPauseAtPeriodEndOk returns a tuple with the PauseAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseAtPeriodEnd

`func (o *SubscriptionsCreateRequest) SetPauseAtPeriodEnd(v bool)`

SetPauseAtPeriodEnd sets PauseAtPeriodEnd field to given value.

### HasPauseAtPeriodEnd

`func (o *SubscriptionsCreateRequest) HasPauseAtPeriodEnd() bool`

HasPauseAtPeriodEnd returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *SubscriptionsCreateRequest) GetCancelAtPeriodEnd() bool`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *SubscriptionsCreateRequest) GetCancelAtPeriodEndOk() (*bool, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *SubscriptionsCreateRequest) SetCancelAtPeriodEnd(v bool)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *SubscriptionsCreateRequest) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetBillingStartsOn

`func (o *SubscriptionsCreateRequest) GetBillingStartsOn() time.Time`

GetBillingStartsOn returns the BillingStartsOn field if non-nil, zero value otherwise.

### GetBillingStartsOnOk

`func (o *SubscriptionsCreateRequest) GetBillingStartsOnOk() (*time.Time, bool)`

GetBillingStartsOnOk returns a tuple with the BillingStartsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartsOn

`func (o *SubscriptionsCreateRequest) SetBillingStartsOn(v time.Time)`

SetBillingStartsOn sets BillingStartsOn field to given value.

### HasBillingStartsOn

`func (o *SubscriptionsCreateRequest) HasBillingStartsOn() bool`

HasBillingStartsOn returns a boolean if a field has been set.

### GetBillingEndStrategy

`func (o *SubscriptionsCreateRequest) GetBillingEndStrategy() string`

GetBillingEndStrategy returns the BillingEndStrategy field if non-nil, zero value otherwise.

### GetBillingEndStrategyOk

`func (o *SubscriptionsCreateRequest) GetBillingEndStrategyOk() (*string, bool)`

GetBillingEndStrategyOk returns a tuple with the BillingEndStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndStrategy

`func (o *SubscriptionsCreateRequest) SetBillingEndStrategy(v string)`

SetBillingEndStrategy sets BillingEndStrategy field to given value.

### HasBillingEndStrategy

`func (o *SubscriptionsCreateRequest) HasBillingEndStrategy() bool`

HasBillingEndStrategy returns a boolean if a field has been set.

### GetBillingEndsOn

`func (o *SubscriptionsCreateRequest) GetBillingEndsOn() time.Time`

GetBillingEndsOn returns the BillingEndsOn field if non-nil, zero value otherwise.

### GetBillingEndsOnOk

`func (o *SubscriptionsCreateRequest) GetBillingEndsOnOk() (*time.Time, bool)`

GetBillingEndsOnOk returns a tuple with the BillingEndsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndsOn

`func (o *SubscriptionsCreateRequest) SetBillingEndsOn(v time.Time)`

SetBillingEndsOn sets BillingEndsOn field to given value.

### HasBillingEndsOn

`func (o *SubscriptionsCreateRequest) HasBillingEndsOn() bool`

HasBillingEndsOn returns a boolean if a field has been set.

### GetBillingMaxCycles

`func (o *SubscriptionsCreateRequest) GetBillingMaxCycles() int32`

GetBillingMaxCycles returns the BillingMaxCycles field if non-nil, zero value otherwise.

### GetBillingMaxCyclesOk

`func (o *SubscriptionsCreateRequest) GetBillingMaxCyclesOk() (*int32, bool)`

GetBillingMaxCyclesOk returns a tuple with the BillingMaxCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMaxCycles

`func (o *SubscriptionsCreateRequest) SetBillingMaxCycles(v int32)`

SetBillingMaxCycles sets BillingMaxCycles field to given value.

### HasBillingMaxCycles

`func (o *SubscriptionsCreateRequest) HasBillingMaxCycles() bool`

HasBillingMaxCycles returns a boolean if a field has been set.

### GetProrationMode

`func (o *SubscriptionsCreateRequest) GetProrationMode() string`

GetProrationMode returns the ProrationMode field if non-nil, zero value otherwise.

### GetProrationModeOk

`func (o *SubscriptionsCreateRequest) GetProrationModeOk() (*string, bool)`

GetProrationModeOk returns a tuple with the ProrationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationMode

`func (o *SubscriptionsCreateRequest) SetProrationMode(v string)`

SetProrationMode sets ProrationMode field to given value.

### HasProrationMode

`func (o *SubscriptionsCreateRequest) HasProrationMode() bool`

HasProrationMode returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscriptionsCreateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscriptionsCreateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscriptionsCreateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscriptionsCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


