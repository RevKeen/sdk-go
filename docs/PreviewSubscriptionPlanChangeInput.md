# PreviewSubscriptionPlanChangeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPlanId** | **string** | The price to move to. A product id resolves to the product&#39;s default price. |
**ProrationBehavior** | Pointer to **string** | As for the change. &#x60;next_invoice&#x60; (default; &#x60;invoice_immediately&#x60; is applied the same way) prorates onto the next renewal invoice; &#x60;none&#x60; previews a change with no proration. | [optional]
**EffectiveDate** | Pointer to **time.Time** | Price the proration as of this date (ISO 8601). Defaults to now. | [optional]

## Methods

### NewPreviewSubscriptionPlanChangeInput

`func NewPreviewSubscriptionPlanChangeInput(newPlanId string, ) *PreviewSubscriptionPlanChangeInput`

NewPreviewSubscriptionPlanChangeInput instantiates a new PreviewSubscriptionPlanChangeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewSubscriptionPlanChangeInputWithDefaults

`func NewPreviewSubscriptionPlanChangeInputWithDefaults() *PreviewSubscriptionPlanChangeInput`

NewPreviewSubscriptionPlanChangeInputWithDefaults instantiates a new PreviewSubscriptionPlanChangeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPlanId

`func (o *PreviewSubscriptionPlanChangeInput) GetNewPlanId() string`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *PreviewSubscriptionPlanChangeInput) GetNewPlanIdOk() (*string, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *PreviewSubscriptionPlanChangeInput) SetNewPlanId(v string)`

SetNewPlanId sets NewPlanId field to given value.


### GetProrationBehavior

`func (o *PreviewSubscriptionPlanChangeInput) GetProrationBehavior() string`

GetProrationBehavior returns the ProrationBehavior field if non-nil, zero value otherwise.

### GetProrationBehaviorOk

`func (o *PreviewSubscriptionPlanChangeInput) GetProrationBehaviorOk() (*string, bool)`

GetProrationBehaviorOk returns a tuple with the ProrationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationBehavior

`func (o *PreviewSubscriptionPlanChangeInput) SetProrationBehavior(v string)`

SetProrationBehavior sets ProrationBehavior field to given value.

### HasProrationBehavior

`func (o *PreviewSubscriptionPlanChangeInput) HasProrationBehavior() bool`

HasProrationBehavior returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *PreviewSubscriptionPlanChangeInput) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *PreviewSubscriptionPlanChangeInput) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *PreviewSubscriptionPlanChangeInput) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *PreviewSubscriptionPlanChangeInput) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
