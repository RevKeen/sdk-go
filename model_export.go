/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-01-15` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Sandbox** | `https://sandbox-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Sandbox** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-01-15
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Export type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Export{}

// Export An async data export job. Create an export, poll for completion, then download the file via the presigned URL.
type Export struct {
	// Unique export job ID
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Export job status. pending: queued. processing: generating file. completed: ready for download. failed: export failed.
	Status string `json:"status"`
	// The type of resource to export.
	ResourceType string `json:"resource_type"`
	// Output file format. csv: Comma-separated values. xlsx: Excel workbook.
	Format string `json:"format"`
	// Presigned download URL (available when status is 'completed'). Expires after 1 hour.
	DownloadUrl NullableString `json:"download_url,omitempty"`
	// Total rows exported (available when completed)
	TotalRows NullableInt32 `json:"total_rows,omitempty"`
	// File size in bytes (available when completed)
	FileSizeBytes NullableInt32 `json:"file_size_bytes,omitempty"`
	// Error message (if status is 'failed')
	ErrorMessage NullableString `json:"error_message,omitempty"`
	// When the export was requested (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// When the export completed (ISO 8601)
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Export Export

// NewExport instantiates a new Export object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExport(id string, object string, status string, resourceType string, format string, createdAt time.Time) *Export {
	this := Export{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.ResourceType = resourceType
	this.Format = format
	this.CreatedAt = createdAt
	return &this
}

// NewExportWithDefaults instantiates a new Export object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportWithDefaults() *Export {
	this := Export{}
	return &this
}

// GetId returns the Id field value
func (o *Export) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Export) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Export) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Export) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Export) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Export) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *Export) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Export) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Export) SetStatus(v string) {
	o.Status = v
}

// GetResourceType returns the ResourceType field value
func (o *Export) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Export) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Export) SetResourceType(v string) {
	o.ResourceType = v
}

// GetFormat returns the Format field value
func (o *Export) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *Export) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *Export) SetFormat(v string) {
	o.Format = v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *Export) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *Export) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *Export) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *Export) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalRows.Get()
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetTotalRowsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalRows.Get(), o.TotalRows.IsSet()
}

// HasTotalRows returns a boolean if a field has been set.
func (o *Export) HasTotalRows() bool {
	if o != nil && o.TotalRows.IsSet() {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given NullableInt32 and assigns it to the TotalRows field.
func (o *Export) SetTotalRows(v int32) {
	o.TotalRows.Set(&v)
}
// SetTotalRowsNil sets the value for TotalRows to be an explicit nil
func (o *Export) SetTotalRowsNil() {
	o.TotalRows.Set(nil)
}

// UnsetTotalRows ensures that no value is present for TotalRows, not even an explicit nil
func (o *Export) UnsetTotalRows() {
	o.TotalRows.Unset()
}

// GetFileSizeBytes returns the FileSizeBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetFileSizeBytes() int32 {
	if o == nil || IsNil(o.FileSizeBytes.Get()) {
		var ret int32
		return ret
	}
	return *o.FileSizeBytes.Get()
}

// GetFileSizeBytesOk returns a tuple with the FileSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetFileSizeBytesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSizeBytes.Get(), o.FileSizeBytes.IsSet()
}

// HasFileSizeBytes returns a boolean if a field has been set.
func (o *Export) HasFileSizeBytes() bool {
	if o != nil && o.FileSizeBytes.IsSet() {
		return true
	}

	return false
}

// SetFileSizeBytes gets a reference to the given NullableInt32 and assigns it to the FileSizeBytes field.
func (o *Export) SetFileSizeBytes(v int32) {
	o.FileSizeBytes.Set(&v)
}
// SetFileSizeBytesNil sets the value for FileSizeBytes to be an explicit nil
func (o *Export) SetFileSizeBytesNil() {
	o.FileSizeBytes.Set(nil)
}

// UnsetFileSizeBytes ensures that no value is present for FileSizeBytes, not even an explicit nil
func (o *Export) UnsetFileSizeBytes() {
	o.FileSizeBytes.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Export) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *Export) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *Export) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *Export) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *Export) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Export) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Export) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Export) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Export) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Export) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *Export) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *Export) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *Export) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

func (o Export) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Export) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["format"] = o.Format
	if o.DownloadUrl.IsSet() {
		toSerialize["download_url"] = o.DownloadUrl.Get()
	}
	if o.TotalRows.IsSet() {
		toSerialize["total_rows"] = o.TotalRows.Get()
	}
	if o.FileSizeBytes.IsSet() {
		toSerialize["file_size_bytes"] = o.FileSizeBytes.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	toSerialize["created_at"] = o.CreatedAt
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Export) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"resource_type",
		"format",
		"created_at",
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

	varExport := _Export{}

	err = json.Unmarshal(data, &varExport)

	if err != nil {
		return err
	}

	*o = Export(varExport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "status")
		delete(additionalProperties, "resource_type")
		delete(additionalProperties, "format")
		delete(additionalProperties, "download_url")
		delete(additionalProperties, "total_rows")
		delete(additionalProperties, "file_size_bytes")
		delete(additionalProperties, "error_message")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExport struct {
	value *Export
	isSet bool
}

func (v NullableExport) Get() *Export {
	return v.value
}

func (v *NullableExport) Set(val *Export) {
	v.value = val
	v.isSet = true
}

func (v NullableExport) IsSet() bool {
	return v.isSet
}

func (v *NullableExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExport(val *Export) *NullableExport {
	return &NullableExport{value: val, isSet: true}
}

func (v NullableExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


