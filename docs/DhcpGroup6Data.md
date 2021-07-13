# DhcpGroup6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DhcpGroup6DataData**](DhcpGroup6DataData.md) |  | [optional] 

## Methods

### NewDhcpGroup6Data

`func NewDhcpGroup6Data() *DhcpGroup6Data`

NewDhcpGroup6Data instantiates a new DhcpGroup6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGroup6DataWithDefaults

`func NewDhcpGroup6DataWithDefaults() *DhcpGroup6Data`

NewDhcpGroup6DataWithDefaults instantiates a new DhcpGroup6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpGroup6Data) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpGroup6Data) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpGroup6Data) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpGroup6Data) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpGroup6Data) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpGroup6Data) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpGroup6Data) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpGroup6Data) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpGroup6Data) GetData() []DhcpGroup6DataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpGroup6Data) GetDataOk() (*[]DhcpGroup6DataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpGroup6Data) SetData(v []DhcpGroup6DataData)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpGroup6Data) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


