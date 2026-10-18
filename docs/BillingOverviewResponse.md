# BillingOverviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverdueInvoices** | [**BillingOverviewResponseOverdueInvoices**](BillingOverviewResponseOverdueInvoices.md) |  |
**Dunning** | [**BillingOverviewResponseDunning**](BillingOverviewResponseDunning.md) |  |
**Usage** | [**BillingOverviewResponseDunning**](BillingOverviewResponseDunning.md) |  |

## Methods

### NewBillingOverviewResponse

`func NewBillingOverviewResponse(overdueInvoices BillingOverviewResponseOverdueInvoices, dunning BillingOverviewResponseDunning, usage BillingOverviewResponseDunning, ) *BillingOverviewResponse`

NewBillingOverviewResponse instantiates a new BillingOverviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingOverviewResponseWithDefaults

`func NewBillingOverviewResponseWithDefaults() *BillingOverviewResponse`

NewBillingOverviewResponseWithDefaults instantiates a new BillingOverviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverdueInvoices

`func (o *BillingOverviewResponse) GetOverdueInvoices() BillingOverviewResponseOverdueInvoices`

GetOverdueInvoices returns the OverdueInvoices field if non-nil, zero value otherwise.

### GetOverdueInvoicesOk

`func (o *BillingOverviewResponse) GetOverdueInvoicesOk() (*BillingOverviewResponseOverdueInvoices, bool)`

GetOverdueInvoicesOk returns a tuple with the OverdueInvoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdueInvoices

`func (o *BillingOverviewResponse) SetOverdueInvoices(v BillingOverviewResponseOverdueInvoices)`

SetOverdueInvoices sets OverdueInvoices field to given value.


### GetDunning

`func (o *BillingOverviewResponse) GetDunning() BillingOverviewResponseDunning`

GetDunning returns the Dunning field if non-nil, zero value otherwise.

### GetDunningOk

`func (o *BillingOverviewResponse) GetDunningOk() (*BillingOverviewResponseDunning, bool)`

GetDunningOk returns a tuple with the Dunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunning

`func (o *BillingOverviewResponse) SetDunning(v BillingOverviewResponseDunning)`

SetDunning sets Dunning field to given value.


### GetUsage

`func (o *BillingOverviewResponse) GetUsage() BillingOverviewResponseDunning`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BillingOverviewResponse) GetUsageOk() (*BillingOverviewResponseDunning, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BillingOverviewResponse) SetUsage(v BillingOverviewResponseDunning)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
