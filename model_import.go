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
	"time"
	"fmt"
)

// checks if the Import type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Import{}

// Import An async data import job. Upload a CSV file to create records in bulk. Poll GET /v2/imports/:id for status and error details.
type Import struct {
	// Unique import job ID
	Id string `json:"id"`
	// Object type
	Object string `json:"object"`
	// Import job status. pending: queued. validating: checking data. processing: creating records. completed: all rows imported. completed_with_errors: some rows failed. failed: import aborted.
	Status string `json:"status"`
	// The type of resource to import.
	ResourceType string `json:"resource_type"`
	// Total rows in the uploaded file
	TotalRows NullableInt32 `json:"total_rows,omitempty"`
	// Number of rows processed so far
	ProcessedRows NullableInt32 `json:"processed_rows,omitempty"`
	// Number of rows successfully imported
	SuccessCount NullableInt32 `json:"success_count,omitempty"`
	// Number of rows that failed
	ErrorCount NullableInt32 `json:"error_count,omitempty"`
	// Detailed error information for failed rows
	Errors []ImportErrorsInner `json:"errors,omitempty"`
	// When the import was created (ISO 8601)
	CreatedAt time.Time `json:"created_at"`
	// When the import completed (ISO 8601)
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Import Import

// NewImport instantiates a new Import object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImport(id string, object string, status string, resourceType string, createdAt time.Time) *Import {
	this := Import{}
	this.Id = id
	this.Object = object
	this.Status = status
	this.ResourceType = resourceType
	this.CreatedAt = createdAt
	return &this
}

// NewImportWithDefaults instantiates a new Import object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportWithDefaults() *Import {
	this := Import{}
	return &this
}

// GetId returns the Id field value
func (o *Import) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Import) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Import) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *Import) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Import) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Import) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *Import) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Import) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Import) SetStatus(v string) {
	o.Status = v
}

// GetResourceType returns the ResourceType field value
func (o *Import) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *Import) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *Import) SetResourceType(v string) {
	o.ResourceType = v
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetTotalRows() int32 {
	if o == nil || IsNil(o.TotalRows.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalRows.Get()
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetTotalRowsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalRows.Get(), o.TotalRows.IsSet()
}

// HasTotalRows returns a boolean if a field has been set.
func (o *Import) HasTotalRows() bool {
	if o != nil && o.TotalRows.IsSet() {
		return true
	}

	return false
}

// SetTotalRows gets a reference to the given NullableInt32 and assigns it to the TotalRows field.
func (o *Import) SetTotalRows(v int32) {
	o.TotalRows.Set(&v)
}
// SetTotalRowsNil sets the value for TotalRows to be an explicit nil
func (o *Import) SetTotalRowsNil() {
	o.TotalRows.Set(nil)
}

// UnsetTotalRows ensures that no value is present for TotalRows, not even an explicit nil
func (o *Import) UnsetTotalRows() {
	o.TotalRows.Unset()
}

// GetProcessedRows returns the ProcessedRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetProcessedRows() int32 {
	if o == nil || IsNil(o.ProcessedRows.Get()) {
		var ret int32
		return ret
	}
	return *o.ProcessedRows.Get()
}

// GetProcessedRowsOk returns a tuple with the ProcessedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetProcessedRowsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessedRows.Get(), o.ProcessedRows.IsSet()
}

// HasProcessedRows returns a boolean if a field has been set.
func (o *Import) HasProcessedRows() bool {
	if o != nil && o.ProcessedRows.IsSet() {
		return true
	}

	return false
}

// SetProcessedRows gets a reference to the given NullableInt32 and assigns it to the ProcessedRows field.
func (o *Import) SetProcessedRows(v int32) {
	o.ProcessedRows.Set(&v)
}
// SetProcessedRowsNil sets the value for ProcessedRows to be an explicit nil
func (o *Import) SetProcessedRowsNil() {
	o.ProcessedRows.Set(nil)
}

// UnsetProcessedRows ensures that no value is present for ProcessedRows, not even an explicit nil
func (o *Import) UnsetProcessedRows() {
	o.ProcessedRows.Unset()
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetSuccessCount() int32 {
	if o == nil || IsNil(o.SuccessCount.Get()) {
		var ret int32
		return ret
	}
	return *o.SuccessCount.Get()
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetSuccessCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuccessCount.Get(), o.SuccessCount.IsSet()
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *Import) HasSuccessCount() bool {
	if o != nil && o.SuccessCount.IsSet() {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given NullableInt32 and assigns it to the SuccessCount field.
func (o *Import) SetSuccessCount(v int32) {
	o.SuccessCount.Set(&v)
}
// SetSuccessCountNil sets the value for SuccessCount to be an explicit nil
func (o *Import) SetSuccessCountNil() {
	o.SuccessCount.Set(nil)
}

// UnsetSuccessCount ensures that no value is present for SuccessCount, not even an explicit nil
func (o *Import) UnsetSuccessCount() {
	o.SuccessCount.Unset()
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetErrorCount() int32 {
	if o == nil || IsNil(o.ErrorCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ErrorCount.Get()
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetErrorCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCount.Get(), o.ErrorCount.IsSet()
}

// HasErrorCount returns a boolean if a field has been set.
func (o *Import) HasErrorCount() bool {
	if o != nil && o.ErrorCount.IsSet() {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given NullableInt32 and assigns it to the ErrorCount field.
func (o *Import) SetErrorCount(v int32) {
	o.ErrorCount.Set(&v)
}
// SetErrorCountNil sets the value for ErrorCount to be an explicit nil
func (o *Import) SetErrorCountNil() {
	o.ErrorCount.Set(nil)
}

// UnsetErrorCount ensures that no value is present for ErrorCount, not even an explicit nil
func (o *Import) UnsetErrorCount() {
	o.ErrorCount.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetErrors() []ImportErrorsInner {
	if o == nil {
		var ret []ImportErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetErrorsOk() ([]ImportErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Import) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ImportErrorsInner and assigns it to the Errors field.
func (o *Import) SetErrors(v []ImportErrorsInner) {
	o.Errors = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Import) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Import) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Import) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Import) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Import) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Import) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *Import) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *Import) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *Import) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

func (o Import) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Import) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["resource_type"] = o.ResourceType
	if o.TotalRows.IsSet() {
		toSerialize["total_rows"] = o.TotalRows.Get()
	}
	if o.ProcessedRows.IsSet() {
		toSerialize["processed_rows"] = o.ProcessedRows.Get()
	}
	if o.SuccessCount.IsSet() {
		toSerialize["success_count"] = o.SuccessCount.Get()
	}
	if o.ErrorCount.IsSet() {
		toSerialize["error_count"] = o.ErrorCount.Get()
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
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

func (o *Import) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"status",
		"resource_type",
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

	varImport := _Import{}

	err = json.Unmarshal(data, &varImport)

	if err != nil {
		return err
	}

	*o = Import(varImport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "status")
		delete(additionalProperties, "resource_type")
		delete(additionalProperties, "total_rows")
		delete(additionalProperties, "processed_rows")
		delete(additionalProperties, "success_count")
		delete(additionalProperties, "error_count")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "completed_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImport struct {
	value *Import
	isSet bool
}

func (v NullableImport) Get() *Import {
	return v.value
}

func (v *NullableImport) Set(val *Import) {
	v.value = val
	v.isSet = true
}

func (v NullableImport) IsSet() bool {
	return v.isSet
}

func (v *NullableImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImport(val *Import) *NullableImport {
	return &NullableImport{value: val, isSet: true}
}

func (v NullableImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


