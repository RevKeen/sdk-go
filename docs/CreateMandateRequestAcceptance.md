# CreateMandateRequestAcceptance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedAt** | **time.Time** |  | 
**PayerAddress** | **map[string]interface{}** |  | 
**GuaranteeVersion** | **string** |  | 
**Declarations** | **map[string]bool** |  | 

## Methods

### NewCreateMandateRequestAcceptance

`func NewCreateMandateRequestAcceptance(acceptedAt time.Time, payerAddress map[string]interface{}, guaranteeVersion string, declarations map[string]bool, ) *CreateMandateRequestAcceptance`

NewCreateMandateRequestAcceptance instantiates a new CreateMandateRequestAcceptance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMandateRequestAcceptanceWithDefaults

`func NewCreateMandateRequestAcceptanceWithDefaults() *CreateMandateRequestAcceptance`

NewCreateMandateRequestAcceptanceWithDefaults instantiates a new CreateMandateRequestAcceptance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedAt

`func (o *CreateMandateRequestAcceptance) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *CreateMandateRequestAcceptance) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *CreateMandateRequestAcceptance) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.


### GetPayerAddress

`func (o *CreateMandateRequestAcceptance) GetPayerAddress() map[string]interface{}`

GetPayerAddress returns the PayerAddress field if non-nil, zero value otherwise.

### GetPayerAddressOk

`func (o *CreateMandateRequestAcceptance) GetPayerAddressOk() (*map[string]interface{}, bool)`

GetPayerAddressOk returns a tuple with the PayerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddress

`func (o *CreateMandateRequestAcceptance) SetPayerAddress(v map[string]interface{})`

SetPayerAddress sets PayerAddress field to given value.


### GetGuaranteeVersion

`func (o *CreateMandateRequestAcceptance) GetGuaranteeVersion() string`

GetGuaranteeVersion returns the GuaranteeVersion field if non-nil, zero value otherwise.

### GetGuaranteeVersionOk

`func (o *CreateMandateRequestAcceptance) GetGuaranteeVersionOk() (*string, bool)`

GetGuaranteeVersionOk returns a tuple with the GuaranteeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteeVersion

`func (o *CreateMandateRequestAcceptance) SetGuaranteeVersion(v string)`

SetGuaranteeVersion sets GuaranteeVersion field to given value.


### GetDeclarations

`func (o *CreateMandateRequestAcceptance) GetDeclarations() map[string]bool`

GetDeclarations returns the Declarations field if non-nil, zero value otherwise.

### GetDeclarationsOk

`func (o *CreateMandateRequestAcceptance) GetDeclarationsOk() (*map[string]bool, bool)`

GetDeclarationsOk returns a tuple with the Declarations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarations

`func (o *CreateMandateRequestAcceptance) SetDeclarations(v map[string]bool)`

SetDeclarations sets Declarations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


