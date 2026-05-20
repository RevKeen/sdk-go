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

// checks if the SetupIntent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetupIntent{}

// SetupIntent A SetupIntent saves a payment method for future use WITHOUT charging the customer. Use it to collect and verify card details before storing them.
type SetupIntent struct {
	// Unique identifier for the setup intent
	Id string `json:"id"`
	// Object type, always 'setup_intent'
	Object string `json:"object"`
	// Public ID visible in API responses (seti_xxx format)
	PublicId string `json:"public_id"`
	// The status of the setup intent
	Status string `json:"status"`
	// ID of the customer this setup intent is for
	CustomerId NullableString `json:"customer_id"`
	// ID of the payment method being set up
	PaymentMethodId NullableString `json:"payment_method_id"`
	// Allowed payment method types for this setup
	PaymentMethodTypes []string `json:"payment_method_types"`
	// Indicates how the payment method will be used
	Usage string `json:"usage"`
	NextAction SetupIntentNextAction `json:"next_action"`
	// Client secret for frontend confirmation
	ClientSecret string `json:"client_secret"`
	// Name of the payment processor that handled this setup intent
	Gateway string `json:"gateway"`
	LastError SetupIntentError `json:"last_error"`
	// Reason for cancellation if canceled
	CancellationReason NullableString `json:"cancellation_reason"`
	// Merchant description for reference
	Description NullableString `json:"description"`
	// Custom metadata attached to the setup intent
	Metadata map[string]interface{} `json:"metadata"`
	// When the setup intent was confirmed
	ConfirmedAt NullableTime `json:"confirmed_at"`
	// When the setup intent was canceled
	CanceledAt NullableTime `json:"canceled_at"`
	// When the setup intent was created
	CreatedAt time.Time `json:"created_at"`
	// When the setup intent was last updated
	UpdatedAt time.Time `json:"updated_at"`
	AdditionalProperties map[string]interface{}
}

type _SetupIntent SetupIntent

// NewSetupIntent instantiates a new SetupIntent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupIntent(id string, object string, publicId string, status string, customerId NullableString, paymentMethodId NullableString, paymentMethodTypes []string, usage string, nextAction SetupIntentNextAction, clientSecret string, gateway string, lastError SetupIntentError, cancellationReason NullableString, description NullableString, metadata map[string]interface{}, confirmedAt NullableTime, canceledAt NullableTime, createdAt time.Time, updatedAt time.Time) *SetupIntent {
	this := SetupIntent{}
	this.Id = id
	this.Object = object
	this.PublicId = publicId
	this.Status = status
	this.CustomerId = customerId
	this.PaymentMethodId = paymentMethodId
	this.PaymentMethodTypes = paymentMethodTypes
	this.Usage = usage
	this.NextAction = nextAction
	this.ClientSecret = clientSecret
	this.Gateway = gateway
	this.LastError = lastError
	this.CancellationReason = cancellationReason
	this.Description = description
	this.Metadata = metadata
	this.ConfirmedAt = confirmedAt
	this.CanceledAt = canceledAt
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewSetupIntentWithDefaults instantiates a new SetupIntent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupIntentWithDefaults() *SetupIntent {
	this := SetupIntent{}
	return &this
}

// GetId returns the Id field value
func (o *SetupIntent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetupIntent) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *SetupIntent) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SetupIntent) SetObject(v string) {
	o.Object = v
}

// GetPublicId returns the PublicId field value
func (o *SetupIntent) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value
func (o *SetupIntent) SetPublicId(v string) {
	o.PublicId = v
}

// GetStatus returns the Status field value
func (o *SetupIntent) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SetupIntent) SetStatus(v string) {
	o.Status = v
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SetupIntent) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// SetCustomerId sets field value
func (o *SetupIntent) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}

// GetPaymentMethodId returns the PaymentMethodId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SetupIntent) GetPaymentMethodId() string {
	if o == nil || o.PaymentMethodId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentMethodId.Get()
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodId.Get(), o.PaymentMethodId.IsSet()
}

// SetPaymentMethodId sets field value
func (o *SetupIntent) SetPaymentMethodId(v string) {
	o.PaymentMethodId.Set(&v)
}

// GetPaymentMethodTypes returns the PaymentMethodTypes field value
func (o *SetupIntent) GetPaymentMethodTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PaymentMethodTypes
}

// GetPaymentMethodTypesOk returns a tuple with the PaymentMethodTypes field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetPaymentMethodTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethodTypes, true
}

// SetPaymentMethodTypes sets field value
func (o *SetupIntent) SetPaymentMethodTypes(v []string) {
	o.PaymentMethodTypes = v
}

// GetUsage returns the Usage field value
func (o *SetupIntent) GetUsage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value
func (o *SetupIntent) SetUsage(v string) {
	o.Usage = v
}

// GetNextAction returns the NextAction field value
func (o *SetupIntent) GetNextAction() SetupIntentNextAction {
	if o == nil {
		var ret SetupIntentNextAction
		return ret
	}

	return o.NextAction
}

// GetNextActionOk returns a tuple with the NextAction field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetNextActionOk() (*SetupIntentNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextAction, true
}

// SetNextAction sets field value
func (o *SetupIntent) SetNextAction(v SetupIntentNextAction) {
	o.NextAction = v
}

// GetClientSecret returns the ClientSecret field value
func (o *SetupIntent) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *SetupIntent) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetGateway returns the Gateway field value
func (o *SetupIntent) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetGatewayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *SetupIntent) SetGateway(v string) {
	o.Gateway = v
}

// GetLastError returns the LastError field value
func (o *SetupIntent) GetLastError() SetupIntentError {
	if o == nil {
		var ret SetupIntentError
		return ret
	}

	return o.LastError
}

// GetLastErrorOk returns a tuple with the LastError field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetLastErrorOk() (*SetupIntentError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastError, true
}

// SetLastError sets field value
func (o *SetupIntent) SetLastError(v SetupIntentError) {
	o.LastError = v
}

// GetCancellationReason returns the CancellationReason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SetupIntent) GetCancellationReason() string {
	if o == nil || o.CancellationReason.Get() == nil {
		var ret string
		return ret
	}

	return *o.CancellationReason.Get()
}

// GetCancellationReasonOk returns a tuple with the CancellationReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetCancellationReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancellationReason.Get(), o.CancellationReason.IsSet()
}

// SetCancellationReason sets field value
func (o *SetupIntent) SetCancellationReason(v string) {
	o.CancellationReason.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SetupIntent) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *SetupIntent) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetMetadata returns the Metadata field value
func (o *SetupIntent) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *SetupIntent) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetConfirmedAt returns the ConfirmedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SetupIntent) GetConfirmedAt() time.Time {
	if o == nil || o.ConfirmedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ConfirmedAt.Get()
}

// GetConfirmedAtOk returns a tuple with the ConfirmedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetConfirmedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfirmedAt.Get(), o.ConfirmedAt.IsSet()
}

// SetConfirmedAt sets field value
func (o *SetupIntent) SetConfirmedAt(v time.Time) {
	o.ConfirmedAt.Set(&v)
}

// GetCanceledAt returns the CanceledAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *SetupIntent) GetCanceledAt() time.Time {
	if o == nil || o.CanceledAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CanceledAt.Get()
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetupIntent) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanceledAt.Get(), o.CanceledAt.IsSet()
}

// SetCanceledAt sets field value
func (o *SetupIntent) SetCanceledAt(v time.Time) {
	o.CanceledAt.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *SetupIntent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SetupIntent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SetupIntent) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SetupIntent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SetupIntent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SetupIntent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetupIntent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["public_id"] = o.PublicId
	toSerialize["status"] = o.Status
	toSerialize["customer_id"] = o.CustomerId.Get()
	toSerialize["payment_method_id"] = o.PaymentMethodId.Get()
	toSerialize["payment_method_types"] = o.PaymentMethodTypes
	toSerialize["usage"] = o.Usage
	toSerialize["next_action"] = o.NextAction
	toSerialize["client_secret"] = o.ClientSecret
	toSerialize["gateway"] = o.Gateway
	toSerialize["last_error"] = o.LastError
	toSerialize["cancellation_reason"] = o.CancellationReason.Get()
	toSerialize["description"] = o.Description.Get()
	toSerialize["metadata"] = o.Metadata
	toSerialize["confirmed_at"] = o.ConfirmedAt.Get()
	toSerialize["canceled_at"] = o.CanceledAt.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SetupIntent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"public_id",
		"status",
		"customer_id",
		"payment_method_id",
		"payment_method_types",
		"usage",
		"next_action",
		"client_secret",
		"gateway",
		"last_error",
		"cancellation_reason",
		"description",
		"metadata",
		"confirmed_at",
		"canceled_at",
		"created_at",
		"updated_at",
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

	varSetupIntent := _SetupIntent{}

	err = json.Unmarshal(data, &varSetupIntent)

	if err != nil {
		return err
	}

	*o = SetupIntent(varSetupIntent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "public_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "payment_method_id")
		delete(additionalProperties, "payment_method_types")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "next_action")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "last_error")
		delete(additionalProperties, "cancellation_reason")
		delete(additionalProperties, "description")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "confirmed_at")
		delete(additionalProperties, "canceled_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSetupIntent struct {
	value *SetupIntent
	isSet bool
}

func (v NullableSetupIntent) Get() *SetupIntent {
	return v.value
}

func (v *NullableSetupIntent) Set(val *SetupIntent) {
	v.value = val
	v.isSet = true
}

func (v NullableSetupIntent) IsSet() bool {
	return v.isSet
}

func (v *NullableSetupIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetupIntent(val *SetupIntent) *NullableSetupIntent {
	return &NullableSetupIntent{value: val, isSet: true}
}

func (v NullableSetupIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetupIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


