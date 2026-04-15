# CreatePaymentLinkRequestLineItemsInnerPriceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**UnitAmount** | **int32** |  | 
**ProductData** | Pointer to [**CreatePaymentLinkRequestLineItemsInnerPriceDataProductData**](CreatePaymentLinkRequestLineItemsInnerPriceDataProductData.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Recurring** | Pointer to [**ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetailsRecurring**](ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetailsRecurring.md) |  | [optional] 
**TaxBehavior** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreatePaymentLinkRequestLineItemsInnerPriceData

`func NewCreatePaymentLinkRequestLineItemsInnerPriceData(unitAmount int32, ) *CreatePaymentLinkRequestLineItemsInnerPriceData`

NewCreatePaymentLinkRequestLineItemsInnerPriceData instantiates a new CreatePaymentLinkRequestLineItemsInnerPriceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentLinkRequestLineItemsInnerPriceDataWithDefaults

`func NewCreatePaymentLinkRequestLineItemsInnerPriceDataWithDefaults() *CreatePaymentLinkRequestLineItemsInnerPriceData`

NewCreatePaymentLinkRequestLineItemsInnerPriceDataWithDefaults instantiates a new CreatePaymentLinkRequestLineItemsInnerPriceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnitAmount

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetProductData

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetProductData() CreatePaymentLinkRequestLineItemsInnerPriceDataProductData`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetProductDataOk() (*CreatePaymentLinkRequestLineItemsInnerPriceDataProductData, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetProductData(v CreatePaymentLinkRequestLineItemsInnerPriceDataProductData)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetProduct

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRecurring

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetRecurring() ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetailsRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetRecurringOk() (*ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetailsRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetRecurring(v ListPaymentLinks200ResponseDataInnerLineItemsInnerProductDetailsRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetTaxBehavior

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### SetTaxBehaviorNil

`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) SetTaxBehaviorNil(b bool)`

 SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil

### UnsetTaxBehavior
`func (o *CreatePaymentLinkRequestLineItemsInnerPriceData) UnsetTaxBehavior()`

UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


