# IpamPool6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerIpamPool6Data**](DataInnerIpamPool6Data.md) |  | [optional] 

## Methods

### NewIpamPool6Data

`func NewIpamPool6Data() *IpamPool6Data`

NewIpamPool6Data instantiates a new IpamPool6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPool6DataWithDefaults

`func NewIpamPool6DataWithDefaults() *IpamPool6Data`

NewIpamPool6DataWithDefaults instantiates a new IpamPool6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *IpamPool6Data) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *IpamPool6Data) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *IpamPool6Data) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *IpamPool6Data) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *IpamPool6Data) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *IpamPool6Data) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *IpamPool6Data) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *IpamPool6Data) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *IpamPool6Data) GetData() []DataInnerIpamPool6Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IpamPool6Data) GetDataOk() (*[]DataInnerIpamPool6Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IpamPool6Data) SetData(v []DataInnerIpamPool6Data)`

SetData sets Data field to given value.

### HasData

`func (o *IpamPool6Data) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


