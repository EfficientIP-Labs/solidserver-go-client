# DeviceInterfaceEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device. | [optional] 
**InterfaceAddr** | Pointer to **string** | The IP addresses you want to associate with the Device Manager port or interface, as follows &lt;b&gt;&amp;lt;ip4_list&amp;gt;,&amp;lt;ip6_list&amp;gt;&lt;/b&gt;. | [optional] 
**InterfaceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify which port or interface to edit. | [optional] 
**InterfaceMac** | Pointer to **string** | The MAC address you want to associate with the Device Manager port or interface. | [optional] 
**InterfaceName** | Pointer to **string** | The name of the Device Manager port or interface, each port or interface must have a unique name. | [optional] 
**InterfaceType** | Pointer to **string** | A way to indicate if the object is either a &lt;b&gt;port&lt;/b&gt; or an &lt;b&gt;interface&lt;/b&gt;. | [optional] 
**InterfaceIp4List** | Pointer to **string** | The list of IPv4 addresses you want to associate to the Device Manager port or interface, in hexadecimal format, as follows: &lt;b&gt;&amp;lt;ip_address&amp;gt;, &amp;lt;ip_address&amp;gt;,...&lt;/b&gt; | [optional] 
**InterfaceAddress6List** | Pointer to **string** | todo(here) :device.interface.edit.input.interface_address6_list. : List of strings | [optional] 
**PortId** | Pointer to **int32** | The database identifier (ID) of a NetChange port you want to associate with the Device Manager port or interface. | [optional] 
**RowState** | Pointer to **int32** | The object activation status.&lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of a space you want to associate with the Device Manager port or interface. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**InterfaceClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**InterfaceClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDeviceInterfaceEditInput

`func NewDeviceInterfaceEditInput() *DeviceInterfaceEditInput`

NewDeviceInterfaceEditInput instantiates a new DeviceInterfaceEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInterfaceEditInputWithDefaults

`func NewDeviceInterfaceEditInputWithDefaults() *DeviceInterfaceEditInput`

NewDeviceInterfaceEditInputWithDefaults instantiates a new DeviceInterfaceEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceInterfaceEditInput) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInterfaceEditInput) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInterfaceEditInput) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceInterfaceEditInput) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceInterfaceEditInput) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceInterfaceEditInput) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceInterfaceEditInput) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceInterfaceEditInput) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInterfaceAddr

`func (o *DeviceInterfaceEditInput) GetInterfaceAddr() string`

GetInterfaceAddr returns the InterfaceAddr field if non-nil, zero value otherwise.

### GetInterfaceAddrOk

`func (o *DeviceInterfaceEditInput) GetInterfaceAddrOk() (*string, bool)`

GetInterfaceAddrOk returns a tuple with the InterfaceAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAddr

`func (o *DeviceInterfaceEditInput) SetInterfaceAddr(v string)`

SetInterfaceAddr sets InterfaceAddr field to given value.

### HasInterfaceAddr

`func (o *DeviceInterfaceEditInput) HasInterfaceAddr() bool`

HasInterfaceAddr returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DeviceInterfaceEditInput) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DeviceInterfaceEditInput) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DeviceInterfaceEditInput) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DeviceInterfaceEditInput) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceMac

`func (o *DeviceInterfaceEditInput) GetInterfaceMac() string`

GetInterfaceMac returns the InterfaceMac field if non-nil, zero value otherwise.

### GetInterfaceMacOk

`func (o *DeviceInterfaceEditInput) GetInterfaceMacOk() (*string, bool)`

GetInterfaceMacOk returns a tuple with the InterfaceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceMac

`func (o *DeviceInterfaceEditInput) SetInterfaceMac(v string)`

SetInterfaceMac sets InterfaceMac field to given value.

### HasInterfaceMac

`func (o *DeviceInterfaceEditInput) HasInterfaceMac() bool`

HasInterfaceMac returns a boolean if a field has been set.

### GetInterfaceName

`func (o *DeviceInterfaceEditInput) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *DeviceInterfaceEditInput) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *DeviceInterfaceEditInput) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *DeviceInterfaceEditInput) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DeviceInterfaceEditInput) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DeviceInterfaceEditInput) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DeviceInterfaceEditInput) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DeviceInterfaceEditInput) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetInterfaceIp4List

`func (o *DeviceInterfaceEditInput) GetInterfaceIp4List() string`

GetInterfaceIp4List returns the InterfaceIp4List field if non-nil, zero value otherwise.

### GetInterfaceIp4ListOk

`func (o *DeviceInterfaceEditInput) GetInterfaceIp4ListOk() (*string, bool)`

GetInterfaceIp4ListOk returns a tuple with the InterfaceIp4List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp4List

`func (o *DeviceInterfaceEditInput) SetInterfaceIp4List(v string)`

SetInterfaceIp4List sets InterfaceIp4List field to given value.

### HasInterfaceIp4List

`func (o *DeviceInterfaceEditInput) HasInterfaceIp4List() bool`

HasInterfaceIp4List returns a boolean if a field has been set.

### GetInterfaceAddress6List

`func (o *DeviceInterfaceEditInput) GetInterfaceAddress6List() string`

GetInterfaceAddress6List returns the InterfaceAddress6List field if non-nil, zero value otherwise.

### GetInterfaceAddress6ListOk

`func (o *DeviceInterfaceEditInput) GetInterfaceAddress6ListOk() (*string, bool)`

GetInterfaceAddress6ListOk returns a tuple with the InterfaceAddress6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAddress6List

`func (o *DeviceInterfaceEditInput) SetInterfaceAddress6List(v string)`

SetInterfaceAddress6List sets InterfaceAddress6List field to given value.

### HasInterfaceAddress6List

`func (o *DeviceInterfaceEditInput) HasInterfaceAddress6List() bool`

HasInterfaceAddress6List returns a boolean if a field has been set.

### GetPortId

`func (o *DeviceInterfaceEditInput) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DeviceInterfaceEditInput) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DeviceInterfaceEditInput) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *DeviceInterfaceEditInput) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRowState

`func (o *DeviceInterfaceEditInput) GetRowState() int32`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DeviceInterfaceEditInput) GetRowStateOk() (*int32, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DeviceInterfaceEditInput) SetRowState(v int32)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DeviceInterfaceEditInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetSpaceId

`func (o *DeviceInterfaceEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DeviceInterfaceEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DeviceInterfaceEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DeviceInterfaceEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DeviceInterfaceEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DeviceInterfaceEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DeviceInterfaceEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DeviceInterfaceEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetInterfaceClassName

`func (o *DeviceInterfaceEditInput) GetInterfaceClassName() string`

GetInterfaceClassName returns the InterfaceClassName field if non-nil, zero value otherwise.

### GetInterfaceClassNameOk

`func (o *DeviceInterfaceEditInput) GetInterfaceClassNameOk() (*string, bool)`

GetInterfaceClassNameOk returns a tuple with the InterfaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceClassName

`func (o *DeviceInterfaceEditInput) SetInterfaceClassName(v string)`

SetInterfaceClassName sets InterfaceClassName field to given value.

### HasInterfaceClassName

`func (o *DeviceInterfaceEditInput) HasInterfaceClassName() bool`

HasInterfaceClassName returns a boolean if a field has been set.

### GetInterfaceClassParameters

`func (o *DeviceInterfaceEditInput) GetInterfaceClassParameters() []ApiClassParameterInputEntry`

GetInterfaceClassParameters returns the InterfaceClassParameters field if non-nil, zero value otherwise.

### GetInterfaceClassParametersOk

`func (o *DeviceInterfaceEditInput) GetInterfaceClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetInterfaceClassParametersOk returns a tuple with the InterfaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceClassParameters

`func (o *DeviceInterfaceEditInput) SetInterfaceClassParameters(v []ApiClassParameterInputEntry)`

SetInterfaceClassParameters sets InterfaceClassParameters field to given value.

### HasInterfaceClassParameters

`func (o *DeviceInterfaceEditInput) HasInterfaceClassParameters() bool`

HasInterfaceClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DeviceInterfaceEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DeviceInterfaceEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DeviceInterfaceEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DeviceInterfaceEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


