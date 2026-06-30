# InvoicesMarginEstimateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMinor** | Pointer to **int32** | Optional amount in minor units. Defaults to the invoice total. | [optional] 
**Rails** | Pointer to **[]string** | Rails to estimate. Defaults to invoice allowed_methods or all rails. | [optional] 
**CardScheme** | Pointer to **string** | Single card scheme to estimate for card rail, e.g. visa or mastercard. | [optional] 
**CardSchemes** | Pointer to **[]string** | Card schemes to estimate for card rail. | [optional] 

## Methods

### NewInvoicesMarginEstimateRequest

`func NewInvoicesMarginEstimateRequest() *InvoicesMarginEstimateRequest`

NewInvoicesMarginEstimateRequest instantiates a new InvoicesMarginEstimateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicesMarginEstimateRequestWithDefaults

`func NewInvoicesMarginEstimateRequestWithDefaults() *InvoicesMarginEstimateRequest`

NewInvoicesMarginEstimateRequestWithDefaults instantiates a new InvoicesMarginEstimateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountMinor

`func (o *InvoicesMarginEstimateRequest) GetAmountMinor() int32`

GetAmountMinor returns the AmountMinor field if non-nil, zero value otherwise.

### GetAmountMinorOk

`func (o *InvoicesMarginEstimateRequest) GetAmountMinorOk() (*int32, bool)`

GetAmountMinorOk returns a tuple with the AmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMinor

`func (o *InvoicesMarginEstimateRequest) SetAmountMinor(v int32)`

SetAmountMinor sets AmountMinor field to given value.

### HasAmountMinor

`func (o *InvoicesMarginEstimateRequest) HasAmountMinor() bool`

HasAmountMinor returns a boolean if a field has been set.

### GetRails

`func (o *InvoicesMarginEstimateRequest) GetRails() []string`

GetRails returns the Rails field if non-nil, zero value otherwise.

### GetRailsOk

`func (o *InvoicesMarginEstimateRequest) GetRailsOk() (*[]string, bool)`

GetRailsOk returns a tuple with the Rails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRails

`func (o *InvoicesMarginEstimateRequest) SetRails(v []string)`

SetRails sets Rails field to given value.

### HasRails

`func (o *InvoicesMarginEstimateRequest) HasRails() bool`

HasRails returns a boolean if a field has been set.

### GetCardScheme

`func (o *InvoicesMarginEstimateRequest) GetCardScheme() string`

GetCardScheme returns the CardScheme field if non-nil, zero value otherwise.

### GetCardSchemeOk

`func (o *InvoicesMarginEstimateRequest) GetCardSchemeOk() (*string, bool)`

GetCardSchemeOk returns a tuple with the CardScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardScheme

`func (o *InvoicesMarginEstimateRequest) SetCardScheme(v string)`

SetCardScheme sets CardScheme field to given value.

### HasCardScheme

`func (o *InvoicesMarginEstimateRequest) HasCardScheme() bool`

HasCardScheme returns a boolean if a field has been set.

### GetCardSchemes

`func (o *InvoicesMarginEstimateRequest) GetCardSchemes() []string`

GetCardSchemes returns the CardSchemes field if non-nil, zero value otherwise.

### GetCardSchemesOk

`func (o *InvoicesMarginEstimateRequest) GetCardSchemesOk() (*[]string, bool)`

GetCardSchemesOk returns a tuple with the CardSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSchemes

`func (o *InvoicesMarginEstimateRequest) SetCardSchemes(v []string)`

SetCardSchemes sets CardSchemes field to given value.

### HasCardSchemes

`func (o *InvoicesMarginEstimateRequest) HasCardSchemes() bool`

HasCardSchemes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


