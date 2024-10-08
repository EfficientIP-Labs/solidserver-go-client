# DhcpScope6DeleteSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerDhcpScope6DeleteSuccess**](DataInnerDhcpScope6DeleteSuccess.md) |  | [optional] 

## Methods

### NewDhcpScope6DeleteSuccess

`func NewDhcpScope6DeleteSuccess() *DhcpScope6DeleteSuccess`

NewDhcpScope6DeleteSuccess instantiates a new DhcpScope6DeleteSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpScope6DeleteSuccessWithDefaults

`func NewDhcpScope6DeleteSuccessWithDefaults() *DhcpScope6DeleteSuccess`

NewDhcpScope6DeleteSuccessWithDefaults instantiates a new DhcpScope6DeleteSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpScope6DeleteSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpScope6DeleteSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpScope6DeleteSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpScope6DeleteSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpScope6DeleteSuccess) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpScope6DeleteSuccess) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpScope6DeleteSuccess) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpScope6DeleteSuccess) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpScope6DeleteSuccess) GetData() []DataInnerDhcpScope6DeleteSuccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpScope6DeleteSuccess) GetDataOk() (*[]DataInnerDhcpScope6DeleteSuccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpScope6DeleteSuccess) SetData(v []DataInnerDhcpScope6DeleteSuccess)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpScope6DeleteSuccess) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


