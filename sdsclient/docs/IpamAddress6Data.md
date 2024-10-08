# IpamAddress6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerIpamAddress6Data**](DataInnerIpamAddress6Data.md) |  | [optional] 

## Methods

### NewIpamAddress6Data

`func NewIpamAddress6Data() *IpamAddress6Data`

NewIpamAddress6Data instantiates a new IpamAddress6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAddress6DataWithDefaults

`func NewIpamAddress6DataWithDefaults() *IpamAddress6Data`

NewIpamAddress6DataWithDefaults instantiates a new IpamAddress6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *IpamAddress6Data) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IpamAddress6Data) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IpamAddress6Data) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *IpamAddress6Data) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *IpamAddress6Data) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *IpamAddress6Data) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *IpamAddress6Data) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *IpamAddress6Data) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *IpamAddress6Data) GetData() []DataInnerIpamAddress6Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IpamAddress6Data) GetDataOk() (*[]DataInnerIpamAddress6Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IpamAddress6Data) SetData(v []DataInnerIpamAddress6Data)`

SetData sets Data field to given value.

### HasData

`func (o *IpamAddress6Data) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


