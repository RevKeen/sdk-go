# TransactionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Data** | [**[]Transaction**](Transaction.md) |  | 
**HasMore** | **bool** |  | 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewTransactionListResponse

`func NewTransactionListResponse(object string, data []Transaction, hasMore bool, ) *TransactionListResponse`

NewTransactionListResponse instantiates a new TransactionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListResponseWithDefaults

`func NewTransactionListResponseWithDefaults() *TransactionListResponse`

NewTransactionListResponseWithDefaults instantiates a new TransactionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *TransactionListResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TransactionListResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TransactionListResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetData

`func (o *TransactionListResponse) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionListResponse) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionListResponse) SetData(v []Transaction)`

SetData sets Data field to given value.


### GetHasMore

`func (o *TransactionListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TransactionListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TransactionListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetTotalCount

`func (o *TransactionListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TransactionListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TransactionListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TransactionListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


