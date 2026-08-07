# SetCartContactInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | Customer email for abandoned-cart recovery. Pass &#x60;null&#x60; to clear. | [optional] 
**PromotionalConsent** | Pointer to **NullableBool** | Whether the customer consented to marketing/recovery email. Records the consent timestamp when set. | [optional] 
**SmsConsent** | Pointer to **NullableBool** | Whether the customer consented to marketing/recovery SMS. Records the consent timestamp when set. | [optional] 

## Methods

### NewSetCartContactInput

`func NewSetCartContactInput() *SetCartContactInput`

NewSetCartContactInput instantiates a new SetCartContactInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCartContactInputWithDefaults

`func NewSetCartContactInputWithDefaults() *SetCartContactInput`

NewSetCartContactInputWithDefaults instantiates a new SetCartContactInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SetCartContactInput) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SetCartContactInput) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SetCartContactInput) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SetCartContactInput) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SetCartContactInput) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SetCartContactInput) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPromotionalConsent

`func (o *SetCartContactInput) GetPromotionalConsent() bool`

GetPromotionalConsent returns the PromotionalConsent field if non-nil, zero value otherwise.

### GetPromotionalConsentOk

`func (o *SetCartContactInput) GetPromotionalConsentOk() (*bool, bool)`

GetPromotionalConsentOk returns a tuple with the PromotionalConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionalConsent

`func (o *SetCartContactInput) SetPromotionalConsent(v bool)`

SetPromotionalConsent sets PromotionalConsent field to given value.

### HasPromotionalConsent

`func (o *SetCartContactInput) HasPromotionalConsent() bool`

HasPromotionalConsent returns a boolean if a field has been set.

### SetPromotionalConsentNil

`func (o *SetCartContactInput) SetPromotionalConsentNil(b bool)`

 SetPromotionalConsentNil sets the value for PromotionalConsent to be an explicit nil

### UnsetPromotionalConsent
`func (o *SetCartContactInput) UnsetPromotionalConsent()`

UnsetPromotionalConsent ensures that no value is present for PromotionalConsent, not even an explicit nil
### GetSmsConsent

`func (o *SetCartContactInput) GetSmsConsent() bool`

GetSmsConsent returns the SmsConsent field if non-nil, zero value otherwise.

### GetSmsConsentOk

`func (o *SetCartContactInput) GetSmsConsentOk() (*bool, bool)`

GetSmsConsentOk returns a tuple with the SmsConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsConsent

`func (o *SetCartContactInput) SetSmsConsent(v bool)`

SetSmsConsent sets SmsConsent field to given value.

### HasSmsConsent

`func (o *SetCartContactInput) HasSmsConsent() bool`

HasSmsConsent returns a boolean if a field has been set.

### SetSmsConsentNil

`func (o *SetCartContactInput) SetSmsConsentNil(b bool)`

 SetSmsConsentNil sets the value for SmsConsent to be an explicit nil

### UnsetSmsConsent
`func (o *SetCartContactInput) UnsetSmsConsent()`

UnsetSmsConsent ensures that no value is present for SmsConsent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


