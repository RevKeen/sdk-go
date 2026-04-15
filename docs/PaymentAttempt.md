# PaymentAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the payment attempt | 
**PublicId** | **NullableString** | External reference ID (payatt_xxx format). Used as order_id in gateway requests. | 
**GatewayOrderId** | **NullableString** | The exact order_id sent to the gateway (NMI). Format: payatt_&lt;public_id&gt;. Primary correlation key for webhook mapping. | 
**GatewayTransactionId** | **NullableString** | Transaction ID returned by the gateway (NMI&#39;s transaction reference). | 
**TransactionId** | **NullableString** | Legacy transaction reference (deprecated - use gateway_transaction_id) | 
**PaymentId** | **NullableString** | Logical payment (PaymentIntent) this attempt belongs to | 
**InvoiceId** | **NullableString** | Associated invoice ID | 
**SubscriptionId** | **NullableString** | Associated subscription ID | 
**CheckoutSessionId** | **NullableString** | Associated checkout session ID | 
**BillingRunId** | **NullableString** | Billing run ID for batch reconciliation | 
**Gateway** | **string** | Payment gateway (e.g., &#39;nmi&#39;, &#39;stripe&#39;) | 
**GatewayCode** | **string** | Raw gateway response code | 
**GatewayMessage** | **string** | Raw gateway response message | 
**Category** | **string** | Normalized error category for the payment attempt | 
**Severity** | **string** | Severity level for logging/monitoring | 
**Retryability** | **string** | Retry behavior recommendation. retry_now: Safe to retry immediately. retry_later: Retry after delay. do_not_retry: Hard decline, don&#39;t retry. | 
**CustomerAction** | **string** | Recommended action for the customer | 
**MerchantAction** | **string** | Recommended action for the merchant | 
**SubscriptionDirective** | **string** | Directive for subscription handling after decline | 
**SafeCustomerMessageKey** | **string** | Template key for customer-facing message | 
**DeclineType** | **NullableString** | Decline classification. soft: May succeed on retry. hard: Will not succeed on retry. | 
**AttemptStatus** | **string** | Final outcome status of the attempt | 
**AvsCode** | **NullableString** | Address Verification System response code | 
**CvvCode** | **NullableString** | CVV verification response code | 
**IssuerCode** | **NullableString** | Issuer-specific response code | 
**AmountCents** | **NullableInt32** | Payment amount in cents | 
**Currency** | **NullableString** | Three-letter ISO currency code | 
**PaymentMethodLast4** | **NullableString** | Last 4 digits of card | 
**PaymentMethodBrand** | **NullableString** | Card brand (visa, mastercard, etc.) | 
**AttemptNumber** | **int32** | Attempt sequence number (1 &#x3D; first attempt) | 
**RetryAttemptNumber** | **int32** | Legacy retry attempt number | 
**MaxRetryAttempts** | **NullableInt32** | Maximum retry attempts allowed | 
**NextRetryAt** | **NullableTime** | Next scheduled retry timestamp | 
**IsScheduledRetry** | **bool** | Whether this was a scheduled retry | 
**IsManualRetry** | **bool** | Whether this was a manual retry | 
**CreatedAt** | **time.Time** | When the attempt was created | 
**UpdatedAt** | **time.Time** | When the attempt was last updated | 

## Methods

### NewPaymentAttempt

`func NewPaymentAttempt(id string, publicId NullableString, gatewayOrderId NullableString, gatewayTransactionId NullableString, transactionId NullableString, paymentId NullableString, invoiceId NullableString, subscriptionId NullableString, checkoutSessionId NullableString, billingRunId NullableString, gateway string, gatewayCode string, gatewayMessage string, category string, severity string, retryability string, customerAction string, merchantAction string, subscriptionDirective string, safeCustomerMessageKey string, declineType NullableString, attemptStatus string, avsCode NullableString, cvvCode NullableString, issuerCode NullableString, amountCents NullableInt32, currency NullableString, paymentMethodLast4 NullableString, paymentMethodBrand NullableString, attemptNumber int32, retryAttemptNumber int32, maxRetryAttempts NullableInt32, nextRetryAt NullableTime, isScheduledRetry bool, isManualRetry bool, createdAt time.Time, updatedAt time.Time, ) *PaymentAttempt`

NewPaymentAttempt instantiates a new PaymentAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAttemptWithDefaults

`func NewPaymentAttemptWithDefaults() *PaymentAttempt`

NewPaymentAttemptWithDefaults instantiates a new PaymentAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentAttempt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentAttempt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentAttempt) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *PaymentAttempt) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *PaymentAttempt) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *PaymentAttempt) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### SetPublicIdNil

`func (o *PaymentAttempt) SetPublicIdNil(b bool)`

 SetPublicIdNil sets the value for PublicId to be an explicit nil

### UnsetPublicId
`func (o *PaymentAttempt) UnsetPublicId()`

UnsetPublicId ensures that no value is present for PublicId, not even an explicit nil
### GetGatewayOrderId

`func (o *PaymentAttempt) GetGatewayOrderId() string`

GetGatewayOrderId returns the GatewayOrderId field if non-nil, zero value otherwise.

### GetGatewayOrderIdOk

`func (o *PaymentAttempt) GetGatewayOrderIdOk() (*string, bool)`

GetGatewayOrderIdOk returns a tuple with the GatewayOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayOrderId

`func (o *PaymentAttempt) SetGatewayOrderId(v string)`

SetGatewayOrderId sets GatewayOrderId field to given value.


### SetGatewayOrderIdNil

`func (o *PaymentAttempt) SetGatewayOrderIdNil(b bool)`

 SetGatewayOrderIdNil sets the value for GatewayOrderId to be an explicit nil

### UnsetGatewayOrderId
`func (o *PaymentAttempt) UnsetGatewayOrderId()`

UnsetGatewayOrderId ensures that no value is present for GatewayOrderId, not even an explicit nil
### GetGatewayTransactionId

`func (o *PaymentAttempt) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *PaymentAttempt) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *PaymentAttempt) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.


### SetGatewayTransactionIdNil

`func (o *PaymentAttempt) SetGatewayTransactionIdNil(b bool)`

 SetGatewayTransactionIdNil sets the value for GatewayTransactionId to be an explicit nil

### UnsetGatewayTransactionId
`func (o *PaymentAttempt) UnsetGatewayTransactionId()`

UnsetGatewayTransactionId ensures that no value is present for GatewayTransactionId, not even an explicit nil
### GetTransactionId

`func (o *PaymentAttempt) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentAttempt) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentAttempt) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### SetTransactionIdNil

`func (o *PaymentAttempt) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *PaymentAttempt) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetPaymentId

`func (o *PaymentAttempt) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentAttempt) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentAttempt) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### SetPaymentIdNil

`func (o *PaymentAttempt) SetPaymentIdNil(b bool)`

 SetPaymentIdNil sets the value for PaymentId to be an explicit nil

### UnsetPaymentId
`func (o *PaymentAttempt) UnsetPaymentId()`

UnsetPaymentId ensures that no value is present for PaymentId, not even an explicit nil
### GetInvoiceId

`func (o *PaymentAttempt) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PaymentAttempt) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PaymentAttempt) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### SetInvoiceIdNil

`func (o *PaymentAttempt) SetInvoiceIdNil(b bool)`

 SetInvoiceIdNil sets the value for InvoiceId to be an explicit nil

### UnsetInvoiceId
`func (o *PaymentAttempt) UnsetInvoiceId()`

UnsetInvoiceId ensures that no value is present for InvoiceId, not even an explicit nil
### GetSubscriptionId

`func (o *PaymentAttempt) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentAttempt) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentAttempt) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### SetSubscriptionIdNil

`func (o *PaymentAttempt) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *PaymentAttempt) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCheckoutSessionId

`func (o *PaymentAttempt) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *PaymentAttempt) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *PaymentAttempt) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.


### SetCheckoutSessionIdNil

`func (o *PaymentAttempt) SetCheckoutSessionIdNil(b bool)`

 SetCheckoutSessionIdNil sets the value for CheckoutSessionId to be an explicit nil

### UnsetCheckoutSessionId
`func (o *PaymentAttempt) UnsetCheckoutSessionId()`

UnsetCheckoutSessionId ensures that no value is present for CheckoutSessionId, not even an explicit nil
### GetBillingRunId

`func (o *PaymentAttempt) GetBillingRunId() string`

GetBillingRunId returns the BillingRunId field if non-nil, zero value otherwise.

### GetBillingRunIdOk

`func (o *PaymentAttempt) GetBillingRunIdOk() (*string, bool)`

GetBillingRunIdOk returns a tuple with the BillingRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRunId

`func (o *PaymentAttempt) SetBillingRunId(v string)`

SetBillingRunId sets BillingRunId field to given value.


### SetBillingRunIdNil

`func (o *PaymentAttempt) SetBillingRunIdNil(b bool)`

 SetBillingRunIdNil sets the value for BillingRunId to be an explicit nil

### UnsetBillingRunId
`func (o *PaymentAttempt) UnsetBillingRunId()`

UnsetBillingRunId ensures that no value is present for BillingRunId, not even an explicit nil
### GetGateway

`func (o *PaymentAttempt) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PaymentAttempt) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PaymentAttempt) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayCode

`func (o *PaymentAttempt) GetGatewayCode() string`

GetGatewayCode returns the GatewayCode field if non-nil, zero value otherwise.

### GetGatewayCodeOk

`func (o *PaymentAttempt) GetGatewayCodeOk() (*string, bool)`

GetGatewayCodeOk returns a tuple with the GatewayCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayCode

`func (o *PaymentAttempt) SetGatewayCode(v string)`

SetGatewayCode sets GatewayCode field to given value.


### GetGatewayMessage

`func (o *PaymentAttempt) GetGatewayMessage() string`

GetGatewayMessage returns the GatewayMessage field if non-nil, zero value otherwise.

### GetGatewayMessageOk

`func (o *PaymentAttempt) GetGatewayMessageOk() (*string, bool)`

GetGatewayMessageOk returns a tuple with the GatewayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMessage

`func (o *PaymentAttempt) SetGatewayMessage(v string)`

SetGatewayMessage sets GatewayMessage field to given value.


### GetCategory

`func (o *PaymentAttempt) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PaymentAttempt) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PaymentAttempt) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSeverity

`func (o *PaymentAttempt) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *PaymentAttempt) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *PaymentAttempt) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetRetryability

`func (o *PaymentAttempt) GetRetryability() string`

GetRetryability returns the Retryability field if non-nil, zero value otherwise.

### GetRetryabilityOk

`func (o *PaymentAttempt) GetRetryabilityOk() (*string, bool)`

GetRetryabilityOk returns a tuple with the Retryability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryability

`func (o *PaymentAttempt) SetRetryability(v string)`

SetRetryability sets Retryability field to given value.


### GetCustomerAction

`func (o *PaymentAttempt) GetCustomerAction() string`

GetCustomerAction returns the CustomerAction field if non-nil, zero value otherwise.

### GetCustomerActionOk

`func (o *PaymentAttempt) GetCustomerActionOk() (*string, bool)`

GetCustomerActionOk returns a tuple with the CustomerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAction

`func (o *PaymentAttempt) SetCustomerAction(v string)`

SetCustomerAction sets CustomerAction field to given value.


### GetMerchantAction

`func (o *PaymentAttempt) GetMerchantAction() string`

GetMerchantAction returns the MerchantAction field if non-nil, zero value otherwise.

### GetMerchantActionOk

`func (o *PaymentAttempt) GetMerchantActionOk() (*string, bool)`

GetMerchantActionOk returns a tuple with the MerchantAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAction

`func (o *PaymentAttempt) SetMerchantAction(v string)`

SetMerchantAction sets MerchantAction field to given value.


### GetSubscriptionDirective

`func (o *PaymentAttempt) GetSubscriptionDirective() string`

GetSubscriptionDirective returns the SubscriptionDirective field if non-nil, zero value otherwise.

### GetSubscriptionDirectiveOk

`func (o *PaymentAttempt) GetSubscriptionDirectiveOk() (*string, bool)`

GetSubscriptionDirectiveOk returns a tuple with the SubscriptionDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDirective

`func (o *PaymentAttempt) SetSubscriptionDirective(v string)`

SetSubscriptionDirective sets SubscriptionDirective field to given value.


### GetSafeCustomerMessageKey

`func (o *PaymentAttempt) GetSafeCustomerMessageKey() string`

GetSafeCustomerMessageKey returns the SafeCustomerMessageKey field if non-nil, zero value otherwise.

### GetSafeCustomerMessageKeyOk

`func (o *PaymentAttempt) GetSafeCustomerMessageKeyOk() (*string, bool)`

GetSafeCustomerMessageKeyOk returns a tuple with the SafeCustomerMessageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeCustomerMessageKey

`func (o *PaymentAttempt) SetSafeCustomerMessageKey(v string)`

SetSafeCustomerMessageKey sets SafeCustomerMessageKey field to given value.


### GetDeclineType

`func (o *PaymentAttempt) GetDeclineType() string`

GetDeclineType returns the DeclineType field if non-nil, zero value otherwise.

### GetDeclineTypeOk

`func (o *PaymentAttempt) GetDeclineTypeOk() (*string, bool)`

GetDeclineTypeOk returns a tuple with the DeclineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineType

`func (o *PaymentAttempt) SetDeclineType(v string)`

SetDeclineType sets DeclineType field to given value.


### SetDeclineTypeNil

`func (o *PaymentAttempt) SetDeclineTypeNil(b bool)`

 SetDeclineTypeNil sets the value for DeclineType to be an explicit nil

### UnsetDeclineType
`func (o *PaymentAttempt) UnsetDeclineType()`

UnsetDeclineType ensures that no value is present for DeclineType, not even an explicit nil
### GetAttemptStatus

`func (o *PaymentAttempt) GetAttemptStatus() string`

GetAttemptStatus returns the AttemptStatus field if non-nil, zero value otherwise.

### GetAttemptStatusOk

`func (o *PaymentAttempt) GetAttemptStatusOk() (*string, bool)`

GetAttemptStatusOk returns a tuple with the AttemptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptStatus

`func (o *PaymentAttempt) SetAttemptStatus(v string)`

SetAttemptStatus sets AttemptStatus field to given value.


### GetAvsCode

`func (o *PaymentAttempt) GetAvsCode() string`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *PaymentAttempt) GetAvsCodeOk() (*string, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *PaymentAttempt) SetAvsCode(v string)`

SetAvsCode sets AvsCode field to given value.


### SetAvsCodeNil

`func (o *PaymentAttempt) SetAvsCodeNil(b bool)`

 SetAvsCodeNil sets the value for AvsCode to be an explicit nil

### UnsetAvsCode
`func (o *PaymentAttempt) UnsetAvsCode()`

UnsetAvsCode ensures that no value is present for AvsCode, not even an explicit nil
### GetCvvCode

`func (o *PaymentAttempt) GetCvvCode() string`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *PaymentAttempt) GetCvvCodeOk() (*string, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *PaymentAttempt) SetCvvCode(v string)`

SetCvvCode sets CvvCode field to given value.


### SetCvvCodeNil

`func (o *PaymentAttempt) SetCvvCodeNil(b bool)`

 SetCvvCodeNil sets the value for CvvCode to be an explicit nil

### UnsetCvvCode
`func (o *PaymentAttempt) UnsetCvvCode()`

UnsetCvvCode ensures that no value is present for CvvCode, not even an explicit nil
### GetIssuerCode

`func (o *PaymentAttempt) GetIssuerCode() string`

GetIssuerCode returns the IssuerCode field if non-nil, zero value otherwise.

### GetIssuerCodeOk

`func (o *PaymentAttempt) GetIssuerCodeOk() (*string, bool)`

GetIssuerCodeOk returns a tuple with the IssuerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCode

`func (o *PaymentAttempt) SetIssuerCode(v string)`

SetIssuerCode sets IssuerCode field to given value.


### SetIssuerCodeNil

`func (o *PaymentAttempt) SetIssuerCodeNil(b bool)`

 SetIssuerCodeNil sets the value for IssuerCode to be an explicit nil

### UnsetIssuerCode
`func (o *PaymentAttempt) UnsetIssuerCode()`

UnsetIssuerCode ensures that no value is present for IssuerCode, not even an explicit nil
### GetAmountCents

`func (o *PaymentAttempt) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *PaymentAttempt) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *PaymentAttempt) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### SetAmountCentsNil

`func (o *PaymentAttempt) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *PaymentAttempt) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetCurrency

`func (o *PaymentAttempt) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAttempt) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAttempt) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *PaymentAttempt) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentAttempt) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPaymentMethodLast4

`func (o *PaymentAttempt) GetPaymentMethodLast4() string`

GetPaymentMethodLast4 returns the PaymentMethodLast4 field if non-nil, zero value otherwise.

### GetPaymentMethodLast4Ok

`func (o *PaymentAttempt) GetPaymentMethodLast4Ok() (*string, bool)`

GetPaymentMethodLast4Ok returns a tuple with the PaymentMethodLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodLast4

`func (o *PaymentAttempt) SetPaymentMethodLast4(v string)`

SetPaymentMethodLast4 sets PaymentMethodLast4 field to given value.


### SetPaymentMethodLast4Nil

`func (o *PaymentAttempt) SetPaymentMethodLast4Nil(b bool)`

 SetPaymentMethodLast4Nil sets the value for PaymentMethodLast4 to be an explicit nil

### UnsetPaymentMethodLast4
`func (o *PaymentAttempt) UnsetPaymentMethodLast4()`

UnsetPaymentMethodLast4 ensures that no value is present for PaymentMethodLast4, not even an explicit nil
### GetPaymentMethodBrand

`func (o *PaymentAttempt) GetPaymentMethodBrand() string`

GetPaymentMethodBrand returns the PaymentMethodBrand field if non-nil, zero value otherwise.

### GetPaymentMethodBrandOk

`func (o *PaymentAttempt) GetPaymentMethodBrandOk() (*string, bool)`

GetPaymentMethodBrandOk returns a tuple with the PaymentMethodBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodBrand

`func (o *PaymentAttempt) SetPaymentMethodBrand(v string)`

SetPaymentMethodBrand sets PaymentMethodBrand field to given value.


### SetPaymentMethodBrandNil

`func (o *PaymentAttempt) SetPaymentMethodBrandNil(b bool)`

 SetPaymentMethodBrandNil sets the value for PaymentMethodBrand to be an explicit nil

### UnsetPaymentMethodBrand
`func (o *PaymentAttempt) UnsetPaymentMethodBrand()`

UnsetPaymentMethodBrand ensures that no value is present for PaymentMethodBrand, not even an explicit nil
### GetAttemptNumber

`func (o *PaymentAttempt) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *PaymentAttempt) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *PaymentAttempt) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.


### GetRetryAttemptNumber

`func (o *PaymentAttempt) GetRetryAttemptNumber() int32`

GetRetryAttemptNumber returns the RetryAttemptNumber field if non-nil, zero value otherwise.

### GetRetryAttemptNumberOk

`func (o *PaymentAttempt) GetRetryAttemptNumberOk() (*int32, bool)`

GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttemptNumber

`func (o *PaymentAttempt) SetRetryAttemptNumber(v int32)`

SetRetryAttemptNumber sets RetryAttemptNumber field to given value.


### GetMaxRetryAttempts

`func (o *PaymentAttempt) GetMaxRetryAttempts() int32`

GetMaxRetryAttempts returns the MaxRetryAttempts field if non-nil, zero value otherwise.

### GetMaxRetryAttemptsOk

`func (o *PaymentAttempt) GetMaxRetryAttemptsOk() (*int32, bool)`

GetMaxRetryAttemptsOk returns a tuple with the MaxRetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryAttempts

`func (o *PaymentAttempt) SetMaxRetryAttempts(v int32)`

SetMaxRetryAttempts sets MaxRetryAttempts field to given value.


### SetMaxRetryAttemptsNil

`func (o *PaymentAttempt) SetMaxRetryAttemptsNil(b bool)`

 SetMaxRetryAttemptsNil sets the value for MaxRetryAttempts to be an explicit nil

### UnsetMaxRetryAttempts
`func (o *PaymentAttempt) UnsetMaxRetryAttempts()`

UnsetMaxRetryAttempts ensures that no value is present for MaxRetryAttempts, not even an explicit nil
### GetNextRetryAt

`func (o *PaymentAttempt) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *PaymentAttempt) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *PaymentAttempt) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### SetNextRetryAtNil

`func (o *PaymentAttempt) SetNextRetryAtNil(b bool)`

 SetNextRetryAtNil sets the value for NextRetryAt to be an explicit nil

### UnsetNextRetryAt
`func (o *PaymentAttempt) UnsetNextRetryAt()`

UnsetNextRetryAt ensures that no value is present for NextRetryAt, not even an explicit nil
### GetIsScheduledRetry

`func (o *PaymentAttempt) GetIsScheduledRetry() bool`

GetIsScheduledRetry returns the IsScheduledRetry field if non-nil, zero value otherwise.

### GetIsScheduledRetryOk

`func (o *PaymentAttempt) GetIsScheduledRetryOk() (*bool, bool)`

GetIsScheduledRetryOk returns a tuple with the IsScheduledRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledRetry

`func (o *PaymentAttempt) SetIsScheduledRetry(v bool)`

SetIsScheduledRetry sets IsScheduledRetry field to given value.


### GetIsManualRetry

`func (o *PaymentAttempt) GetIsManualRetry() bool`

GetIsManualRetry returns the IsManualRetry field if non-nil, zero value otherwise.

### GetIsManualRetryOk

`func (o *PaymentAttempt) GetIsManualRetryOk() (*bool, bool)`

GetIsManualRetryOk returns a tuple with the IsManualRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManualRetry

`func (o *PaymentAttempt) SetIsManualRetry(v bool)`

SetIsManualRetry sets IsManualRetry field to given value.


### GetCreatedAt

`func (o *PaymentAttempt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaymentAttempt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaymentAttempt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PaymentAttempt) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PaymentAttempt) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PaymentAttempt) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


