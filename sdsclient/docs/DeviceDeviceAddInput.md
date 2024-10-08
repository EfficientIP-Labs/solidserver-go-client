# DeviceDeviceAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | The name of the Device Manager device, each device must have a unique name. | [optional] 
**DeviceAddr** | Pointer to **string** | The IP address you want to associate with the Device Manager device, in decimal format. | [optional] 
**DeviceAddressAddr** | Pointer to **string** | The IP address you want to associate with the Device Manager device, in hexadecimal format. | [optional] 
**DeviceSpaceId** | Pointer to **int32** | The database identifier (ID) of a space you want to associate with the Device Manager device. | [optional] 
**DevId** | Pointer to **int32** | The database identifier (ID) of a NetChange network device you want to associate with the Device Manager device. | [optional] 
**RowState** | Pointer to **string** | The object activation status.&lt;ul&gt;&lt;li&gt; If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**DeviceClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**DeviceClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDeviceDeviceAddInput

`func NewDeviceDeviceAddInput() *DeviceDeviceAddInput`

NewDeviceDeviceAddInput instantiates a new DeviceDeviceAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDeviceAddInputWithDefaults

`func NewDeviceDeviceAddInputWithDefaults() *DeviceDeviceAddInput`

NewDeviceDeviceAddInputWithDefaults instantiates a new DeviceDeviceAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *DeviceDeviceAddInput) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceDeviceAddInput) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceDeviceAddInput) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceDeviceAddInput) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceAddr

`func (o *DeviceDeviceAddInput) GetDeviceAddr() string`

GetDeviceAddr returns the DeviceAddr field if non-nil, zero value otherwise.

### GetDeviceAddrOk

`func (o *DeviceDeviceAddInput) GetDeviceAddrOk() (*string, bool)`

GetDeviceAddrOk returns a tuple with the DeviceAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddr

`func (o *DeviceDeviceAddInput) SetDeviceAddr(v string)`

SetDeviceAddr sets DeviceAddr field to given value.

### HasDeviceAddr

`func (o *DeviceDeviceAddInput) HasDeviceAddr() bool`

HasDeviceAddr returns a boolean if a field has been set.

### GetDeviceAddressAddr

`func (o *DeviceDeviceAddInput) GetDeviceAddressAddr() string`

GetDeviceAddressAddr returns the DeviceAddressAddr field if non-nil, zero value otherwise.

### GetDeviceAddressAddrOk

`func (o *DeviceDeviceAddInput) GetDeviceAddressAddrOk() (*string, bool)`

GetDeviceAddressAddrOk returns a tuple with the DeviceAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddressAddr

`func (o *DeviceDeviceAddInput) SetDeviceAddressAddr(v string)`

SetDeviceAddressAddr sets DeviceAddressAddr field to given value.

### HasDeviceAddressAddr

`func (o *DeviceDeviceAddInput) HasDeviceAddressAddr() bool`

HasDeviceAddressAddr returns a boolean if a field has been set.

### GetDeviceSpaceId

`func (o *DeviceDeviceAddInput) GetDeviceSpaceId() int32`

GetDeviceSpaceId returns the DeviceSpaceId field if non-nil, zero value otherwise.

### GetDeviceSpaceIdOk

`func (o *DeviceDeviceAddInput) GetDeviceSpaceIdOk() (*int32, bool)`

GetDeviceSpaceIdOk returns a tuple with the DeviceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpaceId

`func (o *DeviceDeviceAddInput) SetDeviceSpaceId(v int32)`

SetDeviceSpaceId sets DeviceSpaceId field to given value.

### HasDeviceSpaceId

`func (o *DeviceDeviceAddInput) HasDeviceSpaceId() bool`

HasDeviceSpaceId returns a boolean if a field has been set.

### GetDevId

`func (o *DeviceDeviceAddInput) GetDevId() int32`

GetDevId returns the DevId field if non-nil, zero value otherwise.

### GetDevIdOk

`func (o *DeviceDeviceAddInput) GetDevIdOk() (*int32, bool)`

GetDevIdOk returns a tuple with the DevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevId

`func (o *DeviceDeviceAddInput) SetDevId(v int32)`

SetDevId sets DevId field to given value.

### HasDevId

`func (o *DeviceDeviceAddInput) HasDevId() bool`

HasDevId returns a boolean if a field has been set.

### GetRowState

`func (o *DeviceDeviceAddInput) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DeviceDeviceAddInput) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DeviceDeviceAddInput) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DeviceDeviceAddInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DeviceDeviceAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DeviceDeviceAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DeviceDeviceAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DeviceDeviceAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetDeviceClassName

`func (o *DeviceDeviceAddInput) GetDeviceClassName() string`

GetDeviceClassName returns the DeviceClassName field if non-nil, zero value otherwise.

### GetDeviceClassNameOk

`func (o *DeviceDeviceAddInput) GetDeviceClassNameOk() (*string, bool)`

GetDeviceClassNameOk returns a tuple with the DeviceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassName

`func (o *DeviceDeviceAddInput) SetDeviceClassName(v string)`

SetDeviceClassName sets DeviceClassName field to given value.

### HasDeviceClassName

`func (o *DeviceDeviceAddInput) HasDeviceClassName() bool`

HasDeviceClassName returns a boolean if a field has been set.

### GetDeviceClassParameters

`func (o *DeviceDeviceAddInput) GetDeviceClassParameters() []ApiClassParameterInputEntry`

GetDeviceClassParameters returns the DeviceClassParameters field if non-nil, zero value otherwise.

### GetDeviceClassParametersOk

`func (o *DeviceDeviceAddInput) GetDeviceClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetDeviceClassParametersOk returns a tuple with the DeviceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassParameters

`func (o *DeviceDeviceAddInput) SetDeviceClassParameters(v []ApiClassParameterInputEntry)`

SetDeviceClassParameters sets DeviceClassParameters field to given value.

### HasDeviceClassParameters

`func (o *DeviceDeviceAddInput) HasDeviceClassParameters() bool`

HasDeviceClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DeviceDeviceAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DeviceDeviceAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DeviceDeviceAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DeviceDeviceAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


