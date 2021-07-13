# AppApplicationDeleteSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]AppApplicationDeleteSuccessData**](AppApplicationDeleteSuccessData.md) |  | [optional] 

## Methods

### NewAppApplicationDeleteSuccess

`func NewAppApplicationDeleteSuccess() *AppApplicationDeleteSuccess`

NewAppApplicationDeleteSuccess instantiates a new AppApplicationDeleteSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApplicationDeleteSuccessWithDefaults

`func NewAppApplicationDeleteSuccessWithDefaults() *AppApplicationDeleteSuccess`

NewAppApplicationDeleteSuccessWithDefaults instantiates a new AppApplicationDeleteSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AppApplicationDeleteSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AppApplicationDeleteSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AppApplicationDeleteSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AppApplicationDeleteSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *AppApplicationDeleteSuccess) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AppApplicationDeleteSuccess) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AppApplicationDeleteSuccess) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AppApplicationDeleteSuccess) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *AppApplicationDeleteSuccess) GetData() []AppApplicationDeleteSuccessData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppApplicationDeleteSuccess) GetDataOk() (*[]AppApplicationDeleteSuccessData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppApplicationDeleteSuccess) SetData(v []AppApplicationDeleteSuccessData)`

SetData sets Data field to given value.

### HasData

`func (o *AppApplicationDeleteSuccess) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


