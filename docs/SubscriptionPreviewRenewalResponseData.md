# SubscriptionPreviewRenewalResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** |  | 
**CurrentPeriodEnd** | **time.Time** |  | 
**UpcomingRenewals** | [**[]SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner**](SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner.md) |  | 
**NextInvoiceAmountMinor** | **int32** |  | 
**Currency** | **string** |  | 

## Methods

### NewSubscriptionPreviewRenewalResponseData

`func NewSubscriptionPreviewRenewalResponseData(subscriptionId string, currentPeriodEnd time.Time, upcomingRenewals []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, nextInvoiceAmountMinor int32, currency string, ) *SubscriptionPreviewRenewalResponseData`

NewSubscriptionPreviewRenewalResponseData instantiates a new SubscriptionPreviewRenewalResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPreviewRenewalResponseDataWithDefaults

`func NewSubscriptionPreviewRenewalResponseDataWithDefaults() *SubscriptionPreviewRenewalResponseData`

NewSubscriptionPreviewRenewalResponseDataWithDefaults instantiates a new SubscriptionPreviewRenewalResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *SubscriptionPreviewRenewalResponseData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionPreviewRenewalResponseData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionPreviewRenewalResponseData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCurrentPeriodEnd

`func (o *SubscriptionPreviewRenewalResponseData) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *SubscriptionPreviewRenewalResponseData) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *SubscriptionPreviewRenewalResponseData) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetUpcomingRenewals

`func (o *SubscriptionPreviewRenewalResponseData) GetUpcomingRenewals() []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner`

GetUpcomingRenewals returns the UpcomingRenewals field if non-nil, zero value otherwise.

### GetUpcomingRenewalsOk

`func (o *SubscriptionPreviewRenewalResponseData) GetUpcomingRenewalsOk() (*[]SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner, bool)`

GetUpcomingRenewalsOk returns a tuple with the UpcomingRenewals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcomingRenewals

`func (o *SubscriptionPreviewRenewalResponseData) SetUpcomingRenewals(v []SubscriptionPreviewRenewalResponseDataUpcomingRenewalsInner)`

SetUpcomingRenewals sets UpcomingRenewals field to given value.


### GetNextInvoiceAmountMinor

`func (o *SubscriptionPreviewRenewalResponseData) GetNextInvoiceAmountMinor() int32`

GetNextInvoiceAmountMinor returns the NextInvoiceAmountMinor field if non-nil, zero value otherwise.

### GetNextInvoiceAmountMinorOk

`func (o *SubscriptionPreviewRenewalResponseData) GetNextInvoiceAmountMinorOk() (*int32, bool)`

GetNextInvoiceAmountMinorOk returns a tuple with the NextInvoiceAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceAmountMinor

`func (o *SubscriptionPreviewRenewalResponseData) SetNextInvoiceAmountMinor(v int32)`

SetNextInvoiceAmountMinor sets NextInvoiceAmountMinor field to given value.


### GetCurrency

`func (o *SubscriptionPreviewRenewalResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionPreviewRenewalResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionPreviewRenewalResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


