# PreviewCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedCollectionDate** | Pointer to **string** | Preferred collection date (ISO date). Backend resolves the working-day chain. | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription the collection is previewed for, if any | [optional] 

## Methods

### NewPreviewCollectionRequest

`func NewPreviewCollectionRequest() *PreviewCollectionRequest`

NewPreviewCollectionRequest instantiates a new PreviewCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewCollectionRequestWithDefaults

`func NewPreviewCollectionRequestWithDefaults() *PreviewCollectionRequest`

NewPreviewCollectionRequestWithDefaults instantiates a new PreviewCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedCollectionDate

`func (o *PreviewCollectionRequest) GetRequestedCollectionDate() string`

GetRequestedCollectionDate returns the RequestedCollectionDate field if non-nil, zero value otherwise.

### GetRequestedCollectionDateOk

`func (o *PreviewCollectionRequest) GetRequestedCollectionDateOk() (*string, bool)`

GetRequestedCollectionDateOk returns a tuple with the RequestedCollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCollectionDate

`func (o *PreviewCollectionRequest) SetRequestedCollectionDate(v string)`

SetRequestedCollectionDate sets RequestedCollectionDate field to given value.

### HasRequestedCollectionDate

`func (o *PreviewCollectionRequest) HasRequestedCollectionDate() bool`

HasRequestedCollectionDate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PreviewCollectionRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PreviewCollectionRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PreviewCollectionRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *PreviewCollectionRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


