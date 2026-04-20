# TerminalDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this terminal device. Use this as the device_id when initiating payments. | 
**DeviceName** | **NullableString** | Friendly name assigned to this device (e.g., &#39;Front Desk Terminal&#39;) | 
**TerminalSerial** | **NullableString** | Serial number of the paired PAX terminal hardware | 
**TerminalIp** | **NullableString** | The LAN IP address of the terminal as seen by the connector on the merchant&#39;s local network. This is not a publicly reachable address — all commands flow through the RevKeen cloud. | 
**Platform** | **NullableString** | Operating system of the machine running the Terminal Connector application | 
**Status** | **string** | Device connectivity status. online: connector connected and heartbeat within 5 minutes. offline: heartbeat stale or connector disconnected. pairing: connector registered but terminal not yet paired. | 
**TerminalPaired** | **bool** | Whether a PAX terminal has been paired to this connector device. Must be true to accept payments. | 
**TerminalReachable** | **NullableBool** | Whether the connector can currently reach the PAX terminal on the local network | 
**AppVersion** | **NullableString** | Version of the Terminal Connector application | 
**LastHeartbeatAt** | **NullableString** | ISO 8601 timestamp of the last heartbeat received from this device. Devices with no heartbeat in 5 minutes are considered offline. | 

## Methods

### NewTerminalDevice

`func NewTerminalDevice(id string, deviceName NullableString, terminalSerial NullableString, terminalIp NullableString, platform NullableString, status string, terminalPaired bool, terminalReachable NullableBool, appVersion NullableString, lastHeartbeatAt NullableString, ) *TerminalDevice`

NewTerminalDevice instantiates a new TerminalDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminalDeviceWithDefaults

`func NewTerminalDeviceWithDefaults() *TerminalDevice`

NewTerminalDeviceWithDefaults instantiates a new TerminalDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TerminalDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TerminalDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TerminalDevice) SetId(v string)`

SetId sets Id field to given value.


### GetDeviceName

`func (o *TerminalDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *TerminalDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *TerminalDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### SetDeviceNameNil

`func (o *TerminalDevice) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *TerminalDevice) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetTerminalSerial

`func (o *TerminalDevice) GetTerminalSerial() string`

GetTerminalSerial returns the TerminalSerial field if non-nil, zero value otherwise.

### GetTerminalSerialOk

`func (o *TerminalDevice) GetTerminalSerialOk() (*string, bool)`

GetTerminalSerialOk returns a tuple with the TerminalSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalSerial

`func (o *TerminalDevice) SetTerminalSerial(v string)`

SetTerminalSerial sets TerminalSerial field to given value.


### SetTerminalSerialNil

`func (o *TerminalDevice) SetTerminalSerialNil(b bool)`

 SetTerminalSerialNil sets the value for TerminalSerial to be an explicit nil

### UnsetTerminalSerial
`func (o *TerminalDevice) UnsetTerminalSerial()`

UnsetTerminalSerial ensures that no value is present for TerminalSerial, not even an explicit nil
### GetTerminalIp

`func (o *TerminalDevice) GetTerminalIp() string`

GetTerminalIp returns the TerminalIp field if non-nil, zero value otherwise.

### GetTerminalIpOk

`func (o *TerminalDevice) GetTerminalIpOk() (*string, bool)`

GetTerminalIpOk returns a tuple with the TerminalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalIp

`func (o *TerminalDevice) SetTerminalIp(v string)`

SetTerminalIp sets TerminalIp field to given value.


### SetTerminalIpNil

`func (o *TerminalDevice) SetTerminalIpNil(b bool)`

 SetTerminalIpNil sets the value for TerminalIp to be an explicit nil

### UnsetTerminalIp
`func (o *TerminalDevice) UnsetTerminalIp()`

UnsetTerminalIp ensures that no value is present for TerminalIp, not even an explicit nil
### GetPlatform

`func (o *TerminalDevice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *TerminalDevice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *TerminalDevice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### SetPlatformNil

`func (o *TerminalDevice) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *TerminalDevice) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetStatus

`func (o *TerminalDevice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TerminalDevice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TerminalDevice) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTerminalPaired

`func (o *TerminalDevice) GetTerminalPaired() bool`

GetTerminalPaired returns the TerminalPaired field if non-nil, zero value otherwise.

### GetTerminalPairedOk

`func (o *TerminalDevice) GetTerminalPairedOk() (*bool, bool)`

GetTerminalPairedOk returns a tuple with the TerminalPaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalPaired

`func (o *TerminalDevice) SetTerminalPaired(v bool)`

SetTerminalPaired sets TerminalPaired field to given value.


### GetTerminalReachable

`func (o *TerminalDevice) GetTerminalReachable() bool`

GetTerminalReachable returns the TerminalReachable field if non-nil, zero value otherwise.

### GetTerminalReachableOk

`func (o *TerminalDevice) GetTerminalReachableOk() (*bool, bool)`

GetTerminalReachableOk returns a tuple with the TerminalReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalReachable

`func (o *TerminalDevice) SetTerminalReachable(v bool)`

SetTerminalReachable sets TerminalReachable field to given value.


### SetTerminalReachableNil

`func (o *TerminalDevice) SetTerminalReachableNil(b bool)`

 SetTerminalReachableNil sets the value for TerminalReachable to be an explicit nil

### UnsetTerminalReachable
`func (o *TerminalDevice) UnsetTerminalReachable()`

UnsetTerminalReachable ensures that no value is present for TerminalReachable, not even an explicit nil
### GetAppVersion

`func (o *TerminalDevice) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *TerminalDevice) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *TerminalDevice) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.


### SetAppVersionNil

`func (o *TerminalDevice) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *TerminalDevice) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetLastHeartbeatAt

`func (o *TerminalDevice) GetLastHeartbeatAt() string`

GetLastHeartbeatAt returns the LastHeartbeatAt field if non-nil, zero value otherwise.

### GetLastHeartbeatAtOk

`func (o *TerminalDevice) GetLastHeartbeatAtOk() (*string, bool)`

GetLastHeartbeatAtOk returns a tuple with the LastHeartbeatAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatAt

`func (o *TerminalDevice) SetLastHeartbeatAt(v string)`

SetLastHeartbeatAt sets LastHeartbeatAt field to given value.


### SetLastHeartbeatAtNil

`func (o *TerminalDevice) SetLastHeartbeatAtNil(b bool)`

 SetLastHeartbeatAtNil sets the value for LastHeartbeatAt to be an explicit nil

### UnsetLastHeartbeatAt
`func (o *TerminalDevice) UnsetLastHeartbeatAt()`

UnsetLastHeartbeatAt ensures that no value is present for LastHeartbeatAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


