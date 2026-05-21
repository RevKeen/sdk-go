# CardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **NullableString** | The card brand | 
**Last4** | **NullableString** | The last 4 digits of the card number | 
**ExpMonth** | **NullableInt32** | Expiration month (1-12) | 
**ExpYear** | **NullableInt32** | Expiration year | 
**Funding** | **NullableString** | The card funding type | 

## Methods

### NewCardDetails

`func NewCardDetails(brand NullableString, last4 NullableString, expMonth NullableInt32, expYear NullableInt32, funding NullableString, ) *CardDetails`

NewCardDetails instantiates a new CardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardDetailsWithDefaults

`func NewCardDetailsWithDefaults() *CardDetails`

NewCardDetailsWithDefaults instantiates a new CardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CardDetails) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CardDetails) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CardDetails) SetBrand(v string)`

SetBrand sets Brand field to given value.


### SetBrandNil

`func (o *CardDetails) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CardDetails) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetLast4

`func (o *CardDetails) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *CardDetails) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *CardDetails) SetLast4(v string)`

SetLast4 sets Last4 field to given value.


### SetLast4Nil

`func (o *CardDetails) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *CardDetails) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetExpMonth

`func (o *CardDetails) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *CardDetails) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *CardDetails) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.


### SetExpMonthNil

`func (o *CardDetails) SetExpMonthNil(b bool)`

 SetExpMonthNil sets the value for ExpMonth to be an explicit nil

### UnsetExpMonth
`func (o *CardDetails) UnsetExpMonth()`

UnsetExpMonth ensures that no value is present for ExpMonth, not even an explicit nil
### GetExpYear

`func (o *CardDetails) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *CardDetails) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *CardDetails) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.


### SetExpYearNil

`func (o *CardDetails) SetExpYearNil(b bool)`

 SetExpYearNil sets the value for ExpYear to be an explicit nil

### UnsetExpYear
`func (o *CardDetails) UnsetExpYear()`

UnsetExpYear ensures that no value is present for ExpYear, not even an explicit nil
### GetFunding

`func (o *CardDetails) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *CardDetails) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *CardDetails) SetFunding(v string)`

SetFunding sets Funding field to given value.


### SetFundingNil

`func (o *CardDetails) SetFundingNil(b bool)`

 SetFundingNil sets the value for Funding to be an explicit nil

### UnsetFunding
`func (o *CardDetails) UnsetFunding()`

UnsetFunding ensures that no value is present for Funding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


