# DhcpStatic6EditFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewDhcpStatic6EditFailed

`func NewDhcpStatic6EditFailed() *DhcpStatic6EditFailed`

NewDhcpStatic6EditFailed instantiates a new DhcpStatic6EditFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStatic6EditFailedWithDefaults

`func NewDhcpStatic6EditFailedWithDefaults() *DhcpStatic6EditFailed`

NewDhcpStatic6EditFailedWithDefaults instantiates a new DhcpStatic6EditFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpStatic6EditFailed) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpStatic6EditFailed) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpStatic6EditFailed) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpStatic6EditFailed) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpStatic6EditFailed) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpStatic6EditFailed) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpStatic6EditFailed) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpStatic6EditFailed) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpStatic6EditFailed) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpStatic6EditFailed) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpStatic6EditFailed) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DhcpStatic6EditFailed) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


