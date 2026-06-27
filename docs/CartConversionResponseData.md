# CartConversionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartSession** | [**CartSession**](CartSession.md) |  | 
**CheckoutSession** | [**CartCheckoutSession**](CartCheckoutSession.md) |  | 

## Methods

### NewCartConversionResponseData

`func NewCartConversionResponseData(cartSession CartSession, checkoutSession CartCheckoutSession, ) *CartConversionResponseData`

NewCartConversionResponseData instantiates a new CartConversionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCartConversionResponseDataWithDefaults

`func NewCartConversionResponseDataWithDefaults() *CartConversionResponseData`

NewCartConversionResponseDataWithDefaults instantiates a new CartConversionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartSession

`func (o *CartConversionResponseData) GetCartSession() CartSession`

GetCartSession returns the CartSession field if non-nil, zero value otherwise.

### GetCartSessionOk

`func (o *CartConversionResponseData) GetCartSessionOk() (*CartSession, bool)`

GetCartSessionOk returns a tuple with the CartSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartSession

`func (o *CartConversionResponseData) SetCartSession(v CartSession)`

SetCartSession sets CartSession field to given value.


### GetCheckoutSession

`func (o *CartConversionResponseData) GetCheckoutSession() CartCheckoutSession`

GetCheckoutSession returns the CheckoutSession field if non-nil, zero value otherwise.

### GetCheckoutSessionOk

`func (o *CartConversionResponseData) GetCheckoutSessionOk() (*CartCheckoutSession, bool)`

GetCheckoutSessionOk returns a tuple with the CheckoutSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSession

`func (o *CartConversionResponseData) SetCheckoutSession(v CartCheckoutSession)`

SetCheckoutSession sets CheckoutSession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


