# ApplyCartDiscountCodeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | Discount code to apply. Pass &#x60;null&#x60; to clear. Valid discounts are priced into cart totals before checkout handoff. | 

## Methods

### NewApplyCartDiscountCodeInput

`func NewApplyCartDiscountCodeInput(code NullableString, ) *ApplyCartDiscountCodeInput`

NewApplyCartDiscountCodeInput instantiates a new ApplyCartDiscountCodeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyCartDiscountCodeInputWithDefaults

`func NewApplyCartDiscountCodeInputWithDefaults() *ApplyCartDiscountCodeInput`

NewApplyCartDiscountCodeInputWithDefaults instantiates a new ApplyCartDiscountCodeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApplyCartDiscountCodeInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApplyCartDiscountCodeInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApplyCartDiscountCodeInput) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ApplyCartDiscountCodeInput) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ApplyCartDiscountCodeInput) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


