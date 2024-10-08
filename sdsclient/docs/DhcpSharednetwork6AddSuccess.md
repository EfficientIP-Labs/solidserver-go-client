# DhcpSharednetwork6AddSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerDhcpSharednetwork6AddSuccess**](DataInnerDhcpSharednetwork6AddSuccess.md) |  | [optional] 

## Methods

### NewDhcpSharednetwork6AddSuccess

`func NewDhcpSharednetwork6AddSuccess() *DhcpSharednetwork6AddSuccess`

NewDhcpSharednetwork6AddSuccess instantiates a new DhcpSharednetwork6AddSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSharednetwork6AddSuccessWithDefaults

`func NewDhcpSharednetwork6AddSuccessWithDefaults() *DhcpSharednetwork6AddSuccess`

NewDhcpSharednetwork6AddSuccessWithDefaults instantiates a new DhcpSharednetwork6AddSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpSharednetwork6AddSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpSharednetwork6AddSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpSharednetwork6AddSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpSharednetwork6AddSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpSharednetwork6AddSuccess) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpSharednetwork6AddSuccess) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpSharednetwork6AddSuccess) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpSharednetwork6AddSuccess) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpSharednetwork6AddSuccess) GetData() []DataInnerDhcpSharednetwork6AddSuccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpSharednetwork6AddSuccess) GetDataOk() (*[]DataInnerDhcpSharednetwork6AddSuccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpSharednetwork6AddSuccess) SetData(v []DataInnerDhcpSharednetwork6AddSuccess)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpSharednetwork6AddSuccess) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


