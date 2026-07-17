# PaymentLinksCreateRequestLineItemsInnerPriceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**UnitAmount** | **int32** |  | 
**ProductData** | Pointer to [**PaymentLinksCreateRequestLineItemsInnerPriceDataProductData**](PaymentLinksCreateRequestLineItemsInnerPriceDataProductData.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Recurring** | Pointer to [**PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring**](PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring.md) |  | [optional] 
**TaxBehavior** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentLinksCreateRequestLineItemsInnerPriceData

`func NewPaymentLinksCreateRequestLineItemsInnerPriceData(unitAmount int32, ) *PaymentLinksCreateRequestLineItemsInnerPriceData`

NewPaymentLinksCreateRequestLineItemsInnerPriceData instantiates a new PaymentLinksCreateRequestLineItemsInnerPriceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentLinksCreateRequestLineItemsInnerPriceDataWithDefaults

`func NewPaymentLinksCreateRequestLineItemsInnerPriceDataWithDefaults() *PaymentLinksCreateRequestLineItemsInnerPriceData`

NewPaymentLinksCreateRequestLineItemsInnerPriceDataWithDefaults instantiates a new PaymentLinksCreateRequestLineItemsInnerPriceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnitAmount

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.


### GetProductData

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductData() PaymentLinksCreateRequestLineItemsInnerPriceDataProductData`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductDataOk() (*PaymentLinksCreateRequestLineItemsInnerPriceDataProductData, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetProductData(v PaymentLinksCreateRequestLineItemsInnerPriceDataProductData)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetProduct

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRecurring

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetRecurring() PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetRecurringOk() (*PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetRecurring(v PaymentLinksList200ResponseDataInnerLineItemsInnerProductDetailsRecurring)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetTaxBehavior

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetTaxBehavior() string`

GetTaxBehavior returns the TaxBehavior field if non-nil, zero value otherwise.

### GetTaxBehaviorOk

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) GetTaxBehaviorOk() (*string, bool)`

GetTaxBehaviorOk returns a tuple with the TaxBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBehavior

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetTaxBehavior(v string)`

SetTaxBehavior sets TaxBehavior field to given value.

### HasTaxBehavior

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) HasTaxBehavior() bool`

HasTaxBehavior returns a boolean if a field has been set.

### SetTaxBehaviorNil

`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) SetTaxBehaviorNil(b bool)`

 SetTaxBehaviorNil sets the value for TaxBehavior to be an explicit nil

### UnsetTaxBehavior
`func (o *PaymentLinksCreateRequestLineItemsInnerPriceData) UnsetTaxBehavior()`

UnsetTaxBehavior ensures that no value is present for TaxBehavior, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


