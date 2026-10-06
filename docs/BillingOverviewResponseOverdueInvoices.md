# BillingOverviewResponseOverdueInvoices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  |
**Totals** | [**[]BillingOverviewResponseOverdueInvoicesTotalsInner**](BillingOverviewResponseOverdueInvoicesTotalsInner.md) |  |
**Items** | **[]map[string]*interface{}** |  |

## Methods

### NewBillingOverviewResponseOverdueInvoices

`func NewBillingOverviewResponseOverdueInvoices(count int32, totals []BillingOverviewResponseOverdueInvoicesTotalsInner, items []map[string]*interface{}, ) *BillingOverviewResponseOverdueInvoices`

NewBillingOverviewResponseOverdueInvoices instantiates a new BillingOverviewResponseOverdueInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingOverviewResponseOverdueInvoicesWithDefaults

`func NewBillingOverviewResponseOverdueInvoicesWithDefaults() *BillingOverviewResponseOverdueInvoices`

NewBillingOverviewResponseOverdueInvoicesWithDefaults instantiates a new BillingOverviewResponseOverdueInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BillingOverviewResponseOverdueInvoices) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BillingOverviewResponseOverdueInvoices) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BillingOverviewResponseOverdueInvoices) SetCount(v int32)`

SetCount sets Count field to given value.


### GetTotals

`func (o *BillingOverviewResponseOverdueInvoices) GetTotals() []BillingOverviewResponseOverdueInvoicesTotalsInner`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *BillingOverviewResponseOverdueInvoices) GetTotalsOk() (*[]BillingOverviewResponseOverdueInvoicesTotalsInner, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *BillingOverviewResponseOverdueInvoices) SetTotals(v []BillingOverviewResponseOverdueInvoicesTotalsInner)`

SetTotals sets Totals field to given value.


### GetItems

`func (o *BillingOverviewResponseOverdueInvoices) GetItems() []map[string]*interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BillingOverviewResponseOverdueInvoices) GetItemsOk() (*[]map[string]*interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BillingOverviewResponseOverdueInvoices) SetItems(v []map[string]*interface{})`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
