# DhcpScope6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DhcpScope6DataData**](DhcpScope6DataData.md) |  | [optional] 

## Methods

### NewDhcpScope6Data

`func NewDhcpScope6Data() *DhcpScope6Data`

NewDhcpScope6Data instantiates a new DhcpScope6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpScope6DataWithDefaults

`func NewDhcpScope6DataWithDefaults() *DhcpScope6Data`

NewDhcpScope6DataWithDefaults instantiates a new DhcpScope6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpScope6Data) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpScope6Data) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpScope6Data) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpScope6Data) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpScope6Data) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpScope6Data) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpScope6Data) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpScope6Data) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpScope6Data) GetData() []DhcpScope6DataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpScope6Data) GetDataOk() (*[]DhcpScope6DataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpScope6Data) SetData(v []DhcpScope6DataData)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpScope6Data) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


