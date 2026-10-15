# SubscriptionDeleteConflictError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  |
**Message** | **string** |  |
**Details** | [**SubscriptionDeleteConflictErrorDetails**](SubscriptionDeleteConflictErrorDetails.md) |  |

## Methods

### NewSubscriptionDeleteConflictError

`func NewSubscriptionDeleteConflictError(code string, message string, details SubscriptionDeleteConflictErrorDetails, ) *SubscriptionDeleteConflictError`

NewSubscriptionDeleteConflictError instantiates a new SubscriptionDeleteConflictError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDeleteConflictErrorWithDefaults

`func NewSubscriptionDeleteConflictErrorWithDefaults() *SubscriptionDeleteConflictError`

NewSubscriptionDeleteConflictErrorWithDefaults instantiates a new SubscriptionDeleteConflictError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *SubscriptionDeleteConflictError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubscriptionDeleteConflictError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubscriptionDeleteConflictError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *SubscriptionDeleteConflictError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionDeleteConflictError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionDeleteConflictError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *SubscriptionDeleteConflictError) GetDetails() SubscriptionDeleteConflictErrorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SubscriptionDeleteConflictError) GetDetailsOk() (*SubscriptionDeleteConflictErrorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SubscriptionDeleteConflictError) SetDetails(v SubscriptionDeleteConflictErrorDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
