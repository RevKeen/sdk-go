# Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PublicId** | **string** |  | 
**Gateway** | **string** | Name of the payment processor that produced this payout batch | 
**GatewayPayoutId** | **string** |  | 
**GrossAmountMinor** | **float32** | Gross amount in cents | 
**FeesAmountMinor** | **float32** | Fees in cents | 
**NetAmountMinor** | **float32** | Net amount in cents | 
**Currency** | **string** |  | 
**ChargesCount** | **NullableFloat32** |  | 
**RefundsCount** | **NullableFloat32** |  | 
**ChargebacksCount** | **NullableFloat32** |  | 
**Status** | **string** |  | 
**ArrivalDate** | **NullableString** |  | 
**SettledAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 

## Methods

### NewPayout

`func NewPayout(id string, publicId string, gateway string, gatewayPayoutId string, grossAmountMinor float32, feesAmountMinor float32, netAmountMinor float32, currency string, chargesCount NullableFloat32, refundsCount NullableFloat32, chargebacksCount NullableFloat32, status string, arrivalDate NullableString, settledAt NullableString, createdAt string, ) *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payout) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *Payout) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *Payout) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *Payout) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetGateway

`func (o *Payout) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *Payout) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *Payout) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayPayoutId

`func (o *Payout) GetGatewayPayoutId() string`

GetGatewayPayoutId returns the GatewayPayoutId field if non-nil, zero value otherwise.

### GetGatewayPayoutIdOk

`func (o *Payout) GetGatewayPayoutIdOk() (*string, bool)`

GetGatewayPayoutIdOk returns a tuple with the GatewayPayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPayoutId

`func (o *Payout) SetGatewayPayoutId(v string)`

SetGatewayPayoutId sets GatewayPayoutId field to given value.


### GetGrossAmountMinor

`func (o *Payout) GetGrossAmountMinor() float32`

GetGrossAmountMinor returns the GrossAmountMinor field if non-nil, zero value otherwise.

### GetGrossAmountMinorOk

`func (o *Payout) GetGrossAmountMinorOk() (*float32, bool)`

GetGrossAmountMinorOk returns a tuple with the GrossAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmountMinor

`func (o *Payout) SetGrossAmountMinor(v float32)`

SetGrossAmountMinor sets GrossAmountMinor field to given value.


### GetFeesAmountMinor

`func (o *Payout) GetFeesAmountMinor() float32`

GetFeesAmountMinor returns the FeesAmountMinor field if non-nil, zero value otherwise.

### GetFeesAmountMinorOk

`func (o *Payout) GetFeesAmountMinorOk() (*float32, bool)`

GetFeesAmountMinorOk returns a tuple with the FeesAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmountMinor

`func (o *Payout) SetFeesAmountMinor(v float32)`

SetFeesAmountMinor sets FeesAmountMinor field to given value.


### GetNetAmountMinor

`func (o *Payout) GetNetAmountMinor() float32`

GetNetAmountMinor returns the NetAmountMinor field if non-nil, zero value otherwise.

### GetNetAmountMinorOk

`func (o *Payout) GetNetAmountMinorOk() (*float32, bool)`

GetNetAmountMinorOk returns a tuple with the NetAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmountMinor

`func (o *Payout) SetNetAmountMinor(v float32)`

SetNetAmountMinor sets NetAmountMinor field to given value.


### GetCurrency

`func (o *Payout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetChargesCount

`func (o *Payout) GetChargesCount() float32`

GetChargesCount returns the ChargesCount field if non-nil, zero value otherwise.

### GetChargesCountOk

`func (o *Payout) GetChargesCountOk() (*float32, bool)`

GetChargesCountOk returns a tuple with the ChargesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargesCount

`func (o *Payout) SetChargesCount(v float32)`

SetChargesCount sets ChargesCount field to given value.


### SetChargesCountNil

`func (o *Payout) SetChargesCountNil(b bool)`

 SetChargesCountNil sets the value for ChargesCount to be an explicit nil

### UnsetChargesCount
`func (o *Payout) UnsetChargesCount()`

UnsetChargesCount ensures that no value is present for ChargesCount, not even an explicit nil
### GetRefundsCount

`func (o *Payout) GetRefundsCount() float32`

GetRefundsCount returns the RefundsCount field if non-nil, zero value otherwise.

### GetRefundsCountOk

`func (o *Payout) GetRefundsCountOk() (*float32, bool)`

GetRefundsCountOk returns a tuple with the RefundsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundsCount

`func (o *Payout) SetRefundsCount(v float32)`

SetRefundsCount sets RefundsCount field to given value.


### SetRefundsCountNil

`func (o *Payout) SetRefundsCountNil(b bool)`

 SetRefundsCountNil sets the value for RefundsCount to be an explicit nil

### UnsetRefundsCount
`func (o *Payout) UnsetRefundsCount()`

UnsetRefundsCount ensures that no value is present for RefundsCount, not even an explicit nil
### GetChargebacksCount

`func (o *Payout) GetChargebacksCount() float32`

GetChargebacksCount returns the ChargebacksCount field if non-nil, zero value otherwise.

### GetChargebacksCountOk

`func (o *Payout) GetChargebacksCountOk() (*float32, bool)`

GetChargebacksCountOk returns a tuple with the ChargebacksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargebacksCount

`func (o *Payout) SetChargebacksCount(v float32)`

SetChargebacksCount sets ChargebacksCount field to given value.


### SetChargebacksCountNil

`func (o *Payout) SetChargebacksCountNil(b bool)`

 SetChargebacksCountNil sets the value for ChargebacksCount to be an explicit nil

### UnsetChargebacksCount
`func (o *Payout) UnsetChargebacksCount()`

UnsetChargebacksCount ensures that no value is present for ChargebacksCount, not even an explicit nil
### GetStatus

`func (o *Payout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payout) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetArrivalDate

`func (o *Payout) GetArrivalDate() string`

GetArrivalDate returns the ArrivalDate field if non-nil, zero value otherwise.

### GetArrivalDateOk

`func (o *Payout) GetArrivalDateOk() (*string, bool)`

GetArrivalDateOk returns a tuple with the ArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalDate

`func (o *Payout) SetArrivalDate(v string)`

SetArrivalDate sets ArrivalDate field to given value.


### SetArrivalDateNil

`func (o *Payout) SetArrivalDateNil(b bool)`

 SetArrivalDateNil sets the value for ArrivalDate to be an explicit nil

### UnsetArrivalDate
`func (o *Payout) UnsetArrivalDate()`

UnsetArrivalDate ensures that no value is present for ArrivalDate, not even an explicit nil
### GetSettledAt

`func (o *Payout) GetSettledAt() string`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *Payout) GetSettledAtOk() (*string, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *Payout) SetSettledAt(v string)`

SetSettledAt sets SettledAt field to given value.


### SetSettledAtNil

`func (o *Payout) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *Payout) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetCreatedAt

`func (o *Payout) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Payout) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Payout) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


