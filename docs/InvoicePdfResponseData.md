# InvoicePdfResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Presigned URL that downloads the invoice PDF. It expires at expires_at and is a secret link: do not log it or store it. Request a new one when it expires. |
**ExpiresAt** | **time.Time** | When the URL stops working (ISO 8601). Five minutes after the response. |

## Methods

### NewInvoicePdfResponseData

`func NewInvoicePdfResponseData(url string, expiresAt time.Time, ) *InvoicePdfResponseData`

NewInvoicePdfResponseData instantiates a new InvoicePdfResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePdfResponseDataWithDefaults

`func NewInvoicePdfResponseDataWithDefaults() *InvoicePdfResponseData`

NewInvoicePdfResponseDataWithDefaults instantiates a new InvoicePdfResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *InvoicePdfResponseData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InvoicePdfResponseData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InvoicePdfResponseData) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetExpiresAt

`func (o *InvoicePdfResponseData) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvoicePdfResponseData) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvoicePdfResponseData) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
