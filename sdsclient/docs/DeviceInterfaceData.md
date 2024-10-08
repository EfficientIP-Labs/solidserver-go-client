# DeviceInterfaceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | state true/false indicate if action succeed | [optional] 
**Messages** | Pointer to [**[]ApiMessageEntry**](ApiMessageEntry.md) | List or notice/warning/error messages | [optional] 
**Data** | Pointer to [**[]DataInnerDeviceInterfaceData**](DataInnerDeviceInterfaceData.md) |  | [optional] 

## Methods

### NewDeviceInterfaceData

`func NewDeviceInterfaceData() *DeviceInterfaceData`

NewDeviceInterfaceData instantiates a new DeviceInterfaceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInterfaceDataWithDefaults

`func NewDeviceInterfaceDataWithDefaults() *DeviceInterfaceData`

NewDeviceInterfaceDataWithDefaults instantiates a new DeviceInterfaceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeviceInterfaceData) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeviceInterfaceData) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeviceInterfaceData) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeviceInterfaceData) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessages

`func (o *DeviceInterfaceData) GetMessages() []ApiMessageEntry`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *DeviceInterfaceData) GetMessagesOk() (*[]ApiMessageEntry, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *DeviceInterfaceData) SetMessages(v []ApiMessageEntry)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *DeviceInterfaceData) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetData

`func (o *DeviceInterfaceData) GetData() []DataInnerDeviceInterfaceData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceInterfaceData) GetDataOk() (*[]DataInnerDeviceInterfaceData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceInterfaceData) SetData(v []DataInnerDeviceInterfaceData)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceInterfaceData) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


