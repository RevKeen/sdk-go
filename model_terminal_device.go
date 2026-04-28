/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the TerminalDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerminalDevice{}

// TerminalDevice A POS terminal device registered to your merchant account, with connection status and capabilities.
type TerminalDevice struct {
	// Unique identifier for this terminal device. Use this as the device_id when initiating payments.
	Id string `json:"id"`
	// Friendly name assigned to this device (e.g., 'Front Desk Terminal')
	DeviceName NullableString `json:"device_name"`
	// Serial number of the paired PAX terminal hardware
	TerminalSerial NullableString `json:"terminal_serial"`
	// The LAN IP address of the terminal as seen by the connector on the merchant's local network. This is not a publicly reachable address — all commands flow through the RevKeen cloud.
	TerminalIp NullableString `json:"terminal_ip"`
	// Operating system of the machine running the Terminal Connector application
	Platform NullableString `json:"platform"`
	// Device connectivity status. online: connector connected and heartbeat within 5 minutes. offline: heartbeat stale or connector disconnected. pairing: connector registered but terminal not yet paired.
	Status string `json:"status"`
	// Whether a PAX terminal has been paired to this connector device. Must be true to accept payments.
	TerminalPaired bool `json:"terminal_paired"`
	// Whether the connector can currently reach the PAX terminal on the local network
	TerminalReachable NullableBool `json:"terminal_reachable"`
	// Version of the Terminal Connector application
	AppVersion NullableString `json:"app_version"`
	// ISO 8601 timestamp of the last heartbeat received from this device. Devices with no heartbeat in 5 minutes are considered offline.
	LastHeartbeatAt NullableString `json:"last_heartbeat_at"`
	AdditionalProperties map[string]interface{}
}

type _TerminalDevice TerminalDevice

// NewTerminalDevice instantiates a new TerminalDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalDevice(id string, deviceName NullableString, terminalSerial NullableString, terminalIp NullableString, platform NullableString, status string, terminalPaired bool, terminalReachable NullableBool, appVersion NullableString, lastHeartbeatAt NullableString) *TerminalDevice {
	this := TerminalDevice{}
	this.Id = id
	this.DeviceName = deviceName
	this.TerminalSerial = terminalSerial
	this.TerminalIp = terminalIp
	this.Platform = platform
	this.Status = status
	this.TerminalPaired = terminalPaired
	this.TerminalReachable = terminalReachable
	this.AppVersion = appVersion
	this.LastHeartbeatAt = lastHeartbeatAt
	return &this
}

// NewTerminalDeviceWithDefaults instantiates a new TerminalDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalDeviceWithDefaults() *TerminalDevice {
	this := TerminalDevice{}
	return &this
}

// GetId returns the Id field value
func (o *TerminalDevice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TerminalDevice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TerminalDevice) SetId(v string) {
	o.Id = v
}

// GetDeviceName returns the DeviceName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetDeviceName() string {
	if o == nil || o.DeviceName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// SetDeviceName sets field value
func (o *TerminalDevice) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}

// GetTerminalSerial returns the TerminalSerial field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetTerminalSerial() string {
	if o == nil || o.TerminalSerial.Get() == nil {
		var ret string
		return ret
	}

	return *o.TerminalSerial.Get()
}

// GetTerminalSerialOk returns a tuple with the TerminalSerial field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetTerminalSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminalSerial.Get(), o.TerminalSerial.IsSet()
}

// SetTerminalSerial sets field value
func (o *TerminalDevice) SetTerminalSerial(v string) {
	o.TerminalSerial.Set(&v)
}

// GetTerminalIp returns the TerminalIp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetTerminalIp() string {
	if o == nil || o.TerminalIp.Get() == nil {
		var ret string
		return ret
	}

	return *o.TerminalIp.Get()
}

// GetTerminalIpOk returns a tuple with the TerminalIp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetTerminalIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminalIp.Get(), o.TerminalIp.IsSet()
}

// SetTerminalIp sets field value
func (o *TerminalDevice) SetTerminalIp(v string) {
	o.TerminalIp.Set(&v)
}

// GetPlatform returns the Platform field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetPlatform() string {
	if o == nil || o.Platform.Get() == nil {
		var ret string
		return ret
	}

	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// SetPlatform sets field value
func (o *TerminalDevice) SetPlatform(v string) {
	o.Platform.Set(&v)
}

// GetStatus returns the Status field value
func (o *TerminalDevice) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TerminalDevice) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TerminalDevice) SetStatus(v string) {
	o.Status = v
}

// GetTerminalPaired returns the TerminalPaired field value
func (o *TerminalDevice) GetTerminalPaired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TerminalPaired
}

// GetTerminalPairedOk returns a tuple with the TerminalPaired field value
// and a boolean to check if the value has been set.
func (o *TerminalDevice) GetTerminalPairedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminalPaired, true
}

// SetTerminalPaired sets field value
func (o *TerminalDevice) SetTerminalPaired(v bool) {
	o.TerminalPaired = v
}

// GetTerminalReachable returns the TerminalReachable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *TerminalDevice) GetTerminalReachable() bool {
	if o == nil || o.TerminalReachable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.TerminalReachable.Get()
}

// GetTerminalReachableOk returns a tuple with the TerminalReachable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetTerminalReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminalReachable.Get(), o.TerminalReachable.IsSet()
}

// SetTerminalReachable sets field value
func (o *TerminalDevice) SetTerminalReachable(v bool) {
	o.TerminalReachable.Set(&v)
}

// GetAppVersion returns the AppVersion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetAppVersion() string {
	if o == nil || o.AppVersion.Get() == nil {
		var ret string
		return ret
	}

	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// SetAppVersion sets field value
func (o *TerminalDevice) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}

// GetLastHeartbeatAt returns the LastHeartbeatAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TerminalDevice) GetLastHeartbeatAt() string {
	if o == nil || o.LastHeartbeatAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastHeartbeatAt.Get()
}

// GetLastHeartbeatAtOk returns a tuple with the LastHeartbeatAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TerminalDevice) GetLastHeartbeatAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastHeartbeatAt.Get(), o.LastHeartbeatAt.IsSet()
}

// SetLastHeartbeatAt sets field value
func (o *TerminalDevice) SetLastHeartbeatAt(v string) {
	o.LastHeartbeatAt.Set(&v)
}

func (o TerminalDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["device_name"] = o.DeviceName.Get()
	toSerialize["terminal_serial"] = o.TerminalSerial.Get()
	toSerialize["terminal_ip"] = o.TerminalIp.Get()
	toSerialize["platform"] = o.Platform.Get()
	toSerialize["status"] = o.Status
	toSerialize["terminal_paired"] = o.TerminalPaired
	toSerialize["terminal_reachable"] = o.TerminalReachable.Get()
	toSerialize["app_version"] = o.AppVersion.Get()
	toSerialize["last_heartbeat_at"] = o.LastHeartbeatAt.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TerminalDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"device_name",
		"terminal_serial",
		"terminal_ip",
		"platform",
		"status",
		"terminal_paired",
		"terminal_reachable",
		"app_version",
		"last_heartbeat_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTerminalDevice := _TerminalDevice{}

	err = json.Unmarshal(data, &varTerminalDevice)

	if err != nil {
		return err
	}

	*o = TerminalDevice(varTerminalDevice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "device_name")
		delete(additionalProperties, "terminal_serial")
		delete(additionalProperties, "terminal_ip")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "status")
		delete(additionalProperties, "terminal_paired")
		delete(additionalProperties, "terminal_reachable")
		delete(additionalProperties, "app_version")
		delete(additionalProperties, "last_heartbeat_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerminalDevice struct {
	value *TerminalDevice
	isSet bool
}

func (v NullableTerminalDevice) Get() *TerminalDevice {
	return v.value
}

func (v *NullableTerminalDevice) Set(val *TerminalDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalDevice(val *TerminalDevice) *NullableTerminalDevice {
	return &NullableTerminalDevice{value: val, isSet: true}
}

func (v NullableTerminalDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


