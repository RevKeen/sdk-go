# InvoiceUsageExplainResponseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Description** | **NullableString** |  |
**Kind** | **string** |  |
**AmountMinor** | **NullableInt64** |  |
**Currency** | **NullableString** |  |
**MeterId** | **NullableString** |  |
**UsageQuantity** | **NullableString** |  |
**UsageEventCount** | **NullableInt32** |  |

## Methods

### NewInvoiceUsageExplainResponseLinesInner

`func NewInvoiceUsageExplainResponseLinesInner(id string, description NullableString, kind string, amountMinor NullableInt64, currency NullableString, meterId NullableString, usageQuantity NullableString, usageEventCount NullableInt32, ) *InvoiceUsageExplainResponseLinesInner`

NewInvoiceUsageExplainResponseLinesInner instantiates a new InvoiceUsageExplainResponseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUsageExplainResponseLinesInnerWithDefaults

`func NewInvoiceUsageExplainResponseLinesInnerWithDefaults() *InvoiceUsageExplainResponseLinesInner`

NewInvoiceUsageExplainResponseLinesInnerWithDefaults instantiates a new InvoiceUsageExplainResponseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceUsageExplainResponseLinesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceUsageExplainResponseLinesInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *InvoiceUsageExplainResponseLinesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceUsageExplainResponseLinesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKind

`func (o *InvoiceUsageExplainResponseLinesInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *InvoiceUsageExplainResponseLinesInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetAmountMinor

`func (o *InvoiceUsageExplainResponseLinesInner) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoiceUsageExplainResponseLinesInner) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### SetAmountMinorNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetAmountMinorNil(b bool)`

 SetAmountMinorNil sets the value for AmountMinor to be an explicit nil

### UnsetAmountMinor
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetAmountMinor()`

UnsetAmountMinor ensures that no value is present for AmountMinor, not even an explicit nil
### GetCurrency

`func (o *InvoiceUsageExplainResponseLinesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceUsageExplainResponseLinesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetMeterId

`func (o *InvoiceUsageExplainResponseLinesInner) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *InvoiceUsageExplainResponseLinesInner) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### SetMeterIdNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetMeterIdNil(b bool)`

 SetMeterIdNil sets the value for MeterId to be an explicit nil

### UnsetMeterId
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetMeterId()`

UnsetMeterId ensures that no value is present for MeterId, not even an explicit nil
### GetUsageQuantity

`func (o *InvoiceUsageExplainResponseLinesInner) GetUsageQuantity() string`

GetUsageQuantity returns the UsageQuantity field if non-nil, zero value otherwise.

### GetUsageQuantityOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetUsageQuantityOk() (*string, bool)`

GetUsageQuantityOk returns a tuple with the UsageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageQuantity

`func (o *InvoiceUsageExplainResponseLinesInner) SetUsageQuantity(v string)`

SetUsageQuantity sets UsageQuantity field to given value.


### SetUsageQuantityNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetUsageQuantityNil(b bool)`

 SetUsageQuantityNil sets the value for UsageQuantity to be an explicit nil

### UnsetUsageQuantity
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetUsageQuantity()`

UnsetUsageQuantity ensures that no value is present for UsageQuantity, not even an explicit nil
### GetUsageEventCount

`func (o *InvoiceUsageExplainResponseLinesInner) GetUsageEventCount() int32`

GetUsageEventCount returns the UsageEventCount field if non-nil, zero value otherwise.

### GetUsageEventCountOk

`func (o *InvoiceUsageExplainResponseLinesInner) GetUsageEventCountOk() (*int32, bool)`

GetUsageEventCountOk returns a tuple with the UsageEventCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageEventCount

`func (o *InvoiceUsageExplainResponseLinesInner) SetUsageEventCount(v int32)`

SetUsageEventCount sets UsageEventCount field to given value.


### SetUsageEventCountNil

`func (o *InvoiceUsageExplainResponseLinesInner) SetUsageEventCountNil(b bool)`

 SetUsageEventCountNil sets the value for UsageEventCount to be an explicit nil

### UnsetUsageEventCount
`func (o *InvoiceUsageExplainResponseLinesInner) UnsetUsageEventCount()`

UnsetUsageEventCount ensures that no value is present for UsageEventCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
