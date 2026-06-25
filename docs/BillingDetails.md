# BillingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | Full name on the payment method | 
**Email** | **NullableString** | Email address | 
**Phone** | **NullableString** | Phone number | 
**Address** | [**BillingDetailsAddress**](BillingDetailsAddress.md) |  | 

## Methods

### NewBillingDetails

`func NewBillingDetails(name NullableString, email NullableString, phone NullableString, address BillingDetailsAddress, ) *BillingDetails`

NewBillingDetails instantiates a new BillingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDetailsWithDefaults

`func NewBillingDetailsWithDefaults() *BillingDetails`

NewBillingDetailsWithDefaults instantiates a new BillingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillingDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingDetails) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BillingDetails) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BillingDetails) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *BillingDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillingDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillingDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *BillingDetails) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *BillingDetails) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *BillingDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BillingDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BillingDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *BillingDetails) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *BillingDetails) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAddress

`func (o *BillingDetails) GetAddress() BillingDetailsAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillingDetails) GetAddressOk() (*BillingDetailsAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillingDetails) SetAddress(v BillingDetailsAddress)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


