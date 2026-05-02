# ImportErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Row** | **int32** | Row number (1-based) | 
**Field** | Pointer to **string** | Field that caused the error | [optional] 
**Message** | **string** | Error message | 

## Methods

### NewImportErrorsInner

`func NewImportErrorsInner(row int32, message string, ) *ImportErrorsInner`

NewImportErrorsInner instantiates a new ImportErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportErrorsInnerWithDefaults

`func NewImportErrorsInnerWithDefaults() *ImportErrorsInner`

NewImportErrorsInnerWithDefaults instantiates a new ImportErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRow

`func (o *ImportErrorsInner) GetRow() int32`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *ImportErrorsInner) GetRowOk() (*int32, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *ImportErrorsInner) SetRow(v int32)`

SetRow sets Row field to given value.


### GetField

`func (o *ImportErrorsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ImportErrorsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ImportErrorsInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ImportErrorsInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetMessage

`func (o *ImportErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ImportErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ImportErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


