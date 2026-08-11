# PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**UnitAmount** | **int32** |  | 
**Currency** | **string** |  | 
**Recurring** | Pointer to [**PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring**](PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring.md) |  | [optional] 

## Methods

### NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails

`func NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails(name string, unitAmount int32, currency string, ) *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails`

NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails instantiates a new PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsWithDefaults

`func NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsWithDefaults() *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails`

NewPaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsWithDefaults instantiates a new PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUnitAmount

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetCurrency

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetRecurring

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetRecurring() PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) GetRecurringOk() (*PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) SetRecurring(v PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetails) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


