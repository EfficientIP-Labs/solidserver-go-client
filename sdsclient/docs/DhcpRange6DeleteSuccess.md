# DhcpRange6DeleteSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerDhcpRange6DeleteSuccess**](DataInnerDhcpRange6DeleteSuccess.md) |  | [optional] 

## Methods

### NewDhcpRange6DeleteSuccess

`func NewDhcpRange6DeleteSuccess() *DhcpRange6DeleteSuccess`

NewDhcpRange6DeleteSuccess instantiates a new DhcpRange6DeleteSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpRange6DeleteSuccessWithDefaults

`func NewDhcpRange6DeleteSuccessWithDefaults() *DhcpRange6DeleteSuccess`

NewDhcpRange6DeleteSuccessWithDefaults instantiates a new DhcpRange6DeleteSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DhcpRange6DeleteSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DhcpRange6DeleteSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DhcpRange6DeleteSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DhcpRange6DeleteSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DhcpRange6DeleteSuccess) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DhcpRange6DeleteSuccess) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DhcpRange6DeleteSuccess) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DhcpRange6DeleteSuccess) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DhcpRange6DeleteSuccess) GetData() []DataInnerDhcpRange6DeleteSuccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DhcpRange6DeleteSuccess) GetDataOk() (*[]DataInnerDhcpRange6DeleteSuccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DhcpRange6DeleteSuccess) SetData(v []DataInnerDhcpRange6DeleteSuccess)`

SetData sets Data field to given value.

### HasData

`func (o *DhcpRange6DeleteSuccess) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


