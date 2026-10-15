# SubscriptionPlanChangeErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  |
**Code** | **string** | &#x60;parameter_invalid&#x60;: the body failed validation. &#x60;invalid_status&#x60;: only active, trialing or past_due subscriptions can change plan. &#x60;resource_missing&#x60;: &#x60;newPlanId&#x60; names no price or product of yours. &#x60;price_missing&#x60;: the subscription has no current price. &#x60;price_not_recurring&#x60;, &#x60;interval_mismatch&#x60;, &#x60;currency_mismatch&#x60;: the new price is not a recurring price on the subscription&#39;s interval and currency. &#x60;price_amount_missing&#x60;: the new price has no amount above zero. |
**Message** | **string** |  |
**Param** | Pointer to **string** |  | [optional]
**RequestId** | Pointer to **string** |  | [optional]
**Details** | Pointer to **map[string]interface{}** | Extra detail for the refusal, for example the two intervals. | [optional]

## Methods

### NewSubscriptionPlanChangeErrorError

`func NewSubscriptionPlanChangeErrorError(type_ string, code string, message string, ) *SubscriptionPlanChangeErrorError`

NewSubscriptionPlanChangeErrorError instantiates a new SubscriptionPlanChangeErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanChangeErrorErrorWithDefaults

`func NewSubscriptionPlanChangeErrorErrorWithDefaults() *SubscriptionPlanChangeErrorError`

NewSubscriptionPlanChangeErrorErrorWithDefaults instantiates a new SubscriptionPlanChangeErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionPlanChangeErrorError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionPlanChangeErrorError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionPlanChangeErrorError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *SubscriptionPlanChangeErrorError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionPlanChangeErrorError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionPlanChangeErrorError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *SubscriptionPlanChangeErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionPlanChangeErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionPlanChangeErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *SubscriptionPlanChangeErrorError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *SubscriptionPlanChangeErrorError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *SubscriptionPlanChangeErrorError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *SubscriptionPlanChangeErrorError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetRequestId

`func (o *SubscriptionPlanChangeErrorError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SubscriptionPlanChangeErrorError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SubscriptionPlanChangeErrorError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SubscriptionPlanChangeErrorError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDetails

`func (o *SubscriptionPlanChangeErrorError) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SubscriptionPlanChangeErrorError) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SubscriptionPlanChangeErrorError) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SubscriptionPlanChangeErrorError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
