# DhcpStaticData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DhcpStaticDataData**](DhcpStaticDataData.md) |  | [optional] 

## Methods

### NewDhcpStaticData

`func NewDhcpStaticData() *DhcpStaticData`

NewDhcpStaticData instantiates a new DhcpStaticData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStaticDataWithDefaults

`func NewDhcpStaticDataWithDefaults() *DhcpStaticData`

NewDhcpStaticDataWithDefaults instantiates a new DhcpStaticData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpStaticData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpStaticData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpStaticData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpStaticData) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpStaticData) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpStaticData) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpStaticData) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpStaticData) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpStaticData) GetData() []DhcpStaticDataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpStaticData) GetDataOk() (*[]DhcpStaticDataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpStaticData) SetData(v []DhcpStaticDataData)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpStaticData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


