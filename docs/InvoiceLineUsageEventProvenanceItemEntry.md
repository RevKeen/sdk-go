# InvoiceLineUsageEventProvenanceItemEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**AmountMinor** | **int64** |  |
**PreciseAmountMinor** | **NullableString** | Unrounded amount in minor units as a decimal string (REV-6211). Not integer cents; sub-penny rates stay exact. Null means not computable for this pricing scheme; never coerce to 0. |
**Quantity** | **string** |  |
**Currency** | **NullableString** |  |
**RatingVersion** | **int32** |  |
**PriceId** | **string** |  |
**PriceTermsHash** | **NullableString** |  |

## Methods

### NewInvoiceLineUsageEventProvenanceItemEntry

`func NewInvoiceLineUsageEventProvenanceItemEntry(id string, amountMinor int64, preciseAmountMinor NullableString, quantity string, currency NullableString, ratingVersion int32, priceId string, priceTermsHash NullableString, ) *InvoiceLineUsageEventProvenanceItemEntry`

NewInvoiceLineUsageEventProvenanceItemEntry instantiates a new InvoiceLineUsageEventProvenanceItemEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageEventProvenanceItemEntryWithDefaults

`func NewInvoiceLineUsageEventProvenanceItemEntryWithDefaults() *InvoiceLineUsageEventProvenanceItemEntry`

NewInvoiceLineUsageEventProvenanceItemEntryWithDefaults instantiates a new InvoiceLineUsageEventProvenanceItemEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetId(v string)`

SetId sets Id field to given value.


### GetAmountMinor

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetAmountMinor() int64`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetAmountMinorOk() (*int64, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetAmountMinor(v int64)`

SetAmountMinor sets AmountMinor field to given value.


### GetPreciseAmountMinor

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPreciseAmountMinor() string`

GetPreciseAmountMinor returns the PreciseAmountMinor field if non-nil, zero value otherwise.

### GetPreciseAmountMinorOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPreciseAmountMinorOk() (*string, bool)`

GetPreciseAmountMinorOk returns a tuple with the PreciseAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreciseAmountMinor

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetPreciseAmountMinor(v string)`

SetPreciseAmountMinor sets PreciseAmountMinor field to given value.


### SetPreciseAmountMinorNil

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetPreciseAmountMinorNil(b bool)`

 SetPreciseAmountMinorNil sets the value for PreciseAmountMinor to be an explicit nil

### UnsetPreciseAmountMinor
`func (o *InvoiceLineUsageEventProvenanceItemEntry) UnsetPreciseAmountMinor()`

UnsetPreciseAmountMinor ensures that no value is present for PreciseAmountMinor, not even an explicit nil
### GetQuantity

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetCurrency

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InvoiceLineUsageEventProvenanceItemEntry) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetRatingVersion

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetRatingVersion() int32`

GetRatingVersion returns the RatingVersion field if non-nil, zero value otherwise.

### GetRatingVersionOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetRatingVersionOk() (*int32, bool)`

GetRatingVersionOk returns a tuple with the RatingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingVersion

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetRatingVersion(v int32)`

SetRatingVersion sets RatingVersion field to given value.


### GetPriceId

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPriceId() string`

GetPriceId returns the PriceId field if non-nil, zero value otherwise.

### GetPriceIdOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPriceIdOk() (*string, bool)`

GetPriceIdOk returns a tuple with the PriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceId

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetPriceId(v string)`

SetPriceId sets PriceId field to given value.


### GetPriceTermsHash

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPriceTermsHash() string`

GetPriceTermsHash returns the PriceTermsHash field if non-nil, zero value otherwise.

### GetPriceTermsHashOk

`func (o *InvoiceLineUsageEventProvenanceItemEntry) GetPriceTermsHashOk() (*string, bool)`

GetPriceTermsHashOk returns a tuple with the PriceTermsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTermsHash

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetPriceTermsHash(v string)`

SetPriceTermsHash sets PriceTermsHash field to given value.


### SetPriceTermsHashNil

`func (o *InvoiceLineUsageEventProvenanceItemEntry) SetPriceTermsHashNil(b bool)`

 SetPriceTermsHashNil sets the value for PriceTermsHash to be an explicit nil

### UnsetPriceTermsHash
`func (o *InvoiceLineUsageEventProvenanceItemEntry) UnsetPriceTermsHash()`

UnsetPriceTermsHash ensures that no value is present for PriceTermsHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
