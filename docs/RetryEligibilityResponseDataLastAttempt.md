# RetryEligibilityResponseDataLastAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the payment attempt | 
**PublicId** | **string** | External reference ID (payatt_xxx format). Used as order_id in gateway requests. | 
**GatewayOrderId** | **string** | The exact order_id sent to the gateway (NMI). Format: payatt_&lt;public_id&gt;. Primary correlation key for webhook mapping. | 
**GatewayTransactionId** | **string** | Transaction ID returned by the gateway (NMI&#39;s transaction reference). | 
**TransactionId** | **string** | Legacy transaction reference (deprecated - use gateway_transaction_id) | 
**PaymentId** | **string** | Logical payment (PaymentIntent) this attempt belongs to | 
**InvoiceId** | **string** | Associated invoice ID | 
**SubscriptionId** | **string** | Associated subscription ID | 
**CheckoutSessionId** | **string** | Associated checkout session ID | 
**BillingRunId** | **string** | Billing run ID for batch reconciliation | 
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
**AvsCode** | **string** | Address Verification System response code | 
**CvvCode** | **string** | CVV verification response code | 
**IssuerCode** | **string** | Issuer-specific response code | 
**AmountCents** | **int32** | Payment amount in cents | 
**Currency** | **string** | Three-letter ISO currency code | 
**PaymentMethodLast4** | **string** | Last 4 digits of card | 
**PaymentMethodBrand** | **string** | Card brand (visa, mastercard, etc.) | 
**AttemptNumber** | **int32** | Attempt sequence number (1 &#x3D; first attempt) | 
**RetryAttemptNumber** | **int32** | Legacy retry attempt number | 
**MaxRetryAttempts** | **int32** | Maximum retry attempts allowed | 
**NextRetryAt** | **time.Time** | Next scheduled retry timestamp | 
**IsScheduledRetry** | **bool** | Whether this was a scheduled retry | 
**IsManualRetry** | **bool** | Whether this was a manual retry | 
**CreatedAt** | **time.Time** | When the attempt was created | 
**UpdatedAt** | **time.Time** | When the attempt was last updated | 

## Methods

### NewRetryEligibilityResponseDataLastAttempt

`func NewRetryEligibilityResponseDataLastAttempt(id string, publicId string, gatewayOrderId string, gatewayTransactionId string, transactionId string, paymentId string, invoiceId string, subscriptionId string, checkoutSessionId string, billingRunId string, gateway string, gatewayCode string, gatewayMessage string, category string, severity string, retryability string, customerAction string, merchantAction string, subscriptionDirective string, safeCustomerMessageKey string, declineType NullableString, attemptStatus string, avsCode string, cvvCode string, issuerCode string, amountCents int32, currency string, paymentMethodLast4 string, paymentMethodBrand string, attemptNumber int32, retryAttemptNumber int32, maxRetryAttempts int32, nextRetryAt time.Time, isScheduledRetry bool, isManualRetry bool, createdAt time.Time, updatedAt time.Time, ) *RetryEligibilityResponseDataLastAttempt`

NewRetryEligibilityResponseDataLastAttempt instantiates a new RetryEligibilityResponseDataLastAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryEligibilityResponseDataLastAttemptWithDefaults

`func NewRetryEligibilityResponseDataLastAttemptWithDefaults() *RetryEligibilityResponseDataLastAttempt`

NewRetryEligibilityResponseDataLastAttemptWithDefaults instantiates a new RetryEligibilityResponseDataLastAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RetryEligibilityResponseDataLastAttempt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RetryEligibilityResponseDataLastAttempt) SetId(v string)`

SetId sets Id field to given value.


### GetPublicId

`func (o *RetryEligibilityResponseDataLastAttempt) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *RetryEligibilityResponseDataLastAttempt) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetGatewayOrderId

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOrderId() string`

GetGatewayOrderId returns the GatewayOrderId field if non-nil, zero value otherwise.

### GetGatewayOrderIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOrderIdOk() (*string, bool)`

GetGatewayOrderIdOk returns a tuple with the GatewayOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayOrderId

`func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayOrderId(v string)`

SetGatewayOrderId sets GatewayOrderId field to given value.


### GetGatewayTransactionId

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayTransactionId() string`

GetGatewayTransactionId returns the GatewayTransactionId field if non-nil, zero value otherwise.

### GetGatewayTransactionIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayTransactionIdOk() (*string, bool)`

GetGatewayTransactionIdOk returns a tuple with the GatewayTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTransactionId

`func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayTransactionId(v string)`

SetGatewayTransactionId sets GatewayTransactionId field to given value.


### GetTransactionId

`func (o *RetryEligibilityResponseDataLastAttempt) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *RetryEligibilityResponseDataLastAttempt) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetPaymentId

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetInvoiceId

`func (o *RetryEligibilityResponseDataLastAttempt) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RetryEligibilityResponseDataLastAttempt) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSubscriptionId

`func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RetryEligibilityResponseDataLastAttempt) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetCheckoutSessionId

`func (o *RetryEligibilityResponseDataLastAttempt) GetCheckoutSessionId() string`

GetCheckoutSessionId returns the CheckoutSessionId field if non-nil, zero value otherwise.

### GetCheckoutSessionIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCheckoutSessionIdOk() (*string, bool)`

GetCheckoutSessionIdOk returns a tuple with the CheckoutSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutSessionId

`func (o *RetryEligibilityResponseDataLastAttempt) SetCheckoutSessionId(v string)`

SetCheckoutSessionId sets CheckoutSessionId field to given value.


### GetBillingRunId

`func (o *RetryEligibilityResponseDataLastAttempt) GetBillingRunId() string`

GetBillingRunId returns the BillingRunId field if non-nil, zero value otherwise.

### GetBillingRunIdOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetBillingRunIdOk() (*string, bool)`

GetBillingRunIdOk returns a tuple with the BillingRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingRunId

`func (o *RetryEligibilityResponseDataLastAttempt) SetBillingRunId(v string)`

SetBillingRunId sets BillingRunId field to given value.


### GetGateway

`func (o *RetryEligibilityResponseDataLastAttempt) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RetryEligibilityResponseDataLastAttempt) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetGatewayCode

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayCode() string`

GetGatewayCode returns the GatewayCode field if non-nil, zero value otherwise.

### GetGatewayCodeOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayCodeOk() (*string, bool)`

GetGatewayCodeOk returns a tuple with the GatewayCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayCode

`func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayCode(v string)`

SetGatewayCode sets GatewayCode field to given value.


### GetGatewayMessage

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayMessage() string`

GetGatewayMessage returns the GatewayMessage field if non-nil, zero value otherwise.

### GetGatewayMessageOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetGatewayMessageOk() (*string, bool)`

GetGatewayMessageOk returns a tuple with the GatewayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMessage

`func (o *RetryEligibilityResponseDataLastAttempt) SetGatewayMessage(v string)`

SetGatewayMessage sets GatewayMessage field to given value.


### GetCategory

`func (o *RetryEligibilityResponseDataLastAttempt) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RetryEligibilityResponseDataLastAttempt) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSeverity

`func (o *RetryEligibilityResponseDataLastAttempt) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *RetryEligibilityResponseDataLastAttempt) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetRetryability

`func (o *RetryEligibilityResponseDataLastAttempt) GetRetryability() string`

GetRetryability returns the Retryability field if non-nil, zero value otherwise.

### GetRetryabilityOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetRetryabilityOk() (*string, bool)`

GetRetryabilityOk returns a tuple with the Retryability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryability

`func (o *RetryEligibilityResponseDataLastAttempt) SetRetryability(v string)`

SetRetryability sets Retryability field to given value.


### GetCustomerAction

`func (o *RetryEligibilityResponseDataLastAttempt) GetCustomerAction() string`

GetCustomerAction returns the CustomerAction field if non-nil, zero value otherwise.

### GetCustomerActionOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCustomerActionOk() (*string, bool)`

GetCustomerActionOk returns a tuple with the CustomerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAction

`func (o *RetryEligibilityResponseDataLastAttempt) SetCustomerAction(v string)`

SetCustomerAction sets CustomerAction field to given value.


### GetMerchantAction

`func (o *RetryEligibilityResponseDataLastAttempt) GetMerchantAction() string`

GetMerchantAction returns the MerchantAction field if non-nil, zero value otherwise.

### GetMerchantActionOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetMerchantActionOk() (*string, bool)`

GetMerchantActionOk returns a tuple with the MerchantAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAction

`func (o *RetryEligibilityResponseDataLastAttempt) SetMerchantAction(v string)`

SetMerchantAction sets MerchantAction field to given value.


### GetSubscriptionDirective

`func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionDirective() string`

GetSubscriptionDirective returns the SubscriptionDirective field if non-nil, zero value otherwise.

### GetSubscriptionDirectiveOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetSubscriptionDirectiveOk() (*string, bool)`

GetSubscriptionDirectiveOk returns a tuple with the SubscriptionDirective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDirective

`func (o *RetryEligibilityResponseDataLastAttempt) SetSubscriptionDirective(v string)`

SetSubscriptionDirective sets SubscriptionDirective field to given value.


### GetSafeCustomerMessageKey

`func (o *RetryEligibilityResponseDataLastAttempt) GetSafeCustomerMessageKey() string`

GetSafeCustomerMessageKey returns the SafeCustomerMessageKey field if non-nil, zero value otherwise.

### GetSafeCustomerMessageKeyOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetSafeCustomerMessageKeyOk() (*string, bool)`

GetSafeCustomerMessageKeyOk returns a tuple with the SafeCustomerMessageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeCustomerMessageKey

`func (o *RetryEligibilityResponseDataLastAttempt) SetSafeCustomerMessageKey(v string)`

SetSafeCustomerMessageKey sets SafeCustomerMessageKey field to given value.


### GetDeclineType

`func (o *RetryEligibilityResponseDataLastAttempt) GetDeclineType() string`

GetDeclineType returns the DeclineType field if non-nil, zero value otherwise.

### GetDeclineTypeOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetDeclineTypeOk() (*string, bool)`

GetDeclineTypeOk returns a tuple with the DeclineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineType

`func (o *RetryEligibilityResponseDataLastAttempt) SetDeclineType(v string)`

SetDeclineType sets DeclineType field to given value.


### SetDeclineTypeNil

`func (o *RetryEligibilityResponseDataLastAttempt) SetDeclineTypeNil(b bool)`

 SetDeclineTypeNil sets the value for DeclineType to be an explicit nil

### UnsetDeclineType
`func (o *RetryEligibilityResponseDataLastAttempt) UnsetDeclineType()`

UnsetDeclineType ensures that no value is present for DeclineType, not even an explicit nil
### GetAttemptStatus

`func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptStatus() string`

GetAttemptStatus returns the AttemptStatus field if non-nil, zero value otherwise.

### GetAttemptStatusOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptStatusOk() (*string, bool)`

GetAttemptStatusOk returns a tuple with the AttemptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptStatus

`func (o *RetryEligibilityResponseDataLastAttempt) SetAttemptStatus(v string)`

SetAttemptStatus sets AttemptStatus field to given value.


### GetAvsCode

`func (o *RetryEligibilityResponseDataLastAttempt) GetAvsCode() string`

GetAvsCode returns the AvsCode field if non-nil, zero value otherwise.

### GetAvsCodeOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetAvsCodeOk() (*string, bool)`

GetAvsCodeOk returns a tuple with the AvsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsCode

`func (o *RetryEligibilityResponseDataLastAttempt) SetAvsCode(v string)`

SetAvsCode sets AvsCode field to given value.


### GetCvvCode

`func (o *RetryEligibilityResponseDataLastAttempt) GetCvvCode() string`

GetCvvCode returns the CvvCode field if non-nil, zero value otherwise.

### GetCvvCodeOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCvvCodeOk() (*string, bool)`

GetCvvCodeOk returns a tuple with the CvvCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvvCode

`func (o *RetryEligibilityResponseDataLastAttempt) SetCvvCode(v string)`

SetCvvCode sets CvvCode field to given value.


### GetIssuerCode

`func (o *RetryEligibilityResponseDataLastAttempt) GetIssuerCode() string`

GetIssuerCode returns the IssuerCode field if non-nil, zero value otherwise.

### GetIssuerCodeOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetIssuerCodeOk() (*string, bool)`

GetIssuerCodeOk returns a tuple with the IssuerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCode

`func (o *RetryEligibilityResponseDataLastAttempt) SetIssuerCode(v string)`

SetIssuerCode sets IssuerCode field to given value.


### GetAmountCents

`func (o *RetryEligibilityResponseDataLastAttempt) GetAmountCents() int32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetAmountCentsOk() (*int32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *RetryEligibilityResponseDataLastAttempt) SetAmountCents(v int32)`

SetAmountCents sets AmountCents field to given value.


### GetCurrency

`func (o *RetryEligibilityResponseDataLastAttempt) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *RetryEligibilityResponseDataLastAttempt) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethodLast4

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodLast4() string`

GetPaymentMethodLast4 returns the PaymentMethodLast4 field if non-nil, zero value otherwise.

### GetPaymentMethodLast4Ok

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodLast4Ok() (*string, bool)`

GetPaymentMethodLast4Ok returns a tuple with the PaymentMethodLast4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodLast4

`func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentMethodLast4(v string)`

SetPaymentMethodLast4 sets PaymentMethodLast4 field to given value.


### GetPaymentMethodBrand

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodBrand() string`

GetPaymentMethodBrand returns the PaymentMethodBrand field if non-nil, zero value otherwise.

### GetPaymentMethodBrandOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetPaymentMethodBrandOk() (*string, bool)`

GetPaymentMethodBrandOk returns a tuple with the PaymentMethodBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodBrand

`func (o *RetryEligibilityResponseDataLastAttempt) SetPaymentMethodBrand(v string)`

SetPaymentMethodBrand sets PaymentMethodBrand field to given value.


### GetAttemptNumber

`func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *RetryEligibilityResponseDataLastAttempt) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.


### GetRetryAttemptNumber

`func (o *RetryEligibilityResponseDataLastAttempt) GetRetryAttemptNumber() int32`

GetRetryAttemptNumber returns the RetryAttemptNumber field if non-nil, zero value otherwise.

### GetRetryAttemptNumberOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetRetryAttemptNumberOk() (*int32, bool)`

GetRetryAttemptNumberOk returns a tuple with the RetryAttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAttemptNumber

`func (o *RetryEligibilityResponseDataLastAttempt) SetRetryAttemptNumber(v int32)`

SetRetryAttemptNumber sets RetryAttemptNumber field to given value.


### GetMaxRetryAttempts

`func (o *RetryEligibilityResponseDataLastAttempt) GetMaxRetryAttempts() int32`

GetMaxRetryAttempts returns the MaxRetryAttempts field if non-nil, zero value otherwise.

### GetMaxRetryAttemptsOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetMaxRetryAttemptsOk() (*int32, bool)`

GetMaxRetryAttemptsOk returns a tuple with the MaxRetryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryAttempts

`func (o *RetryEligibilityResponseDataLastAttempt) SetMaxRetryAttempts(v int32)`

SetMaxRetryAttempts sets MaxRetryAttempts field to given value.


### GetNextRetryAt

`func (o *RetryEligibilityResponseDataLastAttempt) GetNextRetryAt() time.Time`

GetNextRetryAt returns the NextRetryAt field if non-nil, zero value otherwise.

### GetNextRetryAtOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetNextRetryAtOk() (*time.Time, bool)`

GetNextRetryAtOk returns a tuple with the NextRetryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryAt

`func (o *RetryEligibilityResponseDataLastAttempt) SetNextRetryAt(v time.Time)`

SetNextRetryAt sets NextRetryAt field to given value.


### GetIsScheduledRetry

`func (o *RetryEligibilityResponseDataLastAttempt) GetIsScheduledRetry() bool`

GetIsScheduledRetry returns the IsScheduledRetry field if non-nil, zero value otherwise.

### GetIsScheduledRetryOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetIsScheduledRetryOk() (*bool, bool)`

GetIsScheduledRetryOk returns a tuple with the IsScheduledRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledRetry

`func (o *RetryEligibilityResponseDataLastAttempt) SetIsScheduledRetry(v bool)`

SetIsScheduledRetry sets IsScheduledRetry field to given value.


### GetIsManualRetry

`func (o *RetryEligibilityResponseDataLastAttempt) GetIsManualRetry() bool`

GetIsManualRetry returns the IsManualRetry field if non-nil, zero value otherwise.

### GetIsManualRetryOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetIsManualRetryOk() (*bool, bool)`

GetIsManualRetryOk returns a tuple with the IsManualRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManualRetry

`func (o *RetryEligibilityResponseDataLastAttempt) SetIsManualRetry(v bool)`

SetIsManualRetry sets IsManualRetry field to given value.


### GetCreatedAt

`func (o *RetryEligibilityResponseDataLastAttempt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RetryEligibilityResponseDataLastAttempt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RetryEligibilityResponseDataLastAttempt) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RetryEligibilityResponseDataLastAttempt) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RetryEligibilityResponseDataLastAttempt) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


