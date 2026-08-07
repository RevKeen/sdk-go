/*
RevKeen API

RevKeen is a fintech-grade API for payments, subscriptions, invoices, and billing. The canonical production MCP server is available at `https://mcp.revkeen.com/mcp`.  **API Version:** `2026-05-01` — Pin with the `RevKeen-Version` header.  **Quick Links:** [Full Documentation](https://docs.revkeen.com) | [Authentication](https://docs.revkeen.com/authentication) | [OAuth](https://docs.revkeen.com/oauth) | [SDKs](https://docs.revkeen.com/sdks) | [Webhooks](#webhooks) | [MCP Guide](https://docs.revkeen.com/mcp)  ## Authentication  Two authentication methods are supported:  ### API Keys (recommended for server-to-server REST API integrations)  Send your API key in the `x-api-key` header. Get keys from the [Dashboard](https://app.revkeen.com/settings/api-keys). Use `rk_sandbox_*` for test mode and `rk_live_*` for production.  ### OAuth 2.1 (recommended for MCP and third-party integrations)  Use OAuth 2.1 with PKCE for authorization code flow or client credentials for server-to-server. Tokens are sent via `Authorization: Bearer rk_oauth_*`. See the [OAuth guide](https://docs.revkeen.com/oauth) for setup.  - **Authorization Code + PKCE** — user-facing integrations, MCP hosts - **Client Credentials** — server-to-server, automated workflows - **Dynamic Client Registration** — MCP hosts that auto-register  ## MCP Integration  RevKeen's canonical production MCP server is `https://mcp.revkeen.com/mcp` using Streamable HTTP and OAuth 2.1 bearer tokens.  - **Customer launch surface** — read-first customer v1 tools with least-privilege scopes - **Host setup guide** — see the [MCP guide](https://docs.revkeen.com/mcp) for ChatGPT, Claude, and compatible MCP hosts  ## API Key Scopes  Scopes follow `{resource}:{action}` format (e.g., `invoices:read`, `customers:*`). See [full scope reference](https://docs.revkeen.com/authentication#scopes).  | Category | Scope | Description | |----------|-------|-------------| | **Payments & Checkout** | `checkout:read` | View checkout session details | |  | `checkout:write` | Create and manage checkout sessions | |  | `cart:read` | View cart session details (REV-3511) | |  | `cart:write` | Create and mutate cart sessions, line items, add-ons (REV-3511) | |  | `payment_links:read` | View payment links | |  | `payment_links:write` | Create and manage payment links | |  | `charges:read` | View one-time charges | |  | `charges:write` | Create one-time charges for customers | |  | `payments:read` | View payment details | |  | `payments:write` | Capture or void payments | |  | `payment_intents:read` | View payment intent details | |  | `payment_intents:write` | Create, confirm, capture, and cancel payment intents | |  | `setup_intents:read` | View setup intent details | |  | `setup_intents:write` | Create, confirm, and cancel setup intents | |  | `payment_methods:read` | View saved payment methods | |  | `payment_methods:write` | Attach and detach payment methods | | **Billing** | `invoices:read` | View invoices | |  | `invoices:write` | Create, update, and manage invoices | |  | `subscriptions:read` | View subscriptions | |  | `subscriptions:write` | Create, update, pause, and cancel subscriptions | |  | `subscription_schedules:read` | View subscription schedule details | |  | `subscription_schedules:write` | Create, update, cancel, and release subscription schedules | |  | `orders:read` | View orders | |  | `orders:write` | Create and manage orders | |  | `credit_notes:read` | View credit notes | |  | `credit_notes:write` | Create and void credit notes | | **Products & Pricing** | `products:read` | View product catalog | |  | `products:write` | Create and update products | |  | `prices:read` | View pricing information | |  | `prices:write` | Create and update prices | |  | `discounts:read` | View discount codes | |  | `discounts:write` | Create and manage discount codes | |  | `tax_rates:read` | View tax rate configurations | |  | `tax_rates:write` | Configure tax rates | | **Usage & Metering** | `meters:read` | View meter configurations | |  | `meters:write` | Create and update meters | |  | `usage:read` | View usage events and balances | |  | `usage:write` | Ingest usage events | | **Customers** | `customers:read` | View customer information | |  | `customers:write` | Create and update customers | |  | `entitlements:read` | View customer entitlements / feature access | |  | `entitlements:write` | Grant and revoke customer entitlements | |  | `businesses:read` | View business entities | |  | `businesses:write` | Manage business entities | | **Money Movement** | `refunds:read` | View refund details | |  | `refunds:write` | Issue refunds | |  | `voids:read` | View voided transactions | |  | `voids:write` | Void unsettled transactions | |  | `disputes:read` | View chargebacks and disputes | |  | `disputes:write` | Respond to disputes | |  | `payouts:read` | View payout and settlement data | | **Direct Debit** | `mandates:read` | View Direct Debit mandates and collection status | |  | `mandates:write` | Create, suspend, reinstate, and cancel Direct Debit mandates | | **Terminal** | `terminal:read` | View terminal devices and card-present payments | |  | `terminal:write` | Initiate, cancel, refund, and void terminal payments | | **Data Exchange** | `exports:read` | View and download data exports | |  | `exports:write` | Create data exports | |  | `imports:read` | View import status and history | |  | `imports:write` | Upload and run data imports | | **Analytics & Reporting** | `analytics:read` | View analytics and reports | |  | `finance:read` | View financial reports | | **Communication** | `comms:read` | View SMS and email delivery logs | |  | `comms:write` | Send SMS, email, and WhatsApp messages | |  | `automations:read` | View automations, runs, approvals, and traces | |  | `automations:write` | Create automations and trigger runs | | **Integrations** | `apps:read` | View connected applications | |  | `apps:write` | Manage app connections | |  | `webhooks:read` | View webhook endpoints | |  | `webhooks:write` | Manage webhook endpoints | |  | `integrations:read` | View integration status and sync logs | |  | `integrations:write` | Activate, configure, and sync integrations | |  | `events:read` | View webhook event logs | |  | `events:write` | Resend and test webhook events | |  | `sync:read` | View sync watermarks and state | |  | `sync:write` | Update sync watermarks |  ## Environments  | Environment | Base URL | API Key Prefix | |-------------|----------|----------------| | **Staging** | `https://staging-api.revkeen.com/v2` | `rk_sandbox_*` | | **Production** | `https://api.revkeen.com/v2` | `rk_live_*` |  ## Idempotency  Include `Idempotency-Key` header (UUID) on mutation requests. Keys are valid for 24 hours.  ## Rate Limits  | Plan | Requests/min | Burst | |------|-------------|-------| | **Staging** | 100 | 200 | | **Production** | 1000 | 2000 | | **Enterprise** | Custom | Custom | 

API version: 2026-05-01
Contact: info@revkeen.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package revkeen

import (
	"encoding/json"
	"fmt"
)

// checks if the StorefrontProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorefrontProduct{}

// StorefrontProduct Browser-safe product projection for headless storefronts: display data, active prices, and derived availability only.
type StorefrontProduct struct {
	Id string `json:"id"`
	Object string `json:"object"`
	ProductId NullableString `json:"product_id"`
	Slug NullableString `json:"slug"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Kind string `json:"kind"`
	PricingModel string `json:"pricing_model"`
	Currency string `json:"currency"`
	ImageUrl NullableString `json:"image_url"`
	DefaultPriceId NullableString `json:"default_price_id"`
	Prices []StorefrontPrice `json:"prices"`
	TrialDays int32 `json:"trial_days"`
	UsageMeterId NullableString `json:"usage_meter_id"`
	TaxBehavior NullableString `json:"tax_behavior"`
	TaxCode NullableString `json:"tax_code"`
	Availability StorefrontAvailability `json:"availability"`
	AdditionalProperties map[string]interface{}
}

type _StorefrontProduct StorefrontProduct

// NewStorefrontProduct instantiates a new StorefrontProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorefrontProduct(id string, object string, productId NullableString, slug NullableString, name string, description NullableString, kind string, pricingModel string, currency string, imageUrl NullableString, defaultPriceId NullableString, prices []StorefrontPrice, trialDays int32, usageMeterId NullableString, taxBehavior NullableString, taxCode NullableString, availability StorefrontAvailability) *StorefrontProduct {
	this := StorefrontProduct{}
	this.Id = id
	this.Object = object
	this.ProductId = productId
	this.Slug = slug
	this.Name = name
	this.Description = description
	this.Kind = kind
	this.PricingModel = pricingModel
	this.Currency = currency
	this.ImageUrl = imageUrl
	this.DefaultPriceId = defaultPriceId
	this.Prices = prices
	this.TrialDays = trialDays
	this.UsageMeterId = usageMeterId
	this.TaxBehavior = taxBehavior
	this.TaxCode = taxCode
	this.Availability = availability
	return &this
}

// NewStorefrontProductWithDefaults instantiates a new StorefrontProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorefrontProductWithDefaults() *StorefrontProduct {
	this := StorefrontProduct{}
	return &this
}

// GetId returns the Id field value
func (o *StorefrontProduct) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StorefrontProduct) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *StorefrontProduct) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *StorefrontProduct) SetObject(v string) {
	o.Object = v
}

// GetProductId returns the ProductId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetProductId() string {
	if o == nil || o.ProductId.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProductId.Get()
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId.Get(), o.ProductId.IsSet()
}

// SetProductId sets field value
func (o *StorefrontProduct) SetProductId(v string) {
	o.ProductId.Set(&v)
}

// GetSlug returns the Slug field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetSlug() string {
	if o == nil || o.Slug.Get() == nil {
		var ret string
		return ret
	}

	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// SetSlug sets field value
func (o *StorefrontProduct) SetSlug(v string) {
	o.Slug.Set(&v)
}

// GetName returns the Name field value
func (o *StorefrontProduct) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorefrontProduct) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *StorefrontProduct) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetKind returns the Kind field value
func (o *StorefrontProduct) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *StorefrontProduct) SetKind(v string) {
	o.Kind = v
}

// GetPricingModel returns the PricingModel field value
func (o *StorefrontProduct) GetPricingModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PricingModel
}

// GetPricingModelOk returns a tuple with the PricingModel field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetPricingModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricingModel, true
}

// SetPricingModel sets field value
func (o *StorefrontProduct) SetPricingModel(v string) {
	o.PricingModel = v
}

// GetCurrency returns the Currency field value
func (o *StorefrontProduct) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *StorefrontProduct) SetCurrency(v string) {
	o.Currency = v
}

// GetImageUrl returns the ImageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetImageUrl() string {
	if o == nil || o.ImageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// SetImageUrl sets field value
func (o *StorefrontProduct) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// GetDefaultPriceId returns the DefaultPriceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetDefaultPriceId() string {
	if o == nil || o.DefaultPriceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultPriceId.Get()
}

// GetDefaultPriceIdOk returns a tuple with the DefaultPriceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetDefaultPriceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPriceId.Get(), o.DefaultPriceId.IsSet()
}

// SetDefaultPriceId sets field value
func (o *StorefrontProduct) SetDefaultPriceId(v string) {
	o.DefaultPriceId.Set(&v)
}

// GetPrices returns the Prices field value
func (o *StorefrontProduct) GetPrices() []StorefrontPrice {
	if o == nil {
		var ret []StorefrontPrice
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetPricesOk() ([]StorefrontPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prices, true
}

// SetPrices sets field value
func (o *StorefrontProduct) SetPrices(v []StorefrontPrice) {
	o.Prices = v
}

// GetTrialDays returns the TrialDays field value
func (o *StorefrontProduct) GetTrialDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrialDays
}

// GetTrialDaysOk returns a tuple with the TrialDays field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetTrialDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialDays, true
}

// SetTrialDays sets field value
func (o *StorefrontProduct) SetTrialDays(v int32) {
	o.TrialDays = v
}

// GetUsageMeterId returns the UsageMeterId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetUsageMeterId() string {
	if o == nil || o.UsageMeterId.Get() == nil {
		var ret string
		return ret
	}

	return *o.UsageMeterId.Get()
}

// GetUsageMeterIdOk returns a tuple with the UsageMeterId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetUsageMeterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageMeterId.Get(), o.UsageMeterId.IsSet()
}

// SetUsageMeterId sets field value
func (o *StorefrontProduct) SetUsageMeterId(v string) {
	o.UsageMeterId.Set(&v)
}

// GetTaxBehavior returns the TaxBehavior field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetTaxBehavior() string {
	if o == nil || o.TaxBehavior.Get() == nil {
		var ret string
		return ret
	}

	return *o.TaxBehavior.Get()
}

// GetTaxBehaviorOk returns a tuple with the TaxBehavior field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetTaxBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxBehavior.Get(), o.TaxBehavior.IsSet()
}

// SetTaxBehavior sets field value
func (o *StorefrontProduct) SetTaxBehavior(v string) {
	o.TaxBehavior.Set(&v)
}

// GetTaxCode returns the TaxCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *StorefrontProduct) GetTaxCode() string {
	if o == nil || o.TaxCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.TaxCode.Get()
}

// GetTaxCodeOk returns a tuple with the TaxCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorefrontProduct) GetTaxCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaxCode.Get(), o.TaxCode.IsSet()
}

// SetTaxCode sets field value
func (o *StorefrontProduct) SetTaxCode(v string) {
	o.TaxCode.Set(&v)
}

// GetAvailability returns the Availability field value
func (o *StorefrontProduct) GetAvailability() StorefrontAvailability {
	if o == nil {
		var ret StorefrontAvailability
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *StorefrontProduct) GetAvailabilityOk() (*StorefrontAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *StorefrontProduct) SetAvailability(v StorefrontAvailability) {
	o.Availability = v
}

func (o StorefrontProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorefrontProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["product_id"] = o.ProductId.Get()
	toSerialize["slug"] = o.Slug.Get()
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["kind"] = o.Kind
	toSerialize["pricing_model"] = o.PricingModel
	toSerialize["currency"] = o.Currency
	toSerialize["image_url"] = o.ImageUrl.Get()
	toSerialize["default_price_id"] = o.DefaultPriceId.Get()
	toSerialize["prices"] = o.Prices
	toSerialize["trial_days"] = o.TrialDays
	toSerialize["usage_meter_id"] = o.UsageMeterId.Get()
	toSerialize["tax_behavior"] = o.TaxBehavior.Get()
	toSerialize["tax_code"] = o.TaxCode.Get()
	toSerialize["availability"] = o.Availability

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StorefrontProduct) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"product_id",
		"slug",
		"name",
		"description",
		"kind",
		"pricing_model",
		"currency",
		"image_url",
		"default_price_id",
		"prices",
		"trial_days",
		"usage_meter_id",
		"tax_behavior",
		"tax_code",
		"availability",
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

	varStorefrontProduct := _StorefrontProduct{}

	err = json.Unmarshal(data, &varStorefrontProduct)

	if err != nil {
		return err
	}

	*o = StorefrontProduct(varStorefrontProduct)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "product_id")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "kind")
		delete(additionalProperties, "pricing_model")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "default_price_id")
		delete(additionalProperties, "prices")
		delete(additionalProperties, "trial_days")
		delete(additionalProperties, "usage_meter_id")
		delete(additionalProperties, "tax_behavior")
		delete(additionalProperties, "tax_code")
		delete(additionalProperties, "availability")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorefrontProduct struct {
	value *StorefrontProduct
	isSet bool
}

func (v NullableStorefrontProduct) Get() *StorefrontProduct {
	return v.value
}

func (v *NullableStorefrontProduct) Set(val *StorefrontProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableStorefrontProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableStorefrontProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorefrontProduct(val *StorefrontProduct) *NullableStorefrontProduct {
	return &NullableStorefrontProduct{value: val, isSet: true}
}

func (v NullableStorefrontProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorefrontProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


