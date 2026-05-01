# TerminalDeviceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TerminalDevice**](TerminalDevice.md) |  | 
**Meta** | [**TerminalDeviceListResponseMeta**](TerminalDeviceListResponseMeta.md) |  | 

## Methods

### NewTerminalDeviceListResponse

`func NewTerminalDeviceListResponse(data []TerminalDevice, meta TerminalDeviceListResponseMeta, ) *TerminalDeviceListResponse`

NewTerminalDeviceListResponse instantiates a new TerminalDeviceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalDeviceListResponseWithDefaults

`func NewTerminalDeviceListResponseWithDefaults() *TerminalDeviceListResponse`

NewTerminalDeviceListResponseWithDefaults instantiates a new TerminalDeviceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TerminalDeviceListResponse) GetData() []TerminalDevice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TerminalDeviceListResponse) GetDataOk() (*[]TerminalDevice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TerminalDeviceListResponse) SetData(v []TerminalDevice)`

SetData sets Data field to given value.


### GetMeta

`func (o *TerminalDeviceListResponse) GetMeta() TerminalDeviceListResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TerminalDeviceListResponse) GetMetaOk() (*TerminalDeviceListResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TerminalDeviceListResponse) SetMeta(v TerminalDeviceListResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


