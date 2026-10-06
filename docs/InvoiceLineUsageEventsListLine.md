# InvoiceLineUsageEventsListLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**InvoiceId** | **string** |  |
**MeterId** | **NullableString** |  |
**UsageRecordId** | **NullableString** |  |
**UnitAmountMinor** | **int64** |  |
**TotalAmountMinor** | **int64** |  |
**Currency** | **string** |  |
**UsageQuantity** | **NullableString** |  |
**BilledAmountMinorSum** | **int32** |  |

## Methods

### NewInvoiceLineUsageEventsListLine

`func NewInvoiceLineUsageEventsListLine(id string, invoiceId string, meterId NullableString, usageRecordId NullableString, unitAmountMinor int64, totalAmountMinor int64, currency string, usageQuantity NullableString, billedAmountMinorSum int32, ) *InvoiceLineUsageEventsListLine`

NewInvoiceLineUsageEventsListLine instantiates a new InvoiceLineUsageEventsListLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceLineUsageEventsListLineWithDefaults

`func NewInvoiceLineUsageEventsListLineWithDefaults() *InvoiceLineUsageEventsListLine`

NewInvoiceLineUsageEventsListLineWithDefaults instantiates a new InvoiceLineUsageEventsListLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceLineUsageEventsListLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceLineUsageEventsListLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceLineUsageEventsListLine) SetId(v string)`

SetId sets Id field to given value.


### GetInvoiceId

`func (o *InvoiceLineUsageEventsListLine) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *InvoiceLineUsageEventsListLine) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *InvoiceLineUsageEventsListLine) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetMeterId

`func (o *InvoiceLineUsageEventsListLine) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *InvoiceLineUsageEventsListLine) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *InvoiceLineUsageEventsListLine) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### SetMeterIdNil

`func (o *InvoiceLineUsageEventsListLine) SetMeterIdNil(b bool)`

 SetMeterIdNil sets the value for MeterId to be an explicit nil

### UnsetMeterId
`func (o *InvoiceLineUsageEventsListLine) UnsetMeterId()`

UnsetMeterId ensures that no value is present for MeterId, not even an explicit nil
### GetUsageRecordId

`func (o *InvoiceLineUsageEventsListLine) GetUsageRecordId() string`

GetUsageRecordId returns the UsageRecordId field if non-nil, zero value otherwise.

### GetUsageRecordIdOk

`func (o *InvoiceLineUsageEventsListLine) GetUsageRecordIdOk() (*string, bool)`

GetUsageRecordIdOk returns a tuple with the UsageRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordId

`func (o *InvoiceLineUsageEventsListLine) SetUsageRecordId(v string)`

SetUsageRecordId sets UsageRecordId field to given value.


### SetUsageRecordIdNil

`func (o *InvoiceLineUsageEventsListLine) SetUsageRecordIdNil(b bool)`

 SetUsageRecordIdNil sets the value for UsageRecordId to be an explicit nil

### UnsetUsageRecordId
`func (o *InvoiceLineUsageEventsListLine) UnsetUsageRecordId()`

UnsetUsageRecordId ensures that no value is present for UsageRecordId, not even an explicit nil
### GetUnitAmountMinor

`func (o *InvoiceLineUsageEventsListLine) GetUnitAmountMinor() int64`

GetUnitAmountMinor returns the UnitAmountMinor field if non-nil, zero value otherwise.

### GetUnitAmountMinorOk

`func (o *InvoiceLineUsageEventsListLine) GetUnitAmountMinorOk() (*int64, bool)`

GetUnitAmountMinorOk returns a tuple with the UnitAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountMinor

`func (o *InvoiceLineUsageEventsListLine) SetUnitAmountMinor(v int64)`

SetUnitAmountMinor sets UnitAmountMinor field to given value.


### GetTotalAmountMinor

`func (o *InvoiceLineUsageEventsListLine) GetTotalAmountMinor() int64`

GetTotalAmountMinor returns the TotalAmountMinor field if non-nil, zero value otherwise.

### GetTotalAmountMinorOk

`func (o *InvoiceLineUsageEventsListLine) GetTotalAmountMinorOk() (*int64, bool)`

GetTotalAmountMinorOk returns a tuple with the TotalAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountMinor

`func (o *InvoiceLineUsageEventsListLine) SetTotalAmountMinor(v int64)`

SetTotalAmountMinor sets TotalAmountMinor field to given value.


### GetCurrency

`func (o *InvoiceLineUsageEventsListLine) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceLineUsageEventsListLine) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceLineUsageEventsListLine) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetUsageQuantity

`func (o *InvoiceLineUsageEventsListLine) GetUsageQuantity() string`

GetUsageQuantity returns the UsageQuantity field if non-nil, zero value otherwise.

### GetUsageQuantityOk

`func (o *InvoiceLineUsageEventsListLine) GetUsageQuantityOk() (*string, bool)`

GetUsageQuantityOk returns a tuple with the UsageQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageQuantity

`func (o *InvoiceLineUsageEventsListLine) SetUsageQuantity(v string)`

SetUsageQuantity sets UsageQuantity field to given value.


### SetUsageQuantityNil

`func (o *InvoiceLineUsageEventsListLine) SetUsageQuantityNil(b bool)`

 SetUsageQuantityNil sets the value for UsageQuantity to be an explicit nil

### UnsetUsageQuantity
`func (o *InvoiceLineUsageEventsListLine) UnsetUsageQuantity()`

UnsetUsageQuantity ensures that no value is present for UsageQuantity, not even an explicit nil
### GetBilledAmountMinorSum

`func (o *InvoiceLineUsageEventsListLine) GetBilledAmountMinorSum() int32`

GetBilledAmountMinorSum returns the BilledAmountMinorSum field if non-nil, zero value otherwise.

### GetBilledAmountMinorSumOk

`func (o *InvoiceLineUsageEventsListLine) GetBilledAmountMinorSumOk() (*int32, bool)`

GetBilledAmountMinorSumOk returns a tuple with the BilledAmountMinorSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledAmountMinorSum

`func (o *InvoiceLineUsageEventsListLine) SetBilledAmountMinorSum(v int32)`

SetBilledAmountMinorSum sets BilledAmountMinorSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
